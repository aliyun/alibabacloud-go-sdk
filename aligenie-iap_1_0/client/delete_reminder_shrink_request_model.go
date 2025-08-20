// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteReminderShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceInfoShrink(v string) *DeleteReminderShrinkRequest
	GetDeviceInfoShrink() *string
	SetPayloadShrink(v string) *DeleteReminderShrinkRequest
	GetPayloadShrink() *string
	SetUserInfoShrink(v string) *DeleteReminderShrinkRequest
	GetUserInfoShrink() *string
}

type DeleteReminderShrinkRequest struct {
	// This parameter is required.
	DeviceInfoShrink *string `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty"`
	// This parameter is required.
	PayloadShrink *string `json:"Payload,omitempty" xml:"Payload,omitempty"`
	// This parameter is required.
	UserInfoShrink *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s DeleteReminderShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteReminderShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteReminderShrinkRequest) GetDeviceInfoShrink() *string {
	return s.DeviceInfoShrink
}

func (s *DeleteReminderShrinkRequest) GetPayloadShrink() *string {
	return s.PayloadShrink
}

func (s *DeleteReminderShrinkRequest) GetUserInfoShrink() *string {
	return s.UserInfoShrink
}

func (s *DeleteReminderShrinkRequest) SetDeviceInfoShrink(v string) *DeleteReminderShrinkRequest {
	s.DeviceInfoShrink = &v
	return s
}

func (s *DeleteReminderShrinkRequest) SetPayloadShrink(v string) *DeleteReminderShrinkRequest {
	s.PayloadShrink = &v
	return s
}

func (s *DeleteReminderShrinkRequest) SetUserInfoShrink(v string) *DeleteReminderShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

func (s *DeleteReminderShrinkRequest) Validate() error {
	return dara.Validate(s)
}
