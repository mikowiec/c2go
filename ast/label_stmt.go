package ast

type LabelStmt struct {
	Address  string
	Position string
	Name     string
	Children []Node
}

func parseLabelStmt(line string) *LabelStmt {
	groups := groupsFromRegex(
		"<(?P<position>.*)> '(?P<name>.*)'",
		line,
	)

	return &LabelStmt{
		Address:  groups["address"],
		Position: groups["position"],
		Name:     groups["name"],
		Children: []Node{},
	}
}

// AddChild adds a new child node. Child nodes can then be accessed with the
// Children attribute.
func (n *LabelStmt) AddChild(node Node) {
	n.Children = append(n.Children, node)
}
