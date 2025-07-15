// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRedialCallRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallee(v string) *RedialCallRequest
	GetCallee() *string
	SetCaller(v string) *RedialCallRequest
	GetCaller() *string
	SetDeviceId(v string) *RedialCallRequest
	GetDeviceId() *string
	SetInstanceId(v string) *RedialCallRequest
	GetInstanceId() *string
	SetJobId(v string) *RedialCallRequest
	GetJobId() *string
	SetTags(v string) *RedialCallRequest
	GetTags() *string
	SetTimeoutSeconds(v int32) *RedialCallRequest
	GetTimeoutSeconds() *int32
	SetUserId(v string) *RedialCallRequest
	GetUserId() *string
}

type RedialCallRequest struct {
	// example:
	//
	// 1318888****
	Callee *string `json:"Callee,omitempty" xml:"Callee,omitempty"`
	// example:
	//
	// 0109810****
	Caller   *string `json:"Caller,omitempty" xml:"Caller,omitempty"`
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// abc
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// job-6581536084722****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// a=b
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// example:
	//
	// 30
	TimeoutSeconds *int32 `json:"TimeoutSeconds,omitempty" xml:"TimeoutSeconds,omitempty"`
	// example:
	//
	// samzhang@abc
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s RedialCallRequest) String() string {
	return dara.Prettify(s)
}

func (s RedialCallRequest) GoString() string {
	return s.String()
}

func (s *RedialCallRequest) GetCallee() *string {
	return s.Callee
}

func (s *RedialCallRequest) GetCaller() *string {
	return s.Caller
}

func (s *RedialCallRequest) GetDeviceId() *string {
	return s.DeviceId
}

func (s *RedialCallRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RedialCallRequest) GetJobId() *string {
	return s.JobId
}

func (s *RedialCallRequest) GetTags() *string {
	return s.Tags
}

func (s *RedialCallRequest) GetTimeoutSeconds() *int32 {
	return s.TimeoutSeconds
}

func (s *RedialCallRequest) GetUserId() *string {
	return s.UserId
}

func (s *RedialCallRequest) SetCallee(v string) *RedialCallRequest {
	s.Callee = &v
	return s
}

func (s *RedialCallRequest) SetCaller(v string) *RedialCallRequest {
	s.Caller = &v
	return s
}

func (s *RedialCallRequest) SetDeviceId(v string) *RedialCallRequest {
	s.DeviceId = &v
	return s
}

func (s *RedialCallRequest) SetInstanceId(v string) *RedialCallRequest {
	s.InstanceId = &v
	return s
}

func (s *RedialCallRequest) SetJobId(v string) *RedialCallRequest {
	s.JobId = &v
	return s
}

func (s *RedialCallRequest) SetTags(v string) *RedialCallRequest {
	s.Tags = &v
	return s
}

func (s *RedialCallRequest) SetTimeoutSeconds(v int32) *RedialCallRequest {
	s.TimeoutSeconds = &v
	return s
}

func (s *RedialCallRequest) SetUserId(v string) *RedialCallRequest {
	s.UserId = &v
	return s
}

func (s *RedialCallRequest) Validate() error {
	return dara.Validate(s)
}
