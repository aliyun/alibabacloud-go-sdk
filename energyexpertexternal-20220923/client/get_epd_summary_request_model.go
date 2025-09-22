// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEpdSummaryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetEpdSummaryRequest
	GetCode() *string
	SetProductId(v int64) *GetEpdSummaryRequest
	GetProductId() *int64
	SetProductType(v int64) *GetEpdSummaryRequest
	GetProductType() *int64
}

type GetEpdSummaryRequest struct {
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

func (s GetEpdSummaryRequest) String() string {
	return dara.Prettify(s)
}

func (s GetEpdSummaryRequest) GoString() string {
	return s.String()
}

func (s *GetEpdSummaryRequest) GetCode() *string {
	return s.Code
}

func (s *GetEpdSummaryRequest) GetProductId() *int64 {
	return s.ProductId
}

func (s *GetEpdSummaryRequest) GetProductType() *int64 {
	return s.ProductType
}

func (s *GetEpdSummaryRequest) SetCode(v string) *GetEpdSummaryRequest {
	s.Code = &v
	return s
}

func (s *GetEpdSummaryRequest) SetProductId(v int64) *GetEpdSummaryRequest {
	s.ProductId = &v
	return s
}

func (s *GetEpdSummaryRequest) SetProductType(v int64) *GetEpdSummaryRequest {
	s.ProductType = &v
	return s
}

func (s *GetEpdSummaryRequest) Validate() error {
	return dara.Validate(s)
}
