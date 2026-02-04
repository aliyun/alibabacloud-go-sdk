// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEventTypesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListEventTypesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListEventTypesRequest
	GetPageSize() *int32
}

type ListEventTypesRequest struct {
	// 查询分页页码（从1开始计数）
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 分页大小（从1开始计数）
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListEventTypesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListEventTypesRequest) GoString() string {
	return s.String()
}

func (s *ListEventTypesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListEventTypesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListEventTypesRequest) SetPageNumber(v int32) *ListEventTypesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListEventTypesRequest) SetPageSize(v int32) *ListEventTypesRequest {
	s.PageSize = &v
	return s
}

func (s *ListEventTypesRequest) Validate() error {
	return dara.Validate(s)
}
