// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyTrFirewallV2RoutePolicyScopeShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDestCandidateListShrink(v string) *ModifyTrFirewallV2RoutePolicyScopeShrinkRequest
	GetDestCandidateListShrink() *string
	SetFirewallId(v string) *ModifyTrFirewallV2RoutePolicyScopeShrinkRequest
	GetFirewallId() *string
	SetLang(v string) *ModifyTrFirewallV2RoutePolicyScopeShrinkRequest
	GetLang() *string
	SetShouldRecover(v string) *ModifyTrFirewallV2RoutePolicyScopeShrinkRequest
	GetShouldRecover() *string
	SetSrcCandidateListShrink(v string) *ModifyTrFirewallV2RoutePolicyScopeShrinkRequest
	GetSrcCandidateListShrink() *string
	SetTrFirewallRoutePolicyId(v string) *ModifyTrFirewallV2RoutePolicyScopeShrinkRequest
	GetTrFirewallRoutePolicyId() *string
}

type ModifyTrFirewallV2RoutePolicyScopeShrinkRequest struct {
	// The destination instances.
	DestCandidateListShrink *string `json:"DestCandidateList,omitempty" xml:"DestCandidateList,omitempty"`
	// The ID of the TR firewall instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// vfw-tr-6520de0253bc4669****
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
	// Specifies whether to restore the traffic redirection configuration. Valid values:
	//
	// - true: Rolls back the routes.
	//
	// - false: Revokes the routes.
	//
	// example:
	//
	// false
	ShouldRecover *string `json:"ShouldRecover,omitempty" xml:"ShouldRecover,omitempty"`
	// The source instances.
	SrcCandidateListShrink *string `json:"SrcCandidateList,omitempty" xml:"SrcCandidateList,omitempty"`
	// The ID of the routing policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// policy-4d724d0139df48f1****
	TrFirewallRoutePolicyId *string `json:"TrFirewallRoutePolicyId,omitempty" xml:"TrFirewallRoutePolicyId,omitempty"`
}

func (s ModifyTrFirewallV2RoutePolicyScopeShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyTrFirewallV2RoutePolicyScopeShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyTrFirewallV2RoutePolicyScopeShrinkRequest) GetDestCandidateListShrink() *string {
	return s.DestCandidateListShrink
}

func (s *ModifyTrFirewallV2RoutePolicyScopeShrinkRequest) GetFirewallId() *string {
	return s.FirewallId
}

func (s *ModifyTrFirewallV2RoutePolicyScopeShrinkRequest) GetLang() *string {
	return s.Lang
}

func (s *ModifyTrFirewallV2RoutePolicyScopeShrinkRequest) GetShouldRecover() *string {
	return s.ShouldRecover
}

func (s *ModifyTrFirewallV2RoutePolicyScopeShrinkRequest) GetSrcCandidateListShrink() *string {
	return s.SrcCandidateListShrink
}

func (s *ModifyTrFirewallV2RoutePolicyScopeShrinkRequest) GetTrFirewallRoutePolicyId() *string {
	return s.TrFirewallRoutePolicyId
}

func (s *ModifyTrFirewallV2RoutePolicyScopeShrinkRequest) SetDestCandidateListShrink(v string) *ModifyTrFirewallV2RoutePolicyScopeShrinkRequest {
	s.DestCandidateListShrink = &v
	return s
}

func (s *ModifyTrFirewallV2RoutePolicyScopeShrinkRequest) SetFirewallId(v string) *ModifyTrFirewallV2RoutePolicyScopeShrinkRequest {
	s.FirewallId = &v
	return s
}

func (s *ModifyTrFirewallV2RoutePolicyScopeShrinkRequest) SetLang(v string) *ModifyTrFirewallV2RoutePolicyScopeShrinkRequest {
	s.Lang = &v
	return s
}

func (s *ModifyTrFirewallV2RoutePolicyScopeShrinkRequest) SetShouldRecover(v string) *ModifyTrFirewallV2RoutePolicyScopeShrinkRequest {
	s.ShouldRecover = &v
	return s
}

func (s *ModifyTrFirewallV2RoutePolicyScopeShrinkRequest) SetSrcCandidateListShrink(v string) *ModifyTrFirewallV2RoutePolicyScopeShrinkRequest {
	s.SrcCandidateListShrink = &v
	return s
}

func (s *ModifyTrFirewallV2RoutePolicyScopeShrinkRequest) SetTrFirewallRoutePolicyId(v string) *ModifyTrFirewallV2RoutePolicyScopeShrinkRequest {
	s.TrFirewallRoutePolicyId = &v
	return s
}

func (s *ModifyTrFirewallV2RoutePolicyScopeShrinkRequest) Validate() error {
	return dara.Validate(s)
}
