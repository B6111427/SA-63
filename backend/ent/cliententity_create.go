// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/B6111427/app/ent/booking"
	"github.com/B6111427/app/ent/cliententity"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// ClientEntityCreate is the builder for creating a ClientEntity entity.
type ClientEntityCreate struct {
	config
	mutation *ClientEntityMutation
	hooks    []Hook
}

// SetCLIENTNAME sets the CLIENT_NAME field.
func (cec *ClientEntityCreate) SetCLIENTNAME(s string) *ClientEntityCreate {
	cec.mutation.SetCLIENTNAME(s)
	return cec
}

// SetCLIENTSTATUS sets the CLIENT_STATUS field.
func (cec *ClientEntityCreate) SetCLIENTSTATUS(s string) *ClientEntityCreate {
	cec.mutation.SetCLIENTSTATUS(s)
	return cec
}

// AddBookedIDs adds the booked edge to Booking by ids.
func (cec *ClientEntityCreate) AddBookedIDs(ids ...int) *ClientEntityCreate {
	cec.mutation.AddBookedIDs(ids...)
	return cec
}

// AddBooked adds the booked edges to Booking.
func (cec *ClientEntityCreate) AddBooked(b ...*Booking) *ClientEntityCreate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return cec.AddBookedIDs(ids...)
}

// Mutation returns the ClientEntityMutation object of the builder.
func (cec *ClientEntityCreate) Mutation() *ClientEntityMutation {
	return cec.mutation
}

// Save creates the ClientEntity in the database.
func (cec *ClientEntityCreate) Save(ctx context.Context) (*ClientEntity, error) {
	if _, ok := cec.mutation.CLIENTNAME(); !ok {
		return nil, &ValidationError{Name: "CLIENT_NAME", err: errors.New("ent: missing required field \"CLIENT_NAME\"")}
	}
	if v, ok := cec.mutation.CLIENTNAME(); ok {
		if err := cliententity.CLIENTNAMEValidator(v); err != nil {
			return nil, &ValidationError{Name: "CLIENT_NAME", err: fmt.Errorf("ent: validator failed for field \"CLIENT_NAME\": %w", err)}
		}
	}
	if _, ok := cec.mutation.CLIENTSTATUS(); !ok {
		return nil, &ValidationError{Name: "CLIENT_STATUS", err: errors.New("ent: missing required field \"CLIENT_STATUS\"")}
	}
	if v, ok := cec.mutation.CLIENTSTATUS(); ok {
		if err := cliententity.CLIENTSTATUSValidator(v); err != nil {
			return nil, &ValidationError{Name: "CLIENT_STATUS", err: fmt.Errorf("ent: validator failed for field \"CLIENT_STATUS\": %w", err)}
		}
	}
	var (
		err  error
		node *ClientEntity
	)
	if len(cec.hooks) == 0 {
		node, err = cec.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ClientEntityMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cec.mutation = mutation
			node, err = cec.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cec.hooks) - 1; i >= 0; i-- {
			mut = cec.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cec.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (cec *ClientEntityCreate) SaveX(ctx context.Context) *ClientEntity {
	v, err := cec.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (cec *ClientEntityCreate) sqlSave(ctx context.Context) (*ClientEntity, error) {
	ce, _spec := cec.createSpec()
	if err := sqlgraph.CreateNode(ctx, cec.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	ce.ID = int(id)
	return ce, nil
}

func (cec *ClientEntityCreate) createSpec() (*ClientEntity, *sqlgraph.CreateSpec) {
	var (
		ce    = &ClientEntity{config: cec.config}
		_spec = &sqlgraph.CreateSpec{
			Table: cliententity.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: cliententity.FieldID,
			},
		}
	)
	if value, ok := cec.mutation.CLIENTNAME(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: cliententity.FieldCLIENTNAME,
		})
		ce.CLIENTNAME = value
	}
	if value, ok := cec.mutation.CLIENTSTATUS(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: cliententity.FieldCLIENTSTATUS,
		})
		ce.CLIENTSTATUS = value
	}
	if nodes := cec.mutation.BookedIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   cliententity.BookedTable,
			Columns: []string{cliententity.BookedColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: booking.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return ce, _spec
}