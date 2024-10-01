package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	doubledArray := transformNumbers(numbers, doubleNumber)
	tripledArray := transformNumbers(numbers, tripleNumber)

	fmt.Println("Original", numbers)
	fmt.Println("Doubled", doubledArray)
	fmt.Println("Tripled", tripledArray)
}

func transformNumbers(numbers []int, fn func(int) int) []int {
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
