// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInterceptCallRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceId(v string) *InterceptCallRequest
	GetDeviceId() *string
	SetInstanceId(v string) *InterceptCallRequest
	GetInstanceId() *string
	SetInterceptedUserId(v string) *InterceptCallRequest
	GetInterceptedUserId() *string
	SetJobId(v string) *InterceptCallRequest
	GetJobId() *string
	SetTimeoutSeconds(v int32) *InterceptCallRequest
	GetTimeoutSeconds() *int32
	SetUserId(v string) *InterceptCallRequest
	GetUserId() *string
}

type InterceptCallRequest struct {
	// Device ID. This parameter is meaningless and can be filled in with any value.
	//
	// example:
	//
	// device
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// 实例 ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The agent ID that is forcibly disconnected.
	//
	// This parameter is required.
	//
	// example:
	//
	// agent2@ccc-test
	InterceptedUserId *string `json:"InterceptedUserId,omitempty" xml:"InterceptedUserId,omitempty"`
	// The call ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// job-6538214103685****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// 强拆超时时间，经过指定的时间强拆仍未成功，则取消强拆，正常情况下，强拆操作会马上成功，设置超时时间是为了防止异常发生，此字段选填，默认 30，单位秒。
	//
	// example:
	//
	// 60
	TimeoutSeconds *int32 `json:"TimeoutSeconds,omitempty" xml:"TimeoutSeconds,omitempty"`
	// 发起强拆的坐席。
	//
	// example:
	//
	// agent@ccc-test
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s InterceptCallRequest) String() string {
	return dara.Prettify(s)
}

func (s InterceptCallRequest) GoString() string {
	return s.String()
}

func (s *InterceptCallRequest) GetDeviceId() *string {
	return s.DeviceId
}

func (s *InterceptCallRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *InterceptCallRequest) GetInterceptedUserId() *string {
	return s.InterceptedUserId
}

func (s *InterceptCallRequest) GetJobId() *string {
	return s.JobId
}

func (s *InterceptCallRequest) GetTimeoutSeconds() *int32 {
	return s.TimeoutSeconds
}

func (s *InterceptCallRequest) GetUserId() *string {
	return s.UserId
}

func (s *InterceptCallRequest) SetDeviceId(v string) *InterceptCallRequest {
	s.DeviceId = &v
	return s
}

func (s *InterceptCallRequest) SetInstanceId(v string) *InterceptCallRequest {
	s.InstanceId = &v
	return s
}

func (s *InterceptCallRequest) SetInterceptedUserId(v string) *InterceptCallRequest {
	s.InterceptedUserId = &v
	return s
}

func (s *InterceptCallRequest) SetJobId(v string) *InterceptCallRequest {
	s.JobId = &v
	return s
}

func (s *InterceptCallRequest) SetTimeoutSeconds(v int32) *InterceptCallRequest {
	s.TimeoutSeconds = &v
	return s
}

func (s *InterceptCallRequest) SetUserId(v string) *InterceptCallRequest {
	s.UserId = &v
	return s
}

func (s *InterceptCallRequest) Validate() error {
	return dara.Validate(s)
}
