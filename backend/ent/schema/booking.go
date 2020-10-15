package schema

import (
	"time"

	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Booking holds the schema definition for the Booking entity.
type Booking struct {
	ent.Schema
}

// Fields of the Booking.
func (Booking) Fields() []ent.Field {
	return []ent.Field{
		field.Time("BOOKING_DATE").
			Default(time.Now),
		field.Time("TIME_LEFT").
			Default(time.Now),
	}
}

// Edges of the Booking.
func (Booking) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("Usedby", User.Type).
			Ref("Booking").
			Unique(),
		edge.From("Book", Bookingtype.Type).
			Ref("Booktype").
			Unique(),
		edge.From("Using", ClientEntity.Type).
			Ref("Booked").
			Unique(),
	}
}
