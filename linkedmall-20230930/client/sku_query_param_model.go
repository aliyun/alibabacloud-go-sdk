// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSkuQueryParam interface {
	dara.Model
	String() string
	GoString() string
	SetBuyAmount(v int32) *SkuQueryParam
	GetBuyAmount() *int32
	SetProductId(v string) *SkuQueryParam
	GetProductId() *string
	SetSkuId(v string) *SkuQueryParam
	GetSkuId() *string
}

type SkuQueryParam struct {
	// example:
	//
	// 1
	BuyAmount *int32 `json:"buyAmount,omitempty" xml:"buyAmount,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 660460842235822080
	ProductId *string `json:"productId,omitempty" xml:"productId,omitempty"`
	// skuid
	//
	// This parameter is required.
	//
	// example:
	//
	// 660460842235822081
	SkuId *string `json:"skuId,omitempty" xml:"skuId,omitempty"`
}

func (s SkuQueryParam) String() string {
	return dara.Prettify(s)
}

func (s SkuQueryParam) GoString() string {
	return s.String()
}

func (s *SkuQueryParam) GetBuyAmount() *int32 {
	return s.BuyAmount
}

func (s *SkuQueryParam) GetProductId() *string {
	return s.ProductId
}

func (s *SkuQueryParam) GetSkuId() *string {
	return s.SkuId
}

func (s *SkuQueryParam) SetBuyAmount(v int32) *SkuQueryParam {
	s.BuyAmount = &v
	return s
}

func (s *SkuQueryParam) SetProductId(v string) *SkuQueryParam {
	s.ProductId = &v
	return s
}

func (s *SkuQueryParam) SetSkuId(v string) *SkuQueryParam {
	s.SkuId = &v
	return s
}

func (s *SkuQueryParam) Validate() error {
	return dara.Validate(s)
}
