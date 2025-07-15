// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLockVirtualMFADeviceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *LockVirtualMFADeviceRequest
	GetRegionId() *string
	SetSerialNumber(v string) *LockVirtualMFADeviceRequest
	GetSerialNumber() *string
}

type LockVirtualMFADeviceRequest struct {
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The serial number of the virtual MFA device, which is a unique identifier.
	//
	// You can call the [DescribeVirtualMFADevices](https://help.aliyun.com/document_detail/206210.html) operation to query the serial number of the virtual MFA device bound to AD users.
	//
	// This parameter is required.
	//
	// example:
	//
	// a25f297f-f2e1-4a44-bbf1-5f48a6e5****
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
}

func (s LockVirtualMFADeviceRequest) String() string {
	return dara.Prettify(s)
}

func (s LockVirtualMFADeviceRequest) GoString() string {
	return s.String()
}

func (s *LockVirtualMFADeviceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *LockVirtualMFADeviceRequest) GetSerialNumber() *string {
	return s.SerialNumber
}

func (s *LockVirtualMFADeviceRequest) SetRegionId(v string) *LockVirtualMFADeviceRequest {
	s.RegionId = &v
	return s
}

func (s *LockVirtualMFADeviceRequest) SetSerialNumber(v string) *LockVirtualMFADeviceRequest {
	s.SerialNumber = &v
	return s
}

func (s *LockVirtualMFADeviceRequest) Validate() error {
	return dara.Validate(s)
}
