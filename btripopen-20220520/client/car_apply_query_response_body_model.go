// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCarApplyQueryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplyList(v []*CarApplyQueryResponseBodyApplyList) *CarApplyQueryResponseBody
	GetApplyList() []*CarApplyQueryResponseBodyApplyList
	SetCode(v string) *CarApplyQueryResponseBody
	GetCode() *string
	SetMessage(v string) *CarApplyQueryResponseBody
	GetMessage() *string
	SetRequestId(v string) *CarApplyQueryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CarApplyQueryResponseBody
	GetSuccess() *bool
	SetTotal(v int32) *CarApplyQueryResponseBody
	GetTotal() *int32
	SetTraceId(v string) *CarApplyQueryResponseBody
	GetTraceId() *string
}

type CarApplyQueryResponseBody struct {
	ApplyList []*CarApplyQueryResponseBodyApplyList `json:"apply_list,omitempty" xml:"apply_list,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	Code    *string `json:"code,omitempty" xml:"code,omitempty"`
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 407543AF-2BD9-5890-BD92-9D1AB7218B27
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 6
	Total *int32 `json:"total,omitempty" xml:"total,omitempty"`
	// example:
	//
	// 210bcc3a16583004579056128d33d7
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s CarApplyQueryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CarApplyQueryResponseBody) GoString() string {
	return s.String()
}

func (s *CarApplyQueryResponseBody) GetApplyList() []*CarApplyQueryResponseBodyApplyList {
	return s.ApplyList
}

func (s *CarApplyQueryResponseBody) GetCode() *string {
	return s.Code
}

func (s *CarApplyQueryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CarApplyQueryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CarApplyQueryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CarApplyQueryResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *CarApplyQueryResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *CarApplyQueryResponseBody) SetApplyList(v []*CarApplyQueryResponseBodyApplyList) *CarApplyQueryResponseBody {
	s.ApplyList = v
	return s
}

func (s *CarApplyQueryResponseBody) SetCode(v string) *CarApplyQueryResponseBody {
	s.Code = &v
	return s
}

func (s *CarApplyQueryResponseBody) SetMessage(v string) *CarApplyQueryResponseBody {
	s.Message = &v
	return s
}

func (s *CarApplyQueryResponseBody) SetRequestId(v string) *CarApplyQueryResponseBody {
	s.RequestId = &v
	return s
}

func (s *CarApplyQueryResponseBody) SetSuccess(v bool) *CarApplyQueryResponseBody {
	s.Success = &v
	return s
}

func (s *CarApplyQueryResponseBody) SetTotal(v int32) *CarApplyQueryResponseBody {
	s.Total = &v
	return s
}

func (s *CarApplyQueryResponseBody) SetTraceId(v string) *CarApplyQueryResponseBody {
	s.TraceId = &v
	return s
}

func (s *CarApplyQueryResponseBody) Validate() error {
	return dara.Validate(s)
}

type CarApplyQueryResponseBodyApplyList struct {
	ApproverList []*CarApplyQueryResponseBodyApplyListApproverList `json:"approver_list,omitempty" xml:"approver_list,omitempty" type:"Repeated"`
	BusinessType *string                                           `json:"business_type,omitempty" xml:"business_type,omitempty"`
	// example:
	//
	// depart1
	DepartId   *string `json:"depart_id,omitempty" xml:"depart_id,omitempty"`
	DepartName *string `json:"depart_name,omitempty" xml:"depart_name,omitempty"`
	// example:
	//
	// 2021-03-18T20:26Z
	GmtCreate *string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// example:
	//
	// 2021-03-18T20:26Z
	GmtModified         *string                                            `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	ItineraryList       []*CarApplyQueryResponseBodyApplyListItineraryList `json:"itinerary_list,omitempty" xml:"itinerary_list,omitempty" type:"Repeated"`
	OrderId             *int64                                             `json:"order_id,omitempty" xml:"order_id,omitempty"`
	RelatedThirdApplyId *string                                            `json:"related_third_apply_id,omitempty" xml:"related_third_apply_id,omitempty"`
	// example:
	//
	// 2
	Status     *int32  `json:"status,omitempty" xml:"status,omitempty"`
	StatusDesc *string `json:"status_desc,omitempty" xml:"status_desc,omitempty"`
	// example:
	//
	// 1
	ThirdpartId      *string                                               `json:"thirdpart_id,omitempty" xml:"thirdpart_id,omitempty"`
	TravelerStandard []*CarApplyQueryResponseBodyApplyListTravelerStandard `json:"traveler_standard,omitempty" xml:"traveler_standard,omitempty" type:"Repeated"`
	TripCause        *string                                               `json:"trip_cause,omitempty" xml:"trip_cause,omitempty"`
	TripTitle        *string                                               `json:"trip_title,omitempty" xml:"trip_title,omitempty"`
	// example:
	//
	// user1
	UserId   *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	UserName *string `json:"user_name,omitempty" xml:"user_name,omitempty"`
}

