package collections

import (
	"container/list"
	"errors"
	"fmt"
	"sync"
)

// Define the custom error types for the package
var (
	ErrEmptyQueue = errors.New("Queue is empty")
)

// Deque is a Double Ended Queue. It enables the Stack and Queue data structure property.
// Maxlen can be enabled to make a capped version of Queue.
// If maxlen is set to 0 , deques may grow to an arbitrary length.
// Otherwise, the deque is bounded to the specified maximum length.
// Once a bounded length deque is full, when new items are added, a corresponding number of items are discarded from the opposite end.
type Deque struct {
	sync.RWMutex
	container *list.List
	capacity  int
}

// New method returns a new deque object initialized left-to-right with data from sclice.
// If maxlen is set 0, length of the queue is not specified, the new deque is empty.
func New(maxlen int) *Deque {
	if maxlen < 0 {
		maxlen = 0
	}

	q := Deque{
		container: list.New(),
		capacity:  maxlen,
	}

	return &q
}

// Print the items of the List. Dont print the list unless using for debug purpose.
// The print will take O(n) time to traverse the list
func (q *Deque) String() string {
	return fmt.Sprintf("%+v", q.Items())
}

// Items return the list of item. Dont use this unless using for debug purpose.
// This will take O(n) time to traverse the list
func (q *Deque) Items() []interface{} {
	q.Lock()
	defer q.Unlock()

	items := make([]interface{}, 0, q.container.Len())

	for e := q.container.Front(); e != nil; e = e.Next() {
		items = append(items, e.Value)
	}

	return items
}

// Length return the Queue length
func (q *Deque) Length() int {
	q.Lock()
	defer q.Unlock()

	return q.container.Len()
}

// Append will add an item at the front of the list
func (q *Deque) Append(item interface{}) {
	q.Lock()
	defer q.Unlock()

	if q.capacity > 0 && q.container.Len() == q.capacity {
		firstItem := q.container.Back()
		q.container.Remove(firstItem)
	}

	q.container.PushFront(item)
}

// AppendLeft will add an item at the back of the list
func (q *Deque) AppendLeft(item interface{}) {
	q.Lock()
	defer q.Unlock()

	if q.capacity > 0 && q.container.Len() == q.capacity {
		lastItem := q.container.Front()
		q.container.Remove(lastItem)
	}

	q.container.PushBack(item)
}

// Clear the queue
func (q *Deque) Clear() {
	q.Lock()
	defer q.Unlock()

	q.container.Init()
}

// Count the item in the List
func (q *Deque) Count(item interface{}) int {
	q.Lock()
	defer q.Unlock()

	count := 0
	for e := q.container.Front(); e != nil; e = e.Next() {
		if e.Value == item {
			count++
		}
	}

	return count
}

// Extend will add a list of items in the Queue
func (q *Deque) Extend(items ...interface{}) {
	q.Lock()
	defer q.Unlock()

	for item := range items {
		if q.capacity > 0 && q.container.Len() == q.capacity {
			firstItem := q.container.Back()
			q.container.Remove(firstItem)
		}

		q.container.PushFront(item)
	}
}

// Pop will return the Last insteretd item from the List
// If the list is empty return ErrEmptyQueue
func (q *Deque) Pop() (interface{}, error) {
	q.Lock()
	defer q.Unlock()

	lastItem := q.container.Front()
	if lastItem != nil {
		item := q.container.Remove(lastItem)
		return item, nil
	}

	return nil, ErrEmptyQueue

}

// PopLeft will return the First item from the List
// If the list is empty return ErrEmptyQueue
func (q *Deque) PopLeft() (interface{}, error) {
	q.Lock()
	defer q.Unlock()

	firstItem := q.container.Back()
	if firstItem != nil {
		item := q.container.Remove(firstItem)
		return item, nil
	}

	return nil, ErrEmptyQueue
}

// Remove the item from the Queue
func (q *Deque) Remove(item interface{}) error {
	q.Lock()
	defer q.Unlock()

	e := list.Element{
		Value: item,
	}

	if v := q.container.Remove(&e); v == nil {
		return ErrEmptyQueue
	}

	return nil
}
