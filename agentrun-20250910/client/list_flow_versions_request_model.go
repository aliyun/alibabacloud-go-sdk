// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFlowVersionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListFlowVersionsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListFlowVersionsRequest
	GetPageSize() *int32
}

type ListFlowVersionsRequest struct {
	// 当前页码，从1开始计数
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// 每页返回的记录数量
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListFlowVersionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListFlowVersionsRequest) GoString() string {
	return s.String()
}

func (s *ListFlowVersionsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListFlowVersionsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListFlowVersionsRequest) SetPageNumber(v int32) *ListFlowVersionsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListFlowVersionsRequest) SetPageSize(v int32) *ListFlowVersionsRequest {
	s.PageSize = &v
	return s
}

func (s *ListFlowVersionsRequest) Validate() error {
	return dara.Validate(s)
}
