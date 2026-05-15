// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLlmSmartCallShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationCode(v string) *LlmSmartCallShrinkRequest
	GetApplicationCode() *string
	SetBizParamShrink(v string) *LlmSmartCallShrinkRequest
	GetBizParamShrink() *string
	SetCalledNumber(v string) *LlmSmartCallShrinkRequest
	GetCalledNumber() *string
	SetCallerNumber(v string) *LlmSmartCallShrinkRequest
	GetCallerNumber() *string
	SetCustomerLineCode(v string) *LlmSmartCallShrinkRequest
	GetCustomerLineCode() *string
	SetExtension(v string) *LlmSmartCallShrinkRequest
	GetExtension() *string
	SetOutId(v string) *LlmSmartCallShrinkRequest
	GetOutId() *string
	SetPromptParamShrink(v string) *LlmSmartCallShrinkRequest
	GetPromptParamShrink() *string
	SetSessionTimeout(v int32) *LlmSmartCallShrinkRequest
	GetSessionTimeout() *int32
	SetStartWordParamShrink(v string) *LlmSmartCallShrinkRequest
	GetStartWordParamShrink() *string
	SetTtsSpeed(v int32) *LlmSmartCallShrinkRequest
	GetTtsSpeed() *int32
	SetTtsVoiceCode(v string) *LlmSmartCallShrinkRequest
	GetTtsVoiceCode() *string
	SetTtsVolume(v int32) *LlmSmartCallShrinkRequest
	GetTtsVolume() *int32
}

type LlmSmartCallShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// F32XXX2CF9
	ApplicationCode *string `json:"ApplicationCode,omitempty" xml:"ApplicationCode,omitempty"`
	BizParamShrink  *string `json:"BizParam,omitempty" xml:"BizParam,omitempty"`
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
	OutId                *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	PromptParamShrink    *string `json:"PromptParam,omitempty" xml:"PromptParam,omitempty"`
	SessionTimeout       *int32  `json:"SessionTimeout,omitempty" xml:"SessionTimeout,omitempty"`
	StartWordParamShrink *string `json:"StartWordParam,omitempty" xml:"StartWordParam,omitempty"`
	TtsSpeed             *int32  `json:"TtsSpeed,omitempty" xml:"TtsSpeed,omitempty"`
	TtsVoiceCode         *string `json:"TtsVoiceCode,omitempty" xml:"TtsVoiceCode,omitempty"`
	TtsVolume            *int32  `json:"TtsVolume,omitempty" xml:"TtsVolume,omitempty"`
}

func (s LlmSmartCallShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s LlmSmartCallShrinkRequest) GoString() string {
	return s.String()
}

func (s *LlmSmartCallShrinkRequest) GetApplicationCode() *string {
	return s.ApplicationCode
}

func (s *LlmSmartCallShrinkRequest) GetBizParamShrink() *string {
	return s.BizParamShrink
}

func (s *LlmSmartCallShrinkRequest) GetCalledNumber() *string {
	return s.CalledNumber
}

func (s *LlmSmartCallShrinkRequest) GetCallerNumber() *string {
	return s.CallerNumber
}

func (s *LlmSmartCallShrinkRequest) GetCustomerLineCode() *string {
	return s.CustomerLineCode
}

func (s *LlmSmartCallShrinkRequest) GetExtension() *string {
	return s.Extension
}

func (s *LlmSmartCallShrinkRequest) GetOutId() *string {
	return s.OutId
}

func (s *LlmSmartCallShrinkRequest) GetPromptParamShrink() *string {
	return s.PromptParamShrink
}

func (s *LlmSmartCallShrinkRequest) GetSessionTimeout() *int32 {
	return s.SessionTimeout
}

func (s *LlmSmartCallShrinkRequest) GetStartWordParamShrink() *string {
	return s.StartWordParamShrink
}

func (s *LlmSmartCallShrinkRequest) GetTtsSpeed() *int32 {
	return s.TtsSpeed
}

func (s *LlmSmartCallShrinkRequest) GetTtsVoiceCode() *string {
	return s.TtsVoiceCode
}

func (s *LlmSmartCallShrinkRequest) GetTtsVolume() *int32 {
	return s.TtsVolume
}

func (s *LlmSmartCallShrinkRequest) SetApplicationCode(v string) *LlmSmartCallShrinkRequest {
	s.ApplicationCode = &v
	return s
}

func (s *LlmSmartCallShrinkRequest) SetBizParamShrink(v string) *LlmSmartCallShrinkRequest {
	s.BizParamShrink = &v
	return s
}

func (s *LlmSmartCallShrinkRequest) SetCalledNumber(v string) *LlmSmartCallShrinkRequest {
	s.CalledNumber = &v
	return s
}

func (s *LlmSmartCallShrinkRequest) SetCallerNumber(v string) *LlmSmartCallShrinkRequest {
	s.CallerNumber = &v
	return s
}

func (s *LlmSmartCallShrinkRequest) SetCustomerLineCode(v string) *LlmSmartCallShrinkRequest {
	s.CustomerLineCode = &v
	return s
}

func (s *LlmSmartCallShrinkRequest) SetExtension(v string) *LlmSmartCallShrinkRequest {
	s.Extension = &v
	return s
}

func (s *LlmSmartCallShrinkRequest) SetOutId(v string) *LlmSmartCallShrinkRequest {
	s.OutId = &v
	return s
}

func (s *LlmSmartCallShrinkRequest) SetPromptParamShrink(v string) *LlmSmartCallShrinkRequest {
	s.PromptParamShrink = &v
	return s
}

func (s *LlmSmartCallShrinkRequest) SetSessionTimeout(v int32) *LlmSmartCallShrinkRequest {
	s.SessionTimeout = &v
	return s
}

func (s *LlmSmartCallShrinkRequest) SetStartWordParamShrink(v string) *LlmSmartCallShrinkRequest {
	s.StartWordParamShrink = &v
	return s
}

func (s *LlmSmartCallShrinkRequest) SetTtsSpeed(v int32) *LlmSmartCallShrinkRequest {
	s.TtsSpeed = &v
	return s
}

func (s *LlmSmartCallShrinkRequest) SetTtsVoiceCode(v string) *LlmSmartCallShrinkRequest {
	s.TtsVoiceCode = &v
	return s
}

func (s *LlmSmartCallShrinkRequest) SetTtsVolume(v int32) *LlmSmartCallShrinkRequest {
	s.TtsVolume = &v
	return s
}

func (s *LlmSmartCallShrinkRequest) Validate() error {
	return dara.Validate(s)
}
