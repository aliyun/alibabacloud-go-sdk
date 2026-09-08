// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBargeInCallRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBargedUserId(v string) *BargeInCallRequest
	GetBargedUserId() *string
	SetDeviceId(v string) *BargeInCallRequest
	GetDeviceId() *string
	SetInstanceId(v string) *BargeInCallRequest
	GetInstanceId() *string
	SetJobId(v string) *BargeInCallRequest
	GetJobId() *string
	SetTimeoutSeconds(v int32) *BargeInCallRequest
	GetTimeoutSeconds() *int32
	SetUserId(v string) *BargeInCallRequest
	GetUserId() *string
}

type BargeInCallRequest struct {
	// Agent ID whose call was barged in on.
	//
	// This parameter is required.
	//
	// example:
	//
	// agent2@ccc-test
	BargedUserId *string `json:"BargedUserId,omitempty" xml:"BargedUserId,omitempty"`
	// Device ID. This parameter is meaningless and can be filled in with any value.
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
	// Call ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// job-6538214103685****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// Timeout for force insert, in seconds. If the force insert operation does not succeed within the specified time, it is canceled. Normally, the force insert operation succeeds immediately. The timeout setting is provided to handle abnormal scenarios. This field is optional and defaults to 30 seconds.
	//
	// example:
	//
	// 60
	TimeoutSeconds *int32 `json:"TimeoutSeconds,omitempty" xml:"TimeoutSeconds,omitempty"`
	// Agent ID initiating the force insert.
	//
	// example:
	//
	// agent@ccc-test
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s BargeInCallRequest) String() string {
	return dara.Prettify(s)
}

func (s BargeInCallRequest) GoString() string {
	return s.String()
}

func (s *BargeInCallRequest) GetBargedUserId() *string {
	return s.BargedUserId
}

func (s *BargeInCallRequest) GetDeviceId() *string {
	return s.DeviceId
}

func (s *BargeInCallRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *BargeInCallRequest) GetJobId() *string {
	return s.JobId
}

func (s *BargeInCallRequest) GetTimeoutSeconds() *int32 {
	return s.TimeoutSeconds
}

func (s *BargeInCallRequest) GetUserId() *string {
	return s.UserId
}

func (s *BargeInCallRequest) SetBargedUserId(v string) *BargeInCallRequest {
	s.BargedUserId = &v
	return s
}

func (s *BargeInCallRequest) SetDeviceId(v string) *BargeInCallRequest {
	s.DeviceId = &v
	return s
}

func (s *BargeInCallRequest) SetInstanceId(v string) *BargeInCallRequest {
	s.InstanceId = &v
	return s
}

func (s *BargeInCallRequest) SetJobId(v string) *BargeInCallRequest {
	s.JobId = &v
	return s
}

func (s *BargeInCallRequest) SetTimeoutSeconds(v int32) *BargeInCallRequest {
	s.TimeoutSeconds = &v
	return s
}

func (s *BargeInCallRequest) SetUserId(v string) *BargeInCallRequest {
	s.UserId = &v
	return s
}

func (s *BargeInCallRequest) Validate() error {
	return dara.Validate(s)
}
