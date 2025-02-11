package main

import "fmt"

func main() {
	var number, evenSum, oddSum int
	fmt.Print("Enter a number: ")
	fmt.Scan(&number)

	for i := 1; i <= number; i++ {
		if i%2 == 0 {
			evenSum += i
			continue
		} else {
			oddSum += i
			continue
		}
	}

	fmt.Println("Sum of Even Numbers:", evenSum)
	fmt.Println("Sum of Odd Numbers:", oddSum)
}
