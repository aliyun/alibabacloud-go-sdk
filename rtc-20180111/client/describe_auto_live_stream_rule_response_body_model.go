// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAutoLiveStreamRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeAutoLiveStreamRuleResponseBody
	GetRequestId() *string
	SetRules(v []*DescribeAutoLiveStreamRuleResponseBodyRules) *DescribeAutoLiveStreamRuleResponseBody
	GetRules() []*DescribeAutoLiveStreamRuleResponseBodyRules
}

type DescribeAutoLiveStreamRuleResponseBody struct {
	// example:
	//
	// 069BCB66-CD80-11E8-A82B-A70F78BBDC00
	RequestId *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Rules     []*DescribeAutoLiveStreamRuleResponseBodyRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
}

func (s DescribeAutoLiveStreamRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutoLiveStreamRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAutoLiveStreamRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAutoLiveStreamRuleResponseBody) GetRules() []*DescribeAutoLiveStreamRuleResponseBodyRules {
	return s.Rules
}

func (s *DescribeAutoLiveStreamRuleResponseBody) SetRequestId(v string) *DescribeAutoLiveStreamRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAutoLiveStreamRuleResponseBody) SetRules(v []*DescribeAutoLiveStreamRuleResponseBodyRules) *DescribeAutoLiveStreamRuleResponseBody {
	s.Rules = v
	return s
}

func (s *DescribeAutoLiveStreamRuleResponseBody) Validate() error {
	if s.Rules != nil {
		for _, item := range s.Rules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAutoLiveStreamRuleResponseBodyRules struct {
	// example:
	//
	// http://example.com/callBack
	CallBack          *string   `json:"CallBack,omitempty" xml:"CallBack,omitempty"`
	ChannelIdPrefixes []*string `json:"ChannelIdPrefixes,omitempty" xml:"ChannelIdPrefixes,omitempty" type:"Repeated"`
	ChannelIds        []*string `json:"ChannelIds,omitempty" xml:"ChannelIds,omitempty" type:"Repeated"`
	// example:
	//
	// 2021-08-19T02:53:07Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 20
	MediaEncode *int32 `json:"MediaEncode,omitempty" xml:"MediaEncode,omitempty"`
	// example:
	//
	// rtmp://${domain}/${app}/${stream}
	PlayDomain *string `json:"PlayDomain,omitempty" xml:"PlayDomain,omitempty"`
	// example:
	//
	// 12
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// example:
	//
	// testRule
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// example:
	//
	// disable
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeAutoLiveStreamRuleResponseBodyRules) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutoLiveStreamRuleResponseBodyRules) GoString() string {
	return s.String()
}

func (s *DescribeAutoLiveStreamRuleResponseBodyRules) GetCallBack() *string {
	return s.CallBack
}

func (s *DescribeAutoLiveStreamRuleResponseBodyRules) GetChannelIdPrefixes() []*string {
	return s.ChannelIdPrefixes
}

func (s *DescribeAutoLiveStreamRuleResponseBodyRules) GetChannelIds() []*string {
	return s.ChannelIds
}

func (s *DescribeAutoLiveStreamRuleResponseBodyRules) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeAutoLiveStreamRuleResponseBodyRules) GetMediaEncode() *int32 {
	return s.MediaEncode
}

func (s *DescribeAutoLiveStreamRuleResponseBodyRules) GetPlayDomain() *string {
	return s.PlayDomain
}

func (s *DescribeAutoLiveStreamRuleResponseBodyRules) GetRuleId() *int64 {
	return s.RuleId
}

func (s *DescribeAutoLiveStreamRuleResponseBodyRules) GetRuleName() *string {
	return s.RuleName
}

func (s *DescribeAutoLiveStreamRuleResponseBodyRules) GetStatus() *string {
	return s.Status
}

func (s *DescribeAutoLiveStreamRuleResponseBodyRules) SetCallBack(v string) *DescribeAutoLiveStreamRuleResponseBodyRules {
	s.CallBack = &v
	return s
}

func (s *DescribeAutoLiveStreamRuleResponseBodyRules) SetChannelIdPrefixes(v []*string) *DescribeAutoLiveStreamRuleResponseBodyRules {
	s.ChannelIdPrefixes = v
	return s
}

func (s *DescribeAutoLiveStreamRuleResponseBodyRules) SetChannelIds(v []*string) *DescribeAutoLiveStreamRuleResponseBodyRules {
	s.ChannelIds = v
	return s
}

func (s *DescribeAutoLiveStreamRuleResponseBodyRules) SetCreateTime(v string) *DescribeAutoLiveStreamRuleResponseBodyRules {
	s.CreateTime = &v
	return s
}

func (s *DescribeAutoLiveStreamRuleResponseBodyRules) SetMediaEncode(v int32) *DescribeAutoLiveStreamRuleResponseBodyRules {
	s.MediaEncode = &v
	return s
}

func (s *DescribeAutoLiveStreamRuleResponseBodyRules) SetPlayDomain(v string) *DescribeAutoLiveStreamRuleResponseBodyRules {
	s.PlayDomain = &v
	return s
}

func (s *DescribeAutoLiveStreamRuleResponseBodyRules) SetRuleId(v int64) *DescribeAutoLiveStreamRuleResponseBodyRules {
	s.RuleId = &v
	return s
}

func (s *DescribeAutoLiveStreamRuleResponseBodyRules) SetRuleName(v string) *DescribeAutoLiveStreamRuleResponseBodyRules {
	s.RuleName = &v
	return s
}

func (s *DescribeAutoLiveStreamRuleResponseBodyRules) SetStatus(v string) *DescribeAutoLiveStreamRuleResponseBodyRules {
	s.Status = &v
	return s
}

func (s *DescribeAutoLiveStreamRuleResponseBodyRules) Validate() error {
	return dara.Validate(s)
}
