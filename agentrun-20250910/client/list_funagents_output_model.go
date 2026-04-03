// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFunagentsOutput interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*Funagent) *ListFunagentsOutput
	GetItems() []*Funagent
	SetPageNumber(v int) *ListFunagentsOutput
	GetPageNumber() *int
	SetPageSize(v int) *ListFunagentsOutput
	GetPageSize() *int
	SetTotal(v int) *ListFunagentsOutput
	GetTotal() *int
}

type ListFunagentsOutput struct {
	Items      []*Funagent `json:"items" xml:"items" type:"Repeated"`
	PageNumber *int        `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *int        `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Total      *int        `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListFunagentsOutput) String() string {
	return dara.Prettify(s)
}

func (s ListFunagentsOutput) GoString() string {
	return s.String()
}

func (s *ListFunagentsOutput) GetItems() []*Funagent {
	return s.Items
}

func (s *ListFunagentsOutput) GetPageNumber() *int {
	return s.PageNumber
}

func (s *ListFunagentsOutput) GetPageSize() *int {
	return s.PageSize
}

func (s *ListFunagentsOutput) GetTotal() *int {
	return s.Total
}

func (s *ListFunagentsOutput) SetItems(v []*Funagent) *ListFunagentsOutput {
	s.Items = v
	return s
}

func (s *ListFunagentsOutput) SetPageNumber(v int) *ListFunagentsOutput {
	s.PageNumber = &v
	return s
}

func (s *ListFunagentsOutput) SetPageSize(v int) *ListFunagentsOutput {
	s.PageSize = &v
	return s
}

func (s *ListFunagentsOutput) SetTotal(v int) *ListFunagentsOutput {
	s.Total = &v
	return s
}

func (s *ListFunagentsOutput) Validate() error {
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
