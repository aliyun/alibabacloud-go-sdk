// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iValidatePriceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ValidatePriceResponseBodyData) *ValidatePriceResponseBody
	GetData() *ValidatePriceResponseBodyData
	SetErrorCode(v string) *ValidatePriceResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *ValidatePriceResponseBody
	GetErrorMsg() *string
	SetRequestId(v string) *ValidatePriceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ValidatePriceResponseBody
	GetSuccess() *bool
	SetTracerId(v string) *ValidatePriceResponseBody
	GetTracerId() *string
}

type ValidatePriceResponseBody struct {
	Data *ValidatePriceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// traceId
	TracerId *string `json:"TracerId,omitempty" xml:"TracerId,omitempty"`
}

func (s ValidatePriceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ValidatePriceResponseBody) GoString() string {
	return s.String()
}

func (s *ValidatePriceResponseBody) GetData() *ValidatePriceResponseBodyData {
	return s.Data
}

func (s *ValidatePriceResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ValidatePriceResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *ValidatePriceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ValidatePriceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ValidatePriceResponseBody) GetTracerId() *string {
	return s.TracerId
}

func (s *ValidatePriceResponseBody) SetData(v *ValidatePriceResponseBodyData) *ValidatePriceResponseBody {
	s.Data = v
	return s
}

func (s *ValidatePriceResponseBody) SetErrorCode(v string) *ValidatePriceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ValidatePriceResponseBody) SetErrorMsg(v string) *ValidatePriceResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *ValidatePriceResponseBody) SetRequestId(v string) *ValidatePriceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ValidatePriceResponseBody) SetSuccess(v bool) *ValidatePriceResponseBody {
	s.Success = &v
	return s
}

func (s *ValidatePriceResponseBody) SetTracerId(v string) *ValidatePriceResponseBody {
	s.TracerId = &v
	return s
}

