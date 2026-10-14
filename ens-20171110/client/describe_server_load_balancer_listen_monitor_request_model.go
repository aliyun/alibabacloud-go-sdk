// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeServerLoadBalancerListenMonitorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DescribeServerLoadBalancerListenMonitorRequest
	GetEndTime() *string
	SetLoadBalancerId(v string) *DescribeServerLoadBalancerListenMonitorRequest
	GetLoadBalancerId() *string
	SetProto(v string) *DescribeServerLoadBalancerListenMonitorRequest
	GetProto() *string
	SetStartTime(v string) *DescribeServerLoadBalancerListenMonitorRequest
	GetStartTime() *string
	SetVPort(v string) *DescribeServerLoadBalancerListenMonitorRequest
	GetVPort() *string
}

type DescribeServerLoadBalancerListenMonitorRequest struct {
	// The end time. The maximum interval between StartTime and EndTime cannot exceed 24 hours.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2024-05-16 16:00:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The load balancing instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// lb-5rcvo1n1t3hykfhhjwjgqp****
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// The request protocol, such as http, https, and tcp.
	//
	// example:
	//
	// tcp
	Proto *string `json:"Proto,omitempty" xml:"Proto,omitempty"`
	// The start time in the format of yyyy-MM-dd HH:mm:ss.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2024-05-16 15:00:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The VIP port, such as 80, 8080, and 443.
	//
	// example:
	//
	// 80
	VPort *string `json:"VPort,omitempty" xml:"VPort,omitempty"`
}

func (s DescribeServerLoadBalancerListenMonitorRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeServerLoadBalancerListenMonitorRequest) GoString() string {
	return s.String()
}

func (s *DescribeServerLoadBalancerListenMonitorRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeServerLoadBalancerListenMonitorRequest) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *DescribeServerLoadBalancerListenMonitorRequest) GetProto() *string {
	return s.Proto
}

func (s *DescribeServerLoadBalancerListenMonitorRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeServerLoadBalancerListenMonitorRequest) GetVPort() *string {
	return s.VPort
}

func (s *DescribeServerLoadBalancerListenMonitorRequest) SetEndTime(v string) *DescribeServerLoadBalancerListenMonitorRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeServerLoadBalancerListenMonitorRequest) SetLoadBalancerId(v string) *DescribeServerLoadBalancerListenMonitorRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *DescribeServerLoadBalancerListenMonitorRequest) SetProto(v string) *DescribeServerLoadBalancerListenMonitorRequest {
	s.Proto = &v
	return s
}

func (s *DescribeServerLoadBalancerListenMonitorRequest) SetStartTime(v string) *DescribeServerLoadBalancerListenMonitorRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeServerLoadBalancerListenMonitorRequest) SetVPort(v string) *DescribeServerLoadBalancerListenMonitorRequest {
	s.VPort = &v
	return s
}

func (s *DescribeServerLoadBalancerListenMonitorRequest) Validate() error {
	return dara.Validate(s)
}
