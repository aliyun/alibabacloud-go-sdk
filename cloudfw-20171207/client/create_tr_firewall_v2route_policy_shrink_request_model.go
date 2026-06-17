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
	// The list of destination network instances.
	DestCandidateListShrink *string `json:"DestCandidateList,omitempty" xml:"DestCandidateList,omitempty"`
	// The ID of the VPC firewall instance.
	//
	// example:
	//
	// vfw-tr-f8ce36689b224f77****
	FirewallId *string `json:"FirewallId,omitempty" xml:"FirewallId,omitempty"`
	// The language of the response. Valid values:
	//
	// - **zh*	- (default): Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The description of the routing policy.
	//
	// example:
	//
	// Singapore Point to Multipoint
	PolicyDescription *string `json:"PolicyDescription,omitempty" xml:"PolicyDescription,omitempty"`
	// The name of the routing policy.
	//
	// example:
	//
	// Singapore Point to Multipoint
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// The traffic redirection scenario of the Enterprise Edition transit router. Valid values:
	//
	// - **fullmesh**: full-mesh
	//
	// - **one_to_one**: point-to-point
	//
	// - **end_to_end**: point-to-multipoint
	//
	// example:
	//
	// fullmesh
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	// The list of source network instances.
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
