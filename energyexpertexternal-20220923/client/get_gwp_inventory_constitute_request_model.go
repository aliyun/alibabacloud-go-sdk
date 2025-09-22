// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGwpInventoryConstituteRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetGwpInventoryConstituteRequest
	GetCode() *string
	SetProductId(v int64) *GetGwpInventoryConstituteRequest
	GetProductId() *int64
	SetProductType(v int64) *GetGwpInventoryConstituteRequest
	GetProductType() *int64
}

type GetGwpInventoryConstituteRequest struct {
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

func (s GetGwpInventoryConstituteRequest) String() string {
	return dara.Prettify(s)
}

func (s GetGwpInventoryConstituteRequest) GoString() string {
	return s.String()
}

func (s *GetGwpInventoryConstituteRequest) GetCode() *string {
	return s.Code
}

func (s *GetGwpInventoryConstituteRequest) GetProductId() *int64 {
	return s.ProductId
}

func (s *GetGwpInventoryConstituteRequest) GetProductType() *int64 {
	return s.ProductType
}

func (s *GetGwpInventoryConstituteRequest) SetCode(v string) *GetGwpInventoryConstituteRequest {
	s.Code = &v
	return s
}

func (s *GetGwpInventoryConstituteRequest) SetProductId(v int64) *GetGwpInventoryConstituteRequest {
	s.ProductId = &v
	return s
}

func (s *GetGwpInventoryConstituteRequest) SetProductType(v int64) *GetGwpInventoryConstituteRequest {
	s.ProductType = &v
	return s
}

func (s *GetGwpInventoryConstituteRequest) Validate() error {
	return dara.Validate(s)
}