func (s CarApplyQueryResponseBodyApplyList) String() string {
	return dara.Prettify(s)
}

func (s CarApplyQueryResponseBodyApplyList) GoString() string {
	return s.String()
}

func (s *CarApplyQueryResponseBodyApplyList) GetApproverList() []*CarApplyQueryResponseBodyApplyListApproverList {
	return s.ApproverList
}

func (s *CarApplyQueryResponseBodyApplyList) GetBusinessType() *string {
	return s.BusinessType
}

func (s *CarApplyQueryResponseBodyApplyList) GetDepartId() *string {
	return s.DepartId
}

func (s *CarApplyQueryResponseBodyApplyList) GetDepartName() *string {
	return s.DepartName
}

func (s *CarApplyQueryResponseBodyApplyList) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *CarApplyQueryResponseBodyApplyList) GetGmtModified() *string {
	return s.GmtModified
}

func (s *CarApplyQueryResponseBodyApplyList) GetItineraryList() []*CarApplyQueryResponseBodyApplyListItineraryList {
	return s.ItineraryList
}

func (s *CarApplyQueryResponseBodyApplyList) GetOrderId() *int64 {
	return s.OrderId
}

func (s *CarApplyQueryResponseBodyApplyList) GetRelatedThirdApplyId() *string {
	return s.RelatedThirdApplyId
}

func (s *CarApplyQueryResponseBodyApplyList) GetStatus() *int32 {
	return s.Status
}

func (s *CarApplyQueryResponseBodyApplyList) GetStatusDesc() *string {
	return s.StatusDesc
}

func (s *CarApplyQueryResponseBodyApplyList) GetThirdpartId() *string {
	return s.ThirdpartId
}

func (s *CarApplyQueryResponseBodyApplyList) GetTravelerStandard() []*CarApplyQueryResponseBodyApplyListTravelerStandard {
	return s.TravelerStandard
}

func (s *CarApplyQueryResponseBodyApplyList) GetTripCause() *string {
	return s.TripCause
}

func (s *CarApplyQueryResponseBodyApplyList) GetTripTitle() *string {
	return s.TripTitle
}

func (s *CarApplyQueryResponseBodyApplyList) GetUserId() *string {
	return s.UserId
}

func (s *CarApplyQueryResponseBodyApplyList) GetUserName() *string {
	return s.UserName
}

func (s *CarApplyQueryResponseBodyApplyList) SetApproverList(v []*CarApplyQueryResponseBodyApplyListApproverList) *CarApplyQueryResponseBodyApplyList {
	s.ApproverList = v
	return s
}

func (s *CarApplyQueryResponseBodyApplyList) SetBusinessType(v string) *CarApplyQueryResponseBodyApplyList {
	s.BusinessType = &v
	return s
}

func (s *CarApplyQueryResponseBodyApplyList) SetDepartId(v string) *CarApplyQueryResponseBodyApplyList {
	s.DepartId = &v
	return s
}

