package utils

import (
	"fmt"
)

// StackMin is an interface for a stack that can return the minimum value
type StackMin interface {
	Stack[int]
	Min() int
}

// StackMinValue represents a value in a stack, and an auxiliary min value
// min is used to keep track of the minumum value in the stack at the time of the push and be able to return it in O(1) time
type StackMinValue struct {
	value int
	min   int
}

// stackMin represents a stack data structure implemented with an array that contains a return the minimum value
type stackMin struct {
	values []StackMinValue
}

// NewStackMin returns a new stack with the given capacity
func NewStackMin(capacity int) *stackMin {
	return &stackMin{
		values: make([]StackMinValue, 0, capacity),
	}
}

// Push adds a value on top of the stack
// if the value is less than the current min, it is set as the new min
func (sm *stackMin) Push(value int) {
	if sm.IsEmpty() {
		sm.values = append(sm.values, StackMinValue{value: value, min: value})

		return
	}

	stackMinVal := StackMinValue{value: value, min: sm.Min()}
	if value < sm.Min() {
		stackMinVal.min = value
	}

	sm.values = append(sm.values, stackMinVal)
}

// Pop removes the top value of the stack and returns it
func (sm *stackMin) Pop() int {
	if sm.IsEmpty() {
		// gracefully handle stack underflow
		fmt.Println("empty stack")

		return 0
	}

	value := sm.values[len(sm.values)-1].value
	sm.values = sm.values[:len(sm.values)-1]

	return value
}

// Peek returns the top value of the stack without removing it
func (sm *stackMin) Peek() int {
	if sm.IsEmpty() {
		// gracefully handle stack underflow
		fmt.Println("empty stack")

		return 0
	}

	return sm.values[len(sm.values)-1].value
}

// IsEmpty returns true if the stack is empty, false otherwise
func (sm *stackMin) IsEmpty() bool {
	return len(sm.values) == 0
}

// IsFull returns true if the stack is full, false otherwise
func (sm *stackMin) IsFull(capacity int) bool {
	return len(sm.values) == capacity
}

// Print prints the stack values
func (sm *stackMin) Print(header string) {
	if sm.IsEmpty() {
		// gracefully handle stack underflow
		fmt.Println(header, "\nempty stack")

		return
	}

	if header != "" {
		fmt.Println(header)
	}

	for _, v := range sm.values {
		fmt.Print(v.value, " ")
	}
	fmt.Println()
}

// Len returns the number of values in the stack
func (sm *stackMin) Len() int {
	return len(sm.values)
}

// Min returns the min value of the stack in O(1) time without removing it
func (sm *stackMin) Min() int {
	if sm.IsEmpty() {
		// gracefully handle stack underflow
		fmt.Println("empty stack")

		return 0
	}

	return sm.values[len(sm.values)-1].min
}
