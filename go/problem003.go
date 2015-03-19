/*
Largest prime factor

The prime factors of 13195 are 5, 7, 13 and 29.

What is the largest prime factor of the number 600851475143 ?
*/

package main

import "fmt"

func main() {
	fmt.Println("Answer = ", factors(600851475143))
}

func factors(num int) int {
	i := 2

	for num%i != 0 {
		i++
	}

	if num/i > 1 {
		return factors(num / i)
	} else {
		return num
	}
}
