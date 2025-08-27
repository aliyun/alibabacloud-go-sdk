// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightOrderListQueryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *FlightOrderListQueryResponseBody
	GetCode() *string
	SetMessage(v string) *FlightOrderListQueryResponseBody
	GetMessage() *string
	SetModule(v []*FlightOrderListQueryResponseBodyModule) *FlightOrderListQueryResponseBody
	GetModule() []*FlightOrderListQueryResponseBodyModule
	SetPageInfo(v *FlightOrderListQueryResponseBodyPageInfo) *FlightOrderListQueryResponseBody
	GetPageInfo() *FlightOrderListQueryResponseBodyPageInfo
	SetRequestId(v string) *FlightOrderListQueryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *FlightOrderListQueryResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *FlightOrderListQueryResponseBody
	GetTraceId() *string
}

type FlightOrderListQueryResponseBody struct {
	// example:
	//
	// SUCCESS
	Code     *string                                   `json:"code,omitempty" xml:"code,omitempty"`
	Message  *string                                   `json:"message,omitempty" xml:"message,omitempty"`
	Module   []*FlightOrderListQueryResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Repeated"`
	PageInfo *FlightOrderListQueryResponseBodyPageInfo `json:"page_info,omitempty" xml:"page_info,omitempty" type:"Struct"`
	// example:
	//
	// A5009956-1077-52FB-B520-EA8C7E91D722
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

func (s FlightOrderListQueryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s FlightOrderListQueryResponseBody) GoString() string {
	return s.String()
}

func (s *FlightOrderListQueryResponseBody) GetCode() *string {
	return s.Code
}

func (s *FlightOrderListQueryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *FlightOrderListQueryResponseBody) GetModule() []*FlightOrderListQueryResponseBodyModule {
	return s.Module
}

func (s *FlightOrderListQueryResponseBody) GetPageInfo() *FlightOrderListQueryResponseBodyPageInfo {
	return s.PageInfo
}

func (s *FlightOrderListQueryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *FlightOrderListQueryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *FlightOrderListQueryResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *FlightOrderListQueryResponseBody) SetCode(v string) *FlightOrderListQueryResponseBody {
	s.Code = &v
	return s
}

func (s *FlightOrderListQueryResponseBody) SetMessage(v string) *FlightOrderListQueryResponseBody {
	s.Message = &v
	return s
}

func (s *FlightOrderListQueryResponseBody) SetModule(v []*FlightOrderListQueryResponseBodyModule) *FlightOrderListQueryResponseBody {
	s.Module = v
	return s
}

func (s *FlightOrderListQueryResponseBody) SetPageInfo(v *FlightOrderListQueryResponseBodyPageInfo) *FlightOrderListQueryResponseBody {
	s.PageInfo = v
	return s
}

func (s *FlightOrderListQueryResponseBody) SetRequestId(v string) *FlightOrderListQueryResponseBody {
	s.RequestId = &v
	return s
}

func (s *FlightOrderListQueryResponseBody) SetSuccess(v bool) *FlightOrderListQueryResponseBody {
	s.Success = &v
	return s
}

func (s *FlightOrderListQueryResponseBody) SetTraceId(v string) *FlightOrderListQueryResponseBody {
	s.TraceId = &v
	return s
}

func (s *FlightOrderListQueryResponseBody) Validate() error {
	return dara.Validate(s)
}

type FlightOrderListQueryResponseBodyModule struct {
	// example:
	//
	// 11774
	ApplyId       *int64  `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
	ArrAirport    *string `json:"arr_airport,omitempty" xml:"arr_airport,omitempty"`
	ArrCity       *string `json:"arr_city,omitempty" xml:"arr_city,omitempty"`
	ArrCityAdCode *string `json:"arr_city_ad_code,omitempty" xml:"arr_city_ad_code,omitempty"`
	BtripTitle    *string `json:"btrip_title,omitempty" xml:"btrip_title,omitempty"`
	// example:
	//
	// Y
	CabinClass    *string                                           `json:"cabin_class,omitempty" xml:"cabin_class,omitempty"`
	ContactName   *string                                           `json:"contact_name,omitempty" xml:"contact_name,omitempty"`
	CorpId        *string                                           `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
	CorpName      *string                                           `json:"corp_name,omitempty" xml:"corp_name,omitempty"`
	CostCenter    *FlightOrderListQueryResponseBodyModuleCostCenter `json:"cost_center,omitempty" xml:"cost_center,omitempty" type:"Struct"`
	DepAirport    *string                                           `json:"dep_airport,omitempty" xml:"dep_airport,omitempty"`
	DepCity       *string                                           `json:"dep_city,omitempty" xml:"dep_city,omitempty"`
	DepCityAdCode *string                                           `json:"dep_city_ad_code,omitempty" xml:"dep_city_ad_code,omitempty"`
	// example:
	//
	// 2022-07-20T10:40Z
	DepDate    *string `json:"dep_date,omitempty" xml:"dep_date,omitempty"`
	DepartId   *string `json:"depart_id,omitempty" xml:"depart_id,omitempty"`
	DepartName *string `json:"depart_name,omitempty" xml:"depart_name,omitempty"`
	// example:
	//
	// 30.12%
	Discount *string `json:"discount,omitempty" xml:"discount,omitempty"`
	// example:
	//
	// MU7854
	FlightNo *string `json:"flight_no,omitempty" xml:"flight_no,omitempty"`
	// example:
	//
	// 2022-07-20T10:40Z
	GmtCreate *string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// example:
	//
	// 2022-07-20T10:40Z
	GmtModified *string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// example:
	//
	// 200042
	Id             *int64                                                  `json:"id,omitempty" xml:"id,omitempty"`
	InsureInfoList []*FlightOrderListQueryResponseBodyModuleInsureInfoList `json:"insure_info_list,omitempty" xml:"insure_info_list,omitempty" type:"Repeated"`
	Invoice        *FlightOrderListQueryResponseBodyModuleInvoice          `json:"invoice,omitempty" xml:"invoice,omitempty" type:"Struct"`
	// example:
	//
	// 4
	PassengerCount *int32                                                 `json:"passenger_count,omitempty" xml:"passenger_count,omitempty"`
	PassengerName  *string                                                `json:"passenger_name,omitempty" xml:"passenger_name,omitempty"`
	PriceInfoList  []*FlightOrderListQueryResponseBodyModulePriceInfoList `json:"price_info_list,omitempty" xml:"price_info_list,omitempty" type:"Repeated"`
	// example:
	//
	// CS-PROJECT
	ProjectCode *string `json:"project_code,omitempty" xml:"project_code,omitempty"`
	// example:
	//
	// CS-PROJECT
	ProjectId    *int64  `json:"project_id,omitempty" xml:"project_id,omitempty"`
	ProjectTitle *string `json:"project_title,omitempty" xml:"project_title,omitempty"`
	// example:
	//
	// 2022-07-20T10:40Z
	RetDate *string `json:"ret_date,omitempty" xml:"ret_date,omitempty"`
	// example:
	//
	// 0
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// CS-PROJECT
	ThirdPartProjectId *string `json:"third_part_project_id,omitempty" xml:"third_part_project_id,omitempty"`
	// example:
	//
	// CS-UMN98989
	ThirdpartApplyId    *string `json:"thirdpart_apply_id,omitempty" xml:"thirdpart_apply_id,omitempty"`
	ThirdpartBusinessId *string `json:"thirdpart_business_id,omitempty" xml:"thirdpart_business_id,omitempty"`
	// example:
	//
	// cs9897766
	ThirdpartItineraryId *string `json:"thirdpart_itinerary_id,omitempty" xml:"thirdpart_itinerary_id,omitempty"`
	// example:
	//
	// 0
	TripType          *int32                                                     `json:"trip_type,omitempty" xml:"trip_type,omitempty"`
	UserAffiliateList []*FlightOrderListQueryResponseBodyModuleUserAffiliateList `json:"user_affiliate_list,omitempty" xml:"user_affiliate_list,omitempty" type:"Repeated"`
	UserId            *string                                                    `json:"user_id,omitempty" xml:"user_id,omitempty"`
	UserName          *string                                                    `json:"user_name,omitempty" xml:"user_name,omitempty"`
}

func (s FlightOrderListQueryResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s FlightOrderListQueryResponseBodyModule) GoString() string {
	return s.String()
}

func (s *FlightOrderListQueryResponseBodyModule) GetApplyId() *int64 {
	return s.ApplyId
}

func (s *FlightOrderListQueryResponseBodyModule) GetArrAirport() *string {
	return s.ArrAirport
}

func (s *FlightOrderListQueryResponseBodyModule) GetArrCity() *string {
	return s.ArrCity
}

func (s *FlightOrderListQueryResponseBodyModule) GetArrCityAdCode() *string {
	return s.ArrCityAdCode
}

func (s *FlightOrderListQueryResponseBodyModule) GetBtripTitle() *string {
	return s.BtripTitle
}

func (s *FlightOrderListQueryResponseBodyModule) GetCabinClass() *string {
	return s.CabinClass
}

func (s *FlightOrderListQueryResponseBodyModule) GetContactName() *string {
	return s.ContactName
}

func (s *FlightOrderListQueryResponseBodyModule) GetCorpId() *string {
	return s.CorpId
}

func (s *FlightOrderListQueryResponseBodyModule) GetCorpName() *string {
	return s.CorpName
}

func (s *FlightOrderListQueryResponseBodyModule) GetCostCenter() *FlightOrderListQueryResponseBodyModuleCostCenter {
	return s.CostCenter
}

func (s *FlightOrderListQueryResponseBodyModule) GetDepAirport() *string {
	return s.DepAirport
}

func (s *FlightOrderListQueryResponseBodyModule) GetDepCity() *string {
	return s.DepCity
}

func (s *FlightOrderListQueryResponseBodyModule) GetDepCityAdCode() *string {
	return s.DepCityAdCode
}

func (s *FlightOrderListQueryResponseBodyModule) GetDepDate() *string {
	return s.DepDate
}

func (s *FlightOrderListQueryResponseBodyModule) GetDepartId() *string {
	return s.DepartId
}

func (s *FlightOrderListQueryResponseBodyModule) GetDepartName() *string {
	return s.DepartName
}

func (s *FlightOrderListQueryResponseBodyModule) GetDiscount() *string {
	return s.Discount
}

func (s *FlightOrderListQueryResponseBodyModule) GetFlightNo() *string {
	return s.FlightNo
}

func (s *FlightOrderListQueryResponseBodyModule) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *FlightOrderListQueryResponseBodyModule) GetGmtModified() *string {
	return s.GmtModified
}

