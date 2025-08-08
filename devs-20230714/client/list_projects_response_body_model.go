// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProjectsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*Project) *ListProjectsResponseBody
	GetData() []*Project
	SetPageNumber(v int64) *ListProjectsResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *ListProjectsResponseBody
	GetPageSize() *int64
	SetTotalCount(v int64) *ListProjectsResponseBody
	GetTotalCount() *int64
}

type ListProjectsResponseBody struct {
	Data []*Project `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 100
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// 1
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListProjectsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListProjectsResponseBody) GoString() string {
	return s.String()
}

func (s *ListProjectsResponseBody) GetData() []*Project {
	return s.Data
}

func (s *ListProjectsResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListProjectsResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListProjectsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListProjectsResponseBody) SetData(v []*Project) *ListProjectsResponseBody {
	s.Data = v
	return s
}

func (s *ListProjectsResponseBody) SetPageNumber(v int64) *ListProjectsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListProjectsResponseBody) SetPageSize(v int64) *ListProjectsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListProjectsResponseBody) SetTotalCount(v int64) *ListProjectsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListProjectsResponseBody) Validate() error {
	return dara.Validate(s)
}
