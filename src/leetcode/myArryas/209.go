package myArryas

import "fmt"

/*209. 长度最小的子数组
给定一个含有 n 个正整数的数组和一个正整数 target 。

找出该数组中满足其和 ≥ target 的长度最小的 连续子数组 [numsl, numsl+1, ..., numsr-1, numsr] ，并返回其长度。如果不存在符合条件的子数组，返回 0 。



示例 1：

输入：target = 7, nums = [2,3,1,2,4,3]
输出：2
解释：子数组 [4,3] 是该条件下的长度最小的子数组。
示例 2：

输入：target = 4, nums = [1,4,4]
输出：1
示例 3：

输入：target = 11, nums = [1,1,1,1,1,1,1,1]
输出：0


提示：

1 <= target <= 109
1 <= nums.length <= 105
1 <= nums[i] <= 105


进阶：

如果你已经实现 O(n) 时间复杂度的解法, 请尝试设计一个 O(n log(n)) 时间复杂度的解法。*/

func minSubArrayLen(target int, nums []int) int {
	/*low，high，始终维护当前子数组，一旦子数组值大于或等于target，就不在维护该数组了
	最核心的思想：
	1.对于这种还有双指针的题目，始终只让一种情况做“加法”
	对于另一面的情况，只去维护区间（循环不变量）
	2.类似之前做过得最长回文子串，low的变化是有讲究的。
	（最长回文子串中，low = 重复字符在子串中位置 + 1）
	而本题中，low = low + 1；而且直观上的low = high + 1;(这个点挺易错的)
	维护sum记录当前的子数组和，tempmin维护当前最小，如果有更小，就更新。
	*/
	// 定义初始时左、右指针
	low := 0
	high := 0
	curSectionSum := 0
	curSectionLength := 0
	// 初始时最短长度记录为-1
	minLength := 0
	// 循环条件
	for high < len(nums) {
		curSectionLength = high - low + 1
		if curSectionSum+nums[high] >= target {
			if curSectionLength < minLength || minLength == 0 {
				minLength = curSectionLength
			}
			curSectionSum -= nums[low]
			low++
		} else {
			curSectionSum += nums[high]
			high++
		}
	}
	return minLength
}

func MinSubArrayLen() {
	var nums []int = make([]int, 6)
	nums[0] = 2
	nums[1] = 3
	nums[2] = 1
	nums[3] = 2
	nums[4] = 4
	nums[5] = 3
	result := minSubArrayLen(7, nums)
	fmt.Printf("result: %v", result)
}
