// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBridgeWebCallRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *BridgeWebCallRequest
	GetApplicationId() *string
	SetBusinessUnitId(v string) *BridgeWebCallRequest
	GetBusinessUnitId() *string
	SetCaller(v string) *BridgeWebCallRequest
	GetCaller() *string
	SetDeviceId(v string) *BridgeWebCallRequest
	GetDeviceId() *string
	SetSampleRate(v int32) *BridgeWebCallRequest
	GetSampleRate() *int32
	SetSandbox(v bool) *BridgeWebCallRequest
	GetSandbox() *bool
	SetTags(v string) *BridgeWebCallRequest
	GetTags() *string
	SetTimeoutSeconds(v int32) *BridgeWebCallRequest
	GetTimeoutSeconds() *int32
}

type BridgeWebCallRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// a395011f-a247-400f-bc69-28796749fd52
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-c11iig67g863rih8
	BusinessUnitId *string `json:"BusinessUnitId,omitempty" xml:"BusinessUnitId,omitempty"`
	// example:
	//
	// 13052253537
	Caller *string `json:"Caller,omitempty" xml:"Caller,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 467539456766097392-cn-shenzhen
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// example:
	//
	// 8000
	SampleRate *int32 `json:"SampleRate,omitempty" xml:"SampleRate,omitempty"`
	// example:
	//
	// true
	Sandbox *bool `json:"Sandbox,omitempty" xml:"Sandbox,omitempty"`
	// example:
	//
	// {\\"ENV\\": \\"production\\"}
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// example:
	//
	// 3
	TimeoutSeconds *int32 `json:"TimeoutSeconds,omitempty" xml:"TimeoutSeconds,omitempty"`
}

func (s BridgeWebCallRequest) String() string {
	return dara.Prettify(s)
}

func (s BridgeWebCallRequest) GoString() string {
	return s.String()
}

func (s *BridgeWebCallRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *BridgeWebCallRequest) GetBusinessUnitId() *string {
	return s.BusinessUnitId
}

func (s *BridgeWebCallRequest) GetCaller() *string {
	return s.Caller
}

func (s *BridgeWebCallRequest) GetDeviceId() *string {
	return s.DeviceId
}

func (s *BridgeWebCallRequest) GetSampleRate() *int32 {
	return s.SampleRate
}

func (s *BridgeWebCallRequest) GetSandbox() *bool {
	return s.Sandbox
}

func (s *BridgeWebCallRequest) GetTags() *string {
	return s.Tags
}

func (s *BridgeWebCallRequest) GetTimeoutSeconds() *int32 {
	return s.TimeoutSeconds
}

func (s *BridgeWebCallRequest) SetApplicationId(v string) *BridgeWebCallRequest {
	s.ApplicationId = &v
	return s
}

func (s *BridgeWebCallRequest) SetBusinessUnitId(v string) *BridgeWebCallRequest {
	s.BusinessUnitId = &v
	return s
}

func (s *BridgeWebCallRequest) SetCaller(v string) *BridgeWebCallRequest {
	s.Caller = &v
	return s
}

func (s *BridgeWebCallRequest) SetDeviceId(v string) *BridgeWebCallRequest {
	s.DeviceId = &v
	return s
}

func (s *BridgeWebCallRequest) SetSampleRate(v int32) *BridgeWebCallRequest {
	s.SampleRate = &v
	return s
}

func (s *BridgeWebCallRequest) SetSandbox(v bool) *BridgeWebCallRequest {
	s.Sandbox = &v
	return s
}

func (s *BridgeWebCallRequest) SetTags(v string) *BridgeWebCallRequest {
	s.Tags = &v
	return s
}

func (s *BridgeWebCallRequest) SetTimeoutSeconds(v int32) *BridgeWebCallRequest {
	s.TimeoutSeconds = &v
	return s
}

func (s *BridgeWebCallRequest) Validate() error {
	return dara.Validate(s)
}
