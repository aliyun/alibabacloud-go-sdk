// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyTrFirewallV2RoutePolicyScopeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDestCandidateList(v []*ModifyTrFirewallV2RoutePolicyScopeRequestDestCandidateList) *ModifyTrFirewallV2RoutePolicyScopeRequest
	GetDestCandidateList() []*ModifyTrFirewallV2RoutePolicyScopeRequestDestCandidateList
	SetFirewallId(v string) *ModifyTrFirewallV2RoutePolicyScopeRequest
	GetFirewallId() *string
	SetLang(v string) *ModifyTrFirewallV2RoutePolicyScopeRequest
	GetLang() *string
	SetShouldRecover(v string) *ModifyTrFirewallV2RoutePolicyScopeRequest
	GetShouldRecover() *string
	SetSrcCandidateList(v []*ModifyTrFirewallV2RoutePolicyScopeRequestSrcCandidateList) *ModifyTrFirewallV2RoutePolicyScopeRequest
	GetSrcCandidateList() []*ModifyTrFirewallV2RoutePolicyScopeRequestSrcCandidateList
	SetTrFirewallRoutePolicyId(v string) *ModifyTrFirewallV2RoutePolicyScopeRequest
	GetTrFirewallRoutePolicyId() *string
}

type ModifyTrFirewallV2RoutePolicyScopeRequest struct {
	// The secondary traffic redirection instances.
	DestCandidateList []*ModifyTrFirewallV2RoutePolicyScopeRequestDestCandidateList `json:"DestCandidateList,omitempty" xml:"DestCandidateList,omitempty" type:"Repeated"`
	// The instance ID of the virtual private cloud (VPC) firewall.
	//
	// This parameter is required.
	//
	// example:
	//
	// vfw-tr-6520de0253bc4669bbd9
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
	// Specifies whether to restore the traffic redirection configurations. Valid values:
	//
	// 	- true: roll back
	//
	// 	- false: withdraw
	//
	// example:
	//
	// false
	ShouldRecover *string `json:"ShouldRecover,omitempty" xml:"ShouldRecover,omitempty"`
	// The primary traffic redirection instances.
	SrcCandidateList []*ModifyTrFirewallV2RoutePolicyScopeRequestSrcCandidateList `json:"SrcCandidateList,omitempty" xml:"SrcCandidateList,omitempty" type:"Repeated"`
	// The ID of the routing policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// policy-4d724d0139df48f18091
	TrFirewallRoutePolicyId *string `json:"TrFirewallRoutePolicyId,omitempty" xml:"TrFirewallRoutePolicyId,omitempty"`
}

func (s ModifyTrFirewallV2RoutePolicyScopeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyTrFirewallV2RoutePolicyScopeRequest) GoString() string {
	return s.String()
}

func (s *ModifyTrFirewallV2RoutePolicyScopeRequest) GetDestCandidateList() []*ModifyTrFirewallV2RoutePolicyScopeRequestDestCandidateList {
	return s.DestCandidateList
}

func (s *ModifyTrFirewallV2RoutePolicyScopeRequest) GetFirewallId() *string {
	return s.FirewallId
}

func (s *ModifyTrFirewallV2RoutePolicyScopeRequest) GetLang() *string {
	return s.Lang
}

func (s *ModifyTrFirewallV2RoutePolicyScopeRequest) GetShouldRecover() *string {
	return s.ShouldRecover
}

func (s *ModifyTrFirewallV2RoutePolicyScopeRequest) GetSrcCandidateList() []*ModifyTrFirewallV2RoutePolicyScopeRequestSrcCandidateList {
	return s.SrcCandidateList
}

func (s *ModifyTrFirewallV2RoutePolicyScopeRequest) GetTrFirewallRoutePolicyId() *string {
	return s.TrFirewallRoutePolicyId
}

func (s *ModifyTrFirewallV2RoutePolicyScopeRequest) SetDestCandidateList(v []*ModifyTrFirewallV2RoutePolicyScopeRequestDestCandidateList) *ModifyTrFirewallV2RoutePolicyScopeRequest {
	s.DestCandidateList = v
	return s
}

func (s *ModifyTrFirewallV2RoutePolicyScopeRequest) SetFirewallId(v string) *ModifyTrFirewallV2RoutePolicyScopeRequest {
	s.FirewallId = &v
	return s
}

func (s *ModifyTrFirewallV2RoutePolicyScopeRequest) SetLang(v string) *ModifyTrFirewallV2RoutePolicyScopeRequest {
	s.Lang = &v
	return s
}

