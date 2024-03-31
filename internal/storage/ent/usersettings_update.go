// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/devldavydov/myfood/internal/storage/ent/predicate"
	"github.com/devldavydov/myfood/internal/storage/ent/usersettings"
)

// UserSettingsUpdate is the builder for updating UserSettings entities.
type UserSettingsUpdate struct {
	config
	hooks    []Hook
	mutation *UserSettingsMutation
}

// Where appends a list predicates to the UserSettingsUpdate builder.
func (usu *UserSettingsUpdate) Where(ps ...predicate.UserSettings) *UserSettingsUpdate {
	usu.mutation.Where(ps...)
	return usu
}

// SetUserid sets the "userid" field.
func (usu *UserSettingsUpdate) SetUserid(i int64) *UserSettingsUpdate {
	usu.mutation.ResetUserid()
	usu.mutation.SetUserid(i)
	return usu
}

// SetNillableUserid sets the "userid" field if the given value is not nil.
func (usu *UserSettingsUpdate) SetNillableUserid(i *int64) *UserSettingsUpdate {
	if i != nil {
		usu.SetUserid(*i)
	}
	return usu
}

// AddUserid adds i to the "userid" field.
func (usu *UserSettingsUpdate) AddUserid(i int64) *UserSettingsUpdate {
	usu.mutation.AddUserid(i)
	return usu
}

// SetCalLimit sets the "cal_limit" field.
func (usu *UserSettingsUpdate) SetCalLimit(f float64) *UserSettingsUpdate {
	usu.mutation.ResetCalLimit()
	usu.mutation.SetCalLimit(f)
	return usu
}

// SetNillableCalLimit sets the "cal_limit" field if the given value is not nil.
func (usu *UserSettingsUpdate) SetNillableCalLimit(f *float64) *UserSettingsUpdate {
	if f != nil {
		usu.SetCalLimit(*f)
	}
	return usu
}

// AddCalLimit adds f to the "cal_limit" field.
func (usu *UserSettingsUpdate) AddCalLimit(f float64) *UserSettingsUpdate {
	usu.mutation.AddCalLimit(f)
	return usu
}

// Mutation returns the UserSettingsMutation object of the builder.
func (usu *UserSettingsUpdate) Mutation() *UserSettingsMutation {
	return usu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (usu *UserSettingsUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, usu.sqlSave, usu.mutation, usu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (usu *UserSettingsUpdate) SaveX(ctx context.Context) int {
	affected, err := usu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (usu *UserSettingsUpdate) Exec(ctx context.Context) error {
	_, err := usu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (usu *UserSettingsUpdate) ExecX(ctx context.Context) {
	if err := usu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (usu *UserSettingsUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(usersettings.Table, usersettings.Columns, sqlgraph.NewFieldSpec(usersettings.FieldID, field.TypeInt))
	if ps := usu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := usu.mutation.Userid(); ok {
		_spec.SetField(usersettings.FieldUserid, field.TypeInt64, value)
	}
	if value, ok := usu.mutation.AddedUserid(); ok {
		_spec.AddField(usersettings.FieldUserid, field.TypeInt64, value)
	}
	if value, ok := usu.mutation.CalLimit(); ok {
		_spec.SetField(usersettings.FieldCalLimit, field.TypeFloat64, value)
	}
	if value, ok := usu.mutation.AddedCalLimit(); ok {
		_spec.AddField(usersettings.FieldCalLimit, field.TypeFloat64, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, usu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{usersettings.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	usu.mutation.done = true
	return n, nil
}

// UserSettingsUpdateOne is the builder for updating a single UserSettings entity.
type UserSettingsUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserSettingsMutation
}

// SetUserid sets the "userid" field.
func (usuo *UserSettingsUpdateOne) SetUserid(i int64) *UserSettingsUpdateOne {
	usuo.mutation.ResetUserid()
	usuo.mutation.SetUserid(i)
	return usuo
}

// SetNillableUserid sets the "userid" field if the given value is not nil.
func (usuo *UserSettingsUpdateOne) SetNillableUserid(i *int64) *UserSettingsUpdateOne {
	if i != nil {
		usuo.SetUserid(*i)
	}
	return usuo
}

// AddUserid adds i to the "userid" field.
func (usuo *UserSettingsUpdateOne) AddUserid(i int64) *UserSettingsUpdateOne {
	usuo.mutation.AddUserid(i)
	return usuo
}

// SetCalLimit sets the "cal_limit" field.
func (usuo *UserSettingsUpdateOne) SetCalLimit(f float64) *UserSettingsUpdateOne {
	usuo.mutation.ResetCalLimit()
	usuo.mutation.SetCalLimit(f)
	return usuo
}

// SetNillableCalLimit sets the "cal_limit" field if the given value is not nil.
func (usuo *UserSettingsUpdateOne) SetNillableCalLimit(f *float64) *UserSettingsUpdateOne {
	if f != nil {
		usuo.SetCalLimit(*f)
	}
	return usuo
}

// AddCalLimit adds f to the "cal_limit" field.
func (usuo *UserSettingsUpdateOne) AddCalLimit(f float64) *UserSettingsUpdateOne {
	usuo.mutation.AddCalLimit(f)
	return usuo
}

// Mutation returns the UserSettingsMutation object of the builder.
func (usuo *UserSettingsUpdateOne) Mutation() *UserSettingsMutation {
	return usuo.mutation
}

// Where appends a list predicates to the UserSettingsUpdate builder.
func (usuo *UserSettingsUpdateOne) Where(ps ...predicate.UserSettings) *UserSettingsUpdateOne {
	usuo.mutation.Where(ps...)
	return usuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (usuo *UserSettingsUpdateOne) Select(field string, fields ...string) *UserSettingsUpdateOne {
	usuo.fields = append([]string{field}, fields...)
	return usuo
}

// Save executes the query and returns the updated UserSettings entity.
func (usuo *UserSettingsUpdateOne) Save(ctx context.Context) (*UserSettings, error) {
	return withHooks(ctx, usuo.sqlSave, usuo.mutation, usuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (usuo *UserSettingsUpdateOne) SaveX(ctx context.Context) *UserSettings {
	node, err := usuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (usuo *UserSettingsUpdateOne) Exec(ctx context.Context) error {
	_, err := usuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (usuo *UserSettingsUpdateOne) ExecX(ctx context.Context) {
	if err := usuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (usuo *UserSettingsUpdateOne) sqlSave(ctx context.Context) (_node *UserSettings, err error) {
	_spec := sqlgraph.NewUpdateSpec(usersettings.Table, usersettings.Columns, sqlgraph.NewFieldSpec(usersettings.FieldID, field.TypeInt))
	id, ok := usuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "UserSettings.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := usuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, usersettings.FieldID)
		for _, f := range fields {
			if !usersettings.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != usersettings.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := usuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := usuo.mutation.Userid(); ok {
		_spec.SetField(usersettings.FieldUserid, field.TypeInt64, value)
	}
	if value, ok := usuo.mutation.AddedUserid(); ok {
		_spec.AddField(usersettings.FieldUserid, field.TypeInt64, value)
	}
	if value, ok := usuo.mutation.CalLimit(); ok {
		_spec.SetField(usersettings.FieldCalLimit, field.TypeFloat64, value)
	}
	if value, ok := usuo.mutation.AddedCalLimit(); ok {
		_spec.AddField(usersettings.FieldCalLimit, field.TypeFloat64, value)
	}
	_node = &UserSettings{config: usuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, usuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{usersettings.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	usuo.mutation.done = true
	return _node, nil
}