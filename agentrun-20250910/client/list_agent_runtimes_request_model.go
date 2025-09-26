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
	SetPageNumber(v int32) *ListAgentRuntimesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListAgentRuntimesRequest
	GetPageSize() *int32
}

type ListAgentRuntimesRequest struct {
	// 根据智能体运行时名称进行模糊匹配过滤
	//
	// example:
	//
	// my-runtime
	AgentRuntimeName *string `json:"agentRuntimeName,omitempty" xml:"agentRuntimeName,omitempty"`
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
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
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

func (s *ListAgentRuntimesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListAgentRuntimesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAgentRuntimesRequest) SetAgentRuntimeName(v string) *ListAgentRuntimesRequest {
	s.AgentRuntimeName = &v
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

func (s *ListAgentRuntimesRequest) Validate() error {
	return dara.Validate(s)
}
