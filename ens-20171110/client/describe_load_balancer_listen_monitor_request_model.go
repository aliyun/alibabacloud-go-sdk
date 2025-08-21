// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLoadBalancerListenMonitorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DescribeLoadBalancerListenMonitorRequest
	GetEndTime() *string
	SetLoadBalancerId(v string) *DescribeLoadBalancerListenMonitorRequest
	GetLoadBalancerId() *string
	SetProto(v string) *DescribeLoadBalancerListenMonitorRequest
	GetProto() *string
	SetStartTime(v string) *DescribeLoadBalancerListenMonitorRequest
	GetStartTime() *string
	SetVPort(v string) *DescribeLoadBalancerListenMonitorRequest
	GetVPort() *string
}

type DescribeLoadBalancerListenMonitorRequest struct {
	// The end of the time range to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2024-01-30 08:00:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the ELB instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// lb-5q73cv04zeyh43lh74lp4****
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// The network protocol, such as tcp or udp.
	//
	// example:
	//
	// tcp
	Proto *string `json:"Proto,omitempty" xml:"Proto,omitempty"`
	// The beginning of the time range to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2024-01-15 16:00:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The virtual IP address (VIP) port of the ELB instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// 80
	VPort *string `json:"VPort,omitempty" xml:"VPort,omitempty"`
}

func (s DescribeLoadBalancerListenMonitorRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLoadBalancerListenMonitorRequest) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancerListenMonitorRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeLoadBalancerListenMonitorRequest) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *DescribeLoadBalancerListenMonitorRequest) GetProto() *string {
	return s.Proto
}

func (s *DescribeLoadBalancerListenMonitorRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeLoadBalancerListenMonitorRequest) GetVPort() *string {
	return s.VPort
}

func (s *DescribeLoadBalancerListenMonitorRequest) SetEndTime(v string) *DescribeLoadBalancerListenMonitorRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeLoadBalancerListenMonitorRequest) SetLoadBalancerId(v string) *DescribeLoadBalancerListenMonitorRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *DescribeLoadBalancerListenMonitorRequest) SetProto(v string) *DescribeLoadBalancerListenMonitorRequest {
	s.Proto = &v
	return s
}

func (s *DescribeLoadBalancerListenMonitorRequest) SetStartTime(v string) *DescribeLoadBalancerListenMonitorRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeLoadBalancerListenMonitorRequest) SetVPort(v string) *DescribeLoadBalancerListenMonitorRequest {
	s.VPort = &v
	return s
}

func (s *DescribeLoadBalancerListenMonitorRequest) Validate() error {
	return dara.Validate(s)
}
