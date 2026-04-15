// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFlowEndpointsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListFlowEndpointsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListFlowEndpointsRequest
	GetPageSize() *int32
}

type ListFlowEndpointsRequest struct {
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

func (s ListFlowEndpointsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListFlowEndpointsRequest) GoString() string {
	return s.String()
}

func (s *ListFlowEndpointsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListFlowEndpointsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListFlowEndpointsRequest) SetPageNumber(v int32) *ListFlowEndpointsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListFlowEndpointsRequest) SetPageSize(v int32) *ListFlowEndpointsRequest {
	s.PageSize = &v
	return s
}

func (s *ListFlowEndpointsRequest) Validate() error {
	return dara.Validate(s)
}
