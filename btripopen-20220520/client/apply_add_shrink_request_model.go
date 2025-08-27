// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyAddShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBudget(v int64) *ApplyAddShrinkRequest
	GetBudget() *int64
	SetBudgetMerge(v int32) *ApplyAddShrinkRequest
	GetBudgetMerge() *int32
	SetCarRuleShrink(v string) *ApplyAddShrinkRequest
	GetCarRuleShrink() *string
	SetCorpName(v string) *ApplyAddShrinkRequest
	GetCorpName() *string
	SetDefaultStandardShrink(v string) *ApplyAddShrinkRequest
	GetDefaultStandardShrink() *string
	SetDepartId(v string) *ApplyAddShrinkRequest
	GetDepartId() *string
	SetDepartName(v string) *ApplyAddShrinkRequest
	GetDepartName() *string
	SetExtendField(v string) *ApplyAddShrinkRequest
	GetExtendField() *string
	SetExternalTravelerListShrink(v string) *ApplyAddShrinkRequest
	GetExternalTravelerListShrink() *string
	SetExternalTravelerStandardShrink(v string) *ApplyAddShrinkRequest
	GetExternalTravelerStandardShrink() *string
	SetFlightBudget(v int64) *ApplyAddShrinkRequest
	GetFlightBudget() *int64
	SetHotelBudget(v int64) *ApplyAddShrinkRequest
	GetHotelBudget() *int64
	SetHotelShareShrink(v string) *ApplyAddShrinkRequest
	GetHotelShareShrink() *string
	SetInternationalFlightCabins(v string) *ApplyAddShrinkRequest
	GetInternationalFlightCabins() *string
	SetIntlFlightBudget(v int64) *ApplyAddShrinkRequest
	GetIntlFlightBudget() *int64
	SetIntlHotelBudget(v int64) *ApplyAddShrinkRequest
	GetIntlHotelBudget() *int64
	SetItineraryListShrink(v string) *ApplyAddShrinkRequest
	GetItineraryListShrink() *string
	SetItineraryRule(v int32) *ApplyAddShrinkRequest
	GetItineraryRule() *int32
	SetItinerarySetListShrink(v string) *ApplyAddShrinkRequest
	GetItinerarySetListShrink() *string
	SetLimitTraveler(v int32) *ApplyAddShrinkRequest
	GetLimitTraveler() *int32
	SetMealBudget(v int64) *ApplyAddShrinkRequest
	GetMealBudget() *int64
	SetPaymentDepartmentId(v string) *ApplyAddShrinkRequest
	GetPaymentDepartmentId() *string
	SetPaymentDepartmentName(v string) *ApplyAddShrinkRequest
	GetPaymentDepartmentName() *string
	SetStatus(v int32) *ApplyAddShrinkRequest
	GetStatus() *int32
	SetSubCorpId(v string) *ApplyAddShrinkRequest
	GetSubCorpId() *string
	SetThirdpartApplyId(v string) *ApplyAddShrinkRequest
	GetThirdpartApplyId() *string
	SetThirdpartBusinessId(v string) *ApplyAddShrinkRequest
	GetThirdpartBusinessId() *string
	SetThirdpartDepartId(v string) *ApplyAddShrinkRequest
	GetThirdpartDepartId() *string
	SetTogetherBookRule(v int32) *ApplyAddShrinkRequest
	GetTogetherBookRule() *int32
	SetTrainBudget(v int64) *ApplyAddShrinkRequest
	GetTrainBudget() *int64
	SetTravelerListShrink(v string) *ApplyAddShrinkRequest
	GetTravelerListShrink() *string
	SetTravelerStandardShrink(v string) *ApplyAddShrinkRequest
	GetTravelerStandardShrink() *string
	SetTripCause(v string) *ApplyAddShrinkRequest
	GetTripCause() *string
	SetTripDay(v int32) *ApplyAddShrinkRequest
	GetTripDay() *int32
	SetTripTitle(v string) *ApplyAddShrinkRequest
	GetTripTitle() *string
	SetType(v int32) *ApplyAddShrinkRequest
	GetType() *int32
	SetUnionNo(v string) *ApplyAddShrinkRequest
	GetUnionNo() *string
	SetUserId(v string) *ApplyAddShrinkRequest
	GetUserId() *string
	SetUserName(v string) *ApplyAddShrinkRequest
	GetUserName() *string
	SetVehicleBudget(v int64) *ApplyAddShrinkRequest
	GetVehicleBudget() *int64
}

