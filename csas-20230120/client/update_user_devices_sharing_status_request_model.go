// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUserDevicesSharingStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceTags(v []*string) *UpdateUserDevicesSharingStatusRequest
	GetDeviceTags() []*string
	SetSharingStatus(v bool) *UpdateUserDevicesSharingStatusRequest
	GetSharingStatus() *bool
}

type UpdateUserDevicesSharingStatusRequest struct {
	// This parameter is required.
	DeviceTags []*string `json:"DeviceTags,omitempty" xml:"DeviceTags,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// true
	SharingStatus *bool `json:"SharingStatus,omitempty" xml:"SharingStatus,omitempty"`
}

func (s UpdateUserDevicesSharingStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserDevicesSharingStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateUserDevicesSharingStatusRequest) GetDeviceTags() []*string {
	return s.DeviceTags
}

func (s *UpdateUserDevicesSharingStatusRequest) GetSharingStatus() *bool {
	return s.SharingStatus
}

func (s *UpdateUserDevicesSharingStatusRequest) SetDeviceTags(v []*string) *UpdateUserDevicesSharingStatusRequest {
	s.DeviceTags = v
	return s
}

func (s *UpdateUserDevicesSharingStatusRequest) SetSharingStatus(v bool) *UpdateUserDevicesSharingStatusRequest {
	s.SharingStatus = &v
	return s
}

func (s *UpdateUserDevicesSharingStatusRequest) Validate() error {
	return dara.Validate(s)
}
