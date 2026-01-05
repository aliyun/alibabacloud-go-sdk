// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProductsAsEndUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListProductsAsEndUserResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListProductsAsEndUserResponseBody
	GetPageSize() *int32
	SetProductSummaries(v []*ListProductsAsEndUserResponseBodyProductSummaries) *ListProductsAsEndUserResponseBody
	GetProductSummaries() []*ListProductsAsEndUserResponseBodyProductSummaries
	SetRequestId(v string) *ListProductsAsEndUserResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListProductsAsEndUserResponseBody
	GetTotalCount() *int32
}

type ListProductsAsEndUserResponseBody struct {
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The products.
	ProductSummaries []*ListProductsAsEndUserResponseBodyProductSummaries `json:"ProductSummaries,omitempty" xml:"ProductSummaries,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 0FEEF92D-4052-5202-87D0-3D8EC16F81BF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListProductsAsEndUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListProductsAsEndUserResponseBody) GoString() string {
	return s.String()
}

func (s *ListProductsAsEndUserResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListProductsAsEndUserResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListProductsAsEndUserResponseBody) GetProductSummaries() []*ListProductsAsEndUserResponseBodyProductSummaries {
	return s.ProductSummaries
}

func (s *ListProductsAsEndUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListProductsAsEndUserResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListProductsAsEndUserResponseBody) SetPageNumber(v int32) *ListProductsAsEndUserResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListProductsAsEndUserResponseBody) SetPageSize(v int32) *ListProductsAsEndUserResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListProductsAsEndUserResponseBody) SetProductSummaries(v []*ListProductsAsEndUserResponseBodyProductSummaries) *ListProductsAsEndUserResponseBody {
	s.ProductSummaries = v
	return s
}

func (s *ListProductsAsEndUserResponseBody) SetRequestId(v string) *ListProductsAsEndUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListProductsAsEndUserResponseBody) SetTotalCount(v int32) *ListProductsAsEndUserResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListProductsAsEndUserResponseBody) Validate() error {
	if s.ProductSummaries != nil {
		for _, item := range s.ProductSummaries {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListProductsAsEndUserResponseBodyProductSummaries struct {
	// The time when the product was created.
	//
	// The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ssZ format. The time is displayed in UTC.
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
	// 	- false: The default launch option does not exist. In this case, the PortfolioId parameter is required when the product is launched or when the information about the product instance is updated. For more information about the PortfolioId parameter, see [ListLaunchOptions](~~ListLaunchOptions~~).
	//
	// >  If the product is added to only one product portfolio, the default launch option exists. If the product is added to multiple product portfolios, multiple launch options exist at the same time. However, no default launch options exist.
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
	// The product ID.
	//
	// example:
	//
	// prod-bp18r7q127****
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// The name of the product.
	//
	// example:
	//
	// DEMO-Create an ECS instance
	ProductName *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	// The type of the product.
	//
	// The value is set to Ros, which indicates Resource Orchestration Service (ROS).
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
	// 	- RosTerraformTemplate: the Terraform template that is supported by ROS.
	//
	// 	- RosStandardTemplate: the standard ROS template.
	//
	// example:
	//
	// RosTerraformTemplate
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s ListProductsAsEndUserResponseBodyProductSummaries) String() string {
	return dara.Prettify(s)
}

func (s ListProductsAsEndUserResponseBodyProductSummaries) GoString() string {
	return s.String()
}

func (s *ListProductsAsEndUserResponseBodyProductSummaries) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListProductsAsEndUserResponseBodyProductSummaries) GetDescription() *string {
	return s.Description
}

func (s *ListProductsAsEndUserResponseBodyProductSummaries) GetHasDefaultLaunchOption() *bool {
	return s.HasDefaultLaunchOption
}

func (s *ListProductsAsEndUserResponseBodyProductSummaries) GetProductArn() *string {
	return s.ProductArn
}

func (s *ListProductsAsEndUserResponseBodyProductSummaries) GetProductId() *string {
	return s.ProductId
}

func (s *ListProductsAsEndUserResponseBodyProductSummaries) GetProductName() *string {
	return s.ProductName
}

func (s *ListProductsAsEndUserResponseBodyProductSummaries) GetProductType() *string {
	return s.ProductType
}

func (s *ListProductsAsEndUserResponseBodyProductSummaries) GetProviderName() *string {
	return s.ProviderName
}

func (s *ListProductsAsEndUserResponseBodyProductSummaries) GetTemplateType() *string {
	return s.TemplateType
}

func (s *ListProductsAsEndUserResponseBodyProductSummaries) SetCreateTime(v string) *ListProductsAsEndUserResponseBodyProductSummaries {
	s.CreateTime = &v
	return s
}

func (s *ListProductsAsEndUserResponseBodyProductSummaries) SetDescription(v string) *ListProductsAsEndUserResponseBodyProductSummaries {
	s.Description = &v
	return s
}

func (s *ListProductsAsEndUserResponseBodyProductSummaries) SetHasDefaultLaunchOption(v bool) *ListProductsAsEndUserResponseBodyProductSummaries {
	s.HasDefaultLaunchOption = &v
	return s
}

func (s *ListProductsAsEndUserResponseBodyProductSummaries) SetProductArn(v string) *ListProductsAsEndUserResponseBodyProductSummaries {
	s.ProductArn = &v
	return s
}

func (s *ListProductsAsEndUserResponseBodyProductSummaries) SetProductId(v string) *ListProductsAsEndUserResponseBodyProductSummaries {
	s.ProductId = &v
	return s
}

func (s *ListProductsAsEndUserResponseBodyProductSummaries) SetProductName(v string) *ListProductsAsEndUserResponseBodyProductSummaries {
	s.ProductName = &v
	return s
}

func (s *ListProductsAsEndUserResponseBodyProductSummaries) SetProductType(v string) *ListProductsAsEndUserResponseBodyProductSummaries {
	s.ProductType = &v
	return s
}

func (s *ListProductsAsEndUserResponseBodyProductSummaries) SetProviderName(v string) *ListProductsAsEndUserResponseBodyProductSummaries {
	s.ProviderName = &v
	return s
}

func (s *ListProductsAsEndUserResponseBodyProductSummaries) SetTemplateType(v string) *ListProductsAsEndUserResponseBodyProductSummaries {
	s.TemplateType = &v
	return s
}

func (s *ListProductsAsEndUserResponseBodyProductSummaries) Validate() error {
	return dara.Validate(s)
}
