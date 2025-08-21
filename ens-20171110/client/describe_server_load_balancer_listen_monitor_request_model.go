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
	// The end of the time range to query. The maximum range between StartTime and EndTime is 24 hours.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2024-05-16 16:00:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the ELB instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// lb-5rcvo1n1t3hykfhhjwjgqp****
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// The request protocol, such as http, https, or tcp.
	//
	// example:
	//
	// tcp
	Proto *string `json:"Proto,omitempty" xml:"Proto,omitempty"`
	// The beginning of the time range to query. Specify the time in the yyyy-MM-ddTHH:mm:ssZ format.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2024-05-16 15:00:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The virtual IP address (VIP) port, such as 80, 8080, or 443.
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
