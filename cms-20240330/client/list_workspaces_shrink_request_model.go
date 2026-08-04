// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkspacesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListWorkspacesShrinkRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListWorkspacesShrinkRequest
	GetNextToken() *string
	SetRegion(v string) *ListWorkspacesShrinkRequest
	GetRegion() *string
	SetResourceGroupId(v string) *ListWorkspacesShrinkRequest
	GetResourceGroupId() *string
	SetTagsShrink(v string) *ListWorkspacesShrinkRequest
	GetTagsShrink() *string
	SetWorkspaceName(v string) *ListWorkspacesShrinkRequest
	GetWorkspaceName() *string
	SetWorkspaceNameListShrink(v string) *ListWorkspacesShrinkRequest
	GetWorkspaceNameListShrink() *string
}

type ListWorkspacesShrinkRequest struct {
	// The number of entries per page.
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
	// The pagination token.
	//
	// example:
	//
	// xxxxxxxxxx
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The region.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// cn-heyuan
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-ae******ey
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// The tags.
	TagsShrink *string `json:"tags,omitempty" xml:"tags,omitempty"`
	// The workspace name. Fuzzy match is used.
	//
	// example:
	//
	// workspace-test-001
	WorkspaceName *string `json:"workspaceName,omitempty" xml:"workspaceName,omitempty"`
	// The workspace name. Exact match is used.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// workspace-test-001
	WorkspaceNameListShrink *string `json:"workspaceNameList,omitempty" xml:"workspaceNameList,omitempty"`
}

func (s ListWorkspacesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListWorkspacesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListWorkspacesShrinkRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListWorkspacesShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListWorkspacesShrinkRequest) GetRegion() *string {
	return s.Region
}

func (s *ListWorkspacesShrinkRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListWorkspacesShrinkRequest) GetTagsShrink() *string {
	return s.TagsShrink
}

func (s *ListWorkspacesShrinkRequest) GetWorkspaceName() *string {
	return s.WorkspaceName
}

func (s *ListWorkspacesShrinkRequest) GetWorkspaceNameListShrink() *string {
	return s.WorkspaceNameListShrink
}

func (s *ListWorkspacesShrinkRequest) SetMaxResults(v int32) *ListWorkspacesShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *ListWorkspacesShrinkRequest) SetNextToken(v string) *ListWorkspacesShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListWorkspacesShrinkRequest) SetRegion(v string) *ListWorkspacesShrinkRequest {
	s.Region = &v
	return s
}

func (s *ListWorkspacesShrinkRequest) SetResourceGroupId(v string) *ListWorkspacesShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListWorkspacesShrinkRequest) SetTagsShrink(v string) *ListWorkspacesShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *ListWorkspacesShrinkRequest) SetWorkspaceName(v string) *ListWorkspacesShrinkRequest {
	s.WorkspaceName = &v
	return s
}

func (s *ListWorkspacesShrinkRequest) SetWorkspaceNameListShrink(v string) *ListWorkspacesShrinkRequest {
	s.WorkspaceNameListShrink = &v
	return s
}

func (s *ListWorkspacesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
