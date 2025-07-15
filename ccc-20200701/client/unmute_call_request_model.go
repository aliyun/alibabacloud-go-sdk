// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnmuteCallRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChannelId(v string) *UnmuteCallRequest
	GetChannelId() *string
	SetDeviceId(v string) *UnmuteCallRequest
	GetDeviceId() *string
	SetInstanceId(v string) *UnmuteCallRequest
	GetInstanceId() *string
	SetJobId(v string) *UnmuteCallRequest
	GetJobId() *string
	SetUserId(v string) *UnmuteCallRequest
	GetUserId() *string
}

type UnmuteCallRequest struct {
	// example:
	//
	// ch:user:1390501****->8032****:1609138902226:job-6538214103685****
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
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
	// job-6538214103685****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// agent@ccc-test
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s UnmuteCallRequest) String() string {
	return dara.Prettify(s)
}

func (s UnmuteCallRequest) GoString() string {
	return s.String()
}

func (s *UnmuteCallRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *UnmuteCallRequest) GetDeviceId() *string {
	return s.DeviceId
}

func (s *UnmuteCallRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UnmuteCallRequest) GetJobId() *string {
	return s.JobId
}

func (s *UnmuteCallRequest) GetUserId() *string {
	return s.UserId
}

func (s *UnmuteCallRequest) SetChannelId(v string) *UnmuteCallRequest {
	s.ChannelId = &v
	return s
}

func (s *UnmuteCallRequest) SetDeviceId(v string) *UnmuteCallRequest {
	s.DeviceId = &v
	return s
}

func (s *UnmuteCallRequest) SetInstanceId(v string) *UnmuteCallRequest {
	s.InstanceId = &v
	return s
}

func (s *UnmuteCallRequest) SetJobId(v string) *UnmuteCallRequest {
	s.JobId = &v
	return s
}

func (s *UnmuteCallRequest) SetUserId(v string) *UnmuteCallRequest {
	s.UserId = &v
	return s
}

func (s *UnmuteCallRequest) Validate() error {
	return dara.Validate(s)
}
