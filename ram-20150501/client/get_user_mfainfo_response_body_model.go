// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserMFAInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMFADevice(v *GetUserMFAInfoResponseBodyMFADevice) *GetUserMFAInfoResponseBody
	GetMFADevice() *GetUserMFAInfoResponseBodyMFADevice
	SetRequestId(v string) *GetUserMFAInfoResponseBody
	GetRequestId() *string
}

type GetUserMFAInfoResponseBody struct {
	// The information about the MFA device that is bound to the RAM user.
	MFADevice *GetUserMFAInfoResponseBodyMFADevice `json:"MFADevice,omitempty" xml:"MFADevice,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetUserMFAInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetUserMFAInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserMFAInfoResponseBody) GetMFADevice() *GetUserMFAInfoResponseBodyMFADevice {
	return s.MFADevice
}

func (s *GetUserMFAInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetUserMFAInfoResponseBody) SetMFADevice(v *GetUserMFAInfoResponseBodyMFADevice) *GetUserMFAInfoResponseBody {
	s.MFADevice = v
	return s
}

func (s *GetUserMFAInfoResponseBody) SetRequestId(v string) *GetUserMFAInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserMFAInfoResponseBody) Validate() error {
	if s.MFADevice != nil {
		if err := s.MFADevice.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetUserMFAInfoResponseBodyMFADevice struct {
	// The serial number of the MFA device.
	//
	// example:
	//
	// acs:ram::177242285274****:mfa/test
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	// The type of the MFA device. Valid values:
	//
	// 	- VMFA: virtual MFA device.
	//
	// 	- U2F: Universal 2nd Factor (U2F) security key.
	//
	// example:
	//
	// VMFA
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetUserMFAInfoResponseBodyMFADevice) String() string {
	return dara.Prettify(s)
}

func (s GetUserMFAInfoResponseBodyMFADevice) GoString() string {
	return s.String()
}

func (s *GetUserMFAInfoResponseBodyMFADevice) GetSerialNumber() *string {
	return s.SerialNumber
}

func (s *GetUserMFAInfoResponseBodyMFADevice) GetType() *string {
	return s.Type
}

func (s *GetUserMFAInfoResponseBodyMFADevice) SetSerialNumber(v string) *GetUserMFAInfoResponseBodyMFADevice {
	s.SerialNumber = &v
	return s
}

func (s *GetUserMFAInfoResponseBodyMFADevice) SetType(v string) *GetUserMFAInfoResponseBodyMFADevice {
	s.Type = &v
	return s
}

func (s *GetUserMFAInfoResponseBodyMFADevice) Validate() error {
	return dara.Validate(s)
}
