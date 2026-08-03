// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGlobalHotelValidatePriceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GlobalHotelValidatePriceResponseBodyData) *GlobalHotelValidatePriceResponseBody
	GetData() *GlobalHotelValidatePriceResponseBodyData
	SetErrorCode(v string) *GlobalHotelValidatePriceResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *GlobalHotelValidatePriceResponseBody
	GetErrorMsg() *string
	SetRequestId(v string) *GlobalHotelValidatePriceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GlobalHotelValidatePriceResponseBody
	GetSuccess() *bool
	SetTracerId(v string) *GlobalHotelValidatePriceResponseBody
	GetTracerId() *string
}

type GlobalHotelValidatePriceResponseBody struct {
	Data *GlobalHotelValidatePriceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// CreateOrderFailed
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// 创建订单失败
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// example:
	//
	// 260E4F99-983D-1919-834C-5C42E98E5B2B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success  *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	TracerId *string `json:"TracerId,omitempty" xml:"TracerId,omitempty"`
}

func (s GlobalHotelValidatePriceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GlobalHotelValidatePriceResponseBody) GoString() string {
	return s.String()
}

func (s *GlobalHotelValidatePriceResponseBody) GetData() *GlobalHotelValidatePriceResponseBodyData {
	return s.Data
}

func (s *GlobalHotelValidatePriceResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GlobalHotelValidatePriceResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *GlobalHotelValidatePriceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GlobalHotelValidatePriceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GlobalHotelValidatePriceResponseBody) GetTracerId() *string {
	return s.TracerId
}

func (s *GlobalHotelValidatePriceResponseBody) SetData(v *GlobalHotelValidatePriceResponseBodyData) *GlobalHotelValidatePriceResponseBody {
	s.Data = v
	return s
}

func (s *GlobalHotelValidatePriceResponseBody) SetErrorCode(v string) *GlobalHotelValidatePriceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GlobalHotelValidatePriceResponseBody) SetErrorMsg(v string) *GlobalHotelValidatePriceResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *GlobalHotelValidatePriceResponseBody) SetRequestId(v string) *GlobalHotelValidatePriceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GlobalHotelValidatePriceResponseBody) SetSuccess(v bool) *GlobalHotelValidatePriceResponseBody {
	s.Success = &v
	return s
}

func (s *GlobalHotelValidatePriceResponseBody) SetTracerId(v string) *GlobalHotelValidatePriceResponseBody {
	s.TracerId = &v
	return s
}

func (s *GlobalHotelValidatePriceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GlobalHotelValidatePriceResponseBodyData struct {
	CancellationPolicies []*GlobalHotelValidatePriceResponseBodyDataCancellationPolicies `json:"CancellationPolicies,omitempty" xml:"CancellationPolicies,omitempty" type:"Repeated"`
	// example:
	//
	// itemOffer_123
	ItemOfferId *string                                          `json:"ItemOfferId,omitempty" xml:"ItemOfferId,omitempty"`
	Pricing     *GlobalHotelValidatePriceResponseBodyDataPricing `json:"Pricing,omitempty" xml:"Pricing,omitempty" type:"Struct"`
	TracerId    *string                                          `json:"TracerId,omitempty" xml:"TracerId,omitempty"`
}

func (s GlobalHotelValidatePriceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GlobalHotelValidatePriceResponseBodyData) GoString() string {
	return s.String()
}

func (s *GlobalHotelValidatePriceResponseBodyData) GetCancellationPolicies() []*GlobalHotelValidatePriceResponseBodyDataCancellationPolicies {
	return s.CancellationPolicies
}

func (s *GlobalHotelValidatePriceResponseBodyData) GetItemOfferId() *string {
	return s.ItemOfferId
}

func (s *GlobalHotelValidatePriceResponseBodyData) GetPricing() *GlobalHotelValidatePriceResponseBodyDataPricing {
	return s.Pricing
}

func (s *GlobalHotelValidatePriceResponseBodyData) GetTracerId() *string {
	return s.TracerId
}

func (s *GlobalHotelValidatePriceResponseBodyData) SetCancellationPolicies(v []*GlobalHotelValidatePriceResponseBodyDataCancellationPolicies) *GlobalHotelValidatePriceResponseBodyData {
	s.CancellationPolicies = v
	return s
}

func (s *GlobalHotelValidatePriceResponseBodyData) SetItemOfferId(v string) *GlobalHotelValidatePriceResponseBodyData {
	s.ItemOfferId = &v
	return s
}

func (s *GlobalHotelValidatePriceResponseBodyData) SetPricing(v *GlobalHotelValidatePriceResponseBodyDataPricing) *GlobalHotelValidatePriceResponseBodyData {
	s.Pricing = v
	return s
}

func (s *GlobalHotelValidatePriceResponseBodyData) SetTracerId(v string) *GlobalHotelValidatePriceResponseBodyData {
	s.TracerId = &v
	return s
}

func (s *GlobalHotelValidatePriceResponseBodyData) Validate() error {
	if s.CancellationPolicies != nil {
		for _, item := range s.CancellationPolicies {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Pricing != nil {
		if err := s.Pricing.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GlobalHotelValidatePriceResponseBodyDataCancellationPolicies struct {
	Penalties []*GlobalHotelValidatePriceResponseBodyDataCancellationPoliciesPenalties `json:"Penalties,omitempty" xml:"Penalties,omitempty" type:"Repeated"`
	// example:
	//
	// FREE_CANCELLATION
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	TracerId   *string `json:"TracerId,omitempty" xml:"TracerId,omitempty"`
}

func (s GlobalHotelValidatePriceResponseBodyDataCancellationPolicies) String() string {
	return dara.Prettify(s)
}

func (s GlobalHotelValidatePriceResponseBodyDataCancellationPolicies) GoString() string {
	return s.String()
}

func (s *GlobalHotelValidatePriceResponseBodyDataCancellationPolicies) GetPenalties() []*GlobalHotelValidatePriceResponseBodyDataCancellationPoliciesPenalties {
	return s.Penalties
}

func (s *GlobalHotelValidatePriceResponseBodyDataCancellationPolicies) GetPolicyType() *string {
	return s.PolicyType
}

func (s *GlobalHotelValidatePriceResponseBodyDataCancellationPolicies) GetTracerId() *string {
	return s.TracerId
}

func (s *GlobalHotelValidatePriceResponseBodyDataCancellationPolicies) SetPenalties(v []*GlobalHotelValidatePriceResponseBodyDataCancellationPoliciesPenalties) *GlobalHotelValidatePriceResponseBodyDataCancellationPolicies {
	s.Penalties = v
	return s
}

func (s *GlobalHotelValidatePriceResponseBodyDataCancellationPolicies) SetPolicyType(v string) *GlobalHotelValidatePriceResponseBodyDataCancellationPolicies {
	s.PolicyType = &v
	return s
}

func (s *GlobalHotelValidatePriceResponseBodyDataCancellationPolicies) SetTracerId(v string) *GlobalHotelValidatePriceResponseBodyDataCancellationPolicies {
	s.TracerId = &v
	return s
}

func (s *GlobalHotelValidatePriceResponseBodyDataCancellationPolicies) Validate() error {
	if s.Penalties != nil {
		for _, item := range s.Penalties {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GlobalHotelValidatePriceResponseBodyDataCancellationPoliciesPenalties struct {
	// example:
	//
	// USD
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// example:
	//
	// 1672617600000
	End *int64 `json:"End,omitempty" xml:"End,omitempty"`
	// example:
	//
	// PERCENTAGE
	PenaltyType *string `json:"PenaltyType,omitempty" xml:"PenaltyType,omitempty"`
	// example:
	//
	// 50
	PenaltyValue *string `json:"PenaltyValue,omitempty" xml:"PenaltyValue,omitempty"`
	// example:
	//
	// 1672531200000
	Start    *int64  `json:"Start,omitempty" xml:"Start,omitempty"`
	TracerId *string `json:"TracerId,omitempty" xml:"TracerId,omitempty"`
}

func (s GlobalHotelValidatePriceResponseBodyDataCancellationPoliciesPenalties) String() string {
	return dara.Prettify(s)
}

func (s GlobalHotelValidatePriceResponseBodyDataCancellationPoliciesPenalties) GoString() string {
	return s.String()
}

func (s *GlobalHotelValidatePriceResponseBodyDataCancellationPoliciesPenalties) GetCurrency() *string {
	return s.Currency
}

func (s *GlobalHotelValidatePriceResponseBodyDataCancellationPoliciesPenalties) GetEnd() *int64 {
	return s.End
}

func (s *GlobalHotelValidatePriceResponseBodyDataCancellationPoliciesPenalties) GetPenaltyType() *string {
	return s.PenaltyType
}

func (s *GlobalHotelValidatePriceResponseBodyDataCancellationPoliciesPenalties) GetPenaltyValue() *string {
	return s.PenaltyValue
}

func (s *GlobalHotelValidatePriceResponseBodyDataCancellationPoliciesPenalties) GetStart() *int64 {
	return s.Start
}

func (s *GlobalHotelValidatePriceResponseBodyDataCancellationPoliciesPenalties) GetTracerId() *string {
	return s.TracerId
}

func (s *GlobalHotelValidatePriceResponseBodyDataCancellationPoliciesPenalties) SetCurrency(v string) *GlobalHotelValidatePriceResponseBodyDataCancellationPoliciesPenalties {
	s.Currency = &v
	return s
}

func (s *GlobalHotelValidatePriceResponseBodyDataCancellationPoliciesPenalties) SetEnd(v int64) *GlobalHotelValidatePriceResponseBodyDataCancellationPoliciesPenalties {
	s.End = &v
	return s
}

func (s *GlobalHotelValidatePriceResponseBodyDataCancellationPoliciesPenalties) SetPenaltyType(v string) *GlobalHotelValidatePriceResponseBodyDataCancellationPoliciesPenalties {
	s.PenaltyType = &v
	return s
}

func (s *GlobalHotelValidatePriceResponseBodyDataCancellationPoliciesPenalties) SetPenaltyValue(v string) *GlobalHotelValidatePriceResponseBodyDataCancellationPoliciesPenalties {
	s.PenaltyValue = &v
	return s
}

func (s *GlobalHotelValidatePriceResponseBodyDataCancellationPoliciesPenalties) SetStart(v int64) *GlobalHotelValidatePriceResponseBodyDataCancellationPoliciesPenalties {
	s.Start = &v
	return s
}

func (s *GlobalHotelValidatePriceResponseBodyDataCancellationPoliciesPenalties) SetTracerId(v string) *GlobalHotelValidatePriceResponseBodyDataCancellationPoliciesPenalties {
	s.TracerId = &v
	return s
}

func (s *GlobalHotelValidatePriceResponseBodyDataCancellationPoliciesPenalties) Validate() error {
	return dara.Validate(s)
}

type GlobalHotelValidatePriceResponseBodyDataPricing struct {
	// example:
	//
	// USD
	Currency      *string                                                         `json:"Currency,omitempty" xml:"Currency,omitempty"`
	NightlyPrices []*GlobalHotelValidatePriceResponseBodyDataPricingNightlyPrices `json:"NightlyPrices,omitempty" xml:"NightlyPrices,omitempty" type:"Repeated"`
	// example:
	//
	// 10000
	TotalAmount *string `json:"TotalAmount,omitempty" xml:"TotalAmount,omitempty"`
	TracerId    *string `json:"TracerId,omitempty" xml:"TracerId,omitempty"`
}

func (s GlobalHotelValidatePriceResponseBodyDataPricing) String() string {
	return dara.Prettify(s)
}

func (s GlobalHotelValidatePriceResponseBodyDataPricing) GoString() string {
	return s.String()
}

func (s *GlobalHotelValidatePriceResponseBodyDataPricing) GetCurrency() *string {
	return s.Currency
}

func (s *GlobalHotelValidatePriceResponseBodyDataPricing) GetNightlyPrices() []*GlobalHotelValidatePriceResponseBodyDataPricingNightlyPrices {
	return s.NightlyPrices
}

func (s *GlobalHotelValidatePriceResponseBodyDataPricing) GetTotalAmount() *string {
	return s.TotalAmount
}

func (s *GlobalHotelValidatePriceResponseBodyDataPricing) GetTracerId() *string {
	return s.TracerId
}

func (s *GlobalHotelValidatePriceResponseBodyDataPricing) SetCurrency(v string) *GlobalHotelValidatePriceResponseBodyDataPricing {
	s.Currency = &v
	return s
}

func (s *GlobalHotelValidatePriceResponseBodyDataPricing) SetNightlyPrices(v []*GlobalHotelValidatePriceResponseBodyDataPricingNightlyPrices) *GlobalHotelValidatePriceResponseBodyDataPricing {
	s.NightlyPrices = v
	return s
}

func (s *GlobalHotelValidatePriceResponseBodyDataPricing) SetTotalAmount(v string) *GlobalHotelValidatePriceResponseBodyDataPricing {
	s.TotalAmount = &v
	return s
}

func (s *GlobalHotelValidatePriceResponseBodyDataPricing) SetTracerId(v string) *GlobalHotelValidatePriceResponseBodyDataPricing {
	s.TracerId = &v
	return s
}

func (s *GlobalHotelValidatePriceResponseBodyDataPricing) Validate() error {
	if s.NightlyPrices != nil {
		for _, item := range s.NightlyPrices {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GlobalHotelValidatePriceResponseBodyDataPricingNightlyPrices struct {
	// example:
	//
	// 5000
	Amount *string `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// example:
	//
	// 2026-07-01
	Date     *string `json:"Date,omitempty" xml:"Date,omitempty"`
	TracerId *string `json:"TracerId,omitempty" xml:"TracerId,omitempty"`
}

func (s GlobalHotelValidatePriceResponseBodyDataPricingNightlyPrices) String() string {
	return dara.Prettify(s)
}

func (s GlobalHotelValidatePriceResponseBodyDataPricingNightlyPrices) GoString() string {
	return s.String()
}

func (s *GlobalHotelValidatePriceResponseBodyDataPricingNightlyPrices) GetAmount() *string {
	return s.Amount
}

func (s *GlobalHotelValidatePriceResponseBodyDataPricingNightlyPrices) GetDate() *string {
	return s.Date
}

func (s *GlobalHotelValidatePriceResponseBodyDataPricingNightlyPrices) GetTracerId() *string {
	return s.TracerId
}

func (s *GlobalHotelValidatePriceResponseBodyDataPricingNightlyPrices) SetAmount(v string) *GlobalHotelValidatePriceResponseBodyDataPricingNightlyPrices {
	s.Amount = &v
	return s
}

func (s *GlobalHotelValidatePriceResponseBodyDataPricingNightlyPrices) SetDate(v string) *GlobalHotelValidatePriceResponseBodyDataPricingNightlyPrices {
	s.Date = &v
	return s
}

func (s *GlobalHotelValidatePriceResponseBodyDataPricingNightlyPrices) SetTracerId(v string) *GlobalHotelValidatePriceResponseBodyDataPricingNightlyPrices {
	s.TracerId = &v
	return s
}

func (s *GlobalHotelValidatePriceResponseBodyDataPricingNightlyPrices) Validate() error {
	return dara.Validate(s)
}
