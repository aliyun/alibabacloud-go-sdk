// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeviceGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceGroupId(v string) *GetDeviceGroupRequest
	GetDeviceGroupId() *string
}

type GetDeviceGroupRequest struct {
	// The device label ID. You can obtain this value from:
	//
	// - [ListDeviceGroups](~~ListDeviceGroups~~): Lists device labels.
	//
	// - [CreateDeviceGroup](~~CreateDeviceGroup~~): Creates a device label.
	//
	// This parameter is required.
	//
	// example:
	//
	// device-group-5191cf830a5e****
	DeviceGroupId *string `json:"DeviceGroupId,omitempty" xml:"DeviceGroupId,omitempty"`
}

func (s GetDeviceGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDeviceGroupRequest) GoString() string {
	return s.String()
}

func (s *GetDeviceGroupRequest) GetDeviceGroupId() *string {
	return s.DeviceGroupId
}

func (s *GetDeviceGroupRequest) SetDeviceGroupId(v string) *GetDeviceGroupRequest {
	s.DeviceGroupId = &v
	return s
}

func (s *GetDeviceGroupRequest) Validate() error {
	return dara.Validate(s)
}
