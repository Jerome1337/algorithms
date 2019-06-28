package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		fmt.Println("Is the value a Palimdrome:", isPalindrome(string(os.Args[1])))

		return
	}

	fmt.Println("You must provide a word or number")
}

func isPalindrome(value string) bool {
	valueAsRunes := []rune(value)

	for index, length := 0, len(valueAsRunes)-1; index < len(valueAsRunes)/2; index, length = index+1, length-1 {
		valueAsRunes[index], valueAsRunes[length] = valueAsRunes[length], valueAsRunes[index]
	}

	return value == string(valueAsRunes)
}
