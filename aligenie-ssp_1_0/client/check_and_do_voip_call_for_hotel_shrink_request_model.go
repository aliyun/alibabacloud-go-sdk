// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckAndDoVoipCallForHotelShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizData(v string) *CheckAndDoVoipCallForHotelShrinkRequest
	GetBizData() *string
	SetCalleeNick(v string) *CheckAndDoVoipCallForHotelShrinkRequest
	GetCalleeNick() *string
	SetCalleePhoneNum(v string) *CheckAndDoVoipCallForHotelShrinkRequest
	GetCalleePhoneNum() *string
	SetDeviceInfoShrink(v string) *CheckAndDoVoipCallForHotelShrinkRequest
	GetDeviceInfoShrink() *string
	SetUserInfoShrink(v string) *CheckAndDoVoipCallForHotelShrinkRequest
	GetUserInfoShrink() *string
}

type CheckAndDoVoipCallForHotelShrinkRequest struct {
	BizData        *string `json:"BizData,omitempty" xml:"BizData,omitempty"`
	CalleeNick     *string `json:"CalleeNick,omitempty" xml:"CalleeNick,omitempty"`
	CalleePhoneNum *string `json:"CalleePhoneNum,omitempty" xml:"CalleePhoneNum,omitempty"`
	// This parameter is required.
	DeviceInfoShrink *string `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty"`
	// This parameter is required.
	UserInfoShrink *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s CheckAndDoVoipCallForHotelShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckAndDoVoipCallForHotelShrinkRequest) GoString() string {
	return s.String()
}

func (s *CheckAndDoVoipCallForHotelShrinkRequest) GetBizData() *string {
	return s.BizData
}

func (s *CheckAndDoVoipCallForHotelShrinkRequest) GetCalleeNick() *string {
	return s.CalleeNick
}

func (s *CheckAndDoVoipCallForHotelShrinkRequest) GetCalleePhoneNum() *string {
	return s.CalleePhoneNum
}

func (s *CheckAndDoVoipCallForHotelShrinkRequest) GetDeviceInfoShrink() *string {
	return s.DeviceInfoShrink
}

func (s *CheckAndDoVoipCallForHotelShrinkRequest) GetUserInfoShrink() *string {
	return s.UserInfoShrink
}

func (s *CheckAndDoVoipCallForHotelShrinkRequest) SetBizData(v string) *CheckAndDoVoipCallForHotelShrinkRequest {
	s.BizData = &v
	return s
}

func (s *CheckAndDoVoipCallForHotelShrinkRequest) SetCalleeNick(v string) *CheckAndDoVoipCallForHotelShrinkRequest {
	s.CalleeNick = &v
	return s
}

func (s *CheckAndDoVoipCallForHotelShrinkRequest) SetCalleePhoneNum(v string) *CheckAndDoVoipCallForHotelShrinkRequest {
	s.CalleePhoneNum = &v
	return s
}

func (s *CheckAndDoVoipCallForHotelShrinkRequest) SetDeviceInfoShrink(v string) *CheckAndDoVoipCallForHotelShrinkRequest {
	s.DeviceInfoShrink = &v
	return s
}

func (s *CheckAndDoVoipCallForHotelShrinkRequest) SetUserInfoShrink(v string) *CheckAndDoVoipCallForHotelShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

func (s *CheckAndDoVoipCallForHotelShrinkRequest) Validate() error {
	return dara.Validate(s)
}
