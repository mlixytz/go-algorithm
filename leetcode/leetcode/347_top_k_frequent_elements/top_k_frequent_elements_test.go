package leetcode347

import "testing"

func TestSolution(t *testing.T) {
	if !isEqual(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2), []int{1, 2}) {
		t.Error("test error")
	}
	if !isEqual(topKFrequent([]int{1}, 1), []int{1}) {
		t.Error("test error")
	}
}

func isEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
