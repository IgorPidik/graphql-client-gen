package tweet

// generated by github.com/emicklei/graphql-client-gen/cmd/gcg on Tue, 25 Jan 2022 21:46:00 CET
// DO NOT EDIT

import (
	"time"
)

var (
	_ = time.Now
)

// Tweet is a OBJECT.
type Tweet struct {
	ID         interface{}              `graphql:"id" json:"id,omitempty"`
	Body       *string                  `graphql:"body" json:"body,omitempty"`
	Date       *Date                    `graphql:"date" json:"date,omitempty"`
	Author     *User                    `graphql:"Author" json:"Author,omitempty"`
	Stats      *Stat                    `graphql:"Stats" json:"Stats,omitempty"`
	Responders *TweetRespondersFunction `graphql:"Responders" json:"Responders,omitempty"`
}

// User is a OBJECT.
type User struct {
	ID        interface{} `graphql:"id" json:"id,omitempty"`
	Username  *string     `graphql:"username" json:"username,omitempty"`
	FirstName *string     `graphql:"first_name" json:"first_name,omitempty"`
	LastName  *string     `graphql:"last_name" json:"last_name,omitempty"`
	FullName  *string     `graphql:"full_name" json:"full_name,omitempty"`
	Name      *string     `graphql:"name" json:"name,omitempty"`
	AvatarUrl *Url        `graphql:"avatar_url" json:"avatar_url,omitempty"`
}

// Stat is a OBJECT.
type Stat struct {
	Views     *int32 `graphql:"views" json:"views,omitempty"`
	Likes     *int32 `graphql:"likes" json:"likes,omitempty"`
	Retweets  *int32 `graphql:"retweets" json:"retweets,omitempty"`
	Responses *int32 `graphql:"responses" json:"responses,omitempty"`
}

// Notification is a OBJECT.
type Notification struct {
	ID   *interface{} `graphql:"id" json:"id,omitempty"`
	Date *Date        `graphql:"date" json:"date,omitempty"`
	Type *string      `graphql:"type" json:"type,omitempty"`
}

// Meta is a OBJECT.
type Meta struct {
	Count *int32 `graphql:"count" json:"count,omitempty"`
}
