// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTrainOrderListQueryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *TrainOrderListQueryResponseBody
	GetCode() *string
	SetMessage(v string) *TrainOrderListQueryResponseBody
	GetMessage() *string
	SetModule(v []*TrainOrderListQueryResponseBodyModule) *TrainOrderListQueryResponseBody
	GetModule() []*TrainOrderListQueryResponseBodyModule
	SetPageInfo(v *TrainOrderListQueryResponseBodyPageInfo) *TrainOrderListQueryResponseBody
	GetPageInfo() *TrainOrderListQueryResponseBodyPageInfo
	SetRequestId(v string) *TrainOrderListQueryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *TrainOrderListQueryResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *TrainOrderListQueryResponseBody
	GetTraceId() *string
}

type TrainOrderListQueryResponseBody struct {
	// example:
	//
	// SUCCESS
	Code     *string                                  `json:"code,omitempty" xml:"code,omitempty"`
	Message  *string                                  `json:"message,omitempty" xml:"message,omitempty"`
	Module   []*TrainOrderListQueryResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Repeated"`
	PageInfo *TrainOrderListQueryResponseBodyPageInfo `json:"page_info,omitempty" xml:"page_info,omitempty" type:"Struct"`
	// example:
	//
	// F93C3EBD-17BE-5FE6-BF06-96A6F1AC8DC5
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

func (s TrainOrderListQueryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TrainOrderListQueryResponseBody) GoString() string {
	return s.String()
}

func (s *TrainOrderListQueryResponseBody) GetCode() *string {
	return s.Code
}

func (s *TrainOrderListQueryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *TrainOrderListQueryResponseBody) GetModule() []*TrainOrderListQueryResponseBodyModule {
	return s.Module
}

func (s *TrainOrderListQueryResponseBody) GetPageInfo() *TrainOrderListQueryResponseBodyPageInfo {
	return s.PageInfo
}

func (s *TrainOrderListQueryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TrainOrderListQueryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *TrainOrderListQueryResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *TrainOrderListQueryResponseBody) SetCode(v string) *TrainOrderListQueryResponseBody {
	s.Code = &v
	return s
}

func (s *TrainOrderListQueryResponseBody) SetMessage(v string) *TrainOrderListQueryResponseBody {
	s.Message = &v
	return s
}

func (s *TrainOrderListQueryResponseBody) SetModule(v []*TrainOrderListQueryResponseBodyModule) *TrainOrderListQueryResponseBody {
	s.Module = v
	return s
}

func (s *TrainOrderListQueryResponseBody) SetPageInfo(v *TrainOrderListQueryResponseBodyPageInfo) *TrainOrderListQueryResponseBody {
	s.PageInfo = v
	return s
}

func (s *TrainOrderListQueryResponseBody) SetRequestId(v string) *TrainOrderListQueryResponseBody {
	s.RequestId = &v
	return s
}

func (s *TrainOrderListQueryResponseBody) SetSuccess(v bool) *TrainOrderListQueryResponseBody {
	s.Success = &v
	return s
}

func (s *TrainOrderListQueryResponseBody) SetTraceId(v string) *TrainOrderListQueryResponseBody {
	s.TraceId = &v
	return s
}

func (s *TrainOrderListQueryResponseBody) Validate() error {
	return dara.Validate(s)
}

type TrainOrderListQueryResponseBodyModule struct {
	// example:
	//
	// 11367
	ApplyId       *int64  `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
	ArrCity       *string `json:"arr_city,omitempty" xml:"arr_city,omitempty"`
	ArrCityAdCode *string `json:"arr_city_ad_code,omitempty" xml:"arr_city_ad_code,omitempty"`
	ArrStation    *string `json:"arr_station,omitempty" xml:"arr_station,omitempty"`
	// example:
	//
	// 2022-05-15T22:27Z
	ArrTime       *string                                          `json:"arr_time,omitempty" xml:"arr_time,omitempty"`
	BtripTitle    *string                                          `json:"btrip_title,omitempty" xml:"btrip_title,omitempty"`
	ContactName   *string                                          `json:"contact_name,omitempty" xml:"contact_name,omitempty"`
	CorpId        *string                                          `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
	CorpName      *string                                          `json:"corp_name,omitempty" xml:"corp_name,omitempty"`
	CostCenter    *TrainOrderListQueryResponseBodyModuleCostCenter `json:"cost_center,omitempty" xml:"cost_center,omitempty" type:"Struct"`
	DepCity       *string                                          `json:"dep_city,omitempty" xml:"dep_city,omitempty"`
	DepCityAdCode *string                                          `json:"dep_city_ad_code,omitempty" xml:"dep_city_ad_code,omitempty"`
	DepStation    *string                                          `json:"dep_station,omitempty" xml:"dep_station,omitempty"`
	// example:
	//
	// 2022-05-15T22:27Z
	DepTime    *string `json:"dep_time,omitempty" xml:"dep_time,omitempty"`
	DepartId   *string `json:"depart_id,omitempty" xml:"depart_id,omitempty"`
	DepartName *string `json:"depart_name,omitempty" xml:"depart_name,omitempty"`
	// example:
	//
	// 2022-05-15T22:27Z
	GmtCreate *string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// example:
	//
	// 2022-05-15T22:27Z
	GmtModified *string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// example:
	//
	// 1478652
	Id            *int64                                                `json:"id,omitempty" xml:"id,omitempty"`
	Invoice       *TrainOrderListQueryResponseBodyModuleInvoice         `json:"invoice,omitempty" xml:"invoice,omitempty" type:"Struct"`
	PriceInfoList []*TrainOrderListQueryResponseBodyModulePriceInfoList `json:"price_info_list,omitempty" xml:"price_info_list,omitempty" type:"Repeated"`
	// example:
	//
	// PCXIDF
	ProjectCode *string `json:"project_code,omitempty" xml:"project_code,omitempty"`
	// example:
	//
	// 12425
	ProjectId    *int64  `json:"project_id,omitempty" xml:"project_id,omitempty"`
	ProjectTitle *string `json:"project_title,omitempty" xml:"project_title,omitempty"`
	RiderName    *string `json:"rider_name,omitempty" xml:"rider_name,omitempty"`
	// example:
	//
	// 100
	RunTime  *string `json:"run_time,omitempty" xml:"run_time,omitempty"`
	SeatType *string `json:"seat_type,omitempty" xml:"seat_type,omitempty"`
	// example:
	//
	// 0
	Status              *int32  `json:"status,omitempty" xml:"status,omitempty"`
	ThirdPartBusinessId *string `json:"thirdPart_business_id,omitempty" xml:"thirdPart_business_id,omitempty"`
	// example:
	//
	// CS-SKPFDS
	ThirdPartProjectId *string `json:"third_part_project_id,omitempty" xml:"third_part_project_id,omitempty"`
	// example:
	//
	// CSIODJUSN
	ThirdpartApplyId *string `json:"thirdpart_apply_id,omitempty" xml:"thirdpart_apply_id,omitempty"`
	// example:
	//
	// CS-IT89D
	ThirdpartItineraryId *string `json:"thirdpart_itinerary_id,omitempty" xml:"thirdpart_itinerary_id,omitempty"`
	// example:
	//
	// 1
	TicketCount *int32 `json:"ticket_count,omitempty" xml:"ticket_count,omitempty"`
	// example:
	//
	// CS-663
	TicketNo12306 *string `json:"ticket_no12306,omitempty" xml:"ticket_no12306,omitempty"`
	// example:
	//
	// CS-663
	TrainNumber       *string                                                   `json:"train_number,omitempty" xml:"train_number,omitempty"`
	TrainType         *string                                                   `json:"train_type,omitempty" xml:"train_type,omitempty"`
	UserAffiliateList []*TrainOrderListQueryResponseBodyModuleUserAffiliateList `json:"user_affiliate_list,omitempty" xml:"user_affiliate_list,omitempty" type:"Repeated"`
	UserId            *string                                                   `json:"user_id,omitempty" xml:"user_id,omitempty"`
	UserName          *string                                                   `json:"user_name,omitempty" xml:"user_name,omitempty"`
}

func (s TrainOrderListQueryResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s TrainOrderListQueryResponseBodyModule) GoString() string {
	return s.String()
}

func (s *TrainOrderListQueryResponseBodyModule) GetApplyId() *int64 {
	return s.ApplyId
}

func (s *TrainOrderListQueryResponseBodyModule) GetArrCity() *string {
	return s.ArrCity
}

func (s *TrainOrderListQueryResponseBodyModule) GetArrCityAdCode() *string {
	return s.ArrCityAdCode
}

func (s *TrainOrderListQueryResponseBodyModule) GetArrStation() *string {
	return s.ArrStation
}

func (s *TrainOrderListQueryResponseBodyModule) GetArrTime() *string {
	return s.ArrTime
}

func (s *TrainOrderListQueryResponseBodyModule) GetBtripTitle() *string {
	return s.BtripTitle
}

func (s *TrainOrderListQueryResponseBodyModule) GetContactName() *string {
	return s.ContactName
}

func (s *TrainOrderListQueryResponseBodyModule) GetCorpId() *string {
	return s.CorpId
}

func (s *TrainOrderListQueryResponseBodyModule) GetCorpName() *string {
	return s.CorpName
}

func (s *TrainOrderListQueryResponseBodyModule) GetCostCenter() *TrainOrderListQueryResponseBodyModuleCostCenter {
	return s.CostCenter
}

func (s *TrainOrderListQueryResponseBodyModule) GetDepCity() *string {
	return s.DepCity
}

func (s *TrainOrderListQueryResponseBodyModule) GetDepCityAdCode() *string {
	return s.DepCityAdCode
}

func (s *TrainOrderListQueryResponseBodyModule) GetDepStation() *string {
	return s.DepStation
}

func (s *TrainOrderListQueryResponseBodyModule) GetDepTime() *string {
	return s.DepTime
}

func (s *TrainOrderListQueryResponseBodyModule) GetDepartId() *string {
	return s.DepartId
}

func (s *TrainOrderListQueryResponseBodyModule) GetDepartName() *string {
	return s.DepartName
}

func (s *TrainOrderListQueryResponseBodyModule) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *TrainOrderListQueryResponseBodyModule) GetGmtModified() *string {
	return s.GmtModified
}

