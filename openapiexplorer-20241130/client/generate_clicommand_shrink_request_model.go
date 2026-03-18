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
	// Enable aggregation. If enabled, the CLI automatically reads full data using pagination and aggregates the results.
	//
	// 	Warning:
	//
	// You can use this option only with List operations that support pagination.
	//
	//
	//
	// - true: Enable
	//
	// - false: Disable
	//
	// example:
	//
	// true
	AggregatePagination *bool `json:"aggregatePagination,omitempty" xml:"aggregatePagination,omitempty"`
	// API name.
	//
	// This parameter is required.
	//
	// example:
	//
	// DescribeRegions
	Api *string `json:"api,omitempty" xml:"api,omitempty"`
	// Deprecated
	//
	// Request parameters.
	ApiParamsShrink *string `json:"apiParams,omitempty" xml:"apiParams,omitempty"`
	// API version.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2014-05-26
	ApiVersion *string `json:"apiVersion,omitempty" xml:"apiVersion,omitempty"`
	// API input parameters in JSON format. This parameter has lower priority than apiParams. If you set apiParams, this parameter is ignored.
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
	// Product code.
	//
	// - Call the GetRequestLog operation and get the product code from the response.
	//
	// - Find the product code in the OpenAPI portal URL. For example, the OpenAPI portal URL for Short Message Service is https\\://api.aliyun.com/product/Dysmsapi. The product code for Short Message Service is Dysmsapi. In international regions, the URL is https\\://api.alibabacloud.com/product/Dysmsapi. The product code remains Dysmsapi.
	//
	// This parameter is required.
	//
	// example:
	//
	// Ecs
	Product *string `json:"product,omitempty" xml:"product,omitempty"`
	// Region ID.
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
