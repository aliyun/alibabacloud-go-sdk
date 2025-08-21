// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindDeviceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceInfoShrink(v string) *UnbindDeviceShrinkRequest
	GetDeviceInfoShrink() *string
	SetUserInfoShrink(v string) *UnbindDeviceShrinkRequest
	GetUserInfoShrink() *string
}

type UnbindDeviceShrinkRequest struct {
	// This parameter is required.
	DeviceInfoShrink *string `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty"`
	// This parameter is required.
	UserInfoShrink *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s UnbindDeviceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UnbindDeviceShrinkRequest) GoString() string {
	return s.String()
}

func (s *UnbindDeviceShrinkRequest) GetDeviceInfoShrink() *string {
	return s.DeviceInfoShrink
}

func (s *UnbindDeviceShrinkRequest) GetUserInfoShrink() *string {
	return s.UserInfoShrink
}

func (s *UnbindDeviceShrinkRequest) SetDeviceInfoShrink(v string) *UnbindDeviceShrinkRequest {
	s.DeviceInfoShrink = &v
	return s
}

func (s *UnbindDeviceShrinkRequest) SetUserInfoShrink(v string) *UnbindDeviceShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

func (s *UnbindDeviceShrinkRequest) Validate() error {
	return dara.Validate(s)
}
