// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAiTranslateTransformParameters interface {
	dara.Model
	String() string
	GoString() string
	SetInputField(v *AiTransformField) *AiTranslateTransformParameters
	GetInputField() *AiTransformField
	SetSourceLanguage(v string) *AiTranslateTransformParameters
	GetSourceLanguage() *string
	SetStepName(v string) *AiTranslateTransformParameters
	GetStepName() *string
	SetTargetLanguage(v string) *AiTranslateTransformParameters
	GetTargetLanguage() *string
}

type AiTranslateTransformParameters struct {
	InputField *AiTransformField `json:"InputField,omitempty" xml:"InputField,omitempty"`
	// example:
	//
	// zh
	SourceLanguage *string `json:"SourceLanguage,omitempty" xml:"SourceLanguage,omitempty"`
	// example:
	//
	// translation
	StepName *string `json:"StepName,omitempty" xml:"StepName,omitempty"`
	// example:
	//
	// en
	TargetLanguage *string `json:"TargetLanguage,omitempty" xml:"TargetLanguage,omitempty"`
}

func (s AiTranslateTransformParameters) String() string {
	return dara.Prettify(s)
}

func (s AiTranslateTransformParameters) GoString() string {
	return s.String()
}

func (s *AiTranslateTransformParameters) GetInputField() *AiTransformField {
	return s.InputField
}

func (s *AiTranslateTransformParameters) GetSourceLanguage() *string {
	return s.SourceLanguage
}

func (s *AiTranslateTransformParameters) GetStepName() *string {
	return s.StepName
}

func (s *AiTranslateTransformParameters) GetTargetLanguage() *string {
	return s.TargetLanguage
}

func (s *AiTranslateTransformParameters) SetInputField(v *AiTransformField) *AiTranslateTransformParameters {
	s.InputField = v
	return s
}

func (s *AiTranslateTransformParameters) SetSourceLanguage(v string) *AiTranslateTransformParameters {
	s.SourceLanguage = &v
	return s
}

func (s *AiTranslateTransformParameters) SetStepName(v string) *AiTranslateTransformParameters {
	s.StepName = &v
	return s
}

func (s *AiTranslateTransformParameters) SetTargetLanguage(v string) *AiTranslateTransformParameters {
	s.TargetLanguage = &v
	return s
}

func (s *AiTranslateTransformParameters) Validate() error {
	if s.InputField != nil {
		if err := s.InputField.Validate(); err != nil {
			return err
		}
	}
	return nil
}
