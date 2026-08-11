// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDataHotelsValue interface {
	dara.Model
	String() string
	GoString() string
	SetCheckInDate(v string) *DataHotelsValue
	GetCheckInDate() *string
	SetCheckOutDate(v string) *DataHotelsValue
	GetCheckOutDate() *string
	SetRooms(v []*DataHotelsValueRooms) *DataHotelsValue
	GetRooms() []*DataHotelsValueRooms
	SetStandardRoomId(v string) *DataHotelsValue
	GetStandardRoomId() *string
	SetOffers(v []*DataHotelsValueOffers) *DataHotelsValue
	GetOffers() []*DataHotelsValueOffers
}

type DataHotelsValue struct {
	// The check-in date (yyyy-MM-dd).
	//
	// example:
	//
	// 2026-01-01
	CheckInDate *string `json:"CheckInDate,omitempty" xml:"CheckInDate,omitempty"`
	// The check-out date (yyyy-MM-dd).
	//
	// example:
	//
	// 2026-01-02
	CheckOutDate *string `json:"CheckOutDate,omitempty" xml:"CheckOutDate,omitempty"`
	// The list of available room types for the day.
	Rooms []*DataHotelsValueRooms `json:"Rooms,omitempty" xml:"Rooms,omitempty" type:"Repeated"`
	// The standard room type ID.
	//
	// example:
	//
	// R001
	StandardRoomId *string `json:"StandardRoomId,omitempty" xml:"StandardRoomId,omitempty"`
	// All available offers for the room type.
	Offers []*DataHotelsValueOffers `json:"Offers,omitempty" xml:"Offers,omitempty" type:"Repeated"`
}

func (s DataHotelsValue) String() string {
	return dara.Prettify(s)
}

func (s DataHotelsValue) GoString() string {
	return s.String()
}

func (s *DataHotelsValue) GetCheckInDate() *string {
	return s.CheckInDate
}

func (s *DataHotelsValue) GetCheckOutDate() *string {
	return s.CheckOutDate
}

func (s *DataHotelsValue) GetRooms() []*DataHotelsValueRooms {
	return s.Rooms
}

func (s *DataHotelsValue) GetStandardRoomId() *string {
	return s.StandardRoomId
}

func (s *DataHotelsValue) GetOffers() []*DataHotelsValueOffers {
	return s.Offers
}

func (s *DataHotelsValue) SetCheckInDate(v string) *DataHotelsValue {
	s.CheckInDate = &v
	return s
}

func (s *DataHotelsValue) SetCheckOutDate(v string) *DataHotelsValue {
	s.CheckOutDate = &v
	return s
}

func (s *DataHotelsValue) SetRooms(v []*DataHotelsValueRooms) *DataHotelsValue {
	s.Rooms = v
	return s
}

func (s *DataHotelsValue) SetStandardRoomId(v string) *DataHotelsValue {
	s.StandardRoomId = &v
	return s
}

func (s *DataHotelsValue) SetOffers(v []*DataHotelsValueOffers) *DataHotelsValue {
	s.Offers = v
	return s
}

