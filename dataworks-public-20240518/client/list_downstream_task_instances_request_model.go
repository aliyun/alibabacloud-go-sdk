// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDownstreamTaskInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *ListDownstreamTaskInstancesRequest
	GetId() *int64
	SetPageNumber(v int32) *ListDownstreamTaskInstancesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListDownstreamTaskInstancesRequest
	GetPageSize() *int32
}

type ListDownstreamTaskInstancesRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The page number. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListDownstreamTaskInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDownstreamTaskInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListDownstreamTaskInstancesRequest) GetId() *int64 {
	return s.Id
}

func (s *ListDownstreamTaskInstancesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDownstreamTaskInstancesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDownstreamTaskInstancesRequest) SetId(v int64) *ListDownstreamTaskInstancesRequest {
	s.Id = &v
	return s
}

func (s *ListDownstreamTaskInstancesRequest) SetPageNumber(v int32) *ListDownstreamTaskInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDownstreamTaskInstancesRequest) SetPageSize(v int32) *ListDownstreamTaskInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *ListDownstreamTaskInstancesRequest) Validate() error {
	return dara.Validate(s)
}
