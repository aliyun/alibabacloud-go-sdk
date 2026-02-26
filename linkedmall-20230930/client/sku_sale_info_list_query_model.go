// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSkuSaleInfoListQuery interface {
	dara.Model
	String() string
	GoString() string
	SetDivisionCode(v string) *SkuSaleInfoListQuery
	GetDivisionCode() *string
	SetPurchaserId(v string) *SkuSaleInfoListQuery
	GetPurchaserId() *string
	SetSkuQueryParams(v []*SkuQueryParam) *SkuSaleInfoListQuery
	GetSkuQueryParams() []*SkuQueryParam
}

type SkuSaleInfoListQuery struct {
	// example:
	//
	// 110000
	DivisionCode *string `json:"divisionCode,omitempty" xml:"divisionCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 21000017
	PurchaserId *string `json:"purchaserId,omitempty" xml:"purchaserId,omitempty"`
	// This parameter is required.
	SkuQueryParams []*SkuQueryParam `json:"skuQueryParams,omitempty" xml:"skuQueryParams,omitempty" type:"Repeated"`
}

func (s SkuSaleInfoListQuery) String() string {
	return dara.Prettify(s)
}

func (s SkuSaleInfoListQuery) GoString() string {
	return s.String()
}

func (s *SkuSaleInfoListQuery) GetDivisionCode() *string {
	return s.DivisionCode
}

func (s *SkuSaleInfoListQuery) GetPurchaserId() *string {
	return s.PurchaserId
}

func (s *SkuSaleInfoListQuery) GetSkuQueryParams() []*SkuQueryParam {
	return s.SkuQueryParams
}

func (s *SkuSaleInfoListQuery) SetDivisionCode(v string) *SkuSaleInfoListQuery {
	s.DivisionCode = &v
	return s
}

func (s *SkuSaleInfoListQuery) SetPurchaserId(v string) *SkuSaleInfoListQuery {
	s.PurchaserId = &v
	return s
}

func (s *SkuSaleInfoListQuery) SetSkuQueryParams(v []*SkuQueryParam) *SkuSaleInfoListQuery {
	s.SkuQueryParams = v
	return s
}

func (s *SkuSaleInfoListQuery) Validate() error {
	if s.SkuQueryParams != nil {
		for _, item := range s.SkuQueryParams {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
