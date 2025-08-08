// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListToolsetsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*Toolset) *ListToolsetsResponseBody
	GetData() []*Toolset
	SetPageNumber(v int64) *ListToolsetsResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *ListToolsetsResponseBody
	GetPageSize() *int64
	SetTotalCount(v int64) *ListToolsetsResponseBody
	GetTotalCount() *int64
}

type ListToolsetsResponseBody struct {
	Data []*Toolset `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// 10
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListToolsetsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListToolsetsResponseBody) GoString() string {
	return s.String()
}

func (s *ListToolsetsResponseBody) GetData() []*Toolset {
	return s.Data
}

func (s *ListToolsetsResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListToolsetsResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListToolsetsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListToolsetsResponseBody) SetData(v []*Toolset) *ListToolsetsResponseBody {
	s.Data = v
	return s
}

func (s *ListToolsetsResponseBody) SetPageNumber(v int64) *ListToolsetsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListToolsetsResponseBody) SetPageSize(v int64) *ListToolsetsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListToolsetsResponseBody) SetTotalCount(v int64) *ListToolsetsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListToolsetsResponseBody) Validate() error {
	return dara.Validate(s)
}
