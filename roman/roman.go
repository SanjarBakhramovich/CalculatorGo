package roman

import "errors"

func RomanToArabic(roman string) (uint8, error) {
	romanNums := map[string]uint8{
		"I": 1, "II": 2, "III": 3, "IV": 4,
		"V": 5, "VI": 6, "VII": 7, "VIII": 8,
		"IX": 9,
	}

	if val, ok := romanNums[roman]; ok {
		return val, nil
	}

	return 0, errors.New("неверная римская цифра")
}