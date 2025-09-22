// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGwpInventorySummaryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetGwpInventorySummaryRequest
	GetCode() *string
	SetProductId(v int64) *GetGwpInventorySummaryRequest
	GetProductId() *int64
	SetProductType(v int64) *GetGwpInventorySummaryRequest
	GetProductType() *int64
}

type GetGwpInventorySummaryRequest struct {
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

func (s GetGwpInventorySummaryRequest) String() string {
	return dara.Prettify(s)
}

func (s GetGwpInventorySummaryRequest) GoString() string {
	return s.String()
}

func (s *GetGwpInventorySummaryRequest) GetCode() *string {
	return s.Code
}

func (s *GetGwpInventorySummaryRequest) GetProductId() *int64 {
	return s.ProductId
}

func (s *GetGwpInventorySummaryRequest) GetProductType() *int64 {
	return s.ProductType
}

func (s *GetGwpInventorySummaryRequest) SetCode(v string) *GetGwpInventorySummaryRequest {
	s.Code = &v
	return s
}

func (s *GetGwpInventorySummaryRequest) SetProductId(v int64) *GetGwpInventorySummaryRequest {
	s.ProductId = &v
	return s
}

func (s *GetGwpInventorySummaryRequest) SetProductType(v int64) *GetGwpInventorySummaryRequest {
	s.ProductType = &v
	return s
}

func (s *GetGwpInventorySummaryRequest) Validate() error {
	return dara.Validate(s)
}
