// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgentRuntimesInput interface {
	dara.Model
	String() string
	GoString() string
	SetAgentRuntimeName(v string) *ListAgentRuntimesInput
	GetAgentRuntimeName() *string
	SetPageNumber(v int) *ListAgentRuntimesInput
	GetPageNumber() *int
	SetPageSize(v int) *ListAgentRuntimesInput
	GetPageSize() *int
	SetStatuses(v []*string) *ListAgentRuntimesInput
	GetStatuses() []*string
}

type ListAgentRuntimesInput struct {
	// 按名称过滤
	//
	// example:
	//
	// my-agent-runtime
	AgentRuntimeName *string `json:"agentRuntimeName,omitempty" xml:"agentRuntimeName,omitempty"`
	// 页码
	//
	// example:
	//
	// 1
	PageNumber *int `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// 每页记录数
	//
	// example:
	//
	// 20
	PageSize *int `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// 按状态过滤
	//
	// example:
	//
	// READY,CREATING
	Statuses []*string `json:"statuses,omitempty" xml:"statuses,omitempty" type:"Repeated"`
}

func (s ListAgentRuntimesInput) String() string {
	return dara.Prettify(s)
}

func (s ListAgentRuntimesInput) GoString() string {
	return s.String()
}

func (s *ListAgentRuntimesInput) GetAgentRuntimeName() *string {
	return s.AgentRuntimeName
}

func (s *ListAgentRuntimesInput) GetPageNumber() *int {
	return s.PageNumber
}

func (s *ListAgentRuntimesInput) GetPageSize() *int {
	return s.PageSize
}

func (s *ListAgentRuntimesInput) GetStatuses() []*string {
	return s.Statuses
}

func (s *ListAgentRuntimesInput) SetAgentRuntimeName(v string) *ListAgentRuntimesInput {
	s.AgentRuntimeName = &v
	return s
}

func (s *ListAgentRuntimesInput) SetPageNumber(v int) *ListAgentRuntimesInput {
	s.PageNumber = &v
	return s
}

func (s *ListAgentRuntimesInput) SetPageSize(v int) *ListAgentRuntimesInput {
	s.PageSize = &v
	return s
}

func (s *ListAgentRuntimesInput) SetStatuses(v []*string) *ListAgentRuntimesInput {
	s.Statuses = v
	return s
}

func (s *ListAgentRuntimesInput) Validate() error {
	return dara.Validate(s)
}
