// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSignOutGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceId(v string) *SignOutGroupRequest
	GetDeviceId() *string
	SetInstanceId(v string) *SignOutGroupRequest
	GetInstanceId() *string
	SetUserId(v string) *SignOutGroupRequest
	GetUserId() *string
}

type SignOutGroupRequest struct {
	// example:
	//
	// device
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// agent@ccc-test
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s SignOutGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s SignOutGroupRequest) GoString() string {
	return s.String()
}

func (s *SignOutGroupRequest) GetDeviceId() *string {
	return s.DeviceId
}

func (s *SignOutGroupRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SignOutGroupRequest) GetUserId() *string {
	return s.UserId
}

func (s *SignOutGroupRequest) SetDeviceId(v string) *SignOutGroupRequest {
	s.DeviceId = &v
	return s
}

func (s *SignOutGroupRequest) SetInstanceId(v string) *SignOutGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *SignOutGroupRequest) SetUserId(v string) *SignOutGroupRequest {
	s.UserId = &v
	return s
}

func (s *SignOutGroupRequest) Validate() error {
	return dara.Validate(s)
}
