package b

func FindUniqueValue(numbers []int) int {
	unique := 0
	for _, number := range numbers {

		unique ^= number
	}
	return unique
}
