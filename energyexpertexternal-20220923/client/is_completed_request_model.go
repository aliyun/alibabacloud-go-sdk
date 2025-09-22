// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIsCompletedRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *IsCompletedRequest
	GetCode() *string
	SetProductId(v int64) *IsCompletedRequest
	GetProductId() *int64
	SetProductType(v int64) *IsCompletedRequest
	GetProductType() *int64
}

type IsCompletedRequest struct {
	// The enterprise code.
	//
	// This parameter is required.
	//
	// example:
	//
	// C-20080808-1
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The product id.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1024
	ProductId *int64 `json:"productId,omitempty" xml:"productId,omitempty"`
	// Product type: 1 indicates that the carbon footprint of the product is requested, and 5 indicates that the carbon footprint of the supply chain is requested.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	ProductType *int64 `json:"productType,omitempty" xml:"productType,omitempty"`
}

func (s IsCompletedRequest) String() string {
	return dara.Prettify(s)
}

func (s IsCompletedRequest) GoString() string {
	return s.String()
}

func (s *IsCompletedRequest) GetCode() *string {
	return s.Code
}

func (s *IsCompletedRequest) GetProductId() *int64 {
	return s.ProductId
}

func (s *IsCompletedRequest) GetProductType() *int64 {
	return s.ProductType
}

func (s *IsCompletedRequest) SetCode(v string) *IsCompletedRequest {
	s.Code = &v
	return s
}

func (s *IsCompletedRequest) SetProductId(v int64) *IsCompletedRequest {
	s.ProductId = &v
	return s
}

func (s *IsCompletedRequest) SetProductType(v int64) *IsCompletedRequest {
	s.ProductType = &v
	return s
}

func (s *IsCompletedRequest) Validate() error {
	return dara.Validate(s)
}
