// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateCopilotResponseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLlmParamString(v string) *GenerateCopilotResponseRequest
	GetLlmParamString() *string
}

type GenerateCopilotResponseRequest struct {
	LlmParamString *string `json:"llmParamString,omitempty" xml:"llmParamString,omitempty"`
}

func (s GenerateCopilotResponseRequest) String() string {
	return dara.Prettify(s)
}

func (s GenerateCopilotResponseRequest) GoString() string {
	return s.String()
}

func (s *GenerateCopilotResponseRequest) GetLlmParamString() *string {
	return s.LlmParamString
}

func (s *GenerateCopilotResponseRequest) SetLlmParamString(v string) *GenerateCopilotResponseRequest {
	s.LlmParamString = &v
	return s
}

func (s *GenerateCopilotResponseRequest) Validate() error {
	return dara.Validate(s)
}
