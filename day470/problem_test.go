package day470

import "testing"

// nolint
var testcases = []struct {
	nums        []int
	targetIndex int
	expected    int
}{
	{[]int{4, 1, 3, 5, 6}, 0, 3},
	{[]int{4, 1, 3, 5, 6}, 4, -1},
	{[]int{4, 1, 3, 5, 6}, 1, 0},
}

func TestClosestLargerNumberIndexBrute(t *testing.T) {
	t.Parallel()

	for _, tc := range testcases {
		if result := ClosestLargerNumberIndexBrute(tc.nums, tc.targetIndex); result != tc.expected {
			t.Errorf("Expected %v got %v", tc.expected, result)
		}
	}
}

func TestClosestLargerNumberIndexBruteBoundsCheck(t *testing.T) {
	t.Parallel()

	defer func() {
		if err := recover(); err == nil {
			t.Errorf("It should panic with an index out of bounds")
		}
	}()
	ClosestLargerNumberIndexBrute([]int{1, 2}, 3)
}

func BenchmarkClosestLargerNumberIndexBrute(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			ClosestLargerNumberIndexBrute(tc.nums, tc.targetIndex)
		}
	}
}

func TestPreprocessClosestLargerNumberIndex(t *testing.T) {
	t.Parallel()

	for _, tc := range testcases {
		pre := PreprocessClosestLargerNumberIndex(tc.nums)
		if result := pre[tc.targetIndex]; result != tc.expected {
			t.Errorf("Expected %v got %v", tc.expected, result)
		}
	}
}

func BenchmarkPreprocessClosestLargerNumberIndex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			PreprocessClosestLargerNumberIndex(tc.nums)
		}
	}
}
