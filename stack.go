package stack

type node struct {
	value interface{}
	next  *node
}

// Stack is a Last In First Out data structure
type Stack struct {
	head *node
}

// NewStack Creates a new instance of a Stack
func NewStack() (s *Stack) {
	s = &Stack{}
	return s
}

// Push creates a node with a value v and appends it to the head of the stack
func (s *Stack) Push(v interface{}) {
	node := &node{value: v, next: s.head}
	s.head = node
}

// Pop removes the node at the head position and returns its value and a boolean false if head is nil
func (s *Stack) Pop() (interface{}, bool) {
	if s.head == nil {
		return nil, false
	}
	v := s.head.value
	s.head = s.head.next
	return v, true
}

// Peek returns the value of the node at the head position of the stack
func (s *Stack) Peek() (interface{}, bool) {
	if s.head == nil {
		return nil, false
	}
	return s.head.value, true
}
