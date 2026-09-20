// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPolicyShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActionType(v int32) *ModifyPolicyShrinkRequest
	GetActionType() *int32
	SetContentShrink(v string) *ModifyPolicyShrinkRequest
	GetContentShrink() *string
	SetId(v string) *ModifyPolicyShrinkRequest
	GetId() *string
	SetName(v string) *ModifyPolicyShrinkRequest
	GetName() *string
	SetPortVersion(v string) *ModifyPolicyShrinkRequest
	GetPortVersion() *string
}

type ModifyPolicyShrinkRequest struct {
	// The action type. Valid values:
	//
	// - **10**: modifies the name (Name is required).
	//
	// - **11**: modifies the blacklist timeout period (BlackIpListExpireAt is required). Only IP-specific mitigation policy is supported.
	//
	// - **12**: modifies the switch for whitelisting back-to-origin IP addresses of Anti-DDoS Pro and Anti-DDoS Premium (WhitenGfbrNets is required). Only IP-specific mitigation policy is supported.
	//
	// - **13**: modifies the switch for ICMP Blocking (EnableDropIcmp is required). Only IP-specific mitigation policy is supported.
	//
	// - **20**: adds entries to blacklists and whitelists (WhiteIpList and BlackIpList are optional). Only IP-specific mitigation policy is supported.
	//
	// - **21**: deletes entries from blacklists and whitelists (WhiteIpList and BlackIpList are optional). Only IP-specific mitigation policy is supported.
	//
	// - **22**: clears the whitelist. Only IP-specific mitigation policy is supported.
	//
	// - **23**: clears the blacklist. Only IP-specific mitigation policy is supported.
	//
	// - **30**: modifies the AI-based intelligent protection switch and level (EnableIntelligence and IntelligenceLevel are required). Only IP-specific mitigation policy is supported.
	//
	// - **31**: modifies the Location Blacklist configuration (RegionBlockCountryList and RegionBlockProvinceList are optional). Only IP-specific mitigation policy is supported.
	//
	// - **32**: modifies the source rate limiting configuration (SourceLimit and SourceBlockList are required). Only IP-specific mitigation policy is supported.
	//
	// - **33**: modifies the reflection attack port filtering (ReflectBlockUdpPortList is required). Only IP-specific mitigation policy is supported.
	//
	// - **40**: creates a port blocking rule (PortRuleList is required). Only IP-specific mitigation policy is supported.
	//
	// - **41**: modifies a port blocking rule (PortRuleList is required). Only IP-specific mitigation policy is supported.
	//
	// - **42**: deletes a port blocking rule (PortRuleList is required). Only IP-specific mitigation policy is supported.
	//
	// - **50**: creates a byte-match filter rule (FingerPrintRuleList is required). Only IP-specific mitigation policy is supported.
	//
	// - **51**: modifies a byte-match filter rule (FingerPrintRuleList is required). Only IP-specific mitigation policy is supported.
	//
	// - **52**: deletes a byte-match filter rule (FingerPrintRuleList is required). Only IP-specific mitigation policy is supported.
	//
	// - **60**: modifies the port-specific mitigation switch (EnableL4Defense is required). Only port-specific mitigation policy is supported.
	//
	// - **61**: creates a port-specific mitigation rule (L4RuleList is required). Only port-specific mitigation policy is supported.
	//
	// - **62**: modifies a port-specific mitigation rule (L4RuleList is required). Only port-specific mitigation policy is supported.
	//
	// - **63**: deletes a port-specific mitigation rule (L4RuleList is required). Only port-specific mitigation policy is supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// 11
	ActionType *int32 `json:"ActionType,omitempty" xml:"ActionType,omitempty"`
	// The policy content.
	ContentShrink *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The policy ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// c52c2fa6-fdac-40c4-8753-be7c********
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The policy name.
	//
	// example:
	//
	// demo**
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The version of the port-specific mitigation policy. Valid values:
	//
	// - **Not specified**: Modifies the default surf mitigation engine policy.
	//
	// - **2**: Modifies the new stream mitigation engine policy.
	//
	// > Only port-specific mitigation policies are supported.
	//
	// example:
	//
	// 2
	PortVersion *string `json:"PortVersion,omitempty" xml:"PortVersion,omitempty"`
}

func (s ModifyPolicyShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyPolicyShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyPolicyShrinkRequest) GetActionType() *int32 {
	return s.ActionType
}

func (s *ModifyPolicyShrinkRequest) GetContentShrink() *string {
	return s.ContentShrink
}

func (s *ModifyPolicyShrinkRequest) GetId() *string {
	return s.Id
}

func (s *ModifyPolicyShrinkRequest) GetName() *string {
	return s.Name
}

func (s *ModifyPolicyShrinkRequest) GetPortVersion() *string {
	return s.PortVersion
}

func (s *ModifyPolicyShrinkRequest) SetActionType(v int32) *ModifyPolicyShrinkRequest {
	s.ActionType = &v
	return s
}

func (s *ModifyPolicyShrinkRequest) SetContentShrink(v string) *ModifyPolicyShrinkRequest {
	s.ContentShrink = &v
	return s
}

func (s *ModifyPolicyShrinkRequest) SetId(v string) *ModifyPolicyShrinkRequest {
	s.Id = &v
	return s
}

func (s *ModifyPolicyShrinkRequest) SetName(v string) *ModifyPolicyShrinkRequest {
	s.Name = &v
	return s
}

func (s *ModifyPolicyShrinkRequest) SetPortVersion(v string) *ModifyPolicyShrinkRequest {
	s.PortVersion = &v
	return s
}

func (s *ModifyPolicyShrinkRequest) Validate() error {
	return dara.Validate(s)
}
