package main

import "fmt"

type Stack struct {
	items []int 

}

type Queue struct {
	items []int
}

// Push STACK
func (s *Stack) Push(i int) {
	s.items = append(s.items, i)
}
// Pop STACK
func (s *Stack) Pop() int {
	l := len(s.items) - 1
	targetItem := s.items[l]
	s.items = s.items[:l]
	return targetItem
}

func (q *Queue) Enqueue(i int) {
	q.items = append(q.items, i)
}
// Pop STACK
func (q *Queue) Dequeue() int {
	targetItem := q.items[0]
	q.items = q.items[1:]
	return targetItem
}

func main() {
	stack := Stack{}
	stack.Push(5)
	stack.Push(9)
	stack.Push(8)
	stack.Pop()
	fmt.Println(stack)
	queue := Queue{}
	queue.Enqueue(5)
	queue.Enqueue(8)
	queue.Enqueue(23)
	queue.Enqueue(8)
	fmt.Println(queue.Dequeue())
	fmt.Println(queue)
}