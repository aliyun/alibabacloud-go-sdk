// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkflowTemplatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCollaborationGroupId(v string) *ListWorkflowTemplatesRequest
	GetCollaborationGroupId() *string
	SetKeyword(v string) *ListWorkflowTemplatesRequest
	GetKeyword() *string
	SetMaxResults(v int64) *ListWorkflowTemplatesRequest
	GetMaxResults() *int64
	SetNextToken(v string) *ListWorkflowTemplatesRequest
	GetNextToken() *string
	SetPage(v int64) *ListWorkflowTemplatesRequest
	GetPage() *int64
	SetPageSize(v int64) *ListWorkflowTemplatesRequest
	GetPageSize() *int64
	SetTenantId(v string) *ListWorkflowTemplatesRequest
	GetTenantId() *string
}

type ListWorkflowTemplatesRequest struct {
	// 协作空间 ID
	//
	// This parameter is required.
	//
	// example:
	//
	// group_example
	CollaborationGroupId *string `json:"collaborationGroupId,omitempty" xml:"collaborationGroupId,omitempty"`
	// 模板名称/描述模糊搜索关键词
	//
	// example:
	//
	// example
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// 单页最大返回数量（1~100）；传入时优先于 pageSize
	//
	// example:
	//
	// 1
	MaxResults *int64 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// 翻页令牌，取上次响应返回的 nextToken；传入时优先于 page
	//
	// example:
	//
	// example
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 页码；传入 nextToken 时忽略
	//
	// example:
	//
	// 1
	Page *int64 `json:"page,omitempty" xml:"page,omitempty"`
	// 每页条数（1~100）
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// 租户ID，公共参数；缺省时使用调用方默认租户
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ListWorkflowTemplatesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListWorkflowTemplatesRequest) GoString() string {
	return s.String()
}

func (s *ListWorkflowTemplatesRequest) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *ListWorkflowTemplatesRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListWorkflowTemplatesRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *ListWorkflowTemplatesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListWorkflowTemplatesRequest) GetPage() *int64 {
	return s.Page
}

func (s *ListWorkflowTemplatesRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListWorkflowTemplatesRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListWorkflowTemplatesRequest) SetCollaborationGroupId(v string) *ListWorkflowTemplatesRequest {
	s.CollaborationGroupId = &v
	return s
}

func (s *ListWorkflowTemplatesRequest) SetKeyword(v string) *ListWorkflowTemplatesRequest {
	s.Keyword = &v
	return s
}

func (s *ListWorkflowTemplatesRequest) SetMaxResults(v int64) *ListWorkflowTemplatesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListWorkflowTemplatesRequest) SetNextToken(v string) *ListWorkflowTemplatesRequest {
	s.NextToken = &v
	return s
}

func (s *ListWorkflowTemplatesRequest) SetPage(v int64) *ListWorkflowTemplatesRequest {
	s.Page = &v
	return s
}

func (s *ListWorkflowTemplatesRequest) SetPageSize(v int64) *ListWorkflowTemplatesRequest {
	s.PageSize = &v
	return s
}

func (s *ListWorkflowTemplatesRequest) SetTenantId(v string) *ListWorkflowTemplatesRequest {
	s.TenantId = &v
	return s
}

func (s *ListWorkflowTemplatesRequest) Validate() error {
	return dara.Validate(s)
}
