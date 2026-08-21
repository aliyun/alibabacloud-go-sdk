// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateCopilotResponseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetXDebugId(v string) *GenerateCopilotResponseRequest
	GetXDebugId() *string
	SetLlmParamString(v string) *GenerateCopilotResponseRequest
	GetLlmParamString() *string
	SetXSysomInvokeSource(v string) *GenerateCopilotResponseRequest
	GetXSysomInvokeSource() *string
}

type GenerateCopilotResponseRequest struct {
	XDebugId *string `json:"X-Debug-Id,omitempty" xml:"X-Debug-Id,omitempty"`
	// Input parameters for the copilot LLM service. Refer to the standard LLM API input parameter dict, which needs to be converted to a string and passed into llmParamString
	//
	// example:
	//
	// "{\\"model\\":\\"Qwen-7B-Chat\\",\\"temperature\\":0.9,\\"max_tokens\\":1000,\\"top_p\\":1,\\"frequency_penalty\\":0.0,\\"presence_penalty\\":0.6,\\"messages\\":[{\\"role\\":\\"user\\",\\"content\\":\\"我是os工程师\\"},{\\"role\\":\\"assistant\\",\\"content\\":\\"您好。我是AI语言模型，很高兴为您服 务。有什么我能帮助您的呢\\"},{\\"role\\":\\"user\\",\\"content\\":\\"你知道什么是alinux吗\\"}]}"
	LlmParamString     *string `json:"llmParamString,omitempty" xml:"llmParamString,omitempty"`
	XSysomInvokeSource *string `json:"x-sysom-invoke-source,omitempty" xml:"x-sysom-invoke-source,omitempty"`
}

func (s GenerateCopilotResponseRequest) String() string {
	return dara.Prettify(s)
}

func (s GenerateCopilotResponseRequest) GoString() string {
	return s.String()
}

func (s *GenerateCopilotResponseRequest) GetXDebugId() *string {
	return s.XDebugId
}

func (s *GenerateCopilotResponseRequest) GetLlmParamString() *string {
	return s.LlmParamString
}

func (s *GenerateCopilotResponseRequest) GetXSysomInvokeSource() *string {
	return s.XSysomInvokeSource
}

func (s *GenerateCopilotResponseRequest) SetXDebugId(v string) *GenerateCopilotResponseRequest {
	s.XDebugId = &v
	return s
}

func (s *GenerateCopilotResponseRequest) SetLlmParamString(v string) *GenerateCopilotResponseRequest {
	s.LlmParamString = &v
	return s
}

func (s *GenerateCopilotResponseRequest) SetXSysomInvokeSource(v string) *GenerateCopilotResponseRequest {
	s.XSysomInvokeSource = &v
	return s
}

func (s *GenerateCopilotResponseRequest) Validate() error {
	return dara.Validate(s)
}
