// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetNatFirewallRuleHitCountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAclUuid(v string) *ResetNatFirewallRuleHitCountRequest
	GetAclUuid() *string
	SetLang(v string) *ResetNatFirewallRuleHitCountRequest
	GetLang() *string
	SetNatGatewayId(v string) *ResetNatFirewallRuleHitCountRequest
	GetNatGatewayId() *string
}

type ResetNatFirewallRuleHitCountRequest struct {
	// The UUID of the access control policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3de3aed5-6de7-4ecd-9106-cfe994b9c49f
	AclUuid *string `json:"AclUuid,omitempty" xml:"AclUuid,omitempty"`
	// The language of the content within the request and response. Valid values:
	//
	// 	- **zh*	- (default): Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The ID of the NAT gateway.
	//
	// example:
	//
	// ngw-zm0h3c1exm5bifuorg8c5
	NatGatewayId *string `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
}

func (s ResetNatFirewallRuleHitCountRequest) String() string {
	return dara.Prettify(s)
}

func (s ResetNatFirewallRuleHitCountRequest) GoString() string {
	return s.String()
}

func (s *ResetNatFirewallRuleHitCountRequest) GetAclUuid() *string {
	return s.AclUuid
}

func (s *ResetNatFirewallRuleHitCountRequest) GetLang() *string {
	return s.Lang
}

func (s *ResetNatFirewallRuleHitCountRequest) GetNatGatewayId() *string {
	return s.NatGatewayId
}

func (s *ResetNatFirewallRuleHitCountRequest) SetAclUuid(v string) *ResetNatFirewallRuleHitCountRequest {
	s.AclUuid = &v
	return s
}

func (s *ResetNatFirewallRuleHitCountRequest) SetLang(v string) *ResetNatFirewallRuleHitCountRequest {
	s.Lang = &v
	return s
}

func (s *ResetNatFirewallRuleHitCountRequest) SetNatGatewayId(v string) *ResetNatFirewallRuleHitCountRequest {
	s.NatGatewayId = &v
	return s
}

func (s *ResetNatFirewallRuleHitCountRequest) Validate() error {
	return dara.Validate(s)
}
