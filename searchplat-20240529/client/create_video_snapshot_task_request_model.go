// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVideoSnapshotTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInput(v *CreateVideoSnapshotTaskRequestInput) *CreateVideoSnapshotTaskRequest
	GetInput() *CreateVideoSnapshotTaskRequestInput
	SetOutput(v *CreateVideoSnapshotTaskRequestOutput) *CreateVideoSnapshotTaskRequest
	GetOutput() *CreateVideoSnapshotTaskRequestOutput
	SetParameters(v map[string]interface{}) *CreateVideoSnapshotTaskRequest
	GetParameters() map[string]interface{}
}

type CreateVideoSnapshotTaskRequest struct {
	Input      *CreateVideoSnapshotTaskRequestInput  `json:"input,omitempty" xml:"input,omitempty" type:"Struct"`
	Output     *CreateVideoSnapshotTaskRequestOutput `json:"output,omitempty" xml:"output,omitempty" type:"Struct"`
	Parameters map[string]interface{}                `json:"parameters,omitempty" xml:"parameters,omitempty"`
}

func (s CreateVideoSnapshotTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateVideoSnapshotTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateVideoSnapshotTaskRequest) GetInput() *CreateVideoSnapshotTaskRequestInput {
	return s.Input
}

func (s *CreateVideoSnapshotTaskRequest) GetOutput() *CreateVideoSnapshotTaskRequestOutput {
	return s.Output
}

func (s *CreateVideoSnapshotTaskRequest) GetParameters() map[string]interface{} {
	return s.Parameters
}

func (s *CreateVideoSnapshotTaskRequest) SetInput(v *CreateVideoSnapshotTaskRequestInput) *CreateVideoSnapshotTaskRequest {
	s.Input = v
	return s
}

func (s *CreateVideoSnapshotTaskRequest) SetOutput(v *CreateVideoSnapshotTaskRequestOutput) *CreateVideoSnapshotTaskRequest {
	s.Output = v
	return s
}

func (s *CreateVideoSnapshotTaskRequest) SetParameters(v map[string]interface{}) *CreateVideoSnapshotTaskRequest {
	s.Parameters = v
	return s
}

func (s *CreateVideoSnapshotTaskRequest) Validate() error {
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

type CreateVideoSnapshotTaskRequestInput struct {
	Content  *string `json:"content,omitempty" xml:"content,omitempty"`
	FileName *string `json:"file_name,omitempty" xml:"file_name,omitempty"`
	Oss      *string `json:"oss,omitempty" xml:"oss,omitempty"`
}

func (s CreateVideoSnapshotTaskRequestInput) String() string {
	return dara.Prettify(s)
}

func (s CreateVideoSnapshotTaskRequestInput) GoString() string {
	return s.String()
}

func (s *CreateVideoSnapshotTaskRequestInput) GetContent() *string {
	return s.Content
}

func (s *CreateVideoSnapshotTaskRequestInput) GetFileName() *string {
	return s.FileName
}

func (s *CreateVideoSnapshotTaskRequestInput) GetOss() *string {
	return s.Oss
}

func (s *CreateVideoSnapshotTaskRequestInput) SetContent(v string) *CreateVideoSnapshotTaskRequestInput {
	s.Content = &v
	return s
}

func (s *CreateVideoSnapshotTaskRequestInput) SetFileName(v string) *CreateVideoSnapshotTaskRequestInput {
	s.FileName = &v
	return s
}

func (s *CreateVideoSnapshotTaskRequestInput) SetOss(v string) *CreateVideoSnapshotTaskRequestInput {
	s.Oss = &v
	return s
}

func (s *CreateVideoSnapshotTaskRequestInput) Validate() error {
	return dara.Validate(s)
}

type CreateVideoSnapshotTaskRequestOutput struct {
	Oss  *string `json:"oss,omitempty" xml:"oss,omitempty"`
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateVideoSnapshotTaskRequestOutput) String() string {
	return dara.Prettify(s)
}

func (s CreateVideoSnapshotTaskRequestOutput) GoString() string {
	return s.String()
}

func (s *CreateVideoSnapshotTaskRequestOutput) GetOss() *string {
	return s.Oss
}

func (s *CreateVideoSnapshotTaskRequestOutput) GetType() *string {
	return s.Type
}

func (s *CreateVideoSnapshotTaskRequestOutput) SetOss(v string) *CreateVideoSnapshotTaskRequestOutput {
	s.Oss = &v
	return s
}

func (s *CreateVideoSnapshotTaskRequestOutput) SetType(v string) *CreateVideoSnapshotTaskRequestOutput {
	s.Type = &v
	return s
}

func (s *CreateVideoSnapshotTaskRequestOutput) Validate() error {
	return dara.Validate(s)
}
