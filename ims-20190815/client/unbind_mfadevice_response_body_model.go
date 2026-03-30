// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindMFADeviceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMFADevice(v *UnbindMFADeviceResponseBodyMFADevice) *UnbindMFADeviceResponseBody
	GetMFADevice() *UnbindMFADeviceResponseBodyMFADevice
	SetRequestId(v string) *UnbindMFADeviceResponseBody
	GetRequestId() *string
}

type UnbindMFADeviceResponseBody struct {
	// The information about the MFA device.
	MFADevice *UnbindMFADeviceResponseBodyMFADevice `json:"MFADevice,omitempty" xml:"MFADevice,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// A26CB3E9-1021-452A-AC57-3134B3BA0E4C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnbindMFADeviceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UnbindMFADeviceResponseBody) GoString() string {
	return s.String()
}

func (s *UnbindMFADeviceResponseBody) GetMFADevice() *UnbindMFADeviceResponseBodyMFADevice {
	return s.MFADevice
}

func (s *UnbindMFADeviceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UnbindMFADeviceResponseBody) SetMFADevice(v *UnbindMFADeviceResponseBodyMFADevice) *UnbindMFADeviceResponseBody {
	s.MFADevice = v
	return s
}

func (s *UnbindMFADeviceResponseBody) SetRequestId(v string) *UnbindMFADeviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnbindMFADeviceResponseBody) Validate() error {
	if s.MFADevice != nil {
		if err := s.MFADevice.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UnbindMFADeviceResponseBodyMFADevice struct {
	// The serial number of the MFA device.
	//
	// example:
	//
	// acs:ram::151298381312****:mfa/device001
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
}

func (s UnbindMFADeviceResponseBodyMFADevice) String() string {
	return dara.Prettify(s)
}

func (s UnbindMFADeviceResponseBodyMFADevice) GoString() string {
	return s.String()
}

func (s *UnbindMFADeviceResponseBodyMFADevice) GetSerialNumber() *string {
	return s.SerialNumber
}

func (s *UnbindMFADeviceResponseBodyMFADevice) SetSerialNumber(v string) *UnbindMFADeviceResponseBodyMFADevice {
	s.SerialNumber = &v
	return s
}

func (s *UnbindMFADeviceResponseBodyMFADevice) Validate() error {
	return dara.Validate(s)
}
