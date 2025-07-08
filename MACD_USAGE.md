# MACD Indicator Usage Example

This document demonstrates how to use the MACD (Moving Average Convergence Divergence) indicator in your trading strategies.

## Configuration Example

Here's how to configure MACD indicators in your technology configuration:

```json
{
  "macd": [
    {
      "name": "macd_4h_standard",
      "enable": true,
      "kline_interval": "4h",
      "fast_period": 12,
      "slow_period": 26,
      "signal_period": 9
    },
    {
      "name": "macd_1h_fast",
      "enable": true,
      "kline_interval": "1h", 
      "fast_period": 8,
      "slow_period": 21,
      "signal_period": 5
    }
  ]
}
```

## Strategy Examples

### 1. MACD Golden Cross Strategy (Long)

```javascript
// Long signal: MACD line crosses above signal line
macd_4h_standard.MACD[0] > macd_4h_standard.Signal[0] && 
macd_4h_standard.MACD[1] <= macd_4h_standard.Signal[1]
```

### 2. MACD Death Cross Strategy (Short)

```javascript
// Short signal: MACD line crosses below signal line  
macd_4h_standard.MACD[0] < macd_4h_standard.Signal[0] && 
macd_4h_standard.MACD[1] >= macd_4h_standard.Signal[1]
```

### 3. MACD Histogram Momentum Strategy

```javascript
// Long signal: Histogram turns positive (MACD above signal)
macd_4h_standard.Histogram[0] > 0 && macd_4h_standard.Histogram[1] <= 0
```

### 4. MACD Zero Line Cross Strategy

```javascript
// Long signal: MACD line crosses above zero line
macd_4h_standard.MACD[0] > 0 && macd_4h_standard.MACD[1] <= 0
```

### 5. Multi-Timeframe MACD Strategy

```javascript
// Long signal: Both 4h and 1h MACD show bullish signals
macd_4h_standard.MACD[0] > macd_4h_standard.Signal[0] &&
macd_1h_fast.MACD[0] > macd_1h_fast.Signal[0] &&
macd_4h_standard.Histogram[0] > macd_4h_standard.Histogram[1]
```

## Available MACD Variables

For each MACD indicator (e.g., `macd_4h_standard`), you have access to:

- `macd_4h_standard.KlineInterval` - K-line interval (e.g., "4h")
- `macd_4h_standard.FastPeriod` - Fast EMA period (e.g., 12)
- `macd_4h_standard.SlowPeriod` - Slow EMA period (e.g., 26)
- `macd_4h_standard.SignalPeriod` - Signal line EMA period (e.g., 9)
- `macd_4h_standard.MACD` - Array of MACD line values (newest first)
- `macd_4h_standard.Signal` - Array of signal line values (newest first)
- `macd_4h_standard.Histogram` - Array of histogram values (newest first)

## MACD Signal Interpretation

- **MACD > Signal**: Bullish momentum
- **MACD < Signal**: Bearish momentum
- **MACD > 0**: Price above long-term average
- **MACD < 0**: Price below long-term average
- **Histogram > 0**: MACD above signal line
- **Histogram < 0**: MACD below signal line
- **Increasing Histogram**: Strengthening momentum
- **Decreasing Histogram**: Weakening momentum

## Complete Strategy Configuration Example

```json
{
  "ma": [],
  "ema": [],
  "rsi": [],
  "kc": [],
  "boll": [],
  "atr": [],
  "macd": [
    {
      "name": "macd_main",
      "enable": true,
      "kline_interval": "4h",
      "fast_period": 12,
      "slow_period": 26,
      "signal_period": 9
    }
  ]
}
```

With corresponding strategy rules:

```json
[
  {
    "name": "MACD Long Entry",
    "enable": true,
    "code": "macd_main.MACD[0] > macd_main.Signal[0] && macd_main.MACD[1] <= macd_main.Signal[1] && macd_main.Histogram[0] > 0",
    "type": "long"
  },
  {
    "name": "MACD Short Entry", 
    "enable": true,
    "code": "macd_main.MACD[0] < macd_main.Signal[0] && macd_main.MACD[1] >= macd_main.Signal[1] && macd_main.Histogram[0] < 0",
    "type": "short"
  },
  {
    "name": "MACD Long Exit",
    "enable": true,
    "code": "macd_main.MACD[0] < macd_main.Signal[0] && macd_main.MACD[1] >= macd_main.Signal[1]",
    "type": "close_long"
  },
  {
    "name": "MACD Short Exit",
    "enable": true,
    "code": "macd_main.MACD[0] > macd_main.Signal[0] && macd_main.MACD[1] <= macd_main.Signal[1]", 
    "type": "close_short"
  }
]
```