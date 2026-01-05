// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProductsAsAdminResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListProductsAsAdminResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListProductsAsAdminResponseBody
	GetPageSize() *int32
	SetProductDetails(v []*ListProductsAsAdminResponseBodyProductDetails) *ListProductsAsAdminResponseBody
	GetProductDetails() []*ListProductsAsAdminResponseBodyProductDetails
	SetRequestId(v string) *ListProductsAsAdminResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListProductsAsAdminResponseBody
	GetTotalCount() *int32
}

type ListProductsAsAdminResponseBody struct {
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
	ProductDetails []*ListProductsAsAdminResponseBodyProductDetails `json:"ProductDetails,omitempty" xml:"ProductDetails,omitempty" type:"Repeated"`
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
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListProductsAsAdminResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListProductsAsAdminResponseBody) GoString() string {
	return s.String()
}

func (s *ListProductsAsAdminResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListProductsAsAdminResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListProductsAsAdminResponseBody) GetProductDetails() []*ListProductsAsAdminResponseBodyProductDetails {
	return s.ProductDetails
}

func (s *ListProductsAsAdminResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListProductsAsAdminResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListProductsAsAdminResponseBody) SetPageNumber(v int32) *ListProductsAsAdminResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListProductsAsAdminResponseBody) SetPageSize(v int32) *ListProductsAsAdminResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListProductsAsAdminResponseBody) SetProductDetails(v []*ListProductsAsAdminResponseBodyProductDetails) *ListProductsAsAdminResponseBody {
	s.ProductDetails = v
	return s
}

func (s *ListProductsAsAdminResponseBody) SetRequestId(v string) *ListProductsAsAdminResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListProductsAsAdminResponseBody) SetTotalCount(v int32) *ListProductsAsAdminResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListProductsAsAdminResponseBody) Validate() error {
	if s.ProductDetails != nil {
		for _, item := range s.ProductDetails {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListProductsAsAdminResponseBodyProductDetails struct {
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

func (s ListProductsAsAdminResponseBodyProductDetails) String() string {
	return dara.Prettify(s)
}

func (s ListProductsAsAdminResponseBodyProductDetails) GoString() string {
	return s.String()
}

func (s *ListProductsAsAdminResponseBodyProductDetails) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListProductsAsAdminResponseBodyProductDetails) GetDescription() *string {
	return s.Description
}

func (s *ListProductsAsAdminResponseBodyProductDetails) GetProductArn() *string {
	return s.ProductArn
}

func (s *ListProductsAsAdminResponseBodyProductDetails) GetProductId() *string {
	return s.ProductId
}

func (s *ListProductsAsAdminResponseBodyProductDetails) GetProductName() *string {
	return s.ProductName
}

func (s *ListProductsAsAdminResponseBodyProductDetails) GetProductType() *string {
	return s.ProductType
}

func (s *ListProductsAsAdminResponseBodyProductDetails) GetProviderName() *string {
	return s.ProviderName
}

func (s *ListProductsAsAdminResponseBodyProductDetails) GetTemplateType() *string {
	return s.TemplateType
}

func (s *ListProductsAsAdminResponseBodyProductDetails) SetCreateTime(v string) *ListProductsAsAdminResponseBodyProductDetails {
	s.CreateTime = &v
	return s
}

func (s *ListProductsAsAdminResponseBodyProductDetails) SetDescription(v string) *ListProductsAsAdminResponseBodyProductDetails {
	s.Description = &v
	return s
}

func (s *ListProductsAsAdminResponseBodyProductDetails) SetProductArn(v string) *ListProductsAsAdminResponseBodyProductDetails {
	s.ProductArn = &v
	return s
}

func (s *ListProductsAsAdminResponseBodyProductDetails) SetProductId(v string) *ListProductsAsAdminResponseBodyProductDetails {
	s.ProductId = &v
	return s
}

func (s *ListProductsAsAdminResponseBodyProductDetails) SetProductName(v string) *ListProductsAsAdminResponseBodyProductDetails {
	s.ProductName = &v
	return s
}

func (s *ListProductsAsAdminResponseBodyProductDetails) SetProductType(v string) *ListProductsAsAdminResponseBodyProductDetails {
	s.ProductType = &v
	return s
}

func (s *ListProductsAsAdminResponseBodyProductDetails) SetProviderName(v string) *ListProductsAsAdminResponseBodyProductDetails {
	s.ProviderName = &v
	return s
}

func (s *ListProductsAsAdminResponseBodyProductDetails) SetTemplateType(v string) *ListProductsAsAdminResponseBodyProductDetails {
	s.TemplateType = &v
	return s
}

func (s *ListProductsAsAdminResponseBodyProductDetails) Validate() error {
	return dara.Validate(s)
}
