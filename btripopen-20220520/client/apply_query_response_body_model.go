// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyQueryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ApplyQueryResponseBody
	GetCode() *string
	SetMessage(v string) *ApplyQueryResponseBody
	GetMessage() *string
	SetModule(v *ApplyQueryResponseBodyModule) *ApplyQueryResponseBody
	GetModule() *ApplyQueryResponseBodyModule
	SetRequestId(v string) *ApplyQueryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ApplyQueryResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *ApplyQueryResponseBody
	GetTraceId() *string
}

type ApplyQueryResponseBody struct {
	// example:
	//
	// SUCCESS
	Code    *string                       `json:"code,omitempty" xml:"code,omitempty"`
	Message *string                       `json:"message,omitempty" xml:"message,omitempty"`
	Module  *ApplyQueryResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	// example:
	//
	// B72B39C8-32DE-558D-AD1C-D53F11F6ADFE
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 21041ce316577904808056433edbb2
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s ApplyQueryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ApplyQueryResponseBody) GoString() string {
	return s.String()
}

func (s *ApplyQueryResponseBody) GetCode() *string {
	return s.Code
}

func (s *ApplyQueryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ApplyQueryResponseBody) GetModule() *ApplyQueryResponseBodyModule {
	return s.Module
}

func (s *ApplyQueryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ApplyQueryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ApplyQueryResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *ApplyQueryResponseBody) SetCode(v string) *ApplyQueryResponseBody {
	s.Code = &v
	return s
}

func (s *ApplyQueryResponseBody) SetMessage(v string) *ApplyQueryResponseBody {
	s.Message = &v
	return s
}

func (s *ApplyQueryResponseBody) SetModule(v *ApplyQueryResponseBodyModule) *ApplyQueryResponseBody {
	s.Module = v
	return s
}

func (s *ApplyQueryResponseBody) SetRequestId(v string) *ApplyQueryResponseBody {
	s.RequestId = &v
	return s
}

func (s *ApplyQueryResponseBody) SetSuccess(v bool) *ApplyQueryResponseBody {
	s.Success = &v
	return s
}

func (s *ApplyQueryResponseBody) SetTraceId(v string) *ApplyQueryResponseBody {
	s.TraceId = &v
	return s
}

func (s *ApplyQueryResponseBody) Validate() error {
	return dara.Validate(s)
}

