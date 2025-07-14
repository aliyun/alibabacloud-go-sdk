// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDataInterpretationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *DataInterpretationRequest
	GetData() *string
	SetModelCode(v string) *DataInterpretationRequest
	GetModelCode() *string
	SetPromptForceOverride(v bool) *DataInterpretationRequest
	GetPromptForceOverride() *bool
	SetUserPrompt(v string) *DataInterpretationRequest
	GetUserPrompt() *string
	SetUserQuestion(v string) *DataInterpretationRequest
	GetUserQuestion() *string
}

type DataInterpretationRequest struct {
	// This parameter is required.
	Data                *string `json:"Data,omitempty" xml:"Data,omitempty"`
	ModelCode           *string `json:"ModelCode,omitempty" xml:"ModelCode,omitempty"`
	PromptForceOverride *bool   `json:"PromptForceOverride,omitempty" xml:"PromptForceOverride,omitempty"`
	UserPrompt          *string `json:"UserPrompt,omitempty" xml:"UserPrompt,omitempty"`
	UserQuestion        *string `json:"UserQuestion,omitempty" xml:"UserQuestion,omitempty"`
}

func (s DataInterpretationRequest) String() string {
	return dara.Prettify(s)
}

func (s DataInterpretationRequest) GoString() string {
	return s.String()
}

func (s *DataInterpretationRequest) GetData() *string {
	return s.Data
}

func (s *DataInterpretationRequest) GetModelCode() *string {
	return s.ModelCode
}

func (s *DataInterpretationRequest) GetPromptForceOverride() *bool {
	return s.PromptForceOverride
}

func (s *DataInterpretationRequest) GetUserPrompt() *string {
	return s.UserPrompt
}

func (s *DataInterpretationRequest) GetUserQuestion() *string {
	return s.UserQuestion
}

func (s *DataInterpretationRequest) SetData(v string) *DataInterpretationRequest {
	s.Data = &v
	return s
}

func (s *DataInterpretationRequest) SetModelCode(v string) *DataInterpretationRequest {
	s.ModelCode = &v
	return s
}

func (s *DataInterpretationRequest) SetPromptForceOverride(v bool) *DataInterpretationRequest {
	s.PromptForceOverride = &v
	return s
}

func (s *DataInterpretationRequest) SetUserPrompt(v string) *DataInterpretationRequest {
	s.UserPrompt = &v
	return s
}

func (s *DataInterpretationRequest) SetUserQuestion(v string) *DataInterpretationRequest {
	s.UserQuestion = &v
	return s
}

func (s *DataInterpretationRequest) Validate() error {
	return dara.Validate(s)
}
