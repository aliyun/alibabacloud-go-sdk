// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyModifyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBudget(v int64) *ApplyModifyRequest
	GetBudget() *int64
	SetBudgetMerge(v int32) *ApplyModifyRequest
	GetBudgetMerge() *int32
	SetCarRule(v *ApplyModifyRequestCarRule) *ApplyModifyRequest
	GetCarRule() *ApplyModifyRequestCarRule
	SetCorpName(v string) *ApplyModifyRequest
	GetCorpName() *string
	SetDefaultStandard(v *ApplyModifyRequestDefaultStandard) *ApplyModifyRequest
	GetDefaultStandard() *ApplyModifyRequestDefaultStandard
	SetDepartId(v string) *ApplyModifyRequest
	GetDepartId() *string
	SetDepartName(v string) *ApplyModifyRequest
	GetDepartName() *string
	SetExtendField(v string) *ApplyModifyRequest
	GetExtendField() *string
	SetExternalTravelerList(v []*ApplyModifyRequestExternalTravelerList) *ApplyModifyRequest
	GetExternalTravelerList() []*ApplyModifyRequestExternalTravelerList
	SetExternalTravelerStandard(v *ApplyModifyRequestExternalTravelerStandard) *ApplyModifyRequest
	GetExternalTravelerStandard() *ApplyModifyRequestExternalTravelerStandard
	SetFlightBudget(v int64) *ApplyModifyRequest
	GetFlightBudget() *int64
	SetHotelBudget(v int64) *ApplyModifyRequest
	GetHotelBudget() *int64
	SetHotelShare(v *ApplyModifyRequestHotelShare) *ApplyModifyRequest
	GetHotelShare() *ApplyModifyRequestHotelShare
	SetIntlFlightBudget(v int64) *ApplyModifyRequest
	GetIntlFlightBudget() *int64
	SetIntlHotelBudget(v int64) *ApplyModifyRequest
	GetIntlHotelBudget() *int64
	SetItineraryList(v []*ApplyModifyRequestItineraryList) *ApplyModifyRequest
	GetItineraryList() []*ApplyModifyRequestItineraryList
	SetItineraryRule(v int32) *ApplyModifyRequest
	GetItineraryRule() *int32
	SetItinerarySetList(v []*ApplyModifyRequestItinerarySetList) *ApplyModifyRequest
	GetItinerarySetList() []*ApplyModifyRequestItinerarySetList
	SetLimitTraveler(v int32) *ApplyModifyRequest
	GetLimitTraveler() *int32
	SetMealBudget(v int64) *ApplyModifyRequest
	GetMealBudget() *int64
	SetPaymentDepartmentId(v string) *ApplyModifyRequest
	GetPaymentDepartmentId() *string
	SetPaymentDepartmentName(v string) *ApplyModifyRequest
	GetPaymentDepartmentName() *string
	SetStatus(v int32) *ApplyModifyRequest
	GetStatus() *int32
	SetSubCorpId(v string) *ApplyModifyRequest
	GetSubCorpId() *string
	SetThirdpartApplyId(v string) *ApplyModifyRequest
	GetThirdpartApplyId() *string
	SetThirdpartBusinessId(v string) *ApplyModifyRequest
	GetThirdpartBusinessId() *string
	SetThirdpartDepartId(v string) *ApplyModifyRequest
	GetThirdpartDepartId() *string
	SetTogetherBookRule(v int32) *ApplyModifyRequest
	GetTogetherBookRule() *int32
	SetTrainBudget(v int64) *ApplyModifyRequest
	GetTrainBudget() *int64
	SetTravelerList(v []*ApplyModifyRequestTravelerList) *ApplyModifyRequest
	GetTravelerList() []*ApplyModifyRequestTravelerList
	SetTravelerStandard(v []*ApplyModifyRequestTravelerStandard) *ApplyModifyRequest
	GetTravelerStandard() []*ApplyModifyRequestTravelerStandard
	SetTripCause(v string) *ApplyModifyRequest
	GetTripCause() *string
	SetTripDay(v int32) *ApplyModifyRequest
	GetTripDay() *int32
	SetTripTitle(v string) *ApplyModifyRequest
	GetTripTitle() *string
	SetUnionNo(v string) *ApplyModifyRequest
	GetUnionNo() *string
	SetUserId(v string) *ApplyModifyRequest
	GetUserId() *string
	SetUserName(v string) *ApplyModifyRequest
	GetUserName() *string
	SetVehicleBudget(v int64) *ApplyModifyRequest
	GetVehicleBudget() *int64
}

type ApplyModifyRequest struct {
	Budget                   *int64                                      `json:"budget,omitempty" xml:"budget,omitempty"`
	BudgetMerge              *int32                                      `json:"budget_merge,omitempty" xml:"budget_merge,omitempty"`
	CarRule                  *ApplyModifyRequestCarRule                  `json:"car_rule,omitempty" xml:"car_rule,omitempty" type:"Struct"`
	CorpName                 *string                                     `json:"corp_name,omitempty" xml:"corp_name,omitempty"`
	DefaultStandard          *ApplyModifyRequestDefaultStandard          `json:"default_standard,omitempty" xml:"default_standard,omitempty" type:"Struct"`
	DepartId                 *string                                     `json:"depart_id,omitempty" xml:"depart_id,omitempty"`
	DepartName               *string                                     `json:"depart_name,omitempty" xml:"depart_name,omitempty"`
	ExtendField              *string                                     `json:"extend_field,omitempty" xml:"extend_field,omitempty"`
	ExternalTravelerList     []*ApplyModifyRequestExternalTravelerList   `json:"external_traveler_list,omitempty" xml:"external_traveler_list,omitempty" type:"Repeated"`
	ExternalTravelerStandard *ApplyModifyRequestExternalTravelerStandard `json:"external_traveler_standard,omitempty" xml:"external_traveler_standard,omitempty" type:"Struct"`
	FlightBudget             *int64                                      `json:"flight_budget,omitempty" xml:"flight_budget,omitempty"`
	HotelBudget              *int64                                      `json:"hotel_budget,omitempty" xml:"hotel_budget,omitempty"`
	HotelShare               *ApplyModifyRequestHotelShare               `json:"hotel_share,omitempty" xml:"hotel_share,omitempty" type:"Struct"`
	IntlFlightBudget         *int64                                      `json:"intl_flight_budget,omitempty" xml:"intl_flight_budget,omitempty"`
	IntlHotelBudget          *int64                                      `json:"intl_hotel_budget,omitempty" xml:"intl_hotel_budget,omitempty"`
	ItineraryList            []*ApplyModifyRequestItineraryList          `json:"itinerary_list,omitempty" xml:"itinerary_list,omitempty" type:"Repeated"`
	ItineraryRule            *int32                                      `json:"itinerary_rule,omitempty" xml:"itinerary_rule,omitempty"`
	ItinerarySetList         []*ApplyModifyRequestItinerarySetList       `json:"itinerary_set_list,omitempty" xml:"itinerary_set_list,omitempty" type:"Repeated"`
	LimitTraveler            *int32                                      `json:"limit_traveler,omitempty" xml:"limit_traveler,omitempty"`
	MealBudget               *int64                                      `json:"meal_budget,omitempty" xml:"meal_budget,omitempty"`
	PaymentDepartmentId      *string                                     `json:"payment_department_id,omitempty" xml:"payment_department_id,omitempty"`
	PaymentDepartmentName    *string                                     `json:"payment_department_name,omitempty" xml:"payment_department_name,omitempty"`
	Status                   *int32                                      `json:"status,omitempty" xml:"status,omitempty"`
	SubCorpId                *string                                     `json:"sub_corp_id,omitempty" xml:"sub_corp_id,omitempty"`
	// This parameter is required.
	ThirdpartApplyId    *string                               `json:"thirdpart_apply_id,omitempty" xml:"thirdpart_apply_id,omitempty"`
	ThirdpartBusinessId *string                               `json:"thirdpart_business_id,omitempty" xml:"thirdpart_business_id,omitempty"`
	ThirdpartDepartId   *string                               `json:"thirdpart_depart_id,omitempty" xml:"thirdpart_depart_id,omitempty"`
	TogetherBookRule    *int32                                `json:"together_book_rule,omitempty" xml:"together_book_rule,omitempty"`
	TrainBudget         *int64                                `json:"train_budget,omitempty" xml:"train_budget,omitempty"`
	TravelerList        []*ApplyModifyRequestTravelerList     `json:"traveler_list,omitempty" xml:"traveler_list,omitempty" type:"Repeated"`
	TravelerStandard    []*ApplyModifyRequestTravelerStandard `json:"traveler_standard,omitempty" xml:"traveler_standard,omitempty" type:"Repeated"`
	// This parameter is required.
	TripCause *string `json:"trip_cause,omitempty" xml:"trip_cause,omitempty"`
	TripDay   *int32  `json:"trip_day,omitempty" xml:"trip_day,omitempty"`
	// This parameter is required.
	TripTitle *string `json:"trip_title,omitempty" xml:"trip_title,omitempty"`
	UnionNo   *string `json:"union_no,omitempty" xml:"union_no,omitempty"`
	// This parameter is required.
	UserId        *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	UserName      *string `json:"user_name,omitempty" xml:"user_name,omitempty"`
	VehicleBudget *int64  `json:"vehicle_budget,omitempty" xml:"vehicle_budget,omitempty"`
}

func (s ApplyModifyRequest) String() string {
	return dara.Prettify(s)
}

func (s ApplyModifyRequest) GoString() string {
	return s.String()
}

func (s *ApplyModifyRequest) GetBudget() *int64 {
	return s.Budget
}

func (s *ApplyModifyRequest) GetBudgetMerge() *int32 {
	return s.BudgetMerge
}

func (s *ApplyModifyRequest) GetCarRule() *ApplyModifyRequestCarRule {
	return s.CarRule
}

func (s *ApplyModifyRequest) GetCorpName() *string {
	return s.CorpName
}

func (s *ApplyModifyRequest) GetDefaultStandard() *ApplyModifyRequestDefaultStandard {
	return s.DefaultStandard
}

func (s *ApplyModifyRequest) GetDepartId() *string {
	return s.DepartId
}

func (s *ApplyModifyRequest) GetDepartName() *string {
	return s.DepartName
}

func (s *ApplyModifyRequest) GetExtendField() *string {
	return s.ExtendField
}

func (s *ApplyModifyRequest) GetExternalTravelerList() []*ApplyModifyRequestExternalTravelerList {
	return s.ExternalTravelerList
}

func (s *ApplyModifyRequest) GetExternalTravelerStandard() *ApplyModifyRequestExternalTravelerStandard {
	return s.ExternalTravelerStandard
}

func (s *ApplyModifyRequest) GetFlightBudget() *int64 {
	return s.FlightBudget
}

func (s *ApplyModifyRequest) GetHotelBudget() *int64 {
	return s.HotelBudget
}

func (s *ApplyModifyRequest) GetHotelShare() *ApplyModifyRequestHotelShare {
	return s.HotelShare
}

func (s *ApplyModifyRequest) GetIntlFlightBudget() *int64 {
	return s.IntlFlightBudget
}

func (s *ApplyModifyRequest) GetIntlHotelBudget() *int64 {
	return s.IntlHotelBudget
}

func (s *ApplyModifyRequest) GetItineraryList() []*ApplyModifyRequestItineraryList {
	return s.ItineraryList
}

func (s *ApplyModifyRequest) GetItineraryRule() *int32 {
	return s.ItineraryRule
}

func (s *ApplyModifyRequest) GetItinerarySetList() []*ApplyModifyRequestItinerarySetList {
	return s.ItinerarySetList
}

func (s *ApplyModifyRequest) GetLimitTraveler() *int32 {
	return s.LimitTraveler
}

func (s *ApplyModifyRequest) GetMealBudget() *int64 {
	return s.MealBudget
}

func (s *ApplyModifyRequest) GetPaymentDepartmentId() *string {
	return s.PaymentDepartmentId
}

func (s *ApplyModifyRequest) GetPaymentDepartmentName() *string {
	return s.PaymentDepartmentName
}

func (s *ApplyModifyRequest) GetStatus() *int32 {
	return s.Status
}

func (s *ApplyModifyRequest) GetSubCorpId() *string {
	return s.SubCorpId
}

func (s *ApplyModifyRequest) GetThirdpartApplyId() *string {
	return s.ThirdpartApplyId
}

func (s *ApplyModifyRequest) GetThirdpartBusinessId() *string {
	return s.ThirdpartBusinessId
}

func (s *ApplyModifyRequest) GetThirdpartDepartId() *string {
	return s.ThirdpartDepartId
}

func (s *ApplyModifyRequest) GetTogetherBookRule() *int32 {
	return s.TogetherBookRule
}

func (s *ApplyModifyRequest) GetTrainBudget() *int64 {
	return s.TrainBudget
}

func (s *ApplyModifyRequest) GetTravelerList() []*ApplyModifyRequestTravelerList {
	return s.TravelerList
}

func (s *ApplyModifyRequest) GetTravelerStandard() []*ApplyModifyRequestTravelerStandard {
	return s.TravelerStandard
}

func (s *ApplyModifyRequest) GetTripCause() *string {
	return s.TripCause
}

func (s *ApplyModifyRequest) GetTripDay() *int32 {
	return s.TripDay
}

func (s *ApplyModifyRequest) GetTripTitle() *string {
	return s.TripTitle
}

func (s *ApplyModifyRequest) GetUnionNo() *string {
	return s.UnionNo
}

func (s *ApplyModifyRequest) GetUserId() *string {
	return s.UserId
}

func (s *ApplyModifyRequest) GetUserName() *string {
	return s.UserName
}

func (s *ApplyModifyRequest) GetVehicleBudget() *int64 {
	return s.VehicleBudget
}

func (s *ApplyModifyRequest) SetBudget(v int64) *ApplyModifyRequest {
	s.Budget = &v
	return s
}

func (s *ApplyModifyRequest) SetBudgetMerge(v int32) *ApplyModifyRequest {
	s.BudgetMerge = &v
	return s
}

func (s *ApplyModifyRequest) SetCarRule(v *ApplyModifyRequestCarRule) *ApplyModifyRequest {
	s.CarRule = v
	return s
}

func (s *ApplyModifyRequest) SetCorpName(v string) *ApplyModifyRequest {
	s.CorpName = &v
	return s
}

func (s *ApplyModifyRequest) SetDefaultStandard(v *ApplyModifyRequestDefaultStandard) *ApplyModifyRequest {
	s.DefaultStandard = v
	return s
}

func (s *ApplyModifyRequest) SetDepartId(v string) *ApplyModifyRequest {
	s.DepartId = &v
	return s
}

func (s *ApplyModifyRequest) SetDepartName(v string) *ApplyModifyRequest {
	s.DepartName = &v
	return s
}

func (s *ApplyModifyRequest) SetExtendField(v string) *ApplyModifyRequest {
	s.ExtendField = &v
	return s
}

func (s *ApplyModifyRequest) SetExternalTravelerList(v []*ApplyModifyRequestExternalTravelerList) *ApplyModifyRequest {
	s.ExternalTravelerList = v
	return s
}

