// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRetrieveCallRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChannelId(v string) *RetrieveCallRequest
	GetChannelId() *string
	SetDeviceId(v string) *RetrieveCallRequest
	GetDeviceId() *string
	SetInstanceId(v string) *RetrieveCallRequest
	GetInstanceId() *string
	SetJobId(v string) *RetrieveCallRequest
	GetJobId() *string
	SetUserId(v string) *RetrieveCallRequest
	GetUserId() *string
}

type RetrieveCallRequest struct {
	// This parameter is required.
	//
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

func (s RetrieveCallRequest) String() string {
	return dara.Prettify(s)
}

func (s RetrieveCallRequest) GoString() string {
	return s.String()
}

func (s *RetrieveCallRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *RetrieveCallRequest) GetDeviceId() *string {
	return s.DeviceId
}

func (s *RetrieveCallRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RetrieveCallRequest) GetJobId() *string {
	return s.JobId
}

func (s *RetrieveCallRequest) GetUserId() *string {
	return s.UserId
}

func (s *RetrieveCallRequest) SetChannelId(v string) *RetrieveCallRequest {
	s.ChannelId = &v
	return s
}

func (s *RetrieveCallRequest) SetDeviceId(v string) *RetrieveCallRequest {
	s.DeviceId = &v
	return s
}

func (s *RetrieveCallRequest) SetInstanceId(v string) *RetrieveCallRequest {
	s.InstanceId = &v
	return s
}

func (s *RetrieveCallRequest) SetJobId(v string) *RetrieveCallRequest {
	s.JobId = &v
	return s
}

func (s *RetrieveCallRequest) SetUserId(v string) *RetrieveCallRequest {
	s.UserId = &v
	return s
}

func (s *RetrieveCallRequest) Validate() error {
	return dara.Validate(s)
}