func (s *TrainOrderListQueryResponseBodyModule) GetId() *int64 {
	return s.Id
}

func (s *TrainOrderListQueryResponseBodyModule) GetInvoice() *TrainOrderListQueryResponseBodyModuleInvoice {
	return s.Invoice
}

func (s *TrainOrderListQueryResponseBodyModule) GetPriceInfoList() []*TrainOrderListQueryResponseBodyModulePriceInfoList {
	return s.PriceInfoList
}

func (s *TrainOrderListQueryResponseBodyModule) GetProjectCode() *string {
	return s.ProjectCode
}

func (s *TrainOrderListQueryResponseBodyModule) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *TrainOrderListQueryResponseBodyModule) GetProjectTitle() *string {
	return s.ProjectTitle
}

func (s *TrainOrderListQueryResponseBodyModule) GetRiderName() *string {
	return s.RiderName
}

func (s *TrainOrderListQueryResponseBodyModule) GetRunTime() *string {
	return s.RunTime
}

func (s *TrainOrderListQueryResponseBodyModule) GetSeatType() *string {
	return s.SeatType
}

func (s *TrainOrderListQueryResponseBodyModule) GetStatus() *int32 {
	return s.Status
}

func (s *TrainOrderListQueryResponseBodyModule) GetThirdPartBusinessId() *string {
	return s.ThirdPartBusinessId
}

