// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetReminderShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceInfoShrink(v string) *GetReminderShrinkRequest
	GetDeviceInfoShrink() *string
	SetPayloadShrink(v string) *GetReminderShrinkRequest
	GetPayloadShrink() *string
	SetUserInfoShrink(v string) *GetReminderShrinkRequest
	GetUserInfoShrink() *string
}

type GetReminderShrinkRequest struct {
	// This parameter is required.
	DeviceInfoShrink *string `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty"`
	// This parameter is required.
	PayloadShrink *string `json:"Payload,omitempty" xml:"Payload,omitempty"`
	// This parameter is required.
	UserInfoShrink *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s GetReminderShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetReminderShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetReminderShrinkRequest) GetDeviceInfoShrink() *string {
	return s.DeviceInfoShrink
}

func (s *GetReminderShrinkRequest) GetPayloadShrink() *string {
	return s.PayloadShrink
}

func (s *GetReminderShrinkRequest) GetUserInfoShrink() *string {
	return s.UserInfoShrink
}

func (s *GetReminderShrinkRequest) SetDeviceInfoShrink(v string) *GetReminderShrinkRequest {
	s.DeviceInfoShrink = &v
	return s
}

func (s *GetReminderShrinkRequest) SetPayloadShrink(v string) *GetReminderShrinkRequest {
	s.PayloadShrink = &v
	return s
}

func (s *GetReminderShrinkRequest) SetUserInfoShrink(v string) *GetReminderShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

func (s *GetReminderShrinkRequest) Validate() error {
	return dara.Validate(s)
}
