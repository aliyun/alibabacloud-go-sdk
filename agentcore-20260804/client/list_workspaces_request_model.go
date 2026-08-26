// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkspacesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListWorkspacesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListWorkspacesRequest
	GetNextToken() *string
	SetSkip(v int32) *ListWorkspacesRequest
	GetSkip() *int32
}

type ListWorkspacesRequest struct {
	// The maximum number of records to return per page.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The pagination token used to retrieve the next page.
	//
	// example:
	//
	// d29ya3NwYWNlLW9mZnNldDoyMA
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The number of workspaces to skip.
	//
	// example:
	//
	// 0
	Skip *int32 `json:"skip,omitempty" xml:"skip,omitempty"`
}

func (s ListWorkspacesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListWorkspacesRequest) GoString() string {
	return s.String()
}

func (s *ListWorkspacesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListWorkspacesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListWorkspacesRequest) GetSkip() *int32 {
	return s.Skip
}

func (s *ListWorkspacesRequest) SetMaxResults(v int32) *ListWorkspacesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListWorkspacesRequest) SetNextToken(v string) *ListWorkspacesRequest {
	s.NextToken = &v
	return s
}

func (s *ListWorkspacesRequest) SetSkip(v int32) *ListWorkspacesRequest {
	s.Skip = &v
	return s
}

func (s *ListWorkspacesRequest) Validate() error {
	return dara.Validate(s)
}
