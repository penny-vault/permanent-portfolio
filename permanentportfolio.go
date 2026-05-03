// Copyright 2026
// SPDX-License-Identifier: Apache-2.0
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"
	_ "embed"
	"fmt"
	"strings"
	"time"

	"github.com/penny-vault/pvbt/asset"
	"github.com/penny-vault/pvbt/engine"
	"github.com/penny-vault/pvbt/portfolio"
	"github.com/penny-vault/pvbt/universe"
)

//go:embed README.md
var description string

// expectedSleeves is the canonical four-sleeve construction described by Harry
// Browne: stocks, long-term Treasuries, gold, and short-term Treasuries (cash).
const expectedSleeves = 4

// PermanentPortfolio implements Harry Browne's four-asset, equally-weighted
// portfolio rebalanced once per year. The four sleeves are designed to perform
// in each of the four economic regimes Browne identified: prosperity (stocks),
// deflation (long-term Treasuries), inflation (gold), and recession (cash).
type PermanentPortfolio struct {
	Tickers universe.Universe `pvbt:"tickers" desc:"Comma-separated tickers for the four sleeves: stocks, long-term Treasuries, gold, cash. Order is irrelevant; all four are equally weighted." default:"SPY,TLT,GLD,SHV"`
}

func (s *PermanentPortfolio) Name() string { return "Permanent Portfolio (Harry Browne)" }

func (s *PermanentPortfolio) Setup(_ *engine.Engine) {}

func (s *PermanentPortfolio) Describe() engine.StrategyDescription {
	return engine.StrategyDescription{
		ShortCode:   "pp",
		Description: description,
		Source:      "https://en.wikipedia.org/wiki/Fail-Safe_Investing",
		Version:     "0.1.0",
		VersionDate: time.Date(2026, 5, 2, 0, 0, 0, 0, time.UTC),
		Schedule:    "@monthend",
		Benchmark:   "SPY",
	}
}

func (s *PermanentPortfolio) Compute(ctx context.Context, eng *engine.Engine, _ portfolio.Portfolio, batch *portfolio.Batch) error {
	// Annual rebalance: tradecron has no @yearend directive, so use @monthend
	// and filter to December here.
	if eng.CurrentDate().Month() != time.December {
		return nil
	}

	assets := s.Tickers.Assets(eng.CurrentDate())
	if len(assets) != expectedSleeves {
		return fmt.Errorf("permanent portfolio requires exactly %d tickers, got %d", expectedSleeves, len(assets))
	}

	weight := 1.0 / float64(expectedSleeves)
	members := make(map[asset.Asset]float64, expectedSleeves)
	tickers := make([]string, 0, expectedSleeves)
	for _, a := range assets {
		members[a] = weight
		tickers = append(tickers, a.Ticker)
	}

	justification := "permanent portfolio: 25% each of " + strings.Join(tickers, ", ")
	batch.Annotate("justification", justification)

	allocation := portfolio.Allocation{
		Date:          eng.CurrentDate(),
		Members:       members,
		Justification: justification,
	}

	if err := batch.RebalanceTo(ctx, allocation); err != nil {
		return fmt.Errorf("rebalance failed: %w", err)
	}

	return nil
}
