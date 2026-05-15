// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLlmSmartCallRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationCode(v string) *LlmSmartCallRequest
	GetApplicationCode() *string
	SetBizParam(v map[string]interface{}) *LlmSmartCallRequest
	GetBizParam() map[string]interface{}
	SetCalledNumber(v string) *LlmSmartCallRequest
	GetCalledNumber() *string
	SetCallerNumber(v string) *LlmSmartCallRequest
	GetCallerNumber() *string
	SetCustomerLineCode(v string) *LlmSmartCallRequest
	GetCustomerLineCode() *string
	SetExtension(v string) *LlmSmartCallRequest
	GetExtension() *string
	SetOutId(v string) *LlmSmartCallRequest
	GetOutId() *string
	SetPromptParam(v map[string]interface{}) *LlmSmartCallRequest
	GetPromptParam() map[string]interface{}
	SetSessionTimeout(v int32) *LlmSmartCallRequest
	GetSessionTimeout() *int32
	SetStartWordParam(v map[string]interface{}) *LlmSmartCallRequest
	GetStartWordParam() map[string]interface{}
	SetTtsSpeed(v int32) *LlmSmartCallRequest
	GetTtsSpeed() *int32
	SetTtsVoiceCode(v string) *LlmSmartCallRequest
	GetTtsVoiceCode() *string
	SetTtsVolume(v int32) *LlmSmartCallRequest
	GetTtsVolume() *int32
}

type LlmSmartCallRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// F32XXX2CF9
	ApplicationCode *string                `json:"ApplicationCode,omitempty" xml:"ApplicationCode,omitempty"`
	BizParam        map[string]interface{} `json:"BizParam,omitempty" xml:"BizParam,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 137****0000
	CalledNumber *string `json:"CalledNumber,omitempty" xml:"CalledNumber,omitempty"`
	// example:
	//
	// 057188040000
	CallerNumber     *string `json:"CallerNumber,omitempty" xml:"CallerNumber,omitempty"`
	CustomerLineCode *string `json:"CustomerLineCode,omitempty" xml:"CustomerLineCode,omitempty"`
	Extension        *string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// example:
	//
	// 222356****
	OutId          *string                `json:"OutId,omitempty" xml:"OutId,omitempty"`
	PromptParam    map[string]interface{} `json:"PromptParam,omitempty" xml:"PromptParam,omitempty"`
	SessionTimeout *int32                 `json:"SessionTimeout,omitempty" xml:"SessionTimeout,omitempty"`
	StartWordParam map[string]interface{} `json:"StartWordParam,omitempty" xml:"StartWordParam,omitempty"`
	TtsSpeed       *int32                 `json:"TtsSpeed,omitempty" xml:"TtsSpeed,omitempty"`
	TtsVoiceCode   *string                `json:"TtsVoiceCode,omitempty" xml:"TtsVoiceCode,omitempty"`
	TtsVolume      *int32                 `json:"TtsVolume,omitempty" xml:"TtsVolume,omitempty"`
}

func (s LlmSmartCallRequest) String() string {
	return dara.Prettify(s)
}

func (s LlmSmartCallRequest) GoString() string {
	return s.String()
}

func (s *LlmSmartCallRequest) GetApplicationCode() *string {
	return s.ApplicationCode
}

func (s *LlmSmartCallRequest) GetBizParam() map[string]interface{} {
	return s.BizParam
}

func (s *LlmSmartCallRequest) GetCalledNumber() *string {
	return s.CalledNumber
}

func (s *LlmSmartCallRequest) GetCallerNumber() *string {
	return s.CallerNumber
}

func (s *LlmSmartCallRequest) GetCustomerLineCode() *string {
	return s.CustomerLineCode
}

func (s *LlmSmartCallRequest) GetExtension() *string {
	return s.Extension
}

func (s *LlmSmartCallRequest) GetOutId() *string {
	return s.OutId
}

func (s *LlmSmartCallRequest) GetPromptParam() map[string]interface{} {
	return s.PromptParam
}

func (s *LlmSmartCallRequest) GetSessionTimeout() *int32 {
	return s.SessionTimeout
}

func (s *LlmSmartCallRequest) GetStartWordParam() map[string]interface{} {
	return s.StartWordParam
}

func (s *LlmSmartCallRequest) GetTtsSpeed() *int32 {
	return s.TtsSpeed
}

func (s *LlmSmartCallRequest) GetTtsVoiceCode() *string {
	return s.TtsVoiceCode
}

func (s *LlmSmartCallRequest) GetTtsVolume() *int32 {
	return s.TtsVolume
}

func (s *LlmSmartCallRequest) SetApplicationCode(v string) *LlmSmartCallRequest {
	s.ApplicationCode = &v
	return s
}

func (s *LlmSmartCallRequest) SetBizParam(v map[string]interface{}) *LlmSmartCallRequest {
	s.BizParam = v
	return s
}

func (s *LlmSmartCallRequest) SetCalledNumber(v string) *LlmSmartCallRequest {
	s.CalledNumber = &v
	return s
}

func (s *LlmSmartCallRequest) SetCallerNumber(v string) *LlmSmartCallRequest {
	s.CallerNumber = &v
	return s
}

func (s *LlmSmartCallRequest) SetCustomerLineCode(v string) *LlmSmartCallRequest {
	s.CustomerLineCode = &v
	return s
}

func (s *LlmSmartCallRequest) SetExtension(v string) *LlmSmartCallRequest {
	s.Extension = &v
	return s
}

func (s *LlmSmartCallRequest) SetOutId(v string) *LlmSmartCallRequest {
	s.OutId = &v
	return s
}

func (s *LlmSmartCallRequest) SetPromptParam(v map[string]interface{}) *LlmSmartCallRequest {
	s.PromptParam = v
	return s
}

func (s *LlmSmartCallRequest) SetSessionTimeout(v int32) *LlmSmartCallRequest {
	s.SessionTimeout = &v
	return s
}

func (s *LlmSmartCallRequest) SetStartWordParam(v map[string]interface{}) *LlmSmartCallRequest {
	s.StartWordParam = v
	return s
}

func (s *LlmSmartCallRequest) SetTtsSpeed(v int32) *LlmSmartCallRequest {
	s.TtsSpeed = &v
	return s
}

func (s *LlmSmartCallRequest) SetTtsVoiceCode(v string) *LlmSmartCallRequest {
	s.TtsVoiceCode = &v
	return s
}

func (s *LlmSmartCallRequest) SetTtsVolume(v int32) *LlmSmartCallRequest {
	s.TtsVolume = &v
	return s
}

func (s *LlmSmartCallRequest) Validate() error {
	return dara.Validate(s)
}
