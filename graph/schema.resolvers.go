package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.30

import (
	"context"
	"crypto/rand"
	"fmt"
	"math/big"

	"github.com/mccuskero/go-gqlgen-user-registration/graph/model"
)

// RegisterUser is the resolver for the registerUser field.
func (r *mutationResolver) RegisterUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	rand, _ := rand.Int(rand.Reader, big.NewInt(100))

	user := &model.User{
		ID:   input.UserID,
		Name: fmt.Sprintf("User%d", rand),
	}

	r.users = append(r.users, user)

	return user, nil
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	return r.users, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
