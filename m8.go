/*
Alta Regina Aliyarahmah
2IA25
50422177
*/

package main

import (
	"fmt"
)

type Queue struct {
	items []int
}

func (q *Queue) Enqueue(item int) {
	q.items = append(q.items, item)
}

func (q *Queue) Dequeue() int {
	if len(q.items) == 0 {
		fmt.Println("Antrian kosong")
		return -1
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item
}

func (q *Queue) Display() {
	if len(q.items) == 0 {
		fmt.Println("Antrian kosong")
		return
	}
	fmt.Println("Isi Antrian:")
	for _, item := range q.items {
		fmt.Print(item, " ")
	}
	fmt.Println()
}

func main() {
	queue := Queue{}

	queue.Enqueue(11)
	queue.Enqueue(22)
	queue.Enqueue(33)

	queue.Display()

	item := queue.Dequeue()
	if item != -1 {
		fmt.Println("Elemen yang dihapus:", item)
	}
	queue.Display()
}
