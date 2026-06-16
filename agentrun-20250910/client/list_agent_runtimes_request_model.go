// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgentRuntimesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentRuntimeName(v string) *ListAgentRuntimesRequest
	GetAgentRuntimeName() *string
	SetDiscoveryResourceGroupId(v string) *ListAgentRuntimesRequest
	GetDiscoveryResourceGroupId() *string
	SetPageNumber(v int32) *ListAgentRuntimesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListAgentRuntimesRequest
	GetPageSize() *int32
	SetResourceGroupId(v string) *ListAgentRuntimesRequest
	GetResourceGroupId() *string
	SetSearchMode(v string) *ListAgentRuntimesRequest
	GetSearchMode() *string
	SetStatus(v string) *ListAgentRuntimesRequest
	GetStatus() *string
	SetSystemTags(v string) *ListAgentRuntimesRequest
	GetSystemTags() *string
	SetWorkspaceId(v string) *ListAgentRuntimesRequest
	GetWorkspaceId() *string
	SetWorkspaceIds(v string) *ListAgentRuntimesRequest
	GetWorkspaceIds() *string
}

type ListAgentRuntimesRequest struct {
	// Filters the results by agent runtime name.
	//
	// example:
	//
	// my-runtime
	AgentRuntimeName *string `json:"agentRuntimeName,omitempty" xml:"agentRuntimeName,omitempty"`
	// Deprecated
	//
	// The service discovery resource group ID.
	//
	// example:
	//
	// rg-123456
	DiscoveryResourceGroupId *string `json:"discoveryResourceGroupId,omitempty" xml:"discoveryResourceGroupId,omitempty"`
	// The page number to return.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries to return per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// Deprecated
	//
	// The ID of the resource group. This parameter is deprecated.
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// The search mode.
	//
	// example:
	//
	// fuzzy
	SearchMode *string `json:"searchMode,omitempty" xml:"searchMode,omitempty"`
	// Filters the results by status.
	//
	// example:
	//
	// READY,CREATING
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// Filters the results by system tags. Separate multiple tags with commas. This parameter supports only exact matches.
	//
	// example:
	//
	// acs:ecs:tag1,acs:ecs:tag2
	SystemTags *string `json:"systemTags,omitempty" xml:"systemTags,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// ws-1234567890abcdef
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
	// A comma-separated string of workspace IDs.
	WorkspaceIds *string `json:"workspaceIds,omitempty" xml:"workspaceIds,omitempty"`
}

func (s ListAgentRuntimesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAgentRuntimesRequest) GoString() string {
	return s.String()
}

func (s *ListAgentRuntimesRequest) GetAgentRuntimeName() *string {
	return s.AgentRuntimeName
}

func (s *ListAgentRuntimesRequest) GetDiscoveryResourceGroupId() *string {
	return s.DiscoveryResourceGroupId
}

func (s *ListAgentRuntimesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListAgentRuntimesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAgentRuntimesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListAgentRuntimesRequest) GetSearchMode() *string {
	return s.SearchMode
}

func (s *ListAgentRuntimesRequest) GetStatus() *string {
	return s.Status
}

func (s *ListAgentRuntimesRequest) GetSystemTags() *string {
	return s.SystemTags
}

func (s *ListAgentRuntimesRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListAgentRuntimesRequest) GetWorkspaceIds() *string {
	return s.WorkspaceIds
}

func (s *ListAgentRuntimesRequest) SetAgentRuntimeName(v string) *ListAgentRuntimesRequest {
	s.AgentRuntimeName = &v
	return s
}

func (s *ListAgentRuntimesRequest) SetDiscoveryResourceGroupId(v string) *ListAgentRuntimesRequest {
	s.DiscoveryResourceGroupId = &v
	return s
}

func (s *ListAgentRuntimesRequest) SetPageNumber(v int32) *ListAgentRuntimesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAgentRuntimesRequest) SetPageSize(v int32) *ListAgentRuntimesRequest {
	s.PageSize = &v
	return s
}

func (s *ListAgentRuntimesRequest) SetResourceGroupId(v string) *ListAgentRuntimesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListAgentRuntimesRequest) SetSearchMode(v string) *ListAgentRuntimesRequest {
	s.SearchMode = &v
	return s
}

func (s *ListAgentRuntimesRequest) SetStatus(v string) *ListAgentRuntimesRequest {
	s.Status = &v
	return s
}

func (s *ListAgentRuntimesRequest) SetSystemTags(v string) *ListAgentRuntimesRequest {
	s.SystemTags = &v
	return s
}

func (s *ListAgentRuntimesRequest) SetWorkspaceId(v string) *ListAgentRuntimesRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListAgentRuntimesRequest) SetWorkspaceIds(v string) *ListAgentRuntimesRequest {
	s.WorkspaceIds = &v
	return s
}

func (s *ListAgentRuntimesRequest) Validate() error {
	return dara.Validate(s)
}
