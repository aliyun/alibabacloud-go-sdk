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
	// 父分类 ID；不传时返回企业知识库根目录下的所有分类树
	//
	// example:
	//
	// exampleDirectoryId
	DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
	// 排序字段：name / gmt_create / gmt_modified
	//
	// example:
	//
	// name
	SortField *string `json:"sortField,omitempty" xml:"sortField,omitempty"`
	// 排序方向：asc / desc
	//
	// example:
	//
	// asc
	SortOrder *string `json:"sortOrder,omitempty" xml:"sortOrder,omitempty"`
	// 租户ID，公共参数；winnexo-cli 通过 --tenant-id 显式传入
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
