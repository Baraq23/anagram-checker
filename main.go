package main

import (
	"fmt"
	"os"
	"strings"
)


// AnnagramCheker() takes two strings and returns true if they are anagram.
func AnagramChecker(s1, s2 string) bool {
	return true
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . <word1> <word2")
		os.Exit(1)
	}

	str1 := os.Args[1]
	str2 := os.Args[2]

	

	if isStr := ValidateStr(str1, str2); isStr {
		if str1 == str2 {
			fmt.Println("Provide different words to check")
			os.Exit(0)
		}
		isAnagram := AnagramChecker(str1, str2)
		if isAnagram {
			fmt.Printf("%s is an anagram to %s\n", str1, str2)
		} else {
			fmt.Printf("%s is Not an anagram to %s\n", str1, str2)
		}
	}	
	

}


// Function ValidateStr() takes two strings and check if they only consist of letters
func ValidateStr(s1, s2 string) bool {
	str1 := strings.ToLower(s1)
	str2 := strings.ToLower(s2)

	for _, c := range str1 {
		if !(c >= 'a' && c <= 'z') {
			fmt.Printf("%s is not a word.\n", s1)
			return false
		}
	}
	for _, c := range str2 {
		if !(c >= 'a' && c <= 'z') {
			fmt.Printf("%s is not a word.\n", s2)
			return false
		}
	}
	return true
}
