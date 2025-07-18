// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUserDevicesStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceAction(v string) *UpdateUserDevicesStatusRequest
	GetDeviceAction() *string
	SetDeviceTags(v []*string) *UpdateUserDevicesStatusRequest
	GetDeviceTags() []*string
}

type UpdateUserDevicesStatusRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// Unbound
	DeviceAction *string `json:"DeviceAction,omitempty" xml:"DeviceAction,omitempty"`
	// This parameter is required.
	DeviceTags []*string `json:"DeviceTags,omitempty" xml:"DeviceTags,omitempty" type:"Repeated"`
}

func (s UpdateUserDevicesStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserDevicesStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateUserDevicesStatusRequest) GetDeviceAction() *string {
	return s.DeviceAction
}

func (s *UpdateUserDevicesStatusRequest) GetDeviceTags() []*string {
	return s.DeviceTags
}

func (s *UpdateUserDevicesStatusRequest) SetDeviceAction(v string) *UpdateUserDevicesStatusRequest {
	s.DeviceAction = &v
	return s
}

func (s *UpdateUserDevicesStatusRequest) SetDeviceTags(v []*string) *UpdateUserDevicesStatusRequest {
	s.DeviceTags = v
	return s
}

func (s *UpdateUserDevicesStatusRequest) Validate() error {
	return dara.Validate(s)
}
