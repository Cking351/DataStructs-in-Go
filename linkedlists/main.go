package main

import (
	"fmt"
)


type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
	length int
}

func (l *LinkedList) prepend(n *Node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}

func (l *LinkedList) deleteWithValue(value int) {
	if l.length == 0 {
		return
	}
	
	if l.head.data == value {
		l.head = l.head.next
		l.length--
		return
	}
	
	previousToDelete := l.head
	for previousToDelete.next.data != value {
		if previousToDelete.next.next == nil {
			return
		}
		previousToDelete = previousToDelete.next
	}
	previousToDelete.next = previousToDelete.next.next 
	l.length--
}

func (l LinkedList) printListData() {
	toPrint := l.head
	for l.length != 0 {
		fmt.Printf("%d ", toPrint.data)
		toPrint = toPrint.next
		l.length--
	}
	fmt.Printf("\n")
}

func main() {
	fmt.Println("Hello")
	myList := LinkedList{}
	node1 := &Node{data:48}
	node2 := &Node{data:67}
	node3 := &Node{data:34}
	node4 := &Node{data: 456}
	node5 := &Node{data: 33}
	myList.prepend(node1)
	myList.prepend(node2)
	myList.prepend(node3)
	myList.prepend(node4)
	myList.prepend(node5)
	myList.printListData()

	myList.deleteWithValue(33)
	myList.printListData()
}