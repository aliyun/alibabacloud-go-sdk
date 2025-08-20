// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateReminderShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceInfoShrink(v string) *UpdateReminderShrinkRequest
	GetDeviceInfoShrink() *string
	SetPayloadShrink(v string) *UpdateReminderShrinkRequest
	GetPayloadShrink() *string
	SetUserInfoShrink(v string) *UpdateReminderShrinkRequest
	GetUserInfoShrink() *string
}

type UpdateReminderShrinkRequest struct {
	// This parameter is required.
	DeviceInfoShrink *string `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty"`
	// This parameter is required.
	PayloadShrink *string `json:"Payload,omitempty" xml:"Payload,omitempty"`
	// This parameter is required.
	UserInfoShrink *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s UpdateReminderShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateReminderShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateReminderShrinkRequest) GetDeviceInfoShrink() *string {
	return s.DeviceInfoShrink
}

func (s *UpdateReminderShrinkRequest) GetPayloadShrink() *string {
	return s.PayloadShrink
}

func (s *UpdateReminderShrinkRequest) GetUserInfoShrink() *string {
	return s.UserInfoShrink
}

func (s *UpdateReminderShrinkRequest) SetDeviceInfoShrink(v string) *UpdateReminderShrinkRequest {
	s.DeviceInfoShrink = &v
	return s
}

func (s *UpdateReminderShrinkRequest) SetPayloadShrink(v string) *UpdateReminderShrinkRequest {
	s.PayloadShrink = &v
	return s
}

func (s *UpdateReminderShrinkRequest) SetUserInfoShrink(v string) *UpdateReminderShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

func (s *UpdateReminderShrinkRequest) Validate() error {
	return dara.Validate(s)
}
