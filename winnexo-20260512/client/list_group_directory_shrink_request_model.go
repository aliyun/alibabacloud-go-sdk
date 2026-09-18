// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGroupDirectoryShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirectoryId(v string) *ListGroupDirectoryShrinkRequest
	GetDirectoryId() *string
	SetGroupId(v string) *ListGroupDirectoryShrinkRequest
	GetGroupId() *string
	SetPage(v int64) *ListGroupDirectoryShrinkRequest
	GetPage() *int64
	SetPageSize(v int64) *ListGroupDirectoryShrinkRequest
	GetPageSize() *int64
	SetSortField(v string) *ListGroupDirectoryShrinkRequest
	GetSortField() *string
	SetSortOrder(v string) *ListGroupDirectoryShrinkRequest
	GetSortOrder() *string
	SetSourceStatus(v string) *ListGroupDirectoryShrinkRequest
	GetSourceStatus() *string
	SetSourceTypesShrink(v string) *ListGroupDirectoryShrinkRequest
	GetSourceTypesShrink() *string
	SetTenantId(v string) *ListGroupDirectoryShrinkRequest
	GetTenantId() *string
}

type ListGroupDirectoryShrinkRequest struct {
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
	SourceTypesShrink *string `json:"sourceTypes,omitempty" xml:"sourceTypes,omitempty"`
	// The tenant ID. This is a common parameter. If not specified, the default tenant of the caller is used.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ListGroupDirectoryShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListGroupDirectoryShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListGroupDirectoryShrinkRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *ListGroupDirectoryShrinkRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *ListGroupDirectoryShrinkRequest) GetPage() *int64 {
	return s.Page
}

func (s *ListGroupDirectoryShrinkRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListGroupDirectoryShrinkRequest) GetSortField() *string {
	return s.SortField
}

func (s *ListGroupDirectoryShrinkRequest) GetSortOrder() *string {
	return s.SortOrder
}

func (s *ListGroupDirectoryShrinkRequest) GetSourceStatus() *string {
	return s.SourceStatus
}

func (s *ListGroupDirectoryShrinkRequest) GetSourceTypesShrink() *string {
	return s.SourceTypesShrink
}

func (s *ListGroupDirectoryShrinkRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListGroupDirectoryShrinkRequest) SetDirectoryId(v string) *ListGroupDirectoryShrinkRequest {
	s.DirectoryId = &v
	return s
}

func (s *ListGroupDirectoryShrinkRequest) SetGroupId(v string) *ListGroupDirectoryShrinkRequest {
	s.GroupId = &v
	return s
}

func (s *ListGroupDirectoryShrinkRequest) SetPage(v int64) *ListGroupDirectoryShrinkRequest {
	s.Page = &v
	return s
}

func (s *ListGroupDirectoryShrinkRequest) SetPageSize(v int64) *ListGroupDirectoryShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListGroupDirectoryShrinkRequest) SetSortField(v string) *ListGroupDirectoryShrinkRequest {
	s.SortField = &v
	return s
}

func (s *ListGroupDirectoryShrinkRequest) SetSortOrder(v string) *ListGroupDirectoryShrinkRequest {
	s.SortOrder = &v
	return s
}

func (s *ListGroupDirectoryShrinkRequest) SetSourceStatus(v string) *ListGroupDirectoryShrinkRequest {
	s.SourceStatus = &v
	return s
}

func (s *ListGroupDirectoryShrinkRequest) SetSourceTypesShrink(v string) *ListGroupDirectoryShrinkRequest {
	s.SourceTypesShrink = &v
	return s
}

func (s *ListGroupDirectoryShrinkRequest) SetTenantId(v string) *ListGroupDirectoryShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *ListGroupDirectoryShrinkRequest) Validate() error {
	return dara.Validate(s)
}
