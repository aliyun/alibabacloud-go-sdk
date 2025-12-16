// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFunctionResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *UpdateFunctionResourceRequestData) *UpdateFunctionResourceRequest
	GetData() *UpdateFunctionResourceRequestData
	SetDescription(v string) *UpdateFunctionResourceRequest
	GetDescription() *string
}

type UpdateFunctionResourceRequest struct {
	// The resource data. The data structure varies with the resource type.
	Data *UpdateFunctionResourceRequestData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The description of the resource.
	//
	// example:
	//
	// updated description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s UpdateFunctionResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateFunctionResourceRequest) GoString() string {
	return s.String()
}

func (s *UpdateFunctionResourceRequest) GetData() *UpdateFunctionResourceRequestData {
	return s.Data
}

func (s *UpdateFunctionResourceRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateFunctionResourceRequest) SetData(v *UpdateFunctionResourceRequestData) *UpdateFunctionResourceRequest {
	s.Data = v
	return s
}

func (s *UpdateFunctionResourceRequest) SetDescription(v string) *UpdateFunctionResourceRequest {
	s.Description = &v
	return s
}

func (s *UpdateFunctionResourceRequest) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateFunctionResourceRequestData struct {
	// The content of the file that corresponds to a resource of the raw_file type.
	//
	// example:
	//
	// abc
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The feature generators that correspond to resources of the feature_generator type.
	Generators []*UpdateFunctionResourceRequestDataGenerators `json:"Generators,omitempty" xml:"Generators,omitempty" type:"Repeated"`
}

func (s UpdateFunctionResourceRequestData) String() string {
	return dara.Prettify(s)
}

func (s UpdateFunctionResourceRequestData) GoString() string {
	return s.String()
}

func (s *UpdateFunctionResourceRequestData) GetContent() *string {
	return s.Content
}

func (s *UpdateFunctionResourceRequestData) GetGenerators() []*UpdateFunctionResourceRequestDataGenerators {
	return s.Generators
}

func (s *UpdateFunctionResourceRequestData) SetContent(v string) *UpdateFunctionResourceRequestData {
	s.Content = &v
	return s
}

func (s *UpdateFunctionResourceRequestData) SetGenerators(v []*UpdateFunctionResourceRequestDataGenerators) *UpdateFunctionResourceRequestData {
	s.Generators = v
	return s
}

func (s *UpdateFunctionResourceRequestData) Validate() error {
	if s.Generators != nil {
		for _, item := range s.Generators {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateFunctionResourceRequestDataGenerators struct {
	// The type of the feature generator.
	//
	// example:
	//
	// combo
	Generator *string `json:"Generator,omitempty" xml:"Generator,omitempty"`
	// The input.
	Input *UpdateFunctionResourceRequestDataGeneratorsInput `json:"Input,omitempty" xml:"Input,omitempty" type:"Struct"`
	// The name of the output feature.
	//
	// example:
	//
	// feature1
	Output *string `json:"Output,omitempty" xml:"Output,omitempty"`
}

func (s UpdateFunctionResourceRequestDataGenerators) String() string {
	return dara.Prettify(s)
}

func (s UpdateFunctionResourceRequestDataGenerators) GoString() string {
	return s.String()
}

func (s *UpdateFunctionResourceRequestDataGenerators) GetGenerator() *string {
	return s.Generator
}

func (s *UpdateFunctionResourceRequestDataGenerators) GetInput() *UpdateFunctionResourceRequestDataGeneratorsInput {
	return s.Input
}

func (s *UpdateFunctionResourceRequestDataGenerators) GetOutput() *string {
	return s.Output
}

func (s *UpdateFunctionResourceRequestDataGenerators) SetGenerator(v string) *UpdateFunctionResourceRequestDataGenerators {
	s.Generator = &v
	return s
}

func (s *UpdateFunctionResourceRequestDataGenerators) SetInput(v *UpdateFunctionResourceRequestDataGeneratorsInput) *UpdateFunctionResourceRequestDataGenerators {
	s.Input = v
	return s
}

func (s *UpdateFunctionResourceRequestDataGenerators) SetOutput(v string) *UpdateFunctionResourceRequestDataGenerators {
	s.Output = &v
	return s
}

func (s *UpdateFunctionResourceRequestDataGenerators) Validate() error {
	if s.Input != nil {
		if err := s.Input.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateFunctionResourceRequestDataGeneratorsInput struct {
	// The input features.
	Features []*UpdateFunctionResourceRequestDataGeneratorsInputFeatures `json:"Features,omitempty" xml:"Features,omitempty" type:"Repeated"`
}

func (s UpdateFunctionResourceRequestDataGeneratorsInput) String() string {
	return dara.Prettify(s)
}

func (s UpdateFunctionResourceRequestDataGeneratorsInput) GoString() string {
	return s.String()
}

func (s *UpdateFunctionResourceRequestDataGeneratorsInput) GetFeatures() []*UpdateFunctionResourceRequestDataGeneratorsInputFeatures {
	return s.Features
}

func (s *UpdateFunctionResourceRequestDataGeneratorsInput) SetFeatures(v []*UpdateFunctionResourceRequestDataGeneratorsInputFeatures) *UpdateFunctionResourceRequestDataGeneratorsInput {
	s.Features = v
	return s
}

func (s *UpdateFunctionResourceRequestDataGeneratorsInput) Validate() error {
	if s.Features != nil {
		for _, item := range s.Features {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateFunctionResourceRequestDataGeneratorsInputFeatures struct {
	// The name of the feature.
	//
	// example:
	//
	// system_item_id
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The type of the feature.
	//
	// example:
	//
	// item
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateFunctionResourceRequestDataGeneratorsInputFeatures) String() string {
	return dara.Prettify(s)
}

func (s UpdateFunctionResourceRequestDataGeneratorsInputFeatures) GoString() string {
	return s.String()
}

func (s *UpdateFunctionResourceRequestDataGeneratorsInputFeatures) GetName() *string {
	return s.Name
}

func (s *UpdateFunctionResourceRequestDataGeneratorsInputFeatures) GetType() *string {
	return s.Type
}

func (s *UpdateFunctionResourceRequestDataGeneratorsInputFeatures) SetName(v string) *UpdateFunctionResourceRequestDataGeneratorsInputFeatures {
	s.Name = &v
	return s
}

func (s *UpdateFunctionResourceRequestDataGeneratorsInputFeatures) SetType(v string) *UpdateFunctionResourceRequestDataGeneratorsInputFeatures {
	s.Type = &v
	return s
}

func (s *UpdateFunctionResourceRequestDataGeneratorsInputFeatures) Validate() error {
	return dara.Validate(s)
}
