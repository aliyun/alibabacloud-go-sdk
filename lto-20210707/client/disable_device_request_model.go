// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableDeviceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceId(v string) *DisableDeviceRequest
	GetDeviceId() *string
	SetRegionId(v string) *DisableDeviceRequest
	GetRegionId() *string
}

type DisableDeviceRequest struct {
	// This parameter is required.
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DisableDeviceRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableDeviceRequest) GoString() string {
	return s.String()
}

func (s *DisableDeviceRequest) GetDeviceId() *string {
	return s.DeviceId
}

func (s *DisableDeviceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DisableDeviceRequest) SetDeviceId(v string) *DisableDeviceRequest {
	s.DeviceId = &v
	return s
}

func (s *DisableDeviceRequest) SetRegionId(v string) *DisableDeviceRequest {
	s.RegionId = &v
	return s
}

func (s *DisableDeviceRequest) Validate() error {
	return dara.Validate(s)
}
