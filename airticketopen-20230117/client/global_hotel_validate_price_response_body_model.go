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
	// The business data.
	Data *GlobalHotelValidatePriceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code.
	//
	// example:
	//
	// CreateOrderFailed
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// Failed to create order
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// The unique request ID.
	//
	// example:
	//
	// 260E4F99-983D-1919-834C-5C42E98E5B2B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// TracerId
	//
	// example:
	//
	// TracerId
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
	// The cancellation policies.
	CancellationPolicies []*GlobalHotelValidatePriceResponseBodyDataCancellationPolicies `json:"CancellationPolicies,omitempty" xml:"CancellationPolicies,omitempty" type:"Repeated"`
	// The list of daily prices.
	DailyPrices []*GlobalHotelValidatePriceResponseBodyDataDailyPrices `json:"DailyPrices,omitempty" xml:"DailyPrices,omitempty" type:"Repeated"`
	// The price validation result ID, used for subsequent order creation.
	//
	// example:
	//
	// itemOffer_123
	ItemOfferId *string `json:"ItemOfferId,omitempty" xml:"ItemOfferId,omitempty"`
	// The total selling price.
	TotalPrice *GlobalHotelValidatePriceResponseBodyDataTotalPrice `json:"TotalPrice,omitempty" xml:"TotalPrice,omitempty" type:"Struct"`
	// TracerId
	//
	// example:
	//
	// TracerId
	TracerId *string `json:"TracerId,omitempty" xml:"TracerId,omitempty"`
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

func (s *GlobalHotelValidatePriceResponseBodyData) GetDailyPrices() []*GlobalHotelValidatePriceResponseBodyDataDailyPrices {
	return s.DailyPrices
}

func (s *GlobalHotelValidatePriceResponseBodyData) GetItemOfferId() *string {
	return s.ItemOfferId
}

func (s *GlobalHotelValidatePriceResponseBodyData) GetTotalPrice() *GlobalHotelValidatePriceResponseBodyDataTotalPrice {
	return s.TotalPrice
}

func (s *GlobalHotelValidatePriceResponseBodyData) GetTracerId() *string {
	return s.TracerId
}

func (s *GlobalHotelValidatePriceResponseBodyData) SetCancellationPolicies(v []*GlobalHotelValidatePriceResponseBodyDataCancellationPolicies) *GlobalHotelValidatePriceResponseBodyData {
	s.CancellationPolicies = v
	return s
}

func (s *GlobalHotelValidatePriceResponseBodyData) SetDailyPrices(v []*GlobalHotelValidatePriceResponseBodyDataDailyPrices) *GlobalHotelValidatePriceResponseBodyData {
	s.DailyPrices = v
	return s
}

func (s *GlobalHotelValidatePriceResponseBodyData) SetItemOfferId(v string) *GlobalHotelValidatePriceResponseBodyData {
	s.ItemOfferId = &v
	return s
}

