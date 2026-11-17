// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyNatRateLimitRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBandwidthLimit(v string) *ModifyNatRateLimitRuleRequest
	GetBandwidthLimit() *string
	SetClientToken(v string) *ModifyNatRateLimitRuleRequest
	GetClientToken() *string
	SetDryRun(v string) *ModifyNatRateLimitRuleRequest
	GetDryRun() *string
	SetNatGatewayId(v string) *ModifyNatRateLimitRuleRequest
	GetNatGatewayId() *string
	SetRegionId(v string) *ModifyNatRateLimitRuleRequest
	GetRegionId() *string
	SetRuleDescription(v string) *ModifyNatRateLimitRuleRequest
	GetRuleDescription() *string
	SetRuleId(v string) *ModifyNatRateLimitRuleRequest
	GetRuleId() *string
	SetRuleName(v string) *ModifyNatRateLimitRuleRequest
	GetRuleName() *string
	SetSourceCIDRs(v []*string) *ModifyNatRateLimitRuleRequest
	GetSourceCIDRs() []*string
}

type ModifyNatRateLimitRuleRequest struct {
	// The new bandwidth limit. Unit: Mbit/s.
	//
	// example:
	//
	// 100
	BandwidthLimit *string `json:"BandwidthLimit,omitempty" xml:"BandwidthLimit,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must ensure that it is unique among different requests. The client token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// > If you do not specify this parameter, the system automatically uses the RequestId of the API request as the client token. The RequestId of each API request may be different.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-0016e04115b
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// - **true**: performs a dry run without querying the resource status. The system checks whether your AccessKey pair is valid, the authorization of the Resource Access Management (RAM) user, and whether the required parameters are specified. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// - **false*	- (default): sends a normal request. If the request passes the dry run, a 2xx HTTP status code is returned and the resource status is queried.
	//
	// example:
	//
	// false
	DryRun *string `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the NAT gateway to which the rate limiting rule belongs.
	//
	// example:
	//
	// ngw-bp1b0lic8uz4r6vf2****
	NatGatewayId *string `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
	// The region ID of the NAT gateway.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query region IDs.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The new description of the rate limiting rule.
	//
	// example:
	//
	// test
	RuleDescription *string `json:"RuleDescription,omitempty" xml:"RuleDescription,omitempty"`
	// The ID of the rate limiting rule.
	//
	// example:
	//
	// limit-rule-****
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The new name of the rate limiting rule.
	//
	// example:
	//
	// test-rule
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The new list of source CIDR blocks for the rate limiting rule.
	SourceCIDRs []*string `json:"SourceCIDRs,omitempty" xml:"SourceCIDRs,omitempty" type:"Repeated"`
}

func (s ModifyNatRateLimitRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyNatRateLimitRuleRequest) GoString() string {
	return s.String()
}

func (s *ModifyNatRateLimitRuleRequest) GetBandwidthLimit() *string {
	return s.BandwidthLimit
}

func (s *ModifyNatRateLimitRuleRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyNatRateLimitRuleRequest) GetDryRun() *string {
	return s.DryRun
}

func (s *ModifyNatRateLimitRuleRequest) GetNatGatewayId() *string {
	return s.NatGatewayId
}

func (s *ModifyNatRateLimitRuleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyNatRateLimitRuleRequest) GetRuleDescription() *string {
	return s.RuleDescription
}

func (s *ModifyNatRateLimitRuleRequest) GetRuleId() *string {
	return s.RuleId
}

func (s *ModifyNatRateLimitRuleRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *ModifyNatRateLimitRuleRequest) GetSourceCIDRs() []*string {
	return s.SourceCIDRs
}

func (s *ModifyNatRateLimitRuleRequest) SetBandwidthLimit(v string) *ModifyNatRateLimitRuleRequest {
	s.BandwidthLimit = &v
	return s
}

func (s *ModifyNatRateLimitRuleRequest) SetClientToken(v string) *ModifyNatRateLimitRuleRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyNatRateLimitRuleRequest) SetDryRun(v string) *ModifyNatRateLimitRuleRequest {
	s.DryRun = &v
	return s
}

func (s *ModifyNatRateLimitRuleRequest) SetNatGatewayId(v string) *ModifyNatRateLimitRuleRequest {
	s.NatGatewayId = &v
	return s
}

func (s *ModifyNatRateLimitRuleRequest) SetRegionId(v string) *ModifyNatRateLimitRuleRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyNatRateLimitRuleRequest) SetRuleDescription(v string) *ModifyNatRateLimitRuleRequest {
	s.RuleDescription = &v
	return s
}

func (s *ModifyNatRateLimitRuleRequest) SetRuleId(v string) *ModifyNatRateLimitRuleRequest {
	s.RuleId = &v
	return s
}

func (s *ModifyNatRateLimitRuleRequest) SetRuleName(v string) *ModifyNatRateLimitRuleRequest {
	s.RuleName = &v
	return s
}

func (s *ModifyNatRateLimitRuleRequest) SetSourceCIDRs(v []*string) *ModifyNatRateLimitRuleRequest {
	s.SourceCIDRs = v
	return s
}

func (s *ModifyNatRateLimitRuleRequest) Validate() error {
	return dara.Validate(s)
}
