// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCoachCallRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCoachedUserId(v string) *CoachCallRequest
	GetCoachedUserId() *string
	SetDeviceId(v string) *CoachCallRequest
	GetDeviceId() *string
	SetInstanceId(v string) *CoachCallRequest
	GetInstanceId() *string
	SetJobId(v string) *CoachCallRequest
	GetJobId() *string
	SetTimeoutSeconds(v int32) *CoachCallRequest
	GetTimeoutSeconds() *int32
	SetUserId(v string) *CoachCallRequest
	GetUserId() *string
}

type CoachCallRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// coached-agent@ccc-test
	CoachedUserId *string `json:"CoachedUserId,omitempty" xml:"CoachedUserId,omitempty"`
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
	// job-65382141036853491
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// 30
	TimeoutSeconds *int32 `json:"TimeoutSeconds,omitempty" xml:"TimeoutSeconds,omitempty"`
	// example:
	//
	// agent@ccc-test
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s CoachCallRequest) String() string {
	return dara.Prettify(s)
}

func (s CoachCallRequest) GoString() string {
	return s.String()
}

func (s *CoachCallRequest) GetCoachedUserId() *string {
	return s.CoachedUserId
}

func (s *CoachCallRequest) GetDeviceId() *string {
	return s.DeviceId
}

func (s *CoachCallRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CoachCallRequest) GetJobId() *string {
	return s.JobId
}

func (s *CoachCallRequest) GetTimeoutSeconds() *int32 {
	return s.TimeoutSeconds
}

func (s *CoachCallRequest) GetUserId() *string {
	return s.UserId
}

func (s *CoachCallRequest) SetCoachedUserId(v string) *CoachCallRequest {
	s.CoachedUserId = &v
	return s
}

func (s *CoachCallRequest) SetDeviceId(v string) *CoachCallRequest {
	s.DeviceId = &v
	return s
}

func (s *CoachCallRequest) SetInstanceId(v string) *CoachCallRequest {
	s.InstanceId = &v
	return s
}

func (s *CoachCallRequest) SetJobId(v string) *CoachCallRequest {
	s.JobId = &v
	return s
}

func (s *CoachCallRequest) SetTimeoutSeconds(v int32) *CoachCallRequest {
	s.TimeoutSeconds = &v
	return s
}

func (s *CoachCallRequest) SetUserId(v string) *CoachCallRequest {
	s.UserId = &v
	return s
}

func (s *CoachCallRequest) Validate() error {
	return dara.Validate(s)
}
