// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGlobalHotelQueryOrderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GlobalHotelQueryOrderResponseBodyData) *GlobalHotelQueryOrderResponseBody
	GetData() *GlobalHotelQueryOrderResponseBodyData
	SetErrorCode(v string) *GlobalHotelQueryOrderResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *GlobalHotelQueryOrderResponseBody
	GetErrorMsg() *string
	SetRequestId(v string) *GlobalHotelQueryOrderResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GlobalHotelQueryOrderResponseBody
	GetSuccess() *bool
	SetTracerId(v string) *GlobalHotelQueryOrderResponseBody
	GetTracerId() *string
}

type GlobalHotelQueryOrderResponseBody struct {
	// The business data.
	Data *GlobalHotelQueryOrderResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// The unique ID of the request.
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

func (s GlobalHotelQueryOrderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GlobalHotelQueryOrderResponseBody) GoString() string {
	return s.String()
}

func (s *GlobalHotelQueryOrderResponseBody) GetData() *GlobalHotelQueryOrderResponseBodyData {
	return s.Data
}

func (s *GlobalHotelQueryOrderResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GlobalHotelQueryOrderResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *GlobalHotelQueryOrderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GlobalHotelQueryOrderResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GlobalHotelQueryOrderResponseBody) GetTracerId() *string {
	return s.TracerId
}

func (s *GlobalHotelQueryOrderResponseBody) SetData(v *GlobalHotelQueryOrderResponseBodyData) *GlobalHotelQueryOrderResponseBody {
	s.Data = v
	return s
}

func (s *GlobalHotelQueryOrderResponseBody) SetErrorCode(v string) *GlobalHotelQueryOrderResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GlobalHotelQueryOrderResponseBody) SetErrorMsg(v string) *GlobalHotelQueryOrderResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *GlobalHotelQueryOrderResponseBody) SetRequestId(v string) *GlobalHotelQueryOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *GlobalHotelQueryOrderResponseBody) SetSuccess(v bool) *GlobalHotelQueryOrderResponseBody {
	s.Success = &v
	return s
}

func (s *GlobalHotelQueryOrderResponseBody) SetTracerId(v string) *GlobalHotelQueryOrderResponseBody {
	s.TracerId = &v
	return s
}

func (s *GlobalHotelQueryOrderResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GlobalHotelQueryOrderResponseBodyData struct {
	// The buyer ID.
	//
	// example:
	//
	// 123456
	BuyerId *string `json:"BuyerId,omitempty" xml:"BuyerId,omitempty"`
	// The external order number of the buyer.
	//
	// example:
	//
	// EXT_ORDER_001
	ExternalOrderNo *string `json:"ExternalOrderNo,omitempty" xml:"ExternalOrderNo,omitempty"`
	// The creation time in UTC millisecond timestamp.
	//
	// example:
	//
	// 1672531200000
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The item information.
	ItemInfo *GlobalHotelQueryOrderResponseBodyDataItemInfo `json:"ItemInfo,omitempty" xml:"ItemInfo,omitempty" type:"Struct"`
	// The order number.
	//
	// example:
	//
	// SO202606290001
	OrderNo *string `json:"OrderNo,omitempty" xml:"OrderNo,omitempty"`
	// The payment information.
	Payment *GlobalHotelQueryOrderResponseBodyDataPayment `json:"Payment,omitempty" xml:"Payment,omitempty" type:"Struct"`
	// The list of refund orders.
	RefundOrders []*GlobalHotelQueryOrderResponseBodyDataRefundOrders `json:"RefundOrders,omitempty" xml:"RefundOrders,omitempty" type:"Repeated"`
	// The list of room stays.
	RoomStays []*GlobalHotelQueryOrderResponseBodyDataRoomStays `json:"RoomStays,omitempty" xml:"RoomStays,omitempty" type:"Repeated"`
	// The sales channel.
	//
	// example:
	//
	// POP
	SalesChannel *string `json:"SalesChannel,omitempty" xml:"SalesChannel,omitempty"`
	// The unified order status.
	//
	// example:
	//
	// CONFIRMED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// TracerId
	//
	// example:
	//
	// TracerId
	TracerId *string `json:"TracerId,omitempty" xml:"TracerId,omitempty"`
}

func (s GlobalHotelQueryOrderResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GlobalHotelQueryOrderResponseBodyData) GoString() string {
	return s.String()
}

func (s *GlobalHotelQueryOrderResponseBodyData) GetBuyerId() *string {
	return s.BuyerId
}

func (s *GlobalHotelQueryOrderResponseBodyData) GetExternalOrderNo() *string {
	return s.ExternalOrderNo
}

func (s *GlobalHotelQueryOrderResponseBodyData) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *GlobalHotelQueryOrderResponseBodyData) GetItemInfo() *GlobalHotelQueryOrderResponseBodyDataItemInfo {
	return s.ItemInfo
}

func (s *GlobalHotelQueryOrderResponseBodyData) GetOrderNo() *string {
	return s.OrderNo
}

func (s *GlobalHotelQueryOrderResponseBodyData) GetPayment() *GlobalHotelQueryOrderResponseBodyDataPayment {
	return s.Payment
}

func (s *GlobalHotelQueryOrderResponseBodyData) GetRefundOrders() []*GlobalHotelQueryOrderResponseBodyDataRefundOrders {
	return s.RefundOrders
}

func (s *GlobalHotelQueryOrderResponseBodyData) GetRoomStays() []*GlobalHotelQueryOrderResponseBodyDataRoomStays {
	return s.RoomStays
}

func (s *GlobalHotelQueryOrderResponseBodyData) GetSalesChannel() *string {
	return s.SalesChannel
}

func (s *GlobalHotelQueryOrderResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GlobalHotelQueryOrderResponseBodyData) GetTracerId() *string {
	return s.TracerId
}

func (s *GlobalHotelQueryOrderResponseBodyData) SetBuyerId(v string) *GlobalHotelQueryOrderResponseBodyData {
	s.BuyerId = &v
	return s
}

func (s *GlobalHotelQueryOrderResponseBodyData) SetExternalOrderNo(v string) *GlobalHotelQueryOrderResponseBodyData {
	s.ExternalOrderNo = &v
	return s
}

