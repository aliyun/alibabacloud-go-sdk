// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkspacesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFields(v string) *ListWorkspacesRequest
	GetFields() *string
	SetModuleList(v string) *ListWorkspacesRequest
	GetModuleList() *string
	SetOption(v string) *ListWorkspacesRequest
	GetOption() *string
	SetOrder(v string) *ListWorkspacesRequest
	GetOrder() *string
	SetPageNumber(v int64) *ListWorkspacesRequest
	GetPageNumber() *int64
	SetPageSize(v int32) *ListWorkspacesRequest
	GetPageSize() *int32
	SetResourceGroupId(v string) *ListWorkspacesRequest
	GetResourceGroupId() *string
	SetSortBy(v string) *ListWorkspacesRequest
	GetSortBy() *string
	SetStatus(v string) *ListWorkspacesRequest
	GetStatus() *string
	SetUserId(v string) *ListWorkspacesRequest
	GetUserId() *string
	SetVerbose(v bool) *ListWorkspacesRequest
	GetVerbose() *bool
	SetWorkspaceIds(v string) *ListWorkspacesRequest
	GetWorkspaceIds() *string
	SetWorkspaceName(v string) *ListWorkspacesRequest
	GetWorkspaceName() *string
}

type ListWorkspacesRequest struct {
	// The list of returned fields of workspace details. Used to limit the fields in the returned results. Separate multiple fields with commas (,). Currently, only Id is supported, which is the workspace ID.
	//
	// example:
	//
	// Id
	Fields *string `json:"Fields,omitempty" xml:"Fields,omitempty"`
	// The modules, separated by commas (,). Default value: PAI.
	//
	// example:
	//
	// PAI
	ModuleList *string `json:"ModuleList,omitempty" xml:"ModuleList,omitempty"`
	// The query options. Valid values:
	//
	// 	- GetWorkspaces (default): Obtains a list of Workspaces.
	//
	// 	- GetResourceLimits: Obtains a list of ResourceLimits.
	//
	// example:
	//
	// GetWorkspaces
	Option *string `json:"Option,omitempty" xml:"Option,omitempty"`
	// The order of results (ascending or descending). Valid values:
	//
	// 	- ASC: ascending order. This is the default value.
	//
	// 	- DESC: descending order.
	//
	// example:
	//
	// ASC
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The page number of the workspace list. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Default value: 20.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The resource group ID. To obtain the ID of a resource group, see [View basic information of a resource group](https://help.aliyun.com/document_detail/151181.html).
	//
	// example:
	//
	// rg-acfmwp7rky****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Specifies how to sort the results. Default value: GmtCreateTime. Valid values:
	//
	// 	- GmtCreateTime: Sort by the time when created.
	//
	// 	- GmtModifiedTime: Sort by the time when modified.
	//
	// example:
	//
	// GmtCreateTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The workspace status. Valid values:
	//
	// 	- ENABLED
	//
	// 	- INITIALIZING
	//
	// 	- FAILURE
	//
	// 	- DISABLED
	//
	// 	- FROZEN
	//
	// 	- UPDATING
	//
	// example:
	//
	// ENABLED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// Specifies whether to display workspace details. Valid values:
	//
	// 	- false (default)
	//
	// 	- true
	//
	// example:
	//
	// true
	Verbose *bool `json:"Verbose,omitempty" xml:"Verbose,omitempty"`
	// The workspace IDs. Separate multiple IDs by commas (,).
	//
	// example:
	//
	// 123,234
	WorkspaceIds *string `json:"WorkspaceIds,omitempty" xml:"WorkspaceIds,omitempty"`
	// The name of the workspace.
	//
	// example:
	//
	// abc
	WorkspaceName *string `json:"WorkspaceName,omitempty" xml:"WorkspaceName,omitempty"`
}

func (s ListWorkspacesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListWorkspacesRequest) GoString() string {
	return s.String()
}

func (s *ListWorkspacesRequest) GetFields() *string {
	return s.Fields
}

func (s *ListWorkspacesRequest) GetModuleList() *string {
	return s.ModuleList
}

func (s *ListWorkspacesRequest) GetOption() *string {
	return s.Option
}

func (s *ListWorkspacesRequest) GetOrder() *string {
	return s.Order
}

func (s *ListWorkspacesRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListWorkspacesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListWorkspacesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListWorkspacesRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListWorkspacesRequest) GetStatus() *string {
	return s.Status
}

func (s *ListWorkspacesRequest) GetUserId() *string {
	return s.UserId
}

func (s *ListWorkspacesRequest) GetVerbose() *bool {
	return s.Verbose
}

func (s *ListWorkspacesRequest) GetWorkspaceIds() *string {
	return s.WorkspaceIds
}

func (s *ListWorkspacesRequest) GetWorkspaceName() *string {
	return s.WorkspaceName
}

func (s *ListWorkspacesRequest) SetFields(v string) *ListWorkspacesRequest {
	s.Fields = &v
	return s
}

func (s *ListWorkspacesRequest) SetModuleList(v string) *ListWorkspacesRequest {
	s.ModuleList = &v
	return s
}

func (s *ListWorkspacesRequest) SetOption(v string) *ListWorkspacesRequest {
	s.Option = &v
	return s
}

func (s *ListWorkspacesRequest) SetOrder(v string) *ListWorkspacesRequest {
	s.Order = &v
	return s
}

func (s *ListWorkspacesRequest) SetPageNumber(v int64) *ListWorkspacesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListWorkspacesRequest) SetPageSize(v int32) *ListWorkspacesRequest {
	s.PageSize = &v
	return s
}

func (s *ListWorkspacesRequest) SetResourceGroupId(v string) *ListWorkspacesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListWorkspacesRequest) SetSortBy(v string) *ListWorkspacesRequest {
	s.SortBy = &v
	return s
}

func (s *ListWorkspacesRequest) SetStatus(v string) *ListWorkspacesRequest {
	s.Status = &v
	return s
}

func (s *ListWorkspacesRequest) SetUserId(v string) *ListWorkspacesRequest {
	s.UserId = &v
	return s
}

func (s *ListWorkspacesRequest) SetVerbose(v bool) *ListWorkspacesRequest {
	s.Verbose = &v
	return s
}

func (s *ListWorkspacesRequest) SetWorkspaceIds(v string) *ListWorkspacesRequest {
	s.WorkspaceIds = &v
	return s
}

func (s *ListWorkspacesRequest) SetWorkspaceName(v string) *ListWorkspacesRequest {
	s.WorkspaceName = &v
	return s
}

func (s *ListWorkspacesRequest) Validate() error {
	return dara.Validate(s)
}
