package main

import (
	"errors"
	"fmt"
	"github.com/zggxuexihao/algorithm/datastruct"
)

/**
设计一个有getMin功能的栈
题目：实现一个特殊的栈，在实现栈的基本功能的基础上，再实现返回栈中最小元素的操作
*/

type GetMinStack struct {
	dataStack *datastruct.Stack // 数据栈
	minStack  *datastruct.Stack // 最小栈
}

func NewGetMinStack(size int) *GetMinStack {
	return &GetMinStack{
		dataStack: datastruct.NewStack(size),
		minStack:  datastruct.NewStack(size),
	}
}

// 压栈
func (g *GetMinStack) Push(val int) (err error) {
	// 栈是否为空
	if g.minStack.IsEmpty() {
		// 推入最小栈
		err := g.minStack.Push(val)
		if err != nil {
			return err
		}
	} else {
		// 获取栈的最小数
		min, err := g.getMin()
		if err != nil {
			return err
		}

		// 将最小数与新数做对比
		if min > val {
			// 推入最小栈
			err := g.minStack.Push(val)
			if err != nil {
				return err
			}
		}
	}

	// 推入数据栈
	err = g.dataStack.Push(val)
	if err != nil {
		return err
	}

	return nil
}

// 弹出
func (g *GetMinStack) pop() (val int, err error) {
	if g.dataStack.IsEmpty() {
		return 0, errors.New("error! stack is empty")
	}

	// 弹出数据
	val, err = g.dataStack.Pop()
	if err != nil {
		return 0, err
	}

	min, err := g.getMin()
	if err != nil {
		return 0, err
	}

	// 如果弹出数据等于最小数栈顶，则弹出最小栈的栈顶
	if val == min {
		_, err = g.minStack.Pop()
		if err != nil {
			return 0, err
		}
	}

	return val, err
}

func (g *GetMinStack) getMin() (val int, err error) {
	return g.minStack.First()
}

func main() {
	GetMinObj := NewGetMinStack(10)
	_ = GetMinObj.Push(112)
	_ = GetMinObj.Push(45)
	_ = GetMinObj.Push(13)
	_ = GetMinObj.Push(12)
	_ = GetMinObj.Push(103)
	_ = GetMinObj.Push(20)

	min, _ := GetMinObj.getMin()
	fmt.Println(fmt.Sprintf("min=%v", min))

	one, _ := GetMinObj.pop()
	fmt.Println(fmt.Sprintf("one=%v", one))

	two, _ := GetMinObj.pop()
	fmt.Println(fmt.Sprintf("two=%v", two))

	three, _ := GetMinObj.pop()
	fmt.Println(fmt.Sprintf("three=%v", three))

	twomin, _ := GetMinObj.getMin()
	fmt.Println(fmt.Sprintf("twomin=%v", twomin))
}
