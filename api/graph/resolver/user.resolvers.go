package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/skinnyguy/spectra/api/graph"
	"github.com/skinnyguy/spectra/api/graph/model"
)

func (r *userResolver) FetchUser(ctx context.Context, obj *model.AbstractModel, in model.InputUser) (*model.UserResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *userMutationResolver) AddUser(ctx context.Context, obj *model.AbstractModel, in model.InputUser) (*model.UserResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

// User returns graph.UserResolver implementation.
func (r *Resolver) User() graph.UserResolver { return &userResolver{r} }

// UserMutation returns graph.UserMutationResolver implementation.
func (r *Resolver) UserMutation() graph.UserMutationResolver { return &userMutationResolver{r} }

type userResolver struct{ *Resolver }
type userMutationResolver struct{ *Resolver }
