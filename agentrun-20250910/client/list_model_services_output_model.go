// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListModelServicesOutput interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*ModelService) *ListModelServicesOutput
	GetItems() []*ModelService
	SetPageNumber(v int32) *ListModelServicesOutput
	GetPageNumber() *int32
	SetPageSize(v int32) *ListModelServicesOutput
	GetPageSize() *int32
	SetTotal(v int64) *ListModelServicesOutput
	GetTotal() *int64
}

type ListModelServicesOutput struct {
	Items      []*ModelService `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	PageNumber *int32          `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *int32          `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Total      *int64          `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListModelServicesOutput) String() string {
	return dara.Prettify(s)
}

func (s ListModelServicesOutput) GoString() string {
	return s.String()
}

func (s *ListModelServicesOutput) GetItems() []*ModelService {
	return s.Items
}

func (s *ListModelServicesOutput) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListModelServicesOutput) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListModelServicesOutput) GetTotal() *int64 {
	return s.Total
}

func (s *ListModelServicesOutput) SetItems(v []*ModelService) *ListModelServicesOutput {
	s.Items = v
	return s
}

func (s *ListModelServicesOutput) SetPageNumber(v int32) *ListModelServicesOutput {
	s.PageNumber = &v
	return s
}

func (s *ListModelServicesOutput) SetPageSize(v int32) *ListModelServicesOutput {
	s.PageSize = &v
	return s
}

func (s *ListModelServicesOutput) SetTotal(v int64) *ListModelServicesOutput {
	s.Total = &v
	return s
}

func (s *ListModelServicesOutput) Validate() error {
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
