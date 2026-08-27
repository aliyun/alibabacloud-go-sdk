// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTenantDirectoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirectoryId(v string) *ListTenantDirectoryRequest
	GetDirectoryId() *string
	SetPage(v int64) *ListTenantDirectoryRequest
	GetPage() *int64
	SetPageSize(v int64) *ListTenantDirectoryRequest
	GetPageSize() *int64
	SetSortField(v string) *ListTenantDirectoryRequest
	GetSortField() *string
	SetSortOrder(v string) *ListTenantDirectoryRequest
	GetSortOrder() *string
	SetSourceTypes(v string) *ListTenantDirectoryRequest
	GetSourceTypes() *string
	SetTenantId(v string) *ListTenantDirectoryRequest
	GetTenantId() *string
}

type ListTenantDirectoryRequest struct {
	// The folder ID.
	//
	// example:
	//
	// exampleDirectoryId
	DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
	// The current page number.
	//
	// example:
	//
	// 1
	Page *int64 `json:"page,omitempty" xml:"page,omitempty"`
	// The number of entries per page. Default value: 100. Maximum value: 500.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The field by which the results are sorted. Valid values:
	//
	// - event_time: the event creation time.
	//
	// - event_execute_start_time: the event execution time.
	//
	// - event_execute_finish_time: the event completion time.
	//
	// example:
	//
	// name
	SortField *string `json:"sortField,omitempty" xml:"sortField,omitempty"`
	// The sort order. This parameter takes effect only when sortBy is specified. Valid values: ASC, DESC (case-insensitive).
	//
	// example:
	//
	// asc
	SortOrder *string `json:"sortOrder,omitempty" xml:"sortOrder,omitempty"`
	// The list of source types.
	//
	// example:
	//
	// string_value
	SourceTypes *string `json:"sourceTypes,omitempty" xml:"sourceTypes,omitempty"`
	// The tenant ID. This is a common parameter. If not specified, the default tenant of the caller is used.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ListTenantDirectoryRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTenantDirectoryRequest) GoString() string {
	return s.String()
}

func (s *ListTenantDirectoryRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *ListTenantDirectoryRequest) GetPage() *int64 {
	return s.Page
}

func (s *ListTenantDirectoryRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListTenantDirectoryRequest) GetSortField() *string {
	return s.SortField
}

func (s *ListTenantDirectoryRequest) GetSortOrder() *string {
	return s.SortOrder
}

func (s *ListTenantDirectoryRequest) GetSourceTypes() *string {
	return s.SourceTypes
}

func (s *ListTenantDirectoryRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListTenantDirectoryRequest) SetDirectoryId(v string) *ListTenantDirectoryRequest {
	s.DirectoryId = &v
	return s
}

func (s *ListTenantDirectoryRequest) SetPage(v int64) *ListTenantDirectoryRequest {
	s.Page = &v
	return s
}

func (s *ListTenantDirectoryRequest) SetPageSize(v int64) *ListTenantDirectoryRequest {
	s.PageSize = &v
	return s
}

func (s *ListTenantDirectoryRequest) SetSortField(v string) *ListTenantDirectoryRequest {
	s.SortField = &v
	return s
}

func (s *ListTenantDirectoryRequest) SetSortOrder(v string) *ListTenantDirectoryRequest {
	s.SortOrder = &v
	return s
}

func (s *ListTenantDirectoryRequest) SetSourceTypes(v string) *ListTenantDirectoryRequest {
	s.SourceTypes = &v
	return s
}

func (s *ListTenantDirectoryRequest) SetTenantId(v string) *ListTenantDirectoryRequest {
	s.TenantId = &v
	return s
}

func (s *ListTenantDirectoryRequest) Validate() error {
	return dara.Validate(s)
}
