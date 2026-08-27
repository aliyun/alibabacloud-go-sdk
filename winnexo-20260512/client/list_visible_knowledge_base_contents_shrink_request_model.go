// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVisibleKnowledgeBaseContentsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirectoryId(v string) *ListVisibleKnowledgeBaseContentsShrinkRequest
	GetDirectoryId() *string
	SetOperatingObjectName(v string) *ListVisibleKnowledgeBaseContentsShrinkRequest
	GetOperatingObjectName() *string
	SetPage(v int64) *ListVisibleKnowledgeBaseContentsShrinkRequest
	GetPage() *int64
	SetPageSize(v int64) *ListVisibleKnowledgeBaseContentsShrinkRequest
	GetPageSize() *int64
	SetSortField(v string) *ListVisibleKnowledgeBaseContentsShrinkRequest
	GetSortField() *string
	SetSortOrder(v string) *ListVisibleKnowledgeBaseContentsShrinkRequest
	GetSortOrder() *string
	SetSourceTypesShrink(v string) *ListVisibleKnowledgeBaseContentsShrinkRequest
	GetSourceTypesShrink() *string
	SetTenantId(v string) *ListVisibleKnowledgeBaseContentsShrinkRequest
	GetTenantId() *string
}

type ListVisibleKnowledgeBaseContentsShrinkRequest struct {
	// The directory ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// exampleDirectoryId
	DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
	// The name of the digital employee (operating object name).
	//
	// This parameter is required.
	//
	// example:
	//
	// string_value
	OperatingObjectName *string `json:"operatingObjectName,omitempty" xml:"operatingObjectName,omitempty"`
	// The page number of the results to return. Default value: 1. Minimum value: 1. Maximum value: 200.
	//
	// example:
	//
	// 1
	Page *int64 `json:"page,omitempty" xml:"page,omitempty"`
	// The page size. Maximum value: 100.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The field by which to sort the results. Valid values:
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
	// The sort order. This parameter takes effect only when sortBy is specified. Valid values: ASC, DESC (case-insensitive).
	//
	// example:
	//
	// desc
	SortOrder *string `json:"sortOrder,omitempty" xml:"sortOrder,omitempty"`
	// The list of service source types.
	//
	// example:
	//
	// string_value
	SourceTypesShrink *string `json:"sourceTypes,omitempty" xml:"sourceTypes,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// PiPklI1iSRTm6VFFqlY9VzbgiEiE
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ListVisibleKnowledgeBaseContentsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListVisibleKnowledgeBaseContentsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListVisibleKnowledgeBaseContentsShrinkRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *ListVisibleKnowledgeBaseContentsShrinkRequest) GetOperatingObjectName() *string {
	return s.OperatingObjectName
}

func (s *ListVisibleKnowledgeBaseContentsShrinkRequest) GetPage() *int64 {
	return s.Page
}

func (s *ListVisibleKnowledgeBaseContentsShrinkRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListVisibleKnowledgeBaseContentsShrinkRequest) GetSortField() *string {
	return s.SortField
}

func (s *ListVisibleKnowledgeBaseContentsShrinkRequest) GetSortOrder() *string {
	return s.SortOrder
}

func (s *ListVisibleKnowledgeBaseContentsShrinkRequest) GetSourceTypesShrink() *string {
	return s.SourceTypesShrink
}

func (s *ListVisibleKnowledgeBaseContentsShrinkRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListVisibleKnowledgeBaseContentsShrinkRequest) SetDirectoryId(v string) *ListVisibleKnowledgeBaseContentsShrinkRequest {
	s.DirectoryId = &v
	return s
}

func (s *ListVisibleKnowledgeBaseContentsShrinkRequest) SetOperatingObjectName(v string) *ListVisibleKnowledgeBaseContentsShrinkRequest {
	s.OperatingObjectName = &v
	return s
}

func (s *ListVisibleKnowledgeBaseContentsShrinkRequest) SetPage(v int64) *ListVisibleKnowledgeBaseContentsShrinkRequest {
	s.Page = &v
	return s
}

func (s *ListVisibleKnowledgeBaseContentsShrinkRequest) SetPageSize(v int64) *ListVisibleKnowledgeBaseContentsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListVisibleKnowledgeBaseContentsShrinkRequest) SetSortField(v string) *ListVisibleKnowledgeBaseContentsShrinkRequest {
	s.SortField = &v
	return s
}

func (s *ListVisibleKnowledgeBaseContentsShrinkRequest) SetSortOrder(v string) *ListVisibleKnowledgeBaseContentsShrinkRequest {
	s.SortOrder = &v
	return s
}

func (s *ListVisibleKnowledgeBaseContentsShrinkRequest) SetSourceTypesShrink(v string) *ListVisibleKnowledgeBaseContentsShrinkRequest {
	s.SourceTypesShrink = &v
	return s
}

func (s *ListVisibleKnowledgeBaseContentsShrinkRequest) SetTenantId(v string) *ListVisibleKnowledgeBaseContentsShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *ListVisibleKnowledgeBaseContentsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
