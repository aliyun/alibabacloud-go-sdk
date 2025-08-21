// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAlarmShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceInfoShrink(v string) *GetAlarmShrinkRequest
	GetDeviceInfoShrink() *string
	SetPayloadShrink(v string) *GetAlarmShrinkRequest
	GetPayloadShrink() *string
	SetUserInfoShrink(v string) *GetAlarmShrinkRequest
	GetUserInfoShrink() *string
}

type GetAlarmShrinkRequest struct {
	// This parameter is required.
	DeviceInfoShrink *string `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty"`
	// This parameter is required.
	PayloadShrink *string `json:"Payload,omitempty" xml:"Payload,omitempty"`
	// This parameter is required.
	UserInfoShrink *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s GetAlarmShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAlarmShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetAlarmShrinkRequest) GetDeviceInfoShrink() *string {
	return s.DeviceInfoShrink
}

func (s *GetAlarmShrinkRequest) GetPayloadShrink() *string {
	return s.PayloadShrink
}

func (s *GetAlarmShrinkRequest) GetUserInfoShrink() *string {
	return s.UserInfoShrink
}

func (s *GetAlarmShrinkRequest) SetDeviceInfoShrink(v string) *GetAlarmShrinkRequest {
	s.DeviceInfoShrink = &v
	return s
}

func (s *GetAlarmShrinkRequest) SetPayloadShrink(v string) *GetAlarmShrinkRequest {
	s.PayloadShrink = &v
	return s
}

func (s *GetAlarmShrinkRequest) SetUserInfoShrink(v string) *GetAlarmShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

func (s *GetAlarmShrinkRequest) Validate() error {
	return dara.Validate(s)
}
