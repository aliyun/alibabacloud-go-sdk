// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyListQueryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ApplyListQueryResponseBody
	GetCode() *string
	SetMessage(v string) *ApplyListQueryResponseBody
	GetMessage() *string
	SetModuleList(v []*ApplyListQueryResponseBodyModuleList) *ApplyListQueryResponseBody
	GetModuleList() []*ApplyListQueryResponseBodyModuleList
	SetRequestId(v string) *ApplyListQueryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ApplyListQueryResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *ApplyListQueryResponseBody
	GetTraceId() *string
}

type ApplyListQueryResponseBody struct {
	// example:
	//
	// SUCCESS
	Code       *string                                 `json:"code,omitempty" xml:"code,omitempty"`
	Message    *string                                 `json:"message,omitempty" xml:"message,omitempty"`
	ModuleList []*ApplyListQueryResponseBodyModuleList `json:"module_list,omitempty" xml:"module_list,omitempty" type:"Repeated"`
	// example:
	//
	// C61ECFF6-606B-5F66-B81D-D77369043A5F
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

func (s ApplyListQueryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ApplyListQueryResponseBody) GoString() string {
	return s.String()
}

func (s *ApplyListQueryResponseBody) GetCode() *string {
	return s.Code
}

func (s *ApplyListQueryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ApplyListQueryResponseBody) GetModuleList() []*ApplyListQueryResponseBodyModuleList {
	return s.ModuleList
}

func (s *ApplyListQueryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ApplyListQueryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ApplyListQueryResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *ApplyListQueryResponseBody) SetCode(v string) *ApplyListQueryResponseBody {
	s.Code = &v
	return s
}

func (s *ApplyListQueryResponseBody) SetMessage(v string) *ApplyListQueryResponseBody {
	s.Message = &v
	return s
}

func (s *ApplyListQueryResponseBody) SetModuleList(v []*ApplyListQueryResponseBodyModuleList) *ApplyListQueryResponseBody {
	s.ModuleList = v
	return s
}

func (s *ApplyListQueryResponseBody) SetRequestId(v string) *ApplyListQueryResponseBody {
	s.RequestId = &v
	return s
}

func (s *ApplyListQueryResponseBody) SetSuccess(v bool) *ApplyListQueryResponseBody {
	s.Success = &v
	return s
}

func (s *ApplyListQueryResponseBody) SetTraceId(v string) *ApplyListQueryResponseBody {
	s.TraceId = &v
	return s
}

func (s *ApplyListQueryResponseBody) Validate() error {
	return dara.Validate(s)
}

type ApplyListQueryResponseBodyModuleList struct {
	// example:
	//
	// 201710111505000464651
	ApplyShowId  *string                                             `json:"apply_show_id,omitempty" xml:"apply_show_id,omitempty"`
	ApproverList []*ApplyListQueryResponseBodyModuleListApproverList `json:"approver_list,omitempty" xml:"approver_list,omitempty" type:"Repeated"`
	CarRule      *ApplyListQueryResponseBodyModuleListCarRule        `json:"car_rule,omitempty" xml:"car_rule,omitempty" type:"Struct"`
	// example:
	//
	// corp1
	CorpId *string `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
	// example:
	//
	// abc
	CorpName *string `json:"corp_name,omitempty" xml:"corp_name,omitempty"`
	// example:
	//
	// depart1
	DepartId             *string                                                     `json:"depart_id,omitempty" xml:"depart_id,omitempty"`
	DepartName           *string                                                     `json:"depart_name,omitempty" xml:"depart_name,omitempty"`
	ExternalTravelerList []*ApplyListQueryResponseBodyModuleListExternalTravelerList `json:"external_traveler_list,omitempty" xml:"external_traveler_list,omitempty" type:"Repeated"`
	// example:
	//
	// abc1234
	FlowCode *string `json:"flow_code,omitempty" xml:"flow_code,omitempty"`
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
	// 1476
	Id            *int64                                               `json:"id,omitempty" xml:"id,omitempty"`
	ItineraryList []*ApplyListQueryResponseBodyModuleListItineraryList `json:"itinerary_list,omitempty" xml:"itinerary_list,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	ItineraryRule         *int32                                                  `json:"itinerary_rule,omitempty" xml:"itinerary_rule,omitempty"`
	ItinerarySetList      []*ApplyListQueryResponseBodyModuleListItinerarySetList `json:"itinerary_set_list,omitempty" xml:"itinerary_set_list,omitempty" type:"Repeated"`
	JobNo                 *string                                                 `json:"job_no,omitempty" xml:"job_no,omitempty"`
	PaymentDepartmentId   *string                                                 `json:"payment_department_id,omitempty" xml:"payment_department_id,omitempty"`
	PaymentDepartmentName *string                                                 `json:"payment_department_name,omitempty" xml:"payment_department_name,omitempty"`
	// example:
	//
	// 1
	Status     *int32  `json:"status,omitempty" xml:"status,omitempty"`
	StatusDesc *string `json:"status_desc,omitempty" xml:"status_desc,omitempty"`
	// example:
	//
	// abc
	ThirdpartBusinessId *string `json:"thirdpart_business_id,omitempty" xml:"thirdpart_business_id,omitempty"`
	// example:
	//
	// abc
	ThirdpartId  *string                                             `json:"thirdpart_id,omitempty" xml:"thirdpart_id,omitempty"`
	TravelerList []*ApplyListQueryResponseBodyModuleListTravelerList `json:"traveler_list,omitempty" xml:"traveler_list,omitempty" type:"Repeated"`
	TripCause    *string                                             `json:"trip_cause,omitempty" xml:"trip_cause,omitempty"`
	// example:
	//
	// 1
	TripDay   *int32  `json:"trip_day,omitempty" xml:"trip_day,omitempty"`
	TripTitle *string `json:"trip_title,omitempty" xml:"trip_title,omitempty"`
	// example:
	//
	// 2
	Type *int32 `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// abc
	UnionNo *string `json:"union_no,omitempty" xml:"union_no,omitempty"`
	// example:
	//
	// user1
	UserId   *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	UserName *string `json:"user_name,omitempty" xml:"user_name,omitempty"`
}

func (s ApplyListQueryResponseBodyModuleList) String() string {
	return dara.Prettify(s)
}

func (s ApplyListQueryResponseBodyModuleList) GoString() string {
	return s.String()
}

func (s *ApplyListQueryResponseBodyModuleList) GetApplyShowId() *string {
	return s.ApplyShowId
}

func (s *ApplyListQueryResponseBodyModuleList) GetApproverList() []*ApplyListQueryResponseBodyModuleListApproverList {
	return s.ApproverList
}

func (s *ApplyListQueryResponseBodyModuleList) GetCarRule() *ApplyListQueryResponseBodyModuleListCarRule {
	return s.CarRule
}

func (s *ApplyListQueryResponseBodyModuleList) GetCorpId() *string {
	return s.CorpId
}

func (s *ApplyListQueryResponseBodyModuleList) GetCorpName() *string {
	return s.CorpName
}

func (s *ApplyListQueryResponseBodyModuleList) GetDepartId() *string {
	return s.DepartId
}

func (s *ApplyListQueryResponseBodyModuleList) GetDepartName() *string {
	return s.DepartName
}

func (s *ApplyListQueryResponseBodyModuleList) GetExternalTravelerList() []*ApplyListQueryResponseBodyModuleListExternalTravelerList {
	return s.ExternalTravelerList
}

func (s *ApplyListQueryResponseBodyModuleList) GetFlowCode() *string {
	return s.FlowCode
}

func (s *ApplyListQueryResponseBodyModuleList) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ApplyListQueryResponseBodyModuleList) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ApplyListQueryResponseBodyModuleList) GetId() *int64 {
	return s.Id
}

func (s *ApplyListQueryResponseBodyModuleList) GetItineraryList() []*ApplyListQueryResponseBodyModuleListItineraryList {
	return s.ItineraryList
}

func (s *ApplyListQueryResponseBodyModuleList) GetItineraryRule() *int32 {
	return s.ItineraryRule
}

func (s *ApplyListQueryResponseBodyModuleList) GetItinerarySetList() []*ApplyListQueryResponseBodyModuleListItinerarySetList {
	return s.ItinerarySetList
}

func (s *ApplyListQueryResponseBodyModuleList) GetJobNo() *string {
	return s.JobNo
}

func (s *ApplyListQueryResponseBodyModuleList) GetPaymentDepartmentId() *string {
	return s.PaymentDepartmentId
}

func (s *ApplyListQueryResponseBodyModuleList) GetPaymentDepartmentName() *string {
	return s.PaymentDepartmentName
}

func (s *ApplyListQueryResponseBodyModuleList) GetStatus() *int32 {
	return s.Status
}

func (s *ApplyListQueryResponseBodyModuleList) GetStatusDesc() *string {
	return s.StatusDesc
}

func (s *ApplyListQueryResponseBodyModuleList) GetThirdpartBusinessId() *string {
	return s.ThirdpartBusinessId
}

func (s *ApplyListQueryResponseBodyModuleList) GetThirdpartId() *string {
	return s.ThirdpartId
}

func (s *ApplyListQueryResponseBodyModuleList) GetTravelerList() []*ApplyListQueryResponseBodyModuleListTravelerList {
	return s.TravelerList
}

func (s *ApplyListQueryResponseBodyModuleList) GetTripCause() *string {
	return s.TripCause
}

func (s *ApplyListQueryResponseBodyModuleList) GetTripDay() *int32 {
	return s.TripDay
}

func (s *ApplyListQueryResponseBodyModuleList) GetTripTitle() *string {
	return s.TripTitle
}

func (s *ApplyListQueryResponseBodyModuleList) GetType() *int32 {
	return s.Type
}

func (s *ApplyListQueryResponseBodyModuleList) GetUnionNo() *string {
	return s.UnionNo
}

func (s *ApplyListQueryResponseBodyModuleList) GetUserId() *string {
	return s.UserId
}

func (s *ApplyListQueryResponseBodyModuleList) GetUserName() *string {
	return s.UserName
}

func (s *ApplyListQueryResponseBodyModuleList) SetApplyShowId(v string) *ApplyListQueryResponseBodyModuleList {
	s.ApplyShowId = &v
	return s
}

func (s *ApplyListQueryResponseBodyModuleList) SetApproverList(v []*ApplyListQueryResponseBodyModuleListApproverList) *ApplyListQueryResponseBodyModuleList {
	s.ApproverList = v
	return s
}

func (s *ApplyListQueryResponseBodyModuleList) SetCarRule(v *ApplyListQueryResponseBodyModuleListCarRule) *ApplyListQueryResponseBodyModuleList {
	s.CarRule = v
	return s
}

func (s *ApplyListQueryResponseBodyModuleList) SetCorpId(v string) *ApplyListQueryResponseBodyModuleList {
	s.CorpId = &v
	return s
}

func (s *ApplyListQueryResponseBodyModuleList) SetCorpName(v string) *ApplyListQueryResponseBodyModuleList {
	s.CorpName = &v
	return s
}

func (s *ApplyListQueryResponseBodyModuleList) SetDepartId(v string) *ApplyListQueryResponseBodyModuleList {
	s.DepartId = &v
	return s
}

func (s *ApplyListQueryResponseBodyModuleList) SetDepartName(v string) *ApplyListQueryResponseBodyModuleList {
	s.DepartName = &v
	return s
}

func (s *ApplyListQueryResponseBodyModuleList) SetExternalTravelerList(v []*ApplyListQueryResponseBodyModuleListExternalTravelerList) *ApplyListQueryResponseBodyModuleList {
	s.ExternalTravelerList = v
	return s
}

func (s *ApplyListQueryResponseBodyModuleList) SetFlowCode(v string) *ApplyListQueryResponseBodyModuleList {
	s.FlowCode = &v
	return s
}

func (s *ApplyListQueryResponseBodyModuleList) SetGmtCreate(v string) *ApplyListQueryResponseBodyModuleList {
	s.GmtCreate = &v
	return s
}

func (s *ApplyListQueryResponseBodyModuleList) SetGmtModified(v string) *ApplyListQueryResponseBodyModuleList {
	s.GmtModified = &v
	return s
}

func (s *ApplyListQueryResponseBodyModuleList) SetId(v int64) *ApplyListQueryResponseBodyModuleList {
	s.Id = &v
	return s
}

func (s *ApplyListQueryResponseBodyModuleList) SetItineraryList(v []*ApplyListQueryResponseBodyModuleListItineraryList) *ApplyListQueryResponseBodyModuleList {
	s.ItineraryList = v
	return s
}

func (s *ApplyListQueryResponseBodyModuleList) SetItineraryRule(v int32) *ApplyListQueryResponseBodyModuleList {
	s.ItineraryRule = &v
	return s
}

func (s *ApplyListQueryResponseBodyModuleList) SetItinerarySetList(v []*ApplyListQueryResponseBodyModuleListItinerarySetList) *ApplyListQueryResponseBodyModuleList {
	s.ItinerarySetList = v
	return s
}

func (s *ApplyListQueryResponseBodyModuleList) SetJobNo(v string) *ApplyListQueryResponseBodyModuleList {
	s.JobNo = &v
	return s
}

func (s *ApplyListQueryResponseBodyModuleList) SetPaymentDepartmentId(v string) *ApplyListQueryResponseBodyModuleList {
	s.PaymentDepartmentId = &v
	return s
}

func (s *ApplyListQueryResponseBodyModuleList) SetPaymentDepartmentName(v string) *ApplyListQueryResponseBodyModuleList {
	s.PaymentDepartmentName = &v
	return s
}

func (s *ApplyListQueryResponseBodyModuleList) SetStatus(v int32) *ApplyListQueryResponseBodyModuleList {
	s.Status = &v
	return s
}

func (s *ApplyListQueryResponseBodyModuleList) SetStatusDesc(v string) *ApplyListQueryResponseBodyModuleList {
	s.StatusDesc = &v
	return s
}

func (s *ApplyListQueryResponseBodyModuleList) SetThirdpartBusinessId(v string) *ApplyListQueryResponseBodyModuleList {
	s.ThirdpartBusinessId = &v
	return s
}

func (s *ApplyListQueryResponseBodyModuleList) SetThirdpartId(v string) *ApplyListQueryResponseBodyModuleList {
	s.ThirdpartId = &v
	return s
}

func (s *ApplyListQueryResponseBodyModuleList) SetTravelerList(v []*ApplyListQueryResponseBodyModuleListTravelerList) *ApplyListQueryResponseBodyModuleList {
	s.TravelerList = v
	return s
}

func (s *ApplyListQueryResponseBodyModuleList) SetTripCause(v string) *ApplyListQueryResponseBodyModuleList {
	s.TripCause = &v
	return s
}

func (s *ApplyListQueryResponseBodyModuleList) SetTripDay(v int32) *ApplyListQueryResponseBodyModuleList {
	s.TripDay = &v
	return s
}

func (s *ApplyListQueryResponseBodyModuleList) SetTripTitle(v string) *ApplyListQueryResponseBodyModuleList {
	s.TripTitle = &v
	return s
}

func (s *ApplyListQueryResponseBodyModuleList) SetType(v int32) *ApplyListQueryResponseBodyModuleList {
	s.Type = &v
	return s
}

func (s *ApplyListQueryResponseBodyModuleList) SetUnionNo(v string) *ApplyListQueryResponseBodyModuleList {
	s.UnionNo = &v
	return s
}

func (s *ApplyListQueryResponseBodyModuleList) SetUserId(v string) *ApplyListQueryResponseBodyModuleList {
	s.UserId = &v
	return s
}

func (s *ApplyListQueryResponseBodyModuleList) SetUserName(v string) *ApplyListQueryResponseBodyModuleList {
	s.UserName = &v
	return s
}

func (s *ApplyListQueryResponseBodyModuleList) Validate() error {
	return dara.Validate(s)
}

type ApplyListQueryResponseBodyModuleListApproverList struct {
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
	// 1
	Status     *int32  `json:"status,omitempty" xml:"status,omitempty"`
	StatusDesc *string `json:"status_desc,omitempty" xml:"status_desc,omitempty"`
	// example:
	//
	// user1
	UserId   *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	UserName *string `json:"user_name,omitempty" xml:"user_name,omitempty"`
}

func (s ApplyListQueryResponseBodyModuleListApproverList) String() string {
	return dara.Prettify(s)
}

func (s ApplyListQueryResponseBodyModuleListApproverList) GoString() string {
	return s.String()
}

func (s *ApplyListQueryResponseBodyModuleListApproverList) GetNote() *string {
	return s.Note
}

func (s *ApplyListQueryResponseBodyModuleListApproverList) GetOperateTime() *string {
	return s.OperateTime
}

func (s *ApplyListQueryResponseBodyModuleListApproverList) GetOrder() *int32 {
	return s.Order
}

func (s *ApplyListQueryResponseBodyModuleListApproverList) GetStatus() *int32 {
	return s.Status
}

func (s *ApplyListQueryResponseBodyModuleListApproverList) GetStatusDesc() *string {
	return s.StatusDesc
}

func (s *ApplyListQueryResponseBodyModuleListApproverList) GetUserId() *string {
	return s.UserId
}

func (s *ApplyListQueryResponseBodyModuleListApproverList) GetUserName() *string {
	return s.UserName
}

func (s *ApplyListQueryResponseBodyModuleListApproverList) SetNote(v string) *ApplyListQueryResponseBodyModuleListApproverList {
	s.Note = &v
	return s
}

func (s *ApplyListQueryResponseBodyModuleListApproverList) SetOperateTime(v string) *ApplyListQueryResponseBodyModuleListApproverList {
	s.OperateTime = &v
	return s
}

func (s *ApplyListQueryResponseBodyModuleListApproverList) SetOrder(v int32) *ApplyListQueryResponseBodyModuleListApproverList {
	s.Order = &v
	return s
}

func (s *ApplyListQueryResponseBodyModuleListApproverList) SetStatus(v int32) *ApplyListQueryResponseBodyModuleListApproverList {
	s.Status = &v
	return s
}

func (s *ApplyListQueryResponseBodyModuleListApproverList) SetStatusDesc(v string) *ApplyListQueryResponseBodyModuleListApproverList {
	s.StatusDesc = &v
	return s
}

func (s *ApplyListQueryResponseBodyModuleListApproverList) SetUserId(v string) *ApplyListQueryResponseBodyModuleListApproverList {
	s.UserId = &v
	return s
}

func (s *ApplyListQueryResponseBodyModuleListApproverList) SetUserName(v string) *ApplyListQueryResponseBodyModuleListApproverList {
	s.UserName = &v
	return s
}

func (s *ApplyListQueryResponseBodyModuleListApproverList) Validate() error {
	return dara.Validate(s)
}

type ApplyListQueryResponseBodyModuleListCarRule struct {
	ScenarioTemplateId   *string `json:"scenario_template_id,omitempty" xml:"scenario_template_id,omitempty"`
	ScenarioTemplateName *string `json:"scenario_template_name,omitempty" xml:"scenario_template_name,omitempty"`
}

func (s ApplyListQueryResponseBodyModuleListCarRule) String() string {
	return dara.Prettify(s)
}

func (s ApplyListQueryResponseBodyModuleListCarRule) GoString() string {
	return s.String()
}

func (s *ApplyListQueryResponseBodyModuleListCarRule) GetScenarioTemplateId() *string {
	return s.ScenarioTemplateId
}

func (s *ApplyListQueryResponseBodyModuleListCarRule) GetScenarioTemplateName() *string {
	return s.ScenarioTemplateName
}

func (s *ApplyListQueryResponseBodyModuleListCarRule) SetScenarioTemplateId(v string) *ApplyListQueryResponseBodyModuleListCarRule {
	s.ScenarioTemplateId = &v
	return s
}

func (s *ApplyListQueryResponseBodyModuleListCarRule) SetScenarioTemplateName(v string) *ApplyListQueryResponseBodyModuleListCarRule {
	s.ScenarioTemplateName = &v
	return s
}

func (s *ApplyListQueryResponseBodyModuleListCarRule) Validate() error {
	return dara.Validate(s)
}

type ApplyListQueryResponseBodyModuleListExternalTravelerList struct {
	Attribute             *string `json:"attribute,omitempty" xml:"attribute,omitempty"`
	CostCenterName        *string `json:"cost_center_name,omitempty" xml:"cost_center_name,omitempty"`
	DepartId              *string `json:"depart_id,omitempty" xml:"depart_id,omitempty"`
	ExternalUserId        *string `json:"external_user_id,omitempty" xml:"external_user_id,omitempty"`
	InvoiceName           *string `json:"invoice_name,omitempty" xml:"invoice_name,omitempty"`
	PaymentDepartmentName *string `json:"payment_department_name,omitempty" xml:"payment_department_name,omitempty"`
	ProjectCode           *string `json:"project_code,omitempty" xml:"project_code,omitempty"`
	ProjectTitle          *string `json:"project_title,omitempty" xml:"project_title,omitempty"`
	ThirdpartDepartId     *string `json:"thirdpart_depart_id,omitempty" xml:"thirdpart_depart_id,omitempty"`
	UserName              *string `json:"user_name,omitempty" xml:"user_name,omitempty"`
}

func (s ApplyListQueryResponseBodyModuleListExternalTravelerList) String() string {
	return dara.Prettify(s)
}

func (s ApplyListQueryResponseBodyModuleListExternalTravelerList) GoString() string {
	return s.String()
}

func (s *ApplyListQueryResponseBodyModuleListExternalTravelerList) GetAttribute() *string {
	return s.Attribute
}

func (s *ApplyListQueryResponseBodyModuleListExternalTravelerList) GetCostCenterName() *string {
	return s.CostCenterName
}

func (s *ApplyListQueryResponseBodyModuleListExternalTravelerList) GetDepartId() *string {
	return s.DepartId
}

func (s *ApplyListQueryResponseBodyModuleListExternalTravelerList) GetExternalUserId() *string {
	return s.ExternalUserId
}

func (s *ApplyListQueryResponseBodyModuleListExternalTravelerList) GetInvoiceName() *string {
	return s.InvoiceName
}

func (s *ApplyListQueryResponseBodyModuleListExternalTravelerList) GetPaymentDepartmentName() *string {
	return s.PaymentDepartmentName
}

func (s *ApplyListQueryResponseBodyModuleListExternalTravelerList) GetProjectCode() *string {
	return s.ProjectCode
}

func (s *ApplyListQueryResponseBodyModuleListExternalTravelerList) GetProjectTitle() *string {
	return s.ProjectTitle
}

func (s *ApplyListQueryResponseBodyModuleListExternalTravelerList) GetThirdpartDepartId() *string {
	return s.ThirdpartDepartId
}

func (s *ApplyListQueryResponseBodyModuleListExternalTravelerList) GetUserName() *string {
	return s.UserName
}

func (s *ApplyListQueryResponseBodyModuleListExternalTravelerList) SetAttribute(v string) *ApplyListQueryResponseBodyModuleListExternalTravelerList {
	s.Attribute = &v
	return s
}

func (s *ApplyListQueryResponseBodyModuleListExternalTravelerList) SetCostCenterName(v string) *ApplyListQueryResponseBodyModuleListExternalTravelerList {
	s.CostCenterName = &v
	return s
}

func (s *ApplyListQueryResponseBodyModuleListExternalTravelerList) SetDepartId(v string) *ApplyListQueryResponseBodyModuleListExternalTravelerList {
	s.DepartId = &v
	return s
}

func (s *ApplyListQueryResponseBodyModuleListExternalTravelerList) SetExternalUserId(v string) *ApplyListQueryResponseBodyModuleListExternalTravelerList {
	s.ExternalUserId = &v
	return s
}

func (s *ApplyListQueryResponseBodyModuleListExternalTravelerList) SetInvoiceName(v string) *ApplyListQueryResponseBodyModuleListExternalTravelerList {
	s.InvoiceName = &v
	return s
}

func (s *ApplyListQueryResponseBodyModuleListExternalTravelerList) SetPaymentDepartmentName(v string) *ApplyListQueryResponseBodyModuleListExternalTravelerList {
	s.PaymentDepartmentName = &v
	return s
}

func (s *ApplyListQueryResponseBodyModuleListExternalTravelerList) SetProjectCode(v string) *ApplyListQueryResponseBodyModuleListExternalTravelerList {
	s.ProjectCode = &v
	return s
}

func (s *ApplyListQueryResponseBodyModuleListExternalTravelerList) SetProjectTitle(v string) *ApplyListQueryResponseBodyModuleListExternalTravelerList {
	s.ProjectTitle = &v
	return s
}

func (s *ApplyListQueryResponseBodyModuleListExternalTravelerList) SetThirdpartDepartId(v string) *ApplyListQueryResponseBodyModuleListExternalTravelerList {
	s.ThirdpartDepartId = &v
	return s
}

func (s *ApplyListQueryResponseBodyModuleListExternalTravelerList) SetUserName(v string) *ApplyListQueryResponseBodyModuleListExternalTravelerList {
	s.UserName = &v
	return s
}

func (s *ApplyListQueryResponseBodyModuleListExternalTravelerList) Validate() error {
	return dara.Validate(s)
}

type ApplyListQueryResponseBodyModuleListItineraryList struct {
	ArrCity *string `json:"arr_city,omitempty" xml:"arr_city,omitempty"`
	// example:
	//
	// 2018-09-19T14:03Z
	ArrDate        *string `json:"arr_date,omitempty" xml:"arr_date,omitempty"`
	CostCenterName *string `json:"cost_center_name,omitempty" xml:"cost_center_name,omitempty"`
	DepCity        *string `json:"dep_city,omitempty" xml:"dep_city,omitempty"`
	// example:
	//
	// 2018-09-19T14:03Z
	DepDate     *string `json:"dep_date,omitempty" xml:"dep_date,omitempty"`
	InvoiceName *string `json:"invoice_name,omitempty" xml:"invoice_name,omitempty"`
	// example:
	//
	// abcd
	ItineraryId *string `json:"itinerary_id,omitempty" xml:"itinerary_id,omitempty"`
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
	// 0
	TrafficType *int32 `json:"traffic_type,omitempty" xml:"traffic_type,omitempty"`
	// example:
	//
	// 1
	TripWay *int32 `json:"trip_way,omitempty" xml:"trip_way,omitempty"`
}

func (s ApplyListQueryResponseBodyModuleListItineraryList) String() string {
	return dara.Prettify(s)
}

func (s ApplyListQueryResponseBodyModuleListItineraryList) GoString() string {
	return s.String()
}

func (s *ApplyListQueryResponseBodyModuleListItineraryList) GetArrCity() *string {
	return s.ArrCity
}

func (s *ApplyListQueryResponseBodyModuleListItineraryList) GetArrDate() *string {
	return s.ArrDate
}

func (s *ApplyListQueryResponseBodyModuleListItineraryList) GetCostCenterName() *string {
	return s.CostCenterName
}

func (s *ApplyListQueryResponseBodyModuleListItineraryList) GetDepCity() *string {
	return s.DepCity
}

func (s *ApplyListQueryResponseBodyModuleListItineraryList) GetDepDate() *string {
	return s.DepDate
}

func (s *ApplyListQueryResponseBodyModuleListItineraryList) GetInvoiceName() *string {
	return s.InvoiceName
}

func (s *ApplyListQueryResponseBodyModuleListItineraryList) GetItineraryId() *string {
	return s.ItineraryId
}

func (s *ApplyListQueryResponseBodyModuleListItineraryList) GetProjectCode() *string {
	return s.ProjectCode
}

func (s *ApplyListQueryResponseBodyModuleListItineraryList) GetProjectTitle() *string {
	return s.ProjectTitle
}

func (s *ApplyListQueryResponseBodyModuleListItineraryList) GetThirdpartCostCenterId() *string {
	return s.ThirdpartCostCenterId
}

func (s *ApplyListQueryResponseBodyModuleListItineraryList) GetThirdpartInvoiceId() *string {
	return s.ThirdpartInvoiceId
}

func (s *ApplyListQueryResponseBodyModuleListItineraryList) GetThirdpartItineraryId() *string {
	return s.ThirdpartItineraryId
}

func (s *ApplyListQueryResponseBodyModuleListItineraryList) GetTrafficType() *int32 {
	return s.TrafficType
}

func (s *ApplyListQueryResponseBodyModuleListItineraryList) GetTripWay() *int32 {
	return s.TripWay
}

func (s *ApplyListQueryResponseBodyModuleListItineraryList) SetArrCity(v string) *ApplyListQueryResponseBodyModuleListItineraryList {
	s.ArrCity = &v
	return s
}

func (s *ApplyListQueryResponseBodyModuleListItineraryList) SetArrDate(v string) *ApplyListQueryResponseBodyModuleListItineraryList {
	s.ArrDate = &v
	return s
}

func (s *ApplyListQueryResponseBodyModuleListItineraryList) SetCostCenterName(v string) *ApplyListQueryResponseBodyModuleListItineraryList {
	s.CostCenterName = &v
	return s
}

func (s *ApplyListQueryResponseBodyModuleListItineraryList) SetDepCity(v string) *ApplyListQueryResponseBodyModuleListItineraryList {
	s.DepCity = &v
	return s
}

func (s *ApplyListQueryResponseBodyModuleListItineraryList) SetDepDate(v string) *ApplyListQueryResponseBodyModuleListItineraryList {
	s.DepDate = &v
	return s
}

func (s *ApplyListQueryResponseBodyModuleListItineraryList) SetInvoiceName(v string) *ApplyListQueryResponseBodyModuleListItineraryList {
	s.InvoiceName = &v
	return s
}

func (s *ApplyListQueryResponseBodyModuleListItineraryList) SetItineraryId(v string) *ApplyListQueryResponseBodyModuleListItineraryList {
	s.ItineraryId = &v
	return s
}

func (s *ApplyListQueryResponseBodyModuleListItineraryList) SetProjectCode(v string) *ApplyListQueryResponseBodyModuleListItineraryList {
	s.ProjectCode = &v
	return s
}

func (s *ApplyListQueryResponseBodyModuleListItineraryList) SetProjectTitle(v string) *ApplyListQueryResponseBodyModuleListItineraryList {
	s.ProjectTitle = &v
	return s
}

func (s *ApplyListQueryResponseBodyModuleListItineraryList) SetThirdpartCostCenterId(v string) *ApplyListQueryResponseBodyModuleListItineraryList {
	s.ThirdpartCostCenterId = &v
	return s
}

func (s *ApplyListQueryResponseBodyModuleListItineraryList) SetThirdpartInvoiceId(v string) *ApplyListQueryResponseBodyModuleListItineraryList {
	s.ThirdpartInvoiceId = &v
	return s
}

func (s *ApplyListQueryResponseBodyModuleListItineraryList) SetThirdpartItineraryId(v string) *ApplyListQueryResponseBodyModuleListItineraryList {
	s.ThirdpartItineraryId = &v
	return s
}

func (s *ApplyListQueryResponseBodyModuleListItineraryList) SetTrafficType(v int32) *ApplyListQueryResponseBodyModuleListItineraryList {
	s.TrafficType = &v
	return s
}

func (s *ApplyListQueryResponseBodyModuleListItineraryList) SetTripWay(v int32) *ApplyListQueryResponseBodyModuleListItineraryList {
	s.TripWay = &v
	return s
}

func (s *ApplyListQueryResponseBodyModuleListItineraryList) Validate() error {
	return dara.Validate(s)
}

type ApplyListQueryResponseBodyModuleListItinerarySetList struct {
	// example:
	//
	// 2018-09-19T14:03Z
	ArrDate *string `json:"arr_date,omitempty" xml:"arr_date,omitempty"`
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
	// abcd
	ItineraryId *string `json:"itinerary_id,omitempty" xml:"itinerary_id,omitempty"`
	// example:
	//
	// 12345
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

func (s ApplyListQueryResponseBodyModuleListItinerarySetList) String() string {
	return dara.Prettify(s)
}

func (s ApplyListQueryResponseBodyModuleListItinerarySetList) GoString() string {
	return s.String()
}

func (s *ApplyListQueryResponseBodyModuleListItinerarySetList) GetArrDate() *string {
	return s.ArrDate
}

func (s *ApplyListQueryResponseBodyModuleListItinerarySetList) GetCityCodeSet() *string {
	return s.CityCodeSet
}

func (s *ApplyListQueryResponseBodyModuleListItinerarySetList) GetCitySet() *string {
	return s.CitySet
}

func (s *ApplyListQueryResponseBodyModuleListItinerarySetList) GetCostCenterName() *string {
	return s.CostCenterName
}

func (s *ApplyListQueryResponseBodyModuleListItinerarySetList) GetDepDate() *string {
	return s.DepDate
}

func (s *ApplyListQueryResponseBodyModuleListItinerarySetList) GetInvoiceName() *string {
	return s.InvoiceName
}

func (s *ApplyListQueryResponseBodyModuleListItinerarySetList) GetItineraryId() *string {
	return s.ItineraryId
}

func (s *ApplyListQueryResponseBodyModuleListItinerarySetList) GetProjectCode() *string {
	return s.ProjectCode
}

func (s *ApplyListQueryResponseBodyModuleListItinerarySetList) GetProjectTitle() *string {
	return s.ProjectTitle
}

func (s *ApplyListQueryResponseBodyModuleListItinerarySetList) GetThirdpartCostCenterId() *string {
	return s.ThirdpartCostCenterId
}

func (s *ApplyListQueryResponseBodyModuleListItinerarySetList) GetThirdpartInvoiceId() *string {
	return s.ThirdpartInvoiceId
}

func (s *ApplyListQueryResponseBodyModuleListItinerarySetList) GetThirdpartItineraryId() *string {
	return s.ThirdpartItineraryId
}

func (s *ApplyListQueryResponseBodyModuleListItinerarySetList) GetTrafficType() *int32 {
	return s.TrafficType
}

func (s *ApplyListQueryResponseBodyModuleListItinerarySetList) SetArrDate(v string) *ApplyListQueryResponseBodyModuleListItinerarySetList {
	s.ArrDate = &v
	return s
}

func (s *ApplyListQueryResponseBodyModuleListItinerarySetList) SetCityCodeSet(v string) *ApplyListQueryResponseBodyModuleListItinerarySetList {
	s.CityCodeSet = &v
	return s
}

func (s *ApplyListQueryResponseBodyModuleListItinerarySetList) SetCitySet(v string) *ApplyListQueryResponseBodyModuleListItinerarySetList {
	s.CitySet = &v
	return s
}

func (s *ApplyListQueryResponseBodyModuleListItinerarySetList) SetCostCenterName(v string) *ApplyListQueryResponseBodyModuleListItinerarySetList {
	s.CostCenterName = &v
	return s
}

func (s *ApplyListQueryResponseBodyModuleListItinerarySetList) SetDepDate(v string) *ApplyListQueryResponseBodyModuleListItinerarySetList {
	s.DepDate = &v
	return s
}

func (s *ApplyListQueryResponseBodyModuleListItinerarySetList) SetInvoiceName(v string) *ApplyListQueryResponseBodyModuleListItinerarySetList {
	s.InvoiceName = &v
	return s
}

func (s *ApplyListQueryResponseBodyModuleListItinerarySetList) SetItineraryId(v string) *ApplyListQueryResponseBodyModuleListItinerarySetList {
	s.ItineraryId = &v
	return s
}

func (s *ApplyListQueryResponseBodyModuleListItinerarySetList) SetProjectCode(v string) *ApplyListQueryResponseBodyModuleListItinerarySetList {
	s.ProjectCode = &v
	return s
}

func (s *ApplyListQueryResponseBodyModuleListItinerarySetList) SetProjectTitle(v string) *ApplyListQueryResponseBodyModuleListItinerarySetList {
	s.ProjectTitle = &v
	return s
}

func (s *ApplyListQueryResponseBodyModuleListItinerarySetList) SetThirdpartCostCenterId(v string) *ApplyListQueryResponseBodyModuleListItinerarySetList {
	s.ThirdpartCostCenterId = &v
	return s
}

func (s *ApplyListQueryResponseBodyModuleListItinerarySetList) SetThirdpartInvoiceId(v string) *ApplyListQueryResponseBodyModuleListItinerarySetList {
	s.ThirdpartInvoiceId = &v
	return s
}

func (s *ApplyListQueryResponseBodyModuleListItinerarySetList) SetThirdpartItineraryId(v string) *ApplyListQueryResponseBodyModuleListItinerarySetList {
	s.ThirdpartItineraryId = &v
	return s
}

func (s *ApplyListQueryResponseBodyModuleListItinerarySetList) SetTrafficType(v int32) *ApplyListQueryResponseBodyModuleListItinerarySetList {
	s.TrafficType = &v
	return s
}

func (s *ApplyListQueryResponseBodyModuleListItinerarySetList) Validate() error {
	return dara.Validate(s)
}

type ApplyListQueryResponseBodyModuleListTravelerList struct {
	Attribute             *string                                                       `json:"attribute,omitempty" xml:"attribute,omitempty"`
	CarCitySet            []*ApplyListQueryResponseBodyModuleListTravelerListCarCitySet `json:"car_city_set,omitempty" xml:"car_city_set,omitempty" type:"Repeated"`
	CostCenterName        *string                                                       `json:"cost_center_name,omitempty" xml:"cost_center_name,omitempty"`
	DepartId              *string                                                       `json:"depart_id,omitempty" xml:"depart_id,omitempty"`
	InvoiceName           *string                                                       `json:"invoice_name,omitempty" xml:"invoice_name,omitempty"`
	JobNo                 *string                                                       `json:"job_no,omitempty" xml:"job_no,omitempty"`
	PaymentDepartmentName *string                                                       `json:"payment_department_name,omitempty" xml:"payment_department_name,omitempty"`
	ProjectCode           *string                                                       `json:"project_code,omitempty" xml:"project_code,omitempty"`
	ProjectTitle          *string                                                       `json:"project_title,omitempty" xml:"project_title,omitempty"`
	ThirdpartDepartId     *string                                                       `json:"thirdpart_depart_id,omitempty" xml:"thirdpart_depart_id,omitempty"`
	// example:
	//
	// user1
	UserId   *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	UserName *string `json:"user_name,omitempty" xml:"user_name,omitempty"`
}

func (s ApplyListQueryResponseBodyModuleListTravelerList) String() string {
	return dara.Prettify(s)
}

func (s ApplyListQueryResponseBodyModuleListTravelerList) GoString() string {
	return s.String()
}

func (s *ApplyListQueryResponseBodyModuleListTravelerList) GetAttribute() *string {
	return s.Attribute
}

func (s *ApplyListQueryResponseBodyModuleListTravelerList) GetCarCitySet() []*ApplyListQueryResponseBodyModuleListTravelerListCarCitySet {
	return s.CarCitySet
}

func (s *ApplyListQueryResponseBodyModuleListTravelerList) GetCostCenterName() *string {
	return s.CostCenterName
}

func (s *ApplyListQueryResponseBodyModuleListTravelerList) GetDepartId() *string {
	return s.DepartId
}

func (s *ApplyListQueryResponseBodyModuleListTravelerList) GetInvoiceName() *string {
	return s.InvoiceName
}

func (s *ApplyListQueryResponseBodyModuleListTravelerList) GetJobNo() *string {
	return s.JobNo
}

func (s *ApplyListQueryResponseBodyModuleListTravelerList) GetPaymentDepartmentName() *string {
	return s.PaymentDepartmentName
}

func (s *ApplyListQueryResponseBodyModuleListTravelerList) GetProjectCode() *string {
	return s.ProjectCode
}

func (s *ApplyListQueryResponseBodyModuleListTravelerList) GetProjectTitle() *string {
	return s.ProjectTitle
}

func (s *ApplyListQueryResponseBodyModuleListTravelerList) GetThirdpartDepartId() *string {
	return s.ThirdpartDepartId
}

func (s *ApplyListQueryResponseBodyModuleListTravelerList) GetUserId() *string {
	return s.UserId
}

func (s *ApplyListQueryResponseBodyModuleListTravelerList) GetUserName() *string {
	return s.UserName
}

func (s *ApplyListQueryResponseBodyModuleListTravelerList) SetAttribute(v string) *ApplyListQueryResponseBodyModuleListTravelerList {
	s.Attribute = &v
	return s
}

func (s *ApplyListQueryResponseBodyModuleListTravelerList) SetCarCitySet(v []*ApplyListQueryResponseBodyModuleListTravelerListCarCitySet) *ApplyListQueryResponseBodyModuleListTravelerList {
	s.CarCitySet = v
	return s
}

func (s *ApplyListQueryResponseBodyModuleListTravelerList) SetCostCenterName(v string) *ApplyListQueryResponseBodyModuleListTravelerList {
	s.CostCenterName = &v
	return s
}

func (s *ApplyListQueryResponseBodyModuleListTravelerList) SetDepartId(v string) *ApplyListQueryResponseBodyModuleListTravelerList {
	s.DepartId = &v
	return s
}

func (s *ApplyListQueryResponseBodyModuleListTravelerList) SetInvoiceName(v string) *ApplyListQueryResponseBodyModuleListTravelerList {
	s.InvoiceName = &v
	return s
}

func (s *ApplyListQueryResponseBodyModuleListTravelerList) SetJobNo(v string) *ApplyListQueryResponseBodyModuleListTravelerList {
	s.JobNo = &v
	return s
}

func (s *ApplyListQueryResponseBodyModuleListTravelerList) SetPaymentDepartmentName(v string) *ApplyListQueryResponseBodyModuleListTravelerList {
	s.PaymentDepartmentName = &v
	return s
}

func (s *ApplyListQueryResponseBodyModuleListTravelerList) SetProjectCode(v string) *ApplyListQueryResponseBodyModuleListTravelerList {
	s.ProjectCode = &v
	return s
}

func (s *ApplyListQueryResponseBodyModuleListTravelerList) SetProjectTitle(v string) *ApplyListQueryResponseBodyModuleListTravelerList {
	s.ProjectTitle = &v
	return s
}

func (s *ApplyListQueryResponseBodyModuleListTravelerList) SetThirdpartDepartId(v string) *ApplyListQueryResponseBodyModuleListTravelerList {
	s.ThirdpartDepartId = &v
	return s
}

func (s *ApplyListQueryResponseBodyModuleListTravelerList) SetUserId(v string) *ApplyListQueryResponseBodyModuleListTravelerList {
	s.UserId = &v
	return s
}

func (s *ApplyListQueryResponseBodyModuleListTravelerList) SetUserName(v string) *ApplyListQueryResponseBodyModuleListTravelerList {
	s.UserName = &v
	return s
}

func (s *ApplyListQueryResponseBodyModuleListTravelerList) Validate() error {
	return dara.Validate(s)
}

type ApplyListQueryResponseBodyModuleListTravelerListCarCitySet struct {
	CityCode *string `json:"city_code,omitempty" xml:"city_code,omitempty"`
	CityName *string `json:"city_name,omitempty" xml:"city_name,omitempty"`
}

func (s ApplyListQueryResponseBodyModuleListTravelerListCarCitySet) String() string {
	return dara.Prettify(s)
}

func (s ApplyListQueryResponseBodyModuleListTravelerListCarCitySet) GoString() string {
	return s.String()
}

func (s *ApplyListQueryResponseBodyModuleListTravelerListCarCitySet) GetCityCode() *string {
	return s.CityCode
}

func (s *ApplyListQueryResponseBodyModuleListTravelerListCarCitySet) GetCityName() *string {
	return s.CityName
}

func (s *ApplyListQueryResponseBodyModuleListTravelerListCarCitySet) SetCityCode(v string) *ApplyListQueryResponseBodyModuleListTravelerListCarCitySet {
	s.CityCode = &v
	return s
}

func (s *ApplyListQueryResponseBodyModuleListTravelerListCarCitySet) SetCityName(v string) *ApplyListQueryResponseBodyModuleListTravelerListCarCitySet {
	s.CityName = &v
	return s
}

func (s *ApplyListQueryResponseBodyModuleListTravelerListCarCitySet) Validate() error {
	return dara.Validate(s)
}
