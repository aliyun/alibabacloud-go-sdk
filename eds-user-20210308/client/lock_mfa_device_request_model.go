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
	SetBusinessChannel(v string) *LockMfaDeviceRequest
	GetBusinessChannel() *string
	SetSerialNumber(v string) *LockMfaDeviceRequest
	GetSerialNumber() *string
}

type LockMfaDeviceRequest struct {
	// The address of the AD office network.
	//
	// example:
	//
	// upower.vip
	AdDomain *string `json:"AdDomain,omitempty" xml:"AdDomain,omitempty"`
	// The channel.
	//
	// example:
	//
	// ENTERPRISE
	BusinessChannel *string `json:"BusinessChannel,omitempty" xml:"BusinessChannel,omitempty"`
	// The serial number of the bound virtual MFA device. The user must first complete the MFA binding in the client authentication flow. You can call [DescribeMfaDevices](~~DescribeMfaDevices~~) to query the serial number of the bound device.
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

func (s *LockMfaDeviceRequest) GetBusinessChannel() *string {
	return s.BusinessChannel
}

func (s *LockMfaDeviceRequest) GetSerialNumber() *string {
	return s.SerialNumber
}

func (s *LockMfaDeviceRequest) SetAdDomain(v string) *LockMfaDeviceRequest {
	s.AdDomain = &v
	return s
}

func (s *LockMfaDeviceRequest) SetBusinessChannel(v string) *LockMfaDeviceRequest {
	s.BusinessChannel = &v
	return s
}

func (s *LockMfaDeviceRequest) SetSerialNumber(v string) *LockMfaDeviceRequest {
	s.SerialNumber = &v
	return s
}

func (s *LockMfaDeviceRequest) Validate() error {
	return dara.Validate(s)
}
