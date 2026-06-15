// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFeatureViewRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFields(v []*UpdateFeatureViewRequestFields) *UpdateFeatureViewRequest
	GetFields() []*UpdateFeatureViewRequestFields
}

type UpdateFeatureViewRequest struct {
	// A list of fields.
	//
	// This parameter is required.
	Fields []*UpdateFeatureViewRequestFields `json:"Fields,omitempty" xml:"Fields,omitempty" type:"Repeated"`
}

func (s UpdateFeatureViewRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateFeatureViewRequest) GoString() string {
	return s.String()
}

func (s *UpdateFeatureViewRequest) GetFields() []*UpdateFeatureViewRequestFields {
	return s.Fields
}

func (s *UpdateFeatureViewRequest) SetFields(v []*UpdateFeatureViewRequestFields) *UpdateFeatureViewRequest {
	s.Fields = v
	return s
}

func (s *UpdateFeatureViewRequest) Validate() error {
	if s.Fields != nil {
		for _, item := range s.Fields {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateFeatureViewRequestFields struct {
	// A list of field attributes. Valid values:
	//
	// - `Partition`: Indicates that the field is a partition field.
	//
	// - `PrimaryKey`: Indicates that the field is a primary key.
	//
	// - `EventTime`: Indicates that the field is the event time.
	//
	// This parameter is required.
	Attributes []*string `json:"Attributes,omitempty" xml:"Attributes,omitempty" type:"Repeated"`
	// The name of the field.
	//
	// This parameter is required.
	//
	// example:
	//
	// user_id
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The feature generation configuration.
	Transform []*UpdateFeatureViewRequestFieldsTransform `json:"Transform,omitempty" xml:"Transform,omitempty" type:"Repeated"`
	// The data type of the field.
	//
	// This parameter is required.
	//
	// example:
	//
	// INT32
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateFeatureViewRequestFields) String() string {
	return dara.Prettify(s)
}

func (s UpdateFeatureViewRequestFields) GoString() string {
	return s.String()
}

func (s *UpdateFeatureViewRequestFields) GetAttributes() []*string {
	return s.Attributes
}

func (s *UpdateFeatureViewRequestFields) GetName() *string {
	return s.Name
}

func (s *UpdateFeatureViewRequestFields) GetTransform() []*UpdateFeatureViewRequestFieldsTransform {
	return s.Transform
}

func (s *UpdateFeatureViewRequestFields) GetType() *string {
	return s.Type
}

func (s *UpdateFeatureViewRequestFields) SetAttributes(v []*string) *UpdateFeatureViewRequestFields {
	s.Attributes = v
	return s
}

func (s *UpdateFeatureViewRequestFields) SetName(v string) *UpdateFeatureViewRequestFields {
	s.Name = &v
	return s
}

func (s *UpdateFeatureViewRequestFields) SetTransform(v []*UpdateFeatureViewRequestFieldsTransform) *UpdateFeatureViewRequestFields {
	s.Transform = v
	return s
}

func (s *UpdateFeatureViewRequestFields) SetType(v string) *UpdateFeatureViewRequestFields {
	s.Type = &v
	return s
}

func (s *UpdateFeatureViewRequestFields) Validate() error {
	if s.Transform != nil {
		for _, item := range s.Transform {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateFeatureViewRequestFieldsTransform struct {
	// The input fields.
	//
	// This parameter is required.
	Input []*UpdateFeatureViewRequestFieldsTransformInput `json:"Input,omitempty" xml:"Input,omitempty" type:"Repeated"`
	// The ID of the LLM configuration.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	LLMConfigId *int32 `json:"LLMConfigId,omitempty" xml:"LLMConfigId,omitempty"`
	// The type of feature generation.
	//
	// This parameter is required.
	//
	// example:
	//
	// LLMEmbedding
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateFeatureViewRequestFieldsTransform) String() string {
	return dara.Prettify(s)
}

func (s UpdateFeatureViewRequestFieldsTransform) GoString() string {
	return s.String()
}

func (s *UpdateFeatureViewRequestFieldsTransform) GetInput() []*UpdateFeatureViewRequestFieldsTransformInput {
	return s.Input
}

func (s *UpdateFeatureViewRequestFieldsTransform) GetLLMConfigId() *int32 {
	return s.LLMConfigId
}

func (s *UpdateFeatureViewRequestFieldsTransform) GetType() *string {
	return s.Type
}

func (s *UpdateFeatureViewRequestFieldsTransform) SetInput(v []*UpdateFeatureViewRequestFieldsTransformInput) *UpdateFeatureViewRequestFieldsTransform {
	s.Input = v
	return s
}

func (s *UpdateFeatureViewRequestFieldsTransform) SetLLMConfigId(v int32) *UpdateFeatureViewRequestFieldsTransform {
	s.LLMConfigId = &v
	return s
}

func (s *UpdateFeatureViewRequestFieldsTransform) SetType(v string) *UpdateFeatureViewRequestFieldsTransform {
	s.Type = &v
	return s
}

func (s *UpdateFeatureViewRequestFieldsTransform) Validate() error {
	if s.Input != nil {
		for _, item := range s.Input {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateFeatureViewRequestFieldsTransformInput struct {
	// The modality of the input, such as text or image.
	//
	// example:
	//
	// NONE
	Modality *string `json:"Modality,omitempty" xml:"Modality,omitempty"`
	// The name of the input field.
	//
	// This parameter is required.
	//
	// example:
	//
	// f1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The data type of the input field.
	//
	// This parameter is required.
	//
	// example:
	//
	// STRING
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateFeatureViewRequestFieldsTransformInput) String() string {
	return dara.Prettify(s)
}

func (s UpdateFeatureViewRequestFieldsTransformInput) GoString() string {
	return s.String()
}

func (s *UpdateFeatureViewRequestFieldsTransformInput) GetModality() *string {
	return s.Modality
}

func (s *UpdateFeatureViewRequestFieldsTransformInput) GetName() *string {
	return s.Name
}

func (s *UpdateFeatureViewRequestFieldsTransformInput) GetType() *string {
	return s.Type
}

func (s *UpdateFeatureViewRequestFieldsTransformInput) SetModality(v string) *UpdateFeatureViewRequestFieldsTransformInput {
	s.Modality = &v
	return s
}

func (s *UpdateFeatureViewRequestFieldsTransformInput) SetName(v string) *UpdateFeatureViewRequestFieldsTransformInput {
	s.Name = &v
	return s
}

func (s *UpdateFeatureViewRequestFieldsTransformInput) SetType(v string) *UpdateFeatureViewRequestFieldsTransformInput {
	s.Type = &v
	return s
}

func (s *UpdateFeatureViewRequestFieldsTransformInput) Validate() error {
	return dara.Validate(s)
}
