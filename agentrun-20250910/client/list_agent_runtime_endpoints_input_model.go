// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgentRuntimeEndpointsInput interface {
	dara.Model
	String() string
	GoString() string
	SetEndpointName(v string) *ListAgentRuntimeEndpointsInput
	GetEndpointName() *string
	SetPageNumber(v int) *ListAgentRuntimeEndpointsInput
	GetPageNumber() *int
	SetPageSize(v int) *ListAgentRuntimeEndpointsInput
	GetPageSize() *int
	SetStatuses(v []*string) *ListAgentRuntimeEndpointsInput
	GetStatuses() []*string
}

type ListAgentRuntimeEndpointsInput struct {
	// 按端点名称过滤
	EndpointName *string `json:"endpointName,omitempty" xml:"endpointName,omitempty"`
	// 页码
	PageNumber *int `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// 每页记录数
	PageSize *int `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// 按状态过滤
	Statuses []*string `json:"statuses,omitempty" xml:"statuses,omitempty" type:"Repeated"`
}

func (s ListAgentRuntimeEndpointsInput) String() string {
	return dara.Prettify(s)
}

func (s ListAgentRuntimeEndpointsInput) GoString() string {
	return s.String()
}

func (s *ListAgentRuntimeEndpointsInput) GetEndpointName() *string {
	return s.EndpointName
}

func (s *ListAgentRuntimeEndpointsInput) GetPageNumber() *int {
	return s.PageNumber
}

func (s *ListAgentRuntimeEndpointsInput) GetPageSize() *int {
	return s.PageSize
}

func (s *ListAgentRuntimeEndpointsInput) GetStatuses() []*string {
	return s.Statuses
}

func (s *ListAgentRuntimeEndpointsInput) SetEndpointName(v string) *ListAgentRuntimeEndpointsInput {
	s.EndpointName = &v
	return s
}

func (s *ListAgentRuntimeEndpointsInput) SetPageNumber(v int) *ListAgentRuntimeEndpointsInput {
	s.PageNumber = &v
	return s
}

func (s *ListAgentRuntimeEndpointsInput) SetPageSize(v int) *ListAgentRuntimeEndpointsInput {
	s.PageSize = &v
	return s
}

func (s *ListAgentRuntimeEndpointsInput) SetStatuses(v []*string) *ListAgentRuntimeEndpointsInput {
	s.Statuses = v
	return s
}

func (s *ListAgentRuntimeEndpointsInput) Validate() error {
	return dara.Validate(s)
}
