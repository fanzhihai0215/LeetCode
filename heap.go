package main

type hp struct {
	data         [][2]int
	nums1, nums2 []int
}

func (h hp) Len() int { return len(h.data) }
func (h hp) Less(i, j int) bool {
	return h.nums1[h.data[i][0]]+h.nums2[h.data[i][1]] < h.nums1[h.data[j][0]]+h.nums2[h.data[j][1]]
}
func (h hp) Swap(i, j int)       { h.data[i], h.data[j] = h.data[j], h.data[i] }
func (h *hp) Push(v interface{}) { h.data = append(h.data, v.([2]int)) }
func (h *hp) Pop() interface{}   { a := h.data; v := a[len(a)-1]; h.data = a[:len(a)-1]; return v }
