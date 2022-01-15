package dcp

import (
	"fmt"
	"strconv"
)

// This problem was asked by Microsoft.
// A number is considered perfect if its digits sum up to exactly 10.
// Given a positive integer n, return the n-th perfect number.
// For example, given 1, you should return 19. Given 2, you should return 28
func dcp1(n int) int {
	comp := 10 - n
	perf := fmt.Sprint(n) + fmt.Sprint(comp)
	answer, _ := strconv.Atoi(perf)
	return answer
}
