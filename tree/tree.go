package tree

import (
	"fmt"
	"math"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewNode(val int) *TreeNode {
	return &TreeNode{
		Val:   val,
		Left:  nil,
		Right: nil,
	}
}

// BuildTree 构建树并且返回树的头节点
// arr 树节点的扁平化数组，其中-1表示空节点
// index 当前节点在数组中的下标
// n 数组长度
func BuildTree(arr []int, index, n int) *TreeNode {
	if n == 0 || index >= n || arr[index] == math.MinInt {
		return nil
	}
	return &TreeNode{
		Val:   arr[index],
		Left:  BuildTree(arr, index*2+1, n),
		Right: BuildTree(arr, index*2+2, n),
	}
}

func printTree(root *TreeNode) {
	if root == nil {
		return
	}
	q := []*TreeNode{root}
	for len(q) > 0 {
		size := len(q)
		for i := 0; i < size; i++ {
			node := q[0]
			q = q[1:]
			fmt.Printf("%v ", node.Val)
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		fmt.Println()
	}
}
