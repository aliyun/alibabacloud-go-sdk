// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMealApplyQueryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *MealApplyQueryResponseBody
	GetCode() *string
	SetMessage(v string) *MealApplyQueryResponseBody
	GetMessage() *string
	SetModule(v *MealApplyQueryResponseBodyModule) *MealApplyQueryResponseBody
	GetModule() *MealApplyQueryResponseBodyModule
	SetRequestId(v string) *MealApplyQueryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *MealApplyQueryResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *MealApplyQueryResponseBody
	GetTraceId() *string
}

type MealApplyQueryResponseBody struct {
	// example:
	//
	// 0
	Code    *string                           `json:"code,omitempty" xml:"code,omitempty"`
	Message *string                           `json:"message,omitempty" xml:"message,omitempty"`
	Module  *MealApplyQueryResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	// example:
	//
	// A5009956-1077-52FB-B520-EA8C7E91D722
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// traceId
	//
	// example:
	//
	// 210f07f316603757445272547d959f
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s MealApplyQueryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s MealApplyQueryResponseBody) GoString() string {
	return s.String()
}

func (s *MealApplyQueryResponseBody) GetCode() *string {
	return s.Code
}

func (s *MealApplyQueryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *MealApplyQueryResponseBody) GetModule() *MealApplyQueryResponseBodyModule {
	return s.Module
}

func (s *MealApplyQueryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *MealApplyQueryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *MealApplyQueryResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *MealApplyQueryResponseBody) SetCode(v string) *MealApplyQueryResponseBody {
	s.Code = &v
	return s
}

func (s *MealApplyQueryResponseBody) SetMessage(v string) *MealApplyQueryResponseBody {
	s.Message = &v
	return s
}

func (s *MealApplyQueryResponseBody) SetModule(v *MealApplyQueryResponseBodyModule) *MealApplyQueryResponseBody {
	s.Module = v
	return s
}

func (s *MealApplyQueryResponseBody) SetRequestId(v string) *MealApplyQueryResponseBody {
	s.RequestId = &v
	return s
}

func (s *MealApplyQueryResponseBody) SetSuccess(v bool) *MealApplyQueryResponseBody {
	s.Success = &v
	return s
}

func (s *MealApplyQueryResponseBody) SetTraceId(v string) *MealApplyQueryResponseBody {
	s.TraceId = &v
	return s
}

func (s *MealApplyQueryResponseBody) Validate() error {
	return dara.Validate(s)
}

type MealApplyQueryResponseBodyModule struct {
	ApplyUser *MealApplyQueryResponseBodyModuleApplyUser `json:"apply_user,omitempty" xml:"apply_user,omitempty" type:"Struct"`
	// example:
	//
	// 11376
	CostCenterId *int64 `json:"cost_center_id,omitempty" xml:"cost_center_id,omitempty"`
	// example:
	//
	// 2022-07-04T16:13Z
	GmtCreate *string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// example:
	//
	// 11876
	InvoiceId     *int64                                           `json:"invoice_id,omitempty" xml:"invoice_id,omitempty"`
	ItineraryList []*MealApplyQueryResponseBodyModuleItineraryList `json:"itinerary_list,omitempty" xml:"itinerary_list,omitempty" type:"Repeated"`
	MealAmount    *int64                                           `json:"meal_amount,omitempty" xml:"meal_amount,omitempty"`
	MealCause     *string                                          `json:"meal_cause,omitempty" xml:"meal_cause,omitempty"`
	// example:
	//
	// 11546
	ProjectId *int64 `json:"project_id,omitempty" xml:"project_id,omitempty"`
	// example:
	//
	// 0
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 2024073128454753
	ThirdPartApplyId *string `json:"third_part_apply_id,omitempty" xml:"third_part_apply_id,omitempty"`
	// example:
	//
	// 330000303010292572
	ThirdPartCostCenterId *string `json:"third_part_cost_center_id,omitempty" xml:"third_part_cost_center_id,omitempty"`
	// example:
	//
	// 405009
	ThirdPartInvoiceId *string `json:"third_part_invoice_id,omitempty" xml:"third_part_invoice_id,omitempty"`
	// example:
	//
	// CS-PROJECT
	ThirdPartProjectId *string `json:"third_part_project_id,omitempty" xml:"third_part_project_id,omitempty"`
}

func (s MealApplyQueryResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s MealApplyQueryResponseBodyModule) GoString() string {
	return s.String()
}

