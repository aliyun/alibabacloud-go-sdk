// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBrowsersOutput interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*Browser) *ListBrowsersOutput
	GetItems() []*Browser
	SetPageNumber(v int32) *ListBrowsersOutput
	GetPageNumber() *int32
	SetPageSize(v int32) *ListBrowsersOutput
	GetPageSize() *int32
	SetTotal(v int64) *ListBrowsersOutput
	GetTotal() *int64
}

type ListBrowsersOutput struct {
	Items      []*Browser `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	PageNumber *int32     `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *int32     `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Total      *int64     `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListBrowsersOutput) String() string {
	return dara.Prettify(s)
}

func (s ListBrowsersOutput) GoString() string {
	return s.String()
}

func (s *ListBrowsersOutput) GetItems() []*Browser {
	return s.Items
}

func (s *ListBrowsersOutput) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListBrowsersOutput) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListBrowsersOutput) GetTotal() *int64 {
	return s.Total
}

func (s *ListBrowsersOutput) SetItems(v []*Browser) *ListBrowsersOutput {
	s.Items = v
	return s
}

func (s *ListBrowsersOutput) SetPageNumber(v int32) *ListBrowsersOutput {
	s.PageNumber = &v
	return s
}

func (s *ListBrowsersOutput) SetPageSize(v int32) *ListBrowsersOutput {
	s.PageSize = &v
	return s
}

func (s *ListBrowsersOutput) SetTotal(v int64) *ListBrowsersOutput {
	s.Total = &v
	return s
}

func (s *ListBrowsersOutput) Validate() error {
	return dara.Validate(s)
}
