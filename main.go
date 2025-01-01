package main

import "fmt"

func main() {
	n1 := node{data: 1}
	n2 := node{data: 2}
	n3 := node{data: 3}

	linkedList := doublyLinkedList{
		capacity: 3,
		size:     0,
	}

	linkedList.addLast(&n1)
	linkedList.addLast(&n2)
	linkedList.addLast(&n3)
	err := linkedList.addLast(&n3)
	if err != nil {
		fmt.Println("Error:", err)
	}

	linkedList.print()

	//linkedList.removeFirst()
	linkedList.moveToLast(&n1)
	linkedList.print()
}

type node struct {
	prev *node
	next *node
	data any
}

type doublyLinkedList struct {
	capacity int
	size     int
	first    *node
	last     *node
}

func (doublyLinkedList doublyLinkedList) print() {
	fmt.Println("-------Printing List-------")
	fmt.Printf("linkedList: %v \n", doublyLinkedList)
	var curNode *node = doublyLinkedList.first
	for i := range doublyLinkedList.capacity {
		if curNode == nil {
			break
		}
		fmt.Printf("node: %p - %v \n", curNode, curNode)

		if i == doublyLinkedList.capacity {
			break
		}
		curNode = curNode.next

	}
}

func (doublyLinkedList *doublyLinkedList) addLast(newNode *node) error {
	if doublyLinkedList.capacity == doublyLinkedList.size {
		return fmt.Errorf("list is full")
	}

	if doublyLinkedList.size == 0 {
		doublyLinkedList.first = newNode
		doublyLinkedList.last = newNode
		doublyLinkedList.size += 1
		return nil
	}

	lastNode := doublyLinkedList.last
	lastNode.next = newNode
	newNode.prev = lastNode
	doublyLinkedList.last = newNode
	doublyLinkedList.size += 1

	return nil
}

func (doublyLinkedList *doublyLinkedList) removeFirst() error {
	if doublyLinkedList.size == 0 {
		return fmt.Errorf("list is empty")
	} else if doublyLinkedList.size == 1 {
		doublyLinkedList.first = nil
		doublyLinkedList.last = nil
		doublyLinkedList.size = 0
		return nil
	}

	firstNode := doublyLinkedList.first
	secondNode := firstNode.next

	firstNode.next = nil
	secondNode.prev = nil

	doublyLinkedList.first = secondNode
	doublyLinkedList.size -= 1

	return nil
}

func (doublyLinkedList *doublyLinkedList) moveToLast(node *node) error {
	if doublyLinkedList.size == 0 {
		return fmt.Errorf("list is empty")
	} else if doublyLinkedList.size == 1 {
		return nil
	} else if node == doublyLinkedList.last {
		return nil
	}

	prevNode := node.prev
	nextNode := node.next

	if prevNode == nil {
		doublyLinkedList.first = nextNode
		nextNode.prev = nil
	} else {
		prevNode.next = nextNode
		nextNode.prev = prevNode
	}

	node.prev = doublyLinkedList.last
	doublyLinkedList.last.next = node
	node.next = nil
	doublyLinkedList.last = node

	return nil
}