func (s *ValidatePriceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ValidatePriceResponseBodyData struct {
	CancellationPolicies []*ValidatePriceResponseBodyDataCancellationPolicies `json:"CancellationPolicies,omitempty" xml:"CancellationPolicies,omitempty" type:"Repeated"`
	// example:
	//
	// itemOffer_123
	ItemOfferId *string                               `json:"ItemOfferId,omitempty" xml:"ItemOfferId,omitempty"`
	Pricing     *ValidatePriceResponseBodyDataPricing `json:"Pricing,omitempty" xml:"Pricing,omitempty" type:"Struct"`
	// example:
	//
	// traceId
	TracerId *string `json:"TracerId,omitempty" xml:"TracerId,omitempty"`
}

func (s ValidatePriceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ValidatePriceResponseBodyData) GoString() string {
	return s.String()
}

func (s *ValidatePriceResponseBodyData) GetCancellationPolicies() []*ValidatePriceResponseBodyDataCancellationPolicies {
	return s.CancellationPolicies
}

func (s *ValidatePriceResponseBodyData) GetItemOfferId() *string {
	return s.ItemOfferId
}

func (s *ValidatePriceResponseBodyData) GetPricing() *ValidatePriceResponseBodyDataPricing {
	return s.Pricing
}

func (s *ValidatePriceResponseBodyData) GetTracerId() *string {
	return s.TracerId
}

func (s *ValidatePriceResponseBodyData) SetCancellationPolicies(v []*ValidatePriceResponseBodyDataCancellationPolicies) *ValidatePriceResponseBodyData {
	s.CancellationPolicies = v
	return s
}

func (s *ValidatePriceResponseBodyData) SetItemOfferId(v string) *ValidatePriceResponseBodyData {
	s.ItemOfferId = &v
	return s
}

func (s *ValidatePriceResponseBodyData) SetPricing(v *ValidatePriceResponseBodyDataPricing) *ValidatePriceResponseBodyData {
	s.Pricing = v
	return s
}

func (s *ValidatePriceResponseBodyData) SetTracerId(v string) *ValidatePriceResponseBodyData {
	s.TracerId = &v
	return s
}

func (s *ValidatePriceResponseBodyData) Validate() error {
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

type ValidatePriceResponseBodyDataCancellationPolicies struct {
	Penalties []*ValidatePriceResponseBodyDataCancellationPoliciesPenalties `json:"Penalties,omitempty" xml:"Penalties,omitempty" type:"Repeated"`
	// example:
	//
	// FREE_CANCELLATION
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	// example:
	//
	// traceId
	TracerId *string `json:"TracerId,omitempty" xml:"TracerId,omitempty"`
}

func (s ValidatePriceResponseBodyDataCancellationPolicies) String() string {
	return dara.Prettify(s)
}

func (s ValidatePriceResponseBodyDataCancellationPolicies) GoString() string {
	return s.String()
}

func (s *ValidatePriceResponseBodyDataCancellationPolicies) GetPenalties() []*ValidatePriceResponseBodyDataCancellationPoliciesPenalties {
	return s.Penalties
}

func (s *ValidatePriceResponseBodyDataCancellationPolicies) GetPolicyType() *string {
	return s.PolicyType
}

func (s *ValidatePriceResponseBodyDataCancellationPolicies) GetTracerId() *string {
	return s.TracerId
}

func (s *ValidatePriceResponseBodyDataCancellationPolicies) SetPenalties(v []*ValidatePriceResponseBodyDataCancellationPoliciesPenalties) *ValidatePriceResponseBodyDataCancellationPolicies {
	s.Penalties = v
	return s
}

func (s *ValidatePriceResponseBodyDataCancellationPolicies) SetPolicyType(v string) *ValidatePriceResponseBodyDataCancellationPolicies {
	s.PolicyType = &v
	return s
}

func (s *ValidatePriceResponseBodyDataCancellationPolicies) SetTracerId(v string) *ValidatePriceResponseBodyDataCancellationPolicies {
	s.TracerId = &v
	return s
}

func (s *ValidatePriceResponseBodyDataCancellationPolicies) Validate() error {
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

type ValidatePriceResponseBodyDataCancellationPoliciesPenalties struct {
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
	Start *int64 `json:"Start,omitempty" xml:"Start,omitempty"`
	// example:
	//
	// traceId
	TracerId *string `json:"TracerId,omitempty" xml:"TracerId,omitempty"`
}

func (s ValidatePriceResponseBodyDataCancellationPoliciesPenalties) String() string {
	return dara.Prettify(s)
}

func (s ValidatePriceResponseBodyDataCancellationPoliciesPenalties) GoString() string {
	return s.String()
}

func (s *ValidatePriceResponseBodyDataCancellationPoliciesPenalties) GetCurrency() *string {
	return s.Currency
}

func (s *ValidatePriceResponseBodyDataCancellationPoliciesPenalties) GetEnd() *int64 {
	return s.End
}

func (s *ValidatePriceResponseBodyDataCancellationPoliciesPenalties) GetPenaltyType() *string {
	return s.PenaltyType
}

func (s *ValidatePriceResponseBodyDataCancellationPoliciesPenalties) GetPenaltyValue() *string {
	return s.PenaltyValue
}

func (s *ValidatePriceResponseBodyDataCancellationPoliciesPenalties) GetStart() *int64 {
	return s.Start
}

func (s *ValidatePriceResponseBodyDataCancellationPoliciesPenalties) GetTracerId() *string {
	return s.TracerId
}

func (s *ValidatePriceResponseBodyDataCancellationPoliciesPenalties) SetCurrency(v string) *ValidatePriceResponseBodyDataCancellationPoliciesPenalties {
	s.Currency = &v
	return s
}

func (s *ValidatePriceResponseBodyDataCancellationPoliciesPenalties) SetEnd(v int64) *ValidatePriceResponseBodyDataCancellationPoliciesPenalties {
	s.End = &v
	return s
}

func (s *ValidatePriceResponseBodyDataCancellationPoliciesPenalties) SetPenaltyType(v string) *ValidatePriceResponseBodyDataCancellationPoliciesPenalties {
	s.PenaltyType = &v
	return s
}

func (s *ValidatePriceResponseBodyDataCancellationPoliciesPenalties) SetPenaltyValue(v string) *ValidatePriceResponseBodyDataCancellationPoliciesPenalties {
	s.PenaltyValue = &v
	return s
}

func (s *ValidatePriceResponseBodyDataCancellationPoliciesPenalties) SetStart(v int64) *ValidatePriceResponseBodyDataCancellationPoliciesPenalties {
	s.Start = &v
	return s
}

func (s *ValidatePriceResponseBodyDataCancellationPoliciesPenalties) SetTracerId(v string) *ValidatePriceResponseBodyDataCancellationPoliciesPenalties {
	s.TracerId = &v
	return s
}

func (s *ValidatePriceResponseBodyDataCancellationPoliciesPenalties) Validate() error {
	return dara.Validate(s)
}

type ValidatePriceResponseBodyDataPricing struct {
	// example:
	//
	// USD
	Currency      *string                                              `json:"Currency,omitempty" xml:"Currency,omitempty"`
	NightlyPrices []*ValidatePriceResponseBodyDataPricingNightlyPrices `json:"NightlyPrices,omitempty" xml:"NightlyPrices,omitempty" type:"Repeated"`
	// example:
	//
	// 10000
	TotalAmount *string `json:"TotalAmount,omitempty" xml:"TotalAmount,omitempty"`
	// example:
	//
	// traceId
	TracerId *string `json:"TracerId,omitempty" xml:"TracerId,omitempty"`
}

func (s ValidatePriceResponseBodyDataPricing) String() string {
	return dara.Prettify(s)
}

func (s ValidatePriceResponseBodyDataPricing) GoString() string {
	return s.String()
}

func (s *ValidatePriceResponseBodyDataPricing) GetCurrency() *string {
	return s.Currency
}

func (s *ValidatePriceResponseBodyDataPricing) GetNightlyPrices() []*ValidatePriceResponseBodyDataPricingNightlyPrices {
	return s.NightlyPrices
}

func (s *ValidatePriceResponseBodyDataPricing) GetTotalAmount() *string {
	return s.TotalAmount
}

func (s *ValidatePriceResponseBodyDataPricing) GetTracerId() *string {
	return s.TracerId
}

func (s *ValidatePriceResponseBodyDataPricing) SetCurrency(v string) *ValidatePriceResponseBodyDataPricing {
	s.Currency = &v
	return s
}

func (s *ValidatePriceResponseBodyDataPricing) SetNightlyPrices(v []*ValidatePriceResponseBodyDataPricingNightlyPrices) *ValidatePriceResponseBodyDataPricing {
	s.NightlyPrices = v
	return s
}

func (s *ValidatePriceResponseBodyDataPricing) SetTotalAmount(v string) *ValidatePriceResponseBodyDataPricing {
	s.TotalAmount = &v
	return s
}

func (s *ValidatePriceResponseBodyDataPricing) SetTracerId(v string) *ValidatePriceResponseBodyDataPricing {
	s.TracerId = &v
	return s
}

func (s *ValidatePriceResponseBodyDataPricing) Validate() error {
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

type ValidatePriceResponseBodyDataPricingNightlyPrices struct {
	// example:
	//
	// 5000
	Amount *string `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// example:
	//
	// 2026-07-01
	Date *string `json:"Date,omitempty" xml:"Date,omitempty"`
	// example:
	//
	// traceId
	TracerId *string `json:"TracerId,omitempty" xml:"TracerId,omitempty"`
}

func (s ValidatePriceResponseBodyDataPricingNightlyPrices) String() string {
	return dara.Prettify(s)
}

func (s ValidatePriceResponseBodyDataPricingNightlyPrices) GoString() string {
	return s.String()
}

func (s *ValidatePriceResponseBodyDataPricingNightlyPrices) GetAmount() *string {
	return s.Amount
}

func (s *ValidatePriceResponseBodyDataPricingNightlyPrices) GetDate() *string {
	return s.Date
}

func (s *ValidatePriceResponseBodyDataPricingNightlyPrices) GetTracerId() *string {
	return s.TracerId
}

func (s *ValidatePriceResponseBodyDataPricingNightlyPrices) SetAmount(v string) *ValidatePriceResponseBodyDataPricingNightlyPrices {
	s.Amount = &v
	return s
}

func (s *ValidatePriceResponseBodyDataPricingNightlyPrices) SetDate(v string) *ValidatePriceResponseBodyDataPricingNightlyPrices {
	s.Date = &v
	return s
}

func (s *ValidatePriceResponseBodyDataPricingNightlyPrices) SetTracerId(v string) *ValidatePriceResponseBodyDataPricingNightlyPrices {
	s.TracerId = &v
	return s
}

func (s *ValidatePriceResponseBodyDataPricingNightlyPrices) Validate() error {
	return dara.Validate(s)
}
