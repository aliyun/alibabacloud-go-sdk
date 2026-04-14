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
