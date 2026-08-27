// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPersonalDirectoryContentsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirectoryId(v string) *ListPersonalDirectoryContentsRequest
	GetDirectoryId() *string
	SetOperatingObjectName(v string) *ListPersonalDirectoryContentsRequest
	GetOperatingObjectName() *string
	SetPage(v int64) *ListPersonalDirectoryContentsRequest
	GetPage() *int64
	SetPageSize(v int64) *ListPersonalDirectoryContentsRequest
	GetPageSize() *int64
	SetSortField(v string) *ListPersonalDirectoryContentsRequest
	GetSortField() *string
	SetSortOrder(v string) *ListPersonalDirectoryContentsRequest
	GetSortOrder() *string
	SetSourceTypes(v []*string) *ListPersonalDirectoryContentsRequest
	GetSourceTypes() []*string
	SetTenantId(v string) *ListPersonalDirectoryContentsRequest
	GetTenantId() *string
}

type ListPersonalDirectoryContentsRequest struct {
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
	SourceTypes []*string `json:"sourceTypes,omitempty" xml:"sourceTypes,omitempty" type:"Repeated"`
	// The tenant ID. This is a common parameter. If not specified, the default tenant of the caller is used.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ListPersonalDirectoryContentsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPersonalDirectoryContentsRequest) GoString() string {
	return s.String()
}

func (s *ListPersonalDirectoryContentsRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *ListPersonalDirectoryContentsRequest) GetOperatingObjectName() *string {
	return s.OperatingObjectName
}

func (s *ListPersonalDirectoryContentsRequest) GetPage() *int64 {
	return s.Page
}

func (s *ListPersonalDirectoryContentsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListPersonalDirectoryContentsRequest) GetSortField() *string {
	return s.SortField
}

func (s *ListPersonalDirectoryContentsRequest) GetSortOrder() *string {
	return s.SortOrder
}

func (s *ListPersonalDirectoryContentsRequest) GetSourceTypes() []*string {
	return s.SourceTypes
}

func (s *ListPersonalDirectoryContentsRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListPersonalDirectoryContentsRequest) SetDirectoryId(v string) *ListPersonalDirectoryContentsRequest {
	s.DirectoryId = &v
	return s
}

func (s *ListPersonalDirectoryContentsRequest) SetOperatingObjectName(v string) *ListPersonalDirectoryContentsRequest {
	s.OperatingObjectName = &v
	return s
}

func (s *ListPersonalDirectoryContentsRequest) SetPage(v int64) *ListPersonalDirectoryContentsRequest {
	s.Page = &v
	return s
}

func (s *ListPersonalDirectoryContentsRequest) SetPageSize(v int64) *ListPersonalDirectoryContentsRequest {
	s.PageSize = &v
	return s
}

func (s *ListPersonalDirectoryContentsRequest) SetSortField(v string) *ListPersonalDirectoryContentsRequest {
	s.SortField = &v
	return s
}

func (s *ListPersonalDirectoryContentsRequest) SetSortOrder(v string) *ListPersonalDirectoryContentsRequest {
	s.SortOrder = &v
	return s
}

func (s *ListPersonalDirectoryContentsRequest) SetSourceTypes(v []*string) *ListPersonalDirectoryContentsRequest {
	s.SourceTypes = v
	return s
}

func (s *ListPersonalDirectoryContentsRequest) SetTenantId(v string) *ListPersonalDirectoryContentsRequest {
	s.TenantId = &v
	return s
}

func (s *ListPersonalDirectoryContentsRequest) Validate() error {
	return dara.Validate(s)
}
