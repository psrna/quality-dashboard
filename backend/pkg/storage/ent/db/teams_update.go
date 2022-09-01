// Code generated by entc, DO NOT EDIT.

package db

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/redhat-appstudio/quality-studio/pkg/storage/ent/db/predicate"
	"github.com/redhat-appstudio/quality-studio/pkg/storage/ent/db/repository"
	"github.com/redhat-appstudio/quality-studio/pkg/storage/ent/db/teams"
)

// TeamsUpdate is the builder for updating Teams entities.
type TeamsUpdate struct {
	config
	hooks    []Hook
	mutation *TeamsMutation
}

// Where appends a list predicates to the TeamsUpdate builder.
func (tu *TeamsUpdate) Where(ps ...predicate.Teams) *TeamsUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetTeamName sets the "team_name" field.
func (tu *TeamsUpdate) SetTeamName(s string) *TeamsUpdate {
	tu.mutation.SetTeamName(s)
	return tu
}

// SetDescription sets the "description" field.
func (tu *TeamsUpdate) SetDescription(s string) *TeamsUpdate {
	tu.mutation.SetDescription(s)
	return tu
}

// AddRepositoryIDs adds the "repositories" edge to the Repository entity by IDs.
func (tu *TeamsUpdate) AddRepositoryIDs(ids ...uuid.UUID) *TeamsUpdate {
	tu.mutation.AddRepositoryIDs(ids...)
	return tu
}

// AddRepositories adds the "repositories" edges to the Repository entity.
func (tu *TeamsUpdate) AddRepositories(r ...*Repository) *TeamsUpdate {
	ids := make([]uuid.UUID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return tu.AddRepositoryIDs(ids...)
}

// Mutation returns the TeamsMutation object of the builder.
func (tu *TeamsUpdate) Mutation() *TeamsMutation {
	return tu.mutation
}

// ClearRepositories clears all "repositories" edges to the Repository entity.
func (tu *TeamsUpdate) ClearRepositories() *TeamsUpdate {
	tu.mutation.ClearRepositories()
	return tu
}

// RemoveRepositoryIDs removes the "repositories" edge to Repository entities by IDs.
func (tu *TeamsUpdate) RemoveRepositoryIDs(ids ...uuid.UUID) *TeamsUpdate {
	tu.mutation.RemoveRepositoryIDs(ids...)
	return tu
}

// RemoveRepositories removes "repositories" edges to Repository entities.
func (tu *TeamsUpdate) RemoveRepositories(r ...*Repository) *TeamsUpdate {
	ids := make([]uuid.UUID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return tu.RemoveRepositoryIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TeamsUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(tu.hooks) == 0 {
		affected, err = tu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TeamsMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tu.mutation = mutation
			affected, err = tu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tu.hooks) - 1; i >= 0; i-- {
			if tu.hooks[i] == nil {
				return 0, fmt.Errorf("db: uninitialized hook (forgotten import db/runtime?)")
			}
			mut = tu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TeamsUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TeamsUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TeamsUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tu *TeamsUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   teams.Table,
			Columns: teams.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: teams.FieldID,
			},
		},
	}
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.TeamName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: teams.FieldTeamName,
		})
	}
	if value, ok := tu.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: teams.FieldDescription,
		})
	}
	if tu.mutation.RepositoriesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   teams.RepositoriesTable,
			Columns: []string{teams.RepositoriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: repository.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.RemovedRepositoriesIDs(); len(nodes) > 0 && !tu.mutation.RepositoriesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   teams.RepositoriesTable,
			Columns: []string{teams.RepositoriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: repository.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.RepositoriesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   teams.RepositoriesTable,
			Columns: []string{teams.RepositoriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: repository.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{teams.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// TeamsUpdateOne is the builder for updating a single Teams entity.
type TeamsUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TeamsMutation
}

// SetTeamName sets the "team_name" field.
func (tuo *TeamsUpdateOne) SetTeamName(s string) *TeamsUpdateOne {
	tuo.mutation.SetTeamName(s)
	return tuo
}

// SetDescription sets the "description" field.
func (tuo *TeamsUpdateOne) SetDescription(s string) *TeamsUpdateOne {
	tuo.mutation.SetDescription(s)
	return tuo
}

// AddRepositoryIDs adds the "repositories" edge to the Repository entity by IDs.
func (tuo *TeamsUpdateOne) AddRepositoryIDs(ids ...uuid.UUID) *TeamsUpdateOne {
	tuo.mutation.AddRepositoryIDs(ids...)
	return tuo
}

// AddRepositories adds the "repositories" edges to the Repository entity.
func (tuo *TeamsUpdateOne) AddRepositories(r ...*Repository) *TeamsUpdateOne {
	ids := make([]uuid.UUID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return tuo.AddRepositoryIDs(ids...)
}

// Mutation returns the TeamsMutation object of the builder.
func (tuo *TeamsUpdateOne) Mutation() *TeamsMutation {
	return tuo.mutation
}

// ClearRepositories clears all "repositories" edges to the Repository entity.
func (tuo *TeamsUpdateOne) ClearRepositories() *TeamsUpdateOne {
	tuo.mutation.ClearRepositories()
	return tuo
}

// RemoveRepositoryIDs removes the "repositories" edge to Repository entities by IDs.
func (tuo *TeamsUpdateOne) RemoveRepositoryIDs(ids ...uuid.UUID) *TeamsUpdateOne {
	tuo.mutation.RemoveRepositoryIDs(ids...)
	return tuo
}

// RemoveRepositories removes "repositories" edges to Repository entities.
func (tuo *TeamsUpdateOne) RemoveRepositories(r ...*Repository) *TeamsUpdateOne {
	ids := make([]uuid.UUID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return tuo.RemoveRepositoryIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *TeamsUpdateOne) Select(field string, fields ...string) *TeamsUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Teams entity.
func (tuo *TeamsUpdateOne) Save(ctx context.Context) (*Teams, error) {
	var (
		err  error
		node *Teams
	)
	if len(tuo.hooks) == 0 {
		node, err = tuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TeamsMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tuo.mutation = mutation
			node, err = tuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(tuo.hooks) - 1; i >= 0; i-- {
			if tuo.hooks[i] == nil {
				return nil, fmt.Errorf("db: uninitialized hook (forgotten import db/runtime?)")
			}
			mut = tuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TeamsUpdateOne) SaveX(ctx context.Context) *Teams {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TeamsUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TeamsUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tuo *TeamsUpdateOne) sqlSave(ctx context.Context) (_node *Teams, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   teams.Table,
			Columns: teams.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: teams.FieldID,
			},
		},
	}
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Teams.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, teams.FieldID)
		for _, f := range fields {
			if !teams.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("db: invalid field %q for query", f)}
			}
			if f != teams.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tuo.mutation.TeamName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: teams.FieldTeamName,
		})
	}
	if value, ok := tuo.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: teams.FieldDescription,
		})
	}
	if tuo.mutation.RepositoriesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   teams.RepositoriesTable,
			Columns: []string{teams.RepositoriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: repository.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.RemovedRepositoriesIDs(); len(nodes) > 0 && !tuo.mutation.RepositoriesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   teams.RepositoriesTable,
			Columns: []string{teams.RepositoriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: repository.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.RepositoriesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   teams.RepositoriesTable,
			Columns: []string{teams.RepositoriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: repository.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Teams{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{teams.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
