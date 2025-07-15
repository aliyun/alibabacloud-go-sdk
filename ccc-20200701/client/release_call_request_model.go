// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseCallRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChannelId(v string) *ReleaseCallRequest
	GetChannelId() *string
	SetDeviceId(v string) *ReleaseCallRequest
	GetDeviceId() *string
	SetInstanceId(v string) *ReleaseCallRequest
	GetInstanceId() *string
	SetJobId(v string) *ReleaseCallRequest
	GetJobId() *string
	SetUserId(v string) *ReleaseCallRequest
	GetUserId() *string
}

type ReleaseCallRequest struct {
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
	// This parameter is required.
	//
	// example:
	//
	// job-6538214103685****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// agent@ccc-test
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ReleaseCallRequest) String() string {
	return dara.Prettify(s)
}

func (s ReleaseCallRequest) GoString() string {
	return s.String()
}

func (s *ReleaseCallRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *ReleaseCallRequest) GetDeviceId() *string {
	return s.DeviceId
}

func (s *ReleaseCallRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ReleaseCallRequest) GetJobId() *string {
	return s.JobId
}

func (s *ReleaseCallRequest) GetUserId() *string {
	return s.UserId
}

func (s *ReleaseCallRequest) SetChannelId(v string) *ReleaseCallRequest {
	s.ChannelId = &v
	return s
}

func (s *ReleaseCallRequest) SetDeviceId(v string) *ReleaseCallRequest {
	s.DeviceId = &v
	return s
}

func (s *ReleaseCallRequest) SetInstanceId(v string) *ReleaseCallRequest {
	s.InstanceId = &v
	return s
}

func (s *ReleaseCallRequest) SetJobId(v string) *ReleaseCallRequest {
	s.JobId = &v
	return s
}

func (s *ReleaseCallRequest) SetUserId(v string) *ReleaseCallRequest {
	s.UserId = &v
	return s
}

func (s *ReleaseCallRequest) Validate() error {
	return dara.Validate(s)
}
