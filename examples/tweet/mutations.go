package tweet

// generated by github.com/emicklei/graphql-client-gen/cmd/gcg on Mon, 24 Jan 2022 17:18:52 CET
// DO NOT EDIT

import (
	"context"
	"time"
)

var (
	_ = time.Now
)

// Example implementation: github.com/shurcooL/graphql
type clientInterface interface {
	Mutate(ctx context.Context, m interface{}, variables map[string]interface{}) error
}

var Mutation = mutation{}

type mutation struct{}

// CreateTweet is a Mutation.
func (_ mutation) CreateTweet(
	graphqlClient clientInterface,
	ctx context.Context,
	body string) (Tweet, error) {
	var m struct {
		CreateTweet struct {
			Tweet
		} `graphql:"createTweet(body: $body)"`
	}
	vars := map[string]interface{}{
		"body": body,
	}
	err := graphqlClient.Mutate(ctx, &m, vars)
	return m.CreateTweet.Tweet, err
}

// DeleteTweet is a Mutation.
func (_ mutation) DeleteTweet(
	graphqlClient clientInterface,
	ctx context.Context,
	id interface{}) (Tweet, error) {
	var m struct {
		DeleteTweet struct {
			Tweet
		} `graphql:"deleteTweet(id: $id)"`
	}
	vars := map[string]interface{}{
		"id": id,
	}
	err := graphqlClient.Mutate(ctx, &m, vars)
	return m.DeleteTweet.Tweet, err
}

// MarkTweetRead is a Mutation.
func (_ mutation) MarkTweetRead(
	graphqlClient clientInterface,
	ctx context.Context,
	id interface{}) (bool, error) {
	var m struct {
		MarkTweetRead struct {
			bool
		} `graphql:"markTweetRead(id: $id)"`
	}
	vars := map[string]interface{}{
		"id": id,
	}
	err := graphqlClient.Mutate(ctx, &m, vars)
	return m.MarkTweetRead.bool, err
}
