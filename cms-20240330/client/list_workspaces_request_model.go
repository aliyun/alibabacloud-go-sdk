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
	SetRegion(v string) *ListWorkspacesRequest
	GetRegion() *string
	SetWorkspaceName(v string) *ListWorkspacesRequest
	GetWorkspaceName() *string
	SetWorkspaceNameList(v []*string) *ListWorkspacesRequest
	GetWorkspaceNameList() []*string
}

type ListWorkspacesRequest struct {
	// Page size
	//
	// Default value:
	//
	// 	50
	//
	// Maximum value:
	//
	// 	50
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// Pagination Token
	//
	// example:
	//
	// xxxxxxxxxx
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// Region
	//
	// if can be null:
	// true
	//
	// example:
	//
	// cn-heyuan
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// Workspace name, fuzzy search
	//
	// example:
	//
	// workspace-test-001
	WorkspaceName *string `json:"workspaceName,omitempty" xml:"workspaceName,omitempty"`
	// Workspace name, exact match
	//
	// if can be null:
	// true
	//
	// example:
	//
	// workspace-test-001
	WorkspaceNameList []*string `json:"workspaceNameList,omitempty" xml:"workspaceNameList,omitempty" type:"Repeated"`
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

func (s *ListWorkspacesRequest) GetRegion() *string {
	return s.Region
}

func (s *ListWorkspacesRequest) GetWorkspaceName() *string {
	return s.WorkspaceName
}

func (s *ListWorkspacesRequest) GetWorkspaceNameList() []*string {
	return s.WorkspaceNameList
}

func (s *ListWorkspacesRequest) SetMaxResults(v int32) *ListWorkspacesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListWorkspacesRequest) SetNextToken(v string) *ListWorkspacesRequest {
	s.NextToken = &v
	return s
}

func (s *ListWorkspacesRequest) SetRegion(v string) *ListWorkspacesRequest {
	s.Region = &v
	return s
}

func (s *ListWorkspacesRequest) SetWorkspaceName(v string) *ListWorkspacesRequest {
	s.WorkspaceName = &v
	return s
}

func (s *ListWorkspacesRequest) SetWorkspaceNameList(v []*string) *ListWorkspacesRequest {
	s.WorkspaceNameList = v
	return s
}

func (s *ListWorkspacesRequest) Validate() error {
	return dara.Validate(s)
}
