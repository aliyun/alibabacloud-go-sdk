// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendNotificationForPartnerShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *SendNotificationForPartnerShrinkRequest
	GetBizId() *string
	SetChannelType(v string) *SendNotificationForPartnerShrinkRequest
	GetChannelType() *string
	SetNotifyType(v string) *SendNotificationForPartnerShrinkRequest
	GetNotifyType() *string
	SetNotifycationEventType(v string) *SendNotificationForPartnerShrinkRequest
	GetNotifycationEventType() *string
	SetParamMapShrink(v string) *SendNotificationForPartnerShrinkRequest
	GetParamMapShrink() *string
	SetSendTarget(v string) *SendNotificationForPartnerShrinkRequest
	GetSendTarget() *string
	SetSmartSubChannelsShrink(v string) *SendNotificationForPartnerShrinkRequest
	GetSmartSubChannelsShrink() *string
	SetTrackId(v string) *SendNotificationForPartnerShrinkRequest
	GetTrackId() *string
}

type SendNotificationForPartnerShrinkRequest struct {
	// example:
	//
	// DMP
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// example:
	//
	// 1
	ChannelType *string `json:"ChannelType,omitempty" xml:"ChannelType,omitempty"`
	// example:
	//
	// MESSAGE
	NotifyType *string `json:"NotifyType,omitempty" xml:"NotifyType,omitempty"`
	// example:
	//
	// dmp_website_xxx
	NotifycationEventType *string `json:"NotifycationEventType,omitempty" xml:"NotifycationEventType,omitempty"`
	ParamMapShrink        *string `json:"ParamMap,omitempty" xml:"ParamMap,omitempty"`
	// example:
	//
	// 12312212312
	SendTarget             *string `json:"SendTarget,omitempty" xml:"SendTarget,omitempty"`
	SmartSubChannelsShrink *string `json:"SmartSubChannels,omitempty" xml:"SmartSubChannels,omitempty"`
	// example:
	//
	// 5b29647n-a172-4ccd-ba33-73669896c287
	TrackId *string `json:"TrackId,omitempty" xml:"TrackId,omitempty"`
}

func (s SendNotificationForPartnerShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SendNotificationForPartnerShrinkRequest) GoString() string {
	return s.String()
}

func (s *SendNotificationForPartnerShrinkRequest) GetBizId() *string {
	return s.BizId
}

func (s *SendNotificationForPartnerShrinkRequest) GetChannelType() *string {
	return s.ChannelType
}

func (s *SendNotificationForPartnerShrinkRequest) GetNotifyType() *string {
	return s.NotifyType
}

func (s *SendNotificationForPartnerShrinkRequest) GetNotifycationEventType() *string {
	return s.NotifycationEventType
}

func (s *SendNotificationForPartnerShrinkRequest) GetParamMapShrink() *string {
	return s.ParamMapShrink
}

func (s *SendNotificationForPartnerShrinkRequest) GetSendTarget() *string {
	return s.SendTarget
}

func (s *SendNotificationForPartnerShrinkRequest) GetSmartSubChannelsShrink() *string {
	return s.SmartSubChannelsShrink
}

func (s *SendNotificationForPartnerShrinkRequest) GetTrackId() *string {
	return s.TrackId
}

func (s *SendNotificationForPartnerShrinkRequest) SetBizId(v string) *SendNotificationForPartnerShrinkRequest {
	s.BizId = &v
	return s
}

func (s *SendNotificationForPartnerShrinkRequest) SetChannelType(v string) *SendNotificationForPartnerShrinkRequest {
	s.ChannelType = &v
	return s
}

func (s *SendNotificationForPartnerShrinkRequest) SetNotifyType(v string) *SendNotificationForPartnerShrinkRequest {
	s.NotifyType = &v
	return s
}

func (s *SendNotificationForPartnerShrinkRequest) SetNotifycationEventType(v string) *SendNotificationForPartnerShrinkRequest {
	s.NotifycationEventType = &v
	return s
}

func (s *SendNotificationForPartnerShrinkRequest) SetParamMapShrink(v string) *SendNotificationForPartnerShrinkRequest {
	s.ParamMapShrink = &v
	return s
}

func (s *SendNotificationForPartnerShrinkRequest) SetSendTarget(v string) *SendNotificationForPartnerShrinkRequest {
	s.SendTarget = &v
	return s
}

func (s *SendNotificationForPartnerShrinkRequest) SetSmartSubChannelsShrink(v string) *SendNotificationForPartnerShrinkRequest {
	s.SmartSubChannelsShrink = &v
	return s
}

func (s *SendNotificationForPartnerShrinkRequest) SetTrackId(v string) *SendNotificationForPartnerShrinkRequest {
	s.TrackId = &v
	return s
}

func (s *SendNotificationForPartnerShrinkRequest) Validate() error {
	return dara.Validate(s)
}
