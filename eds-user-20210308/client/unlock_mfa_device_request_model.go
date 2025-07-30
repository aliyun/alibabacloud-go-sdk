// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnlockMfaDeviceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdDomain(v string) *UnlockMfaDeviceRequest
	GetAdDomain() *string
	SetSerialNumber(v string) *UnlockMfaDeviceRequest
	GetSerialNumber() *string
}

type UnlockMfaDeviceRequest struct {
	// The domain of the Active Directory (AD) workspace.
	//
	// example:
	//
	// welab.co.id
	AdDomain *string `json:"AdDomain,omitempty" xml:"AdDomain,omitempty"`
	// The serial number of the virtual MFA device. The serial number is unique for each device.
	//
	// This parameter is required.
	//
	// example:
	//
	// dc856334-446b-4035-bfbc-18af261e****
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
}

func (s UnlockMfaDeviceRequest) String() string {
	return dara.Prettify(s)
}

func (s UnlockMfaDeviceRequest) GoString() string {
	return s.String()
}

func (s *UnlockMfaDeviceRequest) GetAdDomain() *string {
	return s.AdDomain
}

func (s *UnlockMfaDeviceRequest) GetSerialNumber() *string {
	return s.SerialNumber
}

func (s *UnlockMfaDeviceRequest) SetAdDomain(v string) *UnlockMfaDeviceRequest {
	s.AdDomain = &v
	return s
}

func (s *UnlockMfaDeviceRequest) SetSerialNumber(v string) *UnlockMfaDeviceRequest {
	s.SerialNumber = &v
	return s
}

func (s *UnlockMfaDeviceRequest) Validate() error {
	return dara.Validate(s)
}
