// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyModifyShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBudget(v int64) *ApplyModifyShrinkRequest
	GetBudget() *int64
	SetBudgetMerge(v int32) *ApplyModifyShrinkRequest
	GetBudgetMerge() *int32
	SetCarRuleShrink(v string) *ApplyModifyShrinkRequest
	GetCarRuleShrink() *string
	SetCorpName(v string) *ApplyModifyShrinkRequest
	GetCorpName() *string
	SetDefaultStandardShrink(v string) *ApplyModifyShrinkRequest
	GetDefaultStandardShrink() *string
	SetDepartId(v string) *ApplyModifyShrinkRequest
	GetDepartId() *string
	SetDepartName(v string) *ApplyModifyShrinkRequest
	GetDepartName() *string
	SetExtendField(v string) *ApplyModifyShrinkRequest
	GetExtendField() *string
	SetExternalTravelerListShrink(v string) *ApplyModifyShrinkRequest
	GetExternalTravelerListShrink() *string
	SetExternalTravelerStandardShrink(v string) *ApplyModifyShrinkRequest
	GetExternalTravelerStandardShrink() *string
	SetFlightBudget(v int64) *ApplyModifyShrinkRequest
	GetFlightBudget() *int64
	SetHotelBudget(v int64) *ApplyModifyShrinkRequest
	GetHotelBudget() *int64
	SetHotelShareShrink(v string) *ApplyModifyShrinkRequest
	GetHotelShareShrink() *string
	SetIntlFlightBudget(v int64) *ApplyModifyShrinkRequest
	GetIntlFlightBudget() *int64
	SetIntlHotelBudget(v int64) *ApplyModifyShrinkRequest
	GetIntlHotelBudget() *int64
	SetItineraryListShrink(v string) *ApplyModifyShrinkRequest
	GetItineraryListShrink() *string
	SetItineraryRule(v int32) *ApplyModifyShrinkRequest
	GetItineraryRule() *int32
	SetItinerarySetListShrink(v string) *ApplyModifyShrinkRequest
	GetItinerarySetListShrink() *string
	SetLimitTraveler(v int32) *ApplyModifyShrinkRequest
	GetLimitTraveler() *int32
	SetMealBudget(v int64) *ApplyModifyShrinkRequest
	GetMealBudget() *int64
	SetPaymentDepartmentId(v string) *ApplyModifyShrinkRequest
	GetPaymentDepartmentId() *string
	SetPaymentDepartmentName(v string) *ApplyModifyShrinkRequest
	GetPaymentDepartmentName() *string
	SetStatus(v int32) *ApplyModifyShrinkRequest
	GetStatus() *int32
	SetSubCorpId(v string) *ApplyModifyShrinkRequest
	GetSubCorpId() *string
	SetThirdpartApplyId(v string) *ApplyModifyShrinkRequest
	GetThirdpartApplyId() *string
	SetThirdpartBusinessId(v string) *ApplyModifyShrinkRequest
	GetThirdpartBusinessId() *string
	SetThirdpartDepartId(v string) *ApplyModifyShrinkRequest
	GetThirdpartDepartId() *string
	SetTogetherBookRule(v int32) *ApplyModifyShrinkRequest
	GetTogetherBookRule() *int32
	SetTrainBudget(v int64) *ApplyModifyShrinkRequest
	GetTrainBudget() *int64
	SetTravelerListShrink(v string) *ApplyModifyShrinkRequest
	GetTravelerListShrink() *string
	SetTravelerStandardShrink(v string) *ApplyModifyShrinkRequest
	GetTravelerStandardShrink() *string
	SetTripCause(v string) *ApplyModifyShrinkRequest
	GetTripCause() *string
	SetTripDay(v int32) *ApplyModifyShrinkRequest
	GetTripDay() *int32
	SetTripTitle(v string) *ApplyModifyShrinkRequest
	GetTripTitle() *string
	SetUnionNo(v string) *ApplyModifyShrinkRequest
	GetUnionNo() *string
	SetUserId(v string) *ApplyModifyShrinkRequest
	GetUserId() *string
	SetUserName(v string) *ApplyModifyShrinkRequest
	GetUserName() *string
	SetVehicleBudget(v int64) *ApplyModifyShrinkRequest
	GetVehicleBudget() *int64
}

