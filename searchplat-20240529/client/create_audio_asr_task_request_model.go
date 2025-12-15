// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAudioAsrTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInput(v *CreateAudioAsrTaskRequestInput) *CreateAudioAsrTaskRequest
	GetInput() *CreateAudioAsrTaskRequestInput
	SetOutput(v *CreateAudioAsrTaskRequestOutput) *CreateAudioAsrTaskRequest
	GetOutput() *CreateAudioAsrTaskRequestOutput
	SetParameters(v map[string]interface{}) *CreateAudioAsrTaskRequest
	GetParameters() map[string]interface{}
}

type CreateAudioAsrTaskRequest struct {
	Input      *CreateAudioAsrTaskRequestInput  `json:"input,omitempty" xml:"input,omitempty" type:"Struct"`
	Output     *CreateAudioAsrTaskRequestOutput `json:"output,omitempty" xml:"output,omitempty" type:"Struct"`
	Parameters map[string]interface{}           `json:"parameters,omitempty" xml:"parameters,omitempty"`
}

func (s CreateAudioAsrTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAudioAsrTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateAudioAsrTaskRequest) GetInput() *CreateAudioAsrTaskRequestInput {
	return s.Input
}

func (s *CreateAudioAsrTaskRequest) GetOutput() *CreateAudioAsrTaskRequestOutput {
	return s.Output
}

func (s *CreateAudioAsrTaskRequest) GetParameters() map[string]interface{} {
	return s.Parameters
}

func (s *CreateAudioAsrTaskRequest) SetInput(v *CreateAudioAsrTaskRequestInput) *CreateAudioAsrTaskRequest {
	s.Input = v
	return s
}

func (s *CreateAudioAsrTaskRequest) SetOutput(v *CreateAudioAsrTaskRequestOutput) *CreateAudioAsrTaskRequest {
	s.Output = v
	return s
}

func (s *CreateAudioAsrTaskRequest) SetParameters(v map[string]interface{}) *CreateAudioAsrTaskRequest {
	s.Parameters = v
	return s
}

func (s *CreateAudioAsrTaskRequest) Validate() error {
	if s.Input != nil {
		if err := s.Input.Validate(); err != nil {
			return err
		}
	}
	if s.Output != nil {
		if err := s.Output.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateAudioAsrTaskRequestInput struct {
	Content  *string `json:"content,omitempty" xml:"content,omitempty"`
	FileName *string `json:"file_name,omitempty" xml:"file_name,omitempty"`
	Oss      *string `json:"oss,omitempty" xml:"oss,omitempty"`
}

func (s CreateAudioAsrTaskRequestInput) String() string {
	return dara.Prettify(s)
}

func (s CreateAudioAsrTaskRequestInput) GoString() string {
	return s.String()
}

func (s *CreateAudioAsrTaskRequestInput) GetContent() *string {
	return s.Content
}

func (s *CreateAudioAsrTaskRequestInput) GetFileName() *string {
	return s.FileName
}

func (s *CreateAudioAsrTaskRequestInput) GetOss() *string {
	return s.Oss
}

func (s *CreateAudioAsrTaskRequestInput) SetContent(v string) *CreateAudioAsrTaskRequestInput {
	s.Content = &v
	return s
}

func (s *CreateAudioAsrTaskRequestInput) SetFileName(v string) *CreateAudioAsrTaskRequestInput {
	s.FileName = &v
	return s
}

func (s *CreateAudioAsrTaskRequestInput) SetOss(v string) *CreateAudioAsrTaskRequestInput {
	s.Oss = &v
	return s
}

func (s *CreateAudioAsrTaskRequestInput) Validate() error {
	return dara.Validate(s)
}

type CreateAudioAsrTaskRequestOutput struct {
	Oss  *string `json:"oss,omitempty" xml:"oss,omitempty"`
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateAudioAsrTaskRequestOutput) String() string {
	return dara.Prettify(s)
}

func (s CreateAudioAsrTaskRequestOutput) GoString() string {
	return s.String()
}

func (s *CreateAudioAsrTaskRequestOutput) GetOss() *string {
	return s.Oss
}

func (s *CreateAudioAsrTaskRequestOutput) GetType() *string {
	return s.Type
}

func (s *CreateAudioAsrTaskRequestOutput) SetOss(v string) *CreateAudioAsrTaskRequestOutput {
	s.Oss = &v
	return s
}

func (s *CreateAudioAsrTaskRequestOutput) SetType(v string) *CreateAudioAsrTaskRequestOutput {
	s.Type = &v
	return s
}

func (s *CreateAudioAsrTaskRequestOutput) Validate() error {
	return dara.Validate(s)
}
