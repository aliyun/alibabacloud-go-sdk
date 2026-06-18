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
	// LLM application code. View it in the [Application Management](https://aiccs.console.aliyun.com/engine/llmApp) interface.
	//
	// This parameter is required.
	//
	// example:
	//
	// F32******
	ApplicationCode *string `json:"ApplicationCode,omitempty" xml:"ApplicationCode,omitempty"`
	// Business parameters. These parameters are passed to the customer model when invoking the customer model.
	//
	// example:
	//
	// biz_params = {
	//
	//   "user_prompt_params": "{\\"city\\":\\"北京\\"}"
	//
	// }
	BizParamShrink *string `json:"BizParam,omitempty" xml:"BizParam,omitempty"`
	// Called number that receives the intelligent outbound call.
	//
	// This parameter is required.
	//
	// example:
	//
	// 137****0000
	CalledNumber *string `json:"CalledNumber,omitempty" xml:"CalledNumber,omitempty"`
	// Caller number. This parameter is required and supports only numbers from the Chinese mainland. View available numbers in the Voice Service [Number Management](https://dyvmsnext.console.aliyun.com/number/list/normal) interface.
	//
	// example:
	//
	// 132****2054
	CallerNumber *string `json:"CallerNumber,omitempty" xml:"CallerNumber,omitempty"`
	// Customer-provided ingest endpoint encoding.
	//
	// >
	//
	// > - If you use your own line, contact Alibaba Cloud support to enable this feature.
	//
	// > - The line encoding is provided by Alibaba Cloud support. Do not set this parameter if you do not have one.
	//
	// example:
	//
	// SELF_xxxxx_A_NET
	CustomerLineCode *string `json:"CustomerLineCode,omitempty" xml:"CustomerLineCode,omitempty"`
	// The extension number of the X number, up to 5 digits.
	//
	// 	Notice: Fill this field only in AXN extension mode. If no extension number exists, do not fill it.
	//
	// example:
	//
	// 123
	Extension *string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// An ID reserved for the caller. This ID will be returned to the caller through a receipt message. Length: 1–15 bytes.
	//
	// example:
	//
	// 222356****
	OutId *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	// Prompt variable. Go to the [Application Management](https://aiccs.console.aliyun.com/engine/llmApp) interface and click Details to view the prompt variables you created.
	//
	// example:
	//
	// {
	//
	//   "style": "温柔"
	//
	// }
	PromptParamShrink *string `json:"PromptParam,omitempty" xml:"PromptParam,omitempty"`
	// Maximum call duration. The call is automatically disconnected after timeout. Unit: seconds.
	//
	// >
	//
	// >- Maximum value: 3600 s.
	//
	// >- Minimum value: 600 s.
	//
	// example:
	//
	// 1200
	SessionTimeout *int32 `json:"SessionTimeout,omitempty" xml:"SessionTimeout,omitempty"`
	// Start-word variables. Go to the [LLM Application Management](https://aiccs.console.aliyun.com/engine/llmApp) interface and click Details to view the start-word variables of your created LLM application.
	//
	// example:
	//
	// {
	//
	//   "name": "小明",
	//
	//   "address": "浙江省杭州市"
	//
	// }
	StartWordParamShrink *string `json:"StartWordParam,omitempty" xml:"StartWordParam,omitempty"`
	// Voice speed during TTS playback.
	//
	// >
	//
	// > - Value range: -200 to 200. Default value is 0.
	//
	// > - If this parameter is not set, the voice speed configured in the LLM application is used by default.
	//
	// example:
	//
	// 50
	TtsSpeed *int32 `json:"TtsSpeed,omitempty" xml:"TtsSpeed,omitempty"`
	// The voice code used for TTS playback.
	//
	// >
	//
	// > - If no value is set, the voice code configured in the LLM application is used by default.
	//
	// > - You can use the [ListAvailableTts](https://help.aliyun.com/document_detail/2926668.html) API to view all available voice codes.
	//
	// example:
	//
	// V65******
	TtsVoiceCode *string `json:"TtsVoiceCode,omitempty" xml:"TtsVoiceCode,omitempty"`
	// The volume for TTS playback.
	//
	// >
	//
	// > - Value range: 0–100. Default value is 0.
	//
	// > - If no value is set, the volume configured in the LLM application is used by default.
	//
	// example:
	//
	// 10
	TtsVolume *int32 `json:"TtsVolume,omitempty" xml:"TtsVolume,omitempty"`
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
