// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHoldCallRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChannelId(v string) *HoldCallRequest
	GetChannelId() *string
	SetDeviceId(v string) *HoldCallRequest
	GetDeviceId() *string
	SetInstanceId(v string) *HoldCallRequest
	GetInstanceId() *string
	SetJobId(v string) *HoldCallRequest
	GetJobId() *string
	SetMusic(v string) *HoldCallRequest
	GetMusic() *string
	SetUserId(v string) *HoldCallRequest
	GetUserId() *string
}

type HoldCallRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ch:customer:010123****->1318888****:1609255715825:job-6582589278232****
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// example:
	//
	// ACC-YUNBS-1.0.10-****
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
	// job-6582589278232****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	Music *string `json:"Music,omitempty" xml:"Music,omitempty"`
	// example:
	//
	// agent@ccc-test
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s HoldCallRequest) String() string {
	return dara.Prettify(s)
}

func (s HoldCallRequest) GoString() string {
	return s.String()
}

func (s *HoldCallRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *HoldCallRequest) GetDeviceId() *string {
	return s.DeviceId
}

func (s *HoldCallRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *HoldCallRequest) GetJobId() *string {
	return s.JobId
}

func (s *HoldCallRequest) GetMusic() *string {
	return s.Music
}

func (s *HoldCallRequest) GetUserId() *string {
	return s.UserId
}

func (s *HoldCallRequest) SetChannelId(v string) *HoldCallRequest {
	s.ChannelId = &v
	return s
}

func (s *HoldCallRequest) SetDeviceId(v string) *HoldCallRequest {
	s.DeviceId = &v
	return s
}

func (s *HoldCallRequest) SetInstanceId(v string) *HoldCallRequest {
	s.InstanceId = &v
	return s
}

func (s *HoldCallRequest) SetJobId(v string) *HoldCallRequest {
	s.JobId = &v
	return s
}

func (s *HoldCallRequest) SetMusic(v string) *HoldCallRequest {
	s.Music = &v
	return s
}

func (s *HoldCallRequest) SetUserId(v string) *HoldCallRequest {
	s.UserId = &v
	return s
}

func (s *HoldCallRequest) Validate() error {
	return dara.Validate(s)
}
