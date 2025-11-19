// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListModelProxiesOutput interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*ModelProxy) *ListModelProxiesOutput
	GetItems() []*ModelProxy
	SetPageNumber(v int32) *ListModelProxiesOutput
	GetPageNumber() *int32
	SetPageSize(v int32) *ListModelProxiesOutput
	GetPageSize() *int32
	SetTotal(v int64) *ListModelProxiesOutput
	GetTotal() *int64
}

type ListModelProxiesOutput struct {
	Items      []*ModelProxy `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	PageNumber *int32        `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *int32        `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Total      *int64        `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListModelProxiesOutput) String() string {
	return dara.Prettify(s)
}

func (s ListModelProxiesOutput) GoString() string {
	return s.String()
}

func (s *ListModelProxiesOutput) GetItems() []*ModelProxy {
	return s.Items
}

func (s *ListModelProxiesOutput) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListModelProxiesOutput) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListModelProxiesOutput) GetTotal() *int64 {
	return s.Total
}

func (s *ListModelProxiesOutput) SetItems(v []*ModelProxy) *ListModelProxiesOutput {
	s.Items = v
	return s
}

func (s *ListModelProxiesOutput) SetPageNumber(v int32) *ListModelProxiesOutput {
	s.PageNumber = &v
	return s
}

func (s *ListModelProxiesOutput) SetPageSize(v int32) *ListModelProxiesOutput {
	s.PageSize = &v
	return s
}

func (s *ListModelProxiesOutput) SetTotal(v int64) *ListModelProxiesOutput {
	s.Total = &v
	return s
}

func (s *ListModelProxiesOutput) Validate() error {
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
