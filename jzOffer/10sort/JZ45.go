package sort

import (
	"fmt"
	"sort"
	"strings"
)

//把数组排成最小的数
func MinNumber(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		return compareNumber(nums[i], nums[j])
	})
	var res strings.Builder
	for i := 0; i < len(nums); i++ {
		res.WriteString(fmt.Sprintf("%d", nums[i]))
	}
	return res.String()
}

func compareNumber(a, b int) bool {
	str1 := fmt.Sprintf("%d%d", a, b)
	str2 := fmt.Sprintf("%d%d", b, a)
	return str1 < str2
}
