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
	// The direction of the traffic to which the access control policy applies.
	//
	// Valid values:
	//
	// 	- **out**: outbound traffic
	//
	// This parameter is required.
	//
	// example:
	//
	// out
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// The IP version supported by the access control policy. Valid values:
	//
	// 	- **4**: IPv4 (default)
	//
	// example:
	//
	// 4
	IpVersion *string `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	// The language of the content within the request and the response. Valid values:
	//
	// 	- **zh**: Chinese (default)
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The ID of the NAT gateway.
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
