package collections

import (
	"testing"
)

func TestEmptyCounter(t *testing.T) {
	c := Counter{}
	t.Log("Should create an Empty Counter")
	{
		if c.Length() != 0 {
			t.Errorf("Should create an Empty Counter. Expected : 0, but Got : %d", c.Length())
		}
	}
}

func TestCounterCountItems(t *testing.T) {
	sampleItems := []string{"a", "b", "c", "a"}
	c := Counter{}
	t.Log("Should get the disticnt item Count")
	{
		for _, item := range sampleItems {
			c.Add(item)
		}

		if c.Length() != len(sampleItems)-1 {
			t.Errorf("Should get the distinct item Count. Expected : %d, but Got : %d", len(sampleItems)-1, c.Length())
		}
		t.Logf("Should get the distinct item Count %d", c.Length())
	}
}

func TestCounterAddItems(t *testing.T) {
	sampleItems := []string{"a", "b", "c", "a"}
	items := make([]interface{}, 0)
	for _, item := range sampleItems {
		items = append(items, item)
	}

	c := Counter{}
	t.Log("Should create an empty Counter")
	{
		c.AddItems(items...)
		if c.Length() != len(sampleItems)-1 {
			t.Errorf("Should add items in the Counter. Expected : %d, but Got : %d", len(sampleItems)-1, c.Length())
		}
		t.Logf("Should add items %d", c.Length())
	}
}
