// Copyright 2023 The Infratographer Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by entc, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"go.infratographer.com/ipam-api/internal/ent/generated/ipaddress"
	"go.infratographer.com/ipam-api/internal/ent/generated/ipblock"
	"go.infratographer.com/ipam-api/internal/ent/generated/predicate"
	"go.infratographer.com/x/gidx"
)

// IPBlockUpdate is the builder for updating IPBlock entities.
type IPBlockUpdate struct {
	config
	hooks    []Hook
	mutation *IPBlockMutation
}

// Where appends a list predicates to the IPBlockUpdate builder.
func (ibu *IPBlockUpdate) Where(ps ...predicate.IPBlock) *IPBlockUpdate {
	ibu.mutation.Where(ps...)
	return ibu
}

// SetPrefix sets the "prefix" field.
func (ibu *IPBlockUpdate) SetPrefix(s string) *IPBlockUpdate {
	ibu.mutation.SetPrefix(s)
	return ibu
}

// SetAllowAutoSubnet sets the "allow_auto_subnet" field.
func (ibu *IPBlockUpdate) SetAllowAutoSubnet(b bool) *IPBlockUpdate {
	ibu.mutation.SetAllowAutoSubnet(b)
	return ibu
}

// SetNillableAllowAutoSubnet sets the "allow_auto_subnet" field if the given value is not nil.
func (ibu *IPBlockUpdate) SetNillableAllowAutoSubnet(b *bool) *IPBlockUpdate {
	if b != nil {
		ibu.SetAllowAutoSubnet(*b)
	}
	return ibu
}

// SetAllowAutoAllocate sets the "allow_auto_allocate" field.
func (ibu *IPBlockUpdate) SetAllowAutoAllocate(b bool) *IPBlockUpdate {
	ibu.mutation.SetAllowAutoAllocate(b)
	return ibu
}

// SetNillableAllowAutoAllocate sets the "allow_auto_allocate" field if the given value is not nil.
func (ibu *IPBlockUpdate) SetNillableAllowAutoAllocate(b *bool) *IPBlockUpdate {
	if b != nil {
		ibu.SetAllowAutoAllocate(*b)
	}
	return ibu
}

// AddIPAddresIDs adds the "ip_address" edge to the IPAddress entity by IDs.
func (ibu *IPBlockUpdate) AddIPAddresIDs(ids ...gidx.PrefixedID) *IPBlockUpdate {
	ibu.mutation.AddIPAddresIDs(ids...)
	return ibu
}

