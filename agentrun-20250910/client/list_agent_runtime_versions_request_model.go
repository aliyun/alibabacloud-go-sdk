// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgentRuntimeVersionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListAgentRuntimeVersionsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListAgentRuntimeVersionsRequest
	GetPageSize() *int32
}

type ListAgentRuntimeVersionsRequest struct {
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

func (s ListAgentRuntimeVersionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAgentRuntimeVersionsRequest) GoString() string {
	return s.String()
}

func (s *ListAgentRuntimeVersionsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListAgentRuntimeVersionsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAgentRuntimeVersionsRequest) SetPageNumber(v int32) *ListAgentRuntimeVersionsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAgentRuntimeVersionsRequest) SetPageSize(v int32) *ListAgentRuntimeVersionsRequest {
	s.PageSize = &v
	return s
}

func (s *ListAgentRuntimeVersionsRequest) Validate() error {
	return dara.Validate(s)
}
