// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupNameStartWith(v string) *ListGroupsRequest
	GetGroupNameStartWith() *string
	SetMaxResults(v int32) *ListGroupsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListGroupsRequest
	GetNextToken() *string
}

type ListGroupsRequest struct {
	// The prefix of the group name.
	//
	// example:
	//
	// group_xxx
	GroupNameStartWith *string `json:"groupNameStartWith,omitempty" xml:"groupNameStartWith,omitempty"`
	// The number of entries per page. Default value: 20.
	//
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

func (s ListGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListGroupsRequest) GetGroupNameStartWith() *string {
	return s.GroupNameStartWith
}

func (s *ListGroupsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListGroupsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListGroupsRequest) SetGroupNameStartWith(v string) *ListGroupsRequest {
	s.GroupNameStartWith = &v
	return s
}

func (s *ListGroupsRequest) SetMaxResults(v int32) *ListGroupsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListGroupsRequest) SetNextToken(v string) *ListGroupsRequest {
	s.NextToken = &v
	return s
}

func (s *ListGroupsRequest) Validate() error {
	return dara.Validate(s)
}
