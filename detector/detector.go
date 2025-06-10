package detector

import (
	"strconv"
	"strings"
)

type NetworkRule struct {
	Name   string
	Prefix []string
	Length []string // Supports fixed and ranged (e.g., "12-19")
}

var networkRules = []NetworkRule{
	{Name: "American Express", Prefix: []string{"34", "37"}, Length: []string{"15"}},
	{Name: "Diners Club", Prefix: []string{"38", "39"}, Length: []string{"14"}},
	{Name: "Visa", Prefix: []string{"4"}, Length: []string{"13", "16", "19"}},
	{Name: "MasterCard", Prefix: []string{"51", "52", "53", "54", "55"}, Length: []string{"16"}},
	{Name: "Discover", Prefix: []string{"6011", "622126-622925", "644-649", "65"}, Length: []string{"16", "19"}},
	{Name: "Maestro", Prefix: []string{"50", "56-59"}, Length: []string{"12-19"}},
}

func DetectNetwork(cardNumber string) string {
	cardLen := len(cardNumber)

	for _, rule := range networkRules {
		if !matchesLength(cardLen, rule.Length) {
			continue
		}

		for _, prefix := range rule.Prefix {
			if isPrefixMatch(cardNumber, prefix) {
				return rule.Name
			}
		}
	}

	return "Unknown"
}

func isPrefixMatch(cardNumber string, prefixRule string) bool {
	if strings.Contains(prefixRule, "-") {
		bounds := strings.Split(prefixRule, "-")

		if len(bounds) != 2 {
			return false
		}

		low := bounds[0]
		high := bounds[1]
		prefixLen := len(low)

		if len(cardNumber) < prefixLen {
			return false
		}

		prefixVal, err := strconv.Atoi(cardNumber[:prefixLen])
		lowVal, err1 := strconv.Atoi(low)
		highVal, err2 := strconv.Atoi(high)
		if err != nil || err1 != nil || err2 != nil {
			return false
		}

		return prefixVal >= lowVal && prefixVal <= highVal
	}

	return strings.HasPrefix(cardNumber, prefixRule)
}

func matchesLength(length int, rules []string) bool {
	for _, rule := range rules {
		if strings.Contains(rule, "-") {
			bounds := strings.Split(rule, "-")

			if len(bounds) != 2 {
				continue
			}

			min, err1 := strconv.Atoi(bounds[0])
			max, err2 := strconv.Atoi(bounds[1])
			if err1 != nil || err2 != nil {
				continue
			}

			if length >= min && length <= max {
				return true
			}
		} else {
			exact, err := strconv.Atoi(rule)
			if err == nil && length == exact {
				return true
			}
		}
	}

	return false
}
