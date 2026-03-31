// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVirtualMFADeviceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateVirtualMFADeviceResponseBody
	GetRequestId() *string
	SetVirtualMFADevice(v *CreateVirtualMFADeviceResponseBodyVirtualMFADevice) *CreateVirtualMFADeviceResponseBody
	GetVirtualMFADevice() *CreateVirtualMFADeviceResponseBodyVirtualMFADevice
}

type CreateVirtualMFADeviceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the MFA device.
	VirtualMFADevice *CreateVirtualMFADeviceResponseBodyVirtualMFADevice `json:"VirtualMFADevice,omitempty" xml:"VirtualMFADevice,omitempty" type:"Struct"`
}

func (s CreateVirtualMFADeviceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateVirtualMFADeviceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVirtualMFADeviceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateVirtualMFADeviceResponseBody) GetVirtualMFADevice() *CreateVirtualMFADeviceResponseBodyVirtualMFADevice {
	return s.VirtualMFADevice
}

func (s *CreateVirtualMFADeviceResponseBody) SetRequestId(v string) *CreateVirtualMFADeviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateVirtualMFADeviceResponseBody) SetVirtualMFADevice(v *CreateVirtualMFADeviceResponseBodyVirtualMFADevice) *CreateVirtualMFADeviceResponseBody {
	s.VirtualMFADevice = v
	return s
}

func (s *CreateVirtualMFADeviceResponseBody) Validate() error {
	if s.VirtualMFADevice != nil {
		if err := s.VirtualMFADevice.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateVirtualMFADeviceResponseBodyVirtualMFADevice struct {
	// The key of the MFA device.
	//
	// example:
	//
	// DSF98HAD982KJA9SDFNAS9D8FU839B8ADHBGS****
	Base32StringSeed *string `json:"Base32StringSeed,omitempty" xml:"Base32StringSeed,omitempty"`
	// The Base64-encoded QR code, in the PNG format.
	//
	// example:
	//
	// YXNkZmFzZDlmeW5hc2Q5OGZoODd4bXJmcThhaGU5aSBmYXNkZiBzYWRmIGFGIDRxd2VjIGEgdHEz****
	QRCodePNG *string `json:"QRCodePNG,omitempty" xml:"QRCodePNG,omitempty"`
	// The serial number of the MFA device.
	//
	// example:
	//
	// acs:ram::123456789012****:mfa/device001
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
}

func (s CreateVirtualMFADeviceResponseBodyVirtualMFADevice) String() string {
	return dara.Prettify(s)
}

func (s CreateVirtualMFADeviceResponseBodyVirtualMFADevice) GoString() string {
	return s.String()
}

func (s *CreateVirtualMFADeviceResponseBodyVirtualMFADevice) GetBase32StringSeed() *string {
	return s.Base32StringSeed
}

func (s *CreateVirtualMFADeviceResponseBodyVirtualMFADevice) GetQRCodePNG() *string {
	return s.QRCodePNG
}

func (s *CreateVirtualMFADeviceResponseBodyVirtualMFADevice) GetSerialNumber() *string {
	return s.SerialNumber
}

func (s *CreateVirtualMFADeviceResponseBodyVirtualMFADevice) SetBase32StringSeed(v string) *CreateVirtualMFADeviceResponseBodyVirtualMFADevice {
	s.Base32StringSeed = &v
	return s
}

func (s *CreateVirtualMFADeviceResponseBodyVirtualMFADevice) SetQRCodePNG(v string) *CreateVirtualMFADeviceResponseBodyVirtualMFADevice {
	s.QRCodePNG = &v
	return s
}

func (s *CreateVirtualMFADeviceResponseBodyVirtualMFADevice) SetSerialNumber(v string) *CreateVirtualMFADeviceResponseBodyVirtualMFADevice {
	s.SerialNumber = &v
	return s
}

func (s *CreateVirtualMFADeviceResponseBodyVirtualMFADevice) Validate() error {
	return dara.Validate(s)
}
