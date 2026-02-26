// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGood interface {
	dara.Model
	String() string
	GoString() string
	SetGoodName(v string) *Good
	GetGoodName() *string
	SetProductId(v string) *Good
	GetProductId() *string
	SetQuantity(v int32) *Good
	GetQuantity() *int32
	SetSkuId(v string) *Good
	GetSkuId() *string
	SetSkuTitle(v string) *Good
	GetSkuTitle() *string
}

type Good struct {
	GoodName *string `json:"goodName,omitempty" xml:"goodName,omitempty"`
	// example:
	//
	// 6600****6736
	ProductId *string `json:"productId,omitempty" xml:"productId,omitempty"`
	// example:
	//
	// 1
	Quantity *int32  `json:"quantity,omitempty" xml:"quantity,omitempty"`
	SkuId    *string `json:"skuId,omitempty" xml:"skuId,omitempty"`
	SkuTitle *string `json:"skuTitle,omitempty" xml:"skuTitle,omitempty"`
}

func (s Good) String() string {
	return dara.Prettify(s)
}

func (s Good) GoString() string {
	return s.String()
}

func (s *Good) GetGoodName() *string {
	return s.GoodName
}

func (s *Good) GetProductId() *string {
	return s.ProductId
}

func (s *Good) GetQuantity() *int32 {
	return s.Quantity
}

func (s *Good) GetSkuId() *string {
	return s.SkuId
}

func (s *Good) GetSkuTitle() *string {
	return s.SkuTitle
}

func (s *Good) SetGoodName(v string) *Good {
	s.GoodName = &v
	return s
}

func (s *Good) SetProductId(v string) *Good {
	s.ProductId = &v
	return s
}

func (s *Good) SetQuantity(v int32) *Good {
	s.Quantity = &v
	return s
}

func (s *Good) SetSkuId(v string) *Good {
	s.SkuId = &v
	return s
}

func (s *Good) SetSkuTitle(v string) *Good {
	s.SkuTitle = &v
	return s
}

func (s *Good) Validate() error {
	return dara.Validate(s)
}
