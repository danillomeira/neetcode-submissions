func twoSum(nums []int, target int) []int {
    seen := make(map[int]int)

	for index, num := range nums {
		difference := target - num

		if previousIndex, found := seen[difference]; found {
			return []int{previousIndex, index}
		}

		seen[num] = index
	}

	return []int{}
}
