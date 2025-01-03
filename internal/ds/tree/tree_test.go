package tree

import (
	"fmt"
	"testing"

	"github.com/ashtanko/go-algorithms/utils"
	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestTreeSize(t *testing.T) {
	var testCases = []struct {
		tree     *TreeNode
		expected int
	}{
		{
			tree:     &TreeNode{},
			expected: 1,
		},
		{
			tree:     &TreeNode{Val: 1},
			expected: 1,
		},
		{
			tree: &TreeNode{
				Val:   1,
				Left:  &TreeNode{},
				Right: &TreeNode{},
			},
			expected: 3,
		},
		{
			tree: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 2,
					},
				},
				Right: &TreeNode{Val: 5},
			},
			expected: 4,
		},
	}

	for _, tc := range testCases {
		name := fmt.Sprintf("Tree: %d", tc.tree.List())
		t.Run(name, func(t *testing.T) {
			actual := tc.tree.Size()
			utils.Checkf(t, is.Equal(tc.expected, actual), tc)
		})
	}
}

func TestDfs(t *testing.T) {
	tree := &TreeNode{
		Val: 10,
		Left: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
		Right: &TreeNode{
			Val: 15,
			Right: &TreeNode{
				Val: 18,
			},
		},
	}
	actual := Dfs(tree, 7, 15)
	assert.Equal(t, 32, actual)
}

func TestBfs(t *testing.T) {
	testCases := []struct {
		nodes    []*TreeNode
		expected []int
	}{
		{
			[]*TreeNode{
				{
					Val: 0,
				},
			},
			[]int{0},
		},
		{
			[]*TreeNode{
				{
					Val: 0,
					Left: &TreeNode{
						Val: 1,
					},
					Right: &TreeNode{
						Val: 2,
					},
				},
			},
			[]int{0},
		},
	}
	for _, tc := range testCases {
		name := fmt.Sprintf("e: %d", tc.expected)
		t.Run(name, func(t *testing.T) {
			level, _ := Bfs(tc.nodes)
			assert.DeepEqual(t, tc.expected, level)
		})
	}
}

func TestDiagram(t *testing.T) {
	var testCases = []struct {
		tree     *TreeNode
		expected string
	}{
		{
			&TreeNode{
				Val: 10,
			},
			"10\n",
		},
		{
			&TreeNode{
				Val: 0,
				Left: &TreeNode{
					Val: 1,
				},
			},
			"┌──null\n0\n└──1\n",
		},
		{
			&TreeNode{
				Val: 0,
				Left: &TreeNode{
					Val: 1,
				},
				Right: &TreeNode{
					Val: 2,
				},
			},
			"┌──2\n0\n└──1\n",
		},
	}

	for _, tc := range testCases {
		name := fmt.Sprintf("%d %s", tc.tree.List(), tc.expected)
		t.Run(name, func(t *testing.T) {
			actual := tc.tree.Diagram("", "", "")
			assert.DeepEqual(t, tc.expected, actual)
		})
	}
}
