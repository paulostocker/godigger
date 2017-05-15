package godigger

type DbAttr struct {
	Fields         []string
	Conditions     []string
	Joins          []string
	WithAttributes bool
}

func New() DbAttr {
	return DbAttr{nil, nil, nil, false}
}

// Field
func (attr DbAttr) Field(fields ...string) DbAttr {
	attr.Fields = fields
	return attr
}

// Condition
func (attr DbAttr) Condition(conditions ...string) DbAttr {
	attr.Conditions = conditions
	return attr
}

// Join
func (attr DbAttr) Join(joins ...string) DbAttr {
	attr.Joins = joins
	return attr
}
