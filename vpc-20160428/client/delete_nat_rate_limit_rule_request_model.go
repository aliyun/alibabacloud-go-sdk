// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNatRateLimitRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteNatRateLimitRuleRequest
	GetClientToken() *string
	SetNatGatewayId(v string) *DeleteNatRateLimitRuleRequest
	GetNatGatewayId() *string
	SetRegionId(v string) *DeleteNatRateLimitRuleRequest
	GetRegionId() *string
	SetRuleIds(v []*string) *DeleteNatRateLimitRuleRequest
	GetRuleIds() []*string
}

type DeleteNatRateLimitRuleRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must make sure that it is unique among different requests. The client token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- of each API request may be different.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-001****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the NAT gateway to which the rate limit rule belongs.
	//
	// example:
	//
	// ngw-bp1uewa15k4iy5770****
	NatGatewayId *string `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
	// The region ID of the NAT gateway.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to obtain the region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The IDs of the rate limit rules to be deleted.
	RuleIds []*string `json:"RuleIds,omitempty" xml:"RuleIds,omitempty" type:"Repeated"`
}

func (s DeleteNatRateLimitRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteNatRateLimitRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteNatRateLimitRuleRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteNatRateLimitRuleRequest) GetNatGatewayId() *string {
	return s.NatGatewayId
}

func (s *DeleteNatRateLimitRuleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteNatRateLimitRuleRequest) GetRuleIds() []*string {
	return s.RuleIds
}

func (s *DeleteNatRateLimitRuleRequest) SetClientToken(v string) *DeleteNatRateLimitRuleRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteNatRateLimitRuleRequest) SetNatGatewayId(v string) *DeleteNatRateLimitRuleRequest {
	s.NatGatewayId = &v
	return s
}

func (s *DeleteNatRateLimitRuleRequest) SetRegionId(v string) *DeleteNatRateLimitRuleRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteNatRateLimitRuleRequest) SetRuleIds(v []*string) *DeleteNatRateLimitRuleRequest {
	s.RuleIds = v
	return s
}

func (s *DeleteNatRateLimitRuleRequest) Validate() error {
	return dara.Validate(s)
}
