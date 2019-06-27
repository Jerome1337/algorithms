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

		fmt.Println(isInFibonacciSequence(givenNumber))

		return
	}

	fmt.Println(sequence())
}

func sequence() []int {
	lastNumber, currentNumber, maxNumber := 0, 1, 10
	sequence := []int{lastNumber, currentNumber}

	for index := 0; index <= maxNumber; index++ {
		newNumber := lastNumber + currentNumber

		sequence = append(sequence, newNumber)

		lastNumber = currentNumber
		currentNumber = newNumber
	}

	return sequence
}

func isInFibonacciSequence(number int) bool {
	return isSquareNumber(5*number*number+4) ||
		isSquareNumber(5*number*number-4)
}

func isSquareNumber(number int) bool {
	square := int(math.Sqrt(float64(number)))

	return square*square == number
}
