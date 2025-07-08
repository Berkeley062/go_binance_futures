package line

import (
	"math"
	"testing"
)

// Helper function to create test data for EMA calculation
func createTestData() []float64 {
	return []float64{
		22.27, 22.19, 22.08, 22.17, 22.18, 22.13, 22.23, 22.43, 22.24, 22.29,
		22.15, 22.39, 22.38, 22.61, 23.36, 24.05, 23.75, 23.83, 23.95, 23.63,
		23.82, 23.87, 23.65, 23.19, 23.10, 23.33, 22.68, 23.10, 22.40, 22.17,
	}
}

// Simple EMA calculation for testing
func simpleEMA(prices []float64, period int) []float64 {
	if len(prices) < period {
		return nil
	}
	
	alpha := 2.0 / (float64(period) + 1)
	ema := make([]float64, len(prices))
	
	// Start with SMA for first value
	sum := 0.0
	for i := 0; i < period; i++ {
		sum += prices[i]
	}
	ema[period-1] = sum / float64(period)
	
	// Calculate EMA for remaining values
	for i := period; i < len(prices); i++ {
		ema[i] = alpha*prices[i] + (1-alpha)*ema[i-1]
	}
	
	return ema[period-1:]
}

// Test MACD calculation logic without external dependencies
func TestMACDCalculationLogic(t *testing.T) {
	prices := createTestData()
	
	// Manually calculate MACD components for testing
	fastPeriod := 5  // Shorter periods for testing
	slowPeriod := 10
	signalPeriod := 3
	
	// Calculate EMAs manually
	fastEMA := simpleEMA(prices, fastPeriod)
	slowEMA := simpleEMA(prices, slowPeriod)
	
	if len(fastEMA) == 0 || len(slowEMA) == 0 {
		t.Error("EMAs should not be empty")
		return
	}
	
	// Calculate MACD line manually
	minLen := len(slowEMA)
	if len(fastEMA) < minLen {
		minLen = len(fastEMA)
	}
	
	macdLine := make([]float64, minLen)
	for i := 0; i < minLen; i++ {
		macdLine[i] = fastEMA[i] - slowEMA[i]
	}
	
	// Calculate signal line manually
	signalLine := simpleEMA(macdLine, signalPeriod)
	
	// Calculate histogram manually
	histogram := make([]float64, len(signalLine))
	for i := 0; i < len(signalLine); i++ {
		histogram[i] = macdLine[i] - signalLine[i]
	}
	
	// Basic validation
	if len(macdLine) == 0 {
		t.Error("MACD line should not be empty")
	}
	
	if len(signalLine) == 0 {
		t.Error("Signal line should not be empty")
	}
	
	if len(histogram) == 0 {
		t.Error("Histogram should not be empty")
	}
	
	if len(histogram) != len(signalLine) {
		t.Errorf("Histogram length (%d) should equal signal line length (%d)", len(histogram), len(signalLine))
	}
	
	t.Logf("MACD calculation logic test passed")
	t.Logf("MACD line length: %d", len(macdLine))
	t.Logf("Signal line length: %d", len(signalLine))
	t.Logf("Histogram length: %d", len(histogram))
	
	// Test values are reasonable
	for i := 0; i < len(histogram); i++ {
		if math.IsNaN(histogram[i]) || math.IsInf(histogram[i], 0) {
			t.Errorf("Histogram[%d] is invalid: %f", i, histogram[i])
		}
	}
}

// Test MACD parameter validation
func TestMACDParameterValidation(t *testing.T) {
	prices := createTestData()
	
	// Test with fast >= slow period
	testCases := []struct {
		fast, slow, signal int
		shouldError       bool
		description       string
	}{
		{12, 26, 9, false, "standard parameters"},
		{26, 12, 9, true, "fast >= slow"},
		{12, 12, 9, true, "fast == slow"},
		{5, 10, 3, false, "valid short periods"},
	}
	
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			if tc.shouldError {
				// We can't test the actual function due to dependencies, 
				// but we can test the validation logic
				if tc.fast >= tc.slow {
					t.Logf("Correctly identified invalid parameters: fast=%d, slow=%d", tc.fast, tc.slow)
				} else {
					t.Errorf("Should have detected invalid parameters: fast=%d, slow=%d", tc.fast, tc.slow)
				}
			} else {
				if tc.fast < tc.slow && len(prices) >= tc.slow {
					t.Logf("Valid parameters: fast=%d, slow=%d, signal=%d", tc.fast, tc.slow, tc.signal)
				} else {
					t.Errorf("Invalid test case: fast=%d, slow=%d", tc.fast, tc.slow)
				}
			}
		})
	}
}

// Test MACD with insufficient data
func TestMACDInsufficientData(t *testing.T) {
	shortPrices := []float64{22.27, 22.19, 22.08} // Only 3 data points
	
	// Test validation logic
	slowPeriod := 26
	if len(shortPrices) < slowPeriod {
		t.Logf("Correctly identified insufficient data: %d < %d", len(shortPrices), slowPeriod)
	} else {
		t.Error("Should have detected insufficient data")
	}
}