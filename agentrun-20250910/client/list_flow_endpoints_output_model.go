// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFlowEndpointsOutput interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*FlowEndpoint) *ListFlowEndpointsOutput
	GetItems() []*FlowEndpoint
	SetPageNumber(v int) *ListFlowEndpointsOutput
	GetPageNumber() *int
	SetPageSize(v int) *ListFlowEndpointsOutput
	GetPageSize() *int
	SetTotal(v int) *ListFlowEndpointsOutput
	GetTotal() *int
}

type ListFlowEndpointsOutput struct {
	Items      []*FlowEndpoint `json:"items" xml:"items" type:"Repeated"`
	PageNumber *int            `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *int            `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Total      *int            `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListFlowEndpointsOutput) String() string {
	return dara.Prettify(s)
}

func (s ListFlowEndpointsOutput) GoString() string {
	return s.String()
}

func (s *ListFlowEndpointsOutput) GetItems() []*FlowEndpoint {
	return s.Items
}

func (s *ListFlowEndpointsOutput) GetPageNumber() *int {
	return s.PageNumber
}

func (s *ListFlowEndpointsOutput) GetPageSize() *int {
	return s.PageSize
}

func (s *ListFlowEndpointsOutput) GetTotal() *int {
	return s.Total
}

func (s *ListFlowEndpointsOutput) SetItems(v []*FlowEndpoint) *ListFlowEndpointsOutput {
	s.Items = v
	return s
}

func (s *ListFlowEndpointsOutput) SetPageNumber(v int) *ListFlowEndpointsOutput {
	s.PageNumber = &v
	return s
}

func (s *ListFlowEndpointsOutput) SetPageSize(v int) *ListFlowEndpointsOutput {
	s.PageSize = &v
	return s
}

func (s *ListFlowEndpointsOutput) SetTotal(v int) *ListFlowEndpointsOutput {
	s.Total = &v
	return s
}

func (s *ListFlowEndpointsOutput) Validate() error {
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
