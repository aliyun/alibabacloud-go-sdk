// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPersonalDirectoryContentsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirectoryId(v string) *ListPersonalDirectoryContentsShrinkRequest
	GetDirectoryId() *string
	SetOperatingObjectName(v string) *ListPersonalDirectoryContentsShrinkRequest
	GetOperatingObjectName() *string
	SetPage(v int64) *ListPersonalDirectoryContentsShrinkRequest
	GetPage() *int64
	SetPageSize(v int64) *ListPersonalDirectoryContentsShrinkRequest
	GetPageSize() *int64
	SetSortField(v string) *ListPersonalDirectoryContentsShrinkRequest
	GetSortField() *string
	SetSortOrder(v string) *ListPersonalDirectoryContentsShrinkRequest
	GetSortOrder() *string
	SetSourceTypesShrink(v string) *ListPersonalDirectoryContentsShrinkRequest
	GetSourceTypesShrink() *string
	SetTenantId(v string) *ListPersonalDirectoryContentsShrinkRequest
	GetTenantId() *string
}

type ListPersonalDirectoryContentsShrinkRequest struct {
	// The directory ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// exampleDirectoryId
	DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
	// The name of the digital employee.
	//
	// example:
	//
	// string_value
	OperatingObjectName *string `json:"operatingObjectName,omitempty" xml:"operatingObjectName,omitempty"`
	// The page number. Default value: 1. Pages start from page 1.
	//
	// example:
	//
	// 1
	Page *int64 `json:"page,omitempty" xml:"page,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The field by which the results are sorted. Valid values:
	//
	// - event_time: event creation time
	//
	// - event_execute_start_time: event execution time
	//
	// - event_execute_finish_time: event completion time
	//
	// example:
	//
	// name
	SortField *string `json:"sortField,omitempty" xml:"sortField,omitempty"`
	// The sort order. This parameter takes effect when sortBy is specified. Valid values: ASC, DESC (case-insensitive).
	//
	// example:
	//
	// asc
	SortOrder *string `json:"sortOrder,omitempty" xml:"sortOrder,omitempty"`
	// The list of service source types.
	//
	// example:
	//
	// string_value
	SourceTypesShrink *string `json:"sourceTypes,omitempty" xml:"sourceTypes,omitempty"`
	// The tenant ID. This is a common parameter. If not specified, the default tenant of the caller is used.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ListPersonalDirectoryContentsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPersonalDirectoryContentsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListPersonalDirectoryContentsShrinkRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *ListPersonalDirectoryContentsShrinkRequest) GetOperatingObjectName() *string {
	return s.OperatingObjectName
}

func (s *ListPersonalDirectoryContentsShrinkRequest) GetPage() *int64 {
	return s.Page
}

func (s *ListPersonalDirectoryContentsShrinkRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListPersonalDirectoryContentsShrinkRequest) GetSortField() *string {
	return s.SortField
}

func (s *ListPersonalDirectoryContentsShrinkRequest) GetSortOrder() *string {
	return s.SortOrder
}

func (s *ListPersonalDirectoryContentsShrinkRequest) GetSourceTypesShrink() *string {
	return s.SourceTypesShrink
}

func (s *ListPersonalDirectoryContentsShrinkRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListPersonalDirectoryContentsShrinkRequest) SetDirectoryId(v string) *ListPersonalDirectoryContentsShrinkRequest {
	s.DirectoryId = &v
	return s
}

func (s *ListPersonalDirectoryContentsShrinkRequest) SetOperatingObjectName(v string) *ListPersonalDirectoryContentsShrinkRequest {
	s.OperatingObjectName = &v
	return s
}

func (s *ListPersonalDirectoryContentsShrinkRequest) SetPage(v int64) *ListPersonalDirectoryContentsShrinkRequest {
	s.Page = &v
	return s
}

func (s *ListPersonalDirectoryContentsShrinkRequest) SetPageSize(v int64) *ListPersonalDirectoryContentsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListPersonalDirectoryContentsShrinkRequest) SetSortField(v string) *ListPersonalDirectoryContentsShrinkRequest {
	s.SortField = &v
	return s
}

func (s *ListPersonalDirectoryContentsShrinkRequest) SetSortOrder(v string) *ListPersonalDirectoryContentsShrinkRequest {
	s.SortOrder = &v
	return s
}

func (s *ListPersonalDirectoryContentsShrinkRequest) SetSourceTypesShrink(v string) *ListPersonalDirectoryContentsShrinkRequest {
	s.SourceTypesShrink = &v
	return s
}

func (s *ListPersonalDirectoryContentsShrinkRequest) SetTenantId(v string) *ListPersonalDirectoryContentsShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *ListPersonalDirectoryContentsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