func (s *TrainOrderListQueryResponseBodyModule) GetThirdPartProjectId() *string {
	return s.ThirdPartProjectId
}

func (s *TrainOrderListQueryResponseBodyModule) GetThirdpartApplyId() *string {
	return s.ThirdpartApplyId
}

func (s *TrainOrderListQueryResponseBodyModule) GetThirdpartItineraryId() *string {
	return s.ThirdpartItineraryId
}

func (s *TrainOrderListQueryResponseBodyModule) GetTicketCount() *int32 {
	return s.TicketCount
}

func (s *TrainOrderListQueryResponseBodyModule) GetTicketNo12306() *string {
	return s.TicketNo12306
}

func (s *TrainOrderListQueryResponseBodyModule) GetTrainNumber() *string {
	return s.TrainNumber
}

func (s *TrainOrderListQueryResponseBodyModule) GetTrainType() *string {
	return s.TrainType
}

func (s *TrainOrderListQueryResponseBodyModule) GetUserAffiliateList() []*TrainOrderListQueryResponseBodyModuleUserAffiliateList {
	return s.UserAffiliateList
}

func (s *TrainOrderListQueryResponseBodyModule) GetUserId() *string {
	return s.UserId
}

func (s *TrainOrderListQueryResponseBodyModule) GetUserName() *string {
	return s.UserName
}

