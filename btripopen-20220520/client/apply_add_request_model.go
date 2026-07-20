// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyAddRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBudget(v int64) *ApplyAddRequest
	GetBudget() *int64
	SetBudgetMerge(v int32) *ApplyAddRequest
	GetBudgetMerge() *int32
	SetCarRule(v *ApplyAddRequestCarRule) *ApplyAddRequest
	GetCarRule() *ApplyAddRequestCarRule
	SetCorpName(v string) *ApplyAddRequest
	GetCorpName() *string
	SetDefaultStandard(v *ApplyAddRequestDefaultStandard) *ApplyAddRequest
	GetDefaultStandard() *ApplyAddRequestDefaultStandard
	SetDepartId(v string) *ApplyAddRequest
	GetDepartId() *string
	SetDepartName(v string) *ApplyAddRequest
	GetDepartName() *string
	SetExtendField(v string) *ApplyAddRequest
	GetExtendField() *string
	SetExternalTravelerList(v []*ApplyAddRequestExternalTravelerList) *ApplyAddRequest
	GetExternalTravelerList() []*ApplyAddRequestExternalTravelerList
	SetExternalTravelerStandard(v *ApplyAddRequestExternalTravelerStandard) *ApplyAddRequest
	GetExternalTravelerStandard() *ApplyAddRequestExternalTravelerStandard
	SetFlightBudget(v int64) *ApplyAddRequest
	GetFlightBudget() *int64
	SetHotelBudget(v int64) *ApplyAddRequest
	GetHotelBudget() *int64
	SetHotelShare(v *ApplyAddRequestHotelShare) *ApplyAddRequest
	GetHotelShare() *ApplyAddRequestHotelShare
	SetInternationalFlightCabins(v string) *ApplyAddRequest
	GetInternationalFlightCabins() *string
	SetIntlFlightBudget(v int64) *ApplyAddRequest
	GetIntlFlightBudget() *int64
	SetIntlHotelBudget(v int64) *ApplyAddRequest
	GetIntlHotelBudget() *int64
	SetItineraryList(v []*ApplyAddRequestItineraryList) *ApplyAddRequest
	GetItineraryList() []*ApplyAddRequestItineraryList
	SetItineraryRule(v int32) *ApplyAddRequest
	GetItineraryRule() *int32
	SetItinerarySetList(v []*ApplyAddRequestItinerarySetList) *ApplyAddRequest
	GetItinerarySetList() []*ApplyAddRequestItinerarySetList
	SetLimitTraveler(v int32) *ApplyAddRequest
	GetLimitTraveler() *int32
	SetMealBudget(v int64) *ApplyAddRequest
	GetMealBudget() *int64
	SetPaymentDepartmentId(v string) *ApplyAddRequest
	GetPaymentDepartmentId() *string
	SetPaymentDepartmentName(v string) *ApplyAddRequest
	GetPaymentDepartmentName() *string
	SetStatus(v int32) *ApplyAddRequest
	GetStatus() *int32
	SetSubCorpId(v string) *ApplyAddRequest
	GetSubCorpId() *string
	SetThirdpartApplyId(v string) *ApplyAddRequest
	GetThirdpartApplyId() *string
	SetThirdpartBusinessId(v string) *ApplyAddRequest
	GetThirdpartBusinessId() *string
	SetThirdpartDepartId(v string) *ApplyAddRequest
	GetThirdpartDepartId() *string
	SetTogetherBookRule(v int32) *ApplyAddRequest
	GetTogetherBookRule() *int32
	SetTrainBudget(v int64) *ApplyAddRequest
	GetTrainBudget() *int64
	SetTravelerList(v []*ApplyAddRequestTravelerList) *ApplyAddRequest
	GetTravelerList() []*ApplyAddRequestTravelerList
	SetTravelerStandard(v []*ApplyAddRequestTravelerStandard) *ApplyAddRequest
	GetTravelerStandard() []*ApplyAddRequestTravelerStandard
	SetTripCause(v string) *ApplyAddRequest
	GetTripCause() *string
	SetTripDay(v int32) *ApplyAddRequest
	GetTripDay() *int32
	SetTripTitle(v string) *ApplyAddRequest
	GetTripTitle() *string
	SetType(v int32) *ApplyAddRequest
	GetType() *int32
	SetUnionNo(v string) *ApplyAddRequest
	GetUnionNo() *string
	SetUserId(v string) *ApplyAddRequest
	GetUserId() *string
	SetUserName(v string) *ApplyAddRequest
	GetUserName() *string
	SetVehicleBudget(v int64) *ApplyAddRequest
	GetVehicleBudget() *int64
}

type ApplyAddRequest struct {
	// example:
	//
	// 4000
	Budget *int64 `json:"budget,omitempty" xml:"budget,omitempty"`
	// example:
	//
	// 1
	BudgetMerge *int32                  `json:"budget_merge,omitempty" xml:"budget_merge,omitempty"`
	CarRule     *ApplyAddRequestCarRule `json:"car_rule,omitempty" xml:"car_rule,omitempty" type:"Struct"`
	// example:
	//
	// 阿里巴巴
	CorpName        *string                         `json:"corp_name,omitempty" xml:"corp_name,omitempty"`
	DefaultStandard *ApplyAddRequestDefaultStandard `json:"default_standard,omitempty" xml:"default_standard,omitempty" type:"Struct"`
	// example:
	//
	// departId01
	DepartId *string `json:"depart_id,omitempty" xml:"depart_id,omitempty"`
	// example:
	//
	// 采购部
	DepartName *string `json:"depart_name,omitempty" xml:"depart_name,omitempty"`
	// 可将补充描述传入此字段，账单中将会体现此字段的值。可以用于企业的统计和对账
	//
	// example:
	//
	// {"cost_center":"成本中心"}
	ExtendField              *string                                  `json:"extend_field,omitempty" xml:"extend_field,omitempty"`
	ExternalTravelerList     []*ApplyAddRequestExternalTravelerList   `json:"external_traveler_list,omitempty" xml:"external_traveler_list,omitempty" type:"Repeated"`
	ExternalTravelerStandard *ApplyAddRequestExternalTravelerStandard `json:"external_traveler_standard,omitempty" xml:"external_traveler_standard,omitempty" type:"Struct"`
	// example:
	//
	// 1000
	FlightBudget *int64 `json:"flight_budget,omitempty" xml:"flight_budget,omitempty"`
	// example:
	//
	// 1000
	HotelBudget *int64                     `json:"hotel_budget,omitempty" xml:"hotel_budget,omitempty"`
	HotelShare  *ApplyAddRequestHotelShare `json:"hotel_share,omitempty" xml:"hotel_share,omitempty" type:"Struct"`
	// example:
	//
	// Y
	InternationalFlightCabins *string `json:"international_flight_cabins,omitempty" xml:"international_flight_cabins,omitempty"`
	// example:
	//
	// 1000
	IntlFlightBudget *int64 `json:"intl_flight_budget,omitempty" xml:"intl_flight_budget,omitempty"`
	// example:
	//
	// 1000
	IntlHotelBudget *int64                          `json:"intl_hotel_budget,omitempty" xml:"intl_hotel_budget,omitempty"`
	ItineraryList   []*ApplyAddRequestItineraryList `json:"itinerary_list,omitempty" xml:"itinerary_list,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	ItineraryRule    *int32                             `json:"itinerary_rule,omitempty" xml:"itinerary_rule,omitempty"`
	ItinerarySetList []*ApplyAddRequestItinerarySetList `json:"itinerary_set_list,omitempty" xml:"itinerary_set_list,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	LimitTraveler *int32 `json:"limit_traveler,omitempty" xml:"limit_traveler,omitempty"`
	// example:
	//
	// 1000
	MealBudget *int64 `json:"meal_budget,omitempty" xml:"meal_budget,omitempty"`
	// example:
	//
	// 41155
	PaymentDepartmentId *string `json:"payment_department_id,omitempty" xml:"payment_department_id,omitempty"`
	// example:
	//
	// 产品部
	PaymentDepartmentName *string `json:"payment_department_name,omitempty" xml:"payment_department_name,omitempty"`
	// example:
	//
	// 0
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// btrip123
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
	// example:
	//
	// 三方部门id
	ThirdpartDepartId *string `json:"thirdpart_depart_id,omitempty" xml:"thirdpart_depart_id,omitempty"`
	// example:
	//
	// 1
	TogetherBookRule *int32 `json:"together_book_rule,omitempty" xml:"together_book_rule,omitempty"`
	// example:
	//
	// 1000
	TrainBudget      *int64                             `json:"train_budget,omitempty" xml:"train_budget,omitempty"`
	TravelerList     []*ApplyAddRequestTravelerList     `json:"traveler_list,omitempty" xml:"traveler_list,omitempty" type:"Repeated"`
	TravelerStandard []*ApplyAddRequestTravelerStandard `json:"traveler_standard,omitempty" xml:"traveler_standard,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 客服培训
	TripCause *string `json:"trip_cause,omitempty" xml:"trip_cause,omitempty"`
	// example:
	//
	// 1
	TripDay *int32 `json:"trip_day,omitempty" xml:"trip_day,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 客服培训
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
	UserId *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// example:
	//
	// 张三
	UserName *string `json:"user_name,omitempty" xml:"user_name,omitempty"`
	// example:
	//
	// 1000
	VehicleBudget *int64 `json:"vehicle_budget,omitempty" xml:"vehicle_budget,omitempty"`
}

func (s ApplyAddRequest) String() string {
	return dara.Prettify(s)
}

func (s ApplyAddRequest) GoString() string {
	return s.String()
}

func (s *ApplyAddRequest) GetBudget() *int64 {
	return s.Budget
}

func (s *ApplyAddRequest) GetBudgetMerge() *int32 {
	return s.BudgetMerge
}

func (s *ApplyAddRequest) GetCarRule() *ApplyAddRequestCarRule {
	return s.CarRule
}

func (s *ApplyAddRequest) GetCorpName() *string {
	return s.CorpName
}

func (s *ApplyAddRequest) GetDefaultStandard() *ApplyAddRequestDefaultStandard {
	return s.DefaultStandard
}

func (s *ApplyAddRequest) GetDepartId() *string {
	return s.DepartId
}

func (s *ApplyAddRequest) GetDepartName() *string {
	return s.DepartName
}

func (s *ApplyAddRequest) GetExtendField() *string {
	return s.ExtendField
}

func (s *ApplyAddRequest) GetExternalTravelerList() []*ApplyAddRequestExternalTravelerList {
	return s.ExternalTravelerList
}

func (s *ApplyAddRequest) GetExternalTravelerStandard() *ApplyAddRequestExternalTravelerStandard {
	return s.ExternalTravelerStandard
}

func (s *ApplyAddRequest) GetFlightBudget() *int64 {
	return s.FlightBudget
}

func (s *ApplyAddRequest) GetHotelBudget() *int64 {
	return s.HotelBudget
}

func (s *ApplyAddRequest) GetHotelShare() *ApplyAddRequestHotelShare {
	return s.HotelShare
}

func (s *ApplyAddRequest) GetInternationalFlightCabins() *string {
	return s.InternationalFlightCabins
}

func (s *ApplyAddRequest) GetIntlFlightBudget() *int64 {
	return s.IntlFlightBudget
}

func (s *ApplyAddRequest) GetIntlHotelBudget() *int64 {
	return s.IntlHotelBudget
}

func (s *ApplyAddRequest) GetItineraryList() []*ApplyAddRequestItineraryList {
	return s.ItineraryList
}

func (s *ApplyAddRequest) GetItineraryRule() *int32 {
	return s.ItineraryRule
}

func (s *ApplyAddRequest) GetItinerarySetList() []*ApplyAddRequestItinerarySetList {
	return s.ItinerarySetList
}

func (s *ApplyAddRequest) GetLimitTraveler() *int32 {
	return s.LimitTraveler
}

func (s *ApplyAddRequest) GetMealBudget() *int64 {
	return s.MealBudget
}

func (s *ApplyAddRequest) GetPaymentDepartmentId() *string {
	return s.PaymentDepartmentId
}

func (s *ApplyAddRequest) GetPaymentDepartmentName() *string {
	return s.PaymentDepartmentName
}

func (s *ApplyAddRequest) GetStatus() *int32 {
	return s.Status
}

func (s *ApplyAddRequest) GetSubCorpId() *string {
	return s.SubCorpId
}

func (s *ApplyAddRequest) GetThirdpartApplyId() *string {
	return s.ThirdpartApplyId
}

func (s *ApplyAddRequest) GetThirdpartBusinessId() *string {
	return s.ThirdpartBusinessId
}

func (s *ApplyAddRequest) GetThirdpartDepartId() *string {
	return s.ThirdpartDepartId
}

func (s *ApplyAddRequest) GetTogetherBookRule() *int32 {
	return s.TogetherBookRule
}

func (s *ApplyAddRequest) GetTrainBudget() *int64 {
	return s.TrainBudget
}

func (s *ApplyAddRequest) GetTravelerList() []*ApplyAddRequestTravelerList {
	return s.TravelerList
}

func (s *ApplyAddRequest) GetTravelerStandard() []*ApplyAddRequestTravelerStandard {
	return s.TravelerStandard
}

func (s *ApplyAddRequest) GetTripCause() *string {
	return s.TripCause
}

func (s *ApplyAddRequest) GetTripDay() *int32 {
	return s.TripDay
}

func (s *ApplyAddRequest) GetTripTitle() *string {
	return s.TripTitle
}

func (s *ApplyAddRequest) GetType() *int32 {
	return s.Type
}

func (s *ApplyAddRequest) GetUnionNo() *string {
	return s.UnionNo
}

func (s *ApplyAddRequest) GetUserId() *string {
	return s.UserId
}

func (s *ApplyAddRequest) GetUserName() *string {
	return s.UserName
}

func (s *ApplyAddRequest) GetVehicleBudget() *int64 {
	return s.VehicleBudget
}

func (s *ApplyAddRequest) SetBudget(v int64) *ApplyAddRequest {
	s.Budget = &v
	return s
}

func (s *ApplyAddRequest) SetBudgetMerge(v int32) *ApplyAddRequest {
	s.BudgetMerge = &v
	return s
}

func (s *ApplyAddRequest) SetCarRule(v *ApplyAddRequestCarRule) *ApplyAddRequest {
	s.CarRule = v
	return s
}

func (s *ApplyAddRequest) SetCorpName(v string) *ApplyAddRequest {
	s.CorpName = &v
	return s
}

func (s *ApplyAddRequest) SetDefaultStandard(v *ApplyAddRequestDefaultStandard) *ApplyAddRequest {
	s.DefaultStandard = v
	return s
}

func (s *ApplyAddRequest) SetDepartId(v string) *ApplyAddRequest {
	s.DepartId = &v
	return s
}

func (s *ApplyAddRequest) SetDepartName(v string) *ApplyAddRequest {
	s.DepartName = &v
	return s
}

