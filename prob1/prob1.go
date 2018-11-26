// Copyright (c) 2018, Ankit R Gadiya
// BSD 3-Clause License
//
// Project Euler
//
// Problem 1: Multiples of 3 and 5
//
// Q. If we list all the natural numbers below 10 that are multiples of 3 or 5,
//    we get 3, 5, 6 and 9. The sum of these multiples is 23.  Find the sum of
//    all the multiples of 3 or 5 below 1000.
package main

import "fmt"

func main() {
	limit := 1000 - 1
	fmt.Printf("%d\n", sumAP(3, limit)+sumAP(5, limit)-sumAP(15, limit))
}

func sumAP(x, limit int) int {
	n := limit / x
	return x * n * (n + 1) / 2
}
