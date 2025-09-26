// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCodeInterpreterSessionListOut interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*CodeInterpreterSessionOut) *CodeInterpreterSessionListOut
	GetItems() []*CodeInterpreterSessionOut
	SetPageNumber(v int32) *CodeInterpreterSessionListOut
	GetPageNumber() *int32
	SetPageSize(v int32) *CodeInterpreterSessionListOut
	GetPageSize() *int32
	SetTotal(v int64) *CodeInterpreterSessionListOut
	GetTotal() *int64
}

type CodeInterpreterSessionListOut struct {
	Items      []*CodeInterpreterSessionOut `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	PageNumber *int32                       `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *int32                       `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Total      *int64                       `json:"total,omitempty" xml:"total,omitempty"`
}

func (s CodeInterpreterSessionListOut) String() string {
	return dara.Prettify(s)
}

func (s CodeInterpreterSessionListOut) GoString() string {
	return s.String()
}

func (s *CodeInterpreterSessionListOut) GetItems() []*CodeInterpreterSessionOut {
	return s.Items
}

func (s *CodeInterpreterSessionListOut) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *CodeInterpreterSessionListOut) GetPageSize() *int32 {
	return s.PageSize
}

func (s *CodeInterpreterSessionListOut) GetTotal() *int64 {
	return s.Total
}

func (s *CodeInterpreterSessionListOut) SetItems(v []*CodeInterpreterSessionOut) *CodeInterpreterSessionListOut {
	s.Items = v
	return s
}

func (s *CodeInterpreterSessionListOut) SetPageNumber(v int32) *CodeInterpreterSessionListOut {
	s.PageNumber = &v
	return s
}

func (s *CodeInterpreterSessionListOut) SetPageSize(v int32) *CodeInterpreterSessionListOut {
	s.PageSize = &v
	return s
}

func (s *CodeInterpreterSessionListOut) SetTotal(v int64) *CodeInterpreterSessionListOut {
	s.Total = &v
	return s
}

func (s *CodeInterpreterSessionListOut) Validate() error {
	return dara.Validate(s)
}
