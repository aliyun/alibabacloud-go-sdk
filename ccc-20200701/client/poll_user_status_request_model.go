// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPollUserStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceId(v string) *PollUserStatusRequest
	GetDeviceId() *string
	SetInstanceId(v string) *PollUserStatusRequest
	GetInstanceId() *string
	SetUserId(v string) *PollUserStatusRequest
	GetUserId() *string
}

type PollUserStatusRequest struct {
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

func (s PollUserStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s PollUserStatusRequest) GoString() string {
	return s.String()
}

func (s *PollUserStatusRequest) GetDeviceId() *string {
	return s.DeviceId
}

func (s *PollUserStatusRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *PollUserStatusRequest) GetUserId() *string {
	return s.UserId
}

func (s *PollUserStatusRequest) SetDeviceId(v string) *PollUserStatusRequest {
	s.DeviceId = &v
	return s
}

func (s *PollUserStatusRequest) SetInstanceId(v string) *PollUserStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *PollUserStatusRequest) SetUserId(v string) *PollUserStatusRequest {
	s.UserId = &v
	return s
}

func (s *PollUserStatusRequest) Validate() error {
	return dara.Validate(s)
}
