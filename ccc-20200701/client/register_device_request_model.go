// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRegisterDeviceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceId(v string) *RegisterDeviceRequest
	GetDeviceId() *string
	SetInstanceId(v string) *RegisterDeviceRequest
	GetInstanceId() *string
	SetPassword(v string) *RegisterDeviceRequest
	GetPassword() *string
	SetUserId(v string) *RegisterDeviceRequest
	GetUserId() *string
}

type RegisterDeviceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// test-device
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// password
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// example:
	//
	// agent@ccc-test
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s RegisterDeviceRequest) String() string {
	return dara.Prettify(s)
}

func (s RegisterDeviceRequest) GoString() string {
	return s.String()
}

func (s *RegisterDeviceRequest) GetDeviceId() *string {
	return s.DeviceId
}

func (s *RegisterDeviceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RegisterDeviceRequest) GetPassword() *string {
	return s.Password
}

func (s *RegisterDeviceRequest) GetUserId() *string {
	return s.UserId
}

func (s *RegisterDeviceRequest) SetDeviceId(v string) *RegisterDeviceRequest {
	s.DeviceId = &v
	return s
}

func (s *RegisterDeviceRequest) SetInstanceId(v string) *RegisterDeviceRequest {
	s.InstanceId = &v
	return s
}

func (s *RegisterDeviceRequest) SetPassword(v string) *RegisterDeviceRequest {
	s.Password = &v
	return s
}

func (s *RegisterDeviceRequest) SetUserId(v string) *RegisterDeviceRequest {
	s.UserId = &v
	return s
}

func (s *RegisterDeviceRequest) Validate() error {
	return dara.Validate(s)
}
