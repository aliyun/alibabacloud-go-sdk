// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFirewallTemplateRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteFirewallTemplateRulesRequest
	GetClientToken() *string
	SetFirewallTemplateId(v string) *DeleteFirewallTemplateRulesRequest
	GetFirewallTemplateId() *string
	SetFirewallTemplateRuleId(v []*string) *DeleteFirewallTemplateRulesRequest
	GetFirewallTemplateRuleId() []*string
	SetRegionId(v string) *DeleteFirewallTemplateRulesRequest
	GetRegionId() *string
}

type DeleteFirewallTemplateRulesRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The **token*	- can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the firewall template.
	//
	// This parameter is required.
	//
	// example:
	//
	// ft-bcf1a7hrdq717****
	FirewallTemplateId *string `json:"FirewallTemplateId,omitempty" xml:"FirewallTemplateId,omitempty"`
	// The IDs of the firewall template rules.
	//
	// This parameter is required.
	FirewallTemplateRuleId []*string `json:"FirewallTemplateRuleId,omitempty" xml:"FirewallTemplateRuleId,omitempty" type:"Repeated"`
	// The region ID of the simple application server. You can call the [ListRegions](https://help.aliyun.com/document_detail/189315.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteFirewallTemplateRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteFirewallTemplateRulesRequest) GoString() string {
	return s.String()
}

func (s *DeleteFirewallTemplateRulesRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteFirewallTemplateRulesRequest) GetFirewallTemplateId() *string {
	return s.FirewallTemplateId
}

func (s *DeleteFirewallTemplateRulesRequest) GetFirewallTemplateRuleId() []*string {
	return s.FirewallTemplateRuleId
}

func (s *DeleteFirewallTemplateRulesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteFirewallTemplateRulesRequest) SetClientToken(v string) *DeleteFirewallTemplateRulesRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteFirewallTemplateRulesRequest) SetFirewallTemplateId(v string) *DeleteFirewallTemplateRulesRequest {
	s.FirewallTemplateId = &v
	return s
}

func (s *DeleteFirewallTemplateRulesRequest) SetFirewallTemplateRuleId(v []*string) *DeleteFirewallTemplateRulesRequest {
	s.FirewallTemplateRuleId = v
	return s
}

func (s *DeleteFirewallTemplateRulesRequest) SetRegionId(v string) *DeleteFirewallTemplateRulesRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteFirewallTemplateRulesRequest) Validate() error {
	return dara.Validate(s)
}
