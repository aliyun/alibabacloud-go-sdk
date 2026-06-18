// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLlmSmartCallFullDuplexShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationCode(v string) *LlmSmartCallFullDuplexShrinkRequest
	GetApplicationCode() *string
	SetCalledNumber(v string) *LlmSmartCallFullDuplexShrinkRequest
	GetCalledNumber() *string
	SetCallerNumber(v string) *LlmSmartCallFullDuplexShrinkRequest
	GetCallerNumber() *string
	SetOutId(v string) *LlmSmartCallFullDuplexShrinkRequest
	GetOutId() *string
	SetSessionTimeout(v int32) *LlmSmartCallFullDuplexShrinkRequest
	GetSessionTimeout() *int32
	SetStartWordParamShrink(v string) *LlmSmartCallFullDuplexShrinkRequest
	GetStartWordParamShrink() *string
	SetTtsSpeed(v int32) *LlmSmartCallFullDuplexShrinkRequest
	GetTtsSpeed() *int32
	SetTtsVoiceCode(v string) *LlmSmartCallFullDuplexShrinkRequest
	GetTtsVoiceCode() *string
	SetTtsVolume(v int32) *LlmSmartCallFullDuplexShrinkRequest
	GetTtsVolume() *int32
}

type LlmSmartCallFullDuplexShrinkRequest struct {
	// **ApplicationCode*	-
	//
	// The application code.
	//
	// This parameter is required.
	//
	// example:
	//
	// app123
	ApplicationCode *string `json:"ApplicationCode,omitempty" xml:"ApplicationCode,omitempty"`
	// **CalledNumber*	-
	//
	// The phone number that receives the intelligent outbound call.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234567890
	CalledNumber *string `json:"CalledNumber,omitempty" xml:"CalledNumber,omitempty"`
	// **CallerNumber*	-
	//
	// The calling number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0987654321
	CallerNumber *string `json:"CallerNumber,omitempty" xml:"CallerNumber,omitempty"`
	// **OutId*	-
	//
	// The external ID. The value must be 1 to 15 bytes in length.
	//
	// example:
	//
	// out123
	OutId *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	// **SessionTimeout*	-
	//
	// The maximum call duration. The call is automatically ended when the specified duration is exceeded. Unit: seconds. Maximum value: 3600.
	//
	// example:
	//
	// 120
	SessionTimeout *int32 `json:"SessionTimeout,omitempty" xml:"SessionTimeout,omitempty"`
	// **StartWordParam*	-
	//
	// The opening greeting variable. Format: JSON. Specify the variable name and parameter, for example, `{"custom":"XXX"}`.
	//
	// example:
	//
	// {\\"tailnumber\\":\\"7898\\",\\"platformOrder\\":\\"\\",\\"signatory\\":\\"客户\\"}
	StartWordParamShrink *string `json:"StartWordParam,omitempty" xml:"StartWordParam,omitempty"`
	// The voice speed. Valid values: 0 to 100. If this parameter is not specified, the value configured in the application is used.
	//
	// example:
	//
	// 70
	TtsSpeed *int32 `json:"TtsSpeed,omitempty" xml:"TtsSpeed,omitempty"`
	// The voice code from the personalized voice interface. If this parameter is not specified, the value configured in the application is used.
	//
	// example:
	//
	// V745A7CED
	TtsVoiceCode *string `json:"TtsVoiceCode,omitempty" xml:"TtsVoiceCode,omitempty"`
	// The volume. Valid values: -200 to 200. If this parameter is not specified, the value configured in the application is used.
	//
	// example:
	//
	// 100
	TtsVolume *int32 `json:"TtsVolume,omitempty" xml:"TtsVolume,omitempty"`
}

func (s LlmSmartCallFullDuplexShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s LlmSmartCallFullDuplexShrinkRequest) GoString() string {
	return s.String()
}

func (s *LlmSmartCallFullDuplexShrinkRequest) GetApplicationCode() *string {
	return s.ApplicationCode
}

func (s *LlmSmartCallFullDuplexShrinkRequest) GetCalledNumber() *string {
	return s.CalledNumber
}

func (s *LlmSmartCallFullDuplexShrinkRequest) GetCallerNumber() *string {
	return s.CallerNumber
}

func (s *LlmSmartCallFullDuplexShrinkRequest) GetOutId() *string {
	return s.OutId
}

func (s *LlmSmartCallFullDuplexShrinkRequest) GetSessionTimeout() *int32 {
	return s.SessionTimeout
}

func (s *LlmSmartCallFullDuplexShrinkRequest) GetStartWordParamShrink() *string {
	return s.StartWordParamShrink
}

func (s *LlmSmartCallFullDuplexShrinkRequest) GetTtsSpeed() *int32 {
	return s.TtsSpeed
}

func (s *LlmSmartCallFullDuplexShrinkRequest) GetTtsVoiceCode() *string {
	return s.TtsVoiceCode
}

func (s *LlmSmartCallFullDuplexShrinkRequest) GetTtsVolume() *int32 {
	return s.TtsVolume
}

func (s *LlmSmartCallFullDuplexShrinkRequest) SetApplicationCode(v string) *LlmSmartCallFullDuplexShrinkRequest {
	s.ApplicationCode = &v
	return s
}

func (s *LlmSmartCallFullDuplexShrinkRequest) SetCalledNumber(v string) *LlmSmartCallFullDuplexShrinkRequest {
	s.CalledNumber = &v
	return s
}

func (s *LlmSmartCallFullDuplexShrinkRequest) SetCallerNumber(v string) *LlmSmartCallFullDuplexShrinkRequest {
	s.CallerNumber = &v
	return s
}

func (s *LlmSmartCallFullDuplexShrinkRequest) SetOutId(v string) *LlmSmartCallFullDuplexShrinkRequest {
	s.OutId = &v
	return s
}

func (s *LlmSmartCallFullDuplexShrinkRequest) SetSessionTimeout(v int32) *LlmSmartCallFullDuplexShrinkRequest {
	s.SessionTimeout = &v
	return s
}

func (s *LlmSmartCallFullDuplexShrinkRequest) SetStartWordParamShrink(v string) *LlmSmartCallFullDuplexShrinkRequest {
	s.StartWordParamShrink = &v
	return s
}

func (s *LlmSmartCallFullDuplexShrinkRequest) SetTtsSpeed(v int32) *LlmSmartCallFullDuplexShrinkRequest {
	s.TtsSpeed = &v
	return s
}

func (s *LlmSmartCallFullDuplexShrinkRequest) SetTtsVoiceCode(v string) *LlmSmartCallFullDuplexShrinkRequest {
	s.TtsVoiceCode = &v
	return s
}

func (s *LlmSmartCallFullDuplexShrinkRequest) SetTtsVolume(v int32) *LlmSmartCallFullDuplexShrinkRequest {
	s.TtsVolume = &v
	return s
}

func (s *LlmSmartCallFullDuplexShrinkRequest) Validate() error {
	return dara.Validate(s)
}
