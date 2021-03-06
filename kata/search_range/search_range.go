package search_range

func SearchRange(nums []int, target int) []int {
	result := []int{-1, -1}
	for i, num := range nums {
		if num == target {
			if result[0] == -1 {
				result[0] = i
			}
			result[1] = i
		} else if num > target {
			break
		}
	}
  return result
}
