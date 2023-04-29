package datastruct

import "errors"

// 栈功能
type Stack struct {
	data []int // 栈
	size int   // 栈大小
}

func NewStack(size int) *Stack {
	return &Stack{
		data: make([]int, 0),
		size: size,
	}
}

func (s *Stack) Len() int {
	return len(s.data)
}

func (s *Stack) IsEmpty() bool {
	if s.Len() == 0 {
		return true
	}

	return false
}

func (s *Stack) First() (val int, err error) {
	if s.IsEmpty() {
		return 0, errors.New("error! stack is empty")
	}

	return s.data[s.Len()-1], nil
}

func (s *Stack) Push(val int) error {
	if s.Len() == s.size {
		return errors.New("error! stack is full")
	}
	s.data = append(s.data, val)
	return nil
}

func (s *Stack) Pop() (val int, err error) {
	if s.IsEmpty() {
		return 0, errors.New("error! stack is empty")
	}

	// 获取最后一个数据
	val = s.data[s.Len()-1]
	// 将最后一个元素删除
	s.data = s.data[:s.Len()-1]
	return val, nil
}
