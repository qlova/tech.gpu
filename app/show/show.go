// Package show provides user interface tree nodes.
package show

import "qlova.tech/app/then"

type Node interface {
	node()
}

type View struct {
	isNode

	Nodes []Node
	Hints struct {
		Row bool
	}
}

type isNode = Node

type String struct {
	isNode

	Value string
	Hints struct {
		Title bool
	}
}

type pickable interface {
	~string
}

type Picker[T pickable] struct {
	isNode

	Value *T
	Hints struct {
	}
}

type Choice struct {
	isNode

	Steps then.Steps
	Hints struct {
		Button  string
		OnClick bool
		OnTouch bool
		OnHover bool
	}
}