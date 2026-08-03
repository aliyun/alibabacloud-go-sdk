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
	// example:
	//
	// 2026-01-01
	CheckInDate *string `json:"CheckInDate,omitempty" xml:"CheckInDate,omitempty"`
	// example:
	//
	// 2026-01-02
	CheckOutDate *string                 `json:"CheckOutDate,omitempty" xml:"CheckOutDate,omitempty"`
	Rooms        []*DataHotelsValueRooms `json:"Rooms,omitempty" xml:"Rooms,omitempty" type:"Repeated"`
	// example:
	//
	// R001
	StandardRoomId *string                  `json:"StandardRoomId,omitempty" xml:"StandardRoomId,omitempty"`
	Offers         []*DataHotelsValueOffers `json:"Offers,omitempty" xml:"Offers,omitempty" type:"Repeated"`
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
	// example:
	//
	// R001
	StandardRoomId     *string                                 `json:"StandardRoomId,omitempty" xml:"StandardRoomId,omitempty"`
	LowestSellingPrice *DataHotelsValueRoomsLowestSellingPrice `json:"LowestSellingPrice,omitempty" xml:"LowestSellingPrice,omitempty" type:"Struct"`
	Offers             []*DataHotelsValueRoomsOffers           `json:"Offers,omitempty" xml:"Offers,omitempty" type:"Repeated"`
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

func (s *DataHotelsValueRooms) GetLowestSellingPrice() *DataHotelsValueRoomsLowestSellingPrice {
	return s.LowestSellingPrice
}

func (s *DataHotelsValueRooms) GetOffers() []*DataHotelsValueRoomsOffers {
	return s.Offers
}

func (s *DataHotelsValueRooms) SetStandardRoomId(v string) *DataHotelsValueRooms {
	s.StandardRoomId = &v
	return s
}

func (s *DataHotelsValueRooms) SetLowestSellingPrice(v *DataHotelsValueRoomsLowestSellingPrice) *DataHotelsValueRooms {
	s.LowestSellingPrice = v
	return s
}

func (s *DataHotelsValueRooms) SetOffers(v []*DataHotelsValueRoomsOffers) *DataHotelsValueRooms {
	s.Offers = v
	return s
}

