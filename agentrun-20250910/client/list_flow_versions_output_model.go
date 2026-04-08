// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFlowVersionsOutput interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*FlowVersion) *ListFlowVersionsOutput
	GetItems() []*FlowVersion
	SetPageNumber(v int) *ListFlowVersionsOutput
	GetPageNumber() *int
	SetPageSize(v int) *ListFlowVersionsOutput
	GetPageSize() *int
	SetTotal(v int) *ListFlowVersionsOutput
	GetTotal() *int
}

type ListFlowVersionsOutput struct {
	Items      []*FlowVersion `json:"items" xml:"items" type:"Repeated"`
	PageNumber *int           `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *int           `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Total      *int           `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListFlowVersionsOutput) String() string {
	return dara.Prettify(s)
}

func (s ListFlowVersionsOutput) GoString() string {
	return s.String()
}

func (s *ListFlowVersionsOutput) GetItems() []*FlowVersion {
	return s.Items
}

func (s *ListFlowVersionsOutput) GetPageNumber() *int {
	return s.PageNumber
}

func (s *ListFlowVersionsOutput) GetPageSize() *int {
	return s.PageSize
}

func (s *ListFlowVersionsOutput) GetTotal() *int {
	return s.Total
}

func (s *ListFlowVersionsOutput) SetItems(v []*FlowVersion) *ListFlowVersionsOutput {
	s.Items = v
	return s
}

func (s *ListFlowVersionsOutput) SetPageNumber(v int) *ListFlowVersionsOutput {
	s.PageNumber = &v
	return s
}

func (s *ListFlowVersionsOutput) SetPageSize(v int) *ListFlowVersionsOutput {
	s.PageSize = &v
	return s
}

func (s *ListFlowVersionsOutput) SetTotal(v int) *ListFlowVersionsOutput {
	s.Total = &v
	return s
}

func (s *ListFlowVersionsOutput) Validate() error {
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
