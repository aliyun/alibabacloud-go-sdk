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
	// The maximum number of records to return per page. Valid values: 1 to 100. If this parameter is not specified, 10 records are returned by default.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The filter condition for exact match on the username.
	//
	// example:
	//
	// user-01
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The filter condition for fuzzy match on the username. The value can be up to 128 characters in length.
	//
	// example:
	//
	// user
	NameLike *string `json:"nameLike,omitempty" xml:"nameLike,omitempty"`
	// The cursor for querying the next page. Do not specify this parameter in the first call. In subsequent calls, specify the nextToken value returned in the previous call.
	//
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
