// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateProductVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActive(v bool) *CreateProductVersionRequest
	GetActive() *bool
	SetDescription(v string) *CreateProductVersionRequest
	GetDescription() *string
	SetGuidance(v string) *CreateProductVersionRequest
	GetGuidance() *string
	SetProductId(v string) *CreateProductVersionRequest
	GetProductId() *string
	SetProductVersionName(v string) *CreateProductVersionRequest
	GetProductVersionName() *string
	SetTemplateType(v string) *CreateProductVersionRequest
	GetTemplateType() *string
	SetTemplateUrl(v string) *CreateProductVersionRequest
	GetTemplateUrl() *string
}

type CreateProductVersionRequest struct {
	// Specifies whether the product version is active. Valid values:
	//
	// 	- true: The product version is active. This is the default value.
	//
	// 	- false: The product version is inactive.
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
	// The recommendation information. Valid values:
	//
	// 	- Default: No recommendation information is provided. This is the default value.
	//
	// 	- Recommended: the recommendation version.
	//
	// 	- Latest: the latest version.
	//
	// 	- Deprecated: the version that is about to be discontinued.
	//
	// example:
	//
	// Default
	Guidance *string `json:"Guidance,omitempty" xml:"Guidance,omitempty"`
	// The ID of the product to which the product version belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// prod-bp18r7q127****
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// The name of the product version.
	//
	// The value must be 1 to 128 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1.0
	ProductVersionName *string `json:"ProductVersionName,omitempty" xml:"ProductVersionName,omitempty"`
	// The type of the product template. Valid values:
	//
	// 	- RosTerraformTemplate: the Terraform template that is supported by Resource Orchestration Service (ROS).
	//
	// 	- RosStandardTemplate: the standard ROS template.
	//
	// This parameter is required.
	//
	// example:
	//
	// RosTerraformTemplate
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
	// The URL of the template.
	//
	// For more information about how to obtain the URL of a template, see [CreateTemplate](~~CreateTemplate~~).
	//
	// This parameter is required.
	TemplateUrl *string `json:"TemplateUrl,omitempty" xml:"TemplateUrl,omitempty"`
}

func (s CreateProductVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateProductVersionRequest) GoString() string {
	return s.String()
}

func (s *CreateProductVersionRequest) GetActive() *bool {
	return s.Active
}

func (s *CreateProductVersionRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateProductVersionRequest) GetGuidance() *string {
	return s.Guidance
}

func (s *CreateProductVersionRequest) GetProductId() *string {
	return s.ProductId
}

func (s *CreateProductVersionRequest) GetProductVersionName() *string {
	return s.ProductVersionName
}

func (s *CreateProductVersionRequest) GetTemplateType() *string {
	return s.TemplateType
}

func (s *CreateProductVersionRequest) GetTemplateUrl() *string {
	return s.TemplateUrl
}

func (s *CreateProductVersionRequest) SetActive(v bool) *CreateProductVersionRequest {
	s.Active = &v
	return s
}

func (s *CreateProductVersionRequest) SetDescription(v string) *CreateProductVersionRequest {
	s.Description = &v
	return s
}

func (s *CreateProductVersionRequest) SetGuidance(v string) *CreateProductVersionRequest {
	s.Guidance = &v
	return s
}

func (s *CreateProductVersionRequest) SetProductId(v string) *CreateProductVersionRequest {
	s.ProductId = &v
	return s
}

func (s *CreateProductVersionRequest) SetProductVersionName(v string) *CreateProductVersionRequest {
	s.ProductVersionName = &v
	return s
}

func (s *CreateProductVersionRequest) SetTemplateType(v string) *CreateProductVersionRequest {
	s.TemplateType = &v
	return s
}

func (s *CreateProductVersionRequest) SetTemplateUrl(v string) *CreateProductVersionRequest {
	s.TemplateUrl = &v
	return s
}

func (s *CreateProductVersionRequest) Validate() error {
	return dara.Validate(s)
}