type ApplyModifyShrinkRequest struct {
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
	// 001
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
	HotelBudget      *int64  `json:"hotel_budget,omitempty" xml:"hotel_budget,omitempty"`
	HotelShareShrink *string `json:"hotel_share,omitempty" xml:"hotel_share,omitempty"`
	IntlFlightBudget *int64  `json:"intl_flight_budget,omitempty" xml:"intl_flight_budget,omitempty"`
	IntlHotelBudget  *int64  `json:"intl_hotel_budget,omitempty" xml:"intl_hotel_budget,omitempty"`
	// example:
	//
	// 0
	ItineraryListShrink *string `json:"itinerary_list,omitempty" xml:"itinerary_list,omitempty"`
	// example:
	//
	// 0
	ItineraryRule          *int32  `json:"itinerary_rule,omitempty" xml:"itinerary_rule,omitempty"`
	ItinerarySetListShrink *string `json:"itinerary_set_list,omitempty" xml:"itinerary_set_list,omitempty"`
	// example:
	//
	// 0
	LimitTraveler         *int32  `json:"limit_traveler,omitempty" xml:"limit_traveler,omitempty"`
	MealBudget            *int64  `json:"meal_budget,omitempty" xml:"meal_budget,omitempty"`
	PaymentDepartmentId   *string `json:"payment_department_id,omitempty" xml:"payment_department_id,omitempty"`
	PaymentDepartmentName *string `json:"payment_department_name,omitempty" xml:"payment_department_name,omitempty"`
	// example:
	//
	// 1
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// 子企业Id
	//
	// example:
	//
	// btrip123
	SubCorpId *string `json:"sub_corp_id,omitempty" xml:"sub_corp_id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0001A1100000007EX08O
	ThirdpartApplyId *string `json:"thirdpart_apply_id,omitempty" xml:"thirdpart_apply_id,omitempty"`
	// example:
	//
	// 202201413141
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
	// union51415
	UnionNo *string `json:"union_no,omitempty" xml:"union_no,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// thirdpart12138
	UserId   *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	UserName *string `json:"user_name,omitempty" xml:"user_name,omitempty"`
	// example:
	//
	// 1000
	VehicleBudget *int64 `json:"vehicle_budget,omitempty" xml:"vehicle_budget,omitempty"`
}

func (s ApplyModifyShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ApplyModifyShrinkRequest) GoString() string {
	return s.String()
}

func (s *ApplyModifyShrinkRequest) GetBudget() *int64 {
	return s.Budget
}

func (s *ApplyModifyShrinkRequest) GetBudgetMerge() *int32 {
	return s.BudgetMerge
}

func (s *ApplyModifyShrinkRequest) GetCarRuleShrink() *string {
	return s.CarRuleShrink
}

func (s *ApplyModifyShrinkRequest) GetCorpName() *string {
	return s.CorpName
}

func (s *ApplyModifyShrinkRequest) GetDefaultStandardShrink() *string {
	return s.DefaultStandardShrink
}

func (s *ApplyModifyShrinkRequest) GetDepartId() *string {
	return s.DepartId
}

func (s *ApplyModifyShrinkRequest) GetDepartName() *string {
	return s.DepartName
}

func (s *ApplyModifyShrinkRequest) GetExtendField() *string {
	return s.ExtendField
}

func (s *ApplyModifyShrinkRequest) GetExternalTravelerListShrink() *string {
	return s.ExternalTravelerListShrink
}

func (s *ApplyModifyShrinkRequest) GetExternalTravelerStandardShrink() *string {
	return s.ExternalTravelerStandardShrink
}

func (s *ApplyModifyShrinkRequest) GetFlightBudget() *int64 {
	return s.FlightBudget
}

func (s *ApplyModifyShrinkRequest) GetHotelBudget() *int64 {
	return s.HotelBudget
}

func (s *ApplyModifyShrinkRequest) GetHotelShareShrink() *string {
	return s.HotelShareShrink
}

func (s *ApplyModifyShrinkRequest) GetIntlFlightBudget() *int64 {
	return s.IntlFlightBudget
}

func (s *ApplyModifyShrinkRequest) GetIntlHotelBudget() *int64 {
	return s.IntlHotelBudget
}

func (s *ApplyModifyShrinkRequest) GetItineraryListShrink() *string {
	return s.ItineraryListShrink
}

func (s *ApplyModifyShrinkRequest) GetItineraryRule() *int32 {
	return s.ItineraryRule
}

func (s *ApplyModifyShrinkRequest) GetItinerarySetListShrink() *string {
	return s.ItinerarySetListShrink
}

func (s *ApplyModifyShrinkRequest) GetLimitTraveler() *int32 {
	return s.LimitTraveler
}

