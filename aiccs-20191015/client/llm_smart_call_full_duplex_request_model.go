// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLlmSmartCallFullDuplexRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationCode(v string) *LlmSmartCallFullDuplexRequest
	GetApplicationCode() *string
	SetCalledNumber(v string) *LlmSmartCallFullDuplexRequest
	GetCalledNumber() *string
	SetCallerNumber(v string) *LlmSmartCallFullDuplexRequest
	GetCallerNumber() *string
	SetOutId(v string) *LlmSmartCallFullDuplexRequest
	GetOutId() *string
	SetSessionTimeout(v int32) *LlmSmartCallFullDuplexRequest
	GetSessionTimeout() *int32
	SetStartWordParam(v map[string]interface{}) *LlmSmartCallFullDuplexRequest
	GetStartWordParam() map[string]interface{}
	SetTtsSpeed(v int32) *LlmSmartCallFullDuplexRequest
	GetTtsSpeed() *int32
	SetTtsVoiceCode(v string) *LlmSmartCallFullDuplexRequest
	GetTtsVoiceCode() *string
	SetTtsVolume(v int32) *LlmSmartCallFullDuplexRequest
	GetTtsVolume() *int32
}

type LlmSmartCallFullDuplexRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// app123
	ApplicationCode *string `json:"ApplicationCode,omitempty" xml:"ApplicationCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1234567890
	CalledNumber *string `json:"CalledNumber,omitempty" xml:"CalledNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0987654321
	CallerNumber *string `json:"CallerNumber,omitempty" xml:"CallerNumber,omitempty"`
	// example:
	//
	// out123
	OutId *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	// example:
	//
	// 120
	SessionTimeout *int32                 `json:"SessionTimeout,omitempty" xml:"SessionTimeout,omitempty"`
	StartWordParam map[string]interface{} `json:"StartWordParam,omitempty" xml:"StartWordParam,omitempty"`
	// example:
	//
	// 70
	TtsSpeed *int32 `json:"TtsSpeed,omitempty" xml:"TtsSpeed,omitempty"`
	// example:
	//
	// V745A7CED
	TtsVoiceCode *string `json:"TtsVoiceCode,omitempty" xml:"TtsVoiceCode,omitempty"`
	// example:
	//
	// 100
	TtsVolume *int32 `json:"TtsVolume,omitempty" xml:"TtsVolume,omitempty"`
}

func (s LlmSmartCallFullDuplexRequest) String() string {
	return dara.Prettify(s)
}

func (s LlmSmartCallFullDuplexRequest) GoString() string {
	return s.String()
}

func (s *LlmSmartCallFullDuplexRequest) GetApplicationCode() *string {
	return s.ApplicationCode
}

func (s *LlmSmartCallFullDuplexRequest) GetCalledNumber() *string {
	return s.CalledNumber
}

func (s *LlmSmartCallFullDuplexRequest) GetCallerNumber() *string {
	return s.CallerNumber
}

func (s *LlmSmartCallFullDuplexRequest) GetOutId() *string {
	return s.OutId
}

func (s *LlmSmartCallFullDuplexRequest) GetSessionTimeout() *int32 {
	return s.SessionTimeout
}

func (s *LlmSmartCallFullDuplexRequest) GetStartWordParam() map[string]interface{} {
	return s.StartWordParam
}

func (s *LlmSmartCallFullDuplexRequest) GetTtsSpeed() *int32 {
	return s.TtsSpeed
}

func (s *LlmSmartCallFullDuplexRequest) GetTtsVoiceCode() *string {
	return s.TtsVoiceCode
}

func (s *LlmSmartCallFullDuplexRequest) GetTtsVolume() *int32 {
	return s.TtsVolume
}

func (s *LlmSmartCallFullDuplexRequest) SetApplicationCode(v string) *LlmSmartCallFullDuplexRequest {
	s.ApplicationCode = &v
	return s
}

func (s *LlmSmartCallFullDuplexRequest) SetCalledNumber(v string) *LlmSmartCallFullDuplexRequest {
	s.CalledNumber = &v
	return s
}

func (s *LlmSmartCallFullDuplexRequest) SetCallerNumber(v string) *LlmSmartCallFullDuplexRequest {
	s.CallerNumber = &v
	return s
}

func (s *LlmSmartCallFullDuplexRequest) SetOutId(v string) *LlmSmartCallFullDuplexRequest {
	s.OutId = &v
	return s
}

func (s *LlmSmartCallFullDuplexRequest) SetSessionTimeout(v int32) *LlmSmartCallFullDuplexRequest {
	s.SessionTimeout = &v
	return s
}

func (s *LlmSmartCallFullDuplexRequest) SetStartWordParam(v map[string]interface{}) *LlmSmartCallFullDuplexRequest {
	s.StartWordParam = v
	return s
}

func (s *LlmSmartCallFullDuplexRequest) SetTtsSpeed(v int32) *LlmSmartCallFullDuplexRequest {
	s.TtsSpeed = &v
	return s
}

func (s *LlmSmartCallFullDuplexRequest) SetTtsVoiceCode(v string) *LlmSmartCallFullDuplexRequest {
	s.TtsVoiceCode = &v
	return s
}

func (s *LlmSmartCallFullDuplexRequest) SetTtsVolume(v int32) *LlmSmartCallFullDuplexRequest {
	s.TtsVolume = &v
	return s
}

func (s *LlmSmartCallFullDuplexRequest) Validate() error {
	return dara.Validate(s)
}
