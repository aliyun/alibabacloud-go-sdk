// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMessageAppRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNum(v int32) *ListMessageAppRequest
	GetPageNum() *int32
	SetPageSize(v int32) *ListMessageAppRequest
	GetPageSize() *int32
	SetSortType(v int32) *ListMessageAppRequest
	GetSortType() *int32
}

type ListMessageAppRequest struct {
	// The number of the page to return. Default value: 1. Valid values: 1 to 100000.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of applications to return on each page. Default value: 20. Valid values: 1 to 50.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The sort order. Valid values:
	//
	// 	- 0: ascending order by time
	//
	// 	- 1: descending order by time
	//
	// example:
	//
	// 1
	SortType *int32 `json:"SortType,omitempty" xml:"SortType,omitempty"`
}

func (s ListMessageAppRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMessageAppRequest) GoString() string {
	return s.String()
}

func (s *ListMessageAppRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *ListMessageAppRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListMessageAppRequest) GetSortType() *int32 {
	return s.SortType
}

func (s *ListMessageAppRequest) SetPageNum(v int32) *ListMessageAppRequest {
	s.PageNum = &v
	return s
}

func (s *ListMessageAppRequest) SetPageSize(v int32) *ListMessageAppRequest {
	s.PageSize = &v
	return s
}

func (s *ListMessageAppRequest) SetSortType(v int32) *ListMessageAppRequest {
	s.SortType = &v
	return s
}

func (s *ListMessageAppRequest) Validate() error {
	return dara.Validate(s)
}
