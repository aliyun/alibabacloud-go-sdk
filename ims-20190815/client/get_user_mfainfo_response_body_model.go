// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserMFAInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIsMFAEnable(v bool) *GetUserMFAInfoResponseBody
	GetIsMFAEnable() *bool
	SetMFADevice(v *GetUserMFAInfoResponseBodyMFADevice) *GetUserMFAInfoResponseBody
	GetMFADevice() *GetUserMFAInfoResponseBodyMFADevice
	SetRequestId(v string) *GetUserMFAInfoResponseBody
	GetRequestId() *string
}

type GetUserMFAInfoResponseBody struct {
	IsMFAEnable *bool                                `json:"IsMFAEnable,omitempty" xml:"IsMFAEnable,omitempty"`
	MFADevice   *GetUserMFAInfoResponseBodyMFADevice `json:"MFADevice,omitempty" xml:"MFADevice,omitempty" type:"Struct"`
	RequestId   *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetUserMFAInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetUserMFAInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserMFAInfoResponseBody) GetIsMFAEnable() *bool {
	return s.IsMFAEnable
}

func (s *GetUserMFAInfoResponseBody) GetMFADevice() *GetUserMFAInfoResponseBodyMFADevice {
	return s.MFADevice
}

func (s *GetUserMFAInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetUserMFAInfoResponseBody) SetIsMFAEnable(v bool) *GetUserMFAInfoResponseBody {
	s.IsMFAEnable = &v
	return s
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
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	Type         *string `json:"Type,omitempty" xml:"Type,omitempty"`
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