func (s *TrainOrderListQueryResponseBodyModule) SetApplyId(v int64) *TrainOrderListQueryResponseBodyModule {
	s.ApplyId = &v
	return s
}

func (s *TrainOrderListQueryResponseBodyModule) SetArrCity(v string) *TrainOrderListQueryResponseBodyModule {
	s.ArrCity = &v
	return s
}

func (s *TrainOrderListQueryResponseBodyModule) SetArrCityAdCode(v string) *TrainOrderListQueryResponseBodyModule {
	s.ArrCityAdCode = &v
	return s
}

func (s *TrainOrderListQueryResponseBodyModule) SetArrStation(v string) *TrainOrderListQueryResponseBodyModule {
	s.ArrStation = &v
	return s
}

func (s *TrainOrderListQueryResponseBodyModule) SetArrTime(v string) *TrainOrderListQueryResponseBodyModule {
	s.ArrTime = &v
	return s
}

func (s *TrainOrderListQueryResponseBodyModule) SetBtripTitle(v string) *TrainOrderListQueryResponseBodyModule {
	s.BtripTitle = &v
	return s
}

func (s *TrainOrderListQueryResponseBodyModule) SetContactName(v string) *TrainOrderListQueryResponseBodyModule {
	s.ContactName = &v
	return s
}

func (s *TrainOrderListQueryResponseBodyModule) SetCorpId(v string) *TrainOrderListQueryResponseBodyModule {
	s.CorpId = &v
	return s
}

func (s *TrainOrderListQueryResponseBodyModule) SetCorpName(v string) *TrainOrderListQueryResponseBodyModule {
	s.CorpName = &v
	return s
}

func (s *TrainOrderListQueryResponseBodyModule) SetCostCenter(v *TrainOrderListQueryResponseBodyModuleCostCenter) *TrainOrderListQueryResponseBodyModule {
	s.CostCenter = v
	return s
}

func (s *TrainOrderListQueryResponseBodyModule) SetDepCity(v string) *TrainOrderListQueryResponseBodyModule {
	s.DepCity = &v
	return s
}

func (s *TrainOrderListQueryResponseBodyModule) SetDepCityAdCode(v string) *TrainOrderListQueryResponseBodyModule {
	s.DepCityAdCode = &v
	return s
}

func (s *TrainOrderListQueryResponseBodyModule) SetDepStation(v string) *TrainOrderListQueryResponseBodyModule {
	s.DepStation = &v
	return s
}

func (s *TrainOrderListQueryResponseBodyModule) SetDepTime(v string) *TrainOrderListQueryResponseBodyModule {
	s.DepTime = &v
	return s
}

func (s *TrainOrderListQueryResponseBodyModule) SetDepartId(v string) *TrainOrderListQueryResponseBodyModule {
	s.DepartId = &v
	return s
}

func (s *TrainOrderListQueryResponseBodyModule) SetDepartName(v string) *TrainOrderListQueryResponseBodyModule {
	s.DepartName = &v
	return s
}

func (s *TrainOrderListQueryResponseBodyModule) SetGmtCreate(v string) *TrainOrderListQueryResponseBodyModule {
	s.GmtCreate = &v
	return s
}

func (s *TrainOrderListQueryResponseBodyModule) SetGmtModified(v string) *TrainOrderListQueryResponseBodyModule {
	s.GmtModified = &v
	return s
}

func (s *TrainOrderListQueryResponseBodyModule) SetId(v int64) *TrainOrderListQueryResponseBodyModule {
	s.Id = &v
	return s
}

func (s *TrainOrderListQueryResponseBodyModule) SetInvoice(v *TrainOrderListQueryResponseBodyModuleInvoice) *TrainOrderListQueryResponseBodyModule {
	s.Invoice = v
	return s
}

