// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeServerLoadBalancerMonitorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DescribeServerLoadBalancerMonitorRequest
	GetEndTime() *string
	SetLoadBalancerId(v string) *DescribeServerLoadBalancerMonitorRequest
	GetLoadBalancerId() *string
	SetStartTime(v string) *DescribeServerLoadBalancerMonitorRequest
	GetStartTime() *string
}

type DescribeServerLoadBalancerMonitorRequest struct {
	// The end time. The maximum interval between StartTime and EndTime cannot exceed 24 hours.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2024-09-15 17:00:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The load balancing instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// lb-5sc1s9zrui8lpb8u7cl4f****
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// The start time in the format of yyyy-MM-dd HH:mm:ss.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2024-09-15 16:00:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeServerLoadBalancerMonitorRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeServerLoadBalancerMonitorRequest) GoString() string {
	return s.String()
}

func (s *DescribeServerLoadBalancerMonitorRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeServerLoadBalancerMonitorRequest) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *DescribeServerLoadBalancerMonitorRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeServerLoadBalancerMonitorRequest) SetEndTime(v string) *DescribeServerLoadBalancerMonitorRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeServerLoadBalancerMonitorRequest) SetLoadBalancerId(v string) *DescribeServerLoadBalancerMonitorRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *DescribeServerLoadBalancerMonitorRequest) SetStartTime(v string) *DescribeServerLoadBalancerMonitorRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeServerLoadBalancerMonitorRequest) Validate() error {
	return dara.Validate(s)
}
