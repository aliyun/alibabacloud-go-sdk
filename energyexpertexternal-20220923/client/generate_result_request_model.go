// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GenerateResultRequest
	GetCode() *string
	SetProductId(v int64) *GenerateResultRequest
	GetProductId() *int64
	SetProductType(v int64) *GenerateResultRequest
	GetProductType() *int64
}

type GenerateResultRequest struct {
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

func (s GenerateResultRequest) String() string {
	return dara.Prettify(s)
}

func (s GenerateResultRequest) GoString() string {
	return s.String()
}

func (s *GenerateResultRequest) GetCode() *string {
	return s.Code
}

func (s *GenerateResultRequest) GetProductId() *int64 {
	return s.ProductId
}

func (s *GenerateResultRequest) GetProductType() *int64 {
	return s.ProductType
}

func (s *GenerateResultRequest) SetCode(v string) *GenerateResultRequest {
	s.Code = &v
	return s
}

func (s *GenerateResultRequest) SetProductId(v int64) *GenerateResultRequest {
	s.ProductId = &v
	return s
}

func (s *GenerateResultRequest) SetProductType(v int64) *GenerateResultRequest {
	s.ProductType = &v
	return s
}

func (s *GenerateResultRequest) Validate() error {
	return dara.Validate(s)
}
