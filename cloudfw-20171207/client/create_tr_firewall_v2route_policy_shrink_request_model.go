// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTrFirewallV2RoutePolicyShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDestCandidateListShrink(v string) *CreateTrFirewallV2RoutePolicyShrinkRequest
	GetDestCandidateListShrink() *string
	SetFirewallId(v string) *CreateTrFirewallV2RoutePolicyShrinkRequest
	GetFirewallId() *string
	SetLang(v string) *CreateTrFirewallV2RoutePolicyShrinkRequest
	GetLang() *string
	SetPolicyDescription(v string) *CreateTrFirewallV2RoutePolicyShrinkRequest
	GetPolicyDescription() *string
	SetPolicyName(v string) *CreateTrFirewallV2RoutePolicyShrinkRequest
	GetPolicyName() *string
	SetPolicyType(v string) *CreateTrFirewallV2RoutePolicyShrinkRequest
	GetPolicyType() *string
	SetSrcCandidateListShrink(v string) *CreateTrFirewallV2RoutePolicyShrinkRequest
	GetSrcCandidateListShrink() *string
}

type CreateTrFirewallV2RoutePolicyShrinkRequest struct {
	// The secondary traffic redirection instances.
	DestCandidateListShrink *string `json:"DestCandidateList,omitempty" xml:"DestCandidateList,omitempty"`
	// The instance ID of the VPC firewall.
	//
	// example:
	//
	// vfw-tr-f8ce36689b224f77****
	FirewallId *string `json:"FirewallId,omitempty" xml:"FirewallId,omitempty"`
	// The language of the content within the response. Valid values:
	//
	// 	- **zh*	- (default): Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The description of the traffic redirection instance.
	//
	// example:
	//
	// test
	PolicyDescription *string `json:"PolicyDescription,omitempty" xml:"PolicyDescription,omitempty"`
	// The name of the traffic redirection instance.
	//
	// example:
	//
	// TEST_VPC_FW
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// The type of the traffic redirection scenario of the VPC firewall. Valid values:
	//
	// 	- **fullmesh**: interconnected instances
	//
	// 	- **one_to_one**: instance to instance
	//
	// 	- **end_to_end**: instance to instances
	//
	// example:
	//
	// fullmesh
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	// The primary traffic redirection instances.
	SrcCandidateListShrink *string `json:"SrcCandidateList,omitempty" xml:"SrcCandidateList,omitempty"`
}

func (s CreateTrFirewallV2RoutePolicyShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTrFirewallV2RoutePolicyShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateTrFirewallV2RoutePolicyShrinkRequest) GetDestCandidateListShrink() *string {
	return s.DestCandidateListShrink
}

func (s *CreateTrFirewallV2RoutePolicyShrinkRequest) GetFirewallId() *string {
	return s.FirewallId
}

func (s *CreateTrFirewallV2RoutePolicyShrinkRequest) GetLang() *string {
	return s.Lang
}

func (s *CreateTrFirewallV2RoutePolicyShrinkRequest) GetPolicyDescription() *string {
	return s.PolicyDescription
}

func (s *CreateTrFirewallV2RoutePolicyShrinkRequest) GetPolicyName() *string {
	return s.PolicyName
}

func (s *CreateTrFirewallV2RoutePolicyShrinkRequest) GetPolicyType() *string {
	return s.PolicyType
}

func (s *CreateTrFirewallV2RoutePolicyShrinkRequest) GetSrcCandidateListShrink() *string {
	return s.SrcCandidateListShrink
}

func (s *CreateTrFirewallV2RoutePolicyShrinkRequest) SetDestCandidateListShrink(v string) *CreateTrFirewallV2RoutePolicyShrinkRequest {
	s.DestCandidateListShrink = &v
	return s
}

func (s *CreateTrFirewallV2RoutePolicyShrinkRequest) SetFirewallId(v string) *CreateTrFirewallV2RoutePolicyShrinkRequest {
	s.FirewallId = &v
	return s
}

func (s *CreateTrFirewallV2RoutePolicyShrinkRequest) SetLang(v string) *CreateTrFirewallV2RoutePolicyShrinkRequest {
	s.Lang = &v
	return s
}

func (s *CreateTrFirewallV2RoutePolicyShrinkRequest) SetPolicyDescription(v string) *CreateTrFirewallV2RoutePolicyShrinkRequest {
	s.PolicyDescription = &v
	return s
}

func (s *CreateTrFirewallV2RoutePolicyShrinkRequest) SetPolicyName(v string) *CreateTrFirewallV2RoutePolicyShrinkRequest {
	s.PolicyName = &v
	return s
}

func (s *CreateTrFirewallV2RoutePolicyShrinkRequest) SetPolicyType(v string) *CreateTrFirewallV2RoutePolicyShrinkRequest {
	s.PolicyType = &v
	return s
}

func (s *CreateTrFirewallV2RoutePolicyShrinkRequest) SetSrcCandidateListShrink(v string) *CreateTrFirewallV2RoutePolicyShrinkRequest {
	s.SrcCandidateListShrink = &v
	return s
}

func (s *CreateTrFirewallV2RoutePolicyShrinkRequest) Validate() error {
	return dara.Validate(s)
}
