package word_break

import (
	"fmt"
	"testing"

	"github.com/ashtanko/go-algorithms/utils"
	is "gotest.tools/v3/assert/cmp"
)

type testCase struct {
	s        string
	wordDict []string
	expected bool
}

func TestWordBreak(t *testing.T) {

	inputData := []testCase{
		{
			"",
			[]string{"leet"},
			true,
		},
		{
			"",
			[]string{},
			true,
		},
		{
			"a",
			[]string{},
			false,
		},
		{
			"a",
			[]string{"a"},
			true,
		},
		{
			"leetcode",
			[]string{"leet", "code"},
			true,
		},
		{
			"applepenapple",
			[]string{"apple", "pen"},
			true,
		},
		{
			"catsandog",
			[]string{"cats", "dog", "sand", "and", "cat"},
			false,
		},
	}

	testCases := []struct {
		strategy     func(string, []string) bool
		strategyName string
	}{
		{
			strategy:     wordBreakTopDown,
			strategyName: "top down",
		},
		{
			strategy:     wordBreakMemo,
			strategyName: "top down memo",
		},
		{
			strategy:     wordBreak1D,
			strategyName: "DP 1D",
		},
	}

	for _, tc := range testCases {
		for _, i := range inputData {
			name := fmt.Sprintf("%s %s", tc.strategyName, i.s)
			t.Run(name, func(t *testing.T) {
				actual := tc.strategy(i.s, i.wordDict)
				utils.Checkf(t, is.DeepEqual(i.expected, actual), tc)
			})
		}
	}
}
