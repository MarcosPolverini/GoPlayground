package stack

import (
	"fmt"
	"testing"
)

//Exemple of test in go language
func TestStack(t *testing.T) {
	var s Stack
	s.Push("world!")
	s.Push("Hello, ")
	for s.Size() > 0 {
		fmt.Print(s.Pop())
	}
	fmt.Println()
}