func (s *ApplyModifyShrinkRequest) GetMealBudget() *int64 {
	return s.MealBudget
}

func (s *ApplyModifyShrinkRequest) GetPaymentDepartmentId() *string {
	return s.PaymentDepartmentId
}

func (s *ApplyModifyShrinkRequest) GetPaymentDepartmentName() *string {
	return s.PaymentDepartmentName
}

func (s *ApplyModifyShrinkRequest) GetStatus() *int32 {
	return s.Status
}

func (s *ApplyModifyShrinkRequest) GetSubCorpId() *string {
	return s.SubCorpId
}

func (s *ApplyModifyShrinkRequest) GetThirdpartApplyId() *string {
	return s.ThirdpartApplyId
}

func (s *ApplyModifyShrinkRequest) GetThirdpartBusinessId() *string {
	return s.ThirdpartBusinessId
}

func (s *ApplyModifyShrinkRequest) GetThirdpartDepartId() *string {
	return s.ThirdpartDepartId
}

func (s *ApplyModifyShrinkRequest) GetTogetherBookRule() *int32 {
	return s.TogetherBookRule
}

func (s *ApplyModifyShrinkRequest) GetTrainBudget() *int64 {
	return s.TrainBudget
}

func (s *ApplyModifyShrinkRequest) GetTravelerListShrink() *string {
	return s.TravelerListShrink
}

func (s *ApplyModifyShrinkRequest) GetTravelerStandardShrink() *string {
	return s.TravelerStandardShrink
}

func (s *ApplyModifyShrinkRequest) GetTripCause() *string {
	return s.TripCause
}

func (s *ApplyModifyShrinkRequest) GetTripDay() *int32 {
	return s.TripDay
}

func (s *ApplyModifyShrinkRequest) GetTripTitle() *string {
	return s.TripTitle
}

func (s *ApplyModifyShrinkRequest) GetUnionNo() *string {
	return s.UnionNo
}

func (s *ApplyModifyShrinkRequest) GetUserId() *string {
	return s.UserId
}

func (s *ApplyModifyShrinkRequest) GetUserName() *string {
	return s.UserName
}

func (s *ApplyModifyShrinkRequest) GetVehicleBudget() *int64 {
	return s.VehicleBudget
}

func (s *ApplyModifyShrinkRequest) SetBudget(v int64) *ApplyModifyShrinkRequest {
	s.Budget = &v
	return s
}

func (s *ApplyModifyShrinkRequest) SetBudgetMerge(v int32) *ApplyModifyShrinkRequest {
	s.BudgetMerge = &v
	return s
}

func (s *ApplyModifyShrinkRequest) SetCarRuleShrink(v string) *ApplyModifyShrinkRequest {
	s.CarRuleShrink = &v
	return s
}

func (s *ApplyModifyShrinkRequest) SetCorpName(v string) *ApplyModifyShrinkRequest {
	s.CorpName = &v
	return s
}

func (s *ApplyModifyShrinkRequest) SetDefaultStandardShrink(v string) *ApplyModifyShrinkRequest {
	s.DefaultStandardShrink = &v
	return s
}

func (s *ApplyModifyShrinkRequest) SetDepartId(v string) *ApplyModifyShrinkRequest {
	s.DepartId = &v
	return s
}

func (s *ApplyModifyShrinkRequest) SetDepartName(v string) *ApplyModifyShrinkRequest {
	s.DepartName = &v
	return s
}

func (s *ApplyModifyShrinkRequest) SetExtendField(v string) *ApplyModifyShrinkRequest {
	s.ExtendField = &v
	return s
}

func (s *ApplyModifyShrinkRequest) SetExternalTravelerListShrink(v string) *ApplyModifyShrinkRequest {
	s.ExternalTravelerListShrink = &v
	return s
}

func (s *ApplyModifyShrinkRequest) SetExternalTravelerStandardShrink(v string) *ApplyModifyShrinkRequest {
	s.ExternalTravelerStandardShrink = &v
	return s
}

func (s *ApplyModifyShrinkRequest) SetFlightBudget(v int64) *ApplyModifyShrinkRequest {
	s.FlightBudget = &v
	return s
}

func (s *ApplyModifyShrinkRequest) SetHotelBudget(v int64) *ApplyModifyShrinkRequest {
	s.HotelBudget = &v
	return s
}

func (s *ApplyModifyShrinkRequest) SetHotelShareShrink(v string) *ApplyModifyShrinkRequest {
	s.HotelShareShrink = &v
	return s
}

