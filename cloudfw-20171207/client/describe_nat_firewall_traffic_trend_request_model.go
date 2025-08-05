// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNatFirewallTrafficTrendRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *DescribeNatFirewallTrafficTrendRequest
	GetEndTime() *int64
	SetInterval(v int64) *DescribeNatFirewallTrafficTrendRequest
	GetInterval() *int64
	SetNatGatewayId(v string) *DescribeNatFirewallTrafficTrendRequest
	GetNatGatewayId() *string
	SetSrcPrivateIP(v string) *DescribeNatFirewallTrafficTrendRequest
	GetSrcPrivateIP() *string
	SetSrcPublicIP(v string) *DescribeNatFirewallTrafficTrendRequest
	GetSrcPublicIP() *string
	SetStartTime(v int64) *DescribeNatFirewallTrafficTrendRequest
	GetStartTime() *int64
}

type DescribeNatFirewallTrafficTrendRequest struct {
	// The end of the time range to query. The value is a UNIX timestamp that is accurate to seconds.
	//
	// example:
	//
	// 1739330580
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The time interval between the data entries to return. Unit: seconds. Valid values:
	//
	// 	- **60**: 1 minute
	//
	// 	- **1800**: 30 minutes
	//
	// example:
	//
	// 60
	Interval *int64 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The ID of the NAT gateway.
	//
	// example:
	//
	// ngw-xxxxxx
	NatGatewayId *string `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
	// The private IP address of the source.
	//
	// example:
	//
	// 10.100.134.60
	SrcPrivateIP *string `json:"SrcPrivateIP,omitempty" xml:"SrcPrivateIP,omitempty"`
	// The public IP address of the source.
	//
	// example:
	//
	// 47.112.210.136
	SrcPublicIP *string `json:"SrcPublicIP,omitempty" xml:"SrcPublicIP,omitempty"`
	// The beginning of the time range to query. The value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1739326980
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeNatFirewallTrafficTrendRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeNatFirewallTrafficTrendRequest) GoString() string {
	return s.String()
}

func (s *DescribeNatFirewallTrafficTrendRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeNatFirewallTrafficTrendRequest) GetInterval() *int64 {
	return s.Interval
}

func (s *DescribeNatFirewallTrafficTrendRequest) GetNatGatewayId() *string {
	return s.NatGatewayId
}

func (s *DescribeNatFirewallTrafficTrendRequest) GetSrcPrivateIP() *string {
	return s.SrcPrivateIP
}

func (s *DescribeNatFirewallTrafficTrendRequest) GetSrcPublicIP() *string {
	return s.SrcPublicIP
}

func (s *DescribeNatFirewallTrafficTrendRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeNatFirewallTrafficTrendRequest) SetEndTime(v int64) *DescribeNatFirewallTrafficTrendRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeNatFirewallTrafficTrendRequest) SetInterval(v int64) *DescribeNatFirewallTrafficTrendRequest {
	s.Interval = &v
	return s
}

func (s *DescribeNatFirewallTrafficTrendRequest) SetNatGatewayId(v string) *DescribeNatFirewallTrafficTrendRequest {
	s.NatGatewayId = &v
	return s
}

func (s *DescribeNatFirewallTrafficTrendRequest) SetSrcPrivateIP(v string) *DescribeNatFirewallTrafficTrendRequest {
	s.SrcPrivateIP = &v
	return s
}

func (s *DescribeNatFirewallTrafficTrendRequest) SetSrcPublicIP(v string) *DescribeNatFirewallTrafficTrendRequest {
	s.SrcPublicIP = &v
	return s
}

func (s *DescribeNatFirewallTrafficTrendRequest) SetStartTime(v int64) *DescribeNatFirewallTrafficTrendRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeNatFirewallTrafficTrendRequest) Validate() error {
	return dara.Validate(s)
}
