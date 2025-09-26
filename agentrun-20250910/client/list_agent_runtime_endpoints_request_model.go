// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgentRuntimeEndpointsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndpointName(v string) *ListAgentRuntimeEndpointsRequest
	GetEndpointName() *string
	SetPageNumber(v int32) *ListAgentRuntimeEndpointsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListAgentRuntimeEndpointsRequest
	GetPageSize() *int32
}

type ListAgentRuntimeEndpointsRequest struct {
	// 根据端点名称进行模糊匹配过滤
	//
	// example:
	//
	// my-endpoint
	EndpointName *string `json:"endpointName,omitempty" xml:"endpointName,omitempty"`
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

func (s ListAgentRuntimeEndpointsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAgentRuntimeEndpointsRequest) GoString() string {
	return s.String()
}

func (s *ListAgentRuntimeEndpointsRequest) GetEndpointName() *string {
	return s.EndpointName
}

func (s *ListAgentRuntimeEndpointsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListAgentRuntimeEndpointsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAgentRuntimeEndpointsRequest) SetEndpointName(v string) *ListAgentRuntimeEndpointsRequest {
	s.EndpointName = &v
	return s
}

func (s *ListAgentRuntimeEndpointsRequest) SetPageNumber(v int32) *ListAgentRuntimeEndpointsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAgentRuntimeEndpointsRequest) SetPageSize(v int32) *ListAgentRuntimeEndpointsRequest {
	s.PageSize = &v
	return s
}

func (s *ListAgentRuntimeEndpointsRequest) Validate() error {
	return dara.Validate(s)
}
