// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGwpBenchmarkSummaryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetGwpBenchmarkSummaryRequest
	GetCode() *string
	SetProductId(v int64) *GetGwpBenchmarkSummaryRequest
	GetProductId() *int64
	SetProductType(v int64) *GetGwpBenchmarkSummaryRequest
	GetProductType() *int64
}

type GetGwpBenchmarkSummaryRequest struct {
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

func (s GetGwpBenchmarkSummaryRequest) String() string {
	return dara.Prettify(s)
}

func (s GetGwpBenchmarkSummaryRequest) GoString() string {
	return s.String()
}

func (s *GetGwpBenchmarkSummaryRequest) GetCode() *string {
	return s.Code
}

func (s *GetGwpBenchmarkSummaryRequest) GetProductId() *int64 {
	return s.ProductId
}

func (s *GetGwpBenchmarkSummaryRequest) GetProductType() *int64 {
	return s.ProductType
}

func (s *GetGwpBenchmarkSummaryRequest) SetCode(v string) *GetGwpBenchmarkSummaryRequest {
	s.Code = &v
	return s
}

func (s *GetGwpBenchmarkSummaryRequest) SetProductId(v int64) *GetGwpBenchmarkSummaryRequest {
	s.ProductId = &v
	return s
}

func (s *GetGwpBenchmarkSummaryRequest) SetProductType(v int64) *GetGwpBenchmarkSummaryRequest {
	s.ProductType = &v
	return s
}

func (s *GetGwpBenchmarkSummaryRequest) Validate() error {
	return dara.Validate(s)
}
