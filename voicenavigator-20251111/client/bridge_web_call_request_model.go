// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBridgeWebCallRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessChannelId(v string) *BridgeWebCallRequest
	GetAccessChannelId() *string
	SetAccessChannelType(v string) *BridgeWebCallRequest
	GetAccessChannelType() *string
	SetCaller(v string) *BridgeWebCallRequest
	GetCaller() *string
	SetDeviceId(v string) *BridgeWebCallRequest
	GetDeviceId() *string
	SetDraftVersion(v bool) *BridgeWebCallRequest
	GetDraftVersion() *bool
	SetInstanceId(v string) *BridgeWebCallRequest
	GetInstanceId() *string
	SetSampleRate(v int32) *BridgeWebCallRequest
	GetSampleRate() *int32
	SetScriptId(v string) *BridgeWebCallRequest
	GetScriptId() *string
	SetTags(v string) *BridgeWebCallRequest
	GetTags() *string
	SetTimeoutSeconds(v int32) *BridgeWebCallRequest
	GetTimeoutSeconds() *int32
}

type BridgeWebCallRequest struct {
	// example:
	//
	// 33606503-c22c-4547-a51c-dda5e8d87962
	AccessChannelId *string `json:"AccessChannelId,omitempty" xml:"AccessChannelId,omitempty"`
	// example:
	//
	// Text
	AccessChannelType *string `json:"AccessChannelType,omitempty" xml:"AccessChannelType,omitempty"`
	// example:
	//
	// 010123*****
	Caller *string `json:"Caller,omitempty" xml:"Caller,omitempty"`
	// example:
	//
	// d4c38420-****-43bc-b001-56bfc274****
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// example:
	//
	// false
	DraftVersion *bool `json:"DraftVersion,omitempty" xml:"DraftVersion,omitempty"`
	// example:
	//
	// 36e9a4cd-a821-4f78-86fa-a9a4aefeea2e
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 8000
	SampleRate *int32 `json:"SampleRate,omitempty" xml:"SampleRate,omitempty"`
	// example:
	//
	// 0a88e269-01f5-49ac-97af-5830f0ccb271
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
	// example:
	//
	// {"key":"value"}
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// example:
	//
	// 60
	TimeoutSeconds *int32 `json:"TimeoutSeconds,omitempty" xml:"TimeoutSeconds,omitempty"`
}

func (s BridgeWebCallRequest) String() string {
	return dara.Prettify(s)
}

func (s BridgeWebCallRequest) GoString() string {
	return s.String()
}

func (s *BridgeWebCallRequest) GetAccessChannelId() *string {
	return s.AccessChannelId
}

func (s *BridgeWebCallRequest) GetAccessChannelType() *string {
	return s.AccessChannelType
}

func (s *BridgeWebCallRequest) GetCaller() *string {
	return s.Caller
}

func (s *BridgeWebCallRequest) GetDeviceId() *string {
	return s.DeviceId
}

func (s *BridgeWebCallRequest) GetDraftVersion() *bool {
	return s.DraftVersion
}

func (s *BridgeWebCallRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *BridgeWebCallRequest) GetSampleRate() *int32 {
	return s.SampleRate
}

func (s *BridgeWebCallRequest) GetScriptId() *string {
	return s.ScriptId
}

func (s *BridgeWebCallRequest) GetTags() *string {
	return s.Tags
}

func (s *BridgeWebCallRequest) GetTimeoutSeconds() *int32 {
	return s.TimeoutSeconds
}

func (s *BridgeWebCallRequest) SetAccessChannelId(v string) *BridgeWebCallRequest {
	s.AccessChannelId = &v
	return s
}

func (s *BridgeWebCallRequest) SetAccessChannelType(v string) *BridgeWebCallRequest {
	s.AccessChannelType = &v
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

func (s *BridgeWebCallRequest) SetDraftVersion(v bool) *BridgeWebCallRequest {
	s.DraftVersion = &v
	return s
}

func (s *BridgeWebCallRequest) SetInstanceId(v string) *BridgeWebCallRequest {
	s.InstanceId = &v
	return s
}

func (s *BridgeWebCallRequest) SetSampleRate(v int32) *BridgeWebCallRequest {
	s.SampleRate = &v
	return s
}

func (s *BridgeWebCallRequest) SetScriptId(v string) *BridgeWebCallRequest {
	s.ScriptId = &v
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
