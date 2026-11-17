// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNatFirewallPrecheckDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeNatFirewallPrecheckDetailRequest
	GetLang() *string
	SetNatGatewayId(v string) *DescribeNatFirewallPrecheckDetailRequest
	GetNatGatewayId() *string
	SetRegionNo(v string) *DescribeNatFirewallPrecheckDetailRequest
	GetRegionNo() *string
}

type DescribeNatFirewallPrecheckDetailRequest struct {
	// The language type of the request and response messages.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The ID of the NAT gateway. Although this parameter is marked as optional in the API definition, it is required when you call this operation. If this parameter is not specified or its value does not start with `ngw-`, the API returns a 400 error (`ErrorParameters`, error code `-360103`).
	//
	// example:
	//
	// ngw-bp1okz6k7dge****
	NatGatewayId *string `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
}

func (s DescribeNatFirewallPrecheckDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeNatFirewallPrecheckDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeNatFirewallPrecheckDetailRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeNatFirewallPrecheckDetailRequest) GetNatGatewayId() *string {
	return s.NatGatewayId
}

func (s *DescribeNatFirewallPrecheckDetailRequest) GetRegionNo() *string {
	return s.RegionNo
}

func (s *DescribeNatFirewallPrecheckDetailRequest) SetLang(v string) *DescribeNatFirewallPrecheckDetailRequest {
	s.Lang = &v
	return s
}

func (s *DescribeNatFirewallPrecheckDetailRequest) SetNatGatewayId(v string) *DescribeNatFirewallPrecheckDetailRequest {
	s.NatGatewayId = &v
	return s
}

func (s *DescribeNatFirewallPrecheckDetailRequest) SetRegionNo(v string) *DescribeNatFirewallPrecheckDetailRequest {
	s.RegionNo = &v
	return s
}

func (s *DescribeNatFirewallPrecheckDetailRequest) Validate() error {
	return dara.Validate(s)
}
