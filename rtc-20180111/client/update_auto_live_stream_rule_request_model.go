// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAutoLiveStreamRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *UpdateAutoLiveStreamRuleRequest
	GetAppId() *string
	SetCallBack(v string) *UpdateAutoLiveStreamRuleRequest
	GetCallBack() *string
	SetChannelIdPrefixes(v []*string) *UpdateAutoLiveStreamRuleRequest
	GetChannelIdPrefixes() []*string
	SetChannelIds(v []*string) *UpdateAutoLiveStreamRuleRequest
	GetChannelIds() []*string
	SetMediaEncode(v int32) *UpdateAutoLiveStreamRuleRequest
	GetMediaEncode() *int32
	SetOwnerId(v int64) *UpdateAutoLiveStreamRuleRequest
	GetOwnerId() *int64
	SetPlayDomain(v string) *UpdateAutoLiveStreamRuleRequest
	GetPlayDomain() *string
	SetRuleId(v int32) *UpdateAutoLiveStreamRuleRequest
	GetRuleId() *int32
	SetRuleName(v string) *UpdateAutoLiveStreamRuleRequest
	GetRuleName() *string
}

type UpdateAutoLiveStreamRuleRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// eo85****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// http://example.com/callback
	CallBack          *string   `json:"CallBack,omitempty" xml:"CallBack,omitempty"`
	ChannelIdPrefixes []*string `json:"ChannelIdPrefixes,omitempty" xml:"ChannelIdPrefixes,omitempty" type:"Repeated"`
	ChannelIds        []*string `json:"ChannelIds,omitempty" xml:"ChannelIds,omitempty" type:"Repeated"`
	// example:
	//
	// 20
	MediaEncode *int32 `json:"MediaEncode,omitempty" xml:"MediaEncode,omitempty"`
	OwnerId     *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// rtmp://${domain}/${app}/${stream}
	PlayDomain *string `json:"PlayDomain,omitempty" xml:"PlayDomain,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 12
	RuleId *int32 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// example:
	//
	// testRule
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s UpdateAutoLiveStreamRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAutoLiveStreamRuleRequest) GoString() string {
	return s.String()
}

func (s *UpdateAutoLiveStreamRuleRequest) GetAppId() *string {
	return s.AppId
}

func (s *UpdateAutoLiveStreamRuleRequest) GetCallBack() *string {
	return s.CallBack
}

func (s *UpdateAutoLiveStreamRuleRequest) GetChannelIdPrefixes() []*string {
	return s.ChannelIdPrefixes
}

func (s *UpdateAutoLiveStreamRuleRequest) GetChannelIds() []*string {
	return s.ChannelIds
}

func (s *UpdateAutoLiveStreamRuleRequest) GetMediaEncode() *int32 {
	return s.MediaEncode
}

func (s *UpdateAutoLiveStreamRuleRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateAutoLiveStreamRuleRequest) GetPlayDomain() *string {
	return s.PlayDomain
}

func (s *UpdateAutoLiveStreamRuleRequest) GetRuleId() *int32 {
	return s.RuleId
}

func (s *UpdateAutoLiveStreamRuleRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *UpdateAutoLiveStreamRuleRequest) SetAppId(v string) *UpdateAutoLiveStreamRuleRequest {
	s.AppId = &v
	return s
}

func (s *UpdateAutoLiveStreamRuleRequest) SetCallBack(v string) *UpdateAutoLiveStreamRuleRequest {
	s.CallBack = &v
	return s
}

func (s *UpdateAutoLiveStreamRuleRequest) SetChannelIdPrefixes(v []*string) *UpdateAutoLiveStreamRuleRequest {
	s.ChannelIdPrefixes = v
	return s
}

func (s *UpdateAutoLiveStreamRuleRequest) SetChannelIds(v []*string) *UpdateAutoLiveStreamRuleRequest {
	s.ChannelIds = v
	return s
}

func (s *UpdateAutoLiveStreamRuleRequest) SetMediaEncode(v int32) *UpdateAutoLiveStreamRuleRequest {
	s.MediaEncode = &v
	return s
}

func (s *UpdateAutoLiveStreamRuleRequest) SetOwnerId(v int64) *UpdateAutoLiveStreamRuleRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateAutoLiveStreamRuleRequest) SetPlayDomain(v string) *UpdateAutoLiveStreamRuleRequest {
	s.PlayDomain = &v
	return s
}

func (s *UpdateAutoLiveStreamRuleRequest) SetRuleId(v int32) *UpdateAutoLiveStreamRuleRequest {
	s.RuleId = &v
	return s
}

func (s *UpdateAutoLiveStreamRuleRequest) SetRuleName(v string) *UpdateAutoLiveStreamRuleRequest {
	s.RuleName = &v
	return s
}

func (s *UpdateAutoLiveStreamRuleRequest) Validate() error {
	return dara.Validate(s)
}
