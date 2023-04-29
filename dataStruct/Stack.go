package dataStruct

import "errors"

// 栈功能
type Stack struct {
	data []int
	num int
}

func NewStack() *Stack {
	return &Stack{data: make([]int, 0)}
}

func (s *Stack) push(val int) error  {
	if s.num == len(s.data) {
		return errors.New("error! stack is full")
	}
	s.data = append(s.data, val)
	return nil
}

func (s *Stack) pop() (num int, err error) {
	if len(s.data) == 0 {
		return 0, errors.New("error! stack is empty")
	}

	num = s.data[len(s.data) - 1]

	return num, nil
}



