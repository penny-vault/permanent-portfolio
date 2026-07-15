# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [0.2.4] - 2026-07-14

### Changed
- Upgrade pvbt dependency to v0.12.1
- Update other dependencies to latest

## [0.2.3] - 2026-05-28

### Changed
- Upgrade pvbt dependency to v0.10.3
- Update other dependencies to latest

## [0.2.2] - 2026-05-26

### Changed
- Upgrade pvbt dependency to v0.10.2
- Update other dependencies to latest

## [0.2.1] - 2026-05-05

### Changed
- Upgrade pvbt dependency to v0.9.3

## [0.2.0] - 2026-05-04

### Changed
- Upgrade pvbt dependency to v0.9.2

## [0.1.0] - 2026-05-02

### Added
- Initial release of the Permanent Portfolio (Harry Browne) strategy
- Equally-weighted four-sleeve allocation (stocks, long-term Treasuries, gold,
  short-term Treasuries) rebalanced annually on the last trading day of December
- `tickers` parameter (default `SPY,TLT,GLD,SHV`) accepts a comma-separated list
  of four tickers; all are equally weighted at 25% each

[0.1.0]: https://github.com/penny-vault/permanent-portfolio/releases/tag/v0.1.0
[0.2.0]: https://github.com/penny-vault/permanent-portfolio/compare/v0.1.0...v0.2.0
[0.2.1]: https://github.com/penny-vault/permanent-portfolio/compare/v0.2.0...v0.2.1
[0.2.3]: https://github.com/penny-vault/permanent-portfolio/compare/v0.2.2...v0.2.3
[0.2.4]: https://github.com/penny-vault/permanent-portfolio/compare/v0.2.3...v0.2.4
