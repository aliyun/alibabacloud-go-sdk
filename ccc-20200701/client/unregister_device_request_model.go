// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnregisterDeviceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *UnregisterDeviceRequest
	GetInstanceId() *string
	SetUserId(v string) *UnregisterDeviceRequest
	GetUserId() *string
}

type UnregisterDeviceRequest struct {
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

func (s UnregisterDeviceRequest) String() string {
	return dara.Prettify(s)
}

func (s UnregisterDeviceRequest) GoString() string {
	return s.String()
}

func (s *UnregisterDeviceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UnregisterDeviceRequest) GetUserId() *string {
	return s.UserId
}

func (s *UnregisterDeviceRequest) SetInstanceId(v string) *UnregisterDeviceRequest {
	s.InstanceId = &v
	return s
}

func (s *UnregisterDeviceRequest) SetUserId(v string) *UnregisterDeviceRequest {
	s.UserId = &v
	return s
}

func (s *UnregisterDeviceRequest) Validate() error {
	return dara.Validate(s)
}
