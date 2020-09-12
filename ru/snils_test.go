package ru

import (
	"fmt"
	"testing"
)

func TestCalculateChecksum(t *testing.T) {
	cases := []struct {
		number, checksum int
	}{
		{112233445, 95},
		{87654303, 0},
	}
	for i, c := range cases {
		t.Run(fmt.Sprintf("test case %d", i), func(t *testing.T) {
			if calculateSNILSChecksum(c.number) != c.checksum {
				t.Fatal("invalid checksum")
			}
		})

	}
}

func TestRandomSNILS(t *testing.T) {
	for i := 0; i < 1000; i++ {
		s := RandomSNILS()
		_, err := ParseSNILS(s.Full())
		if err != nil {
			t.Fatal(err)
		}
	}

}
