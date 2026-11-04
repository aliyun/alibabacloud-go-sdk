// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveFaceDeviceFromDeviceGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceGroupId(v string) *RemoveFaceDeviceFromDeviceGroupRequest
	GetDeviceGroupId() *string
	SetDeviceName(v string) *RemoveFaceDeviceFromDeviceGroupRequest
	GetDeviceName() *string
	SetIotInstanceId(v string) *RemoveFaceDeviceFromDeviceGroupRequest
	GetIotInstanceId() *string
	SetIsolationId(v string) *RemoveFaceDeviceFromDeviceGroupRequest
	GetIsolationId() *string
	SetProductKey(v string) *RemoveFaceDeviceFromDeviceGroupRequest
	GetProductKey() *string
}

type RemoveFaceDeviceFromDeviceGroupRequest struct {
	// The device group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// TestDeviceGroupId
	DeviceGroupId *string `json:"DeviceGroupId,omitempty" xml:"DeviceGroupId,omitempty"`
	// The device name.
	//
	// This parameter is required.
	//
	// example:
	//
	// TestDeviceName
	DeviceName *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// TestIotInstanceId
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
	// TestProductKey
	ProductKey *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
}

func (s RemoveFaceDeviceFromDeviceGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveFaceDeviceFromDeviceGroupRequest) GoString() string {
	return s.String()
}

func (s *RemoveFaceDeviceFromDeviceGroupRequest) GetDeviceGroupId() *string {
	return s.DeviceGroupId
}

func (s *RemoveFaceDeviceFromDeviceGroupRequest) GetDeviceName() *string {
	return s.DeviceName
}

func (s *RemoveFaceDeviceFromDeviceGroupRequest) GetIotInstanceId() *string {
	return s.IotInstanceId
}

func (s *RemoveFaceDeviceFromDeviceGroupRequest) GetIsolationId() *string {
	return s.IsolationId
}

func (s *RemoveFaceDeviceFromDeviceGroupRequest) GetProductKey() *string {
	return s.ProductKey
}

func (s *RemoveFaceDeviceFromDeviceGroupRequest) SetDeviceGroupId(v string) *RemoveFaceDeviceFromDeviceGroupRequest {
	s.DeviceGroupId = &v
	return s
}

func (s *RemoveFaceDeviceFromDeviceGroupRequest) SetDeviceName(v string) *RemoveFaceDeviceFromDeviceGroupRequest {
	s.DeviceName = &v
	return s
}

func (s *RemoveFaceDeviceFromDeviceGroupRequest) SetIotInstanceId(v string) *RemoveFaceDeviceFromDeviceGroupRequest {
	s.IotInstanceId = &v
	return s
}

func (s *RemoveFaceDeviceFromDeviceGroupRequest) SetIsolationId(v string) *RemoveFaceDeviceFromDeviceGroupRequest {
	s.IsolationId = &v
	return s
}

func (s *RemoveFaceDeviceFromDeviceGroupRequest) SetProductKey(v string) *RemoveFaceDeviceFromDeviceGroupRequest {
	s.ProductKey = &v
	return s
}

func (s *RemoveFaceDeviceFromDeviceGroupRequest) Validate() error {
	return dara.Validate(s)
}
