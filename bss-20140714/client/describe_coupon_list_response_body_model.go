// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCouponListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCoupons(v *DescribeCouponListResponseBodyCoupons) *DescribeCouponListResponseBody
	GetCoupons() *DescribeCouponListResponseBodyCoupons
	SetRequestId(v string) *DescribeCouponListResponseBody
	GetRequestId() *string
}

type DescribeCouponListResponseBody struct {
	Coupons   *DescribeCouponListResponseBodyCoupons `json:"Coupons,omitempty" xml:"Coupons,omitempty" type:"Struct"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeCouponListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCouponListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCouponListResponseBody) GetCoupons() *DescribeCouponListResponseBodyCoupons {
	return s.Coupons
}

func (s *DescribeCouponListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCouponListResponseBody) SetCoupons(v *DescribeCouponListResponseBodyCoupons) *DescribeCouponListResponseBody {
	s.Coupons = v
	return s
}

func (s *DescribeCouponListResponseBody) SetRequestId(v string) *DescribeCouponListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCouponListResponseBody) Validate() error {
	if s.Coupons != nil {
		if err := s.Coupons.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCouponListResponseBodyCoupons struct {
	Coupon []*DescribeCouponListResponseBodyCouponsCoupon `json:"Coupon,omitempty" xml:"Coupon,omitempty" type:"Repeated"`
}

func (s DescribeCouponListResponseBodyCoupons) String() string {
	return dara.Prettify(s)
}

func (s DescribeCouponListResponseBodyCoupons) GoString() string {
	return s.String()
}

func (s *DescribeCouponListResponseBodyCoupons) GetCoupon() []*DescribeCouponListResponseBodyCouponsCoupon {
	return s.Coupon
}

func (s *DescribeCouponListResponseBodyCoupons) SetCoupon(v []*DescribeCouponListResponseBodyCouponsCoupon) *DescribeCouponListResponseBodyCoupons {
	s.Coupon = v
	return s
}

func (s *DescribeCouponListResponseBodyCoupons) Validate() error {
	if s.Coupon != nil {
		for _, item := range s.Coupon {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCouponListResponseBodyCouponsCoupon struct {
	Application      *string                                                  `json:"Application,omitempty" xml:"Application,omitempty"`
	BalanceAmount    *string                                                  `json:"BalanceAmount,omitempty" xml:"BalanceAmount,omitempty"`
	CouponNumber     *string                                                  `json:"CouponNumber,omitempty" xml:"CouponNumber,omitempty"`
	CouponTemplateId *int64                                                   `json:"CouponTemplateId,omitempty" xml:"CouponTemplateId,omitempty"`
	CreationTime     *string                                                  `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	DeliveryTime     *string                                                  `json:"DeliveryTime,omitempty" xml:"DeliveryTime,omitempty"`
	Description      *string                                                  `json:"Description,omitempty" xml:"Description,omitempty"`
	ExpiredAmount    *string                                                  `json:"ExpiredAmount,omitempty" xml:"ExpiredAmount,omitempty"`
	ExpiredTime      *string                                                  `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	FrozenAmount     *string                                                  `json:"FrozenAmount,omitempty" xml:"FrozenAmount,omitempty"`
	ModificationTime *string                                                  `json:"ModificationTime,omitempty" xml:"ModificationTime,omitempty"`
	PriceLimit       *string                                                  `json:"PriceLimit,omitempty" xml:"PriceLimit,omitempty"`
	ProductCodes     *DescribeCouponListResponseBodyCouponsCouponProductCodes `json:"ProductCodes,omitempty" xml:"ProductCodes,omitempty" type:"Struct"`
	Status           *string                                                  `json:"Status,omitempty" xml:"Status,omitempty"`
	TotalAmount      *string                                                  `json:"TotalAmount,omitempty" xml:"TotalAmount,omitempty"`
	TradeTypes       *DescribeCouponListResponseBodyCouponsCouponTradeTypes   `json:"TradeTypes,omitempty" xml:"TradeTypes,omitempty" type:"Struct"`
}

func (s DescribeCouponListResponseBodyCouponsCoupon) String() string {
	return dara.Prettify(s)
}

func (s DescribeCouponListResponseBodyCouponsCoupon) GoString() string {
	return s.String()
}

func (s *DescribeCouponListResponseBodyCouponsCoupon) GetApplication() *string {
	return s.Application
}

func (s *DescribeCouponListResponseBodyCouponsCoupon) GetBalanceAmount() *string {
	return s.BalanceAmount
}

func (s *DescribeCouponListResponseBodyCouponsCoupon) GetCouponNumber() *string {
	return s.CouponNumber
}

func (s *DescribeCouponListResponseBodyCouponsCoupon) GetCouponTemplateId() *int64 {
	return s.CouponTemplateId
}

func (s *DescribeCouponListResponseBodyCouponsCoupon) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeCouponListResponseBodyCouponsCoupon) GetDeliveryTime() *string {
	return s.DeliveryTime
}

func (s *DescribeCouponListResponseBodyCouponsCoupon) GetDescription() *string {
	return s.Description
}

func (s *DescribeCouponListResponseBodyCouponsCoupon) GetExpiredAmount() *string {
	return s.ExpiredAmount
}

func (s *DescribeCouponListResponseBodyCouponsCoupon) GetExpiredTime() *string {
	return s.ExpiredTime
}

func (s *DescribeCouponListResponseBodyCouponsCoupon) GetFrozenAmount() *string {
	return s.FrozenAmount
}

func (s *DescribeCouponListResponseBodyCouponsCoupon) GetModificationTime() *string {
	return s.ModificationTime
}

func (s *DescribeCouponListResponseBodyCouponsCoupon) GetPriceLimit() *string {
	return s.PriceLimit
}

func (s *DescribeCouponListResponseBodyCouponsCoupon) GetProductCodes() *DescribeCouponListResponseBodyCouponsCouponProductCodes {
	return s.ProductCodes
}

func (s *DescribeCouponListResponseBodyCouponsCoupon) GetStatus() *string {
	return s.Status
}

func (s *DescribeCouponListResponseBodyCouponsCoupon) GetTotalAmount() *string {
	return s.TotalAmount
}

func (s *DescribeCouponListResponseBodyCouponsCoupon) GetTradeTypes() *DescribeCouponListResponseBodyCouponsCouponTradeTypes {
	return s.TradeTypes
}

func (s *DescribeCouponListResponseBodyCouponsCoupon) SetApplication(v string) *DescribeCouponListResponseBodyCouponsCoupon {
	s.Application = &v
	return s
}

func (s *DescribeCouponListResponseBodyCouponsCoupon) SetBalanceAmount(v string) *DescribeCouponListResponseBodyCouponsCoupon {
	s.BalanceAmount = &v
	return s
}

func (s *DescribeCouponListResponseBodyCouponsCoupon) SetCouponNumber(v string) *DescribeCouponListResponseBodyCouponsCoupon {
	s.CouponNumber = &v
	return s
}

func (s *DescribeCouponListResponseBodyCouponsCoupon) SetCouponTemplateId(v int64) *DescribeCouponListResponseBodyCouponsCoupon {
	s.CouponTemplateId = &v
	return s
}

func (s *DescribeCouponListResponseBodyCouponsCoupon) SetCreationTime(v string) *DescribeCouponListResponseBodyCouponsCoupon {
	s.CreationTime = &v
	return s
}

func (s *DescribeCouponListResponseBodyCouponsCoupon) SetDeliveryTime(v string) *DescribeCouponListResponseBodyCouponsCoupon {
	s.DeliveryTime = &v
	return s
}

func (s *DescribeCouponListResponseBodyCouponsCoupon) SetDescription(v string) *DescribeCouponListResponseBodyCouponsCoupon {
	s.Description = &v
	return s
}

func (s *DescribeCouponListResponseBodyCouponsCoupon) SetExpiredAmount(v string) *DescribeCouponListResponseBodyCouponsCoupon {
	s.ExpiredAmount = &v
	return s
}

func (s *DescribeCouponListResponseBodyCouponsCoupon) SetExpiredTime(v string) *DescribeCouponListResponseBodyCouponsCoupon {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeCouponListResponseBodyCouponsCoupon) SetFrozenAmount(v string) *DescribeCouponListResponseBodyCouponsCoupon {
	s.FrozenAmount = &v
	return s
}

func (s *DescribeCouponListResponseBodyCouponsCoupon) SetModificationTime(v string) *DescribeCouponListResponseBodyCouponsCoupon {
	s.ModificationTime = &v
	return s
}

func (s *DescribeCouponListResponseBodyCouponsCoupon) SetPriceLimit(v string) *DescribeCouponListResponseBodyCouponsCoupon {
	s.PriceLimit = &v
	return s
}

func (s *DescribeCouponListResponseBodyCouponsCoupon) SetProductCodes(v *DescribeCouponListResponseBodyCouponsCouponProductCodes) *DescribeCouponListResponseBodyCouponsCoupon {
	s.ProductCodes = v
	return s
}

func (s *DescribeCouponListResponseBodyCouponsCoupon) SetStatus(v string) *DescribeCouponListResponseBodyCouponsCoupon {
	s.Status = &v
	return s
}

func (s *DescribeCouponListResponseBodyCouponsCoupon) SetTotalAmount(v string) *DescribeCouponListResponseBodyCouponsCoupon {
	s.TotalAmount = &v
	return s
}

func (s *DescribeCouponListResponseBodyCouponsCoupon) SetTradeTypes(v *DescribeCouponListResponseBodyCouponsCouponTradeTypes) *DescribeCouponListResponseBodyCouponsCoupon {
	s.TradeTypes = v
	return s
}

func (s *DescribeCouponListResponseBodyCouponsCoupon) Validate() error {
	if s.ProductCodes != nil {
		if err := s.ProductCodes.Validate(); err != nil {
			return err
		}
	}
	if s.TradeTypes != nil {
		if err := s.TradeTypes.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCouponListResponseBodyCouponsCouponProductCodes struct {
	ProductCode []*string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty" type:"Repeated"`
}

