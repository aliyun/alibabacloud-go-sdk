// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPcrInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetPcrInfoRequest
	GetCode() *string
	SetProductId(v string) *GetPcrInfoRequest
	GetProductId() *string
	SetProductType(v int64) *GetPcrInfoRequest
	GetProductType() *int64
}

type GetPcrInfoRequest struct {
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
	ProductId *string `json:"productId,omitempty" xml:"productId,omitempty"`
	// Product type: 1 indicates that the carbon footprint of the product is requested, and 5 indicates that the carbon footprint of the supply chain is requested.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	ProductType *int64 `json:"productType,omitempty" xml:"productType,omitempty"`
}

func (s GetPcrInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPcrInfoRequest) GoString() string {
	return s.String()
}

func (s *GetPcrInfoRequest) GetCode() *string {
	return s.Code
}

func (s *GetPcrInfoRequest) GetProductId() *string {
	return s.ProductId
}

func (s *GetPcrInfoRequest) GetProductType() *int64 {
	return s.ProductType
}

func (s *GetPcrInfoRequest) SetCode(v string) *GetPcrInfoRequest {
	s.Code = &v
	return s
}

func (s *GetPcrInfoRequest) SetProductId(v string) *GetPcrInfoRequest {
	s.ProductId = &v
	return s
}

func (s *GetPcrInfoRequest) SetProductType(v int64) *GetPcrInfoRequest {
	s.ProductType = &v
	return s
}

func (s *GetPcrInfoRequest) Validate() error {
	return dara.Validate(s)
}
