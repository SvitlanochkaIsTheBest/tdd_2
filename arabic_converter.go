package main

import (
	"errors"
)

var romanValues = map[rune]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func toArabic(roman string) (int, error) {
	arabic := 0

	for i := 0; i < len(roman); i++ {
		value, exists := romanValues[int32(roman[i])]
		if !exists {
			return 0, errors.New("invalid Roman numeral")
		}

		if i+1 < len(roman) {
			nextValue, exists := romanValues[int32(roman[i+1])]
			if !exists {
				return 0, errors.New("invalid Roman numeral")
			}
			if nextValue > value {
				arabic -= value
				continue
			}
		}
		arabic += value
	}

	if arabic < 1 || arabic > 3999 {
		return -1, errors.New("number out of range (1-3999)")
	}
	return arabic, nil
}
