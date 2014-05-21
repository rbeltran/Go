package main

import "fmt"

type LinkedList struct {
	head *ListNode
	tail *ListNode
	size int
}

type ListNode struct {
	value interface{}
	next *ListNode
	prev *ListNode
}

func (list *LinkedList) Size() int {
	return list.size
}

// func NewLinkedList() *LinkedList {
// 	return &LinkedList{}
// }

func( list *LinkedList ) Add( val interface{} ) {
	if list.size  == 0 {
		list.head = &ListNode{ val, nil, nil }
		list.tail = list.head
	} else {
		newNode := &ListNode{ val, nil, nil }
		list.tail.next = newNode
		newNode.prev = list.tail
		list.tail = newNode
	}
	list.size++
}

func( list *LinkedList ) Print() {
	node := list.head
	for node != nil {
		fmt.Println(node.value, " \n ")
		node = node.next
	}
}

func main() {

	list := new(LinkedList)
	list.Add("hello")
	list.Add("good")
	list.Add("sir")
	list.Add(1)
	list.Add(2.1)

	list.Print()
}