// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRemindersShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceInfoShrink(v string) *ListRemindersShrinkRequest
	GetDeviceInfoShrink() *string
	SetPayloadShrink(v string) *ListRemindersShrinkRequest
	GetPayloadShrink() *string
	SetUserInfoShrink(v string) *ListRemindersShrinkRequest
	GetUserInfoShrink() *string
}

type ListRemindersShrinkRequest struct {
	// This parameter is required.
	DeviceInfoShrink *string `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty"`
	// This parameter is required.
	PayloadShrink *string `json:"Payload,omitempty" xml:"Payload,omitempty"`
	// This parameter is required.
	UserInfoShrink *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s ListRemindersShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRemindersShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListRemindersShrinkRequest) GetDeviceInfoShrink() *string {
	return s.DeviceInfoShrink
}

func (s *ListRemindersShrinkRequest) GetPayloadShrink() *string {
	return s.PayloadShrink
}

func (s *ListRemindersShrinkRequest) GetUserInfoShrink() *string {
	return s.UserInfoShrink
}

func (s *ListRemindersShrinkRequest) SetDeviceInfoShrink(v string) *ListRemindersShrinkRequest {
	s.DeviceInfoShrink = &v
	return s
}

func (s *ListRemindersShrinkRequest) SetPayloadShrink(v string) *ListRemindersShrinkRequest {
	s.PayloadShrink = &v
	return s
}

func (s *ListRemindersShrinkRequest) SetUserInfoShrink(v string) *ListRemindersShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

func (s *ListRemindersShrinkRequest) Validate() error {
	return dara.Validate(s)
}
