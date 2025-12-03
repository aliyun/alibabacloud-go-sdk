// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUsersForGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListUsersForGroupRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListUsersForGroupRequest
	GetNextToken() *string
}

type ListUsersForGroupRequest struct {
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

func (s ListUsersForGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s ListUsersForGroupRequest) GoString() string {
	return s.String()
}

func (s *ListUsersForGroupRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListUsersForGroupRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListUsersForGroupRequest) SetMaxResults(v int32) *ListUsersForGroupRequest {
	s.MaxResults = &v
	return s
}

func (s *ListUsersForGroupRequest) SetNextToken(v string) *ListUsersForGroupRequest {
	s.NextToken = &v
	return s
}

func (s *ListUsersForGroupRequest) Validate() error {
	return dara.Validate(s)
}
