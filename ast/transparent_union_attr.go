package ast

type TransparentUnionAttr struct {
	Address  string
	Position string
	Children []Node
}

func parseTransparentUnionAttr(line string) *TransparentUnionAttr {
	groups := groupsFromRegex(`<(?P<position>.*)>`, line)

	return &TransparentUnionAttr{
		Address:  groups["address"],
		Position: groups["position"],
		Children: []Node{},
	}
}

// AddChild adds a new child node. Child nodes can then be accessed with the
// Children attribute.
func (n *TransparentUnionAttr) AddChild(node Node) {
	n.Children = append(n.Children, node)
}
