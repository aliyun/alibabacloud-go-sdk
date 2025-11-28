// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableDeviceGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceGroupId(v string) *DisableDeviceGroupRequest
	GetDeviceGroupId() *string
	SetRegionId(v string) *DisableDeviceGroupRequest
	GetRegionId() *string
}

type DisableDeviceGroupRequest struct {
	// This parameter is required.
	DeviceGroupId *string `json:"DeviceGroupId,omitempty" xml:"DeviceGroupId,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DisableDeviceGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableDeviceGroupRequest) GoString() string {
	return s.String()
}

func (s *DisableDeviceGroupRequest) GetDeviceGroupId() *string {
	return s.DeviceGroupId
}

func (s *DisableDeviceGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DisableDeviceGroupRequest) SetDeviceGroupId(v string) *DisableDeviceGroupRequest {
	s.DeviceGroupId = &v
	return s
}

func (s *DisableDeviceGroupRequest) SetRegionId(v string) *DisableDeviceGroupRequest {
	s.RegionId = &v
	return s
}

func (s *DisableDeviceGroupRequest) Validate() error {
	return dara.Validate(s)
}