func (s *FlightOrderListQueryResponseBodyModule) GetId() *int64 {
	return s.Id
}

func (s *FlightOrderListQueryResponseBodyModule) GetInsureInfoList() []*FlightOrderListQueryResponseBodyModuleInsureInfoList {
	return s.InsureInfoList
}

func (s *FlightOrderListQueryResponseBodyModule) GetInvoice() *FlightOrderListQueryResponseBodyModuleInvoice {
	return s.Invoice
}

func (s *FlightOrderListQueryResponseBodyModule) GetPassengerCount() *int32 {
	return s.PassengerCount
}

func (s *FlightOrderListQueryResponseBodyModule) GetPassengerName() *string {
	return s.PassengerName
}

func (s *FlightOrderListQueryResponseBodyModule) GetPriceInfoList() []*FlightOrderListQueryResponseBodyModulePriceInfoList {
	return s.PriceInfoList
}

func (s *FlightOrderListQueryResponseBodyModule) GetProjectCode() *string {
	return s.ProjectCode
}

func (s *FlightOrderListQueryResponseBodyModule) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *FlightOrderListQueryResponseBodyModule) GetProjectTitle() *string {
	return s.ProjectTitle
}

func (s *FlightOrderListQueryResponseBodyModule) GetRetDate() *string {
	return s.RetDate
}

