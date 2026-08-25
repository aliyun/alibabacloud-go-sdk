// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUsersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListUsersRequest
	GetMaxResults() *int32
	SetName(v string) *ListUsersRequest
	GetName() *string
	SetNameLike(v string) *ListUsersRequest
	GetNameLike() *string
	SetNextToken(v string) *ListUsersRequest
	GetNextToken() *string
}

type ListUsersRequest struct {
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// user-01
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// user
	NameLike *string `json:"nameLike,omitempty" xml:"nameLike,omitempty"`
	// example:
	//
	// dXNlci1vZmZzZXQ6MTA
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListUsersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListUsersRequest) GoString() string {
	return s.String()
}

func (s *ListUsersRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListUsersRequest) GetName() *string {
	return s.Name
}

func (s *ListUsersRequest) GetNameLike() *string {
	return s.NameLike
}

func (s *ListUsersRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListUsersRequest) SetMaxResults(v int32) *ListUsersRequest {
	s.MaxResults = &v
	return s
}

func (s *ListUsersRequest) SetName(v string) *ListUsersRequest {
	s.Name = &v
	return s
}

func (s *ListUsersRequest) SetNameLike(v string) *ListUsersRequest {
	s.NameLike = &v
	return s
}

func (s *ListUsersRequest) SetNextToken(v string) *ListUsersRequest {
	s.NextToken = &v
	return s
}

func (s *ListUsersRequest) Validate() error {
	return dara.Validate(s)
}
