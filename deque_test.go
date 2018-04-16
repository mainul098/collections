package collections

import (
	"testing"
)

func TestDequeAppend(t *testing.T) {
	sampleSize := 100
	q := New(-1)
	t.Log("Should create a dequeue with no capacity")
	{
		for i := 0; i < sampleSize; i++ {
			q.Append(i)
		}

		if q.Length() != sampleSize {
			t.Errorf("Queue should be infinite sized. Expected %d but Got %d", sampleSize, q.Length())
		}
	}
	t.Logf("Create a Queue. Final queue length %d", q.Length())
}

func TestDequeExtend(t *testing.T) {

	// Conversion between a []string and a []interface{} requires the memory layout to be changed and happens in O(n) time.
	// Converting a type to an interface{} requires O(1) time. If they made this for loop unnecessary,
	// the compiler would still need to insert it.
	sampleItems := []string{"hello", "test", "world", "lots", "of", "hack"}
	items := make([]interface{}, 0)
	for item := range sampleItems {
		items = append(items, item)
	}

	q := New(-1)
	t.Log("Should extend the deque with a array items")
	{
		q.Extend(items...)
		if q.Length() != len(sampleItems) {
			t.Errorf("Queue should be infinite sized. Expected %d but Got %d", len(sampleItems), q.Length())
		}
	}
	t.Logf("Extend the Queue with a slice of items. Final queue length %d", q.Length())
}

func TestDequeItems(t *testing.T) {
	sampleSize := 100
	q := New(0)
	t.Log("Should get the dequeue items")
	{
		for i := 0; i < sampleSize; i++ {
			q.Append(i)
		}
	}
	t.Logf("Return the items in Queue %d", q.Items())
}

func TestFixedDequeAppend(t *testing.T) {
	maxlen := 3
	sampleSize := 100
	q := New(maxlen)
	t.Logf("Should create a dequeue with capacity of %d", maxlen)
	{
		for i := 1; i < sampleSize; i++ {
			q.Append(i)
		}

		if q.Length() != maxlen {
			t.Errorf("Queue should be fixed sized. Expected %d but Got %d", maxlen, q.Length())
		}
	}
	t.Logf("Create a fixed sized Queue %d", q.Length())
}

func TestClearDeque(t *testing.T) {
	sampleSize := 100
	q := New(0)
	t.Log("Should Clear the Deque")
	{
		for i := 0; i < sampleSize; i++ {
			q.Append(i)
		}

		q.Clear()
		l := q.Length()
		if l != 0 {
			t.Errorf("Should clear the Dequeu. Expected : 0 but Got : %d", l)
		}
	}

	t.Log("Should clear the Deque")
}

func TestQueue(t *testing.T) {
	sampleItems := []int{2, 3, 4, 5, 6}
	q := New(0)

	t.Log("Should Enqueue and Dequeue elements on the Queue")
	{
		for _, v := range sampleItems {
			q.Append(v)
		}

		if q.Length() != len(sampleItems) {
			t.Error("Should enqueue all element in the Queue.")
		}

		t.Log("Should dequeue element from the Queue")
		{
			v, _ := q.PopLeft()
			if v != sampleItems[0] {
				t.Errorf("Should dequeue the first item from the Queue. Expected %d, but Got : %d.", sampleItems[0], v)
			}
			t.Logf("Should dequeue the first item %d", v)
		}
	}
}

func TestStack(t *testing.T) {
	sampleItems := []int{2, 3, 4, 5, 6}
	q := New(0)

	t.Log("Should Push and Pop elements on the Stack")
	{
		for _, v := range sampleItems {
			q.Append(v)
		}

		if q.Length() != len(sampleItems) {
			t.Error("Should Push all element in the Stack.")
		}

		t.Log("Should pop the last element from the Stack")
		{
			v, _ := q.Pop()
			if v != sampleItems[len(sampleItems)-1] {
				t.Errorf("Should pop the last item from the Stack. Expected %d, but Got : %d.", sampleItems[0], v)
			}
			t.Logf("Should pop the last item %d", v)
		}
	}
}
