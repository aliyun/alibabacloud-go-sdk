// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLoadBalancerListenersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *DescribeLoadBalancerListenersRequest
	GetDescription() *string
	SetListenerPort(v int32) *DescribeLoadBalancerListenersRequest
	GetListenerPort() *int32
	SetLoadBalancerId(v string) *DescribeLoadBalancerListenersRequest
	GetLoadBalancerId() *string
	SetPageNumber(v int32) *DescribeLoadBalancerListenersRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeLoadBalancerListenersRequest
	GetPageSize() *int32
}

type DescribeLoadBalancerListenersRequest struct {
	// The description of the image.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The listener port.
	//
	// example:
	//
	// 80
	ListenerPort *int32 `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	// The ID of the ELB instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// lb-5s7crik3yo3p5****
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeLoadBalancerListenersRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLoadBalancerListenersRequest) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancerListenersRequest) GetDescription() *string {
	return s.Description
}

func (s *DescribeLoadBalancerListenersRequest) GetListenerPort() *int32 {
	return s.ListenerPort
}

func (s *DescribeLoadBalancerListenersRequest) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *DescribeLoadBalancerListenersRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeLoadBalancerListenersRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeLoadBalancerListenersRequest) SetDescription(v string) *DescribeLoadBalancerListenersRequest {
	s.Description = &v
	return s
}

func (s *DescribeLoadBalancerListenersRequest) SetListenerPort(v int32) *DescribeLoadBalancerListenersRequest {
	s.ListenerPort = &v
	return s
}

func (s *DescribeLoadBalancerListenersRequest) SetLoadBalancerId(v string) *DescribeLoadBalancerListenersRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *DescribeLoadBalancerListenersRequest) SetPageNumber(v int32) *DescribeLoadBalancerListenersRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeLoadBalancerListenersRequest) SetPageSize(v int32) *DescribeLoadBalancerListenersRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeLoadBalancerListenersRequest) Validate() error {
	return dara.Validate(s)
}
