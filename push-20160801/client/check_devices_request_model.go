// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckDevicesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppKey(v int64) *CheckDevicesRequest
	GetAppKey() *int64
	SetDeviceIds(v string) *CheckDevicesRequest
	GetDeviceIds() *string
}

type CheckDevicesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 23419851
	AppKey *int64 `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ae296f3b04a58a05b30c95f****,ae296f3b04a58a05b30c95f****,ae296f3b04a58a05b30c95f****
	DeviceIds *string `json:"DeviceIds,omitempty" xml:"DeviceIds,omitempty"`
}

func (s CheckDevicesRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckDevicesRequest) GoString() string {
	return s.String()
}

func (s *CheckDevicesRequest) GetAppKey() *int64 {
	return s.AppKey
}

func (s *CheckDevicesRequest) GetDeviceIds() *string {
	return s.DeviceIds
}

func (s *CheckDevicesRequest) SetAppKey(v int64) *CheckDevicesRequest {
	s.AppKey = &v
	return s
}

func (s *CheckDevicesRequest) SetDeviceIds(v string) *CheckDevicesRequest {
	s.DeviceIds = &v
	return s
}

func (s *CheckDevicesRequest) Validate() error {
	return dara.Validate(s)
}
