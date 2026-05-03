# Permanent Portfolio (Harry Browne)

The **Permanent Portfolio**, designed by free-market author and investment advisor Harry Browne in the 1980s, is a long-only, equally-weighted allocation across four asset sleeves -- each chosen to perform in one of the four economic regimes Browne identified. The portfolio is rebalanced once per year and is intended to be a "set and forget" baseline: low-maintenance, robust to forecast errors, and resilient across decades of changing macro conditions.

## The four sleeves

| Sleeve | Default ticker | Performs best during | Why |
|--------|---------------|---------------------|-----|
| Stocks | `SPY` | Prosperity | Equity multiples expand when growth is strong and capital is plentiful. |
| Long-term Treasuries | `TLT` | Deflation | Long-duration bonds rally when rates fall in response to slowing growth and falling prices. |
| Gold | `GLD` | Inflation | Gold preserves real purchasing power when fiat currency depreciates. |
| Short-term Treasuries (cash) | `SHV` | Recession | Cash holds nominal value and is the safe pool that buys the other three sleeves at the rebalance. |

## Rules

1. Hold each sleeve at 25% of portfolio value.
2. On the last trading day of each calendar year, rebalance back to the 25/25/25/25 target.
3. No tactical signals, no timing, no momentum overlay -- the four-sleeve diversification *is* the risk management.

The strategy trades at most once per year per sleeve.

## Parameters

- **Tickers** (default `SPY,TLT,GLD,SHV`): comma-separated list of exactly four tickers in the order stocks, long-bonds, gold, cash. Order does not affect behavior -- all four are equally weighted regardless. Substitution lets you express the same portfolio with different ETF wrappers (e.g. `VTI,EDV,GLD,VGSH` for low-expense alternatives).

## Rationale for the design

Browne's insight is that the macroeconomic future is unknowable, but every period falls into one of four regimes (prosperity, deflation, inflation, recession). Each sleeve is a near-perfect hedge for one regime. By weighting the four sleeves equally, the portfolio always has roughly a quarter of its capital allocated to whatever regime is currently rewarding the market -- without requiring the investor to forecast which regime that is. The annual rebalance is a buy-low / sell-high mechanic against the prior year's regime.

## Why annual rebalancing

More frequent rebalancing churns the portfolio without materially improving the long-run risk/return profile of a static-weight strategy, while annual rebalancing is tax-efficient (long-term capital gains apply to any sales) and keeps trading costs negligible. Browne's original recommendation was to rebalance only when an asset deviated more than 35/15 percentage points from target; the simpler annual cadence used here is a close approximation that avoids state-dependent triggers.

## Source

Harry Browne, *Fail-Safe Investing: Lifelong Financial Security in 30 Minutes* (1999). The strategy is widely documented; see also <https://en.wikipedia.org/wiki/Fail-Safe_Investing>.
