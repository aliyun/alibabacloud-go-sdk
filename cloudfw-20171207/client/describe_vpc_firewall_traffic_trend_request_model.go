// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpcFirewallTrafficTrendRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DescribeVpcFirewallTrafficTrendRequest
	GetEndTime() *string
	SetLang(v string) *DescribeVpcFirewallTrafficTrendRequest
	GetLang() *string
	SetPeerVpcId(v string) *DescribeVpcFirewallTrafficTrendRequest
	GetPeerVpcId() *string
	SetPrivateIP(v string) *DescribeVpcFirewallTrafficTrendRequest
	GetPrivateIP() *string
	SetStartTime(v string) *DescribeVpcFirewallTrafficTrendRequest
	GetStartTime() *string
	SetVpcId(v string) *DescribeVpcFirewallTrafficTrendRequest
	GetVpcId() *string
}

type DescribeVpcFirewallTrafficTrendRequest struct {
	// The end time. The value is a UNIX timestamp. Unit: seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1767493189
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The language of the request and response. Valid values:
	//
	// - **zh*	- (default): Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The instance ID of the peer VPC instance.
	//
	// example:
	//
	// vpc-j6c7mscdg4hfqanfi****
	PeerVpcId *string `json:"PeerVpcId,omitempty" xml:"PeerVpcId,omitempty"`
	// The private IP address.
	//
	// example:
	//
	// 10.21.186.XXX
	PrivateIP *string `json:"PrivateIP,omitempty" xml:"PrivateIP,omitempty"`
	// The start time. The value is a UNIX timestamp. Unit: seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1780626107
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The instance ID of the VPC-connected instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-j6cnofg4lg1ujx3xj****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeVpcFirewallTrafficTrendRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcFirewallTrafficTrendRequest) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallTrafficTrendRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeVpcFirewallTrafficTrendRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeVpcFirewallTrafficTrendRequest) GetPeerVpcId() *string {
	return s.PeerVpcId
}

func (s *DescribeVpcFirewallTrafficTrendRequest) GetPrivateIP() *string {
	return s.PrivateIP
}

func (s *DescribeVpcFirewallTrafficTrendRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeVpcFirewallTrafficTrendRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeVpcFirewallTrafficTrendRequest) SetEndTime(v string) *DescribeVpcFirewallTrafficTrendRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeVpcFirewallTrafficTrendRequest) SetLang(v string) *DescribeVpcFirewallTrafficTrendRequest {
	s.Lang = &v
	return s
}

func (s *DescribeVpcFirewallTrafficTrendRequest) SetPeerVpcId(v string) *DescribeVpcFirewallTrafficTrendRequest {
	s.PeerVpcId = &v
	return s
}

func (s *DescribeVpcFirewallTrafficTrendRequest) SetPrivateIP(v string) *DescribeVpcFirewallTrafficTrendRequest {
	s.PrivateIP = &v
	return s
}

func (s *DescribeVpcFirewallTrafficTrendRequest) SetStartTime(v string) *DescribeVpcFirewallTrafficTrendRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeVpcFirewallTrafficTrendRequest) SetVpcId(v string) *DescribeVpcFirewallTrafficTrendRequest {
	s.VpcId = &v
	return s
}

func (s *DescribeVpcFirewallTrafficTrendRequest) Validate() error {
	return dara.Validate(s)
}
