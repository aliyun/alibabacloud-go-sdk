// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendNotificationForPartnerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *SendNotificationForPartnerRequest
	GetBizId() *string
	SetChannelType(v string) *SendNotificationForPartnerRequest
	GetChannelType() *string
	SetNotifyType(v string) *SendNotificationForPartnerRequest
	GetNotifyType() *string
	SetNotifycationEventType(v string) *SendNotificationForPartnerRequest
	GetNotifycationEventType() *string
	SetParamMap(v map[string]*string) *SendNotificationForPartnerRequest
	GetParamMap() map[string]*string
	SetSendTarget(v string) *SendNotificationForPartnerRequest
	GetSendTarget() *string
	SetSmartSubChannels(v []*string) *SendNotificationForPartnerRequest
	GetSmartSubChannels() []*string
	SetTrackId(v string) *SendNotificationForPartnerRequest
	GetTrackId() *string
}

type SendNotificationForPartnerRequest struct {
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
	NotifycationEventType *string            `json:"NotifycationEventType,omitempty" xml:"NotifycationEventType,omitempty"`
	ParamMap              map[string]*string `json:"ParamMap,omitempty" xml:"ParamMap,omitempty"`
	// example:
	//
	// 12312212312
	SendTarget       *string   `json:"SendTarget,omitempty" xml:"SendTarget,omitempty"`
	SmartSubChannels []*string `json:"SmartSubChannels,omitempty" xml:"SmartSubChannels,omitempty" type:"Repeated"`
	// example:
	//
	// 5b29647n-a172-4ccd-ba33-73669896c287
	TrackId *string `json:"TrackId,omitempty" xml:"TrackId,omitempty"`
}

func (s SendNotificationForPartnerRequest) String() string {
	return dara.Prettify(s)
}

func (s SendNotificationForPartnerRequest) GoString() string {
	return s.String()
}

func (s *SendNotificationForPartnerRequest) GetBizId() *string {
	return s.BizId
}

func (s *SendNotificationForPartnerRequest) GetChannelType() *string {
	return s.ChannelType
}

func (s *SendNotificationForPartnerRequest) GetNotifyType() *string {
	return s.NotifyType
}

func (s *SendNotificationForPartnerRequest) GetNotifycationEventType() *string {
	return s.NotifycationEventType
}

func (s *SendNotificationForPartnerRequest) GetParamMap() map[string]*string {
	return s.ParamMap
}

func (s *SendNotificationForPartnerRequest) GetSendTarget() *string {
	return s.SendTarget
}

func (s *SendNotificationForPartnerRequest) GetSmartSubChannels() []*string {
	return s.SmartSubChannels
}

func (s *SendNotificationForPartnerRequest) GetTrackId() *string {
	return s.TrackId
}

func (s *SendNotificationForPartnerRequest) SetBizId(v string) *SendNotificationForPartnerRequest {
	s.BizId = &v
	return s
}

func (s *SendNotificationForPartnerRequest) SetChannelType(v string) *SendNotificationForPartnerRequest {
	s.ChannelType = &v
	return s
}

func (s *SendNotificationForPartnerRequest) SetNotifyType(v string) *SendNotificationForPartnerRequest {
	s.NotifyType = &v
	return s
}

func (s *SendNotificationForPartnerRequest) SetNotifycationEventType(v string) *SendNotificationForPartnerRequest {
	s.NotifycationEventType = &v
	return s
}

func (s *SendNotificationForPartnerRequest) SetParamMap(v map[string]*string) *SendNotificationForPartnerRequest {
	s.ParamMap = v
	return s
}

func (s *SendNotificationForPartnerRequest) SetSendTarget(v string) *SendNotificationForPartnerRequest {
	s.SendTarget = &v
	return s
}

func (s *SendNotificationForPartnerRequest) SetSmartSubChannels(v []*string) *SendNotificationForPartnerRequest {
	s.SmartSubChannels = v
	return s
}

func (s *SendNotificationForPartnerRequest) SetTrackId(v string) *SendNotificationForPartnerRequest {
	s.TrackId = &v
	return s
}

func (s *SendNotificationForPartnerRequest) Validate() error {
	return dara.Validate(s)
}
