// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListKnowledgeBaseDirectoriesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirectoryId(v string) *ListKnowledgeBaseDirectoriesRequest
	GetDirectoryId() *string
	SetSortField(v string) *ListKnowledgeBaseDirectoriesRequest
	GetSortField() *string
	SetSortOrder(v string) *ListKnowledgeBaseDirectoriesRequest
	GetSortOrder() *string
	SetTenantId(v string) *ListKnowledgeBaseDirectoriesRequest
	GetTenantId() *string
}

type ListKnowledgeBaseDirectoriesRequest struct {
	// The directory ID.
	//
	// example:
	//
	// exampleDirectoryId
	DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
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
	// The sort order. This parameter takes effect only when sortBy is specified. Valid values: ASC, DESC (case-insensitive).
	//
	// example:
	//
	// asc
	SortOrder *string `json:"sortOrder,omitempty" xml:"sortOrder,omitempty"`
	// The tenant ID. This is a common parameter. You can pass it explicitly by using --tenant-id in winnexo-cli.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ListKnowledgeBaseDirectoriesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListKnowledgeBaseDirectoriesRequest) GoString() string {
	return s.String()
}

func (s *ListKnowledgeBaseDirectoriesRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *ListKnowledgeBaseDirectoriesRequest) GetSortField() *string {
	return s.SortField
}

func (s *ListKnowledgeBaseDirectoriesRequest) GetSortOrder() *string {
	return s.SortOrder
}

func (s *ListKnowledgeBaseDirectoriesRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListKnowledgeBaseDirectoriesRequest) SetDirectoryId(v string) *ListKnowledgeBaseDirectoriesRequest {
	s.DirectoryId = &v
	return s
}

func (s *ListKnowledgeBaseDirectoriesRequest) SetSortField(v string) *ListKnowledgeBaseDirectoriesRequest {
	s.SortField = &v
	return s
}

func (s *ListKnowledgeBaseDirectoriesRequest) SetSortOrder(v string) *ListKnowledgeBaseDirectoriesRequest {
	s.SortOrder = &v
	return s
}

func (s *ListKnowledgeBaseDirectoriesRequest) SetTenantId(v string) *ListKnowledgeBaseDirectoriesRequest {
	s.TenantId = &v
	return s
}

func (s *ListKnowledgeBaseDirectoriesRequest) Validate() error {
	return dara.Validate(s)
}