type ApplyAddShrinkRequest struct {
	// example:
	//
	// 4000
	Budget *int64 `json:"budget,omitempty" xml:"budget,omitempty"`
	// example:
	//
	// 1
	BudgetMerge           *int32  `json:"budget_merge,omitempty" xml:"budget_merge,omitempty"`
	CarRuleShrink         *string `json:"car_rule,omitempty" xml:"car_rule,omitempty"`
	CorpName              *string `json:"corp_name,omitempty" xml:"corp_name,omitempty"`
	DefaultStandardShrink *string `json:"default_standard,omitempty" xml:"default_standard,omitempty"`
	// example:
	//
	// departId01
	DepartId   *string `json:"depart_id,omitempty" xml:"depart_id,omitempty"`
	DepartName *string `json:"depart_name,omitempty" xml:"depart_name,omitempty"`
	// 可将补充描述传入此字段，账单中将会体现此字段的值。可以用于企业的统计和对账
	//
	// example:
	//
	// {"cost_center":"成本中心"}
	ExtendField                    *string `json:"extend_field,omitempty" xml:"extend_field,omitempty"`
	ExternalTravelerListShrink     *string `json:"external_traveler_list,omitempty" xml:"external_traveler_list,omitempty"`
	ExternalTravelerStandardShrink *string `json:"external_traveler_standard,omitempty" xml:"external_traveler_standard,omitempty"`
	// example:
	//
	// 1000
	FlightBudget *int64 `json:"flight_budget,omitempty" xml:"flight_budget,omitempty"`
	// example:
	//
	// 1000
	HotelBudget               *int64  `json:"hotel_budget,omitempty" xml:"hotel_budget,omitempty"`
	HotelShareShrink          *string `json:"hotel_share,omitempty" xml:"hotel_share,omitempty"`
	InternationalFlightCabins *string `json:"international_flight_cabins,omitempty" xml:"international_flight_cabins,omitempty"`
	IntlFlightBudget          *int64  `json:"intl_flight_budget,omitempty" xml:"intl_flight_budget,omitempty"`
	IntlHotelBudget           *int64  `json:"intl_hotel_budget,omitempty" xml:"intl_hotel_budget,omitempty"`
	ItineraryListShrink       *string `json:"itinerary_list,omitempty" xml:"itinerary_list,omitempty"`
	// example:
	//
	// 0
	ItineraryRule          *int32  `json:"itinerary_rule,omitempty" xml:"itinerary_rule,omitempty"`
	ItinerarySetListShrink *string `json:"itinerary_set_list,omitempty" xml:"itinerary_set_list,omitempty"`
	// example:
	//
	// 1
	LimitTraveler         *int32  `json:"limit_traveler,omitempty" xml:"limit_traveler,omitempty"`
	MealBudget            *int64  `json:"meal_budget,omitempty" xml:"meal_budget,omitempty"`
	PaymentDepartmentId   *string `json:"payment_department_id,omitempty" xml:"payment_department_id,omitempty"`
	PaymentDepartmentName *string `json:"payment_department_name,omitempty" xml:"payment_department_name,omitempty"`
	// example:
	//
	// 0
	Status    *int32  `json:"status,omitempty" xml:"status,omitempty"`
	SubCorpId *string `json:"sub_corp_id,omitempty" xml:"sub_corp_id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 12345
	ThirdpartApplyId *string `json:"thirdpart_apply_id,omitempty" xml:"thirdpart_apply_id,omitempty"`
	// example:
	//
	// 00714131
	ThirdpartBusinessId *string `json:"thirdpart_business_id,omitempty" xml:"thirdpart_business_id,omitempty"`
	ThirdpartDepartId   *string `json:"thirdpart_depart_id,omitempty" xml:"thirdpart_depart_id,omitempty"`
	// example:
	//
	// 1
	TogetherBookRule *int32 `json:"together_book_rule,omitempty" xml:"together_book_rule,omitempty"`
	// example:
	//
	// 1000
	TrainBudget            *int64  `json:"train_budget,omitempty" xml:"train_budget,omitempty"`
	TravelerListShrink     *string `json:"traveler_list,omitempty" xml:"traveler_list,omitempty"`
	TravelerStandardShrink *string `json:"traveler_standard,omitempty" xml:"traveler_standard,omitempty"`
	// This parameter is required.
	TripCause *string `json:"trip_cause,omitempty" xml:"trip_cause,omitempty"`
	// example:
	//
	// 1
	TripDay *int32 `json:"trip_day,omitempty" xml:"trip_day,omitempty"`
	// This parameter is required.
	TripTitle *string `json:"trip_title,omitempty" xml:"trip_title,omitempty"`
	// example:
	//
	// 1
	Type *int32 `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// union001
	UnionNo *string `json:"union_no,omitempty" xml:"union_no,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// thridpart12138
	UserId   *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	UserName *string `json:"user_name,omitempty" xml:"user_name,omitempty"`
	// example:
	//
	// 1000
	VehicleBudget *int64 `json:"vehicle_budget,omitempty" xml:"vehicle_budget,omitempty"`
}

func (s ApplyAddShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ApplyAddShrinkRequest) GoString() string {
	return s.String()
}

func (s *ApplyAddShrinkRequest) GetBudget() *int64 {
	return s.Budget
}

func (s *ApplyAddShrinkRequest) GetBudgetMerge() *int32 {
	return s.BudgetMerge
}

func (s *ApplyAddShrinkRequest) GetCarRuleShrink() *string {
	return s.CarRuleShrink
}

func (s *ApplyAddShrinkRequest) GetCorpName() *string {
	return s.CorpName
}

func (s *ApplyAddShrinkRequest) GetDefaultStandardShrink() *string {
	return s.DefaultStandardShrink
}

func (s *ApplyAddShrinkRequest) GetDepartId() *string {
	return s.DepartId
}

func (s *ApplyAddShrinkRequest) GetDepartName() *string {
	return s.DepartName
}

func (s *ApplyAddShrinkRequest) GetExtendField() *string {
	return s.ExtendField
}

func (s *ApplyAddShrinkRequest) GetExternalTravelerListShrink() *string {
	return s.ExternalTravelerListShrink
}

func (s *ApplyAddShrinkRequest) GetExternalTravelerStandardShrink() *string {
	return s.ExternalTravelerStandardShrink
}

func (s *ApplyAddShrinkRequest) GetFlightBudget() *int64 {
	return s.FlightBudget
}

func (s *ApplyAddShrinkRequest) GetHotelBudget() *int64 {
	return s.HotelBudget
}

func (s *ApplyAddShrinkRequest) GetHotelShareShrink() *string {
	return s.HotelShareShrink
}

func (s *ApplyAddShrinkRequest) GetInternationalFlightCabins() *string {
	return s.InternationalFlightCabins
}

func (s *ApplyAddShrinkRequest) GetIntlFlightBudget() *int64 {
	return s.IntlFlightBudget
}

func (s *ApplyAddShrinkRequest) GetIntlHotelBudget() *int64 {
	return s.IntlHotelBudget
}

func (s *ApplyAddShrinkRequest) GetItineraryListShrink() *string {
	return s.ItineraryListShrink
}

func (s *ApplyAddShrinkRequest) GetItineraryRule() *int32 {
	return s.ItineraryRule
}

func (s *ApplyAddShrinkRequest) GetItinerarySetListShrink() *string {
	return s.ItinerarySetListShrink
}

func (s *ApplyAddShrinkRequest) GetLimitTraveler() *int32 {
	return s.LimitTraveler
}

func (s *ApplyAddShrinkRequest) GetMealBudget() *int64 {
	return s.MealBudget
}

func (s *ApplyAddShrinkRequest) GetPaymentDepartmentId() *string {
	return s.PaymentDepartmentId
}

func (s *ApplyAddShrinkRequest) GetPaymentDepartmentName() *string {
	return s.PaymentDepartmentName
}

func (s *ApplyAddShrinkRequest) GetStatus() *int32 {
	return s.Status
}

func (s *ApplyAddShrinkRequest) GetSubCorpId() *string {
	return s.SubCorpId
}

func (s *ApplyAddShrinkRequest) GetThirdpartApplyId() *string {
	return s.ThirdpartApplyId
}

func (s *ApplyAddShrinkRequest) GetThirdpartBusinessId() *string {
	return s.ThirdpartBusinessId
}

func (s *ApplyAddShrinkRequest) GetThirdpartDepartId() *string {
	return s.ThirdpartDepartId
}

func (s *ApplyAddShrinkRequest) GetTogetherBookRule() *int32 {
	return s.TogetherBookRule
}

func (s *ApplyAddShrinkRequest) GetTrainBudget() *int64 {
	return s.TrainBudget
}

func (s *ApplyAddShrinkRequest) GetTravelerListShrink() *string {
	return s.TravelerListShrink
}

func (s *ApplyAddShrinkRequest) GetTravelerStandardShrink() *string {
	return s.TravelerStandardShrink
}

func (s *ApplyAddShrinkRequest) GetTripCause() *string {
	return s.TripCause
}

func (s *ApplyAddShrinkRequest) GetTripDay() *int32 {
	return s.TripDay
}

func (s *ApplyAddShrinkRequest) GetTripTitle() *string {
	return s.TripTitle
}

func (s *ApplyAddShrinkRequest) GetType() *int32 {
	return s.Type
}

func (s *ApplyAddShrinkRequest) GetUnionNo() *string {
	return s.UnionNo
}

func (s *ApplyAddShrinkRequest) GetUserId() *string {
	return s.UserId
}

func (s *ApplyAddShrinkRequest) GetUserName() *string {
	return s.UserName
}

func (s *ApplyAddShrinkRequest) GetVehicleBudget() *int64 {
	return s.VehicleBudget
}

func (s *ApplyAddShrinkRequest) SetBudget(v int64) *ApplyAddShrinkRequest {
	s.Budget = &v
	return s
}

func (s *ApplyAddShrinkRequest) SetBudgetMerge(v int32) *ApplyAddShrinkRequest {
	s.BudgetMerge = &v
	return s
}

func (s *ApplyAddShrinkRequest) SetCarRuleShrink(v string) *ApplyAddShrinkRequest {
	s.CarRuleShrink = &v
	return s
}

func (s *ApplyAddShrinkRequest) SetCorpName(v string) *ApplyAddShrinkRequest {
	s.CorpName = &v
	return s
}

func (s *ApplyAddShrinkRequest) SetDefaultStandardShrink(v string) *ApplyAddShrinkRequest {
	s.DefaultStandardShrink = &v
	return s
}

func (s *ApplyAddShrinkRequest) SetDepartId(v string) *ApplyAddShrinkRequest {
	s.DepartId = &v
	return s
}

func (s *ApplyAddShrinkRequest) SetDepartName(v string) *ApplyAddShrinkRequest {
	s.DepartName = &v
	return s
}

func (s *ApplyAddShrinkRequest) SetExtendField(v string) *ApplyAddShrinkRequest {
	s.ExtendField = &v
	return s
}

func (s *ApplyAddShrinkRequest) SetExternalTravelerListShrink(v string) *ApplyAddShrinkRequest {
	s.ExternalTravelerListShrink = &v
	return s
}

func (s *ApplyAddShrinkRequest) SetExternalTravelerStandardShrink(v string) *ApplyAddShrinkRequest {
	s.ExternalTravelerStandardShrink = &v
	return s
}

func (s *ApplyAddShrinkRequest) SetFlightBudget(v int64) *ApplyAddShrinkRequest {
	s.FlightBudget = &v
	return s
}

func (s *ApplyAddShrinkRequest) SetHotelBudget(v int64) *ApplyAddShrinkRequest {
	s.HotelBudget = &v
	return s
}

func (s *ApplyAddShrinkRequest) SetHotelShareShrink(v string) *ApplyAddShrinkRequest {
	s.HotelShareShrink = &v
	return s
}

func (s *ApplyAddShrinkRequest) SetInternationalFlightCabins(v string) *ApplyAddShrinkRequest {
	s.InternationalFlightCabins = &v
	return s
}

func (s *ApplyAddShrinkRequest) SetIntlFlightBudget(v int64) *ApplyAddShrinkRequest {
	s.IntlFlightBudget = &v
	return s
}

func (s *ApplyAddShrinkRequest) SetIntlHotelBudget(v int64) *ApplyAddShrinkRequest {
	s.IntlHotelBudget = &v
	return s
}

func (s *ApplyAddShrinkRequest) SetItineraryListShrink(v string) *ApplyAddShrinkRequest {
	s.ItineraryListShrink = &v
	return s
}

func (s *ApplyAddShrinkRequest) SetItineraryRule(v int32) *ApplyAddShrinkRequest {
	s.ItineraryRule = &v
	return s
}

func (s *ApplyAddShrinkRequest) SetItinerarySetListShrink(v string) *ApplyAddShrinkRequest {
	s.ItinerarySetListShrink = &v
	return s
}

func (s *ApplyAddShrinkRequest) SetLimitTraveler(v int32) *ApplyAddShrinkRequest {
	s.LimitTraveler = &v
	return s
}

func (s *ApplyAddShrinkRequest) SetMealBudget(v int64) *ApplyAddShrinkRequest {
	s.MealBudget = &v
	return s
}

func (s *ApplyAddShrinkRequest) SetPaymentDepartmentId(v string) *ApplyAddShrinkRequest {
	s.PaymentDepartmentId = &v
	return s
}

func (s *ApplyAddShrinkRequest) SetPaymentDepartmentName(v string) *ApplyAddShrinkRequest {
	s.PaymentDepartmentName = &v
	return s
}

func (s *ApplyAddShrinkRequest) SetStatus(v int32) *ApplyAddShrinkRequest {
	s.Status = &v
	return s
}

func (s *ApplyAddShrinkRequest) SetSubCorpId(v string) *ApplyAddShrinkRequest {
	s.SubCorpId = &v
	return s
}

func (s *ApplyAddShrinkRequest) SetThirdpartApplyId(v string) *ApplyAddShrinkRequest {
	s.ThirdpartApplyId = &v
	return s
}

func (s *ApplyAddShrinkRequest) SetThirdpartBusinessId(v string) *ApplyAddShrinkRequest {
	s.ThirdpartBusinessId = &v
	return s
}

func (s *ApplyAddShrinkRequest) SetThirdpartDepartId(v string) *ApplyAddShrinkRequest {
	s.ThirdpartDepartId = &v
	return s
}

func (s *ApplyAddShrinkRequest) SetTogetherBookRule(v int32) *ApplyAddShrinkRequest {
	s.TogetherBookRule = &v
	return s
}

func (s *ApplyAddShrinkRequest) SetTrainBudget(v int64) *ApplyAddShrinkRequest {
	s.TrainBudget = &v
	return s
}

func (s *ApplyAddShrinkRequest) SetTravelerListShrink(v string) *ApplyAddShrinkRequest {
	s.TravelerListShrink = &v
	return s
}

func (s *ApplyAddShrinkRequest) SetTravelerStandardShrink(v string) *ApplyAddShrinkRequest {
	s.TravelerStandardShrink = &v
	return s
}

func (s *ApplyAddShrinkRequest) SetTripCause(v string) *ApplyAddShrinkRequest {
	s.TripCause = &v
	return s
}

func (s *ApplyAddShrinkRequest) SetTripDay(v int32) *ApplyAddShrinkRequest {
	s.TripDay = &v
	return s
}

func (s *ApplyAddShrinkRequest) SetTripTitle(v string) *ApplyAddShrinkRequest {
	s.TripTitle = &v
	return s
}

func (s *ApplyAddShrinkRequest) SetType(v int32) *ApplyAddShrinkRequest {
	s.Type = &v
	return s
}

func (s *ApplyAddShrinkRequest) SetUnionNo(v string) *ApplyAddShrinkRequest {
	s.UnionNo = &v
	return s
}

func (s *ApplyAddShrinkRequest) SetUserId(v string) *ApplyAddShrinkRequest {
	s.UserId = &v
	return s
}

func (s *ApplyAddShrinkRequest) SetUserName(v string) *ApplyAddShrinkRequest {
	s.UserName = &v
	return s
}

func (s *ApplyAddShrinkRequest) SetVehicleBudget(v int64) *ApplyAddShrinkRequest {
	s.VehicleBudget = &v
	return s
}

func (s *ApplyAddShrinkRequest) Validate() error {
	return dara.Validate(s)
}
