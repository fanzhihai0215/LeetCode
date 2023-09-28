package list

import "fmt"

// ListNode 单向链表 头指针、头节点（方便链表增删的）、首元节点（第一个有效数据节点）
type ListNode struct {
	Val  int       //数据域
	Next *ListNode //指针域
}

func NewNode(val int) *ListNode {
	return &ListNode{
		Val:  val,
		Next: nil,
	}
}

func ShowList(p *ListNode) { //遍历链表值
	for p != nil {
		fmt.Println(*p)
		p = p.Next
	}
	fmt.Println()
}

//创建一个链表通过数组, 首字母小写表示private
func createLikedList(arr []int) *ListNode {
	head := ListNode{}
	p := &head
	for _, v := range arr {
		//创建新的节点
		node := ListNode{Val: v, Next: nil}
		p.Next = &node
		p = p.Next
	}
	return head.Next
}
