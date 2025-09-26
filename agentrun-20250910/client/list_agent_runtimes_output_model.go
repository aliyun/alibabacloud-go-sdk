// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgentRuntimesOutput interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*AgentRuntime) *ListAgentRuntimesOutput
	GetItems() []*AgentRuntime
	SetPageNumber(v int) *ListAgentRuntimesOutput
	GetPageNumber() *int
	SetPageSize(v int) *ListAgentRuntimesOutput
	GetPageSize() *int
	SetTotal(v int) *ListAgentRuntimesOutput
	GetTotal() *int
}

type ListAgentRuntimesOutput struct {
	Items      []*AgentRuntime `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	PageNumber *int            `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *int            `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Total      *int            `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListAgentRuntimesOutput) String() string {
	return dara.Prettify(s)
}

func (s ListAgentRuntimesOutput) GoString() string {
	return s.String()
}

func (s *ListAgentRuntimesOutput) GetItems() []*AgentRuntime {
	return s.Items
}

func (s *ListAgentRuntimesOutput) GetPageNumber() *int {
	return s.PageNumber
}

func (s *ListAgentRuntimesOutput) GetPageSize() *int {
	return s.PageSize
}

func (s *ListAgentRuntimesOutput) GetTotal() *int {
	return s.Total
}

func (s *ListAgentRuntimesOutput) SetItems(v []*AgentRuntime) *ListAgentRuntimesOutput {
	s.Items = v
	return s
}

func (s *ListAgentRuntimesOutput) SetPageNumber(v int) *ListAgentRuntimesOutput {
	s.PageNumber = &v
	return s
}

func (s *ListAgentRuntimesOutput) SetPageSize(v int) *ListAgentRuntimesOutput {
	s.PageSize = &v
	return s
}

func (s *ListAgentRuntimesOutput) SetTotal(v int) *ListAgentRuntimesOutput {
	s.Total = &v
	return s
}

func (s *ListAgentRuntimesOutput) Validate() error {
	return dara.Validate(s)
}
