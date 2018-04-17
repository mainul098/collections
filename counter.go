package collections

// Counter typr to Count the items
type Counter map[interface{}]int

// Add an item in the Counter
func (c Counter) Add(item interface{}) *Counter {
	c[item]++
	return &c
}

// AddItems add a slice of items in the Counter
func (c Counter) AddItems(items ...interface{}) *Counter {
	for _, item := range items {
		c[item]++
	}

	return &c
}

// Length of the Counter
func (c Counter) Length() int {
	return len(c)
}
