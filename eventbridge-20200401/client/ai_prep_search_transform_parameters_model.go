// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAiPrepSearchTransformParameters interface {
	dara.Model
	String() string
	GoString() string
	SetInputField(v *AiTransformField) *AiPrepSearchTransformParameters
	GetInputField() *AiTransformField
	SetMaxChunkSize(v int32) *AiPrepSearchTransformParameters
	GetMaxChunkSize() *int32
	SetStepName(v string) *AiPrepSearchTransformParameters
	GetStepName() *string
}

type AiPrepSearchTransformParameters struct {
	// The input text field.
	InputField *AiTransformField `json:"InputField,omitempty" xml:"InputField,omitempty"`
	// The maximum number of chunks. Valid values: 1 to 1000. Default value: 20.
	//
	// example:
	//
	// 20
	MaxChunkSize *int32 `json:"MaxChunkSize,omitempty" xml:"MaxChunkSize,omitempty"`
	// The field name attached to the CloudEvent for output. Default value: transform0.
	//
	// example:
	//
	// chunks
	StepName *string `json:"StepName,omitempty" xml:"StepName,omitempty"`
}

func (s AiPrepSearchTransformParameters) String() string {
	return dara.Prettify(s)
}

func (s AiPrepSearchTransformParameters) GoString() string {
	return s.String()
}

func (s *AiPrepSearchTransformParameters) GetInputField() *AiTransformField {
	return s.InputField
}

func (s *AiPrepSearchTransformParameters) GetMaxChunkSize() *int32 {
	return s.MaxChunkSize
}

func (s *AiPrepSearchTransformParameters) GetStepName() *string {
	return s.StepName
}

func (s *AiPrepSearchTransformParameters) SetInputField(v *AiTransformField) *AiPrepSearchTransformParameters {
	s.InputField = v
	return s
}

func (s *AiPrepSearchTransformParameters) SetMaxChunkSize(v int32) *AiPrepSearchTransformParameters {
	s.MaxChunkSize = &v
	return s
}

func (s *AiPrepSearchTransformParameters) SetStepName(v string) *AiPrepSearchTransformParameters {
	s.StepName = &v
	return s
}

func (s *AiPrepSearchTransformParameters) Validate() error {
	if s.InputField != nil {
		if err := s.InputField.Validate(); err != nil {
			return err
		}
	}
	return nil
}
