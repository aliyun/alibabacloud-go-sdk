// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeviceControlShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPayloadShrink(v string) *DeviceControlShrinkRequest
	GetPayloadShrink() *string
	SetUserInfoShrink(v string) *DeviceControlShrinkRequest
	GetUserInfoShrink() *string
}

type DeviceControlShrinkRequest struct {
	PayloadShrink  *string `json:"Payload,omitempty" xml:"Payload,omitempty"`
	UserInfoShrink *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s DeviceControlShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeviceControlShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeviceControlShrinkRequest) GetPayloadShrink() *string {
	return s.PayloadShrink
}

func (s *DeviceControlShrinkRequest) GetUserInfoShrink() *string {
	return s.UserInfoShrink
}

func (s *DeviceControlShrinkRequest) SetPayloadShrink(v string) *DeviceControlShrinkRequest {
	s.PayloadShrink = &v
	return s
}

func (s *DeviceControlShrinkRequest) SetUserInfoShrink(v string) *DeviceControlShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

func (s *DeviceControlShrinkRequest) Validate() error {
	return dara.Validate(s)
}