func (s *ApplyModifyRequest) SetExternalTravelerStandard(v *ApplyModifyRequestExternalTravelerStandard) *ApplyModifyRequest {
	s.ExternalTravelerStandard = v
	return s
}

func (s *ApplyModifyRequest) SetFlightBudget(v int64) *ApplyModifyRequest {
	s.FlightBudget = &v
	return s
}

func (s *ApplyModifyRequest) SetHotelBudget(v int64) *ApplyModifyRequest {
	s.HotelBudget = &v
	return s
}

func (s *ApplyModifyRequest) SetHotelShare(v *ApplyModifyRequestHotelShare) *ApplyModifyRequest {
	s.HotelShare = v
	return s
}

func (s *ApplyModifyRequest) SetIntlFlightBudget(v int64) *ApplyModifyRequest {
	s.IntlFlightBudget = &v
	return s
}

func (s *ApplyModifyRequest) SetIntlHotelBudget(v int64) *ApplyModifyRequest {
	s.IntlHotelBudget = &v
	return s
}

func (s *ApplyModifyRequest) SetItineraryList(v []*ApplyModifyRequestItineraryList) *ApplyModifyRequest {
	s.ItineraryList = v
	return s
}

func (s *ApplyModifyRequest) SetItineraryRule(v int32) *ApplyModifyRequest {
	s.ItineraryRule = &v
	return s
}

func (s *ApplyModifyRequest) SetItinerarySetList(v []*ApplyModifyRequestItinerarySetList) *ApplyModifyRequest {
	s.ItinerarySetList = v
	return s
}

func (s *ApplyModifyRequest) SetLimitTraveler(v int32) *ApplyModifyRequest {
	s.LimitTraveler = &v
	return s
}

func (s *ApplyModifyRequest) SetMealBudget(v int64) *ApplyModifyRequest {
	s.MealBudget = &v
	return s
}

func (s *ApplyModifyRequest) SetPaymentDepartmentId(v string) *ApplyModifyRequest {
	s.PaymentDepartmentId = &v
	return s
}

func (s *ApplyModifyRequest) SetPaymentDepartmentName(v string) *ApplyModifyRequest {
	s.PaymentDepartmentName = &v
	return s
}

func (s *ApplyModifyRequest) SetStatus(v int32) *ApplyModifyRequest {
	s.Status = &v
	return s
}

func (s *ApplyModifyRequest) SetSubCorpId(v string) *ApplyModifyRequest {
	s.SubCorpId = &v
	return s
}

func (s *ApplyModifyRequest) SetThirdpartApplyId(v string) *ApplyModifyRequest {
	s.ThirdpartApplyId = &v
	return s
}

func (s *ApplyModifyRequest) SetThirdpartBusinessId(v string) *ApplyModifyRequest {
	s.ThirdpartBusinessId = &v
	return s
}

func (s *ApplyModifyRequest) SetThirdpartDepartId(v string) *ApplyModifyRequest {
	s.ThirdpartDepartId = &v
	return s
}

func (s *ApplyModifyRequest) SetTogetherBookRule(v int32) *ApplyModifyRequest {
	s.TogetherBookRule = &v
	return s
}

func (s *ApplyModifyRequest) SetTrainBudget(v int64) *ApplyModifyRequest {
	s.TrainBudget = &v
	return s
}

func (s *ApplyModifyRequest) SetTravelerList(v []*ApplyModifyRequestTravelerList) *ApplyModifyRequest {
	s.TravelerList = v
	return s
}

func (s *ApplyModifyRequest) SetTravelerStandard(v []*ApplyModifyRequestTravelerStandard) *ApplyModifyRequest {
	s.TravelerStandard = v
	return s
}

func (s *ApplyModifyRequest) SetTripCause(v string) *ApplyModifyRequest {
	s.TripCause = &v
	return s
}

func (s *ApplyModifyRequest) SetTripDay(v int32) *ApplyModifyRequest {
	s.TripDay = &v
	return s
}

func (s *ApplyModifyRequest) SetTripTitle(v string) *ApplyModifyRequest {
	s.TripTitle = &v
	return s
}

func (s *ApplyModifyRequest) SetUnionNo(v string) *ApplyModifyRequest {
	s.UnionNo = &v
	return s
}

func (s *ApplyModifyRequest) SetUserId(v string) *ApplyModifyRequest {
	s.UserId = &v
	return s
}

func (s *ApplyModifyRequest) SetUserName(v string) *ApplyModifyRequest {
	s.UserName = &v
	return s
}

func (s *ApplyModifyRequest) SetVehicleBudget(v int64) *ApplyModifyRequest {
	s.VehicleBudget = &v
	return s
}

