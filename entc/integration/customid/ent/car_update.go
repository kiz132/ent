// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/entc/integration/customid/ent/car"
	"github.com/facebookincubator/ent/entc/integration/customid/ent/pet"
	"github.com/facebookincubator/ent/entc/integration/customid/ent/predicate"
	"github.com/facebookincubator/ent/schema/field"
)

// CarUpdate is the builder for updating Car entities.
type CarUpdate struct {
	config
	model        *string
	owner        map[string]struct{}
	clearedOwner bool
	predicates   []predicate.Car
}

// Where adds a new predicate for the builder.
func (cu *CarUpdate) Where(ps ...predicate.Car) *CarUpdate {
	cu.predicates = append(cu.predicates, ps...)
	return cu
}

// SetModel sets the model field.
func (cu *CarUpdate) SetModel(s string) *CarUpdate {
	cu.model = &s
	return cu
}

// SetOwnerID sets the owner edge to Pet by id.
func (cu *CarUpdate) SetOwnerID(id string) *CarUpdate {
	if cu.owner == nil {
		cu.owner = make(map[string]struct{})
	}
	cu.owner[id] = struct{}{}
	return cu
}

// SetNillableOwnerID sets the owner edge to Pet by id if the given value is not nil.
func (cu *CarUpdate) SetNillableOwnerID(id *string) *CarUpdate {
	if id != nil {
		cu = cu.SetOwnerID(*id)
	}
	return cu
}

// SetOwner sets the owner edge to Pet.
func (cu *CarUpdate) SetOwner(p *Pet) *CarUpdate {
	return cu.SetOwnerID(p.ID)
}

// ClearOwner clears the owner edge to Pet.
func (cu *CarUpdate) ClearOwner() *CarUpdate {
	cu.clearedOwner = true
	return cu
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (cu *CarUpdate) Save(ctx context.Context) (int, error) {
	if len(cu.owner) > 1 {
		return 0, errors.New("ent: multiple assignments on a unique edge \"owner\"")
	}
	return cu.sqlSave(ctx)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CarUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CarUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CarUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cu *CarUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   car.Table,
			Columns: car.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: car.FieldID,
			},
		},
	}
	if ps := cu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value := cu.model; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: car.FieldModel,
		})
	}
	if cu.clearedOwner {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   car.OwnerTable,
			Columns: []string{car.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: pet.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.owner; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   car.OwnerTable,
			Columns: []string{car.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: pet.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{car.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// CarUpdateOne is the builder for updating a single Car entity.
type CarUpdateOne struct {
	config
	id           int
	model        *string
	owner        map[string]struct{}
	clearedOwner bool
}

// SetModel sets the model field.
func (cuo *CarUpdateOne) SetModel(s string) *CarUpdateOne {
	cuo.model = &s
	return cuo
}

// SetOwnerID sets the owner edge to Pet by id.
func (cuo *CarUpdateOne) SetOwnerID(id string) *CarUpdateOne {
	if cuo.owner == nil {
		cuo.owner = make(map[string]struct{})
	}
	cuo.owner[id] = struct{}{}
	return cuo
}

// SetNillableOwnerID sets the owner edge to Pet by id if the given value is not nil.
func (cuo *CarUpdateOne) SetNillableOwnerID(id *string) *CarUpdateOne {
	if id != nil {
		cuo = cuo.SetOwnerID(*id)
	}
	return cuo
}

// SetOwner sets the owner edge to Pet.
func (cuo *CarUpdateOne) SetOwner(p *Pet) *CarUpdateOne {
	return cuo.SetOwnerID(p.ID)
}

// ClearOwner clears the owner edge to Pet.
func (cuo *CarUpdateOne) ClearOwner() *CarUpdateOne {
	cuo.clearedOwner = true
	return cuo
}

// Save executes the query and returns the updated entity.
func (cuo *CarUpdateOne) Save(ctx context.Context) (*Car, error) {
	if len(cuo.owner) > 1 {
		return nil, errors.New("ent: multiple assignments on a unique edge \"owner\"")
	}
	return cuo.sqlSave(ctx)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CarUpdateOne) SaveX(ctx context.Context) *Car {
	c, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return c
}

// Exec executes the query on the entity.
func (cuo *CarUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CarUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cuo *CarUpdateOne) sqlSave(ctx context.Context) (c *Car, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   car.Table,
			Columns: car.Columns,
			ID: &sqlgraph.FieldSpec{
				Value:  cuo.id,
				Type:   field.TypeInt,
				Column: car.FieldID,
			},
		},
	}
	if value := cuo.model; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: car.FieldModel,
		})
	}
	if cuo.clearedOwner {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   car.OwnerTable,
			Columns: []string{car.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: pet.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.owner; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   car.OwnerTable,
			Columns: []string{car.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: pet.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	c = &Car{config: cuo.config}
	_spec.Assign = c.assignValues
	_spec.ScanValues = c.scanValues()
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{car.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return c, nil
}
