// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLoadBalancerRegionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListLoadBalancerRegionsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListLoadBalancerRegionsRequest
	GetPageSize() *int32
}

type ListLoadBalancerRegionsRequest struct {
	// Page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Page size.
	//
	// example:
	//
	// 1024
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListLoadBalancerRegionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListLoadBalancerRegionsRequest) GoString() string {
	return s.String()
}

func (s *ListLoadBalancerRegionsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListLoadBalancerRegionsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListLoadBalancerRegionsRequest) SetPageNumber(v int32) *ListLoadBalancerRegionsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListLoadBalancerRegionsRequest) SetPageSize(v int32) *ListLoadBalancerRegionsRequest {
	s.PageSize = &v
	return s
}

func (s *ListLoadBalancerRegionsRequest) Validate() error {
	return dara.Validate(s)
}
