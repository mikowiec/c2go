package ast

type IntegerLiteral struct {
	Address  string
	Position string
	Type     string
	Value    string
	Children []Node
}

func parseIntegerLiteral(line string) *IntegerLiteral {
	groups := groupsFromRegex(
		"<(?P<position>.*)> '(?P<type>.*?)' (?P<value>\\d+)",
		line,
	)

	return &IntegerLiteral{
		Address:  groups["address"],
		Position: groups["position"],
		Type:     groups["type"],
		Value:    groups["value"],
		Children: []Node{},
	}
}

// AddChild adds a new child node. Child nodes can then be accessed with the
// Children attribute.
func (n *IntegerLiteral) AddChild(node Node) {
	n.Children = append(n.Children, node)
}
