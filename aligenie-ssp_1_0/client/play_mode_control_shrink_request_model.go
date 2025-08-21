// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPlayModeControlShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceInfoShrink(v string) *PlayModeControlShrinkRequest
	GetDeviceInfoShrink() *string
	SetOpenPlayModeControlRequestShrink(v string) *PlayModeControlShrinkRequest
	GetOpenPlayModeControlRequestShrink() *string
	SetUserInfoShrink(v string) *PlayModeControlShrinkRequest
	GetUserInfoShrink() *string
}

type PlayModeControlShrinkRequest struct {
	// This parameter is required.
	DeviceInfoShrink *string `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty"`
	// This parameter is required.
	OpenPlayModeControlRequestShrink *string `json:"OpenPlayModeControlRequest,omitempty" xml:"OpenPlayModeControlRequest,omitempty"`
	// This parameter is required.
	UserInfoShrink *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s PlayModeControlShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s PlayModeControlShrinkRequest) GoString() string {
	return s.String()
}

func (s *PlayModeControlShrinkRequest) GetDeviceInfoShrink() *string {
	return s.DeviceInfoShrink
}

func (s *PlayModeControlShrinkRequest) GetOpenPlayModeControlRequestShrink() *string {
	return s.OpenPlayModeControlRequestShrink
}

func (s *PlayModeControlShrinkRequest) GetUserInfoShrink() *string {
	return s.UserInfoShrink
}

func (s *PlayModeControlShrinkRequest) SetDeviceInfoShrink(v string) *PlayModeControlShrinkRequest {
	s.DeviceInfoShrink = &v
	return s
}

func (s *PlayModeControlShrinkRequest) SetOpenPlayModeControlRequestShrink(v string) *PlayModeControlShrinkRequest {
	s.OpenPlayModeControlRequestShrink = &v
	return s
}

func (s *PlayModeControlShrinkRequest) SetUserInfoShrink(v string) *PlayModeControlShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

func (s *PlayModeControlShrinkRequest) Validate() error {
	return dara.Validate(s)
}
