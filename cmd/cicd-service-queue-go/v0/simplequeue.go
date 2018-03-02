package main

import (
	"fmt"
)

type Queue struct {
	elements []interface{}
}

func NewQueue() *Queue {
	q := new(Queue)
	q.elements = make([]interface{}, 0)
	return q
}

func (q *Queue) Enqueue(e interface{}) *Queue {
	q.elements = append(q.elements, e)
	return q
}

func (q *Queue) Dequeue() (e interface{}) {
	e, q.elements = q.elements[0], q.elements[1:len(q.elements)]
	return
}

func (q *Queue) Size() int {
	return len(q.elements)
}

func main() {
	q := NewQueue()

	q.Enqueue(1).Enqueue(2).Enqueue(3).Enqueue("Hello")
	for q.Size() > 0 {
		fmt.Println(q.Dequeue())
	}

}