func (s *ApplyAddRequest) SetExtendField(v string) *ApplyAddRequest {
	s.ExtendField = &v
	return s
}

func (s *ApplyAddRequest) SetExternalTravelerList(v []*ApplyAddRequestExternalTravelerList) *ApplyAddRequest {
	s.ExternalTravelerList = v
	return s
}

func (s *ApplyAddRequest) SetExternalTravelerStandard(v *ApplyAddRequestExternalTravelerStandard) *ApplyAddRequest {
	s.ExternalTravelerStandard = v
	return s
}

func (s *ApplyAddRequest) SetFlightBudget(v int64) *ApplyAddRequest {
	s.FlightBudget = &v
	return s
}

func (s *ApplyAddRequest) SetHotelBudget(v int64) *ApplyAddRequest {
	s.HotelBudget = &v
	return s
}

func (s *ApplyAddRequest) SetHotelShare(v *ApplyAddRequestHotelShare) *ApplyAddRequest {
	s.HotelShare = v
	return s
}

func (s *ApplyAddRequest) SetInternationalFlightCabins(v string) *ApplyAddRequest {
	s.InternationalFlightCabins = &v
	return s
}

func (s *ApplyAddRequest) SetIntlFlightBudget(v int64) *ApplyAddRequest {
	s.IntlFlightBudget = &v
	return s
}

func (s *ApplyAddRequest) SetIntlHotelBudget(v int64) *ApplyAddRequest {
	s.IntlHotelBudget = &v
	return s
}

func (s *ApplyAddRequest) SetItineraryList(v []*ApplyAddRequestItineraryList) *ApplyAddRequest {
	s.ItineraryList = v
	return s
}

func (s *ApplyAddRequest) SetItineraryRule(v int32) *ApplyAddRequest {
	s.ItineraryRule = &v
	return s
}

func (s *ApplyAddRequest) SetItinerarySetList(v []*ApplyAddRequestItinerarySetList) *ApplyAddRequest {
	s.ItinerarySetList = v
	return s
}

func (s *ApplyAddRequest) SetLimitTraveler(v int32) *ApplyAddRequest {
	s.LimitTraveler = &v
	return s
}

func (s *ApplyAddRequest) SetMealBudget(v int64) *ApplyAddRequest {
	s.MealBudget = &v
	return s
}

func (s *ApplyAddRequest) SetPaymentDepartmentId(v string) *ApplyAddRequest {
	s.PaymentDepartmentId = &v
	return s
}

func (s *ApplyAddRequest) SetPaymentDepartmentName(v string) *ApplyAddRequest {
	s.PaymentDepartmentName = &v
	return s
}

func (s *ApplyAddRequest) SetStatus(v int32) *ApplyAddRequest {
	s.Status = &v
	return s
}

func (s *ApplyAddRequest) SetSubCorpId(v string) *ApplyAddRequest {
	s.SubCorpId = &v
	return s
}

func (s *ApplyAddRequest) SetThirdpartApplyId(v string) *ApplyAddRequest {
	s.ThirdpartApplyId = &v
	return s
}

func (s *ApplyAddRequest) SetThirdpartBusinessId(v string) *ApplyAddRequest {
	s.ThirdpartBusinessId = &v
	return s
}

func (s *ApplyAddRequest) SetThirdpartDepartId(v string) *ApplyAddRequest {
	s.ThirdpartDepartId = &v
	return s
}

func (s *ApplyAddRequest) SetTogetherBookRule(v int32) *ApplyAddRequest {
	s.TogetherBookRule = &v
	return s
}

func (s *ApplyAddRequest) SetTrainBudget(v int64) *ApplyAddRequest {
	s.TrainBudget = &v
	return s
}

func (s *ApplyAddRequest) SetTravelerList(v []*ApplyAddRequestTravelerList) *ApplyAddRequest {
	s.TravelerList = v
	return s
}

func (s *ApplyAddRequest) SetTravelerStandard(v []*ApplyAddRequestTravelerStandard) *ApplyAddRequest {
	s.TravelerStandard = v
	return s
}

func (s *ApplyAddRequest) SetTripCause(v string) *ApplyAddRequest {
	s.TripCause = &v
	return s
}

func (s *ApplyAddRequest) SetTripDay(v int32) *ApplyAddRequest {
	s.TripDay = &v
	return s
}

func (s *ApplyAddRequest) SetTripTitle(v string) *ApplyAddRequest {
	s.TripTitle = &v
	return s
}

func (s *ApplyAddRequest) SetType(v int32) *ApplyAddRequest {
	s.Type = &v
	return s
}

func (s *ApplyAddRequest) SetUnionNo(v string) *ApplyAddRequest {
	s.UnionNo = &v
	return s
}

func (s *ApplyAddRequest) SetUserId(v string) *ApplyAddRequest {
	s.UserId = &v
	return s
}

func (s *ApplyAddRequest) SetUserName(v string) *ApplyAddRequest {
	s.UserName = &v
	return s
}

func (s *ApplyAddRequest) SetVehicleBudget(v int64) *ApplyAddRequest {
	s.VehicleBudget = &v
	return s
}