func (s *FlightOrderListQueryResponseBodyModule) GetStatus() *int32 {
	return s.Status
}

func (s *FlightOrderListQueryResponseBodyModule) GetThirdPartProjectId() *string {
	return s.ThirdPartProjectId
}

func (s *FlightOrderListQueryResponseBodyModule) GetThirdpartApplyId() *string {
	return s.ThirdpartApplyId
}

func (s *FlightOrderListQueryResponseBodyModule) GetThirdpartBusinessId() *string {
	return s.ThirdpartBusinessId
}

func (s *FlightOrderListQueryResponseBodyModule) GetThirdpartItineraryId() *string {
	return s.ThirdpartItineraryId
}

func (s *FlightOrderListQueryResponseBodyModule) GetTripType() *int32 {
	return s.TripType
}

func (s *FlightOrderListQueryResponseBodyModule) GetUserAffiliateList() []*FlightOrderListQueryResponseBodyModuleUserAffiliateList {
	return s.UserAffiliateList
}

func (s *FlightOrderListQueryResponseBodyModule) GetUserId() *string {
	return s.UserId
}

func (s *FlightOrderListQueryResponseBodyModule) GetUserName() *string {
	return s.UserName
}

func (s *FlightOrderListQueryResponseBodyModule) SetApplyId(v int64) *FlightOrderListQueryResponseBodyModule {
	s.ApplyId = &v
	return s
}

