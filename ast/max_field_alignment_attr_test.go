package ast

import (
	"testing"
)

func TestMaxFieldAlignmentAttr(t *testing.T) {
	nodes := map[string]Node{
		`0x7fd4b7063ac0 <<invalid sloc>> Implicit 32`: &MaxFieldAlignmentAttr{
			Address:  "0x7fd4b7063ac0",
			Position: "<invalid sloc>",
			Size:     32,
			Children: []Node{},
		},
		`0x7fd4b7063ac0 <<invalid sloc>> Implicit 8`: &MaxFieldAlignmentAttr{
			Address:  "0x7fd4b7063ac0",
			Position: "<invalid sloc>",
			Size:     8,
			Children: []Node{},
		},
	}

	runNodeTests(t, nodes)
}
