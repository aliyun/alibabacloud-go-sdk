// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateCLICommandRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAggregatePagination(v bool) *GenerateCLICommandRequest
	GetAggregatePagination() *bool
	SetApi(v string) *GenerateCLICommandRequest
	GetApi() *string
	SetApiParams(v map[string]interface{}) *GenerateCLICommandRequest
	GetApiParams() map[string]interface{}
	SetApiVersion(v string) *GenerateCLICommandRequest
	GetApiVersion() *string
	SetJsonApiParams(v string) *GenerateCLICommandRequest
	GetJsonApiParams() *string
	SetProduct(v string) *GenerateCLICommandRequest
	GetProduct() *string
	SetRegionId(v string) *GenerateCLICommandRequest
	GetRegionId() *string
}

type GenerateCLICommandRequest struct {
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
	ApiParams map[string]interface{} `json:"apiParams,omitempty" xml:"apiParams,omitempty"`
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

func (s GenerateCLICommandRequest) String() string {
	return dara.Prettify(s)
}

func (s GenerateCLICommandRequest) GoString() string {
	return s.String()
}

func (s *GenerateCLICommandRequest) GetAggregatePagination() *bool {
	return s.AggregatePagination
}

func (s *GenerateCLICommandRequest) GetApi() *string {
	return s.Api
}

func (s *GenerateCLICommandRequest) GetApiParams() map[string]interface{} {
	return s.ApiParams
}

func (s *GenerateCLICommandRequest) GetApiVersion() *string {
	return s.ApiVersion
}

func (s *GenerateCLICommandRequest) GetJsonApiParams() *string {
	return s.JsonApiParams
}

func (s *GenerateCLICommandRequest) GetProduct() *string {
	return s.Product
}

func (s *GenerateCLICommandRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GenerateCLICommandRequest) SetAggregatePagination(v bool) *GenerateCLICommandRequest {
	s.AggregatePagination = &v
	return s
}

func (s *GenerateCLICommandRequest) SetApi(v string) *GenerateCLICommandRequest {
	s.Api = &v
	return s
}

func (s *GenerateCLICommandRequest) SetApiParams(v map[string]interface{}) *GenerateCLICommandRequest {
	s.ApiParams = v
	return s
}

func (s *GenerateCLICommandRequest) SetApiVersion(v string) *GenerateCLICommandRequest {
	s.ApiVersion = &v
	return s
}

func (s *GenerateCLICommandRequest) SetJsonApiParams(v string) *GenerateCLICommandRequest {
	s.JsonApiParams = &v
	return s
}

func (s *GenerateCLICommandRequest) SetProduct(v string) *GenerateCLICommandRequest {
	s.Product = &v
	return s
}

func (s *GenerateCLICommandRequest) SetRegionId(v string) *GenerateCLICommandRequest {
	s.RegionId = &v
	return s
}

func (s *GenerateCLICommandRequest) Validate() error {
	return dara.Validate(s)
}