func (s *FlightOrderListQueryResponseBodyModule) SetArrAirport(v string) *FlightOrderListQueryResponseBodyModule {
	s.ArrAirport = &v
	return s
}

func (s *FlightOrderListQueryResponseBodyModule) SetArrCity(v string) *FlightOrderListQueryResponseBodyModule {
	s.ArrCity = &v
	return s
}

func (s *FlightOrderListQueryResponseBodyModule) SetArrCityAdCode(v string) *FlightOrderListQueryResponseBodyModule {
	s.ArrCityAdCode = &v
	return s
}

func (s *FlightOrderListQueryResponseBodyModule) SetBtripTitle(v string) *FlightOrderListQueryResponseBodyModule {
	s.BtripTitle = &v
	return s
}

func (s *FlightOrderListQueryResponseBodyModule) SetCabinClass(v string) *FlightOrderListQueryResponseBodyModule {
	s.CabinClass = &v
	return s
}

func (s *FlightOrderListQueryResponseBodyModule) SetContactName(v string) *FlightOrderListQueryResponseBodyModule {
	s.ContactName = &v
	return s
}

func (s *FlightOrderListQueryResponseBodyModule) SetCorpId(v string) *FlightOrderListQueryResponseBodyModule {
	s.CorpId = &v
	return s
}

func (s *FlightOrderListQueryResponseBodyModule) SetCorpName(v string) *FlightOrderListQueryResponseBodyModule {
	s.CorpName = &v
	return s
}

func (s *FlightOrderListQueryResponseBodyModule) SetCostCenter(v *FlightOrderListQueryResponseBodyModuleCostCenter) *FlightOrderListQueryResponseBodyModule {
	s.CostCenter = v
	return s
}

func (s *FlightOrderListQueryResponseBodyModule) SetDepAirport(v string) *FlightOrderListQueryResponseBodyModule {
	s.DepAirport = &v
	return s
}

func (s *FlightOrderListQueryResponseBodyModule) SetDepCity(v string) *FlightOrderListQueryResponseBodyModule {
	s.DepCity = &v
	return s
}

func (s *FlightOrderListQueryResponseBodyModule) SetDepCityAdCode(v string) *FlightOrderListQueryResponseBodyModule {
	s.DepCityAdCode = &v
	return s
}

func (s *FlightOrderListQueryResponseBodyModule) SetDepDate(v string) *FlightOrderListQueryResponseBodyModule {
	s.DepDate = &v
	return s
}

func (s *FlightOrderListQueryResponseBodyModule) SetDepartId(v string) *FlightOrderListQueryResponseBodyModule {
	s.DepartId = &v
	return s
}

func (s *FlightOrderListQueryResponseBodyModule) SetDepartName(v string) *FlightOrderListQueryResponseBodyModule {
	s.DepartName = &v
	return s
}

func (s *FlightOrderListQueryResponseBodyModule) SetDiscount(v string) *FlightOrderListQueryResponseBodyModule {
	s.Discount = &v
	return s
}

func (s *FlightOrderListQueryResponseBodyModule) SetFlightNo(v string) *FlightOrderListQueryResponseBodyModule {
	s.FlightNo = &v
	return s
}

func (s *FlightOrderListQueryResponseBodyModule) SetGmtCreate(v string) *FlightOrderListQueryResponseBodyModule {
	s.GmtCreate = &v
	return s
}

func (s *FlightOrderListQueryResponseBodyModule) SetGmtModified(v string) *FlightOrderListQueryResponseBodyModule {
	s.GmtModified = &v
	return s
}

func (s *FlightOrderListQueryResponseBodyModule) SetId(v int64) *FlightOrderListQueryResponseBodyModule {
	s.Id = &v
	return s
}

func (s *FlightOrderListQueryResponseBodyModule) SetInsureInfoList(v []*FlightOrderListQueryResponseBodyModuleInsureInfoList) *FlightOrderListQueryResponseBodyModule {
	s.InsureInfoList = v
	return s
}

func (s *FlightOrderListQueryResponseBodyModule) SetInvoice(v *FlightOrderListQueryResponseBodyModuleInvoice) *FlightOrderListQueryResponseBodyModule {
	s.Invoice = v
	return s
}

