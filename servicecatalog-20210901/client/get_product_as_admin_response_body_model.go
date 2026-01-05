// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProductAsAdminResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetProductDetail(v *GetProductAsAdminResponseBodyProductDetail) *GetProductAsAdminResponseBody
	GetProductDetail() *GetProductAsAdminResponseBodyProductDetail
	SetRequestId(v string) *GetProductAsAdminResponseBody
	GetRequestId() *string
	SetTagOptions(v []*GetProductAsAdminResponseBodyTagOptions) *GetProductAsAdminResponseBody
	GetTagOptions() []*GetProductAsAdminResponseBodyTagOptions
}

type GetProductAsAdminResponseBody struct {
	// The information about the product.
	ProductDetail *GetProductAsAdminResponseBodyProductDetail `json:"ProductDetail,omitempty" xml:"ProductDetail,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 0FEEF92D-4052-5202-87D0-3D8EC16F81BF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The tag options associated with the product.
	TagOptions []*GetProductAsAdminResponseBodyTagOptions `json:"TagOptions,omitempty" xml:"TagOptions,omitempty" type:"Repeated"`
}

func (s GetProductAsAdminResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetProductAsAdminResponseBody) GoString() string {
	return s.String()
}

func (s *GetProductAsAdminResponseBody) GetProductDetail() *GetProductAsAdminResponseBodyProductDetail {
	return s.ProductDetail
}

func (s *GetProductAsAdminResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetProductAsAdminResponseBody) GetTagOptions() []*GetProductAsAdminResponseBodyTagOptions {
	return s.TagOptions
}

func (s *GetProductAsAdminResponseBody) SetProductDetail(v *GetProductAsAdminResponseBodyProductDetail) *GetProductAsAdminResponseBody {
	s.ProductDetail = v
	return s
}

func (s *GetProductAsAdminResponseBody) SetRequestId(v string) *GetProductAsAdminResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetProductAsAdminResponseBody) SetTagOptions(v []*GetProductAsAdminResponseBodyTagOptions) *GetProductAsAdminResponseBody {
	s.TagOptions = v
	return s
}

func (s *GetProductAsAdminResponseBody) Validate() error {
	if s.ProductDetail != nil {
		if err := s.ProductDetail.Validate(); err != nil {
			return err
		}
	}
	if s.TagOptions != nil {
		for _, item := range s.TagOptions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetProductAsAdminResponseBodyProductDetail struct {
	// The creation time.
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
	// This is a product description.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
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
	// DEMO-Create an ECS instance
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
	// 	- RosTerraformTemplate: the Terraform template that is supported by Resource Orchestration Service (ROS).
	//
	// 	- RosStandardTemplate: the standard ROS template.
	//
	// example:
	//
	// RosTerraformTemplate
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s GetProductAsAdminResponseBodyProductDetail) String() string {
	return dara.Prettify(s)
}

func (s GetProductAsAdminResponseBodyProductDetail) GoString() string {
	return s.String()
}

func (s *GetProductAsAdminResponseBodyProductDetail) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetProductAsAdminResponseBodyProductDetail) GetDescription() *string {
	return s.Description
}

func (s *GetProductAsAdminResponseBodyProductDetail) GetProductArn() *string {
	return s.ProductArn
}

func (s *GetProductAsAdminResponseBodyProductDetail) GetProductId() *string {
	return s.ProductId
}

func (s *GetProductAsAdminResponseBodyProductDetail) GetProductName() *string {
	return s.ProductName
}

func (s *GetProductAsAdminResponseBodyProductDetail) GetProductType() *string {
	return s.ProductType
}

func (s *GetProductAsAdminResponseBodyProductDetail) GetProviderName() *string {
	return s.ProviderName
}

func (s *GetProductAsAdminResponseBodyProductDetail) GetTemplateType() *string {
	return s.TemplateType
}

func (s *GetProductAsAdminResponseBodyProductDetail) SetCreateTime(v string) *GetProductAsAdminResponseBodyProductDetail {
	s.CreateTime = &v
	return s
}

func (s *GetProductAsAdminResponseBodyProductDetail) SetDescription(v string) *GetProductAsAdminResponseBodyProductDetail {
	s.Description = &v
	return s
}

func (s *GetProductAsAdminResponseBodyProductDetail) SetProductArn(v string) *GetProductAsAdminResponseBodyProductDetail {
	s.ProductArn = &v
	return s
}

func (s *GetProductAsAdminResponseBodyProductDetail) SetProductId(v string) *GetProductAsAdminResponseBodyProductDetail {
	s.ProductId = &v
	return s
}

func (s *GetProductAsAdminResponseBodyProductDetail) SetProductName(v string) *GetProductAsAdminResponseBodyProductDetail {
	s.ProductName = &v
	return s
}

func (s *GetProductAsAdminResponseBodyProductDetail) SetProductType(v string) *GetProductAsAdminResponseBodyProductDetail {
	s.ProductType = &v
	return s
}

func (s *GetProductAsAdminResponseBodyProductDetail) SetProviderName(v string) *GetProductAsAdminResponseBodyProductDetail {
	s.ProviderName = &v
	return s
}

func (s *GetProductAsAdminResponseBodyProductDetail) SetTemplateType(v string) *GetProductAsAdminResponseBodyProductDetail {
	s.TemplateType = &v
	return s
}

func (s *GetProductAsAdminResponseBodyProductDetail) Validate() error {
	return dara.Validate(s)
}

type GetProductAsAdminResponseBodyTagOptions struct {
	// Indicates whether the tag option is enabled. Valid values:
	//
	// 	- true (default)
	//
	// 	- false
	//
	// example:
	//
	// true
	Active *bool `json:"Active,omitempty" xml:"Active,omitempty"`
	// The key of the tag option.
	//
	// example:
	//
	// k1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The ID of the owner of the tag option.
	//
	// example:
	//
	// 133413081827****
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The ID of the tag option.
	//
	// example:
	//
	// tag-bp1q65xd3j****
	TagOptionId *string `json:"TagOptionId,omitempty" xml:"TagOptionId,omitempty"`
	// The value of the tag option.
	//
	// example:
	//
	// v1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetProductAsAdminResponseBodyTagOptions) String() string {
	return dara.Prettify(s)
}

func (s GetProductAsAdminResponseBodyTagOptions) GoString() string {
	return s.String()
}

func (s *GetProductAsAdminResponseBodyTagOptions) GetActive() *bool {
	return s.Active
}

func (s *GetProductAsAdminResponseBodyTagOptions) GetKey() *string {
	return s.Key
}

func (s *GetProductAsAdminResponseBodyTagOptions) GetOwner() *string {
	return s.Owner
}

func (s *GetProductAsAdminResponseBodyTagOptions) GetTagOptionId() *string {
	return s.TagOptionId
}

func (s *GetProductAsAdminResponseBodyTagOptions) GetValue() *string {
	return s.Value
}

func (s *GetProductAsAdminResponseBodyTagOptions) SetActive(v bool) *GetProductAsAdminResponseBodyTagOptions {
	s.Active = &v
	return s
}

func (s *GetProductAsAdminResponseBodyTagOptions) SetKey(v string) *GetProductAsAdminResponseBodyTagOptions {
	s.Key = &v
	return s
}

func (s *GetProductAsAdminResponseBodyTagOptions) SetOwner(v string) *GetProductAsAdminResponseBodyTagOptions {
	s.Owner = &v
	return s
}

func (s *GetProductAsAdminResponseBodyTagOptions) SetTagOptionId(v string) *GetProductAsAdminResponseBodyTagOptions {
	s.TagOptionId = &v
	return s
}

func (s *GetProductAsAdminResponseBodyTagOptions) SetValue(v string) *GetProductAsAdminResponseBodyTagOptions {
	s.Value = &v
	return s
}

func (s *GetProductAsAdminResponseBodyTagOptions) Validate() error {
	return dara.Validate(s)
}
