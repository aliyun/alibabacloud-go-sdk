// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iProductSaleInfoQuery interface {
	dara.Model
	String() string
	GoString() string
	SetDistributorShopId(v string) *ProductSaleInfoQuery
	GetDistributorShopId() *string
	SetDivisionCode(v string) *ProductSaleInfoQuery
	GetDivisionCode() *string
}

type ProductSaleInfoQuery struct {
	// This parameter is required.
	//
	// example:
	//
	// 22000009
	DistributorShopId *string `json:"distributorShopId,omitempty" xml:"distributorShopId,omitempty"`
	// example:
	//
	// 110000
	DivisionCode *string `json:"divisionCode,omitempty" xml:"divisionCode,omitempty"`
}

func (s ProductSaleInfoQuery) String() string {
	return dara.Prettify(s)
}

func (s ProductSaleInfoQuery) GoString() string {
	return s.String()
}

func (s *ProductSaleInfoQuery) GetDistributorShopId() *string {
	return s.DistributorShopId
}

func (s *ProductSaleInfoQuery) GetDivisionCode() *string {
	return s.DivisionCode
}

func (s *ProductSaleInfoQuery) SetDistributorShopId(v string) *ProductSaleInfoQuery {
	s.DistributorShopId = &v
	return s
}

func (s *ProductSaleInfoQuery) SetDivisionCode(v string) *ProductSaleInfoQuery {
	s.DivisionCode = &v
	return s
}

func (s *ProductSaleInfoQuery) Validate() error {
	return dara.Validate(s)
}
