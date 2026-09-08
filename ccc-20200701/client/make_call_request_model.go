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
	// Callee number. For internal calls, specify the target agent\\"s extension number in this field. For outbound calls, specify the customer\\"s phone number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1318888****
	Callee *string `json:"Callee,omitempty" xml:"Callee,omitempty"`
	// Caller number. This parameter is invalid for internal calls. For outbound calls, specify an outbound number available to the current agent. Ensure that the number supports outbound calling and that the agent has permission to use it. Permission can be granted in two ways: either by attaching the number to the skill group the agent signed into, or by setting the number as the agent\\"s personal outbound number.
	//
	// example:
	//
	// 010989****
	Caller *string `json:"Caller,omitempty" xml:"Caller,omitempty"`
	// Device ID. This field is meaningless and can be filled with any value.
	//
	// example:
	//
	// device
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// Flash SMS configuration
	//
	// example:
	//
	// {\\"applicationId\\":\\"6bd18325-ea7f-4881-8902-4d06283d3b3b\\",\\"templateId\\":\\"1722217249064\\"}
	FlashSmsVariables *string `json:"FlashSmsVariables,omitempty" xml:"FlashSmsVariables,omitempty"`
	// Instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The desensitized callee number. If this field is not empty, it indicates that the callee number must be desensitized. The desensitization rule is defined by the customer. You only need to enter the desensitized callee number here. Using a desensitized callee number means that in certain scenarios, you will see the desensitized callee number and cannot view the real callee number.
	//
	// example:
	//
	// 131****8888
	MaskedCallee *string `json:"MaskedCallee,omitempty" xml:"MaskedCallee,omitempty"`
	// Media type. The default value is AUDIO. Other valid values include VIDEO.
	//
	// example:
	//
	// AUDIO
	MediaType *string `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
	// Ingest endpoint data. The customer does not need to concern themselves with this.
	//
	// example:
	//
	// tags
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// Timeout. If the call is not answered within the time specified by this parameter, the system automatically hangs up. Valid values range from 30 to 300 seconds.
	//
	// example:
	//
	// 30
	TimeoutSeconds *int32 `json:"TimeoutSeconds,omitempty" xml:"TimeoutSeconds,omitempty"`
	// Agent ID initiating the outbound call. This field is optional. If not specified, the system uses the agent mapped to the current RAM user by default.
	//
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
