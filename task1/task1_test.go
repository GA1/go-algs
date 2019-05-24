package task1

import (
	"testing"
)

func TestSolveTask1(t *testing.T) {
	tests := []struct {
		input []int
		expected int
	}{
		{
			input: []int{1, 1},
			expected: 0,
		},
		{
			input: []int{7, 3},
			expected: 0,
		},
		{
			input: []int{7, 2, 13},
			expected: 11,
		},
		{
			input: []int{3, 6, 2},
			expected: 3,
		},
	}

	for _, test := range tests {
		result := Solution(test.input)
		if result != test.expected {
			t.Errorf("Solution is incorrect, got: %d, expected: %d", result, test.expected)
		}
	}
}
