// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGroupDirectoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirectoryId(v string) *ListGroupDirectoryRequest
	GetDirectoryId() *string
	SetGroupId(v string) *ListGroupDirectoryRequest
	GetGroupId() *string
	SetPage(v int64) *ListGroupDirectoryRequest
	GetPage() *int64
	SetPageSize(v int64) *ListGroupDirectoryRequest
	GetPageSize() *int64
	SetSortField(v string) *ListGroupDirectoryRequest
	GetSortField() *string
	SetSortOrder(v string) *ListGroupDirectoryRequest
	GetSortOrder() *string
	SetSourceStatus(v string) *ListGroupDirectoryRequest
	GetSourceStatus() *string
	SetSourceTypes(v []*string) *ListGroupDirectoryRequest
	GetSourceTypes() []*string
	SetTenantId(v string) *ListGroupDirectoryRequest
	GetTenantId() *string
}

type ListGroupDirectoryRequest struct {
	// The ID of a visible directory within the space. If omitted or set to root, the internal root is queried. On the first query, the existing service initialization for the root directory is used.
	//
	// example:
	//
	// dir_example
	DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
	// The collaboration space ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// group_example
	GroupId *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// The page number, starting from 1.
	//
	// example:
	//
	// 1
	Page *int64 `json:"page,omitempty" xml:"page,omitempty"`
	// The number of entries per page. Valid values: 1 to 100. Default value: 20.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The sort field within the group. Valid values: name, gmt_create, and gmt_modified. Directories are listed first.
	//
	// example:
	//
	// name
	SortField *string `json:"sortField,omitempty" xml:"sortField,omitempty"`
	// The sort order within the group. Valid values: asc and desc. Directories are always listed first.
	//
	// example:
	//
	// asc
	SortOrder *string `json:"sortOrder,omitempty" xml:"sortOrder,omitempty"`
	// The resource status filter. Physical subdirectories are retained. Immediate reference directories are not returned when a status filter is set. This follows the existing behavior.
	//
	// example:
	//
	// READY
	SourceStatus *string `json:"sourceStatus,omitempty" xml:"sourceStatus,omitempty"`
	// The array of resource types. If values are specified, only resources are returned. If the array is empty or omitted, no type-based filtering is applied, and the existing resource type filtering logic is used.
	//
	// example:
	//
	// ["TEXT"]
	SourceTypes []*string `json:"sourceTypes,omitempty" xml:"sourceTypes,omitempty" type:"Repeated"`
	// The tenant ID. This is a common parameter. If not specified, the default tenant of the caller is used.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ListGroupDirectoryRequest) String() string {
	return dara.Prettify(s)
}

func (s ListGroupDirectoryRequest) GoString() string {
	return s.String()
}

func (s *ListGroupDirectoryRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *ListGroupDirectoryRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *ListGroupDirectoryRequest) GetPage() *int64 {
	return s.Page
}

func (s *ListGroupDirectoryRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListGroupDirectoryRequest) GetSortField() *string {
	return s.SortField
}

func (s *ListGroupDirectoryRequest) GetSortOrder() *string {
	return s.SortOrder
}

func (s *ListGroupDirectoryRequest) GetSourceStatus() *string {
	return s.SourceStatus
}

func (s *ListGroupDirectoryRequest) GetSourceTypes() []*string {
	return s.SourceTypes
}

func (s *ListGroupDirectoryRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListGroupDirectoryRequest) SetDirectoryId(v string) *ListGroupDirectoryRequest {
	s.DirectoryId = &v
	return s
}

func (s *ListGroupDirectoryRequest) SetGroupId(v string) *ListGroupDirectoryRequest {
	s.GroupId = &v
	return s
}

func (s *ListGroupDirectoryRequest) SetPage(v int64) *ListGroupDirectoryRequest {
	s.Page = &v
	return s
}

func (s *ListGroupDirectoryRequest) SetPageSize(v int64) *ListGroupDirectoryRequest {
	s.PageSize = &v
	return s
}

func (s *ListGroupDirectoryRequest) SetSortField(v string) *ListGroupDirectoryRequest {
	s.SortField = &v
	return s
}

func (s *ListGroupDirectoryRequest) SetSortOrder(v string) *ListGroupDirectoryRequest {
	s.SortOrder = &v
	return s
}

func (s *ListGroupDirectoryRequest) SetSourceStatus(v string) *ListGroupDirectoryRequest {
	s.SourceStatus = &v
	return s
}

func (s *ListGroupDirectoryRequest) SetSourceTypes(v []*string) *ListGroupDirectoryRequest {
	s.SourceTypes = v
	return s
}

func (s *ListGroupDirectoryRequest) SetTenantId(v string) *ListGroupDirectoryRequest {
	s.TenantId = &v
	return s
}

func (s *ListGroupDirectoryRequest) Validate() error {
	return dara.Validate(s)
}
