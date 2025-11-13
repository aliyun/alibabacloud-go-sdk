// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetModifyBEClusterInquiryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetModifyBEClusterInquiryResponseBodyData) *GetModifyBEClusterInquiryResponseBody
	GetData() *GetModifyBEClusterInquiryResponseBodyData
	SetRequestId(v string) *GetModifyBEClusterInquiryResponseBody
	GetRequestId() *string
}

type GetModifyBEClusterInquiryResponseBody struct {
	// The information returned.
	Data *GetModifyBEClusterInquiryResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 06758CAB-1204-5852-A471-29C87D5C1D0F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetModifyBEClusterInquiryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetModifyBEClusterInquiryResponseBody) GoString() string {
	return s.String()
}

func (s *GetModifyBEClusterInquiryResponseBody) GetData() *GetModifyBEClusterInquiryResponseBodyData {
	return s.Data
}

func (s *GetModifyBEClusterInquiryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetModifyBEClusterInquiryResponseBody) SetData(v *GetModifyBEClusterInquiryResponseBodyData) *GetModifyBEClusterInquiryResponseBody {
	s.Data = v
	return s
}

func (s *GetModifyBEClusterInquiryResponseBody) SetRequestId(v string) *GetModifyBEClusterInquiryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetModifyBEClusterInquiryResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetModifyBEClusterInquiryResponseBodyData struct {
	// The currency.
	//
	// example:
	//
	// CNY
	Currency           *string                                                        `json:"Currency,omitempty" xml:"Currency,omitempty"`
	OptionalPromotions []*GetModifyBEClusterInquiryResponseBodyDataOptionalPromotions `json:"OptionalPromotions,omitempty" xml:"OptionalPromotions,omitempty" type:"Repeated"`
	PricingRules       map[string]*string                                             `json:"PricingRules,omitempty" xml:"PricingRules,omitempty"`
	// The estimated refund amount when the subscription cluster of a subscription instance is changed to a pay-as-you-go cluster.
	//
	// example:
	//
	// -100
	RefundAmount *string `json:"RefundAmount,omitempty" xml:"RefundAmount,omitempty"`
	// The amount of money.
	//
	// example:
	//
	// 1.76
	TradeAmount *string `json:"TradeAmount,omitempty" xml:"TradeAmount,omitempty"`
}

func (s GetModifyBEClusterInquiryResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetModifyBEClusterInquiryResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetModifyBEClusterInquiryResponseBodyData) GetCurrency() *string {
	return s.Currency
}

func (s *GetModifyBEClusterInquiryResponseBodyData) GetOptionalPromotions() []*GetModifyBEClusterInquiryResponseBodyDataOptionalPromotions {
	return s.OptionalPromotions
}

func (s *GetModifyBEClusterInquiryResponseBodyData) GetPricingRules() map[string]*string {
	return s.PricingRules
}

func (s *GetModifyBEClusterInquiryResponseBodyData) GetRefundAmount() *string {
	return s.RefundAmount
}

func (s *GetModifyBEClusterInquiryResponseBodyData) GetTradeAmount() *string {
	return s.TradeAmount
}

func (s *GetModifyBEClusterInquiryResponseBodyData) SetCurrency(v string) *GetModifyBEClusterInquiryResponseBodyData {
	s.Currency = &v
	return s
}

func (s *GetModifyBEClusterInquiryResponseBodyData) SetOptionalPromotions(v []*GetModifyBEClusterInquiryResponseBodyDataOptionalPromotions) *GetModifyBEClusterInquiryResponseBodyData {
	s.OptionalPromotions = v
	return s
}

func (s *GetModifyBEClusterInquiryResponseBodyData) SetPricingRules(v map[string]*string) *GetModifyBEClusterInquiryResponseBodyData {
	s.PricingRules = v
	return s
}

func (s *GetModifyBEClusterInquiryResponseBodyData) SetRefundAmount(v string) *GetModifyBEClusterInquiryResponseBodyData {
	s.RefundAmount = &v
	return s
}

func (s *GetModifyBEClusterInquiryResponseBodyData) SetTradeAmount(v string) *GetModifyBEClusterInquiryResponseBodyData {
	s.TradeAmount = &v
	return s
}

func (s *GetModifyBEClusterInquiryResponseBodyData) Validate() error {
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

type GetModifyBEClusterInquiryResponseBodyDataOptionalPromotions struct {
	// example:
	//
	// 100
	CanPromFee *string `json:"CanPromFee,omitempty" xml:"CanPromFee,omitempty"`
	// example:
	//
	// 789asdf
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
	// 732211480132
	PromotionOptionNo *string `json:"PromotionOptionNo,omitempty" xml:"PromotionOptionNo,omitempty"`
}

func (s GetModifyBEClusterInquiryResponseBodyDataOptionalPromotions) String() string {
	return dara.Prettify(s)
}

func (s GetModifyBEClusterInquiryResponseBodyDataOptionalPromotions) GoString() string {
	return s.String()
}

func (s *GetModifyBEClusterInquiryResponseBodyDataOptionalPromotions) GetCanPromFee() *string {
	return s.CanPromFee
}

func (s *GetModifyBEClusterInquiryResponseBodyDataOptionalPromotions) GetOptionCode() *string {
	return s.OptionCode
}

func (s *GetModifyBEClusterInquiryResponseBodyDataOptionalPromotions) GetPromotionDesc() *string {
	return s.PromotionDesc
}

func (s *GetModifyBEClusterInquiryResponseBodyDataOptionalPromotions) GetPromotionName() *string {
	return s.PromotionName
}

func (s *GetModifyBEClusterInquiryResponseBodyDataOptionalPromotions) GetPromotionOptionNo() *string {
	return s.PromotionOptionNo
}

func (s *GetModifyBEClusterInquiryResponseBodyDataOptionalPromotions) SetCanPromFee(v string) *GetModifyBEClusterInquiryResponseBodyDataOptionalPromotions {
	s.CanPromFee = &v
	return s
}

func (s *GetModifyBEClusterInquiryResponseBodyDataOptionalPromotions) SetOptionCode(v string) *GetModifyBEClusterInquiryResponseBodyDataOptionalPromotions {
	s.OptionCode = &v
	return s
}

func (s *GetModifyBEClusterInquiryResponseBodyDataOptionalPromotions) SetPromotionDesc(v string) *GetModifyBEClusterInquiryResponseBodyDataOptionalPromotions {
	s.PromotionDesc = &v
	return s
}

func (s *GetModifyBEClusterInquiryResponseBodyDataOptionalPromotions) SetPromotionName(v string) *GetModifyBEClusterInquiryResponseBodyDataOptionalPromotions {
	s.PromotionName = &v
	return s
}

func (s *GetModifyBEClusterInquiryResponseBodyDataOptionalPromotions) SetPromotionOptionNo(v string) *GetModifyBEClusterInquiryResponseBodyDataOptionalPromotions {
	s.PromotionOptionNo = &v
	return s
}

func (s *GetModifyBEClusterInquiryResponseBodyDataOptionalPromotions) Validate() error {
	return dara.Validate(s)
}
