package ru

import (
	"fmt"
	"github.com/mr-tron/papers"
	"math/rand"
	"strconv"
	"strings"
	"unicode"
)

type SNILS struct {
	value string
}

// NewSNILS returns SNILS without format validation.
// Input value should have format `12345678900` (just 11 digits)
// It's useful if you ensure that your input has correct format and value (checksum).
func NewSNILS(s string) SNILS {
	return SNILS{s}
}

func ParseSNILS(s string) (SNILS, error) {
	s = extractDigits(s)
	if len(s) == 9 {
		return SNILS{value: s}, nil
	}
	if len(s) != 11 {
		return SNILS{}, fmt.Errorf("%w: too many digits in SNILS '%v'", papers.ParsingError, s)
	}

	intSnils, _ := strconv.Atoi(s[:9])
	cheksum, _ := strconv.Atoi(s[9:])
	if intSnils < 1001998 {
		return SNILS{}, fmt.Errorf("%w: SNILS with number less than 001-001-998 shoudn't has checksum, but '%v' has checksum", papers.ParsingError, s)
	}
	calculatedCheksum := calculateSNILSChecksum(intSnils)
	if calculatedCheksum != cheksum {
		return SNILS{}, fmt.Errorf("%w: invalid checksum %d. should be %d", papers.ParsingError, cheksum, calculatedCheksum)
	}
	return SNILS{value: s}, nil
}

func RandomSNILS() SNILS {
	number := rand.Intn(98998000) + 1001998
	snils := int64(calculateSNILSChecksum(number)) + int64(number)*100
	return SNILS{value: fmt.Sprintf("%011d", snils)}
}

func (s SNILS) String() string {
	return s.Full()
}

func (s SNILS) Short() string {
	return s.value
}

func (s SNILS) Full() string {
	if len(s.value) < 11 {
		return s.value
	}
	return fmt.Sprintf("%s-%s-%s %s", s.value[:3], s.value[3:6], s.value[6:9], s.value[9:11])
}

func calculateSNILSChecksum(snils int) int {
	var sum int
	for i := 1; i < 10; i++ {
		sum += (snils % 10) * i
		snils /= 10
	}
	if sum < 100 {
		return sum
	}
	if sum > 101 {
		return sum % 101
	}
	return 0
}

func extractDigits(s string) string {
	b := strings.Builder{}
	for _, r := range s {
		if unicode.IsDigit(r) {
			b.WriteRune(r)
		}
	}
	return b.String()
}
