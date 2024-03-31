// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/devldavydov/myfood/internal/storage/ent/journal"
	"github.com/devldavydov/myfood/internal/storage/ent/predicate"
)

// JournalDelete is the builder for deleting a Journal entity.
type JournalDelete struct {
	config
	hooks    []Hook
	mutation *JournalMutation
}

// Where appends a list predicates to the JournalDelete builder.
func (jd *JournalDelete) Where(ps ...predicate.Journal) *JournalDelete {
	jd.mutation.Where(ps...)
	return jd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (jd *JournalDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, jd.sqlExec, jd.mutation, jd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (jd *JournalDelete) ExecX(ctx context.Context) int {
	n, err := jd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (jd *JournalDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(journal.Table, sqlgraph.NewFieldSpec(journal.FieldID, field.TypeInt))
	if ps := jd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, jd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	jd.mutation.done = true
	return affected, err
}

// JournalDeleteOne is the builder for deleting a single Journal entity.
type JournalDeleteOne struct {
	jd *JournalDelete
}

// Where appends a list predicates to the JournalDelete builder.
func (jdo *JournalDeleteOne) Where(ps ...predicate.Journal) *JournalDeleteOne {
	jdo.jd.mutation.Where(ps...)
	return jdo
}

// Exec executes the deletion query.
func (jdo *JournalDeleteOne) Exec(ctx context.Context) error {
	n, err := jdo.jd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{journal.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (jdo *JournalDeleteOne) ExecX(ctx context.Context) {
	if err := jdo.Exec(ctx); err != nil {
		panic(err)
	}
}