// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAiClassifyTransformParameters interface {
	dara.Model
	String() string
	GoString() string
	SetInputField(v *AiTransformField) *AiClassifyTransformParameters
	GetInputField() *AiTransformField
	SetInstruction(v string) *AiClassifyTransformParameters
	GetInstruction() *string
	SetLabels(v []*string) *AiClassifyTransformParameters
	GetLabels() []*string
	SetOutputMode(v string) *AiClassifyTransformParameters
	GetOutputMode() *string
	SetStepName(v string) *AiClassifyTransformParameters
	GetStepName() *string
}

type AiClassifyTransformParameters struct {
	// The input text field.
	InputField *AiTransformField `json:"InputField,omitempty" xml:"InputField,omitempty"`
	// The classification constraints provided to the model, such as priority rules or how to categorize uncertain cases. If left empty, classification is performed based on Labels only.
	//
	// example:
	//
	// Classify as bug when crash or exception is mentioned, classify as other when uncertain
	Instruction *string `json:"Instruction,omitempty" xml:"Instruction,omitempty"`
	// The candidate classification labels. The classification result must fall within this list. Specify at least two labels.
	//
	// example:
	//
	// ["bug","feature","question","other"]
	Labels []*string `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// The output mode. Valid values: single: single-label. multi: multi-label. Default value: single.
	//
	// example:
	//
	// single
	OutputMode *string `json:"OutputMode,omitempty" xml:"OutputMode,omitempty"`
	// The field name in the CloudEvent to which the output is attached. Default value: transform0.
	//
	// example:
	//
	// classify_result
	StepName *string `json:"StepName,omitempty" xml:"StepName,omitempty"`
}

func (s AiClassifyTransformParameters) String() string {
	return dara.Prettify(s)
}

func (s AiClassifyTransformParameters) GoString() string {
	return s.String()
}

func (s *AiClassifyTransformParameters) GetInputField() *AiTransformField {
	return s.InputField
}

func (s *AiClassifyTransformParameters) GetInstruction() *string {
	return s.Instruction
}

func (s *AiClassifyTransformParameters) GetLabels() []*string {
	return s.Labels
}

func (s *AiClassifyTransformParameters) GetOutputMode() *string {
	return s.OutputMode
}

func (s *AiClassifyTransformParameters) GetStepName() *string {
	return s.StepName
}

func (s *AiClassifyTransformParameters) SetInputField(v *AiTransformField) *AiClassifyTransformParameters {
	s.InputField = v
	return s
}

func (s *AiClassifyTransformParameters) SetInstruction(v string) *AiClassifyTransformParameters {
	s.Instruction = &v
	return s
}

func (s *AiClassifyTransformParameters) SetLabels(v []*string) *AiClassifyTransformParameters {
	s.Labels = v
	return s
}

func (s *AiClassifyTransformParameters) SetOutputMode(v string) *AiClassifyTransformParameters {
	s.OutputMode = &v
	return s
}

func (s *AiClassifyTransformParameters) SetStepName(v string) *AiClassifyTransformParameters {
	s.StepName = &v
	return s
}

func (s *AiClassifyTransformParameters) Validate() error {
	if s.InputField != nil {
		if err := s.InputField.Validate(); err != nil {
			return err
		}
	}
	return nil
}
