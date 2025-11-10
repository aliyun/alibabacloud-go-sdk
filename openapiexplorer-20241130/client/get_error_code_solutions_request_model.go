// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetErrorCodeSolutionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *GetErrorCodeSolutionsRequest
	GetAcceptLanguage() *string
	SetErrorCode(v string) *GetErrorCodeSolutionsRequest
	GetErrorCode() *string
	SetErrorMessage(v string) *GetErrorCodeSolutionsRequest
	GetErrorMessage() *string
	SetProduct(v string) *GetErrorCodeSolutionsRequest
	GetProduct() *string
}

type GetErrorCodeSolutionsRequest struct {
	// The language of the solution. Valid values: zh-CN and en-US. Not all of the solutions are available in English. If you set this parameter to en-US, but the corresponding solution is actually not available in English, no response is returned.
	//
	// example:
	//
	// zh-CN
	AcceptLanguage *string `json:"acceptLanguage,omitempty" xml:"acceptLanguage,omitempty"`
	// The error code based on which you want to query a solution.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234-56789012
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// The error message for which you want to query a solution. This parameter must be configured together with the errorCode parameter.
	//
	// example:
	//
	// An error occurred while processing your request.
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// The product code. You can use one of the following methods to query a product code:
	//
	// 	- Call the GetRequestLog operation to query a product code from the response.
	//
	// 	- Query the code of a product in the OpenAPI Explorer URL of the product. For example, the OpenAPI Explorer URL of Short Message Service (SMS) is https://api.alibabacloud.com/product/Dysmsapi. Therefore, the product code of SMS is Dysmsapi.
	//
	// example:
	//
	// oss
	Product *string `json:"product,omitempty" xml:"product,omitempty"`
}

func (s GetErrorCodeSolutionsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetErrorCodeSolutionsRequest) GoString() string {
	return s.String()
}

func (s *GetErrorCodeSolutionsRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *GetErrorCodeSolutionsRequest) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetErrorCodeSolutionsRequest) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetErrorCodeSolutionsRequest) GetProduct() *string {
	return s.Product
}

func (s *GetErrorCodeSolutionsRequest) SetAcceptLanguage(v string) *GetErrorCodeSolutionsRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *GetErrorCodeSolutionsRequest) SetErrorCode(v string) *GetErrorCodeSolutionsRequest {
	s.ErrorCode = &v
	return s
}

func (s *GetErrorCodeSolutionsRequest) SetErrorMessage(v string) *GetErrorCodeSolutionsRequest {
	s.ErrorMessage = &v
	return s
}

func (s *GetErrorCodeSolutionsRequest) SetProduct(v string) *GetErrorCodeSolutionsRequest {
	s.Product = &v
	return s
}

func (s *GetErrorCodeSolutionsRequest) Validate() error {
	return dara.Validate(s)
}
