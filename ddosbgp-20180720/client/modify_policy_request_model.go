// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActionType(v int32) *ModifyPolicyRequest
	GetActionType() *int32
	SetContent(v *ModifyPolicyRequestContent) *ModifyPolicyRequest
	GetContent() *ModifyPolicyRequestContent
	SetId(v string) *ModifyPolicyRequest
	GetId() *string
	SetName(v string) *ModifyPolicyRequest
	GetName() *string
	SetPortVersion(v string) *ModifyPolicyRequest
	GetPortVersion() *string
}

type ModifyPolicyRequest struct {
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
	Content *ModifyPolicyRequestContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
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

func (s ModifyPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyPolicyRequest) GoString() string {
	return s.String()
}

func (s *ModifyPolicyRequest) GetActionType() *int32 {
	return s.ActionType
}

func (s *ModifyPolicyRequest) GetContent() *ModifyPolicyRequestContent {
	return s.Content
}

func (s *ModifyPolicyRequest) GetId() *string {
	return s.Id
}

func (s *ModifyPolicyRequest) GetName() *string {
	return s.Name
}

func (s *ModifyPolicyRequest) GetPortVersion() *string {
	return s.PortVersion
}

func (s *ModifyPolicyRequest) SetActionType(v int32) *ModifyPolicyRequest {
	s.ActionType = &v
	return s
}

func (s *ModifyPolicyRequest) SetContent(v *ModifyPolicyRequestContent) *ModifyPolicyRequest {
	s.Content = v
	return s
}

func (s *ModifyPolicyRequest) SetId(v string) *ModifyPolicyRequest {
	s.Id = &v
	return s
}

func (s *ModifyPolicyRequest) SetName(v string) *ModifyPolicyRequest {
	s.Name = &v
	return s
}

func (s *ModifyPolicyRequest) SetPortVersion(v string) *ModifyPolicyRequest {
	s.PortVersion = &v
	return s
}

func (s *ModifyPolicyRequest) Validate() error {
	if s.Content != nil {
		if err := s.Content.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyPolicyRequestContent struct {
	// The IP addresses in the blacklist.
	BlackIpList []*string `json:"BlackIpList,omitempty" xml:"BlackIpList,omitempty" type:"Repeated"`
	// The validity period of the IP address blacklist. The value is a UNIX timestamp.
	//
	// example:
	//
	// 1716878000
	BlackIpListExpireAt *int64 `json:"BlackIpListExpireAt,omitempty" xml:"BlackIpListExpireAt,omitempty"`
	// Specifies whether to enable ICMP blocking.
	//
	// example:
	//
	// true
	EnableDropIcmp *bool `json:"EnableDropIcmp,omitempty" xml:"EnableDropIcmp,omitempty"`
	// Specifies whether to enable intelligent protection.
	//
	// example:
	//
	// true
	EnableIntelligence *bool `json:"EnableIntelligence,omitempty" xml:"EnableIntelligence,omitempty"`
	// Specifies whether to enable port-specific mitigation.
	//
	// example:
	//
	// true
	EnableL4Defense *bool `json:"EnableL4Defense,omitempty" xml:"EnableL4Defense,omitempty"`
	// The byte-match filter rules.
	FingerPrintRuleList []*ModifyPolicyRequestContentFingerPrintRuleList `json:"FingerPrintRuleList,omitempty" xml:"FingerPrintRuleList,omitempty" type:"Repeated"`
	// The level of intelligent protection. Valid values:
	//
	// 	- **default**: normal.
	//
	// 	- **hard**: strict.
	//
	// 	- **weak**: loose.
	//
	// example:
	//
	// default
	IntelligenceLevel *string `json:"IntelligenceLevel,omitempty" xml:"IntelligenceLevel,omitempty"`
	// The port-specific mitigation rules.
	L4RuleList []*ModifyPolicyRequestContentL4RuleList `json:"L4RuleList,omitempty" xml:"L4RuleList,omitempty" type:"Repeated"`
	// The port blocking rules.
	PortRuleList []*ModifyPolicyRequestContentPortRuleList `json:"PortRuleList,omitempty" xml:"PortRuleList,omitempty" type:"Repeated"`
	// The ports whose traffic is filtered out by the filtering policies for UDP reflection attacks.
	ReflectBlockUdpPortList []*int32 `json:"ReflectBlockUdpPortList,omitempty" xml:"ReflectBlockUdpPortList,omitempty" type:"Repeated"`
	// The countries in the location blacklist.
	RegionBlockCountryList []*int32 `json:"RegionBlockCountryList,omitempty" xml:"RegionBlockCountryList,omitempty" type:"Repeated"`
	// The provinces in the location blacklist.
	RegionBlockProvinceList []*int32 `json:"RegionBlockProvinceList,omitempty" xml:"RegionBlockProvinceList,omitempty" type:"Repeated"`
	// The source IP addresses that are added to the blacklist.
	SourceBlockList []*ModifyPolicyRequestContentSourceBlockList `json:"SourceBlockList,omitempty" xml:"SourceBlockList,omitempty" type:"Repeated"`
	// The settings for source rate limiting.
	SourceLimit *ModifyPolicyRequestContentSourceLimit `json:"SourceLimit,omitempty" xml:"SourceLimit,omitempty" type:"Struct"`
	// The IP addresses in the whitelist.
	WhiteIpList []*string `json:"WhiteIpList,omitempty" xml:"WhiteIpList,omitempty" type:"Repeated"`
	// Specifies whether to add back-to-origin CIDR blocks of Anti-DDoS Proxy to the whitelist.
	//
	// example:
	//
	// false
	WhitenGfbrNets *bool `json:"WhitenGfbrNets,omitempty" xml:"WhitenGfbrNets,omitempty"`
}

func (s ModifyPolicyRequestContent) String() string {
	return dara.Prettify(s)
}

func (s ModifyPolicyRequestContent) GoString() string {
	return s.String()
}

func (s *ModifyPolicyRequestContent) GetBlackIpList() []*string {
	return s.BlackIpList
}

func (s *ModifyPolicyRequestContent) GetBlackIpListExpireAt() *int64 {
	return s.BlackIpListExpireAt
}

func (s *ModifyPolicyRequestContent) GetEnableDropIcmp() *bool {
	return s.EnableDropIcmp
}

func (s *ModifyPolicyRequestContent) GetEnableIntelligence() *bool {
	return s.EnableIntelligence
}

func (s *ModifyPolicyRequestContent) GetEnableL4Defense() *bool {
	return s.EnableL4Defense
}

func (s *ModifyPolicyRequestContent) GetFingerPrintRuleList() []*ModifyPolicyRequestContentFingerPrintRuleList {
	return s.FingerPrintRuleList
}

func (s *ModifyPolicyRequestContent) GetIntelligenceLevel() *string {
	return s.IntelligenceLevel
}

func (s *ModifyPolicyRequestContent) GetL4RuleList() []*ModifyPolicyRequestContentL4RuleList {
	return s.L4RuleList
}

func (s *ModifyPolicyRequestContent) GetPortRuleList() []*ModifyPolicyRequestContentPortRuleList {
	return s.PortRuleList
}

func (s *ModifyPolicyRequestContent) GetReflectBlockUdpPortList() []*int32 {
	return s.ReflectBlockUdpPortList
}

func (s *ModifyPolicyRequestContent) GetRegionBlockCountryList() []*int32 {
	return s.RegionBlockCountryList
}

func (s *ModifyPolicyRequestContent) GetRegionBlockProvinceList() []*int32 {
	return s.RegionBlockProvinceList
}

func (s *ModifyPolicyRequestContent) GetSourceBlockList() []*ModifyPolicyRequestContentSourceBlockList {
	return s.SourceBlockList
}

func (s *ModifyPolicyRequestContent) GetSourceLimit() *ModifyPolicyRequestContentSourceLimit {
	return s.SourceLimit
}

func (s *ModifyPolicyRequestContent) GetWhiteIpList() []*string {
	return s.WhiteIpList
}

func (s *ModifyPolicyRequestContent) GetWhitenGfbrNets() *bool {
	return s.WhitenGfbrNets
}

func (s *ModifyPolicyRequestContent) SetBlackIpList(v []*string) *ModifyPolicyRequestContent {
	s.BlackIpList = v
	return s
}

func (s *ModifyPolicyRequestContent) SetBlackIpListExpireAt(v int64) *ModifyPolicyRequestContent {
	s.BlackIpListExpireAt = &v
	return s
}

func (s *ModifyPolicyRequestContent) SetEnableDropIcmp(v bool) *ModifyPolicyRequestContent {
	s.EnableDropIcmp = &v
	return s
}

func (s *ModifyPolicyRequestContent) SetEnableIntelligence(v bool) *ModifyPolicyRequestContent {
	s.EnableIntelligence = &v
	return s
}

func (s *ModifyPolicyRequestContent) SetEnableL4Defense(v bool) *ModifyPolicyRequestContent {
	s.EnableL4Defense = &v
	return s
}

func (s *ModifyPolicyRequestContent) SetFingerPrintRuleList(v []*ModifyPolicyRequestContentFingerPrintRuleList) *ModifyPolicyRequestContent {
	s.FingerPrintRuleList = v
	return s
}

func (s *ModifyPolicyRequestContent) SetIntelligenceLevel(v string) *ModifyPolicyRequestContent {
	s.IntelligenceLevel = &v
	return s
}

func (s *ModifyPolicyRequestContent) SetL4RuleList(v []*ModifyPolicyRequestContentL4RuleList) *ModifyPolicyRequestContent {
	s.L4RuleList = v
	return s
}

func (s *ModifyPolicyRequestContent) SetPortRuleList(v []*ModifyPolicyRequestContentPortRuleList) *ModifyPolicyRequestContent {
	s.PortRuleList = v
	return s
}

func (s *ModifyPolicyRequestContent) SetReflectBlockUdpPortList(v []*int32) *ModifyPolicyRequestContent {
	s.ReflectBlockUdpPortList = v
	return s
}

func (s *ModifyPolicyRequestContent) SetRegionBlockCountryList(v []*int32) *ModifyPolicyRequestContent {
	s.RegionBlockCountryList = v
	return s
}

func (s *ModifyPolicyRequestContent) SetRegionBlockProvinceList(v []*int32) *ModifyPolicyRequestContent {
	s.RegionBlockProvinceList = v
	return s
}

func (s *ModifyPolicyRequestContent) SetSourceBlockList(v []*ModifyPolicyRequestContentSourceBlockList) *ModifyPolicyRequestContent {
	s.SourceBlockList = v
	return s
}

func (s *ModifyPolicyRequestContent) SetSourceLimit(v *ModifyPolicyRequestContentSourceLimit) *ModifyPolicyRequestContent {
	s.SourceLimit = v
	return s
}

func (s *ModifyPolicyRequestContent) SetWhiteIpList(v []*string) *ModifyPolicyRequestContent {
	s.WhiteIpList = v
	return s
}

func (s *ModifyPolicyRequestContent) SetWhitenGfbrNets(v bool) *ModifyPolicyRequestContent {
	s.WhitenGfbrNets = &v
	return s
}

func (s *ModifyPolicyRequestContent) Validate() error {
	if s.FingerPrintRuleList != nil {
		for _, item := range s.FingerPrintRuleList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.L4RuleList != nil {
		for _, item := range s.L4RuleList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.PortRuleList != nil {
		for _, item := range s.PortRuleList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SourceBlockList != nil {
		for _, item := range s.SourceBlockList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SourceLimit != nil {
		if err := s.SourceLimit.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyPolicyRequestContentFingerPrintRuleList struct {
	// The end of the destination port range. Valid values: **0*	- to **65535**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 65535
	DstPortEnd *int32 `json:"DstPortEnd,omitempty" xml:"DstPortEnd,omitempty"`
	// The start of the destination port range. Valid values: **0*	- to **65535**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	DstPortStart *int32 `json:"DstPortStart,omitempty" xml:"DstPortStart,omitempty"`
	// The ID of the rule.
	//
	// example:
	//
	// 5fbe941f-a0cf-4a49-9c7c-8fac********
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The action triggered if the rule is matched. Valid values:
	//
	// 	- **accept**: allows the traffic that matches the conditions in the byte-match filter rule.
	//
	// 	- **drop**: discards the traffic that matches the conditions in the byte-match filter rule.
	//
	// 	- **ip_rate**: limits rates on the source IP address whose traffic matches the conditions in the byte-match filter rule. The rate limit is specified by **RateValue**.
	//
	// 	- **session_rate**: limits the number of sessions from the source IP address whose traffic matches the conditions in the byte-match filter rule. The rate limit is specified by **RateValue**.
	//
	// This parameter is required.
	//
	// example:
	//
	// drop
	MatchAction *string `json:"MatchAction,omitempty" xml:"MatchAction,omitempty"`
	// The maximum packet length. Valid values: **1*	- to **1500**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1500
	MaxPktLen *int32 `json:"MaxPktLen,omitempty" xml:"MaxPktLen,omitempty"`
	// The minimum packet length. Valid values: **1*	- to **1500**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	MinPktLen *int32 `json:"MinPktLen,omitempty" xml:"MinPktLen,omitempty"`
	// The offset. Valid values: **0*	- to **1500**.
	//
	// example:
	//
	// 0
	Offset *int32 `json:"Offset,omitempty" xml:"Offset,omitempty"`
	// The payload. The value is a hexadecimal string.
	//
	// example:
	//
	// abcd
	PayloadBytes *string `json:"PayloadBytes,omitempty" xml:"PayloadBytes,omitempty"`
	// The type of the protocol. Valid values:
	//
	// 	- **tcp**
	//
	// 	- **udp**
	//
	// This parameter is required.
	//
	// example:
	//
	// udp
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The rate limit. Valid values: **1*	- to **100000**.
	//
	// >  This parameter is required when **MatchAction*	- is set to **ip_rate*	- or **session_rate**.
	//
	// example:
	//
	// 100
	RateValue *int32 `json:"RateValue,omitempty" xml:"RateValue,omitempty"`
	// The sequence number that indicates the order for the rule to take effect. The value is an integer.
	//
	// >  A smaller number indicates a higher priority.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	SeqNo *int32 `json:"SeqNo,omitempty" xml:"SeqNo,omitempty"`
	// The end of the source port range. Valid values: **0*	- to **65535**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 65535
	SrcPortEnd *int32 `json:"SrcPortEnd,omitempty" xml:"SrcPortEnd,omitempty"`
	// The start of the source port range. Valid values: **0*	- to **65535**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	SrcPortStart *int32 `json:"SrcPortStart,omitempty" xml:"SrcPortStart,omitempty"`
}

func (s ModifyPolicyRequestContentFingerPrintRuleList) String() string {
	return dara.Prettify(s)
}

func (s ModifyPolicyRequestContentFingerPrintRuleList) GoString() string {
	return s.String()
}

func (s *ModifyPolicyRequestContentFingerPrintRuleList) GetDstPortEnd() *int32 {
	return s.DstPortEnd
}

func (s *ModifyPolicyRequestContentFingerPrintRuleList) GetDstPortStart() *int32 {
	return s.DstPortStart
}

func (s *ModifyPolicyRequestContentFingerPrintRuleList) GetId() *string {
	return s.Id
}

func (s *ModifyPolicyRequestContentFingerPrintRuleList) GetMatchAction() *string {
	return s.MatchAction
}

func (s *ModifyPolicyRequestContentFingerPrintRuleList) GetMaxPktLen() *int32 {
	return s.MaxPktLen
}

func (s *ModifyPolicyRequestContentFingerPrintRuleList) GetMinPktLen() *int32 {
	return s.MinPktLen
}

func (s *ModifyPolicyRequestContentFingerPrintRuleList) GetOffset() *int32 {
	return s.Offset
}

func (s *ModifyPolicyRequestContentFingerPrintRuleList) GetPayloadBytes() *string {
	return s.PayloadBytes
}

func (s *ModifyPolicyRequestContentFingerPrintRuleList) GetProtocol() *string {
	return s.Protocol
}

func (s *ModifyPolicyRequestContentFingerPrintRuleList) GetRateValue() *int32 {
	return s.RateValue
}

func (s *ModifyPolicyRequestContentFingerPrintRuleList) GetSeqNo() *int32 {
	return s.SeqNo
}

func (s *ModifyPolicyRequestContentFingerPrintRuleList) GetSrcPortEnd() *int32 {
	return s.SrcPortEnd
}

func (s *ModifyPolicyRequestContentFingerPrintRuleList) GetSrcPortStart() *int32 {
	return s.SrcPortStart
}

func (s *ModifyPolicyRequestContentFingerPrintRuleList) SetDstPortEnd(v int32) *ModifyPolicyRequestContentFingerPrintRuleList {
	s.DstPortEnd = &v
	return s
}

func (s *ModifyPolicyRequestContentFingerPrintRuleList) SetDstPortStart(v int32) *ModifyPolicyRequestContentFingerPrintRuleList {
	s.DstPortStart = &v
	return s
}

func (s *ModifyPolicyRequestContentFingerPrintRuleList) SetId(v string) *ModifyPolicyRequestContentFingerPrintRuleList {
	s.Id = &v
	return s
}

func (s *ModifyPolicyRequestContentFingerPrintRuleList) SetMatchAction(v string) *ModifyPolicyRequestContentFingerPrintRuleList {
	s.MatchAction = &v
	return s
}

func (s *ModifyPolicyRequestContentFingerPrintRuleList) SetMaxPktLen(v int32) *ModifyPolicyRequestContentFingerPrintRuleList {
	s.MaxPktLen = &v
	return s
}

func (s *ModifyPolicyRequestContentFingerPrintRuleList) SetMinPktLen(v int32) *ModifyPolicyRequestContentFingerPrintRuleList {
	s.MinPktLen = &v
	return s
}

func (s *ModifyPolicyRequestContentFingerPrintRuleList) SetOffset(v int32) *ModifyPolicyRequestContentFingerPrintRuleList {
	s.Offset = &v
	return s
}

func (s *ModifyPolicyRequestContentFingerPrintRuleList) SetPayloadBytes(v string) *ModifyPolicyRequestContentFingerPrintRuleList {
	s.PayloadBytes = &v
	return s
}

func (s *ModifyPolicyRequestContentFingerPrintRuleList) SetProtocol(v string) *ModifyPolicyRequestContentFingerPrintRuleList {
	s.Protocol = &v
	return s
}

func (s *ModifyPolicyRequestContentFingerPrintRuleList) SetRateValue(v int32) *ModifyPolicyRequestContentFingerPrintRuleList {
	s.RateValue = &v
	return s
}

func (s *ModifyPolicyRequestContentFingerPrintRuleList) SetSeqNo(v int32) *ModifyPolicyRequestContentFingerPrintRuleList {
	s.SeqNo = &v
	return s
}

func (s *ModifyPolicyRequestContentFingerPrintRuleList) SetSrcPortEnd(v int32) *ModifyPolicyRequestContentFingerPrintRuleList {
	s.SrcPortEnd = &v
	return s
}

func (s *ModifyPolicyRequestContentFingerPrintRuleList) SetSrcPortStart(v int32) *ModifyPolicyRequestContentFingerPrintRuleList {
	s.SrcPortStart = &v
	return s
}

func (s *ModifyPolicyRequestContentFingerPrintRuleList) Validate() error {
	return dara.Validate(s)
}

type ModifyPolicyRequestContentL4RuleList struct {
	// The action that is specified in the rule. Valid value:
	//
	// 	- **2**: The traffic is discarded.
	//
	// example:
	//
	// 2
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// The match conditions.
	ConditionList []*ModifyPolicyRequestContentL4RuleListConditionList `json:"ConditionList,omitempty" xml:"ConditionList,omitempty" type:"Repeated"`
	// The minimum number of bytes in a session to trigger matching. Valid values: **0*	- to **2048**.
	//
	// example:
	//
	// 0
	Limited *int32 `json:"Limited,omitempty" xml:"Limited,omitempty"`
	// The condition based on which an action is performed. Valid values:
	//
	// 	- **0**: If the rule is matched, the action specified in the rule is performed.
	//
	// 	- **1**: If the rule is not matched, the action specified in the rule is performed.
	//
	// example:
	//
	// 0
	Match *string `json:"Match,omitempty" xml:"Match,omitempty"`
	// The type of the rule. Valid values:
	//
	// 	- **char**: string match.
	//
	// 	- **hex**: hexadecimal string match.
	//
	// example:
	//
	// char
	Method *string `json:"Method,omitempty" xml:"Method,omitempty"`
	// The name of the rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// test****
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The priority of the rule. Valid values: **1*	- to **100**.
	//
	// >  A smaller value indicates a higher priority.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
}

func (s ModifyPolicyRequestContentL4RuleList) String() string {
	return dara.Prettify(s)
}

func (s ModifyPolicyRequestContentL4RuleList) GoString() string {
	return s.String()
}

func (s *ModifyPolicyRequestContentL4RuleList) GetAction() *string {
	return s.Action
}

func (s *ModifyPolicyRequestContentL4RuleList) GetConditionList() []*ModifyPolicyRequestContentL4RuleListConditionList {
	return s.ConditionList
}

func (s *ModifyPolicyRequestContentL4RuleList) GetLimited() *int32 {
	return s.Limited
}

func (s *ModifyPolicyRequestContentL4RuleList) GetMatch() *string {
	return s.Match
}

func (s *ModifyPolicyRequestContentL4RuleList) GetMethod() *string {
	return s.Method
}

func (s *ModifyPolicyRequestContentL4RuleList) GetName() *string {
	return s.Name
}

func (s *ModifyPolicyRequestContentL4RuleList) GetPriority() *int32 {
	return s.Priority
}

func (s *ModifyPolicyRequestContentL4RuleList) SetAction(v string) *ModifyPolicyRequestContentL4RuleList {
	s.Action = &v
	return s
}

func (s *ModifyPolicyRequestContentL4RuleList) SetConditionList(v []*ModifyPolicyRequestContentL4RuleListConditionList) *ModifyPolicyRequestContentL4RuleList {
	s.ConditionList = v
	return s
}

func (s *ModifyPolicyRequestContentL4RuleList) SetLimited(v int32) *ModifyPolicyRequestContentL4RuleList {
	s.Limited = &v
	return s
}

func (s *ModifyPolicyRequestContentL4RuleList) SetMatch(v string) *ModifyPolicyRequestContentL4RuleList {
	s.Match = &v
	return s
}

func (s *ModifyPolicyRequestContentL4RuleList) SetMethod(v string) *ModifyPolicyRequestContentL4RuleList {
	s.Method = &v
	return s
}

func (s *ModifyPolicyRequestContentL4RuleList) SetName(v string) *ModifyPolicyRequestContentL4RuleList {
	s.Name = &v
	return s
}

func (s *ModifyPolicyRequestContentL4RuleList) SetPriority(v int32) *ModifyPolicyRequestContentL4RuleList {
	s.Priority = &v
	return s
}

func (s *ModifyPolicyRequestContentL4RuleList) Validate() error {
	if s.ConditionList != nil {
		for _, item := range s.ConditionList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ModifyPolicyRequestContentL4RuleListConditionList struct {
	// The term that is used for matching.
	//
	// >  If Method is set to **char**, the value of this parameter must be ASCII strings. If Method is set to **hex**, the value of this parameter must be hexadecimal strings. Maximum length: 2,048.
	//
	// example:
	//
	// abcd
	Arg     *string `json:"Arg,omitempty" xml:"Arg,omitempty"`
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The number of bytes from the start position for matching. Valid values: **1*	- to **2048**.
	//
	// example:
	//
	// 1200
	Depth   *int32                                                   `json:"Depth,omitempty" xml:"Depth,omitempty"`
	Encode  *string                                                  `json:"Encode,omitempty" xml:"Encode,omitempty"`
	Offset  *ModifyPolicyRequestContentL4RuleListConditionListOffset `json:"Offset,omitempty" xml:"Offset,omitempty" type:"Struct"`
	Pattern *string                                                  `json:"Pattern,omitempty" xml:"Pattern,omitempty"`
	// The start position for matching. Valid values: **0*	- to **2047**.
	//
	// example:
	//
	// 0
	Position *int32 `json:"Position,omitempty" xml:"Position,omitempty"`
}

func (s ModifyPolicyRequestContentL4RuleListConditionList) String() string {
	return dara.Prettify(s)
}

func (s ModifyPolicyRequestContentL4RuleListConditionList) GoString() string {
	return s.String()
}

func (s *ModifyPolicyRequestContentL4RuleListConditionList) GetArg() *string {
	return s.Arg
}

func (s *ModifyPolicyRequestContentL4RuleListConditionList) GetContent() *string {
	return s.Content
}

func (s *ModifyPolicyRequestContentL4RuleListConditionList) GetDepth() *int32 {
	return s.Depth
}

func (s *ModifyPolicyRequestContentL4RuleListConditionList) GetEncode() *string {
	return s.Encode
}

func (s *ModifyPolicyRequestContentL4RuleListConditionList) GetOffset() *ModifyPolicyRequestContentL4RuleListConditionListOffset {
	return s.Offset
}

func (s *ModifyPolicyRequestContentL4RuleListConditionList) GetPattern() *string {
	return s.Pattern
}

func (s *ModifyPolicyRequestContentL4RuleListConditionList) GetPosition() *int32 {
	return s.Position
}

func (s *ModifyPolicyRequestContentL4RuleListConditionList) SetArg(v string) *ModifyPolicyRequestContentL4RuleListConditionList {
	s.Arg = &v
	return s
}

func (s *ModifyPolicyRequestContentL4RuleListConditionList) SetContent(v string) *ModifyPolicyRequestContentL4RuleListConditionList {
	s.Content = &v
	return s
}

func (s *ModifyPolicyRequestContentL4RuleListConditionList) SetDepth(v int32) *ModifyPolicyRequestContentL4RuleListConditionList {
	s.Depth = &v
	return s
}

func (s *ModifyPolicyRequestContentL4RuleListConditionList) SetEncode(v string) *ModifyPolicyRequestContentL4RuleListConditionList {
	s.Encode = &v
	return s
}

func (s *ModifyPolicyRequestContentL4RuleListConditionList) SetOffset(v *ModifyPolicyRequestContentL4RuleListConditionListOffset) *ModifyPolicyRequestContentL4RuleListConditionList {
	s.Offset = v
	return s
}

func (s *ModifyPolicyRequestContentL4RuleListConditionList) SetPattern(v string) *ModifyPolicyRequestContentL4RuleListConditionList {
	s.Pattern = &v
	return s
}

func (s *ModifyPolicyRequestContentL4RuleListConditionList) SetPosition(v int32) *ModifyPolicyRequestContentL4RuleListConditionList {
	s.Position = &v
	return s
}

func (s *ModifyPolicyRequestContentL4RuleListConditionList) Validate() error {
	if s.Offset != nil {
		if err := s.Offset.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyPolicyRequestContentL4RuleListConditionListOffset struct {
	End   *int32 `json:"End,omitempty" xml:"End,omitempty"`
	Start *int32 `json:"Start,omitempty" xml:"Start,omitempty"`
}

func (s ModifyPolicyRequestContentL4RuleListConditionListOffset) String() string {
	return dara.Prettify(s)
}

func (s ModifyPolicyRequestContentL4RuleListConditionListOffset) GoString() string {
	return s.String()
}

func (s *ModifyPolicyRequestContentL4RuleListConditionListOffset) GetEnd() *int32 {
	return s.End
}

func (s *ModifyPolicyRequestContentL4RuleListConditionListOffset) GetStart() *int32 {
	return s.Start
}

func (s *ModifyPolicyRequestContentL4RuleListConditionListOffset) SetEnd(v int32) *ModifyPolicyRequestContentL4RuleListConditionListOffset {
	s.End = &v
	return s
}

func (s *ModifyPolicyRequestContentL4RuleListConditionListOffset) SetStart(v int32) *ModifyPolicyRequestContentL4RuleListConditionListOffset {
	s.Start = &v
	return s
}

func (s *ModifyPolicyRequestContentL4RuleListConditionListOffset) Validate() error {
	return dara.Validate(s)
}

type ModifyPolicyRequestContentPortRuleList struct {
	// The end of the destination port range. Valid values: **0*	- to **65535**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 65535
	DstPortEnd *int32 `json:"DstPortEnd,omitempty" xml:"DstPortEnd,omitempty"`
	// The start of the destination port range. Valid values: **0*	- to **65535**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	DstPortStart *int32 `json:"DstPortStart,omitempty" xml:"DstPortStart,omitempty"`
	// The ID of the rule.
	//
	// example:
	//
	// c52c2fa6-fdac-40c4-8753-be7c*********
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The action triggered if the rule is matched. Valid values:
	//
	// 	- **drop**: The traffic is discarded.
	//
	// This parameter is required.
	//
	// example:
	//
	// drop
	MatchAction *string `json:"MatchAction,omitempty" xml:"MatchAction,omitempty"`
	// The type of the protocol. Valid values:
	//
	// 	- **tcp**
	//
	// 	- **udp**
	//
	// This parameter is required.
	//
	// example:
	//
	// tcp
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The sequence number that indicates the order for the rule to take effect. The value is an integer.
	//
	// >  A smaller number indicates a higher priority.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	SeqNo *int32 `json:"SeqNo,omitempty" xml:"SeqNo,omitempty"`
	// The end of the source port range. Valid values: **0*	- to **65535**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 65535
	SrcPortEnd *int32 `json:"SrcPortEnd,omitempty" xml:"SrcPortEnd,omitempty"`
	// The start of the source port range. Valid values: **0*	- to **65535**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	SrcPortStart *int32 `json:"SrcPortStart,omitempty" xml:"SrcPortStart,omitempty"`
}

func (s ModifyPolicyRequestContentPortRuleList) String() string {
	return dara.Prettify(s)
}

func (s ModifyPolicyRequestContentPortRuleList) GoString() string {
	return s.String()
}

func (s *ModifyPolicyRequestContentPortRuleList) GetDstPortEnd() *int32 {
	return s.DstPortEnd
}

func (s *ModifyPolicyRequestContentPortRuleList) GetDstPortStart() *int32 {
	return s.DstPortStart
}

func (s *ModifyPolicyRequestContentPortRuleList) GetId() *string {
	return s.Id
}

func (s *ModifyPolicyRequestContentPortRuleList) GetMatchAction() *string {
	return s.MatchAction
}

func (s *ModifyPolicyRequestContentPortRuleList) GetProtocol() *string {
	return s.Protocol
}

func (s *ModifyPolicyRequestContentPortRuleList) GetSeqNo() *int32 {
	return s.SeqNo
}

func (s *ModifyPolicyRequestContentPortRuleList) GetSrcPortEnd() *int32 {
	return s.SrcPortEnd
}

func (s *ModifyPolicyRequestContentPortRuleList) GetSrcPortStart() *int32 {
	return s.SrcPortStart
}

func (s *ModifyPolicyRequestContentPortRuleList) SetDstPortEnd(v int32) *ModifyPolicyRequestContentPortRuleList {
	s.DstPortEnd = &v
	return s
}

func (s *ModifyPolicyRequestContentPortRuleList) SetDstPortStart(v int32) *ModifyPolicyRequestContentPortRuleList {
	s.DstPortStart = &v
	return s
}

func (s *ModifyPolicyRequestContentPortRuleList) SetId(v string) *ModifyPolicyRequestContentPortRuleList {
	s.Id = &v
	return s
}

func (s *ModifyPolicyRequestContentPortRuleList) SetMatchAction(v string) *ModifyPolicyRequestContentPortRuleList {
	s.MatchAction = &v
	return s
}

func (s *ModifyPolicyRequestContentPortRuleList) SetProtocol(v string) *ModifyPolicyRequestContentPortRuleList {
	s.Protocol = &v
	return s
}

func (s *ModifyPolicyRequestContentPortRuleList) SetSeqNo(v int32) *ModifyPolicyRequestContentPortRuleList {
	s.SeqNo = &v
	return s
}

func (s *ModifyPolicyRequestContentPortRuleList) SetSrcPortEnd(v int32) *ModifyPolicyRequestContentPortRuleList {
	s.SrcPortEnd = &v
	return s
}

func (s *ModifyPolicyRequestContentPortRuleList) SetSrcPortStart(v int32) *ModifyPolicyRequestContentPortRuleList {
	s.SrcPortStart = &v
	return s
}

func (s *ModifyPolicyRequestContentPortRuleList) Validate() error {
	return dara.Validate(s)
}

type ModifyPolicyRequestContentSourceBlockList struct {
	// The validity period of the blacklist to which the source IP address is added. Unit: seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 120
	BlockExpireSeconds *int32 `json:"BlockExpireSeconds,omitempty" xml:"BlockExpireSeconds,omitempty"`
	// The statistical period during which the system collects data on source IP addresses to determine whether to add the source IP addresses to the blacklist. Unit: seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 60
	EverySeconds *int32 `json:"EverySeconds,omitempty" xml:"EverySeconds,omitempty"`
	// The number of times that the source IP address exceeds a limit in a statistical period.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5
	ExceedLimitTimes *int32 `json:"ExceedLimitTimes,omitempty" xml:"ExceedLimitTimes,omitempty"`
	// The type of the source rate limit. Valid values:
	//
	// 	- **3**: the pps limit on source IP addresses.
	//
	// 	- **4**: the bandwidth limit on source IP addresses.
	//
	// 	- **5**: the pps limit on source SYN packets.
	//
	// 	- **6**: the bandwidth limit on source SYN packets.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ModifyPolicyRequestContentSourceBlockList) String() string {
	return dara.Prettify(s)
}

func (s ModifyPolicyRequestContentSourceBlockList) GoString() string {
	return s.String()
}

func (s *ModifyPolicyRequestContentSourceBlockList) GetBlockExpireSeconds() *int32 {
	return s.BlockExpireSeconds
}

func (s *ModifyPolicyRequestContentSourceBlockList) GetEverySeconds() *int32 {
	return s.EverySeconds
}

func (s *ModifyPolicyRequestContentSourceBlockList) GetExceedLimitTimes() *int32 {
	return s.ExceedLimitTimes
}

func (s *ModifyPolicyRequestContentSourceBlockList) GetType() *int32 {
	return s.Type
}

func (s *ModifyPolicyRequestContentSourceBlockList) SetBlockExpireSeconds(v int32) *ModifyPolicyRequestContentSourceBlockList {
	s.BlockExpireSeconds = &v
	return s
}

func (s *ModifyPolicyRequestContentSourceBlockList) SetEverySeconds(v int32) *ModifyPolicyRequestContentSourceBlockList {
	s.EverySeconds = &v
	return s
}

func (s *ModifyPolicyRequestContentSourceBlockList) SetExceedLimitTimes(v int32) *ModifyPolicyRequestContentSourceBlockList {
	s.ExceedLimitTimes = &v
	return s
}

func (s *ModifyPolicyRequestContentSourceBlockList) SetType(v int32) *ModifyPolicyRequestContentSourceBlockList {
	s.Type = &v
	return s
}

func (s *ModifyPolicyRequestContentSourceBlockList) Validate() error {
	return dara.Validate(s)
}

type ModifyPolicyRequestContentSourceLimit struct {
	// The bandwidth limit on source IP addresses. Unit: bytes per second.
	//
	// example:
	//
	// 2048
	Bps *int32 `json:"Bps,omitempty" xml:"Bps,omitempty"`
	// The packets per second (pps) limit on source IP addresses.
	//
	// example:
	//
	// 64
	Pps *int32 `json:"Pps,omitempty" xml:"Pps,omitempty"`
	// The bandwidth limit on source SYN packets. Unit: bytes per second.
	//
	// example:
	//
	// 2048
	SynBps *int32 `json:"SynBps,omitempty" xml:"SynBps,omitempty"`
	// The pps limit on source SYN packets.
	//
	// example:
	//
	// 64
	SynPps *int32 `json:"SynPps,omitempty" xml:"SynPps,omitempty"`
}

func (s ModifyPolicyRequestContentSourceLimit) String() string {
	return dara.Prettify(s)
}

func (s ModifyPolicyRequestContentSourceLimit) GoString() string {
	return s.String()
}

func (s *ModifyPolicyRequestContentSourceLimit) GetBps() *int32 {
	return s.Bps
}

func (s *ModifyPolicyRequestContentSourceLimit) GetPps() *int32 {
	return s.Pps
}

func (s *ModifyPolicyRequestContentSourceLimit) GetSynBps() *int32 {
	return s.SynBps
}

func (s *ModifyPolicyRequestContentSourceLimit) GetSynPps() *int32 {
	return s.SynPps
}

func (s *ModifyPolicyRequestContentSourceLimit) SetBps(v int32) *ModifyPolicyRequestContentSourceLimit {
	s.Bps = &v
	return s
}

func (s *ModifyPolicyRequestContentSourceLimit) SetPps(v int32) *ModifyPolicyRequestContentSourceLimit {
	s.Pps = &v
	return s
}

func (s *ModifyPolicyRequestContentSourceLimit) SetSynBps(v int32) *ModifyPolicyRequestContentSourceLimit {
	s.SynBps = &v
	return s
}

func (s *ModifyPolicyRequestContentSourceLimit) SetSynPps(v int32) *ModifyPolicyRequestContentSourceLimit {
	s.SynPps = &v
	return s
}

func (s *ModifyPolicyRequestContentSourceLimit) Validate() error {
	return dara.Validate(s)
}
