// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserVisibleKnowledgeBaseContentsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirectoryId(v string) *ListUserVisibleKnowledgeBaseContentsRequest
	GetDirectoryId() *string
	SetKeyword(v string) *ListUserVisibleKnowledgeBaseContentsRequest
	GetKeyword() *string
	SetPage(v int64) *ListUserVisibleKnowledgeBaseContentsRequest
	GetPage() *int64
	SetPageSize(v int64) *ListUserVisibleKnowledgeBaseContentsRequest
	GetPageSize() *int64
	SetSortField(v string) *ListUserVisibleKnowledgeBaseContentsRequest
	GetSortField() *string
	SetSortOrder(v string) *ListUserVisibleKnowledgeBaseContentsRequest
	GetSortOrder() *string
	SetSourceTypes(v string) *ListUserVisibleKnowledgeBaseContentsRequest
	GetSourceTypes() *string
	SetTenantId(v string) *ListUserVisibleKnowledgeBaseContentsRequest
	GetTenantId() *string
}

type ListUserVisibleKnowledgeBaseContentsRequest struct {
	// The directory ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// dir_example
	DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
	// The search keyword. Supports fuzzy match on version names.
	//
	// example:
	//
	// Product description
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	Page *int64 `json:"page,omitempty" xml:"page,omitempty"`
	// The page size.
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
	// FILE,WEB_PAGE
	SourceTypes *string `json:"sourceTypes,omitempty" xml:"sourceTypes,omitempty"`
	// The tenant ID. This is a common parameter. Pass it explicitly through --tenant-id in winnexo-cli.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ListUserVisibleKnowledgeBaseContentsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListUserVisibleKnowledgeBaseContentsRequest) GoString() string {
	return s.String()
}

func (s *ListUserVisibleKnowledgeBaseContentsRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *ListUserVisibleKnowledgeBaseContentsRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListUserVisibleKnowledgeBaseContentsRequest) GetPage() *int64 {
	return s.Page
}

func (s *ListUserVisibleKnowledgeBaseContentsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListUserVisibleKnowledgeBaseContentsRequest) GetSortField() *string {
	return s.SortField
}

func (s *ListUserVisibleKnowledgeBaseContentsRequest) GetSortOrder() *string {
	return s.SortOrder
}

func (s *ListUserVisibleKnowledgeBaseContentsRequest) GetSourceTypes() *string {
	return s.SourceTypes
}

func (s *ListUserVisibleKnowledgeBaseContentsRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListUserVisibleKnowledgeBaseContentsRequest) SetDirectoryId(v string) *ListUserVisibleKnowledgeBaseContentsRequest {
	s.DirectoryId = &v
	return s
}

func (s *ListUserVisibleKnowledgeBaseContentsRequest) SetKeyword(v string) *ListUserVisibleKnowledgeBaseContentsRequest {
	s.Keyword = &v
	return s
}

func (s *ListUserVisibleKnowledgeBaseContentsRequest) SetPage(v int64) *ListUserVisibleKnowledgeBaseContentsRequest {
	s.Page = &v
	return s
}

func (s *ListUserVisibleKnowledgeBaseContentsRequest) SetPageSize(v int64) *ListUserVisibleKnowledgeBaseContentsRequest {
	s.PageSize = &v
	return s
}

func (s *ListUserVisibleKnowledgeBaseContentsRequest) SetSortField(v string) *ListUserVisibleKnowledgeBaseContentsRequest {
	s.SortField = &v
	return s
}

func (s *ListUserVisibleKnowledgeBaseContentsRequest) SetSortOrder(v string) *ListUserVisibleKnowledgeBaseContentsRequest {
	s.SortOrder = &v
	return s
}

func (s *ListUserVisibleKnowledgeBaseContentsRequest) SetSourceTypes(v string) *ListUserVisibleKnowledgeBaseContentsRequest {
	s.SourceTypes = &v
	return s
}

func (s *ListUserVisibleKnowledgeBaseContentsRequest) SetTenantId(v string) *ListUserVisibleKnowledgeBaseContentsRequest {
	s.TenantId = &v
	return s
}

func (s *ListUserVisibleKnowledgeBaseContentsRequest) Validate() error {
	return dara.Validate(s)
}
