// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAlarmsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceInfoShrink(v string) *DeleteAlarmsShrinkRequest
	GetDeviceInfoShrink() *string
	SetPayloadShrink(v string) *DeleteAlarmsShrinkRequest
	GetPayloadShrink() *string
	SetUserInfoShrink(v string) *DeleteAlarmsShrinkRequest
	GetUserInfoShrink() *string
}

type DeleteAlarmsShrinkRequest struct {
	// This parameter is required.
	DeviceInfoShrink *string `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty"`
	// This parameter is required.
	PayloadShrink *string `json:"Payload,omitempty" xml:"Payload,omitempty"`
	// This parameter is required.
	UserInfoShrink *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s DeleteAlarmsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAlarmsShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteAlarmsShrinkRequest) GetDeviceInfoShrink() *string {
	return s.DeviceInfoShrink
}

func (s *DeleteAlarmsShrinkRequest) GetPayloadShrink() *string {
	return s.PayloadShrink
}

func (s *DeleteAlarmsShrinkRequest) GetUserInfoShrink() *string {
	return s.UserInfoShrink
}

func (s *DeleteAlarmsShrinkRequest) SetDeviceInfoShrink(v string) *DeleteAlarmsShrinkRequest {
	s.DeviceInfoShrink = &v
	return s
}

func (s *DeleteAlarmsShrinkRequest) SetPayloadShrink(v string) *DeleteAlarmsShrinkRequest {
	s.PayloadShrink = &v
	return s
}

func (s *DeleteAlarmsShrinkRequest) SetUserInfoShrink(v string) *DeleteAlarmsShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

func (s *DeleteAlarmsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
