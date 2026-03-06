// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProductEndpointsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProduct(v string) *GetProductEndpointsRequest
	GetProduct() *string
}

type GetProductEndpointsRequest struct {
	// The product code.
	//
	// - Call the GetRequestLog operation and find the product code in the response.
	//
	// - Find the product code in the URL of the OpenAPI Portal page for the product. For example, <props="china">the URL for the Short Message Service (SMS) OpenAPI Portal page is https\\://api.aliyun.com/product/Dysmsapi. The product code is Dysmsapi.
	//
	//   <props="intl">the URL for the Short Message Service (SMS) OpenAPI Portal page is https\\://api.alibabacloud.com/product/Dysmsapi. The product code is Dysmsapi.
	//
	// This parameter is required.
	//
	// example:
	//
	// Ecs
	Product *string `json:"product,omitempty" xml:"product,omitempty"`
}

func (s GetProductEndpointsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetProductEndpointsRequest) GoString() string {
	return s.String()
}

func (s *GetProductEndpointsRequest) GetProduct() *string {
	return s.Product
}

func (s *GetProductEndpointsRequest) SetProduct(v string) *GetProductEndpointsRequest {
	s.Product = &v
	return s
}

func (s *GetProductEndpointsRequest) Validate() error {
	return dara.Validate(s)
}