type ApplyQueryResponseBodyModule struct {
	// example:
	//
	// 201710111505000464651
	ApplyShowId  *string                                     `json:"apply_show_id,omitempty" xml:"apply_show_id,omitempty"`
	ApproverList []*ApplyQueryResponseBodyModuleApproverList `json:"approver_list,omitempty" xml:"approver_list,omitempty" type:"Repeated"`
	// example:
	//
	// 100000
	Budget *int64 `json:"budget,omitempty" xml:"budget,omitempty"`
	// example:
	//
	// 1
	BudgetMerge *int32                               `json:"budget_merge,omitempty" xml:"budget_merge,omitempty"`
	CarRule     *ApplyQueryResponseBodyModuleCarRule `json:"car_rule,omitempty" xml:"car_rule,omitempty" type:"Struct"`
	// example:
	//
	// corpid
	CorpId   *string `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
	CorpName *string `json:"corp_name,omitempty" xml:"corp_name,omitempty"`
	// example:
	//
	// dept1
	DepartId *string `json:"depart_id,omitempty" xml:"depart_id,omitempty"`
	// example:
	//
	// adv
	DepartName *string `json:"depart_name,omitempty" xml:"depart_name,omitempty"`
	// 补充描述，账单中将会体现此字段的值。可以用于企业的统计和对账
	//
	// example:
	//
	// {"cost_center":"成本中心"}
	ExtendField          *string                                             `json:"extend_field,omitempty" xml:"extend_field,omitempty"`
	ExternalTravelerList []*ApplyQueryResponseBodyModuleExternalTravelerList `json:"external_traveler_list,omitempty" xml:"external_traveler_list,omitempty" type:"Repeated"`
	// example:
	//
	// 10000
	FlightBudget *int64 `json:"flight_budget,omitempty" xml:"flight_budget,omitempty"`
	// example:
	//
	// 2018-09-19T14:03Z
	GmtCreate *string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// example:
	//
	// 2018-09-19T14:03Z
	GmtModified *string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// example:
	//
	// 100000
	HotelBudget *int64                                  `json:"hotel_budget,omitempty" xml:"hotel_budget,omitempty"`
	HotelShare  *ApplyQueryResponseBodyModuleHotelShare `json:"hotel_share,omitempty" xml:"hotel_share,omitempty" type:"Struct"`
	// example:
	//
	// 3298
	Id               *int64                                       `json:"id,omitempty" xml:"id,omitempty"`
	IntlFlightBudget *int64                                       `json:"intl_flight_budget,omitempty" xml:"intl_flight_budget,omitempty"`
	IntlHotelBudget  *int64                                       `json:"intl_hotel_budget,omitempty" xml:"intl_hotel_budget,omitempty"`
	ItineraryList    []*ApplyQueryResponseBodyModuleItineraryList `json:"itinerary_list,omitempty" xml:"itinerary_list,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	ItineraryRule    *int32                                          `json:"itinerary_rule,omitempty" xml:"itinerary_rule,omitempty"`
	ItinerarySetList []*ApplyQueryResponseBodyModuleItinerarySetList `json:"itinerary_set_list,omitempty" xml:"itinerary_set_list,omitempty" type:"Repeated"`
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
	Status     *int32  `json:"status,omitempty" xml:"status,omitempty"`
	StatusDesc *string `json:"status_desc,omitempty" xml:"status_desc,omitempty"`
	// example:
	//
	// 1213
	ThirdpartBusinessId *string `json:"thirdpart_business_id,omitempty" xml:"thirdpart_business_id,omitempty"`
	// example:
	//
	// 1214254
	ThirdpartId *string `json:"thirdpart_id,omitempty" xml:"thirdpart_id,omitempty"`
	// example:
	//
	// 1
	TogetherBookRule *int32 `json:"together_book_rule,omitempty" xml:"together_book_rule,omitempty"`
	// example:
	//
	// 10000
	TrainBudget  *int64                                      `json:"train_budget,omitempty" xml:"train_budget,omitempty"`
	TravelerList []*ApplyQueryResponseBodyModuleTravelerList `json:"traveler_list,omitempty" xml:"traveler_list,omitempty" type:"Repeated"`
	TripCause    *string                                     `json:"trip_cause,omitempty" xml:"trip_cause,omitempty"`
	// example:
	//
	// 2
	TripDay   *int32  `json:"trip_day,omitempty" xml:"trip_day,omitempty"`
	TripTitle *string `json:"trip_title,omitempty" xml:"trip_title,omitempty"`
	// example:
	//
	// 2
	Type *int32 `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// abd123
	UnionNo *string `json:"union_no,omitempty" xml:"union_no,omitempty"`
	// example:
	//
	// user1
	UserId   *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	UserName *string `json:"user_name,omitempty" xml:"user_name,omitempty"`
	// example:
	//
	// 10000
	VehicleBudget *int64 `json:"vehicle_budget,omitempty" xml:"vehicle_budget,omitempty"`
}

func (s ApplyQueryResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s ApplyQueryResponseBodyModule) GoString() string {
	return s.String()
}

func (s *ApplyQueryResponseBodyModule) GetApplyShowId() *string {
	return s.ApplyShowId
}

func (s *ApplyQueryResponseBodyModule) GetApproverList() []*ApplyQueryResponseBodyModuleApproverList {
	return s.ApproverList
}

func (s *ApplyQueryResponseBodyModule) GetBudget() *int64 {
	return s.Budget
}

func (s *ApplyQueryResponseBodyModule) GetBudgetMerge() *int32 {
	return s.BudgetMerge
}

func (s *ApplyQueryResponseBodyModule) GetCarRule() *ApplyQueryResponseBodyModuleCarRule {
	return s.CarRule
}

func (s *ApplyQueryResponseBodyModule) GetCorpId() *string {
	return s.CorpId
}

func (s *ApplyQueryResponseBodyModule) GetCorpName() *string {
	return s.CorpName
}

func (s *ApplyQueryResponseBodyModule) GetDepartId() *string {
	return s.DepartId
}

func (s *ApplyQueryResponseBodyModule) GetDepartName() *string {
	return s.DepartName
}

func (s *ApplyQueryResponseBodyModule) GetExtendField() *string {
	return s.ExtendField
}

func (s *ApplyQueryResponseBodyModule) GetExternalTravelerList() []*ApplyQueryResponseBodyModuleExternalTravelerList {
	return s.ExternalTravelerList
}

func (s *ApplyQueryResponseBodyModule) GetFlightBudget() *int64 {
	return s.FlightBudget
}

func (s *ApplyQueryResponseBodyModule) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ApplyQueryResponseBodyModule) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ApplyQueryResponseBodyModule) GetHotelBudget() *int64 {
	return s.HotelBudget
}

func (s *ApplyQueryResponseBodyModule) GetHotelShare() *ApplyQueryResponseBodyModuleHotelShare {
	return s.HotelShare
}

func (s *ApplyQueryResponseBodyModule) GetId() *int64 {
	return s.Id
}

func (s *ApplyQueryResponseBodyModule) GetIntlFlightBudget() *int64 {
	return s.IntlFlightBudget
}

func (s *ApplyQueryResponseBodyModule) GetIntlHotelBudget() *int64 {
	return s.IntlHotelBudget
}

func (s *ApplyQueryResponseBodyModule) GetItineraryList() []*ApplyQueryResponseBodyModuleItineraryList {
	return s.ItineraryList
}

func (s *ApplyQueryResponseBodyModule) GetItineraryRule() *int32 {
	return s.ItineraryRule
}

func (s *ApplyQueryResponseBodyModule) GetItinerarySetList() []*ApplyQueryResponseBodyModuleItinerarySetList {
	return s.ItinerarySetList
}

func (s *ApplyQueryResponseBodyModule) GetLimitTraveler() *int32 {
	return s.LimitTraveler
}

func (s *ApplyQueryResponseBodyModule) GetMealBudget() *int64 {
	return s.MealBudget
}

func (s *ApplyQueryResponseBodyModule) GetPaymentDepartmentId() *string {
	return s.PaymentDepartmentId
}

func (s *ApplyQueryResponseBodyModule) GetPaymentDepartmentName() *string {
	return s.PaymentDepartmentName
}

func (s *ApplyQueryResponseBodyModule) GetStatus() *int32 {
	return s.Status
}

func (s *ApplyQueryResponseBodyModule) GetStatusDesc() *string {
	return s.StatusDesc
}

func (s *ApplyQueryResponseBodyModule) GetThirdpartBusinessId() *string {
	return s.ThirdpartBusinessId
}

func (s *ApplyQueryResponseBodyModule) GetThirdpartId() *string {
	return s.ThirdpartId
}

func (s *ApplyQueryResponseBodyModule) GetTogetherBookRule() *int32 {
	return s.TogetherBookRule
}

func (s *ApplyQueryResponseBodyModule) GetTrainBudget() *int64 {
	return s.TrainBudget
}

func (s *ApplyQueryResponseBodyModule) GetTravelerList() []*ApplyQueryResponseBodyModuleTravelerList {
	return s.TravelerList
}

func (s *ApplyQueryResponseBodyModule) GetTripCause() *string {
	return s.TripCause
}

func (s *ApplyQueryResponseBodyModule) GetTripDay() *int32 {
	return s.TripDay
}

func (s *ApplyQueryResponseBodyModule) GetTripTitle() *string {
	return s.TripTitle
}

func (s *ApplyQueryResponseBodyModule) GetType() *int32 {
	return s.Type
}

func (s *ApplyQueryResponseBodyModule) GetUnionNo() *string {
	return s.UnionNo
}

func (s *ApplyQueryResponseBodyModule) GetUserId() *string {
	return s.UserId
}

func (s *ApplyQueryResponseBodyModule) GetUserName() *string {
	return s.UserName
}

func (s *ApplyQueryResponseBodyModule) GetVehicleBudget() *int64 {
	return s.VehicleBudget
}

func (s *ApplyQueryResponseBodyModule) SetApplyShowId(v string) *ApplyQueryResponseBodyModule {
	s.ApplyShowId = &v
	return s
}

func (s *ApplyQueryResponseBodyModule) SetApproverList(v []*ApplyQueryResponseBodyModuleApproverList) *ApplyQueryResponseBodyModule {
	s.ApproverList = v
	return s
}

func (s *ApplyQueryResponseBodyModule) SetBudget(v int64) *ApplyQueryResponseBodyModule {
	s.Budget = &v
	return s
}

func (s *ApplyQueryResponseBodyModule) SetBudgetMerge(v int32) *ApplyQueryResponseBodyModule {
	s.BudgetMerge = &v
	return s
}

func (s *ApplyQueryResponseBodyModule) SetCarRule(v *ApplyQueryResponseBodyModuleCarRule) *ApplyQueryResponseBodyModule {
	s.CarRule = v
	return s
}

func (s *ApplyQueryResponseBodyModule) SetCorpId(v string) *ApplyQueryResponseBodyModule {
	s.CorpId = &v
	return s
}

func (s *ApplyQueryResponseBodyModule) SetCorpName(v string) *ApplyQueryResponseBodyModule {
	s.CorpName = &v
	return s
}

func (s *ApplyQueryResponseBodyModule) SetDepartId(v string) *ApplyQueryResponseBodyModule {
	s.DepartId = &v
	return s
}

func (s *ApplyQueryResponseBodyModule) SetDepartName(v string) *ApplyQueryResponseBodyModule {
	s.DepartName = &v
	return s
}

func (s *ApplyQueryResponseBodyModule) SetExtendField(v string) *ApplyQueryResponseBodyModule {
	s.ExtendField = &v
	return s
}

func (s *ApplyQueryResponseBodyModule) SetExternalTravelerList(v []*ApplyQueryResponseBodyModuleExternalTravelerList) *ApplyQueryResponseBodyModule {
	s.ExternalTravelerList = v
	return s
}

func (s *ApplyQueryResponseBodyModule) SetFlightBudget(v int64) *ApplyQueryResponseBodyModule {
	s.FlightBudget = &v
	return s
}

func (s *ApplyQueryResponseBodyModule) SetGmtCreate(v string) *ApplyQueryResponseBodyModule {
	s.GmtCreate = &v
	return s
}

func (s *ApplyQueryResponseBodyModule) SetGmtModified(v string) *ApplyQueryResponseBodyModule {
	s.GmtModified = &v
	return s
}

func (s *ApplyQueryResponseBodyModule) SetHotelBudget(v int64) *ApplyQueryResponseBodyModule {
	s.HotelBudget = &v
	return s
}

func (s *ApplyQueryResponseBodyModule) SetHotelShare(v *ApplyQueryResponseBodyModuleHotelShare) *ApplyQueryResponseBodyModule {
	s.HotelShare = v
	return s
}

func (s *ApplyQueryResponseBodyModule) SetId(v int64) *ApplyQueryResponseBodyModule {
	s.Id = &v
	return s
}

func (s *ApplyQueryResponseBodyModule) SetIntlFlightBudget(v int64) *ApplyQueryResponseBodyModule {
	s.IntlFlightBudget = &v
	return s
}

func (s *ApplyQueryResponseBodyModule) SetIntlHotelBudget(v int64) *ApplyQueryResponseBodyModule {
	s.IntlHotelBudget = &v
	return s
}

func (s *ApplyQueryResponseBodyModule) SetItineraryList(v []*ApplyQueryResponseBodyModuleItineraryList) *ApplyQueryResponseBodyModule {
	s.ItineraryList = v
	return s
}

func (s *ApplyQueryResponseBodyModule) SetItineraryRule(v int32) *ApplyQueryResponseBodyModule {
	s.ItineraryRule = &v
	return s
}

func (s *ApplyQueryResponseBodyModule) SetItinerarySetList(v []*ApplyQueryResponseBodyModuleItinerarySetList) *ApplyQueryResponseBodyModule {
	s.ItinerarySetList = v
	return s
}

func (s *ApplyQueryResponseBodyModule) SetLimitTraveler(v int32) *ApplyQueryResponseBodyModule {
	s.LimitTraveler = &v
	return s
}

func (s *ApplyQueryResponseBodyModule) SetMealBudget(v int64) *ApplyQueryResponseBodyModule {
	s.MealBudget = &v
	return s
}

func (s *ApplyQueryResponseBodyModule) SetPaymentDepartmentId(v string) *ApplyQueryResponseBodyModule {
	s.PaymentDepartmentId = &v
	return s
}

func (s *ApplyQueryResponseBodyModule) SetPaymentDepartmentName(v string) *ApplyQueryResponseBodyModule {
	s.PaymentDepartmentName = &v
	return s
}

func (s *ApplyQueryResponseBodyModule) SetStatus(v int32) *ApplyQueryResponseBodyModule {
	s.Status = &v
	return s
}

func (s *ApplyQueryResponseBodyModule) SetStatusDesc(v string) *ApplyQueryResponseBodyModule {
	s.StatusDesc = &v
	return s
}

func (s *ApplyQueryResponseBodyModule) SetThirdpartBusinessId(v string) *ApplyQueryResponseBodyModule {
	s.ThirdpartBusinessId = &v
	return s
}

func (s *ApplyQueryResponseBodyModule) SetThirdpartId(v string) *ApplyQueryResponseBodyModule {
	s.ThirdpartId = &v
	return s
}

func (s *ApplyQueryResponseBodyModule) SetTogetherBookRule(v int32) *ApplyQueryResponseBodyModule {
	s.TogetherBookRule = &v
	return s
}

func (s *ApplyQueryResponseBodyModule) SetTrainBudget(v int64) *ApplyQueryResponseBodyModule {
	s.TrainBudget = &v
	return s
}

func (s *ApplyQueryResponseBodyModule) SetTravelerList(v []*ApplyQueryResponseBodyModuleTravelerList) *ApplyQueryResponseBodyModule {
	s.TravelerList = v
	return s
}

func (s *ApplyQueryResponseBodyModule) SetTripCause(v string) *ApplyQueryResponseBodyModule {
	s.TripCause = &v
	return s
}

func (s *ApplyQueryResponseBodyModule) SetTripDay(v int32) *ApplyQueryResponseBodyModule {
	s.TripDay = &v
	return s
}

func (s *ApplyQueryResponseBodyModule) SetTripTitle(v string) *ApplyQueryResponseBodyModule {
	s.TripTitle = &v
	return s
}

func (s *ApplyQueryResponseBodyModule) SetType(v int32) *ApplyQueryResponseBodyModule {
	s.Type = &v
	return s
}

func (s *ApplyQueryResponseBodyModule) SetUnionNo(v string) *ApplyQueryResponseBodyModule {
	s.UnionNo = &v
	return s
}

func (s *ApplyQueryResponseBodyModule) SetUserId(v string) *ApplyQueryResponseBodyModule {
	s.UserId = &v
	return s
}

func (s *ApplyQueryResponseBodyModule) SetUserName(v string) *ApplyQueryResponseBodyModule {
	s.UserName = &v
	return s
}

func (s *ApplyQueryResponseBodyModule) SetVehicleBudget(v int64) *ApplyQueryResponseBodyModule {
	s.VehicleBudget = &v
	return s
}

func (s *ApplyQueryResponseBodyModule) Validate() error {
	return dara.Validate(s)
}

type ApplyQueryResponseBodyModuleApproverList struct {
	Note *string `json:"note,omitempty" xml:"note,omitempty"`
	// example:
	//
	// 2018-09-19T14:03Z
	OperateTime *string `json:"operate_time,omitempty" xml:"operate_time,omitempty"`
	// example:
	//
	// 1
	Order *int32 `json:"order,omitempty" xml:"order,omitempty"`
	// example:
	//
	// 0
	Status     *int32  `json:"status,omitempty" xml:"status,omitempty"`
	StatusDesc *string `json:"status_desc,omitempty" xml:"status_desc,omitempty"`
	// example:
	//
	// user1
	UserId   *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	UserName *string `json:"user_name,omitempty" xml:"user_name,omitempty"`
}

func (s ApplyQueryResponseBodyModuleApproverList) String() string {
	return dara.Prettify(s)
}

func (s ApplyQueryResponseBodyModuleApproverList) GoString() string {
	return s.String()
}

func (s *ApplyQueryResponseBodyModuleApproverList) GetNote() *string {
	return s.Note
}

func (s *ApplyQueryResponseBodyModuleApproverList) GetOperateTime() *string {
	return s.OperateTime
}

func (s *ApplyQueryResponseBodyModuleApproverList) GetOrder() *int32 {
	return s.Order
}

func (s *ApplyQueryResponseBodyModuleApproverList) GetStatus() *int32 {
	return s.Status
}

func (s *ApplyQueryResponseBodyModuleApproverList) GetStatusDesc() *string {
	return s.StatusDesc
}

func (s *ApplyQueryResponseBodyModuleApproverList) GetUserId() *string {
	return s.UserId
}

func (s *ApplyQueryResponseBodyModuleApproverList) GetUserName() *string {
	return s.UserName
}

func (s *ApplyQueryResponseBodyModuleApproverList) SetNote(v string) *ApplyQueryResponseBodyModuleApproverList {
	s.Note = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleApproverList) SetOperateTime(v string) *ApplyQueryResponseBodyModuleApproverList {
	s.OperateTime = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleApproverList) SetOrder(v int32) *ApplyQueryResponseBodyModuleApproverList {
	s.Order = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleApproverList) SetStatus(v int32) *ApplyQueryResponseBodyModuleApproverList {
	s.Status = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleApproverList) SetStatusDesc(v string) *ApplyQueryResponseBodyModuleApproverList {
	s.StatusDesc = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleApproverList) SetUserId(v string) *ApplyQueryResponseBodyModuleApproverList {
	s.UserId = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleApproverList) SetUserName(v string) *ApplyQueryResponseBodyModuleApproverList {
	s.UserName = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleApproverList) Validate() error {
	return dara.Validate(s)
}

type ApplyQueryResponseBodyModuleCarRule struct {
	ScenarioTemplateId   *string `json:"scenario_template_id,omitempty" xml:"scenario_template_id,omitempty"`
	ScenarioTemplateName *string `json:"scenario_template_name,omitempty" xml:"scenario_template_name,omitempty"`
}

func (s ApplyQueryResponseBodyModuleCarRule) String() string {
	return dara.Prettify(s)
}

func (s ApplyQueryResponseBodyModuleCarRule) GoString() string {
	return s.String()
}

func (s *ApplyQueryResponseBodyModuleCarRule) GetScenarioTemplateId() *string {
	return s.ScenarioTemplateId
}

func (s *ApplyQueryResponseBodyModuleCarRule) GetScenarioTemplateName() *string {
	return s.ScenarioTemplateName
}

func (s *ApplyQueryResponseBodyModuleCarRule) SetScenarioTemplateId(v string) *ApplyQueryResponseBodyModuleCarRule {
	s.ScenarioTemplateId = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleCarRule) SetScenarioTemplateName(v string) *ApplyQueryResponseBodyModuleCarRule {
	s.ScenarioTemplateName = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleCarRule) Validate() error {
	return dara.Validate(s)
}

type ApplyQueryResponseBodyModuleExternalTravelerList struct {
	Attribute *string `json:"attribute,omitempty" xml:"attribute,omitempty"`
	// example:
	//
	// 1
	BusinessDiscount *int32  `json:"business_discount,omitempty" xml:"business_discount,omitempty"`
	CostCenterName   *string `json:"cost_center_name,omitempty" xml:"cost_center_name,omitempty"`
	DepartId         *string `json:"depart_id,omitempty" xml:"depart_id,omitempty"`
	// example:
	//
	// 1
	EconomyDiscount *int32  `json:"economy_discount,omitempty" xml:"economy_discount,omitempty"`
	ExternalUserId  *string `json:"external_user_id,omitempty" xml:"external_user_id,omitempty"`
	// example:
	//
	// 1
	FirstDiscount *int32 `json:"first_discount,omitempty" xml:"first_discount,omitempty"`
	// example:
	//
	// F
	FlightCabins          *string                                                           `json:"flight_cabins,omitempty" xml:"flight_cabins,omitempty"`
	FlightIntlRuleCode    *int64                                                            `json:"flight_intl_rule_code,omitempty" xml:"flight_intl_rule_code,omitempty"`
	FlightRuleCode        *int64                                                            `json:"flight_rule_code,omitempty" xml:"flight_rule_code,omitempty"`
	HotelCitys            []*ApplyQueryResponseBodyModuleExternalTravelerListHotelCitys     `json:"hotel_citys,omitempty" xml:"hotel_citys,omitempty" type:"Repeated"`
	HotelIntlCitys        []*ApplyQueryResponseBodyModuleExternalTravelerListHotelIntlCitys `json:"hotel_intl_citys,omitempty" xml:"hotel_intl_citys,omitempty" type:"Repeated"`
	HotelIntlRuleCode     *int64                                                            `json:"hotel_intl_rule_code,omitempty" xml:"hotel_intl_rule_code,omitempty"`
	HotelRuleCode         *int64                                                            `json:"hotel_rule_code,omitempty" xml:"hotel_rule_code,omitempty"`
	InvoiceName           *string                                                           `json:"invoice_name,omitempty" xml:"invoice_name,omitempty"`
	PaymentDepartmentId   *string                                                           `json:"payment_department_id,omitempty" xml:"payment_department_id,omitempty"`
	PaymentDepartmentName *string                                                           `json:"payment_department_name,omitempty" xml:"payment_department_name,omitempty"`
	// example:
	//
	// 1
	PremiumEconomyDiscount *int32  `json:"premium_economy_discount,omitempty" xml:"premium_economy_discount,omitempty"`
	ProjectCode            *string `json:"project_code,omitempty" xml:"project_code,omitempty"`
	ProjectTitle           *string `json:"project_title,omitempty" xml:"project_title,omitempty"`
	// example:
	//
	// 1
	ReserveType           *int32  `json:"reserve_type,omitempty" xml:"reserve_type,omitempty"`
	ThirdPartInvoiceId    *string `json:"third_part_invoice_id,omitempty" xml:"third_part_invoice_id,omitempty"`
	ThirdpartCostCenterId *string `json:"thirdpart_cost_center_id,omitempty" xml:"thirdpart_cost_center_id,omitempty"`
	ThirdpartDepartId     *string `json:"thirdpart_depart_id,omitempty" xml:"thirdpart_depart_id,omitempty"`
	TrainRuleCode         *int64  `json:"train_rule_code,omitempty" xml:"train_rule_code,omitempty"`
	// example:
	//
	// 1
	TrainSeats *string `json:"train_seats,omitempty" xml:"train_seats,omitempty"`
	UserName   *string `json:"user_name,omitempty" xml:"user_name,omitempty"`
}

func (s ApplyQueryResponseBodyModuleExternalTravelerList) String() string {
	return dara.Prettify(s)
}

func (s ApplyQueryResponseBodyModuleExternalTravelerList) GoString() string {
	return s.String()
}

func (s *ApplyQueryResponseBodyModuleExternalTravelerList) GetAttribute() *string {
	return s.Attribute
}

func (s *ApplyQueryResponseBodyModuleExternalTravelerList) GetBusinessDiscount() *int32 {
	return s.BusinessDiscount
}

func (s *ApplyQueryResponseBodyModuleExternalTravelerList) GetCostCenterName() *string {
	return s.CostCenterName
}

func (s *ApplyQueryResponseBodyModuleExternalTravelerList) GetDepartId() *string {
	return s.DepartId
}

func (s *ApplyQueryResponseBodyModuleExternalTravelerList) GetEconomyDiscount() *int32 {
	return s.EconomyDiscount
}

func (s *ApplyQueryResponseBodyModuleExternalTravelerList) GetExternalUserId() *string {
	return s.ExternalUserId
}

func (s *ApplyQueryResponseBodyModuleExternalTravelerList) GetFirstDiscount() *int32 {
	return s.FirstDiscount
}

func (s *ApplyQueryResponseBodyModuleExternalTravelerList) GetFlightCabins() *string {
	return s.FlightCabins
}

func (s *ApplyQueryResponseBodyModuleExternalTravelerList) GetFlightIntlRuleCode() *int64 {
	return s.FlightIntlRuleCode
}

func (s *ApplyQueryResponseBodyModuleExternalTravelerList) GetFlightRuleCode() *int64 {
	return s.FlightRuleCode
}

func (s *ApplyQueryResponseBodyModuleExternalTravelerList) GetHotelCitys() []*ApplyQueryResponseBodyModuleExternalTravelerListHotelCitys {
	return s.HotelCitys
}

func (s *ApplyQueryResponseBodyModuleExternalTravelerList) GetHotelIntlCitys() []*ApplyQueryResponseBodyModuleExternalTravelerListHotelIntlCitys {
	return s.HotelIntlCitys
}

func (s *ApplyQueryResponseBodyModuleExternalTravelerList) GetHotelIntlRuleCode() *int64 {
	return s.HotelIntlRuleCode
}

func (s *ApplyQueryResponseBodyModuleExternalTravelerList) GetHotelRuleCode() *int64 {
	return s.HotelRuleCode
}

func (s *ApplyQueryResponseBodyModuleExternalTravelerList) GetInvoiceName() *string {
	return s.InvoiceName
}

func (s *ApplyQueryResponseBodyModuleExternalTravelerList) GetPaymentDepartmentId() *string {
	return s.PaymentDepartmentId
}

func (s *ApplyQueryResponseBodyModuleExternalTravelerList) GetPaymentDepartmentName() *string {
	return s.PaymentDepartmentName
}

func (s *ApplyQueryResponseBodyModuleExternalTravelerList) GetPremiumEconomyDiscount() *int32 {
	return s.PremiumEconomyDiscount
}

func (s *ApplyQueryResponseBodyModuleExternalTravelerList) GetProjectCode() *string {
	return s.ProjectCode
}

func (s *ApplyQueryResponseBodyModuleExternalTravelerList) GetProjectTitle() *string {
	return s.ProjectTitle
}

func (s *ApplyQueryResponseBodyModuleExternalTravelerList) GetReserveType() *int32 {
	return s.ReserveType
}

func (s *ApplyQueryResponseBodyModuleExternalTravelerList) GetThirdPartInvoiceId() *string {
	return s.ThirdPartInvoiceId
}

func (s *ApplyQueryResponseBodyModuleExternalTravelerList) GetThirdpartCostCenterId() *string {
	return s.ThirdpartCostCenterId
}

func (s *ApplyQueryResponseBodyModuleExternalTravelerList) GetThirdpartDepartId() *string {
	return s.ThirdpartDepartId
}

func (s *ApplyQueryResponseBodyModuleExternalTravelerList) GetTrainRuleCode() *int64 {
	return s.TrainRuleCode
}

func (s *ApplyQueryResponseBodyModuleExternalTravelerList) GetTrainSeats() *string {
	return s.TrainSeats
}

func (s *ApplyQueryResponseBodyModuleExternalTravelerList) GetUserName() *string {
	return s.UserName
}

func (s *ApplyQueryResponseBodyModuleExternalTravelerList) SetAttribute(v string) *ApplyQueryResponseBodyModuleExternalTravelerList {
	s.Attribute = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleExternalTravelerList) SetBusinessDiscount(v int32) *ApplyQueryResponseBodyModuleExternalTravelerList {
	s.BusinessDiscount = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleExternalTravelerList) SetCostCenterName(v string) *ApplyQueryResponseBodyModuleExternalTravelerList {
	s.CostCenterName = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleExternalTravelerList) SetDepartId(v string) *ApplyQueryResponseBodyModuleExternalTravelerList {
	s.DepartId = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleExternalTravelerList) SetEconomyDiscount(v int32) *ApplyQueryResponseBodyModuleExternalTravelerList {
	s.EconomyDiscount = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleExternalTravelerList) SetExternalUserId(v string) *ApplyQueryResponseBodyModuleExternalTravelerList {
	s.ExternalUserId = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleExternalTravelerList) SetFirstDiscount(v int32) *ApplyQueryResponseBodyModuleExternalTravelerList {
	s.FirstDiscount = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleExternalTravelerList) SetFlightCabins(v string) *ApplyQueryResponseBodyModuleExternalTravelerList {
	s.FlightCabins = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleExternalTravelerList) SetFlightIntlRuleCode(v int64) *ApplyQueryResponseBodyModuleExternalTravelerList {
	s.FlightIntlRuleCode = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleExternalTravelerList) SetFlightRuleCode(v int64) *ApplyQueryResponseBodyModuleExternalTravelerList {
	s.FlightRuleCode = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleExternalTravelerList) SetHotelCitys(v []*ApplyQueryResponseBodyModuleExternalTravelerListHotelCitys) *ApplyQueryResponseBodyModuleExternalTravelerList {
	s.HotelCitys = v
	return s
}

func (s *ApplyQueryResponseBodyModuleExternalTravelerList) SetHotelIntlCitys(v []*ApplyQueryResponseBodyModuleExternalTravelerListHotelIntlCitys) *ApplyQueryResponseBodyModuleExternalTravelerList {
	s.HotelIntlCitys = v
	return s
}

func (s *ApplyQueryResponseBodyModuleExternalTravelerList) SetHotelIntlRuleCode(v int64) *ApplyQueryResponseBodyModuleExternalTravelerList {
	s.HotelIntlRuleCode = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleExternalTravelerList) SetHotelRuleCode(v int64) *ApplyQueryResponseBodyModuleExternalTravelerList {
	s.HotelRuleCode = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleExternalTravelerList) SetInvoiceName(v string) *ApplyQueryResponseBodyModuleExternalTravelerList {
	s.InvoiceName = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleExternalTravelerList) SetPaymentDepartmentId(v string) *ApplyQueryResponseBodyModuleExternalTravelerList {
	s.PaymentDepartmentId = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleExternalTravelerList) SetPaymentDepartmentName(v string) *ApplyQueryResponseBodyModuleExternalTravelerList {
	s.PaymentDepartmentName = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleExternalTravelerList) SetPremiumEconomyDiscount(v int32) *ApplyQueryResponseBodyModuleExternalTravelerList {
	s.PremiumEconomyDiscount = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleExternalTravelerList) SetProjectCode(v string) *ApplyQueryResponseBodyModuleExternalTravelerList {
	s.ProjectCode = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleExternalTravelerList) SetProjectTitle(v string) *ApplyQueryResponseBodyModuleExternalTravelerList {
	s.ProjectTitle = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleExternalTravelerList) SetReserveType(v int32) *ApplyQueryResponseBodyModuleExternalTravelerList {
	s.ReserveType = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleExternalTravelerList) SetThirdPartInvoiceId(v string) *ApplyQueryResponseBodyModuleExternalTravelerList {
	s.ThirdPartInvoiceId = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleExternalTravelerList) SetThirdpartCostCenterId(v string) *ApplyQueryResponseBodyModuleExternalTravelerList {
	s.ThirdpartCostCenterId = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleExternalTravelerList) SetThirdpartDepartId(v string) *ApplyQueryResponseBodyModuleExternalTravelerList {
	s.ThirdpartDepartId = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleExternalTravelerList) SetTrainRuleCode(v int64) *ApplyQueryResponseBodyModuleExternalTravelerList {
	s.TrainRuleCode = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleExternalTravelerList) SetTrainSeats(v string) *ApplyQueryResponseBodyModuleExternalTravelerList {
	s.TrainSeats = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleExternalTravelerList) SetUserName(v string) *ApplyQueryResponseBodyModuleExternalTravelerList {
	s.UserName = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleExternalTravelerList) Validate() error {
	return dara.Validate(s)
}

type ApplyQueryResponseBodyModuleExternalTravelerListHotelCitys struct {
	// example:
	//
	// 0
	CityCode *string `json:"city_code,omitempty" xml:"city_code,omitempty"`
	CityName *string `json:"city_name,omitempty" xml:"city_name,omitempty"`
	// example:
	//
	// 100000
	Fee *int64 `json:"fee,omitempty" xml:"fee,omitempty"`
}

func (s ApplyQueryResponseBodyModuleExternalTravelerListHotelCitys) String() string {
	return dara.Prettify(s)
}

func (s ApplyQueryResponseBodyModuleExternalTravelerListHotelCitys) GoString() string {
	return s.String()
}

func (s *ApplyQueryResponseBodyModuleExternalTravelerListHotelCitys) GetCityCode() *string {
	return s.CityCode
}

func (s *ApplyQueryResponseBodyModuleExternalTravelerListHotelCitys) GetCityName() *string {
	return s.CityName
}

func (s *ApplyQueryResponseBodyModuleExternalTravelerListHotelCitys) GetFee() *int64 {
	return s.Fee
}

func (s *ApplyQueryResponseBodyModuleExternalTravelerListHotelCitys) SetCityCode(v string) *ApplyQueryResponseBodyModuleExternalTravelerListHotelCitys {
	s.CityCode = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleExternalTravelerListHotelCitys) SetCityName(v string) *ApplyQueryResponseBodyModuleExternalTravelerListHotelCitys {
	s.CityName = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleExternalTravelerListHotelCitys) SetFee(v int64) *ApplyQueryResponseBodyModuleExternalTravelerListHotelCitys {
	s.Fee = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleExternalTravelerListHotelCitys) Validate() error {
	return dara.Validate(s)
}

type ApplyQueryResponseBodyModuleExternalTravelerListHotelIntlCitys struct {
	CityCode *string `json:"city_code,omitempty" xml:"city_code,omitempty"`
	CityName *string `json:"city_name,omitempty" xml:"city_name,omitempty"`
	Fee      *int64  `json:"fee,omitempty" xml:"fee,omitempty"`
}

func (s ApplyQueryResponseBodyModuleExternalTravelerListHotelIntlCitys) String() string {
	return dara.Prettify(s)
}

func (s ApplyQueryResponseBodyModuleExternalTravelerListHotelIntlCitys) GoString() string {
	return s.String()
}

func (s *ApplyQueryResponseBodyModuleExternalTravelerListHotelIntlCitys) GetCityCode() *string {
	return s.CityCode
}

func (s *ApplyQueryResponseBodyModuleExternalTravelerListHotelIntlCitys) GetCityName() *string {
	return s.CityName
}

func (s *ApplyQueryResponseBodyModuleExternalTravelerListHotelIntlCitys) GetFee() *int64 {
	return s.Fee
}

func (s *ApplyQueryResponseBodyModuleExternalTravelerListHotelIntlCitys) SetCityCode(v string) *ApplyQueryResponseBodyModuleExternalTravelerListHotelIntlCitys {
	s.CityCode = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleExternalTravelerListHotelIntlCitys) SetCityName(v string) *ApplyQueryResponseBodyModuleExternalTravelerListHotelIntlCitys {
	s.CityName = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleExternalTravelerListHotelIntlCitys) SetFee(v int64) *ApplyQueryResponseBodyModuleExternalTravelerListHotelIntlCitys {
	s.Fee = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleExternalTravelerListHotelIntlCitys) Validate() error {
	return dara.Validate(s)
}

type ApplyQueryResponseBodyModuleHotelShare struct {
	// example:
	//
	// 70
	Param *string `json:"param,omitempty" xml:"param,omitempty"`
	// example:
	//
	// 1
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ApplyQueryResponseBodyModuleHotelShare) String() string {
	return dara.Prettify(s)
}

func (s ApplyQueryResponseBodyModuleHotelShare) GoString() string {
	return s.String()
}

func (s *ApplyQueryResponseBodyModuleHotelShare) GetParam() *string {
	return s.Param
}

func (s *ApplyQueryResponseBodyModuleHotelShare) GetType() *string {
	return s.Type
}

func (s *ApplyQueryResponseBodyModuleHotelShare) SetParam(v string) *ApplyQueryResponseBodyModuleHotelShare {
	s.Param = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleHotelShare) SetType(v string) *ApplyQueryResponseBodyModuleHotelShare {
	s.Type = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleHotelShare) Validate() error {
	return dara.Validate(s)
}

type ApplyQueryResponseBodyModuleItineraryList struct {
	ArrCity *string `json:"arr_city,omitempty" xml:"arr_city,omitempty"`
	// example:
	//
	// 330100
	ArrCityCode *string `json:"arr_city_code,omitempty" xml:"arr_city_code,omitempty"`
	// example:
	//
	// 2018-09-19T14:03Z
	ArrDate   *string `json:"arr_date,omitempty" xml:"arr_date,omitempty"`
	Attribute *string `json:"attribute,omitempty" xml:"attribute,omitempty"`
	// example:
	//
	// accac
	CostCenterName *string `json:"cost_center_name,omitempty" xml:"cost_center_name,omitempty"`
	DepCity        *string `json:"dep_city,omitempty" xml:"dep_city,omitempty"`
	// example:
	//
	// 330100
	DepCityCode *string `json:"dep_city_code,omitempty" xml:"dep_city_code,omitempty"`
	// example:
	//
	// 2018-09-19T14:03Z
	DepDate                 *string                                                           `json:"dep_date,omitempty" xml:"dep_date,omitempty"`
	InvoiceName             *string                                                           `json:"invoice_name,omitempty" xml:"invoice_name,omitempty"`
	ItineraryId             *string                                                           `json:"itinerary_id,omitempty" xml:"itinerary_id,omitempty"`
	ItineraryTravelStandard *ApplyQueryResponseBodyModuleItineraryListItineraryTravelStandard `json:"itinerary_travel_standard,omitempty" xml:"itinerary_travel_standard,omitempty" type:"Struct"`
	// example:
	//
	// xm1
	ProjectCode           *string `json:"project_code,omitempty" xml:"project_code,omitempty"`
	ProjectTitle          *string `json:"project_title,omitempty" xml:"project_title,omitempty"`
	ThirdpartCostCenterId *string `json:"thirdpart_cost_center_id,omitempty" xml:"thirdpart_cost_center_id,omitempty"`
	ThirdpartInvoiceId    *string `json:"thirdpart_invoice_id,omitempty" xml:"thirdpart_invoice_id,omitempty"`
	ThirdpartItineraryId  *string `json:"thirdpart_itinerary_id,omitempty" xml:"thirdpart_itinerary_id,omitempty"`
	// example:
	//
	// 1
	TrafficType *int32 `json:"traffic_type,omitempty" xml:"traffic_type,omitempty"`
	// example:
	//
	// 1
	TripWay *int32 `json:"trip_way,omitempty" xml:"trip_way,omitempty"`
}

func (s ApplyQueryResponseBodyModuleItineraryList) String() string {
	return dara.Prettify(s)
}

func (s ApplyQueryResponseBodyModuleItineraryList) GoString() string {
	return s.String()
}

func (s *ApplyQueryResponseBodyModuleItineraryList) GetArrCity() *string {
	return s.ArrCity
}

func (s *ApplyQueryResponseBodyModuleItineraryList) GetArrCityCode() *string {
	return s.ArrCityCode
}

func (s *ApplyQueryResponseBodyModuleItineraryList) GetArrDate() *string {
	return s.ArrDate
}

func (s *ApplyQueryResponseBodyModuleItineraryList) GetAttribute() *string {
	return s.Attribute
}

func (s *ApplyQueryResponseBodyModuleItineraryList) GetCostCenterName() *string {
	return s.CostCenterName
}

func (s *ApplyQueryResponseBodyModuleItineraryList) GetDepCity() *string {
	return s.DepCity
}

func (s *ApplyQueryResponseBodyModuleItineraryList) GetDepCityCode() *string {
	return s.DepCityCode
}

func (s *ApplyQueryResponseBodyModuleItineraryList) GetDepDate() *string {
	return s.DepDate
}

func (s *ApplyQueryResponseBodyModuleItineraryList) GetInvoiceName() *string {
	return s.InvoiceName
}

func (s *ApplyQueryResponseBodyModuleItineraryList) GetItineraryId() *string {
	return s.ItineraryId
}

func (s *ApplyQueryResponseBodyModuleItineraryList) GetItineraryTravelStandard() *ApplyQueryResponseBodyModuleItineraryListItineraryTravelStandard {
	return s.ItineraryTravelStandard
}

func (s *ApplyQueryResponseBodyModuleItineraryList) GetProjectCode() *string {
	return s.ProjectCode
}

func (s *ApplyQueryResponseBodyModuleItineraryList) GetProjectTitle() *string {
	return s.ProjectTitle
}

func (s *ApplyQueryResponseBodyModuleItineraryList) GetThirdpartCostCenterId() *string {
	return s.ThirdpartCostCenterId
}

func (s *ApplyQueryResponseBodyModuleItineraryList) GetThirdpartInvoiceId() *string {
	return s.ThirdpartInvoiceId
}

func (s *ApplyQueryResponseBodyModuleItineraryList) GetThirdpartItineraryId() *string {
	return s.ThirdpartItineraryId
}

func (s *ApplyQueryResponseBodyModuleItineraryList) GetTrafficType() *int32 {
	return s.TrafficType
}

func (s *ApplyQueryResponseBodyModuleItineraryList) GetTripWay() *int32 {
	return s.TripWay
}

func (s *ApplyQueryResponseBodyModuleItineraryList) SetArrCity(v string) *ApplyQueryResponseBodyModuleItineraryList {
	s.ArrCity = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleItineraryList) SetArrCityCode(v string) *ApplyQueryResponseBodyModuleItineraryList {
	s.ArrCityCode = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleItineraryList) SetArrDate(v string) *ApplyQueryResponseBodyModuleItineraryList {
	s.ArrDate = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleItineraryList) SetAttribute(v string) *ApplyQueryResponseBodyModuleItineraryList {
	s.Attribute = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleItineraryList) SetCostCenterName(v string) *ApplyQueryResponseBodyModuleItineraryList {
	s.CostCenterName = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleItineraryList) SetDepCity(v string) *ApplyQueryResponseBodyModuleItineraryList {
	s.DepCity = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleItineraryList) SetDepCityCode(v string) *ApplyQueryResponseBodyModuleItineraryList {
	s.DepCityCode = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleItineraryList) SetDepDate(v string) *ApplyQueryResponseBodyModuleItineraryList {
	s.DepDate = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleItineraryList) SetInvoiceName(v string) *ApplyQueryResponseBodyModuleItineraryList {
	s.InvoiceName = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleItineraryList) SetItineraryId(v string) *ApplyQueryResponseBodyModuleItineraryList {
	s.ItineraryId = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleItineraryList) SetItineraryTravelStandard(v *ApplyQueryResponseBodyModuleItineraryListItineraryTravelStandard) *ApplyQueryResponseBodyModuleItineraryList {
	s.ItineraryTravelStandard = v
	return s
}

func (s *ApplyQueryResponseBodyModuleItineraryList) SetProjectCode(v string) *ApplyQueryResponseBodyModuleItineraryList {
	s.ProjectCode = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleItineraryList) SetProjectTitle(v string) *ApplyQueryResponseBodyModuleItineraryList {
	s.ProjectTitle = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleItineraryList) SetThirdpartCostCenterId(v string) *ApplyQueryResponseBodyModuleItineraryList {
	s.ThirdpartCostCenterId = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleItineraryList) SetThirdpartInvoiceId(v string) *ApplyQueryResponseBodyModuleItineraryList {
	s.ThirdpartInvoiceId = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleItineraryList) SetThirdpartItineraryId(v string) *ApplyQueryResponseBodyModuleItineraryList {
	s.ThirdpartItineraryId = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleItineraryList) SetTrafficType(v int32) *ApplyQueryResponseBodyModuleItineraryList {
	s.TrafficType = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleItineraryList) SetTripWay(v int32) *ApplyQueryResponseBodyModuleItineraryList {
	s.TripWay = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleItineraryList) Validate() error {
	return dara.Validate(s)
}

type ApplyQueryResponseBodyModuleItineraryListItineraryTravelStandard struct {
	HotelAvailableNightsPerDay *int32 `json:"hotel_available_nights_per_day,omitempty" xml:"hotel_available_nights_per_day,omitempty"`
}

func (s ApplyQueryResponseBodyModuleItineraryListItineraryTravelStandard) String() string {
	return dara.Prettify(s)
}

func (s ApplyQueryResponseBodyModuleItineraryListItineraryTravelStandard) GoString() string {
	return s.String()
}

func (s *ApplyQueryResponseBodyModuleItineraryListItineraryTravelStandard) GetHotelAvailableNightsPerDay() *int32 {
	return s.HotelAvailableNightsPerDay
}

func (s *ApplyQueryResponseBodyModuleItineraryListItineraryTravelStandard) SetHotelAvailableNightsPerDay(v int32) *ApplyQueryResponseBodyModuleItineraryListItineraryTravelStandard {
	s.HotelAvailableNightsPerDay = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleItineraryListItineraryTravelStandard) Validate() error {
	return dara.Validate(s)
}

type ApplyQueryResponseBodyModuleItinerarySetList struct {
	// example:
	//
	// 2018-09-19T14:03Z
	ArrDate   *string `json:"arr_date,omitempty" xml:"arr_date,omitempty"`
	Attribute *string `json:"attribute,omitempty" xml:"attribute,omitempty"`
	// example:
	//
	// BJS，HGH
	CityCodeSet    *string `json:"city_code_set,omitempty" xml:"city_code_set,omitempty"`
	CitySet        *string `json:"city_set,omitempty" xml:"city_set,omitempty"`
	CostCenterName *string `json:"cost_center_name,omitempty" xml:"cost_center_name,omitempty"`
	// example:
	//
	// 2018-09-19T14:03Z
	DepDate     *string `json:"dep_date,omitempty" xml:"dep_date,omitempty"`
	InvoiceName *string `json:"invoice_name,omitempty" xml:"invoice_name,omitempty"`
	// example:
	//
	// 12345
	ItineraryId             *string                                                              `json:"itinerary_id,omitempty" xml:"itinerary_id,omitempty"`
	ItineraryTravelStandard *ApplyQueryResponseBodyModuleItinerarySetListItineraryTravelStandard `json:"itinerary_travel_standard,omitempty" xml:"itinerary_travel_standard,omitempty" type:"Struct"`
	// example:
	//
	// projecttow
	ProjectCode           *string `json:"project_code,omitempty" xml:"project_code,omitempty"`
	ProjectTitle          *string `json:"project_title,omitempty" xml:"project_title,omitempty"`
	ThirdpartCostCenterId *string `json:"thirdpart_cost_center_id,omitempty" xml:"thirdpart_cost_center_id,omitempty"`
	ThirdpartInvoiceId    *string `json:"thirdpart_invoice_id,omitempty" xml:"thirdpart_invoice_id,omitempty"`
	ThirdpartItineraryId  *string `json:"thirdpart_itinerary_id,omitempty" xml:"thirdpart_itinerary_id,omitempty"`
	// example:
	//
	// 0
	TrafficType *int32 `json:"traffic_type,omitempty" xml:"traffic_type,omitempty"`
}

func (s ApplyQueryResponseBodyModuleItinerarySetList) String() string {
	return dara.Prettify(s)
}

func (s ApplyQueryResponseBodyModuleItinerarySetList) GoString() string {
	return s.String()
}

func (s *ApplyQueryResponseBodyModuleItinerarySetList) GetArrDate() *string {
	return s.ArrDate
}

func (s *ApplyQueryResponseBodyModuleItinerarySetList) GetAttribute() *string {
	return s.Attribute
}

func (s *ApplyQueryResponseBodyModuleItinerarySetList) GetCityCodeSet() *string {
	return s.CityCodeSet
}

func (s *ApplyQueryResponseBodyModuleItinerarySetList) GetCitySet() *string {
	return s.CitySet
}

func (s *ApplyQueryResponseBodyModuleItinerarySetList) GetCostCenterName() *string {
	return s.CostCenterName
}

func (s *ApplyQueryResponseBodyModuleItinerarySetList) GetDepDate() *string {
	return s.DepDate
}

func (s *ApplyQueryResponseBodyModuleItinerarySetList) GetInvoiceName() *string {
	return s.InvoiceName
}

func (s *ApplyQueryResponseBodyModuleItinerarySetList) GetItineraryId() *string {
	return s.ItineraryId
}

func (s *ApplyQueryResponseBodyModuleItinerarySetList) GetItineraryTravelStandard() *ApplyQueryResponseBodyModuleItinerarySetListItineraryTravelStandard {
	return s.ItineraryTravelStandard
}

func (s *ApplyQueryResponseBodyModuleItinerarySetList) GetProjectCode() *string {
	return s.ProjectCode
}

func (s *ApplyQueryResponseBodyModuleItinerarySetList) GetProjectTitle() *string {
	return s.ProjectTitle
}

func (s *ApplyQueryResponseBodyModuleItinerarySetList) GetThirdpartCostCenterId() *string {
	return s.ThirdpartCostCenterId
}

func (s *ApplyQueryResponseBodyModuleItinerarySetList) GetThirdpartInvoiceId() *string {
	return s.ThirdpartInvoiceId
}

func (s *ApplyQueryResponseBodyModuleItinerarySetList) GetThirdpartItineraryId() *string {
	return s.ThirdpartItineraryId
}

func (s *ApplyQueryResponseBodyModuleItinerarySetList) GetTrafficType() *int32 {
	return s.TrafficType
}

func (s *ApplyQueryResponseBodyModuleItinerarySetList) SetArrDate(v string) *ApplyQueryResponseBodyModuleItinerarySetList {
	s.ArrDate = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleItinerarySetList) SetAttribute(v string) *ApplyQueryResponseBodyModuleItinerarySetList {
	s.Attribute = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleItinerarySetList) SetCityCodeSet(v string) *ApplyQueryResponseBodyModuleItinerarySetList {
	s.CityCodeSet = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleItinerarySetList) SetCitySet(v string) *ApplyQueryResponseBodyModuleItinerarySetList {
	s.CitySet = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleItinerarySetList) SetCostCenterName(v string) *ApplyQueryResponseBodyModuleItinerarySetList {
	s.CostCenterName = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleItinerarySetList) SetDepDate(v string) *ApplyQueryResponseBodyModuleItinerarySetList {
	s.DepDate = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleItinerarySetList) SetInvoiceName(v string) *ApplyQueryResponseBodyModuleItinerarySetList {
	s.InvoiceName = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleItinerarySetList) SetItineraryId(v string) *ApplyQueryResponseBodyModuleItinerarySetList {
	s.ItineraryId = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleItinerarySetList) SetItineraryTravelStandard(v *ApplyQueryResponseBodyModuleItinerarySetListItineraryTravelStandard) *ApplyQueryResponseBodyModuleItinerarySetList {
	s.ItineraryTravelStandard = v
	return s
}

func (s *ApplyQueryResponseBodyModuleItinerarySetList) SetProjectCode(v string) *ApplyQueryResponseBodyModuleItinerarySetList {
	s.ProjectCode = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleItinerarySetList) SetProjectTitle(v string) *ApplyQueryResponseBodyModuleItinerarySetList {
	s.ProjectTitle = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleItinerarySetList) SetThirdpartCostCenterId(v string) *ApplyQueryResponseBodyModuleItinerarySetList {
	s.ThirdpartCostCenterId = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleItinerarySetList) SetThirdpartInvoiceId(v string) *ApplyQueryResponseBodyModuleItinerarySetList {
	s.ThirdpartInvoiceId = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleItinerarySetList) SetThirdpartItineraryId(v string) *ApplyQueryResponseBodyModuleItinerarySetList {
	s.ThirdpartItineraryId = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleItinerarySetList) SetTrafficType(v int32) *ApplyQueryResponseBodyModuleItinerarySetList {
	s.TrafficType = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleItinerarySetList) Validate() error {
	return dara.Validate(s)
}

type ApplyQueryResponseBodyModuleItinerarySetListItineraryTravelStandard struct {
	HotelAvailableNightsPerDay *int32 `json:"hotel_available_nights_per_day,omitempty" xml:"hotel_available_nights_per_day,omitempty"`
}

func (s ApplyQueryResponseBodyModuleItinerarySetListItineraryTravelStandard) String() string {
	return dara.Prettify(s)
}

func (s ApplyQueryResponseBodyModuleItinerarySetListItineraryTravelStandard) GoString() string {
	return s.String()
}

func (s *ApplyQueryResponseBodyModuleItinerarySetListItineraryTravelStandard) GetHotelAvailableNightsPerDay() *int32 {
	return s.HotelAvailableNightsPerDay
}

func (s *ApplyQueryResponseBodyModuleItinerarySetListItineraryTravelStandard) SetHotelAvailableNightsPerDay(v int32) *ApplyQueryResponseBodyModuleItinerarySetListItineraryTravelStandard {
	s.HotelAvailableNightsPerDay = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleItinerarySetListItineraryTravelStandard) Validate() error {
	return dara.Validate(s)
}

type ApplyQueryResponseBodyModuleTravelerList struct {
	Attribute *string `json:"attribute,omitempty" xml:"attribute,omitempty"`
	// example:
	//
	// 1
	BusinessDiscount *int32                                                `json:"business_discount,omitempty" xml:"business_discount,omitempty"`
	CarCitySet       []*ApplyQueryResponseBodyModuleTravelerListCarCitySet `json:"car_city_set,omitempty" xml:"car_city_set,omitempty" type:"Repeated"`
	CostCenterName   *string                                               `json:"cost_center_name,omitempty" xml:"cost_center_name,omitempty"`
	DepartId         *string                                               `json:"depart_id,omitempty" xml:"depart_id,omitempty"`
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
	FlightCabins          *string                                                   `json:"flight_cabins,omitempty" xml:"flight_cabins,omitempty"`
	FlightIntlRuleCode    *int64                                                    `json:"flight_intl_rule_code,omitempty" xml:"flight_intl_rule_code,omitempty"`
	FlightRuleCode        *int64                                                    `json:"flight_rule_code,omitempty" xml:"flight_rule_code,omitempty"`
	HotelCitys            []*ApplyQueryResponseBodyModuleTravelerListHotelCitys     `json:"hotel_citys,omitempty" xml:"hotel_citys,omitempty" type:"Repeated"`
	HotelIntlCitys        []*ApplyQueryResponseBodyModuleTravelerListHotelIntlCitys `json:"hotel_intl_citys,omitempty" xml:"hotel_intl_citys,omitempty" type:"Repeated"`
	HotelIntlRuleCode     *int64                                                    `json:"hotel_intl_rule_code,omitempty" xml:"hotel_intl_rule_code,omitempty"`
	HotelRuleCode         *int64                                                    `json:"hotel_rule_code,omitempty" xml:"hotel_rule_code,omitempty"`
	InvoiceName           *string                                                   `json:"invoice_name,omitempty" xml:"invoice_name,omitempty"`
	PaymentDepartmentId   *string                                                   `json:"payment_department_id,omitempty" xml:"payment_department_id,omitempty"`
	PaymentDepartmentName *string                                                   `json:"payment_department_name,omitempty" xml:"payment_department_name,omitempty"`
	// example:
	//
	// 1
	PremiumEconomyDiscount *int32  `json:"premium_economy_discount,omitempty" xml:"premium_economy_discount,omitempty"`
	ProjectCode            *string `json:"project_code,omitempty" xml:"project_code,omitempty"`
	ProjectTitle           *string `json:"project_title,omitempty" xml:"project_title,omitempty"`
	// example:
	//
	// 1
	ReserveType           *int32  `json:"reserve_type,omitempty" xml:"reserve_type,omitempty"`
	ThirdPartInvoiceId    *string `json:"third_part_invoice_id,omitempty" xml:"third_part_invoice_id,omitempty"`
	ThirdpartCostCenterId *string `json:"thirdpart_cost_center_id,omitempty" xml:"thirdpart_cost_center_id,omitempty"`
	ThirdpartDepartId     *string `json:"thirdpart_depart_id,omitempty" xml:"thirdpart_depart_id,omitempty"`
	TrainRuleCode         *int64  `json:"train_rule_code,omitempty" xml:"train_rule_code,omitempty"`
	// example:
	//
	// 1
	TrainSeats *string `json:"train_seats,omitempty" xml:"train_seats,omitempty"`
	// example:
	//
	// 3423
	UserId   *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	UserName *string `json:"user_name,omitempty" xml:"user_name,omitempty"`
}

func (s ApplyQueryResponseBodyModuleTravelerList) String() string {
	return dara.Prettify(s)
}

func (s ApplyQueryResponseBodyModuleTravelerList) GoString() string {
	return s.String()
}

func (s *ApplyQueryResponseBodyModuleTravelerList) GetAttribute() *string {
	return s.Attribute
}

func (s *ApplyQueryResponseBodyModuleTravelerList) GetBusinessDiscount() *int32 {
	return s.BusinessDiscount
}

func (s *ApplyQueryResponseBodyModuleTravelerList) GetCarCitySet() []*ApplyQueryResponseBodyModuleTravelerListCarCitySet {
	return s.CarCitySet
}

func (s *ApplyQueryResponseBodyModuleTravelerList) GetCostCenterName() *string {
	return s.CostCenterName
}

func (s *ApplyQueryResponseBodyModuleTravelerList) GetDepartId() *string {
	return s.DepartId
}

func (s *ApplyQueryResponseBodyModuleTravelerList) GetEconomyDiscount() *int32 {
	return s.EconomyDiscount
}

func (s *ApplyQueryResponseBodyModuleTravelerList) GetFirstDiscount() *int32 {
	return s.FirstDiscount
}

func (s *ApplyQueryResponseBodyModuleTravelerList) GetFlightCabins() *string {
	return s.FlightCabins
}

func (s *ApplyQueryResponseBodyModuleTravelerList) GetFlightIntlRuleCode() *int64 {
	return s.FlightIntlRuleCode
}

func (s *ApplyQueryResponseBodyModuleTravelerList) GetFlightRuleCode() *int64 {
	return s.FlightRuleCode
}

func (s *ApplyQueryResponseBodyModuleTravelerList) GetHotelCitys() []*ApplyQueryResponseBodyModuleTravelerListHotelCitys {
	return s.HotelCitys
}

func (s *ApplyQueryResponseBodyModuleTravelerList) GetHotelIntlCitys() []*ApplyQueryResponseBodyModuleTravelerListHotelIntlCitys {
	return s.HotelIntlCitys
}

func (s *ApplyQueryResponseBodyModuleTravelerList) GetHotelIntlRuleCode() *int64 {
	return s.HotelIntlRuleCode
}

func (s *ApplyQueryResponseBodyModuleTravelerList) GetHotelRuleCode() *int64 {
	return s.HotelRuleCode
}

func (s *ApplyQueryResponseBodyModuleTravelerList) GetInvoiceName() *string {
	return s.InvoiceName
}

func (s *ApplyQueryResponseBodyModuleTravelerList) GetPaymentDepartmentId() *string {
	return s.PaymentDepartmentId
}

func (s *ApplyQueryResponseBodyModuleTravelerList) GetPaymentDepartmentName() *string {
	return s.PaymentDepartmentName
}

func (s *ApplyQueryResponseBodyModuleTravelerList) GetPremiumEconomyDiscount() *int32 {
	return s.PremiumEconomyDiscount
}

func (s *ApplyQueryResponseBodyModuleTravelerList) GetProjectCode() *string {
	return s.ProjectCode
}

func (s *ApplyQueryResponseBodyModuleTravelerList) GetProjectTitle() *string {
	return s.ProjectTitle
}

func (s *ApplyQueryResponseBodyModuleTravelerList) GetReserveType() *int32 {
	return s.ReserveType
}

func (s *ApplyQueryResponseBodyModuleTravelerList) GetThirdPartInvoiceId() *string {
	return s.ThirdPartInvoiceId
}

func (s *ApplyQueryResponseBodyModuleTravelerList) GetThirdpartCostCenterId() *string {
	return s.ThirdpartCostCenterId
}

func (s *ApplyQueryResponseBodyModuleTravelerList) GetThirdpartDepartId() *string {
	return s.ThirdpartDepartId
}

func (s *ApplyQueryResponseBodyModuleTravelerList) GetTrainRuleCode() *int64 {
	return s.TrainRuleCode
}

func (s *ApplyQueryResponseBodyModuleTravelerList) GetTrainSeats() *string {
	return s.TrainSeats
}

func (s *ApplyQueryResponseBodyModuleTravelerList) GetUserId() *string {
	return s.UserId
}

func (s *ApplyQueryResponseBodyModuleTravelerList) GetUserName() *string {
	return s.UserName
}

func (s *ApplyQueryResponseBodyModuleTravelerList) SetAttribute(v string) *ApplyQueryResponseBodyModuleTravelerList {
	s.Attribute = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleTravelerList) SetBusinessDiscount(v int32) *ApplyQueryResponseBodyModuleTravelerList {
	s.BusinessDiscount = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleTravelerList) SetCarCitySet(v []*ApplyQueryResponseBodyModuleTravelerListCarCitySet) *ApplyQueryResponseBodyModuleTravelerList {
	s.CarCitySet = v
	return s
}

func (s *ApplyQueryResponseBodyModuleTravelerList) SetCostCenterName(v string) *ApplyQueryResponseBodyModuleTravelerList {
	s.CostCenterName = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleTravelerList) SetDepartId(v string) *ApplyQueryResponseBodyModuleTravelerList {
	s.DepartId = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleTravelerList) SetEconomyDiscount(v int32) *ApplyQueryResponseBodyModuleTravelerList {
	s.EconomyDiscount = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleTravelerList) SetFirstDiscount(v int32) *ApplyQueryResponseBodyModuleTravelerList {
	s.FirstDiscount = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleTravelerList) SetFlightCabins(v string) *ApplyQueryResponseBodyModuleTravelerList {
	s.FlightCabins = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleTravelerList) SetFlightIntlRuleCode(v int64) *ApplyQueryResponseBodyModuleTravelerList {
	s.FlightIntlRuleCode = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleTravelerList) SetFlightRuleCode(v int64) *ApplyQueryResponseBodyModuleTravelerList {
	s.FlightRuleCode = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleTravelerList) SetHotelCitys(v []*ApplyQueryResponseBodyModuleTravelerListHotelCitys) *ApplyQueryResponseBodyModuleTravelerList {
	s.HotelCitys = v
	return s
}

func (s *ApplyQueryResponseBodyModuleTravelerList) SetHotelIntlCitys(v []*ApplyQueryResponseBodyModuleTravelerListHotelIntlCitys) *ApplyQueryResponseBodyModuleTravelerList {
	s.HotelIntlCitys = v
	return s
}

func (s *ApplyQueryResponseBodyModuleTravelerList) SetHotelIntlRuleCode(v int64) *ApplyQueryResponseBodyModuleTravelerList {
	s.HotelIntlRuleCode = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleTravelerList) SetHotelRuleCode(v int64) *ApplyQueryResponseBodyModuleTravelerList {
	s.HotelRuleCode = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleTravelerList) SetInvoiceName(v string) *ApplyQueryResponseBodyModuleTravelerList {
	s.InvoiceName = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleTravelerList) SetPaymentDepartmentId(v string) *ApplyQueryResponseBodyModuleTravelerList {
	s.PaymentDepartmentId = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleTravelerList) SetPaymentDepartmentName(v string) *ApplyQueryResponseBodyModuleTravelerList {
	s.PaymentDepartmentName = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleTravelerList) SetPremiumEconomyDiscount(v int32) *ApplyQueryResponseBodyModuleTravelerList {
	s.PremiumEconomyDiscount = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleTravelerList) SetProjectCode(v string) *ApplyQueryResponseBodyModuleTravelerList {
	s.ProjectCode = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleTravelerList) SetProjectTitle(v string) *ApplyQueryResponseBodyModuleTravelerList {
	s.ProjectTitle = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleTravelerList) SetReserveType(v int32) *ApplyQueryResponseBodyModuleTravelerList {
	s.ReserveType = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleTravelerList) SetThirdPartInvoiceId(v string) *ApplyQueryResponseBodyModuleTravelerList {
	s.ThirdPartInvoiceId = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleTravelerList) SetThirdpartCostCenterId(v string) *ApplyQueryResponseBodyModuleTravelerList {
	s.ThirdpartCostCenterId = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleTravelerList) SetThirdpartDepartId(v string) *ApplyQueryResponseBodyModuleTravelerList {
	s.ThirdpartDepartId = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleTravelerList) SetTrainRuleCode(v int64) *ApplyQueryResponseBodyModuleTravelerList {
	s.TrainRuleCode = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleTravelerList) SetTrainSeats(v string) *ApplyQueryResponseBodyModuleTravelerList {
	s.TrainSeats = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleTravelerList) SetUserId(v string) *ApplyQueryResponseBodyModuleTravelerList {
	s.UserId = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleTravelerList) SetUserName(v string) *ApplyQueryResponseBodyModuleTravelerList {
	s.UserName = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleTravelerList) Validate() error {
	return dara.Validate(s)
}

type ApplyQueryResponseBodyModuleTravelerListCarCitySet struct {
	CityCode *string `json:"city_code,omitempty" xml:"city_code,omitempty"`
	CityName *string `json:"city_name,omitempty" xml:"city_name,omitempty"`
}

func (s ApplyQueryResponseBodyModuleTravelerListCarCitySet) String() string {
	return dara.Prettify(s)
}

func (s ApplyQueryResponseBodyModuleTravelerListCarCitySet) GoString() string {
	return s.String()
}

func (s *ApplyQueryResponseBodyModuleTravelerListCarCitySet) GetCityCode() *string {
	return s.CityCode
}

func (s *ApplyQueryResponseBodyModuleTravelerListCarCitySet) GetCityName() *string {
	return s.CityName
}

func (s *ApplyQueryResponseBodyModuleTravelerListCarCitySet) SetCityCode(v string) *ApplyQueryResponseBodyModuleTravelerListCarCitySet {
	s.CityCode = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleTravelerListCarCitySet) SetCityName(v string) *ApplyQueryResponseBodyModuleTravelerListCarCitySet {
	s.CityName = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleTravelerListCarCitySet) Validate() error {
	return dara.Validate(s)
}

type ApplyQueryResponseBodyModuleTravelerListHotelCitys struct {
	// example:
	//
	// 0
	CityCode *string `json:"city_code,omitempty" xml:"city_code,omitempty"`
	CityName *string `json:"city_name,omitempty" xml:"city_name,omitempty"`
	// example:
	//
	// 100000
	Fee *int64 `json:"fee,omitempty" xml:"fee,omitempty"`
}

func (s ApplyQueryResponseBodyModuleTravelerListHotelCitys) String() string {
	return dara.Prettify(s)
}

func (s ApplyQueryResponseBodyModuleTravelerListHotelCitys) GoString() string {
	return s.String()
}

func (s *ApplyQueryResponseBodyModuleTravelerListHotelCitys) GetCityCode() *string {
	return s.CityCode
}

func (s *ApplyQueryResponseBodyModuleTravelerListHotelCitys) GetCityName() *string {
	return s.CityName
}

func (s *ApplyQueryResponseBodyModuleTravelerListHotelCitys) GetFee() *int64 {
	return s.Fee
}

func (s *ApplyQueryResponseBodyModuleTravelerListHotelCitys) SetCityCode(v string) *ApplyQueryResponseBodyModuleTravelerListHotelCitys {
	s.CityCode = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleTravelerListHotelCitys) SetCityName(v string) *ApplyQueryResponseBodyModuleTravelerListHotelCitys {
	s.CityName = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleTravelerListHotelCitys) SetFee(v int64) *ApplyQueryResponseBodyModuleTravelerListHotelCitys {
	s.Fee = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleTravelerListHotelCitys) Validate() error {
	return dara.Validate(s)
}

type ApplyQueryResponseBodyModuleTravelerListHotelIntlCitys struct {
	CityCode *string `json:"city_code,omitempty" xml:"city_code,omitempty"`
	CityName *string `json:"city_name,omitempty" xml:"city_name,omitempty"`
	Fee      *int64  `json:"fee,omitempty" xml:"fee,omitempty"`
}

func (s ApplyQueryResponseBodyModuleTravelerListHotelIntlCitys) String() string {
	return dara.Prettify(s)
}

func (s ApplyQueryResponseBodyModuleTravelerListHotelIntlCitys) GoString() string {
	return s.String()
}

func (s *ApplyQueryResponseBodyModuleTravelerListHotelIntlCitys) GetCityCode() *string {
	return s.CityCode
}

func (s *ApplyQueryResponseBodyModuleTravelerListHotelIntlCitys) GetCityName() *string {
	return s.CityName
}

func (s *ApplyQueryResponseBodyModuleTravelerListHotelIntlCitys) GetFee() *int64 {
	return s.Fee
}

func (s *ApplyQueryResponseBodyModuleTravelerListHotelIntlCitys) SetCityCode(v string) *ApplyQueryResponseBodyModuleTravelerListHotelIntlCitys {
	s.CityCode = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleTravelerListHotelIntlCitys) SetCityName(v string) *ApplyQueryResponseBodyModuleTravelerListHotelIntlCitys {
	s.CityName = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleTravelerListHotelIntlCitys) SetFee(v int64) *ApplyQueryResponseBodyModuleTravelerListHotelIntlCitys {
	s.Fee = &v
	return s
}

func (s *ApplyQueryResponseBodyModuleTravelerListHotelIntlCitys) Validate() error {
	return dara.Validate(s)
}
