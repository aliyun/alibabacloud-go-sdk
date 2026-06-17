// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNatFirewallPolicyPriorUsedRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirection(v string) *DescribeNatFirewallPolicyPriorUsedRequest
	GetDirection() *string
	SetIpVersion(v string) *DescribeNatFirewallPolicyPriorUsedRequest
	GetIpVersion() *string
	SetLang(v string) *DescribeNatFirewallPolicyPriorUsedRequest
	GetLang() *string
	SetNatGatewayId(v string) *DescribeNatFirewallPolicyPriorUsedRequest
	GetNatGatewayId() *string
}

type DescribeNatFirewallPolicyPriorUsedRequest struct {
	// The traffic direction of the access control policy.
	//
	// Valid value:
	//
	// - **out**: outbound traffic.
	//
	// This parameter is required.
	//
	// example:
	//
	// out
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// The IP version. Valid value:
	//
	// - **4*	- (default): IPv4
	//
	// example:
	//
	// 4
	IpVersion *string `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	// The language of the request and response. Valid values:
	//
	// - **zh*	- (default): Chinese.
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The ID of the NAT Gateway.
	//
	// This parameter is required.
	//
	// example:
	//
	// ngw-xxxxxx
	NatGatewayId *string `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
}

func (s DescribeNatFirewallPolicyPriorUsedRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeNatFirewallPolicyPriorUsedRequest) GoString() string {
	return s.String()
}

func (s *DescribeNatFirewallPolicyPriorUsedRequest) GetDirection() *string {
	return s.Direction
}

func (s *DescribeNatFirewallPolicyPriorUsedRequest) GetIpVersion() *string {
	return s.IpVersion
}

func (s *DescribeNatFirewallPolicyPriorUsedRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeNatFirewallPolicyPriorUsedRequest) GetNatGatewayId() *string {
	return s.NatGatewayId
}

func (s *DescribeNatFirewallPolicyPriorUsedRequest) SetDirection(v string) *DescribeNatFirewallPolicyPriorUsedRequest {
	s.Direction = &v
	return s
}

func (s *DescribeNatFirewallPolicyPriorUsedRequest) SetIpVersion(v string) *DescribeNatFirewallPolicyPriorUsedRequest {
	s.IpVersion = &v
	return s
}

func (s *DescribeNatFirewallPolicyPriorUsedRequest) SetLang(v string) *DescribeNatFirewallPolicyPriorUsedRequest {
	s.Lang = &v
	return s
}

func (s *DescribeNatFirewallPolicyPriorUsedRequest) SetNatGatewayId(v string) *DescribeNatFirewallPolicyPriorUsedRequest {
	s.NatGatewayId = &v
	return s
}

func (s *DescribeNatFirewallPolicyPriorUsedRequest) Validate() error {
	return dara.Validate(s)
}
