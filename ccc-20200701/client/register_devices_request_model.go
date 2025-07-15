// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRegisterDevicesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceId(v string) *RegisterDevicesRequest
	GetDeviceId() *string
	SetInstanceId(v string) *RegisterDevicesRequest
	GetInstanceId() *string
	SetPassword(v string) *RegisterDevicesRequest
	GetPassword() *string
	SetUserIdListJson(v string) *RegisterDevicesRequest
	GetUserIdListJson() *string
}

type RegisterDevicesRequest struct {
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
	// ["user-test@ccc-test"]
	UserIdListJson *string `json:"UserIdListJson,omitempty" xml:"UserIdListJson,omitempty"`
}

func (s RegisterDevicesRequest) String() string {
	return dara.Prettify(s)
}

func (s RegisterDevicesRequest) GoString() string {
	return s.String()
}

func (s *RegisterDevicesRequest) GetDeviceId() *string {
	return s.DeviceId
}

func (s *RegisterDevicesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RegisterDevicesRequest) GetPassword() *string {
	return s.Password
}

func (s *RegisterDevicesRequest) GetUserIdListJson() *string {
	return s.UserIdListJson
}

func (s *RegisterDevicesRequest) SetDeviceId(v string) *RegisterDevicesRequest {
	s.DeviceId = &v
	return s
}

func (s *RegisterDevicesRequest) SetInstanceId(v string) *RegisterDevicesRequest {
	s.InstanceId = &v
	return s
}

func (s *RegisterDevicesRequest) SetPassword(v string) *RegisterDevicesRequest {
	s.Password = &v
	return s
}

func (s *RegisterDevicesRequest) SetUserIdListJson(v string) *RegisterDevicesRequest {
	s.UserIdListJson = &v
	return s
}

func (s *RegisterDevicesRequest) Validate() error {
	return dara.Validate(s)
}
