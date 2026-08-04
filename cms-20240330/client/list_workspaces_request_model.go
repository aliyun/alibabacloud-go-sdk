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
	SetResourceGroupId(v string) *ListWorkspacesRequest
	GetResourceGroupId() *string
	SetTags(v []*ListWorkspacesRequestTags) *ListWorkspacesRequest
	GetTags() []*ListWorkspacesRequestTags
	SetWorkspaceName(v string) *ListWorkspacesRequest
	GetWorkspaceName() *string
	SetWorkspaceNameList(v []*string) *ListWorkspacesRequest
	GetWorkspaceNameList() []*string
}

type ListWorkspacesRequest struct {
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
	Tags []*ListWorkspacesRequestTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
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

func (s *ListWorkspacesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListWorkspacesRequest) GetTags() []*ListWorkspacesRequestTags {
	return s.Tags
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

func (s *ListWorkspacesRequest) SetResourceGroupId(v string) *ListWorkspacesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListWorkspacesRequest) SetTags(v []*ListWorkspacesRequestTags) *ListWorkspacesRequest {
	s.Tags = v
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
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListWorkspacesRequestTags struct {
	// The tag key.
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// The tag value.
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ListWorkspacesRequestTags) String() string {
	return dara.Prettify(s)
}

func (s ListWorkspacesRequestTags) GoString() string {
	return s.String()
}

func (s *ListWorkspacesRequestTags) GetKey() *string {
	return s.Key
}

func (s *ListWorkspacesRequestTags) GetValue() *string {
	return s.Value
}

func (s *ListWorkspacesRequestTags) SetKey(v string) *ListWorkspacesRequestTags {
	s.Key = &v
	return s
}

func (s *ListWorkspacesRequestTags) SetValue(v string) *ListWorkspacesRequestTags {
	s.Value = &v
	return s
}

func (s *ListWorkspacesRequestTags) Validate() error {
	return dara.Validate(s)
}
