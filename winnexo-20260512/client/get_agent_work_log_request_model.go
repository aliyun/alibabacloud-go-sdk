// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentWorkLogRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNodeId(v string) *GetAgentWorkLogRequest
	GetNodeId() *string
	SetPage(v int64) *GetAgentWorkLogRequest
	GetPage() *int64
	SetPageSize(v int64) *GetAgentWorkLogRequest
	GetPageSize() *int64
	SetTenantId(v string) *GetAgentWorkLogRequest
	GetTenantId() *string
	SetWorkItemId(v string) *GetAgentWorkLogRequest
	GetWorkItemId() *string
}

type GetAgentWorkLogRequest struct {
	// 可选节点 key
	//
	// example:
	//
	// step_1
	NodeId *string `json:"nodeId,omitempty" xml:"nodeId,omitempty"`
	// 页码，从 1 开始
	//
	// example:
	//
	// 1
	Page *int64 `json:"page,omitempty" xml:"page,omitempty"`
	// 每页条数，最大 100
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// 租户ID，缺省使用调用方默认租户
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// 工作项ID
	//
	// This parameter is required.
	//
	// example:
	//
	// wi_example
	WorkItemId *string `json:"workItemId,omitempty" xml:"workItemId,omitempty"`
}

func (s GetAgentWorkLogRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAgentWorkLogRequest) GoString() string {
	return s.String()
}

func (s *GetAgentWorkLogRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *GetAgentWorkLogRequest) GetPage() *int64 {
	return s.Page
}

func (s *GetAgentWorkLogRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *GetAgentWorkLogRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *GetAgentWorkLogRequest) GetWorkItemId() *string {
	return s.WorkItemId
}

func (s *GetAgentWorkLogRequest) SetNodeId(v string) *GetAgentWorkLogRequest {
	s.NodeId = &v
	return s
}

func (s *GetAgentWorkLogRequest) SetPage(v int64) *GetAgentWorkLogRequest {
	s.Page = &v
	return s
}

func (s *GetAgentWorkLogRequest) SetPageSize(v int64) *GetAgentWorkLogRequest {
	s.PageSize = &v
	return s
}

func (s *GetAgentWorkLogRequest) SetTenantId(v string) *GetAgentWorkLogRequest {
	s.TenantId = &v
	return s
}

func (s *GetAgentWorkLogRequest) SetWorkItemId(v string) *GetAgentWorkLogRequest {
	s.WorkItemId = &v
	return s
}

func (s *GetAgentWorkLogRequest) Validate() error {
	return dara.Validate(s)
}
