// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddFaceDeviceToDeviceGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceGroupId(v string) *AddFaceDeviceToDeviceGroupRequest
	GetDeviceGroupId() *string
	SetDeviceName(v string) *AddFaceDeviceToDeviceGroupRequest
	GetDeviceName() *string
	SetIotInstanceId(v string) *AddFaceDeviceToDeviceGroupRequest
	GetIotInstanceId() *string
	SetIsolationId(v string) *AddFaceDeviceToDeviceGroupRequest
	GetIsolationId() *string
	SetProductKey(v string) *AddFaceDeviceToDeviceGroupRequest
	GetProductKey() *string
}

type AddFaceDeviceToDeviceGroupRequest struct {
	// The device group ID. You can obtain the ID by calling the AddFaceDeviceGroup operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// kvjyhurc
	DeviceGroupId *string `json:"DeviceGroupId,omitempty" xml:"DeviceGroupId,omitempty"`
	// The device name.
	//
	// This parameter is required.
	//
	// example:
	//
	// TestDevice
	DeviceName *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	// The IoT instance ID.
	//
	// example:
	//
	// TestInstanceId
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	// The isolation ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// TestIsolationId
	IsolationId *string `json:"IsolationId,omitempty" xml:"IsolationId,omitempty"`
	// The product key.
	//
	// This parameter is required.
	//
	// example:
	//
	// a1Tp7KYZHHK
	ProductKey *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
}

func (s AddFaceDeviceToDeviceGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s AddFaceDeviceToDeviceGroupRequest) GoString() string {
	return s.String()
}

func (s *AddFaceDeviceToDeviceGroupRequest) GetDeviceGroupId() *string {
	return s.DeviceGroupId
}

func (s *AddFaceDeviceToDeviceGroupRequest) GetDeviceName() *string {
	return s.DeviceName
}

func (s *AddFaceDeviceToDeviceGroupRequest) GetIotInstanceId() *string {
	return s.IotInstanceId
}

func (s *AddFaceDeviceToDeviceGroupRequest) GetIsolationId() *string {
	return s.IsolationId
}

func (s *AddFaceDeviceToDeviceGroupRequest) GetProductKey() *string {
	return s.ProductKey
}

func (s *AddFaceDeviceToDeviceGroupRequest) SetDeviceGroupId(v string) *AddFaceDeviceToDeviceGroupRequest {
	s.DeviceGroupId = &v
	return s
}

func (s *AddFaceDeviceToDeviceGroupRequest) SetDeviceName(v string) *AddFaceDeviceToDeviceGroupRequest {
	s.DeviceName = &v
	return s
}

func (s *AddFaceDeviceToDeviceGroupRequest) SetIotInstanceId(v string) *AddFaceDeviceToDeviceGroupRequest {
	s.IotInstanceId = &v
	return s
}

func (s *AddFaceDeviceToDeviceGroupRequest) SetIsolationId(v string) *AddFaceDeviceToDeviceGroupRequest {
	s.IsolationId = &v
	return s
}

func (s *AddFaceDeviceToDeviceGroupRequest) SetProductKey(v string) *AddFaceDeviceToDeviceGroupRequest {
	s.ProductKey = &v
	return s
}

func (s *AddFaceDeviceToDeviceGroupRequest) Validate() error {
	return dara.Validate(s)
}
