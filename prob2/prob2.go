// Copyright (c) 2018, Ankit R Gadiya
// BSD 3-Clause License
//
// Project Euler
//
// Problem 2: Even Fibonacci numbers
//
// Q. Each new term in the Fibonacci sequence is generated by adding the
//    previous two terms. By starting with 1 and 2, the first 10 terms will be:
//
//    1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...
//
//    By considering the terms in the Fibonacci sequence whose values do not
//    exceed four million, find the sum of the even-valued terms.
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("%d\n", sumEvenFib(4000000))
}

func sumEvenFib(limit int) int {
	if limit == 0 || limit == 1 {
		return limit
	}

	phi3 := math.Pow((1 + (math.Sqrt(5))) / 2, 3)
	a := 2
	fibSum := 0

	for a < limit {
		fibSum += a
		a = int(math.Round(float64(a) * phi3))
	}

	return fibSum
}