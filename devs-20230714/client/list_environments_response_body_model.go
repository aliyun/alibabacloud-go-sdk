// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEnvironmentsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*Environment) *ListEnvironmentsResponseBody
	GetData() []*Environment
	SetPageNumber(v int64) *ListEnvironmentsResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *ListEnvironmentsResponseBody
	GetPageSize() *int64
	SetTotalCount(v int64) *ListEnvironmentsResponseBody
	GetTotalCount() *int64
}

type ListEnvironmentsResponseBody struct {
	Data []*Environment `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 50
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// 100
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListEnvironmentsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListEnvironmentsResponseBody) GoString() string {
	return s.String()
}

func (s *ListEnvironmentsResponseBody) GetData() []*Environment {
	return s.Data
}

func (s *ListEnvironmentsResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListEnvironmentsResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListEnvironmentsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListEnvironmentsResponseBody) SetData(v []*Environment) *ListEnvironmentsResponseBody {
	s.Data = v
	return s
}

func (s *ListEnvironmentsResponseBody) SetPageNumber(v int64) *ListEnvironmentsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListEnvironmentsResponseBody) SetPageSize(v int64) *ListEnvironmentsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListEnvironmentsResponseBody) SetTotalCount(v int64) *ListEnvironmentsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListEnvironmentsResponseBody) Validate() error {
	return dara.Validate(s)
}
