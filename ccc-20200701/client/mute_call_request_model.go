// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMuteCallRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChannelId(v string) *MuteCallRequest
	GetChannelId() *string
	SetDeviceId(v string) *MuteCallRequest
	GetDeviceId() *string
	SetInstanceId(v string) *MuteCallRequest
	GetInstanceId() *string
	SetJobId(v string) *MuteCallRequest
	GetJobId() *string
	SetUserId(v string) *MuteCallRequest
	GetUserId() *string
}

type MuteCallRequest struct {
	// The ID of the voice channel to be muted. This parameter is optional and defaults to empty. If it is empty, the voice channel of the agent corresponding to UserId will be muted.
	//
	// example:
	//
	// ch:user:1318888****->8001****:1609253204816:job-6581536084722****
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
	// This parameter is required.
	//
	// example:
	//
	// job-6581536084722****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The agent ID to be muted.
	//
	// example:
	//
	// agent@ccc-test
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s MuteCallRequest) String() string {
	return dara.Prettify(s)
}

func (s MuteCallRequest) GoString() string {
	return s.String()
}

func (s *MuteCallRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *MuteCallRequest) GetDeviceId() *string {
	return s.DeviceId
}

func (s *MuteCallRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *MuteCallRequest) GetJobId() *string {
	return s.JobId
}

func (s *MuteCallRequest) GetUserId() *string {
	return s.UserId
}

func (s *MuteCallRequest) SetChannelId(v string) *MuteCallRequest {
	s.ChannelId = &v
	return s
}

func (s *MuteCallRequest) SetDeviceId(v string) *MuteCallRequest {
	s.DeviceId = &v
	return s
}

func (s *MuteCallRequest) SetInstanceId(v string) *MuteCallRequest {
	s.InstanceId = &v
	return s
}

func (s *MuteCallRequest) SetJobId(v string) *MuteCallRequest {
	s.JobId = &v
	return s
}

func (s *MuteCallRequest) SetUserId(v string) *MuteCallRequest {
	s.UserId = &v
	return s
}

func (s *MuteCallRequest) Validate() error {
	return dara.Validate(s)
}