func (s *ModifyTrFirewallV2RoutePolicyScopeRequest) SetShouldRecover(v string) *ModifyTrFirewallV2RoutePolicyScopeRequest {
	s.ShouldRecover = &v
	return s
}

func (s *ModifyTrFirewallV2RoutePolicyScopeRequest) SetSrcCandidateList(v []*ModifyTrFirewallV2RoutePolicyScopeRequestSrcCandidateList) *ModifyTrFirewallV2RoutePolicyScopeRequest {
	s.SrcCandidateList = v
	return s
}

func (s *ModifyTrFirewallV2RoutePolicyScopeRequest) SetTrFirewallRoutePolicyId(v string) *ModifyTrFirewallV2RoutePolicyScopeRequest {
	s.TrFirewallRoutePolicyId = &v
	return s
}

func (s *ModifyTrFirewallV2RoutePolicyScopeRequest) Validate() error {
	return dara.Validate(s)
}

type ModifyTrFirewallV2RoutePolicyScopeRequestDestCandidateList struct {
	// The ID of the traffic redirection instance.
	//
	// example:
	//
	// vpc-2ze9epancaw8t4shajuzi
	CandidateId *string `json:"CandidateId,omitempty" xml:"CandidateId,omitempty"`
	// The type of the traffic redirection instance.
	//
	// example:
	//
	// VPC
	CandidateType *string `json:"CandidateType,omitempty" xml:"CandidateType,omitempty"`
}

func (s ModifyTrFirewallV2RoutePolicyScopeRequestDestCandidateList) String() string {
	return dara.Prettify(s)
}

func (s ModifyTrFirewallV2RoutePolicyScopeRequestDestCandidateList) GoString() string {
	return s.String()
}

func (s *ModifyTrFirewallV2RoutePolicyScopeRequestDestCandidateList) GetCandidateId() *string {
	return s.CandidateId
}

func (s *ModifyTrFirewallV2RoutePolicyScopeRequestDestCandidateList) GetCandidateType() *string {
	return s.CandidateType
}

func (s *ModifyTrFirewallV2RoutePolicyScopeRequestDestCandidateList) SetCandidateId(v string) *ModifyTrFirewallV2RoutePolicyScopeRequestDestCandidateList {
	s.CandidateId = &v
	return s
}

func (s *ModifyTrFirewallV2RoutePolicyScopeRequestDestCandidateList) SetCandidateType(v string) *ModifyTrFirewallV2RoutePolicyScopeRequestDestCandidateList {
	s.CandidateType = &v
	return s
}

func (s *ModifyTrFirewallV2RoutePolicyScopeRequestDestCandidateList) Validate() error {
	return dara.Validate(s)
}

type ModifyTrFirewallV2RoutePolicyScopeRequestSrcCandidateList struct {
	// The ID of the traffic redirection instance.
	//
	// example:
	//
	// vpc-2ze9epancaw8t4shajuzi
	CandidateId *string `json:"CandidateId,omitempty" xml:"CandidateId,omitempty"`
	// The type of the traffic redirection instance.
	//
	// example:
	//
	// VPC
	CandidateType *string `json:"CandidateType,omitempty" xml:"CandidateType,omitempty"`
}

func (s ModifyTrFirewallV2RoutePolicyScopeRequestSrcCandidateList) String() string {
	return dara.Prettify(s)
}

func (s ModifyTrFirewallV2RoutePolicyScopeRequestSrcCandidateList) GoString() string {
	return s.String()
}

func (s *ModifyTrFirewallV2RoutePolicyScopeRequestSrcCandidateList) GetCandidateId() *string {
	return s.CandidateId
}

func (s *ModifyTrFirewallV2RoutePolicyScopeRequestSrcCandidateList) GetCandidateType() *string {
	return s.CandidateType
}

func (s *ModifyTrFirewallV2RoutePolicyScopeRequestSrcCandidateList) SetCandidateId(v string) *ModifyTrFirewallV2RoutePolicyScopeRequestSrcCandidateList {
	s.CandidateId = &v
	return s
}

func (s *ModifyTrFirewallV2RoutePolicyScopeRequestSrcCandidateList) SetCandidateType(v string) *ModifyTrFirewallV2RoutePolicyScopeRequestSrcCandidateList {
	s.CandidateType = &v
	return s
}

func (s *ModifyTrFirewallV2RoutePolicyScopeRequestSrcCandidateList) Validate() error {
	return dara.Validate(s)
}
