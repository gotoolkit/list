package list

// Node 链表节点
type Node struct {
	next, prev *Node

	list  *List
	Value interface{}
}

// List 链表
type List struct {
	root Node
	len  int
}

// Init 初始化链表
func (l *List) Init() *List {
	l.root.next = &l.root
	l.root.prev = &l.root
	l.len = 0
	return l
}

// New 初始化链表
func New() *List {
	return new(List).Init()
}

// Len 链表长度
func (l *List) Len() int {
	return l.len
}

// Front 链表最前节点
func (l *List) Front() *Node {
	if l.len == 0 {
		return nil
	}
	return l.root.next
}

// Back 链表最后节点
func (l *List) Back() *Node {
	if l.len == 0 {
		return nil
	}
	return l.root.prev
}

func (l *List) lazyInit() {
	if l.root.next == nil {
		l.Init()
	}
}

func (l *List) insert(n, at *Node) *Node {
	x := at.next
	at.next = n
	n.prev = at
	n.next = x
	x.prev = n
	n.list = l
	l.len++
	return n
}

func (l *List) insertValue(v interface{}, at *Node) *Node {
	return l.insert(&Node{Value: v}, at)
}

func (l *List) remove(n *Node) *Node {
	n.prev.next = n.next
	n.next.prev = n.prev
	n.next = nil
	n.prev = nil
	n.list = nil
	l.len--
	return n
}

// MoveToFront 移动节点到最前端
func (l *List) MoveToFront(n *Node) {
	if n.list != l || l.root.prev == n {
		return
	}
	l.insert(l.remove(n), &l.root)
}

// Remove 链表移除节点
func (l *List) Remove(n *Node) interface{} {
	if n.list == l {
		l.remove(n)
	}
	return n.Value
}

// PushFront 插入值为v的新节点到最前段
func (l *List) PushFront(v interface{}) *Node {
	l.lazyInit()
	return l.insertValue(v, &l.root)
}

// PushBack 插入值为v的新节点到最后段
func (l *List) PushBack(v interface{}) *Node {
	l.lazyInit()
	return l.insertValue(v, l.root.prev)
}
