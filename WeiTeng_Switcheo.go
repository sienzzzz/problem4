package main

import (
	"fmt"
)

// Arithmetic Progression using three different methods in Go
// Can produce both positive and negative integer

// Summary
// 1) The first method is the most efficient one as it dont bothered by how big is the value of n
// 2) Loop is efficient, but as n value goes up, it will take more time
// 3) Recursion is less efficient with a risk of a stack overflow when n value gets bigger
// - A stack overflow occurs when a program's call stack, limited by a predefined size is exhausted due to excessive function calls,
// leading to runtime error and potential program crash

// Direct Formula
// The formula only involves simple arithmetic operations, so the time for it to compute is the same regardless of the input value
// This method is straight forward, and it is faster when dealing with big value of n
func sum_to_n_a(n int) int {
	var result int
	if n >= 0 {
		result = n * (n + 1) / 2
	} else if n < 0 {
		result = -n * (n - 1) / 2
	}
	return result
}

// Loop
// When the value of n is big, it will take more slightly longer time to compute compared to the first method
// It is straightforward and easy to understand
func sum_to_n_b(n int) int {
	var sum int

	if n >= 0 {
		sum = 0
		for i := 1; i <= n; i++ {
			sum += i
		}
	} else if n < 0 {
		for i := 1; i <= abs(n); i++ {
			sum -= i
		}
	}

	return sum
}

// To produce absolute value for method 2
func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// Recursion
// Worst case among three, it may cause stack overflow when n is big due to the overhead of function calls
func sum_to_n_c(n int) int {
	var result int
	if n == 0 {
		return 0
	}
	if n > 0 {
		result = n + sum_to_n_c(n-1)
	} else if n < 0 {
		result = n + sum_to_n_c(n+1)
	}
	return result
}

// To get user input
func GetInput() int {
	var input int
	fmt.Print("Enter either positive or negative integer: ")
	fmt.Scan(&input)
	return input
}

func main() {
	n := GetInput()

	// Method 1: Direct Formula
	result_a := sum_to_n_a(n)
	fmt.Printf("Method 1 - The sum of %d arithmetic progression is: %d \n", n, result_a)

	// Method 2: Loop Iteration
	result_b := sum_to_n_b(n)
	fmt.Printf("Method 2 - The sum of %d arithmetic progression is: %d \n", n, result_b)

	// Method 3: Recursion
	result_c := sum_to_n_c(n)
	fmt.Printf("Method 3 - The sum of %d arithmetic progression is: %d \n", n, result_c)
}