func (s *GlobalHotelValidatePriceResponseBodyData) SetTotalPrice(v *GlobalHotelValidatePriceResponseBodyDataTotalPrice) *GlobalHotelValidatePriceResponseBodyData {
	s.TotalPrice = v
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
	if s.DailyPrices != nil {
		for _, item := range s.DailyPrices {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.TotalPrice != nil {
		if err := s.TotalPrice.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GlobalHotelValidatePriceResponseBodyDataCancellationPolicies struct {
	// The list of cancellation penalty details.
	Penalties []*GlobalHotelValidatePriceResponseBodyDataCancellationPoliciesPenalties `json:"Penalties,omitempty" xml:"Penalties,omitempty" type:"Repeated"`
	// The cancellation policy type (FREE_CANCEL/CONDITIONAL/NON_REFUNDABLE).
	//
	// example:
	//
	// FREE_CANCELLATION
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	// TracerId
	//
	// example:
	//
	// TracerId
	TracerId *string `json:"TracerId,omitempty" xml:"TracerId,omitempty"`
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
	// The currency code. This field has a value only when the penalty type is AMOUNT.
	//
	// example:
	//
	// USD
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// The effective end time as a UTC millisecond timestamp.
	//
	// example:
	//
	// 1672617600000
	End *string `json:"End,omitempty" xml:"End,omitempty"`
	// The penalty type (PERCENT/NIGHTS/NON_CANCELLABLE).
	//
	// example:
	//
	// PERCENTAGE
	PenaltyType *string `json:"PenaltyType,omitempty" xml:"PenaltyType,omitempty"`
	// The penalty value (percentage, amount, or number of nights). This field is not present when PenaltyType is NON_CANCELLABLE.
	//
	// example:
	//
	// 50
	PenaltyValue *string `json:"PenaltyValue,omitempty" xml:"PenaltyValue,omitempty"`
	// The effective start time as a UTC millisecond timestamp.
	//
	// example:
	//
	// 1672531200000
	Start *string `json:"Start,omitempty" xml:"Start,omitempty"`
	// TracerId
	//
	// example:
	//
	// TracerId
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

func (s *GlobalHotelValidatePriceResponseBodyDataCancellationPoliciesPenalties) GetEnd() *string {
	return s.End
}

func (s *GlobalHotelValidatePriceResponseBodyDataCancellationPoliciesPenalties) GetPenaltyType() *string {
	return s.PenaltyType
}

func (s *GlobalHotelValidatePriceResponseBodyDataCancellationPoliciesPenalties) GetPenaltyValue() *string {
	return s.PenaltyValue
}

func (s *GlobalHotelValidatePriceResponseBodyDataCancellationPoliciesPenalties) GetStart() *string {
	return s.Start
}

func (s *GlobalHotelValidatePriceResponseBodyDataCancellationPoliciesPenalties) GetTracerId() *string {
	return s.TracerId
}

func (s *GlobalHotelValidatePriceResponseBodyDataCancellationPoliciesPenalties) SetCurrency(v string) *GlobalHotelValidatePriceResponseBodyDataCancellationPoliciesPenalties {
	s.Currency = &v
	return s
}

func (s *GlobalHotelValidatePriceResponseBodyDataCancellationPoliciesPenalties) SetEnd(v string) *GlobalHotelValidatePriceResponseBodyDataCancellationPoliciesPenalties {
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

func (s *GlobalHotelValidatePriceResponseBodyDataCancellationPoliciesPenalties) SetStart(v string) *GlobalHotelValidatePriceResponseBodyDataCancellationPoliciesPenalties {
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

type GlobalHotelValidatePriceResponseBodyDataDailyPrices struct {
	// The date in yyyy-MM-dd format, in the local time zone of the hotel.
	//
	// example:
	//
	// 2026-08-16
	Date *string `json:"Date,omitempty" xml:"Date,omitempty"`
	// The price for the night.
	Price *GlobalHotelValidatePriceResponseBodyDataDailyPricesPrice `json:"Price,omitempty" xml:"Price,omitempty" type:"Struct"`
	// null
	//
	// example:
	//
	// null
	TracerId *string `json:"TracerId,omitempty" xml:"TracerId,omitempty"`
}

func (s GlobalHotelValidatePriceResponseBodyDataDailyPrices) String() string {
	return dara.Prettify(s)
}

func (s GlobalHotelValidatePriceResponseBodyDataDailyPrices) GoString() string {
	return s.String()
}

func (s *GlobalHotelValidatePriceResponseBodyDataDailyPrices) GetDate() *string {
	return s.Date
}

func (s *GlobalHotelValidatePriceResponseBodyDataDailyPrices) GetPrice() *GlobalHotelValidatePriceResponseBodyDataDailyPricesPrice {
	return s.Price
}

func (s *GlobalHotelValidatePriceResponseBodyDataDailyPrices) GetTracerId() *string {
	return s.TracerId
}

func (s *GlobalHotelValidatePriceResponseBodyDataDailyPrices) SetDate(v string) *GlobalHotelValidatePriceResponseBodyDataDailyPrices {
	s.Date = &v
	return s
}

func (s *GlobalHotelValidatePriceResponseBodyDataDailyPrices) SetPrice(v *GlobalHotelValidatePriceResponseBodyDataDailyPricesPrice) *GlobalHotelValidatePriceResponseBodyDataDailyPrices {
	s.Price = v
	return s
}

func (s *GlobalHotelValidatePriceResponseBodyDataDailyPrices) SetTracerId(v string) *GlobalHotelValidatePriceResponseBodyDataDailyPrices {
	s.TracerId = &v
	return s
}

func (s *GlobalHotelValidatePriceResponseBodyDataDailyPrices) Validate() error {
	if s.Price != nil {
		if err := s.Price.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GlobalHotelValidatePriceResponseBodyDataDailyPricesPrice struct {
	// The amount in the smallest currency unit.
	//
	// example:
	//
	// 574
	Amount *string `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// The currency code (ISO 4217).
	//
	// example:
	//
	// USD
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// null
	//
	// example:
	//
	// null
	TracerId *string `json:"TracerId,omitempty" xml:"TracerId,omitempty"`
}

func (s GlobalHotelValidatePriceResponseBodyDataDailyPricesPrice) String() string {
	return dara.Prettify(s)
}

func (s GlobalHotelValidatePriceResponseBodyDataDailyPricesPrice) GoString() string {
	return s.String()
}

func (s *GlobalHotelValidatePriceResponseBodyDataDailyPricesPrice) GetAmount() *string {
	return s.Amount
}

func (s *GlobalHotelValidatePriceResponseBodyDataDailyPricesPrice) GetCurrency() *string {
	return s.Currency
}

func (s *GlobalHotelValidatePriceResponseBodyDataDailyPricesPrice) GetTracerId() *string {
	return s.TracerId
}

func (s *GlobalHotelValidatePriceResponseBodyDataDailyPricesPrice) SetAmount(v string) *GlobalHotelValidatePriceResponseBodyDataDailyPricesPrice {
	s.Amount = &v
	return s
}

func (s *GlobalHotelValidatePriceResponseBodyDataDailyPricesPrice) SetCurrency(v string) *GlobalHotelValidatePriceResponseBodyDataDailyPricesPrice {
	s.Currency = &v
	return s
}

func (s *GlobalHotelValidatePriceResponseBodyDataDailyPricesPrice) SetTracerId(v string) *GlobalHotelValidatePriceResponseBodyDataDailyPricesPrice {
	s.TracerId = &v
	return s
}

func (s *GlobalHotelValidatePriceResponseBodyDataDailyPricesPrice) Validate() error {
	return dara.Validate(s)
}

type GlobalHotelValidatePriceResponseBodyDataTotalPrice struct {
	// The amount in the smallest currency unit.
	//
	// example:
	//
	// 574
	Amount *string `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// The currency code (ISO 4217).
	//
	// example:
	//
	// USD
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// null
	//
	// example:
	//
	// null
	TracerId *string `json:"TracerId,omitempty" xml:"TracerId,omitempty"`
}

func (s GlobalHotelValidatePriceResponseBodyDataTotalPrice) String() string {
	return dara.Prettify(s)
}

func (s GlobalHotelValidatePriceResponseBodyDataTotalPrice) GoString() string {
	return s.String()
}

func (s *GlobalHotelValidatePriceResponseBodyDataTotalPrice) GetAmount() *string {
	return s.Amount
}

func (s *GlobalHotelValidatePriceResponseBodyDataTotalPrice) GetCurrency() *string {
	return s.Currency
}

func (s *GlobalHotelValidatePriceResponseBodyDataTotalPrice) GetTracerId() *string {
	return s.TracerId
}

func (s *GlobalHotelValidatePriceResponseBodyDataTotalPrice) SetAmount(v string) *GlobalHotelValidatePriceResponseBodyDataTotalPrice {
	s.Amount = &v
	return s
}

func (s *GlobalHotelValidatePriceResponseBodyDataTotalPrice) SetCurrency(v string) *GlobalHotelValidatePriceResponseBodyDataTotalPrice {
	s.Currency = &v
	return s
}

func (s *GlobalHotelValidatePriceResponseBodyDataTotalPrice) SetTracerId(v string) *GlobalHotelValidatePriceResponseBodyDataTotalPrice {
	s.TracerId = &v
	return s
}

func (s *GlobalHotelValidatePriceResponseBodyDataTotalPrice) Validate() error {
	return dara.Validate(s)
}
