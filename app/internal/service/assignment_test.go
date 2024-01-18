package service

import (
	"errors"
	"fmt"
	"testing"
)

func TestGeneratePermutations(t *testing.T) {
	//setup env
	svc := NewService()

	t.Run("GeneratePermutations() should get response success", func(t *testing.T) {
		testCasesGeneratePermutation := []struct {
			input    string
			expected []string
		}{
			{"abc", []string{"abc", "acb", "bac", "bca", "cab", "cba"}},
			{"", []string{""}},
			{"a", []string{"a"}},
			{"123", []string{"123", "132", "213", "231", "312", "321"}},
		}

		for _, tc := range testCasesGeneratePermutation {
			response := svc.GeneratePermutations(tc.input)
			t.Logf(fmt.Sprintf("Your function should return:%s", response))
		}
	})

	t.Run("FindOdd() should get response success", func(t *testing.T) {
		testCases := []struct {
			input    []int
			expected int
			err      error
		}{
			{[]int{1, 2, 2, 3, 1, 3, 4, 5, 4}, 5, nil},                                                      // Valid case with odd occurrences of 5
			{[]int{1, 2, 2, 3, 1, 3, 4, 4}, 0, errors.New("Number not found in Odd : [1 2 2 3 1 3 4 4]\n")}, // No number with odd occurrences
			{[]int{}, 0, errors.New("Number not found in Odd : []\n")},                                      // Empty input
		}

		for _, tc := range testCases {
			result, _ := svc.FindOdd(tc.input)

			// Check if the result is equal to the expected value
			if result != tc.expected {
				t.Errorf("For input %v, expected %d, but got %d", tc.input, tc.expected, result)
			}
		}
	})

	t.Run("CheckCountSmileys() should get response success", func(t *testing.T) {
		testCases := []struct {
			input    []string
			expected int
		}{
			{[]string{":)", ":D", ";-D", ":~)"}, 4},      // Valid smiley faces
			{[]string{":(", ":>", ":}", ":]"}, 0},        // Invalid smiley faces
			{[]string{";)", ":]", ":)", ":D"}, 2},        // Mix of valid and invalid smiley faces
			{[]string{";~D", ":]", ":)", ":D"}, 1},       // Valid and invalid with noses
			{[]string{";)", ":]", ":)", ":D", ":)D"}, 3}, // Multiple occurrences of the same smiley
		}

		for _, tc := range testCases {
			response := svc.CheckCountSmileys(tc.input)
			t.Logf(fmt.Sprintf("Your function should return:%v", response))
		}
	})
}
