// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAlarmsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceInfoShrink(v string) *ListAlarmsShrinkRequest
	GetDeviceInfoShrink() *string
	SetPayloadShrink(v string) *ListAlarmsShrinkRequest
	GetPayloadShrink() *string
	SetUserInfoShrink(v string) *ListAlarmsShrinkRequest
	GetUserInfoShrink() *string
}

type ListAlarmsShrinkRequest struct {
	// This parameter is required.
	DeviceInfoShrink *string `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty"`
	// This parameter is required.
	PayloadShrink *string `json:"Payload,omitempty" xml:"Payload,omitempty"`
	// This parameter is required.
	UserInfoShrink *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s ListAlarmsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAlarmsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListAlarmsShrinkRequest) GetDeviceInfoShrink() *string {
	return s.DeviceInfoShrink
}

func (s *ListAlarmsShrinkRequest) GetPayloadShrink() *string {
	return s.PayloadShrink
}

func (s *ListAlarmsShrinkRequest) GetUserInfoShrink() *string {
	return s.UserInfoShrink
}

func (s *ListAlarmsShrinkRequest) SetDeviceInfoShrink(v string) *ListAlarmsShrinkRequest {
	s.DeviceInfoShrink = &v
	return s
}

func (s *ListAlarmsShrinkRequest) SetPayloadShrink(v string) *ListAlarmsShrinkRequest {
	s.PayloadShrink = &v
	return s
}

func (s *ListAlarmsShrinkRequest) SetUserInfoShrink(v string) *ListAlarmsShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

func (s *ListAlarmsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
