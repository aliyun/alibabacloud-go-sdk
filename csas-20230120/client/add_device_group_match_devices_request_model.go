// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDeviceGroupMatchDevicesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDevTags(v []*string) *AddDeviceGroupMatchDevicesRequest
	GetDevTags() []*string
	SetDeviceGroupId(v string) *AddDeviceGroupMatchDevicesRequest
	GetDeviceGroupId() *string
}

type AddDeviceGroupMatchDevicesRequest struct {
	// The IDs of the terminal devices to add. You must specify at least one ID, and the IDs must be unique.
	//
	// This parameter is required.
	DevTags []*string `json:"DevTags,omitempty" xml:"DevTags,omitempty" type:"Repeated"`
	// The ID of the device tag.
	//
	// This parameter is required.
	//
	// example:
	//
	// device-group-5191cf830a5e****
	DeviceGroupId *string `json:"DeviceGroupId,omitempty" xml:"DeviceGroupId,omitempty"`
}

func (s AddDeviceGroupMatchDevicesRequest) String() string {
	return dara.Prettify(s)
}

func (s AddDeviceGroupMatchDevicesRequest) GoString() string {
	return s.String()
}

func (s *AddDeviceGroupMatchDevicesRequest) GetDevTags() []*string {
	return s.DevTags
}

func (s *AddDeviceGroupMatchDevicesRequest) GetDeviceGroupId() *string {
	return s.DeviceGroupId
}

func (s *AddDeviceGroupMatchDevicesRequest) SetDevTags(v []*string) *AddDeviceGroupMatchDevicesRequest {
	s.DevTags = v
	return s
}

func (s *AddDeviceGroupMatchDevicesRequest) SetDeviceGroupId(v string) *AddDeviceGroupMatchDevicesRequest {
	s.DeviceGroupId = &v
	return s
}

func (s *AddDeviceGroupMatchDevicesRequest) Validate() error {
	return dara.Validate(s)
}