// AddIPAddress adds the "ip_address" edges to the IPAddress entity.
func (ibu *IPBlockUpdate) AddIPAddress(i ...*IPAddress) *IPBlockUpdate {
	ids := make([]gidx.PrefixedID, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return ibu.AddIPAddresIDs(ids...)
}

// Mutation returns the IPBlockMutation object of the builder.
func (ibu *IPBlockUpdate) Mutation() *IPBlockMutation {
	return ibu.mutation
}

// ClearIPAddress clears all "ip_address" edges to the IPAddress entity.
func (ibu *IPBlockUpdate) ClearIPAddress() *IPBlockUpdate {
	ibu.mutation.ClearIPAddress()
	return ibu
}

// RemoveIPAddresIDs removes the "ip_address" edge to IPAddress entities by IDs.
func (ibu *IPBlockUpdate) RemoveIPAddresIDs(ids ...gidx.PrefixedID) *IPBlockUpdate {
	ibu.mutation.RemoveIPAddresIDs(ids...)
	return ibu
}

// RemoveIPAddress removes "ip_address" edges to IPAddress entities.
func (ibu *IPBlockUpdate) RemoveIPAddress(i ...*IPAddress) *IPBlockUpdate {
	ids := make([]gidx.PrefixedID, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return ibu.RemoveIPAddresIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ibu *IPBlockUpdate) Save(ctx context.Context) (int, error) {
	ibu.defaults()
	return withHooks(ctx, ibu.sqlSave, ibu.mutation, ibu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ibu *IPBlockUpdate) SaveX(ctx context.Context) int {
	affected, err := ibu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ibu *IPBlockUpdate) Exec(ctx context.Context) error {
	_, err := ibu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ibu *IPBlockUpdate) ExecX(ctx context.Context) {
	if err := ibu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ibu *IPBlockUpdate) defaults() {
	if _, ok := ibu.mutation.UpdatedAt(); !ok {
		v := ipblock.UpdateDefaultUpdatedAt()
		ibu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ibu *IPBlockUpdate) check() error {
	if v, ok := ibu.mutation.Prefix(); ok {
		if err := ipblock.PrefixValidator(v); err != nil {
			return &ValidationError{Name: "prefix", err: fmt.Errorf(`generated: validator failed for field "IPBlock.prefix": %w`, err)}
		}
	}
	if _, ok := ibu.mutation.IPBlockTypeID(); ibu.mutation.IPBlockTypeCleared() && !ok {
		return errors.New(`generated: clearing a required unique edge "IPBlock.ip_block_type"`)
	}
	return nil
}

func (ibu *IPBlockUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := ibu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(ipblock.Table, ipblock.Columns, sqlgraph.NewFieldSpec(ipblock.FieldID, field.TypeString))
	if ps := ibu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ibu.mutation.UpdatedAt(); ok {
		_spec.SetField(ipblock.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := ibu.mutation.Prefix(); ok {
		_spec.SetField(ipblock.FieldPrefix, field.TypeString, value)
	}
	if value, ok := ibu.mutation.AllowAutoSubnet(); ok {
		_spec.SetField(ipblock.FieldAllowAutoSubnet, field.TypeBool, value)
	}
	if value, ok := ibu.mutation.AllowAutoAllocate(); ok {
		_spec.SetField(ipblock.FieldAllowAutoAllocate, field.TypeBool, value)
	}
	if ibu.mutation.IPAddressCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   ipblock.IPAddressTable,
			Columns: []string{ipblock.IPAddressColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(ipaddress.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ibu.mutation.RemovedIPAddressIDs(); len(nodes) > 0 && !ibu.mutation.IPAddressCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   ipblock.IPAddressTable,
			Columns: []string{ipblock.IPAddressColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(ipaddress.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ibu.mutation.IPAddressIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   ipblock.IPAddressTable,
			Columns: []string{ipblock.IPAddressColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(ipaddress.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ibu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{ipblock.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ibu.mutation.done = true
	return n, nil
}

// IPBlockUpdateOne is the builder for updating a single IPBlock entity.
type IPBlockUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *IPBlockMutation
}

// SetPrefix sets the "prefix" field.
func (ibuo *IPBlockUpdateOne) SetPrefix(s string) *IPBlockUpdateOne {
	ibuo.mutation.SetPrefix(s)
	return ibuo
}

// SetAllowAutoSubnet sets the "allow_auto_subnet" field.
func (ibuo *IPBlockUpdateOne) SetAllowAutoSubnet(b bool) *IPBlockUpdateOne {
	ibuo.mutation.SetAllowAutoSubnet(b)
	return ibuo
}

// SetNillableAllowAutoSubnet sets the "allow_auto_subnet" field if the given value is not nil.
func (ibuo *IPBlockUpdateOne) SetNillableAllowAutoSubnet(b *bool) *IPBlockUpdateOne {
	if b != nil {
		ibuo.SetAllowAutoSubnet(*b)
	}
	return ibuo
}

// SetAllowAutoAllocate sets the "allow_auto_allocate" field.
func (ibuo *IPBlockUpdateOne) SetAllowAutoAllocate(b bool) *IPBlockUpdateOne {
	ibuo.mutation.SetAllowAutoAllocate(b)
	return ibuo
}

// SetNillableAllowAutoAllocate sets the "allow_auto_allocate" field if the given value is not nil.
func (ibuo *IPBlockUpdateOne) SetNillableAllowAutoAllocate(b *bool) *IPBlockUpdateOne {
	if b != nil {
		ibuo.SetAllowAutoAllocate(*b)
	}
	return ibuo
}

// AddIPAddresIDs adds the "ip_address" edge to the IPAddress entity by IDs.
func (ibuo *IPBlockUpdateOne) AddIPAddresIDs(ids ...gidx.PrefixedID) *IPBlockUpdateOne {
	ibuo.mutation.AddIPAddresIDs(ids...)
	return ibuo
}

// AddIPAddress adds the "ip_address" edges to the IPAddress entity.
func (ibuo *IPBlockUpdateOne) AddIPAddress(i ...*IPAddress) *IPBlockUpdateOne {
	ids := make([]gidx.PrefixedID, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return ibuo.AddIPAddresIDs(ids...)
}

// Mutation returns the IPBlockMutation object of the builder.
func (ibuo *IPBlockUpdateOne) Mutation() *IPBlockMutation {
	return ibuo.mutation
}

// ClearIPAddress clears all "ip_address" edges to the IPAddress entity.
func (ibuo *IPBlockUpdateOne) ClearIPAddress() *IPBlockUpdateOne {
	ibuo.mutation.ClearIPAddress()
	return ibuo
}

// RemoveIPAddresIDs removes the "ip_address" edge to IPAddress entities by IDs.
func (ibuo *IPBlockUpdateOne) RemoveIPAddresIDs(ids ...gidx.PrefixedID) *IPBlockUpdateOne {
	ibuo.mutation.RemoveIPAddresIDs(ids...)
	return ibuo
}

// RemoveIPAddress removes "ip_address" edges to IPAddress entities.
func (ibuo *IPBlockUpdateOne) RemoveIPAddress(i ...*IPAddress) *IPBlockUpdateOne {
	ids := make([]gidx.PrefixedID, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return ibuo.RemoveIPAddresIDs(ids...)
}

// Where appends a list predicates to the IPBlockUpdate builder.
func (ibuo *IPBlockUpdateOne) Where(ps ...predicate.IPBlock) *IPBlockUpdateOne {
	ibuo.mutation.Where(ps...)
	return ibuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ibuo *IPBlockUpdateOne) Select(field string, fields ...string) *IPBlockUpdateOne {
	ibuo.fields = append([]string{field}, fields...)
	return ibuo
}

// Save executes the query and returns the updated IPBlock entity.
func (ibuo *IPBlockUpdateOne) Save(ctx context.Context) (*IPBlock, error) {
	ibuo.defaults()
	return withHooks(ctx, ibuo.sqlSave, ibuo.mutation, ibuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ibuo *IPBlockUpdateOne) SaveX(ctx context.Context) *IPBlock {
	node, err := ibuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ibuo *IPBlockUpdateOne) Exec(ctx context.Context) error {
	_, err := ibuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ibuo *IPBlockUpdateOne) ExecX(ctx context.Context) {
	if err := ibuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ibuo *IPBlockUpdateOne) defaults() {
	if _, ok := ibuo.mutation.UpdatedAt(); !ok {
		v := ipblock.UpdateDefaultUpdatedAt()
		ibuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ibuo *IPBlockUpdateOne) check() error {
	if v, ok := ibuo.mutation.Prefix(); ok {
		if err := ipblock.PrefixValidator(v); err != nil {
			return &ValidationError{Name: "prefix", err: fmt.Errorf(`generated: validator failed for field "IPBlock.prefix": %w`, err)}
		}
	}
	if _, ok := ibuo.mutation.IPBlockTypeID(); ibuo.mutation.IPBlockTypeCleared() && !ok {
		return errors.New(`generated: clearing a required unique edge "IPBlock.ip_block_type"`)
	}
	return nil
}

func (ibuo *IPBlockUpdateOne) sqlSave(ctx context.Context) (_node *IPBlock, err error) {
	if err := ibuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(ipblock.Table, ipblock.Columns, sqlgraph.NewFieldSpec(ipblock.FieldID, field.TypeString))
	id, ok := ibuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`generated: missing "IPBlock.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ibuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, ipblock.FieldID)
		for _, f := range fields {
			if !ipblock.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
			}
			if f != ipblock.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ibuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ibuo.mutation.UpdatedAt(); ok {
		_spec.SetField(ipblock.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := ibuo.mutation.Prefix(); ok {
		_spec.SetField(ipblock.FieldPrefix, field.TypeString, value)
	}
	if value, ok := ibuo.mutation.AllowAutoSubnet(); ok {
		_spec.SetField(ipblock.FieldAllowAutoSubnet, field.TypeBool, value)
	}
	if value, ok := ibuo.mutation.AllowAutoAllocate(); ok {
		_spec.SetField(ipblock.FieldAllowAutoAllocate, field.TypeBool, value)
	}
	if ibuo.mutation.IPAddressCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   ipblock.IPAddressTable,
			Columns: []string{ipblock.IPAddressColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(ipaddress.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ibuo.mutation.RemovedIPAddressIDs(); len(nodes) > 0 && !ibuo.mutation.IPAddressCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   ipblock.IPAddressTable,
			Columns: []string{ipblock.IPAddressColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(ipaddress.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ibuo.mutation.IPAddressIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   ipblock.IPAddressTable,
			Columns: []string{ipblock.IPAddressColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(ipaddress.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &IPBlock{config: ibuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ibuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{ipblock.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ibuo.mutation.done = true
	return _node, nil
}
