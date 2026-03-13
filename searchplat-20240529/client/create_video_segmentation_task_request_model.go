// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVideoSegmentationTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInput(v *CreateVideoSegmentationTaskRequestInput) *CreateVideoSegmentationTaskRequest
	GetInput() *CreateVideoSegmentationTaskRequestInput
	SetOutput(v *CreateVideoSegmentationTaskRequestOutput) *CreateVideoSegmentationTaskRequest
	GetOutput() *CreateVideoSegmentationTaskRequestOutput
	SetParameters(v map[string]interface{}) *CreateVideoSegmentationTaskRequest
	GetParameters() map[string]interface{}
}

type CreateVideoSegmentationTaskRequest struct {
	Input      *CreateVideoSegmentationTaskRequestInput  `json:"input,omitempty" xml:"input,omitempty" type:"Struct"`
	Output     *CreateVideoSegmentationTaskRequestOutput `json:"output,omitempty" xml:"output,omitempty" type:"Struct"`
	Parameters map[string]interface{}                    `json:"parameters,omitempty" xml:"parameters,omitempty"`
}

func (s CreateVideoSegmentationTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateVideoSegmentationTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateVideoSegmentationTaskRequest) GetInput() *CreateVideoSegmentationTaskRequestInput {
	return s.Input
}

func (s *CreateVideoSegmentationTaskRequest) GetOutput() *CreateVideoSegmentationTaskRequestOutput {
	return s.Output
}

func (s *CreateVideoSegmentationTaskRequest) GetParameters() map[string]interface{} {
	return s.Parameters
}

func (s *CreateVideoSegmentationTaskRequest) SetInput(v *CreateVideoSegmentationTaskRequestInput) *CreateVideoSegmentationTaskRequest {
	s.Input = v
	return s
}

func (s *CreateVideoSegmentationTaskRequest) SetOutput(v *CreateVideoSegmentationTaskRequestOutput) *CreateVideoSegmentationTaskRequest {
	s.Output = v
	return s
}

func (s *CreateVideoSegmentationTaskRequest) SetParameters(v map[string]interface{}) *CreateVideoSegmentationTaskRequest {
	s.Parameters = v
	return s
}

func (s *CreateVideoSegmentationTaskRequest) Validate() error {
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

type CreateVideoSegmentationTaskRequestInput struct {
	FileName *string `json:"file_name,omitempty" xml:"file_name,omitempty"`
	Oss      *string `json:"oss,omitempty" xml:"oss,omitempty"`
	Url      *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s CreateVideoSegmentationTaskRequestInput) String() string {
	return dara.Prettify(s)
}

func (s CreateVideoSegmentationTaskRequestInput) GoString() string {
	return s.String()
}

func (s *CreateVideoSegmentationTaskRequestInput) GetFileName() *string {
	return s.FileName
}

func (s *CreateVideoSegmentationTaskRequestInput) GetOss() *string {
	return s.Oss
}

func (s *CreateVideoSegmentationTaskRequestInput) GetUrl() *string {
	return s.Url
}

func (s *CreateVideoSegmentationTaskRequestInput) SetFileName(v string) *CreateVideoSegmentationTaskRequestInput {
	s.FileName = &v
	return s
}

func (s *CreateVideoSegmentationTaskRequestInput) SetOss(v string) *CreateVideoSegmentationTaskRequestInput {
	s.Oss = &v
	return s
}

func (s *CreateVideoSegmentationTaskRequestInput) SetUrl(v string) *CreateVideoSegmentationTaskRequestInput {
	s.Url = &v
	return s
}

func (s *CreateVideoSegmentationTaskRequestInput) Validate() error {
	return dara.Validate(s)
}

type CreateVideoSegmentationTaskRequestOutput struct {
	Oss  *string `json:"oss,omitempty" xml:"oss,omitempty"`
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateVideoSegmentationTaskRequestOutput) String() string {
	return dara.Prettify(s)
}

func (s CreateVideoSegmentationTaskRequestOutput) GoString() string {
	return s.String()
}

func (s *CreateVideoSegmentationTaskRequestOutput) GetOss() *string {
	return s.Oss
}

func (s *CreateVideoSegmentationTaskRequestOutput) GetType() *string {
	return s.Type
}

func (s *CreateVideoSegmentationTaskRequestOutput) SetOss(v string) *CreateVideoSegmentationTaskRequestOutput {
	s.Oss = &v
	return s
}

func (s *CreateVideoSegmentationTaskRequestOutput) SetType(v string) *CreateVideoSegmentationTaskRequestOutput {
	s.Type = &v
	return s
}

func (s *CreateVideoSegmentationTaskRequestOutput) Validate() error {
	return dara.Validate(s)
}
