// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCustomDomainsOutput interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*CustomDomain) *ListCustomDomainsOutput
	GetItems() []*CustomDomain
	SetPageNumber(v int32) *ListCustomDomainsOutput
	GetPageNumber() *int32
	SetPageSize(v int32) *ListCustomDomainsOutput
	GetPageSize() *int32
	SetTotal(v int32) *ListCustomDomainsOutput
	GetTotal() *int32
}

type ListCustomDomainsOutput struct {
	Items      []*CustomDomain `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	PageNumber *int32          `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *int32          `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Total      *int32          `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListCustomDomainsOutput) String() string {
	return dara.Prettify(s)
}

func (s ListCustomDomainsOutput) GoString() string {
	return s.String()
}

func (s *ListCustomDomainsOutput) GetItems() []*CustomDomain {
	return s.Items
}

func (s *ListCustomDomainsOutput) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCustomDomainsOutput) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCustomDomainsOutput) GetTotal() *int32 {
	return s.Total
}

func (s *ListCustomDomainsOutput) SetItems(v []*CustomDomain) *ListCustomDomainsOutput {
	s.Items = v
	return s
}

func (s *ListCustomDomainsOutput) SetPageNumber(v int32) *ListCustomDomainsOutput {
	s.PageNumber = &v
	return s
}

func (s *ListCustomDomainsOutput) SetPageSize(v int32) *ListCustomDomainsOutput {
	s.PageSize = &v
	return s
}

func (s *ListCustomDomainsOutput) SetTotal(v int32) *ListCustomDomainsOutput {
	s.Total = &v
	return s
}

func (s *ListCustomDomainsOutput) Validate() error {
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
