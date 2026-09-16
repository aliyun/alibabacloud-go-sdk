// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAiFilterTransformParameters interface {
	dara.Model
	String() string
	GoString() string
	SetCondition(v string) *AiFilterTransformParameters
	GetCondition() *string
	SetInputField(v *AiTransformField) *AiFilterTransformParameters
	GetInputField() *AiTransformField
	SetOnMismatch(v string) *AiFilterTransformParameters
	GetOnMismatch() *string
	SetStepName(v string) *AiFilterTransformParameters
	GetStepName() *string
}

type AiFilterTransformParameters struct {
	// The retention condition described in natural language. The model uses this condition to determine whether an event matches.
	//
	// example:
	//
	// Retain only content related to user complaints or refunds
	Condition *string `json:"Condition,omitempty" xml:"Condition,omitempty"`
	// The input text field.
	InputField *AiTransformField `json:"InputField,omitempty" xml:"InputField,omitempty"`
	// The behavior when a mismatch occurs. Valid values: discard (default): discards the event. forward: forwards the event as-is.
	//
	// example:
	//
	// discard
	OnMismatch *string `json:"OnMismatch,omitempty" xml:"OnMismatch,omitempty"`
	// The field name in the CloudEvent to which the output is attached. Default value: transform0.
	//
	// example:
	//
	// filter_result
	StepName *string `json:"StepName,omitempty" xml:"StepName,omitempty"`
}

func (s AiFilterTransformParameters) String() string {
	return dara.Prettify(s)
}

func (s AiFilterTransformParameters) GoString() string {
	return s.String()
}

func (s *AiFilterTransformParameters) GetCondition() *string {
	return s.Condition
}

func (s *AiFilterTransformParameters) GetInputField() *AiTransformField {
	return s.InputField
}

func (s *AiFilterTransformParameters) GetOnMismatch() *string {
	return s.OnMismatch
}

func (s *AiFilterTransformParameters) GetStepName() *string {
	return s.StepName
}

func (s *AiFilterTransformParameters) SetCondition(v string) *AiFilterTransformParameters {
	s.Condition = &v
	return s
}

func (s *AiFilterTransformParameters) SetInputField(v *AiTransformField) *AiFilterTransformParameters {
	s.InputField = v
	return s
}

func (s *AiFilterTransformParameters) SetOnMismatch(v string) *AiFilterTransformParameters {
	s.OnMismatch = &v
	return s
}

func (s *AiFilterTransformParameters) SetStepName(v string) *AiFilterTransformParameters {
	s.StepName = &v
	return s
}

func (s *AiFilterTransformParameters) Validate() error {
	if s.InputField != nil {
		if err := s.InputField.Validate(); err != nil {
			return err
		}
	}
	return nil
}
