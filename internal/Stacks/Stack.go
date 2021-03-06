package Stacks

import "fmt"

type Stack struct {
	A 		[]interface{}
	top 	int
	size 	int
}

func New(size int) *Stack {
	return &Stack{
		A: nil,
		top: -1,
		size: size,
	}
}

func (s Stack) IsFull() (bool) {
	return s.top == s.size - 1
}

func (s Stack) IsEmpty() (bool) {
	return s.top == -1
}

func (s *Stack) Push(key interface{}) {
	if s.IsFull() {
		fmt.Printf("Stack is full. Cannot push element %v.\n", key)
		return
	}
	s.A = append(s.A, key)
	s.top++
}

func (s *Stack) Top() (interface{}) {
	if s.IsEmpty() {
		fmt.Println("Stack is empty.")
		return -1
	}
	return s.A[s.top]
}

func (s *Stack) Pop() {
	if s.IsEmpty() {
		fmt.Println("Stack is empty.")
		return
	}
	s.A = s.A[:len(s.A) - 1]	
	s.top--
}

