// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAutoLiveStreamRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *CreateAutoLiveStreamRuleRequest
	GetAppId() *string
	SetCallBack(v string) *CreateAutoLiveStreamRuleRequest
	GetCallBack() *string
	SetChannelIdPrefixes(v []*string) *CreateAutoLiveStreamRuleRequest
	GetChannelIdPrefixes() []*string
	SetChannelIds(v []*string) *CreateAutoLiveStreamRuleRequest
	GetChannelIds() []*string
	SetMediaEncode(v int32) *CreateAutoLiveStreamRuleRequest
	GetMediaEncode() *int32
	SetOwnerId(v int64) *CreateAutoLiveStreamRuleRequest
	GetOwnerId() *int64
	SetPlayDomain(v string) *CreateAutoLiveStreamRuleRequest
	GetPlayDomain() *string
	SetRuleName(v string) *CreateAutoLiveStreamRuleRequest
	GetRuleName() *string
}

type CreateAutoLiveStreamRuleRequest struct {
	// The application ID, which can be created and queried in the console.
	//
	// This parameter is required.
	//
	// example:
	//
	// eo85****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The callback URL used to notify the start and end of the relay. For details about the callback content, see the CallBack callback example below.
	//
	// example:
	//
	// http://example.com/callback
	CallBack *string `json:"CallBack,omitempty" xml:"CallBack,omitempty"`
	// The list of channel ID prefixes that are fuzzy matched.
	ChannelIdPrefixes []*string `json:"ChannelIdPrefixes,omitempty" xml:"ChannelIdPrefixes,omitempty" type:"Repeated"`
	// The list of channel IDs that are exactly matched.
	ChannelIds []*string `json:"ChannelIds,omitempty" xml:"ChannelIds,omitempty" type:"Repeated"`
	// The encoding option. For more information, see the MediaEncode enumeration values below.
	//
	// example:
	//
	// 20
	MediaEncode *int32 `json:"MediaEncode,omitempty" xml:"MediaEncode,omitempty"`
	OwnerId     *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The streaming domain name for automatic relay. The streaming URL of an automatic relay task is in the format rtmp://${domain}/${app}/${stream}.
	//
	// - domain: your live streaming domain name.
	//
	// - app: live.
	//
	// - stream: urlencode(${AppId}_${ChannelId}_${UserId}_${SourceType}), where SourceType indicates the type of the input source.
	//
	// This parameter is required.
	//
	// example:
	//
	// rtmp://${domain}/${app}/${stream}
	PlayDomain *string `json:"PlayDomain,omitempty" xml:"PlayDomain,omitempty"`
	// The name of the relay rule.
	//
	// example:
	//
	// testRule
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s CreateAutoLiveStreamRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAutoLiveStreamRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateAutoLiveStreamRuleRequest) GetAppId() *string {
	return s.AppId
}

func (s *CreateAutoLiveStreamRuleRequest) GetCallBack() *string {
	return s.CallBack
}

func (s *CreateAutoLiveStreamRuleRequest) GetChannelIdPrefixes() []*string {
	return s.ChannelIdPrefixes
}

func (s *CreateAutoLiveStreamRuleRequest) GetChannelIds() []*string {
	return s.ChannelIds
}

func (s *CreateAutoLiveStreamRuleRequest) GetMediaEncode() *int32 {
	return s.MediaEncode
}

func (s *CreateAutoLiveStreamRuleRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateAutoLiveStreamRuleRequest) GetPlayDomain() *string {
	return s.PlayDomain
}

func (s *CreateAutoLiveStreamRuleRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *CreateAutoLiveStreamRuleRequest) SetAppId(v string) *CreateAutoLiveStreamRuleRequest {
	s.AppId = &v
	return s
}

func (s *CreateAutoLiveStreamRuleRequest) SetCallBack(v string) *CreateAutoLiveStreamRuleRequest {
	s.CallBack = &v
	return s
}

func (s *CreateAutoLiveStreamRuleRequest) SetChannelIdPrefixes(v []*string) *CreateAutoLiveStreamRuleRequest {
	s.ChannelIdPrefixes = v
	return s
}

func (s *CreateAutoLiveStreamRuleRequest) SetChannelIds(v []*string) *CreateAutoLiveStreamRuleRequest {
	s.ChannelIds = v
	return s
}

func (s *CreateAutoLiveStreamRuleRequest) SetMediaEncode(v int32) *CreateAutoLiveStreamRuleRequest {
	s.MediaEncode = &v
	return s
}

func (s *CreateAutoLiveStreamRuleRequest) SetOwnerId(v int64) *CreateAutoLiveStreamRuleRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateAutoLiveStreamRuleRequest) SetPlayDomain(v string) *CreateAutoLiveStreamRuleRequest {
	s.PlayDomain = &v
	return s
}

func (s *CreateAutoLiveStreamRuleRequest) SetRuleName(v string) *CreateAutoLiveStreamRuleRequest {
	s.RuleName = &v
	return s
}

func (s *CreateAutoLiveStreamRuleRequest) Validate() error {
	return dara.Validate(s)
}
