// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveMfaDeviceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdDomain(v string) *RemoveMfaDeviceRequest
	GetAdDomain() *string
	SetSerialNumber(v string) *RemoveMfaDeviceRequest
	GetSerialNumber() *string
}

type RemoveMfaDeviceRequest struct {
	// The address of the AD office network.
	//
	// example:
	//
	// alpha.lftltd.net
	AdDomain *string `json:"AdDomain,omitempty" xml:"AdDomain,omitempty"`
	// The serial number of the virtual MFA device. You can call the [DescribeMfaDevices](~~DescribeMfaDevices~~) operation to get this information.
	//
	// This parameter is required.
	//
	// example:
	//
	// dc856334-446b-4035-bfbc-18af261e****
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
}

func (s RemoveMfaDeviceRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveMfaDeviceRequest) GoString() string {
	return s.String()
}

func (s *RemoveMfaDeviceRequest) GetAdDomain() *string {
	return s.AdDomain
}

func (s *RemoveMfaDeviceRequest) GetSerialNumber() *string {
	return s.SerialNumber
}

func (s *RemoveMfaDeviceRequest) SetAdDomain(v string) *RemoveMfaDeviceRequest {
	s.AdDomain = &v
	return s
}

func (s *RemoveMfaDeviceRequest) SetSerialNumber(v string) *RemoveMfaDeviceRequest {
	s.SerialNumber = &v
	return s
}

func (s *RemoveMfaDeviceRequest) Validate() error {
	return dara.Validate(s)
}
