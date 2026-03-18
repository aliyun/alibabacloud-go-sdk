// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPolicyList(v []*ListPolicyResponseBodyPolicyList) *ListPolicyResponseBody
	GetPolicyList() []*ListPolicyResponseBodyPolicyList
	SetRequestId(v string) *ListPolicyResponseBody
	GetRequestId() *string
	SetTotal(v int64) *ListPolicyResponseBody
	GetTotal() *int64
}

type ListPolicyResponseBody struct {
	// The policies.
	PolicyList []*ListPolicyResponseBodyPolicyList `json:"PolicyList,omitempty" xml:"PolicyList,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// B4B379C2-9319-4C6B-B579-FE36831B09F4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of policies.
	//
	// example:
	//
	// 10
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *ListPolicyResponseBody) GetPolicyList() []*ListPolicyResponseBodyPolicyList {
	return s.PolicyList
}

func (s *ListPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPolicyResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *ListPolicyResponseBody) SetPolicyList(v []*ListPolicyResponseBodyPolicyList) *ListPolicyResponseBody {
	s.PolicyList = v
	return s
}

func (s *ListPolicyResponseBody) SetRequestId(v string) *ListPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPolicyResponseBody) SetTotal(v int64) *ListPolicyResponseBody {
	s.Total = &v
	return s
}

func (s *ListPolicyResponseBody) Validate() error {
	if s.PolicyList != nil {
		for _, item := range s.PolicyList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListPolicyResponseBodyPolicyList struct {
	// The number of protected objects that are added to the policy.
	//
	// example:
	//
	// 0
	AttachedCount *int32 `json:"AttachedCount,omitempty" xml:"AttachedCount,omitempty"`
	// The content of the policy.
	Content *ListPolicyResponseBodyPolicyListContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The ID of the policy.
	//
	// example:
	//
	// 877afbdf-3982-4d36-9886-f043********
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the policy.
	//
	// example:
	//
	// test**
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The remarks of the policy.
	//
	// example:
	//
	// test
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The type of the policy. Valid values:
	//
	// 	- **default**: the default mitigation policy.
	//
	// 	- **l3**: IP-specific mitigation policies.
	//
	// 	- **l4**: port-specific mitigation policies.
	//
	// example:
	//
	// l3
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListPolicyResponseBodyPolicyList) String() string {
	return dara.Prettify(s)
}

func (s ListPolicyResponseBodyPolicyList) GoString() string {
	return s.String()
}

func (s *ListPolicyResponseBodyPolicyList) GetAttachedCount() *int32 {
	return s.AttachedCount
}

func (s *ListPolicyResponseBodyPolicyList) GetContent() *ListPolicyResponseBodyPolicyListContent {
	return s.Content
}

func (s *ListPolicyResponseBodyPolicyList) GetId() *string {
	return s.Id
}

func (s *ListPolicyResponseBodyPolicyList) GetName() *string {
	return s.Name
}

func (s *ListPolicyResponseBodyPolicyList) GetRemark() *string {
	return s.Remark
}

func (s *ListPolicyResponseBodyPolicyList) GetType() *string {
	return s.Type
}

func (s *ListPolicyResponseBodyPolicyList) SetAttachedCount(v int32) *ListPolicyResponseBodyPolicyList {
	s.AttachedCount = &v
	return s
}

func (s *ListPolicyResponseBodyPolicyList) SetContent(v *ListPolicyResponseBodyPolicyListContent) *ListPolicyResponseBodyPolicyList {
	s.Content = v
	return s
}

func (s *ListPolicyResponseBodyPolicyList) SetId(v string) *ListPolicyResponseBodyPolicyList {
	s.Id = &v
	return s
}

func (s *ListPolicyResponseBodyPolicyList) SetName(v string) *ListPolicyResponseBodyPolicyList {
	s.Name = &v
	return s
}

func (s *ListPolicyResponseBodyPolicyList) SetRemark(v string) *ListPolicyResponseBodyPolicyList {
	s.Remark = &v
	return s
}

func (s *ListPolicyResponseBodyPolicyList) SetType(v string) *ListPolicyResponseBodyPolicyList {
	s.Type = &v
	return s
}

func (s *ListPolicyResponseBodyPolicyList) Validate() error {
	if s.Content != nil {
		if err := s.Content.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListPolicyResponseBodyPolicyListContent struct {
	// The validity period of the IP address blacklist. The value is a UNIX timestamp.
	//
	// example:
	//
	// 1716878000
	BlackIpListExpireAt *int64 `json:"BlackIpListExpireAt,omitempty" xml:"BlackIpListExpireAt,omitempty"`
	// Indicates whether ICMP blocking is enabled.
	//
	// example:
	//
	// false
	EnableDropIcmp *bool `json:"EnableDropIcmp,omitempty" xml:"EnableDropIcmp,omitempty"`
	// Indicates whether intelligent protection is enabled.
	//
	// example:
	//
	// true
	EnableIntelligence *bool `json:"EnableIntelligence,omitempty" xml:"EnableIntelligence,omitempty"`
	// Indicates whether port-specific mitigation is enabled.
	//
	// example:
	//
	// true
	EnableL4Defense *bool `json:"EnableL4Defense,omitempty" xml:"EnableL4Defense,omitempty"`
	// The byte-match filter rules.
	FingerPrintRuleList []*ListPolicyResponseBodyPolicyListContentFingerPrintRuleList `json:"FingerPrintRuleList,omitempty" xml:"FingerPrintRuleList,omitempty" type:"Repeated"`
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
	L4RuleList []*ListPolicyResponseBodyPolicyListContentL4RuleList `json:"L4RuleList,omitempty" xml:"L4RuleList,omitempty" type:"Repeated"`
	// The port blocking rules.
	PortRuleList []*ListPolicyResponseBodyPolicyListContentPortRuleList `json:"PortRuleList,omitempty" xml:"PortRuleList,omitempty" type:"Repeated"`
	PortVersion  *string                                                `json:"PortVersion,omitempty" xml:"PortVersion,omitempty"`
	// The ports whose traffic is filtered out by the filtering policies for UDP reflection attacks.
	ReflectBlockUdpPortList []*int32 `json:"ReflectBlockUdpPortList,omitempty" xml:"ReflectBlockUdpPortList,omitempty" type:"Repeated"`
	// The countries in the location blacklist.
	RegionBlockCountryList []*int32 `json:"RegionBlockCountryList,omitempty" xml:"RegionBlockCountryList,omitempty" type:"Repeated"`
	// The provinces in the location blacklist.
	RegionBlockProvinceList []*int32 `json:"RegionBlockProvinceList,omitempty" xml:"RegionBlockProvinceList,omitempty" type:"Repeated"`
	// The source IP addresses that are added to the blacklist.
	SourceBlockList []*ListPolicyResponseBodyPolicyListContentSourceBlockList `json:"SourceBlockList,omitempty" xml:"SourceBlockList,omitempty" type:"Repeated"`
	// The settings for source rate limiting.
	SourceLimit *ListPolicyResponseBodyPolicyListContentSourceLimit `json:"SourceLimit,omitempty" xml:"SourceLimit,omitempty" type:"Struct"`
	// Indicates whether back-to-origin CIDR blocks of Anti-DDoS Proxy are added to the whitelist.
	//
	// example:
	//
	// false
	WhitenGfbrNets *bool `json:"WhitenGfbrNets,omitempty" xml:"WhitenGfbrNets,omitempty"`
}

func (s ListPolicyResponseBodyPolicyListContent) String() string {
	return dara.Prettify(s)
}

func (s ListPolicyResponseBodyPolicyListContent) GoString() string {
	return s.String()
}

func (s *ListPolicyResponseBodyPolicyListContent) GetBlackIpListExpireAt() *int64 {
	return s.BlackIpListExpireAt
}

func (s *ListPolicyResponseBodyPolicyListContent) GetEnableDropIcmp() *bool {
	return s.EnableDropIcmp
}

func (s *ListPolicyResponseBodyPolicyListContent) GetEnableIntelligence() *bool {
	return s.EnableIntelligence
}

func (s *ListPolicyResponseBodyPolicyListContent) GetEnableL4Defense() *bool {
	return s.EnableL4Defense
}

func (s *ListPolicyResponseBodyPolicyListContent) GetFingerPrintRuleList() []*ListPolicyResponseBodyPolicyListContentFingerPrintRuleList {
	return s.FingerPrintRuleList
}

func (s *ListPolicyResponseBodyPolicyListContent) GetIntelligenceLevel() *string {
	return s.IntelligenceLevel
}

func (s *ListPolicyResponseBodyPolicyListContent) GetL4RuleList() []*ListPolicyResponseBodyPolicyListContentL4RuleList {
	return s.L4RuleList
}

func (s *ListPolicyResponseBodyPolicyListContent) GetPortRuleList() []*ListPolicyResponseBodyPolicyListContentPortRuleList {
	return s.PortRuleList
}

func (s *ListPolicyResponseBodyPolicyListContent) GetPortVersion() *string {
	return s.PortVersion
}

func (s *ListPolicyResponseBodyPolicyListContent) GetReflectBlockUdpPortList() []*int32 {
	return s.ReflectBlockUdpPortList
}

func (s *ListPolicyResponseBodyPolicyListContent) GetRegionBlockCountryList() []*int32 {
	return s.RegionBlockCountryList
}

func (s *ListPolicyResponseBodyPolicyListContent) GetRegionBlockProvinceList() []*int32 {
	return s.RegionBlockProvinceList
}

func (s *ListPolicyResponseBodyPolicyListContent) GetSourceBlockList() []*ListPolicyResponseBodyPolicyListContentSourceBlockList {
	return s.SourceBlockList
}

func (s *ListPolicyResponseBodyPolicyListContent) GetSourceLimit() *ListPolicyResponseBodyPolicyListContentSourceLimit {
	return s.SourceLimit
}

func (s *ListPolicyResponseBodyPolicyListContent) GetWhitenGfbrNets() *bool {
	return s.WhitenGfbrNets
}

func (s *ListPolicyResponseBodyPolicyListContent) SetBlackIpListExpireAt(v int64) *ListPolicyResponseBodyPolicyListContent {
	s.BlackIpListExpireAt = &v
	return s
}

func (s *ListPolicyResponseBodyPolicyListContent) SetEnableDropIcmp(v bool) *ListPolicyResponseBodyPolicyListContent {
	s.EnableDropIcmp = &v
	return s
}

func (s *ListPolicyResponseBodyPolicyListContent) SetEnableIntelligence(v bool) *ListPolicyResponseBodyPolicyListContent {
	s.EnableIntelligence = &v
	return s
}

func (s *ListPolicyResponseBodyPolicyListContent) SetEnableL4Defense(v bool) *ListPolicyResponseBodyPolicyListContent {
	s.EnableL4Defense = &v
	return s
}

func (s *ListPolicyResponseBodyPolicyListContent) SetFingerPrintRuleList(v []*ListPolicyResponseBodyPolicyListContentFingerPrintRuleList) *ListPolicyResponseBodyPolicyListContent {
	s.FingerPrintRuleList = v
	return s
}

func (s *ListPolicyResponseBodyPolicyListContent) SetIntelligenceLevel(v string) *ListPolicyResponseBodyPolicyListContent {
	s.IntelligenceLevel = &v
	return s
}

func (s *ListPolicyResponseBodyPolicyListContent) SetL4RuleList(v []*ListPolicyResponseBodyPolicyListContentL4RuleList) *ListPolicyResponseBodyPolicyListContent {
	s.L4RuleList = v
	return s
}

func (s *ListPolicyResponseBodyPolicyListContent) SetPortRuleList(v []*ListPolicyResponseBodyPolicyListContentPortRuleList) *ListPolicyResponseBodyPolicyListContent {
	s.PortRuleList = v
	return s
}

func (s *ListPolicyResponseBodyPolicyListContent) SetPortVersion(v string) *ListPolicyResponseBodyPolicyListContent {
	s.PortVersion = &v
	return s
}

func (s *ListPolicyResponseBodyPolicyListContent) SetReflectBlockUdpPortList(v []*int32) *ListPolicyResponseBodyPolicyListContent {
	s.ReflectBlockUdpPortList = v
	return s
}

func (s *ListPolicyResponseBodyPolicyListContent) SetRegionBlockCountryList(v []*int32) *ListPolicyResponseBodyPolicyListContent {
	s.RegionBlockCountryList = v
	return s
}

func (s *ListPolicyResponseBodyPolicyListContent) SetRegionBlockProvinceList(v []*int32) *ListPolicyResponseBodyPolicyListContent {
	s.RegionBlockProvinceList = v
	return s
}

func (s *ListPolicyResponseBodyPolicyListContent) SetSourceBlockList(v []*ListPolicyResponseBodyPolicyListContentSourceBlockList) *ListPolicyResponseBodyPolicyListContent {
	s.SourceBlockList = v
	return s
}

func (s *ListPolicyResponseBodyPolicyListContent) SetSourceLimit(v *ListPolicyResponseBodyPolicyListContentSourceLimit) *ListPolicyResponseBodyPolicyListContent {
	s.SourceLimit = v
	return s
}

func (s *ListPolicyResponseBodyPolicyListContent) SetWhitenGfbrNets(v bool) *ListPolicyResponseBodyPolicyListContent {
	s.WhitenGfbrNets = &v
	return s
}

func (s *ListPolicyResponseBodyPolicyListContent) Validate() error {
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

type ListPolicyResponseBodyPolicyListContentFingerPrintRuleList struct {
	// The end of the destination port range. Valid values: **0*	- to **65535**.
	//
	// example:
	//
	// 65535
	DstPortEnd *int32 `json:"DstPortEnd,omitempty" xml:"DstPortEnd,omitempty"`
	// The start of the destination port range. Valid values: **0*	- to **65535**.
	//
	// example:
	//
	// 0
	DstPortStart *int32 `json:"DstPortStart,omitempty" xml:"DstPortStart,omitempty"`
	// The ID of the rule.
	//
	// example:
	//
	// 2c0b09cd-a565-4481-9acb-418b********
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
	// example:
	//
	// drop
	MatchAction *string `json:"MatchAction,omitempty" xml:"MatchAction,omitempty"`
	// The maximum packet length. Valid values: **1*	- to **1500**.
	//
	// example:
	//
	// 1500
	MaxPktLen *int32 `json:"MaxPktLen,omitempty" xml:"MaxPktLen,omitempty"`
	// The minimum packet length. Valid values: **1*	- to **1500**.
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
	// 1000
	RateValue *int32 `json:"RateValue,omitempty" xml:"RateValue,omitempty"`
	// The sequence number that indicates the order for the rule to take effect. The value is an integer.
	//
	// example:
	//
	// 1
	SeqNo *int32 `json:"SeqNo,omitempty" xml:"SeqNo,omitempty"`
	// The end of the source port range. Valid values: **0*	- to **65535**.
	//
	// example:
	//
	// 65535
	SrcPortEnd *int32 `json:"SrcPortEnd,omitempty" xml:"SrcPortEnd,omitempty"`
	// The start of the source port range. Valid values: **0*	- to **65535**.
	//
	// example:
	//
	// 0
	SrcPortStart *int32 `json:"SrcPortStart,omitempty" xml:"SrcPortStart,omitempty"`
}

func (s ListPolicyResponseBodyPolicyListContentFingerPrintRuleList) String() string {
	return dara.Prettify(s)
}

func (s ListPolicyResponseBodyPolicyListContentFingerPrintRuleList) GoString() string {
	return s.String()
}

func (s *ListPolicyResponseBodyPolicyListContentFingerPrintRuleList) GetDstPortEnd() *int32 {
	return s.DstPortEnd
}

func (s *ListPolicyResponseBodyPolicyListContentFingerPrintRuleList) GetDstPortStart() *int32 {
	return s.DstPortStart
}

func (s *ListPolicyResponseBodyPolicyListContentFingerPrintRuleList) GetId() *string {
	return s.Id
}

func (s *ListPolicyResponseBodyPolicyListContentFingerPrintRuleList) GetMatchAction() *string {
	return s.MatchAction
}

func (s *ListPolicyResponseBodyPolicyListContentFingerPrintRuleList) GetMaxPktLen() *int32 {
	return s.MaxPktLen
}

func (s *ListPolicyResponseBodyPolicyListContentFingerPrintRuleList) GetMinPktLen() *int32 {
	return s.MinPktLen
}

func (s *ListPolicyResponseBodyPolicyListContentFingerPrintRuleList) GetOffset() *int32 {
	return s.Offset
}

func (s *ListPolicyResponseBodyPolicyListContentFingerPrintRuleList) GetPayloadBytes() *string {
	return s.PayloadBytes
}

func (s *ListPolicyResponseBodyPolicyListContentFingerPrintRuleList) GetProtocol() *string {
	return s.Protocol
}

func (s *ListPolicyResponseBodyPolicyListContentFingerPrintRuleList) GetRateValue() *int32 {
	return s.RateValue
}

func (s *ListPolicyResponseBodyPolicyListContentFingerPrintRuleList) GetSeqNo() *int32 {
	return s.SeqNo
}

func (s *ListPolicyResponseBodyPolicyListContentFingerPrintRuleList) GetSrcPortEnd() *int32 {
	return s.SrcPortEnd
}

func (s *ListPolicyResponseBodyPolicyListContentFingerPrintRuleList) GetSrcPortStart() *int32 {
	return s.SrcPortStart
}

func (s *ListPolicyResponseBodyPolicyListContentFingerPrintRuleList) SetDstPortEnd(v int32) *ListPolicyResponseBodyPolicyListContentFingerPrintRuleList {
	s.DstPortEnd = &v
	return s
}

func (s *ListPolicyResponseBodyPolicyListContentFingerPrintRuleList) SetDstPortStart(v int32) *ListPolicyResponseBodyPolicyListContentFingerPrintRuleList {
	s.DstPortStart = &v
	return s
}

func (s *ListPolicyResponseBodyPolicyListContentFingerPrintRuleList) SetId(v string) *ListPolicyResponseBodyPolicyListContentFingerPrintRuleList {
	s.Id = &v
	return s
}

func (s *ListPolicyResponseBodyPolicyListContentFingerPrintRuleList) SetMatchAction(v string) *ListPolicyResponseBodyPolicyListContentFingerPrintRuleList {
	s.MatchAction = &v
	return s
}

func (s *ListPolicyResponseBodyPolicyListContentFingerPrintRuleList) SetMaxPktLen(v int32) *ListPolicyResponseBodyPolicyListContentFingerPrintRuleList {
	s.MaxPktLen = &v
	return s
}

func (s *ListPolicyResponseBodyPolicyListContentFingerPrintRuleList) SetMinPktLen(v int32) *ListPolicyResponseBodyPolicyListContentFingerPrintRuleList {
	s.MinPktLen = &v
	return s
}

func (s *ListPolicyResponseBodyPolicyListContentFingerPrintRuleList) SetOffset(v int32) *ListPolicyResponseBodyPolicyListContentFingerPrintRuleList {
	s.Offset = &v
	return s
}

func (s *ListPolicyResponseBodyPolicyListContentFingerPrintRuleList) SetPayloadBytes(v string) *ListPolicyResponseBodyPolicyListContentFingerPrintRuleList {
	s.PayloadBytes = &v
	return s
}

func (s *ListPolicyResponseBodyPolicyListContentFingerPrintRuleList) SetProtocol(v string) *ListPolicyResponseBodyPolicyListContentFingerPrintRuleList {
	s.Protocol = &v
	return s
}

func (s *ListPolicyResponseBodyPolicyListContentFingerPrintRuleList) SetRateValue(v int32) *ListPolicyResponseBodyPolicyListContentFingerPrintRuleList {
	s.RateValue = &v
	return s
}

func (s *ListPolicyResponseBodyPolicyListContentFingerPrintRuleList) SetSeqNo(v int32) *ListPolicyResponseBodyPolicyListContentFingerPrintRuleList {
	s.SeqNo = &v
	return s
}

func (s *ListPolicyResponseBodyPolicyListContentFingerPrintRuleList) SetSrcPortEnd(v int32) *ListPolicyResponseBodyPolicyListContentFingerPrintRuleList {
	s.SrcPortEnd = &v
	return s
}

func (s *ListPolicyResponseBodyPolicyListContentFingerPrintRuleList) SetSrcPortStart(v int32) *ListPolicyResponseBodyPolicyListContentFingerPrintRuleList {
	s.SrcPortStart = &v
	return s
}

func (s *ListPolicyResponseBodyPolicyListContentFingerPrintRuleList) Validate() error {
	return dara.Validate(s)
}

type ListPolicyResponseBodyPolicyListContentL4RuleList struct {
	// The action that is specified in the rule. Valid value:
	//
	// 	- **2**: The traffic is discarded.
	//
	// example:
	//
	// 2
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// The match conditions.
	ConditionList []*ListPolicyResponseBodyPolicyListContentL4RuleListConditionList `json:"ConditionList,omitempty" xml:"ConditionList,omitempty" type:"Repeated"`
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
	// 1
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
	// example:
	//
	// test**
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The priority of the rule.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
}

func (s ListPolicyResponseBodyPolicyListContentL4RuleList) String() string {
	return dara.Prettify(s)
}

func (s ListPolicyResponseBodyPolicyListContentL4RuleList) GoString() string {
	return s.String()
}

func (s *ListPolicyResponseBodyPolicyListContentL4RuleList) GetAction() *string {
	return s.Action
}

func (s *ListPolicyResponseBodyPolicyListContentL4RuleList) GetConditionList() []*ListPolicyResponseBodyPolicyListContentL4RuleListConditionList {
	return s.ConditionList
}

func (s *ListPolicyResponseBodyPolicyListContentL4RuleList) GetLimited() *int32 {
	return s.Limited
}

func (s *ListPolicyResponseBodyPolicyListContentL4RuleList) GetMatch() *string {
	return s.Match
}

func (s *ListPolicyResponseBodyPolicyListContentL4RuleList) GetMethod() *string {
	return s.Method
}

func (s *ListPolicyResponseBodyPolicyListContentL4RuleList) GetName() *string {
	return s.Name
}

func (s *ListPolicyResponseBodyPolicyListContentL4RuleList) GetPriority() *int32 {
	return s.Priority
}

func (s *ListPolicyResponseBodyPolicyListContentL4RuleList) SetAction(v string) *ListPolicyResponseBodyPolicyListContentL4RuleList {
	s.Action = &v
	return s
}

func (s *ListPolicyResponseBodyPolicyListContentL4RuleList) SetConditionList(v []*ListPolicyResponseBodyPolicyListContentL4RuleListConditionList) *ListPolicyResponseBodyPolicyListContentL4RuleList {
	s.ConditionList = v
	return s
}

func (s *ListPolicyResponseBodyPolicyListContentL4RuleList) SetLimited(v int32) *ListPolicyResponseBodyPolicyListContentL4RuleList {
	s.Limited = &v
	return s
}

func (s *ListPolicyResponseBodyPolicyListContentL4RuleList) SetMatch(v string) *ListPolicyResponseBodyPolicyListContentL4RuleList {
	s.Match = &v
	return s
}

func (s *ListPolicyResponseBodyPolicyListContentL4RuleList) SetMethod(v string) *ListPolicyResponseBodyPolicyListContentL4RuleList {
	s.Method = &v
	return s
}

func (s *ListPolicyResponseBodyPolicyListContentL4RuleList) SetName(v string) *ListPolicyResponseBodyPolicyListContentL4RuleList {
	s.Name = &v
	return s
}

func (s *ListPolicyResponseBodyPolicyListContentL4RuleList) SetPriority(v int32) *ListPolicyResponseBodyPolicyListContentL4RuleList {
	s.Priority = &v
	return s
}

func (s *ListPolicyResponseBodyPolicyListContentL4RuleList) Validate() error {
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

type ListPolicyResponseBodyPolicyListContentL4RuleListConditionList struct {
	// The term that is used for matching.
	//
	// >  If Method is set to **char**, the value of this parameter must be ASCII strings. If Method is set to **hex**, the value of this parameter must be hexadecimal strings. Maximum length: 2,048.
	//
	// example:
	//
	// test
	Arg     *string `json:"Arg,omitempty" xml:"Arg,omitempty"`
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The number of bytes from the start position for matching. Valid values: **1*	- to **2048**.
	//
	// example:
	//
	// 32
	Depth   *int32                                                                `json:"Depth,omitempty" xml:"Depth,omitempty"`
	Encode  *string                                                               `json:"Encode,omitempty" xml:"Encode,omitempty"`
	Offset  *ListPolicyResponseBodyPolicyListContentL4RuleListConditionListOffset `json:"Offset,omitempty" xml:"Offset,omitempty" type:"Struct"`
	Pattern *string                                                               `json:"Pattern,omitempty" xml:"Pattern,omitempty"`
	// The start position for matching. Valid values: **0*	- to **2047**.
	//
	// example:
	//
	// 0
	Position *int32 `json:"Position,omitempty" xml:"Position,omitempty"`
}

func (s ListPolicyResponseBodyPolicyListContentL4RuleListConditionList) String() string {
	return dara.Prettify(s)
}

func (s ListPolicyResponseBodyPolicyListContentL4RuleListConditionList) GoString() string {
	return s.String()
}

func (s *ListPolicyResponseBodyPolicyListContentL4RuleListConditionList) GetArg() *string {
	return s.Arg
}

func (s *ListPolicyResponseBodyPolicyListContentL4RuleListConditionList) GetContent() *string {
	return s.Content
}

func (s *ListPolicyResponseBodyPolicyListContentL4RuleListConditionList) GetDepth() *int32 {
	return s.Depth
}

func (s *ListPolicyResponseBodyPolicyListContentL4RuleListConditionList) GetEncode() *string {
	return s.Encode
}

func (s *ListPolicyResponseBodyPolicyListContentL4RuleListConditionList) GetOffset() *ListPolicyResponseBodyPolicyListContentL4RuleListConditionListOffset {
	return s.Offset
}

func (s *ListPolicyResponseBodyPolicyListContentL4RuleListConditionList) GetPattern() *string {
	return s.Pattern
}

func (s *ListPolicyResponseBodyPolicyListContentL4RuleListConditionList) GetPosition() *int32 {
	return s.Position
}

func (s *ListPolicyResponseBodyPolicyListContentL4RuleListConditionList) SetArg(v string) *ListPolicyResponseBodyPolicyListContentL4RuleListConditionList {
	s.Arg = &v
	return s
}

func (s *ListPolicyResponseBodyPolicyListContentL4RuleListConditionList) SetContent(v string) *ListPolicyResponseBodyPolicyListContentL4RuleListConditionList {
	s.Content = &v
	return s
}

func (s *ListPolicyResponseBodyPolicyListContentL4RuleListConditionList) SetDepth(v int32) *ListPolicyResponseBodyPolicyListContentL4RuleListConditionList {
	s.Depth = &v
	return s
}

func (s *ListPolicyResponseBodyPolicyListContentL4RuleListConditionList) SetEncode(v string) *ListPolicyResponseBodyPolicyListContentL4RuleListConditionList {
	s.Encode = &v
	return s
}

func (s *ListPolicyResponseBodyPolicyListContentL4RuleListConditionList) SetOffset(v *ListPolicyResponseBodyPolicyListContentL4RuleListConditionListOffset) *ListPolicyResponseBodyPolicyListContentL4RuleListConditionList {
	s.Offset = v
	return s
}

func (s *ListPolicyResponseBodyPolicyListContentL4RuleListConditionList) SetPattern(v string) *ListPolicyResponseBodyPolicyListContentL4RuleListConditionList {
	s.Pattern = &v
	return s
}

func (s *ListPolicyResponseBodyPolicyListContentL4RuleListConditionList) SetPosition(v int32) *ListPolicyResponseBodyPolicyListContentL4RuleListConditionList {
	s.Position = &v
	return s
}

func (s *ListPolicyResponseBodyPolicyListContentL4RuleListConditionList) Validate() error {
	if s.Offset != nil {
		if err := s.Offset.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListPolicyResponseBodyPolicyListContentL4RuleListConditionListOffset struct {
	End   *int32 `json:"End,omitempty" xml:"End,omitempty"`
	Start *int32 `json:"Start,omitempty" xml:"Start,omitempty"`
}

func (s ListPolicyResponseBodyPolicyListContentL4RuleListConditionListOffset) String() string {
	return dara.Prettify(s)
}

func (s ListPolicyResponseBodyPolicyListContentL4RuleListConditionListOffset) GoString() string {
	return s.String()
}

func (s *ListPolicyResponseBodyPolicyListContentL4RuleListConditionListOffset) GetEnd() *int32 {
	return s.End
}

func (s *ListPolicyResponseBodyPolicyListContentL4RuleListConditionListOffset) GetStart() *int32 {
	return s.Start
}

func (s *ListPolicyResponseBodyPolicyListContentL4RuleListConditionListOffset) SetEnd(v int32) *ListPolicyResponseBodyPolicyListContentL4RuleListConditionListOffset {
	s.End = &v
	return s
}

func (s *ListPolicyResponseBodyPolicyListContentL4RuleListConditionListOffset) SetStart(v int32) *ListPolicyResponseBodyPolicyListContentL4RuleListConditionListOffset {
	s.Start = &v
	return s
}

func (s *ListPolicyResponseBodyPolicyListContentL4RuleListConditionListOffset) Validate() error {
	return dara.Validate(s)
}

type ListPolicyResponseBodyPolicyListContentPortRuleList struct {
	// The end of the destination port range. Valid values: **0*	- to **65535**.
	//
	// example:
	//
	// 65535
	DstPortEnd *int32 `json:"DstPortEnd,omitempty" xml:"DstPortEnd,omitempty"`
	// The start of the destination port range. Valid values: **0*	- to **65535**.
	//
	// example:
	//
	// 0
	DstPortStart *int32 `json:"DstPortStart,omitempty" xml:"DstPortStart,omitempty"`
	// The ID of the rule.
	//
	// example:
	//
	// 8f3c3062-6c20-425d-8405-2bd1********
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The action triggered if the rule is matched. Valid value:
	//
	// 	- **drop**: The traffic is discarded.
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
	// example:
	//
	// udp
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The sequence number that indicates the order for the rule to take effect. The value is an integer.
	//
	// example:
	//
	// 1
	SeqNo *int32 `json:"SeqNo,omitempty" xml:"SeqNo,omitempty"`
	// The end of the source port range. Valid values: **0*	- to **65535**.
	//
	// example:
	//
	// 65535
	SrcPortEnd *int32 `json:"SrcPortEnd,omitempty" xml:"SrcPortEnd,omitempty"`
	// The start of the source port range. Valid values: **0*	- to **65535**.
	//
	// example:
	//
	// 0
	SrcPortStart *int32 `json:"SrcPortStart,omitempty" xml:"SrcPortStart,omitempty"`
}

func (s ListPolicyResponseBodyPolicyListContentPortRuleList) String() string {
	return dara.Prettify(s)
}

func (s ListPolicyResponseBodyPolicyListContentPortRuleList) GoString() string {
	return s.String()
}

func (s *ListPolicyResponseBodyPolicyListContentPortRuleList) GetDstPortEnd() *int32 {
	return s.DstPortEnd
}

func (s *ListPolicyResponseBodyPolicyListContentPortRuleList) GetDstPortStart() *int32 {
	return s.DstPortStart
}

func (s *ListPolicyResponseBodyPolicyListContentPortRuleList) GetId() *string {
	return s.Id
}

func (s *ListPolicyResponseBodyPolicyListContentPortRuleList) GetMatchAction() *string {
	return s.MatchAction
}

func (s *ListPolicyResponseBodyPolicyListContentPortRuleList) GetProtocol() *string {
	return s.Protocol
}

func (s *ListPolicyResponseBodyPolicyListContentPortRuleList) GetSeqNo() *int32 {
	return s.SeqNo
}

func (s *ListPolicyResponseBodyPolicyListContentPortRuleList) GetSrcPortEnd() *int32 {
	return s.SrcPortEnd
}

func (s *ListPolicyResponseBodyPolicyListContentPortRuleList) GetSrcPortStart() *int32 {
	return s.SrcPortStart
}

func (s *ListPolicyResponseBodyPolicyListContentPortRuleList) SetDstPortEnd(v int32) *ListPolicyResponseBodyPolicyListContentPortRuleList {
	s.DstPortEnd = &v
	return s
}

func (s *ListPolicyResponseBodyPolicyListContentPortRuleList) SetDstPortStart(v int32) *ListPolicyResponseBodyPolicyListContentPortRuleList {
	s.DstPortStart = &v
	return s
}

func (s *ListPolicyResponseBodyPolicyListContentPortRuleList) SetId(v string) *ListPolicyResponseBodyPolicyListContentPortRuleList {
	s.Id = &v
	return s
}

func (s *ListPolicyResponseBodyPolicyListContentPortRuleList) SetMatchAction(v string) *ListPolicyResponseBodyPolicyListContentPortRuleList {
	s.MatchAction = &v
	return s
}

func (s *ListPolicyResponseBodyPolicyListContentPortRuleList) SetProtocol(v string) *ListPolicyResponseBodyPolicyListContentPortRuleList {
	s.Protocol = &v
	return s
}

func (s *ListPolicyResponseBodyPolicyListContentPortRuleList) SetSeqNo(v int32) *ListPolicyResponseBodyPolicyListContentPortRuleList {
	s.SeqNo = &v
	return s
}

func (s *ListPolicyResponseBodyPolicyListContentPortRuleList) SetSrcPortEnd(v int32) *ListPolicyResponseBodyPolicyListContentPortRuleList {
	s.SrcPortEnd = &v
	return s
}

func (s *ListPolicyResponseBodyPolicyListContentPortRuleList) SetSrcPortStart(v int32) *ListPolicyResponseBodyPolicyListContentPortRuleList {
	s.SrcPortStart = &v
	return s
}

func (s *ListPolicyResponseBodyPolicyListContentPortRuleList) Validate() error {
	return dara.Validate(s)
}

type ListPolicyResponseBodyPolicyListContentSourceBlockList struct {
	// The validity period of the blacklist to which the source IP address is added. Unit: seconds.
	//
	// example:
	//
	// 120
	BlockExpireSeconds *int32 `json:"BlockExpireSeconds,omitempty" xml:"BlockExpireSeconds,omitempty"`
	// The statistical period during which the system collects data on source IP addresses to determine whether to add the source IP addresses to the blacklist. Unit: seconds.
	//
	// example:
	//
	// 60
	EverySeconds *int32 `json:"EverySeconds,omitempty" xml:"EverySeconds,omitempty"`
	// The number of times that the source IP address exceeds a limit in a statistical period.
	//
	// example:
	//
	// 5
	ExceedLimitTimes *int32 `json:"ExceedLimitTimes,omitempty" xml:"ExceedLimitTimes,omitempty"`
	// The type of the source rate limit. Valid values:
	//
	// 	- **3**: the PPS limit on source IP addresses.
	//
	// 	- **4**: the bandwidth limit on source IP addresses.
	//
	// 	- **5**: the PPS limit on source SYN packets.
	//
	// 	- **6**: the bandwidth limit on source SYN packets.
	//
	// example:
	//
	// 3
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListPolicyResponseBodyPolicyListContentSourceBlockList) String() string {
	return dara.Prettify(s)
}

func (s ListPolicyResponseBodyPolicyListContentSourceBlockList) GoString() string {
	return s.String()
}

func (s *ListPolicyResponseBodyPolicyListContentSourceBlockList) GetBlockExpireSeconds() *int32 {
	return s.BlockExpireSeconds
}

func (s *ListPolicyResponseBodyPolicyListContentSourceBlockList) GetEverySeconds() *int32 {
	return s.EverySeconds
}

func (s *ListPolicyResponseBodyPolicyListContentSourceBlockList) GetExceedLimitTimes() *int32 {
	return s.ExceedLimitTimes
}

func (s *ListPolicyResponseBodyPolicyListContentSourceBlockList) GetType() *int32 {
	return s.Type
}

func (s *ListPolicyResponseBodyPolicyListContentSourceBlockList) SetBlockExpireSeconds(v int32) *ListPolicyResponseBodyPolicyListContentSourceBlockList {
	s.BlockExpireSeconds = &v
	return s
}

func (s *ListPolicyResponseBodyPolicyListContentSourceBlockList) SetEverySeconds(v int32) *ListPolicyResponseBodyPolicyListContentSourceBlockList {
	s.EverySeconds = &v
	return s
}

func (s *ListPolicyResponseBodyPolicyListContentSourceBlockList) SetExceedLimitTimes(v int32) *ListPolicyResponseBodyPolicyListContentSourceBlockList {
	s.ExceedLimitTimes = &v
	return s
}

func (s *ListPolicyResponseBodyPolicyListContentSourceBlockList) SetType(v int32) *ListPolicyResponseBodyPolicyListContentSourceBlockList {
	s.Type = &v
	return s
}

func (s *ListPolicyResponseBodyPolicyListContentSourceBlockList) Validate() error {
	return dara.Validate(s)
}

type ListPolicyResponseBodyPolicyListContentSourceLimit struct {
	// The bandwidth limit on source IP addresses. Unit: bytes per second.
	//
	// example:
	//
	// 2048
	Bps *int32 `json:"Bps,omitempty" xml:"Bps,omitempty"`
	// The packets per second (PPS) limit on source IP addresses.
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
	// The PPS limit on source SYN packets.
	//
	// example:
	//
	// 64
	SynPps *int32 `json:"SynPps,omitempty" xml:"SynPps,omitempty"`
}

func (s ListPolicyResponseBodyPolicyListContentSourceLimit) String() string {
	return dara.Prettify(s)
}

func (s ListPolicyResponseBodyPolicyListContentSourceLimit) GoString() string {
	return s.String()
}

func (s *ListPolicyResponseBodyPolicyListContentSourceLimit) GetBps() *int32 {
	return s.Bps
}

func (s *ListPolicyResponseBodyPolicyListContentSourceLimit) GetPps() *int32 {
	return s.Pps
}

func (s *ListPolicyResponseBodyPolicyListContentSourceLimit) GetSynBps() *int32 {
	return s.SynBps
}

func (s *ListPolicyResponseBodyPolicyListContentSourceLimit) GetSynPps() *int32 {
	return s.SynPps
}

func (s *ListPolicyResponseBodyPolicyListContentSourceLimit) SetBps(v int32) *ListPolicyResponseBodyPolicyListContentSourceLimit {
	s.Bps = &v
	return s
}

func (s *ListPolicyResponseBodyPolicyListContentSourceLimit) SetPps(v int32) *ListPolicyResponseBodyPolicyListContentSourceLimit {
	s.Pps = &v
	return s
}

func (s *ListPolicyResponseBodyPolicyListContentSourceLimit) SetSynBps(v int32) *ListPolicyResponseBodyPolicyListContentSourceLimit {
	s.SynBps = &v
	return s
}

func (s *ListPolicyResponseBodyPolicyListContentSourceLimit) SetSynPps(v int32) *ListPolicyResponseBodyPolicyListContentSourceLimit {
	s.SynPps = &v
	return s
}

func (s *ListPolicyResponseBodyPolicyListContentSourceLimit) Validate() error {
	return dara.Validate(s)
}
