// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourcePackagePriceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetResourcePackagePriceResponseBody
	GetCode() *string
	SetData(v *GetResourcePackagePriceResponseBodyData) *GetResourcePackagePriceResponseBody
	GetData() *GetResourcePackagePriceResponseBodyData
	SetMessage(v string) *GetResourcePackagePriceResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetResourcePackagePriceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetResourcePackagePriceResponseBody
	GetSuccess() *bool
}

type GetResourcePackagePriceResponseBody struct {
	// The status code.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *GetResourcePackagePriceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The message returned.
	//
	// example:
	//
	// Successful!
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// BBEF51A3-E933-4F40-A534-C673CBDB9C80
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetResourcePackagePriceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetResourcePackagePriceResponseBody) GoString() string {
	return s.String()
}

func (s *GetResourcePackagePriceResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetResourcePackagePriceResponseBody) GetData() *GetResourcePackagePriceResponseBodyData {
	return s.Data
}

func (s *GetResourcePackagePriceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetResourcePackagePriceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetResourcePackagePriceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetResourcePackagePriceResponseBody) SetCode(v string) *GetResourcePackagePriceResponseBody {
	s.Code = &v
	return s
}

func (s *GetResourcePackagePriceResponseBody) SetData(v *GetResourcePackagePriceResponseBodyData) *GetResourcePackagePriceResponseBody {
	s.Data = v
	return s
}

func (s *GetResourcePackagePriceResponseBody) SetMessage(v string) *GetResourcePackagePriceResponseBody {
	s.Message = &v
	return s
}

func (s *GetResourcePackagePriceResponseBody) SetRequestId(v string) *GetResourcePackagePriceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetResourcePackagePriceResponseBody) SetSuccess(v bool) *GetResourcePackagePriceResponseBody {
	s.Success = &v
	return s
}

func (s *GetResourcePackagePriceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetResourcePackagePriceResponseBodyData struct {
	// The type of the currency.
	//
	// example:
	//
	// CNY
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// The discounted amount. Unit: CNY.
	//
	// example:
	//
	// 215040
	DiscountPrice *float32 `json:"DiscountPrice,omitempty" xml:"DiscountPrice,omitempty"`
	// The original price. Unit: CNY.
	//
	// example:
	//
	// 1290240
	OriginalPrice *float32 `json:"OriginalPrice,omitempty" xml:"OriginalPrice,omitempty"`
	// The details of the discount.
	Promotions *GetResourcePackagePriceResponseBodyDataPromotions `json:"Promotions,omitempty" xml:"Promotions,omitempty" type:"Struct"`
	// The price at which the transaction is made. Unit: CNY.
	//
	// example:
	//
	// 1075200
	TradePrice *float32 `json:"TradePrice,omitempty" xml:"TradePrice,omitempty"`
}

func (s GetResourcePackagePriceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetResourcePackagePriceResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetResourcePackagePriceResponseBodyData) GetCurrency() *string {
	return s.Currency
}

func (s *GetResourcePackagePriceResponseBodyData) GetDiscountPrice() *float32 {
	return s.DiscountPrice
}

func (s *GetResourcePackagePriceResponseBodyData) GetOriginalPrice() *float32 {
	return s.OriginalPrice
}

func (s *GetResourcePackagePriceResponseBodyData) GetPromotions() *GetResourcePackagePriceResponseBodyDataPromotions {
	return s.Promotions
}

func (s *GetResourcePackagePriceResponseBodyData) GetTradePrice() *float32 {
	return s.TradePrice
}

func (s *GetResourcePackagePriceResponseBodyData) SetCurrency(v string) *GetResourcePackagePriceResponseBodyData {
	s.Currency = &v
	return s
}

func (s *GetResourcePackagePriceResponseBodyData) SetDiscountPrice(v float32) *GetResourcePackagePriceResponseBodyData {
	s.DiscountPrice = &v
	return s
}

func (s *GetResourcePackagePriceResponseBodyData) SetOriginalPrice(v float32) *GetResourcePackagePriceResponseBodyData {
	s.OriginalPrice = &v
	return s
}

func (s *GetResourcePackagePriceResponseBodyData) SetPromotions(v *GetResourcePackagePriceResponseBodyDataPromotions) *GetResourcePackagePriceResponseBodyData {
	s.Promotions = v
	return s
}

func (s *GetResourcePackagePriceResponseBodyData) SetTradePrice(v float32) *GetResourcePackagePriceResponseBodyData {
	s.TradePrice = &v
	return s
}

func (s *GetResourcePackagePriceResponseBodyData) Validate() error {
	if s.Promotions != nil {
		if err := s.Promotions.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetResourcePackagePriceResponseBodyDataPromotions struct {
	Promotion []*GetResourcePackagePriceResponseBodyDataPromotionsPromotion `json:"Promotion,omitempty" xml:"Promotion,omitempty" type:"Repeated"`
}

func (s GetResourcePackagePriceResponseBodyDataPromotions) String() string {
	return dara.Prettify(s)
}

func (s GetResourcePackagePriceResponseBodyDataPromotions) GoString() string {
	return s.String()
}

func (s *GetResourcePackagePriceResponseBodyDataPromotions) GetPromotion() []*GetResourcePackagePriceResponseBodyDataPromotionsPromotion {
	return s.Promotion
}

func (s *GetResourcePackagePriceResponseBodyDataPromotions) SetPromotion(v []*GetResourcePackagePriceResponseBodyDataPromotionsPromotion) *GetResourcePackagePriceResponseBodyDataPromotions {
	s.Promotion = v
	return s
}

func (s *GetResourcePackagePriceResponseBodyDataPromotions) Validate() error {
	if s.Promotion != nil {
		for _, item := range s.Promotion {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetResourcePackagePriceResponseBodyDataPromotionsPromotion struct {
	// The ID of the promotion.
	//
	// example:
	//
	// 1000680914
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The description of the discount.
	//
	// example:
	//
	// A discount of 17% is offered if you purchase a resource plan for six months.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetResourcePackagePriceResponseBodyDataPromotionsPromotion) String() string {
	return dara.Prettify(s)
}

func (s GetResourcePackagePriceResponseBodyDataPromotionsPromotion) GoString() string {
	return s.String()
}

func (s *GetResourcePackagePriceResponseBodyDataPromotionsPromotion) GetId() *int64 {
	return s.Id
}

func (s *GetResourcePackagePriceResponseBodyDataPromotionsPromotion) GetName() *string {
	return s.Name
}

func (s *GetResourcePackagePriceResponseBodyDataPromotionsPromotion) SetId(v int64) *GetResourcePackagePriceResponseBodyDataPromotionsPromotion {
	s.Id = &v
	return s
}

func (s *GetResourcePackagePriceResponseBodyDataPromotionsPromotion) SetName(v string) *GetResourcePackagePriceResponseBodyDataPromotionsPromotion {
	s.Name = &v
	return s
}

func (s *GetResourcePackagePriceResponseBodyDataPromotionsPromotion) Validate() error {
	return dara.Validate(s)
}