func (s *GlobalHotelQueryOrderResponseBodyData) SetGmtCreate(v int64) *GlobalHotelQueryOrderResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *GlobalHotelQueryOrderResponseBodyData) SetItemInfo(v *GlobalHotelQueryOrderResponseBodyDataItemInfo) *GlobalHotelQueryOrderResponseBodyData {
	s.ItemInfo = v
	return s
}

func (s *GlobalHotelQueryOrderResponseBodyData) SetOrderNo(v string) *GlobalHotelQueryOrderResponseBodyData {
	s.OrderNo = &v
	return s
}

func (s *GlobalHotelQueryOrderResponseBodyData) SetPayment(v *GlobalHotelQueryOrderResponseBodyDataPayment) *GlobalHotelQueryOrderResponseBodyData {
	s.Payment = v
	return s
}

func (s *GlobalHotelQueryOrderResponseBodyData) SetRefundOrders(v []*GlobalHotelQueryOrderResponseBodyDataRefundOrders) *GlobalHotelQueryOrderResponseBodyData {
	s.RefundOrders = v
	return s
}

func (s *GlobalHotelQueryOrderResponseBodyData) SetRoomStays(v []*GlobalHotelQueryOrderResponseBodyDataRoomStays) *GlobalHotelQueryOrderResponseBodyData {
	s.RoomStays = v
	return s
}

func (s *GlobalHotelQueryOrderResponseBodyData) SetSalesChannel(v string) *GlobalHotelQueryOrderResponseBodyData {
	s.SalesChannel = &v
	return s
}

func (s *GlobalHotelQueryOrderResponseBodyData) SetStatus(v string) *GlobalHotelQueryOrderResponseBodyData {
	s.Status = &v
	return s
}

func (s *GlobalHotelQueryOrderResponseBodyData) SetTracerId(v string) *GlobalHotelQueryOrderResponseBodyData {
	s.TracerId = &v
	return s
}

