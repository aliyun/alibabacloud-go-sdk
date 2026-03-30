// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVirtualMFADeviceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSerialNumber(v string) *DeleteVirtualMFADeviceRequest
	GetSerialNumber() *string
}

type DeleteVirtualMFADeviceRequest struct {
	// The serial number of the MFA device.
	//
	// This parameter is required.
	//
	// example:
	//
	// acs:ram::123456789012****:mfa/device002
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
}

func (s DeleteVirtualMFADeviceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteVirtualMFADeviceRequest) GoString() string {
	return s.String()
}

func (s *DeleteVirtualMFADeviceRequest) GetSerialNumber() *string {
	return s.SerialNumber
}

func (s *DeleteVirtualMFADeviceRequest) SetSerialNumber(v string) *DeleteVirtualMFADeviceRequest {
	s.SerialNumber = &v
	return s
}

func (s *DeleteVirtualMFADeviceRequest) Validate() error {
	return dara.Validate(s)
}