func (s *MealApplyQueryResponseBodyModule) GetApplyUser() *MealApplyQueryResponseBodyModuleApplyUser {
	return s.ApplyUser
}

func (s *MealApplyQueryResponseBodyModule) GetCostCenterId() *int64 {
	return s.CostCenterId
}

func (s *MealApplyQueryResponseBodyModule) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *MealApplyQueryResponseBodyModule) GetInvoiceId() *int64 {
	return s.InvoiceId
}

func (s *MealApplyQueryResponseBodyModule) GetItineraryList() []*MealApplyQueryResponseBodyModuleItineraryList {
	return s.ItineraryList
}

func (s *MealApplyQueryResponseBodyModule) GetMealAmount() *int64 {
	return s.MealAmount
}

func (s *MealApplyQueryResponseBodyModule) GetMealCause() *string {
	return s.MealCause
}

func (s *MealApplyQueryResponseBodyModule) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *MealApplyQueryResponseBodyModule) GetStatus() *int32 {
	return s.Status
}

func (s *MealApplyQueryResponseBodyModule) GetThirdPartApplyId() *string {
	return s.ThirdPartApplyId
}

func (s *MealApplyQueryResponseBodyModule) GetThirdPartCostCenterId() *string {
	return s.ThirdPartCostCenterId
}

func (s *MealApplyQueryResponseBodyModule) GetThirdPartInvoiceId() *string {
	return s.ThirdPartInvoiceId
}

func (s *MealApplyQueryResponseBodyModule) GetThirdPartProjectId() *string {
	return s.ThirdPartProjectId
}

func (s *MealApplyQueryResponseBodyModule) SetApplyUser(v *MealApplyQueryResponseBodyModuleApplyUser) *MealApplyQueryResponseBodyModule {
	s.ApplyUser = v
	return s
}

func (s *MealApplyQueryResponseBodyModule) SetCostCenterId(v int64) *MealApplyQueryResponseBodyModule {
	s.CostCenterId = &v
	return s
}

func (s *MealApplyQueryResponseBodyModule) SetGmtCreate(v string) *MealApplyQueryResponseBodyModule {
	s.GmtCreate = &v
	return s
}

func (s *MealApplyQueryResponseBodyModule) SetInvoiceId(v int64) *MealApplyQueryResponseBodyModule {
	s.InvoiceId = &v
	return s
}

func (s *MealApplyQueryResponseBodyModule) SetItineraryList(v []*MealApplyQueryResponseBodyModuleItineraryList) *MealApplyQueryResponseBodyModule {
	s.ItineraryList = v
	return s
}

func (s *MealApplyQueryResponseBodyModule) SetMealAmount(v int64) *MealApplyQueryResponseBodyModule {
	s.MealAmount = &v
	return s
}

func (s *MealApplyQueryResponseBodyModule) SetMealCause(v string) *MealApplyQueryResponseBodyModule {
	s.MealCause = &v
	return s
}

func (s *MealApplyQueryResponseBodyModule) SetProjectId(v int64) *MealApplyQueryResponseBodyModule {
	s.ProjectId = &v
	return s
}

func (s *MealApplyQueryResponseBodyModule) SetStatus(v int32) *MealApplyQueryResponseBodyModule {
	s.Status = &v
	return s
}

func (s *MealApplyQueryResponseBodyModule) SetThirdPartApplyId(v string) *MealApplyQueryResponseBodyModule {
	s.ThirdPartApplyId = &v
	return s
}

func (s *MealApplyQueryResponseBodyModule) SetThirdPartCostCenterId(v string) *MealApplyQueryResponseBodyModule {
	s.ThirdPartCostCenterId = &v
	return s
}

func (s *MealApplyQueryResponseBodyModule) SetThirdPartInvoiceId(v string) *MealApplyQueryResponseBodyModule {
	s.ThirdPartInvoiceId = &v
	return s
}

func (s *MealApplyQueryResponseBodyModule) SetThirdPartProjectId(v string) *MealApplyQueryResponseBodyModule {
	s.ThirdPartProjectId = &v
	return s
}

func (s *MealApplyQueryResponseBodyModule) Validate() error {
	return dara.Validate(s)
}

type MealApplyQueryResponseBodyModuleApplyUser struct {
	// example:
	//
	// 2123
	UserId   *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	UserName *string `json:"user_name,omitempty" xml:"user_name,omitempty"`
}

func (s MealApplyQueryResponseBodyModuleApplyUser) String() string {
	return dara.Prettify(s)
}

