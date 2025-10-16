// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgentRuntimeVersionsOutput interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*AgentRuntimeVersion) *ListAgentRuntimeVersionsOutput
	GetItems() []*AgentRuntimeVersion
	SetPageNumber(v int) *ListAgentRuntimeVersionsOutput
	GetPageNumber() *int
	SetPageSize(v int) *ListAgentRuntimeVersionsOutput
	GetPageSize() *int
	SetTotal(v int) *ListAgentRuntimeVersionsOutput
	GetTotal() *int
}

type ListAgentRuntimeVersionsOutput struct {
	Items      []*AgentRuntimeVersion `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	PageNumber *int                   `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *int                   `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Total      *int                   `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListAgentRuntimeVersionsOutput) String() string {
	return dara.Prettify(s)
}

func (s ListAgentRuntimeVersionsOutput) GoString() string {
	return s.String()
}

func (s *ListAgentRuntimeVersionsOutput) GetItems() []*AgentRuntimeVersion {
	return s.Items
}

func (s *ListAgentRuntimeVersionsOutput) GetPageNumber() *int {
	return s.PageNumber
}

func (s *ListAgentRuntimeVersionsOutput) GetPageSize() *int {
	return s.PageSize
}

func (s *ListAgentRuntimeVersionsOutput) GetTotal() *int {
	return s.Total
}

func (s *ListAgentRuntimeVersionsOutput) SetItems(v []*AgentRuntimeVersion) *ListAgentRuntimeVersionsOutput {
	s.Items = v
	return s
}

func (s *ListAgentRuntimeVersionsOutput) SetPageNumber(v int) *ListAgentRuntimeVersionsOutput {
	s.PageNumber = &v
	return s
}

func (s *ListAgentRuntimeVersionsOutput) SetPageSize(v int) *ListAgentRuntimeVersionsOutput {
	s.PageSize = &v
	return s
}

func (s *ListAgentRuntimeVersionsOutput) SetTotal(v int) *ListAgentRuntimeVersionsOutput {
	s.Total = &v
	return s
}

func (s *ListAgentRuntimeVersionsOutput) Validate() error {
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
