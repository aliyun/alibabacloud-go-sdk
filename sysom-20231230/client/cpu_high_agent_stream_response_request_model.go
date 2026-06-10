// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCpuHighAgentStreamResponseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLlmParamString(v string) *CpuHighAgentStreamResponseRequest
	GetLlmParamString() *string
}

type CpuHighAgentStreamResponseRequest struct {
	// Input parameter for interfacing with the high-CPU agent service. Refer to the standard LLM API input parameter dictionary, convert it into a string, and pass it in the `llmParamString` field.
	//
	// example:
	//
	// "llmParamString": "{\\"messages\\": [{\\"role\\": \\"user\\", \\"content\\": \\"用户12345的机器i-67890，最近2分钟CPU使用率高，请结合最近2分钟的火焰图信息，分析原因\\"}]}"
	LlmParamString *string `json:"llmParamString,omitempty" xml:"llmParamString,omitempty"`
}

func (s CpuHighAgentStreamResponseRequest) String() string {
	return dara.Prettify(s)
}

func (s CpuHighAgentStreamResponseRequest) GoString() string {
	return s.String()
}

func (s *CpuHighAgentStreamResponseRequest) GetLlmParamString() *string {
	return s.LlmParamString
}

func (s *CpuHighAgentStreamResponseRequest) SetLlmParamString(v string) *CpuHighAgentStreamResponseRequest {
	s.LlmParamString = &v
	return s
}

func (s *CpuHighAgentStreamResponseRequest) Validate() error {
	return dara.Validate(s)
}
