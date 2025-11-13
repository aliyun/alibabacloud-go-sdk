// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCreateBEClusterInquiryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetCreateBEClusterInquiryResponseBodyData) *GetCreateBEClusterInquiryResponseBody
	GetData() *GetCreateBEClusterInquiryResponseBodyData
	SetRequestId(v string) *GetCreateBEClusterInquiryResponseBody
	GetRequestId() *string
}

type GetCreateBEClusterInquiryResponseBody struct {
	// The information returned.
	Data *GetCreateBEClusterInquiryResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 06758CAB-1204-5852-A471-29C87D5C1D0F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetCreateBEClusterInquiryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCreateBEClusterInquiryResponseBody) GoString() string {
	return s.String()
}

func (s *GetCreateBEClusterInquiryResponseBody) GetData() *GetCreateBEClusterInquiryResponseBodyData {
	return s.Data
}

func (s *GetCreateBEClusterInquiryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCreateBEClusterInquiryResponseBody) SetData(v *GetCreateBEClusterInquiryResponseBodyData) *GetCreateBEClusterInquiryResponseBody {
	s.Data = v
	return s
}

func (s *GetCreateBEClusterInquiryResponseBody) SetRequestId(v string) *GetCreateBEClusterInquiryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCreateBEClusterInquiryResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetCreateBEClusterInquiryResponseBodyData struct {
	// The currency.
	//
	// example:
	//
	// CNY
	Currency           *string                                                        `json:"Currency,omitempty" xml:"Currency,omitempty"`
	OptionalPromotions []*GetCreateBEClusterInquiryResponseBodyDataOptionalPromotions `json:"OptionalPromotions,omitempty" xml:"OptionalPromotions,omitempty" type:"Repeated"`
	PricingRules       map[string]*string                                             `json:"PricingRules,omitempty" xml:"PricingRules,omitempty"`
	// The amount of money.
	//
	// example:
	//
	// 1.76
	TradeAmount *string `json:"TradeAmount,omitempty" xml:"TradeAmount,omitempty"`
}

func (s GetCreateBEClusterInquiryResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetCreateBEClusterInquiryResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetCreateBEClusterInquiryResponseBodyData) GetCurrency() *string {
	return s.Currency
}

func (s *GetCreateBEClusterInquiryResponseBodyData) GetOptionalPromotions() []*GetCreateBEClusterInquiryResponseBodyDataOptionalPromotions {
	return s.OptionalPromotions
}

func (s *GetCreateBEClusterInquiryResponseBodyData) GetPricingRules() map[string]*string {
	return s.PricingRules
}

func (s *GetCreateBEClusterInquiryResponseBodyData) GetTradeAmount() *string {
	return s.TradeAmount
}

func (s *GetCreateBEClusterInquiryResponseBodyData) SetCurrency(v string) *GetCreateBEClusterInquiryResponseBodyData {
	s.Currency = &v
	return s
}

func (s *GetCreateBEClusterInquiryResponseBodyData) SetOptionalPromotions(v []*GetCreateBEClusterInquiryResponseBodyDataOptionalPromotions) *GetCreateBEClusterInquiryResponseBodyData {
	s.OptionalPromotions = v
	return s
}

func (s *GetCreateBEClusterInquiryResponseBodyData) SetPricingRules(v map[string]*string) *GetCreateBEClusterInquiryResponseBodyData {
	s.PricingRules = v
	return s
}

func (s *GetCreateBEClusterInquiryResponseBodyData) SetTradeAmount(v string) *GetCreateBEClusterInquiryResponseBodyData {
	s.TradeAmount = &v
	return s
}

func (s *GetCreateBEClusterInquiryResponseBodyData) Validate() error {
	if s.OptionalPromotions != nil {
		for _, item := range s.OptionalPromotions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetCreateBEClusterInquiryResponseBodyDataOptionalPromotions struct {
	// example:
	//
	// 100
	CanPromFee *string `json:"CanPromFee,omitempty" xml:"CanPromFee,omitempty"`
	// example:
	//
	// 7afs9d
	OptionCode *string `json:"OptionCode,omitempty" xml:"OptionCode,omitempty"`
	// example:
	//
	// 通用优惠券可抵扣100
	PromotionDesc *string `json:"PromotionDesc,omitempty" xml:"PromotionDesc,omitempty"`
	// example:
	//
	// 通用优惠券
	PromotionName *string `json:"PromotionName,omitempty" xml:"PromotionName,omitempty"`
	// example:
	//
	// 775735400028
	PromotionOptionNo *string `json:"PromotionOptionNo,omitempty" xml:"PromotionOptionNo,omitempty"`
}

func (s GetCreateBEClusterInquiryResponseBodyDataOptionalPromotions) String() string {
	return dara.Prettify(s)
}

func (s GetCreateBEClusterInquiryResponseBodyDataOptionalPromotions) GoString() string {
	return s.String()
}

func (s *GetCreateBEClusterInquiryResponseBodyDataOptionalPromotions) GetCanPromFee() *string {
	return s.CanPromFee
}

func (s *GetCreateBEClusterInquiryResponseBodyDataOptionalPromotions) GetOptionCode() *string {
	return s.OptionCode
}

func (s *GetCreateBEClusterInquiryResponseBodyDataOptionalPromotions) GetPromotionDesc() *string {
	return s.PromotionDesc
}

func (s *GetCreateBEClusterInquiryResponseBodyDataOptionalPromotions) GetPromotionName() *string {
	return s.PromotionName
}

func (s *GetCreateBEClusterInquiryResponseBodyDataOptionalPromotions) GetPromotionOptionNo() *string {
	return s.PromotionOptionNo
}

func (s *GetCreateBEClusterInquiryResponseBodyDataOptionalPromotions) SetCanPromFee(v string) *GetCreateBEClusterInquiryResponseBodyDataOptionalPromotions {
	s.CanPromFee = &v
	return s
}

func (s *GetCreateBEClusterInquiryResponseBodyDataOptionalPromotions) SetOptionCode(v string) *GetCreateBEClusterInquiryResponseBodyDataOptionalPromotions {
	s.OptionCode = &v
	return s
}

func (s *GetCreateBEClusterInquiryResponseBodyDataOptionalPromotions) SetPromotionDesc(v string) *GetCreateBEClusterInquiryResponseBodyDataOptionalPromotions {
	s.PromotionDesc = &v
	return s
}

func (s *GetCreateBEClusterInquiryResponseBodyDataOptionalPromotions) SetPromotionName(v string) *GetCreateBEClusterInquiryResponseBodyDataOptionalPromotions {
	s.PromotionName = &v
	return s
}

func (s *GetCreateBEClusterInquiryResponseBodyDataOptionalPromotions) SetPromotionOptionNo(v string) *GetCreateBEClusterInquiryResponseBodyDataOptionalPromotions {
	s.PromotionOptionNo = &v
	return s
}

func (s *GetCreateBEClusterInquiryResponseBodyDataOptionalPromotions) Validate() error {
	return dara.Validate(s)
}
