// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFlowsOutput interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*Flow) *ListFlowsOutput
	GetItems() []*Flow
	SetPageNumber(v int) *ListFlowsOutput
	GetPageNumber() *int
	SetPageSize(v int) *ListFlowsOutput
	GetPageSize() *int
	SetTotal(v int) *ListFlowsOutput
	GetTotal() *int
}

type ListFlowsOutput struct {
	Items      []*Flow `json:"items" xml:"items" type:"Repeated"`
	PageNumber *int    `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *int    `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Total      *int    `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListFlowsOutput) String() string {
	return dara.Prettify(s)
}

func (s ListFlowsOutput) GoString() string {
	return s.String()
}

func (s *ListFlowsOutput) GetItems() []*Flow {
	return s.Items
}

func (s *ListFlowsOutput) GetPageNumber() *int {
	return s.PageNumber
}

func (s *ListFlowsOutput) GetPageSize() *int {
	return s.PageSize
}

func (s *ListFlowsOutput) GetTotal() *int {
	return s.Total
}

func (s *ListFlowsOutput) SetItems(v []*Flow) *ListFlowsOutput {
	s.Items = v
	return s
}

func (s *ListFlowsOutput) SetPageNumber(v int) *ListFlowsOutput {
	s.PageNumber = &v
	return s
}

func (s *ListFlowsOutput) SetPageSize(v int) *ListFlowsOutput {
	s.PageSize = &v
	return s
}

func (s *ListFlowsOutput) SetTotal(v int) *ListFlowsOutput {
	s.Total = &v
	return s
}

func (s *ListFlowsOutput) Validate() error {
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
