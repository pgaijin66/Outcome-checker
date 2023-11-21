package main

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		expectedOutput string
		expectedErr    error
	}{
		{
			name:           "Player X wins most of the time",
			input:          "XOXOX",
			expectedOutput: "X",
			expectedErr:    nil,
		},
		{
			name:           "Player O wins most of the time",
			input:          "OXOXO",
			expectedOutput: "O",
			expectedErr:    nil,
		},
		{
			name:           "Player X wins most of the time in a row",
			input:          "OXXXOXO",
			expectedOutput: "X",
			expectedErr:    nil,
		},
		{
			name:           "Player O wins most of the time in a row",
			input:          "OXOOOOXO",
			expectedOutput: "O",
			expectedErr:    nil,
		},
		{
			name:           "Both players win the same number of times",
			input:          "OOXXOXOXOX",
			expectedOutput: "tie",
		},
		{
			name:           "Both players win the same consecutive times in a row",
			input:          "OOOXXX",
			expectedOutput: "tie",
			expectedErr:    nil,
		},
		{
			name:           "Accept only valid characters",
			input:          "129BSD",
			expectedOutput: "",
			expectedErr:    ErrInvalidChars,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := Solution(tc.input)
			if err != nil {
				if !reflect.DeepEqual(tc.expectedErr.Error(), err.Error()) {
					t.Fatalf("%s: expected: %#v, got: %#v", tc.name, tc.expectedOutput, got)
				}
			}

			if !reflect.DeepEqual(got, tc.expectedOutput) {
				t.Fatalf("%s: expected: %#v, got: %#v", tc.name, tc.expectedOutput, got)
			}
		})
	}

}
