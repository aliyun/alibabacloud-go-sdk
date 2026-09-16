// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAiRedactTransformParameters interface {
	dara.Model
	String() string
	GoString() string
	SetEntities(v []*string) *AiRedactTransformParameters
	GetEntities() []*string
	SetInputField(v *AiTransformField) *AiRedactTransformParameters
	GetInputField() *AiTransformField
	SetMaskChar(v string) *AiRedactTransformParameters
	GetMaskChar() *string
	SetMode(v string) *AiRedactTransformParameters
	GetMode() *string
	SetStepName(v string) *AiRedactTransformParameters
	GetStepName() *string
}

type AiRedactTransformParameters struct {
	// The entity types to identify and mask in the text, such as phone numbers, ID card numbers, and email addresses.
	//
	// example:
	//
	// ["PHONE","ID_CARD","EMAIL"]
	Entities []*string `json:"Entities,omitempty" xml:"Entities,omitempty" type:"Repeated"`
	// The input text field.
	InputField *AiTransformField `json:"InputField,omitempty" xml:"InputField,omitempty"`
	// The mask character used in mask mode. Default value: *.
	//
	// example:
	//
	// *
	MaskChar *string `json:"MaskChar,omitempty" xml:"MaskChar,omitempty"`
	// The masking mode. Valid values: mask, replace, and remove.
	//
	// example:
	//
	// mask
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The field name appended to the CloudEvent for output. Default value: transform0.
	//
	// example:
	//
	// redact_result
	StepName *string `json:"StepName,omitempty" xml:"StepName,omitempty"`
}

func (s AiRedactTransformParameters) String() string {
	return dara.Prettify(s)
}

func (s AiRedactTransformParameters) GoString() string {
	return s.String()
}

func (s *AiRedactTransformParameters) GetEntities() []*string {
	return s.Entities
}

func (s *AiRedactTransformParameters) GetInputField() *AiTransformField {
	return s.InputField
}

func (s *AiRedactTransformParameters) GetMaskChar() *string {
	return s.MaskChar
}

func (s *AiRedactTransformParameters) GetMode() *string {
	return s.Mode
}

func (s *AiRedactTransformParameters) GetStepName() *string {
	return s.StepName
}

func (s *AiRedactTransformParameters) SetEntities(v []*string) *AiRedactTransformParameters {
	s.Entities = v
	return s
}

func (s *AiRedactTransformParameters) SetInputField(v *AiTransformField) *AiRedactTransformParameters {
	s.InputField = v
	return s
}

func (s *AiRedactTransformParameters) SetMaskChar(v string) *AiRedactTransformParameters {
	s.MaskChar = &v
	return s
}

func (s *AiRedactTransformParameters) SetMode(v string) *AiRedactTransformParameters {
	s.Mode = &v
	return s
}

func (s *AiRedactTransformParameters) SetStepName(v string) *AiRedactTransformParameters {
	s.StepName = &v
	return s
}

func (s *AiRedactTransformParameters) Validate() error {
	if s.InputField != nil {
		if err := s.InputField.Validate(); err != nil {
			return err
		}
	}
	return nil
}
