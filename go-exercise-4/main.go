package main

import "fmt"

func main() {
	var result int
	var n int

	for {
		fmt.Print("Enter a positive number: ")
		fmt.Scan(&n)

		if n >= 0 {
			break
		}
		fmt.Println("Invalid input! Please enter a non-negative number.")

	}

	for i := 1; i <= n; i++ {
		if result == 0 {
			result = i
			continue
		} else {
			result *= i
			continue
		}
	}

	fmt.Println("Factorial of 5 is:", result)
}
