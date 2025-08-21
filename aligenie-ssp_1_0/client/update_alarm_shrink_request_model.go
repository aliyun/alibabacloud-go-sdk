// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAlarmShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceInfoShrink(v string) *UpdateAlarmShrinkRequest
	GetDeviceInfoShrink() *string
	SetPayloadShrink(v string) *UpdateAlarmShrinkRequest
	GetPayloadShrink() *string
	SetUserInfoShrink(v string) *UpdateAlarmShrinkRequest
	GetUserInfoShrink() *string
}

type UpdateAlarmShrinkRequest struct {
	// This parameter is required.
	DeviceInfoShrink *string `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty"`
	// This parameter is required.
	PayloadShrink *string `json:"Payload,omitempty" xml:"Payload,omitempty"`
	// This parameter is required.
	UserInfoShrink *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s UpdateAlarmShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAlarmShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateAlarmShrinkRequest) GetDeviceInfoShrink() *string {
	return s.DeviceInfoShrink
}

func (s *UpdateAlarmShrinkRequest) GetPayloadShrink() *string {
	return s.PayloadShrink
}

func (s *UpdateAlarmShrinkRequest) GetUserInfoShrink() *string {
	return s.UserInfoShrink
}

func (s *UpdateAlarmShrinkRequest) SetDeviceInfoShrink(v string) *UpdateAlarmShrinkRequest {
	s.DeviceInfoShrink = &v
	return s
}

func (s *UpdateAlarmShrinkRequest) SetPayloadShrink(v string) *UpdateAlarmShrinkRequest {
	s.PayloadShrink = &v
	return s
}

func (s *UpdateAlarmShrinkRequest) SetUserInfoShrink(v string) *UpdateAlarmShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

func (s *UpdateAlarmShrinkRequest) Validate() error {
	return dara.Validate(s)
}
