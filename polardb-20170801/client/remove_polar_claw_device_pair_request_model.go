// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemovePolarClawDevicePairRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *RemovePolarClawDevicePairRequest
	GetApplicationId() *string
	SetDeviceId(v string) *RemovePolarClawDevicePairRequest
	GetDeviceId() *string
}

type RemovePolarClawDevicePairRequest struct {
	// The application ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pa-**************
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The device ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// device-mac-789
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
}

func (s RemovePolarClawDevicePairRequest) String() string {
	return dara.Prettify(s)
}

func (s RemovePolarClawDevicePairRequest) GoString() string {
	return s.String()
}

func (s *RemovePolarClawDevicePairRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *RemovePolarClawDevicePairRequest) GetDeviceId() *string {
	return s.DeviceId
}

func (s *RemovePolarClawDevicePairRequest) SetApplicationId(v string) *RemovePolarClawDevicePairRequest {
	s.ApplicationId = &v
	return s
}

func (s *RemovePolarClawDevicePairRequest) SetDeviceId(v string) *RemovePolarClawDevicePairRequest {
	s.DeviceId = &v
	return s
}

func (s *RemovePolarClawDevicePairRequest) Validate() error {
	return dara.Validate(s)
}
