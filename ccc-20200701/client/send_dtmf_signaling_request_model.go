// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendDtmfSignalingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChannelId(v string) *SendDtmfSignalingRequest
	GetChannelId() *string
	SetDeviceId(v string) *SendDtmfSignalingRequest
	GetDeviceId() *string
	SetDtmf(v string) *SendDtmfSignalingRequest
	GetDtmf() *string
	SetInstanceId(v string) *SendDtmfSignalingRequest
	GetInstanceId() *string
	SetJobId(v string) *SendDtmfSignalingRequest
	GetJobId() *string
	SetUserId(v string) *SendDtmfSignalingRequest
	GetUserId() *string
}

type SendDtmfSignalingRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ch:customer:0108989****->1318888****:1609234221870:job-6573574060089****
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// example:
	//
	// device
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 5
	Dtmf *string `json:"Dtmf,omitempty" xml:"Dtmf,omitempty"`
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
	// job-6573574060089****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// agent@ccc-test
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s SendDtmfSignalingRequest) String() string {
	return dara.Prettify(s)
}

func (s SendDtmfSignalingRequest) GoString() string {
	return s.String()
}

func (s *SendDtmfSignalingRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *SendDtmfSignalingRequest) GetDeviceId() *string {
	return s.DeviceId
}

func (s *SendDtmfSignalingRequest) GetDtmf() *string {
	return s.Dtmf
}

func (s *SendDtmfSignalingRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SendDtmfSignalingRequest) GetJobId() *string {
	return s.JobId
}

func (s *SendDtmfSignalingRequest) GetUserId() *string {
	return s.UserId
}

func (s *SendDtmfSignalingRequest) SetChannelId(v string) *SendDtmfSignalingRequest {
	s.ChannelId = &v
	return s
}

func (s *SendDtmfSignalingRequest) SetDeviceId(v string) *SendDtmfSignalingRequest {
	s.DeviceId = &v
	return s
}

func (s *SendDtmfSignalingRequest) SetDtmf(v string) *SendDtmfSignalingRequest {
	s.Dtmf = &v
	return s
}

func (s *SendDtmfSignalingRequest) SetInstanceId(v string) *SendDtmfSignalingRequest {
	s.InstanceId = &v
	return s
}

func (s *SendDtmfSignalingRequest) SetJobId(v string) *SendDtmfSignalingRequest {
	s.JobId = &v
	return s
}

func (s *SendDtmfSignalingRequest) SetUserId(v string) *SendDtmfSignalingRequest {
	s.UserId = &v
	return s
}

func (s *SendDtmfSignalingRequest) Validate() error {
	return dara.Validate(s)
}
