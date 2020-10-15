// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/B6111427/app/ent/bookingtype"
	"github.com/facebookincubator/ent/dialect/sql"
)

// Bookingtype is the model entity for the Bookingtype schema.
type Bookingtype struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// BOOKTYPENAME holds the value of the "BOOKTYPE_NAME" field.
	BOOKTYPENAME string `json:"BOOKTYPE_NAME,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the BookingtypeQuery when eager-loading is set.
	Edges BookingtypeEdges `json:"edges"`
}

// BookingtypeEdges holds the relations/edges for other nodes in the graph.
type BookingtypeEdges struct {
	// Booktype holds the value of the booktype edge.
	Booktype []*Booking
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// BooktypeOrErr returns the Booktype value or an error if the edge
// was not loaded in eager-loading.
func (e BookingtypeEdges) BooktypeOrErr() ([]*Booking, error) {
	if e.loadedTypes[0] {
		return e.Booktype, nil
	}
	return nil, &NotLoadedError{edge: "booktype"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Bookingtype) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // BOOKTYPE_NAME
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Bookingtype fields.
func (b *Bookingtype) assignValues(values ...interface{}) error {
	if m, n := len(values), len(bookingtype.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	b.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field BOOKTYPE_NAME", values[0])
	} else if value.Valid {
		b.BOOKTYPENAME = value.String
	}
	return nil
}

// QueryBooktype queries the booktype edge of the Bookingtype.
func (b *Bookingtype) QueryBooktype() *BookingQuery {
	return (&BookingtypeClient{config: b.config}).QueryBooktype(b)
}

// Update returns a builder for updating this Bookingtype.
// Note that, you need to call Bookingtype.Unwrap() before calling this method, if this Bookingtype
// was returned from a transaction, and the transaction was committed or rolled back.
func (b *Bookingtype) Update() *BookingtypeUpdateOne {
	return (&BookingtypeClient{config: b.config}).UpdateOne(b)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (b *Bookingtype) Unwrap() *Bookingtype {
	tx, ok := b.config.driver.(*txDriver)
	if !ok {
		panic("ent: Bookingtype is not a transactional entity")
	}
	b.config.driver = tx.drv
	return b
}

// String implements the fmt.Stringer.
func (b *Bookingtype) String() string {
	var builder strings.Builder
	builder.WriteString("Bookingtype(")
	builder.WriteString(fmt.Sprintf("id=%v", b.ID))
	builder.WriteString(", BOOKTYPE_NAME=")
	builder.WriteString(b.BOOKTYPENAME)
	builder.WriteByte(')')
	return builder.String()
}

// Bookingtypes is a parsable slice of Bookingtype.
type Bookingtypes []*Bookingtype

func (b Bookingtypes) config(cfg config) {
	for _i := range b {
		b[_i].config = cfg
	}
}