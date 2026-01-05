// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProductAsEndUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetProductSummary(v *GetProductAsEndUserResponseBodyProductSummary) *GetProductAsEndUserResponseBody
	GetProductSummary() *GetProductAsEndUserResponseBodyProductSummary
	SetRequestId(v string) *GetProductAsEndUserResponseBody
	GetRequestId() *string
}

type GetProductAsEndUserResponseBody struct {
	// The information about the product.
	ProductSummary *GetProductAsEndUserResponseBodyProductSummary `json:"ProductSummary,omitempty" xml:"ProductSummary,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 0FEEF92D-4052-5202-87D0-3D8EC16F81BF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetProductAsEndUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetProductAsEndUserResponseBody) GoString() string {
	return s.String()
}

func (s *GetProductAsEndUserResponseBody) GetProductSummary() *GetProductAsEndUserResponseBodyProductSummary {
	return s.ProductSummary
}

func (s *GetProductAsEndUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetProductAsEndUserResponseBody) SetProductSummary(v *GetProductAsEndUserResponseBodyProductSummary) *GetProductAsEndUserResponseBody {
	s.ProductSummary = v
	return s
}

func (s *GetProductAsEndUserResponseBody) SetRequestId(v string) *GetProductAsEndUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetProductAsEndUserResponseBody) Validate() error {
	if s.ProductSummary != nil {
		if err := s.ProductSummary.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetProductAsEndUserResponseBodyProductSummary struct {
	// The time when the product was created.
	//
	// The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-04-12T06:10:37Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the product.
	//
	// example:
	//
	// The description of the product.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Indicates whether the default launch option exists. Valid values:
	//
	// 	- true: The default launch option exists. In this case, the PortfolioId parameter is not required when the product is launched or when the information about the product instance is updated.
	//
	// 	- false: The default launch option does not exist. In this case, the PortfolioId parameter is required when the product is launched or when the information about the product instance is updated. For more information about how to obtain the value of the PortfolioId parameter, see [ListLaunchOptions](~~ListLaunchOptions~~).
	//
	// > If the product is added to only one product portfolio, the default launch option exists. If the product is added to multiple product portfolios, multiple launch options exist at the same time. However, no default launch options exist.
	//
	// example:
	//
	// true
	HasDefaultLaunchOption *bool `json:"HasDefaultLaunchOption,omitempty" xml:"HasDefaultLaunchOption,omitempty"`
	// The Alibaba Cloud Resource Name (ARN) of the product.
	//
	// example:
	//
	// acs:servicecatalog:cn-hangzhou:146611588617****:product/prod-bp18r7q127****
	ProductArn *string `json:"ProductArn,omitempty" xml:"ProductArn,omitempty"`
	// The ID of the product.
	//
	// example:
	//
	// prod-bp18r7q127****
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// The name of the product.
	//
	// example:
	//
	// DEMO-Create an ECS instance.
	ProductName *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	// The type of the product.
	//
	// The value is fixed as Ros, which indicates Resource Orchestration Service (ROS).
	//
	// example:
	//
	// Ros
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The provider of the product.
	//
	// example:
	//
	// IT team
	ProviderName *string `json:"ProviderName,omitempty" xml:"ProviderName,omitempty"`
	// The type of the product template. Valid values:
	//
	// 	- RosTerraformTemplate: The Terraform template that is supported by Resource Orchestration Service (ROS).
	//
	// 	- RosStandardTemplate: The standard ROS template.
	//
	// example:
	//
	// RosTerraformTemplate
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s GetProductAsEndUserResponseBodyProductSummary) String() string {
	return dara.Prettify(s)
}

func (s GetProductAsEndUserResponseBodyProductSummary) GoString() string {
	return s.String()
}

func (s *GetProductAsEndUserResponseBodyProductSummary) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetProductAsEndUserResponseBodyProductSummary) GetDescription() *string {
	return s.Description
}

func (s *GetProductAsEndUserResponseBodyProductSummary) GetHasDefaultLaunchOption() *bool {
	return s.HasDefaultLaunchOption
}

func (s *GetProductAsEndUserResponseBodyProductSummary) GetProductArn() *string {
	return s.ProductArn
}

func (s *GetProductAsEndUserResponseBodyProductSummary) GetProductId() *string {
	return s.ProductId
}

func (s *GetProductAsEndUserResponseBodyProductSummary) GetProductName() *string {
	return s.ProductName
}

func (s *GetProductAsEndUserResponseBodyProductSummary) GetProductType() *string {
	return s.ProductType
}

func (s *GetProductAsEndUserResponseBodyProductSummary) GetProviderName() *string {
	return s.ProviderName
}

func (s *GetProductAsEndUserResponseBodyProductSummary) GetTemplateType() *string {
	return s.TemplateType
}

func (s *GetProductAsEndUserResponseBodyProductSummary) SetCreateTime(v string) *GetProductAsEndUserResponseBodyProductSummary {
	s.CreateTime = &v
	return s
}

func (s *GetProductAsEndUserResponseBodyProductSummary) SetDescription(v string) *GetProductAsEndUserResponseBodyProductSummary {
	s.Description = &v
	return s
}

func (s *GetProductAsEndUserResponseBodyProductSummary) SetHasDefaultLaunchOption(v bool) *GetProductAsEndUserResponseBodyProductSummary {
	s.HasDefaultLaunchOption = &v
	return s
}

func (s *GetProductAsEndUserResponseBodyProductSummary) SetProductArn(v string) *GetProductAsEndUserResponseBodyProductSummary {
	s.ProductArn = &v
	return s
}

func (s *GetProductAsEndUserResponseBodyProductSummary) SetProductId(v string) *GetProductAsEndUserResponseBodyProductSummary {
	s.ProductId = &v
	return s
}

func (s *GetProductAsEndUserResponseBodyProductSummary) SetProductName(v string) *GetProductAsEndUserResponseBodyProductSummary {
	s.ProductName = &v
	return s
}

func (s *GetProductAsEndUserResponseBodyProductSummary) SetProductType(v string) *GetProductAsEndUserResponseBodyProductSummary {
	s.ProductType = &v
	return s
}

func (s *GetProductAsEndUserResponseBodyProductSummary) SetProviderName(v string) *GetProductAsEndUserResponseBodyProductSummary {
	s.ProviderName = &v
	return s
}

func (s *GetProductAsEndUserResponseBodyProductSummary) SetTemplateType(v string) *GetProductAsEndUserResponseBodyProductSummary {
	s.TemplateType = &v
	return s
}

func (s *GetProductAsEndUserResponseBodyProductSummary) Validate() error {
	return dara.Validate(s)
}
