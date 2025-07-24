// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPromotion interface {
	dara.Model
	String() string
	GoString() string
	SetProductCode(v string) *Promotion
	GetProductCode() *string
	SetPromotionDesc(v string) *Promotion
	GetPromotionDesc() *string
	SetPromotionName(v string) *Promotion
	GetPromotionName() *string
	SetPromotionOptionCode(v string) *Promotion
	GetPromotionOptionCode() *string
	SetPromotionOptionNo(v string) *Promotion
	GetPromotionOptionNo() *string
}

type Promotion struct {
	// 产品码。
	//
	// example:
	//
	// ecs
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// 优惠券描述。
	//
	// example:
	//
	// 5元优惠券（有效期至23年8月11日）
	PromotionDesc *string `json:"PromotionDesc,omitempty" xml:"PromotionDesc,omitempty"`
	// 优惠券名称。
	//
	// example:
	//
	// 5元优惠券
	PromotionName *string `json:"PromotionName,omitempty" xml:"PromotionName,omitempty"`
	// 优惠券码。
	//
	// example:
	//
	// youhui_quan
	PromotionOptionCode *string `json:"PromotionOptionCode,omitempty" xml:"PromotionOptionCode,omitempty"`
	// 优惠券号。
	//
	// This parameter is required.
	//
	// example:
	//
	// ABC123
	PromotionOptionNo *string `json:"PromotionOptionNo,omitempty" xml:"PromotionOptionNo,omitempty"`
}

func (s Promotion) String() string {
	return dara.Prettify(s)
}

func (s Promotion) GoString() string {
	return s.String()
}

func (s *Promotion) GetProductCode() *string {
	return s.ProductCode
}

func (s *Promotion) GetPromotionDesc() *string {
	return s.PromotionDesc
}

func (s *Promotion) GetPromotionName() *string {
	return s.PromotionName
}

func (s *Promotion) GetPromotionOptionCode() *string {
	return s.PromotionOptionCode
}

func (s *Promotion) GetPromotionOptionNo() *string {
	return s.PromotionOptionNo
}

func (s *Promotion) SetProductCode(v string) *Promotion {
	s.ProductCode = &v
	return s
}

func (s *Promotion) SetPromotionDesc(v string) *Promotion {
	s.PromotionDesc = &v
	return s
}

func (s *Promotion) SetPromotionName(v string) *Promotion {
	s.PromotionName = &v
	return s
}

func (s *Promotion) SetPromotionOptionCode(v string) *Promotion {
	s.PromotionOptionCode = &v
	return s
}

func (s *Promotion) SetPromotionOptionNo(v string) *Promotion {
	s.PromotionOptionNo = &v
	return s
}

func (s *Promotion) Validate() error {
	return dara.Validate(s)
}
