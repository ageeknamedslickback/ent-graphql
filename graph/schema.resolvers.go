package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"ent-graphql/ent"
	"ent-graphql/graph/generated"
	"ent-graphql/graph/model"
)

func (*mutationResolver) CreateTodo(ctx context.Context, todo model.TodoInput) (*ent.Todo, error) {
	client := ent.FromContext(ctx)
	return client.Todo.
		Create().
		SetText(todo.Text).
		SetStatus(todo.Status).
		SetNillablePriority(todo.Priority).
		SetNillableParentID(todo.Parent).
		Save(ctx)
}

func (r *queryResolver) Todos(ctx context.Context) ([]*ent.Todo, error) {
	return r.client.Debug().Todo.Query().All(ctx)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
