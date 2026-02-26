// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iProductQuery interface {
	dara.Model
	String() string
	GoString() string
	SetDistributorShopId(v string) *ProductQuery
	GetDistributorShopId() *string
	SetDivisionCode(v string) *ProductQuery
	GetDivisionCode() *string
}

type ProductQuery struct {
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

func (s ProductQuery) String() string {
	return dara.Prettify(s)
}

func (s ProductQuery) GoString() string {
	return s.String()
}

func (s *ProductQuery) GetDistributorShopId() *string {
	return s.DistributorShopId
}

func (s *ProductQuery) GetDivisionCode() *string {
	return s.DivisionCode
}

func (s *ProductQuery) SetDistributorShopId(v string) *ProductQuery {
	s.DistributorShopId = &v
	return s
}

func (s *ProductQuery) SetDivisionCode(v string) *ProductQuery {
	s.DivisionCode = &v
	return s
}

func (s *ProductQuery) Validate() error {
	return dara.Validate(s)
}
