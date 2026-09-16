// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAiExtractTransformParameters interface {
	dara.Model
	String() string
	GoString() string
	SetExtractSchema(v string) *AiExtractTransformParameters
	GetExtractSchema() *string
	SetInputField(v *AiTransformField) *AiExtractTransformParameters
	GetInputField() *AiTransformField
	SetStepName(v string) *AiExtractTransformParameters
	GetStepName() *string
}

type AiExtractTransformParameters struct {
	// The JSON Schema of the extraction results. Pass a serialized JSON object string or JSON Schema text.
	//
	// example:
	//
	// {"type":"object","properties":{"orderId":{"type":"string"},"amount":{"type":"number"}}}
	ExtractSchema *string `json:"ExtractSchema,omitempty" xml:"ExtractSchema,omitempty"`
	// The input text field.
	InputField *AiTransformField `json:"InputField,omitempty" xml:"InputField,omitempty"`
	// The field name attached to the CloudEvent for output. Default value: transform0.
	//
	// example:
	//
	// extract_result
	StepName *string `json:"StepName,omitempty" xml:"StepName,omitempty"`
}

func (s AiExtractTransformParameters) String() string {
	return dara.Prettify(s)
}

func (s AiExtractTransformParameters) GoString() string {
	return s.String()
}

func (s *AiExtractTransformParameters) GetExtractSchema() *string {
	return s.ExtractSchema
}

func (s *AiExtractTransformParameters) GetInputField() *AiTransformField {
	return s.InputField
}

func (s *AiExtractTransformParameters) GetStepName() *string {
	return s.StepName
}

func (s *AiExtractTransformParameters) SetExtractSchema(v string) *AiExtractTransformParameters {
	s.ExtractSchema = &v
	return s
}

func (s *AiExtractTransformParameters) SetInputField(v *AiTransformField) *AiExtractTransformParameters {
	s.InputField = v
	return s
}

func (s *AiExtractTransformParameters) SetStepName(v string) *AiExtractTransformParameters {
	s.StepName = &v
	return s
}

func (s *AiExtractTransformParameters) Validate() error {
	if s.InputField != nil {
		if err := s.InputField.Validate(); err != nil {
			return err
		}
	}
	return nil
}
