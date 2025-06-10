package detector

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDetectNetwork(t *testing.T) {
	tests := []struct {
		cardNumber string
		expected   string
	}{
		{"341234567890123", "American Express"},
		{"371234567890123", "American Express"},
		{"38123456789012", "Diners Club"},
		{"4123456789012", "Visa"},
		{"4123456789012345", "Visa"},
		{"4123456789012345678", "Visa"},
		{"5112345678901234", "MasterCard"},
		{"5212345678901234", "MasterCard"},
		{"5312345678901234", "MasterCard"},
		{"5412345678901234", "MasterCard"},
		{"5512345678901234", "MasterCard"},
		{"6011123456789012345", "Discover"},
		{"6221261234567890", "Discover"},
		{"6441234567890123", "Discover"},
		{"6512345678901234", "Discover"},
		{"5034567890123456", "Maestro"},
		{"582126123456", "Maestro"},
		{"5034567890123456789", "Maestro"},
	}

	for _, tc := range tests {
		t.Run(tc.cardNumber, func(t *testing.T) {
			got := DetectNetwork(tc.cardNumber)
			assert.Equal(t, tc.expected, got, "Expected network for card number %s to be %s, but got %s", tc.cardNumber, tc.expected, got)
		})
	}
}
