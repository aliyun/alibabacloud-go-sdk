// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBridgeRtcCallRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallee(v string) *BridgeRtcCallRequest
	GetCallee() *string
	SetCaller(v string) *BridgeRtcCallRequest
	GetCaller() *string
	SetDeviceId(v string) *BridgeRtcCallRequest
	GetDeviceId() *string
	SetInstanceId(v string) *BridgeRtcCallRequest
	GetInstanceId() *string
	SetServiceProvider(v string) *BridgeRtcCallRequest
	GetServiceProvider() *string
	SetTags(v string) *BridgeRtcCallRequest
	GetTags() *string
	SetTimeoutSeconds(v int32) *BridgeRtcCallRequest
	GetTimeoutSeconds() *int32
	SetUserId(v string) *BridgeRtcCallRequest
	GetUserId() *string
	SetVideoEnabled(v bool) *BridgeRtcCallRequest
	GetVideoEnabled() *bool
}

type BridgeRtcCallRequest struct {
	// This parameter is required.
	Callee *string `json:"Callee,omitempty" xml:"Callee,omitempty"`
	Caller *string `json:"Caller,omitempty" xml:"Caller,omitempty"`
	// This parameter is required.
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// This parameter is required.
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ServiceProvider *string `json:"ServiceProvider,omitempty" xml:"ServiceProvider,omitempty"`
	Tags            *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	TimeoutSeconds  *int32  `json:"TimeoutSeconds,omitempty" xml:"TimeoutSeconds,omitempty"`
	UserId          *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	VideoEnabled    *bool   `json:"VideoEnabled,omitempty" xml:"VideoEnabled,omitempty"`
}

func (s BridgeRtcCallRequest) String() string {
	return dara.Prettify(s)
}

func (s BridgeRtcCallRequest) GoString() string {
	return s.String()
}

func (s *BridgeRtcCallRequest) GetCallee() *string {
	return s.Callee
}

func (s *BridgeRtcCallRequest) GetCaller() *string {
	return s.Caller
}

func (s *BridgeRtcCallRequest) GetDeviceId() *string {
	return s.DeviceId
}

func (s *BridgeRtcCallRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *BridgeRtcCallRequest) GetServiceProvider() *string {
	return s.ServiceProvider
}

func (s *BridgeRtcCallRequest) GetTags() *string {
	return s.Tags
}

func (s *BridgeRtcCallRequest) GetTimeoutSeconds() *int32 {
	return s.TimeoutSeconds
}

func (s *BridgeRtcCallRequest) GetUserId() *string {
	return s.UserId
}

func (s *BridgeRtcCallRequest) GetVideoEnabled() *bool {
	return s.VideoEnabled
}

func (s *BridgeRtcCallRequest) SetCallee(v string) *BridgeRtcCallRequest {
	s.Callee = &v
	return s
}

func (s *BridgeRtcCallRequest) SetCaller(v string) *BridgeRtcCallRequest {
	s.Caller = &v
	return s
}

func (s *BridgeRtcCallRequest) SetDeviceId(v string) *BridgeRtcCallRequest {
	s.DeviceId = &v
	return s
}

func (s *BridgeRtcCallRequest) SetInstanceId(v string) *BridgeRtcCallRequest {
	s.InstanceId = &v
	return s
}

func (s *BridgeRtcCallRequest) SetServiceProvider(v string) *BridgeRtcCallRequest {
	s.ServiceProvider = &v
	return s
}

func (s *BridgeRtcCallRequest) SetTags(v string) *BridgeRtcCallRequest {
	s.Tags = &v
	return s
}

func (s *BridgeRtcCallRequest) SetTimeoutSeconds(v int32) *BridgeRtcCallRequest {
	s.TimeoutSeconds = &v
	return s
}

func (s *BridgeRtcCallRequest) SetUserId(v string) *BridgeRtcCallRequest {
	s.UserId = &v
	return s
}

func (s *BridgeRtcCallRequest) SetVideoEnabled(v bool) *BridgeRtcCallRequest {
	s.VideoEnabled = &v
	return s
}

func (s *BridgeRtcCallRequest) Validate() error {
	return dara.Validate(s)
}
