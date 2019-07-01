package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) > 1 {
		givenNumber, _ := strconv.Atoi(os.Args[1])

		fmt.Println(isPrimeNumber(givenNumber))

		return
	}

	fmt.Println(sequence())
}

func isPrimeNumber(number int) bool {
	for index := 2; index <= int(math.Floor(float64(number)/2)); index++ {
		if number%index == 0 {
			return false
		}
	}
	return number > 1
}

func sequence() []int {
	var sequence []int
	sequenceLength := 100

	for index := 0; index <= sequenceLength; index++ {
		if isPrimeNumber(index) {
			sequence = append(sequence, index)
		}
	}

	return sequence
}
