// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMobileRecommendShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBotId(v string) *MobileRecommendShrinkRequest
	GetBotId() *string
	SetCount(v string) *MobileRecommendShrinkRequest
	GetCount() *string
	SetDeviceInfoShrink(v string) *MobileRecommendShrinkRequest
	GetDeviceInfoShrink() *string
	SetStyle(v string) *MobileRecommendShrinkRequest
	GetStyle() *string
	SetType(v string) *MobileRecommendShrinkRequest
	GetType() *string
	SetUserInfoShrink(v string) *MobileRecommendShrinkRequest
	GetUserInfoShrink() *string
}

type MobileRecommendShrinkRequest struct {
	// example:
	//
	// 10
	BotId *string `json:"BotId,omitempty" xml:"BotId,omitempty"`
	// example:
	//
	// 6
	Count *string `json:"Count,omitempty" xml:"Count,omitempty"`
	// This parameter is required.
	DeviceInfoShrink *string `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty"`
	Style            *string `json:"Style,omitempty" xml:"Style,omitempty"`
	// example:
	//
	// DAILY_REC
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// This parameter is required.
	UserInfoShrink *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s MobileRecommendShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s MobileRecommendShrinkRequest) GoString() string {
	return s.String()
}

func (s *MobileRecommendShrinkRequest) GetBotId() *string {
	return s.BotId
}

func (s *MobileRecommendShrinkRequest) GetCount() *string {
	return s.Count
}

func (s *MobileRecommendShrinkRequest) GetDeviceInfoShrink() *string {
	return s.DeviceInfoShrink
}

func (s *MobileRecommendShrinkRequest) GetStyle() *string {
	return s.Style
}

func (s *MobileRecommendShrinkRequest) GetType() *string {
	return s.Type
}

func (s *MobileRecommendShrinkRequest) GetUserInfoShrink() *string {
	return s.UserInfoShrink
}

func (s *MobileRecommendShrinkRequest) SetBotId(v string) *MobileRecommendShrinkRequest {
	s.BotId = &v
	return s
}

func (s *MobileRecommendShrinkRequest) SetCount(v string) *MobileRecommendShrinkRequest {
	s.Count = &v
	return s
}

func (s *MobileRecommendShrinkRequest) SetDeviceInfoShrink(v string) *MobileRecommendShrinkRequest {
	s.DeviceInfoShrink = &v
	return s
}

func (s *MobileRecommendShrinkRequest) SetStyle(v string) *MobileRecommendShrinkRequest {
	s.Style = &v
	return s
}

func (s *MobileRecommendShrinkRequest) SetType(v string) *MobileRecommendShrinkRequest {
	s.Type = &v
	return s
}

func (s *MobileRecommendShrinkRequest) SetUserInfoShrink(v string) *MobileRecommendShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

func (s *MobileRecommendShrinkRequest) Validate() error {
	return dara.Validate(s)
}
