// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProjectsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListProjectsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListProjectsRequest
	GetNextToken() *string
	SetSearchVal(v string) *ListProjectsRequest
	GetSearchVal() *string
	SetWorkspaceId(v string) *ListProjectsRequest
	GetWorkspaceId() *string
}

type ListProjectsRequest struct {
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// 123abc***
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// test
	SearchVal *string `json:"searchVal,omitempty" xml:"searchVal,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 12345
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s ListProjectsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListProjectsRequest) GoString() string {
	return s.String()
}

func (s *ListProjectsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListProjectsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListProjectsRequest) GetSearchVal() *string {
	return s.SearchVal
}

func (s *ListProjectsRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListProjectsRequest) SetMaxResults(v int32) *ListProjectsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListProjectsRequest) SetNextToken(v string) *ListProjectsRequest {
	s.NextToken = &v
	return s
}

func (s *ListProjectsRequest) SetSearchVal(v string) *ListProjectsRequest {
	s.SearchVal = &v
	return s
}

func (s *ListProjectsRequest) SetWorkspaceId(v string) *ListProjectsRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListProjectsRequest) Validate() error {
	return dara.Validate(s)
}
