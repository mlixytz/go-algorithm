package leetcode043

import (
	"strconv"
)

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	ans := make([]int, len(num1)+len(num2))
	carray := 0
	for i := len(num1) - 1; i >= 0; i-- {
		carray = 0
		for j := len(num2) - 1; j >= 0; j-- {
			a, _ := strconv.Atoi(num1[i : i+1])
			b, _ := strconv.Atoi(num2[j : j+1])
			mul := a * b
			mul += carray + ans[i+j+1]
			if mul >= 10 {
				carray = mul / 10
				ans[i+j+1] = mul % 10
			} else {
				carray = 0
				ans[i+j+1] = mul
			}
		}
		ans[i] = carray
	}
	i := 0
	for i < len(ans) && ans[i] == 0 {
		i++
	}
	result := ""
	for j := i; j < len(ans); j++ {
		result += strconv.Itoa(ans[j])
	}
	return result
}
