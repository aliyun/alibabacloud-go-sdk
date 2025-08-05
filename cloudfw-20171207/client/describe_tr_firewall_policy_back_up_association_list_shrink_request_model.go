// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTrFirewallPolicyBackUpAssociationListShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCandidateListShrink(v string) *DescribeTrFirewallPolicyBackUpAssociationListShrinkRequest
	GetCandidateListShrink() *string
	SetFirewallId(v string) *DescribeTrFirewallPolicyBackUpAssociationListShrinkRequest
	GetFirewallId() *string
	SetLang(v string) *DescribeTrFirewallPolicyBackUpAssociationListShrinkRequest
	GetLang() *string
	SetTrFirewallRoutePolicyId(v string) *DescribeTrFirewallPolicyBackUpAssociationListShrinkRequest
	GetTrFirewallRoutePolicyId() *string
}

type DescribeTrFirewallPolicyBackUpAssociationListShrinkRequest struct {
	// The traffic redirection instances.
	CandidateListShrink *string `json:"CandidateList,omitempty" xml:"CandidateList,omitempty"`
	// The instance ID of the VPC firewall.
	//
	// example:
	//
	// vfw-tr-8b268ce1b26e4c68****
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
	// The ID of the routing policy.
	//
	// example:
	//
	// policy-5dcafb12ff794a56****
	TrFirewallRoutePolicyId *string `json:"TrFirewallRoutePolicyId,omitempty" xml:"TrFirewallRoutePolicyId,omitempty"`
}

func (s DescribeTrFirewallPolicyBackUpAssociationListShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeTrFirewallPolicyBackUpAssociationListShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeTrFirewallPolicyBackUpAssociationListShrinkRequest) GetCandidateListShrink() *string {
	return s.CandidateListShrink
}

func (s *DescribeTrFirewallPolicyBackUpAssociationListShrinkRequest) GetFirewallId() *string {
	return s.FirewallId
}

func (s *DescribeTrFirewallPolicyBackUpAssociationListShrinkRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeTrFirewallPolicyBackUpAssociationListShrinkRequest) GetTrFirewallRoutePolicyId() *string {
	return s.TrFirewallRoutePolicyId
}

func (s *DescribeTrFirewallPolicyBackUpAssociationListShrinkRequest) SetCandidateListShrink(v string) *DescribeTrFirewallPolicyBackUpAssociationListShrinkRequest {
	s.CandidateListShrink = &v
	return s
}

func (s *DescribeTrFirewallPolicyBackUpAssociationListShrinkRequest) SetFirewallId(v string) *DescribeTrFirewallPolicyBackUpAssociationListShrinkRequest {
	s.FirewallId = &v
	return s
}

func (s *DescribeTrFirewallPolicyBackUpAssociationListShrinkRequest) SetLang(v string) *DescribeTrFirewallPolicyBackUpAssociationListShrinkRequest {
	s.Lang = &v
	return s
}

func (s *DescribeTrFirewallPolicyBackUpAssociationListShrinkRequest) SetTrFirewallRoutePolicyId(v string) *DescribeTrFirewallPolicyBackUpAssociationListShrinkRequest {
	s.TrFirewallRoutePolicyId = &v
	return s
}

func (s *DescribeTrFirewallPolicyBackUpAssociationListShrinkRequest) Validate() error {
	return dara.Validate(s)
}
