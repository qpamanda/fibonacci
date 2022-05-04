package main

import "fmt"

// The Fibonacci sequence is a series of numbers in which each number is the sum of the two that precede it.
// Starting at 0 and 1, the sequence looks like this: 0, 1, 1, 2, 3, 5, 8, 13, 21, 34, and so on forever.
// The Fibonacci sequence can be described using a mathematical equation: Xn+2= Xn+1 + Xn.

func fibonacci(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	return fibonacci(n-2) + fibonacci(n-1)
}

func fibonacciGenerator(n int) []int {

	output := make([]int, 0)

	if n == 1 {
		output = append(output, 0)
	} else if n == 2 {
		output = append(output, 0, 1)
	} else {
		output = append(output, 0, 1)

		for i := 2; i <= n; i++ {
			output = append(output, output[len(output)-2]+output[len(output)-1])
		}
	}
	return output
}

func main() {

	num := 10

	result := fibonacci(num)
	fmt.Printf("The output for Fibonacci series of %d is %d\n", num, result)

	output := fibonacciGenerator(num)
	fmt.Printf("The Fibonacci series of %d are %d\n", num, output)
}