func (s *DataHotelsValue) Validate() error {
	if s.Rooms != nil {
		for _, item := range s.Rooms {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Offers != nil {
		for _, item := range s.Offers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DataHotelsValueRooms struct {
	// The standard room type ID.
	//
	// example:
	//
	// R001
	StandardRoomId *string `json:"StandardRoomId,omitempty" xml:"StandardRoomId,omitempty"`
	// The lowest price for the room type on the day.
	LowestPrice *DataHotelsValueRoomsLowestPrice `json:"LowestPrice,omitempty" xml:"LowestPrice,omitempty" type:"Struct"`
	// The list of all available offers for the room type.
	Offers []*DataHotelsValueRoomsOffers `json:"Offers,omitempty" xml:"Offers,omitempty" type:"Repeated"`
}

func (s DataHotelsValueRooms) String() string {
	return dara.Prettify(s)
}

func (s DataHotelsValueRooms) GoString() string {
	return s.String()
}

func (s *DataHotelsValueRooms) GetStandardRoomId() *string {
	return s.StandardRoomId
}

func (s *DataHotelsValueRooms) GetLowestPrice() *DataHotelsValueRoomsLowestPrice {
	return s.LowestPrice
}

func (s *DataHotelsValueRooms) GetOffers() []*DataHotelsValueRoomsOffers {
	return s.Offers
}

func (s *DataHotelsValueRooms) SetStandardRoomId(v string) *DataHotelsValueRooms {
	s.StandardRoomId = &v
	return s
}

func (s *DataHotelsValueRooms) SetLowestPrice(v *DataHotelsValueRoomsLowestPrice) *DataHotelsValueRooms {
	s.LowestPrice = v
	return s
}

func (s *DataHotelsValueRooms) SetOffers(v []*DataHotelsValueRoomsOffers) *DataHotelsValueRooms {
	s.Offers = v
	return s
}

func (s *DataHotelsValueRooms) Validate() error {
	if s.LowestPrice != nil {
		if err := s.LowestPrice.Validate(); err != nil {
			return err
		}
	}
	if s.Offers != nil {
		for _, item := range s.Offers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DataHotelsValueRoomsLowestPrice struct {
	// The amount in the smallest currency unit.
	//
	// example:
	//
	// 287
	Amount *string `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// The currency code.
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

func (s DataHotelsValueRoomsLowestPrice) String() string {
	return dara.Prettify(s)
}

func (s DataHotelsValueRoomsLowestPrice) GoString() string {
	return s.String()
}

func (s *DataHotelsValueRoomsLowestPrice) GetAmount() *string {
	return s.Amount
}

func (s *DataHotelsValueRoomsLowestPrice) GetCurrency() *string {
	return s.Currency
}

func (s *DataHotelsValueRoomsLowestPrice) GetTracerId() *string {
	return s.TracerId
}

func (s *DataHotelsValueRoomsLowestPrice) SetAmount(v string) *DataHotelsValueRoomsLowestPrice {
	s.Amount = &v
	return s
}

func (s *DataHotelsValueRoomsLowestPrice) SetCurrency(v string) *DataHotelsValueRoomsLowestPrice {
	s.Currency = &v
	return s
}

func (s *DataHotelsValueRoomsLowestPrice) SetTracerId(v string) *DataHotelsValueRoomsLowestPrice {
	s.TracerId = &v
	return s
}

func (s *DataHotelsValueRoomsLowestPrice) Validate() error {
	return dara.Validate(s)
}

type DataHotelsValueRoomsOffers struct {
	// The item-level offer identifier (price verification key, pass through as-is).
	//
	// example:
	//
	// itemOffer_123
	ItemOfferKey *string `json:"ItemOfferKey,omitempty" xml:"ItemOfferKey,omitempty"`
	// The rate plan name.
	//
	// example:
	//
	// Breakfast included
	RatePlanName *string `json:"RatePlanName,omitempty" xml:"RatePlanName,omitempty"`
	// The meal type.
	//
	// example:
	//
	// BREAKFAST
	MealType *string `json:"MealType,omitempty" xml:"MealType,omitempty"`
	// The number of meals included.
	//
	// example:
	//
	// 2
	MealCount *int32 `json:"MealCount,omitempty" xml:"MealCount,omitempty"`
	// The cancellation policy.
	CancelPolicy *DataHotelsValueRoomsOffersCancelPolicy `json:"CancelPolicy,omitempty" xml:"CancelPolicy,omitempty" type:"Struct"`
	// The total selling price.
	TotalPrice *DataHotelsValueRoomsOffersTotalPrice `json:"TotalPrice,omitempty" xml:"TotalPrice,omitempty" type:"Struct"`
	// The list of daily prices.
	DailyPrices []*DataHotelsValueRoomsOffersDailyPrices `json:"DailyPrices,omitempty" xml:"DailyPrices,omitempty" type:"Repeated"`
	// The number of available rooms.
	//
	// example:
	//
	// 5
	AvailableRooms *int32 `json:"AvailableRooms,omitempty" xml:"AvailableRooms,omitempty"`
	// The maximum number of guests allowed.
	//
	// example:
	//
	// 3
	MaxOccupancy *int32 `json:"MaxOccupancy,omitempty" xml:"MaxOccupancy,omitempty"`
	// The confirmation type (INSTANT_CONFIRM/NON_INSTANT_CONFIRM).
	//
	// example:
	//
	// INSTANT_CONFIRM
	ConfirmType *string `json:"ConfirmType,omitempty" xml:"ConfirmType,omitempty"`
}

func (s DataHotelsValueRoomsOffers) String() string {
	return dara.Prettify(s)
}

func (s DataHotelsValueRoomsOffers) GoString() string {
	return s.String()
}

func (s *DataHotelsValueRoomsOffers) GetItemOfferKey() *string {
	return s.ItemOfferKey
}

func (s *DataHotelsValueRoomsOffers) GetRatePlanName() *string {
	return s.RatePlanName
}

func (s *DataHotelsValueRoomsOffers) GetMealType() *string {
	return s.MealType
}

func (s *DataHotelsValueRoomsOffers) GetMealCount() *int32 {
	return s.MealCount
}

func (s *DataHotelsValueRoomsOffers) GetCancelPolicy() *DataHotelsValueRoomsOffersCancelPolicy {
	return s.CancelPolicy
}

func (s *DataHotelsValueRoomsOffers) GetTotalPrice() *DataHotelsValueRoomsOffersTotalPrice {
	return s.TotalPrice
}

func (s *DataHotelsValueRoomsOffers) GetDailyPrices() []*DataHotelsValueRoomsOffersDailyPrices {
	return s.DailyPrices
}

func (s *DataHotelsValueRoomsOffers) GetAvailableRooms() *int32 {
	return s.AvailableRooms
}

func (s *DataHotelsValueRoomsOffers) GetMaxOccupancy() *int32 {
	return s.MaxOccupancy
}

func (s *DataHotelsValueRoomsOffers) GetConfirmType() *string {
	return s.ConfirmType
}

func (s *DataHotelsValueRoomsOffers) SetItemOfferKey(v string) *DataHotelsValueRoomsOffers {
	s.ItemOfferKey = &v
	return s
}

func (s *DataHotelsValueRoomsOffers) SetRatePlanName(v string) *DataHotelsValueRoomsOffers {
	s.RatePlanName = &v
	return s
}

func (s *DataHotelsValueRoomsOffers) SetMealType(v string) *DataHotelsValueRoomsOffers {
	s.MealType = &v
	return s
}

func (s *DataHotelsValueRoomsOffers) SetMealCount(v int32) *DataHotelsValueRoomsOffers {
	s.MealCount = &v
	return s
}

func (s *DataHotelsValueRoomsOffers) SetCancelPolicy(v *DataHotelsValueRoomsOffersCancelPolicy) *DataHotelsValueRoomsOffers {
	s.CancelPolicy = v
	return s
}

func (s *DataHotelsValueRoomsOffers) SetTotalPrice(v *DataHotelsValueRoomsOffersTotalPrice) *DataHotelsValueRoomsOffers {
	s.TotalPrice = v
	return s
}

func (s *DataHotelsValueRoomsOffers) SetDailyPrices(v []*DataHotelsValueRoomsOffersDailyPrices) *DataHotelsValueRoomsOffers {
	s.DailyPrices = v
	return s
}

func (s *DataHotelsValueRoomsOffers) SetAvailableRooms(v int32) *DataHotelsValueRoomsOffers {
	s.AvailableRooms = &v
	return s
}

func (s *DataHotelsValueRoomsOffers) SetMaxOccupancy(v int32) *DataHotelsValueRoomsOffers {
	s.MaxOccupancy = &v
	return s
}

func (s *DataHotelsValueRoomsOffers) SetConfirmType(v string) *DataHotelsValueRoomsOffers {
	s.ConfirmType = &v
	return s
}

func (s *DataHotelsValueRoomsOffers) Validate() error {
	if s.CancelPolicy != nil {
		if err := s.CancelPolicy.Validate(); err != nil {
			return err
		}
	}
	if s.TotalPrice != nil {
		if err := s.TotalPrice.Validate(); err != nil {
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
	return nil
}

type DataHotelsValueRoomsOffersCancelPolicy struct {
	// The policy type (NON_REFUNDABLE/FREE_CANCELLATION/PARTIAL_REFUND).
	//
	// example:
	//
	// FREE_CANCELLATION
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	// The list of penalty details.
	Penalties []*DataHotelsValueRoomsOffersCancelPolicyPenalties `json:"Penalties,omitempty" xml:"Penalties,omitempty" type:"Repeated"`
	// TraceId
	//
	// example:
	//
	// TraceId
	TracerId *string `json:"TracerId,omitempty" xml:"TracerId,omitempty"`
}

func (s DataHotelsValueRoomsOffersCancelPolicy) String() string {
	return dara.Prettify(s)
}

func (s DataHotelsValueRoomsOffersCancelPolicy) GoString() string {
	return s.String()
}

func (s *DataHotelsValueRoomsOffersCancelPolicy) GetPolicyType() *string {
	return s.PolicyType
}

func (s *DataHotelsValueRoomsOffersCancelPolicy) GetPenalties() []*DataHotelsValueRoomsOffersCancelPolicyPenalties {
	return s.Penalties
}

func (s *DataHotelsValueRoomsOffersCancelPolicy) GetTracerId() *string {
	return s.TracerId
}

func (s *DataHotelsValueRoomsOffersCancelPolicy) SetPolicyType(v string) *DataHotelsValueRoomsOffersCancelPolicy {
	s.PolicyType = &v
	return s
}

func (s *DataHotelsValueRoomsOffersCancelPolicy) SetPenalties(v []*DataHotelsValueRoomsOffersCancelPolicyPenalties) *DataHotelsValueRoomsOffersCancelPolicy {
	s.Penalties = v
	return s
}

func (s *DataHotelsValueRoomsOffersCancelPolicy) SetTracerId(v string) *DataHotelsValueRoomsOffersCancelPolicy {
	s.TracerId = &v
	return s
}

func (s *DataHotelsValueRoomsOffersCancelPolicy) Validate() error {
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

type DataHotelsValueRoomsOffersCancelPolicyPenalties struct {
	// The effective start time (UTC millisecond timestamp).
	//
	// example:
	//
	// 1672531200000
	Start *string `json:"Start,omitempty" xml:"Start,omitempty"`
	// The effective end time (UTC millisecond timestamp).
	//
	// example:
	//
	// 1672617600000
	End *string `json:"End,omitempty" xml:"End,omitempty"`
	// The penalty type (PERCENTAGE/AMOUNT/NIGHTS).
	//
	// example:
	//
	// PERCENTAGE
	PenaltyType *string `json:"PenaltyType,omitempty" xml:"PenaltyType,omitempty"`
	// The penalty value (percentage/amount/nights).
	//
	// example:
	//
	// 50
	PenaltyValue *string `json:"PenaltyValue,omitempty" xml:"PenaltyValue,omitempty"`
	// The currency code (present only when the penalty type is AMOUNT).
	//
	// example:
	//
	// USD
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// traceId
	//
	// example:
	//
	// TraceId
	TracerId *string `json:"TracerId,omitempty" xml:"TracerId,omitempty"`
}

func (s DataHotelsValueRoomsOffersCancelPolicyPenalties) String() string {
	return dara.Prettify(s)
}

func (s DataHotelsValueRoomsOffersCancelPolicyPenalties) GoString() string {
	return s.String()
}

func (s *DataHotelsValueRoomsOffersCancelPolicyPenalties) GetStart() *string {
	return s.Start
}

func (s *DataHotelsValueRoomsOffersCancelPolicyPenalties) GetEnd() *string {
	return s.End
}

func (s *DataHotelsValueRoomsOffersCancelPolicyPenalties) GetPenaltyType() *string {
	return s.PenaltyType
}

func (s *DataHotelsValueRoomsOffersCancelPolicyPenalties) GetPenaltyValue() *string {
	return s.PenaltyValue
}

func (s *DataHotelsValueRoomsOffersCancelPolicyPenalties) GetCurrency() *string {
	return s.Currency
}

func (s *DataHotelsValueRoomsOffersCancelPolicyPenalties) GetTracerId() *string {
	return s.TracerId
}

func (s *DataHotelsValueRoomsOffersCancelPolicyPenalties) SetStart(v string) *DataHotelsValueRoomsOffersCancelPolicyPenalties {
	s.Start = &v
	return s
}

func (s *DataHotelsValueRoomsOffersCancelPolicyPenalties) SetEnd(v string) *DataHotelsValueRoomsOffersCancelPolicyPenalties {
	s.End = &v
	return s
}

func (s *DataHotelsValueRoomsOffersCancelPolicyPenalties) SetPenaltyType(v string) *DataHotelsValueRoomsOffersCancelPolicyPenalties {
	s.PenaltyType = &v
	return s
}

func (s *DataHotelsValueRoomsOffersCancelPolicyPenalties) SetPenaltyValue(v string) *DataHotelsValueRoomsOffersCancelPolicyPenalties {
	s.PenaltyValue = &v
	return s
}

func (s *DataHotelsValueRoomsOffersCancelPolicyPenalties) SetCurrency(v string) *DataHotelsValueRoomsOffersCancelPolicyPenalties {
	s.Currency = &v
	return s
}

func (s *DataHotelsValueRoomsOffersCancelPolicyPenalties) SetTracerId(v string) *DataHotelsValueRoomsOffersCancelPolicyPenalties {
	s.TracerId = &v
	return s
}

func (s *DataHotelsValueRoomsOffersCancelPolicyPenalties) Validate() error {
	return dara.Validate(s)
}

type DataHotelsValueRoomsOffersTotalPrice struct {
	// The amount in the smallest currency unit.
	//
	// example:
	//
	// 287
	Amount *string `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// The currency code.
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

func (s DataHotelsValueRoomsOffersTotalPrice) String() string {
	return dara.Prettify(s)
}

func (s DataHotelsValueRoomsOffersTotalPrice) GoString() string {
	return s.String()
}

func (s *DataHotelsValueRoomsOffersTotalPrice) GetAmount() *string {
	return s.Amount
}

func (s *DataHotelsValueRoomsOffersTotalPrice) GetCurrency() *string {
	return s.Currency
}

func (s *DataHotelsValueRoomsOffersTotalPrice) GetTracerId() *string {
	return s.TracerId
}

func (s *DataHotelsValueRoomsOffersTotalPrice) SetAmount(v string) *DataHotelsValueRoomsOffersTotalPrice {
	s.Amount = &v
	return s
}

func (s *DataHotelsValueRoomsOffersTotalPrice) SetCurrency(v string) *DataHotelsValueRoomsOffersTotalPrice {
	s.Currency = &v
	return s
}

func (s *DataHotelsValueRoomsOffersTotalPrice) SetTracerId(v string) *DataHotelsValueRoomsOffersTotalPrice {
	s.TracerId = &v
	return s
}

func (s *DataHotelsValueRoomsOffersTotalPrice) Validate() error {
	return dara.Validate(s)
}

type DataHotelsValueRoomsOffersDailyPrices struct {
	// The check-in date (yyyy-MM-dd, time zone: hotel local time zone).
	//
	// example:
	//
	// 2026-08-16
	Date *string `json:"Date,omitempty" xml:"Date,omitempty"`
	// The price for the day.
	Price *DataHotelsValueRoomsOffersDailyPricesPrice `json:"Price,omitempty" xml:"Price,omitempty" type:"Struct"`
	// null
	//
	// example:
	//
	// null
	TracerId *string `json:"TracerId,omitempty" xml:"TracerId,omitempty"`
}

func (s DataHotelsValueRoomsOffersDailyPrices) String() string {
	return dara.Prettify(s)
}

func (s DataHotelsValueRoomsOffersDailyPrices) GoString() string {
	return s.String()
}

func (s *DataHotelsValueRoomsOffersDailyPrices) GetDate() *string {
	return s.Date
}

func (s *DataHotelsValueRoomsOffersDailyPrices) GetPrice() *DataHotelsValueRoomsOffersDailyPricesPrice {
	return s.Price
}

func (s *DataHotelsValueRoomsOffersDailyPrices) GetTracerId() *string {
	return s.TracerId
}

func (s *DataHotelsValueRoomsOffersDailyPrices) SetDate(v string) *DataHotelsValueRoomsOffersDailyPrices {
	s.Date = &v
	return s
}

func (s *DataHotelsValueRoomsOffersDailyPrices) SetPrice(v *DataHotelsValueRoomsOffersDailyPricesPrice) *DataHotelsValueRoomsOffersDailyPrices {
	s.Price = v
	return s
}

func (s *DataHotelsValueRoomsOffersDailyPrices) SetTracerId(v string) *DataHotelsValueRoomsOffersDailyPrices {
	s.TracerId = &v
	return s
}

func (s *DataHotelsValueRoomsOffersDailyPrices) Validate() error {
	if s.Price != nil {
		if err := s.Price.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DataHotelsValueRoomsOffersDailyPricesPrice struct {
	// The amount in the smallest currency unit.
	//
	// example:
	//
	// 287
	Amount *string `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// The currency code.
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

func (s DataHotelsValueRoomsOffersDailyPricesPrice) String() string {
	return dara.Prettify(s)
}

func (s DataHotelsValueRoomsOffersDailyPricesPrice) GoString() string {
	return s.String()
}

func (s *DataHotelsValueRoomsOffersDailyPricesPrice) GetAmount() *string {
	return s.Amount
}

func (s *DataHotelsValueRoomsOffersDailyPricesPrice) GetCurrency() *string {
	return s.Currency
}

func (s *DataHotelsValueRoomsOffersDailyPricesPrice) GetTracerId() *string {
	return s.TracerId
}

func (s *DataHotelsValueRoomsOffersDailyPricesPrice) SetAmount(v string) *DataHotelsValueRoomsOffersDailyPricesPrice {
	s.Amount = &v
	return s
}

func (s *DataHotelsValueRoomsOffersDailyPricesPrice) SetCurrency(v string) *DataHotelsValueRoomsOffersDailyPricesPrice {
	s.Currency = &v
	return s
}

func (s *DataHotelsValueRoomsOffersDailyPricesPrice) SetTracerId(v string) *DataHotelsValueRoomsOffersDailyPricesPrice {
	s.TracerId = &v
	return s
}

func (s *DataHotelsValueRoomsOffersDailyPricesPrice) Validate() error {
	return dara.Validate(s)
}

type DataHotelsValueOffers struct {
	// The item offer key used for price verification. Pass through this value as-is.
	//
	// example:
	//
	// itemOffer_123
	ItemOfferKey *string `json:"ItemOfferKey,omitempty" xml:"ItemOfferKey,omitempty"`
	// The rate plan name.
	//
	// example:
	//
	// Breakfast included
	RatePlanName *string `json:"RatePlanName,omitempty" xml:"RatePlanName,omitempty"`
	// The meal type.
	//
	// example:
	//
	// BREAKFAST
	MealType *string `json:"MealType,omitempty" xml:"MealType,omitempty"`
	// The number of meals included.
	//
	// example:
	//
	// 2
	MealCount *int32 `json:"MealCount,omitempty" xml:"MealCount,omitempty"`
	// The cancellation policy.
	CancelPolicy *DataHotelsValueOffersCancelPolicy `json:"CancelPolicy,omitempty" xml:"CancelPolicy,omitempty" type:"Struct"`
	// The total selling price.
	TotalPrice *DataHotelsValueOffersTotalPrice `json:"TotalPrice,omitempty" xml:"TotalPrice,omitempty" type:"Struct"`
	// The list of daily prices.
	DailyPrices []*DataHotelsValueOffersDailyPrices `json:"DailyPrices,omitempty" xml:"DailyPrices,omitempty" type:"Repeated"`
	// The number of available rooms.
	//
	// example:
	//
	// 5
	AvailableRooms *int32 `json:"AvailableRooms,omitempty" xml:"AvailableRooms,omitempty"`
	// The maximum number of guests allowed.
	//
	// example:
	//
	// 3
	MaxOccupancy *int32 `json:"MaxOccupancy,omitempty" xml:"MaxOccupancy,omitempty"`
	// The confirmation type (INSTANT_CONFIRM/NON_INSTANT_CONFIRM).
	//
	// example:
	//
	// INSTANT_CONFIRM
	ConfirmType *string `json:"ConfirmType,omitempty" xml:"ConfirmType,omitempty"`
}

func (s DataHotelsValueOffers) String() string {
	return dara.Prettify(s)
}

func (s DataHotelsValueOffers) GoString() string {
	return s.String()
}

func (s *DataHotelsValueOffers) GetItemOfferKey() *string {
	return s.ItemOfferKey
}

func (s *DataHotelsValueOffers) GetRatePlanName() *string {
	return s.RatePlanName
}

func (s *DataHotelsValueOffers) GetMealType() *string {
	return s.MealType
}

func (s *DataHotelsValueOffers) GetMealCount() *int32 {
	return s.MealCount
}

func (s *DataHotelsValueOffers) GetCancelPolicy() *DataHotelsValueOffersCancelPolicy {
	return s.CancelPolicy
}

func (s *DataHotelsValueOffers) GetTotalPrice() *DataHotelsValueOffersTotalPrice {
	return s.TotalPrice
}

func (s *DataHotelsValueOffers) GetDailyPrices() []*DataHotelsValueOffersDailyPrices {
	return s.DailyPrices
}

func (s *DataHotelsValueOffers) GetAvailableRooms() *int32 {
	return s.AvailableRooms
}

func (s *DataHotelsValueOffers) GetMaxOccupancy() *int32 {
	return s.MaxOccupancy
}

func (s *DataHotelsValueOffers) GetConfirmType() *string {
	return s.ConfirmType
}

func (s *DataHotelsValueOffers) SetItemOfferKey(v string) *DataHotelsValueOffers {
	s.ItemOfferKey = &v
	return s
}

func (s *DataHotelsValueOffers) SetRatePlanName(v string) *DataHotelsValueOffers {
	s.RatePlanName = &v
	return s
}

func (s *DataHotelsValueOffers) SetMealType(v string) *DataHotelsValueOffers {
	s.MealType = &v
	return s
}

func (s *DataHotelsValueOffers) SetMealCount(v int32) *DataHotelsValueOffers {
	s.MealCount = &v
	return s
}

func (s *DataHotelsValueOffers) SetCancelPolicy(v *DataHotelsValueOffersCancelPolicy) *DataHotelsValueOffers {
	s.CancelPolicy = v
	return s
}

func (s *DataHotelsValueOffers) SetTotalPrice(v *DataHotelsValueOffersTotalPrice) *DataHotelsValueOffers {
	s.TotalPrice = v
	return s
}

func (s *DataHotelsValueOffers) SetDailyPrices(v []*DataHotelsValueOffersDailyPrices) *DataHotelsValueOffers {
	s.DailyPrices = v
	return s
}

func (s *DataHotelsValueOffers) SetAvailableRooms(v int32) *DataHotelsValueOffers {
	s.AvailableRooms = &v
	return s
}

func (s *DataHotelsValueOffers) SetMaxOccupancy(v int32) *DataHotelsValueOffers {
	s.MaxOccupancy = &v
	return s
}

func (s *DataHotelsValueOffers) SetConfirmType(v string) *DataHotelsValueOffers {
	s.ConfirmType = &v
	return s
}

func (s *DataHotelsValueOffers) Validate() error {
	if s.CancelPolicy != nil {
		if err := s.CancelPolicy.Validate(); err != nil {
			return err
		}
	}
	if s.TotalPrice != nil {
		if err := s.TotalPrice.Validate(); err != nil {
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
	return nil
}

type DataHotelsValueOffersCancelPolicy struct {
	// The policy type (NON_REFUNDABLE/FREE_CANCELLATION/PARTIAL_REFUND).
	//
	// example:
	//
	// FREE_CANCELLATION
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	// The list of penalty details.
	Penalties []*DataHotelsValueOffersCancelPolicyPenalties `json:"Penalties,omitempty" xml:"Penalties,omitempty" type:"Repeated"`
	// TraceId
	//
	// example:
	//
	// TraceId
	TracerId *string `json:"TracerId,omitempty" xml:"TracerId,omitempty"`
}

func (s DataHotelsValueOffersCancelPolicy) String() string {
	return dara.Prettify(s)
}

func (s DataHotelsValueOffersCancelPolicy) GoString() string {
	return s.String()
}

func (s *DataHotelsValueOffersCancelPolicy) GetPolicyType() *string {
	return s.PolicyType
}

func (s *DataHotelsValueOffersCancelPolicy) GetPenalties() []*DataHotelsValueOffersCancelPolicyPenalties {
	return s.Penalties
}

func (s *DataHotelsValueOffersCancelPolicy) GetTracerId() *string {
	return s.TracerId
}

func (s *DataHotelsValueOffersCancelPolicy) SetPolicyType(v string) *DataHotelsValueOffersCancelPolicy {
	s.PolicyType = &v
	return s
}

func (s *DataHotelsValueOffersCancelPolicy) SetPenalties(v []*DataHotelsValueOffersCancelPolicyPenalties) *DataHotelsValueOffersCancelPolicy {
	s.Penalties = v
	return s
}

func (s *DataHotelsValueOffersCancelPolicy) SetTracerId(v string) *DataHotelsValueOffersCancelPolicy {
	s.TracerId = &v
	return s
}

func (s *DataHotelsValueOffersCancelPolicy) Validate() error {
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

type DataHotelsValueOffersCancelPolicyPenalties struct {
	// The effective start time (UTC millisecond timestamp).
	//
	// example:
	//
	// 1672531200000
	Start *string `json:"Start,omitempty" xml:"Start,omitempty"`
	// The effective end time (UTC millisecond timestamp).
	//
	// example:
	//
	// 1672617600000
	End *string `json:"End,omitempty" xml:"End,omitempty"`
	// The penalty type (PERCENTAGE/AMOUNT/NIGHTS).
	//
	// example:
	//
	// PERCENTAGE
	PenaltyType *string `json:"PenaltyType,omitempty" xml:"PenaltyType,omitempty"`
	// The penalty value (percentage/amount/nights).
	//
	// example:
	//
	// 50
	PenaltyValue *string `json:"PenaltyValue,omitempty" xml:"PenaltyValue,omitempty"`
	// The currency code. This field has a value only when the penalty type is AMOUNT.
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

func (s DataHotelsValueOffersCancelPolicyPenalties) String() string {
	return dara.Prettify(s)
}

func (s DataHotelsValueOffersCancelPolicyPenalties) GoString() string {
	return s.String()
}

func (s *DataHotelsValueOffersCancelPolicyPenalties) GetStart() *string {
	return s.Start
}

func (s *DataHotelsValueOffersCancelPolicyPenalties) GetEnd() *string {
	return s.End
}

func (s *DataHotelsValueOffersCancelPolicyPenalties) GetPenaltyType() *string {
	return s.PenaltyType
}

func (s *DataHotelsValueOffersCancelPolicyPenalties) GetPenaltyValue() *string {
	return s.PenaltyValue
}

func (s *DataHotelsValueOffersCancelPolicyPenalties) GetCurrency() *string {
	return s.Currency
}

func (s *DataHotelsValueOffersCancelPolicyPenalties) GetTracerId() *string {
	return s.TracerId
}

func (s *DataHotelsValueOffersCancelPolicyPenalties) SetStart(v string) *DataHotelsValueOffersCancelPolicyPenalties {
	s.Start = &v
	return s
}

func (s *DataHotelsValueOffersCancelPolicyPenalties) SetEnd(v string) *DataHotelsValueOffersCancelPolicyPenalties {
	s.End = &v
	return s
}

func (s *DataHotelsValueOffersCancelPolicyPenalties) SetPenaltyType(v string) *DataHotelsValueOffersCancelPolicyPenalties {
	s.PenaltyType = &v
	return s
}

func (s *DataHotelsValueOffersCancelPolicyPenalties) SetPenaltyValue(v string) *DataHotelsValueOffersCancelPolicyPenalties {
	s.PenaltyValue = &v
	return s
}

func (s *DataHotelsValueOffersCancelPolicyPenalties) SetCurrency(v string) *DataHotelsValueOffersCancelPolicyPenalties {
	s.Currency = &v
	return s
}

func (s *DataHotelsValueOffersCancelPolicyPenalties) SetTracerId(v string) *DataHotelsValueOffersCancelPolicyPenalties {
	s.TracerId = &v
	return s
}

func (s *DataHotelsValueOffersCancelPolicyPenalties) Validate() error {
	return dara.Validate(s)
}

type DataHotelsValueOffersTotalPrice struct {
	// The amount in the smallest currency unit.
	//
	// example:
	//
	// 287
	Amount *string `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// The currency code.
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

func (s DataHotelsValueOffersTotalPrice) String() string {
	return dara.Prettify(s)
}

func (s DataHotelsValueOffersTotalPrice) GoString() string {
	return s.String()
}

func (s *DataHotelsValueOffersTotalPrice) GetAmount() *string {
	return s.Amount
}

func (s *DataHotelsValueOffersTotalPrice) GetCurrency() *string {
	return s.Currency
}

func (s *DataHotelsValueOffersTotalPrice) GetTracerId() *string {
	return s.TracerId
}

func (s *DataHotelsValueOffersTotalPrice) SetAmount(v string) *DataHotelsValueOffersTotalPrice {
	s.Amount = &v
	return s
}

func (s *DataHotelsValueOffersTotalPrice) SetCurrency(v string) *DataHotelsValueOffersTotalPrice {
	s.Currency = &v
	return s
}

func (s *DataHotelsValueOffersTotalPrice) SetTracerId(v string) *DataHotelsValueOffersTotalPrice {
	s.TracerId = &v
	return s
}

func (s *DataHotelsValueOffersTotalPrice) Validate() error {
	return dara.Validate(s)
}

type DataHotelsValueOffersDailyPrices struct {
	// The check-in date in yyyy-MM-dd format, based on the hotel local time zone.
	//
	// example:
	//
	// 2026-08-16
	Date *string `json:"Date,omitempty" xml:"Date,omitempty"`
	// The price for the day.
	Price *DataHotelsValueOffersDailyPricesPrice `json:"Price,omitempty" xml:"Price,omitempty" type:"Struct"`
	// null
	//
	// example:
	//
	// null
	TracerId *string `json:"TracerId,omitempty" xml:"TracerId,omitempty"`
}

func (s DataHotelsValueOffersDailyPrices) String() string {
	return dara.Prettify(s)
}

func (s DataHotelsValueOffersDailyPrices) GoString() string {
	return s.String()
}

func (s *DataHotelsValueOffersDailyPrices) GetDate() *string {
	return s.Date
}

func (s *DataHotelsValueOffersDailyPrices) GetPrice() *DataHotelsValueOffersDailyPricesPrice {
	return s.Price
}

func (s *DataHotelsValueOffersDailyPrices) GetTracerId() *string {
	return s.TracerId
}

func (s *DataHotelsValueOffersDailyPrices) SetDate(v string) *DataHotelsValueOffersDailyPrices {
	s.Date = &v
	return s
}

func (s *DataHotelsValueOffersDailyPrices) SetPrice(v *DataHotelsValueOffersDailyPricesPrice) *DataHotelsValueOffersDailyPrices {
	s.Price = v
	return s
}

func (s *DataHotelsValueOffersDailyPrices) SetTracerId(v string) *DataHotelsValueOffersDailyPrices {
	s.TracerId = &v
	return s
}

func (s *DataHotelsValueOffersDailyPrices) Validate() error {
	if s.Price != nil {
		if err := s.Price.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DataHotelsValueOffersDailyPricesPrice struct {
	// The amount in the smallest currency unit.
	//
	// example:
	//
	// 287
	Amount *string `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// The currency code.
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

func (s DataHotelsValueOffersDailyPricesPrice) String() string {
	return dara.Prettify(s)
}

func (s DataHotelsValueOffersDailyPricesPrice) GoString() string {
	return s.String()
}

func (s *DataHotelsValueOffersDailyPricesPrice) GetAmount() *string {
	return s.Amount
}

func (s *DataHotelsValueOffersDailyPricesPrice) GetCurrency() *string {
	return s.Currency
}

func (s *DataHotelsValueOffersDailyPricesPrice) GetTracerId() *string {
	return s.TracerId
}

func (s *DataHotelsValueOffersDailyPricesPrice) SetAmount(v string) *DataHotelsValueOffersDailyPricesPrice {
	s.Amount = &v
	return s
}

func (s *DataHotelsValueOffersDailyPricesPrice) SetCurrency(v string) *DataHotelsValueOffersDailyPricesPrice {
	s.Currency = &v
	return s
}

func (s *DataHotelsValueOffersDailyPricesPrice) SetTracerId(v string) *DataHotelsValueOffersDailyPricesPrice {
	s.TracerId = &v
	return s
}

func (s *DataHotelsValueOffersDailyPricesPrice) Validate() error {
	return dara.Validate(s)
}
