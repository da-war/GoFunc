package main

import "fmt"

type funco func(int) int

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	moreNumbers := []int{2, 2, 3, 4, 5}
	doubledArray := transformNumbers(numbers, doubleNumber)
	tripledArray := transformNumbers(numbers, tripleNumber)

	fmt.Println("Original", numbers)
	fmt.Println("Doubled", doubledArray)
	fmt.Println("Tripled", tripledArray)

	// Function returning function
	fn := transformFunReturn(moreNumbers)
	fmt.Println("Function returning function", transformNumbers(moreNumbers, fn))
}

func transformNumbers(numbers []int, fn funco) []int {
	transformed := []int{}
	for i := range numbers {
		transformed = append(transformed, fn(numbers[i]))
	}
	return transformed
}

func doubleNumbers(numbers []int) []int {
	dNumber := []int{}
	for i := range numbers {
		dNumber = append(dNumber, numbers[i]*2)
	}
	fmt.Println("Doubled", dNumber)
	return dNumber

}

func doubleNumber(n int) int {
	return n * 2
}

func tripleNumber(n int) int {
	return n * 3
}

func transformFunReturn(numbers []int) funco {
	if numbers[0] == 1 {
		return doubleNumber
	} else {
		return tripleNumber
	}
}
