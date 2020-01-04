package quicksort

func sort(arrayNums []int) []int {
	if len(arrayNums) < 2 {
		return arrayNums
	} else {
		var less []int
		var greater []int
		pivot := arrayNums[0]

		for _, element := range arrayNums[1:] {
			if element <= pivot {
				less = append(less, element)
			}
			if element > pivot {
				greater = append(greater, element)
			}

		}

		var output []int
		output = append(sort(less), pivot)
		output = append(output, sort(greater)...)
		return output
	}

}
