// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveDeviceGroupMatchDevicesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDevTags(v []*string) *RemoveDeviceGroupMatchDevicesRequest
	GetDevTags() []*string
	SetDeviceGroupId(v string) *RemoveDeviceGroupMatchDevicesRequest
	GetDeviceGroupId() *string
}

type RemoveDeviceGroupMatchDevicesRequest struct {
	// The collection of terminal device IDs to be removed. At least one ID must be specified, and duplicate values are not allowed.
	//
	// This parameter is required.
	DevTags []*string `json:"DevTags,omitempty" xml:"DevTags,omitempty" type:"Repeated"`
	// The device label ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// device-group-5191cf830a5e****
	DeviceGroupId *string `json:"DeviceGroupId,omitempty" xml:"DeviceGroupId,omitempty"`
}

func (s RemoveDeviceGroupMatchDevicesRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveDeviceGroupMatchDevicesRequest) GoString() string {
	return s.String()
}

func (s *RemoveDeviceGroupMatchDevicesRequest) GetDevTags() []*string {
	return s.DevTags
}

func (s *RemoveDeviceGroupMatchDevicesRequest) GetDeviceGroupId() *string {
	return s.DeviceGroupId
}

func (s *RemoveDeviceGroupMatchDevicesRequest) SetDevTags(v []*string) *RemoveDeviceGroupMatchDevicesRequest {
	s.DevTags = v
	return s
}

func (s *RemoveDeviceGroupMatchDevicesRequest) SetDeviceGroupId(v string) *RemoveDeviceGroupMatchDevicesRequest {
	s.DeviceGroupId = &v
	return s
}

func (s *RemoveDeviceGroupMatchDevicesRequest) Validate() error {
	return dara.Validate(s)
}