func (s *FlightOrderListQueryResponseBodyModule) SetPassengerCount(v int32) *FlightOrderListQueryResponseBodyModule {
	s.PassengerCount = &v
	return s
}

func (s *FlightOrderListQueryResponseBodyModule) SetPassengerName(v string) *FlightOrderListQueryResponseBodyModule {
	s.PassengerName = &v
	return s
}

func (s *FlightOrderListQueryResponseBodyModule) SetPriceInfoList(v []*FlightOrderListQueryResponseBodyModulePriceInfoList) *FlightOrderListQueryResponseBodyModule {
	s.PriceInfoList = v
	return s
}

func (s *FlightOrderListQueryResponseBodyModule) SetProjectCode(v string) *FlightOrderListQueryResponseBodyModule {
	s.ProjectCode = &v
	return s
}

func (s *FlightOrderListQueryResponseBodyModule) SetProjectId(v int64) *FlightOrderListQueryResponseBodyModule {
	s.ProjectId = &v
	return s
}

func (s *FlightOrderListQueryResponseBodyModule) SetProjectTitle(v string) *FlightOrderListQueryResponseBodyModule {
	s.ProjectTitle = &v
	return s
}

func (s *FlightOrderListQueryResponseBodyModule) SetRetDate(v string) *FlightOrderListQueryResponseBodyModule {
	s.RetDate = &v
	return s
}

func (s *FlightOrderListQueryResponseBodyModule) SetStatus(v int32) *FlightOrderListQueryResponseBodyModule {
	s.Status = &v
	return s
}

func (s *FlightOrderListQueryResponseBodyModule) SetThirdPartProjectId(v string) *FlightOrderListQueryResponseBodyModule {
	s.ThirdPartProjectId = &v
	return s
}

func (s *FlightOrderListQueryResponseBodyModule) SetThirdpartApplyId(v string) *FlightOrderListQueryResponseBodyModule {
	s.ThirdpartApplyId = &v
	return s
}

func (s *FlightOrderListQueryResponseBodyModule) SetThirdpartBusinessId(v string) *FlightOrderListQueryResponseBodyModule {
	s.ThirdpartBusinessId = &v
	return s
}

func (s *FlightOrderListQueryResponseBodyModule) SetThirdpartItineraryId(v string) *FlightOrderListQueryResponseBodyModule {
	s.ThirdpartItineraryId = &v
	return s
}

func (s *FlightOrderListQueryResponseBodyModule) SetTripType(v int32) *FlightOrderListQueryResponseBodyModule {
	s.TripType = &v
	return s
}

func (s *FlightOrderListQueryResponseBodyModule) SetUserAffiliateList(v []*FlightOrderListQueryResponseBodyModuleUserAffiliateList) *FlightOrderListQueryResponseBodyModule {
	s.UserAffiliateList = v
	return s
}

func (s *FlightOrderListQueryResponseBodyModule) SetUserId(v string) *FlightOrderListQueryResponseBodyModule {
	s.UserId = &v
	return s
}

func (s *FlightOrderListQueryResponseBodyModule) SetUserName(v string) *FlightOrderListQueryResponseBodyModule {
	s.UserName = &v
	return s
}

func (s *FlightOrderListQueryResponseBodyModule) Validate() error {
	return dara.Validate(s)
}

