// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgentRuntimeEndpointsOutput interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*AgentRuntimeEndpoint) *ListAgentRuntimeEndpointsOutput
	GetItems() []*AgentRuntimeEndpoint
	SetPageNumber(v int) *ListAgentRuntimeEndpointsOutput
	GetPageNumber() *int
	SetPageSize(v int) *ListAgentRuntimeEndpointsOutput
	GetPageSize() *int
	SetTotal(v int) *ListAgentRuntimeEndpointsOutput
	GetTotal() *int
}

type ListAgentRuntimeEndpointsOutput struct {
	Items      []*AgentRuntimeEndpoint `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	PageNumber *int                    `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *int                    `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Total      *int                    `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListAgentRuntimeEndpointsOutput) String() string {
	return dara.Prettify(s)
}

func (s ListAgentRuntimeEndpointsOutput) GoString() string {
	return s.String()
}

func (s *ListAgentRuntimeEndpointsOutput) GetItems() []*AgentRuntimeEndpoint {
	return s.Items
}

func (s *ListAgentRuntimeEndpointsOutput) GetPageNumber() *int {
	return s.PageNumber
}

func (s *ListAgentRuntimeEndpointsOutput) GetPageSize() *int {
	return s.PageSize
}

func (s *ListAgentRuntimeEndpointsOutput) GetTotal() *int {
	return s.Total
}

func (s *ListAgentRuntimeEndpointsOutput) SetItems(v []*AgentRuntimeEndpoint) *ListAgentRuntimeEndpointsOutput {
	s.Items = v
	return s
}

func (s *ListAgentRuntimeEndpointsOutput) SetPageNumber(v int) *ListAgentRuntimeEndpointsOutput {
	s.PageNumber = &v
	return s
}

func (s *ListAgentRuntimeEndpointsOutput) SetPageSize(v int) *ListAgentRuntimeEndpointsOutput {
	s.PageSize = &v
	return s
}

func (s *ListAgentRuntimeEndpointsOutput) SetTotal(v int) *ListAgentRuntimeEndpointsOutput {
	s.Total = &v
	return s
}

func (s *ListAgentRuntimeEndpointsOutput) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
