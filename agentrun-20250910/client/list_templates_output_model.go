// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTemplatesOutput interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*Template) *ListTemplatesOutput
	GetItems() []*Template
	SetPageNumber(v int32) *ListTemplatesOutput
	GetPageNumber() *int32
	SetPageSize(v int32) *ListTemplatesOutput
	GetPageSize() *int32
	SetTotal(v int32) *ListTemplatesOutput
	GetTotal() *int32
}

type ListTemplatesOutput struct {
	// This parameter is required.
	Items []*Template `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// This parameter is required.
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// This parameter is required.
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// This parameter is required.
	Total *int32 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListTemplatesOutput) String() string {
	return dara.Prettify(s)
}

func (s ListTemplatesOutput) GoString() string {
	return s.String()
}

func (s *ListTemplatesOutput) GetItems() []*Template {
	return s.Items
}

func (s *ListTemplatesOutput) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListTemplatesOutput) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListTemplatesOutput) GetTotal() *int32 {
	return s.Total
}

func (s *ListTemplatesOutput) SetItems(v []*Template) *ListTemplatesOutput {
	s.Items = v
	return s
}

func (s *ListTemplatesOutput) SetPageNumber(v int32) *ListTemplatesOutput {
	s.PageNumber = &v
	return s
}

func (s *ListTemplatesOutput) SetPageSize(v int32) *ListTemplatesOutput {
	s.PageSize = &v
	return s
}

func (s *ListTemplatesOutput) SetTotal(v int32) *ListTemplatesOutput {
	s.Total = &v
	return s
}

func (s *ListTemplatesOutput) Validate() error {
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
