// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVisibleKnowledgeBaseContentsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirectoryId(v string) *ListVisibleKnowledgeBaseContentsRequest
	GetDirectoryId() *string
	SetOperatingObjectName(v string) *ListVisibleKnowledgeBaseContentsRequest
	GetOperatingObjectName() *string
	SetPage(v int64) *ListVisibleKnowledgeBaseContentsRequest
	GetPage() *int64
	SetPageSize(v int64) *ListVisibleKnowledgeBaseContentsRequest
	GetPageSize() *int64
	SetSortField(v string) *ListVisibleKnowledgeBaseContentsRequest
	GetSortField() *string
	SetSortOrder(v string) *ListVisibleKnowledgeBaseContentsRequest
	GetSortOrder() *string
	SetSourceTypes(v []*string) *ListVisibleKnowledgeBaseContentsRequest
	GetSourceTypes() []*string
	SetTenantId(v string) *ListVisibleKnowledgeBaseContentsRequest
	GetTenantId() *string
}

type ListVisibleKnowledgeBaseContentsRequest struct {
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
	SourceTypes []*string `json:"sourceTypes,omitempty" xml:"sourceTypes,omitempty" type:"Repeated"`
	// The tenant ID.
	//
	// example:
	//
	// PiPklI1iSRTm6VFFqlY9VzbgiEiE
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ListVisibleKnowledgeBaseContentsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListVisibleKnowledgeBaseContentsRequest) GoString() string {
	return s.String()
}

func (s *ListVisibleKnowledgeBaseContentsRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *ListVisibleKnowledgeBaseContentsRequest) GetOperatingObjectName() *string {
	return s.OperatingObjectName
}

func (s *ListVisibleKnowledgeBaseContentsRequest) GetPage() *int64 {
	return s.Page
}

func (s *ListVisibleKnowledgeBaseContentsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListVisibleKnowledgeBaseContentsRequest) GetSortField() *string {
	return s.SortField
}

func (s *ListVisibleKnowledgeBaseContentsRequest) GetSortOrder() *string {
	return s.SortOrder
}

func (s *ListVisibleKnowledgeBaseContentsRequest) GetSourceTypes() []*string {
	return s.SourceTypes
}

func (s *ListVisibleKnowledgeBaseContentsRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListVisibleKnowledgeBaseContentsRequest) SetDirectoryId(v string) *ListVisibleKnowledgeBaseContentsRequest {
	s.DirectoryId = &v
	return s
}

func (s *ListVisibleKnowledgeBaseContentsRequest) SetOperatingObjectName(v string) *ListVisibleKnowledgeBaseContentsRequest {
	s.OperatingObjectName = &v
	return s
}

func (s *ListVisibleKnowledgeBaseContentsRequest) SetPage(v int64) *ListVisibleKnowledgeBaseContentsRequest {
	s.Page = &v
	return s
}

func (s *ListVisibleKnowledgeBaseContentsRequest) SetPageSize(v int64) *ListVisibleKnowledgeBaseContentsRequest {
	s.PageSize = &v
	return s
}

func (s *ListVisibleKnowledgeBaseContentsRequest) SetSortField(v string) *ListVisibleKnowledgeBaseContentsRequest {
	s.SortField = &v
	return s
}

func (s *ListVisibleKnowledgeBaseContentsRequest) SetSortOrder(v string) *ListVisibleKnowledgeBaseContentsRequest {
	s.SortOrder = &v
	return s
}

func (s *ListVisibleKnowledgeBaseContentsRequest) SetSourceTypes(v []*string) *ListVisibleKnowledgeBaseContentsRequest {
	s.SourceTypes = v
	return s
}

func (s *ListVisibleKnowledgeBaseContentsRequest) SetTenantId(v string) *ListVisibleKnowledgeBaseContentsRequest {
	s.TenantId = &v
	return s
}

func (s *ListVisibleKnowledgeBaseContentsRequest) Validate() error {
	return dara.Validate(s)
}