func (s MealApplyQueryResponseBodyModuleApplyUser) GoString() string {
	return s.String()
}

func (s *MealApplyQueryResponseBodyModuleApplyUser) GetUserId() *string {
	return s.UserId
}

func (s *MealApplyQueryResponseBodyModuleApplyUser) GetUserName() *string {
	return s.UserName
}

func (s *MealApplyQueryResponseBodyModuleApplyUser) SetUserId(v string) *MealApplyQueryResponseBodyModuleApplyUser {
	s.UserId = &v
	return s
}

func (s *MealApplyQueryResponseBodyModuleApplyUser) SetUserName(v string) *MealApplyQueryResponseBodyModuleApplyUser {
	s.UserName = &v
	return s
}

func (s *MealApplyQueryResponseBodyModuleApplyUser) Validate() error {
	return dara.Validate(s)
}

type MealApplyQueryResponseBodyModuleItineraryList struct {
	Cities               []*MealApplyQueryResponseBodyModuleItineraryListCities `json:"cities,omitempty" xml:"cities,omitempty" type:"Repeated"`
	EndDate              *string                                                `json:"end_date,omitempty" xml:"end_date,omitempty"`
	StartDate            *string                                                `json:"start_date,omitempty" xml:"start_date,omitempty"`
	ThirdpartItineraryId *string                                                `json:"thirdpart_itinerary_id,omitempty" xml:"thirdpart_itinerary_id,omitempty"`
}

func (s MealApplyQueryResponseBodyModuleItineraryList) String() string {
	return dara.Prettify(s)
}

func (s MealApplyQueryResponseBodyModuleItineraryList) GoString() string {
	return s.String()
}

func (s *MealApplyQueryResponseBodyModuleItineraryList) GetCities() []*MealApplyQueryResponseBodyModuleItineraryListCities {
	return s.Cities
}

func (s *MealApplyQueryResponseBodyModuleItineraryList) GetEndDate() *string {
	return s.EndDate
}

func (s *MealApplyQueryResponseBodyModuleItineraryList) GetStartDate() *string {
	return s.StartDate
}

func (s *MealApplyQueryResponseBodyModuleItineraryList) GetThirdpartItineraryId() *string {
	return s.ThirdpartItineraryId
}

func (s *MealApplyQueryResponseBodyModuleItineraryList) SetCities(v []*MealApplyQueryResponseBodyModuleItineraryListCities) *MealApplyQueryResponseBodyModuleItineraryList {
	s.Cities = v
	return s
}

func (s *MealApplyQueryResponseBodyModuleItineraryList) SetEndDate(v string) *MealApplyQueryResponseBodyModuleItineraryList {
	s.EndDate = &v
	return s
}

func (s *MealApplyQueryResponseBodyModuleItineraryList) SetStartDate(v string) *MealApplyQueryResponseBodyModuleItineraryList {
	s.StartDate = &v
	return s
}

func (s *MealApplyQueryResponseBodyModuleItineraryList) SetThirdpartItineraryId(v string) *MealApplyQueryResponseBodyModuleItineraryList {
	s.ThirdpartItineraryId = &v
	return s
}

func (s *MealApplyQueryResponseBodyModuleItineraryList) Validate() error {
	return dara.Validate(s)
}

type MealApplyQueryResponseBodyModuleItineraryListCities struct {
	CityCode *string `json:"city_code,omitempty" xml:"city_code,omitempty"`
	CityName *string `json:"city_name,omitempty" xml:"city_name,omitempty"`
}

func (s MealApplyQueryResponseBodyModuleItineraryListCities) String() string {
	return dara.Prettify(s)
}

func (s MealApplyQueryResponseBodyModuleItineraryListCities) GoString() string {
	return s.String()
}

func (s *MealApplyQueryResponseBodyModuleItineraryListCities) GetCityCode() *string {
	return s.CityCode
}

func (s *MealApplyQueryResponseBodyModuleItineraryListCities) GetCityName() *string {
	return s.CityName
}

func (s *MealApplyQueryResponseBodyModuleItineraryListCities) SetCityCode(v string) *MealApplyQueryResponseBodyModuleItineraryListCities {
	s.CityCode = &v
	return s
}

func (s *MealApplyQueryResponseBodyModuleItineraryListCities) SetCityName(v string) *MealApplyQueryResponseBodyModuleItineraryListCities {
	s.CityName = &v
	return s
}

func (s *MealApplyQueryResponseBodyModuleItineraryListCities) Validate() error {
	return dara.Validate(s)
}
