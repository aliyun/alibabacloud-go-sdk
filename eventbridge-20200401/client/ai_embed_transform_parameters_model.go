// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAiEmbedTransformParameters interface {
	dara.Model
	String() string
	GoString() string
	SetDimension(v int32) *AiEmbedTransformParameters
	GetDimension() *int32
	SetInputField(v *AiTransformField) *AiEmbedTransformParameters
	GetInputField() *AiTransformField
	SetModel(v string) *AiEmbedTransformParameters
	GetModel() *string
	SetStepName(v string) *AiEmbedTransformParameters
	GetStepName() *string
}

type AiEmbedTransformParameters struct {
	// The vector dimensions. Must be a dimension supported by the selected model. If not specified, the default value of the model is used (1024 for most models, 1536 for v1/v2/async).
	//
	// example:
	//
	// 1024
	Dimension *int32 `json:"Dimension,omitempty" xml:"Dimension,omitempty"`
	// The input text field.
	InputField *AiTransformField `json:"InputField,omitempty" xml:"InputField,omitempty"`
	// The embedding model. Default value: text-embedding-v4.
	//
	// example:
	//
	// text-embedding-v4
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// The field name in the CloudEvent to which the output is attached. Default value: transform0.
	//
	// example:
	//
	// embedding
	StepName *string `json:"StepName,omitempty" xml:"StepName,omitempty"`
}

func (s AiEmbedTransformParameters) String() string {
	return dara.Prettify(s)
}

func (s AiEmbedTransformParameters) GoString() string {
	return s.String()
}

func (s *AiEmbedTransformParameters) GetDimension() *int32 {
	return s.Dimension
}

func (s *AiEmbedTransformParameters) GetInputField() *AiTransformField {
	return s.InputField
}

func (s *AiEmbedTransformParameters) GetModel() *string {
	return s.Model
}

func (s *AiEmbedTransformParameters) GetStepName() *string {
	return s.StepName
}

func (s *AiEmbedTransformParameters) SetDimension(v int32) *AiEmbedTransformParameters {
	s.Dimension = &v
	return s
}

func (s *AiEmbedTransformParameters) SetInputField(v *AiTransformField) *AiEmbedTransformParameters {
	s.InputField = v
	return s
}

func (s *AiEmbedTransformParameters) SetModel(v string) *AiEmbedTransformParameters {
	s.Model = &v
	return s
}

func (s *AiEmbedTransformParameters) SetStepName(v string) *AiEmbedTransformParameters {
	s.StepName = &v
	return s
}

func (s *AiEmbedTransformParameters) Validate() error {
	if s.InputField != nil {
		if err := s.InputField.Validate(); err != nil {
			return err
		}
	}
	return nil
}
