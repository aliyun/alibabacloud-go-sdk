// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAlarmShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceInfoShrink(v string) *CreateAlarmShrinkRequest
	GetDeviceInfoShrink() *string
	SetPayloadShrink(v string) *CreateAlarmShrinkRequest
	GetPayloadShrink() *string
	SetUserInfoShrink(v string) *CreateAlarmShrinkRequest
	GetUserInfoShrink() *string
}

type CreateAlarmShrinkRequest struct {
	// This parameter is required.
	DeviceInfoShrink *string `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty"`
	// This parameter is required.
	PayloadShrink *string `json:"Payload,omitempty" xml:"Payload,omitempty"`
	// This parameter is required.
	UserInfoShrink *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s CreateAlarmShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAlarmShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateAlarmShrinkRequest) GetDeviceInfoShrink() *string {
	return s.DeviceInfoShrink
}

func (s *CreateAlarmShrinkRequest) GetPayloadShrink() *string {
	return s.PayloadShrink
}

func (s *CreateAlarmShrinkRequest) GetUserInfoShrink() *string {
	return s.UserInfoShrink
}

func (s *CreateAlarmShrinkRequest) SetDeviceInfoShrink(v string) *CreateAlarmShrinkRequest {
	s.DeviceInfoShrink = &v
	return s
}

func (s *CreateAlarmShrinkRequest) SetPayloadShrink(v string) *CreateAlarmShrinkRequest {
	s.PayloadShrink = &v
	return s
}

func (s *CreateAlarmShrinkRequest) SetUserInfoShrink(v string) *CreateAlarmShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

func (s *CreateAlarmShrinkRequest) Validate() error {
	return dara.Validate(s)
}
