// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryOrderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *QueryOrderResponseBodyData) *QueryOrderResponseBody
	GetData() *QueryOrderResponseBodyData
	SetErrorCode(v string) *QueryOrderResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *QueryOrderResponseBody
	GetErrorMsg() *string
	SetRequestId(v string) *QueryOrderResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryOrderResponseBody
	GetSuccess() *bool
	SetTracerId(v string) *QueryOrderResponseBody
	GetTracerId() *string
}

type QueryOrderResponseBody struct {
	Data *QueryOrderResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// TracerId
	TracerId *string `json:"TracerId,omitempty" xml:"TracerId,omitempty"`
}

func (s QueryOrderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryOrderResponseBody) GoString() string {
	return s.String()
}

func (s *QueryOrderResponseBody) GetData() *QueryOrderResponseBodyData {
	return s.Data
}

func (s *QueryOrderResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *QueryOrderResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *QueryOrderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryOrderResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryOrderResponseBody) GetTracerId() *string {
	return s.TracerId
}

func (s *QueryOrderResponseBody) SetData(v *QueryOrderResponseBodyData) *QueryOrderResponseBody {
	s.Data = v
	return s
}

func (s *QueryOrderResponseBody) SetErrorCode(v string) *QueryOrderResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *QueryOrderResponseBody) SetErrorMsg(v string) *QueryOrderResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *QueryOrderResponseBody) SetRequestId(v string) *QueryOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryOrderResponseBody) SetSuccess(v bool) *QueryOrderResponseBody {
	s.Success = &v
	return s
}

func (s *QueryOrderResponseBody) SetTracerId(v string) *QueryOrderResponseBody {
	s.TracerId = &v
	return s
}

