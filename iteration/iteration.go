package iteration

import (
	"strings"
)

func Repeat(character string, numberOfTimes int) string {
	var repeated string
	for i := 0; i < numberOfTimes; i++ {
		repeated += character
	}
	return repeated
}

func SubstringInString(word, substring string) bool {
	return strings.Contains(word, substring)
}
