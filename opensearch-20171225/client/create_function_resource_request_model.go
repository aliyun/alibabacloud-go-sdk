// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFunctionResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *CreateFunctionResourceRequestData) *CreateFunctionResourceRequest
	GetData() *CreateFunctionResourceRequestData
	SetDescription(v string) *CreateFunctionResourceRequest
	GetDescription() *string
	SetResourceName(v string) *CreateFunctionResourceRequest
	GetResourceName() *string
	SetResourceType(v string) *CreateFunctionResourceRequest
	GetResourceType() *string
}

type CreateFunctionResourceRequest struct {
	// The resource data. The data structure varies with the resource type.
	Data *CreateFunctionResourceRequestData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The description of the resource.
	//
	// example:
	//
	// ""
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the resource.
	//
	// example:
	//
	// fg_jsoon
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// The resource type.
	//
	// Valid values:
	//
	// 	- feature_generator
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- raw_file
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// feature_generator
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s CreateFunctionResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateFunctionResourceRequest) GoString() string {
	return s.String()
}

func (s *CreateFunctionResourceRequest) GetData() *CreateFunctionResourceRequestData {
	return s.Data
}

func (s *CreateFunctionResourceRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateFunctionResourceRequest) GetResourceName() *string {
	return s.ResourceName
}

func (s *CreateFunctionResourceRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *CreateFunctionResourceRequest) SetData(v *CreateFunctionResourceRequestData) *CreateFunctionResourceRequest {
	s.Data = v
	return s
}

func (s *CreateFunctionResourceRequest) SetDescription(v string) *CreateFunctionResourceRequest {
	s.Description = &v
	return s
}

func (s *CreateFunctionResourceRequest) SetResourceName(v string) *CreateFunctionResourceRequest {
	s.ResourceName = &v
	return s
}

func (s *CreateFunctionResourceRequest) SetResourceType(v string) *CreateFunctionResourceRequest {
	s.ResourceType = &v
	return s
}

func (s *CreateFunctionResourceRequest) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateFunctionResourceRequestData struct {
	// The content of the file that corresponds to a resource of the raw_file type.
	//
	// example:
	//
	// "abc"
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The feature generators that correspond to resources of the feature_generator type.
	Generators []*CreateFunctionResourceRequestDataGenerators `json:"Generators,omitempty" xml:"Generators,omitempty" type:"Repeated"`
}

func (s CreateFunctionResourceRequestData) String() string {
	return dara.Prettify(s)
}

func (s CreateFunctionResourceRequestData) GoString() string {
	return s.String()
}

func (s *CreateFunctionResourceRequestData) GetContent() *string {
	return s.Content
}

func (s *CreateFunctionResourceRequestData) GetGenerators() []*CreateFunctionResourceRequestDataGenerators {
	return s.Generators
}

func (s *CreateFunctionResourceRequestData) SetContent(v string) *CreateFunctionResourceRequestData {
	s.Content = &v
	return s
}

func (s *CreateFunctionResourceRequestData) SetGenerators(v []*CreateFunctionResourceRequestDataGenerators) *CreateFunctionResourceRequestData {
	s.Generators = v
	return s
}

func (s *CreateFunctionResourceRequestData) Validate() error {
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

type CreateFunctionResourceRequestDataGenerators struct {
	// The type of the feature generator.
	//
	// Valid values:
	//
	// 	- lookup
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- sequence
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- overlap
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- raw
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- combo
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- id
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// id
	Generator *string `json:"Generator,omitempty" xml:"Generator,omitempty"`
	// The input.
	Input *CreateFunctionResourceRequestDataGeneratorsInput `json:"Input,omitempty" xml:"Input,omitempty" type:"Struct"`
	// The name of the output feature.
	//
	// example:
	//
	// item_id_feature
	Output *string `json:"Output,omitempty" xml:"Output,omitempty"`
}

func (s CreateFunctionResourceRequestDataGenerators) String() string {
	return dara.Prettify(s)
}

func (s CreateFunctionResourceRequestDataGenerators) GoString() string {
	return s.String()
}

func (s *CreateFunctionResourceRequestDataGenerators) GetGenerator() *string {
	return s.Generator
}

func (s *CreateFunctionResourceRequestDataGenerators) GetInput() *CreateFunctionResourceRequestDataGeneratorsInput {
	return s.Input
}

func (s *CreateFunctionResourceRequestDataGenerators) GetOutput() *string {
	return s.Output
}

func (s *CreateFunctionResourceRequestDataGenerators) SetGenerator(v string) *CreateFunctionResourceRequestDataGenerators {
	s.Generator = &v
	return s
}

func (s *CreateFunctionResourceRequestDataGenerators) SetInput(v *CreateFunctionResourceRequestDataGeneratorsInput) *CreateFunctionResourceRequestDataGenerators {
	s.Input = v
	return s
}

func (s *CreateFunctionResourceRequestDataGenerators) SetOutput(v string) *CreateFunctionResourceRequestDataGenerators {
	s.Output = &v
	return s
}

func (s *CreateFunctionResourceRequestDataGenerators) Validate() error {
	if s.Input != nil {
		if err := s.Input.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateFunctionResourceRequestDataGeneratorsInput struct {
	// The input features.
	Features []*CreateFunctionResourceRequestDataGeneratorsInputFeatures `json:"Features,omitempty" xml:"Features,omitempty" type:"Repeated"`
}

func (s CreateFunctionResourceRequestDataGeneratorsInput) String() string {
	return dara.Prettify(s)
}

func (s CreateFunctionResourceRequestDataGeneratorsInput) GoString() string {
	return s.String()
}

func (s *CreateFunctionResourceRequestDataGeneratorsInput) GetFeatures() []*CreateFunctionResourceRequestDataGeneratorsInputFeatures {
	return s.Features
}

func (s *CreateFunctionResourceRequestDataGeneratorsInput) SetFeatures(v []*CreateFunctionResourceRequestDataGeneratorsInputFeatures) *CreateFunctionResourceRequestDataGeneratorsInput {
	s.Features = v
	return s
}

func (s *CreateFunctionResourceRequestDataGeneratorsInput) Validate() error {
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

type CreateFunctionResourceRequestDataGeneratorsInputFeatures struct {
	// The name of the feature.
	//
	// example:
	//
	// system_item_id
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The type of the feature.
	//
	// Valid values:
	//
	// 	- item
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- user
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// item
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateFunctionResourceRequestDataGeneratorsInputFeatures) String() string {
	return dara.Prettify(s)
}

func (s CreateFunctionResourceRequestDataGeneratorsInputFeatures) GoString() string {
	return s.String()
}

func (s *CreateFunctionResourceRequestDataGeneratorsInputFeatures) GetName() *string {
	return s.Name
}

func (s *CreateFunctionResourceRequestDataGeneratorsInputFeatures) GetType() *string {
	return s.Type
}

func (s *CreateFunctionResourceRequestDataGeneratorsInputFeatures) SetName(v string) *CreateFunctionResourceRequestDataGeneratorsInputFeatures {
	s.Name = &v
	return s
}

func (s *CreateFunctionResourceRequestDataGeneratorsInputFeatures) SetType(v string) *CreateFunctionResourceRequestDataGeneratorsInputFeatures {
	s.Type = &v
	return s
}

func (s *CreateFunctionResourceRequestDataGeneratorsInputFeatures) Validate() error {
	return dara.Validate(s)
}
