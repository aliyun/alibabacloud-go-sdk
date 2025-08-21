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
	// The end of the time range to query. The maximum range between StartTime and EndTime is 24 hours.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2024-09-15 17:00:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the ELB instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// lb-5sc1s9zrui8lpb8u7cl4f****
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// The beginning of the time range to query. Specify the time in the yyyy-MM-ddTHH:mm:ssZ format.
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