func (s *TrainOrderListQueryResponseBodyModule) SetPriceInfoList(v []*TrainOrderListQueryResponseBodyModulePriceInfoList) *TrainOrderListQueryResponseBodyModule {
	s.PriceInfoList = v
	return s
}

func (s *TrainOrderListQueryResponseBodyModule) SetProjectCode(v string) *TrainOrderListQueryResponseBodyModule {
	s.ProjectCode = &v
	return s
}

func (s *TrainOrderListQueryResponseBodyModule) SetProjectId(v int64) *TrainOrderListQueryResponseBodyModule {
	s.ProjectId = &v
	return s
}

func (s *TrainOrderListQueryResponseBodyModule) SetProjectTitle(v string) *TrainOrderListQueryResponseBodyModule {
	s.ProjectTitle = &v
	return s
}

func (s *TrainOrderListQueryResponseBodyModule) SetRiderName(v string) *TrainOrderListQueryResponseBodyModule {
	s.RiderName = &v
	return s
}

func (s *TrainOrderListQueryResponseBodyModule) SetRunTime(v string) *TrainOrderListQueryResponseBodyModule {
	s.RunTime = &v
	return s
}

func (s *TrainOrderListQueryResponseBodyModule) SetSeatType(v string) *TrainOrderListQueryResponseBodyModule {
	s.SeatType = &v
	return s
}

func (s *TrainOrderListQueryResponseBodyModule) SetStatus(v int32) *TrainOrderListQueryResponseBodyModule {
	s.Status = &v
	return s
}

func (s *TrainOrderListQueryResponseBodyModule) SetThirdPartBusinessId(v string) *TrainOrderListQueryResponseBodyModule {
	s.ThirdPartBusinessId = &v
	return s
}

func (s *TrainOrderListQueryResponseBodyModule) SetThirdPartProjectId(v string) *TrainOrderListQueryResponseBodyModule {
	s.ThirdPartProjectId = &v
	return s
}

func (s *TrainOrderListQueryResponseBodyModule) SetThirdpartApplyId(v string) *TrainOrderListQueryResponseBodyModule {
	s.ThirdpartApplyId = &v
	return s
}

func (s *TrainOrderListQueryResponseBodyModule) SetThirdpartItineraryId(v string) *TrainOrderListQueryResponseBodyModule {
	s.ThirdpartItineraryId = &v
	return s
}

func (s *TrainOrderListQueryResponseBodyModule) SetTicketCount(v int32) *TrainOrderListQueryResponseBodyModule {
	s.TicketCount = &v
	return s
}

func (s *TrainOrderListQueryResponseBodyModule) SetTicketNo12306(v string) *TrainOrderListQueryResponseBodyModule {
	s.TicketNo12306 = &v
	return s
}

func (s *TrainOrderListQueryResponseBodyModule) SetTrainNumber(v string) *TrainOrderListQueryResponseBodyModule {
	s.TrainNumber = &v
	return s
}

func (s *TrainOrderListQueryResponseBodyModule) SetTrainType(v string) *TrainOrderListQueryResponseBodyModule {
	s.TrainType = &v
	return s
}

func (s *TrainOrderListQueryResponseBodyModule) SetUserAffiliateList(v []*TrainOrderListQueryResponseBodyModuleUserAffiliateList) *TrainOrderListQueryResponseBodyModule {
	s.UserAffiliateList = v
	return s
}

func (s *TrainOrderListQueryResponseBodyModule) SetUserId(v string) *TrainOrderListQueryResponseBodyModule {
	s.UserId = &v
	return s
}

func (s *TrainOrderListQueryResponseBodyModule) SetUserName(v string) *TrainOrderListQueryResponseBodyModule {
	s.UserName = &v
	return s
}

func (s *TrainOrderListQueryResponseBodyModule) Validate() error {
	return dara.Validate(s)
}

