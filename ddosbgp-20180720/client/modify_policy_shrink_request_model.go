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
	// The type of the action. Valid values:
	//
	// 	- **10**: modifies the name. If you specify this value, `Name` is required.
	//
	// 	- **11**: modifies the blacklist validity period. If you specify this value, `BlackIpListExpireAt` is required. Only IP-specific mitigation policies support this value.
	//
	// 	- **12**: changes the status of the feature of adding back-to-origin CIDR blocks of Anti-DDoS Proxy to the whitelist. If you specify this value, `WhitenGfbrNets` is required. Only IP-specific mitigation policies support this value.
	//
	// 	- **13**: changes the status of the ICMP blocking feature. If you specify this value, `EnableDropIcmp` is required. Only IP-specific mitigation policies support this value.
	//
	// 	- **20**: adds IP addresses to the blacklist or the whitelist. If you specify this value, you must specify at least one of `WhiteIpList` and `BlackIpList`. Only IP-specific mitigation policies support this value.
	//
	// 	- **21**: removes IP addresses from the blacklist or the whitelist. If you specify this value, at least one of `WhiteIpList` and `BlackIpList` is required. Only IP-specific mitigation policies support this value.
	//
	// 	- **22**: clears the whitelist. Only IP-specific mitigation policies support this value.
	//
	// 	- **23**: clears the blacklist. Only IP-specific mitigation policies support this value.
	//
	// 	- **30**: modifies the status and level of intelligent protection. If you specify this value, `EnableIntelligence` and `IntelligenceLevel` are required. Only IP-specific mitigation policies support this value.
	//
	// 	- **31**: modifies the location blacklist settings. If you specify this value, one of `RegionBlockCountryList` and `RegionBlockProvinceList` is required. Only IP-specific mitigation policies support this value.
	//
	// 	- **32**: modifies the settings for source rate limiting. If you specify this value, `SourceLimit` and `SourceBlockList` are required. Only IP-specific mitigation policies support this value.
	//
	// 	- **33**: modifies the settings for reflection attack filtering. If you specify this value, `ReflectBlockUdpPortList` is required. Only IP-specific mitigation policies support this value.
	//
	// 	- **40**: creates a port blocking rule. If you specify this value, `PortRuleList` is required. Only IP-specific mitigation policies support this value.
	//
	// 	- **41**: modifies the port blocking rule. If you specify this value, `PortRuleList` is required. Only IP-specific mitigation policies support this value.
	//
	// 	- **42**: deletes the port blocking rule. If you specify this value, `PortRuleList` is required. Only IP-specific mitigation policies support this value.
	//
	// 	- **50**: creates a byte-match filter rule. If you specify this value, `FingerPrintRuleList` is required. Only IP-specific mitigation policies support this value.
	//
	// 	- **51**: modifies the byte-match filter rule. If you specify this value, `FingerPrintRuleList` is required. Only IP-specific mitigation policies support this value.
	//
	// 	- **52**: deletes the byte-match filter rule. If you specify this value, `FingerPrintRuleList` is required. Only IP-specific mitigation policies support this value.
	//
	// 	- **60**: changes the status of the port-specific mitigation feature. If you specify this value, `EnableL4Defense` is required. Only port-specific mitigation policies support this value.
	//
	// 	- **61**: creates a port-specific mitigation rule. If you specify this value, `L4RuleList` is required. Only port-specific mitigation policies support this value.
	//
	// 	- **62**: modifies the port-specific mitigation rule. If you specify this value, `L4RuleList` is required. Only port-specific mitigation policies support this value.
	//
	// 	- **63**: deletes the port-specific mitigation rule. If you specify this value, `L4RuleList` is required. Only port-specific mitigation policies support this value.
	//
	// This parameter is required.
	//
	// example:
	//
	// 11
	ActionType *int32 `json:"ActionType,omitempty" xml:"ActionType,omitempty"`
	// The policy content.
	ContentShrink *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The ID of the policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// c52c2fa6-fdac-40c4-8753-be7c********
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the policy.
	//
	// example:
	//
	// demo**
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
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
