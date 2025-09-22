// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEpdInventoryConstituteRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetEpdInventoryConstituteRequest
	GetCode() *string
	SetProductId(v int64) *GetEpdInventoryConstituteRequest
	GetProductId() *int64
	SetProductType(v int64) *GetEpdInventoryConstituteRequest
	GetProductType() *int64
}

type GetEpdInventoryConstituteRequest struct {
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

func (s GetEpdInventoryConstituteRequest) String() string {
	return dara.Prettify(s)
}

func (s GetEpdInventoryConstituteRequest) GoString() string {
	return s.String()
}

func (s *GetEpdInventoryConstituteRequest) GetCode() *string {
	return s.Code
}

func (s *GetEpdInventoryConstituteRequest) GetProductId() *int64 {
	return s.ProductId
}

func (s *GetEpdInventoryConstituteRequest) GetProductType() *int64 {
	return s.ProductType
}

func (s *GetEpdInventoryConstituteRequest) SetCode(v string) *GetEpdInventoryConstituteRequest {
	s.Code = &v
	return s
}

func (s *GetEpdInventoryConstituteRequest) SetProductId(v int64) *GetEpdInventoryConstituteRequest {
	s.ProductId = &v
	return s
}

func (s *GetEpdInventoryConstituteRequest) SetProductType(v int64) *GetEpdInventoryConstituteRequest {
	s.ProductType = &v
	return s
}

func (s *GetEpdInventoryConstituteRequest) Validate() error {
	return dara.Validate(s)
}
