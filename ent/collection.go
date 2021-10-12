// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
)

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (c *CarQuery) CollectFields(ctx context.Context, satisfies ...string) *CarQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		c = c.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return c
}

func (c *CarQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *CarQuery {
	return c
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (gr *GroupQuery) CollectFields(ctx context.Context, satisfies ...string) *GroupQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		gr = gr.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return gr
}

func (gr *GroupQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *GroupQuery {
	return gr
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (u *UserQuery) CollectFields(ctx context.Context, satisfies ...string) *UserQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		u = u.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return u
}

func (u *UserQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *UserQuery {
	return u
}