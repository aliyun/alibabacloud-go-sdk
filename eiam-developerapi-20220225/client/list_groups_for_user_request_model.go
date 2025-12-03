// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGroupsForUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListGroupsForUserRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListGroupsForUserRequest
	GetNextToken() *string
}

type ListGroupsForUserRequest struct {
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// nextToken
	//
	// example:
	//
	// NTxxx
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListGroupsForUserRequest) String() string {
	return dara.Prettify(s)
}

func (s ListGroupsForUserRequest) GoString() string {
	return s.String()
}

func (s *ListGroupsForUserRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListGroupsForUserRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListGroupsForUserRequest) SetMaxResults(v int32) *ListGroupsForUserRequest {
	s.MaxResults = &v
	return s
}

func (s *ListGroupsForUserRequest) SetNextToken(v string) *ListGroupsForUserRequest {
	s.NextToken = &v
	return s
}

func (s *ListGroupsForUserRequest) Validate() error {
	return dara.Validate(s)
}
