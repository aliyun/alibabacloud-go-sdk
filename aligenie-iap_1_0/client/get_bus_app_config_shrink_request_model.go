// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBusAppConfigShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceInfoShrink(v string) *GetBusAppConfigShrinkRequest
	GetDeviceInfoShrink() *string
	SetPayloadShrink(v string) *GetBusAppConfigShrinkRequest
	GetPayloadShrink() *string
	SetUserInfoShrink(v string) *GetBusAppConfigShrinkRequest
	GetUserInfoShrink() *string
}

type GetBusAppConfigShrinkRequest struct {
	DeviceInfoShrink *string `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty"`
	PayloadShrink    *string `json:"Payload,omitempty" xml:"Payload,omitempty"`
	// This parameter is required.
	UserInfoShrink *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s GetBusAppConfigShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetBusAppConfigShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetBusAppConfigShrinkRequest) GetDeviceInfoShrink() *string {
	return s.DeviceInfoShrink
}

func (s *GetBusAppConfigShrinkRequest) GetPayloadShrink() *string {
	return s.PayloadShrink
}

func (s *GetBusAppConfigShrinkRequest) GetUserInfoShrink() *string {
	return s.UserInfoShrink
}

func (s *GetBusAppConfigShrinkRequest) SetDeviceInfoShrink(v string) *GetBusAppConfigShrinkRequest {
	s.DeviceInfoShrink = &v
	return s
}

func (s *GetBusAppConfigShrinkRequest) SetPayloadShrink(v string) *GetBusAppConfigShrinkRequest {
	s.PayloadShrink = &v
	return s
}

func (s *GetBusAppConfigShrinkRequest) SetUserInfoShrink(v string) *GetBusAppConfigShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

func (s *GetBusAppConfigShrinkRequest) Validate() error {
	return dara.Validate(s)
}
