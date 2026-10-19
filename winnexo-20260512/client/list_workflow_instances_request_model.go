// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkflowInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessId(v string) *ListWorkflowInstancesRequest
	GetBusinessId() *string
	SetCollaborationGroupId(v string) *ListWorkflowInstancesRequest
	GetCollaborationGroupId() *string
	SetInitiator(v string) *ListWorkflowInstancesRequest
	GetInitiator() *string
	SetKeyword(v string) *ListWorkflowInstancesRequest
	GetKeyword() *string
	SetPage(v int64) *ListWorkflowInstancesRequest
	GetPage() *int64
	SetPageSize(v int64) *ListWorkflowInstancesRequest
	GetPageSize() *int64
	SetSortBy(v string) *ListWorkflowInstancesRequest
	GetSortBy() *string
	SetSortOrder(v string) *ListWorkflowInstancesRequest
	GetSortOrder() *string
	SetStatus(v string) *ListWorkflowInstancesRequest
	GetStatus() *string
	SetTenantId(v string) *ListWorkflowInstancesRequest
	GetTenantId() *string
}

type ListWorkflowInstancesRequest struct {
	// 精确需求编号
	//
	// example:
	//
	// REQ-1001
	BusinessId *string `json:"businessId,omitempty" xml:"businessId,omitempty"`
	// 协作空间 ID
	//
	// This parameter is required.
	//
	// example:
	//
	// cg_20261016_abc123
	CollaborationGroupId *string `json:"collaborationGroupId,omitempty" xml:"collaborationGroupId,omitempty"`
	// 精确外部发起人标签
	//
	// example:
	//
	// 外部发起人张三
	Initiator *string `json:"initiator,omitempty" xml:"initiator,omitempty"`
	// 标题或需求编号关键词
	//
	// example:
	//
	// 季度报告
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// 页码，从 1 开始
	//
	// example:
	//
	// 1
	Page *int64 `json:"page,omitempty" xml:"page,omitempty"`
	// 每页数量 1-100
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// 排序字段
	//
	// example:
	//
	// GMT_CREATE
	SortBy *string `json:"sortBy,omitempty" xml:"sortBy,omitempty"`
	// 排序方向
	//
	// example:
	//
	// DESC
	SortOrder *string `json:"sortOrder,omitempty" xml:"sortOrder,omitempty"`
	// Running/Completed/Failed/Terminated 筛选
	//
	// example:
	//
	// Running
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// 租户ID，公共参数；winnexo-cli 通过 --tenant-id 显式传入
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ListWorkflowInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListWorkflowInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListWorkflowInstancesRequest) GetBusinessId() *string {
	return s.BusinessId
}

func (s *ListWorkflowInstancesRequest) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *ListWorkflowInstancesRequest) GetInitiator() *string {
	return s.Initiator
}

func (s *ListWorkflowInstancesRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListWorkflowInstancesRequest) GetPage() *int64 {
	return s.Page
}

func (s *ListWorkflowInstancesRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListWorkflowInstancesRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListWorkflowInstancesRequest) GetSortOrder() *string {
	return s.SortOrder
}

func (s *ListWorkflowInstancesRequest) GetStatus() *string {
	return s.Status
}

func (s *ListWorkflowInstancesRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListWorkflowInstancesRequest) SetBusinessId(v string) *ListWorkflowInstancesRequest {
	s.BusinessId = &v
	return s
}

func (s *ListWorkflowInstancesRequest) SetCollaborationGroupId(v string) *ListWorkflowInstancesRequest {
	s.CollaborationGroupId = &v
	return s
}

func (s *ListWorkflowInstancesRequest) SetInitiator(v string) *ListWorkflowInstancesRequest {
	s.Initiator = &v
	return s
}

func (s *ListWorkflowInstancesRequest) SetKeyword(v string) *ListWorkflowInstancesRequest {
	s.Keyword = &v
	return s
}

func (s *ListWorkflowInstancesRequest) SetPage(v int64) *ListWorkflowInstancesRequest {
	s.Page = &v
	return s
}

func (s *ListWorkflowInstancesRequest) SetPageSize(v int64) *ListWorkflowInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *ListWorkflowInstancesRequest) SetSortBy(v string) *ListWorkflowInstancesRequest {
	s.SortBy = &v
	return s
}

func (s *ListWorkflowInstancesRequest) SetSortOrder(v string) *ListWorkflowInstancesRequest {
	s.SortOrder = &v
	return s
}

func (s *ListWorkflowInstancesRequest) SetStatus(v string) *ListWorkflowInstancesRequest {
	s.Status = &v
	return s
}

func (s *ListWorkflowInstancesRequest) SetTenantId(v string) *ListWorkflowInstancesRequest {
	s.TenantId = &v
	return s
}

func (s *ListWorkflowInstancesRequest) Validate() error {
	return dara.Validate(s)
}
