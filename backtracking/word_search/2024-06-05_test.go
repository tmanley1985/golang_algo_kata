package word_search

import (
	"testing"
)

func TestWordExists(t *testing.T) {
    for _, tc := range TestCases { // Import the test cases from the separate file
        result := WordExists_2024_06_05(tc.Board, tc.Word)
        if result != tc.Expected {
            t.Errorf("WordExists(%v, %s) = %t; want %t", tc.Board, tc.Word, result, tc.Expected)
        }
    }
}
