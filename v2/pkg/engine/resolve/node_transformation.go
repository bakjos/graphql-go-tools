package resolve

import (
	"slices"

	pipeline "github.com/jensneuse/pipeline/pkg/pipe"
)

type Transformation struct {
	Nullable        bool
	UseParentObject bool
	Path            []string
	InnerValue      Node
	Pipeline        *pipeline.Pipeline
}

func (_ *Transformation) NodeKind() NodeKind {
	return NodeKindTransformation
}

func (c *Transformation) Copy() Node {
	return &Transformation{
		Pipeline:        c.Pipeline,
		Nullable:        c.Nullable,
		Path:            c.Path,
		UseParentObject: c.UseParentObject,
	}
}

func (c *Transformation) NodePath() []string {
	return c.Path
}

func (c *Transformation) NodeNullable() bool {
	return c.Nullable
}

func (c *Transformation) Equals(n Node) bool {
	other, ok := n.(*Transformation)
	if !ok {
		return false
	}

	if !slices.Equal(c.Path, other.Path) {
		return false
	}

	if c.Pipeline != other.Pipeline {
		return false
	}

	if c.Nullable != other.Nullable {
		return false
	}

	if c.UseParentObject != other.UseParentObject {
		return false
	}

	return true
}
