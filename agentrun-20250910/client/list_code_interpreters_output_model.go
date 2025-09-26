// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCodeInterpretersOutput interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*CodeInterpreter) *ListCodeInterpretersOutput
	GetItems() []*CodeInterpreter
	SetPageNumber(v int32) *ListCodeInterpretersOutput
	GetPageNumber() *int32
	SetPageSize(v int32) *ListCodeInterpretersOutput
	GetPageSize() *int32
	SetTotal(v int64) *ListCodeInterpretersOutput
	GetTotal() *int64
}

type ListCodeInterpretersOutput struct {
	Items      []*CodeInterpreter `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	PageNumber *int32             `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *int32             `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Total      *int64             `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListCodeInterpretersOutput) String() string {
	return dara.Prettify(s)
}

func (s ListCodeInterpretersOutput) GoString() string {
	return s.String()
}

func (s *ListCodeInterpretersOutput) GetItems() []*CodeInterpreter {
	return s.Items
}

func (s *ListCodeInterpretersOutput) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCodeInterpretersOutput) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCodeInterpretersOutput) GetTotal() *int64 {
	return s.Total
}

func (s *ListCodeInterpretersOutput) SetItems(v []*CodeInterpreter) *ListCodeInterpretersOutput {
	s.Items = v
	return s
}

func (s *ListCodeInterpretersOutput) SetPageNumber(v int32) *ListCodeInterpretersOutput {
	s.PageNumber = &v
	return s
}

func (s *ListCodeInterpretersOutput) SetPageSize(v int32) *ListCodeInterpretersOutput {
	s.PageSize = &v
	return s
}

func (s *ListCodeInterpretersOutput) SetTotal(v int64) *ListCodeInterpretersOutput {
	s.Total = &v
	return s
}

func (s *ListCodeInterpretersOutput) Validate() error {
	return dara.Validate(s)
}
