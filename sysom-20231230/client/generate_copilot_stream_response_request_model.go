// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateCopilotStreamResponseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLlmParamString(v string) *GenerateCopilotStreamResponseRequest
	GetLlmParamString() *string
}

type GenerateCopilotStreamResponseRequest struct {
	LlmParamString *string `json:"llmParamString,omitempty" xml:"llmParamString,omitempty"`
}

func (s GenerateCopilotStreamResponseRequest) String() string {
	return dara.Prettify(s)
}

func (s GenerateCopilotStreamResponseRequest) GoString() string {
	return s.String()
}

func (s *GenerateCopilotStreamResponseRequest) GetLlmParamString() *string {
	return s.LlmParamString
}

func (s *GenerateCopilotStreamResponseRequest) SetLlmParamString(v string) *GenerateCopilotStreamResponseRequest {
	s.LlmParamString = &v
	return s
}

func (s *GenerateCopilotStreamResponseRequest) Validate() error {
	return dara.Validate(s)
}
