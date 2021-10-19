package linkedlist

import "sync"

type LinkedList interface {
	First() Item
	Add(Item)
	Reverse()
}

type linkedList struct {
	first Item
	count uint64

	sync.RWMutex
}

// New - new linked list create
func New() LinkedList {
	return &linkedList{}
}

// First - return first linked list item
func (ll *linkedList) First() (i Item) {
	ll.RLock()
	i = ll.first
	ll.RUnlock()
	return
}

// Add - add new item to linked list
func (ll *linkedList) Add(i Item) {
	if ll.first == nil {
		ll.first = i
		ll.count++
		return
	}

	var last Item
	for item := ll.first; item != nil; item = item.Next() {
		last = item
	}

	ll.Lock()
	last.SetNext(i)
	ll.count++
	ll.Unlock()
}

func swapItem(prev, item Item) (last Item) {
	next := item.Next()
	if next != nil {
		last = swapItem(item, next)
		next.SetNext(item)
	} else {
		last = item
	}
	return
}

// Reverse - reverse a linked list
func (ll *linkedList) Reverse() {
	last := swapItem(nil, ll.First())
	ll.first.SetNext(nil)
	ll.first = last
}
