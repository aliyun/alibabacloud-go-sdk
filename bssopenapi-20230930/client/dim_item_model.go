// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDimItem interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DimItem
	GetCode() *string
	SetName(v string) *DimItem
	GetName() *string
	SetPageInfo(v *DimItemPageInfo) *DimItem
	GetPageInfo() *DimItemPageInfo
	SetValues(v []*ItemValues) *DimItem
	GetValues() []*ItemValues
}

type DimItem struct {
	Code     *string          `json:"Code,omitempty" xml:"Code,omitempty"`
	Name     *string          `json:"Name,omitempty" xml:"Name,omitempty"`
	PageInfo *DimItemPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	Values   []*ItemValues    `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s DimItem) String() string {
	return dara.Prettify(s)
}

func (s DimItem) GoString() string {
	return s.String()
}

func (s *DimItem) GetCode() *string {
	return s.Code
}

func (s *DimItem) GetName() *string {
	return s.Name
}

func (s *DimItem) GetPageInfo() *DimItemPageInfo {
	return s.PageInfo
}

func (s *DimItem) GetValues() []*ItemValues {
	return s.Values
}

func (s *DimItem) SetCode(v string) *DimItem {
	s.Code = &v
	return s
}

func (s *DimItem) SetName(v string) *DimItem {
	s.Name = &v
	return s
}

func (s *DimItem) SetPageInfo(v *DimItemPageInfo) *DimItem {
	s.PageInfo = v
	return s
}

func (s *DimItem) SetValues(v []*ItemValues) *DimItem {
	s.Values = v
	return s
}

func (s *DimItem) Validate() error {
	if s.PageInfo != nil {
		if err := s.PageInfo.Validate(); err != nil {
			return err
		}
	}
	if s.Values != nil {
		for _, item := range s.Values {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DimItemPageInfo struct {
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total       *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DimItemPageInfo) String() string {
	return dara.Prettify(s)
}

func (s DimItemPageInfo) GoString() string {
	return s.String()
}

func (s *DimItemPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DimItemPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DimItemPageInfo) GetTotal() *int32 {
	return s.Total
}

func (s *DimItemPageInfo) SetCurrentPage(v int32) *DimItemPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *DimItemPageInfo) SetPageSize(v int32) *DimItemPageInfo {
	s.PageSize = &v
	return s
}

func (s *DimItemPageInfo) SetTotal(v int32) *DimItemPageInfo {
	s.Total = &v
	return s
}

func (s *DimItemPageInfo) Validate() error {
	return dara.Validate(s)
}
