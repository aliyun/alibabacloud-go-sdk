// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMonitorCallRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceId(v string) *MonitorCallRequest
	GetDeviceId() *string
	SetInstanceId(v string) *MonitorCallRequest
	GetInstanceId() *string
	SetMonitoredUserId(v string) *MonitorCallRequest
	GetMonitoredUserId() *string
	SetTimeoutSeconds(v int32) *MonitorCallRequest
	GetTimeoutSeconds() *int32
	SetUserId(v string) *MonitorCallRequest
	GetUserId() *string
}

type MonitorCallRequest struct {
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
	// The ID of the agent being monitored.
	//
	// This parameter is required.
	//
	// example:
	//
	// agent2@ccc-test
	MonitoredUserId *string `json:"MonitoredUserId,omitempty" xml:"MonitoredUserId,omitempty"`
	// The timeout period for the listening operation, in seconds. If the listening operation does not succeed within the specified time, it is canceled. Normally, the listening operation succeeds immediately. The timeout setting is provided to handle abnormal scenarios. This field is optional and defaults to 30 seconds.
	//
	// example:
	//
	// 30
	TimeoutSeconds *int32 `json:"TimeoutSeconds,omitempty" xml:"TimeoutSeconds,omitempty"`
	// Agent ID.
	//
	// example:
	//
	// agent@ccc-test
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s MonitorCallRequest) String() string {
	return dara.Prettify(s)
}

func (s MonitorCallRequest) GoString() string {
	return s.String()
}

func (s *MonitorCallRequest) GetDeviceId() *string {
	return s.DeviceId
}

func (s *MonitorCallRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *MonitorCallRequest) GetMonitoredUserId() *string {
	return s.MonitoredUserId
}

func (s *MonitorCallRequest) GetTimeoutSeconds() *int32 {
	return s.TimeoutSeconds
}

func (s *MonitorCallRequest) GetUserId() *string {
	return s.UserId
}

func (s *MonitorCallRequest) SetDeviceId(v string) *MonitorCallRequest {
	s.DeviceId = &v
	return s
}

func (s *MonitorCallRequest) SetInstanceId(v string) *MonitorCallRequest {
	s.InstanceId = &v
	return s
}

func (s *MonitorCallRequest) SetMonitoredUserId(v string) *MonitorCallRequest {
	s.MonitoredUserId = &v
	return s
}

func (s *MonitorCallRequest) SetTimeoutSeconds(v int32) *MonitorCallRequest {
	s.TimeoutSeconds = &v
	return s
}

func (s *MonitorCallRequest) SetUserId(v string) *MonitorCallRequest {
	s.UserId = &v
	return s
}

func (s *MonitorCallRequest) Validate() error {
	return dara.Validate(s)
}
