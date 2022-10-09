package myArryas

import "fmt"

/*239. 滑动窗口最大值
给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。

返回 滑动窗口中的最大值 。



示例 1：

输入：nums = [1,3,-1,-3,5,3,6,7], k = 3
输出：[3,3,5,5,6,7]
解释：
滑动窗口的位置                最大值
---------------               -----
[1  3  -1] -3  5  3  6  7       3
1 [3  -1  -3] 5  3  6  7       3
1  3 [-1  -3  5] 3  6  7       5
1  3  -1 [-3  5  3] 6  7       5
1  3  -1  -3 [5  3  6] 7       6
1  3  -1  -3  5 [3  6  7]      7
示例 2：

输入：nums = [1], k = 1
输出：[1]


提示：

1 <= nums.length <= 105
-104 <= nums[i] <= 104
1 <= k <= nums.length
通过次数365,137提交次数729,871*/

func MaxSlidingWindow() {
	//输入：nums = [1,3,-1,-3,5,3,6,7], k = 3
	//输出：[3,3,5,5,6,7]
	//[1,-1]
	//	1
	//[7,2,4]
	//	2
	nums := make([]int, 3)
	nums[0] = 7
	nums[1] = 2
	nums[2] = 4
	newNums := maxSlidingWindow(nums, 2)
	fmt.Printf("newNums: %v", newNums)
}

func maxSlidingWindow(nums []int, k int) []int {
	if k == 1 {
		return nums
	}
	newLen := len(nums) - k + 1
	newNums := make([]int, newLen)
	newIndex := 0
	left := 0
	right := left + k - 1

	// 标记一个是否选举标志位,初始时需要选举
	needElection := 0
	// 确定首个区间中最大值（选举）；移动时确保最大值是否在最左侧；如果最左侧被移除则重新选举；如果不在最左侧，则直接比较当前最大和新加入者。
	thisSectionMaxIndex := 0
	// 首次选举
	for j := left; j <= right; j++ {
		if nums[thisSectionMaxIndex] <= nums[j] {
			// 找最大值
			thisSectionMaxIndex = j
			// 选举出当前最大值以及最大下标，最大值记录到newNums中
		}
	}
	// 首次选举结束.
	newNums[newIndex] = nums[thisSectionMaxIndex]
	newIndex++
	right++
	left++
	// 直接从第一次移动开始
	for ; newIndex < newLen; newIndex++ {
		if thisSectionMaxIndex+1 == left {
			// 本次移动中，left刚好大于等于thisSectionMaxIndex，说明需要重新选举
			needElection = 1
		} else {
			// 不需要选举，只需要将新加入的value与当前进行比较
			if nums[thisSectionMaxIndex] < nums[right] {
				thisSectionMaxIndex = right
			}
			newNums[newIndex] = nums[thisSectionMaxIndex]
		}
		if needElection == 1 {
			// 选举当前区间最大值所在下标
			thisSectionMaxIndex = left
			for j := left; j <= right; j++ {
				if nums[thisSectionMaxIndex] <= nums[j] {
					thisSectionMaxIndex = j
					// 选举出当前最大值以及最大下标，最大值记录到newNums中
					newNums[newIndex] = nums[thisSectionMaxIndex]
					//newIndex++
				}
			}
			// 选举完置为1
			needElection = 0
		}
		left++
		right++
	}
	return newNums
}
