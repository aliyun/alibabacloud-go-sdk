// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUpstreamTaskInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *ListUpstreamTaskInstancesRequest
	GetId() *int64
	SetPageNumber(v int32) *ListUpstreamTaskInstancesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListUpstreamTaskInstancesRequest
	GetPageSize() *int32
}

type ListUpstreamTaskInstancesRequest struct {
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

func (s ListUpstreamTaskInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListUpstreamTaskInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListUpstreamTaskInstancesRequest) GetId() *int64 {
	return s.Id
}

func (s *ListUpstreamTaskInstancesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListUpstreamTaskInstancesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListUpstreamTaskInstancesRequest) SetId(v int64) *ListUpstreamTaskInstancesRequest {
	s.Id = &v
	return s
}

func (s *ListUpstreamTaskInstancesRequest) SetPageNumber(v int32) *ListUpstreamTaskInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListUpstreamTaskInstancesRequest) SetPageSize(v int32) *ListUpstreamTaskInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *ListUpstreamTaskInstancesRequest) Validate() error {
	return dara.Validate(s)
}
