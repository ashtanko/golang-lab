package has_all_codes

import (
	"fmt"
	"github.com/ashtanko/go-algorithms/utils"
	is "gotest.tools/v3/assert/cmp"
	"testing"
)

func TestHasAllCodes(t *testing.T) {
	testCases := []struct {
		s        string
		k        int
		expected bool
	}{
		{
			"00110110",
			2,
			true,
		},
		{
			"00110",
			2,
			true,
		},
		{
			"0110",
			1,
			true,
		},
		{
			"0110",
			2,
			false,
		},
		{
			"0000000001011100",
			4,
			false,
		},
		{
			"11000110100010010000000110011000101111100110110110100111011111010101000110001010001111010001100000011010001100000000010011011011011111101101011111111000111000000000100011111010010010001110010111100011011000010111000000010111111001100110011011001000101010011101101010000010001010010110010001100101110111001101001011001011010011101001001101101000011111001110101011011101110011010110111001100010010011011001111000110110011101011101011011000011110100011100010000011011101000010101010010011010100011000011100111101001111100110110000000100111111101111010010100010101110110000101111001011010011010011001001010010001101010000011010001100011000010010101110111101101000111000011001101101010110101000111010011011001000100011111101111001001011010111010111001001010001110101000111100010010101100100111001001110001100101000100000000101101000101100010100010100100110000101000100010111001100101110101001101001001110010101111111010111010011001111011011100001101111010000000111110110100110000010101011100110101111000101110011110010110100011011001101111000101011101110100011000000110100001111111011001110110010110110101001001000001010001000010101011110000001000011011000111110001100110011010001000000010001100101010011011010011101101011000101111110100011011111110000100000111110011111001110110111011011110000101111101011100001011110010111011110101000111111100111000101100101000110111101000011110101010101000000101010101100001011011000000100100011110010001110100111111001100010001000001010110101011001000000010111000010010111011110110010000100011101110101101000101111110010011110010000110001001101000001100100010010110010111001101001111000000111010111110110110111111011000010001111111111001001110111100100011101100111101000010111100101111111100011101101100010001111010101100000101100111110100001011011111011110111101010100011001001111100001001101010000001101110000110011101110110111101111111101000011001010110101100000000000101000000000101111110010001110111111110010101101001001010001010000001010001000000000111011101100010011001001010110011100110000000011100011001100",
			8,
			true,
		},
	}

	for _, tc := range testCases {
		name := fmt.Sprintf("s: %s k: %d", tc.s, tc.k)
		t.Run(name, func(t *testing.T) {
			actual := hasAllCodes(tc.s, tc.k)
			utils.Checkf(t, is.Equal(tc.expected, actual), tc)
		})
	}
}