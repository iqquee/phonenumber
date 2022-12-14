package phonenumber

import (
	"errors"
)

func Parse(number string) (string, error) {

	constAlpha := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j",
		"k", "l", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
		"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O",
		"P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

	var newNum string
	var avail bool
	for _, v := range number {
		newChar := string(v)
		for _, x := range constAlpha {
			if newChar == x {
				avail = true
				break
			}
		}
	}

	if avail {
		return "", errors.New("invalid phone number")
	} else {
		//verify if the number is upto 11 digits
		if len(number) == 11 {

			num := trimLeftChar(number)
			constNum := "234"
			newNum = constNum + num

		} else {
			return "", errors.New("invalid phone number")
		}
	}

	return newNum, nil
}

func trimLeftChar(s string) string {
	for i := range s {
		if i > 0 {
			return s[i:]
		}
	}
	return s[:0]
}