func (s *QueryOrderResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryOrderResponseBodyData struct {
	// example:
	//
	// 123456
	BuyerId *string `json:"BuyerId,omitempty" xml:"BuyerId,omitempty"`
	// example:
	//
	// EXT_ORDER_001
	ExternalOrderNo *string `json:"ExternalOrderNo,omitempty" xml:"ExternalOrderNo,omitempty"`
	// example:
	//
	// 1672531200000
	GmtCreate *int64                              `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	ItemInfo  *QueryOrderResponseBodyDataItemInfo `json:"ItemInfo,omitempty" xml:"ItemInfo,omitempty" type:"Struct"`
	// example:
	//
	// SO202606290001
	OrderNo      *string                                   `json:"OrderNo,omitempty" xml:"OrderNo,omitempty"`
	Payment      *QueryOrderResponseBodyDataPayment        `json:"Payment,omitempty" xml:"Payment,omitempty" type:"Struct"`
	RefundOrders []*QueryOrderResponseBodyDataRefundOrders `json:"RefundOrders,omitempty" xml:"RefundOrders,omitempty" type:"Repeated"`
	RoomStays    []*QueryOrderResponseBodyDataRoomStays    `json:"RoomStays,omitempty" xml:"RoomStays,omitempty" type:"Repeated"`
	// example:
	//
	// POP
	SalesChannel *string `json:"SalesChannel,omitempty" xml:"SalesChannel,omitempty"`
	// example:
	//
	// CONFIRMED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// TracerId
	TracerId *string `json:"TracerId,omitempty" xml:"TracerId,omitempty"`
}

func (s QueryOrderResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryOrderResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryOrderResponseBodyData) GetBuyerId() *string {
	return s.BuyerId
}

func (s *QueryOrderResponseBodyData) GetExternalOrderNo() *string {
	return s.ExternalOrderNo
}

func (s *QueryOrderResponseBodyData) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *QueryOrderResponseBodyData) GetItemInfo() *QueryOrderResponseBodyDataItemInfo {
	return s.ItemInfo
}

func (s *QueryOrderResponseBodyData) GetOrderNo() *string {
	return s.OrderNo
}

func (s *QueryOrderResponseBodyData) GetPayment() *QueryOrderResponseBodyDataPayment {
	return s.Payment
}

func (s *QueryOrderResponseBodyData) GetRefundOrders() []*QueryOrderResponseBodyDataRefundOrders {
	return s.RefundOrders
}

func (s *QueryOrderResponseBodyData) GetRoomStays() []*QueryOrderResponseBodyDataRoomStays {
	return s.RoomStays
}

func (s *QueryOrderResponseBodyData) GetSalesChannel() *string {
	return s.SalesChannel
}

func (s *QueryOrderResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *QueryOrderResponseBodyData) GetTracerId() *string {
	return s.TracerId
}

func (s *QueryOrderResponseBodyData) SetBuyerId(v string) *QueryOrderResponseBodyData {
	s.BuyerId = &v
	return s
}

func (s *QueryOrderResponseBodyData) SetExternalOrderNo(v string) *QueryOrderResponseBodyData {
	s.ExternalOrderNo = &v
	return s
}

func (s *QueryOrderResponseBodyData) SetGmtCreate(v int64) *QueryOrderResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *QueryOrderResponseBodyData) SetItemInfo(v *QueryOrderResponseBodyDataItemInfo) *QueryOrderResponseBodyData {
	s.ItemInfo = v
	return s
}

func (s *QueryOrderResponseBodyData) SetOrderNo(v string) *QueryOrderResponseBodyData {
	s.OrderNo = &v
	return s
}

func (s *QueryOrderResponseBodyData) SetPayment(v *QueryOrderResponseBodyDataPayment) *QueryOrderResponseBodyData {
	s.Payment = v
	return s
}

func (s *QueryOrderResponseBodyData) SetRefundOrders(v []*QueryOrderResponseBodyDataRefundOrders) *QueryOrderResponseBodyData {
	s.RefundOrders = v
	return s
}

func (s *QueryOrderResponseBodyData) SetRoomStays(v []*QueryOrderResponseBodyDataRoomStays) *QueryOrderResponseBodyData {
	s.RoomStays = v
	return s
}

func (s *QueryOrderResponseBodyData) SetSalesChannel(v string) *QueryOrderResponseBodyData {
	s.SalesChannel = &v
	return s
}

func (s *QueryOrderResponseBodyData) SetStatus(v string) *QueryOrderResponseBodyData {
	s.Status = &v
	return s
}

func (s *QueryOrderResponseBodyData) SetTracerId(v string) *QueryOrderResponseBodyData {
	s.TracerId = &v
	return s
}

func (s *QueryOrderResponseBodyData) Validate() error {
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

type QueryOrderResponseBodyDataItemInfo struct {
	CancelPolicy *QueryOrderResponseBodyDataItemInfoCancelPolicy `json:"CancelPolicy,omitempty" xml:"CancelPolicy,omitempty" type:"Struct"`
	// example:
	//
	// 2026-07-01
	CheckIn *string `json:"CheckIn,omitempty" xml:"CheckIn,omitempty"`
	// example:
	//
	// 2
	CheckInNumber *int32 `json:"CheckInNumber,omitempty" xml:"CheckInNumber,omitempty"`
	// example:
	//
	// 2026-07-03
	CheckOut    *string                                          `json:"CheckOut,omitempty" xml:"CheckOut,omitempty"`
	DailyPrices []*QueryOrderResponseBodyDataItemInfoDailyPrices `json:"DailyPrices,omitempty" xml:"DailyPrices,omitempty" type:"Repeated"`
	Meal        *QueryOrderResponseBodyDataItemInfoMeal          `json:"Meal,omitempty" xml:"Meal,omitempty" type:"Struct"`
	// example:
	//
	// 1
	RoomCount         *int32                                               `json:"RoomCount,omitempty" xml:"RoomCount,omitempty"`
	SellingTotalPrice *QueryOrderResponseBodyDataItemInfoSellingTotalPrice `json:"SellingTotalPrice,omitempty" xml:"SellingTotalPrice,omitempty" type:"Struct"`
}

func (s QueryOrderResponseBodyDataItemInfo) String() string {
	return dara.Prettify(s)
}

func (s QueryOrderResponseBodyDataItemInfo) GoString() string {
	return s.String()
}

func (s *QueryOrderResponseBodyDataItemInfo) GetCancelPolicy() *QueryOrderResponseBodyDataItemInfoCancelPolicy {
	return s.CancelPolicy
}

func (s *QueryOrderResponseBodyDataItemInfo) GetCheckIn() *string {
	return s.CheckIn
}

func (s *QueryOrderResponseBodyDataItemInfo) GetCheckInNumber() *int32 {
	return s.CheckInNumber
}

func (s *QueryOrderResponseBodyDataItemInfo) GetCheckOut() *string {
	return s.CheckOut
}

func (s *QueryOrderResponseBodyDataItemInfo) GetDailyPrices() []*QueryOrderResponseBodyDataItemInfoDailyPrices {
	return s.DailyPrices
}

func (s *QueryOrderResponseBodyDataItemInfo) GetMeal() *QueryOrderResponseBodyDataItemInfoMeal {
	return s.Meal
}

func (s *QueryOrderResponseBodyDataItemInfo) GetRoomCount() *int32 {
	return s.RoomCount
}

func (s *QueryOrderResponseBodyDataItemInfo) GetSellingTotalPrice() *QueryOrderResponseBodyDataItemInfoSellingTotalPrice {
	return s.SellingTotalPrice
}

func (s *QueryOrderResponseBodyDataItemInfo) SetCancelPolicy(v *QueryOrderResponseBodyDataItemInfoCancelPolicy) *QueryOrderResponseBodyDataItemInfo {
	s.CancelPolicy = v
	return s
}

func (s *QueryOrderResponseBodyDataItemInfo) SetCheckIn(v string) *QueryOrderResponseBodyDataItemInfo {
	s.CheckIn = &v
	return s
}

func (s *QueryOrderResponseBodyDataItemInfo) SetCheckInNumber(v int32) *QueryOrderResponseBodyDataItemInfo {
	s.CheckInNumber = &v
	return s
}

func (s *QueryOrderResponseBodyDataItemInfo) SetCheckOut(v string) *QueryOrderResponseBodyDataItemInfo {
	s.CheckOut = &v
	return s
}

func (s *QueryOrderResponseBodyDataItemInfo) SetDailyPrices(v []*QueryOrderResponseBodyDataItemInfoDailyPrices) *QueryOrderResponseBodyDataItemInfo {
	s.DailyPrices = v
	return s
}

func (s *QueryOrderResponseBodyDataItemInfo) SetMeal(v *QueryOrderResponseBodyDataItemInfoMeal) *QueryOrderResponseBodyDataItemInfo {
	s.Meal = v
	return s
}

func (s *QueryOrderResponseBodyDataItemInfo) SetRoomCount(v int32) *QueryOrderResponseBodyDataItemInfo {
	s.RoomCount = &v
	return s
}

func (s *QueryOrderResponseBodyDataItemInfo) SetSellingTotalPrice(v *QueryOrderResponseBodyDataItemInfoSellingTotalPrice) *QueryOrderResponseBodyDataItemInfo {
	s.SellingTotalPrice = v
	return s
}

func (s *QueryOrderResponseBodyDataItemInfo) Validate() error {
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

type QueryOrderResponseBodyDataItemInfoCancelPolicy struct {
	Penalties []*QueryOrderResponseBodyDataItemInfoCancelPolicyPenalties `json:"Penalties,omitempty" xml:"Penalties,omitempty" type:"Repeated"`
	// example:
	//
	// FREE_CANCELLATION
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	// example:
	//
	// TracerId
	TracerId *string `json:"TracerId,omitempty" xml:"TracerId,omitempty"`
}

func (s QueryOrderResponseBodyDataItemInfoCancelPolicy) String() string {
	return dara.Prettify(s)
}

func (s QueryOrderResponseBodyDataItemInfoCancelPolicy) GoString() string {
	return s.String()
}

func (s *QueryOrderResponseBodyDataItemInfoCancelPolicy) GetPenalties() []*QueryOrderResponseBodyDataItemInfoCancelPolicyPenalties {
	return s.Penalties
}

func (s *QueryOrderResponseBodyDataItemInfoCancelPolicy) GetPolicyType() *string {
	return s.PolicyType
}

func (s *QueryOrderResponseBodyDataItemInfoCancelPolicy) GetTracerId() *string {
	return s.TracerId
}

func (s *QueryOrderResponseBodyDataItemInfoCancelPolicy) SetPenalties(v []*QueryOrderResponseBodyDataItemInfoCancelPolicyPenalties) *QueryOrderResponseBodyDataItemInfoCancelPolicy {
	s.Penalties = v
	return s
}

func (s *QueryOrderResponseBodyDataItemInfoCancelPolicy) SetPolicyType(v string) *QueryOrderResponseBodyDataItemInfoCancelPolicy {
	s.PolicyType = &v
	return s
}

func (s *QueryOrderResponseBodyDataItemInfoCancelPolicy) SetTracerId(v string) *QueryOrderResponseBodyDataItemInfoCancelPolicy {
	s.TracerId = &v
	return s
}

func (s *QueryOrderResponseBodyDataItemInfoCancelPolicy) Validate() error {
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

type QueryOrderResponseBodyDataItemInfoCancelPolicyPenalties struct {
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
	// TracerId
	TracerId *string `json:"TracerId,omitempty" xml:"TracerId,omitempty"`
}

func (s QueryOrderResponseBodyDataItemInfoCancelPolicyPenalties) String() string {
	return dara.Prettify(s)
}

func (s QueryOrderResponseBodyDataItemInfoCancelPolicyPenalties) GoString() string {
	return s.String()
}

func (s *QueryOrderResponseBodyDataItemInfoCancelPolicyPenalties) GetCurrency() *string {
	return s.Currency
}

func (s *QueryOrderResponseBodyDataItemInfoCancelPolicyPenalties) GetEnd() *int64 {
	return s.End
}

func (s *QueryOrderResponseBodyDataItemInfoCancelPolicyPenalties) GetPenaltyType() *string {
	return s.PenaltyType
}

func (s *QueryOrderResponseBodyDataItemInfoCancelPolicyPenalties) GetPenaltyValue() *string {
	return s.PenaltyValue
}

func (s *QueryOrderResponseBodyDataItemInfoCancelPolicyPenalties) GetStart() *int64 {
	return s.Start
}

func (s *QueryOrderResponseBodyDataItemInfoCancelPolicyPenalties) GetTracerId() *string {
	return s.TracerId
}

func (s *QueryOrderResponseBodyDataItemInfoCancelPolicyPenalties) SetCurrency(v string) *QueryOrderResponseBodyDataItemInfoCancelPolicyPenalties {
	s.Currency = &v
	return s
}

func (s *QueryOrderResponseBodyDataItemInfoCancelPolicyPenalties) SetEnd(v int64) *QueryOrderResponseBodyDataItemInfoCancelPolicyPenalties {
	s.End = &v
	return s
}

func (s *QueryOrderResponseBodyDataItemInfoCancelPolicyPenalties) SetPenaltyType(v string) *QueryOrderResponseBodyDataItemInfoCancelPolicyPenalties {
	s.PenaltyType = &v
	return s
}

func (s *QueryOrderResponseBodyDataItemInfoCancelPolicyPenalties) SetPenaltyValue(v string) *QueryOrderResponseBodyDataItemInfoCancelPolicyPenalties {
	s.PenaltyValue = &v
	return s
}

func (s *QueryOrderResponseBodyDataItemInfoCancelPolicyPenalties) SetStart(v int64) *QueryOrderResponseBodyDataItemInfoCancelPolicyPenalties {
	s.Start = &v
	return s
}

func (s *QueryOrderResponseBodyDataItemInfoCancelPolicyPenalties) SetTracerId(v string) *QueryOrderResponseBodyDataItemInfoCancelPolicyPenalties {
	s.TracerId = &v
	return s
}

func (s *QueryOrderResponseBodyDataItemInfoCancelPolicyPenalties) Validate() error {
	return dara.Validate(s)
}

type QueryOrderResponseBodyDataItemInfoDailyPrices struct {
	// example:
	//
	// LocalDate
	Date  *string                                             `json:"Date,omitempty" xml:"Date,omitempty"`
	Price *QueryOrderResponseBodyDataItemInfoDailyPricesPrice `json:"Price,omitempty" xml:"Price,omitempty" type:"Struct"`
}

func (s QueryOrderResponseBodyDataItemInfoDailyPrices) String() string {
	return dara.Prettify(s)
}

func (s QueryOrderResponseBodyDataItemInfoDailyPrices) GoString() string {
	return s.String()
}

func (s *QueryOrderResponseBodyDataItemInfoDailyPrices) GetDate() *string {
	return s.Date
}

func (s *QueryOrderResponseBodyDataItemInfoDailyPrices) GetPrice() *QueryOrderResponseBodyDataItemInfoDailyPricesPrice {
	return s.Price
}

func (s *QueryOrderResponseBodyDataItemInfoDailyPrices) SetDate(v string) *QueryOrderResponseBodyDataItemInfoDailyPrices {
	s.Date = &v
	return s
}

func (s *QueryOrderResponseBodyDataItemInfoDailyPrices) SetPrice(v *QueryOrderResponseBodyDataItemInfoDailyPricesPrice) *QueryOrderResponseBodyDataItemInfoDailyPrices {
	s.Price = v
	return s
}

func (s *QueryOrderResponseBodyDataItemInfoDailyPrices) Validate() error {
	if s.Price != nil {
		if err := s.Price.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryOrderResponseBodyDataItemInfoDailyPricesPrice struct {
	// example:
	//
	// 1
	Cent     *int64                                                      `json:"Cent,omitempty" xml:"Cent,omitempty"`
	Currency *QueryOrderResponseBodyDataItemInfoDailyPricesPriceCurrency `json:"Currency,omitempty" xml:"Currency,omitempty" type:"Struct"`
}

func (s QueryOrderResponseBodyDataItemInfoDailyPricesPrice) String() string {
	return dara.Prettify(s)
}

func (s QueryOrderResponseBodyDataItemInfoDailyPricesPrice) GoString() string {
	return s.String()
}

func (s *QueryOrderResponseBodyDataItemInfoDailyPricesPrice) GetCent() *int64 {
	return s.Cent
}

func (s *QueryOrderResponseBodyDataItemInfoDailyPricesPrice) GetCurrency() *QueryOrderResponseBodyDataItemInfoDailyPricesPriceCurrency {
	return s.Currency
}

func (s *QueryOrderResponseBodyDataItemInfoDailyPricesPrice) SetCent(v int64) *QueryOrderResponseBodyDataItemInfoDailyPricesPrice {
	s.Cent = &v
	return s
}

func (s *QueryOrderResponseBodyDataItemInfoDailyPricesPrice) SetCurrency(v *QueryOrderResponseBodyDataItemInfoDailyPricesPriceCurrency) *QueryOrderResponseBodyDataItemInfoDailyPricesPrice {
	s.Currency = v
	return s
}

func (s *QueryOrderResponseBodyDataItemInfoDailyPricesPrice) Validate() error {
	if s.Currency != nil {
		if err := s.Currency.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryOrderResponseBodyDataItemInfoDailyPricesPriceCurrency struct {
	// example:
	//
	// USD
	CurrencyCode *string `json:"CurrencyCode,omitempty" xml:"CurrencyCode,omitempty"`
	// example:
	//
	// 1
	DefaultFractionDigits *int32 `json:"DefaultFractionDigits,omitempty" xml:"DefaultFractionDigits,omitempty"`
	// example:
	//
	// 1
	NumericCode *int32 `json:"NumericCode,omitempty" xml:"NumericCode,omitempty"`
}

func (s QueryOrderResponseBodyDataItemInfoDailyPricesPriceCurrency) String() string {
	return dara.Prettify(s)
}

func (s QueryOrderResponseBodyDataItemInfoDailyPricesPriceCurrency) GoString() string {
	return s.String()
}

func (s *QueryOrderResponseBodyDataItemInfoDailyPricesPriceCurrency) GetCurrencyCode() *string {
	return s.CurrencyCode
}

func (s *QueryOrderResponseBodyDataItemInfoDailyPricesPriceCurrency) GetDefaultFractionDigits() *int32 {
	return s.DefaultFractionDigits
}

func (s *QueryOrderResponseBodyDataItemInfoDailyPricesPriceCurrency) GetNumericCode() *int32 {
	return s.NumericCode
}

func (s *QueryOrderResponseBodyDataItemInfoDailyPricesPriceCurrency) SetCurrencyCode(v string) *QueryOrderResponseBodyDataItemInfoDailyPricesPriceCurrency {
	s.CurrencyCode = &v
	return s
}

func (s *QueryOrderResponseBodyDataItemInfoDailyPricesPriceCurrency) SetDefaultFractionDigits(v int32) *QueryOrderResponseBodyDataItemInfoDailyPricesPriceCurrency {
	s.DefaultFractionDigits = &v
	return s
}

func (s *QueryOrderResponseBodyDataItemInfoDailyPricesPriceCurrency) SetNumericCode(v int32) *QueryOrderResponseBodyDataItemInfoDailyPricesPriceCurrency {
	s.NumericCode = &v
	return s
}

func (s *QueryOrderResponseBodyDataItemInfoDailyPricesPriceCurrency) Validate() error {
	return dara.Validate(s)
}

type QueryOrderResponseBodyDataItemInfoMeal struct {
	// example:
	//
	// 含早餐
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// BREAKFAST
	MealType *string `json:"MealType,omitempty" xml:"MealType,omitempty"`
	// example:
	//
	// TracerId
	TracerId *string `json:"TracerId,omitempty" xml:"TracerId,omitempty"`
}

func (s QueryOrderResponseBodyDataItemInfoMeal) String() string {
	return dara.Prettify(s)
}

func (s QueryOrderResponseBodyDataItemInfoMeal) GoString() string {
	return s.String()
}

func (s *QueryOrderResponseBodyDataItemInfoMeal) GetDescription() *string {
	return s.Description
}

func (s *QueryOrderResponseBodyDataItemInfoMeal) GetMealType() *string {
	return s.MealType
}

func (s *QueryOrderResponseBodyDataItemInfoMeal) GetTracerId() *string {
	return s.TracerId
}

func (s *QueryOrderResponseBodyDataItemInfoMeal) SetDescription(v string) *QueryOrderResponseBodyDataItemInfoMeal {
	s.Description = &v
	return s
}

func (s *QueryOrderResponseBodyDataItemInfoMeal) SetMealType(v string) *QueryOrderResponseBodyDataItemInfoMeal {
	s.MealType = &v
	return s
}

func (s *QueryOrderResponseBodyDataItemInfoMeal) SetTracerId(v string) *QueryOrderResponseBodyDataItemInfoMeal {
	s.TracerId = &v
	return s
}

func (s *QueryOrderResponseBodyDataItemInfoMeal) Validate() error {
	return dara.Validate(s)
}

type QueryOrderResponseBodyDataItemInfoSellingTotalPrice struct {
	// example:
	//
	// 10000
	Amount *string `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// example:
	//
	// USD
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// example:
	//
	// TraceId
	TracerId *string `json:"TracerId,omitempty" xml:"TracerId,omitempty"`
}

func (s QueryOrderResponseBodyDataItemInfoSellingTotalPrice) String() string {
	return dara.Prettify(s)
}

func (s QueryOrderResponseBodyDataItemInfoSellingTotalPrice) GoString() string {
	return s.String()
}

func (s *QueryOrderResponseBodyDataItemInfoSellingTotalPrice) GetAmount() *string {
	return s.Amount
}

func (s *QueryOrderResponseBodyDataItemInfoSellingTotalPrice) GetCurrency() *string {
	return s.Currency
}

func (s *QueryOrderResponseBodyDataItemInfoSellingTotalPrice) GetTracerId() *string {
	return s.TracerId
}

func (s *QueryOrderResponseBodyDataItemInfoSellingTotalPrice) SetAmount(v string) *QueryOrderResponseBodyDataItemInfoSellingTotalPrice {
	s.Amount = &v
	return s
}

func (s *QueryOrderResponseBodyDataItemInfoSellingTotalPrice) SetCurrency(v string) *QueryOrderResponseBodyDataItemInfoSellingTotalPrice {
	s.Currency = &v
	return s
}

func (s *QueryOrderResponseBodyDataItemInfoSellingTotalPrice) SetTracerId(v string) *QueryOrderResponseBodyDataItemInfoSellingTotalPrice {
	s.TracerId = &v
	return s
}

func (s *QueryOrderResponseBodyDataItemInfoSellingTotalPrice) Validate() error {
	return dara.Validate(s)
}

type QueryOrderResponseBodyDataPayment struct {
	Amount *QueryOrderResponseBodyDataPaymentAmount `json:"Amount,omitempty" xml:"Amount,omitempty" type:"Struct"`
	// example:
	//
	// 1672531200000
	GmtPaid *int64 `json:"GmtPaid,omitempty" xml:"GmtPaid,omitempty"`
	// example:
	//
	// BALANCE
	PaymentMethod *string `json:"PaymentMethod,omitempty" xml:"PaymentMethod,omitempty"`
	// example:
	//
	// PAY202606290001
	PaymentTransactionId *string `json:"PaymentTransactionId,omitempty" xml:"PaymentTransactionId,omitempty"`
}

func (s QueryOrderResponseBodyDataPayment) String() string {
	return dara.Prettify(s)
}

func (s QueryOrderResponseBodyDataPayment) GoString() string {
	return s.String()
}

func (s *QueryOrderResponseBodyDataPayment) GetAmount() *QueryOrderResponseBodyDataPaymentAmount {
	return s.Amount
}

func (s *QueryOrderResponseBodyDataPayment) GetGmtPaid() *int64 {
	return s.GmtPaid
}

func (s *QueryOrderResponseBodyDataPayment) GetPaymentMethod() *string {
	return s.PaymentMethod
}

func (s *QueryOrderResponseBodyDataPayment) GetPaymentTransactionId() *string {
	return s.PaymentTransactionId
}

func (s *QueryOrderResponseBodyDataPayment) SetAmount(v *QueryOrderResponseBodyDataPaymentAmount) *QueryOrderResponseBodyDataPayment {
	s.Amount = v
	return s
}

func (s *QueryOrderResponseBodyDataPayment) SetGmtPaid(v int64) *QueryOrderResponseBodyDataPayment {
	s.GmtPaid = &v
	return s
}

func (s *QueryOrderResponseBodyDataPayment) SetPaymentMethod(v string) *QueryOrderResponseBodyDataPayment {
	s.PaymentMethod = &v
	return s
}

func (s *QueryOrderResponseBodyDataPayment) SetPaymentTransactionId(v string) *QueryOrderResponseBodyDataPayment {
	s.PaymentTransactionId = &v
	return s
}

func (s *QueryOrderResponseBodyDataPayment) Validate() error {
	if s.Amount != nil {
		if err := s.Amount.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryOrderResponseBodyDataPaymentAmount struct {
	// example:
	//
	// 10000
	Amount *string `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// example:
	//
	// USD
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// example:
	//
	// TracerId
	TracerId *string `json:"TracerId,omitempty" xml:"TracerId,omitempty"`
}

func (s QueryOrderResponseBodyDataPaymentAmount) String() string {
	return dara.Prettify(s)
}

func (s QueryOrderResponseBodyDataPaymentAmount) GoString() string {
	return s.String()
}

func (s *QueryOrderResponseBodyDataPaymentAmount) GetAmount() *string {
	return s.Amount
}

func (s *QueryOrderResponseBodyDataPaymentAmount) GetCurrency() *string {
	return s.Currency
}

func (s *QueryOrderResponseBodyDataPaymentAmount) GetTracerId() *string {
	return s.TracerId
}

func (s *QueryOrderResponseBodyDataPaymentAmount) SetAmount(v string) *QueryOrderResponseBodyDataPaymentAmount {
	s.Amount = &v
	return s
}

func (s *QueryOrderResponseBodyDataPaymentAmount) SetCurrency(v string) *QueryOrderResponseBodyDataPaymentAmount {
	s.Currency = &v
	return s
}

func (s *QueryOrderResponseBodyDataPaymentAmount) SetTracerId(v string) *QueryOrderResponseBodyDataPaymentAmount {
	s.TracerId = &v
	return s
}

func (s *QueryOrderResponseBodyDataPaymentAmount) Validate() error {
	return dara.Validate(s)
}

type QueryOrderResponseBodyDataRefundOrders struct {
	// example:
	//
	// 1672531200000
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// TXN001
	RefundTransactionId *string `json:"RefundTransactionId,omitempty" xml:"RefundTransactionId,omitempty"`
	// example:
	//
	// 不符合条件
	RejectReason *string `json:"RejectReason,omitempty" xml:"RejectReason,omitempty"`
	// example:
	//
	// RF202606290001
	SellRefundOrderNo *string `json:"SellRefundOrderNo,omitempty" xml:"SellRefundOrderNo,omitempty"`
	// example:
	//
	// REFUNDED
	Status             *string                                                   `json:"Status,omitempty" xml:"Status,omitempty"`
	TotalPenaltyAmount *QueryOrderResponseBodyDataRefundOrdersTotalPenaltyAmount `json:"TotalPenaltyAmount,omitempty" xml:"TotalPenaltyAmount,omitempty" type:"Struct"`
	TotalRefundAmount  *QueryOrderResponseBodyDataRefundOrdersTotalRefundAmount  `json:"TotalRefundAmount,omitempty" xml:"TotalRefundAmount,omitempty" type:"Struct"`
}

func (s QueryOrderResponseBodyDataRefundOrders) String() string {
	return dara.Prettify(s)
}

func (s QueryOrderResponseBodyDataRefundOrders) GoString() string {
	return s.String()
}

func (s *QueryOrderResponseBodyDataRefundOrders) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *QueryOrderResponseBodyDataRefundOrders) GetRefundTransactionId() *string {
	return s.RefundTransactionId
}

func (s *QueryOrderResponseBodyDataRefundOrders) GetRejectReason() *string {
	return s.RejectReason
}

func (s *QueryOrderResponseBodyDataRefundOrders) GetSellRefundOrderNo() *string {
	return s.SellRefundOrderNo
}

func (s *QueryOrderResponseBodyDataRefundOrders) GetStatus() *string {
	return s.Status
}

func (s *QueryOrderResponseBodyDataRefundOrders) GetTotalPenaltyAmount() *QueryOrderResponseBodyDataRefundOrdersTotalPenaltyAmount {
	return s.TotalPenaltyAmount
}

func (s *QueryOrderResponseBodyDataRefundOrders) GetTotalRefundAmount() *QueryOrderResponseBodyDataRefundOrdersTotalRefundAmount {
	return s.TotalRefundAmount
}

func (s *QueryOrderResponseBodyDataRefundOrders) SetGmtCreate(v int64) *QueryOrderResponseBodyDataRefundOrders {
	s.GmtCreate = &v
	return s
}

func (s *QueryOrderResponseBodyDataRefundOrders) SetRefundTransactionId(v string) *QueryOrderResponseBodyDataRefundOrders {
	s.RefundTransactionId = &v
	return s
}

func (s *QueryOrderResponseBodyDataRefundOrders) SetRejectReason(v string) *QueryOrderResponseBodyDataRefundOrders {
	s.RejectReason = &v
	return s
}

func (s *QueryOrderResponseBodyDataRefundOrders) SetSellRefundOrderNo(v string) *QueryOrderResponseBodyDataRefundOrders {
	s.SellRefundOrderNo = &v
	return s
}

func (s *QueryOrderResponseBodyDataRefundOrders) SetStatus(v string) *QueryOrderResponseBodyDataRefundOrders {
	s.Status = &v
	return s
}

func (s *QueryOrderResponseBodyDataRefundOrders) SetTotalPenaltyAmount(v *QueryOrderResponseBodyDataRefundOrdersTotalPenaltyAmount) *QueryOrderResponseBodyDataRefundOrders {
	s.TotalPenaltyAmount = v
	return s
}

func (s *QueryOrderResponseBodyDataRefundOrders) SetTotalRefundAmount(v *QueryOrderResponseBodyDataRefundOrdersTotalRefundAmount) *QueryOrderResponseBodyDataRefundOrders {
	s.TotalRefundAmount = v
	return s
}

func (s *QueryOrderResponseBodyDataRefundOrders) Validate() error {
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

type QueryOrderResponseBodyDataRefundOrdersTotalPenaltyAmount struct {
	// example:
	//
	// 10000
	Amount *string `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// example:
	//
	// USD
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// example:
	//
	// TraceId
	TracerId *string `json:"TracerId,omitempty" xml:"TracerId,omitempty"`
}

func (s QueryOrderResponseBodyDataRefundOrdersTotalPenaltyAmount) String() string {
	return dara.Prettify(s)
}

func (s QueryOrderResponseBodyDataRefundOrdersTotalPenaltyAmount) GoString() string {
	return s.String()
}

func (s *QueryOrderResponseBodyDataRefundOrdersTotalPenaltyAmount) GetAmount() *string {
	return s.Amount
}

func (s *QueryOrderResponseBodyDataRefundOrdersTotalPenaltyAmount) GetCurrency() *string {
	return s.Currency
}

func (s *QueryOrderResponseBodyDataRefundOrdersTotalPenaltyAmount) GetTracerId() *string {
	return s.TracerId
}

func (s *QueryOrderResponseBodyDataRefundOrdersTotalPenaltyAmount) SetAmount(v string) *QueryOrderResponseBodyDataRefundOrdersTotalPenaltyAmount {
	s.Amount = &v
	return s
}

func (s *QueryOrderResponseBodyDataRefundOrdersTotalPenaltyAmount) SetCurrency(v string) *QueryOrderResponseBodyDataRefundOrdersTotalPenaltyAmount {
	s.Currency = &v
	return s
}

func (s *QueryOrderResponseBodyDataRefundOrdersTotalPenaltyAmount) SetTracerId(v string) *QueryOrderResponseBodyDataRefundOrdersTotalPenaltyAmount {
	s.TracerId = &v
	return s
}

func (s *QueryOrderResponseBodyDataRefundOrdersTotalPenaltyAmount) Validate() error {
	return dara.Validate(s)
}

type QueryOrderResponseBodyDataRefundOrdersTotalRefundAmount struct {
	// example:
	//
	// 10000
	Amount *string `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// example:
	//
	// USD
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// example:
	//
	// TraceId
	TracerId *string `json:"TracerId,omitempty" xml:"TracerId,omitempty"`
}

func (s QueryOrderResponseBodyDataRefundOrdersTotalRefundAmount) String() string {
	return dara.Prettify(s)
}

func (s QueryOrderResponseBodyDataRefundOrdersTotalRefundAmount) GoString() string {
	return s.String()
}

func (s *QueryOrderResponseBodyDataRefundOrdersTotalRefundAmount) GetAmount() *string {
	return s.Amount
}

func (s *QueryOrderResponseBodyDataRefundOrdersTotalRefundAmount) GetCurrency() *string {
	return s.Currency
}

func (s *QueryOrderResponseBodyDataRefundOrdersTotalRefundAmount) GetTracerId() *string {
	return s.TracerId
}

func (s *QueryOrderResponseBodyDataRefundOrdersTotalRefundAmount) SetAmount(v string) *QueryOrderResponseBodyDataRefundOrdersTotalRefundAmount {
	s.Amount = &v
	return s
}

func (s *QueryOrderResponseBodyDataRefundOrdersTotalRefundAmount) SetCurrency(v string) *QueryOrderResponseBodyDataRefundOrdersTotalRefundAmount {
	s.Currency = &v
	return s
}

func (s *QueryOrderResponseBodyDataRefundOrdersTotalRefundAmount) SetTracerId(v string) *QueryOrderResponseBodyDataRefundOrdersTotalRefundAmount {
	s.TracerId = &v
	return s
}

func (s *QueryOrderResponseBodyDataRefundOrdersTotalRefundAmount) Validate() error {
	return dara.Validate(s)
}

type QueryOrderResponseBodyDataRoomStays struct {
	// example:
	//
	// CONF001
	ConfirmationId *string                                      `json:"ConfirmationId,omitempty" xml:"ConfirmationId,omitempty"`
	Guests         []*QueryOrderResponseBodyDataRoomStaysGuests `json:"Guests,omitempty" xml:"Guests,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	RoomIndex *int32 `json:"RoomIndex,omitempty" xml:"RoomIndex,omitempty"`
	// example:
	//
	// 豪华大床房
	RoomName *string `json:"RoomName,omitempty" xml:"RoomName,omitempty"`
}

func (s QueryOrderResponseBodyDataRoomStays) String() string {
	return dara.Prettify(s)
}

func (s QueryOrderResponseBodyDataRoomStays) GoString() string {
	return s.String()
}

func (s *QueryOrderResponseBodyDataRoomStays) GetConfirmationId() *string {
	return s.ConfirmationId
}

func (s *QueryOrderResponseBodyDataRoomStays) GetGuests() []*QueryOrderResponseBodyDataRoomStaysGuests {
	return s.Guests
}

func (s *QueryOrderResponseBodyDataRoomStays) GetRoomIndex() *int32 {
	return s.RoomIndex
}

func (s *QueryOrderResponseBodyDataRoomStays) GetRoomName() *string {
	return s.RoomName
}

func (s *QueryOrderResponseBodyDataRoomStays) SetConfirmationId(v string) *QueryOrderResponseBodyDataRoomStays {
	s.ConfirmationId = &v
	return s
}

func (s *QueryOrderResponseBodyDataRoomStays) SetGuests(v []*QueryOrderResponseBodyDataRoomStaysGuests) *QueryOrderResponseBodyDataRoomStays {
	s.Guests = v
	return s
}

func (s *QueryOrderResponseBodyDataRoomStays) SetRoomIndex(v int32) *QueryOrderResponseBodyDataRoomStays {
	s.RoomIndex = &v
	return s
}

func (s *QueryOrderResponseBodyDataRoomStays) SetRoomName(v string) *QueryOrderResponseBodyDataRoomStays {
	s.RoomName = &v
	return s
}

func (s *QueryOrderResponseBodyDataRoomStays) Validate() error {
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

type QueryOrderResponseBodyDataRoomStaysGuests struct {
	// example:
	//
	// John
	FirstName *string `json:"FirstName,omitempty" xml:"FirstName,omitempty"`
	// example:
	//
	// Doe
	LastName *string `json:"LastName,omitempty" xml:"LastName,omitempty"`
	// example:
	//
	// TracerId
	TracerId *string `json:"TracerId,omitempty" xml:"TracerId,omitempty"`
}

func (s QueryOrderResponseBodyDataRoomStaysGuests) String() string {
	return dara.Prettify(s)
}

func (s QueryOrderResponseBodyDataRoomStaysGuests) GoString() string {
	return s.String()
}

func (s *QueryOrderResponseBodyDataRoomStaysGuests) GetFirstName() *string {
	return s.FirstName
}

func (s *QueryOrderResponseBodyDataRoomStaysGuests) GetLastName() *string {
	return s.LastName
}

func (s *QueryOrderResponseBodyDataRoomStaysGuests) GetTracerId() *string {
	return s.TracerId
}

func (s *QueryOrderResponseBodyDataRoomStaysGuests) SetFirstName(v string) *QueryOrderResponseBodyDataRoomStaysGuests {
	s.FirstName = &v
	return s
}

func (s *QueryOrderResponseBodyDataRoomStaysGuests) SetLastName(v string) *QueryOrderResponseBodyDataRoomStaysGuests {
	s.LastName = &v
	return s
}

func (s *QueryOrderResponseBodyDataRoomStaysGuests) SetTracerId(v string) *QueryOrderResponseBodyDataRoomStaysGuests {
	s.TracerId = &v
	return s
}

func (s *QueryOrderResponseBodyDataRoomStaysGuests) Validate() error {
	return dara.Validate(s)
}
