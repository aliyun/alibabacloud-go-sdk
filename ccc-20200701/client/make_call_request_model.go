// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMakeCallRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallee(v string) *MakeCallRequest
	GetCallee() *string
	SetCaller(v string) *MakeCallRequest
	GetCaller() *string
	SetDeviceId(v string) *MakeCallRequest
	GetDeviceId() *string
	SetFlashSmsVariables(v string) *MakeCallRequest
	GetFlashSmsVariables() *string
	SetInstanceId(v string) *MakeCallRequest
	GetInstanceId() *string
	SetMaskedCallee(v string) *MakeCallRequest
	GetMaskedCallee() *string
	SetMediaType(v string) *MakeCallRequest
	GetMediaType() *string
	SetTags(v string) *MakeCallRequest
	GetTags() *string
	SetTimeoutSeconds(v int32) *MakeCallRequest
	GetTimeoutSeconds() *int32
	SetUserId(v string) *MakeCallRequest
	GetUserId() *string
}

type MakeCallRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1318888****
	Callee *string `json:"Callee,omitempty" xml:"Callee,omitempty"`
	// example:
	//
	// 010989****
	Caller *string `json:"Caller,omitempty" xml:"Caller,omitempty"`
	// example:
	//
	// device
	DeviceId          *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	FlashSmsVariables *string `json:"FlashSmsVariables,omitempty" xml:"FlashSmsVariables,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 131****8888
	MaskedCallee *string `json:"MaskedCallee,omitempty" xml:"MaskedCallee,omitempty"`
	MediaType    *string `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
	// example:
	//
	// tags
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// example:
	//
	// 30
	TimeoutSeconds *int32 `json:"TimeoutSeconds,omitempty" xml:"TimeoutSeconds,omitempty"`
	// example:
	//
	// agent@ccc-test
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s MakeCallRequest) String() string {
	return dara.Prettify(s)
}

func (s MakeCallRequest) GoString() string {
	return s.String()
}

func (s *MakeCallRequest) GetCallee() *string {
	return s.Callee
}

func (s *MakeCallRequest) GetCaller() *string {
	return s.Caller
}

func (s *MakeCallRequest) GetDeviceId() *string {
	return s.DeviceId
}

func (s *MakeCallRequest) GetFlashSmsVariables() *string {
	return s.FlashSmsVariables
}

func (s *MakeCallRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *MakeCallRequest) GetMaskedCallee() *string {
	return s.MaskedCallee
}

func (s *MakeCallRequest) GetMediaType() *string {
	return s.MediaType
}

func (s *MakeCallRequest) GetTags() *string {
	return s.Tags
}

func (s *MakeCallRequest) GetTimeoutSeconds() *int32 {
	return s.TimeoutSeconds
}

func (s *MakeCallRequest) GetUserId() *string {
	return s.UserId
}

func (s *MakeCallRequest) SetCallee(v string) *MakeCallRequest {
	s.Callee = &v
	return s
}

func (s *MakeCallRequest) SetCaller(v string) *MakeCallRequest {
	s.Caller = &v
	return s
}

func (s *MakeCallRequest) SetDeviceId(v string) *MakeCallRequest {
	s.DeviceId = &v
	return s
}

func (s *MakeCallRequest) SetFlashSmsVariables(v string) *MakeCallRequest {
	s.FlashSmsVariables = &v
	return s
}

func (s *MakeCallRequest) SetInstanceId(v string) *MakeCallRequest {
	s.InstanceId = &v
	return s
}

func (s *MakeCallRequest) SetMaskedCallee(v string) *MakeCallRequest {
	s.MaskedCallee = &v
	return s
}

func (s *MakeCallRequest) SetMediaType(v string) *MakeCallRequest {
	s.MediaType = &v
	return s
}

func (s *MakeCallRequest) SetTags(v string) *MakeCallRequest {
	s.Tags = &v
	return s
}

func (s *MakeCallRequest) SetTimeoutSeconds(v int32) *MakeCallRequest {
	s.TimeoutSeconds = &v
	return s
}

func (s *MakeCallRequest) SetUserId(v string) *MakeCallRequest {
	s.UserId = &v
	return s
}

func (s *MakeCallRequest) Validate() error {
	return dara.Validate(s)
}
