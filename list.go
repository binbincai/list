package list

type ListNode struct {
    Val interface{}
    Prev *ListNode
    Next *ListNode
}

func NewListNode(val interface{}) *ListNode {
    var node ListNode
    node.Val = val
    return &node
}

type List struct {
    Head *ListNode
}

func NewList() *List {
    var list List
    list.Head = NewListNode(nil)
    list.Head.Prev = list.Head
    list.Head.Next = list.Head
    return &list
}

func (list *List) Empty() bool {
    return list.Head.Next == list.Head
}

func (list *List) Push(node *ListNode) {
    node.Next = list.Head
    node.Prev = list.Head.Prev
    list.Head.Prev.Next = node
    list.Head.Prev = node
}

func (list *List) PushFront(node *ListNode) {
    node.Prev = list.Head
    node.Next = list.Head.Next
    list.Head.Next.Prev = node
    list.Head.Next = node
}

func RemoveNode(node *ListNode) {
    node.Next.Prev = node.Prev
    node.Prev.Next = node.Next
    node.Next = nil
    node.Prev = nil
}

func (list *List) Pop() *ListNode {
    if list.Empty() {
        return nil
    }
    node := list.Head.Prev
    RemoveNode(node)
    return node
}

func (list *List) PopFront() *ListNode {
    if list.Empty() {
        return nil
    }
    node := list.Head.Next
    RemoveNode(node)
    return node
}

func (list *List) First() *ListNode {
    return list.Head.Next
}

func (list *List) Last() *ListNode {
    return list.Head.Prev
}
