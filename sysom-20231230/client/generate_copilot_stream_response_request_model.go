// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateCopilotStreamResponseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetXDebugId(v string) *GenerateCopilotStreamResponseRequest
	GetXDebugId() *string
	SetLlmParamString(v string) *GenerateCopilotStreamResponseRequest
	GetLlmParamString() *string
	SetXSysomInvokeSource(v string) *GenerateCopilotStreamResponseRequest
	GetXSysomInvokeSource() *string
}

type GenerateCopilotStreamResponseRequest struct {
	XDebugId *string `json:"X-Debug-Id,omitempty" xml:"X-Debug-Id,omitempty"`
	// The input parameter for the Copilot LLM service. Refer to the standard LLM interface input parameter dict. Convert it to a string and pass it to llmParamString.
	//
	// example:
	//
	// "{\\"model\\":\\"Qwen-7B-Chat\\",\\"temperature\\":0.9,\\"max_tokens\\":1000,\\"top_p\\":1,\\"frequency_penalty\\":0.0,\\"presence_penalty\\":0.6,\\"messages\\":[{\\"role\\":\\"user\\",\\"content\\":\\"I am an OS engineer\\"},{\\"role\\":\\"assistant\\",\\"content\\":\\"Hello. I am an AI language model, happy to help you. What can I do for you?\\"},{\\"role\\":\\"user\\",\\"content\\":\\"Do you know what Alinux is?\\"}]}"
	LlmParamString     *string `json:"llmParamString,omitempty" xml:"llmParamString,omitempty"`
	XSysomInvokeSource *string `json:"x-sysom-invoke-source,omitempty" xml:"x-sysom-invoke-source,omitempty"`
}

func (s GenerateCopilotStreamResponseRequest) String() string {
	return dara.Prettify(s)
}

func (s GenerateCopilotStreamResponseRequest) GoString() string {
	return s.String()
}

func (s *GenerateCopilotStreamResponseRequest) GetXDebugId() *string {
	return s.XDebugId
}

func (s *GenerateCopilotStreamResponseRequest) GetLlmParamString() *string {
	return s.LlmParamString
}

func (s *GenerateCopilotStreamResponseRequest) GetXSysomInvokeSource() *string {
	return s.XSysomInvokeSource
}

func (s *GenerateCopilotStreamResponseRequest) SetXDebugId(v string) *GenerateCopilotStreamResponseRequest {
	s.XDebugId = &v
	return s
}

func (s *GenerateCopilotStreamResponseRequest) SetLlmParamString(v string) *GenerateCopilotStreamResponseRequest {
	s.LlmParamString = &v
	return s
}

func (s *GenerateCopilotStreamResponseRequest) SetXSysomInvokeSource(v string) *GenerateCopilotStreamResponseRequest {
	s.XSysomInvokeSource = &v
	return s
}

func (s *GenerateCopilotStreamResponseRequest) Validate() error {
	return dara.Validate(s)
}
