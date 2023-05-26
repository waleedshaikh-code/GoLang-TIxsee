package main

import "fmt"

// STACK (PUSH AND POP)

// type Stack struct {
// 	items []int
// }

// // Push

// func (s *Stack) Push(i int) {
// 	s.items = append(s.items, i)
// }

// // pop will remove the value in the end
// // and returns the removed value
// func (s *Stack) Pop() int {
// 	l := len(s.items) - 1
// 	toRemove := s.items[l]
// 	s.items = s.items[:l]
// 	return toRemove
// }

// func main() {
// 	myStack := Stack{}
// 	fmt.Println("EMPTY ARRAY \n", myStack)
// 	myStack.Push(100)
// 	myStack.Push(200)
// 	myStack.Push(300)
// 	fmt.Println("UPDATED VALUE AFTER PUSH: \n", myStack)
// 	myStack.Pop()
// 	fmt.Println("UPDATED VALUE AFTER POP: \n", myStack)

// }

// QUEUE

// SIMPLE QUEUE

type Queue struct {
	items []int
}

// ENQUEUE add A Value in the end

func (q *Queue) Enqueue(i int) {
	q.items = append(q.items, i)
}

// DEQEUE
// DEQEUE REMOVES A VALUE AT THE FRONT AND RETURNS THE REMOVAL VALUE
func (q *Queue) Dequeue() int {
	toRemove := q.items[0]
	q.items = q.items[1:]
	return toRemove

}

func main() {
	myQueue := Queue{}
	fmt.Println("Empty QUEUE ARRAY: \n", myQueue)
	myQueue.Enqueue(100)
	myQueue.Enqueue(200)
	myQueue.Enqueue(300)
	fmt.Println("AFTER ENQUEUE THE VALUE: \n", myQueue)
	myQueue.Dequeue()
	fmt.Println("FIRST VALUE *100* has been Removed(DEQUEUE) \n", myQueue)
}
