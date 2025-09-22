// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGwpBenchmarkListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetGwpBenchmarkListRequest
	GetCode() *string
	SetProductId(v int64) *GetGwpBenchmarkListRequest
	GetProductId() *int64
	SetProductType(v int64) *GetGwpBenchmarkListRequest
	GetProductType() *int64
}

type GetGwpBenchmarkListRequest struct {
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

func (s GetGwpBenchmarkListRequest) String() string {
	return dara.Prettify(s)
}

func (s GetGwpBenchmarkListRequest) GoString() string {
	return s.String()
}

func (s *GetGwpBenchmarkListRequest) GetCode() *string {
	return s.Code
}

func (s *GetGwpBenchmarkListRequest) GetProductId() *int64 {
	return s.ProductId
}

func (s *GetGwpBenchmarkListRequest) GetProductType() *int64 {
	return s.ProductType
}

func (s *GetGwpBenchmarkListRequest) SetCode(v string) *GetGwpBenchmarkListRequest {
	s.Code = &v
	return s
}

func (s *GetGwpBenchmarkListRequest) SetProductId(v int64) *GetGwpBenchmarkListRequest {
	s.ProductId = &v
	return s
}

func (s *GetGwpBenchmarkListRequest) SetProductType(v int64) *GetGwpBenchmarkListRequest {
	s.ProductType = &v
	return s
}

func (s *GetGwpBenchmarkListRequest) Validate() error {
	return dara.Validate(s)
}
