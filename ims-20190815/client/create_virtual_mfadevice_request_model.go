// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVirtualMFADeviceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetVirtualMFADeviceName(v string) *CreateVirtualMFADeviceRequest
	GetVirtualMFADeviceName() *string
}

type CreateVirtualMFADeviceRequest struct {
	// This parameter is required.
	VirtualMFADeviceName *string `json:"VirtualMFADeviceName,omitempty" xml:"VirtualMFADeviceName,omitempty"`
}

func (s CreateVirtualMFADeviceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateVirtualMFADeviceRequest) GoString() string {
	return s.String()
}

func (s *CreateVirtualMFADeviceRequest) GetVirtualMFADeviceName() *string {
	return s.VirtualMFADeviceName
}

func (s *CreateVirtualMFADeviceRequest) SetVirtualMFADeviceName(v string) *CreateVirtualMFADeviceRequest {
	s.VirtualMFADeviceName = &v
	return s
}

func (s *CreateVirtualMFADeviceRequest) Validate() error {
	return dara.Validate(s)
}
