package sum

func Sum(numbers []int) (sum int) {
	for _, number := range numbers {
		sum += number
	}
	
	return
}

func SumAll(numbersToSum ...[]int) (sums []int) {
	// sums := make([]int, len(numbersToSum))
	
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	
	return
}