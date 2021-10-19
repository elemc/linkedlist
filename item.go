package linkedlist

import "fmt"

type Item interface {
	Next() Item
	SetNext(Item)
	String() string
}

type item struct {
	value interface{}
	next  *item
}

// NewItem - create new item with nil next value
func NewItem(value interface{}) Item {
	return &item{
		value: value,
		next:  nil,
	}
}

// Next - return next item
func (i *item) Next() (ri Item) {
	// FIXME: this is workaround, explain this code below
	if i.next == nil {
		return nil
	}
	ri = i.next
	return
}

// SetNext - set next item
func (i *item) SetNext(next Item) {
	if next == nil {
		i.next = nil
		return
	}
	switch t := next.(type) {
	case *item:
		i.next = t
	}
}

// String - return item value as a string
func (i *item) String() string {
	return fmt.Sprintf("%#v", i.value)
}
