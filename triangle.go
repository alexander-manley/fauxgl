package fauxgl

type Triangle struct {
	V1, V2, V3 Vertex
}

func NewTriangle(v1, v2, v3 Vertex) *Triangle {
	t := Triangle{v1, v2, v3}
	t.FixNormals()
	return &t
}

func (t *Triangle) Normal() Vector {
	e1 := t.V2.Position.Sub(t.V1.Position)
	e2 := t.V3.Position.Sub(t.V1.Position)
	return e1.Cross(e2).Normalize()
}

func (t *Triangle) FixNormals() {
	n := t.Normal()
	zero := Vector{}
	if t.V1.Normal == zero {
		t.V1.Normal = n
	}
	if t.V2.Normal == zero {
		t.V2.Normal = n
	}
	if t.V3.Normal == zero {
		t.V3.Normal = n
	}
}