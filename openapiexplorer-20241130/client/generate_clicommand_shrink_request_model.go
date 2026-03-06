// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateCLICommandShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAggregatePagination(v bool) *GenerateCLICommandShrinkRequest
	GetAggregatePagination() *bool
	SetApi(v string) *GenerateCLICommandShrinkRequest
	GetApi() *string
	SetApiParamsShrink(v string) *GenerateCLICommandShrinkRequest
	GetApiParamsShrink() *string
	SetApiVersion(v string) *GenerateCLICommandShrinkRequest
	GetApiVersion() *string
	SetJsonApiParams(v string) *GenerateCLICommandShrinkRequest
	GetJsonApiParams() *string
	SetProduct(v string) *GenerateCLICommandShrinkRequest
	GetProduct() *string
	SetRegionId(v string) *GenerateCLICommandShrinkRequest
	GetRegionId() *string
}

type GenerateCLICommandShrinkRequest struct {
	// Specifies whether to enable aggregation. If you enable this feature, the CLI automatically reads full data by page and aggregates the results.
	//
	// 	Warning:
	//
	// Only List operations that support paging can use this switch.
	//
	//
	//
	// - true: enables aggregation.
	//
	// - false: disables aggregation.
	//
	// example:
	//
	// true
	AggregatePagination *bool `json:"aggregatePagination,omitempty" xml:"aggregatePagination,omitempty"`
	// The name of the API.
	//
	// This parameter is required.
	//
	// example:
	//
	// DescribeRegions
	Api *string `json:"api,omitempty" xml:"api,omitempty"`
	// Deprecated
	//
	// The request parameters.
	ApiParamsShrink *string `json:"apiParams,omitempty" xml:"apiParams,omitempty"`
	// The version of the API.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2014-05-26
	ApiVersion *string `json:"apiVersion,omitempty" xml:"apiVersion,omitempty"`
	// The request parameters for the API in JSON format. This parameter has a lower priority than \\`apiParams\\`. If \\`apiParams\\` is set, this parameter is ignored.
	//
	// example:
	//
	// {
	//
	//     "InstanceChargeType": "PostPaid",
	//
	//     "ResourceType": "instance",
	//
	//     "AcceptLanguage": "en-US"
	//
	// }
	JsonApiParams *string `json:"jsonApiParams,omitempty" xml:"jsonApiParams,omitempty"`
	// The product code.
	//
	// - Call the GetRequestLog operation to obtain the product code from the response.
	//
	// - Find the product code in the URL of the product in OpenAPI Portal. For example, <props="china">the OpenAPI Portal URL for Short Message Service is https\\://api.aliyun.com/product/Dysmsapi. The product code for Short Message Service is Dysmsapi.
	//
	//   <props="intl">the OpenAPI Portal URL for Short Message Service is https\\://api.alibabacloud.com/product/Dysmsapi. The product code for Short Message Service is Dysmsapi.
	//
	// This parameter is required.
	//
	// example:
	//
	// Ecs
	Product *string `json:"product,omitempty" xml:"product,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s GenerateCLICommandShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GenerateCLICommandShrinkRequest) GoString() string {
	return s.String()
}

func (s *GenerateCLICommandShrinkRequest) GetAggregatePagination() *bool {
	return s.AggregatePagination
}

func (s *GenerateCLICommandShrinkRequest) GetApi() *string {
	return s.Api
}

func (s *GenerateCLICommandShrinkRequest) GetApiParamsShrink() *string {
	return s.ApiParamsShrink
}

func (s *GenerateCLICommandShrinkRequest) GetApiVersion() *string {
	return s.ApiVersion
}

func (s *GenerateCLICommandShrinkRequest) GetJsonApiParams() *string {
	return s.JsonApiParams
}

func (s *GenerateCLICommandShrinkRequest) GetProduct() *string {
	return s.Product
}

func (s *GenerateCLICommandShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GenerateCLICommandShrinkRequest) SetAggregatePagination(v bool) *GenerateCLICommandShrinkRequest {
	s.AggregatePagination = &v
	return s
}

func (s *GenerateCLICommandShrinkRequest) SetApi(v string) *GenerateCLICommandShrinkRequest {
	s.Api = &v
	return s
}

func (s *GenerateCLICommandShrinkRequest) SetApiParamsShrink(v string) *GenerateCLICommandShrinkRequest {
	s.ApiParamsShrink = &v
	return s
}

func (s *GenerateCLICommandShrinkRequest) SetApiVersion(v string) *GenerateCLICommandShrinkRequest {
	s.ApiVersion = &v
	return s
}

func (s *GenerateCLICommandShrinkRequest) SetJsonApiParams(v string) *GenerateCLICommandShrinkRequest {
	s.JsonApiParams = &v
	return s
}

func (s *GenerateCLICommandShrinkRequest) SetProduct(v string) *GenerateCLICommandShrinkRequest {
	s.Product = &v
	return s
}

func (s *GenerateCLICommandShrinkRequest) SetRegionId(v string) *GenerateCLICommandShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *GenerateCLICommandShrinkRequest) Validate() error {
	return dara.Validate(s)
}