func (s *CarApplyQueryResponseBodyApplyList) SetDepartName(v string) *CarApplyQueryResponseBodyApplyList {
	s.DepartName = &v
	return s
}

func (s *CarApplyQueryResponseBodyApplyList) SetGmtCreate(v string) *CarApplyQueryResponseBodyApplyList {
	s.GmtCreate = &v
	return s
}

func (s *CarApplyQueryResponseBodyApplyList) SetGmtModified(v string) *CarApplyQueryResponseBodyApplyList {
	s.GmtModified = &v
	return s
}

func (s *CarApplyQueryResponseBodyApplyList) SetItineraryList(v []*CarApplyQueryResponseBodyApplyListItineraryList) *CarApplyQueryResponseBodyApplyList {
	s.ItineraryList = v
	return s
}

func (s *CarApplyQueryResponseBodyApplyList) SetOrderId(v int64) *CarApplyQueryResponseBodyApplyList {
	s.OrderId = &v
	return s
}

func (s *CarApplyQueryResponseBodyApplyList) SetRelatedThirdApplyId(v string) *CarApplyQueryResponseBodyApplyList {
	s.RelatedThirdApplyId = &v
	return s
}

func (s *CarApplyQueryResponseBodyApplyList) SetStatus(v int32) *CarApplyQueryResponseBodyApplyList {
	s.Status = &v
	return s
}

func (s *CarApplyQueryResponseBodyApplyList) SetStatusDesc(v string) *CarApplyQueryResponseBodyApplyList {
	s.StatusDesc = &v
	return s
}

func (s *CarApplyQueryResponseBodyApplyList) SetThirdpartId(v string) *CarApplyQueryResponseBodyApplyList {
	s.ThirdpartId = &v
	return s
}

func (s *CarApplyQueryResponseBodyApplyList) SetTravelerStandard(v []*CarApplyQueryResponseBodyApplyListTravelerStandard) *CarApplyQueryResponseBodyApplyList {
	s.TravelerStandard = v
	return s
}

func (s *CarApplyQueryResponseBodyApplyList) SetTripCause(v string) *CarApplyQueryResponseBodyApplyList {
	s.TripCause = &v
	return s
}

func (s *CarApplyQueryResponseBodyApplyList) SetTripTitle(v string) *CarApplyQueryResponseBodyApplyList {
	s.TripTitle = &v
	return s
}

func (s *CarApplyQueryResponseBodyApplyList) SetUserId(v string) *CarApplyQueryResponseBodyApplyList {
	s.UserId = &v
	return s
}

func (s *CarApplyQueryResponseBodyApplyList) SetUserName(v string) *CarApplyQueryResponseBodyApplyList {
	s.UserName = &v
	return s
}

func (s *CarApplyQueryResponseBodyApplyList) Validate() error {
	return dara.Validate(s)
}