func (s *ApplyAddRequest) Validate() error {
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

type ApplyAddRequestCarRule struct {
	// example:
	//
	// 1234567
	ScenarioTemplateId *string `json:"scenario_template_id,omitempty" xml:"scenario_template_id,omitempty"`
	// example:
	//
	// 测试场景模板
	ScenarioTemplateName *string `json:"scenario_template_name,omitempty" xml:"scenario_template_name,omitempty"`
}

func (s ApplyAddRequestCarRule) String() string {
	return dara.Prettify(s)
}

func (s ApplyAddRequestCarRule) GoString() string {
	return s.String()
}

func (s *ApplyAddRequestCarRule) GetScenarioTemplateId() *string {
	return s.ScenarioTemplateId
}

func (s *ApplyAddRequestCarRule) GetScenarioTemplateName() *string {
	return s.ScenarioTemplateName
}

func (s *ApplyAddRequestCarRule) SetScenarioTemplateId(v string) *ApplyAddRequestCarRule {
	s.ScenarioTemplateId = &v
	return s
}

func (s *ApplyAddRequestCarRule) SetScenarioTemplateName(v string) *ApplyAddRequestCarRule {
	s.ScenarioTemplateName = &v
	return s
}

func (s *ApplyAddRequestCarRule) Validate() error {
	return dara.Validate(s)
}

type ApplyAddRequestDefaultStandard struct {
	// example:
	//
	// 3
	BusinessDiscount *int32 `json:"business_discount,omitempty" xml:"business_discount,omitempty"`
	// example:
	//
	// 3
	EconomyDiscount *int32 `json:"economy_discount,omitempty" xml:"economy_discount,omitempty"`
	// example:
	//
	// 3
	FirstDiscount *int32 `json:"first_discount,omitempty" xml:"first_discount,omitempty"`
	// example:
	//
	// Y
	FlightCabins *string `json:"flight_cabins,omitempty" xml:"flight_cabins,omitempty"`
	// example:
	//
	// 100132
	FlightIntlRuleCode *int64 `json:"flight_intl_rule_code,omitempty" xml:"flight_intl_rule_code,omitempty"`
	// example:
	//
	// 100132
	FlightRuleCode *int64                                          `json:"flight_rule_code,omitempty" xml:"flight_rule_code,omitempty"`
	HotelCitys     []*ApplyAddRequestDefaultStandardHotelCitys     `json:"hotel_citys,omitempty" xml:"hotel_citys,omitempty" type:"Repeated"`
	HotelIntlCitys []*ApplyAddRequestDefaultStandardHotelIntlCitys `json:"hotel_intl_citys,omitempty" xml:"hotel_intl_citys,omitempty" type:"Repeated"`
	// example:
	//
	// 100132
	HotelIntlRuleCode *int64 `json:"hotel_intl_rule_code,omitempty" xml:"hotel_intl_rule_code,omitempty"`
	// example:
	//
	// 100132
	HotelRuleCode *int64 `json:"hotel_rule_code,omitempty" xml:"hotel_rule_code,omitempty"`
	// example:
	//
	// F
	InternationalFlightCabins *string `json:"international_flight_cabins,omitempty" xml:"international_flight_cabins,omitempty"`
	// example:
	//
	// 9
	PremiumEconomyDiscount *int32 `json:"premium_economy_discount,omitempty" xml:"premium_economy_discount,omitempty"`
	// example:
	//
	// 1
	ReserveType *int32 `json:"reserve_type,omitempty" xml:"reserve_type,omitempty"`
	// example:
	//
	// 100132
	TrainRuleCode *int64 `json:"train_rule_code,omitempty" xml:"train_rule_code,omitempty"`
	// example:
	//
	// 1000
	TrainSeats *string `json:"train_seats,omitempty" xml:"train_seats,omitempty"`
}

func (s ApplyAddRequestDefaultStandard) String() string {
	return dara.Prettify(s)
}

func (s ApplyAddRequestDefaultStandard) GoString() string {
	return s.String()
}

func (s *ApplyAddRequestDefaultStandard) GetBusinessDiscount() *int32 {
	return s.BusinessDiscount
}

func (s *ApplyAddRequestDefaultStandard) GetEconomyDiscount() *int32 {
	return s.EconomyDiscount
}

func (s *ApplyAddRequestDefaultStandard) GetFirstDiscount() *int32 {
	return s.FirstDiscount
}

func (s *ApplyAddRequestDefaultStandard) GetFlightCabins() *string {
	return s.FlightCabins
}

func (s *ApplyAddRequestDefaultStandard) GetFlightIntlRuleCode() *int64 {
	return s.FlightIntlRuleCode
}

func (s *ApplyAddRequestDefaultStandard) GetFlightRuleCode() *int64 {
	return s.FlightRuleCode
}

func (s *ApplyAddRequestDefaultStandard) GetHotelCitys() []*ApplyAddRequestDefaultStandardHotelCitys {
	return s.HotelCitys
}

func (s *ApplyAddRequestDefaultStandard) GetHotelIntlCitys() []*ApplyAddRequestDefaultStandardHotelIntlCitys {
	return s.HotelIntlCitys
}

func (s *ApplyAddRequestDefaultStandard) GetHotelIntlRuleCode() *int64 {
	return s.HotelIntlRuleCode
}

func (s *ApplyAddRequestDefaultStandard) GetHotelRuleCode() *int64 {
	return s.HotelRuleCode
}

func (s *ApplyAddRequestDefaultStandard) GetInternationalFlightCabins() *string {
	return s.InternationalFlightCabins
}

func (s *ApplyAddRequestDefaultStandard) GetPremiumEconomyDiscount() *int32 {
	return s.PremiumEconomyDiscount
}

func (s *ApplyAddRequestDefaultStandard) GetReserveType() *int32 {
	return s.ReserveType
}

func (s *ApplyAddRequestDefaultStandard) GetTrainRuleCode() *int64 {
	return s.TrainRuleCode
}

func (s *ApplyAddRequestDefaultStandard) GetTrainSeats() *string {
	return s.TrainSeats
}

func (s *ApplyAddRequestDefaultStandard) SetBusinessDiscount(v int32) *ApplyAddRequestDefaultStandard {
	s.BusinessDiscount = &v
	return s
}

func (s *ApplyAddRequestDefaultStandard) SetEconomyDiscount(v int32) *ApplyAddRequestDefaultStandard {
	s.EconomyDiscount = &v
	return s
}

func (s *ApplyAddRequestDefaultStandard) SetFirstDiscount(v int32) *ApplyAddRequestDefaultStandard {
	s.FirstDiscount = &v
	return s
}

func (s *ApplyAddRequestDefaultStandard) SetFlightCabins(v string) *ApplyAddRequestDefaultStandard {
	s.FlightCabins = &v
	return s
}

func (s *ApplyAddRequestDefaultStandard) SetFlightIntlRuleCode(v int64) *ApplyAddRequestDefaultStandard {
	s.FlightIntlRuleCode = &v
	return s
}

func (s *ApplyAddRequestDefaultStandard) SetFlightRuleCode(v int64) *ApplyAddRequestDefaultStandard {
	s.FlightRuleCode = &v
	return s
}

func (s *ApplyAddRequestDefaultStandard) SetHotelCitys(v []*ApplyAddRequestDefaultStandardHotelCitys) *ApplyAddRequestDefaultStandard {
	s.HotelCitys = v
	return s
}

func (s *ApplyAddRequestDefaultStandard) SetHotelIntlCitys(v []*ApplyAddRequestDefaultStandardHotelIntlCitys) *ApplyAddRequestDefaultStandard {
	s.HotelIntlCitys = v
	return s
}

func (s *ApplyAddRequestDefaultStandard) SetHotelIntlRuleCode(v int64) *ApplyAddRequestDefaultStandard {
	s.HotelIntlRuleCode = &v
	return s
}

func (s *ApplyAddRequestDefaultStandard) SetHotelRuleCode(v int64) *ApplyAddRequestDefaultStandard {
	s.HotelRuleCode = &v
	return s
}

func (s *ApplyAddRequestDefaultStandard) SetInternationalFlightCabins(v string) *ApplyAddRequestDefaultStandard {
	s.InternationalFlightCabins = &v
	return s
}

func (s *ApplyAddRequestDefaultStandard) SetPremiumEconomyDiscount(v int32) *ApplyAddRequestDefaultStandard {
	s.PremiumEconomyDiscount = &v
	return s
}

func (s *ApplyAddRequestDefaultStandard) SetReserveType(v int32) *ApplyAddRequestDefaultStandard {
	s.ReserveType = &v
	return s
}

func (s *ApplyAddRequestDefaultStandard) SetTrainRuleCode(v int64) *ApplyAddRequestDefaultStandard {
	s.TrainRuleCode = &v
	return s
}

func (s *ApplyAddRequestDefaultStandard) SetTrainSeats(v string) *ApplyAddRequestDefaultStandard {
	s.TrainSeats = &v
	return s
}

func (s *ApplyAddRequestDefaultStandard) Validate() error {
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

type ApplyAddRequestDefaultStandardHotelCitys struct {
	// example:
	//
	// 360100
	CityCode *string `json:"city_code,omitempty" xml:"city_code,omitempty"`
	// example:
	//
	// 杭州
	CityName *string `json:"city_name,omitempty" xml:"city_name,omitempty"`
	// example:
	//
	// 10000
	Fee *int64 `json:"fee,omitempty" xml:"fee,omitempty"`
}

func (s ApplyAddRequestDefaultStandardHotelCitys) String() string {
	return dara.Prettify(s)
}

func (s ApplyAddRequestDefaultStandardHotelCitys) GoString() string {
	return s.String()
}

func (s *ApplyAddRequestDefaultStandardHotelCitys) GetCityCode() *string {
	return s.CityCode
}

func (s *ApplyAddRequestDefaultStandardHotelCitys) GetCityName() *string {
	return s.CityName
}

func (s *ApplyAddRequestDefaultStandardHotelCitys) GetFee() *int64 {
	return s.Fee
}

func (s *ApplyAddRequestDefaultStandardHotelCitys) SetCityCode(v string) *ApplyAddRequestDefaultStandardHotelCitys {
	s.CityCode = &v
	return s
}

func (s *ApplyAddRequestDefaultStandardHotelCitys) SetCityName(v string) *ApplyAddRequestDefaultStandardHotelCitys {
	s.CityName = &v
	return s
}

func (s *ApplyAddRequestDefaultStandardHotelCitys) SetFee(v int64) *ApplyAddRequestDefaultStandardHotelCitys {
	s.Fee = &v
	return s
}

func (s *ApplyAddRequestDefaultStandardHotelCitys) Validate() error {
	return dara.Validate(s)
}

type ApplyAddRequestDefaultStandardHotelIntlCitys struct {
	// example:
	//
	// 210200
	CityCode *string `json:"city_code,omitempty" xml:"city_code,omitempty"`
	// example:
	//
	// 香港
	CityName *string `json:"city_name,omitempty" xml:"city_name,omitempty"`
	// example:
	//
	// 100000
	Fee *int64 `json:"fee,omitempty" xml:"fee,omitempty"`
}

func (s ApplyAddRequestDefaultStandardHotelIntlCitys) String() string {
	return dara.Prettify(s)
}

func (s ApplyAddRequestDefaultStandardHotelIntlCitys) GoString() string {
	return s.String()
}

func (s *ApplyAddRequestDefaultStandardHotelIntlCitys) GetCityCode() *string {
	return s.CityCode
}

func (s *ApplyAddRequestDefaultStandardHotelIntlCitys) GetCityName() *string {
	return s.CityName
}

func (s *ApplyAddRequestDefaultStandardHotelIntlCitys) GetFee() *int64 {
	return s.Fee
}

func (s *ApplyAddRequestDefaultStandardHotelIntlCitys) SetCityCode(v string) *ApplyAddRequestDefaultStandardHotelIntlCitys {
	s.CityCode = &v
	return s
}

func (s *ApplyAddRequestDefaultStandardHotelIntlCitys) SetCityName(v string) *ApplyAddRequestDefaultStandardHotelIntlCitys {
	s.CityName = &v
	return s
}

func (s *ApplyAddRequestDefaultStandardHotelIntlCitys) SetFee(v int64) *ApplyAddRequestDefaultStandardHotelIntlCitys {
	s.Fee = &v
	return s
}

func (s *ApplyAddRequestDefaultStandardHotelIntlCitys) Validate() error {
	return dara.Validate(s)
}

type ApplyAddRequestExternalTravelerList struct {
	// example:
	//
	// “{"name":"张三"}”
	Attribute *string `json:"attribute,omitempty" xml:"attribute,omitempty"`
	// example:
	//
	// 414522
	CostCenterId *int64 `json:"cost_center_id,omitempty" xml:"cost_center_id,omitempty"`
	// example:
	//
	// 371164
	ExternalUserId *string `json:"external_user_id,omitempty" xml:"external_user_id,omitempty"`
	// example:
	//
	// 4451
	InvoiceId *int64 `json:"invoice_id,omitempty" xml:"invoice_id,omitempty"`
	// example:
	//
	// 141125
	PaymentDepartmentId *string `json:"payment_department_id,omitempty" xml:"payment_department_id,omitempty"`
	// example:
	//
	// 产品部
	PaymentDepartmentName *string `json:"payment_department_name,omitempty" xml:"payment_department_name,omitempty"`
	// example:
	//
	// acs
	ProjectCode *string `json:"project_code,omitempty" xml:"project_code,omitempty"`
	// example:
	//
	// “成本项目”
	ProjectTitle *string `json:"project_title,omitempty" xml:"project_title,omitempty"`
	// example:
	//
	// 91130124566177980M
	ThirdPartInvoiceId *string `json:"third_part_invoice_id,omitempty" xml:"third_part_invoice_id,omitempty"`
	// example:
	//
	// HD155
	ThirdpartCostCenterId *string `json:"thirdpart_cost_center_id,omitempty" xml:"thirdpart_cost_center_id,omitempty"`
	// example:
	//
	// DHDI2208051REIO6BK
	ThirdpartDepartId *string `json:"thirdpart_depart_id,omitempty" xml:"thirdpart_depart_id,omitempty"`
	// example:
	//
	// 李四
	UserName *string `json:"user_name,omitempty" xml:"user_name,omitempty"`
	// example:
	//
	// SUN/MENGXUAN
	UserNameEn *string `json:"user_name_en,omitempty" xml:"user_name_en,omitempty"`
}

func (s ApplyAddRequestExternalTravelerList) String() string {
	return dara.Prettify(s)
}

func (s ApplyAddRequestExternalTravelerList) GoString() string {
	return s.String()
}

func (s *ApplyAddRequestExternalTravelerList) GetAttribute() *string {
	return s.Attribute
}

func (s *ApplyAddRequestExternalTravelerList) GetCostCenterId() *int64 {
	return s.CostCenterId
}

func (s *ApplyAddRequestExternalTravelerList) GetExternalUserId() *string {
	return s.ExternalUserId
}

func (s *ApplyAddRequestExternalTravelerList) GetInvoiceId() *int64 {
	return s.InvoiceId
}

func (s *ApplyAddRequestExternalTravelerList) GetPaymentDepartmentId() *string {
	return s.PaymentDepartmentId
}

func (s *ApplyAddRequestExternalTravelerList) GetPaymentDepartmentName() *string {
	return s.PaymentDepartmentName
}

func (s *ApplyAddRequestExternalTravelerList) GetProjectCode() *string {
	return s.ProjectCode
}

func (s *ApplyAddRequestExternalTravelerList) GetProjectTitle() *string {
	return s.ProjectTitle
}

func (s *ApplyAddRequestExternalTravelerList) GetThirdPartInvoiceId() *string {
	return s.ThirdPartInvoiceId
}

func (s *ApplyAddRequestExternalTravelerList) GetThirdpartCostCenterId() *string {
	return s.ThirdpartCostCenterId
}

func (s *ApplyAddRequestExternalTravelerList) GetThirdpartDepartId() *string {
	return s.ThirdpartDepartId
}

func (s *ApplyAddRequestExternalTravelerList) GetUserName() *string {
	return s.UserName
}

func (s *ApplyAddRequestExternalTravelerList) GetUserNameEn() *string {
	return s.UserNameEn
}

func (s *ApplyAddRequestExternalTravelerList) SetAttribute(v string) *ApplyAddRequestExternalTravelerList {
	s.Attribute = &v
	return s
}

func (s *ApplyAddRequestExternalTravelerList) SetCostCenterId(v int64) *ApplyAddRequestExternalTravelerList {
	s.CostCenterId = &v
	return s
}

func (s *ApplyAddRequestExternalTravelerList) SetExternalUserId(v string) *ApplyAddRequestExternalTravelerList {
	s.ExternalUserId = &v
	return s
}

func (s *ApplyAddRequestExternalTravelerList) SetInvoiceId(v int64) *ApplyAddRequestExternalTravelerList {
	s.InvoiceId = &v
	return s
}

func (s *ApplyAddRequestExternalTravelerList) SetPaymentDepartmentId(v string) *ApplyAddRequestExternalTravelerList {
	s.PaymentDepartmentId = &v
	return s
}

func (s *ApplyAddRequestExternalTravelerList) SetPaymentDepartmentName(v string) *ApplyAddRequestExternalTravelerList {
	s.PaymentDepartmentName = &v
	return s
}

func (s *ApplyAddRequestExternalTravelerList) SetProjectCode(v string) *ApplyAddRequestExternalTravelerList {
	s.ProjectCode = &v
	return s
}

func (s *ApplyAddRequestExternalTravelerList) SetProjectTitle(v string) *ApplyAddRequestExternalTravelerList {
	s.ProjectTitle = &v
	return s
}

func (s *ApplyAddRequestExternalTravelerList) SetThirdPartInvoiceId(v string) *ApplyAddRequestExternalTravelerList {
	s.ThirdPartInvoiceId = &v
	return s
}

func (s *ApplyAddRequestExternalTravelerList) SetThirdpartCostCenterId(v string) *ApplyAddRequestExternalTravelerList {
	s.ThirdpartCostCenterId = &v
	return s
}

func (s *ApplyAddRequestExternalTravelerList) SetThirdpartDepartId(v string) *ApplyAddRequestExternalTravelerList {
	s.ThirdpartDepartId = &v
	return s
}

func (s *ApplyAddRequestExternalTravelerList) SetUserName(v string) *ApplyAddRequestExternalTravelerList {
	s.UserName = &v
	return s
}

func (s *ApplyAddRequestExternalTravelerList) SetUserNameEn(v string) *ApplyAddRequestExternalTravelerList {
	s.UserNameEn = &v
	return s
}

func (s *ApplyAddRequestExternalTravelerList) Validate() error {
	return dara.Validate(s)
}

type ApplyAddRequestExternalTravelerStandard struct {
	// example:
	//
	// 1
	BusinessDiscount *int32 `json:"business_discount,omitempty" xml:"business_discount,omitempty"`
	// example:
	//
	// 1
	EconomyDiscount *int32 `json:"economy_discount,omitempty" xml:"economy_discount,omitempty"`
	// example:
	//
	// 1
	FirstDiscount *int32 `json:"first_discount,omitempty" xml:"first_discount,omitempty"`
	// example:
	//
	// F
	FlightCabins *string `json:"flight_cabins,omitempty" xml:"flight_cabins,omitempty"`
	// example:
	//
	// 100132
	FlightIntlRuleCode *int64 `json:"flight_intl_rule_code,omitempty" xml:"flight_intl_rule_code,omitempty"`
	// example:
	//
	// 100132
	FlightRuleCode *int64                                                   `json:"flight_rule_code,omitempty" xml:"flight_rule_code,omitempty"`
	HotelCitys     []*ApplyAddRequestExternalTravelerStandardHotelCitys     `json:"hotel_citys,omitempty" xml:"hotel_citys,omitempty" type:"Repeated"`
	HotelIntlCitys []*ApplyAddRequestExternalTravelerStandardHotelIntlCitys `json:"hotel_intl_citys,omitempty" xml:"hotel_intl_citys,omitempty" type:"Repeated"`
	// example:
	//
	// 100132
	HotelIntlRuleCode *int64 `json:"hotel_intl_rule_code,omitempty" xml:"hotel_intl_rule_code,omitempty"`
	// example:
	//
	// 100132
	HotelRuleCode *int64 `json:"hotel_rule_code,omitempty" xml:"hotel_rule_code,omitempty"`
	// example:
	//
	// F
	InternationalFlightCabins *string `json:"international_flight_cabins,omitempty" xml:"international_flight_cabins,omitempty"`
	// example:
	//
	// 1
	PremiumEconomyDiscount *int32 `json:"premium_economy_discount,omitempty" xml:"premium_economy_discount,omitempty"`
	// example:
	//
	// 0
	ReserveType *int32 `json:"reserve_type,omitempty" xml:"reserve_type,omitempty"`
	// example:
	//
	// 100132
	TrainRuleCode *int64 `json:"train_rule_code,omitempty" xml:"train_rule_code,omitempty"`
	// example:
	//
	// 0
	TrainSeats *string `json:"train_seats,omitempty" xml:"train_seats,omitempty"`
}

func (s ApplyAddRequestExternalTravelerStandard) String() string {
	return dara.Prettify(s)
}

func (s ApplyAddRequestExternalTravelerStandard) GoString() string {
	return s.String()
}

func (s *ApplyAddRequestExternalTravelerStandard) GetBusinessDiscount() *int32 {
	return s.BusinessDiscount
}

func (s *ApplyAddRequestExternalTravelerStandard) GetEconomyDiscount() *int32 {
	return s.EconomyDiscount
}

func (s *ApplyAddRequestExternalTravelerStandard) GetFirstDiscount() *int32 {
	return s.FirstDiscount
}

func (s *ApplyAddRequestExternalTravelerStandard) GetFlightCabins() *string {
	return s.FlightCabins
}

func (s *ApplyAddRequestExternalTravelerStandard) GetFlightIntlRuleCode() *int64 {
	return s.FlightIntlRuleCode
}

func (s *ApplyAddRequestExternalTravelerStandard) GetFlightRuleCode() *int64 {
	return s.FlightRuleCode
}

func (s *ApplyAddRequestExternalTravelerStandard) GetHotelCitys() []*ApplyAddRequestExternalTravelerStandardHotelCitys {
	return s.HotelCitys
}

func (s *ApplyAddRequestExternalTravelerStandard) GetHotelIntlCitys() []*ApplyAddRequestExternalTravelerStandardHotelIntlCitys {
	return s.HotelIntlCitys
}

func (s *ApplyAddRequestExternalTravelerStandard) GetHotelIntlRuleCode() *int64 {
	return s.HotelIntlRuleCode
}

func (s *ApplyAddRequestExternalTravelerStandard) GetHotelRuleCode() *int64 {
	return s.HotelRuleCode
}

func (s *ApplyAddRequestExternalTravelerStandard) GetInternationalFlightCabins() *string {
	return s.InternationalFlightCabins
}

func (s *ApplyAddRequestExternalTravelerStandard) GetPremiumEconomyDiscount() *int32 {
	return s.PremiumEconomyDiscount
}

func (s *ApplyAddRequestExternalTravelerStandard) GetReserveType() *int32 {
	return s.ReserveType
}

func (s *ApplyAddRequestExternalTravelerStandard) GetTrainRuleCode() *int64 {
	return s.TrainRuleCode
}

func (s *ApplyAddRequestExternalTravelerStandard) GetTrainSeats() *string {
	return s.TrainSeats
}

func (s *ApplyAddRequestExternalTravelerStandard) SetBusinessDiscount(v int32) *ApplyAddRequestExternalTravelerStandard {
	s.BusinessDiscount = &v
	return s
}

func (s *ApplyAddRequestExternalTravelerStandard) SetEconomyDiscount(v int32) *ApplyAddRequestExternalTravelerStandard {
	s.EconomyDiscount = &v
	return s
}

func (s *ApplyAddRequestExternalTravelerStandard) SetFirstDiscount(v int32) *ApplyAddRequestExternalTravelerStandard {
	s.FirstDiscount = &v
	return s
}

func (s *ApplyAddRequestExternalTravelerStandard) SetFlightCabins(v string) *ApplyAddRequestExternalTravelerStandard {
	s.FlightCabins = &v
	return s
}

func (s *ApplyAddRequestExternalTravelerStandard) SetFlightIntlRuleCode(v int64) *ApplyAddRequestExternalTravelerStandard {
	s.FlightIntlRuleCode = &v
	return s
}

func (s *ApplyAddRequestExternalTravelerStandard) SetFlightRuleCode(v int64) *ApplyAddRequestExternalTravelerStandard {
	s.FlightRuleCode = &v
	return s
}

func (s *ApplyAddRequestExternalTravelerStandard) SetHotelCitys(v []*ApplyAddRequestExternalTravelerStandardHotelCitys) *ApplyAddRequestExternalTravelerStandard {
	s.HotelCitys = v
	return s
}

func (s *ApplyAddRequestExternalTravelerStandard) SetHotelIntlCitys(v []*ApplyAddRequestExternalTravelerStandardHotelIntlCitys) *ApplyAddRequestExternalTravelerStandard {
	s.HotelIntlCitys = v
	return s
}

func (s *ApplyAddRequestExternalTravelerStandard) SetHotelIntlRuleCode(v int64) *ApplyAddRequestExternalTravelerStandard {
	s.HotelIntlRuleCode = &v
	return s
}

func (s *ApplyAddRequestExternalTravelerStandard) SetHotelRuleCode(v int64) *ApplyAddRequestExternalTravelerStandard {
	s.HotelRuleCode = &v
	return s
}

func (s *ApplyAddRequestExternalTravelerStandard) SetInternationalFlightCabins(v string) *ApplyAddRequestExternalTravelerStandard {
	s.InternationalFlightCabins = &v
	return s
}

func (s *ApplyAddRequestExternalTravelerStandard) SetPremiumEconomyDiscount(v int32) *ApplyAddRequestExternalTravelerStandard {
	s.PremiumEconomyDiscount = &v
	return s
}

func (s *ApplyAddRequestExternalTravelerStandard) SetReserveType(v int32) *ApplyAddRequestExternalTravelerStandard {
	s.ReserveType = &v
	return s
}

func (s *ApplyAddRequestExternalTravelerStandard) SetTrainRuleCode(v int64) *ApplyAddRequestExternalTravelerStandard {
	s.TrainRuleCode = &v
	return s
}

func (s *ApplyAddRequestExternalTravelerStandard) SetTrainSeats(v string) *ApplyAddRequestExternalTravelerStandard {
	s.TrainSeats = &v
	return s
}

func (s *ApplyAddRequestExternalTravelerStandard) Validate() error {
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

type ApplyAddRequestExternalTravelerStandardHotelCitys struct {
	// example:
	//
	// 0
	CityCode *string `json:"city_code,omitempty" xml:"city_code,omitempty"`
	// example:
	//
	// 北京
	CityName *string `json:"city_name,omitempty" xml:"city_name,omitempty"`
	// example:
	//
	// 1009
	Fee *int64 `json:"fee,omitempty" xml:"fee,omitempty"`
}

func (s ApplyAddRequestExternalTravelerStandardHotelCitys) String() string {
	return dara.Prettify(s)
}

func (s ApplyAddRequestExternalTravelerStandardHotelCitys) GoString() string {
	return s.String()
}

func (s *ApplyAddRequestExternalTravelerStandardHotelCitys) GetCityCode() *string {
	return s.CityCode
}

func (s *ApplyAddRequestExternalTravelerStandardHotelCitys) GetCityName() *string {
	return s.CityName
}

func (s *ApplyAddRequestExternalTravelerStandardHotelCitys) GetFee() *int64 {
	return s.Fee
}

func (s *ApplyAddRequestExternalTravelerStandardHotelCitys) SetCityCode(v string) *ApplyAddRequestExternalTravelerStandardHotelCitys {
	s.CityCode = &v
	return s
}

func (s *ApplyAddRequestExternalTravelerStandardHotelCitys) SetCityName(v string) *ApplyAddRequestExternalTravelerStandardHotelCitys {
	s.CityName = &v
	return s
}

func (s *ApplyAddRequestExternalTravelerStandardHotelCitys) SetFee(v int64) *ApplyAddRequestExternalTravelerStandardHotelCitys {
	s.Fee = &v
	return s
}

func (s *ApplyAddRequestExternalTravelerStandardHotelCitys) Validate() error {
	return dara.Validate(s)
}

type ApplyAddRequestExternalTravelerStandardHotelIntlCitys struct {
	// example:
	//
	// 0
	CityCode *string `json:"city_code,omitempty" xml:"city_code,omitempty"`
	// example:
	//
	// 北京
	CityName *string `json:"city_name,omitempty" xml:"city_name,omitempty"`
	// example:
	//
	// 1000
	Fee *int64 `json:"fee,omitempty" xml:"fee,omitempty"`
}

func (s ApplyAddRequestExternalTravelerStandardHotelIntlCitys) String() string {
	return dara.Prettify(s)
}

func (s ApplyAddRequestExternalTravelerStandardHotelIntlCitys) GoString() string {
	return s.String()
}

func (s *ApplyAddRequestExternalTravelerStandardHotelIntlCitys) GetCityCode() *string {
	return s.CityCode
}

func (s *ApplyAddRequestExternalTravelerStandardHotelIntlCitys) GetCityName() *string {
	return s.CityName
}

func (s *ApplyAddRequestExternalTravelerStandardHotelIntlCitys) GetFee() *int64 {
	return s.Fee
}

func (s *ApplyAddRequestExternalTravelerStandardHotelIntlCitys) SetCityCode(v string) *ApplyAddRequestExternalTravelerStandardHotelIntlCitys {
	s.CityCode = &v
	return s
}

func (s *ApplyAddRequestExternalTravelerStandardHotelIntlCitys) SetCityName(v string) *ApplyAddRequestExternalTravelerStandardHotelIntlCitys {
	s.CityName = &v
	return s
}

func (s *ApplyAddRequestExternalTravelerStandardHotelIntlCitys) SetFee(v int64) *ApplyAddRequestExternalTravelerStandardHotelIntlCitys {
	s.Fee = &v
	return s
}

func (s *ApplyAddRequestExternalTravelerStandardHotelIntlCitys) Validate() error {
	return dara.Validate(s)
}

type ApplyAddRequestHotelShare struct {
	// example:
	//
	// 70
	Param *string `json:"param,omitempty" xml:"param,omitempty"`
	// example:
	//
	// 1
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ApplyAddRequestHotelShare) String() string {
	return dara.Prettify(s)
}

func (s ApplyAddRequestHotelShare) GoString() string {
	return s.String()
}

func (s *ApplyAddRequestHotelShare) GetParam() *string {
	return s.Param
}

func (s *ApplyAddRequestHotelShare) GetType() *string {
	return s.Type
}

func (s *ApplyAddRequestHotelShare) SetParam(v string) *ApplyAddRequestHotelShare {
	s.Param = &v
	return s
}

func (s *ApplyAddRequestHotelShare) SetType(v string) *ApplyAddRequestHotelShare {
	s.Type = &v
	return s
}

func (s *ApplyAddRequestHotelShare) Validate() error {
	return dara.Validate(s)
}

type ApplyAddRequestItineraryList struct {
	// This parameter is required.
	//
	// example:
	//
	// 杭州
	ArrCity *string `json:"arr_city,omitempty" xml:"arr_city,omitempty"`
	// example:
	//
	// HGH
	ArrCityCode *string `json:"arr_city_code,omitempty" xml:"arr_city_code,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2017-01-02 00:00:00
	ArrDate *string `json:"arr_date,omitempty" xml:"arr_date,omitempty"`
	// example:
	//
	// "{"name":"张三"}"
	Attribute *string `json:"attribute,omitempty" xml:"attribute,omitempty"`
	// example:
	//
	// 12138
	CostCenterId *int64 `json:"cost_center_id,omitempty" xml:"cost_center_id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 北京
	DepCity *string `json:"dep_city,omitempty" xml:"dep_city,omitempty"`
	// example:
	//
	// BJS
	DepCityCode *string `json:"dep_city_code,omitempty" xml:"dep_city_code,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2017-01-01 00:00:00
	DepDate *string `json:"dep_date,omitempty" xml:"dep_date,omitempty"`
	// example:
	//
	// 34711
	InvoiceId *int64 `json:"invoice_id,omitempty" xml:"invoice_id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123456
	ItineraryId             *string                                              `json:"itinerary_id,omitempty" xml:"itinerary_id,omitempty"`
	ItineraryTravelStandard *ApplyAddRequestItineraryListItineraryTravelStandard `json:"itinerary_travel_standard,omitempty" xml:"itinerary_travel_standard,omitempty" type:"Struct"`
	// example:
	//
	// true
	NeedHotel *bool `json:"need_hotel,omitempty" xml:"need_hotel,omitempty"`
	// example:
	//
	// true
	NeedTraffic *bool `json:"need_traffic,omitempty" xml:"need_traffic,omitempty"`
	// example:
	//
	// projecttow
	ProjectCode *string `json:"project_code,omitempty" xml:"project_code,omitempty"`
	// example:
	//
	// 项目1
	ProjectTitle              *string   `json:"project_title,omitempty" xml:"project_title,omitempty"`
	ProvinceTravelCityAdcodes []*string `json:"province_travel_city_adcodes,omitempty" xml:"province_travel_city_adcodes,omitempty" type:"Repeated"`
	// example:
	//
	// thirdpart34711
	ThirdPartInvoiceId *string `json:"third_part_invoice_id,omitempty" xml:"third_part_invoice_id,omitempty"`
	// example:
	//
	// thridpart12138
	ThirdpartCostCenterId *string `json:"thirdpart_cost_center_id,omitempty" xml:"thirdpart_cost_center_id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0
	TrafficType *int32 `json:"traffic_type,omitempty" xml:"traffic_type,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0
	TripWay *int32 `json:"trip_way,omitempty" xml:"trip_way,omitempty"`
}

func (s ApplyAddRequestItineraryList) String() string {
	return dara.Prettify(s)
}

func (s ApplyAddRequestItineraryList) GoString() string {
	return s.String()
}

func (s *ApplyAddRequestItineraryList) GetArrCity() *string {
	return s.ArrCity
}

func (s *ApplyAddRequestItineraryList) GetArrCityCode() *string {
	return s.ArrCityCode
}

func (s *ApplyAddRequestItineraryList) GetArrDate() *string {
	return s.ArrDate
}

func (s *ApplyAddRequestItineraryList) GetAttribute() *string {
	return s.Attribute
}

func (s *ApplyAddRequestItineraryList) GetCostCenterId() *int64 {
	return s.CostCenterId
}

func (s *ApplyAddRequestItineraryList) GetDepCity() *string {
	return s.DepCity
}

func (s *ApplyAddRequestItineraryList) GetDepCityCode() *string {
	return s.DepCityCode
}

func (s *ApplyAddRequestItineraryList) GetDepDate() *string {
	return s.DepDate
}

func (s *ApplyAddRequestItineraryList) GetInvoiceId() *int64 {
	return s.InvoiceId
}

func (s *ApplyAddRequestItineraryList) GetItineraryId() *string {
	return s.ItineraryId
}

func (s *ApplyAddRequestItineraryList) GetItineraryTravelStandard() *ApplyAddRequestItineraryListItineraryTravelStandard {
	return s.ItineraryTravelStandard
}

func (s *ApplyAddRequestItineraryList) GetNeedHotel() *bool {
	return s.NeedHotel
}

func (s *ApplyAddRequestItineraryList) GetNeedTraffic() *bool {
	return s.NeedTraffic
}

func (s *ApplyAddRequestItineraryList) GetProjectCode() *string {
	return s.ProjectCode
}

func (s *ApplyAddRequestItineraryList) GetProjectTitle() *string {
	return s.ProjectTitle
}

func (s *ApplyAddRequestItineraryList) GetProvinceTravelCityAdcodes() []*string {
	return s.ProvinceTravelCityAdcodes
}

func (s *ApplyAddRequestItineraryList) GetThirdPartInvoiceId() *string {
	return s.ThirdPartInvoiceId
}

func (s *ApplyAddRequestItineraryList) GetThirdpartCostCenterId() *string {
	return s.ThirdpartCostCenterId
}

func (s *ApplyAddRequestItineraryList) GetTrafficType() *int32 {
	return s.TrafficType
}

func (s *ApplyAddRequestItineraryList) GetTripWay() *int32 {
	return s.TripWay
}

func (s *ApplyAddRequestItineraryList) SetArrCity(v string) *ApplyAddRequestItineraryList {
	s.ArrCity = &v
	return s
}

func (s *ApplyAddRequestItineraryList) SetArrCityCode(v string) *ApplyAddRequestItineraryList {
	s.ArrCityCode = &v
	return s
}

func (s *ApplyAddRequestItineraryList) SetArrDate(v string) *ApplyAddRequestItineraryList {
	s.ArrDate = &v
	return s
}

func (s *ApplyAddRequestItineraryList) SetAttribute(v string) *ApplyAddRequestItineraryList {
	s.Attribute = &v
	return s
}

func (s *ApplyAddRequestItineraryList) SetCostCenterId(v int64) *ApplyAddRequestItineraryList {
	s.CostCenterId = &v
	return s
}

func (s *ApplyAddRequestItineraryList) SetDepCity(v string) *ApplyAddRequestItineraryList {
	s.DepCity = &v
	return s
}

func (s *ApplyAddRequestItineraryList) SetDepCityCode(v string) *ApplyAddRequestItineraryList {
	s.DepCityCode = &v
	return s
}

func (s *ApplyAddRequestItineraryList) SetDepDate(v string) *ApplyAddRequestItineraryList {
	s.DepDate = &v
	return s
}

func (s *ApplyAddRequestItineraryList) SetInvoiceId(v int64) *ApplyAddRequestItineraryList {
	s.InvoiceId = &v
	return s
}

func (s *ApplyAddRequestItineraryList) SetItineraryId(v string) *ApplyAddRequestItineraryList {
	s.ItineraryId = &v
	return s
}

func (s *ApplyAddRequestItineraryList) SetItineraryTravelStandard(v *ApplyAddRequestItineraryListItineraryTravelStandard) *ApplyAddRequestItineraryList {
	s.ItineraryTravelStandard = v
	return s
}

func (s *ApplyAddRequestItineraryList) SetNeedHotel(v bool) *ApplyAddRequestItineraryList {
	s.NeedHotel = &v
	return s
}

func (s *ApplyAddRequestItineraryList) SetNeedTraffic(v bool) *ApplyAddRequestItineraryList {
	s.NeedTraffic = &v
	return s
}

func (s *ApplyAddRequestItineraryList) SetProjectCode(v string) *ApplyAddRequestItineraryList {
	s.ProjectCode = &v
	return s
}

func (s *ApplyAddRequestItineraryList) SetProjectTitle(v string) *ApplyAddRequestItineraryList {
	s.ProjectTitle = &v
	return s
}

func (s *ApplyAddRequestItineraryList) SetProvinceTravelCityAdcodes(v []*string) *ApplyAddRequestItineraryList {
	s.ProvinceTravelCityAdcodes = v
	return s
}

func (s *ApplyAddRequestItineraryList) SetThirdPartInvoiceId(v string) *ApplyAddRequestItineraryList {
	s.ThirdPartInvoiceId = &v
	return s
}

func (s *ApplyAddRequestItineraryList) SetThirdpartCostCenterId(v string) *ApplyAddRequestItineraryList {
	s.ThirdpartCostCenterId = &v
	return s
}

func (s *ApplyAddRequestItineraryList) SetTrafficType(v int32) *ApplyAddRequestItineraryList {
	s.TrafficType = &v
	return s
}

func (s *ApplyAddRequestItineraryList) SetTripWay(v int32) *ApplyAddRequestItineraryList {
	s.TripWay = &v
	return s
}

func (s *ApplyAddRequestItineraryList) Validate() error {
	if s.ItineraryTravelStandard != nil {
		if err := s.ItineraryTravelStandard.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ApplyAddRequestItineraryListItineraryTravelStandard struct {
	// example:
	//
	// 2
	HotelAvailableNightsPerDay *int32 `json:"hotel_available_nights_per_day,omitempty" xml:"hotel_available_nights_per_day,omitempty"`
}

func (s ApplyAddRequestItineraryListItineraryTravelStandard) String() string {
	return dara.Prettify(s)
}

func (s ApplyAddRequestItineraryListItineraryTravelStandard) GoString() string {
	return s.String()
}

func (s *ApplyAddRequestItineraryListItineraryTravelStandard) GetHotelAvailableNightsPerDay() *int32 {
	return s.HotelAvailableNightsPerDay
}

func (s *ApplyAddRequestItineraryListItineraryTravelStandard) SetHotelAvailableNightsPerDay(v int32) *ApplyAddRequestItineraryListItineraryTravelStandard {
	s.HotelAvailableNightsPerDay = &v
	return s
}

func (s *ApplyAddRequestItineraryListItineraryTravelStandard) Validate() error {
	return dara.Validate(s)
}

type ApplyAddRequestItinerarySetList struct {
	// This parameter is required.
	//
	// example:
	//
	// 2017-01-01 00:00:00
	ArrDate *string `json:"arr_date,omitempty" xml:"arr_date,omitempty"`
	// example:
	//
	// "{"name":"张三"}"
	Attribute *string `json:"attribute,omitempty" xml:"attribute,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// BJS,HGH
	CityCodeSet *string `json:"city_code_set,omitempty" xml:"city_code_set,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 北京，杭州
	CitySet *string `json:"city_set,omitempty" xml:"city_set,omitempty"`
	// example:
	//
	// 12345
	CostCenterId *int64 `json:"cost_center_id,omitempty" xml:"cost_center_id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2017-01-01 00:00:00
	DepDate *string `json:"dep_date,omitempty" xml:"dep_date,omitempty"`
	// example:
	//
	// 12345
	InvoiceId *int64 `json:"invoice_id,omitempty" xml:"invoice_id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123
	ItineraryId             *string                                                 `json:"itinerary_id,omitempty" xml:"itinerary_id,omitempty"`
	ItineraryTravelStandard *ApplyAddRequestItinerarySetListItineraryTravelStandard `json:"itinerary_travel_standard,omitempty" xml:"itinerary_travel_standard,omitempty" type:"Struct"`
	// example:
	//
	// projecttow
	ProjectCode *string `json:"project_code,omitempty" xml:"project_code,omitempty"`
	// example:
	//
	// 项目1
	ProjectTitle              *string   `json:"project_title,omitempty" xml:"project_title,omitempty"`
	ProvinceTravelCityAdcodes []*string `json:"province_travel_city_adcodes,omitempty" xml:"province_travel_city_adcodes,omitempty" type:"Repeated"`
	// example:
	//
	// thridpart12138
	ThirdPartInvoiceId *string `json:"third_part_invoice_id,omitempty" xml:"third_part_invoice_id,omitempty"`
	// example:
	//
	// thridpart12138
	ThirdpartCostCenterId *string `json:"thirdpart_cost_center_id,omitempty" xml:"thirdpart_cost_center_id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0
	TrafficType *int32 `json:"traffic_type,omitempty" xml:"traffic_type,omitempty"`
}

func (s ApplyAddRequestItinerarySetList) String() string {
	return dara.Prettify(s)
}

func (s ApplyAddRequestItinerarySetList) GoString() string {
	return s.String()
}

func (s *ApplyAddRequestItinerarySetList) GetArrDate() *string {
	return s.ArrDate
}

func (s *ApplyAddRequestItinerarySetList) GetAttribute() *string {
	return s.Attribute
}

func (s *ApplyAddRequestItinerarySetList) GetCityCodeSet() *string {
	return s.CityCodeSet
}

func (s *ApplyAddRequestItinerarySetList) GetCitySet() *string {
	return s.CitySet
}

func (s *ApplyAddRequestItinerarySetList) GetCostCenterId() *int64 {
	return s.CostCenterId
}

func (s *ApplyAddRequestItinerarySetList) GetDepDate() *string {
	return s.DepDate
}

func (s *ApplyAddRequestItinerarySetList) GetInvoiceId() *int64 {
	return s.InvoiceId
}

func (s *ApplyAddRequestItinerarySetList) GetItineraryId() *string {
	return s.ItineraryId
}

func (s *ApplyAddRequestItinerarySetList) GetItineraryTravelStandard() *ApplyAddRequestItinerarySetListItineraryTravelStandard {
	return s.ItineraryTravelStandard
}

func (s *ApplyAddRequestItinerarySetList) GetProjectCode() *string {
	return s.ProjectCode
}

func (s *ApplyAddRequestItinerarySetList) GetProjectTitle() *string {
	return s.ProjectTitle
}

func (s *ApplyAddRequestItinerarySetList) GetProvinceTravelCityAdcodes() []*string {
	return s.ProvinceTravelCityAdcodes
}

func (s *ApplyAddRequestItinerarySetList) GetThirdPartInvoiceId() *string {
	return s.ThirdPartInvoiceId
}

func (s *ApplyAddRequestItinerarySetList) GetThirdpartCostCenterId() *string {
	return s.ThirdpartCostCenterId
}

func (s *ApplyAddRequestItinerarySetList) GetTrafficType() *int32 {
	return s.TrafficType
}

func (s *ApplyAddRequestItinerarySetList) SetArrDate(v string) *ApplyAddRequestItinerarySetList {
	s.ArrDate = &v
	return s
}

func (s *ApplyAddRequestItinerarySetList) SetAttribute(v string) *ApplyAddRequestItinerarySetList {
	s.Attribute = &v
	return s
}

func (s *ApplyAddRequestItinerarySetList) SetCityCodeSet(v string) *ApplyAddRequestItinerarySetList {
	s.CityCodeSet = &v
	return s
}

func (s *ApplyAddRequestItinerarySetList) SetCitySet(v string) *ApplyAddRequestItinerarySetList {
	s.CitySet = &v
	return s
}

func (s *ApplyAddRequestItinerarySetList) SetCostCenterId(v int64) *ApplyAddRequestItinerarySetList {
	s.CostCenterId = &v
	return s
}

func (s *ApplyAddRequestItinerarySetList) SetDepDate(v string) *ApplyAddRequestItinerarySetList {
	s.DepDate = &v
	return s
}

func (s *ApplyAddRequestItinerarySetList) SetInvoiceId(v int64) *ApplyAddRequestItinerarySetList {
	s.InvoiceId = &v
	return s
}

func (s *ApplyAddRequestItinerarySetList) SetItineraryId(v string) *ApplyAddRequestItinerarySetList {
	s.ItineraryId = &v
	return s
}

func (s *ApplyAddRequestItinerarySetList) SetItineraryTravelStandard(v *ApplyAddRequestItinerarySetListItineraryTravelStandard) *ApplyAddRequestItinerarySetList {
	s.ItineraryTravelStandard = v
	return s
}

func (s *ApplyAddRequestItinerarySetList) SetProjectCode(v string) *ApplyAddRequestItinerarySetList {
	s.ProjectCode = &v
	return s
}

func (s *ApplyAddRequestItinerarySetList) SetProjectTitle(v string) *ApplyAddRequestItinerarySetList {
	s.ProjectTitle = &v
	return s
}

func (s *ApplyAddRequestItinerarySetList) SetProvinceTravelCityAdcodes(v []*string) *ApplyAddRequestItinerarySetList {
	s.ProvinceTravelCityAdcodes = v
	return s
}

func (s *ApplyAddRequestItinerarySetList) SetThirdPartInvoiceId(v string) *ApplyAddRequestItinerarySetList {
	s.ThirdPartInvoiceId = &v
	return s
}

func (s *ApplyAddRequestItinerarySetList) SetThirdpartCostCenterId(v string) *ApplyAddRequestItinerarySetList {
	s.ThirdpartCostCenterId = &v
	return s
}

func (s *ApplyAddRequestItinerarySetList) SetTrafficType(v int32) *ApplyAddRequestItinerarySetList {
	s.TrafficType = &v
	return s
}

func (s *ApplyAddRequestItinerarySetList) Validate() error {
	if s.ItineraryTravelStandard != nil {
		if err := s.ItineraryTravelStandard.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ApplyAddRequestItinerarySetListItineraryTravelStandard struct {
	// example:
	//
	// 2
	HotelAvailableNightsPerDay *int32 `json:"hotel_available_nights_per_day,omitempty" xml:"hotel_available_nights_per_day,omitempty"`
}

func (s ApplyAddRequestItinerarySetListItineraryTravelStandard) String() string {
	return dara.Prettify(s)
}

func (s ApplyAddRequestItinerarySetListItineraryTravelStandard) GoString() string {
	return s.String()
}

func (s *ApplyAddRequestItinerarySetListItineraryTravelStandard) GetHotelAvailableNightsPerDay() *int32 {
	return s.HotelAvailableNightsPerDay
}

func (s *ApplyAddRequestItinerarySetListItineraryTravelStandard) SetHotelAvailableNightsPerDay(v int32) *ApplyAddRequestItinerarySetListItineraryTravelStandard {
	s.HotelAvailableNightsPerDay = &v
	return s
}

func (s *ApplyAddRequestItinerarySetListItineraryTravelStandard) Validate() error {
	return dara.Validate(s)
}

type ApplyAddRequestTravelerList struct {
	// example:
	//
	// “{"name":"张三"}”
	Attribute *string `json:"attribute,omitempty" xml:"attribute,omitempty"`
	// example:
	//
	// 112711
	CostCenterId *int64 `json:"cost_center_id,omitempty" xml:"cost_center_id,omitempty"`
	// example:
	//
	// 11251
	InvoiceId *int64 `json:"invoice_id,omitempty" xml:"invoice_id,omitempty"`
	// example:
	//
	// 1142
	PaymentDepartmentId *string `json:"payment_department_id,omitempty" xml:"payment_department_id,omitempty"`
	// example:
	//
	// 产品部
	PaymentDepartmentName *string `json:"payment_department_name,omitempty" xml:"payment_department_name,omitempty"`
	// example:
	//
	// acs
	ProjectCode *string `json:"project_code,omitempty" xml:"project_code,omitempty"`
	// example:
	//
	// "成本项目"
	ProjectTitle *string `json:"project_title,omitempty" xml:"project_title,omitempty"`
	// example:
	//
	// 517492
	ThirdPartInvoiceId *string `json:"third_part_invoice_id,omitempty" xml:"third_part_invoice_id,omitempty"`
	// example:
	//
	// 441154
	ThirdpartCostCenterId *string `json:"thirdpart_cost_center_id,omitempty" xml:"thirdpart_cost_center_id,omitempty"`
	// example:
	//
	// 101128
	ThirdpartDepartId *string `json:"thirdpart_depart_id,omitempty" xml:"thirdpart_depart_id,omitempty"`
	// example:
	//
	// wu51531
	UserId *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// example:
	//
	// 王武
	UserName *string `json:"user_name,omitempty" xml:"user_name,omitempty"`
}

func (s ApplyAddRequestTravelerList) String() string {
	return dara.Prettify(s)
}

func (s ApplyAddRequestTravelerList) GoString() string {
	return s.String()
}

func (s *ApplyAddRequestTravelerList) GetAttribute() *string {
	return s.Attribute
}

func (s *ApplyAddRequestTravelerList) GetCostCenterId() *int64 {
	return s.CostCenterId
}

func (s *ApplyAddRequestTravelerList) GetInvoiceId() *int64 {
	return s.InvoiceId
}

func (s *ApplyAddRequestTravelerList) GetPaymentDepartmentId() *string {
	return s.PaymentDepartmentId
}

func (s *ApplyAddRequestTravelerList) GetPaymentDepartmentName() *string {
	return s.PaymentDepartmentName
}

func (s *ApplyAddRequestTravelerList) GetProjectCode() *string {
	return s.ProjectCode
}

func (s *ApplyAddRequestTravelerList) GetProjectTitle() *string {
	return s.ProjectTitle
}

func (s *ApplyAddRequestTravelerList) GetThirdPartInvoiceId() *string {
	return s.ThirdPartInvoiceId
}

func (s *ApplyAddRequestTravelerList) GetThirdpartCostCenterId() *string {
	return s.ThirdpartCostCenterId
}

func (s *ApplyAddRequestTravelerList) GetThirdpartDepartId() *string {
	return s.ThirdpartDepartId
}

func (s *ApplyAddRequestTravelerList) GetUserId() *string {
	return s.UserId
}

func (s *ApplyAddRequestTravelerList) GetUserName() *string {
	return s.UserName
}

func (s *ApplyAddRequestTravelerList) SetAttribute(v string) *ApplyAddRequestTravelerList {
	s.Attribute = &v
	return s
}

func (s *ApplyAddRequestTravelerList) SetCostCenterId(v int64) *ApplyAddRequestTravelerList {
	s.CostCenterId = &v
	return s
}

func (s *ApplyAddRequestTravelerList) SetInvoiceId(v int64) *ApplyAddRequestTravelerList {
	s.InvoiceId = &v
	return s
}

func (s *ApplyAddRequestTravelerList) SetPaymentDepartmentId(v string) *ApplyAddRequestTravelerList {
	s.PaymentDepartmentId = &v
	return s
}

func (s *ApplyAddRequestTravelerList) SetPaymentDepartmentName(v string) *ApplyAddRequestTravelerList {
	s.PaymentDepartmentName = &v
	return s
}

func (s *ApplyAddRequestTravelerList) SetProjectCode(v string) *ApplyAddRequestTravelerList {
	s.ProjectCode = &v
	return s
}

func (s *ApplyAddRequestTravelerList) SetProjectTitle(v string) *ApplyAddRequestTravelerList {
	s.ProjectTitle = &v
	return s
}

func (s *ApplyAddRequestTravelerList) SetThirdPartInvoiceId(v string) *ApplyAddRequestTravelerList {
	s.ThirdPartInvoiceId = &v
	return s
}

func (s *ApplyAddRequestTravelerList) SetThirdpartCostCenterId(v string) *ApplyAddRequestTravelerList {
	s.ThirdpartCostCenterId = &v
	return s
}

func (s *ApplyAddRequestTravelerList) SetThirdpartDepartId(v string) *ApplyAddRequestTravelerList {
	s.ThirdpartDepartId = &v
	return s
}

func (s *ApplyAddRequestTravelerList) SetUserId(v string) *ApplyAddRequestTravelerList {
	s.UserId = &v
	return s
}

func (s *ApplyAddRequestTravelerList) SetUserName(v string) *ApplyAddRequestTravelerList {
	s.UserName = &v
	return s
}

func (s *ApplyAddRequestTravelerList) Validate() error {
	return dara.Validate(s)
}

type ApplyAddRequestTravelerStandard struct {
	// example:
	//
	// 1
	BusinessDiscount *int32                                       `json:"business_discount,omitempty" xml:"business_discount,omitempty"`
	CarCitySet       []*ApplyAddRequestTravelerStandardCarCitySet `json:"car_city_set,omitempty" xml:"car_city_set,omitempty" type:"Repeated"`
	CarStandard      *ApplyAddRequestTravelerStandardCarStandard  `json:"car_standard,omitempty" xml:"car_standard,omitempty" type:"Struct"`
	// example:
	//
	// 1
	EconomyDiscount *int32 `json:"economy_discount,omitempty" xml:"economy_discount,omitempty"`
	// example:
	//
	// 1
	FirstDiscount *int32 `json:"first_discount,omitempty" xml:"first_discount,omitempty"`
	// example:
	//
	// Y
	FlightCabins *string `json:"flight_cabins,omitempty" xml:"flight_cabins,omitempty"`
	// example:
	//
	// 100132
	FlightIntlRuleCode *int64 `json:"flight_intl_rule_code,omitempty" xml:"flight_intl_rule_code,omitempty"`
	// example:
	//
	// 100132
	FlightRuleCode *int64                                           `json:"flight_rule_code,omitempty" xml:"flight_rule_code,omitempty"`
	HotelCitys     []*ApplyAddRequestTravelerStandardHotelCitys     `json:"hotel_citys,omitempty" xml:"hotel_citys,omitempty" type:"Repeated"`
	HotelIntlCitys []*ApplyAddRequestTravelerStandardHotelIntlCitys `json:"hotel_intl_citys,omitempty" xml:"hotel_intl_citys,omitempty" type:"Repeated"`
	// example:
	//
	// 100132
	HotelIntlRuleCode *int64 `json:"hotel_intl_rule_code,omitempty" xml:"hotel_intl_rule_code,omitempty"`
	// example:
	//
	// 100132
	HotelRuleCode *int64 `json:"hotel_rule_code,omitempty" xml:"hotel_rule_code,omitempty"`
	// example:
	//
	// F
	InternationalFlightCabins *string `json:"international_flight_cabins,omitempty" xml:"international_flight_cabins,omitempty"`
	// example:
	//
	// 1
	PremiumEconomyDiscount *int32 `json:"premium_economy_discount,omitempty" xml:"premium_economy_discount,omitempty"`
	// example:
	//
	// 0
	ReserveType *int32 `json:"reserve_type,omitempty" xml:"reserve_type,omitempty"`
	// example:
	//
	// 100132
	TrainRuleCode *int64 `json:"train_rule_code,omitempty" xml:"train_rule_code,omitempty"`
	// example:
	//
	// 1
	TrainSeats *string `json:"train_seats,omitempty" xml:"train_seats,omitempty"`
	// example:
	//
	// wfffeng
	UserId *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

func (s ApplyAddRequestTravelerStandard) String() string {
	return dara.Prettify(s)
}

func (s ApplyAddRequestTravelerStandard) GoString() string {
	return s.String()
}

func (s *ApplyAddRequestTravelerStandard) GetBusinessDiscount() *int32 {
	return s.BusinessDiscount
}

func (s *ApplyAddRequestTravelerStandard) GetCarCitySet() []*ApplyAddRequestTravelerStandardCarCitySet {
	return s.CarCitySet
}

func (s *ApplyAddRequestTravelerStandard) GetCarStandard() *ApplyAddRequestTravelerStandardCarStandard {
	return s.CarStandard
}

func (s *ApplyAddRequestTravelerStandard) GetEconomyDiscount() *int32 {
	return s.EconomyDiscount
}

func (s *ApplyAddRequestTravelerStandard) GetFirstDiscount() *int32 {
	return s.FirstDiscount
}

func (s *ApplyAddRequestTravelerStandard) GetFlightCabins() *string {
	return s.FlightCabins
}

func (s *ApplyAddRequestTravelerStandard) GetFlightIntlRuleCode() *int64 {
	return s.FlightIntlRuleCode
}

func (s *ApplyAddRequestTravelerStandard) GetFlightRuleCode() *int64 {
	return s.FlightRuleCode
}

func (s *ApplyAddRequestTravelerStandard) GetHotelCitys() []*ApplyAddRequestTravelerStandardHotelCitys {
	return s.HotelCitys
}

func (s *ApplyAddRequestTravelerStandard) GetHotelIntlCitys() []*ApplyAddRequestTravelerStandardHotelIntlCitys {
	return s.HotelIntlCitys
}

func (s *ApplyAddRequestTravelerStandard) GetHotelIntlRuleCode() *int64 {
	return s.HotelIntlRuleCode
}

func (s *ApplyAddRequestTravelerStandard) GetHotelRuleCode() *int64 {
	return s.HotelRuleCode
}

func (s *ApplyAddRequestTravelerStandard) GetInternationalFlightCabins() *string {
	return s.InternationalFlightCabins
}

func (s *ApplyAddRequestTravelerStandard) GetPremiumEconomyDiscount() *int32 {
	return s.PremiumEconomyDiscount
}

func (s *ApplyAddRequestTravelerStandard) GetReserveType() *int32 {
	return s.ReserveType
}

func (s *ApplyAddRequestTravelerStandard) GetTrainRuleCode() *int64 {
	return s.TrainRuleCode
}

func (s *ApplyAddRequestTravelerStandard) GetTrainSeats() *string {
	return s.TrainSeats
}

func (s *ApplyAddRequestTravelerStandard) GetUserId() *string {
	return s.UserId
}

func (s *ApplyAddRequestTravelerStandard) SetBusinessDiscount(v int32) *ApplyAddRequestTravelerStandard {
	s.BusinessDiscount = &v
	return s
}

func (s *ApplyAddRequestTravelerStandard) SetCarCitySet(v []*ApplyAddRequestTravelerStandardCarCitySet) *ApplyAddRequestTravelerStandard {
	s.CarCitySet = v
	return s
}

func (s *ApplyAddRequestTravelerStandard) SetCarStandard(v *ApplyAddRequestTravelerStandardCarStandard) *ApplyAddRequestTravelerStandard {
	s.CarStandard = v
	return s
}

func (s *ApplyAddRequestTravelerStandard) SetEconomyDiscount(v int32) *ApplyAddRequestTravelerStandard {
	s.EconomyDiscount = &v
	return s
}

func (s *ApplyAddRequestTravelerStandard) SetFirstDiscount(v int32) *ApplyAddRequestTravelerStandard {
	s.FirstDiscount = &v
	return s
}

func (s *ApplyAddRequestTravelerStandard) SetFlightCabins(v string) *ApplyAddRequestTravelerStandard {
	s.FlightCabins = &v
	return s
}

func (s *ApplyAddRequestTravelerStandard) SetFlightIntlRuleCode(v int64) *ApplyAddRequestTravelerStandard {
	s.FlightIntlRuleCode = &v
	return s
}

func (s *ApplyAddRequestTravelerStandard) SetFlightRuleCode(v int64) *ApplyAddRequestTravelerStandard {
	s.FlightRuleCode = &v
	return s
}

func (s *ApplyAddRequestTravelerStandard) SetHotelCitys(v []*ApplyAddRequestTravelerStandardHotelCitys) *ApplyAddRequestTravelerStandard {
	s.HotelCitys = v
	return s
}

func (s *ApplyAddRequestTravelerStandard) SetHotelIntlCitys(v []*ApplyAddRequestTravelerStandardHotelIntlCitys) *ApplyAddRequestTravelerStandard {
	s.HotelIntlCitys = v
	return s
}

func (s *ApplyAddRequestTravelerStandard) SetHotelIntlRuleCode(v int64) *ApplyAddRequestTravelerStandard {
	s.HotelIntlRuleCode = &v
	return s
}

func (s *ApplyAddRequestTravelerStandard) SetHotelRuleCode(v int64) *ApplyAddRequestTravelerStandard {
	s.HotelRuleCode = &v
	return s
}

func (s *ApplyAddRequestTravelerStandard) SetInternationalFlightCabins(v string) *ApplyAddRequestTravelerStandard {
	s.InternationalFlightCabins = &v
	return s
}

func (s *ApplyAddRequestTravelerStandard) SetPremiumEconomyDiscount(v int32) *ApplyAddRequestTravelerStandard {
	s.PremiumEconomyDiscount = &v
	return s
}

func (s *ApplyAddRequestTravelerStandard) SetReserveType(v int32) *ApplyAddRequestTravelerStandard {
	s.ReserveType = &v
	return s
}

func (s *ApplyAddRequestTravelerStandard) SetTrainRuleCode(v int64) *ApplyAddRequestTravelerStandard {
	s.TrainRuleCode = &v
	return s
}

func (s *ApplyAddRequestTravelerStandard) SetTrainSeats(v string) *ApplyAddRequestTravelerStandard {
	s.TrainSeats = &v
	return s
}

func (s *ApplyAddRequestTravelerStandard) SetUserId(v string) *ApplyAddRequestTravelerStandard {
	s.UserId = &v
	return s
}

func (s *ApplyAddRequestTravelerStandard) Validate() error {
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

type ApplyAddRequestTravelerStandardCarCitySet struct {
	// This parameter is required.
	//
	// example:
	//
	// 110100，330100
	CityCode *string `json:"city_code,omitempty" xml:"city_code,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 北京，杭州
	CityName *string `json:"city_name,omitempty" xml:"city_name,omitempty"`
}

func (s ApplyAddRequestTravelerStandardCarCitySet) String() string {
	return dara.Prettify(s)
}

func (s ApplyAddRequestTravelerStandardCarCitySet) GoString() string {
	return s.String()
}

func (s *ApplyAddRequestTravelerStandardCarCitySet) GetCityCode() *string {
	return s.CityCode
}

func (s *ApplyAddRequestTravelerStandardCarCitySet) GetCityName() *string {
	return s.CityName
}

func (s *ApplyAddRequestTravelerStandardCarCitySet) SetCityCode(v string) *ApplyAddRequestTravelerStandardCarCitySet {
	s.CityCode = &v
	return s
}

func (s *ApplyAddRequestTravelerStandardCarCitySet) SetCityName(v string) *ApplyAddRequestTravelerStandardCarCitySet {
	s.CityName = &v
	return s
}

func (s *ApplyAddRequestTravelerStandardCarCitySet) Validate() error {
	return dara.Validate(s)
}

type ApplyAddRequestTravelerStandardCarStandard struct {
	BookAllowInfo         *ApplyAddRequestTravelerStandardCarStandardBookAllowInfo         `json:"book_allow_info,omitempty" xml:"book_allow_info,omitempty" type:"Struct"`
	CarHelper             *ApplyAddRequestTravelerStandardCarStandardCarHelper             `json:"car_helper,omitempty" xml:"car_helper,omitempty" type:"Struct"`
	CarTimeControl        *ApplyAddRequestTravelerStandardCarStandardCarTimeControl        `json:"car_time_control,omitempty" xml:"car_time_control,omitempty" type:"Struct"`
	CityControlInfo       *ApplyAddRequestTravelerStandardCarStandardCityControlInfo       `json:"city_control_info,omitempty" xml:"city_control_info,omitempty" type:"Struct"`
	CrossCityInfo         *ApplyAddRequestTravelerStandardCarStandardCrossCityInfo         `json:"cross_city_info,omitempty" xml:"cross_city_info,omitempty" type:"Struct"`
	ElectronicFenceInfo   *ApplyAddRequestTravelerStandardCarStandardElectronicFenceInfo   `json:"electronic_fence_info,omitempty" xml:"electronic_fence_info,omitempty" type:"Struct"`
	LevelCodes            *string                                                          `json:"level_codes,omitempty" xml:"level_codes,omitempty"`
	ModifyDestinationInfo *ApplyAddRequestTravelerStandardCarStandardModifyDestinationInfo `json:"modify_destination_info,omitempty" xml:"modify_destination_info,omitempty" type:"Struct"`
	TimesTotal            *int32                                                           `json:"times_total,omitempty" xml:"times_total,omitempty"`
	TimesType             *int32                                                           `json:"times_type,omitempty" xml:"times_type,omitempty"`
}

func (s ApplyAddRequestTravelerStandardCarStandard) String() string {
	return dara.Prettify(s)
}

func (s ApplyAddRequestTravelerStandardCarStandard) GoString() string {
	return s.String()
}

func (s *ApplyAddRequestTravelerStandardCarStandard) GetBookAllowInfo() *ApplyAddRequestTravelerStandardCarStandardBookAllowInfo {
	return s.BookAllowInfo
}

func (s *ApplyAddRequestTravelerStandardCarStandard) GetCarHelper() *ApplyAddRequestTravelerStandardCarStandardCarHelper {
	return s.CarHelper
}

func (s *ApplyAddRequestTravelerStandardCarStandard) GetCarTimeControl() *ApplyAddRequestTravelerStandardCarStandardCarTimeControl {
	return s.CarTimeControl
}

func (s *ApplyAddRequestTravelerStandardCarStandard) GetCityControlInfo() *ApplyAddRequestTravelerStandardCarStandardCityControlInfo {
	return s.CityControlInfo
}

func (s *ApplyAddRequestTravelerStandardCarStandard) GetCrossCityInfo() *ApplyAddRequestTravelerStandardCarStandardCrossCityInfo {
	return s.CrossCityInfo
}

func (s *ApplyAddRequestTravelerStandardCarStandard) GetElectronicFenceInfo() *ApplyAddRequestTravelerStandardCarStandardElectronicFenceInfo {
	return s.ElectronicFenceInfo
}

func (s *ApplyAddRequestTravelerStandardCarStandard) GetLevelCodes() *string {
	return s.LevelCodes
}

func (s *ApplyAddRequestTravelerStandardCarStandard) GetModifyDestinationInfo() *ApplyAddRequestTravelerStandardCarStandardModifyDestinationInfo {
	return s.ModifyDestinationInfo
}

func (s *ApplyAddRequestTravelerStandardCarStandard) GetTimesTotal() *int32 {
	return s.TimesTotal
}

func (s *ApplyAddRequestTravelerStandardCarStandard) GetTimesType() *int32 {
	return s.TimesType
}

func (s *ApplyAddRequestTravelerStandardCarStandard) SetBookAllowInfo(v *ApplyAddRequestTravelerStandardCarStandardBookAllowInfo) *ApplyAddRequestTravelerStandardCarStandard {
	s.BookAllowInfo = v
	return s
}

func (s *ApplyAddRequestTravelerStandardCarStandard) SetCarHelper(v *ApplyAddRequestTravelerStandardCarStandardCarHelper) *ApplyAddRequestTravelerStandardCarStandard {
	s.CarHelper = v
	return s
}

func (s *ApplyAddRequestTravelerStandardCarStandard) SetCarTimeControl(v *ApplyAddRequestTravelerStandardCarStandardCarTimeControl) *ApplyAddRequestTravelerStandardCarStandard {
	s.CarTimeControl = v
	return s
}

func (s *ApplyAddRequestTravelerStandardCarStandard) SetCityControlInfo(v *ApplyAddRequestTravelerStandardCarStandardCityControlInfo) *ApplyAddRequestTravelerStandardCarStandard {
	s.CityControlInfo = v
	return s
}

func (s *ApplyAddRequestTravelerStandardCarStandard) SetCrossCityInfo(v *ApplyAddRequestTravelerStandardCarStandardCrossCityInfo) *ApplyAddRequestTravelerStandardCarStandard {
	s.CrossCityInfo = v
	return s
}

func (s *ApplyAddRequestTravelerStandardCarStandard) SetElectronicFenceInfo(v *ApplyAddRequestTravelerStandardCarStandardElectronicFenceInfo) *ApplyAddRequestTravelerStandardCarStandard {
	s.ElectronicFenceInfo = v
	return s
}

func (s *ApplyAddRequestTravelerStandardCarStandard) SetLevelCodes(v string) *ApplyAddRequestTravelerStandardCarStandard {
	s.LevelCodes = &v
	return s
}

func (s *ApplyAddRequestTravelerStandardCarStandard) SetModifyDestinationInfo(v *ApplyAddRequestTravelerStandardCarStandardModifyDestinationInfo) *ApplyAddRequestTravelerStandardCarStandard {
	s.ModifyDestinationInfo = v
	return s
}

func (s *ApplyAddRequestTravelerStandardCarStandard) SetTimesTotal(v int32) *ApplyAddRequestTravelerStandardCarStandard {
	s.TimesTotal = &v
	return s
}

func (s *ApplyAddRequestTravelerStandardCarStandard) SetTimesType(v int32) *ApplyAddRequestTravelerStandardCarStandard {
	s.TimesType = &v
	return s
}

func (s *ApplyAddRequestTravelerStandardCarStandard) Validate() error {
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

type ApplyAddRequestTravelerStandardCarStandardBookAllowInfo struct {
	// This parameter is required.
	BookAllow *bool `json:"book_allow,omitempty" xml:"book_allow,omitempty"`
}

func (s ApplyAddRequestTravelerStandardCarStandardBookAllowInfo) String() string {
	return dara.Prettify(s)
}

func (s ApplyAddRequestTravelerStandardCarStandardBookAllowInfo) GoString() string {
	return s.String()
}

func (s *ApplyAddRequestTravelerStandardCarStandardBookAllowInfo) GetBookAllow() *bool {
	return s.BookAllow
}

func (s *ApplyAddRequestTravelerStandardCarStandardBookAllowInfo) SetBookAllow(v bool) *ApplyAddRequestTravelerStandardCarStandardBookAllowInfo {
	s.BookAllow = &v
	return s
}

func (s *ApplyAddRequestTravelerStandardCarStandardBookAllowInfo) Validate() error {
	return dara.Validate(s)
}

type ApplyAddRequestTravelerStandardCarStandardCarHelper struct {
	// This parameter is required.
	CarHelperType *string `json:"car_helper_type,omitempty" xml:"car_helper_type,omitempty"`
}

func (s ApplyAddRequestTravelerStandardCarStandardCarHelper) String() string {
	return dara.Prettify(s)
}

func (s ApplyAddRequestTravelerStandardCarStandardCarHelper) GoString() string {
	return s.String()
}

func (s *ApplyAddRequestTravelerStandardCarStandardCarHelper) GetCarHelperType() *string {
	return s.CarHelperType
}

func (s *ApplyAddRequestTravelerStandardCarStandardCarHelper) SetCarHelperType(v string) *ApplyAddRequestTravelerStandardCarStandardCarHelper {
	s.CarHelperType = &v
	return s
}

func (s *ApplyAddRequestTravelerStandardCarStandardCarHelper) Validate() error {
	return dara.Validate(s)
}

type ApplyAddRequestTravelerStandardCarStandardCarTimeControl struct {
	// This parameter is required.
	TimeLimit []*ApplyAddRequestTravelerStandardCarStandardCarTimeControlTimeLimit `json:"time_limit,omitempty" xml:"time_limit,omitempty" type:"Repeated"`
	// This parameter is required.
	TimeSwitch *bool `json:"time_switch,omitempty" xml:"time_switch,omitempty"`
}

func (s ApplyAddRequestTravelerStandardCarStandardCarTimeControl) String() string {
	return dara.Prettify(s)
}

func (s ApplyAddRequestTravelerStandardCarStandardCarTimeControl) GoString() string {
	return s.String()
}

func (s *ApplyAddRequestTravelerStandardCarStandardCarTimeControl) GetTimeLimit() []*ApplyAddRequestTravelerStandardCarStandardCarTimeControlTimeLimit {
	return s.TimeLimit
}

func (s *ApplyAddRequestTravelerStandardCarStandardCarTimeControl) GetTimeSwitch() *bool {
	return s.TimeSwitch
}

func (s *ApplyAddRequestTravelerStandardCarStandardCarTimeControl) SetTimeLimit(v []*ApplyAddRequestTravelerStandardCarStandardCarTimeControlTimeLimit) *ApplyAddRequestTravelerStandardCarStandardCarTimeControl {
	s.TimeLimit = v
	return s
}

func (s *ApplyAddRequestTravelerStandardCarStandardCarTimeControl) SetTimeSwitch(v bool) *ApplyAddRequestTravelerStandardCarStandardCarTimeControl {
	s.TimeSwitch = &v
	return s
}

func (s *ApplyAddRequestTravelerStandardCarStandardCarTimeControl) Validate() error {
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

type ApplyAddRequestTravelerStandardCarStandardCarTimeControlTimeLimit struct {
	// This parameter is required.
	EndTime *string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// This parameter is required.
	StartTime *string `json:"start_time,omitempty" xml:"start_time,omitempty"`
}

func (s ApplyAddRequestTravelerStandardCarStandardCarTimeControlTimeLimit) String() string {
	return dara.Prettify(s)
}

func (s ApplyAddRequestTravelerStandardCarStandardCarTimeControlTimeLimit) GoString() string {
	return s.String()
}

func (s *ApplyAddRequestTravelerStandardCarStandardCarTimeControlTimeLimit) GetEndTime() *string {
	return s.EndTime
}

func (s *ApplyAddRequestTravelerStandardCarStandardCarTimeControlTimeLimit) GetStartTime() *string {
	return s.StartTime
}

func (s *ApplyAddRequestTravelerStandardCarStandardCarTimeControlTimeLimit) SetEndTime(v string) *ApplyAddRequestTravelerStandardCarStandardCarTimeControlTimeLimit {
	s.EndTime = &v
	return s
}

func (s *ApplyAddRequestTravelerStandardCarStandardCarTimeControlTimeLimit) SetStartTime(v string) *ApplyAddRequestTravelerStandardCarStandardCarTimeControlTimeLimit {
	s.StartTime = &v
	return s
}

func (s *ApplyAddRequestTravelerStandardCarStandardCarTimeControlTimeLimit) Validate() error {
	return dara.Validate(s)
}

type ApplyAddRequestTravelerStandardCarStandardCityControlInfo struct {
	// This parameter is required.
	CityControlType *int32 `json:"city_control_type,omitempty" xml:"city_control_type,omitempty"`
	// This parameter is required.
	CityInfos []*ApplyAddRequestTravelerStandardCarStandardCityControlInfoCityInfos `json:"city_infos,omitempty" xml:"city_infos,omitempty" type:"Repeated"`
}

func (s ApplyAddRequestTravelerStandardCarStandardCityControlInfo) String() string {
	return dara.Prettify(s)
}

func (s ApplyAddRequestTravelerStandardCarStandardCityControlInfo) GoString() string {
	return s.String()
}

func (s *ApplyAddRequestTravelerStandardCarStandardCityControlInfo) GetCityControlType() *int32 {
	return s.CityControlType
}

func (s *ApplyAddRequestTravelerStandardCarStandardCityControlInfo) GetCityInfos() []*ApplyAddRequestTravelerStandardCarStandardCityControlInfoCityInfos {
	return s.CityInfos
}

func (s *ApplyAddRequestTravelerStandardCarStandardCityControlInfo) SetCityControlType(v int32) *ApplyAddRequestTravelerStandardCarStandardCityControlInfo {
	s.CityControlType = &v
	return s
}

func (s *ApplyAddRequestTravelerStandardCarStandardCityControlInfo) SetCityInfos(v []*ApplyAddRequestTravelerStandardCarStandardCityControlInfoCityInfos) *ApplyAddRequestTravelerStandardCarStandardCityControlInfo {
	s.CityInfos = v
	return s
}

func (s *ApplyAddRequestTravelerStandardCarStandardCityControlInfo) Validate() error {
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

type ApplyAddRequestTravelerStandardCarStandardCityControlInfoCityInfos struct {
	// This parameter is required.
	Adcode *string `json:"adcode,omitempty" xml:"adcode,omitempty"`
	// This parameter is required.
	CityCode *string `json:"city_code,omitempty" xml:"city_code,omitempty"`
	// This parameter is required.
	CityName *string `json:"city_name,omitempty" xml:"city_name,omitempty"`
}

func (s ApplyAddRequestTravelerStandardCarStandardCityControlInfoCityInfos) String() string {
	return dara.Prettify(s)
}

func (s ApplyAddRequestTravelerStandardCarStandardCityControlInfoCityInfos) GoString() string {
	return s.String()
}

func (s *ApplyAddRequestTravelerStandardCarStandardCityControlInfoCityInfos) GetAdcode() *string {
	return s.Adcode
}

func (s *ApplyAddRequestTravelerStandardCarStandardCityControlInfoCityInfos) GetCityCode() *string {
	return s.CityCode
}

func (s *ApplyAddRequestTravelerStandardCarStandardCityControlInfoCityInfos) GetCityName() *string {
	return s.CityName
}

func (s *ApplyAddRequestTravelerStandardCarStandardCityControlInfoCityInfos) SetAdcode(v string) *ApplyAddRequestTravelerStandardCarStandardCityControlInfoCityInfos {
	s.Adcode = &v
	return s
}

func (s *ApplyAddRequestTravelerStandardCarStandardCityControlInfoCityInfos) SetCityCode(v string) *ApplyAddRequestTravelerStandardCarStandardCityControlInfoCityInfos {
	s.CityCode = &v
	return s
}

func (s *ApplyAddRequestTravelerStandardCarStandardCityControlInfoCityInfos) SetCityName(v string) *ApplyAddRequestTravelerStandardCarStandardCityControlInfoCityInfos {
	s.CityName = &v
	return s
}

func (s *ApplyAddRequestTravelerStandardCarStandardCityControlInfoCityInfos) Validate() error {
	return dara.Validate(s)
}

type ApplyAddRequestTravelerStandardCarStandardCrossCityInfo struct {
	CrossCityList []*ApplyAddRequestTravelerStandardCarStandardCrossCityInfoCrossCityList `json:"cross_city_list,omitempty" xml:"cross_city_list,omitempty" type:"Repeated"`
	// This parameter is required.
	CrossCityType *int32 `json:"cross_city_type,omitempty" xml:"cross_city_type,omitempty"`
}

func (s ApplyAddRequestTravelerStandardCarStandardCrossCityInfo) String() string {
	return dara.Prettify(s)
}

func (s ApplyAddRequestTravelerStandardCarStandardCrossCityInfo) GoString() string {
	return s.String()
}

func (s *ApplyAddRequestTravelerStandardCarStandardCrossCityInfo) GetCrossCityList() []*ApplyAddRequestTravelerStandardCarStandardCrossCityInfoCrossCityList {
	return s.CrossCityList
}

func (s *ApplyAddRequestTravelerStandardCarStandardCrossCityInfo) GetCrossCityType() *int32 {
	return s.CrossCityType
}

func (s *ApplyAddRequestTravelerStandardCarStandardCrossCityInfo) SetCrossCityList(v []*ApplyAddRequestTravelerStandardCarStandardCrossCityInfoCrossCityList) *ApplyAddRequestTravelerStandardCarStandardCrossCityInfo {
	s.CrossCityList = v
	return s
}

func (s *ApplyAddRequestTravelerStandardCarStandardCrossCityInfo) SetCrossCityType(v int32) *ApplyAddRequestTravelerStandardCarStandardCrossCityInfo {
	s.CrossCityType = &v
	return s
}

func (s *ApplyAddRequestTravelerStandardCarStandardCrossCityInfo) Validate() error {
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

type ApplyAddRequestTravelerStandardCarStandardCrossCityInfoCrossCityList struct {
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

func (s ApplyAddRequestTravelerStandardCarStandardCrossCityInfoCrossCityList) String() string {
	return dara.Prettify(s)
}

func (s ApplyAddRequestTravelerStandardCarStandardCrossCityInfoCrossCityList) GoString() string {
	return s.String()
}

func (s *ApplyAddRequestTravelerStandardCarStandardCrossCityInfoCrossCityList) GetFromAdcode() *string {
	return s.FromAdcode
}

func (s *ApplyAddRequestTravelerStandardCarStandardCrossCityInfoCrossCityList) GetFromCityCode() *string {
	return s.FromCityCode
}

func (s *ApplyAddRequestTravelerStandardCarStandardCrossCityInfoCrossCityList) GetFromCityName() *string {
	return s.FromCityName
}

func (s *ApplyAddRequestTravelerStandardCarStandardCrossCityInfoCrossCityList) GetToAdcode() *string {
	return s.ToAdcode
}

func (s *ApplyAddRequestTravelerStandardCarStandardCrossCityInfoCrossCityList) GetToCityCode() *string {
	return s.ToCityCode
}

func (s *ApplyAddRequestTravelerStandardCarStandardCrossCityInfoCrossCityList) GetToCityName() *string {
	return s.ToCityName
}

func (s *ApplyAddRequestTravelerStandardCarStandardCrossCityInfoCrossCityList) SetFromAdcode(v string) *ApplyAddRequestTravelerStandardCarStandardCrossCityInfoCrossCityList {
	s.FromAdcode = &v
	return s
}

func (s *ApplyAddRequestTravelerStandardCarStandardCrossCityInfoCrossCityList) SetFromCityCode(v string) *ApplyAddRequestTravelerStandardCarStandardCrossCityInfoCrossCityList {
	s.FromCityCode = &v
	return s
}

func (s *ApplyAddRequestTravelerStandardCarStandardCrossCityInfoCrossCityList) SetFromCityName(v string) *ApplyAddRequestTravelerStandardCarStandardCrossCityInfoCrossCityList {
	s.FromCityName = &v
	return s
}

func (s *ApplyAddRequestTravelerStandardCarStandardCrossCityInfoCrossCityList) SetToAdcode(v string) *ApplyAddRequestTravelerStandardCarStandardCrossCityInfoCrossCityList {
	s.ToAdcode = &v
	return s
}

func (s *ApplyAddRequestTravelerStandardCarStandardCrossCityInfoCrossCityList) SetToCityCode(v string) *ApplyAddRequestTravelerStandardCarStandardCrossCityInfoCrossCityList {
	s.ToCityCode = &v
	return s
}

func (s *ApplyAddRequestTravelerStandardCarStandardCrossCityInfoCrossCityList) SetToCityName(v string) *ApplyAddRequestTravelerStandardCarStandardCrossCityInfoCrossCityList {
	s.ToCityName = &v
	return s
}

func (s *ApplyAddRequestTravelerStandardCarStandardCrossCityInfoCrossCityList) Validate() error {
	return dara.Validate(s)
}

type ApplyAddRequestTravelerStandardCarStandardElectronicFenceInfo struct {
	// This parameter is required.
	ElectronicFenceLocationsFrom []*ApplyAddRequestTravelerStandardCarStandardElectronicFenceInfoElectronicFenceLocationsFrom `json:"electronic_fence_locations_from,omitempty" xml:"electronic_fence_locations_from,omitempty" type:"Repeated"`
	// This parameter is required.
	ElectronicFenceLocationsTo []*ApplyAddRequestTravelerStandardCarStandardElectronicFenceInfoElectronicFenceLocationsTo `json:"electronic_fence_locations_to,omitempty" xml:"electronic_fence_locations_to,omitempty" type:"Repeated"`
	// This parameter is required.
	ElectronicFenceType *int32 `json:"electronic_fence_type,omitempty" xml:"electronic_fence_type,omitempty"`
}

func (s ApplyAddRequestTravelerStandardCarStandardElectronicFenceInfo) String() string {
	return dara.Prettify(s)
}

func (s ApplyAddRequestTravelerStandardCarStandardElectronicFenceInfo) GoString() string {
	return s.String()
}

func (s *ApplyAddRequestTravelerStandardCarStandardElectronicFenceInfo) GetElectronicFenceLocationsFrom() []*ApplyAddRequestTravelerStandardCarStandardElectronicFenceInfoElectronicFenceLocationsFrom {
	return s.ElectronicFenceLocationsFrom
}

func (s *ApplyAddRequestTravelerStandardCarStandardElectronicFenceInfo) GetElectronicFenceLocationsTo() []*ApplyAddRequestTravelerStandardCarStandardElectronicFenceInfoElectronicFenceLocationsTo {
	return s.ElectronicFenceLocationsTo
}

func (s *ApplyAddRequestTravelerStandardCarStandardElectronicFenceInfo) GetElectronicFenceType() *int32 {
	return s.ElectronicFenceType
}

func (s *ApplyAddRequestTravelerStandardCarStandardElectronicFenceInfo) SetElectronicFenceLocationsFrom(v []*ApplyAddRequestTravelerStandardCarStandardElectronicFenceInfoElectronicFenceLocationsFrom) *ApplyAddRequestTravelerStandardCarStandardElectronicFenceInfo {
	s.ElectronicFenceLocationsFrom = v
	return s
}

func (s *ApplyAddRequestTravelerStandardCarStandardElectronicFenceInfo) SetElectronicFenceLocationsTo(v []*ApplyAddRequestTravelerStandardCarStandardElectronicFenceInfoElectronicFenceLocationsTo) *ApplyAddRequestTravelerStandardCarStandardElectronicFenceInfo {
	s.ElectronicFenceLocationsTo = v
	return s
}

func (s *ApplyAddRequestTravelerStandardCarStandardElectronicFenceInfo) SetElectronicFenceType(v int32) *ApplyAddRequestTravelerStandardCarStandardElectronicFenceInfo {
	s.ElectronicFenceType = &v
	return s
}

func (s *ApplyAddRequestTravelerStandardCarStandardElectronicFenceInfo) Validate() error {
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

type ApplyAddRequestTravelerStandardCarStandardElectronicFenceInfoElectronicFenceLocationsFrom struct {
	// This parameter is required.
	Address *string `json:"address,omitempty" xml:"address,omitempty"`
	// This parameter is required.
	Latitude *string `json:"latitude,omitempty" xml:"latitude,omitempty"`
	// This parameter is required.
	Longitude *string `json:"longitude,omitempty" xml:"longitude,omitempty"`
	// This parameter is required.
	Radius *int32 `json:"radius,omitempty" xml:"radius,omitempty"`
}

func (s ApplyAddRequestTravelerStandardCarStandardElectronicFenceInfoElectronicFenceLocationsFrom) String() string {
	return dara.Prettify(s)
}

func (s ApplyAddRequestTravelerStandardCarStandardElectronicFenceInfoElectronicFenceLocationsFrom) GoString() string {
	return s.String()
}

func (s *ApplyAddRequestTravelerStandardCarStandardElectronicFenceInfoElectronicFenceLocationsFrom) GetAddress() *string {
	return s.Address
}

func (s *ApplyAddRequestTravelerStandardCarStandardElectronicFenceInfoElectronicFenceLocationsFrom) GetLatitude() *string {
	return s.Latitude
}

func (s *ApplyAddRequestTravelerStandardCarStandardElectronicFenceInfoElectronicFenceLocationsFrom) GetLongitude() *string {
	return s.Longitude
}

func (s *ApplyAddRequestTravelerStandardCarStandardElectronicFenceInfoElectronicFenceLocationsFrom) GetRadius() *int32 {
	return s.Radius
}

func (s *ApplyAddRequestTravelerStandardCarStandardElectronicFenceInfoElectronicFenceLocationsFrom) SetAddress(v string) *ApplyAddRequestTravelerStandardCarStandardElectronicFenceInfoElectronicFenceLocationsFrom {
	s.Address = &v
	return s
}

func (s *ApplyAddRequestTravelerStandardCarStandardElectronicFenceInfoElectronicFenceLocationsFrom) SetLatitude(v string) *ApplyAddRequestTravelerStandardCarStandardElectronicFenceInfoElectronicFenceLocationsFrom {
	s.Latitude = &v
	return s
}

func (s *ApplyAddRequestTravelerStandardCarStandardElectronicFenceInfoElectronicFenceLocationsFrom) SetLongitude(v string) *ApplyAddRequestTravelerStandardCarStandardElectronicFenceInfoElectronicFenceLocationsFrom {
	s.Longitude = &v
	return s
}

func (s *ApplyAddRequestTravelerStandardCarStandardElectronicFenceInfoElectronicFenceLocationsFrom) SetRadius(v int32) *ApplyAddRequestTravelerStandardCarStandardElectronicFenceInfoElectronicFenceLocationsFrom {
	s.Radius = &v
	return s
}

func (s *ApplyAddRequestTravelerStandardCarStandardElectronicFenceInfoElectronicFenceLocationsFrom) Validate() error {
	return dara.Validate(s)
}

type ApplyAddRequestTravelerStandardCarStandardElectronicFenceInfoElectronicFenceLocationsTo struct {
	// This parameter is required.
	Address *string `json:"address,omitempty" xml:"address,omitempty"`
	// This parameter is required.
	Latitude *string `json:"latitude,omitempty" xml:"latitude,omitempty"`
	// This parameter is required.
	Longitude *string `json:"longitude,omitempty" xml:"longitude,omitempty"`
	// This parameter is required.
	Radius *int32 `json:"radius,omitempty" xml:"radius,omitempty"`
}

func (s ApplyAddRequestTravelerStandardCarStandardElectronicFenceInfoElectronicFenceLocationsTo) String() string {
	return dara.Prettify(s)
}

func (s ApplyAddRequestTravelerStandardCarStandardElectronicFenceInfoElectronicFenceLocationsTo) GoString() string {
	return s.String()
}

func (s *ApplyAddRequestTravelerStandardCarStandardElectronicFenceInfoElectronicFenceLocationsTo) GetAddress() *string {
	return s.Address
}

func (s *ApplyAddRequestTravelerStandardCarStandardElectronicFenceInfoElectronicFenceLocationsTo) GetLatitude() *string {
	return s.Latitude
}

func (s *ApplyAddRequestTravelerStandardCarStandardElectronicFenceInfoElectronicFenceLocationsTo) GetLongitude() *string {
	return s.Longitude
}

func (s *ApplyAddRequestTravelerStandardCarStandardElectronicFenceInfoElectronicFenceLocationsTo) GetRadius() *int32 {
	return s.Radius
}

func (s *ApplyAddRequestTravelerStandardCarStandardElectronicFenceInfoElectronicFenceLocationsTo) SetAddress(v string) *ApplyAddRequestTravelerStandardCarStandardElectronicFenceInfoElectronicFenceLocationsTo {
	s.Address = &v
	return s
}

func (s *ApplyAddRequestTravelerStandardCarStandardElectronicFenceInfoElectronicFenceLocationsTo) SetLatitude(v string) *ApplyAddRequestTravelerStandardCarStandardElectronicFenceInfoElectronicFenceLocationsTo {
	s.Latitude = &v
	return s
}

func (s *ApplyAddRequestTravelerStandardCarStandardElectronicFenceInfoElectronicFenceLocationsTo) SetLongitude(v string) *ApplyAddRequestTravelerStandardCarStandardElectronicFenceInfoElectronicFenceLocationsTo {
	s.Longitude = &v
	return s
}

func (s *ApplyAddRequestTravelerStandardCarStandardElectronicFenceInfoElectronicFenceLocationsTo) SetRadius(v int32) *ApplyAddRequestTravelerStandardCarStandardElectronicFenceInfoElectronicFenceLocationsTo {
	s.Radius = &v
	return s
}

func (s *ApplyAddRequestTravelerStandardCarStandardElectronicFenceInfoElectronicFenceLocationsTo) Validate() error {
	return dara.Validate(s)
}

type ApplyAddRequestTravelerStandardCarStandardModifyDestinationInfo struct {
	// This parameter is required.
	ModifyDestination *bool `json:"modify_destination,omitempty" xml:"modify_destination,omitempty"`
}

func (s ApplyAddRequestTravelerStandardCarStandardModifyDestinationInfo) String() string {
	return dara.Prettify(s)
}

func (s ApplyAddRequestTravelerStandardCarStandardModifyDestinationInfo) GoString() string {
	return s.String()
}

func (s *ApplyAddRequestTravelerStandardCarStandardModifyDestinationInfo) GetModifyDestination() *bool {
	return s.ModifyDestination
}

func (s *ApplyAddRequestTravelerStandardCarStandardModifyDestinationInfo) SetModifyDestination(v bool) *ApplyAddRequestTravelerStandardCarStandardModifyDestinationInfo {
	s.ModifyDestination = &v
	return s
}

func (s *ApplyAddRequestTravelerStandardCarStandardModifyDestinationInfo) Validate() error {
	return dara.Validate(s)
}

type ApplyAddRequestTravelerStandardHotelCitys struct {
	// example:
	//
	// 0
	CityCode *string `json:"city_code,omitempty" xml:"city_code,omitempty"`
	// example:
	//
	// 北京
	CityName *string `json:"city_name,omitempty" xml:"city_name,omitempty"`
	// example:
	//
	// 1009
	Fee *int64 `json:"fee,omitempty" xml:"fee,omitempty"`
}

func (s ApplyAddRequestTravelerStandardHotelCitys) String() string {
	return dara.Prettify(s)
}

func (s ApplyAddRequestTravelerStandardHotelCitys) GoString() string {
	return s.String()
}

func (s *ApplyAddRequestTravelerStandardHotelCitys) GetCityCode() *string {
	return s.CityCode
}

func (s *ApplyAddRequestTravelerStandardHotelCitys) GetCityName() *string {
	return s.CityName
}

func (s *ApplyAddRequestTravelerStandardHotelCitys) GetFee() *int64 {
	return s.Fee
}

func (s *ApplyAddRequestTravelerStandardHotelCitys) SetCityCode(v string) *ApplyAddRequestTravelerStandardHotelCitys {
	s.CityCode = &v
	return s
}

func (s *ApplyAddRequestTravelerStandardHotelCitys) SetCityName(v string) *ApplyAddRequestTravelerStandardHotelCitys {
	s.CityName = &v
	return s
}

func (s *ApplyAddRequestTravelerStandardHotelCitys) SetFee(v int64) *ApplyAddRequestTravelerStandardHotelCitys {
	s.Fee = &v
	return s
}

func (s *ApplyAddRequestTravelerStandardHotelCitys) Validate() error {
	return dara.Validate(s)
}

type ApplyAddRequestTravelerStandardHotelIntlCitys struct {
	// example:
	//
	// 0
	CityCode *string `json:"city_code,omitempty" xml:"city_code,omitempty"`
	// example:
	//
	// 北京
	CityName *string `json:"city_name,omitempty" xml:"city_name,omitempty"`
	// example:
	//
	// 1009
	Fee *int64 `json:"fee,omitempty" xml:"fee,omitempty"`
}

func (s ApplyAddRequestTravelerStandardHotelIntlCitys) String() string {
	return dara.Prettify(s)
}

func (s ApplyAddRequestTravelerStandardHotelIntlCitys) GoString() string {
	return s.String()
}

func (s *ApplyAddRequestTravelerStandardHotelIntlCitys) GetCityCode() *string {
	return s.CityCode
}

func (s *ApplyAddRequestTravelerStandardHotelIntlCitys) GetCityName() *string {
	return s.CityName
}

func (s *ApplyAddRequestTravelerStandardHotelIntlCitys) GetFee() *int64 {
	return s.Fee
}

func (s *ApplyAddRequestTravelerStandardHotelIntlCitys) SetCityCode(v string) *ApplyAddRequestTravelerStandardHotelIntlCitys {
	s.CityCode = &v
	return s
}

func (s *ApplyAddRequestTravelerStandardHotelIntlCitys) SetCityName(v string) *ApplyAddRequestTravelerStandardHotelIntlCitys {
	s.CityName = &v
	return s
}

func (s *ApplyAddRequestTravelerStandardHotelIntlCitys) SetFee(v int64) *ApplyAddRequestTravelerStandardHotelIntlCitys {
	s.Fee = &v
	return s
}

func (s *ApplyAddRequestTravelerStandardHotelIntlCitys) Validate() error {
	return dara.Validate(s)
}
