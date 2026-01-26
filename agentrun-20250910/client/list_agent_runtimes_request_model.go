// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgentRuntimesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentRuntimeName(v string) *ListAgentRuntimesRequest
	GetAgentRuntimeName() *string
	SetDiscoveryResourceGroupId(v string) *ListAgentRuntimesRequest
	GetDiscoveryResourceGroupId() *string
	SetPageNumber(v int32) *ListAgentRuntimesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListAgentRuntimesRequest
	GetPageSize() *int32
	SetResourceGroupId(v string) *ListAgentRuntimesRequest
	GetResourceGroupId() *string
	SetSearchMode(v string) *ListAgentRuntimesRequest
	GetSearchMode() *string
	SetStatus(v string) *ListAgentRuntimesRequest
	GetStatus() *string
}

type ListAgentRuntimesRequest struct {
	// 根据智能体运行时名称进行模糊匹配过滤
	//
	// example:
	//
	// my-runtime
	AgentRuntimeName *string `json:"agentRuntimeName,omitempty" xml:"agentRuntimeName,omitempty"`
	// 用于服务发现的资源组标识符
	//
	// example:
	//
	// rg-123456
	DiscoveryResourceGroupId *string `json:"discoveryResourceGroupId,omitempty" xml:"discoveryResourceGroupId,omitempty"`
	// 当前页码，从1开始计数
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// 每页返回的记录数量
	//
	// example:
	//
	// 10
	PageSize        *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// 查询模式，支持精确查询和模糊查询
	//
	// example:
	//
	// fuzzy
	SearchMode *string `json:"searchMode,omitempty" xml:"searchMode,omitempty"`
	// 根据状态进行过滤，多个状态用逗号分隔，支持精确匹配
	//
	// example:
	//
	// READY,CREATING
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ListAgentRuntimesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAgentRuntimesRequest) GoString() string {
	return s.String()
}

func (s *ListAgentRuntimesRequest) GetAgentRuntimeName() *string {
	return s.AgentRuntimeName
}

func (s *ListAgentRuntimesRequest) GetDiscoveryResourceGroupId() *string {
	return s.DiscoveryResourceGroupId
}

func (s *ListAgentRuntimesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListAgentRuntimesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAgentRuntimesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListAgentRuntimesRequest) GetSearchMode() *string {
	return s.SearchMode
}

func (s *ListAgentRuntimesRequest) GetStatus() *string {
	return s.Status
}

func (s *ListAgentRuntimesRequest) SetAgentRuntimeName(v string) *ListAgentRuntimesRequest {
	s.AgentRuntimeName = &v
	return s
}

func (s *ListAgentRuntimesRequest) SetDiscoveryResourceGroupId(v string) *ListAgentRuntimesRequest {
	s.DiscoveryResourceGroupId = &v
	return s
}

func (s *ListAgentRuntimesRequest) SetPageNumber(v int32) *ListAgentRuntimesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAgentRuntimesRequest) SetPageSize(v int32) *ListAgentRuntimesRequest {
	s.PageSize = &v
	return s
}

func (s *ListAgentRuntimesRequest) SetResourceGroupId(v string) *ListAgentRuntimesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListAgentRuntimesRequest) SetSearchMode(v string) *ListAgentRuntimesRequest {
	s.SearchMode = &v
	return s
}

func (s *ListAgentRuntimesRequest) SetStatus(v string) *ListAgentRuntimesRequest {
	s.Status = &v
	return s
}

func (s *ListAgentRuntimesRequest) Validate() error {
	return dara.Validate(s)
}
