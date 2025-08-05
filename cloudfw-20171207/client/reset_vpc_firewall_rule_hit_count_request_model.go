// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetVpcFirewallRuleHitCountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAclUuid(v string) *ResetVpcFirewallRuleHitCountRequest
	GetAclUuid() *string
	SetLang(v string) *ResetVpcFirewallRuleHitCountRequest
	GetLang() *string
}

type ResetVpcFirewallRuleHitCountRequest struct {
	// The ID of the access control policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// 00281255-d220-4db1-8f4f-c4df221a****
	AclUuid *string `json:"AclUuid,omitempty" xml:"AclUuid,omitempty"`
	// The natural language of the request and response.
	//
	// Valid values:
	//
	// - **zh**: Chinese (default)
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s ResetVpcFirewallRuleHitCountRequest) String() string {
	return dara.Prettify(s)
}

func (s ResetVpcFirewallRuleHitCountRequest) GoString() string {
	return s.String()
}

func (s *ResetVpcFirewallRuleHitCountRequest) GetAclUuid() *string {
	return s.AclUuid
}

func (s *ResetVpcFirewallRuleHitCountRequest) GetLang() *string {
	return s.Lang
}

func (s *ResetVpcFirewallRuleHitCountRequest) SetAclUuid(v string) *ResetVpcFirewallRuleHitCountRequest {
	s.AclUuid = &v
	return s
}

func (s *ResetVpcFirewallRuleHitCountRequest) SetLang(v string) *ResetVpcFirewallRuleHitCountRequest {
	s.Lang = &v
	return s
}

func (s *ResetVpcFirewallRuleHitCountRequest) Validate() error {
	return dara.Validate(s)
}
