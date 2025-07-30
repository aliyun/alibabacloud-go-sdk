// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLockMfaDeviceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdDomain(v string) *LockMfaDeviceRequest
	GetAdDomain() *string
	SetSerialNumber(v string) *LockMfaDeviceRequest
	GetSerialNumber() *string
}

type LockMfaDeviceRequest struct {
	// The domain of the Active Directory (AD) workspace.
	//
	// example:
	//
	// pg-jifenn.com
	AdDomain *string `json:"AdDomain,omitempty" xml:"AdDomain,omitempty"`
	// The serial number of the virtual MFA device. The serial number is unique for each device.
	//
	// example:
	//
	// dc856334-446b-4035-bfbc-18af261e****
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
}

func (s LockMfaDeviceRequest) String() string {
	return dara.Prettify(s)
}

func (s LockMfaDeviceRequest) GoString() string {
	return s.String()
}

func (s *LockMfaDeviceRequest) GetAdDomain() *string {
	return s.AdDomain
}

func (s *LockMfaDeviceRequest) GetSerialNumber() *string {
	return s.SerialNumber
}

func (s *LockMfaDeviceRequest) SetAdDomain(v string) *LockMfaDeviceRequest {
	s.AdDomain = &v
	return s
}

func (s *LockMfaDeviceRequest) SetSerialNumber(v string) *LockMfaDeviceRequest {
	s.SerialNumber = &v
	return s
}

func (s *LockMfaDeviceRequest) Validate() error {
	return dara.Validate(s)
}
