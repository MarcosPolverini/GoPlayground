package stack

//Stack structure
type Stack struct {
	data []string
}

// Push adds x to the top of the stack.
func (s *Stack) Push(x string) {
	s.data = append(s.data, x)
}

// Pop removes and returns the top element of the stack.
// (s *Stack) == this in Java
func (s *Stack) Pop() string {
	n := len(s.data) - 1
	res := s.data[n]
	s.data[n] = ""
	s.data = s.data[:n]
	return res
}

// Size returns the number of elements in the stack.
func (s *Stack) Size() int {
	return len(s.data)
}
