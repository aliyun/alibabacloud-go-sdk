// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPolicyContentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v *ModifyPolicyContentRequestContent) *ModifyPolicyContentRequest
	GetContent() *ModifyPolicyContentRequestContent
	SetId(v string) *ModifyPolicyContentRequest
	GetId() *string
	SetName(v string) *ModifyPolicyContentRequest
	GetName() *string
	SetPortVersion(v string) *ModifyPolicyContentRequest
	GetPortVersion() *string
}

type ModifyPolicyContentRequest struct {
	// The policy content.
	Content *ModifyPolicyContentRequestContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The ID of the policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// 83967609-7ea5-4f6d-a6ea-380b09e****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the policy.
	//
	// example:
	//
	// demo**
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PortVersion *string `json:"PortVersion,omitempty" xml:"PortVersion,omitempty"`
}

func (s ModifyPolicyContentRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyPolicyContentRequest) GoString() string {
	return s.String()
}

func (s *ModifyPolicyContentRequest) GetContent() *ModifyPolicyContentRequestContent {
	return s.Content
}

func (s *ModifyPolicyContentRequest) GetId() *string {
	return s.Id
}

func (s *ModifyPolicyContentRequest) GetName() *string {
	return s.Name
}

func (s *ModifyPolicyContentRequest) GetPortVersion() *string {
	return s.PortVersion
}

func (s *ModifyPolicyContentRequest) SetContent(v *ModifyPolicyContentRequestContent) *ModifyPolicyContentRequest {
	s.Content = v
	return s
}

func (s *ModifyPolicyContentRequest) SetId(v string) *ModifyPolicyContentRequest {
	s.Id = &v
	return s
}

func (s *ModifyPolicyContentRequest) SetName(v string) *ModifyPolicyContentRequest {
	s.Name = &v
	return s
}

func (s *ModifyPolicyContentRequest) SetPortVersion(v string) *ModifyPolicyContentRequest {
	s.PortVersion = &v
	return s
}

