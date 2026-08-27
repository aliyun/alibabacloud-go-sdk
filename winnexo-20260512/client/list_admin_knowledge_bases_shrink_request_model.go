// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAdminKnowledgeBasesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirectoryId(v string) *ListAdminKnowledgeBasesShrinkRequest
	GetDirectoryId() *string
	SetKeyword(v string) *ListAdminKnowledgeBasesShrinkRequest
	GetKeyword() *string
	SetPage(v int64) *ListAdminKnowledgeBasesShrinkRequest
	GetPage() *int64
	SetPageSize(v int64) *ListAdminKnowledgeBasesShrinkRequest
	GetPageSize() *int64
	SetSortField(v string) *ListAdminKnowledgeBasesShrinkRequest
	GetSortField() *string
	SetSortOrder(v string) *ListAdminKnowledgeBasesShrinkRequest
	GetSortOrder() *string
	SetSourceTypesShrink(v string) *ListAdminKnowledgeBasesShrinkRequest
	GetSourceTypesShrink() *string
	SetTenantId(v string) *ListAdminKnowledgeBasesShrinkRequest
	GetTenantId() *string
}

type ListAdminKnowledgeBasesShrinkRequest struct {
	// The directory ID.
	//
	// example:
	//
	// exampleDirectoryId
	DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
	// The keyword for fuzzy search on form component data.
	//
	// example:
	//
	// SampleKeyword
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
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
	SourceTypesShrink *string `json:"sourceTypes,omitempty" xml:"sourceTypes,omitempty"`
	// The tenant ID to which the task belongs.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ListAdminKnowledgeBasesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAdminKnowledgeBasesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListAdminKnowledgeBasesShrinkRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *ListAdminKnowledgeBasesShrinkRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListAdminKnowledgeBasesShrinkRequest) GetPage() *int64 {
	return s.Page
}

func (s *ListAdminKnowledgeBasesShrinkRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListAdminKnowledgeBasesShrinkRequest) GetSortField() *string {
	return s.SortField
}

func (s *ListAdminKnowledgeBasesShrinkRequest) GetSortOrder() *string {
	return s.SortOrder
}

func (s *ListAdminKnowledgeBasesShrinkRequest) GetSourceTypesShrink() *string {
	return s.SourceTypesShrink
}

func (s *ListAdminKnowledgeBasesShrinkRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListAdminKnowledgeBasesShrinkRequest) SetDirectoryId(v string) *ListAdminKnowledgeBasesShrinkRequest {
	s.DirectoryId = &v
	return s
}

func (s *ListAdminKnowledgeBasesShrinkRequest) SetKeyword(v string) *ListAdminKnowledgeBasesShrinkRequest {
	s.Keyword = &v
	return s
}

func (s *ListAdminKnowledgeBasesShrinkRequest) SetPage(v int64) *ListAdminKnowledgeBasesShrinkRequest {
	s.Page = &v
	return s
}

func (s *ListAdminKnowledgeBasesShrinkRequest) SetPageSize(v int64) *ListAdminKnowledgeBasesShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListAdminKnowledgeBasesShrinkRequest) SetSortField(v string) *ListAdminKnowledgeBasesShrinkRequest {
	s.SortField = &v
	return s
}

func (s *ListAdminKnowledgeBasesShrinkRequest) SetSortOrder(v string) *ListAdminKnowledgeBasesShrinkRequest {
	s.SortOrder = &v
	return s
}

func (s *ListAdminKnowledgeBasesShrinkRequest) SetSourceTypesShrink(v string) *ListAdminKnowledgeBasesShrinkRequest {
	s.SourceTypesShrink = &v
	return s
}

func (s *ListAdminKnowledgeBasesShrinkRequest) SetTenantId(v string) *ListAdminKnowledgeBasesShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *ListAdminKnowledgeBasesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
