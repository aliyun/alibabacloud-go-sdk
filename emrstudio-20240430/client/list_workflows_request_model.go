// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkflowsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListWorkflowsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListWorkflowsRequest
	GetNextToken() *string
	SetSearchVal(v string) *ListWorkflowsRequest
	GetSearchVal() *string
	SetWorkspaceId(v string) *ListWorkflowsRequest
	GetWorkspaceId() *string
}

type ListWorkflowsRequest struct {
	// example:
	//
	// 20
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

func (s ListWorkflowsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListWorkflowsRequest) GoString() string {
	return s.String()
}

func (s *ListWorkflowsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListWorkflowsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListWorkflowsRequest) GetSearchVal() *string {
	return s.SearchVal
}

func (s *ListWorkflowsRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListWorkflowsRequest) SetMaxResults(v int32) *ListWorkflowsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListWorkflowsRequest) SetNextToken(v string) *ListWorkflowsRequest {
	s.NextToken = &v
	return s
}

func (s *ListWorkflowsRequest) SetSearchVal(v string) *ListWorkflowsRequest {
	s.SearchVal = &v
	return s
}

func (s *ListWorkflowsRequest) SetWorkspaceId(v string) *ListWorkflowsRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListWorkflowsRequest) Validate() error {
	return dara.Validate(s)
}