func (s *ApplyModifyShrinkRequest) SetIntlFlightBudget(v int64) *ApplyModifyShrinkRequest {
	s.IntlFlightBudget = &v
	return s
}

func (s *ApplyModifyShrinkRequest) SetIntlHotelBudget(v int64) *ApplyModifyShrinkRequest {
	s.IntlHotelBudget = &v
	return s
}

func (s *ApplyModifyShrinkRequest) SetItineraryListShrink(v string) *ApplyModifyShrinkRequest {
	s.ItineraryListShrink = &v
	return s
}

func (s *ApplyModifyShrinkRequest) SetItineraryRule(v int32) *ApplyModifyShrinkRequest {
	s.ItineraryRule = &v
	return s
}

func (s *ApplyModifyShrinkRequest) SetItinerarySetListShrink(v string) *ApplyModifyShrinkRequest {
	s.ItinerarySetListShrink = &v
	return s
}

func (s *ApplyModifyShrinkRequest) SetLimitTraveler(v int32) *ApplyModifyShrinkRequest {
	s.LimitTraveler = &v
	return s
}

func (s *ApplyModifyShrinkRequest) SetMealBudget(v int64) *ApplyModifyShrinkRequest {
	s.MealBudget = &v
	return s
}

func (s *ApplyModifyShrinkRequest) SetPaymentDepartmentId(v string) *ApplyModifyShrinkRequest {
	s.PaymentDepartmentId = &v
	return s
}

func (s *ApplyModifyShrinkRequest) SetPaymentDepartmentName(v string) *ApplyModifyShrinkRequest {
	s.PaymentDepartmentName = &v
	return s
}

func (s *ApplyModifyShrinkRequest) SetStatus(v int32) *ApplyModifyShrinkRequest {
	s.Status = &v
	return s
}

func (s *ApplyModifyShrinkRequest) SetSubCorpId(v string) *ApplyModifyShrinkRequest {
	s.SubCorpId = &v
	return s
}

func (s *ApplyModifyShrinkRequest) SetThirdpartApplyId(v string) *ApplyModifyShrinkRequest {
	s.ThirdpartApplyId = &v
	return s
}

func (s *ApplyModifyShrinkRequest) SetThirdpartBusinessId(v string) *ApplyModifyShrinkRequest {
	s.ThirdpartBusinessId = &v
	return s
}

func (s *ApplyModifyShrinkRequest) SetThirdpartDepartId(v string) *ApplyModifyShrinkRequest {
	s.ThirdpartDepartId = &v
	return s
}

func (s *ApplyModifyShrinkRequest) SetTogetherBookRule(v int32) *ApplyModifyShrinkRequest {
	s.TogetherBookRule = &v
	return s
}

func (s *ApplyModifyShrinkRequest) SetTrainBudget(v int64) *ApplyModifyShrinkRequest {
	s.TrainBudget = &v
	return s
}

func (s *ApplyModifyShrinkRequest) SetTravelerListShrink(v string) *ApplyModifyShrinkRequest {
	s.TravelerListShrink = &v
	return s
}

func (s *ApplyModifyShrinkRequest) SetTravelerStandardShrink(v string) *ApplyModifyShrinkRequest {
	s.TravelerStandardShrink = &v
	return s
}

func (s *ApplyModifyShrinkRequest) SetTripCause(v string) *ApplyModifyShrinkRequest {
	s.TripCause = &v
	return s
}

func (s *ApplyModifyShrinkRequest) SetTripDay(v int32) *ApplyModifyShrinkRequest {
	s.TripDay = &v
	return s
}

func (s *ApplyModifyShrinkRequest) SetTripTitle(v string) *ApplyModifyShrinkRequest {
	s.TripTitle = &v
	return s
}

func (s *ApplyModifyShrinkRequest) SetUnionNo(v string) *ApplyModifyShrinkRequest {
	s.UnionNo = &v
	return s
}

func (s *ApplyModifyShrinkRequest) SetUserId(v string) *ApplyModifyShrinkRequest {
	s.UserId = &v
	return s
}

func (s *ApplyModifyShrinkRequest) SetUserName(v string) *ApplyModifyShrinkRequest {
	s.UserName = &v
	return s
}

func (s *ApplyModifyShrinkRequest) SetVehicleBudget(v int64) *ApplyModifyShrinkRequest {
	s.VehicleBudget = &v
	return s
}

func (s *ApplyModifyShrinkRequest) Validate() error {
	return dara.Validate(s)
}
