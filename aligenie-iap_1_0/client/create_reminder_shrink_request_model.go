// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateReminderShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceInfoShrink(v string) *CreateReminderShrinkRequest
	GetDeviceInfoShrink() *string
	SetPayloadShrink(v string) *CreateReminderShrinkRequest
	GetPayloadShrink() *string
	SetUserInfoShrink(v string) *CreateReminderShrinkRequest
	GetUserInfoShrink() *string
}

type CreateReminderShrinkRequest struct {
	// This parameter is required.
	DeviceInfoShrink *string `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty"`
	// This parameter is required.
	PayloadShrink *string `json:"Payload,omitempty" xml:"Payload,omitempty"`
	// This parameter is required.
	UserInfoShrink *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s CreateReminderShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateReminderShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateReminderShrinkRequest) GetDeviceInfoShrink() *string {
	return s.DeviceInfoShrink
}

func (s *CreateReminderShrinkRequest) GetPayloadShrink() *string {
	return s.PayloadShrink
}

func (s *CreateReminderShrinkRequest) GetUserInfoShrink() *string {
	return s.UserInfoShrink
}

func (s *CreateReminderShrinkRequest) SetDeviceInfoShrink(v string) *CreateReminderShrinkRequest {
	s.DeviceInfoShrink = &v
	return s
}

func (s *CreateReminderShrinkRequest) SetPayloadShrink(v string) *CreateReminderShrinkRequest {
	s.PayloadShrink = &v
	return s
}

func (s *CreateReminderShrinkRequest) SetUserInfoShrink(v string) *CreateReminderShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

func (s *CreateReminderShrinkRequest) Validate() error {
	return dara.Validate(s)
}
