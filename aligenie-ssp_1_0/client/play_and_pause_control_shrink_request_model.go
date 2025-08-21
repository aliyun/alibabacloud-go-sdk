// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPlayAndPauseControlShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceInfoShrink(v string) *PlayAndPauseControlShrinkRequest
	GetDeviceInfoShrink() *string
	SetOpenPlayAndPauseControlParamShrink(v string) *PlayAndPauseControlShrinkRequest
	GetOpenPlayAndPauseControlParamShrink() *string
	SetUserInfoShrink(v string) *PlayAndPauseControlShrinkRequest
	GetUserInfoShrink() *string
}

type PlayAndPauseControlShrinkRequest struct {
	// This parameter is required.
	DeviceInfoShrink *string `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty"`
	// This parameter is required.
	OpenPlayAndPauseControlParamShrink *string `json:"OpenPlayAndPauseControlParam,omitempty" xml:"OpenPlayAndPauseControlParam,omitempty"`
	// This parameter is required.
	UserInfoShrink *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s PlayAndPauseControlShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s PlayAndPauseControlShrinkRequest) GoString() string {
	return s.String()
}

func (s *PlayAndPauseControlShrinkRequest) GetDeviceInfoShrink() *string {
	return s.DeviceInfoShrink
}

func (s *PlayAndPauseControlShrinkRequest) GetOpenPlayAndPauseControlParamShrink() *string {
	return s.OpenPlayAndPauseControlParamShrink
}

func (s *PlayAndPauseControlShrinkRequest) GetUserInfoShrink() *string {
	return s.UserInfoShrink
}

func (s *PlayAndPauseControlShrinkRequest) SetDeviceInfoShrink(v string) *PlayAndPauseControlShrinkRequest {
	s.DeviceInfoShrink = &v
	return s
}

func (s *PlayAndPauseControlShrinkRequest) SetOpenPlayAndPauseControlParamShrink(v string) *PlayAndPauseControlShrinkRequest {
	s.OpenPlayAndPauseControlParamShrink = &v
	return s
}

func (s *PlayAndPauseControlShrinkRequest) SetUserInfoShrink(v string) *PlayAndPauseControlShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

func (s *PlayAndPauseControlShrinkRequest) Validate() error {
	return dara.Validate(s)
}