func (s *DataHotelsValueRooms) Validate() error {
	if s.LowestSellingPrice != nil {
		if err := s.LowestSellingPrice.Validate(); err != nil {
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

type DataHotelsValueRoomsLowestSellingPrice struct {
	// example:
	//
	// 100.00
	Amount *float64 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// example:
	//
	// USD
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// example:
	//
	// TraceId
	TracerId *string `json:"TracerId,omitempty" xml:"TracerId,omitempty"`
}

func (s DataHotelsValueRoomsLowestSellingPrice) String() string {
	return dara.Prettify(s)
}

func (s DataHotelsValueRoomsLowestSellingPrice) GoString() string {
	return s.String()
}

func (s *DataHotelsValueRoomsLowestSellingPrice) GetAmount() *float64 {
	return s.Amount
}

func (s *DataHotelsValueRoomsLowestSellingPrice) GetCurrency() *string {
	return s.Currency
}

func (s *DataHotelsValueRoomsLowestSellingPrice) GetTracerId() *string {
	return s.TracerId
}

func (s *DataHotelsValueRoomsLowestSellingPrice) SetAmount(v float64) *DataHotelsValueRoomsLowestSellingPrice {
	s.Amount = &v
	return s
}

func (s *DataHotelsValueRoomsLowestSellingPrice) SetCurrency(v string) *DataHotelsValueRoomsLowestSellingPrice {
	s.Currency = &v
	return s
}

func (s *DataHotelsValueRoomsLowestSellingPrice) SetTracerId(v string) *DataHotelsValueRoomsLowestSellingPrice {
	s.TracerId = &v
	return s
}

func (s *DataHotelsValueRoomsLowestSellingPrice) Validate() error {
	return dara.Validate(s)
}

type DataHotelsValueRoomsOffers struct {
	// example:
	//
	// itemOffer_123
	ItemOfferKey *string `json:"ItemOfferKey,omitempty" xml:"ItemOfferKey,omitempty"`
	// example:
	//
	// 含早房
	RatePlanName *string `json:"RatePlanName,omitempty" xml:"RatePlanName,omitempty"`
	// example:
	//
	// BREAKFAST
	MealType *string `json:"MealType,omitempty" xml:"MealType,omitempty"`
	// example:
	//
	// 2
	MealCount          *int32                                          `json:"MealCount,omitempty" xml:"MealCount,omitempty"`
	CancelPolicy       *DataHotelsValueRoomsOffersCancelPolicy         `json:"CancelPolicy,omitempty" xml:"CancelPolicy,omitempty" type:"Struct"`
	SellingTotalPrice  *DataHotelsValueRoomsOffersSellingTotalPrice    `json:"SellingTotalPrice,omitempty" xml:"SellingTotalPrice,omitempty" type:"Struct"`
	SellingDailyPrices []*DataHotelsValueRoomsOffersSellingDailyPrices `json:"SellingDailyPrices,omitempty" xml:"SellingDailyPrices,omitempty" type:"Repeated"`
	// example:
	//
	// 5
	AvailableRooms *int32 `json:"AvailableRooms,omitempty" xml:"AvailableRooms,omitempty"`
	// example:
	//
	// 3
	MaxOccupancy *int32 `json:"MaxOccupancy,omitempty" xml:"MaxOccupancy,omitempty"`
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

func (s *DataHotelsValueRoomsOffers) GetSellingTotalPrice() *DataHotelsValueRoomsOffersSellingTotalPrice {
	return s.SellingTotalPrice
}

func (s *DataHotelsValueRoomsOffers) GetSellingDailyPrices() []*DataHotelsValueRoomsOffersSellingDailyPrices {
	return s.SellingDailyPrices
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

func (s *DataHotelsValueRoomsOffers) SetSellingTotalPrice(v *DataHotelsValueRoomsOffersSellingTotalPrice) *DataHotelsValueRoomsOffers {
	s.SellingTotalPrice = v
	return s
}

func (s *DataHotelsValueRoomsOffers) SetSellingDailyPrices(v []*DataHotelsValueRoomsOffersSellingDailyPrices) *DataHotelsValueRoomsOffers {
	s.SellingDailyPrices = v
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
	if s.SellingTotalPrice != nil {
		if err := s.SellingTotalPrice.Validate(); err != nil {
			return err
		}
	}
	if s.SellingDailyPrices != nil {
		for _, item := range s.SellingDailyPrices {
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
	// example:
	//
	// FREE_CANCELLATION
	PolicyType *string                                            `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	Penalties  []*DataHotelsValueRoomsOffersCancelPolicyPenalties `json:"Penalties,omitempty" xml:"Penalties,omitempty" type:"Repeated"`
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
	// example:
	//
	// 1672531200000
	Start *int64 `json:"Start,omitempty" xml:"Start,omitempty"`
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
	// USD
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
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

func (s *DataHotelsValueRoomsOffersCancelPolicyPenalties) GetStart() *int64 {
	return s.Start
}

func (s *DataHotelsValueRoomsOffersCancelPolicyPenalties) GetEnd() *int64 {
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

func (s *DataHotelsValueRoomsOffersCancelPolicyPenalties) SetStart(v int64) *DataHotelsValueRoomsOffersCancelPolicyPenalties {
	s.Start = &v
	return s
}

func (s *DataHotelsValueRoomsOffersCancelPolicyPenalties) SetEnd(v int64) *DataHotelsValueRoomsOffersCancelPolicyPenalties {
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

type DataHotelsValueRoomsOffersSellingTotalPrice struct {
	// example:
	//
	// 100.00
	Amount *float64 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// example:
	//
	// USD
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// example:
	//
	// TraceId
	TracerId *string `json:"TracerId,omitempty" xml:"TracerId,omitempty"`
}

func (s DataHotelsValueRoomsOffersSellingTotalPrice) String() string {
	return dara.Prettify(s)
}

func (s DataHotelsValueRoomsOffersSellingTotalPrice) GoString() string {
	return s.String()
}

func (s *DataHotelsValueRoomsOffersSellingTotalPrice) GetAmount() *float64 {
	return s.Amount
}

func (s *DataHotelsValueRoomsOffersSellingTotalPrice) GetCurrency() *string {
	return s.Currency
}

func (s *DataHotelsValueRoomsOffersSellingTotalPrice) GetTracerId() *string {
	return s.TracerId
}

func (s *DataHotelsValueRoomsOffersSellingTotalPrice) SetAmount(v float64) *DataHotelsValueRoomsOffersSellingTotalPrice {
	s.Amount = &v
	return s
}

func (s *DataHotelsValueRoomsOffersSellingTotalPrice) SetCurrency(v string) *DataHotelsValueRoomsOffersSellingTotalPrice {
	s.Currency = &v
	return s
}

func (s *DataHotelsValueRoomsOffersSellingTotalPrice) SetTracerId(v string) *DataHotelsValueRoomsOffersSellingTotalPrice {
	s.TracerId = &v
	return s
}

func (s *DataHotelsValueRoomsOffersSellingTotalPrice) Validate() error {
	return dara.Validate(s)
}

type DataHotelsValueRoomsOffersSellingDailyPrices struct {
	// example:
	//
	// 2026-07-01
	Date  *string                                            `json:"Date,omitempty" xml:"Date,omitempty"`
	Price *DataHotelsValueRoomsOffersSellingDailyPricesPrice `json:"Price,omitempty" xml:"Price,omitempty" type:"Struct"`
	// example:
	//
	// TraceId
	TracerId *string `json:"TracerId,omitempty" xml:"TracerId,omitempty"`
}

func (s DataHotelsValueRoomsOffersSellingDailyPrices) String() string {
	return dara.Prettify(s)
}

func (s DataHotelsValueRoomsOffersSellingDailyPrices) GoString() string {
	return s.String()
}

func (s *DataHotelsValueRoomsOffersSellingDailyPrices) GetDate() *string {
	return s.Date
}

func (s *DataHotelsValueRoomsOffersSellingDailyPrices) GetPrice() *DataHotelsValueRoomsOffersSellingDailyPricesPrice {
	return s.Price
}

func (s *DataHotelsValueRoomsOffersSellingDailyPrices) GetTracerId() *string {
	return s.TracerId
}

func (s *DataHotelsValueRoomsOffersSellingDailyPrices) SetDate(v string) *DataHotelsValueRoomsOffersSellingDailyPrices {
	s.Date = &v
	return s
}

func (s *DataHotelsValueRoomsOffersSellingDailyPrices) SetPrice(v *DataHotelsValueRoomsOffersSellingDailyPricesPrice) *DataHotelsValueRoomsOffersSellingDailyPrices {
	s.Price = v
	return s
}

func (s *DataHotelsValueRoomsOffersSellingDailyPrices) SetTracerId(v string) *DataHotelsValueRoomsOffersSellingDailyPrices {
	s.TracerId = &v
	return s
}

func (s *DataHotelsValueRoomsOffersSellingDailyPrices) Validate() error {
	if s.Price != nil {
		if err := s.Price.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DataHotelsValueRoomsOffersSellingDailyPricesPrice struct {
	// example:
	//
	// 100.00
	Amount *float64 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// example:
	//
	// USD
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// example:
	//
	// TraceId
	TracerId *string `json:"TracerId,omitempty" xml:"TracerId,omitempty"`
}

func (s DataHotelsValueRoomsOffersSellingDailyPricesPrice) String() string {
	return dara.Prettify(s)
}

func (s DataHotelsValueRoomsOffersSellingDailyPricesPrice) GoString() string {
	return s.String()
}

func (s *DataHotelsValueRoomsOffersSellingDailyPricesPrice) GetAmount() *float64 {
	return s.Amount
}

func (s *DataHotelsValueRoomsOffersSellingDailyPricesPrice) GetCurrency() *string {
	return s.Currency
}

func (s *DataHotelsValueRoomsOffersSellingDailyPricesPrice) GetTracerId() *string {
	return s.TracerId
}

func (s *DataHotelsValueRoomsOffersSellingDailyPricesPrice) SetAmount(v float64) *DataHotelsValueRoomsOffersSellingDailyPricesPrice {
	s.Amount = &v
	return s
}

func (s *DataHotelsValueRoomsOffersSellingDailyPricesPrice) SetCurrency(v string) *DataHotelsValueRoomsOffersSellingDailyPricesPrice {
	s.Currency = &v
	return s
}

func (s *DataHotelsValueRoomsOffersSellingDailyPricesPrice) SetTracerId(v string) *DataHotelsValueRoomsOffersSellingDailyPricesPrice {
	s.TracerId = &v
	return s
}

func (s *DataHotelsValueRoomsOffersSellingDailyPricesPrice) Validate() error {
	return dara.Validate(s)
}

type DataHotelsValueOffers struct {
	// example:
	//
	// itemOffer_123
	ItemOfferKey *string `json:"ItemOfferKey,omitempty" xml:"ItemOfferKey,omitempty"`
	// example:
	//
	// 含早房
	RatePlanName *string `json:"RatePlanName,omitempty" xml:"RatePlanName,omitempty"`
	// example:
	//
	// BREAKFAST
	MealType *string `json:"MealType,omitempty" xml:"MealType,omitempty"`
	// example:
	//
	// 2
	MealCount          *int32                                     `json:"MealCount,omitempty" xml:"MealCount,omitempty"`
	CancelPolicy       *DataHotelsValueOffersCancelPolicy         `json:"CancelPolicy,omitempty" xml:"CancelPolicy,omitempty" type:"Struct"`
	SellingTotalPrice  *DataHotelsValueOffersSellingTotalPrice    `json:"SellingTotalPrice,omitempty" xml:"SellingTotalPrice,omitempty" type:"Struct"`
	SellingDailyPrices []*DataHotelsValueOffersSellingDailyPrices `json:"SellingDailyPrices,omitempty" xml:"SellingDailyPrices,omitempty" type:"Repeated"`
	// example:
	//
	// 5
	AvailableRooms *int32 `json:"AvailableRooms,omitempty" xml:"AvailableRooms,omitempty"`
	// example:
	//
	// 3
	MaxOccupancy *int32 `json:"MaxOccupancy,omitempty" xml:"MaxOccupancy,omitempty"`
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

func (s *DataHotelsValueOffers) GetSellingTotalPrice() *DataHotelsValueOffersSellingTotalPrice {
	return s.SellingTotalPrice
}

func (s *DataHotelsValueOffers) GetSellingDailyPrices() []*DataHotelsValueOffersSellingDailyPrices {
	return s.SellingDailyPrices
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

func (s *DataHotelsValueOffers) SetSellingTotalPrice(v *DataHotelsValueOffersSellingTotalPrice) *DataHotelsValueOffers {
	s.SellingTotalPrice = v
	return s
}

func (s *DataHotelsValueOffers) SetSellingDailyPrices(v []*DataHotelsValueOffersSellingDailyPrices) *DataHotelsValueOffers {
	s.SellingDailyPrices = v
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
	if s.SellingTotalPrice != nil {
		if err := s.SellingTotalPrice.Validate(); err != nil {
			return err
		}
	}
	if s.SellingDailyPrices != nil {
		for _, item := range s.SellingDailyPrices {
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
	// example:
	//
	// FREE_CANCELLATION
	PolicyType *string                                       `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	Penalties  []*DataHotelsValueOffersCancelPolicyPenalties `json:"Penalties,omitempty" xml:"Penalties,omitempty" type:"Repeated"`
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
	// example:
	//
	// 1672531200000
	Start *int64 `json:"Start,omitempty" xml:"Start,omitempty"`
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
	// USD
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
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

func (s *DataHotelsValueOffersCancelPolicyPenalties) GetStart() *int64 {
	return s.Start
}

func (s *DataHotelsValueOffersCancelPolicyPenalties) GetEnd() *int64 {
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

func (s *DataHotelsValueOffersCancelPolicyPenalties) SetStart(v int64) *DataHotelsValueOffersCancelPolicyPenalties {
	s.Start = &v
	return s
}

func (s *DataHotelsValueOffersCancelPolicyPenalties) SetEnd(v int64) *DataHotelsValueOffersCancelPolicyPenalties {
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

type DataHotelsValueOffersSellingTotalPrice struct {
	// example:
	//
	// 100.00
	Amount *float64 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// example:
	//
	// USD
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// example:
	//
	// TraceId
	TracerId *string `json:"TracerId,omitempty" xml:"TracerId,omitempty"`
}

func (s DataHotelsValueOffersSellingTotalPrice) String() string {
	return dara.Prettify(s)
}

func (s DataHotelsValueOffersSellingTotalPrice) GoString() string {
	return s.String()
}

func (s *DataHotelsValueOffersSellingTotalPrice) GetAmount() *float64 {
	return s.Amount
}

func (s *DataHotelsValueOffersSellingTotalPrice) GetCurrency() *string {
	return s.Currency
}

func (s *DataHotelsValueOffersSellingTotalPrice) GetTracerId() *string {
	return s.TracerId
}

func (s *DataHotelsValueOffersSellingTotalPrice) SetAmount(v float64) *DataHotelsValueOffersSellingTotalPrice {
	s.Amount = &v
	return s
}

func (s *DataHotelsValueOffersSellingTotalPrice) SetCurrency(v string) *DataHotelsValueOffersSellingTotalPrice {
	s.Currency = &v
	return s
}

func (s *DataHotelsValueOffersSellingTotalPrice) SetTracerId(v string) *DataHotelsValueOffersSellingTotalPrice {
	s.TracerId = &v
	return s
}

func (s *DataHotelsValueOffersSellingTotalPrice) Validate() error {
	return dara.Validate(s)
}

type DataHotelsValueOffersSellingDailyPrices struct {
	// example:
	//
	// 2026-07-01
	Date  *string                                       `json:"Date,omitempty" xml:"Date,omitempty"`
	Price *DataHotelsValueOffersSellingDailyPricesPrice `json:"Price,omitempty" xml:"Price,omitempty" type:"Struct"`
	// example:
	//
	// TraceId
	TracerId *string `json:"TracerId,omitempty" xml:"TracerId,omitempty"`
}

func (s DataHotelsValueOffersSellingDailyPrices) String() string {
	return dara.Prettify(s)
}

func (s DataHotelsValueOffersSellingDailyPrices) GoString() string {
	return s.String()
}

func (s *DataHotelsValueOffersSellingDailyPrices) GetDate() *string {
	return s.Date
}

func (s *DataHotelsValueOffersSellingDailyPrices) GetPrice() *DataHotelsValueOffersSellingDailyPricesPrice {
	return s.Price
}

func (s *DataHotelsValueOffersSellingDailyPrices) GetTracerId() *string {
	return s.TracerId
}

func (s *DataHotelsValueOffersSellingDailyPrices) SetDate(v string) *DataHotelsValueOffersSellingDailyPrices {
	s.Date = &v
	return s
}

func (s *DataHotelsValueOffersSellingDailyPrices) SetPrice(v *DataHotelsValueOffersSellingDailyPricesPrice) *DataHotelsValueOffersSellingDailyPrices {
	s.Price = v
	return s
}

func (s *DataHotelsValueOffersSellingDailyPrices) SetTracerId(v string) *DataHotelsValueOffersSellingDailyPrices {
	s.TracerId = &v
	return s
}

func (s *DataHotelsValueOffersSellingDailyPrices) Validate() error {
	if s.Price != nil {
		if err := s.Price.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DataHotelsValueOffersSellingDailyPricesPrice struct {
	// example:
	//
	// 100.00
	Amount *float64 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// example:
	//
	// USD
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// example:
	//
	// TraceId
	TracerId *string `json:"TracerId,omitempty" xml:"TracerId,omitempty"`
}

func (s DataHotelsValueOffersSellingDailyPricesPrice) String() string {
	return dara.Prettify(s)
}

func (s DataHotelsValueOffersSellingDailyPricesPrice) GoString() string {
	return s.String()
}

func (s *DataHotelsValueOffersSellingDailyPricesPrice) GetAmount() *float64 {
	return s.Amount
}

func (s *DataHotelsValueOffersSellingDailyPricesPrice) GetCurrency() *string {
	return s.Currency
}

func (s *DataHotelsValueOffersSellingDailyPricesPrice) GetTracerId() *string {
	return s.TracerId
}

func (s *DataHotelsValueOffersSellingDailyPricesPrice) SetAmount(v float64) *DataHotelsValueOffersSellingDailyPricesPrice {
	s.Amount = &v
	return s
}

func (s *DataHotelsValueOffersSellingDailyPricesPrice) SetCurrency(v string) *DataHotelsValueOffersSellingDailyPricesPrice {
	s.Currency = &v
	return s
}

func (s *DataHotelsValueOffersSellingDailyPricesPrice) SetTracerId(v string) *DataHotelsValueOffersSellingDailyPricesPrice {
	s.TracerId = &v
	return s
}

func (s *DataHotelsValueOffersSellingDailyPricesPrice) Validate() error {
	return dara.Validate(s)
}
