package schema

import (
	"regexp"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Match(regexp.MustCompile("[a-zA-Z_]+$")),
		field.String("email").Unique().Match(regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z.\-]+\.[a-z]{2,}$`)),
		field.Int("age").Positive().Optional(),
		field.String("phone").MinLen(10).MaxLen(10),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
