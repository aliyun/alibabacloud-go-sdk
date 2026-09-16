// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAiSentimentAnalysisTransformParameters interface {
	dara.Model
	String() string
	GoString() string
	SetAspects(v []*string) *AiSentimentAnalysisTransformParameters
	GetAspects() []*string
	SetInputField(v *AiTransformField) *AiSentimentAnalysisTransformParameters
	GetInputField() *AiTransformField
	SetStepName(v string) *AiSentimentAnalysisTransformParameters
	GetStepName() *string
}

type AiSentimentAnalysisTransformParameters struct {
	// Performs emotion analysis on each specified aspect separately. If left empty, performs overall emotion analysis on the entire text.
	//
	// example:
	//
	// ["price","logistics","customer service"]
	Aspects []*string `json:"Aspects,omitempty" xml:"Aspects,omitempty" type:"Repeated"`
	// The input text field.
	InputField *AiTransformField `json:"InputField,omitempty" xml:"InputField,omitempty"`
	// The field name attached to the CloudEvent for output. Default value: transform0.
	//
	// example:
	//
	// sentiment
	StepName *string `json:"StepName,omitempty" xml:"StepName,omitempty"`
}

func (s AiSentimentAnalysisTransformParameters) String() string {
	return dara.Prettify(s)
}

func (s AiSentimentAnalysisTransformParameters) GoString() string {
	return s.String()
}

func (s *AiSentimentAnalysisTransformParameters) GetAspects() []*string {
	return s.Aspects
}

func (s *AiSentimentAnalysisTransformParameters) GetInputField() *AiTransformField {
	return s.InputField
}

func (s *AiSentimentAnalysisTransformParameters) GetStepName() *string {
	return s.StepName
}

func (s *AiSentimentAnalysisTransformParameters) SetAspects(v []*string) *AiSentimentAnalysisTransformParameters {
	s.Aspects = v
	return s
}

func (s *AiSentimentAnalysisTransformParameters) SetInputField(v *AiTransformField) *AiSentimentAnalysisTransformParameters {
	s.InputField = v
	return s
}

func (s *AiSentimentAnalysisTransformParameters) SetStepName(v string) *AiSentimentAnalysisTransformParameters {
	s.StepName = &v
	return s
}

func (s *AiSentimentAnalysisTransformParameters) Validate() error {
	if s.InputField != nil {
		if err := s.InputField.Validate(); err != nil {
			return err
		}
	}
	return nil
}
