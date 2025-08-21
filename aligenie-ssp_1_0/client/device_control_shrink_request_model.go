// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeviceControlShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetControlRequestShrink(v string) *DeviceControlShrinkRequest
	GetControlRequestShrink() *string
	SetDeviceInfoShrink(v string) *DeviceControlShrinkRequest
	GetDeviceInfoShrink() *string
}

type DeviceControlShrinkRequest struct {
	ControlRequestShrink *string `json:"ControlRequest,omitempty" xml:"ControlRequest,omitempty"`
	// This parameter is required.
	DeviceInfoShrink *string `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty"`
}

func (s DeviceControlShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeviceControlShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeviceControlShrinkRequest) GetControlRequestShrink() *string {
	return s.ControlRequestShrink
}

func (s *DeviceControlShrinkRequest) GetDeviceInfoShrink() *string {
	return s.DeviceInfoShrink
}

func (s *DeviceControlShrinkRequest) SetControlRequestShrink(v string) *DeviceControlShrinkRequest {
	s.ControlRequestShrink = &v
	return s
}

func (s *DeviceControlShrinkRequest) SetDeviceInfoShrink(v string) *DeviceControlShrinkRequest {
	s.DeviceInfoShrink = &v
	return s
}

func (s *DeviceControlShrinkRequest) Validate() error {
	return dara.Validate(s)
}