func (s *GlobalHotelQueryOrderResponseBodyData) Validate() error {
	if s.ItemInfo != nil {
		if err := s.ItemInfo.Validate(); err != nil {
			return err
		}
	}
	if s.Payment != nil {
		if err := s.Payment.Validate(); err != nil {
			return err
		}
	}
	if s.RefundOrders != nil {
		for _, item := range s.RefundOrders {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.RoomStays != nil {
		for _, item := range s.RoomStays {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GlobalHotelQueryOrderResponseBodyDataItemInfo struct {
	// The cancellation policy.
	CancelPolicy *GlobalHotelQueryOrderResponseBodyDataItemInfoCancelPolicy `json:"CancelPolicy,omitempty" xml:"CancelPolicy,omitempty" type:"Struct"`
	// The check-in date in yyyy-MM-dd format.
	//
	// example:
	//
	// 2026-07-01
	CheckIn *string `json:"CheckIn,omitempty" xml:"CheckIn,omitempty"`
	// The number of guests checking in.
	//
	// example:
	//
	// 2
	CheckInNumber *int32 `json:"CheckInNumber,omitempty" xml:"CheckInNumber,omitempty"`
	// The check-out date in yyyy-MM-dd format.
	//
	// example:
	//
	// 2026-07-03
	CheckOut *string `json:"CheckOut,omitempty" xml:"CheckOut,omitempty"`
	// The list of nightly rates.
	DailyPrices []*GlobalHotelQueryOrderResponseBodyDataItemInfoDailyPrices `json:"DailyPrices,omitempty" xml:"DailyPrices,omitempty" type:"Repeated"`
	// The meal information.
	Meal *GlobalHotelQueryOrderResponseBodyDataItemInfoMeal `json:"Meal,omitempty" xml:"Meal,omitempty" type:"Struct"`
	// The number of rooms.
	//
	// example:
	//
	// 1
	RoomCount *int32 `json:"RoomCount,omitempty" xml:"RoomCount,omitempty"`
	// The total selling price.
	SellingTotalPrice *GlobalHotelQueryOrderResponseBodyDataItemInfoSellingTotalPrice `json:"SellingTotalPrice,omitempty" xml:"SellingTotalPrice,omitempty" type:"Struct"`
}

func (s GlobalHotelQueryOrderResponseBodyDataItemInfo) String() string {
	return dara.Prettify(s)
}

func (s GlobalHotelQueryOrderResponseBodyDataItemInfo) GoString() string {
	return s.String()
}

func (s *GlobalHotelQueryOrderResponseBodyDataItemInfo) GetCancelPolicy() *GlobalHotelQueryOrderResponseBodyDataItemInfoCancelPolicy {
	return s.CancelPolicy
}

func (s *GlobalHotelQueryOrderResponseBodyDataItemInfo) GetCheckIn() *string {
	return s.CheckIn
}

func (s *GlobalHotelQueryOrderResponseBodyDataItemInfo) GetCheckInNumber() *int32 {
	return s.CheckInNumber
}

func (s *GlobalHotelQueryOrderResponseBodyDataItemInfo) GetCheckOut() *string {
	return s.CheckOut
}

func (s *GlobalHotelQueryOrderResponseBodyDataItemInfo) GetDailyPrices() []*GlobalHotelQueryOrderResponseBodyDataItemInfoDailyPrices {
	return s.DailyPrices
}

func (s *GlobalHotelQueryOrderResponseBodyDataItemInfo) GetMeal() *GlobalHotelQueryOrderResponseBodyDataItemInfoMeal {
	return s.Meal
}

func (s *GlobalHotelQueryOrderResponseBodyDataItemInfo) GetRoomCount() *int32 {
	return s.RoomCount
}

func (s *GlobalHotelQueryOrderResponseBodyDataItemInfo) GetSellingTotalPrice() *GlobalHotelQueryOrderResponseBodyDataItemInfoSellingTotalPrice {
	return s.SellingTotalPrice
}

func (s *GlobalHotelQueryOrderResponseBodyDataItemInfo) SetCancelPolicy(v *GlobalHotelQueryOrderResponseBodyDataItemInfoCancelPolicy) *GlobalHotelQueryOrderResponseBodyDataItemInfo {
	s.CancelPolicy = v
	return s
}

func (s *GlobalHotelQueryOrderResponseBodyDataItemInfo) SetCheckIn(v string) *GlobalHotelQueryOrderResponseBodyDataItemInfo {
	s.CheckIn = &v
	return s
}

func (s *GlobalHotelQueryOrderResponseBodyDataItemInfo) SetCheckInNumber(v int32) *GlobalHotelQueryOrderResponseBodyDataItemInfo {
	s.CheckInNumber = &v
	return s
}

func (s *GlobalHotelQueryOrderResponseBodyDataItemInfo) SetCheckOut(v string) *GlobalHotelQueryOrderResponseBodyDataItemInfo {
	s.CheckOut = &v
	return s
}

func (s *GlobalHotelQueryOrderResponseBodyDataItemInfo) SetDailyPrices(v []*GlobalHotelQueryOrderResponseBodyDataItemInfoDailyPrices) *GlobalHotelQueryOrderResponseBodyDataItemInfo {
	s.DailyPrices = v
	return s
}

func (s *GlobalHotelQueryOrderResponseBodyDataItemInfo) SetMeal(v *GlobalHotelQueryOrderResponseBodyDataItemInfoMeal) *GlobalHotelQueryOrderResponseBodyDataItemInfo {
	s.Meal = v
	return s
}

func (s *GlobalHotelQueryOrderResponseBodyDataItemInfo) SetRoomCount(v int32) *GlobalHotelQueryOrderResponseBodyDataItemInfo {
	s.RoomCount = &v
	return s
}

func (s *GlobalHotelQueryOrderResponseBodyDataItemInfo) SetSellingTotalPrice(v *GlobalHotelQueryOrderResponseBodyDataItemInfoSellingTotalPrice) *GlobalHotelQueryOrderResponseBodyDataItemInfo {
	s.SellingTotalPrice = v
	return s
}

func (s *GlobalHotelQueryOrderResponseBodyDataItemInfo) Validate() error {
	if s.CancelPolicy != nil {
		if err := s.CancelPolicy.Validate(); err != nil {
			return err
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
	if s.Meal != nil {
		if err := s.Meal.Validate(); err != nil {
			return err
		}
	}
	if s.SellingTotalPrice != nil {
		if err := s.SellingTotalPrice.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GlobalHotelQueryOrderResponseBodyDataItemInfoCancelPolicy struct {
	// The list of cancellation penalties.
	Penalties []*GlobalHotelQueryOrderResponseBodyDataItemInfoCancelPolicyPenalties `json:"Penalties,omitempty" xml:"Penalties,omitempty" type:"Repeated"`
	// The cancellation policy type.
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

func (s GlobalHotelQueryOrderResponseBodyDataItemInfoCancelPolicy) String() string {
	return dara.Prettify(s)
}

func (s GlobalHotelQueryOrderResponseBodyDataItemInfoCancelPolicy) GoString() string {
	return s.String()
}

func (s *GlobalHotelQueryOrderResponseBodyDataItemInfoCancelPolicy) GetPenalties() []*GlobalHotelQueryOrderResponseBodyDataItemInfoCancelPolicyPenalties {
	return s.Penalties
}

func (s *GlobalHotelQueryOrderResponseBodyDataItemInfoCancelPolicy) GetPolicyType() *string {
	return s.PolicyType
}

func (s *GlobalHotelQueryOrderResponseBodyDataItemInfoCancelPolicy) GetTracerId() *string {
	return s.TracerId
}

func (s *GlobalHotelQueryOrderResponseBodyDataItemInfoCancelPolicy) SetPenalties(v []*GlobalHotelQueryOrderResponseBodyDataItemInfoCancelPolicyPenalties) *GlobalHotelQueryOrderResponseBodyDataItemInfoCancelPolicy {
	s.Penalties = v
	return s
}

func (s *GlobalHotelQueryOrderResponseBodyDataItemInfoCancelPolicy) SetPolicyType(v string) *GlobalHotelQueryOrderResponseBodyDataItemInfoCancelPolicy {
	s.PolicyType = &v
	return s
}

func (s *GlobalHotelQueryOrderResponseBodyDataItemInfoCancelPolicy) SetTracerId(v string) *GlobalHotelQueryOrderResponseBodyDataItemInfoCancelPolicy {
	s.TracerId = &v
	return s
}

func (s *GlobalHotelQueryOrderResponseBodyDataItemInfoCancelPolicy) Validate() error {
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

type GlobalHotelQueryOrderResponseBodyDataItemInfoCancelPolicyPenalties struct {
	// The currency code. This parameter is valid only when the penalty type is AMOUNT.
	//
	// example:
	//
	// USD
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// The effective end time in UTC millisecond timestamp.
	//
	// example:
	//
	// 1672617600000
	End *int64 `json:"End,omitempty" xml:"End,omitempty"`
	// The penalty type.
	//
	// example:
	//
	// PERCENTAGE
	PenaltyType *string `json:"PenaltyType,omitempty" xml:"PenaltyType,omitempty"`
	// The penalty value, which can be a percentage, amount, or number of nights.
	//
	// example:
	//
	// 50
	PenaltyValue *string `json:"PenaltyValue,omitempty" xml:"PenaltyValue,omitempty"`
	// The effective start time in UTC millisecond timestamp.
	//
	// example:
	//
	// 1672531200000
	Start *int64 `json:"Start,omitempty" xml:"Start,omitempty"`
	// TracerId
	//
	// example:
	//
	// TracerId
	TracerId *string `json:"TracerId,omitempty" xml:"TracerId,omitempty"`
}

func (s GlobalHotelQueryOrderResponseBodyDataItemInfoCancelPolicyPenalties) String() string {
	return dara.Prettify(s)
}

func (s GlobalHotelQueryOrderResponseBodyDataItemInfoCancelPolicyPenalties) GoString() string {
	return s.String()
}

func (s *GlobalHotelQueryOrderResponseBodyDataItemInfoCancelPolicyPenalties) GetCurrency() *string {
	return s.Currency
}

func (s *GlobalHotelQueryOrderResponseBodyDataItemInfoCancelPolicyPenalties) GetEnd() *int64 {
	return s.End
}

func (s *GlobalHotelQueryOrderResponseBodyDataItemInfoCancelPolicyPenalties) GetPenaltyType() *string {
	return s.PenaltyType
}

func (s *GlobalHotelQueryOrderResponseBodyDataItemInfoCancelPolicyPenalties) GetPenaltyValue() *string {
	return s.PenaltyValue
}

func (s *GlobalHotelQueryOrderResponseBodyDataItemInfoCancelPolicyPenalties) GetStart() *int64 {
	return s.Start
}

func (s *GlobalHotelQueryOrderResponseBodyDataItemInfoCancelPolicyPenalties) GetTracerId() *string {
	return s.TracerId
}

func (s *GlobalHotelQueryOrderResponseBodyDataItemInfoCancelPolicyPenalties) SetCurrency(v string) *GlobalHotelQueryOrderResponseBodyDataItemInfoCancelPolicyPenalties {
	s.Currency = &v
	return s
}

func (s *GlobalHotelQueryOrderResponseBodyDataItemInfoCancelPolicyPenalties) SetEnd(v int64) *GlobalHotelQueryOrderResponseBodyDataItemInfoCancelPolicyPenalties {
	s.End = &v
	return s
}

func (s *GlobalHotelQueryOrderResponseBodyDataItemInfoCancelPolicyPenalties) SetPenaltyType(v string) *GlobalHotelQueryOrderResponseBodyDataItemInfoCancelPolicyPenalties {
	s.PenaltyType = &v
	return s
}

func (s *GlobalHotelQueryOrderResponseBodyDataItemInfoCancelPolicyPenalties) SetPenaltyValue(v string) *GlobalHotelQueryOrderResponseBodyDataItemInfoCancelPolicyPenalties {
	s.PenaltyValue = &v
	return s
}

func (s *GlobalHotelQueryOrderResponseBodyDataItemInfoCancelPolicyPenalties) SetStart(v int64) *GlobalHotelQueryOrderResponseBodyDataItemInfoCancelPolicyPenalties {
	s.Start = &v
	return s
}

func (s *GlobalHotelQueryOrderResponseBodyDataItemInfoCancelPolicyPenalties) SetTracerId(v string) *GlobalHotelQueryOrderResponseBodyDataItemInfoCancelPolicyPenalties {
	s.TracerId = &v
	return s
}

func (s *GlobalHotelQueryOrderResponseBodyDataItemInfoCancelPolicyPenalties) Validate() error {
	return dara.Validate(s)
}

type GlobalHotelQueryOrderResponseBodyDataItemInfoDailyPrices struct {
	// LocalDate
	//
	// example:
	//
	// LocalDate
	Date *string `json:"Date,omitempty" xml:"Date,omitempty"`
	// The price.
	Price *GlobalHotelQueryOrderResponseBodyDataItemInfoDailyPricesPrice `json:"Price,omitempty" xml:"Price,omitempty" type:"Struct"`
}

func (s GlobalHotelQueryOrderResponseBodyDataItemInfoDailyPrices) String() string {
	return dara.Prettify(s)
}

func (s GlobalHotelQueryOrderResponseBodyDataItemInfoDailyPrices) GoString() string {
	return s.String()
}

func (s *GlobalHotelQueryOrderResponseBodyDataItemInfoDailyPrices) GetDate() *string {
	return s.Date
}

func (s *GlobalHotelQueryOrderResponseBodyDataItemInfoDailyPrices) GetPrice() *GlobalHotelQueryOrderResponseBodyDataItemInfoDailyPricesPrice {
	return s.Price
}

func (s *GlobalHotelQueryOrderResponseBodyDataItemInfoDailyPrices) SetDate(v string) *GlobalHotelQueryOrderResponseBodyDataItemInfoDailyPrices {
	s.Date = &v
	return s
}

func (s *GlobalHotelQueryOrderResponseBodyDataItemInfoDailyPrices) SetPrice(v *GlobalHotelQueryOrderResponseBodyDataItemInfoDailyPricesPrice) *GlobalHotelQueryOrderResponseBodyDataItemInfoDailyPrices {
	s.Price = v
	return s
}

func (s *GlobalHotelQueryOrderResponseBodyDataItemInfoDailyPrices) Validate() error {
	if s.Price != nil {
		if err := s.Price.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GlobalHotelQueryOrderResponseBodyDataItemInfoDailyPricesPrice struct {
	// cent
	//
	// example:
	//
	// 1
	Cent *int64 `json:"Cent,omitempty" xml:"Cent,omitempty"`
	// The currency.
	Currency *GlobalHotelQueryOrderResponseBodyDataItemInfoDailyPricesPriceCurrency `json:"Currency,omitempty" xml:"Currency,omitempty" type:"Struct"`
}

func (s GlobalHotelQueryOrderResponseBodyDataItemInfoDailyPricesPrice) String() string {
	return dara.Prettify(s)
}

func (s GlobalHotelQueryOrderResponseBodyDataItemInfoDailyPricesPrice) GoString() string {
	return s.String()
}

func (s *GlobalHotelQueryOrderResponseBodyDataItemInfoDailyPricesPrice) GetCent() *int64 {
	return s.Cent
}

func (s *GlobalHotelQueryOrderResponseBodyDataItemInfoDailyPricesPrice) GetCurrency() *GlobalHotelQueryOrderResponseBodyDataItemInfoDailyPricesPriceCurrency {
	return s.Currency
}

func (s *GlobalHotelQueryOrderResponseBodyDataItemInfoDailyPricesPrice) SetCent(v int64) *GlobalHotelQueryOrderResponseBodyDataItemInfoDailyPricesPrice {
	s.Cent = &v
	return s
}

func (s *GlobalHotelQueryOrderResponseBodyDataItemInfoDailyPricesPrice) SetCurrency(v *GlobalHotelQueryOrderResponseBodyDataItemInfoDailyPricesPriceCurrency) *GlobalHotelQueryOrderResponseBodyDataItemInfoDailyPricesPrice {
	s.Currency = v
	return s
}

func (s *GlobalHotelQueryOrderResponseBodyDataItemInfoDailyPricesPrice) Validate() error {
	if s.Currency != nil {
		if err := s.Currency.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GlobalHotelQueryOrderResponseBodyDataItemInfoDailyPricesPriceCurrency struct {
	// The currency code.
	//
	// example:
	//
	// USD
	CurrencyCode *string `json:"CurrencyCode,omitempty" xml:"CurrencyCode,omitempty"`
	// DefaultFractionDigits
	//
	// example:
	//
	// 1
	DefaultFractionDigits *int32 `json:"DefaultFractionDigits,omitempty" xml:"DefaultFractionDigits,omitempty"`
	// NumericCode
	//
	// example:
	//
	// 1
	NumericCode *int32 `json:"NumericCode,omitempty" xml:"NumericCode,omitempty"`
}

func (s GlobalHotelQueryOrderResponseBodyDataItemInfoDailyPricesPriceCurrency) String() string {
	return dara.Prettify(s)
}

func (s GlobalHotelQueryOrderResponseBodyDataItemInfoDailyPricesPriceCurrency) GoString() string {
	return s.String()
}

func (s *GlobalHotelQueryOrderResponseBodyDataItemInfoDailyPricesPriceCurrency) GetCurrencyCode() *string {
	return s.CurrencyCode
}

func (s *GlobalHotelQueryOrderResponseBodyDataItemInfoDailyPricesPriceCurrency) GetDefaultFractionDigits() *int32 {
	return s.DefaultFractionDigits
}

func (s *GlobalHotelQueryOrderResponseBodyDataItemInfoDailyPricesPriceCurrency) GetNumericCode() *int32 {
	return s.NumericCode
}

func (s *GlobalHotelQueryOrderResponseBodyDataItemInfoDailyPricesPriceCurrency) SetCurrencyCode(v string) *GlobalHotelQueryOrderResponseBodyDataItemInfoDailyPricesPriceCurrency {
	s.CurrencyCode = &v
	return s
}

func (s *GlobalHotelQueryOrderResponseBodyDataItemInfoDailyPricesPriceCurrency) SetDefaultFractionDigits(v int32) *GlobalHotelQueryOrderResponseBodyDataItemInfoDailyPricesPriceCurrency {
	s.DefaultFractionDigits = &v
	return s
}

func (s *GlobalHotelQueryOrderResponseBodyDataItemInfoDailyPricesPriceCurrency) SetNumericCode(v int32) *GlobalHotelQueryOrderResponseBodyDataItemInfoDailyPricesPriceCurrency {
	s.NumericCode = &v
	return s
}

func (s *GlobalHotelQueryOrderResponseBodyDataItemInfoDailyPricesPriceCurrency) Validate() error {
	return dara.Validate(s)
}

type GlobalHotelQueryOrderResponseBodyDataItemInfoMeal struct {
	// The description.
	//
	// example:
	//
	// Breakfast included
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The meal type.
	//
	// example:
	//
	// BREAKFAST
	MealType *string `json:"MealType,omitempty" xml:"MealType,omitempty"`
	// TracerId
	//
	// example:
	//
	// TracerId
	TracerId *string `json:"TracerId,omitempty" xml:"TracerId,omitempty"`
}

func (s GlobalHotelQueryOrderResponseBodyDataItemInfoMeal) String() string {
	return dara.Prettify(s)
}

func (s GlobalHotelQueryOrderResponseBodyDataItemInfoMeal) GoString() string {
	return s.String()
}

func (s *GlobalHotelQueryOrderResponseBodyDataItemInfoMeal) GetDescription() *string {
	return s.Description
}

func (s *GlobalHotelQueryOrderResponseBodyDataItemInfoMeal) GetMealType() *string {
	return s.MealType
}

func (s *GlobalHotelQueryOrderResponseBodyDataItemInfoMeal) GetTracerId() *string {
	return s.TracerId
}

func (s *GlobalHotelQueryOrderResponseBodyDataItemInfoMeal) SetDescription(v string) *GlobalHotelQueryOrderResponseBodyDataItemInfoMeal {
	s.Description = &v
	return s
}

func (s *GlobalHotelQueryOrderResponseBodyDataItemInfoMeal) SetMealType(v string) *GlobalHotelQueryOrderResponseBodyDataItemInfoMeal {
	s.MealType = &v
	return s
}

func (s *GlobalHotelQueryOrderResponseBodyDataItemInfoMeal) SetTracerId(v string) *GlobalHotelQueryOrderResponseBodyDataItemInfoMeal {
	s.TracerId = &v
	return s
}

func (s *GlobalHotelQueryOrderResponseBodyDataItemInfoMeal) Validate() error {
	return dara.Validate(s)
}

type GlobalHotelQueryOrderResponseBodyDataItemInfoSellingTotalPrice struct {
	// The amount in the smallest currency unit.
	//
	// example:
	//
	// 10000
	Amount *string `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// The currency code in ISO 4217 format.
	//
	// example:
	//
	// USD
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// TracerId
	//
	// example:
	//
	// TracerId
	TracerId *string `json:"TracerId,omitempty" xml:"TracerId,omitempty"`
}

func (s GlobalHotelQueryOrderResponseBodyDataItemInfoSellingTotalPrice) String() string {
	return dara.Prettify(s)
}

func (s GlobalHotelQueryOrderResponseBodyDataItemInfoSellingTotalPrice) GoString() string {
	return s.String()
}

func (s *GlobalHotelQueryOrderResponseBodyDataItemInfoSellingTotalPrice) GetAmount() *string {
	return s.Amount
}

func (s *GlobalHotelQueryOrderResponseBodyDataItemInfoSellingTotalPrice) GetCurrency() *string {
	return s.Currency
}

func (s *GlobalHotelQueryOrderResponseBodyDataItemInfoSellingTotalPrice) GetTracerId() *string {
	return s.TracerId
}

func (s *GlobalHotelQueryOrderResponseBodyDataItemInfoSellingTotalPrice) SetAmount(v string) *GlobalHotelQueryOrderResponseBodyDataItemInfoSellingTotalPrice {
	s.Amount = &v
	return s
}

func (s *GlobalHotelQueryOrderResponseBodyDataItemInfoSellingTotalPrice) SetCurrency(v string) *GlobalHotelQueryOrderResponseBodyDataItemInfoSellingTotalPrice {
	s.Currency = &v
	return s
}

func (s *GlobalHotelQueryOrderResponseBodyDataItemInfoSellingTotalPrice) SetTracerId(v string) *GlobalHotelQueryOrderResponseBodyDataItemInfoSellingTotalPrice {
	s.TracerId = &v
	return s
}

func (s *GlobalHotelQueryOrderResponseBodyDataItemInfoSellingTotalPrice) Validate() error {
	return dara.Validate(s)
}

type GlobalHotelQueryOrderResponseBodyDataPayment struct {
	// The payment amount.
	Amount *GlobalHotelQueryOrderResponseBodyDataPaymentAmount `json:"Amount,omitempty" xml:"Amount,omitempty" type:"Struct"`
	// The payment completion time in UTC millisecond timestamp.
	//
	// example:
	//
	// 1672531200000
	GmtPaid *int64 `json:"GmtPaid,omitempty" xml:"GmtPaid,omitempty"`
	// The payment method.
	//
	// example:
	//
	// BALANCE
	PaymentMethod *string `json:"PaymentMethod,omitempty" xml:"PaymentMethod,omitempty"`
	// The payment transaction ID.
	//
	// example:
	//
	// PAY202606290001
	PaymentTransactionId *string `json:"PaymentTransactionId,omitempty" xml:"PaymentTransactionId,omitempty"`
}

func (s GlobalHotelQueryOrderResponseBodyDataPayment) String() string {
	return dara.Prettify(s)
}

func (s GlobalHotelQueryOrderResponseBodyDataPayment) GoString() string {
	return s.String()
}

func (s *GlobalHotelQueryOrderResponseBodyDataPayment) GetAmount() *GlobalHotelQueryOrderResponseBodyDataPaymentAmount {
	return s.Amount
}

func (s *GlobalHotelQueryOrderResponseBodyDataPayment) GetGmtPaid() *int64 {
	return s.GmtPaid
}

func (s *GlobalHotelQueryOrderResponseBodyDataPayment) GetPaymentMethod() *string {
	return s.PaymentMethod
}

func (s *GlobalHotelQueryOrderResponseBodyDataPayment) GetPaymentTransactionId() *string {
	return s.PaymentTransactionId
}

func (s *GlobalHotelQueryOrderResponseBodyDataPayment) SetAmount(v *GlobalHotelQueryOrderResponseBodyDataPaymentAmount) *GlobalHotelQueryOrderResponseBodyDataPayment {
	s.Amount = v
	return s
}

func (s *GlobalHotelQueryOrderResponseBodyDataPayment) SetGmtPaid(v int64) *GlobalHotelQueryOrderResponseBodyDataPayment {
	s.GmtPaid = &v
	return s
}

func (s *GlobalHotelQueryOrderResponseBodyDataPayment) SetPaymentMethod(v string) *GlobalHotelQueryOrderResponseBodyDataPayment {
	s.PaymentMethod = &v
	return s
}

func (s *GlobalHotelQueryOrderResponseBodyDataPayment) SetPaymentTransactionId(v string) *GlobalHotelQueryOrderResponseBodyDataPayment {
	s.PaymentTransactionId = &v
	return s
}

func (s *GlobalHotelQueryOrderResponseBodyDataPayment) Validate() error {
	if s.Amount != nil {
		if err := s.Amount.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GlobalHotelQueryOrderResponseBodyDataPaymentAmount struct {
	// The amount in the smallest currency unit.
	//
	// example:
	//
	// 10000
	Amount *string `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// The currency code in ISO 4217 format.
	//
	// example:
	//
	// USD
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// TracerId
	//
	// example:
	//
	// TracerId
	TracerId *string `json:"TracerId,omitempty" xml:"TracerId,omitempty"`
}

func (s GlobalHotelQueryOrderResponseBodyDataPaymentAmount) String() string {
	return dara.Prettify(s)
}

func (s GlobalHotelQueryOrderResponseBodyDataPaymentAmount) GoString() string {
	return s.String()
}

func (s *GlobalHotelQueryOrderResponseBodyDataPaymentAmount) GetAmount() *string {
	return s.Amount
}

func (s *GlobalHotelQueryOrderResponseBodyDataPaymentAmount) GetCurrency() *string {
	return s.Currency
}

func (s *GlobalHotelQueryOrderResponseBodyDataPaymentAmount) GetTracerId() *string {
	return s.TracerId
}

func (s *GlobalHotelQueryOrderResponseBodyDataPaymentAmount) SetAmount(v string) *GlobalHotelQueryOrderResponseBodyDataPaymentAmount {
	s.Amount = &v
	return s
}

func (s *GlobalHotelQueryOrderResponseBodyDataPaymentAmount) SetCurrency(v string) *GlobalHotelQueryOrderResponseBodyDataPaymentAmount {
	s.Currency = &v
	return s
}

func (s *GlobalHotelQueryOrderResponseBodyDataPaymentAmount) SetTracerId(v string) *GlobalHotelQueryOrderResponseBodyDataPaymentAmount {
	s.TracerId = &v
	return s
}

func (s *GlobalHotelQueryOrderResponseBodyDataPaymentAmount) Validate() error {
	return dara.Validate(s)
}

type GlobalHotelQueryOrderResponseBodyDataRefundOrders struct {
	// The creation time of the refund order, in UTC millisecond timestamp.
	//
	// example:
	//
	// 1672531200000
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The refund transaction ID.
	//
	// example:
	//
	// TXN001
	RefundTransactionId *string `json:"RefundTransactionId,omitempty" xml:"RefundTransactionId,omitempty"`
	// The reason for rejection.
	//
	// example:
	//
	// Supplier rejected
	RejectReason *string `json:"RejectReason,omitempty" xml:"RejectReason,omitempty"`
	// The external refund order number.
	//
	// example:
	//
	// RF202606290001
	SellRefundOrderNo *string `json:"SellRefundOrderNo,omitempty" xml:"SellRefundOrderNo,omitempty"`
	// The unified refund status.
	//
	// example:
	//
	// REFUNDED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The penalty amount on the sales side.
	TotalPenaltyAmount *GlobalHotelQueryOrderResponseBodyDataRefundOrdersTotalPenaltyAmount `json:"TotalPenaltyAmount,omitempty" xml:"TotalPenaltyAmount,omitempty" type:"Struct"`
	// The actual refund amount.
	TotalRefundAmount *GlobalHotelQueryOrderResponseBodyDataRefundOrdersTotalRefundAmount `json:"TotalRefundAmount,omitempty" xml:"TotalRefundAmount,omitempty" type:"Struct"`
}

func (s GlobalHotelQueryOrderResponseBodyDataRefundOrders) String() string {
	return dara.Prettify(s)
}

func (s GlobalHotelQueryOrderResponseBodyDataRefundOrders) GoString() string {
	return s.String()
}

func (s *GlobalHotelQueryOrderResponseBodyDataRefundOrders) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *GlobalHotelQueryOrderResponseBodyDataRefundOrders) GetRefundTransactionId() *string {
	return s.RefundTransactionId
}

func (s *GlobalHotelQueryOrderResponseBodyDataRefundOrders) GetRejectReason() *string {
	return s.RejectReason
}

func (s *GlobalHotelQueryOrderResponseBodyDataRefundOrders) GetSellRefundOrderNo() *string {
	return s.SellRefundOrderNo
}

func (s *GlobalHotelQueryOrderResponseBodyDataRefundOrders) GetStatus() *string {
	return s.Status
}

func (s *GlobalHotelQueryOrderResponseBodyDataRefundOrders) GetTotalPenaltyAmount() *GlobalHotelQueryOrderResponseBodyDataRefundOrdersTotalPenaltyAmount {
	return s.TotalPenaltyAmount
}

func (s *GlobalHotelQueryOrderResponseBodyDataRefundOrders) GetTotalRefundAmount() *GlobalHotelQueryOrderResponseBodyDataRefundOrdersTotalRefundAmount {
	return s.TotalRefundAmount
}

func (s *GlobalHotelQueryOrderResponseBodyDataRefundOrders) SetGmtCreate(v int64) *GlobalHotelQueryOrderResponseBodyDataRefundOrders {
	s.GmtCreate = &v
	return s
}

func (s *GlobalHotelQueryOrderResponseBodyDataRefundOrders) SetRefundTransactionId(v string) *GlobalHotelQueryOrderResponseBodyDataRefundOrders {
	s.RefundTransactionId = &v
	return s
}

func (s *GlobalHotelQueryOrderResponseBodyDataRefundOrders) SetRejectReason(v string) *GlobalHotelQueryOrderResponseBodyDataRefundOrders {
	s.RejectReason = &v
	return s
}

func (s *GlobalHotelQueryOrderResponseBodyDataRefundOrders) SetSellRefundOrderNo(v string) *GlobalHotelQueryOrderResponseBodyDataRefundOrders {
	s.SellRefundOrderNo = &v
	return s
}

func (s *GlobalHotelQueryOrderResponseBodyDataRefundOrders) SetStatus(v string) *GlobalHotelQueryOrderResponseBodyDataRefundOrders {
	s.Status = &v
	return s
}

func (s *GlobalHotelQueryOrderResponseBodyDataRefundOrders) SetTotalPenaltyAmount(v *GlobalHotelQueryOrderResponseBodyDataRefundOrdersTotalPenaltyAmount) *GlobalHotelQueryOrderResponseBodyDataRefundOrders {
	s.TotalPenaltyAmount = v
	return s
}

func (s *GlobalHotelQueryOrderResponseBodyDataRefundOrders) SetTotalRefundAmount(v *GlobalHotelQueryOrderResponseBodyDataRefundOrdersTotalRefundAmount) *GlobalHotelQueryOrderResponseBodyDataRefundOrders {
	s.TotalRefundAmount = v
	return s
}

func (s *GlobalHotelQueryOrderResponseBodyDataRefundOrders) Validate() error {
	if s.TotalPenaltyAmount != nil {
		if err := s.TotalPenaltyAmount.Validate(); err != nil {
			return err
		}
	}
	if s.TotalRefundAmount != nil {
		if err := s.TotalRefundAmount.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GlobalHotelQueryOrderResponseBodyDataRefundOrdersTotalPenaltyAmount struct {
	// The amount in the smallest currency unit.
	//
	// example:
	//
	// 10000
	Amount *string `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// The currency code in ISO 4217 format.
	//
	// example:
	//
	// USD
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// TraceId
	//
	// example:
	//
	// TraceId
	TracerId *string `json:"TracerId,omitempty" xml:"TracerId,omitempty"`
}

func (s GlobalHotelQueryOrderResponseBodyDataRefundOrdersTotalPenaltyAmount) String() string {
	return dara.Prettify(s)
}

func (s GlobalHotelQueryOrderResponseBodyDataRefundOrdersTotalPenaltyAmount) GoString() string {
	return s.String()
}

func (s *GlobalHotelQueryOrderResponseBodyDataRefundOrdersTotalPenaltyAmount) GetAmount() *string {
	return s.Amount
}

func (s *GlobalHotelQueryOrderResponseBodyDataRefundOrdersTotalPenaltyAmount) GetCurrency() *string {
	return s.Currency
}

func (s *GlobalHotelQueryOrderResponseBodyDataRefundOrdersTotalPenaltyAmount) GetTracerId() *string {
	return s.TracerId
}

func (s *GlobalHotelQueryOrderResponseBodyDataRefundOrdersTotalPenaltyAmount) SetAmount(v string) *GlobalHotelQueryOrderResponseBodyDataRefundOrdersTotalPenaltyAmount {
	s.Amount = &v
	return s
}

func (s *GlobalHotelQueryOrderResponseBodyDataRefundOrdersTotalPenaltyAmount) SetCurrency(v string) *GlobalHotelQueryOrderResponseBodyDataRefundOrdersTotalPenaltyAmount {
	s.Currency = &v
	return s
}

func (s *GlobalHotelQueryOrderResponseBodyDataRefundOrdersTotalPenaltyAmount) SetTracerId(v string) *GlobalHotelQueryOrderResponseBodyDataRefundOrdersTotalPenaltyAmount {
	s.TracerId = &v
	return s
}

func (s *GlobalHotelQueryOrderResponseBodyDataRefundOrdersTotalPenaltyAmount) Validate() error {
	return dara.Validate(s)
}

type GlobalHotelQueryOrderResponseBodyDataRefundOrdersTotalRefundAmount struct {
	// The amount in the smallest currency unit.
	//
	// example:
	//
	// 10000
	Amount *string `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// The currency code in ISO 4217 format.
	//
	// example:
	//
	// USD
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// TraceId
	//
	// example:
	//
	// TraceId
	TracerId *string `json:"TracerId,omitempty" xml:"TracerId,omitempty"`
}

func (s GlobalHotelQueryOrderResponseBodyDataRefundOrdersTotalRefundAmount) String() string {
	return dara.Prettify(s)
}

func (s GlobalHotelQueryOrderResponseBodyDataRefundOrdersTotalRefundAmount) GoString() string {
	return s.String()
}

func (s *GlobalHotelQueryOrderResponseBodyDataRefundOrdersTotalRefundAmount) GetAmount() *string {
	return s.Amount
}

func (s *GlobalHotelQueryOrderResponseBodyDataRefundOrdersTotalRefundAmount) GetCurrency() *string {
	return s.Currency
}

func (s *GlobalHotelQueryOrderResponseBodyDataRefundOrdersTotalRefundAmount) GetTracerId() *string {
	return s.TracerId
}

func (s *GlobalHotelQueryOrderResponseBodyDataRefundOrdersTotalRefundAmount) SetAmount(v string) *GlobalHotelQueryOrderResponseBodyDataRefundOrdersTotalRefundAmount {
	s.Amount = &v
	return s
}

func (s *GlobalHotelQueryOrderResponseBodyDataRefundOrdersTotalRefundAmount) SetCurrency(v string) *GlobalHotelQueryOrderResponseBodyDataRefundOrdersTotalRefundAmount {
	s.Currency = &v
	return s
}

func (s *GlobalHotelQueryOrderResponseBodyDataRefundOrdersTotalRefundAmount) SetTracerId(v string) *GlobalHotelQueryOrderResponseBodyDataRefundOrdersTotalRefundAmount {
	s.TracerId = &v
	return s
}

func (s *GlobalHotelQueryOrderResponseBodyDataRefundOrdersTotalRefundAmount) Validate() error {
	return dara.Validate(s)
}

type GlobalHotelQueryOrderResponseBodyDataRoomStays struct {
	// The room confirmation ID.
	//
	// example:
	//
	// CONF001
	ConfirmationId *string `json:"ConfirmationId,omitempty" xml:"ConfirmationId,omitempty"`
	// The list of guests.
	Guests []*GlobalHotelQueryOrderResponseBodyDataRoomStaysGuests `json:"Guests,omitempty" xml:"Guests,omitempty" type:"Repeated"`
	// The room index, starting from 1.
	//
	// example:
	//
	// 1
	RoomIndex *int32 `json:"RoomIndex,omitempty" xml:"RoomIndex,omitempty"`
	// The delivery status. Valid values: PENDING_CHECKIN, CHECKED_IN, CHECKED_OUT, and CANCELLED. The value is null before the delivery is created.
	//
	// example:
	//
	// CHECKED_IN
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GlobalHotelQueryOrderResponseBodyDataRoomStays) String() string {
	return dara.Prettify(s)
}

func (s GlobalHotelQueryOrderResponseBodyDataRoomStays) GoString() string {
	return s.String()
}

func (s *GlobalHotelQueryOrderResponseBodyDataRoomStays) GetConfirmationId() *string {
	return s.ConfirmationId
}

func (s *GlobalHotelQueryOrderResponseBodyDataRoomStays) GetGuests() []*GlobalHotelQueryOrderResponseBodyDataRoomStaysGuests {
	return s.Guests
}

func (s *GlobalHotelQueryOrderResponseBodyDataRoomStays) GetRoomIndex() *int32 {
	return s.RoomIndex
}

func (s *GlobalHotelQueryOrderResponseBodyDataRoomStays) GetStatus() *string {
	return s.Status
}

func (s *GlobalHotelQueryOrderResponseBodyDataRoomStays) SetConfirmationId(v string) *GlobalHotelQueryOrderResponseBodyDataRoomStays {
	s.ConfirmationId = &v
	return s
}

func (s *GlobalHotelQueryOrderResponseBodyDataRoomStays) SetGuests(v []*GlobalHotelQueryOrderResponseBodyDataRoomStaysGuests) *GlobalHotelQueryOrderResponseBodyDataRoomStays {
	s.Guests = v
	return s
}

func (s *GlobalHotelQueryOrderResponseBodyDataRoomStays) SetRoomIndex(v int32) *GlobalHotelQueryOrderResponseBodyDataRoomStays {
	s.RoomIndex = &v
	return s
}

func (s *GlobalHotelQueryOrderResponseBodyDataRoomStays) SetStatus(v string) *GlobalHotelQueryOrderResponseBodyDataRoomStays {
	s.Status = &v
	return s
}

func (s *GlobalHotelQueryOrderResponseBodyDataRoomStays) Validate() error {
	if s.Guests != nil {
		for _, item := range s.Guests {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GlobalHotelQueryOrderResponseBodyDataRoomStaysGuests struct {
	// The first name of the guest.
	//
	// example:
	//
	// John
	FirstName *string `json:"FirstName,omitempty" xml:"FirstName,omitempty"`
	// The last name of the guest.
	//
	// example:
	//
	// Doe
	LastName *string `json:"LastName,omitempty" xml:"LastName,omitempty"`
	// TraceId
	//
	// example:
	//
	// TraceId
	TracerId *string `json:"TracerId,omitempty" xml:"TracerId,omitempty"`
}

func (s GlobalHotelQueryOrderResponseBodyDataRoomStaysGuests) String() string {
	return dara.Prettify(s)
}

func (s GlobalHotelQueryOrderResponseBodyDataRoomStaysGuests) GoString() string {
	return s.String()
}

func (s *GlobalHotelQueryOrderResponseBodyDataRoomStaysGuests) GetFirstName() *string {
	return s.FirstName
}

func (s *GlobalHotelQueryOrderResponseBodyDataRoomStaysGuests) GetLastName() *string {
	return s.LastName
}

func (s *GlobalHotelQueryOrderResponseBodyDataRoomStaysGuests) GetTracerId() *string {
	return s.TracerId
}

func (s *GlobalHotelQueryOrderResponseBodyDataRoomStaysGuests) SetFirstName(v string) *GlobalHotelQueryOrderResponseBodyDataRoomStaysGuests {
	s.FirstName = &v
	return s
}

func (s *GlobalHotelQueryOrderResponseBodyDataRoomStaysGuests) SetLastName(v string) *GlobalHotelQueryOrderResponseBodyDataRoomStaysGuests {
	s.LastName = &v
	return s
}

func (s *GlobalHotelQueryOrderResponseBodyDataRoomStaysGuests) SetTracerId(v string) *GlobalHotelQueryOrderResponseBodyDataRoomStaysGuests {
	s.TracerId = &v
	return s
}

func (s *GlobalHotelQueryOrderResponseBodyDataRoomStaysGuests) Validate() error {
	return dara.Validate(s)
}