type CarApplyQueryResponseBodyApplyListApproverList struct {
	Note *string `json:"note,omitempty" xml:"note,omitempty"`
	// example:
	//
	// 2021-03-18T20:26Z
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

func (s CarApplyQueryResponseBodyApplyListApproverList) String() string {
	return dara.Prettify(s)
}

func (s CarApplyQueryResponseBodyApplyListApproverList) GoString() string {
	return s.String()
}

func (s *CarApplyQueryResponseBodyApplyListApproverList) GetNote() *string {
	return s.Note
}

func (s *CarApplyQueryResponseBodyApplyListApproverList) GetOperateTime() *string {
	return s.OperateTime
}

func (s *CarApplyQueryResponseBodyApplyListApproverList) GetOrder() *int32 {
	return s.Order
}

func (s *CarApplyQueryResponseBodyApplyListApproverList) GetStatus() *int32 {
	return s.Status
}

func (s *CarApplyQueryResponseBodyApplyListApproverList) GetStatusDesc() *string {
	return s.StatusDesc
}

func (s *CarApplyQueryResponseBodyApplyListApproverList) GetUserId() *string {
	return s.UserId
}

func (s *CarApplyQueryResponseBodyApplyListApproverList) GetUserName() *string {
	return s.UserName
}

func (s *CarApplyQueryResponseBodyApplyListApproverList) SetNote(v string) *CarApplyQueryResponseBodyApplyListApproverList {
	s.Note = &v
	return s
}

func (s *CarApplyQueryResponseBodyApplyListApproverList) SetOperateTime(v string) *CarApplyQueryResponseBodyApplyListApproverList {
	s.OperateTime = &v
	return s
}

func (s *CarApplyQueryResponseBodyApplyListApproverList) SetOrder(v int32) *CarApplyQueryResponseBodyApplyListApproverList {
	s.Order = &v
	return s
}

func (s *CarApplyQueryResponseBodyApplyListApproverList) SetStatus(v int32) *CarApplyQueryResponseBodyApplyListApproverList {
	s.Status = &v
	return s
}

func (s *CarApplyQueryResponseBodyApplyListApproverList) SetStatusDesc(v string) *CarApplyQueryResponseBodyApplyListApproverList {
	s.StatusDesc = &v
	return s
}

func (s *CarApplyQueryResponseBodyApplyListApproverList) SetUserId(v string) *CarApplyQueryResponseBodyApplyListApproverList {
	s.UserId = &v
	return s
}

func (s *CarApplyQueryResponseBodyApplyListApproverList) SetUserName(v string) *CarApplyQueryResponseBodyApplyListApproverList {
	s.UserName = &v
	return s
}

func (s *CarApplyQueryResponseBodyApplyListApproverList) Validate() error {
	return dara.Validate(s)
}

type CarApplyQueryResponseBodyApplyListItineraryList struct {
	ArrCity *string `json:"arr_city,omitempty" xml:"arr_city,omitempty"`
	// example:
	//
	// HGH
	ArrCityCode *string `json:"arr_city_code,omitempty" xml:"arr_city_code,omitempty"`
	// example:
	//
	// 2021-03-18T20:26Z
	ArrDate *string `json:"arr_date,omitempty" xml:"arr_date,omitempty"`
	// example:
	//
	// 1
	CostCenterId   *int64  `json:"cost_center_id,omitempty" xml:"cost_center_id,omitempty"`
	CostCenterName *string `json:"cost_center_name,omitempty" xml:"cost_center_name,omitempty"`
	DepCity        *string `json:"dep_city,omitempty" xml:"dep_city,omitempty"`
	// example:
	//
	// HGH
	DepCityCode *string `json:"dep_city_code,omitempty" xml:"dep_city_code,omitempty"`
	// example:
	//
	// 2021-03-18T20:26Z
	DepDate *string `json:"dep_date,omitempty" xml:"dep_date,omitempty"`
	// example:
	//
	// 1
	InvoiceId   *int64  `json:"invoice_id,omitempty" xml:"invoice_id,omitempty"`
	InvoiceName *string `json:"invoice_name,omitempty" xml:"invoice_name,omitempty"`
	// example:
	//
	// 1
	ItineraryId *string `json:"itinerary_id,omitempty" xml:"itinerary_id,omitempty"`
	// example:
	//
	// xm1
	ProjectCode  *string `json:"project_code,omitempty" xml:"project_code,omitempty"`
	ProjectTitle *string `json:"project_title,omitempty" xml:"project_title,omitempty"`
	// example:
	//
	// 4
	TrafficType *int32 `json:"traffic_type,omitempty" xml:"traffic_type,omitempty"`
}

func (s CarApplyQueryResponseBodyApplyListItineraryList) String() string {
	return dara.Prettify(s)
}

func (s CarApplyQueryResponseBodyApplyListItineraryList) GoString() string {
	return s.String()
}

func (s *CarApplyQueryResponseBodyApplyListItineraryList) GetArrCity() *string {
	return s.ArrCity
}

func (s *CarApplyQueryResponseBodyApplyListItineraryList) GetArrCityCode() *string {
	return s.ArrCityCode
}

func (s *CarApplyQueryResponseBodyApplyListItineraryList) GetArrDate() *string {
	return s.ArrDate
}

func (s *CarApplyQueryResponseBodyApplyListItineraryList) GetCostCenterId() *int64 {
	return s.CostCenterId
}

func (s *CarApplyQueryResponseBodyApplyListItineraryList) GetCostCenterName() *string {
	return s.CostCenterName
}

func (s *CarApplyQueryResponseBodyApplyListItineraryList) GetDepCity() *string {
	return s.DepCity
}

func (s *CarApplyQueryResponseBodyApplyListItineraryList) GetDepCityCode() *string {
	return s.DepCityCode
}

func (s *CarApplyQueryResponseBodyApplyListItineraryList) GetDepDate() *string {
	return s.DepDate
}

func (s *CarApplyQueryResponseBodyApplyListItineraryList) GetInvoiceId() *int64 {
	return s.InvoiceId
}

func (s *CarApplyQueryResponseBodyApplyListItineraryList) GetInvoiceName() *string {
	return s.InvoiceName
}

func (s *CarApplyQueryResponseBodyApplyListItineraryList) GetItineraryId() *string {
	return s.ItineraryId
}

func (s *CarApplyQueryResponseBodyApplyListItineraryList) GetProjectCode() *string {
	return s.ProjectCode
}

func (s *CarApplyQueryResponseBodyApplyListItineraryList) GetProjectTitle() *string {
	return s.ProjectTitle
}

func (s *CarApplyQueryResponseBodyApplyListItineraryList) GetTrafficType() *int32 {
	return s.TrafficType
}

func (s *CarApplyQueryResponseBodyApplyListItineraryList) SetArrCity(v string) *CarApplyQueryResponseBodyApplyListItineraryList {
	s.ArrCity = &v
	return s
}

func (s *CarApplyQueryResponseBodyApplyListItineraryList) SetArrCityCode(v string) *CarApplyQueryResponseBodyApplyListItineraryList {
	s.ArrCityCode = &v
	return s
}

func (s *CarApplyQueryResponseBodyApplyListItineraryList) SetArrDate(v string) *CarApplyQueryResponseBodyApplyListItineraryList {
	s.ArrDate = &v
	return s
}

func (s *CarApplyQueryResponseBodyApplyListItineraryList) SetCostCenterId(v int64) *CarApplyQueryResponseBodyApplyListItineraryList {
	s.CostCenterId = &v
	return s
}

func (s *CarApplyQueryResponseBodyApplyListItineraryList) SetCostCenterName(v string) *CarApplyQueryResponseBodyApplyListItineraryList {
	s.CostCenterName = &v
	return s
}

func (s *CarApplyQueryResponseBodyApplyListItineraryList) SetDepCity(v string) *CarApplyQueryResponseBodyApplyListItineraryList {
	s.DepCity = &v
	return s
}

func (s *CarApplyQueryResponseBodyApplyListItineraryList) SetDepCityCode(v string) *CarApplyQueryResponseBodyApplyListItineraryList {
	s.DepCityCode = &v
	return s
}

func (s *CarApplyQueryResponseBodyApplyListItineraryList) SetDepDate(v string) *CarApplyQueryResponseBodyApplyListItineraryList {
	s.DepDate = &v
	return s
}

func (s *CarApplyQueryResponseBodyApplyListItineraryList) SetInvoiceId(v int64) *CarApplyQueryResponseBodyApplyListItineraryList {
	s.InvoiceId = &v
	return s
}

func (s *CarApplyQueryResponseBodyApplyListItineraryList) SetInvoiceName(v string) *CarApplyQueryResponseBodyApplyListItineraryList {
	s.InvoiceName = &v
	return s
}

func (s *CarApplyQueryResponseBodyApplyListItineraryList) SetItineraryId(v string) *CarApplyQueryResponseBodyApplyListItineraryList {
	s.ItineraryId = &v
	return s
}

func (s *CarApplyQueryResponseBodyApplyListItineraryList) SetProjectCode(v string) *CarApplyQueryResponseBodyApplyListItineraryList {
	s.ProjectCode = &v
	return s
}

func (s *CarApplyQueryResponseBodyApplyListItineraryList) SetProjectTitle(v string) *CarApplyQueryResponseBodyApplyListItineraryList {
	s.ProjectTitle = &v
	return s
}

func (s *CarApplyQueryResponseBodyApplyListItineraryList) SetTrafficType(v int32) *CarApplyQueryResponseBodyApplyListItineraryList {
	s.TrafficType = &v
	return s
}

func (s *CarApplyQueryResponseBodyApplyListItineraryList) Validate() error {
	return dara.Validate(s)
}

type CarApplyQueryResponseBodyApplyListTravelerStandard struct {
	CarCitySet []*CarApplyQueryResponseBodyApplyListTravelerStandardCarCitySet `json:"car_city_set,omitempty" xml:"car_city_set,omitempty" type:"Repeated"`
	UserId     *string                                                         `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

func (s CarApplyQueryResponseBodyApplyListTravelerStandard) String() string {
	return dara.Prettify(s)
}

func (s CarApplyQueryResponseBodyApplyListTravelerStandard) GoString() string {
	return s.String()
}

func (s *CarApplyQueryResponseBodyApplyListTravelerStandard) GetCarCitySet() []*CarApplyQueryResponseBodyApplyListTravelerStandardCarCitySet {
	return s.CarCitySet
}

func (s *CarApplyQueryResponseBodyApplyListTravelerStandard) GetUserId() *string {
	return s.UserId
}

func (s *CarApplyQueryResponseBodyApplyListTravelerStandard) SetCarCitySet(v []*CarApplyQueryResponseBodyApplyListTravelerStandardCarCitySet) *CarApplyQueryResponseBodyApplyListTravelerStandard {
	s.CarCitySet = v
	return s
}

func (s *CarApplyQueryResponseBodyApplyListTravelerStandard) SetUserId(v string) *CarApplyQueryResponseBodyApplyListTravelerStandard {
	s.UserId = &v
	return s
}

func (s *CarApplyQueryResponseBodyApplyListTravelerStandard) Validate() error {
	return dara.Validate(s)
}

type CarApplyQueryResponseBodyApplyListTravelerStandardCarCitySet struct {
	CityCode *string `json:"city_code,omitempty" xml:"city_code,omitempty"`
	CityName *string `json:"city_name,omitempty" xml:"city_name,omitempty"`
}

func (s CarApplyQueryResponseBodyApplyListTravelerStandardCarCitySet) String() string {
	return dara.Prettify(s)
}

func (s CarApplyQueryResponseBodyApplyListTravelerStandardCarCitySet) GoString() string {
	return s.String()
}

func (s *CarApplyQueryResponseBodyApplyListTravelerStandardCarCitySet) GetCityCode() *string {
	return s.CityCode
}

func (s *CarApplyQueryResponseBodyApplyListTravelerStandardCarCitySet) GetCityName() *string {
	return s.CityName
}

func (s *CarApplyQueryResponseBodyApplyListTravelerStandardCarCitySet) SetCityCode(v string) *CarApplyQueryResponseBodyApplyListTravelerStandardCarCitySet {
	s.CityCode = &v
	return s
}

func (s *CarApplyQueryResponseBodyApplyListTravelerStandardCarCitySet) SetCityName(v string) *CarApplyQueryResponseBodyApplyListTravelerStandardCarCitySet {
	s.CityName = &v
	return s
}

func (s *CarApplyQueryResponseBodyApplyListTravelerStandardCarCitySet) Validate() error {
	return dara.Validate(s)
}
