// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVirtualMFADeviceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *DeleteVirtualMFADeviceRequest
	GetRegionId() *string
	SetSerialNumber(v string) *DeleteVirtualMFADeviceRequest
	GetSerialNumber() *string
}

type DeleteVirtualMFADeviceRequest struct {
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
	// You can call the [DescribeVirtualMFADevices](https://help.aliyun.com/document_detail/206210.html) operation to query the serial number of the virtual MFA device that is bound by AD users.
	//
	// This parameter is required.
	//
	// example:
	//
	// a25f297f-f2e1-4a44-bbf1-5f48a6e5****
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
}

func (s DeleteVirtualMFADeviceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteVirtualMFADeviceRequest) GoString() string {
	return s.String()
}

func (s *DeleteVirtualMFADeviceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteVirtualMFADeviceRequest) GetSerialNumber() *string {
	return s.SerialNumber
}

func (s *DeleteVirtualMFADeviceRequest) SetRegionId(v string) *DeleteVirtualMFADeviceRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteVirtualMFADeviceRequest) SetSerialNumber(v string) *DeleteVirtualMFADeviceRequest {
	s.SerialNumber = &v
	return s
}

func (s *DeleteVirtualMFADeviceRequest) Validate() error {
	return dara.Validate(s)
}
