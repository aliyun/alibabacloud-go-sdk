// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWeatherShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceInfoShrink(v string) *GetWeatherShrinkRequest
	GetDeviceInfoShrink() *string
	SetPayloadShrink(v string) *GetWeatherShrinkRequest
	GetPayloadShrink() *string
	SetUserInfoShrink(v string) *GetWeatherShrinkRequest
	GetUserInfoShrink() *string
}

type GetWeatherShrinkRequest struct {
	// This parameter is required.
	DeviceInfoShrink *string `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty"`
	// if can be null:
	// false
	PayloadShrink *string `json:"Payload,omitempty" xml:"Payload,omitempty"`
	// This parameter is required.
	UserInfoShrink *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s GetWeatherShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetWeatherShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetWeatherShrinkRequest) GetDeviceInfoShrink() *string {
	return s.DeviceInfoShrink
}

func (s *GetWeatherShrinkRequest) GetPayloadShrink() *string {
	return s.PayloadShrink
}

func (s *GetWeatherShrinkRequest) GetUserInfoShrink() *string {
	return s.UserInfoShrink
}

func (s *GetWeatherShrinkRequest) SetDeviceInfoShrink(v string) *GetWeatherShrinkRequest {
	s.DeviceInfoShrink = &v
	return s
}

func (s *GetWeatherShrinkRequest) SetPayloadShrink(v string) *GetWeatherShrinkRequest {
	s.PayloadShrink = &v
	return s
}

func (s *GetWeatherShrinkRequest) SetUserInfoShrink(v string) *GetWeatherShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

func (s *GetWeatherShrinkRequest) Validate() error {
	return dara.Validate(s)
}
