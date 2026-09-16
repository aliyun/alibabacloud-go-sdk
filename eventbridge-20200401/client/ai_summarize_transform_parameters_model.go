// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAiSummarizeTransformParameters interface {
	dara.Model
	String() string
	GoString() string
	SetInputField(v *AiTransformField) *AiSummarizeTransformParameters
	GetInputField() *AiTransformField
	SetMaxLength(v int32) *AiSummarizeTransformParameters
	GetMaxLength() *int32
	SetStepName(v string) *AiSummarizeTransformParameters
	GetStepName() *string
}

type AiSummarizeTransformParameters struct {
	// The input text field.
	InputField *AiTransformField `json:"InputField,omitempty" xml:"InputField,omitempty"`
	// The maximum summary length. The value must be a positive integer. Default value: 200.
	//
	// example:
	//
	// 200
	MaxLength *int32 `json:"MaxLength,omitempty" xml:"MaxLength,omitempty"`
	// The field name attached to the CloudEvent for output. Default value: transform0.
	//
	// example:
	//
	// summary
	StepName *string `json:"StepName,omitempty" xml:"StepName,omitempty"`
}

func (s AiSummarizeTransformParameters) String() string {
	return dara.Prettify(s)
}

func (s AiSummarizeTransformParameters) GoString() string {
	return s.String()
}

func (s *AiSummarizeTransformParameters) GetInputField() *AiTransformField {
	return s.InputField
}

func (s *AiSummarizeTransformParameters) GetMaxLength() *int32 {
	return s.MaxLength
}

func (s *AiSummarizeTransformParameters) GetStepName() *string {
	return s.StepName
}

func (s *AiSummarizeTransformParameters) SetInputField(v *AiTransformField) *AiSummarizeTransformParameters {
	s.InputField = v
	return s
}

func (s *AiSummarizeTransformParameters) SetMaxLength(v int32) *AiSummarizeTransformParameters {
	s.MaxLength = &v
	return s
}

func (s *AiSummarizeTransformParameters) SetStepName(v string) *AiSummarizeTransformParameters {
	s.StepName = &v
	return s
}

func (s *AiSummarizeTransformParameters) Validate() error {
	if s.InputField != nil {
		if err := s.InputField.Validate(); err != nil {
			return err
		}
	}
	return nil
}
