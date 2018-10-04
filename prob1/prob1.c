/*
 * Copyright (c) 2018, Ankit R Gadiya
 * BSD 3-Clause License
 *
 * Project Euler
 *
 * Problem 1: Multiples of 3 and 5
 *
 * Q. If we list all the natural numbers below 10 that are multiples of 3 or 5,
 *    we get 3, 5, 6 and 9. The sum of these multiples is 23.  Find the sum of
 *    all the multiples of 3 or 5 below 1000.
 */
#include <stdio.h>

#define LIMIT 1000 - 1

int sumAP(int x, int limit);

int main(void)
{
	printf("%d\n", sumAP(3, LIMIT) + sumAP(5, LIMIT) - sumAP(15, LIMIT));

	return 0;
}

int sumAP(int x, int limit)
{
	int n = (int) limit / x;
	return (int) (x * n * (n + 1) / 2);
}
