package list

import "testing"

func checkListLen(t *testing.T, l *List, len int) bool {
	if n := l.Len(); n != len {
		t.Errorf("l.Len() = %d want %d ", n, len)
		return false
	}
	return true
}

func checkListPointers(t *testing.T, l *List, ns []*Node) {
	root := &l.root
	if !checkListLen(t, l, len(ns)) {
		return
	}

	if len(ns) == 0 {
		if l.root.next != nil && l.root.prev != root {
			t.Errorf("l.root.next = %p", l.root.next)
		}
		return
	}
}
func TestList(t *testing.T) {
	l := New()
	checkListPointers(t, l, []*Node{})

}