type TrainOrderListQueryResponseBodyModuleCostCenter struct {
	CorpId *string `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
	// example:
	//
	// 11643
	Id   *int64  `json:"id,omitempty" xml:"id,omitempty"`
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// CS112234
	Number *string `json:"number,omitempty" xml:"number,omitempty"`
}

func (s TrainOrderListQueryResponseBodyModuleCostCenter) String() string {
	return dara.Prettify(s)
}

func (s TrainOrderListQueryResponseBodyModuleCostCenter) GoString() string {
	return s.String()
}

func (s *TrainOrderListQueryResponseBodyModuleCostCenter) GetCorpId() *string {
	return s.CorpId
}

func (s *TrainOrderListQueryResponseBodyModuleCostCenter) GetId() *int64 {
	return s.Id
}

func (s *TrainOrderListQueryResponseBodyModuleCostCenter) GetName() *string {
	return s.Name
}

func (s *TrainOrderListQueryResponseBodyModuleCostCenter) GetNumber() *string {
	return s.Number
}

func (s *TrainOrderListQueryResponseBodyModuleCostCenter) SetCorpId(v string) *TrainOrderListQueryResponseBodyModuleCostCenter {
	s.CorpId = &v
	return s
}

func (s *TrainOrderListQueryResponseBodyModuleCostCenter) SetId(v int64) *TrainOrderListQueryResponseBodyModuleCostCenter {
	s.Id = &v
	return s
}

func (s *TrainOrderListQueryResponseBodyModuleCostCenter) SetName(v string) *TrainOrderListQueryResponseBodyModuleCostCenter {
	s.Name = &v
	return s
}

func (s *TrainOrderListQueryResponseBodyModuleCostCenter) SetNumber(v string) *TrainOrderListQueryResponseBodyModuleCostCenter {
	s.Number = &v
	return s
}

func (s *TrainOrderListQueryResponseBodyModuleCostCenter) Validate() error {
	return dara.Validate(s)
}

type TrainOrderListQueryResponseBodyModuleInvoice struct {
	// example:
	//
	// 11324
	Id    *int64  `json:"id,omitempty" xml:"id,omitempty"`
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s TrainOrderListQueryResponseBodyModuleInvoice) String() string {
	return dara.Prettify(s)
}

func (s TrainOrderListQueryResponseBodyModuleInvoice) GoString() string {
	return s.String()
}

func (s *TrainOrderListQueryResponseBodyModuleInvoice) GetId() *int64 {
	return s.Id
}

func (s *TrainOrderListQueryResponseBodyModuleInvoice) GetTitle() *string {
	return s.Title
}

func (s *TrainOrderListQueryResponseBodyModuleInvoice) SetId(v int64) *TrainOrderListQueryResponseBodyModuleInvoice {
	s.Id = &v
	return s
}

func (s *TrainOrderListQueryResponseBodyModuleInvoice) SetTitle(v string) *TrainOrderListQueryResponseBodyModuleInvoice {
	s.Title = &v
	return s
}

func (s *TrainOrderListQueryResponseBodyModuleInvoice) Validate() error {
	return dara.Validate(s)
}

type TrainOrderListQueryResponseBodyModulePriceInfoList struct {
	// example:
	//
	// 1
	CategoryCode *int32 `json:"category_code,omitempty" xml:"category_code,omitempty"`
	// example:
	//
	// 1
	CategoryType *int32  `json:"category_type,omitempty" xml:"category_type,omitempty"`
	EndCity      *string `json:"end_city,omitempty" xml:"end_city,omitempty"`
	// example:
	//
	// 2022-05-15T22:27Z
	EndTime *string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// example:
	//
	// 2022-05-15T22:27Z
	GmtCreate *string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// example:
	//
	// CS-663
	OriginalTrainNo *string `json:"original_train_no,omitempty" xml:"original_train_no,omitempty"`
	PassengerName   *string `json:"passenger_name,omitempty" xml:"passenger_name,omitempty"`
	// example:
	//
	// 1
	PayType *int32 `json:"pay_type,omitempty" xml:"pay_type,omitempty"`
	// example:
	//
	// 100
	Price     *float64 `json:"price,omitempty" xml:"price,omitempty"`
	SeatType  *string  `json:"seat_type,omitempty" xml:"seat_type,omitempty"`
	StartCity *string  `json:"start_city,omitempty" xml:"start_city,omitempty"`
	// example:
	//
	// 2022-05-15T22:27Z
	StartTime *string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// example:
	//
	// cs1165734212
	TradeId *string `json:"trade_id,omitempty" xml:"trade_id,omitempty"`
	// example:
	//
	// Z1521
	TrainNo *string `json:"train_no,omitempty" xml:"train_no,omitempty"`
	// example:
	//
	// 1
	Type *int32 `json:"type,omitempty" xml:"type,omitempty"`
}

func (s TrainOrderListQueryResponseBodyModulePriceInfoList) String() string {
	return dara.Prettify(s)
}

func (s TrainOrderListQueryResponseBodyModulePriceInfoList) GoString() string {
	return s.String()
}

func (s *TrainOrderListQueryResponseBodyModulePriceInfoList) GetCategoryCode() *int32 {
	return s.CategoryCode
}

func (s *TrainOrderListQueryResponseBodyModulePriceInfoList) GetCategoryType() *int32 {
	return s.CategoryType
}

func (s *TrainOrderListQueryResponseBodyModulePriceInfoList) GetEndCity() *string {
	return s.EndCity
}

func (s *TrainOrderListQueryResponseBodyModulePriceInfoList) GetEndTime() *string {
	return s.EndTime
}

func (s *TrainOrderListQueryResponseBodyModulePriceInfoList) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *TrainOrderListQueryResponseBodyModulePriceInfoList) GetOriginalTrainNo() *string {
	return s.OriginalTrainNo
}

func (s *TrainOrderListQueryResponseBodyModulePriceInfoList) GetPassengerName() *string {
	return s.PassengerName
}

func (s *TrainOrderListQueryResponseBodyModulePriceInfoList) GetPayType() *int32 {
	return s.PayType
}

func (s *TrainOrderListQueryResponseBodyModulePriceInfoList) GetPrice() *float64 {
	return s.Price
}

func (s *TrainOrderListQueryResponseBodyModulePriceInfoList) GetSeatType() *string {
	return s.SeatType
}

func (s *TrainOrderListQueryResponseBodyModulePriceInfoList) GetStartCity() *string {
	return s.StartCity
}

func (s *TrainOrderListQueryResponseBodyModulePriceInfoList) GetStartTime() *string {
	return s.StartTime
}

func (s *TrainOrderListQueryResponseBodyModulePriceInfoList) GetTradeId() *string {
	return s.TradeId
}

func (s *TrainOrderListQueryResponseBodyModulePriceInfoList) GetTrainNo() *string {
	return s.TrainNo
}

func (s *TrainOrderListQueryResponseBodyModulePriceInfoList) GetType() *int32 {
	return s.Type
}

func (s *TrainOrderListQueryResponseBodyModulePriceInfoList) SetCategoryCode(v int32) *TrainOrderListQueryResponseBodyModulePriceInfoList {
	s.CategoryCode = &v
	return s
}

func (s *TrainOrderListQueryResponseBodyModulePriceInfoList) SetCategoryType(v int32) *TrainOrderListQueryResponseBodyModulePriceInfoList {
	s.CategoryType = &v
	return s
}

func (s *TrainOrderListQueryResponseBodyModulePriceInfoList) SetEndCity(v string) *TrainOrderListQueryResponseBodyModulePriceInfoList {
	s.EndCity = &v
	return s
}

func (s *TrainOrderListQueryResponseBodyModulePriceInfoList) SetEndTime(v string) *TrainOrderListQueryResponseBodyModulePriceInfoList {
	s.EndTime = &v
	return s
}

func (s *TrainOrderListQueryResponseBodyModulePriceInfoList) SetGmtCreate(v string) *TrainOrderListQueryResponseBodyModulePriceInfoList {
	s.GmtCreate = &v
	return s
}

func (s *TrainOrderListQueryResponseBodyModulePriceInfoList) SetOriginalTrainNo(v string) *TrainOrderListQueryResponseBodyModulePriceInfoList {
	s.OriginalTrainNo = &v
	return s
}

func (s *TrainOrderListQueryResponseBodyModulePriceInfoList) SetPassengerName(v string) *TrainOrderListQueryResponseBodyModulePriceInfoList {
	s.PassengerName = &v
	return s
}

func (s *TrainOrderListQueryResponseBodyModulePriceInfoList) SetPayType(v int32) *TrainOrderListQueryResponseBodyModulePriceInfoList {
	s.PayType = &v
	return s
}

func (s *TrainOrderListQueryResponseBodyModulePriceInfoList) SetPrice(v float64) *TrainOrderListQueryResponseBodyModulePriceInfoList {
	s.Price = &v
	return s
}

func (s *TrainOrderListQueryResponseBodyModulePriceInfoList) SetSeatType(v string) *TrainOrderListQueryResponseBodyModulePriceInfoList {
	s.SeatType = &v
	return s
}

func (s *TrainOrderListQueryResponseBodyModulePriceInfoList) SetStartCity(v string) *TrainOrderListQueryResponseBodyModulePriceInfoList {
	s.StartCity = &v
	return s
}

func (s *TrainOrderListQueryResponseBodyModulePriceInfoList) SetStartTime(v string) *TrainOrderListQueryResponseBodyModulePriceInfoList {
	s.StartTime = &v
	return s
}

func (s *TrainOrderListQueryResponseBodyModulePriceInfoList) SetTradeId(v string) *TrainOrderListQueryResponseBodyModulePriceInfoList {
	s.TradeId = &v
	return s
}

func (s *TrainOrderListQueryResponseBodyModulePriceInfoList) SetTrainNo(v string) *TrainOrderListQueryResponseBodyModulePriceInfoList {
	s.TrainNo = &v
	return s
}

func (s *TrainOrderListQueryResponseBodyModulePriceInfoList) SetType(v int32) *TrainOrderListQueryResponseBodyModulePriceInfoList {
	s.Type = &v
	return s
}

func (s *TrainOrderListQueryResponseBodyModulePriceInfoList) Validate() error {
	return dara.Validate(s)
}

type TrainOrderListQueryResponseBodyModuleUserAffiliateList struct {
	UserId   *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	UserName *string `json:"user_name,omitempty" xml:"user_name,omitempty"`
}

func (s TrainOrderListQueryResponseBodyModuleUserAffiliateList) String() string {
	return dara.Prettify(s)
}

func (s TrainOrderListQueryResponseBodyModuleUserAffiliateList) GoString() string {
	return s.String()
}

func (s *TrainOrderListQueryResponseBodyModuleUserAffiliateList) GetUserId() *string {
	return s.UserId
}

func (s *TrainOrderListQueryResponseBodyModuleUserAffiliateList) GetUserName() *string {
	return s.UserName
}

func (s *TrainOrderListQueryResponseBodyModuleUserAffiliateList) SetUserId(v string) *TrainOrderListQueryResponseBodyModuleUserAffiliateList {
	s.UserId = &v
	return s
}

func (s *TrainOrderListQueryResponseBodyModuleUserAffiliateList) SetUserName(v string) *TrainOrderListQueryResponseBodyModuleUserAffiliateList {
	s.UserName = &v
	return s
}

func (s *TrainOrderListQueryResponseBodyModuleUserAffiliateList) Validate() error {
	return dara.Validate(s)
}

type TrainOrderListQueryResponseBodyPageInfo struct {
	// example:
	//
	// 1
	Page *int32 `json:"page,omitempty" xml:"page,omitempty"`
	// example:
	//
	// 50
	PageSize *int32 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// example:
	//
	// 100
	TotalNumber *int32 `json:"total_number,omitempty" xml:"total_number,omitempty"`
}

func (s TrainOrderListQueryResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s TrainOrderListQueryResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *TrainOrderListQueryResponseBodyPageInfo) GetPage() *int32 {
	return s.Page
}

func (s *TrainOrderListQueryResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *TrainOrderListQueryResponseBodyPageInfo) GetTotalNumber() *int32 {
	return s.TotalNumber
}

func (s *TrainOrderListQueryResponseBodyPageInfo) SetPage(v int32) *TrainOrderListQueryResponseBodyPageInfo {
	s.Page = &v
	return s
}

func (s *TrainOrderListQueryResponseBodyPageInfo) SetPageSize(v int32) *TrainOrderListQueryResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *TrainOrderListQueryResponseBodyPageInfo) SetTotalNumber(v int32) *TrainOrderListQueryResponseBodyPageInfo {
	s.TotalNumber = &v
	return s
}

func (s *TrainOrderListQueryResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}
