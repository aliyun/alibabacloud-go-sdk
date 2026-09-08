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
	// The channel ID to unmute. This parameter is optional and defaults to empty. If empty, the system unmutes the channel associated with the agent specified by UserId.
	//
	// example:
	//
	// ch:user:1390501****->8032****:1609138902226:job-6538214103685****
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// Device ID. This parameter is meaningless and can be filled with any value.
	//
	// example:
	//
	// device
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// Instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The call ID.
	//
	// example:
	//
	// job-6538214103685****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The agent ID to unmute. If not specified, defaults to the agent mapped to the current RAM account.
	//
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
