// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApiDefinitionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiVersion(v string) *ListApiDefinitionsRequest
	GetApiVersion() *string
	SetProduct(v string) *ListApiDefinitionsRequest
	GetProduct() *string
}

type ListApiDefinitionsRequest struct {
	// The version of the API.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2014-05-26
	ApiVersion *string `json:"apiVersion,omitempty" xml:"apiVersion,omitempty"`
	// The product code.
	//
	// - Call the GetRequestLog operation to obtain the product code from the response.
	//
	// - Find the product code in the URL of the OpenAPI Portal. For example, <props="china">the URL of the OpenAPI Portal for Short Message Service is https\\://api.aliyun.com/product/Dysmsapi. The product code for Short Message Service is Dysmsapi.
	//
	//   <props="intl">the URL of the OpenAPI Portal for Short Message Service is https\\://api.alibabacloud.com/product/Dysmsapi. The product code for Short Message Service is Dysmsapi.
	//
	// This parameter is required.
	//
	// example:
	//
	// Ecs
	Product *string `json:"product,omitempty" xml:"product,omitempty"`
}

func (s ListApiDefinitionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListApiDefinitionsRequest) GoString() string {
	return s.String()
}

func (s *ListApiDefinitionsRequest) GetApiVersion() *string {
	return s.ApiVersion
}

func (s *ListApiDefinitionsRequest) GetProduct() *string {
	return s.Product
}

func (s *ListApiDefinitionsRequest) SetApiVersion(v string) *ListApiDefinitionsRequest {
	s.ApiVersion = &v
	return s
}

func (s *ListApiDefinitionsRequest) SetProduct(v string) *ListApiDefinitionsRequest {
	s.Product = &v
	return s
}

func (s *ListApiDefinitionsRequest) Validate() error {
	return dara.Validate(s)
}
