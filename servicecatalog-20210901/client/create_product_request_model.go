// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateProductRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateProductRequest
	GetDescription() *string
	SetProductName(v string) *CreateProductRequest
	GetProductName() *string
	SetProductType(v string) *CreateProductRequest
	GetProductType() *string
	SetProductVersionParameters(v *CreateProductRequestProductVersionParameters) *CreateProductRequest
	GetProductVersionParameters() *CreateProductRequestProductVersionParameters
	SetProviderName(v string) *CreateProductRequest
	GetProviderName() *string
	SetTemplateType(v string) *CreateProductRequest
	GetTemplateType() *string
}

type CreateProductRequest struct {
	// The description of the product.
	//
	// The value must be 1 to 128 characters in length.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// The description of the product.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the product.
	//
	// The value must be 1 to 128 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// DEMO-Create an ECS instance
	ProductName *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	// The type of the product.
	//
	// Set the value to Ros, which specifies Resource Orchestration Service (ROS).
	//
	// This parameter is required.
	//
	// example:
	//
	// Ros
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The information about the product version.
	ProductVersionParameters *CreateProductRequestProductVersionParameters `json:"ProductVersionParameters,omitempty" xml:"ProductVersionParameters,omitempty" type:"Struct"`
	// The provider of the product.
	//
	// The value must be 1 to 128 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// IT team
	ProviderName *string `json:"ProviderName,omitempty" xml:"ProviderName,omitempty"`
	// The type of the product template. Valid values:
	//
	// 	- RosTerraformTemplate: the Terraform template that is supported by ROS.
	//
	// 	- RosStandardTemplate: the standard ROS template.
	//
	// example:
	//
	// RosTerraformTemplate
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s CreateProductRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateProductRequest) GoString() string {
	return s.String()
}

func (s *CreateProductRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateProductRequest) GetProductName() *string {
	return s.ProductName
}

func (s *CreateProductRequest) GetProductType() *string {
	return s.ProductType
}

func (s *CreateProductRequest) GetProductVersionParameters() *CreateProductRequestProductVersionParameters {
	return s.ProductVersionParameters
}

func (s *CreateProductRequest) GetProviderName() *string {
	return s.ProviderName
}

func (s *CreateProductRequest) GetTemplateType() *string {
	return s.TemplateType
}

func (s *CreateProductRequest) SetDescription(v string) *CreateProductRequest {
	s.Description = &v
	return s
}

func (s *CreateProductRequest) SetProductName(v string) *CreateProductRequest {
	s.ProductName = &v
	return s
}

func (s *CreateProductRequest) SetProductType(v string) *CreateProductRequest {
	s.ProductType = &v
	return s
}

func (s *CreateProductRequest) SetProductVersionParameters(v *CreateProductRequestProductVersionParameters) *CreateProductRequest {
	s.ProductVersionParameters = v
	return s
}

func (s *CreateProductRequest) SetProviderName(v string) *CreateProductRequest {
	s.ProviderName = &v
	return s
}

func (s *CreateProductRequest) SetTemplateType(v string) *CreateProductRequest {
	s.TemplateType = &v
	return s
}

func (s *CreateProductRequest) Validate() error {
	if s.ProductVersionParameters != nil {
		if err := s.ProductVersionParameters.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateProductRequestProductVersionParameters struct {
	// Specifies whether to enable the product version. Valid values:
	//
	// 	- true (default)
	//
	// 	- false
	//
	// example:
	//
	// true
	Active *bool `json:"Active,omitempty" xml:"Active,omitempty"`
	// The description of the product version.
	//
	// The value must be 1 to 128 characters in length.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// The description of the product version.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The recommended product version. Valid values:
	//
	// 	- Default (default): No product version is recommended.
	//
	// 	- Recommended: the stable version.
	//
	// 	- Latest: the latest version.
	//
	// 	- Deprecated: the version that is about to be deprecated.
	//
	// example:
	//
	// Default
	Guidance *string `json:"Guidance,omitempty" xml:"Guidance,omitempty"`
	// The name of the product version.
	//
	// The value must be 1 to 128 characters in length.
	//
	// example:
	//
	// 1.0
	ProductVersionName *string `json:"ProductVersionName,omitempty" xml:"ProductVersionName,omitempty"`
	// The type of the template.
	//
	// Set the value to RosTerraformTemplate, which indicates the Terraform template that is supported by Resource Orchestration Service (ROS).
	//
	// example:
	//
	// RosTerraformTemplate
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
	// The URL of the template.
	//
	// To obtain the URL of a template, you can call the [CreateTemplate](~~CreateTemplate~~) operation.
	TemplateUrl *string `json:"TemplateUrl,omitempty" xml:"TemplateUrl,omitempty"`
}

func (s CreateProductRequestProductVersionParameters) String() string {
	return dara.Prettify(s)
}

func (s CreateProductRequestProductVersionParameters) GoString() string {
	return s.String()
}

func (s *CreateProductRequestProductVersionParameters) GetActive() *bool {
	return s.Active
}

func (s *CreateProductRequestProductVersionParameters) GetDescription() *string {
	return s.Description
}

func (s *CreateProductRequestProductVersionParameters) GetGuidance() *string {
	return s.Guidance
}

func (s *CreateProductRequestProductVersionParameters) GetProductVersionName() *string {
	return s.ProductVersionName
}

func (s *CreateProductRequestProductVersionParameters) GetTemplateType() *string {
	return s.TemplateType
}

func (s *CreateProductRequestProductVersionParameters) GetTemplateUrl() *string {
	return s.TemplateUrl
}

func (s *CreateProductRequestProductVersionParameters) SetActive(v bool) *CreateProductRequestProductVersionParameters {
	s.Active = &v
	return s
}

func (s *CreateProductRequestProductVersionParameters) SetDescription(v string) *CreateProductRequestProductVersionParameters {
	s.Description = &v
	return s
}

func (s *CreateProductRequestProductVersionParameters) SetGuidance(v string) *CreateProductRequestProductVersionParameters {
	s.Guidance = &v
	return s
}

func (s *CreateProductRequestProductVersionParameters) SetProductVersionName(v string) *CreateProductRequestProductVersionParameters {
	s.ProductVersionName = &v
	return s
}

func (s *CreateProductRequestProductVersionParameters) SetTemplateType(v string) *CreateProductRequestProductVersionParameters {
	s.TemplateType = &v
	return s
}

func (s *CreateProductRequestProductVersionParameters) SetTemplateUrl(v string) *CreateProductRequestProductVersionParameters {
	s.TemplateUrl = &v
	return s
}

func (s *CreateProductRequestProductVersionParameters) Validate() error {
	return dara.Validate(s)
}