type FlightOrderListQueryResponseBodyModuleCostCenter struct {
	CorpId *string `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
	// example:
	//
	// 44632
	Id   *int64  `json:"id,omitempty" xml:"id,omitempty"`
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// NM98767
	Number *string `json:"number,omitempty" xml:"number,omitempty"`
}

func (s FlightOrderListQueryResponseBodyModuleCostCenter) String() string {
	return dara.Prettify(s)
}

func (s FlightOrderListQueryResponseBodyModuleCostCenter) GoString() string {
	return s.String()
}

func (s *FlightOrderListQueryResponseBodyModuleCostCenter) GetCorpId() *string {
	return s.CorpId
}

func (s *FlightOrderListQueryResponseBodyModuleCostCenter) GetId() *int64 {
	return s.Id
}

func (s *FlightOrderListQueryResponseBodyModuleCostCenter) GetName() *string {
	return s.Name
}

func (s *FlightOrderListQueryResponseBodyModuleCostCenter) GetNumber() *string {
	return s.Number
}

func (s *FlightOrderListQueryResponseBodyModuleCostCenter) SetCorpId(v string) *FlightOrderListQueryResponseBodyModuleCostCenter {
	s.CorpId = &v
	return s
}

func (s *FlightOrderListQueryResponseBodyModuleCostCenter) SetId(v int64) *FlightOrderListQueryResponseBodyModuleCostCenter {
	s.Id = &v
	return s
}

func (s *FlightOrderListQueryResponseBodyModuleCostCenter) SetName(v string) *FlightOrderListQueryResponseBodyModuleCostCenter {
	s.Name = &v
	return s
}

func (s *FlightOrderListQueryResponseBodyModuleCostCenter) SetNumber(v string) *FlightOrderListQueryResponseBodyModuleCostCenter {
	s.Number = &v
	return s
}

func (s *FlightOrderListQueryResponseBodyModuleCostCenter) Validate() error {
	return dara.Validate(s)
}

type FlightOrderListQueryResponseBodyModuleInsureInfoList struct {
	// example:
	//
	// KJ-879657
	InsureNo *string `json:"insure_no,omitempty" xml:"insure_no,omitempty"`
	Name     *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 1
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
}

func (s FlightOrderListQueryResponseBodyModuleInsureInfoList) String() string {
	return dara.Prettify(s)
}

func (s FlightOrderListQueryResponseBodyModuleInsureInfoList) GoString() string {
	return s.String()
}

func (s *FlightOrderListQueryResponseBodyModuleInsureInfoList) GetInsureNo() *string {
	return s.InsureNo
}

func (s *FlightOrderListQueryResponseBodyModuleInsureInfoList) GetName() *string {
	return s.Name
}

func (s *FlightOrderListQueryResponseBodyModuleInsureInfoList) GetStatus() *int32 {
	return s.Status
}

func (s *FlightOrderListQueryResponseBodyModuleInsureInfoList) SetInsureNo(v string) *FlightOrderListQueryResponseBodyModuleInsureInfoList {
	s.InsureNo = &v
	return s
}

func (s *FlightOrderListQueryResponseBodyModuleInsureInfoList) SetName(v string) *FlightOrderListQueryResponseBodyModuleInsureInfoList {
	s.Name = &v
	return s
}

func (s *FlightOrderListQueryResponseBodyModuleInsureInfoList) SetStatus(v int32) *FlightOrderListQueryResponseBodyModuleInsureInfoList {
	s.Status = &v
	return s
}

func (s *FlightOrderListQueryResponseBodyModuleInsureInfoList) Validate() error {
	return dara.Validate(s)
}

type FlightOrderListQueryResponseBodyModuleInvoice struct {
	// example:
	//
	// 7304
	Id    *int64  `json:"id,omitempty" xml:"id,omitempty"`
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s FlightOrderListQueryResponseBodyModuleInvoice) String() string {
	return dara.Prettify(s)
}

func (s FlightOrderListQueryResponseBodyModuleInvoice) GoString() string {
	return s.String()
}

func (s *FlightOrderListQueryResponseBodyModuleInvoice) GetId() *int64 {
	return s.Id
}

func (s *FlightOrderListQueryResponseBodyModuleInvoice) GetTitle() *string {
	return s.Title
}

func (s *FlightOrderListQueryResponseBodyModuleInvoice) SetId(v int64) *FlightOrderListQueryResponseBodyModuleInvoice {
	s.Id = &v
	return s
}

func (s *FlightOrderListQueryResponseBodyModuleInvoice) SetTitle(v string) *FlightOrderListQueryResponseBodyModuleInvoice {
	s.Title = &v
	return s
}

func (s *FlightOrderListQueryResponseBodyModuleInvoice) Validate() error {
	return dara.Validate(s)
}

type FlightOrderListQueryResponseBodyModulePriceInfoList struct {
	// example:
	//
	// 1
	CategoryCode *int32 `json:"category_code,omitempty" xml:"category_code,omitempty"`
	// example:
	//
	// 1
	CategoryType *int32 `json:"category_type,omitempty" xml:"category_type,omitempty"`
	// example:
	//
	// MU5354
	ChangeFlightNo *string `json:"change_flight_no,omitempty" xml:"change_flight_no,omitempty"`
	// example:
	//
	// 12%
	Discount *string `json:"discount,omitempty" xml:"discount,omitempty"`
	// example:
	//
	// 2022-07-20T10:40Z
	EndTime *string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// example:
	//
	// 2022-07-20T10:40Z
	GmtCreate *string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// example:
	//
	// 00-7687
	OriginalTicketNo *string `json:"original_ticket_no,omitempty" xml:"original_ticket_no,omitempty"`
	PassengerName    *string `json:"passenger_name,omitempty" xml:"passenger_name,omitempty"`
	// example:
	//
	// 1
	PayType *int32 `json:"pay_type,omitempty" xml:"pay_type,omitempty"`
	// example:
	//
	// 100
	Price *float64 `json:"price,omitempty" xml:"price,omitempty"`
	// example:
	//
	// 2022-07-20T10:40Z
	StartTime *string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// example:
	//
	// A-135767
	TicketNo *string `json:"ticket_no,omitempty" xml:"ticket_no,omitempty"`
	// example:
	//
	// f98236773
	TradeId *string `json:"trade_id,omitempty" xml:"trade_id,omitempty"`
	// example:
	//
	// 1
	Type *int32 `json:"type,omitempty" xml:"type,omitempty"`
}

func (s FlightOrderListQueryResponseBodyModulePriceInfoList) String() string {
	return dara.Prettify(s)
}

func (s FlightOrderListQueryResponseBodyModulePriceInfoList) GoString() string {
	return s.String()
}

func (s *FlightOrderListQueryResponseBodyModulePriceInfoList) GetCategoryCode() *int32 {
	return s.CategoryCode
}

func (s *FlightOrderListQueryResponseBodyModulePriceInfoList) GetCategoryType() *int32 {
	return s.CategoryType
}

func (s *FlightOrderListQueryResponseBodyModulePriceInfoList) GetChangeFlightNo() *string {
	return s.ChangeFlightNo
}

func (s *FlightOrderListQueryResponseBodyModulePriceInfoList) GetDiscount() *string {
	return s.Discount
}

func (s *FlightOrderListQueryResponseBodyModulePriceInfoList) GetEndTime() *string {
	return s.EndTime
}

func (s *FlightOrderListQueryResponseBodyModulePriceInfoList) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *FlightOrderListQueryResponseBodyModulePriceInfoList) GetOriginalTicketNo() *string {
	return s.OriginalTicketNo
}

func (s *FlightOrderListQueryResponseBodyModulePriceInfoList) GetPassengerName() *string {
	return s.PassengerName
}

func (s *FlightOrderListQueryResponseBodyModulePriceInfoList) GetPayType() *int32 {
	return s.PayType
}

func (s *FlightOrderListQueryResponseBodyModulePriceInfoList) GetPrice() *float64 {
	return s.Price
}

func (s *FlightOrderListQueryResponseBodyModulePriceInfoList) GetStartTime() *string {
	return s.StartTime
}

func (s *FlightOrderListQueryResponseBodyModulePriceInfoList) GetTicketNo() *string {
	return s.TicketNo
}

func (s *FlightOrderListQueryResponseBodyModulePriceInfoList) GetTradeId() *string {
	return s.TradeId
}

func (s *FlightOrderListQueryResponseBodyModulePriceInfoList) GetType() *int32 {
	return s.Type
}

func (s *FlightOrderListQueryResponseBodyModulePriceInfoList) SetCategoryCode(v int32) *FlightOrderListQueryResponseBodyModulePriceInfoList {
	s.CategoryCode = &v
	return s
}

func (s *FlightOrderListQueryResponseBodyModulePriceInfoList) SetCategoryType(v int32) *FlightOrderListQueryResponseBodyModulePriceInfoList {
	s.CategoryType = &v
	return s
}

func (s *FlightOrderListQueryResponseBodyModulePriceInfoList) SetChangeFlightNo(v string) *FlightOrderListQueryResponseBodyModulePriceInfoList {
	s.ChangeFlightNo = &v
	return s
}

func (s *FlightOrderListQueryResponseBodyModulePriceInfoList) SetDiscount(v string) *FlightOrderListQueryResponseBodyModulePriceInfoList {
	s.Discount = &v
	return s
}

func (s *FlightOrderListQueryResponseBodyModulePriceInfoList) SetEndTime(v string) *FlightOrderListQueryResponseBodyModulePriceInfoList {
	s.EndTime = &v
	return s
}

func (s *FlightOrderListQueryResponseBodyModulePriceInfoList) SetGmtCreate(v string) *FlightOrderListQueryResponseBodyModulePriceInfoList {
	s.GmtCreate = &v
	return s
}

func (s *FlightOrderListQueryResponseBodyModulePriceInfoList) SetOriginalTicketNo(v string) *FlightOrderListQueryResponseBodyModulePriceInfoList {
	s.OriginalTicketNo = &v
	return s
}

func (s *FlightOrderListQueryResponseBodyModulePriceInfoList) SetPassengerName(v string) *FlightOrderListQueryResponseBodyModulePriceInfoList {
	s.PassengerName = &v
	return s
}

func (s *FlightOrderListQueryResponseBodyModulePriceInfoList) SetPayType(v int32) *FlightOrderListQueryResponseBodyModulePriceInfoList {
	s.PayType = &v
	return s
}

func (s *FlightOrderListQueryResponseBodyModulePriceInfoList) SetPrice(v float64) *FlightOrderListQueryResponseBodyModulePriceInfoList {
	s.Price = &v
	return s
}

func (s *FlightOrderListQueryResponseBodyModulePriceInfoList) SetStartTime(v string) *FlightOrderListQueryResponseBodyModulePriceInfoList {
	s.StartTime = &v
	return s
}

func (s *FlightOrderListQueryResponseBodyModulePriceInfoList) SetTicketNo(v string) *FlightOrderListQueryResponseBodyModulePriceInfoList {
	s.TicketNo = &v
	return s
}

func (s *FlightOrderListQueryResponseBodyModulePriceInfoList) SetTradeId(v string) *FlightOrderListQueryResponseBodyModulePriceInfoList {
	s.TradeId = &v
	return s
}

func (s *FlightOrderListQueryResponseBodyModulePriceInfoList) SetType(v int32) *FlightOrderListQueryResponseBodyModulePriceInfoList {
	s.Type = &v
	return s
}

func (s *FlightOrderListQueryResponseBodyModulePriceInfoList) Validate() error {
	return dara.Validate(s)
}

type FlightOrderListQueryResponseBodyModuleUserAffiliateList struct {
	UserId   *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	UserName *string `json:"user_name,omitempty" xml:"user_name,omitempty"`
}

func (s FlightOrderListQueryResponseBodyModuleUserAffiliateList) String() string {
	return dara.Prettify(s)
}

func (s FlightOrderListQueryResponseBodyModuleUserAffiliateList) GoString() string {
	return s.String()
}

func (s *FlightOrderListQueryResponseBodyModuleUserAffiliateList) GetUserId() *string {
	return s.UserId
}

func (s *FlightOrderListQueryResponseBodyModuleUserAffiliateList) GetUserName() *string {
	return s.UserName
}

func (s *FlightOrderListQueryResponseBodyModuleUserAffiliateList) SetUserId(v string) *FlightOrderListQueryResponseBodyModuleUserAffiliateList {
	s.UserId = &v
	return s
}

func (s *FlightOrderListQueryResponseBodyModuleUserAffiliateList) SetUserName(v string) *FlightOrderListQueryResponseBodyModuleUserAffiliateList {
	s.UserName = &v
	return s
}

func (s *FlightOrderListQueryResponseBodyModuleUserAffiliateList) Validate() error {
	return dara.Validate(s)
}

type FlightOrderListQueryResponseBodyPageInfo struct {
	// example:
	//
	// 1
	Page *int32 `json:"page,omitempty" xml:"page,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// example:
	//
	// 100
	TotalNumber *int32 `json:"total_number,omitempty" xml:"total_number,omitempty"`
}

func (s FlightOrderListQueryResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s FlightOrderListQueryResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *FlightOrderListQueryResponseBodyPageInfo) GetPage() *int32 {
	return s.Page
}

func (s *FlightOrderListQueryResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *FlightOrderListQueryResponseBodyPageInfo) GetTotalNumber() *int32 {
	return s.TotalNumber
}

func (s *FlightOrderListQueryResponseBodyPageInfo) SetPage(v int32) *FlightOrderListQueryResponseBodyPageInfo {
	s.Page = &v
	return s
}

func (s *FlightOrderListQueryResponseBodyPageInfo) SetPageSize(v int32) *FlightOrderListQueryResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *FlightOrderListQueryResponseBodyPageInfo) SetTotalNumber(v int32) *FlightOrderListQueryResponseBodyPageInfo {
	s.TotalNumber = &v
	return s
}

func (s *FlightOrderListQueryResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}