func (s DescribeCouponListResponseBodyCouponsCouponProductCodes) String() string {
	return dara.Prettify(s)
}

func (s DescribeCouponListResponseBodyCouponsCouponProductCodes) GoString() string {
	return s.String()
}

func (s *DescribeCouponListResponseBodyCouponsCouponProductCodes) GetProductCode() []*string {
	return s.ProductCode
}

func (s *DescribeCouponListResponseBodyCouponsCouponProductCodes) SetProductCode(v []*string) *DescribeCouponListResponseBodyCouponsCouponProductCodes {
	s.ProductCode = v
	return s
}

func (s *DescribeCouponListResponseBodyCouponsCouponProductCodes) Validate() error {
	return dara.Validate(s)
}

type DescribeCouponListResponseBodyCouponsCouponTradeTypes struct {
	TradeType []*string `json:"TradeType,omitempty" xml:"TradeType,omitempty" type:"Repeated"`
}

func (s DescribeCouponListResponseBodyCouponsCouponTradeTypes) String() string {
	return dara.Prettify(s)
}

func (s DescribeCouponListResponseBodyCouponsCouponTradeTypes) GoString() string {
	return s.String()
}

func (s *DescribeCouponListResponseBodyCouponsCouponTradeTypes) GetTradeType() []*string {
	return s.TradeType
}

func (s *DescribeCouponListResponseBodyCouponsCouponTradeTypes) SetTradeType(v []*string) *DescribeCouponListResponseBodyCouponsCouponTradeTypes {
	s.TradeType = v
	return s
}

func (s *DescribeCouponListResponseBodyCouponsCouponTradeTypes) Validate() error {
	return dara.Validate(s)
}
