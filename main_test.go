package main

import (
	"testing"
)

// TestAnagramChecker tests the AnagramChecker function.
func TestAnagramChecker(t *testing.T) {
	tests := []struct {
		s1       string
		s2       string
		expected bool
	}{
		{"listen", "silent", true},
		{"triangle", "integral", true},
		{"apple", "pale", false},
		{"abc", "cab", true},
		{"abc", "cba", true},
		{"hello", "world", false},
		{"anagram", "nagaram", true},
	}

	for _, test := range tests {
		result := AnagramChecker(test.s1, test.s2)
		if result != test.expected {
			t.Errorf("AnagramChecker(%s, %s) = %v; want %v", test.s1, test.s2, result, test.expected)
		}
	}
}

// TestValidateStr tests the ValidateStr function.
func TestValidateStr(t *testing.T) {
	tests := []struct {
		s1       string
		s2       string
		expected bool
	}{
		{"hello", "world", true},
		{"Hello", "World", true},
		{"123", "abc", false},
		{"!@#", "def", false},
		{"abc", "def", true},
	}

	for _, test := range tests {
		result := ValidateStr(test.s1, test.s2)
		if result != test.expected {
			t.Errorf("ValidateStr(%s, %s) = %v; want %v", test.s1, test.s2, result, test.expected)
		}
	}
}


// TestFibonacciNum tests the FibonacciNum function
// func TestFibonacciNum(t *testing.T) {
// 	tests := []struct {
// 		input    int
// 		expected int
// 	}{
// 		{0, 0},
// 		{1, 1},
// 		{2, 1},
// 		{3, 2},
// 		{4, 3},
// 		{5, 5},
// 		{10, 55},
// 		{20, 6765},
// 		{30, 832040},
// 	}

// 	for _, test := range tests {
// 		result := FibonacciNum(test.input)
// 		if result != test.expected {
// 			t.Errorf("FibonacciNum(%d) = %d; expected %d", test.input, result, test.expected)
// 		}
// 	}
// }
