package main

type WrapNode struct {
	node  *TreeNode
	level int
}

type Queue struct {
	Q []*WrapNode
}

func (Q *Queue) push(node *TreeNode, level int) {
	n := &WrapNode{
		node:  node,
		level: level,
	}
	Q.Q = append(Q.Q, n)
}

func (Q *Queue) pop() (*TreeNode, int) {
	n := Q.Q[0]
	Q.Q = Q.Q[1:]
	return n.node, n.level
}

func (Q *Queue) isempty() bool {
	if len(Q.Q) == 0 {
		return true
	} else {
		return false
	}
}

func levelOrder(root *TreeNode) [][]int {
	var q Queue
	var result [][]int

	if root == nil {
		return result
	}

	q.push(root, 0)

	for !q.isempty() {
		node, level := q.pop()

		if node.Left != nil {
			q.push(node.Left, level+1)
		}

		if node.Right != nil {
			q.push(node.Right, level+1)
		}

		if level >= len(result) {
			result = append(result, []int{})
		}
		result[level] = append(result[level], node.Val)
	}

	return result
}