func (s *ModifyPolicyContentRequest) Validate() error {
	if s.Content != nil {
		if err := s.Content.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyPolicyContentRequestContent struct {
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
	FingerPrintRuleList []*ModifyPolicyContentRequestContentFingerPrintRuleList `json:"FingerPrintRuleList,omitempty" xml:"FingerPrintRuleList,omitempty" type:"Repeated"`
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
	L4RuleList []*ModifyPolicyContentRequestContentL4RuleList `json:"L4RuleList,omitempty" xml:"L4RuleList,omitempty" type:"Repeated"`
	// The port blocking rules.
	PortRuleList []*ModifyPolicyContentRequestContentPortRuleList `json:"PortRuleList,omitempty" xml:"PortRuleList,omitempty" type:"Repeated"`
	// The ports whose traffic is filtered out by the filtering policies for UDP reflection attacks.
	ReflectBlockUdpPortList []*int32 `json:"ReflectBlockUdpPortList,omitempty" xml:"ReflectBlockUdpPortList,omitempty" type:"Repeated"`
	// The countries in the location blacklist.
	RegionBlockCountryList []*int32 `json:"RegionBlockCountryList,omitempty" xml:"RegionBlockCountryList,omitempty" type:"Repeated"`
	// The provinces in the location blacklist.
	RegionBlockProvinceList []*int32 `json:"RegionBlockProvinceList,omitempty" xml:"RegionBlockProvinceList,omitempty" type:"Repeated"`
	// The source IP addresses that are added to the blacklist.
	SourceBlockList []*ModifyPolicyContentRequestContentSourceBlockList `json:"SourceBlockList,omitempty" xml:"SourceBlockList,omitempty" type:"Repeated"`
	// The settings for source rate limiting.
	SourceLimit *ModifyPolicyContentRequestContentSourceLimit `json:"SourceLimit,omitempty" xml:"SourceLimit,omitempty" type:"Struct"`
	// Specifies whether to add back-to-origin CIDR blocks of Anti-DDoS Proxy to the whitelist.
	//
	// example:
	//
	// false
	WhitenGfbrNets *bool `json:"WhitenGfbrNets,omitempty" xml:"WhitenGfbrNets,omitempty"`
}

func (s ModifyPolicyContentRequestContent) String() string {
	return dara.Prettify(s)
}

func (s ModifyPolicyContentRequestContent) GoString() string {
	return s.String()
}

func (s *ModifyPolicyContentRequestContent) GetBlackIpListExpireAt() *int64 {
	return s.BlackIpListExpireAt
}

func (s *ModifyPolicyContentRequestContent) GetEnableDropIcmp() *bool {
	return s.EnableDropIcmp
}

func (s *ModifyPolicyContentRequestContent) GetEnableIntelligence() *bool {
	return s.EnableIntelligence
}

func (s *ModifyPolicyContentRequestContent) GetEnableL4Defense() *bool {
	return s.EnableL4Defense
}

func (s *ModifyPolicyContentRequestContent) GetFingerPrintRuleList() []*ModifyPolicyContentRequestContentFingerPrintRuleList {
	return s.FingerPrintRuleList
}

func (s *ModifyPolicyContentRequestContent) GetIntelligenceLevel() *string {
	return s.IntelligenceLevel
}

func (s *ModifyPolicyContentRequestContent) GetL4RuleList() []*ModifyPolicyContentRequestContentL4RuleList {
	return s.L4RuleList
}

func (s *ModifyPolicyContentRequestContent) GetPortRuleList() []*ModifyPolicyContentRequestContentPortRuleList {
	return s.PortRuleList
}

func (s *ModifyPolicyContentRequestContent) GetReflectBlockUdpPortList() []*int32 {
	return s.ReflectBlockUdpPortList
}

func (s *ModifyPolicyContentRequestContent) GetRegionBlockCountryList() []*int32 {
	return s.RegionBlockCountryList
}

func (s *ModifyPolicyContentRequestContent) GetRegionBlockProvinceList() []*int32 {
	return s.RegionBlockProvinceList
}

func (s *ModifyPolicyContentRequestContent) GetSourceBlockList() []*ModifyPolicyContentRequestContentSourceBlockList {
	return s.SourceBlockList
}

func (s *ModifyPolicyContentRequestContent) GetSourceLimit() *ModifyPolicyContentRequestContentSourceLimit {
	return s.SourceLimit
}

func (s *ModifyPolicyContentRequestContent) GetWhitenGfbrNets() *bool {
	return s.WhitenGfbrNets
}

func (s *ModifyPolicyContentRequestContent) SetBlackIpListExpireAt(v int64) *ModifyPolicyContentRequestContent {
	s.BlackIpListExpireAt = &v
	return s
}

func (s *ModifyPolicyContentRequestContent) SetEnableDropIcmp(v bool) *ModifyPolicyContentRequestContent {
	s.EnableDropIcmp = &v
	return s
}

func (s *ModifyPolicyContentRequestContent) SetEnableIntelligence(v bool) *ModifyPolicyContentRequestContent {
	s.EnableIntelligence = &v
	return s
}

func (s *ModifyPolicyContentRequestContent) SetEnableL4Defense(v bool) *ModifyPolicyContentRequestContent {
	s.EnableL4Defense = &v
	return s
}

func (s *ModifyPolicyContentRequestContent) SetFingerPrintRuleList(v []*ModifyPolicyContentRequestContentFingerPrintRuleList) *ModifyPolicyContentRequestContent {
	s.FingerPrintRuleList = v
	return s
}

func (s *ModifyPolicyContentRequestContent) SetIntelligenceLevel(v string) *ModifyPolicyContentRequestContent {
	s.IntelligenceLevel = &v
	return s
}

func (s *ModifyPolicyContentRequestContent) SetL4RuleList(v []*ModifyPolicyContentRequestContentL4RuleList) *ModifyPolicyContentRequestContent {
	s.L4RuleList = v
	return s
}

func (s *ModifyPolicyContentRequestContent) SetPortRuleList(v []*ModifyPolicyContentRequestContentPortRuleList) *ModifyPolicyContentRequestContent {
	s.PortRuleList = v
	return s
}

func (s *ModifyPolicyContentRequestContent) SetReflectBlockUdpPortList(v []*int32) *ModifyPolicyContentRequestContent {
	s.ReflectBlockUdpPortList = v
	return s
}

func (s *ModifyPolicyContentRequestContent) SetRegionBlockCountryList(v []*int32) *ModifyPolicyContentRequestContent {
	s.RegionBlockCountryList = v
	return s
}

func (s *ModifyPolicyContentRequestContent) SetRegionBlockProvinceList(v []*int32) *ModifyPolicyContentRequestContent {
	s.RegionBlockProvinceList = v
	return s
}

func (s *ModifyPolicyContentRequestContent) SetSourceBlockList(v []*ModifyPolicyContentRequestContentSourceBlockList) *ModifyPolicyContentRequestContent {
	s.SourceBlockList = v
	return s
}

func (s *ModifyPolicyContentRequestContent) SetSourceLimit(v *ModifyPolicyContentRequestContentSourceLimit) *ModifyPolicyContentRequestContent {
	s.SourceLimit = v
	return s
}

func (s *ModifyPolicyContentRequestContent) SetWhitenGfbrNets(v bool) *ModifyPolicyContentRequestContent {
	s.WhitenGfbrNets = &v
	return s
}

func (s *ModifyPolicyContentRequestContent) Validate() error {
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

type ModifyPolicyContentRequestContentFingerPrintRuleList struct {
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
	// 83967609-7ea5-4f6d-a6ea-380b09e****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The action triggered if the rule is matched. Valid values:
	//
	// 	- **permit**: allows the traffic that matches the conditions in the byte-match filter rule.
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
	// The protocol type. Valid values:
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

func (s ModifyPolicyContentRequestContentFingerPrintRuleList) String() string {
	return dara.Prettify(s)
}

func (s ModifyPolicyContentRequestContentFingerPrintRuleList) GoString() string {
	return s.String()
}

func (s *ModifyPolicyContentRequestContentFingerPrintRuleList) GetDstPortEnd() *int32 {
	return s.DstPortEnd
}

func (s *ModifyPolicyContentRequestContentFingerPrintRuleList) GetDstPortStart() *int32 {
	return s.DstPortStart
}

func (s *ModifyPolicyContentRequestContentFingerPrintRuleList) GetId() *string {
	return s.Id
}

func (s *ModifyPolicyContentRequestContentFingerPrintRuleList) GetMatchAction() *string {
	return s.MatchAction
}

func (s *ModifyPolicyContentRequestContentFingerPrintRuleList) GetMaxPktLen() *int32 {
	return s.MaxPktLen
}

func (s *ModifyPolicyContentRequestContentFingerPrintRuleList) GetMinPktLen() *int32 {
	return s.MinPktLen
}

func (s *ModifyPolicyContentRequestContentFingerPrintRuleList) GetOffset() *int32 {
	return s.Offset
}

func (s *ModifyPolicyContentRequestContentFingerPrintRuleList) GetPayloadBytes() *string {
	return s.PayloadBytes
}

func (s *ModifyPolicyContentRequestContentFingerPrintRuleList) GetProtocol() *string {
	return s.Protocol
}

func (s *ModifyPolicyContentRequestContentFingerPrintRuleList) GetRateValue() *int32 {
	return s.RateValue
}

func (s *ModifyPolicyContentRequestContentFingerPrintRuleList) GetSeqNo() *int32 {
	return s.SeqNo
}

func (s *ModifyPolicyContentRequestContentFingerPrintRuleList) GetSrcPortEnd() *int32 {
	return s.SrcPortEnd
}

func (s *ModifyPolicyContentRequestContentFingerPrintRuleList) GetSrcPortStart() *int32 {
	return s.SrcPortStart
}

func (s *ModifyPolicyContentRequestContentFingerPrintRuleList) SetDstPortEnd(v int32) *ModifyPolicyContentRequestContentFingerPrintRuleList {
	s.DstPortEnd = &v
	return s
}

func (s *ModifyPolicyContentRequestContentFingerPrintRuleList) SetDstPortStart(v int32) *ModifyPolicyContentRequestContentFingerPrintRuleList {
	s.DstPortStart = &v
	return s
}

func (s *ModifyPolicyContentRequestContentFingerPrintRuleList) SetId(v string) *ModifyPolicyContentRequestContentFingerPrintRuleList {
	s.Id = &v
	return s
}

func (s *ModifyPolicyContentRequestContentFingerPrintRuleList) SetMatchAction(v string) *ModifyPolicyContentRequestContentFingerPrintRuleList {
	s.MatchAction = &v
	return s
}

func (s *ModifyPolicyContentRequestContentFingerPrintRuleList) SetMaxPktLen(v int32) *ModifyPolicyContentRequestContentFingerPrintRuleList {
	s.MaxPktLen = &v
	return s
}

func (s *ModifyPolicyContentRequestContentFingerPrintRuleList) SetMinPktLen(v int32) *ModifyPolicyContentRequestContentFingerPrintRuleList {
	s.MinPktLen = &v
	return s
}

func (s *ModifyPolicyContentRequestContentFingerPrintRuleList) SetOffset(v int32) *ModifyPolicyContentRequestContentFingerPrintRuleList {
	s.Offset = &v
	return s
}

func (s *ModifyPolicyContentRequestContentFingerPrintRuleList) SetPayloadBytes(v string) *ModifyPolicyContentRequestContentFingerPrintRuleList {
	s.PayloadBytes = &v
	return s
}

func (s *ModifyPolicyContentRequestContentFingerPrintRuleList) SetProtocol(v string) *ModifyPolicyContentRequestContentFingerPrintRuleList {
	s.Protocol = &v
	return s
}

func (s *ModifyPolicyContentRequestContentFingerPrintRuleList) SetRateValue(v int32) *ModifyPolicyContentRequestContentFingerPrintRuleList {
	s.RateValue = &v
	return s
}

func (s *ModifyPolicyContentRequestContentFingerPrintRuleList) SetSeqNo(v int32) *ModifyPolicyContentRequestContentFingerPrintRuleList {
	s.SeqNo = &v
	return s
}

func (s *ModifyPolicyContentRequestContentFingerPrintRuleList) SetSrcPortEnd(v int32) *ModifyPolicyContentRequestContentFingerPrintRuleList {
	s.SrcPortEnd = &v
	return s
}

func (s *ModifyPolicyContentRequestContentFingerPrintRuleList) SetSrcPortStart(v int32) *ModifyPolicyContentRequestContentFingerPrintRuleList {
	s.SrcPortStart = &v
	return s
}

func (s *ModifyPolicyContentRequestContentFingerPrintRuleList) Validate() error {
	return dara.Validate(s)
}

type ModifyPolicyContentRequestContentL4RuleList struct {
	// The action that is specified in the rule. Valid value:
	//
	// 	- **2**: The traffic is discarded.
	//
	// example:
	//
	// 2
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// The match conditions.
	ConditionList []*ModifyPolicyContentRequestContentL4RuleListConditionList `json:"ConditionList,omitempty" xml:"ConditionList,omitempty" type:"Repeated"`
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
	// test**
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The priority of the rule. Valid values: 1 to 100.
	//
	// >  A smaller value indicates a higher priority.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
}

func (s ModifyPolicyContentRequestContentL4RuleList) String() string {
	return dara.Prettify(s)
}

func (s ModifyPolicyContentRequestContentL4RuleList) GoString() string {
	return s.String()
}

func (s *ModifyPolicyContentRequestContentL4RuleList) GetAction() *string {
	return s.Action
}

func (s *ModifyPolicyContentRequestContentL4RuleList) GetConditionList() []*ModifyPolicyContentRequestContentL4RuleListConditionList {
	return s.ConditionList
}

func (s *ModifyPolicyContentRequestContentL4RuleList) GetLimited() *int32 {
	return s.Limited
}

func (s *ModifyPolicyContentRequestContentL4RuleList) GetMatch() *string {
	return s.Match
}

func (s *ModifyPolicyContentRequestContentL4RuleList) GetMethod() *string {
	return s.Method
}

func (s *ModifyPolicyContentRequestContentL4RuleList) GetName() *string {
	return s.Name
}

func (s *ModifyPolicyContentRequestContentL4RuleList) GetPriority() *int32 {
	return s.Priority
}

func (s *ModifyPolicyContentRequestContentL4RuleList) SetAction(v string) *ModifyPolicyContentRequestContentL4RuleList {
	s.Action = &v
	return s
}

func (s *ModifyPolicyContentRequestContentL4RuleList) SetConditionList(v []*ModifyPolicyContentRequestContentL4RuleListConditionList) *ModifyPolicyContentRequestContentL4RuleList {
	s.ConditionList = v
	return s
}

func (s *ModifyPolicyContentRequestContentL4RuleList) SetLimited(v int32) *ModifyPolicyContentRequestContentL4RuleList {
	s.Limited = &v
	return s
}

func (s *ModifyPolicyContentRequestContentL4RuleList) SetMatch(v string) *ModifyPolicyContentRequestContentL4RuleList {
	s.Match = &v
	return s
}

func (s *ModifyPolicyContentRequestContentL4RuleList) SetMethod(v string) *ModifyPolicyContentRequestContentL4RuleList {
	s.Method = &v
	return s
}

func (s *ModifyPolicyContentRequestContentL4RuleList) SetName(v string) *ModifyPolicyContentRequestContentL4RuleList {
	s.Name = &v
	return s
}

func (s *ModifyPolicyContentRequestContentL4RuleList) SetPriority(v int32) *ModifyPolicyContentRequestContentL4RuleList {
	s.Priority = &v
	return s
}

func (s *ModifyPolicyContentRequestContentL4RuleList) Validate() error {
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

type ModifyPolicyContentRequestContentL4RuleListConditionList struct {
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
	Depth   *int32                                                          `json:"Depth,omitempty" xml:"Depth,omitempty"`
	Encode  *string                                                         `json:"Encode,omitempty" xml:"Encode,omitempty"`
	Offset  *ModifyPolicyContentRequestContentL4RuleListConditionListOffset `json:"Offset,omitempty" xml:"Offset,omitempty" type:"Struct"`
	Pattern *string                                                         `json:"Pattern,omitempty" xml:"Pattern,omitempty"`
	// The start position for matching. Valid values: **0*	- to **2047**.
	//
	// example:
	//
	// 0
	Position *int32 `json:"Position,omitempty" xml:"Position,omitempty"`
}

func (s ModifyPolicyContentRequestContentL4RuleListConditionList) String() string {
	return dara.Prettify(s)
}

func (s ModifyPolicyContentRequestContentL4RuleListConditionList) GoString() string {
	return s.String()
}

func (s *ModifyPolicyContentRequestContentL4RuleListConditionList) GetArg() *string {
	return s.Arg
}

func (s *ModifyPolicyContentRequestContentL4RuleListConditionList) GetContent() *string {
	return s.Content
}

func (s *ModifyPolicyContentRequestContentL4RuleListConditionList) GetDepth() *int32 {
	return s.Depth
}

func (s *ModifyPolicyContentRequestContentL4RuleListConditionList) GetEncode() *string {
	return s.Encode
}

func (s *ModifyPolicyContentRequestContentL4RuleListConditionList) GetOffset() *ModifyPolicyContentRequestContentL4RuleListConditionListOffset {
	return s.Offset
}

func (s *ModifyPolicyContentRequestContentL4RuleListConditionList) GetPattern() *string {
	return s.Pattern
}

func (s *ModifyPolicyContentRequestContentL4RuleListConditionList) GetPosition() *int32 {
	return s.Position
}

func (s *ModifyPolicyContentRequestContentL4RuleListConditionList) SetArg(v string) *ModifyPolicyContentRequestContentL4RuleListConditionList {
	s.Arg = &v
	return s
}

func (s *ModifyPolicyContentRequestContentL4RuleListConditionList) SetContent(v string) *ModifyPolicyContentRequestContentL4RuleListConditionList {
	s.Content = &v
	return s
}

func (s *ModifyPolicyContentRequestContentL4RuleListConditionList) SetDepth(v int32) *ModifyPolicyContentRequestContentL4RuleListConditionList {
	s.Depth = &v
	return s
}

func (s *ModifyPolicyContentRequestContentL4RuleListConditionList) SetEncode(v string) *ModifyPolicyContentRequestContentL4RuleListConditionList {
	s.Encode = &v
	return s
}

func (s *ModifyPolicyContentRequestContentL4RuleListConditionList) SetOffset(v *ModifyPolicyContentRequestContentL4RuleListConditionListOffset) *ModifyPolicyContentRequestContentL4RuleListConditionList {
	s.Offset = v
	return s
}

func (s *ModifyPolicyContentRequestContentL4RuleListConditionList) SetPattern(v string) *ModifyPolicyContentRequestContentL4RuleListConditionList {
	s.Pattern = &v
	return s
}

func (s *ModifyPolicyContentRequestContentL4RuleListConditionList) SetPosition(v int32) *ModifyPolicyContentRequestContentL4RuleListConditionList {
	s.Position = &v
	return s
}

func (s *ModifyPolicyContentRequestContentL4RuleListConditionList) Validate() error {
	if s.Offset != nil {
		if err := s.Offset.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyPolicyContentRequestContentL4RuleListConditionListOffset struct {
	End   *int32 `json:"End,omitempty" xml:"End,omitempty"`
	Start *int32 `json:"Start,omitempty" xml:"Start,omitempty"`
}

func (s ModifyPolicyContentRequestContentL4RuleListConditionListOffset) String() string {
	return dara.Prettify(s)
}

func (s ModifyPolicyContentRequestContentL4RuleListConditionListOffset) GoString() string {
	return s.String()
}

func (s *ModifyPolicyContentRequestContentL4RuleListConditionListOffset) GetEnd() *int32 {
	return s.End
}

func (s *ModifyPolicyContentRequestContentL4RuleListConditionListOffset) GetStart() *int32 {
	return s.Start
}

func (s *ModifyPolicyContentRequestContentL4RuleListConditionListOffset) SetEnd(v int32) *ModifyPolicyContentRequestContentL4RuleListConditionListOffset {
	s.End = &v
	return s
}

func (s *ModifyPolicyContentRequestContentL4RuleListConditionListOffset) SetStart(v int32) *ModifyPolicyContentRequestContentL4RuleListConditionListOffset {
	s.Start = &v
	return s
}

func (s *ModifyPolicyContentRequestContentL4RuleListConditionListOffset) Validate() error {
	return dara.Validate(s)
}

type ModifyPolicyContentRequestContentPortRuleList struct {
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
	// 412a7312-58ff-4e32-a202-0ab0*******
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
	// The protocol type. Valid values:
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

func (s ModifyPolicyContentRequestContentPortRuleList) String() string {
	return dara.Prettify(s)
}

func (s ModifyPolicyContentRequestContentPortRuleList) GoString() string {
	return s.String()
}

func (s *ModifyPolicyContentRequestContentPortRuleList) GetDstPortEnd() *int32 {
	return s.DstPortEnd
}

func (s *ModifyPolicyContentRequestContentPortRuleList) GetDstPortStart() *int32 {
	return s.DstPortStart
}

func (s *ModifyPolicyContentRequestContentPortRuleList) GetId() *string {
	return s.Id
}

func (s *ModifyPolicyContentRequestContentPortRuleList) GetMatchAction() *string {
	return s.MatchAction
}

func (s *ModifyPolicyContentRequestContentPortRuleList) GetProtocol() *string {
	return s.Protocol
}

func (s *ModifyPolicyContentRequestContentPortRuleList) GetSeqNo() *int32 {
	return s.SeqNo
}

func (s *ModifyPolicyContentRequestContentPortRuleList) GetSrcPortEnd() *int32 {
	return s.SrcPortEnd
}

func (s *ModifyPolicyContentRequestContentPortRuleList) GetSrcPortStart() *int32 {
	return s.SrcPortStart
}

func (s *ModifyPolicyContentRequestContentPortRuleList) SetDstPortEnd(v int32) *ModifyPolicyContentRequestContentPortRuleList {
	s.DstPortEnd = &v
	return s
}

func (s *ModifyPolicyContentRequestContentPortRuleList) SetDstPortStart(v int32) *ModifyPolicyContentRequestContentPortRuleList {
	s.DstPortStart = &v
	return s
}

func (s *ModifyPolicyContentRequestContentPortRuleList) SetId(v string) *ModifyPolicyContentRequestContentPortRuleList {
	s.Id = &v
	return s
}

func (s *ModifyPolicyContentRequestContentPortRuleList) SetMatchAction(v string) *ModifyPolicyContentRequestContentPortRuleList {
	s.MatchAction = &v
	return s
}

func (s *ModifyPolicyContentRequestContentPortRuleList) SetProtocol(v string) *ModifyPolicyContentRequestContentPortRuleList {
	s.Protocol = &v
	return s
}

func (s *ModifyPolicyContentRequestContentPortRuleList) SetSeqNo(v int32) *ModifyPolicyContentRequestContentPortRuleList {
	s.SeqNo = &v
	return s
}

func (s *ModifyPolicyContentRequestContentPortRuleList) SetSrcPortEnd(v int32) *ModifyPolicyContentRequestContentPortRuleList {
	s.SrcPortEnd = &v
	return s
}

func (s *ModifyPolicyContentRequestContentPortRuleList) SetSrcPortStart(v int32) *ModifyPolicyContentRequestContentPortRuleList {
	s.SrcPortStart = &v
	return s
}

func (s *ModifyPolicyContentRequestContentPortRuleList) Validate() error {
	return dara.Validate(s)
}

type ModifyPolicyContentRequestContentSourceBlockList struct {
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

func (s ModifyPolicyContentRequestContentSourceBlockList) String() string {
	return dara.Prettify(s)
}

func (s ModifyPolicyContentRequestContentSourceBlockList) GoString() string {
	return s.String()
}

func (s *ModifyPolicyContentRequestContentSourceBlockList) GetBlockExpireSeconds() *int32 {
	return s.BlockExpireSeconds
}

func (s *ModifyPolicyContentRequestContentSourceBlockList) GetEverySeconds() *int32 {
	return s.EverySeconds
}

func (s *ModifyPolicyContentRequestContentSourceBlockList) GetExceedLimitTimes() *int32 {
	return s.ExceedLimitTimes
}

func (s *ModifyPolicyContentRequestContentSourceBlockList) GetType() *int32 {
	return s.Type
}

func (s *ModifyPolicyContentRequestContentSourceBlockList) SetBlockExpireSeconds(v int32) *ModifyPolicyContentRequestContentSourceBlockList {
	s.BlockExpireSeconds = &v
	return s
}

func (s *ModifyPolicyContentRequestContentSourceBlockList) SetEverySeconds(v int32) *ModifyPolicyContentRequestContentSourceBlockList {
	s.EverySeconds = &v
	return s
}

func (s *ModifyPolicyContentRequestContentSourceBlockList) SetExceedLimitTimes(v int32) *ModifyPolicyContentRequestContentSourceBlockList {
	s.ExceedLimitTimes = &v
	return s
}

func (s *ModifyPolicyContentRequestContentSourceBlockList) SetType(v int32) *ModifyPolicyContentRequestContentSourceBlockList {
	s.Type = &v
	return s
}

func (s *ModifyPolicyContentRequestContentSourceBlockList) Validate() error {
	return dara.Validate(s)
}

type ModifyPolicyContentRequestContentSourceLimit struct {
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

func (s ModifyPolicyContentRequestContentSourceLimit) String() string {
	return dara.Prettify(s)
}

func (s ModifyPolicyContentRequestContentSourceLimit) GoString() string {
	return s.String()
}

func (s *ModifyPolicyContentRequestContentSourceLimit) GetBps() *int32 {
	return s.Bps
}

func (s *ModifyPolicyContentRequestContentSourceLimit) GetPps() *int32 {
	return s.Pps
}

func (s *ModifyPolicyContentRequestContentSourceLimit) GetSynBps() *int32 {
	return s.SynBps
}

func (s *ModifyPolicyContentRequestContentSourceLimit) GetSynPps() *int32 {
	return s.SynPps
}

func (s *ModifyPolicyContentRequestContentSourceLimit) SetBps(v int32) *ModifyPolicyContentRequestContentSourceLimit {
	s.Bps = &v
	return s
}

func (s *ModifyPolicyContentRequestContentSourceLimit) SetPps(v int32) *ModifyPolicyContentRequestContentSourceLimit {
	s.Pps = &v
	return s
}

func (s *ModifyPolicyContentRequestContentSourceLimit) SetSynBps(v int32) *ModifyPolicyContentRequestContentSourceLimit {
	s.SynBps = &v
	return s
}

func (s *ModifyPolicyContentRequestContentSourceLimit) SetSynPps(v int32) *ModifyPolicyContentRequestContentSourceLimit {
	s.SynPps = &v
	return s
}

func (s *ModifyPolicyContentRequestContentSourceLimit) Validate() error {
	return dara.Validate(s)
}
