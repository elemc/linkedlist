package linkedlist_test

import (
	"fmt"
	linkedlist "linked-list"
	"testing"
)

func TestLinkedList_Reverse(t *testing.T) {
	ll := linkedlist.New()
	ll.Add(linkedlist.NewItem(1))
	ll.Add(linkedlist.NewItem(2))
	ll.Add(linkedlist.NewItem(3))
	ll.Add(linkedlist.NewItem(4))

	for item := ll.First(); item != nil; item = item.Next() {
		fmt.Println(item.String())
	}

	ll.Reverse()
	for item := ll.First(); item != nil; item = item.Next() {
		fmt.Println(item.String())
	}
}

func BenchmarkLinkedList_Reverse(b *testing.B) {
	ll := linkedlist.New()
	for i := 0; i < b.N; i++ {
		ll.Add(linkedlist.NewItem(i))
	}
	ll.Reverse()
}
