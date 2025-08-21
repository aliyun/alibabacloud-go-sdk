// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iProgressControlShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceInfoShrink(v string) *ProgressControlShrinkRequest
	GetDeviceInfoShrink() *string
	SetOpenProgressControlRequestShrink(v string) *ProgressControlShrinkRequest
	GetOpenProgressControlRequestShrink() *string
	SetUserInfoShrink(v string) *ProgressControlShrinkRequest
	GetUserInfoShrink() *string
}

type ProgressControlShrinkRequest struct {
	// This parameter is required.
	DeviceInfoShrink *string `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty"`
	// This parameter is required.
	OpenProgressControlRequestShrink *string `json:"OpenProgressControlRequest,omitempty" xml:"OpenProgressControlRequest,omitempty"`
	// This parameter is required.
	UserInfoShrink *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s ProgressControlShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ProgressControlShrinkRequest) GoString() string {
	return s.String()
}

func (s *ProgressControlShrinkRequest) GetDeviceInfoShrink() *string {
	return s.DeviceInfoShrink
}

func (s *ProgressControlShrinkRequest) GetOpenProgressControlRequestShrink() *string {
	return s.OpenProgressControlRequestShrink
}

func (s *ProgressControlShrinkRequest) GetUserInfoShrink() *string {
	return s.UserInfoShrink
}

func (s *ProgressControlShrinkRequest) SetDeviceInfoShrink(v string) *ProgressControlShrinkRequest {
	s.DeviceInfoShrink = &v
	return s
}

func (s *ProgressControlShrinkRequest) SetOpenProgressControlRequestShrink(v string) *ProgressControlShrinkRequest {
	s.OpenProgressControlRequestShrink = &v
	return s
}

func (s *ProgressControlShrinkRequest) SetUserInfoShrink(v string) *ProgressControlShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

func (s *ProgressControlShrinkRequest) Validate() error {
	return dara.Validate(s)
}
