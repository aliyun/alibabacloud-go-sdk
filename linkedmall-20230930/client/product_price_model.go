// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iProductPrice interface {
	dara.Model
	String() string
	GoString() string
	SetFundAmountMoney(v string) *ProductPrice
	GetFundAmountMoney() *string
}

type ProductPrice struct {
	// example:
	//
	// 120
	FundAmountMoney *string `json:"fundAmountMoney,omitempty" xml:"fundAmountMoney,omitempty"`
}

func (s ProductPrice) String() string {
	return dara.Prettify(s)
}

func (s ProductPrice) GoString() string {
	return s.String()
}

func (s *ProductPrice) GetFundAmountMoney() *string {
	return s.FundAmountMoney
}

func (s *ProductPrice) SetFundAmountMoney(v string) *ProductPrice {
	s.FundAmountMoney = &v
	return s
}

func (s *ProductPrice) Validate() error {
	return dara.Validate(s)
}
