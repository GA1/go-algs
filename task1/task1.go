package task1

func SolveTask1(A []int) int {
	var min = A[0]
	var bestProfit = A[1] - A[0]
	for i := 2; i < len(A); i++ {
		if A[i - 1] < min {
			min = A[i - 1]
		}
		if bestProfit < A[i] - min {
			bestProfit = A[i] - min
		}
	}
	if bestProfit > 0 {
		return bestProfit
	} else {
		return 0
	}
}