func (s *ApplyModifyRequest) Validate() error {
	if s.CarRule != nil {
		if err := s.CarRule.Validate(); err != nil {
			return err
		}
	}
	if s.DefaultStandard != nil {
		if err := s.DefaultStandard.Validate(); err != nil {
			return err
		}
	}
	if s.ExternalTravelerList != nil {
		for _, item := range s.ExternalTravelerList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ExternalTravelerStandard != nil {
		if err := s.ExternalTravelerStandard.Validate(); err != nil {
			return err
		}
	}
	if s.HotelShare != nil {
		if err := s.HotelShare.Validate(); err != nil {
			return err
		}
	}
	if s.ItineraryList != nil {
		for _, item := range s.ItineraryList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ItinerarySetList != nil {
		for _, item := range s.ItinerarySetList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.TravelerList != nil {
		for _, item := range s.TravelerList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.TravelerStandard != nil {
		for _, item := range s.TravelerStandard {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ApplyModifyRequestCarRule struct {
	ScenarioTemplateId   *string `json:"scenario_template_id,omitempty" xml:"scenario_template_id,omitempty"`
	ScenarioTemplateName *string `json:"scenario_template_name,omitempty" xml:"scenario_template_name,omitempty"`
}

func (s ApplyModifyRequestCarRule) String() string {
	return dara.Prettify(s)
}

func (s ApplyModifyRequestCarRule) GoString() string {
	return s.String()
}

func (s *ApplyModifyRequestCarRule) GetScenarioTemplateId() *string {
	return s.ScenarioTemplateId
}

func (s *ApplyModifyRequestCarRule) GetScenarioTemplateName() *string {
	return s.ScenarioTemplateName
}

func (s *ApplyModifyRequestCarRule) SetScenarioTemplateId(v string) *ApplyModifyRequestCarRule {
	s.ScenarioTemplateId = &v
	return s
}

func (s *ApplyModifyRequestCarRule) SetScenarioTemplateName(v string) *ApplyModifyRequestCarRule {
	s.ScenarioTemplateName = &v
	return s
}

func (s *ApplyModifyRequestCarRule) Validate() error {
	return dara.Validate(s)
}

type ApplyModifyRequestDefaultStandard struct {
	BusinessDiscount          *int32                                             `json:"business_discount,omitempty" xml:"business_discount,omitempty"`
	EconomyDiscount           *int32                                             `json:"economy_discount,omitempty" xml:"economy_discount,omitempty"`
	FirstDiscount             *int32                                             `json:"first_discount,omitempty" xml:"first_discount,omitempty"`
	FlightCabins              *string                                            `json:"flight_cabins,omitempty" xml:"flight_cabins,omitempty"`
	FlightIntlRuleCode        *int64                                             `json:"flight_intl_rule_code,omitempty" xml:"flight_intl_rule_code,omitempty"`
	FlightRuleCode            *int64                                             `json:"flight_rule_code,omitempty" xml:"flight_rule_code,omitempty"`
	HotelCitys                []*ApplyModifyRequestDefaultStandardHotelCitys     `json:"hotel_citys,omitempty" xml:"hotel_citys,omitempty" type:"Repeated"`
	HotelIntlCitys            []*ApplyModifyRequestDefaultStandardHotelIntlCitys `json:"hotel_intl_citys,omitempty" xml:"hotel_intl_citys,omitempty" type:"Repeated"`
	HotelIntlRuleCode         *int64                                             `json:"hotel_intl_rule_code,omitempty" xml:"hotel_intl_rule_code,omitempty"`
	HotelRuleCode             *int64                                             `json:"hotel_rule_code,omitempty" xml:"hotel_rule_code,omitempty"`
	InternationalFlightCabins *string                                            `json:"international_flight_cabins,omitempty" xml:"international_flight_cabins,omitempty"`
	PremiumEconomyDiscount    *int32                                             `json:"premium_economy_discount,omitempty" xml:"premium_economy_discount,omitempty"`
	ReserveType               *int32                                             `json:"reserve_type,omitempty" xml:"reserve_type,omitempty"`
	TrainRuleCode             *int64                                             `json:"train_rule_code,omitempty" xml:"train_rule_code,omitempty"`
	TrainSeats                *string                                            `json:"train_seats,omitempty" xml:"train_seats,omitempty"`
}

func (s ApplyModifyRequestDefaultStandard) String() string {
	return dara.Prettify(s)
}

func (s ApplyModifyRequestDefaultStandard) GoString() string {
	return s.String()
}

func (s *ApplyModifyRequestDefaultStandard) GetBusinessDiscount() *int32 {
	return s.BusinessDiscount
}

func (s *ApplyModifyRequestDefaultStandard) GetEconomyDiscount() *int32 {
	return s.EconomyDiscount
}

func (s *ApplyModifyRequestDefaultStandard) GetFirstDiscount() *int32 {
	return s.FirstDiscount
}

func (s *ApplyModifyRequestDefaultStandard) GetFlightCabins() *string {
	return s.FlightCabins
}

func (s *ApplyModifyRequestDefaultStandard) GetFlightIntlRuleCode() *int64 {
	return s.FlightIntlRuleCode
}

func (s *ApplyModifyRequestDefaultStandard) GetFlightRuleCode() *int64 {
	return s.FlightRuleCode
}

func (s *ApplyModifyRequestDefaultStandard) GetHotelCitys() []*ApplyModifyRequestDefaultStandardHotelCitys {
	return s.HotelCitys
}

func (s *ApplyModifyRequestDefaultStandard) GetHotelIntlCitys() []*ApplyModifyRequestDefaultStandardHotelIntlCitys {
	return s.HotelIntlCitys
}

func (s *ApplyModifyRequestDefaultStandard) GetHotelIntlRuleCode() *int64 {
	return s.HotelIntlRuleCode
}

func (s *ApplyModifyRequestDefaultStandard) GetHotelRuleCode() *int64 {
	return s.HotelRuleCode
}

func (s *ApplyModifyRequestDefaultStandard) GetInternationalFlightCabins() *string {
	return s.InternationalFlightCabins
}

func (s *ApplyModifyRequestDefaultStandard) GetPremiumEconomyDiscount() *int32 {
	return s.PremiumEconomyDiscount
}

func (s *ApplyModifyRequestDefaultStandard) GetReserveType() *int32 {
	return s.ReserveType
}

func (s *ApplyModifyRequestDefaultStandard) GetTrainRuleCode() *int64 {
	return s.TrainRuleCode
}

func (s *ApplyModifyRequestDefaultStandard) GetTrainSeats() *string {
	return s.TrainSeats
}

func (s *ApplyModifyRequestDefaultStandard) SetBusinessDiscount(v int32) *ApplyModifyRequestDefaultStandard {
	s.BusinessDiscount = &v
	return s
}

func (s *ApplyModifyRequestDefaultStandard) SetEconomyDiscount(v int32) *ApplyModifyRequestDefaultStandard {
	s.EconomyDiscount = &v
	return s
}

func (s *ApplyModifyRequestDefaultStandard) SetFirstDiscount(v int32) *ApplyModifyRequestDefaultStandard {
	s.FirstDiscount = &v
	return s
}

func (s *ApplyModifyRequestDefaultStandard) SetFlightCabins(v string) *ApplyModifyRequestDefaultStandard {
	s.FlightCabins = &v
	return s
}

func (s *ApplyModifyRequestDefaultStandard) SetFlightIntlRuleCode(v int64) *ApplyModifyRequestDefaultStandard {
	s.FlightIntlRuleCode = &v
	return s
}

func (s *ApplyModifyRequestDefaultStandard) SetFlightRuleCode(v int64) *ApplyModifyRequestDefaultStandard {
	s.FlightRuleCode = &v
	return s
}

func (s *ApplyModifyRequestDefaultStandard) SetHotelCitys(v []*ApplyModifyRequestDefaultStandardHotelCitys) *ApplyModifyRequestDefaultStandard {
	s.HotelCitys = v
	return s
}

func (s *ApplyModifyRequestDefaultStandard) SetHotelIntlCitys(v []*ApplyModifyRequestDefaultStandardHotelIntlCitys) *ApplyModifyRequestDefaultStandard {
	s.HotelIntlCitys = v
	return s
}

func (s *ApplyModifyRequestDefaultStandard) SetHotelIntlRuleCode(v int64) *ApplyModifyRequestDefaultStandard {
	s.HotelIntlRuleCode = &v
	return s
}

func (s *ApplyModifyRequestDefaultStandard) SetHotelRuleCode(v int64) *ApplyModifyRequestDefaultStandard {
	s.HotelRuleCode = &v
	return s
}

func (s *ApplyModifyRequestDefaultStandard) SetInternationalFlightCabins(v string) *ApplyModifyRequestDefaultStandard {
	s.InternationalFlightCabins = &v
	return s
}

func (s *ApplyModifyRequestDefaultStandard) SetPremiumEconomyDiscount(v int32) *ApplyModifyRequestDefaultStandard {
	s.PremiumEconomyDiscount = &v
	return s
}

func (s *ApplyModifyRequestDefaultStandard) SetReserveType(v int32) *ApplyModifyRequestDefaultStandard {
	s.ReserveType = &v
	return s
}

func (s *ApplyModifyRequestDefaultStandard) SetTrainRuleCode(v int64) *ApplyModifyRequestDefaultStandard {
	s.TrainRuleCode = &v
	return s
}

func (s *ApplyModifyRequestDefaultStandard) SetTrainSeats(v string) *ApplyModifyRequestDefaultStandard {
	s.TrainSeats = &v
	return s
}

func (s *ApplyModifyRequestDefaultStandard) Validate() error {
	if s.HotelCitys != nil {
		for _, item := range s.HotelCitys {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.HotelIntlCitys != nil {
		for _, item := range s.HotelIntlCitys {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ApplyModifyRequestDefaultStandardHotelCitys struct {
	CityCode *string `json:"city_code,omitempty" xml:"city_code,omitempty"`
	CityName *string `json:"city_name,omitempty" xml:"city_name,omitempty"`
	Fee      *int64  `json:"fee,omitempty" xml:"fee,omitempty"`
}

func (s ApplyModifyRequestDefaultStandardHotelCitys) String() string {
	return dara.Prettify(s)
}

func (s ApplyModifyRequestDefaultStandardHotelCitys) GoString() string {
	return s.String()
}

func (s *ApplyModifyRequestDefaultStandardHotelCitys) GetCityCode() *string {
	return s.CityCode
}

func (s *ApplyModifyRequestDefaultStandardHotelCitys) GetCityName() *string {
	return s.CityName
}

func (s *ApplyModifyRequestDefaultStandardHotelCitys) GetFee() *int64 {
	return s.Fee
}

func (s *ApplyModifyRequestDefaultStandardHotelCitys) SetCityCode(v string) *ApplyModifyRequestDefaultStandardHotelCitys {
	s.CityCode = &v
	return s
}

func (s *ApplyModifyRequestDefaultStandardHotelCitys) SetCityName(v string) *ApplyModifyRequestDefaultStandardHotelCitys {
	s.CityName = &v
	return s
}

func (s *ApplyModifyRequestDefaultStandardHotelCitys) SetFee(v int64) *ApplyModifyRequestDefaultStandardHotelCitys {
	s.Fee = &v
	return s
}

func (s *ApplyModifyRequestDefaultStandardHotelCitys) Validate() error {
	return dara.Validate(s)
}

type ApplyModifyRequestDefaultStandardHotelIntlCitys struct {
	CityCode *string `json:"city_code,omitempty" xml:"city_code,omitempty"`
	CityName *string `json:"city_name,omitempty" xml:"city_name,omitempty"`
	Fee      *int64  `json:"fee,omitempty" xml:"fee,omitempty"`
}

func (s ApplyModifyRequestDefaultStandardHotelIntlCitys) String() string {
	return dara.Prettify(s)
}

func (s ApplyModifyRequestDefaultStandardHotelIntlCitys) GoString() string {
	return s.String()
}

func (s *ApplyModifyRequestDefaultStandardHotelIntlCitys) GetCityCode() *string {
	return s.CityCode
}

func (s *ApplyModifyRequestDefaultStandardHotelIntlCitys) GetCityName() *string {
	return s.CityName
}

func (s *ApplyModifyRequestDefaultStandardHotelIntlCitys) GetFee() *int64 {
	return s.Fee
}

func (s *ApplyModifyRequestDefaultStandardHotelIntlCitys) SetCityCode(v string) *ApplyModifyRequestDefaultStandardHotelIntlCitys {
	s.CityCode = &v
	return s
}

func (s *ApplyModifyRequestDefaultStandardHotelIntlCitys) SetCityName(v string) *ApplyModifyRequestDefaultStandardHotelIntlCitys {
	s.CityName = &v
	return s
}

func (s *ApplyModifyRequestDefaultStandardHotelIntlCitys) SetFee(v int64) *ApplyModifyRequestDefaultStandardHotelIntlCitys {
	s.Fee = &v
	return s
}

func (s *ApplyModifyRequestDefaultStandardHotelIntlCitys) Validate() error {
	return dara.Validate(s)
}

type ApplyModifyRequestExternalTravelerList struct {
	Attribute             *string `json:"attribute,omitempty" xml:"attribute,omitempty"`
	CostCenterId          *int64  `json:"cost_center_id,omitempty" xml:"cost_center_id,omitempty"`
	ExternalUserId        *string `json:"external_user_id,omitempty" xml:"external_user_id,omitempty"`
	InvoiceId             *int64  `json:"invoice_id,omitempty" xml:"invoice_id,omitempty"`
	PaymentDepartmentId   *string `json:"payment_department_id,omitempty" xml:"payment_department_id,omitempty"`
	PaymentDepartmentName *string `json:"payment_department_name,omitempty" xml:"payment_department_name,omitempty"`
	ProjectCode           *string `json:"project_code,omitempty" xml:"project_code,omitempty"`
	ProjectTitle          *string `json:"project_title,omitempty" xml:"project_title,omitempty"`
	ThirdPartInvoiceId    *string `json:"third_part_invoice_id,omitempty" xml:"third_part_invoice_id,omitempty"`
	ThirdpartCostCenterId *string `json:"thirdpart_cost_center_id,omitempty" xml:"thirdpart_cost_center_id,omitempty"`
	ThirdpartDepartId     *string `json:"thirdpart_depart_id,omitempty" xml:"thirdpart_depart_id,omitempty"`
	UserName              *string `json:"user_name,omitempty" xml:"user_name,omitempty"`
	UserNameEn            *string `json:"user_name_en,omitempty" xml:"user_name_en,omitempty"`
}

func (s ApplyModifyRequestExternalTravelerList) String() string {
	return dara.Prettify(s)
}

func (s ApplyModifyRequestExternalTravelerList) GoString() string {
	return s.String()
}

func (s *ApplyModifyRequestExternalTravelerList) GetAttribute() *string {
	return s.Attribute
}

func (s *ApplyModifyRequestExternalTravelerList) GetCostCenterId() *int64 {
	return s.CostCenterId
}

func (s *ApplyModifyRequestExternalTravelerList) GetExternalUserId() *string {
	return s.ExternalUserId
}

func (s *ApplyModifyRequestExternalTravelerList) GetInvoiceId() *int64 {
	return s.InvoiceId
}

func (s *ApplyModifyRequestExternalTravelerList) GetPaymentDepartmentId() *string {
	return s.PaymentDepartmentId
}

func (s *ApplyModifyRequestExternalTravelerList) GetPaymentDepartmentName() *string {
	return s.PaymentDepartmentName
}

func (s *ApplyModifyRequestExternalTravelerList) GetProjectCode() *string {
	return s.ProjectCode
}

func (s *ApplyModifyRequestExternalTravelerList) GetProjectTitle() *string {
	return s.ProjectTitle
}

func (s *ApplyModifyRequestExternalTravelerList) GetThirdPartInvoiceId() *string {
	return s.ThirdPartInvoiceId
}

func (s *ApplyModifyRequestExternalTravelerList) GetThirdpartCostCenterId() *string {
	return s.ThirdpartCostCenterId
}

func (s *ApplyModifyRequestExternalTravelerList) GetThirdpartDepartId() *string {
	return s.ThirdpartDepartId
}

func (s *ApplyModifyRequestExternalTravelerList) GetUserName() *string {
	return s.UserName
}

func (s *ApplyModifyRequestExternalTravelerList) GetUserNameEn() *string {
	return s.UserNameEn
}

func (s *ApplyModifyRequestExternalTravelerList) SetAttribute(v string) *ApplyModifyRequestExternalTravelerList {
	s.Attribute = &v
	return s
}

func (s *ApplyModifyRequestExternalTravelerList) SetCostCenterId(v int64) *ApplyModifyRequestExternalTravelerList {
	s.CostCenterId = &v
	return s
}

func (s *ApplyModifyRequestExternalTravelerList) SetExternalUserId(v string) *ApplyModifyRequestExternalTravelerList {
	s.ExternalUserId = &v
	return s
}

func (s *ApplyModifyRequestExternalTravelerList) SetInvoiceId(v int64) *ApplyModifyRequestExternalTravelerList {
	s.InvoiceId = &v
	return s
}

func (s *ApplyModifyRequestExternalTravelerList) SetPaymentDepartmentId(v string) *ApplyModifyRequestExternalTravelerList {
	s.PaymentDepartmentId = &v
	return s
}

func (s *ApplyModifyRequestExternalTravelerList) SetPaymentDepartmentName(v string) *ApplyModifyRequestExternalTravelerList {
	s.PaymentDepartmentName = &v
	return s
}

func (s *ApplyModifyRequestExternalTravelerList) SetProjectCode(v string) *ApplyModifyRequestExternalTravelerList {
	s.ProjectCode = &v
	return s
}

func (s *ApplyModifyRequestExternalTravelerList) SetProjectTitle(v string) *ApplyModifyRequestExternalTravelerList {
	s.ProjectTitle = &v
	return s
}

func (s *ApplyModifyRequestExternalTravelerList) SetThirdPartInvoiceId(v string) *ApplyModifyRequestExternalTravelerList {
	s.ThirdPartInvoiceId = &v
	return s
}

func (s *ApplyModifyRequestExternalTravelerList) SetThirdpartCostCenterId(v string) *ApplyModifyRequestExternalTravelerList {
	s.ThirdpartCostCenterId = &v
	return s
}

func (s *ApplyModifyRequestExternalTravelerList) SetThirdpartDepartId(v string) *ApplyModifyRequestExternalTravelerList {
	s.ThirdpartDepartId = &v
	return s
}

func (s *ApplyModifyRequestExternalTravelerList) SetUserName(v string) *ApplyModifyRequestExternalTravelerList {
	s.UserName = &v
	return s
}

func (s *ApplyModifyRequestExternalTravelerList) SetUserNameEn(v string) *ApplyModifyRequestExternalTravelerList {
	s.UserNameEn = &v
	return s
}

func (s *ApplyModifyRequestExternalTravelerList) Validate() error {
	return dara.Validate(s)
}

type ApplyModifyRequestExternalTravelerStandard struct {
	BusinessDiscount          *int32                                                      `json:"business_discount,omitempty" xml:"business_discount,omitempty"`
	EconomyDiscount           *int32                                                      `json:"economy_discount,omitempty" xml:"economy_discount,omitempty"`
	FirstDiscount             *int32                                                      `json:"first_discount,omitempty" xml:"first_discount,omitempty"`
	FlightCabins              *string                                                     `json:"flight_cabins,omitempty" xml:"flight_cabins,omitempty"`
	FlightIntlRuleCode        *int64                                                      `json:"flight_intl_rule_code,omitempty" xml:"flight_intl_rule_code,omitempty"`
	FlightRuleCode            *int64                                                      `json:"flight_rule_code,omitempty" xml:"flight_rule_code,omitempty"`
	HotelCitys                []*ApplyModifyRequestExternalTravelerStandardHotelCitys     `json:"hotel_citys,omitempty" xml:"hotel_citys,omitempty" type:"Repeated"`
	HotelIntlCitys            []*ApplyModifyRequestExternalTravelerStandardHotelIntlCitys `json:"hotel_intl_citys,omitempty" xml:"hotel_intl_citys,omitempty" type:"Repeated"`
	HotelIntlRuleCode         *int64                                                      `json:"hotel_intl_rule_code,omitempty" xml:"hotel_intl_rule_code,omitempty"`
	HotelRuleCode             *int64                                                      `json:"hotel_rule_code,omitempty" xml:"hotel_rule_code,omitempty"`
	InternationalFlightCabins *string                                                     `json:"international_flight_cabins,omitempty" xml:"international_flight_cabins,omitempty"`
	PremiumEconomyDiscount    *int32                                                      `json:"premium_economy_discount,omitempty" xml:"premium_economy_discount,omitempty"`
	ReserveType               *int32                                                      `json:"reserve_type,omitempty" xml:"reserve_type,omitempty"`
	TrainRuleCode             *int64                                                      `json:"train_rule_code,omitempty" xml:"train_rule_code,omitempty"`
	TrainSeats                *string                                                     `json:"train_seats,omitempty" xml:"train_seats,omitempty"`
}

func (s ApplyModifyRequestExternalTravelerStandard) String() string {
	return dara.Prettify(s)
}

func (s ApplyModifyRequestExternalTravelerStandard) GoString() string {
	return s.String()
}

func (s *ApplyModifyRequestExternalTravelerStandard) GetBusinessDiscount() *int32 {
	return s.BusinessDiscount
}

func (s *ApplyModifyRequestExternalTravelerStandard) GetEconomyDiscount() *int32 {
	return s.EconomyDiscount
}

func (s *ApplyModifyRequestExternalTravelerStandard) GetFirstDiscount() *int32 {
	return s.FirstDiscount
}

func (s *ApplyModifyRequestExternalTravelerStandard) GetFlightCabins() *string {
	return s.FlightCabins
}

func (s *ApplyModifyRequestExternalTravelerStandard) GetFlightIntlRuleCode() *int64 {
	return s.FlightIntlRuleCode
}

func (s *ApplyModifyRequestExternalTravelerStandard) GetFlightRuleCode() *int64 {
	return s.FlightRuleCode
}

func (s *ApplyModifyRequestExternalTravelerStandard) GetHotelCitys() []*ApplyModifyRequestExternalTravelerStandardHotelCitys {
	return s.HotelCitys
}

func (s *ApplyModifyRequestExternalTravelerStandard) GetHotelIntlCitys() []*ApplyModifyRequestExternalTravelerStandardHotelIntlCitys {
	return s.HotelIntlCitys
}

func (s *ApplyModifyRequestExternalTravelerStandard) GetHotelIntlRuleCode() *int64 {
	return s.HotelIntlRuleCode
}

func (s *ApplyModifyRequestExternalTravelerStandard) GetHotelRuleCode() *int64 {
	return s.HotelRuleCode
}

func (s *ApplyModifyRequestExternalTravelerStandard) GetInternationalFlightCabins() *string {
	return s.InternationalFlightCabins
}

func (s *ApplyModifyRequestExternalTravelerStandard) GetPremiumEconomyDiscount() *int32 {
	return s.PremiumEconomyDiscount
}

func (s *ApplyModifyRequestExternalTravelerStandard) GetReserveType() *int32 {
	return s.ReserveType
}

func (s *ApplyModifyRequestExternalTravelerStandard) GetTrainRuleCode() *int64 {
	return s.TrainRuleCode
}

func (s *ApplyModifyRequestExternalTravelerStandard) GetTrainSeats() *string {
	return s.TrainSeats
}

func (s *ApplyModifyRequestExternalTravelerStandard) SetBusinessDiscount(v int32) *ApplyModifyRequestExternalTravelerStandard {
	s.BusinessDiscount = &v
	return s
}

func (s *ApplyModifyRequestExternalTravelerStandard) SetEconomyDiscount(v int32) *ApplyModifyRequestExternalTravelerStandard {
	s.EconomyDiscount = &v
	return s
}

func (s *ApplyModifyRequestExternalTravelerStandard) SetFirstDiscount(v int32) *ApplyModifyRequestExternalTravelerStandard {
	s.FirstDiscount = &v
	return s
}

func (s *ApplyModifyRequestExternalTravelerStandard) SetFlightCabins(v string) *ApplyModifyRequestExternalTravelerStandard {
	s.FlightCabins = &v
	return s
}

func (s *ApplyModifyRequestExternalTravelerStandard) SetFlightIntlRuleCode(v int64) *ApplyModifyRequestExternalTravelerStandard {
	s.FlightIntlRuleCode = &v
	return s
}

func (s *ApplyModifyRequestExternalTravelerStandard) SetFlightRuleCode(v int64) *ApplyModifyRequestExternalTravelerStandard {
	s.FlightRuleCode = &v
	return s
}

func (s *ApplyModifyRequestExternalTravelerStandard) SetHotelCitys(v []*ApplyModifyRequestExternalTravelerStandardHotelCitys) *ApplyModifyRequestExternalTravelerStandard {
	s.HotelCitys = v
	return s
}

func (s *ApplyModifyRequestExternalTravelerStandard) SetHotelIntlCitys(v []*ApplyModifyRequestExternalTravelerStandardHotelIntlCitys) *ApplyModifyRequestExternalTravelerStandard {
	s.HotelIntlCitys = v
	return s
}

func (s *ApplyModifyRequestExternalTravelerStandard) SetHotelIntlRuleCode(v int64) *ApplyModifyRequestExternalTravelerStandard {
	s.HotelIntlRuleCode = &v
	return s
}

func (s *ApplyModifyRequestExternalTravelerStandard) SetHotelRuleCode(v int64) *ApplyModifyRequestExternalTravelerStandard {
	s.HotelRuleCode = &v
	return s
}

func (s *ApplyModifyRequestExternalTravelerStandard) SetInternationalFlightCabins(v string) *ApplyModifyRequestExternalTravelerStandard {
	s.InternationalFlightCabins = &v
	return s
}

func (s *ApplyModifyRequestExternalTravelerStandard) SetPremiumEconomyDiscount(v int32) *ApplyModifyRequestExternalTravelerStandard {
	s.PremiumEconomyDiscount = &v
	return s
}

func (s *ApplyModifyRequestExternalTravelerStandard) SetReserveType(v int32) *ApplyModifyRequestExternalTravelerStandard {
	s.ReserveType = &v
	return s
}

func (s *ApplyModifyRequestExternalTravelerStandard) SetTrainRuleCode(v int64) *ApplyModifyRequestExternalTravelerStandard {
	s.TrainRuleCode = &v
	return s
}

func (s *ApplyModifyRequestExternalTravelerStandard) SetTrainSeats(v string) *ApplyModifyRequestExternalTravelerStandard {
	s.TrainSeats = &v
	return s
}

func (s *ApplyModifyRequestExternalTravelerStandard) Validate() error {
	if s.HotelCitys != nil {
		for _, item := range s.HotelCitys {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.HotelIntlCitys != nil {
		for _, item := range s.HotelIntlCitys {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ApplyModifyRequestExternalTravelerStandardHotelCitys struct {
	CityCode *string `json:"city_code,omitempty" xml:"city_code,omitempty"`
	CityName *string `json:"city_name,omitempty" xml:"city_name,omitempty"`
	Fee      *int64  `json:"fee,omitempty" xml:"fee,omitempty"`
}

func (s ApplyModifyRequestExternalTravelerStandardHotelCitys) String() string {
	return dara.Prettify(s)
}

func (s ApplyModifyRequestExternalTravelerStandardHotelCitys) GoString() string {
	return s.String()
}

func (s *ApplyModifyRequestExternalTravelerStandardHotelCitys) GetCityCode() *string {
	return s.CityCode
}

func (s *ApplyModifyRequestExternalTravelerStandardHotelCitys) GetCityName() *string {
	return s.CityName
}

func (s *ApplyModifyRequestExternalTravelerStandardHotelCitys) GetFee() *int64 {
	return s.Fee
}

func (s *ApplyModifyRequestExternalTravelerStandardHotelCitys) SetCityCode(v string) *ApplyModifyRequestExternalTravelerStandardHotelCitys {
	s.CityCode = &v
	return s
}

func (s *ApplyModifyRequestExternalTravelerStandardHotelCitys) SetCityName(v string) *ApplyModifyRequestExternalTravelerStandardHotelCitys {
	s.CityName = &v
	return s
}

func (s *ApplyModifyRequestExternalTravelerStandardHotelCitys) SetFee(v int64) *ApplyModifyRequestExternalTravelerStandardHotelCitys {
	s.Fee = &v
	return s
}

func (s *ApplyModifyRequestExternalTravelerStandardHotelCitys) Validate() error {
	return dara.Validate(s)
}

type ApplyModifyRequestExternalTravelerStandardHotelIntlCitys struct {
	CityCode *string `json:"city_code,omitempty" xml:"city_code,omitempty"`
	CityName *string `json:"city_name,omitempty" xml:"city_name,omitempty"`
	Fee      *int64  `json:"fee,omitempty" xml:"fee,omitempty"`
}

func (s ApplyModifyRequestExternalTravelerStandardHotelIntlCitys) String() string {
	return dara.Prettify(s)
}

func (s ApplyModifyRequestExternalTravelerStandardHotelIntlCitys) GoString() string {
	return s.String()
}

func (s *ApplyModifyRequestExternalTravelerStandardHotelIntlCitys) GetCityCode() *string {
	return s.CityCode
}

func (s *ApplyModifyRequestExternalTravelerStandardHotelIntlCitys) GetCityName() *string {
	return s.CityName
}

func (s *ApplyModifyRequestExternalTravelerStandardHotelIntlCitys) GetFee() *int64 {
	return s.Fee
}

func (s *ApplyModifyRequestExternalTravelerStandardHotelIntlCitys) SetCityCode(v string) *ApplyModifyRequestExternalTravelerStandardHotelIntlCitys {
	s.CityCode = &v
	return s
}

func (s *ApplyModifyRequestExternalTravelerStandardHotelIntlCitys) SetCityName(v string) *ApplyModifyRequestExternalTravelerStandardHotelIntlCitys {
	s.CityName = &v
	return s
}

func (s *ApplyModifyRequestExternalTravelerStandardHotelIntlCitys) SetFee(v int64) *ApplyModifyRequestExternalTravelerStandardHotelIntlCitys {
	s.Fee = &v
	return s
}

func (s *ApplyModifyRequestExternalTravelerStandardHotelIntlCitys) Validate() error {
	return dara.Validate(s)
}

type ApplyModifyRequestHotelShare struct {
	Param *string `json:"param,omitempty" xml:"param,omitempty"`
	Type  *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ApplyModifyRequestHotelShare) String() string {
	return dara.Prettify(s)
}

func (s ApplyModifyRequestHotelShare) GoString() string {
	return s.String()
}

func (s *ApplyModifyRequestHotelShare) GetParam() *string {
	return s.Param
}

func (s *ApplyModifyRequestHotelShare) GetType() *string {
	return s.Type
}

func (s *ApplyModifyRequestHotelShare) SetParam(v string) *ApplyModifyRequestHotelShare {
	s.Param = &v
	return s
}

func (s *ApplyModifyRequestHotelShare) SetType(v string) *ApplyModifyRequestHotelShare {
	s.Type = &v
	return s
}

func (s *ApplyModifyRequestHotelShare) Validate() error {
	return dara.Validate(s)
}

type ApplyModifyRequestItineraryList struct {
	// This parameter is required.
	ArrCity     *string `json:"arr_city,omitempty" xml:"arr_city,omitempty"`
	ArrCityCode *string `json:"arr_city_code,omitempty" xml:"arr_city_code,omitempty"`
	// This parameter is required.
	ArrDate      *string `json:"arr_date,omitempty" xml:"arr_date,omitempty"`
	Attribute    *string `json:"attribute,omitempty" xml:"attribute,omitempty"`
	CostCenterId *int64  `json:"cost_center_id,omitempty" xml:"cost_center_id,omitempty"`
	// This parameter is required.
	DepCity     *string `json:"dep_city,omitempty" xml:"dep_city,omitempty"`
	DepCityCode *string `json:"dep_city_code,omitempty" xml:"dep_city_code,omitempty"`
	// This parameter is required.
	DepDate   *string `json:"dep_date,omitempty" xml:"dep_date,omitempty"`
	InvoiceId *int64  `json:"invoice_id,omitempty" xml:"invoice_id,omitempty"`
	// This parameter is required.
	ItineraryId               *string                                                 `json:"itinerary_id,omitempty" xml:"itinerary_id,omitempty"`
	ItineraryTravelStandard   *ApplyModifyRequestItineraryListItineraryTravelStandard `json:"itinerary_travel_standard,omitempty" xml:"itinerary_travel_standard,omitempty" type:"Struct"`
	NeedHotel                 *bool                                                   `json:"need_hotel,omitempty" xml:"need_hotel,omitempty"`
	NeedTraffic               *bool                                                   `json:"need_traffic,omitempty" xml:"need_traffic,omitempty"`
	ProjectCode               *string                                                 `json:"project_code,omitempty" xml:"project_code,omitempty"`
	ProjectTitle              *string                                                 `json:"project_title,omitempty" xml:"project_title,omitempty"`
	ProvinceTravelCityAdcodes []*string                                               `json:"province_travel_city_adcodes,omitempty" xml:"province_travel_city_adcodes,omitempty" type:"Repeated"`
	ThirdPartInvoiceId        *string                                                 `json:"third_part_invoice_id,omitempty" xml:"third_part_invoice_id,omitempty"`
	ThirdpartCostCenterId     *string                                                 `json:"thirdpart_cost_center_id,omitempty" xml:"thirdpart_cost_center_id,omitempty"`
	// This parameter is required.
	TrafficType *int32 `json:"traffic_type,omitempty" xml:"traffic_type,omitempty"`
	// This parameter is required.
	TripWay *int32 `json:"trip_way,omitempty" xml:"trip_way,omitempty"`
}

func (s ApplyModifyRequestItineraryList) String() string {
	return dara.Prettify(s)
}

func (s ApplyModifyRequestItineraryList) GoString() string {
	return s.String()
}

func (s *ApplyModifyRequestItineraryList) GetArrCity() *string {
	return s.ArrCity
}

func (s *ApplyModifyRequestItineraryList) GetArrCityCode() *string {
	return s.ArrCityCode
}

func (s *ApplyModifyRequestItineraryList) GetArrDate() *string {
	return s.ArrDate
}

func (s *ApplyModifyRequestItineraryList) GetAttribute() *string {
	return s.Attribute
}

func (s *ApplyModifyRequestItineraryList) GetCostCenterId() *int64 {
	return s.CostCenterId
}

func (s *ApplyModifyRequestItineraryList) GetDepCity() *string {
	return s.DepCity
}

func (s *ApplyModifyRequestItineraryList) GetDepCityCode() *string {
	return s.DepCityCode
}

func (s *ApplyModifyRequestItineraryList) GetDepDate() *string {
	return s.DepDate
}

func (s *ApplyModifyRequestItineraryList) GetInvoiceId() *int64 {
	return s.InvoiceId
}

func (s *ApplyModifyRequestItineraryList) GetItineraryId() *string {
	return s.ItineraryId
}

func (s *ApplyModifyRequestItineraryList) GetItineraryTravelStandard() *ApplyModifyRequestItineraryListItineraryTravelStandard {
	return s.ItineraryTravelStandard
}

func (s *ApplyModifyRequestItineraryList) GetNeedHotel() *bool {
	return s.NeedHotel
}

func (s *ApplyModifyRequestItineraryList) GetNeedTraffic() *bool {
	return s.NeedTraffic
}

func (s *ApplyModifyRequestItineraryList) GetProjectCode() *string {
	return s.ProjectCode
}

func (s *ApplyModifyRequestItineraryList) GetProjectTitle() *string {
	return s.ProjectTitle
}

func (s *ApplyModifyRequestItineraryList) GetProvinceTravelCityAdcodes() []*string {
	return s.ProvinceTravelCityAdcodes
}

func (s *ApplyModifyRequestItineraryList) GetThirdPartInvoiceId() *string {
	return s.ThirdPartInvoiceId
}

func (s *ApplyModifyRequestItineraryList) GetThirdpartCostCenterId() *string {
	return s.ThirdpartCostCenterId
}

func (s *ApplyModifyRequestItineraryList) GetTrafficType() *int32 {
	return s.TrafficType
}

func (s *ApplyModifyRequestItineraryList) GetTripWay() *int32 {
	return s.TripWay
}

func (s *ApplyModifyRequestItineraryList) SetArrCity(v string) *ApplyModifyRequestItineraryList {
	s.ArrCity = &v
	return s
}

func (s *ApplyModifyRequestItineraryList) SetArrCityCode(v string) *ApplyModifyRequestItineraryList {
	s.ArrCityCode = &v
	return s
}

func (s *ApplyModifyRequestItineraryList) SetArrDate(v string) *ApplyModifyRequestItineraryList {
	s.ArrDate = &v
	return s
}

func (s *ApplyModifyRequestItineraryList) SetAttribute(v string) *ApplyModifyRequestItineraryList {
	s.Attribute = &v
	return s
}

func (s *ApplyModifyRequestItineraryList) SetCostCenterId(v int64) *ApplyModifyRequestItineraryList {
	s.CostCenterId = &v
	return s
}

func (s *ApplyModifyRequestItineraryList) SetDepCity(v string) *ApplyModifyRequestItineraryList {
	s.DepCity = &v
	return s
}

func (s *ApplyModifyRequestItineraryList) SetDepCityCode(v string) *ApplyModifyRequestItineraryList {
	s.DepCityCode = &v
	return s
}

func (s *ApplyModifyRequestItineraryList) SetDepDate(v string) *ApplyModifyRequestItineraryList {
	s.DepDate = &v
	return s
}

func (s *ApplyModifyRequestItineraryList) SetInvoiceId(v int64) *ApplyModifyRequestItineraryList {
	s.InvoiceId = &v
	return s
}

func (s *ApplyModifyRequestItineraryList) SetItineraryId(v string) *ApplyModifyRequestItineraryList {
	s.ItineraryId = &v
	return s
}

func (s *ApplyModifyRequestItineraryList) SetItineraryTravelStandard(v *ApplyModifyRequestItineraryListItineraryTravelStandard) *ApplyModifyRequestItineraryList {
	s.ItineraryTravelStandard = v
	return s
}

func (s *ApplyModifyRequestItineraryList) SetNeedHotel(v bool) *ApplyModifyRequestItineraryList {
	s.NeedHotel = &v
	return s
}

func (s *ApplyModifyRequestItineraryList) SetNeedTraffic(v bool) *ApplyModifyRequestItineraryList {
	s.NeedTraffic = &v
	return s
}

func (s *ApplyModifyRequestItineraryList) SetProjectCode(v string) *ApplyModifyRequestItineraryList {
	s.ProjectCode = &v
	return s
}

func (s *ApplyModifyRequestItineraryList) SetProjectTitle(v string) *ApplyModifyRequestItineraryList {
	s.ProjectTitle = &v
	return s
}

func (s *ApplyModifyRequestItineraryList) SetProvinceTravelCityAdcodes(v []*string) *ApplyModifyRequestItineraryList {
	s.ProvinceTravelCityAdcodes = v
	return s
}

func (s *ApplyModifyRequestItineraryList) SetThirdPartInvoiceId(v string) *ApplyModifyRequestItineraryList {
	s.ThirdPartInvoiceId = &v
	return s
}

func (s *ApplyModifyRequestItineraryList) SetThirdpartCostCenterId(v string) *ApplyModifyRequestItineraryList {
	s.ThirdpartCostCenterId = &v
	return s
}

func (s *ApplyModifyRequestItineraryList) SetTrafficType(v int32) *ApplyModifyRequestItineraryList {
	s.TrafficType = &v
	return s
}

func (s *ApplyModifyRequestItineraryList) SetTripWay(v int32) *ApplyModifyRequestItineraryList {
	s.TripWay = &v
	return s
}

func (s *ApplyModifyRequestItineraryList) Validate() error {
	if s.ItineraryTravelStandard != nil {
		if err := s.ItineraryTravelStandard.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ApplyModifyRequestItineraryListItineraryTravelStandard struct {
	HotelAvailableNightsPerDay *int32 `json:"hotel_available_nights_per_day,omitempty" xml:"hotel_available_nights_per_day,omitempty"`
}

func (s ApplyModifyRequestItineraryListItineraryTravelStandard) String() string {
	return dara.Prettify(s)
}

func (s ApplyModifyRequestItineraryListItineraryTravelStandard) GoString() string {
	return s.String()
}

func (s *ApplyModifyRequestItineraryListItineraryTravelStandard) GetHotelAvailableNightsPerDay() *int32 {
	return s.HotelAvailableNightsPerDay
}

func (s *ApplyModifyRequestItineraryListItineraryTravelStandard) SetHotelAvailableNightsPerDay(v int32) *ApplyModifyRequestItineraryListItineraryTravelStandard {
	s.HotelAvailableNightsPerDay = &v
	return s
}

func (s *ApplyModifyRequestItineraryListItineraryTravelStandard) Validate() error {
	return dara.Validate(s)
}

type ApplyModifyRequestItinerarySetList struct {
	// example:
	//
	// 2017-01-01 00:00:00
	ArrDate      *string `json:"arr_date,omitempty" xml:"arr_date,omitempty"`
	Attribute    *string `json:"attribute,omitempty" xml:"attribute,omitempty"`
	CityCodeSet  *string `json:"city_code_set,omitempty" xml:"city_code_set,omitempty"`
	CitySet      *string `json:"city_set,omitempty" xml:"city_set,omitempty"`
	CostCenterId *int64  `json:"cost_center_id,omitempty" xml:"cost_center_id,omitempty"`
	// example:
	//
	// 2017-01-01 00:00:00
	DepDate                   *string                                                    `json:"dep_date,omitempty" xml:"dep_date,omitempty"`
	InvoiceId                 *int64                                                     `json:"invoice_id,omitempty" xml:"invoice_id,omitempty"`
	ItineraryId               *string                                                    `json:"itinerary_id,omitempty" xml:"itinerary_id,omitempty"`
	ItineraryTravelStandard   *ApplyModifyRequestItinerarySetListItineraryTravelStandard `json:"itinerary_travel_standard,omitempty" xml:"itinerary_travel_standard,omitempty" type:"Struct"`
	ProjectCode               *string                                                    `json:"project_code,omitempty" xml:"project_code,omitempty"`
	ProjectTitle              *string                                                    `json:"project_title,omitempty" xml:"project_title,omitempty"`
	ProvinceTravelCityAdcodes []*string                                                  `json:"province_travel_city_adcodes,omitempty" xml:"province_travel_city_adcodes,omitempty" type:"Repeated"`
	ThirdPartInvoiceId        *string                                                    `json:"third_part_invoice_id,omitempty" xml:"third_part_invoice_id,omitempty"`
	ThirdpartCostCenterId     *string                                                    `json:"thirdpart_cost_center_id,omitempty" xml:"thirdpart_cost_center_id,omitempty"`
	TrafficType               *int32                                                     `json:"traffic_type,omitempty" xml:"traffic_type,omitempty"`
}

func (s ApplyModifyRequestItinerarySetList) String() string {
	return dara.Prettify(s)
}

func (s ApplyModifyRequestItinerarySetList) GoString() string {
	return s.String()
}

func (s *ApplyModifyRequestItinerarySetList) GetArrDate() *string {
	return s.ArrDate
}

func (s *ApplyModifyRequestItinerarySetList) GetAttribute() *string {
	return s.Attribute
}

func (s *ApplyModifyRequestItinerarySetList) GetCityCodeSet() *string {
	return s.CityCodeSet
}

func (s *ApplyModifyRequestItinerarySetList) GetCitySet() *string {
	return s.CitySet
}

func (s *ApplyModifyRequestItinerarySetList) GetCostCenterId() *int64 {
	return s.CostCenterId
}

func (s *ApplyModifyRequestItinerarySetList) GetDepDate() *string {
	return s.DepDate
}

func (s *ApplyModifyRequestItinerarySetList) GetInvoiceId() *int64 {
	return s.InvoiceId
}

func (s *ApplyModifyRequestItinerarySetList) GetItineraryId() *string {
	return s.ItineraryId
}

func (s *ApplyModifyRequestItinerarySetList) GetItineraryTravelStandard() *ApplyModifyRequestItinerarySetListItineraryTravelStandard {
	return s.ItineraryTravelStandard
}

func (s *ApplyModifyRequestItinerarySetList) GetProjectCode() *string {
	return s.ProjectCode
}

func (s *ApplyModifyRequestItinerarySetList) GetProjectTitle() *string {
	return s.ProjectTitle
}

func (s *ApplyModifyRequestItinerarySetList) GetProvinceTravelCityAdcodes() []*string {
	return s.ProvinceTravelCityAdcodes
}

func (s *ApplyModifyRequestItinerarySetList) GetThirdPartInvoiceId() *string {
	return s.ThirdPartInvoiceId
}

func (s *ApplyModifyRequestItinerarySetList) GetThirdpartCostCenterId() *string {
	return s.ThirdpartCostCenterId
}

func (s *ApplyModifyRequestItinerarySetList) GetTrafficType() *int32 {
	return s.TrafficType
}

func (s *ApplyModifyRequestItinerarySetList) SetArrDate(v string) *ApplyModifyRequestItinerarySetList {
	s.ArrDate = &v
	return s
}

func (s *ApplyModifyRequestItinerarySetList) SetAttribute(v string) *ApplyModifyRequestItinerarySetList {
	s.Attribute = &v
	return s
}

func (s *ApplyModifyRequestItinerarySetList) SetCityCodeSet(v string) *ApplyModifyRequestItinerarySetList {
	s.CityCodeSet = &v
	return s
}

func (s *ApplyModifyRequestItinerarySetList) SetCitySet(v string) *ApplyModifyRequestItinerarySetList {
	s.CitySet = &v
	return s
}

func (s *ApplyModifyRequestItinerarySetList) SetCostCenterId(v int64) *ApplyModifyRequestItinerarySetList {
	s.CostCenterId = &v
	return s
}

func (s *ApplyModifyRequestItinerarySetList) SetDepDate(v string) *ApplyModifyRequestItinerarySetList {
	s.DepDate = &v
	return s
}

func (s *ApplyModifyRequestItinerarySetList) SetInvoiceId(v int64) *ApplyModifyRequestItinerarySetList {
	s.InvoiceId = &v
	return s
}

func (s *ApplyModifyRequestItinerarySetList) SetItineraryId(v string) *ApplyModifyRequestItinerarySetList {
	s.ItineraryId = &v
	return s
}

func (s *ApplyModifyRequestItinerarySetList) SetItineraryTravelStandard(v *ApplyModifyRequestItinerarySetListItineraryTravelStandard) *ApplyModifyRequestItinerarySetList {
	s.ItineraryTravelStandard = v
	return s
}

func (s *ApplyModifyRequestItinerarySetList) SetProjectCode(v string) *ApplyModifyRequestItinerarySetList {
	s.ProjectCode = &v
	return s
}

func (s *ApplyModifyRequestItinerarySetList) SetProjectTitle(v string) *ApplyModifyRequestItinerarySetList {
	s.ProjectTitle = &v
	return s
}

func (s *ApplyModifyRequestItinerarySetList) SetProvinceTravelCityAdcodes(v []*string) *ApplyModifyRequestItinerarySetList {
	s.ProvinceTravelCityAdcodes = v
	return s
}

func (s *ApplyModifyRequestItinerarySetList) SetThirdPartInvoiceId(v string) *ApplyModifyRequestItinerarySetList {
	s.ThirdPartInvoiceId = &v
	return s
}

func (s *ApplyModifyRequestItinerarySetList) SetThirdpartCostCenterId(v string) *ApplyModifyRequestItinerarySetList {
	s.ThirdpartCostCenterId = &v
	return s
}

func (s *ApplyModifyRequestItinerarySetList) SetTrafficType(v int32) *ApplyModifyRequestItinerarySetList {
	s.TrafficType = &v
	return s
}

func (s *ApplyModifyRequestItinerarySetList) Validate() error {
	if s.ItineraryTravelStandard != nil {
		if err := s.ItineraryTravelStandard.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ApplyModifyRequestItinerarySetListItineraryTravelStandard struct {
	HotelAvailableNightsPerDay *int32 `json:"hotel_available_nights_per_day,omitempty" xml:"hotel_available_nights_per_day,omitempty"`
}

func (s ApplyModifyRequestItinerarySetListItineraryTravelStandard) String() string {
	return dara.Prettify(s)
}

func (s ApplyModifyRequestItinerarySetListItineraryTravelStandard) GoString() string {
	return s.String()
}

func (s *ApplyModifyRequestItinerarySetListItineraryTravelStandard) GetHotelAvailableNightsPerDay() *int32 {
	return s.HotelAvailableNightsPerDay
}

func (s *ApplyModifyRequestItinerarySetListItineraryTravelStandard) SetHotelAvailableNightsPerDay(v int32) *ApplyModifyRequestItinerarySetListItineraryTravelStandard {
	s.HotelAvailableNightsPerDay = &v
	return s
}

func (s *ApplyModifyRequestItinerarySetListItineraryTravelStandard) Validate() error {
	return dara.Validate(s)
}

type ApplyModifyRequestTravelerList struct {
	Attribute             *string `json:"attribute,omitempty" xml:"attribute,omitempty"`
	CostCenterId          *int64  `json:"cost_center_id,omitempty" xml:"cost_center_id,omitempty"`
	InvoiceId             *int64  `json:"invoice_id,omitempty" xml:"invoice_id,omitempty"`
	PaymentDepartmentId   *string `json:"payment_department_id,omitempty" xml:"payment_department_id,omitempty"`
	PaymentDepartmentName *string `json:"payment_department_name,omitempty" xml:"payment_department_name,omitempty"`
	ProjectCode           *string `json:"project_code,omitempty" xml:"project_code,omitempty"`
	ProjectTitle          *string `json:"project_title,omitempty" xml:"project_title,omitempty"`
	ThirdPartInvoiceId    *string `json:"third_part_invoice_id,omitempty" xml:"third_part_invoice_id,omitempty"`
	ThirdpartCostCenterId *string `json:"thirdpart_cost_center_id,omitempty" xml:"thirdpart_cost_center_id,omitempty"`
	ThirdpartDepartId     *string `json:"thirdpart_depart_id,omitempty" xml:"thirdpart_depart_id,omitempty"`
	UserId                *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	UserName              *string `json:"user_name,omitempty" xml:"user_name,omitempty"`
}

func (s ApplyModifyRequestTravelerList) String() string {
	return dara.Prettify(s)
}

func (s ApplyModifyRequestTravelerList) GoString() string {
	return s.String()
}

func (s *ApplyModifyRequestTravelerList) GetAttribute() *string {
	return s.Attribute
}

func (s *ApplyModifyRequestTravelerList) GetCostCenterId() *int64 {
	return s.CostCenterId
}

func (s *ApplyModifyRequestTravelerList) GetInvoiceId() *int64 {
	return s.InvoiceId
}

func (s *ApplyModifyRequestTravelerList) GetPaymentDepartmentId() *string {
	return s.PaymentDepartmentId
}

func (s *ApplyModifyRequestTravelerList) GetPaymentDepartmentName() *string {
	return s.PaymentDepartmentName
}

func (s *ApplyModifyRequestTravelerList) GetProjectCode() *string {
	return s.ProjectCode
}

func (s *ApplyModifyRequestTravelerList) GetProjectTitle() *string {
	return s.ProjectTitle
}

func (s *ApplyModifyRequestTravelerList) GetThirdPartInvoiceId() *string {
	return s.ThirdPartInvoiceId
}

func (s *ApplyModifyRequestTravelerList) GetThirdpartCostCenterId() *string {
	return s.ThirdpartCostCenterId
}

func (s *ApplyModifyRequestTravelerList) GetThirdpartDepartId() *string {
	return s.ThirdpartDepartId
}

func (s *ApplyModifyRequestTravelerList) GetUserId() *string {
	return s.UserId
}

func (s *ApplyModifyRequestTravelerList) GetUserName() *string {
	return s.UserName
}

func (s *ApplyModifyRequestTravelerList) SetAttribute(v string) *ApplyModifyRequestTravelerList {
	s.Attribute = &v
	return s
}

func (s *ApplyModifyRequestTravelerList) SetCostCenterId(v int64) *ApplyModifyRequestTravelerList {
	s.CostCenterId = &v
	return s
}

func (s *ApplyModifyRequestTravelerList) SetInvoiceId(v int64) *ApplyModifyRequestTravelerList {
	s.InvoiceId = &v
	return s
}

func (s *ApplyModifyRequestTravelerList) SetPaymentDepartmentId(v string) *ApplyModifyRequestTravelerList {
	s.PaymentDepartmentId = &v
	return s
}

func (s *ApplyModifyRequestTravelerList) SetPaymentDepartmentName(v string) *ApplyModifyRequestTravelerList {
	s.PaymentDepartmentName = &v
	return s
}

func (s *ApplyModifyRequestTravelerList) SetProjectCode(v string) *ApplyModifyRequestTravelerList {
	s.ProjectCode = &v
	return s
}

func (s *ApplyModifyRequestTravelerList) SetProjectTitle(v string) *ApplyModifyRequestTravelerList {
	s.ProjectTitle = &v
	return s
}

func (s *ApplyModifyRequestTravelerList) SetThirdPartInvoiceId(v string) *ApplyModifyRequestTravelerList {
	s.ThirdPartInvoiceId = &v
	return s
}

func (s *ApplyModifyRequestTravelerList) SetThirdpartCostCenterId(v string) *ApplyModifyRequestTravelerList {
	s.ThirdpartCostCenterId = &v
	return s
}

func (s *ApplyModifyRequestTravelerList) SetThirdpartDepartId(v string) *ApplyModifyRequestTravelerList {
	s.ThirdpartDepartId = &v
	return s
}

func (s *ApplyModifyRequestTravelerList) SetUserId(v string) *ApplyModifyRequestTravelerList {
	s.UserId = &v
	return s
}

func (s *ApplyModifyRequestTravelerList) SetUserName(v string) *ApplyModifyRequestTravelerList {
	s.UserName = &v
	return s
}

func (s *ApplyModifyRequestTravelerList) Validate() error {
	return dara.Validate(s)
}

type ApplyModifyRequestTravelerStandard struct {
	BusinessDiscount          *int32                                              `json:"business_discount,omitempty" xml:"business_discount,omitempty"`
	CarCitySet                []*ApplyModifyRequestTravelerStandardCarCitySet     `json:"car_city_set,omitempty" xml:"car_city_set,omitempty" type:"Repeated"`
	CarStandard               *ApplyModifyRequestTravelerStandardCarStandard      `json:"car_standard,omitempty" xml:"car_standard,omitempty" type:"Struct"`
	EconomyDiscount           *int32                                              `json:"economy_discount,omitempty" xml:"economy_discount,omitempty"`
	FirstDiscount             *int32                                              `json:"first_discount,omitempty" xml:"first_discount,omitempty"`
	FlightCabins              *string                                             `json:"flight_cabins,omitempty" xml:"flight_cabins,omitempty"`
	FlightIntlRuleCode        *int64                                              `json:"flight_intl_rule_code,omitempty" xml:"flight_intl_rule_code,omitempty"`
	FlightRuleCode            *int64                                              `json:"flight_rule_code,omitempty" xml:"flight_rule_code,omitempty"`
	HotelCitys                []*ApplyModifyRequestTravelerStandardHotelCitys     `json:"hotel_citys,omitempty" xml:"hotel_citys,omitempty" type:"Repeated"`
	HotelIntlCitys            []*ApplyModifyRequestTravelerStandardHotelIntlCitys `json:"hotel_intl_citys,omitempty" xml:"hotel_intl_citys,omitempty" type:"Repeated"`
	HotelIntlRuleCode         *int64                                              `json:"hotel_intl_rule_code,omitempty" xml:"hotel_intl_rule_code,omitempty"`
	HotelRuleCode             *int64                                              `json:"hotel_rule_code,omitempty" xml:"hotel_rule_code,omitempty"`
	InternationalFlightCabins *string                                             `json:"international_flight_cabins,omitempty" xml:"international_flight_cabins,omitempty"`
	PremiumEconomyDiscount    *int32                                              `json:"premium_economy_discount,omitempty" xml:"premium_economy_discount,omitempty"`
	ReserveType               *int32                                              `json:"reserve_type,omitempty" xml:"reserve_type,omitempty"`
	TrainRuleCode             *int64                                              `json:"train_rule_code,omitempty" xml:"train_rule_code,omitempty"`
	TrainSeats                *string                                             `json:"train_seats,omitempty" xml:"train_seats,omitempty"`
	UserId                    *string                                             `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

func (s ApplyModifyRequestTravelerStandard) String() string {
	return dara.Prettify(s)
}

func (s ApplyModifyRequestTravelerStandard) GoString() string {
	return s.String()
}

func (s *ApplyModifyRequestTravelerStandard) GetBusinessDiscount() *int32 {
	return s.BusinessDiscount
}

func (s *ApplyModifyRequestTravelerStandard) GetCarCitySet() []*ApplyModifyRequestTravelerStandardCarCitySet {
	return s.CarCitySet
}

func (s *ApplyModifyRequestTravelerStandard) GetCarStandard() *ApplyModifyRequestTravelerStandardCarStandard {
	return s.CarStandard
}

func (s *ApplyModifyRequestTravelerStandard) GetEconomyDiscount() *int32 {
	return s.EconomyDiscount
}

func (s *ApplyModifyRequestTravelerStandard) GetFirstDiscount() *int32 {
	return s.FirstDiscount
}

func (s *ApplyModifyRequestTravelerStandard) GetFlightCabins() *string {
	return s.FlightCabins
}

func (s *ApplyModifyRequestTravelerStandard) GetFlightIntlRuleCode() *int64 {
	return s.FlightIntlRuleCode
}

func (s *ApplyModifyRequestTravelerStandard) GetFlightRuleCode() *int64 {
	return s.FlightRuleCode
}

func (s *ApplyModifyRequestTravelerStandard) GetHotelCitys() []*ApplyModifyRequestTravelerStandardHotelCitys {
	return s.HotelCitys
}

func (s *ApplyModifyRequestTravelerStandard) GetHotelIntlCitys() []*ApplyModifyRequestTravelerStandardHotelIntlCitys {
	return s.HotelIntlCitys
}

func (s *ApplyModifyRequestTravelerStandard) GetHotelIntlRuleCode() *int64 {
	return s.HotelIntlRuleCode
}

func (s *ApplyModifyRequestTravelerStandard) GetHotelRuleCode() *int64 {
	return s.HotelRuleCode
}

func (s *ApplyModifyRequestTravelerStandard) GetInternationalFlightCabins() *string {
	return s.InternationalFlightCabins
}

func (s *ApplyModifyRequestTravelerStandard) GetPremiumEconomyDiscount() *int32 {
	return s.PremiumEconomyDiscount
}

func (s *ApplyModifyRequestTravelerStandard) GetReserveType() *int32 {
	return s.ReserveType
}

func (s *ApplyModifyRequestTravelerStandard) GetTrainRuleCode() *int64 {
	return s.TrainRuleCode
}

func (s *ApplyModifyRequestTravelerStandard) GetTrainSeats() *string {
	return s.TrainSeats
}

func (s *ApplyModifyRequestTravelerStandard) GetUserId() *string {
	return s.UserId
}

func (s *ApplyModifyRequestTravelerStandard) SetBusinessDiscount(v int32) *ApplyModifyRequestTravelerStandard {
	s.BusinessDiscount = &v
	return s
}

func (s *ApplyModifyRequestTravelerStandard) SetCarCitySet(v []*ApplyModifyRequestTravelerStandardCarCitySet) *ApplyModifyRequestTravelerStandard {
	s.CarCitySet = v
	return s
}

func (s *ApplyModifyRequestTravelerStandard) SetCarStandard(v *ApplyModifyRequestTravelerStandardCarStandard) *ApplyModifyRequestTravelerStandard {
	s.CarStandard = v
	return s
}

func (s *ApplyModifyRequestTravelerStandard) SetEconomyDiscount(v int32) *ApplyModifyRequestTravelerStandard {
	s.EconomyDiscount = &v
	return s
}

func (s *ApplyModifyRequestTravelerStandard) SetFirstDiscount(v int32) *ApplyModifyRequestTravelerStandard {
	s.FirstDiscount = &v
	return s
}

func (s *ApplyModifyRequestTravelerStandard) SetFlightCabins(v string) *ApplyModifyRequestTravelerStandard {
	s.FlightCabins = &v
	return s
}

func (s *ApplyModifyRequestTravelerStandard) SetFlightIntlRuleCode(v int64) *ApplyModifyRequestTravelerStandard {
	s.FlightIntlRuleCode = &v
	return s
}

func (s *ApplyModifyRequestTravelerStandard) SetFlightRuleCode(v int64) *ApplyModifyRequestTravelerStandard {
	s.FlightRuleCode = &v
	return s
}

func (s *ApplyModifyRequestTravelerStandard) SetHotelCitys(v []*ApplyModifyRequestTravelerStandardHotelCitys) *ApplyModifyRequestTravelerStandard {
	s.HotelCitys = v
	return s
}

func (s *ApplyModifyRequestTravelerStandard) SetHotelIntlCitys(v []*ApplyModifyRequestTravelerStandardHotelIntlCitys) *ApplyModifyRequestTravelerStandard {
	s.HotelIntlCitys = v
	return s
}

func (s *ApplyModifyRequestTravelerStandard) SetHotelIntlRuleCode(v int64) *ApplyModifyRequestTravelerStandard {
	s.HotelIntlRuleCode = &v
	return s
}

func (s *ApplyModifyRequestTravelerStandard) SetHotelRuleCode(v int64) *ApplyModifyRequestTravelerStandard {
	s.HotelRuleCode = &v
	return s
}

func (s *ApplyModifyRequestTravelerStandard) SetInternationalFlightCabins(v string) *ApplyModifyRequestTravelerStandard {
	s.InternationalFlightCabins = &v
	return s
}

func (s *ApplyModifyRequestTravelerStandard) SetPremiumEconomyDiscount(v int32) *ApplyModifyRequestTravelerStandard {
	s.PremiumEconomyDiscount = &v
	return s
}

func (s *ApplyModifyRequestTravelerStandard) SetReserveType(v int32) *ApplyModifyRequestTravelerStandard {
	s.ReserveType = &v
	return s
}

func (s *ApplyModifyRequestTravelerStandard) SetTrainRuleCode(v int64) *ApplyModifyRequestTravelerStandard {
	s.TrainRuleCode = &v
	return s
}

func (s *ApplyModifyRequestTravelerStandard) SetTrainSeats(v string) *ApplyModifyRequestTravelerStandard {
	s.TrainSeats = &v
	return s
}

func (s *ApplyModifyRequestTravelerStandard) SetUserId(v string) *ApplyModifyRequestTravelerStandard {
	s.UserId = &v
	return s
}

func (s *ApplyModifyRequestTravelerStandard) Validate() error {
	if s.CarCitySet != nil {
		for _, item := range s.CarCitySet {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.CarStandard != nil {
		if err := s.CarStandard.Validate(); err != nil {
			return err
		}
	}
	if s.HotelCitys != nil {
		for _, item := range s.HotelCitys {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.HotelIntlCitys != nil {
		for _, item := range s.HotelIntlCitys {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ApplyModifyRequestTravelerStandardCarCitySet struct {
	// This parameter is required.
	CityCode *string `json:"city_code,omitempty" xml:"city_code,omitempty"`
	// This parameter is required.
	CityName *string `json:"city_name,omitempty" xml:"city_name,omitempty"`
}

func (s ApplyModifyRequestTravelerStandardCarCitySet) String() string {
	return dara.Prettify(s)
}

func (s ApplyModifyRequestTravelerStandardCarCitySet) GoString() string {
	return s.String()
}

func (s *ApplyModifyRequestTravelerStandardCarCitySet) GetCityCode() *string {
	return s.CityCode
}

func (s *ApplyModifyRequestTravelerStandardCarCitySet) GetCityName() *string {
	return s.CityName
}

func (s *ApplyModifyRequestTravelerStandardCarCitySet) SetCityCode(v string) *ApplyModifyRequestTravelerStandardCarCitySet {
	s.CityCode = &v
	return s
}

func (s *ApplyModifyRequestTravelerStandardCarCitySet) SetCityName(v string) *ApplyModifyRequestTravelerStandardCarCitySet {
	s.CityName = &v
	return s
}

func (s *ApplyModifyRequestTravelerStandardCarCitySet) Validate() error {
	return dara.Validate(s)
}

type ApplyModifyRequestTravelerStandardCarStandard struct {
	BookAllowInfo         *ApplyModifyRequestTravelerStandardCarStandardBookAllowInfo         `json:"book_allow_info,omitempty" xml:"book_allow_info,omitempty" type:"Struct"`
	CarHelper             *ApplyModifyRequestTravelerStandardCarStandardCarHelper             `json:"car_helper,omitempty" xml:"car_helper,omitempty" type:"Struct"`
	CarTimeControl        *ApplyModifyRequestTravelerStandardCarStandardCarTimeControl        `json:"car_time_control,omitempty" xml:"car_time_control,omitempty" type:"Struct"`
	CityControlInfo       *ApplyModifyRequestTravelerStandardCarStandardCityControlInfo       `json:"city_control_info,omitempty" xml:"city_control_info,omitempty" type:"Struct"`
	CrossCityInfo         *ApplyModifyRequestTravelerStandardCarStandardCrossCityInfo         `json:"cross_city_info,omitempty" xml:"cross_city_info,omitempty" type:"Struct"`
	ElectronicFenceInfo   *ApplyModifyRequestTravelerStandardCarStandardElectronicFenceInfo   `json:"electronic_fence_info,omitempty" xml:"electronic_fence_info,omitempty" type:"Struct"`
	LevelCodes            *string                                                             `json:"level_codes,omitempty" xml:"level_codes,omitempty"`
	ModifyDestinationInfo *ApplyModifyRequestTravelerStandardCarStandardModifyDestinationInfo `json:"modify_destination_info,omitempty" xml:"modify_destination_info,omitempty" type:"Struct"`
	TimesTotal            *int32                                                              `json:"times_total,omitempty" xml:"times_total,omitempty"`
	TimesType             *int32                                                              `json:"times_type,omitempty" xml:"times_type,omitempty"`
}

func (s ApplyModifyRequestTravelerStandardCarStandard) String() string {
	return dara.Prettify(s)
}

func (s ApplyModifyRequestTravelerStandardCarStandard) GoString() string {
	return s.String()
}

func (s *ApplyModifyRequestTravelerStandardCarStandard) GetBookAllowInfo() *ApplyModifyRequestTravelerStandardCarStandardBookAllowInfo {
	return s.BookAllowInfo
}

func (s *ApplyModifyRequestTravelerStandardCarStandard) GetCarHelper() *ApplyModifyRequestTravelerStandardCarStandardCarHelper {
	return s.CarHelper
}

func (s *ApplyModifyRequestTravelerStandardCarStandard) GetCarTimeControl() *ApplyModifyRequestTravelerStandardCarStandardCarTimeControl {
	return s.CarTimeControl
}

func (s *ApplyModifyRequestTravelerStandardCarStandard) GetCityControlInfo() *ApplyModifyRequestTravelerStandardCarStandardCityControlInfo {
	return s.CityControlInfo
}

func (s *ApplyModifyRequestTravelerStandardCarStandard) GetCrossCityInfo() *ApplyModifyRequestTravelerStandardCarStandardCrossCityInfo {
	return s.CrossCityInfo
}

func (s *ApplyModifyRequestTravelerStandardCarStandard) GetElectronicFenceInfo() *ApplyModifyRequestTravelerStandardCarStandardElectronicFenceInfo {
	return s.ElectronicFenceInfo
}

func (s *ApplyModifyRequestTravelerStandardCarStandard) GetLevelCodes() *string {
	return s.LevelCodes
}

func (s *ApplyModifyRequestTravelerStandardCarStandard) GetModifyDestinationInfo() *ApplyModifyRequestTravelerStandardCarStandardModifyDestinationInfo {
	return s.ModifyDestinationInfo
}

func (s *ApplyModifyRequestTravelerStandardCarStandard) GetTimesTotal() *int32 {
	return s.TimesTotal
}

func (s *ApplyModifyRequestTravelerStandardCarStandard) GetTimesType() *int32 {
	return s.TimesType
}

func (s *ApplyModifyRequestTravelerStandardCarStandard) SetBookAllowInfo(v *ApplyModifyRequestTravelerStandardCarStandardBookAllowInfo) *ApplyModifyRequestTravelerStandardCarStandard {
	s.BookAllowInfo = v
	return s
}

func (s *ApplyModifyRequestTravelerStandardCarStandard) SetCarHelper(v *ApplyModifyRequestTravelerStandardCarStandardCarHelper) *ApplyModifyRequestTravelerStandardCarStandard {
	s.CarHelper = v
	return s
}

func (s *ApplyModifyRequestTravelerStandardCarStandard) SetCarTimeControl(v *ApplyModifyRequestTravelerStandardCarStandardCarTimeControl) *ApplyModifyRequestTravelerStandardCarStandard {
	s.CarTimeControl = v
	return s
}

func (s *ApplyModifyRequestTravelerStandardCarStandard) SetCityControlInfo(v *ApplyModifyRequestTravelerStandardCarStandardCityControlInfo) *ApplyModifyRequestTravelerStandardCarStandard {
	s.CityControlInfo = v
	return s
}

func (s *ApplyModifyRequestTravelerStandardCarStandard) SetCrossCityInfo(v *ApplyModifyRequestTravelerStandardCarStandardCrossCityInfo) *ApplyModifyRequestTravelerStandardCarStandard {
	s.CrossCityInfo = v
	return s
}

func (s *ApplyModifyRequestTravelerStandardCarStandard) SetElectronicFenceInfo(v *ApplyModifyRequestTravelerStandardCarStandardElectronicFenceInfo) *ApplyModifyRequestTravelerStandardCarStandard {
	s.ElectronicFenceInfo = v
	return s
}

func (s *ApplyModifyRequestTravelerStandardCarStandard) SetLevelCodes(v string) *ApplyModifyRequestTravelerStandardCarStandard {
	s.LevelCodes = &v
	return s
}

func (s *ApplyModifyRequestTravelerStandardCarStandard) SetModifyDestinationInfo(v *ApplyModifyRequestTravelerStandardCarStandardModifyDestinationInfo) *ApplyModifyRequestTravelerStandardCarStandard {
	s.ModifyDestinationInfo = v
	return s
}

func (s *ApplyModifyRequestTravelerStandardCarStandard) SetTimesTotal(v int32) *ApplyModifyRequestTravelerStandardCarStandard {
	s.TimesTotal = &v
	return s
}

func (s *ApplyModifyRequestTravelerStandardCarStandard) SetTimesType(v int32) *ApplyModifyRequestTravelerStandardCarStandard {
	s.TimesType = &v
	return s
}

func (s *ApplyModifyRequestTravelerStandardCarStandard) Validate() error {
	if s.BookAllowInfo != nil {
		if err := s.BookAllowInfo.Validate(); err != nil {
			return err
		}
	}
	if s.CarHelper != nil {
		if err := s.CarHelper.Validate(); err != nil {
			return err
		}
	}
	if s.CarTimeControl != nil {
		if err := s.CarTimeControl.Validate(); err != nil {
			return err
		}
	}
	if s.CityControlInfo != nil {
		if err := s.CityControlInfo.Validate(); err != nil {
			return err
		}
	}
	if s.CrossCityInfo != nil {
		if err := s.CrossCityInfo.Validate(); err != nil {
			return err
		}
	}
	if s.ElectronicFenceInfo != nil {
		if err := s.ElectronicFenceInfo.Validate(); err != nil {
			return err
		}
	}
	if s.ModifyDestinationInfo != nil {
		if err := s.ModifyDestinationInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ApplyModifyRequestTravelerStandardCarStandardBookAllowInfo struct {
	// This parameter is required.
	BookAllow *bool `json:"book_allow,omitempty" xml:"book_allow,omitempty"`
}

func (s ApplyModifyRequestTravelerStandardCarStandardBookAllowInfo) String() string {
	return dara.Prettify(s)
}

func (s ApplyModifyRequestTravelerStandardCarStandardBookAllowInfo) GoString() string {
	return s.String()
}

func (s *ApplyModifyRequestTravelerStandardCarStandardBookAllowInfo) GetBookAllow() *bool {
	return s.BookAllow
}

func (s *ApplyModifyRequestTravelerStandardCarStandardBookAllowInfo) SetBookAllow(v bool) *ApplyModifyRequestTravelerStandardCarStandardBookAllowInfo {
	s.BookAllow = &v
	return s
}

func (s *ApplyModifyRequestTravelerStandardCarStandardBookAllowInfo) Validate() error {
	return dara.Validate(s)
}

type ApplyModifyRequestTravelerStandardCarStandardCarHelper struct {
	// This parameter is required.
	CarHelperType *string `json:"car_helper_type,omitempty" xml:"car_helper_type,omitempty"`
}

func (s ApplyModifyRequestTravelerStandardCarStandardCarHelper) String() string {
	return dara.Prettify(s)
}

func (s ApplyModifyRequestTravelerStandardCarStandardCarHelper) GoString() string {
	return s.String()
}

func (s *ApplyModifyRequestTravelerStandardCarStandardCarHelper) GetCarHelperType() *string {
	return s.CarHelperType
}

func (s *ApplyModifyRequestTravelerStandardCarStandardCarHelper) SetCarHelperType(v string) *ApplyModifyRequestTravelerStandardCarStandardCarHelper {
	s.CarHelperType = &v
	return s
}

func (s *ApplyModifyRequestTravelerStandardCarStandardCarHelper) Validate() error {
	return dara.Validate(s)
}

type ApplyModifyRequestTravelerStandardCarStandardCarTimeControl struct {
	// This parameter is required.
	TimeLimit []*ApplyModifyRequestTravelerStandardCarStandardCarTimeControlTimeLimit `json:"time_limit,omitempty" xml:"time_limit,omitempty" type:"Repeated"`
	// This parameter is required.
	TimeSwitch *bool `json:"time_switch,omitempty" xml:"time_switch,omitempty"`
}

func (s ApplyModifyRequestTravelerStandardCarStandardCarTimeControl) String() string {
	return dara.Prettify(s)
}

func (s ApplyModifyRequestTravelerStandardCarStandardCarTimeControl) GoString() string {
	return s.String()
}

func (s *ApplyModifyRequestTravelerStandardCarStandardCarTimeControl) GetTimeLimit() []*ApplyModifyRequestTravelerStandardCarStandardCarTimeControlTimeLimit {
	return s.TimeLimit
}

func (s *ApplyModifyRequestTravelerStandardCarStandardCarTimeControl) GetTimeSwitch() *bool {
	return s.TimeSwitch
}

func (s *ApplyModifyRequestTravelerStandardCarStandardCarTimeControl) SetTimeLimit(v []*ApplyModifyRequestTravelerStandardCarStandardCarTimeControlTimeLimit) *ApplyModifyRequestTravelerStandardCarStandardCarTimeControl {
	s.TimeLimit = v
	return s
}

func (s *ApplyModifyRequestTravelerStandardCarStandardCarTimeControl) SetTimeSwitch(v bool) *ApplyModifyRequestTravelerStandardCarStandardCarTimeControl {
	s.TimeSwitch = &v
	return s
}

func (s *ApplyModifyRequestTravelerStandardCarStandardCarTimeControl) Validate() error {
	if s.TimeLimit != nil {
		for _, item := range s.TimeLimit {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ApplyModifyRequestTravelerStandardCarStandardCarTimeControlTimeLimit struct {
	// This parameter is required.
	EndTime *string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// This parameter is required.
	StartTime *string `json:"start_time,omitempty" xml:"start_time,omitempty"`
}

func (s ApplyModifyRequestTravelerStandardCarStandardCarTimeControlTimeLimit) String() string {
	return dara.Prettify(s)
}

func (s ApplyModifyRequestTravelerStandardCarStandardCarTimeControlTimeLimit) GoString() string {
	return s.String()
}

func (s *ApplyModifyRequestTravelerStandardCarStandardCarTimeControlTimeLimit) GetEndTime() *string {
	return s.EndTime
}

func (s *ApplyModifyRequestTravelerStandardCarStandardCarTimeControlTimeLimit) GetStartTime() *string {
	return s.StartTime
}

func (s *ApplyModifyRequestTravelerStandardCarStandardCarTimeControlTimeLimit) SetEndTime(v string) *ApplyModifyRequestTravelerStandardCarStandardCarTimeControlTimeLimit {
	s.EndTime = &v
	return s
}

func (s *ApplyModifyRequestTravelerStandardCarStandardCarTimeControlTimeLimit) SetStartTime(v string) *ApplyModifyRequestTravelerStandardCarStandardCarTimeControlTimeLimit {
	s.StartTime = &v
	return s
}

func (s *ApplyModifyRequestTravelerStandardCarStandardCarTimeControlTimeLimit) Validate() error {
	return dara.Validate(s)
}

type ApplyModifyRequestTravelerStandardCarStandardCityControlInfo struct {
	// This parameter is required.
	CityControlType *int32 `json:"city_control_type,omitempty" xml:"city_control_type,omitempty"`
	// This parameter is required.
	CityInfos []*ApplyModifyRequestTravelerStandardCarStandardCityControlInfoCityInfos `json:"city_infos,omitempty" xml:"city_infos,omitempty" type:"Repeated"`
}

func (s ApplyModifyRequestTravelerStandardCarStandardCityControlInfo) String() string {
	return dara.Prettify(s)
}

func (s ApplyModifyRequestTravelerStandardCarStandardCityControlInfo) GoString() string {
	return s.String()
}

func (s *ApplyModifyRequestTravelerStandardCarStandardCityControlInfo) GetCityControlType() *int32 {
	return s.CityControlType
}

func (s *ApplyModifyRequestTravelerStandardCarStandardCityControlInfo) GetCityInfos() []*ApplyModifyRequestTravelerStandardCarStandardCityControlInfoCityInfos {
	return s.CityInfos
}

func (s *ApplyModifyRequestTravelerStandardCarStandardCityControlInfo) SetCityControlType(v int32) *ApplyModifyRequestTravelerStandardCarStandardCityControlInfo {
	s.CityControlType = &v
	return s
}

func (s *ApplyModifyRequestTravelerStandardCarStandardCityControlInfo) SetCityInfos(v []*ApplyModifyRequestTravelerStandardCarStandardCityControlInfoCityInfos) *ApplyModifyRequestTravelerStandardCarStandardCityControlInfo {
	s.CityInfos = v
	return s
}

func (s *ApplyModifyRequestTravelerStandardCarStandardCityControlInfo) Validate() error {
	if s.CityInfos != nil {
		for _, item := range s.CityInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ApplyModifyRequestTravelerStandardCarStandardCityControlInfoCityInfos struct {
	// This parameter is required.
	Adcode *string `json:"adcode,omitempty" xml:"adcode,omitempty"`
	// This parameter is required.
	CityCode *string `json:"city_code,omitempty" xml:"city_code,omitempty"`
	// This parameter is required.
	CityName *string `json:"city_name,omitempty" xml:"city_name,omitempty"`
}

func (s ApplyModifyRequestTravelerStandardCarStandardCityControlInfoCityInfos) String() string {
	return dara.Prettify(s)
}

func (s ApplyModifyRequestTravelerStandardCarStandardCityControlInfoCityInfos) GoString() string {
	return s.String()
}

func (s *ApplyModifyRequestTravelerStandardCarStandardCityControlInfoCityInfos) GetAdcode() *string {
	return s.Adcode
}

func (s *ApplyModifyRequestTravelerStandardCarStandardCityControlInfoCityInfos) GetCityCode() *string {
	return s.CityCode
}

func (s *ApplyModifyRequestTravelerStandardCarStandardCityControlInfoCityInfos) GetCityName() *string {
	return s.CityName
}

func (s *ApplyModifyRequestTravelerStandardCarStandardCityControlInfoCityInfos) SetAdcode(v string) *ApplyModifyRequestTravelerStandardCarStandardCityControlInfoCityInfos {
	s.Adcode = &v
	return s
}

func (s *ApplyModifyRequestTravelerStandardCarStandardCityControlInfoCityInfos) SetCityCode(v string) *ApplyModifyRequestTravelerStandardCarStandardCityControlInfoCityInfos {
	s.CityCode = &v
	return s
}

func (s *ApplyModifyRequestTravelerStandardCarStandardCityControlInfoCityInfos) SetCityName(v string) *ApplyModifyRequestTravelerStandardCarStandardCityControlInfoCityInfos {
	s.CityName = &v
	return s
}

func (s *ApplyModifyRequestTravelerStandardCarStandardCityControlInfoCityInfos) Validate() error {
	return dara.Validate(s)
}

type ApplyModifyRequestTravelerStandardCarStandardCrossCityInfo struct {
	CrossCityList []*ApplyModifyRequestTravelerStandardCarStandardCrossCityInfoCrossCityList `json:"cross_city_list,omitempty" xml:"cross_city_list,omitempty" type:"Repeated"`
	// This parameter is required.
	CrossCityType *int32 `json:"cross_city_type,omitempty" xml:"cross_city_type,omitempty"`
}

func (s ApplyModifyRequestTravelerStandardCarStandardCrossCityInfo) String() string {
	return dara.Prettify(s)
}

func (s ApplyModifyRequestTravelerStandardCarStandardCrossCityInfo) GoString() string {
	return s.String()
}

func (s *ApplyModifyRequestTravelerStandardCarStandardCrossCityInfo) GetCrossCityList() []*ApplyModifyRequestTravelerStandardCarStandardCrossCityInfoCrossCityList {
	return s.CrossCityList
}

func (s *ApplyModifyRequestTravelerStandardCarStandardCrossCityInfo) GetCrossCityType() *int32 {
	return s.CrossCityType
}

func (s *ApplyModifyRequestTravelerStandardCarStandardCrossCityInfo) SetCrossCityList(v []*ApplyModifyRequestTravelerStandardCarStandardCrossCityInfoCrossCityList) *ApplyModifyRequestTravelerStandardCarStandardCrossCityInfo {
	s.CrossCityList = v
	return s
}

func (s *ApplyModifyRequestTravelerStandardCarStandardCrossCityInfo) SetCrossCityType(v int32) *ApplyModifyRequestTravelerStandardCarStandardCrossCityInfo {
	s.CrossCityType = &v
	return s
}

func (s *ApplyModifyRequestTravelerStandardCarStandardCrossCityInfo) Validate() error {
	if s.CrossCityList != nil {
		for _, item := range s.CrossCityList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ApplyModifyRequestTravelerStandardCarStandardCrossCityInfoCrossCityList struct {
	// This parameter is required.
	FromAdcode *string `json:"from_adcode,omitempty" xml:"from_adcode,omitempty"`
	// This parameter is required.
	FromCityCode *string `json:"from_city_code,omitempty" xml:"from_city_code,omitempty"`
	// This parameter is required.
	FromCityName *string `json:"from_city_name,omitempty" xml:"from_city_name,omitempty"`
	// This parameter is required.
	ToAdcode *string `json:"to_adcode,omitempty" xml:"to_adcode,omitempty"`
	// This parameter is required.
	ToCityCode *string `json:"to_city_code,omitempty" xml:"to_city_code,omitempty"`
	// This parameter is required.
	ToCityName *string `json:"to_city_name,omitempty" xml:"to_city_name,omitempty"`
}

func (s ApplyModifyRequestTravelerStandardCarStandardCrossCityInfoCrossCityList) String() string {
	return dara.Prettify(s)
}

func (s ApplyModifyRequestTravelerStandardCarStandardCrossCityInfoCrossCityList) GoString() string {
	return s.String()
}

func (s *ApplyModifyRequestTravelerStandardCarStandardCrossCityInfoCrossCityList) GetFromAdcode() *string {
	return s.FromAdcode
}

func (s *ApplyModifyRequestTravelerStandardCarStandardCrossCityInfoCrossCityList) GetFromCityCode() *string {
	return s.FromCityCode
}

func (s *ApplyModifyRequestTravelerStandardCarStandardCrossCityInfoCrossCityList) GetFromCityName() *string {
	return s.FromCityName
}

func (s *ApplyModifyRequestTravelerStandardCarStandardCrossCityInfoCrossCityList) GetToAdcode() *string {
	return s.ToAdcode
}

func (s *ApplyModifyRequestTravelerStandardCarStandardCrossCityInfoCrossCityList) GetToCityCode() *string {
	return s.ToCityCode
}

func (s *ApplyModifyRequestTravelerStandardCarStandardCrossCityInfoCrossCityList) GetToCityName() *string {
	return s.ToCityName
}

func (s *ApplyModifyRequestTravelerStandardCarStandardCrossCityInfoCrossCityList) SetFromAdcode(v string) *ApplyModifyRequestTravelerStandardCarStandardCrossCityInfoCrossCityList {
	s.FromAdcode = &v
	return s
}

func (s *ApplyModifyRequestTravelerStandardCarStandardCrossCityInfoCrossCityList) SetFromCityCode(v string) *ApplyModifyRequestTravelerStandardCarStandardCrossCityInfoCrossCityList {
	s.FromCityCode = &v
	return s
}

func (s *ApplyModifyRequestTravelerStandardCarStandardCrossCityInfoCrossCityList) SetFromCityName(v string) *ApplyModifyRequestTravelerStandardCarStandardCrossCityInfoCrossCityList {
	s.FromCityName = &v
	return s
}

func (s *ApplyModifyRequestTravelerStandardCarStandardCrossCityInfoCrossCityList) SetToAdcode(v string) *ApplyModifyRequestTravelerStandardCarStandardCrossCityInfoCrossCityList {
	s.ToAdcode = &v
	return s
}

func (s *ApplyModifyRequestTravelerStandardCarStandardCrossCityInfoCrossCityList) SetToCityCode(v string) *ApplyModifyRequestTravelerStandardCarStandardCrossCityInfoCrossCityList {
	s.ToCityCode = &v
	return s
}

func (s *ApplyModifyRequestTravelerStandardCarStandardCrossCityInfoCrossCityList) SetToCityName(v string) *ApplyModifyRequestTravelerStandardCarStandardCrossCityInfoCrossCityList {
	s.ToCityName = &v
	return s
}

func (s *ApplyModifyRequestTravelerStandardCarStandardCrossCityInfoCrossCityList) Validate() error {
	return dara.Validate(s)
}

type ApplyModifyRequestTravelerStandardCarStandardElectronicFenceInfo struct {
	// This parameter is required.
	ElectronicFenceLocationsFrom []*ApplyModifyRequestTravelerStandardCarStandardElectronicFenceInfoElectronicFenceLocationsFrom `json:"electronic_fence_locations_from,omitempty" xml:"electronic_fence_locations_from,omitempty" type:"Repeated"`
	// This parameter is required.
	ElectronicFenceLocationsTo []*ApplyModifyRequestTravelerStandardCarStandardElectronicFenceInfoElectronicFenceLocationsTo `json:"electronic_fence_locations_to,omitempty" xml:"electronic_fence_locations_to,omitempty" type:"Repeated"`
	// This parameter is required.
	ElectronicFenceType *int32 `json:"electronic_fence_type,omitempty" xml:"electronic_fence_type,omitempty"`
}

func (s ApplyModifyRequestTravelerStandardCarStandardElectronicFenceInfo) String() string {
	return dara.Prettify(s)
}

func (s ApplyModifyRequestTravelerStandardCarStandardElectronicFenceInfo) GoString() string {
	return s.String()
}

func (s *ApplyModifyRequestTravelerStandardCarStandardElectronicFenceInfo) GetElectronicFenceLocationsFrom() []*ApplyModifyRequestTravelerStandardCarStandardElectronicFenceInfoElectronicFenceLocationsFrom {
	return s.ElectronicFenceLocationsFrom
}

func (s *ApplyModifyRequestTravelerStandardCarStandardElectronicFenceInfo) GetElectronicFenceLocationsTo() []*ApplyModifyRequestTravelerStandardCarStandardElectronicFenceInfoElectronicFenceLocationsTo {
	return s.ElectronicFenceLocationsTo
}

func (s *ApplyModifyRequestTravelerStandardCarStandardElectronicFenceInfo) GetElectronicFenceType() *int32 {
	return s.ElectronicFenceType
}

func (s *ApplyModifyRequestTravelerStandardCarStandardElectronicFenceInfo) SetElectronicFenceLocationsFrom(v []*ApplyModifyRequestTravelerStandardCarStandardElectronicFenceInfoElectronicFenceLocationsFrom) *ApplyModifyRequestTravelerStandardCarStandardElectronicFenceInfo {
	s.ElectronicFenceLocationsFrom = v
	return s
}

func (s *ApplyModifyRequestTravelerStandardCarStandardElectronicFenceInfo) SetElectronicFenceLocationsTo(v []*ApplyModifyRequestTravelerStandardCarStandardElectronicFenceInfoElectronicFenceLocationsTo) *ApplyModifyRequestTravelerStandardCarStandardElectronicFenceInfo {
	s.ElectronicFenceLocationsTo = v
	return s
}

func (s *ApplyModifyRequestTravelerStandardCarStandardElectronicFenceInfo) SetElectronicFenceType(v int32) *ApplyModifyRequestTravelerStandardCarStandardElectronicFenceInfo {
	s.ElectronicFenceType = &v
	return s
}

func (s *ApplyModifyRequestTravelerStandardCarStandardElectronicFenceInfo) Validate() error {
	if s.ElectronicFenceLocationsFrom != nil {
		for _, item := range s.ElectronicFenceLocationsFrom {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ElectronicFenceLocationsTo != nil {
		for _, item := range s.ElectronicFenceLocationsTo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ApplyModifyRequestTravelerStandardCarStandardElectronicFenceInfoElectronicFenceLocationsFrom struct {
	// This parameter is required.
	Address *string `json:"address,omitempty" xml:"address,omitempty"`
	// This parameter is required.
	Latitude *string `json:"latitude,omitempty" xml:"latitude,omitempty"`
	// This parameter is required.
	Longitude *string `json:"longitude,omitempty" xml:"longitude,omitempty"`
	// This parameter is required.
	Radius *int32 `json:"radius,omitempty" xml:"radius,omitempty"`
}

func (s ApplyModifyRequestTravelerStandardCarStandardElectronicFenceInfoElectronicFenceLocationsFrom) String() string {
	return dara.Prettify(s)
}

func (s ApplyModifyRequestTravelerStandardCarStandardElectronicFenceInfoElectronicFenceLocationsFrom) GoString() string {
	return s.String()
}

func (s *ApplyModifyRequestTravelerStandardCarStandardElectronicFenceInfoElectronicFenceLocationsFrom) GetAddress() *string {
	return s.Address
}

func (s *ApplyModifyRequestTravelerStandardCarStandardElectronicFenceInfoElectronicFenceLocationsFrom) GetLatitude() *string {
	return s.Latitude
}

func (s *ApplyModifyRequestTravelerStandardCarStandardElectronicFenceInfoElectronicFenceLocationsFrom) GetLongitude() *string {
	return s.Longitude
}

func (s *ApplyModifyRequestTravelerStandardCarStandardElectronicFenceInfoElectronicFenceLocationsFrom) GetRadius() *int32 {
	return s.Radius
}

func (s *ApplyModifyRequestTravelerStandardCarStandardElectronicFenceInfoElectronicFenceLocationsFrom) SetAddress(v string) *ApplyModifyRequestTravelerStandardCarStandardElectronicFenceInfoElectronicFenceLocationsFrom {
	s.Address = &v
	return s
}

func (s *ApplyModifyRequestTravelerStandardCarStandardElectronicFenceInfoElectronicFenceLocationsFrom) SetLatitude(v string) *ApplyModifyRequestTravelerStandardCarStandardElectronicFenceInfoElectronicFenceLocationsFrom {
	s.Latitude = &v
	return s
}

func (s *ApplyModifyRequestTravelerStandardCarStandardElectronicFenceInfoElectronicFenceLocationsFrom) SetLongitude(v string) *ApplyModifyRequestTravelerStandardCarStandardElectronicFenceInfoElectronicFenceLocationsFrom {
	s.Longitude = &v
	return s
}

func (s *ApplyModifyRequestTravelerStandardCarStandardElectronicFenceInfoElectronicFenceLocationsFrom) SetRadius(v int32) *ApplyModifyRequestTravelerStandardCarStandardElectronicFenceInfoElectronicFenceLocationsFrom {
	s.Radius = &v
	return s
}

func (s *ApplyModifyRequestTravelerStandardCarStandardElectronicFenceInfoElectronicFenceLocationsFrom) Validate() error {
	return dara.Validate(s)
}

type ApplyModifyRequestTravelerStandardCarStandardElectronicFenceInfoElectronicFenceLocationsTo struct {
	// This parameter is required.
	Address *string `json:"address,omitempty" xml:"address,omitempty"`
	// This parameter is required.
	Latitude *string `json:"latitude,omitempty" xml:"latitude,omitempty"`
	// This parameter is required.
	Longitude *string `json:"longitude,omitempty" xml:"longitude,omitempty"`
	// This parameter is required.
	Radius *int32 `json:"radius,omitempty" xml:"radius,omitempty"`
}

func (s ApplyModifyRequestTravelerStandardCarStandardElectronicFenceInfoElectronicFenceLocationsTo) String() string {
	return dara.Prettify(s)
}

func (s ApplyModifyRequestTravelerStandardCarStandardElectronicFenceInfoElectronicFenceLocationsTo) GoString() string {
	return s.String()
}

func (s *ApplyModifyRequestTravelerStandardCarStandardElectronicFenceInfoElectronicFenceLocationsTo) GetAddress() *string {
	return s.Address
}

func (s *ApplyModifyRequestTravelerStandardCarStandardElectronicFenceInfoElectronicFenceLocationsTo) GetLatitude() *string {
	return s.Latitude
}

func (s *ApplyModifyRequestTravelerStandardCarStandardElectronicFenceInfoElectronicFenceLocationsTo) GetLongitude() *string {
	return s.Longitude
}

func (s *ApplyModifyRequestTravelerStandardCarStandardElectronicFenceInfoElectronicFenceLocationsTo) GetRadius() *int32 {
	return s.Radius
}

func (s *ApplyModifyRequestTravelerStandardCarStandardElectronicFenceInfoElectronicFenceLocationsTo) SetAddress(v string) *ApplyModifyRequestTravelerStandardCarStandardElectronicFenceInfoElectronicFenceLocationsTo {
	s.Address = &v
	return s
}

func (s *ApplyModifyRequestTravelerStandardCarStandardElectronicFenceInfoElectronicFenceLocationsTo) SetLatitude(v string) *ApplyModifyRequestTravelerStandardCarStandardElectronicFenceInfoElectronicFenceLocationsTo {
	s.Latitude = &v
	return s
}

func (s *ApplyModifyRequestTravelerStandardCarStandardElectronicFenceInfoElectronicFenceLocationsTo) SetLongitude(v string) *ApplyModifyRequestTravelerStandardCarStandardElectronicFenceInfoElectronicFenceLocationsTo {
	s.Longitude = &v
	return s
}

func (s *ApplyModifyRequestTravelerStandardCarStandardElectronicFenceInfoElectronicFenceLocationsTo) SetRadius(v int32) *ApplyModifyRequestTravelerStandardCarStandardElectronicFenceInfoElectronicFenceLocationsTo {
	s.Radius = &v
	return s
}

func (s *ApplyModifyRequestTravelerStandardCarStandardElectronicFenceInfoElectronicFenceLocationsTo) Validate() error {
	return dara.Validate(s)
}

type ApplyModifyRequestTravelerStandardCarStandardModifyDestinationInfo struct {
	// This parameter is required.
	ModifyDestination *bool `json:"modify_destination,omitempty" xml:"modify_destination,omitempty"`
}

func (s ApplyModifyRequestTravelerStandardCarStandardModifyDestinationInfo) String() string {
	return dara.Prettify(s)
}

func (s ApplyModifyRequestTravelerStandardCarStandardModifyDestinationInfo) GoString() string {
	return s.String()
}

func (s *ApplyModifyRequestTravelerStandardCarStandardModifyDestinationInfo) GetModifyDestination() *bool {
	return s.ModifyDestination
}

func (s *ApplyModifyRequestTravelerStandardCarStandardModifyDestinationInfo) SetModifyDestination(v bool) *ApplyModifyRequestTravelerStandardCarStandardModifyDestinationInfo {
	s.ModifyDestination = &v
	return s
}

func (s *ApplyModifyRequestTravelerStandardCarStandardModifyDestinationInfo) Validate() error {
	return dara.Validate(s)
}

type ApplyModifyRequestTravelerStandardHotelCitys struct {
	CityCode *string `json:"city_code,omitempty" xml:"city_code,omitempty"`
	CityName *string `json:"city_name,omitempty" xml:"city_name,omitempty"`
	Fee      *int64  `json:"fee,omitempty" xml:"fee,omitempty"`
}

func (s ApplyModifyRequestTravelerStandardHotelCitys) String() string {
	return dara.Prettify(s)
}

func (s ApplyModifyRequestTravelerStandardHotelCitys) GoString() string {
	return s.String()
}

func (s *ApplyModifyRequestTravelerStandardHotelCitys) GetCityCode() *string {
	return s.CityCode
}

func (s *ApplyModifyRequestTravelerStandardHotelCitys) GetCityName() *string {
	return s.CityName
}

func (s *ApplyModifyRequestTravelerStandardHotelCitys) GetFee() *int64 {
	return s.Fee
}

func (s *ApplyModifyRequestTravelerStandardHotelCitys) SetCityCode(v string) *ApplyModifyRequestTravelerStandardHotelCitys {
	s.CityCode = &v
	return s
}

func (s *ApplyModifyRequestTravelerStandardHotelCitys) SetCityName(v string) *ApplyModifyRequestTravelerStandardHotelCitys {
	s.CityName = &v
	return s
}

func (s *ApplyModifyRequestTravelerStandardHotelCitys) SetFee(v int64) *ApplyModifyRequestTravelerStandardHotelCitys {
	s.Fee = &v
	return s
}

func (s *ApplyModifyRequestTravelerStandardHotelCitys) Validate() error {
	return dara.Validate(s)
}

type ApplyModifyRequestTravelerStandardHotelIntlCitys struct {
	CityCode *string `json:"city_code,omitempty" xml:"city_code,omitempty"`
	CityName *string `json:"city_name,omitempty" xml:"city_name,omitempty"`
	Fee      *int64  `json:"fee,omitempty" xml:"fee,omitempty"`
}

func (s ApplyModifyRequestTravelerStandardHotelIntlCitys) String() string {
	return dara.Prettify(s)
}

func (s ApplyModifyRequestTravelerStandardHotelIntlCitys) GoString() string {
	return s.String()
}

func (s *ApplyModifyRequestTravelerStandardHotelIntlCitys) GetCityCode() *string {
	return s.CityCode
}

func (s *ApplyModifyRequestTravelerStandardHotelIntlCitys) GetCityName() *string {
	return s.CityName
}

func (s *ApplyModifyRequestTravelerStandardHotelIntlCitys) GetFee() *int64 {
	return s.Fee
}

func (s *ApplyModifyRequestTravelerStandardHotelIntlCitys) SetCityCode(v string) *ApplyModifyRequestTravelerStandardHotelIntlCitys {
	s.CityCode = &v
	return s
}

func (s *ApplyModifyRequestTravelerStandardHotelIntlCitys) SetCityName(v string) *ApplyModifyRequestTravelerStandardHotelIntlCitys {
	s.CityName = &v
	return s
}

func (s *ApplyModifyRequestTravelerStandardHotelIntlCitys) SetFee(v int64) *ApplyModifyRequestTravelerStandardHotelIntlCitys {
	s.Fee = &v
	return s
}

func (s *ApplyModifyRequestTravelerStandardHotelIntlCitys) Validate() error {
	return dara.Validate(s)
}
