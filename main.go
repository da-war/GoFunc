package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	doubleNumbers(numbers)
	fmt.Println("Original", numbers)
}

func doubleNumbers(numbers []int) []int {
	dNumber := []int{}
	for i := range numbers {
		dNumber = append(dNumber, numbers[i]*2)
	}
	fmt.Println("Doubled", dNumber)
	return dNumber

}
