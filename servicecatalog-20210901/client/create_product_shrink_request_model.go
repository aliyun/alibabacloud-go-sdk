// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateProductShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateProductShrinkRequest
	GetDescription() *string
	SetProductName(v string) *CreateProductShrinkRequest
	GetProductName() *string
	SetProductType(v string) *CreateProductShrinkRequest
	GetProductType() *string
	SetProductVersionParametersShrink(v string) *CreateProductShrinkRequest
	GetProductVersionParametersShrink() *string
	SetProviderName(v string) *CreateProductShrinkRequest
	GetProviderName() *string
	SetTemplateType(v string) *CreateProductShrinkRequest
	GetTemplateType() *string
}

type CreateProductShrinkRequest struct {
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
	ProductVersionParametersShrink *string `json:"ProductVersionParameters,omitempty" xml:"ProductVersionParameters,omitempty"`
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

func (s CreateProductShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateProductShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateProductShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateProductShrinkRequest) GetProductName() *string {
	return s.ProductName
}

func (s *CreateProductShrinkRequest) GetProductType() *string {
	return s.ProductType
}

func (s *CreateProductShrinkRequest) GetProductVersionParametersShrink() *string {
	return s.ProductVersionParametersShrink
}

func (s *CreateProductShrinkRequest) GetProviderName() *string {
	return s.ProviderName
}

func (s *CreateProductShrinkRequest) GetTemplateType() *string {
	return s.TemplateType
}

func (s *CreateProductShrinkRequest) SetDescription(v string) *CreateProductShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateProductShrinkRequest) SetProductName(v string) *CreateProductShrinkRequest {
	s.ProductName = &v
	return s
}

func (s *CreateProductShrinkRequest) SetProductType(v string) *CreateProductShrinkRequest {
	s.ProductType = &v
	return s
}

func (s *CreateProductShrinkRequest) SetProductVersionParametersShrink(v string) *CreateProductShrinkRequest {
	s.ProductVersionParametersShrink = &v
	return s
}

func (s *CreateProductShrinkRequest) SetProviderName(v string) *CreateProductShrinkRequest {
	s.ProviderName = &v
	return s
}

func (s *CreateProductShrinkRequest) SetTemplateType(v string) *CreateProductShrinkRequest {
	s.TemplateType = &v
	return s
}

func (s *CreateProductShrinkRequest) Validate() error {
	return dara.Validate(s)
}
