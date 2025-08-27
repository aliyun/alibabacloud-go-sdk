// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightOrderListQueryV2ResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *FlightOrderListQueryV2ResponseBody
	GetCode() *string
	SetMessage(v string) *FlightOrderListQueryV2ResponseBody
	GetMessage() *string
	SetModule(v []*FlightOrderListQueryV2ResponseBodyModule) *FlightOrderListQueryV2ResponseBody
	GetModule() []*FlightOrderListQueryV2ResponseBodyModule
	SetPageInfo(v *FlightOrderListQueryV2ResponseBodyPageInfo) *FlightOrderListQueryV2ResponseBody
	GetPageInfo() *FlightOrderListQueryV2ResponseBodyPageInfo
	SetRequestId(v string) *FlightOrderListQueryV2ResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *FlightOrderListQueryV2ResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *FlightOrderListQueryV2ResponseBody
	GetTraceId() *string
}

type FlightOrderListQueryV2ResponseBody struct {
	// example:
	//
	// 0
	Code     *string                                     `json:"code,omitempty" xml:"code,omitempty"`
	Message  *string                                     `json:"message,omitempty" xml:"message,omitempty"`
	Module   []*FlightOrderListQueryV2ResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Repeated"`
	PageInfo *FlightOrderListQueryV2ResponseBodyPageInfo `json:"pageInfo,omitempty" xml:"pageInfo,omitempty" type:"Struct"`
	// example:
	//
	// C61ECFF6-****-****-****-D77369043A5F
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 21041ce********056433edbb2
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s FlightOrderListQueryV2ResponseBody) String() string {
	return dara.Prettify(s)
}

func (s FlightOrderListQueryV2ResponseBody) GoString() string {
	return s.String()
}

func (s *FlightOrderListQueryV2ResponseBody) GetCode() *string {
	return s.Code
}

func (s *FlightOrderListQueryV2ResponseBody) GetMessage() *string {
	return s.Message
}

func (s *FlightOrderListQueryV2ResponseBody) GetModule() []*FlightOrderListQueryV2ResponseBodyModule {
	return s.Module
}

func (s *FlightOrderListQueryV2ResponseBody) GetPageInfo() *FlightOrderListQueryV2ResponseBodyPageInfo {
	return s.PageInfo
}

func (s *FlightOrderListQueryV2ResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *FlightOrderListQueryV2ResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *FlightOrderListQueryV2ResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *FlightOrderListQueryV2ResponseBody) SetCode(v string) *FlightOrderListQueryV2ResponseBody {
	s.Code = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBody) SetMessage(v string) *FlightOrderListQueryV2ResponseBody {
	s.Message = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBody) SetModule(v []*FlightOrderListQueryV2ResponseBodyModule) *FlightOrderListQueryV2ResponseBody {
	s.Module = v
	return s
}

func (s *FlightOrderListQueryV2ResponseBody) SetPageInfo(v *FlightOrderListQueryV2ResponseBodyPageInfo) *FlightOrderListQueryV2ResponseBody {
	s.PageInfo = v
	return s
}

func (s *FlightOrderListQueryV2ResponseBody) SetRequestId(v string) *FlightOrderListQueryV2ResponseBody {
	s.RequestId = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBody) SetSuccess(v bool) *FlightOrderListQueryV2ResponseBody {
	s.Success = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBody) SetTraceId(v string) *FlightOrderListQueryV2ResponseBody {
	s.TraceId = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBody) Validate() error {
	return dara.Validate(s)
}

type FlightOrderListQueryV2ResponseBodyModule struct {
	Approve                *FlightOrderListQueryV2ResponseBodyModuleApprove                  `json:"approve,omitempty" xml:"approve,omitempty" type:"Struct"`
	CorpId                 *string                                                           `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
	CorpName               *string                                                           `json:"corp_name,omitempty" xml:"corp_name,omitempty"`
	DepartId               *string                                                           `json:"depart_id,omitempty" xml:"depart_id,omitempty"`
	DepartName             *string                                                           `json:"depart_name,omitempty" xml:"depart_name,omitempty"`
	FlightOrderTicketList  []*FlightOrderListQueryV2ResponseBodyModuleFlightOrderTicketList  `json:"flight_order_ticket_list,omitempty" xml:"flight_order_ticket_list,omitempty" type:"Repeated"`
	FlightOrderUserFeeList []*FlightOrderListQueryV2ResponseBodyModuleFlightOrderUserFeeList `json:"flight_order_user_fee_list,omitempty" xml:"flight_order_user_fee_list,omitempty" type:"Repeated"`
	FlightRefundApplyList  []*FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyList  `json:"flight_refund_apply_list,omitempty" xml:"flight_refund_apply_list,omitempty" type:"Repeated"`
	FlightReshopApplyList  []*FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyList  `json:"flight_reshop_apply_list,omitempty" xml:"flight_reshop_apply_list,omitempty" type:"Repeated"`
	FlightSegmentList      []*FlightOrderListQueryV2ResponseBodyModuleFlightSegmentList      `json:"flight_segment_list,omitempty" xml:"flight_segment_list,omitempty" type:"Repeated"`
	// example:
	//
	// yyyy-MM-dd HH:mm:ss
	GmtCreate *string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// example:
	//
	// yyyy-MM-dd HH:mm:ss
	GmtModified *string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// example:
	//
	// 200042
	Id             *string                                                   `json:"id,omitempty" xml:"id,omitempty"`
	InsureInfoList []*FlightOrderListQueryV2ResponseBodyModuleInsureInfoList `json:"insure_info_list,omitempty" xml:"insure_info_list,omitempty" type:"Repeated"`
	// example:
	//
	// true
	MixPay             *bool    `json:"mix_pay,omitempty" xml:"mix_pay,omitempty"`
	OrderReserveAmount *float64 `json:"order_reserve_amount,omitempty" xml:"order_reserve_amount,omitempty"`
	// example:
	//
	// 1
	PassengerCount *int32 `json:"passenger_count,omitempty" xml:"passenger_count,omitempty"`
	// example:
	//
	// yyyy-MM-dd HH:mm:ss
	PayTime       *string                                                  `json:"pay_time,omitempty" xml:"pay_time,omitempty"`
	PriceInfoList []*FlightOrderListQueryV2ResponseBodyModulePriceInfoList `json:"price_info_list,omitempty" xml:"price_info_list,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	Status   *int32  `json:"status,omitempty" xml:"status,omitempty"`
	Supplier *string `json:"supplier,omitempty" xml:"supplier,omitempty"`
	// example:
	//
	// cs9897766
	ThirdpartItineraryId      []*string `json:"thirdpart_itinerary_id,omitempty" xml:"thirdpart_itinerary_id,omitempty" type:"Repeated"`
	TicketCorpReserveAmount   *float64  `json:"ticket_corp_reserve_amount,omitempty" xml:"ticket_corp_reserve_amount,omitempty"`
	TicketPersonReserveAmount *float64  `json:"ticket_person_reserve_amount,omitempty" xml:"ticket_person_reserve_amount,omitempty"`
	// example:
	//
	// 1
	TripMode *int32 `json:"trip_mode,omitempty" xml:"trip_mode,omitempty"`
	// example:
	//
	// 0
	TripType          *int32                                                       `json:"trip_type,omitempty" xml:"trip_type,omitempty"`
	UserAffiliateList []*FlightOrderListQueryV2ResponseBodyModuleUserAffiliateList `json:"user_affiliate_list,omitempty" xml:"user_affiliate_list,omitempty" type:"Repeated"`
	UserId            *string                                                      `json:"user_id,omitempty" xml:"user_id,omitempty"`
	UserName          *string                                                      `json:"user_name,omitempty" xml:"user_name,omitempty"`
}

func (s FlightOrderListQueryV2ResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s FlightOrderListQueryV2ResponseBodyModule) GoString() string {
	return s.String()
}

func (s *FlightOrderListQueryV2ResponseBodyModule) GetApprove() *FlightOrderListQueryV2ResponseBodyModuleApprove {
	return s.Approve
}

func (s *FlightOrderListQueryV2ResponseBodyModule) GetCorpId() *string {
	return s.CorpId
}

func (s *FlightOrderListQueryV2ResponseBodyModule) GetCorpName() *string {
	return s.CorpName
}

func (s *FlightOrderListQueryV2ResponseBodyModule) GetDepartId() *string {
	return s.DepartId
}

func (s *FlightOrderListQueryV2ResponseBodyModule) GetDepartName() *string {
	return s.DepartName
}

func (s *FlightOrderListQueryV2ResponseBodyModule) GetFlightOrderTicketList() []*FlightOrderListQueryV2ResponseBodyModuleFlightOrderTicketList {
	return s.FlightOrderTicketList
}

func (s *FlightOrderListQueryV2ResponseBodyModule) GetFlightOrderUserFeeList() []*FlightOrderListQueryV2ResponseBodyModuleFlightOrderUserFeeList {
	return s.FlightOrderUserFeeList
}

func (s *FlightOrderListQueryV2ResponseBodyModule) GetFlightRefundApplyList() []*FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyList {
	return s.FlightRefundApplyList
}

func (s *FlightOrderListQueryV2ResponseBodyModule) GetFlightReshopApplyList() []*FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyList {
	return s.FlightReshopApplyList
}

func (s *FlightOrderListQueryV2ResponseBodyModule) GetFlightSegmentList() []*FlightOrderListQueryV2ResponseBodyModuleFlightSegmentList {
	return s.FlightSegmentList
}

func (s *FlightOrderListQueryV2ResponseBodyModule) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *FlightOrderListQueryV2ResponseBodyModule) GetGmtModified() *string {
	return s.GmtModified
}

func (s *FlightOrderListQueryV2ResponseBodyModule) GetId() *string {
	return s.Id
}

func (s *FlightOrderListQueryV2ResponseBodyModule) GetInsureInfoList() []*FlightOrderListQueryV2ResponseBodyModuleInsureInfoList {
	return s.InsureInfoList
}

func (s *FlightOrderListQueryV2ResponseBodyModule) GetMixPay() *bool {
	return s.MixPay
}

func (s *FlightOrderListQueryV2ResponseBodyModule) GetOrderReserveAmount() *float64 {
	return s.OrderReserveAmount
}

func (s *FlightOrderListQueryV2ResponseBodyModule) GetPassengerCount() *int32 {
	return s.PassengerCount
}

func (s *FlightOrderListQueryV2ResponseBodyModule) GetPayTime() *string {
	return s.PayTime
}

func (s *FlightOrderListQueryV2ResponseBodyModule) GetPriceInfoList() []*FlightOrderListQueryV2ResponseBodyModulePriceInfoList {
	return s.PriceInfoList
}

func (s *FlightOrderListQueryV2ResponseBodyModule) GetStatus() *int32 {
	return s.Status
}

func (s *FlightOrderListQueryV2ResponseBodyModule) GetSupplier() *string {
	return s.Supplier
}

func (s *FlightOrderListQueryV2ResponseBodyModule) GetThirdpartItineraryId() []*string {
	return s.ThirdpartItineraryId
}

func (s *FlightOrderListQueryV2ResponseBodyModule) GetTicketCorpReserveAmount() *float64 {
	return s.TicketCorpReserveAmount
}

func (s *FlightOrderListQueryV2ResponseBodyModule) GetTicketPersonReserveAmount() *float64 {
	return s.TicketPersonReserveAmount
}

func (s *FlightOrderListQueryV2ResponseBodyModule) GetTripMode() *int32 {
	return s.TripMode
}

func (s *FlightOrderListQueryV2ResponseBodyModule) GetTripType() *int32 {
	return s.TripType
}

func (s *FlightOrderListQueryV2ResponseBodyModule) GetUserAffiliateList() []*FlightOrderListQueryV2ResponseBodyModuleUserAffiliateList {
	return s.UserAffiliateList
}

func (s *FlightOrderListQueryV2ResponseBodyModule) GetUserId() *string {
	return s.UserId
}

func (s *FlightOrderListQueryV2ResponseBodyModule) GetUserName() *string {
	return s.UserName
}

func (s *FlightOrderListQueryV2ResponseBodyModule) SetApprove(v *FlightOrderListQueryV2ResponseBodyModuleApprove) *FlightOrderListQueryV2ResponseBodyModule {
	s.Approve = v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModule) SetCorpId(v string) *FlightOrderListQueryV2ResponseBodyModule {
	s.CorpId = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModule) SetCorpName(v string) *FlightOrderListQueryV2ResponseBodyModule {
	s.CorpName = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModule) SetDepartId(v string) *FlightOrderListQueryV2ResponseBodyModule {
	s.DepartId = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModule) SetDepartName(v string) *FlightOrderListQueryV2ResponseBodyModule {
	s.DepartName = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModule) SetFlightOrderTicketList(v []*FlightOrderListQueryV2ResponseBodyModuleFlightOrderTicketList) *FlightOrderListQueryV2ResponseBodyModule {
	s.FlightOrderTicketList = v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModule) SetFlightOrderUserFeeList(v []*FlightOrderListQueryV2ResponseBodyModuleFlightOrderUserFeeList) *FlightOrderListQueryV2ResponseBodyModule {
	s.FlightOrderUserFeeList = v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModule) SetFlightRefundApplyList(v []*FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyList) *FlightOrderListQueryV2ResponseBodyModule {
	s.FlightRefundApplyList = v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModule) SetFlightReshopApplyList(v []*FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyList) *FlightOrderListQueryV2ResponseBodyModule {
	s.FlightReshopApplyList = v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModule) SetFlightSegmentList(v []*FlightOrderListQueryV2ResponseBodyModuleFlightSegmentList) *FlightOrderListQueryV2ResponseBodyModule {
	s.FlightSegmentList = v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModule) SetGmtCreate(v string) *FlightOrderListQueryV2ResponseBodyModule {
	s.GmtCreate = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModule) SetGmtModified(v string) *FlightOrderListQueryV2ResponseBodyModule {
	s.GmtModified = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModule) SetId(v string) *FlightOrderListQueryV2ResponseBodyModule {
	s.Id = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModule) SetInsureInfoList(v []*FlightOrderListQueryV2ResponseBodyModuleInsureInfoList) *FlightOrderListQueryV2ResponseBodyModule {
	s.InsureInfoList = v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModule) SetMixPay(v bool) *FlightOrderListQueryV2ResponseBodyModule {
	s.MixPay = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModule) SetOrderReserveAmount(v float64) *FlightOrderListQueryV2ResponseBodyModule {
	s.OrderReserveAmount = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModule) SetPassengerCount(v int32) *FlightOrderListQueryV2ResponseBodyModule {
	s.PassengerCount = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModule) SetPayTime(v string) *FlightOrderListQueryV2ResponseBodyModule {
	s.PayTime = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModule) SetPriceInfoList(v []*FlightOrderListQueryV2ResponseBodyModulePriceInfoList) *FlightOrderListQueryV2ResponseBodyModule {
	s.PriceInfoList = v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModule) SetStatus(v int32) *FlightOrderListQueryV2ResponseBodyModule {
	s.Status = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModule) SetSupplier(v string) *FlightOrderListQueryV2ResponseBodyModule {
	s.Supplier = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModule) SetThirdpartItineraryId(v []*string) *FlightOrderListQueryV2ResponseBodyModule {
	s.ThirdpartItineraryId = v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModule) SetTicketCorpReserveAmount(v float64) *FlightOrderListQueryV2ResponseBodyModule {
	s.TicketCorpReserveAmount = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModule) SetTicketPersonReserveAmount(v float64) *FlightOrderListQueryV2ResponseBodyModule {
	s.TicketPersonReserveAmount = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModule) SetTripMode(v int32) *FlightOrderListQueryV2ResponseBodyModule {
	s.TripMode = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModule) SetTripType(v int32) *FlightOrderListQueryV2ResponseBodyModule {
	s.TripType = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModule) SetUserAffiliateList(v []*FlightOrderListQueryV2ResponseBodyModuleUserAffiliateList) *FlightOrderListQueryV2ResponseBodyModule {
	s.UserAffiliateList = v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModule) SetUserId(v string) *FlightOrderListQueryV2ResponseBodyModule {
	s.UserId = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModule) SetUserName(v string) *FlightOrderListQueryV2ResponseBodyModule {
	s.UserName = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModule) Validate() error {
	return dara.Validate(s)
}

type FlightOrderListQueryV2ResponseBodyModuleApprove struct {
	// example:
	//
	// test1234
	ApproveId  *int64  `json:"approve_id,omitempty" xml:"approve_id,omitempty"`
	BtripTitle *string `json:"btrip_title,omitempty" xml:"btrip_title,omitempty"`
	// example:
	//
	// 1233544
	ExceedApproveId *string `json:"exceed_approve_id,omitempty" xml:"exceed_approve_id,omitempty"`
	// example:
	//
	// 100231431
	ThirdpartApproveId *string `json:"thirdpart_approve_id,omitempty" xml:"thirdpart_approve_id,omitempty"`
	// example:
	//
	// test123
	ThirdpartExceedApproveId *string `json:"thirdpart_exceed_approve_id,omitempty" xml:"thirdpart_exceed_approve_id,omitempty"`
}

func (s FlightOrderListQueryV2ResponseBodyModuleApprove) String() string {
	return dara.Prettify(s)
}

func (s FlightOrderListQueryV2ResponseBodyModuleApprove) GoString() string {
	return s.String()
}

func (s *FlightOrderListQueryV2ResponseBodyModuleApprove) GetApproveId() *int64 {
	return s.ApproveId
}

func (s *FlightOrderListQueryV2ResponseBodyModuleApprove) GetBtripTitle() *string {
	return s.BtripTitle
}

func (s *FlightOrderListQueryV2ResponseBodyModuleApprove) GetExceedApproveId() *string {
	return s.ExceedApproveId
}

func (s *FlightOrderListQueryV2ResponseBodyModuleApprove) GetThirdpartApproveId() *string {
	return s.ThirdpartApproveId
}

func (s *FlightOrderListQueryV2ResponseBodyModuleApprove) GetThirdpartExceedApproveId() *string {
	return s.ThirdpartExceedApproveId
}

func (s *FlightOrderListQueryV2ResponseBodyModuleApprove) SetApproveId(v int64) *FlightOrderListQueryV2ResponseBodyModuleApprove {
	s.ApproveId = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleApprove) SetBtripTitle(v string) *FlightOrderListQueryV2ResponseBodyModuleApprove {
	s.BtripTitle = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleApprove) SetExceedApproveId(v string) *FlightOrderListQueryV2ResponseBodyModuleApprove {
	s.ExceedApproveId = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleApprove) SetThirdpartApproveId(v string) *FlightOrderListQueryV2ResponseBodyModuleApprove {
	s.ThirdpartApproveId = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleApprove) SetThirdpartExceedApproveId(v string) *FlightOrderListQueryV2ResponseBodyModuleApprove {
	s.ThirdpartExceedApproveId = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleApprove) Validate() error {
	return dara.Validate(s)
}

type FlightOrderListQueryV2ResponseBodyModuleFlightOrderTicketList struct {
	FlightList   []*FlightOrderListQueryV2ResponseBodyModuleFlightOrderTicketListFlightList `json:"flight_list,omitempty" xml:"flight_list,omitempty" type:"Repeated"`
	TicketNoList []*string                                                                  `json:"ticket_no_list,omitempty" xml:"ticket_no_list,omitempty" type:"Repeated"`
	UserId       *string                                                                    `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

func (s FlightOrderListQueryV2ResponseBodyModuleFlightOrderTicketList) String() string {
	return dara.Prettify(s)
}

func (s FlightOrderListQueryV2ResponseBodyModuleFlightOrderTicketList) GoString() string {
	return s.String()
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightOrderTicketList) GetFlightList() []*FlightOrderListQueryV2ResponseBodyModuleFlightOrderTicketListFlightList {
	return s.FlightList
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightOrderTicketList) GetTicketNoList() []*string {
	return s.TicketNoList
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightOrderTicketList) GetUserId() *string {
	return s.UserId
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightOrderTicketList) SetFlightList(v []*FlightOrderListQueryV2ResponseBodyModuleFlightOrderTicketListFlightList) *FlightOrderListQueryV2ResponseBodyModuleFlightOrderTicketList {
	s.FlightList = v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightOrderTicketList) SetTicketNoList(v []*string) *FlightOrderListQueryV2ResponseBodyModuleFlightOrderTicketList {
	s.TicketNoList = v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightOrderTicketList) SetUserId(v string) *FlightOrderListQueryV2ResponseBodyModuleFlightOrderTicketList {
	s.UserId = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightOrderTicketList) Validate() error {
	return dara.Validate(s)
}

type FlightOrderListQueryV2ResponseBodyModuleFlightOrderTicketListFlightList struct {
	// example:
	//
	// 2022-07-20T10:40Z
	ArrTime *string `json:"arr_time,omitempty" xml:"arr_time,omitempty"`
	// example:
	//
	// V
	Cabin *string `json:"cabin,omitempty" xml:"cabin,omitempty"`
	// example:
	//
	// Y
	CabinClass     *string `json:"cabin_class,omitempty" xml:"cabin_class,omitempty"`
	CabinClassName *string `json:"cabin_class_name,omitempty" xml:"cabin_class_name,omitempty"`
	// example:
	//
	// 2022-07-20T10:40Z
	DepTime *string `json:"dep_time,omitempty" xml:"dep_time,omitempty"`
	// example:
	//
	// CA8572
	FlightNo *string `json:"flight_no,omitempty" xml:"flight_no,omitempty"`
}

func (s FlightOrderListQueryV2ResponseBodyModuleFlightOrderTicketListFlightList) String() string {
	return dara.Prettify(s)
}

func (s FlightOrderListQueryV2ResponseBodyModuleFlightOrderTicketListFlightList) GoString() string {
	return s.String()
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightOrderTicketListFlightList) GetArrTime() *string {
	return s.ArrTime
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightOrderTicketListFlightList) GetCabin() *string {
	return s.Cabin
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightOrderTicketListFlightList) GetCabinClass() *string {
	return s.CabinClass
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightOrderTicketListFlightList) GetCabinClassName() *string {
	return s.CabinClassName
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightOrderTicketListFlightList) GetDepTime() *string {
	return s.DepTime
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightOrderTicketListFlightList) GetFlightNo() *string {
	return s.FlightNo
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightOrderTicketListFlightList) SetArrTime(v string) *FlightOrderListQueryV2ResponseBodyModuleFlightOrderTicketListFlightList {
	s.ArrTime = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightOrderTicketListFlightList) SetCabin(v string) *FlightOrderListQueryV2ResponseBodyModuleFlightOrderTicketListFlightList {
	s.Cabin = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightOrderTicketListFlightList) SetCabinClass(v string) *FlightOrderListQueryV2ResponseBodyModuleFlightOrderTicketListFlightList {
	s.CabinClass = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightOrderTicketListFlightList) SetCabinClassName(v string) *FlightOrderListQueryV2ResponseBodyModuleFlightOrderTicketListFlightList {
	s.CabinClassName = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightOrderTicketListFlightList) SetDepTime(v string) *FlightOrderListQueryV2ResponseBodyModuleFlightOrderTicketListFlightList {
	s.DepTime = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightOrderTicketListFlightList) SetFlightNo(v string) *FlightOrderListQueryV2ResponseBodyModuleFlightOrderTicketListFlightList {
	s.FlightNo = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightOrderTicketListFlightList) Validate() error {
	return dara.Validate(s)
}

type FlightOrderListQueryV2ResponseBodyModuleFlightOrderUserFeeList struct {
	BuildFee        *float64 `json:"build_fee,omitempty" xml:"build_fee,omitempty"`
	CorpPayAmount   *float64 `json:"corp_pay_amount,omitempty" xml:"corp_pay_amount,omitempty"`
	OilFee          *float64 `json:"oil_fee,omitempty" xml:"oil_fee,omitempty"`
	PersonPayAmount *float64 `json:"person_pay_amount,omitempty" xml:"person_pay_amount,omitempty"`
	TicketPrice     *float64 `json:"ticket_price,omitempty" xml:"ticket_price,omitempty"`
	UserId          *string  `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

func (s FlightOrderListQueryV2ResponseBodyModuleFlightOrderUserFeeList) String() string {
	return dara.Prettify(s)
}

func (s FlightOrderListQueryV2ResponseBodyModuleFlightOrderUserFeeList) GoString() string {
	return s.String()
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightOrderUserFeeList) GetBuildFee() *float64 {
	return s.BuildFee
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightOrderUserFeeList) GetCorpPayAmount() *float64 {
	return s.CorpPayAmount
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightOrderUserFeeList) GetOilFee() *float64 {
	return s.OilFee
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightOrderUserFeeList) GetPersonPayAmount() *float64 {
	return s.PersonPayAmount
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightOrderUserFeeList) GetTicketPrice() *float64 {
	return s.TicketPrice
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightOrderUserFeeList) GetUserId() *string {
	return s.UserId
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightOrderUserFeeList) SetBuildFee(v float64) *FlightOrderListQueryV2ResponseBodyModuleFlightOrderUserFeeList {
	s.BuildFee = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightOrderUserFeeList) SetCorpPayAmount(v float64) *FlightOrderListQueryV2ResponseBodyModuleFlightOrderUserFeeList {
	s.CorpPayAmount = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightOrderUserFeeList) SetOilFee(v float64) *FlightOrderListQueryV2ResponseBodyModuleFlightOrderUserFeeList {
	s.OilFee = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightOrderUserFeeList) SetPersonPayAmount(v float64) *FlightOrderListQueryV2ResponseBodyModuleFlightOrderUserFeeList {
	s.PersonPayAmount = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightOrderUserFeeList) SetTicketPrice(v float64) *FlightOrderListQueryV2ResponseBodyModuleFlightOrderUserFeeList {
	s.TicketPrice = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightOrderUserFeeList) SetUserId(v string) *FlightOrderListQueryV2ResponseBodyModuleFlightOrderUserFeeList {
	s.UserId = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightOrderUserFeeList) Validate() error {
	return dara.Validate(s)
}

type FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyList struct {
	FlightRefundApplyTicketList []*FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundApplyTicketList `json:"flight_refund_apply_ticket_list,omitempty" xml:"flight_refund_apply_ticket_list,omitempty" type:"Repeated"`
	FlightRefundSegmentList     []*FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundSegmentList     `json:"flight_refund_segment_list,omitempty" xml:"flight_refund_segment_list,omitempty" type:"Repeated"`
	FlightRefundUserFeeList     []*FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundUserFeeList     `json:"flight_refund_user_fee_list,omitempty" xml:"flight_refund_user_fee_list,omitempty" type:"Repeated"`
	// example:
	//
	// 232213
	RefundApplyId *string `json:"refund_apply_id,omitempty" xml:"refund_apply_id,omitempty"`
	// example:
	//
	// 2024122312
	RefundApproveId         *string  `json:"refund_approve_id,omitempty" xml:"refund_approve_id,omitempty"`
	RefundCorpTotalAmount   *float64 `json:"refund_corp_total_amount,omitempty" xml:"refund_corp_total_amount,omitempty"`
	RefundHandFee           *float64 `json:"refund_hand_fee,omitempty" xml:"refund_hand_fee,omitempty"`
	RefundPersonTotalAmount *float64 `json:"refund_person_total_amount,omitempty" xml:"refund_person_total_amount,omitempty"`
	RefundReason            *string  `json:"refund_reason,omitempty" xml:"refund_reason,omitempty"`
	// example:
	//
	// 0
	RefundReasonCode  *string  `json:"refund_reason_code,omitempty" xml:"refund_reason_code,omitempty"`
	RefundTotalAmount *float64 `json:"refund_total_amount,omitempty" xml:"refund_total_amount,omitempty"`
	// example:
	//
	// 232218
	RelateRefundApplyId *string   `json:"relate_refund_apply_id,omitempty" xml:"relate_refund_apply_id,omitempty"`
	UserIdList          []*string `json:"user_id_list,omitempty" xml:"user_id_list,omitempty" type:"Repeated"`
}

func (s FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyList) String() string {
	return dara.Prettify(s)
}

func (s FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyList) GoString() string {
	return s.String()
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyList) GetFlightRefundApplyTicketList() []*FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundApplyTicketList {
	return s.FlightRefundApplyTicketList
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyList) GetFlightRefundSegmentList() []*FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundSegmentList {
	return s.FlightRefundSegmentList
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyList) GetFlightRefundUserFeeList() []*FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundUserFeeList {
	return s.FlightRefundUserFeeList
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyList) GetRefundApplyId() *string {
	return s.RefundApplyId
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyList) GetRefundApproveId() *string {
	return s.RefundApproveId
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyList) GetRefundCorpTotalAmount() *float64 {
	return s.RefundCorpTotalAmount
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyList) GetRefundHandFee() *float64 {
	return s.RefundHandFee
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyList) GetRefundPersonTotalAmount() *float64 {
	return s.RefundPersonTotalAmount
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyList) GetRefundReason() *string {
	return s.RefundReason
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyList) GetRefundReasonCode() *string {
	return s.RefundReasonCode
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyList) GetRefundTotalAmount() *float64 {
	return s.RefundTotalAmount
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyList) GetRelateRefundApplyId() *string {
	return s.RelateRefundApplyId
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyList) GetUserIdList() []*string {
	return s.UserIdList
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyList) SetFlightRefundApplyTicketList(v []*FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundApplyTicketList) *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyList {
	s.FlightRefundApplyTicketList = v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyList) SetFlightRefundSegmentList(v []*FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundSegmentList) *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyList {
	s.FlightRefundSegmentList = v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyList) SetFlightRefundUserFeeList(v []*FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundUserFeeList) *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyList {
	s.FlightRefundUserFeeList = v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyList) SetRefundApplyId(v string) *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyList {
	s.RefundApplyId = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyList) SetRefundApproveId(v string) *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyList {
	s.RefundApproveId = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyList) SetRefundCorpTotalAmount(v float64) *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyList {
	s.RefundCorpTotalAmount = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyList) SetRefundHandFee(v float64) *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyList {
	s.RefundHandFee = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyList) SetRefundPersonTotalAmount(v float64) *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyList {
	s.RefundPersonTotalAmount = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyList) SetRefundReason(v string) *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyList {
	s.RefundReason = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyList) SetRefundReasonCode(v string) *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyList {
	s.RefundReasonCode = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyList) SetRefundTotalAmount(v float64) *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyList {
	s.RefundTotalAmount = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyList) SetRelateRefundApplyId(v string) *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyList {
	s.RelateRefundApplyId = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyList) SetUserIdList(v []*string) *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyList {
	s.UserIdList = v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyList) Validate() error {
	return dara.Validate(s)
}

type FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundApplyTicketList struct {
	FlightList   []*FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundApplyTicketListFlightList `json:"flight_list,omitempty" xml:"flight_list,omitempty" type:"Repeated"`
	TicketNoList []*string                                                                                             `json:"ticket_no_list,omitempty" xml:"ticket_no_list,omitempty" type:"Repeated"`
	UserId       *string                                                                                               `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

func (s FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundApplyTicketList) String() string {
	return dara.Prettify(s)
}

func (s FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundApplyTicketList) GoString() string {
	return s.String()
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundApplyTicketList) GetFlightList() []*FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundApplyTicketListFlightList {
	return s.FlightList
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundApplyTicketList) GetTicketNoList() []*string {
	return s.TicketNoList
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundApplyTicketList) GetUserId() *string {
	return s.UserId
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundApplyTicketList) SetFlightList(v []*FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundApplyTicketListFlightList) *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundApplyTicketList {
	s.FlightList = v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundApplyTicketList) SetTicketNoList(v []*string) *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundApplyTicketList {
	s.TicketNoList = v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundApplyTicketList) SetUserId(v string) *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundApplyTicketList {
	s.UserId = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundApplyTicketList) Validate() error {
	return dara.Validate(s)
}

type FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundApplyTicketListFlightList struct {
	// example:
	//
	// 2022-07-20T10:40Z
	ArrTime *string `json:"arr_time,omitempty" xml:"arr_time,omitempty"`
	// example:
	//
	// V
	Cabin *string `json:"cabin,omitempty" xml:"cabin,omitempty"`
	// example:
	//
	// Y
	CabinClass     *string `json:"cabin_class,omitempty" xml:"cabin_class,omitempty"`
	CabinClassName *string `json:"cabin_class_name,omitempty" xml:"cabin_class_name,omitempty"`
	// example:
	//
	// 2022-07-20T10:40Z
	DepTime *string `json:"dep_time,omitempty" xml:"dep_time,omitempty"`
	// example:
	//
	// HU7052
	FlightNo *string `json:"flight_no,omitempty" xml:"flight_no,omitempty"`
}

func (s FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundApplyTicketListFlightList) String() string {
	return dara.Prettify(s)
}

func (s FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundApplyTicketListFlightList) GoString() string {
	return s.String()
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundApplyTicketListFlightList) GetArrTime() *string {
	return s.ArrTime
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundApplyTicketListFlightList) GetCabin() *string {
	return s.Cabin
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundApplyTicketListFlightList) GetCabinClass() *string {
	return s.CabinClass
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundApplyTicketListFlightList) GetCabinClassName() *string {
	return s.CabinClassName
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundApplyTicketListFlightList) GetDepTime() *string {
	return s.DepTime
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundApplyTicketListFlightList) GetFlightNo() *string {
	return s.FlightNo
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundApplyTicketListFlightList) SetArrTime(v string) *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundApplyTicketListFlightList {
	s.ArrTime = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundApplyTicketListFlightList) SetCabin(v string) *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundApplyTicketListFlightList {
	s.Cabin = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundApplyTicketListFlightList) SetCabinClass(v string) *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundApplyTicketListFlightList {
	s.CabinClass = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundApplyTicketListFlightList) SetCabinClassName(v string) *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundApplyTicketListFlightList {
	s.CabinClassName = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundApplyTicketListFlightList) SetDepTime(v string) *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundApplyTicketListFlightList {
	s.DepTime = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundApplyTicketListFlightList) SetFlightNo(v string) *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundApplyTicketListFlightList {
	s.FlightNo = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundApplyTicketListFlightList) Validate() error {
	return dara.Validate(s)
}

type FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundSegmentList struct {
	// example:
	//
	// CZ
	AirlineCode *string `json:"airline_code,omitempty" xml:"airline_code,omitempty"`
	AirlineName *string `json:"airline_name,omitempty" xml:"airline_name,omitempty"`
	ArrApt      *string `json:"arr_apt,omitempty" xml:"arr_apt,omitempty"`
	// example:
	//
	// PKX
	ArrAptCode *string `json:"arr_apt_code,omitempty" xml:"arr_apt_code,omitempty"`
	ArrCity    *string `json:"arr_city,omitempty" xml:"arr_city,omitempty"`
	// example:
	//
	// BJS
	ArrCityCode *string `json:"arr_city_code,omitempty" xml:"arr_city_code,omitempty"`
	// example:
	//
	// T1
	ArrTerminal *string `json:"arr_terminal,omitempty" xml:"arr_terminal,omitempty"`
	// example:
	//
	// 2022-07-20T10:40Z
	ArrTime *string `json:"arr_time,omitempty" xml:"arr_time,omitempty"`
	DepApt  *string `json:"dep_apt,omitempty" xml:"dep_apt,omitempty"`
	// example:
	//
	// HGH
	DepAptCode *string `json:"dep_apt_code,omitempty" xml:"dep_apt_code,omitempty"`
	DepCity    *string `json:"dep_city,omitempty" xml:"dep_city,omitempty"`
	// example:
	//
	// HGH
	DepCityCode *string `json:"dep_city_code,omitempty" xml:"dep_city_code,omitempty"`
	// example:
	//
	// T1
	DepTerminal *string `json:"dep_terminal,omitempty" xml:"dep_terminal,omitempty"`
	// example:
	//
	// 2022-07-20T10:40Z
	DepTime    *string `json:"dep_time,omitempty" xml:"dep_time,omitempty"`
	FlightMile *int32  `json:"flight_mile,omitempty" xml:"flight_mile,omitempty"`
	// example:
	//
	// MU5619
	FlightNo *string `json:"flight_no,omitempty" xml:"flight_no,omitempty"`
	// example:
	//
	// 0
	JourneyIndex *int32 `json:"journey_index,omitempty" xml:"journey_index,omitempty"`
	// example:
	//
	// 0
	SegmentIndex *int32    `json:"segment_index,omitempty" xml:"segment_index,omitempty"`
	StopCity     []*string `json:"stop_city,omitempty" xml:"stop_city,omitempty" type:"Repeated"`
}

func (s FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundSegmentList) String() string {
	return dara.Prettify(s)
}

func (s FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundSegmentList) GoString() string {
	return s.String()
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundSegmentList) GetAirlineCode() *string {
	return s.AirlineCode
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundSegmentList) GetAirlineName() *string {
	return s.AirlineName
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundSegmentList) GetArrApt() *string {
	return s.ArrApt
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundSegmentList) GetArrAptCode() *string {
	return s.ArrAptCode
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundSegmentList) GetArrCity() *string {
	return s.ArrCity
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundSegmentList) GetArrCityCode() *string {
	return s.ArrCityCode
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundSegmentList) GetArrTerminal() *string {
	return s.ArrTerminal
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundSegmentList) GetArrTime() *string {
	return s.ArrTime
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundSegmentList) GetDepApt() *string {
	return s.DepApt
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundSegmentList) GetDepAptCode() *string {
	return s.DepAptCode
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundSegmentList) GetDepCity() *string {
	return s.DepCity
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundSegmentList) GetDepCityCode() *string {
	return s.DepCityCode
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundSegmentList) GetDepTerminal() *string {
	return s.DepTerminal
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundSegmentList) GetDepTime() *string {
	return s.DepTime
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundSegmentList) GetFlightMile() *int32 {
	return s.FlightMile
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundSegmentList) GetFlightNo() *string {
	return s.FlightNo
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundSegmentList) GetJourneyIndex() *int32 {
	return s.JourneyIndex
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundSegmentList) GetSegmentIndex() *int32 {
	return s.SegmentIndex
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundSegmentList) GetStopCity() []*string {
	return s.StopCity
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundSegmentList) SetAirlineCode(v string) *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundSegmentList {
	s.AirlineCode = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundSegmentList) SetAirlineName(v string) *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundSegmentList {
	s.AirlineName = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundSegmentList) SetArrApt(v string) *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundSegmentList {
	s.ArrApt = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundSegmentList) SetArrAptCode(v string) *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundSegmentList {
	s.ArrAptCode = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundSegmentList) SetArrCity(v string) *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundSegmentList {
	s.ArrCity = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundSegmentList) SetArrCityCode(v string) *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundSegmentList {
	s.ArrCityCode = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundSegmentList) SetArrTerminal(v string) *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundSegmentList {
	s.ArrTerminal = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundSegmentList) SetArrTime(v string) *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundSegmentList {
	s.ArrTime = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundSegmentList) SetDepApt(v string) *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundSegmentList {
	s.DepApt = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundSegmentList) SetDepAptCode(v string) *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundSegmentList {
	s.DepAptCode = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundSegmentList) SetDepCity(v string) *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundSegmentList {
	s.DepCity = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundSegmentList) SetDepCityCode(v string) *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundSegmentList {
	s.DepCityCode = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundSegmentList) SetDepTerminal(v string) *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundSegmentList {
	s.DepTerminal = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundSegmentList) SetDepTime(v string) *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundSegmentList {
	s.DepTime = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundSegmentList) SetFlightMile(v int32) *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundSegmentList {
	s.FlightMile = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundSegmentList) SetFlightNo(v string) *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundSegmentList {
	s.FlightNo = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundSegmentList) SetJourneyIndex(v int32) *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundSegmentList {
	s.JourneyIndex = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundSegmentList) SetSegmentIndex(v int32) *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundSegmentList {
	s.SegmentIndex = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundSegmentList) SetStopCity(v []*string) *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundSegmentList {
	s.StopCity = v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundSegmentList) Validate() error {
	return dara.Validate(s)
}

type FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundUserFeeList struct {
	AlreadyUseAmount                 *float64 `json:"already_use_amount,omitempty" xml:"already_use_amount,omitempty"`
	NonRefundableReshopChangeAmount  *float64 `json:"non_refundable_reshop_change_amount,omitempty" xml:"non_refundable_reshop_change_amount,omitempty"`
	NonRefundableReshopUpgradeAmount *float64 `json:"non_refundable_reshop_upgrade_amount,omitempty" xml:"non_refundable_reshop_upgrade_amount,omitempty"`
	RefundAmount                     *float64 `json:"refund_amount,omitempty" xml:"refund_amount,omitempty"`
	RefundCorpAmount                 *float64 `json:"refund_corp_amount,omitempty" xml:"refund_corp_amount,omitempty"`
	RefundHandFee                    *float64 `json:"refund_hand_fee,omitempty" xml:"refund_hand_fee,omitempty"`
	RefundPersonAmount               *float64 `json:"refund_person_amount,omitempty" xml:"refund_person_amount,omitempty"`
	UserId                           *string  `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

func (s FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundUserFeeList) String() string {
	return dara.Prettify(s)
}

func (s FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundUserFeeList) GoString() string {
	return s.String()
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundUserFeeList) GetAlreadyUseAmount() *float64 {
	return s.AlreadyUseAmount
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundUserFeeList) GetNonRefundableReshopChangeAmount() *float64 {
	return s.NonRefundableReshopChangeAmount
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundUserFeeList) GetNonRefundableReshopUpgradeAmount() *float64 {
	return s.NonRefundableReshopUpgradeAmount
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundUserFeeList) GetRefundAmount() *float64 {
	return s.RefundAmount
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundUserFeeList) GetRefundCorpAmount() *float64 {
	return s.RefundCorpAmount
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundUserFeeList) GetRefundHandFee() *float64 {
	return s.RefundHandFee
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundUserFeeList) GetRefundPersonAmount() *float64 {
	return s.RefundPersonAmount
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundUserFeeList) GetUserId() *string {
	return s.UserId
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundUserFeeList) SetAlreadyUseAmount(v float64) *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundUserFeeList {
	s.AlreadyUseAmount = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundUserFeeList) SetNonRefundableReshopChangeAmount(v float64) *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundUserFeeList {
	s.NonRefundableReshopChangeAmount = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundUserFeeList) SetNonRefundableReshopUpgradeAmount(v float64) *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundUserFeeList {
	s.NonRefundableReshopUpgradeAmount = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundUserFeeList) SetRefundAmount(v float64) *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundUserFeeList {
	s.RefundAmount = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundUserFeeList) SetRefundCorpAmount(v float64) *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundUserFeeList {
	s.RefundCorpAmount = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundUserFeeList) SetRefundHandFee(v float64) *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundUserFeeList {
	s.RefundHandFee = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundUserFeeList) SetRefundPersonAmount(v float64) *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundUserFeeList {
	s.RefundPersonAmount = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundUserFeeList) SetUserId(v string) *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundUserFeeList {
	s.UserId = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightRefundApplyListFlightRefundUserFeeList) Validate() error {
	return dara.Validate(s)
}

type FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyList struct {
	FlightReshopApplyTicketList []*FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyListFlightReshopApplyTicketList `json:"flight_reshop_apply_ticket_list,omitempty" xml:"flight_reshop_apply_ticket_list,omitempty" type:"Repeated"`
	FlightReshopSegmentList     []*FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyListFlightReshopSegmentList     `json:"flight_reshop_segment_list,omitempty" xml:"flight_reshop_segment_list,omitempty" type:"Repeated"`
	FlightReshopUserFeeList     []*FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyListFlightReshopUserFeeList     `json:"flight_reshop_user_fee_list,omitempty" xml:"flight_reshop_user_fee_list,omitempty" type:"Repeated"`
	// example:
	//
	// 100231231
	RelateReshopApplyId *int64 `json:"relate_reshop_apply_id,omitempty" xml:"relate_reshop_apply_id,omitempty"`
	// example:
	//
	// 123232323
	ReshopApplyId *int64 `json:"reshop_apply_id,omitempty" xml:"reshop_apply_id,omitempty"`
	// example:
	//
	// 122312
	ReshopApproveId         *string  `json:"reshop_approve_id,omitempty" xml:"reshop_approve_id,omitempty"`
	ReshopCorpTotalAmount   *float64 `json:"reshop_corp_total_amount,omitempty" xml:"reshop_corp_total_amount,omitempty"`
	ReshopPersonTotalAmount *float64 `json:"reshop_person_total_amount,omitempty" xml:"reshop_person_total_amount,omitempty"`
	ReshopReason            *string  `json:"reshop_reason,omitempty" xml:"reshop_reason,omitempty"`
	// example:
	//
	// 1002
	ReshopReasonCode  *string   `json:"reshop_reason_code,omitempty" xml:"reshop_reason_code,omitempty"`
	ReshopTotalAmount *float64  `json:"reshop_total_amount,omitempty" xml:"reshop_total_amount,omitempty"`
	UserIdList        []*string `json:"user_id_list,omitempty" xml:"user_id_list,omitempty" type:"Repeated"`
}

func (s FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyList) String() string {
	return dara.Prettify(s)
}

func (s FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyList) GoString() string {
	return s.String()
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyList) GetFlightReshopApplyTicketList() []*FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyListFlightReshopApplyTicketList {
	return s.FlightReshopApplyTicketList
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyList) GetFlightReshopSegmentList() []*FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyListFlightReshopSegmentList {
	return s.FlightReshopSegmentList
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyList) GetFlightReshopUserFeeList() []*FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyListFlightReshopUserFeeList {
	return s.FlightReshopUserFeeList
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyList) GetRelateReshopApplyId() *int64 {
	return s.RelateReshopApplyId
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyList) GetReshopApplyId() *int64 {
	return s.ReshopApplyId
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyList) GetReshopApproveId() *string {
	return s.ReshopApproveId
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyList) GetReshopCorpTotalAmount() *float64 {
	return s.ReshopCorpTotalAmount
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyList) GetReshopPersonTotalAmount() *float64 {
	return s.ReshopPersonTotalAmount
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyList) GetReshopReason() *string {
	return s.ReshopReason
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyList) GetReshopReasonCode() *string {
	return s.ReshopReasonCode
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyList) GetReshopTotalAmount() *float64 {
	return s.ReshopTotalAmount
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyList) GetUserIdList() []*string {
	return s.UserIdList
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyList) SetFlightReshopApplyTicketList(v []*FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyListFlightReshopApplyTicketList) *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyList {
	s.FlightReshopApplyTicketList = v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyList) SetFlightReshopSegmentList(v []*FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyListFlightReshopSegmentList) *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyList {
	s.FlightReshopSegmentList = v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyList) SetFlightReshopUserFeeList(v []*FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyListFlightReshopUserFeeList) *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyList {
	s.FlightReshopUserFeeList = v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyList) SetRelateReshopApplyId(v int64) *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyList {
	s.RelateReshopApplyId = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyList) SetReshopApplyId(v int64) *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyList {
	s.ReshopApplyId = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyList) SetReshopApproveId(v string) *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyList {
	s.ReshopApproveId = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyList) SetReshopCorpTotalAmount(v float64) *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyList {
	s.ReshopCorpTotalAmount = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyList) SetReshopPersonTotalAmount(v float64) *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyList {
	s.ReshopPersonTotalAmount = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyList) SetReshopReason(v string) *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyList {
	s.ReshopReason = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyList) SetReshopReasonCode(v string) *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyList {
	s.ReshopReasonCode = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyList) SetReshopTotalAmount(v float64) *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyList {
	s.ReshopTotalAmount = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyList) SetUserIdList(v []*string) *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyList {
	s.UserIdList = v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyList) Validate() error {
	return dara.Validate(s)
}

type FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyListFlightReshopApplyTicketList struct {
	FlightList         []*FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyListFlightReshopApplyTicketListFlightList `json:"flight_list,omitempty" xml:"flight_list,omitempty" type:"Repeated"`
	RelateTicketNoList []*string                                                                                             `json:"relate_ticket_no_list,omitempty" xml:"relate_ticket_no_list,omitempty" type:"Repeated"`
	TicketNoList       []*string                                                                                             `json:"ticket_no_list,omitempty" xml:"ticket_no_list,omitempty" type:"Repeated"`
	// example:
	//
	// alitrip123
	UserId *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

func (s FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyListFlightReshopApplyTicketList) String() string {
	return dara.Prettify(s)
}

func (s FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyListFlightReshopApplyTicketList) GoString() string {
	return s.String()
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyListFlightReshopApplyTicketList) GetFlightList() []*FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyListFlightReshopApplyTicketListFlightList {
	return s.FlightList
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyListFlightReshopApplyTicketList) GetRelateTicketNoList() []*string {
	return s.RelateTicketNoList
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyListFlightReshopApplyTicketList) GetTicketNoList() []*string {
	return s.TicketNoList
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyListFlightReshopApplyTicketList) GetUserId() *string {
	return s.UserId
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyListFlightReshopApplyTicketList) SetFlightList(v []*FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyListFlightReshopApplyTicketListFlightList) *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyListFlightReshopApplyTicketList {
	s.FlightList = v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyListFlightReshopApplyTicketList) SetRelateTicketNoList(v []*string) *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyListFlightReshopApplyTicketList {
	s.RelateTicketNoList = v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyListFlightReshopApplyTicketList) SetTicketNoList(v []*string) *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyListFlightReshopApplyTicketList {
	s.TicketNoList = v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyListFlightReshopApplyTicketList) SetUserId(v string) *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyListFlightReshopApplyTicketList {
	s.UserId = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyListFlightReshopApplyTicketList) Validate() error {
	return dara.Validate(s)
}

type FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyListFlightReshopApplyTicketListFlightList struct {
	// example:
	//
	// 2022-07-20T10:40Z
	ArrTime *string `json:"arr_time,omitempty" xml:"arr_time,omitempty"`
	// example:
	//
	// V
	Cabin *string `json:"cabin,omitempty" xml:"cabin,omitempty"`
	// example:
	//
	// Y
	CabinClass     *string `json:"cabin_class,omitempty" xml:"cabin_class,omitempty"`
	CabinClassName *string `json:"cabin_class_name,omitempty" xml:"cabin_class_name,omitempty"`
	// example:
	//
	// 2022-07-20T10:40Z
	DepTime *string `json:"dep_time,omitempty" xml:"dep_time,omitempty"`
	// example:
	//
	// MU1398
	FlightNo *string `json:"flight_no,omitempty" xml:"flight_no,omitempty"`
}

func (s FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyListFlightReshopApplyTicketListFlightList) String() string {
	return dara.Prettify(s)
}

func (s FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyListFlightReshopApplyTicketListFlightList) GoString() string {
	return s.String()
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyListFlightReshopApplyTicketListFlightList) GetArrTime() *string {
	return s.ArrTime
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyListFlightReshopApplyTicketListFlightList) GetCabin() *string {
	return s.Cabin
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyListFlightReshopApplyTicketListFlightList) GetCabinClass() *string {
	return s.CabinClass
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyListFlightReshopApplyTicketListFlightList) GetCabinClassName() *string {
	return s.CabinClassName
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyListFlightReshopApplyTicketListFlightList) GetDepTime() *string {
	return s.DepTime
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyListFlightReshopApplyTicketListFlightList) GetFlightNo() *string {
	return s.FlightNo
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyListFlightReshopApplyTicketListFlightList) SetArrTime(v string) *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyListFlightReshopApplyTicketListFlightList {
	s.ArrTime = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyListFlightReshopApplyTicketListFlightList) SetCabin(v string) *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyListFlightReshopApplyTicketListFlightList {
	s.Cabin = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyListFlightReshopApplyTicketListFlightList) SetCabinClass(v string) *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyListFlightReshopApplyTicketListFlightList {
	s.CabinClass = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyListFlightReshopApplyTicketListFlightList) SetCabinClassName(v string) *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyListFlightReshopApplyTicketListFlightList {
	s.CabinClassName = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyListFlightReshopApplyTicketListFlightList) SetDepTime(v string) *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyListFlightReshopApplyTicketListFlightList {
	s.DepTime = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyListFlightReshopApplyTicketListFlightList) SetFlightNo(v string) *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyListFlightReshopApplyTicketListFlightList {
	s.FlightNo = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyListFlightReshopApplyTicketListFlightList) Validate() error {
	return dara.Validate(s)
}

type FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyListFlightReshopSegmentList struct {
	// example:
	//
	// CA
	AirlineCode *string `json:"airline_code,omitempty" xml:"airline_code,omitempty"`
	AirlineName *string `json:"airline_name,omitempty" xml:"airline_name,omitempty"`
	ArrApt      *string `json:"arr_apt,omitempty" xml:"arr_apt,omitempty"`
	// example:
	//
	// PEX
	ArrAptCode *string `json:"arr_apt_code,omitempty" xml:"arr_apt_code,omitempty"`
	ArrCity    *string `json:"arr_city,omitempty" xml:"arr_city,omitempty"`
	// example:
	//
	// BJS
	ArrCityCode *string `json:"arr_city_code,omitempty" xml:"arr_city_code,omitempty"`
	// example:
	//
	// T1
	ArrTerminal *string `json:"arr_terminal,omitempty" xml:"arr_terminal,omitempty"`
	// example:
	//
	// 2022-07-20T10:40Z
	ArrTime *string `json:"arr_time,omitempty" xml:"arr_time,omitempty"`
	DepApt  *string `json:"dep_apt,omitempty" xml:"dep_apt,omitempty"`
	// example:
	//
	// HGH
	DepAptCode *string `json:"dep_apt_code,omitempty" xml:"dep_apt_code,omitempty"`
	DepCity    *string `json:"dep_city,omitempty" xml:"dep_city,omitempty"`
	// example:
	//
	// HGH
	DepCityCode *string `json:"dep_city_code,omitempty" xml:"dep_city_code,omitempty"`
	// example:
	//
	// T1
	DepTerminal *string `json:"dep_terminal,omitempty" xml:"dep_terminal,omitempty"`
	// example:
	//
	// 2022-07-20T10:40Z
	DepTime    *string `json:"dep_time,omitempty" xml:"dep_time,omitempty"`
	FlightMile *int32  `json:"flight_mile,omitempty" xml:"flight_mile,omitempty"`
	// example:
	//
	// CA3358
	FlightNo *string `json:"flight_no,omitempty" xml:"flight_no,omitempty"`
	// example:
	//
	// 0
	JourneyIndex *int32 `json:"journey_index,omitempty" xml:"journey_index,omitempty"`
	// example:
	//
	// 0
	SegmentIndex *int32    `json:"segment_index,omitempty" xml:"segment_index,omitempty"`
	StopCity     []*string `json:"stop_city,omitempty" xml:"stop_city,omitempty" type:"Repeated"`
}

func (s FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyListFlightReshopSegmentList) String() string {
	return dara.Prettify(s)
}

func (s FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyListFlightReshopSegmentList) GoString() string {
	return s.String()
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyListFlightReshopSegmentList) GetAirlineCode() *string {
	return s.AirlineCode
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyListFlightReshopSegmentList) GetAirlineName() *string {
	return s.AirlineName
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyListFlightReshopSegmentList) GetArrApt() *string {
	return s.ArrApt
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyListFlightReshopSegmentList) GetArrAptCode() *string {
	return s.ArrAptCode
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyListFlightReshopSegmentList) GetArrCity() *string {
	return s.ArrCity
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyListFlightReshopSegmentList) GetArrCityCode() *string {
	return s.ArrCityCode
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyListFlightReshopSegmentList) GetArrTerminal() *string {
	return s.ArrTerminal
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyListFlightReshopSegmentList) GetArrTime() *string {
	return s.ArrTime
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyListFlightReshopSegmentList) GetDepApt() *string {
	return s.DepApt
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyListFlightReshopSegmentList) GetDepAptCode() *string {
	return s.DepAptCode
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyListFlightReshopSegmentList) GetDepCity() *string {
	return s.DepCity
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyListFlightReshopSegmentList) GetDepCityCode() *string {
	return s.DepCityCode
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyListFlightReshopSegmentList) GetDepTerminal() *string {
	return s.DepTerminal
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyListFlightReshopSegmentList) GetDepTime() *string {
	return s.DepTime
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyListFlightReshopSegmentList) GetFlightMile() *int32 {
	return s.FlightMile
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyListFlightReshopSegmentList) GetFlightNo() *string {
	return s.FlightNo
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyListFlightReshopSegmentList) GetJourneyIndex() *int32 {
	return s.JourneyIndex
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyListFlightReshopSegmentList) GetSegmentIndex() *int32 {
	return s.SegmentIndex
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyListFlightReshopSegmentList) GetStopCity() []*string {
	return s.StopCity
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyListFlightReshopSegmentList) SetAirlineCode(v string) *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyListFlightReshopSegmentList {
	s.AirlineCode = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyListFlightReshopSegmentList) SetAirlineName(v string) *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyListFlightReshopSegmentList {
	s.AirlineName = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyListFlightReshopSegmentList) SetArrApt(v string) *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyListFlightReshopSegmentList {
	s.ArrApt = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyListFlightReshopSegmentList) SetArrAptCode(v string) *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyListFlightReshopSegmentList {
	s.ArrAptCode = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyListFlightReshopSegmentList) SetArrCity(v string) *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyListFlightReshopSegmentList {
	s.ArrCity = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyListFlightReshopSegmentList) SetArrCityCode(v string) *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyListFlightReshopSegmentList {
	s.ArrCityCode = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyListFlightReshopSegmentList) SetArrTerminal(v string) *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyListFlightReshopSegmentList {
	s.ArrTerminal = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyListFlightReshopSegmentList) SetArrTime(v string) *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyListFlightReshopSegmentList {
	s.ArrTime = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyListFlightReshopSegmentList) SetDepApt(v string) *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyListFlightReshopSegmentList {
	s.DepApt = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyListFlightReshopSegmentList) SetDepAptCode(v string) *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyListFlightReshopSegmentList {
	s.DepAptCode = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyListFlightReshopSegmentList) SetDepCity(v string) *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyListFlightReshopSegmentList {
	s.DepCity = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyListFlightReshopSegmentList) SetDepCityCode(v string) *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyListFlightReshopSegmentList {
	s.DepCityCode = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyListFlightReshopSegmentList) SetDepTerminal(v string) *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyListFlightReshopSegmentList {
	s.DepTerminal = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyListFlightReshopSegmentList) SetDepTime(v string) *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyListFlightReshopSegmentList {
	s.DepTime = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyListFlightReshopSegmentList) SetFlightMile(v int32) *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyListFlightReshopSegmentList {
	s.FlightMile = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyListFlightReshopSegmentList) SetFlightNo(v string) *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyListFlightReshopSegmentList {
	s.FlightNo = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyListFlightReshopSegmentList) SetJourneyIndex(v int32) *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyListFlightReshopSegmentList {
	s.JourneyIndex = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyListFlightReshopSegmentList) SetSegmentIndex(v int32) *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyListFlightReshopSegmentList {
	s.SegmentIndex = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyListFlightReshopSegmentList) SetStopCity(v []*string) *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyListFlightReshopSegmentList {
	s.StopCity = v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyListFlightReshopSegmentList) Validate() error {
	return dara.Validate(s)
}

type FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyListFlightReshopUserFeeList struct {
	ChangeFee          *float64 `json:"change_fee,omitempty" xml:"change_fee,omitempty"`
	ReshopCorpAmount   *float64 `json:"reshop_corp_amount,omitempty" xml:"reshop_corp_amount,omitempty"`
	ReshopPersonAmount *float64 `json:"reshop_person_amount,omitempty" xml:"reshop_person_amount,omitempty"`
	UpgradeFee         *float64 `json:"upgrade_fee,omitempty" xml:"upgrade_fee,omitempty"`
	// example:
	//
	// alitrip123
	UserId *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

func (s FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyListFlightReshopUserFeeList) String() string {
	return dara.Prettify(s)
}

func (s FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyListFlightReshopUserFeeList) GoString() string {
	return s.String()
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyListFlightReshopUserFeeList) GetChangeFee() *float64 {
	return s.ChangeFee
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyListFlightReshopUserFeeList) GetReshopCorpAmount() *float64 {
	return s.ReshopCorpAmount
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyListFlightReshopUserFeeList) GetReshopPersonAmount() *float64 {
	return s.ReshopPersonAmount
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyListFlightReshopUserFeeList) GetUpgradeFee() *float64 {
	return s.UpgradeFee
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyListFlightReshopUserFeeList) GetUserId() *string {
	return s.UserId
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyListFlightReshopUserFeeList) SetChangeFee(v float64) *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyListFlightReshopUserFeeList {
	s.ChangeFee = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyListFlightReshopUserFeeList) SetReshopCorpAmount(v float64) *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyListFlightReshopUserFeeList {
	s.ReshopCorpAmount = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyListFlightReshopUserFeeList) SetReshopPersonAmount(v float64) *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyListFlightReshopUserFeeList {
	s.ReshopPersonAmount = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyListFlightReshopUserFeeList) SetUpgradeFee(v float64) *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyListFlightReshopUserFeeList {
	s.UpgradeFee = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyListFlightReshopUserFeeList) SetUserId(v string) *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyListFlightReshopUserFeeList {
	s.UserId = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightReshopApplyListFlightReshopUserFeeList) Validate() error {
	return dara.Validate(s)
}

type FlightOrderListQueryV2ResponseBodyModuleFlightSegmentList struct {
	// example:
	//
	// CZ
	AirlineCode *string `json:"airline_code,omitempty" xml:"airline_code,omitempty"`
	AirlineName *string `json:"airline_name,omitempty" xml:"airline_name,omitempty"`
	ArrApt      *string `json:"arr_apt,omitempty" xml:"arr_apt,omitempty"`
	// example:
	//
	// PEK
	ArrAptCode *string `json:"arr_apt_code,omitempty" xml:"arr_apt_code,omitempty"`
	ArrCity    *string `json:"arr_city,omitempty" xml:"arr_city,omitempty"`
	// example:
	//
	// BJS
	ArrCityCode *string `json:"arr_city_code,omitempty" xml:"arr_city_code,omitempty"`
	// example:
	//
	// T1
	ArrTerminal *string `json:"arr_terminal,omitempty" xml:"arr_terminal,omitempty"`
	// example:
	//
	// 2022-07-20T10:40Z
	ArrTime *string `json:"arr_time,omitempty" xml:"arr_time,omitempty"`
	DepApt  *string `json:"dep_apt,omitempty" xml:"dep_apt,omitempty"`
	// example:
	//
	// HGH
	DepAptCode *string `json:"dep_apt_code,omitempty" xml:"dep_apt_code,omitempty"`
	DepCity    *string `json:"dep_city,omitempty" xml:"dep_city,omitempty"`
	// example:
	//
	// HGH
	DepCityCode *string `json:"dep_city_code,omitempty" xml:"dep_city_code,omitempty"`
	// example:
	//
	// T1
	DepTerminal *string `json:"dep_terminal,omitempty" xml:"dep_terminal,omitempty"`
	// example:
	//
	// 2022-07-20T10:40Z
	DepTime    *string `json:"dep_time,omitempty" xml:"dep_time,omitempty"`
	FlightMile *int32  `json:"flight_mile,omitempty" xml:"flight_mile,omitempty"`
	// example:
	//
	// CZ2891
	FlightNo *string `json:"flight_no,omitempty" xml:"flight_no,omitempty"`
	// example:
	//
	// 0
	JourneyIndex *int32 `json:"journey_index,omitempty" xml:"journey_index,omitempty"`
	// example:
	//
	// 0
	SegmentIndex *int32    `json:"segment_index,omitempty" xml:"segment_index,omitempty"`
	StopCity     []*string `json:"stop_city,omitempty" xml:"stop_city,omitempty" type:"Repeated"`
}

func (s FlightOrderListQueryV2ResponseBodyModuleFlightSegmentList) String() string {
	return dara.Prettify(s)
}

func (s FlightOrderListQueryV2ResponseBodyModuleFlightSegmentList) GoString() string {
	return s.String()
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightSegmentList) GetAirlineCode() *string {
	return s.AirlineCode
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightSegmentList) GetAirlineName() *string {
	return s.AirlineName
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightSegmentList) GetArrApt() *string {
	return s.ArrApt
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightSegmentList) GetArrAptCode() *string {
	return s.ArrAptCode
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightSegmentList) GetArrCity() *string {
	return s.ArrCity
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightSegmentList) GetArrCityCode() *string {
	return s.ArrCityCode
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightSegmentList) GetArrTerminal() *string {
	return s.ArrTerminal
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightSegmentList) GetArrTime() *string {
	return s.ArrTime
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightSegmentList) GetDepApt() *string {
	return s.DepApt
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightSegmentList) GetDepAptCode() *string {
	return s.DepAptCode
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightSegmentList) GetDepCity() *string {
	return s.DepCity
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightSegmentList) GetDepCityCode() *string {
	return s.DepCityCode
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightSegmentList) GetDepTerminal() *string {
	return s.DepTerminal
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightSegmentList) GetDepTime() *string {
	return s.DepTime
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightSegmentList) GetFlightMile() *int32 {
	return s.FlightMile
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightSegmentList) GetFlightNo() *string {
	return s.FlightNo
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightSegmentList) GetJourneyIndex() *int32 {
	return s.JourneyIndex
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightSegmentList) GetSegmentIndex() *int32 {
	return s.SegmentIndex
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightSegmentList) GetStopCity() []*string {
	return s.StopCity
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightSegmentList) SetAirlineCode(v string) *FlightOrderListQueryV2ResponseBodyModuleFlightSegmentList {
	s.AirlineCode = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightSegmentList) SetAirlineName(v string) *FlightOrderListQueryV2ResponseBodyModuleFlightSegmentList {
	s.AirlineName = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightSegmentList) SetArrApt(v string) *FlightOrderListQueryV2ResponseBodyModuleFlightSegmentList {
	s.ArrApt = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightSegmentList) SetArrAptCode(v string) *FlightOrderListQueryV2ResponseBodyModuleFlightSegmentList {
	s.ArrAptCode = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightSegmentList) SetArrCity(v string) *FlightOrderListQueryV2ResponseBodyModuleFlightSegmentList {
	s.ArrCity = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightSegmentList) SetArrCityCode(v string) *FlightOrderListQueryV2ResponseBodyModuleFlightSegmentList {
	s.ArrCityCode = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightSegmentList) SetArrTerminal(v string) *FlightOrderListQueryV2ResponseBodyModuleFlightSegmentList {
	s.ArrTerminal = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightSegmentList) SetArrTime(v string) *FlightOrderListQueryV2ResponseBodyModuleFlightSegmentList {
	s.ArrTime = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightSegmentList) SetDepApt(v string) *FlightOrderListQueryV2ResponseBodyModuleFlightSegmentList {
	s.DepApt = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightSegmentList) SetDepAptCode(v string) *FlightOrderListQueryV2ResponseBodyModuleFlightSegmentList {
	s.DepAptCode = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightSegmentList) SetDepCity(v string) *FlightOrderListQueryV2ResponseBodyModuleFlightSegmentList {
	s.DepCity = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightSegmentList) SetDepCityCode(v string) *FlightOrderListQueryV2ResponseBodyModuleFlightSegmentList {
	s.DepCityCode = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightSegmentList) SetDepTerminal(v string) *FlightOrderListQueryV2ResponseBodyModuleFlightSegmentList {
	s.DepTerminal = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightSegmentList) SetDepTime(v string) *FlightOrderListQueryV2ResponseBodyModuleFlightSegmentList {
	s.DepTime = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightSegmentList) SetFlightMile(v int32) *FlightOrderListQueryV2ResponseBodyModuleFlightSegmentList {
	s.FlightMile = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightSegmentList) SetFlightNo(v string) *FlightOrderListQueryV2ResponseBodyModuleFlightSegmentList {
	s.FlightNo = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightSegmentList) SetJourneyIndex(v int32) *FlightOrderListQueryV2ResponseBodyModuleFlightSegmentList {
	s.JourneyIndex = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightSegmentList) SetSegmentIndex(v int32) *FlightOrderListQueryV2ResponseBodyModuleFlightSegmentList {
	s.SegmentIndex = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightSegmentList) SetStopCity(v []*string) *FlightOrderListQueryV2ResponseBodyModuleFlightSegmentList {
	s.StopCity = v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleFlightSegmentList) Validate() error {
	return dara.Validate(s)
}

type FlightOrderListQueryV2ResponseBodyModuleInsureInfoList struct {
	// example:
	//
	// 1002308231
	InsureId          *string   `json:"insure_id,omitempty" xml:"insure_id,omitempty"`
	InsureOrderAmount *float64  `json:"insure_order_amount,omitempty" xml:"insure_order_amount,omitempty"`
	InsurePrice       *float64  `json:"insure_price,omitempty" xml:"insure_price,omitempty"`
	InsureType        *string   `json:"insure_type,omitempty" xml:"insure_type,omitempty"`
	NameList          []*string `json:"name_list,omitempty" xml:"name_list,omitempty" type:"Repeated"`
	// example:
	//
	// 5
	Number *int32 `json:"number,omitempty" xml:"number,omitempty"`
	// example:
	//
	// 0
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
}

func (s FlightOrderListQueryV2ResponseBodyModuleInsureInfoList) String() string {
	return dara.Prettify(s)
}

func (s FlightOrderListQueryV2ResponseBodyModuleInsureInfoList) GoString() string {
	return s.String()
}

func (s *FlightOrderListQueryV2ResponseBodyModuleInsureInfoList) GetInsureId() *string {
	return s.InsureId
}

func (s *FlightOrderListQueryV2ResponseBodyModuleInsureInfoList) GetInsureOrderAmount() *float64 {
	return s.InsureOrderAmount
}

func (s *FlightOrderListQueryV2ResponseBodyModuleInsureInfoList) GetInsurePrice() *float64 {
	return s.InsurePrice
}

func (s *FlightOrderListQueryV2ResponseBodyModuleInsureInfoList) GetInsureType() *string {
	return s.InsureType
}

func (s *FlightOrderListQueryV2ResponseBodyModuleInsureInfoList) GetNameList() []*string {
	return s.NameList
}

func (s *FlightOrderListQueryV2ResponseBodyModuleInsureInfoList) GetNumber() *int32 {
	return s.Number
}

func (s *FlightOrderListQueryV2ResponseBodyModuleInsureInfoList) GetStatus() *int32 {
	return s.Status
}

func (s *FlightOrderListQueryV2ResponseBodyModuleInsureInfoList) SetInsureId(v string) *FlightOrderListQueryV2ResponseBodyModuleInsureInfoList {
	s.InsureId = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleInsureInfoList) SetInsureOrderAmount(v float64) *FlightOrderListQueryV2ResponseBodyModuleInsureInfoList {
	s.InsureOrderAmount = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleInsureInfoList) SetInsurePrice(v float64) *FlightOrderListQueryV2ResponseBodyModuleInsureInfoList {
	s.InsurePrice = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleInsureInfoList) SetInsureType(v string) *FlightOrderListQueryV2ResponseBodyModuleInsureInfoList {
	s.InsureType = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleInsureInfoList) SetNameList(v []*string) *FlightOrderListQueryV2ResponseBodyModuleInsureInfoList {
	s.NameList = v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleInsureInfoList) SetNumber(v int32) *FlightOrderListQueryV2ResponseBodyModuleInsureInfoList {
	s.Number = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleInsureInfoList) SetStatus(v int32) *FlightOrderListQueryV2ResponseBodyModuleInsureInfoList {
	s.Status = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleInsureInfoList) Validate() error {
	return dara.Validate(s)
}

type FlightOrderListQueryV2ResponseBodyModulePriceInfoList struct {
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
	// yyyy-MM-dd HH:mm:ss
	GmtCreate *string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// example:
	//
	// 1
	PayType *int32   `json:"pay_type,omitempty" xml:"pay_type,omitempty"`
	Price   *float64 `json:"price,omitempty" xml:"price,omitempty"`
	// example:
	//
	// 175549295
	SubOrderId *string `json:"sub_order_id,omitempty" xml:"sub_order_id,omitempty"`
	// example:
	//
	// f98236773
	TradeId *string `json:"trade_id,omitempty" xml:"trade_id,omitempty"`
	// example:
	//
	// 1
	Type *int32 `json:"type,omitempty" xml:"type,omitempty"`
}

func (s FlightOrderListQueryV2ResponseBodyModulePriceInfoList) String() string {
	return dara.Prettify(s)
}

func (s FlightOrderListQueryV2ResponseBodyModulePriceInfoList) GoString() string {
	return s.String()
}

func (s *FlightOrderListQueryV2ResponseBodyModulePriceInfoList) GetCategoryCode() *int32 {
	return s.CategoryCode
}

func (s *FlightOrderListQueryV2ResponseBodyModulePriceInfoList) GetCategoryType() *int32 {
	return s.CategoryType
}

func (s *FlightOrderListQueryV2ResponseBodyModulePriceInfoList) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *FlightOrderListQueryV2ResponseBodyModulePriceInfoList) GetPayType() *int32 {
	return s.PayType
}

func (s *FlightOrderListQueryV2ResponseBodyModulePriceInfoList) GetPrice() *float64 {
	return s.Price
}

func (s *FlightOrderListQueryV2ResponseBodyModulePriceInfoList) GetSubOrderId() *string {
	return s.SubOrderId
}

func (s *FlightOrderListQueryV2ResponseBodyModulePriceInfoList) GetTradeId() *string {
	return s.TradeId
}

func (s *FlightOrderListQueryV2ResponseBodyModulePriceInfoList) GetType() *int32 {
	return s.Type
}

func (s *FlightOrderListQueryV2ResponseBodyModulePriceInfoList) SetCategoryCode(v int32) *FlightOrderListQueryV2ResponseBodyModulePriceInfoList {
	s.CategoryCode = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModulePriceInfoList) SetCategoryType(v int32) *FlightOrderListQueryV2ResponseBodyModulePriceInfoList {
	s.CategoryType = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModulePriceInfoList) SetGmtCreate(v string) *FlightOrderListQueryV2ResponseBodyModulePriceInfoList {
	s.GmtCreate = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModulePriceInfoList) SetPayType(v int32) *FlightOrderListQueryV2ResponseBodyModulePriceInfoList {
	s.PayType = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModulePriceInfoList) SetPrice(v float64) *FlightOrderListQueryV2ResponseBodyModulePriceInfoList {
	s.Price = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModulePriceInfoList) SetSubOrderId(v string) *FlightOrderListQueryV2ResponseBodyModulePriceInfoList {
	s.SubOrderId = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModulePriceInfoList) SetTradeId(v string) *FlightOrderListQueryV2ResponseBodyModulePriceInfoList {
	s.TradeId = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModulePriceInfoList) SetType(v int32) *FlightOrderListQueryV2ResponseBodyModulePriceInfoList {
	s.Type = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModulePriceInfoList) Validate() error {
	return dara.Validate(s)
}

type FlightOrderListQueryV2ResponseBodyModuleUserAffiliateList struct {
	CostCenter *FlightOrderListQueryV2ResponseBodyModuleUserAffiliateListCostCenter `json:"cost_center,omitempty" xml:"cost_center,omitempty" type:"Struct"`
	Department *FlightOrderListQueryV2ResponseBodyModuleUserAffiliateListDepartment `json:"department,omitempty" xml:"department,omitempty" type:"Struct"`
	Invoice    *FlightOrderListQueryV2ResponseBodyModuleUserAffiliateListInvoice    `json:"invoice,omitempty" xml:"invoice,omitempty" type:"Struct"`
	Project    *FlightOrderListQueryV2ResponseBodyModuleUserAffiliateListProject    `json:"project,omitempty" xml:"project,omitempty" type:"Struct"`
	UserId     *string                                                              `json:"user_id,omitempty" xml:"user_id,omitempty"`
	UserName   *string                                                              `json:"user_name,omitempty" xml:"user_name,omitempty"`
}

func (s FlightOrderListQueryV2ResponseBodyModuleUserAffiliateList) String() string {
	return dara.Prettify(s)
}

func (s FlightOrderListQueryV2ResponseBodyModuleUserAffiliateList) GoString() string {
	return s.String()
}

func (s *FlightOrderListQueryV2ResponseBodyModuleUserAffiliateList) GetCostCenter() *FlightOrderListQueryV2ResponseBodyModuleUserAffiliateListCostCenter {
	return s.CostCenter
}

func (s *FlightOrderListQueryV2ResponseBodyModuleUserAffiliateList) GetDepartment() *FlightOrderListQueryV2ResponseBodyModuleUserAffiliateListDepartment {
	return s.Department
}

func (s *FlightOrderListQueryV2ResponseBodyModuleUserAffiliateList) GetInvoice() *FlightOrderListQueryV2ResponseBodyModuleUserAffiliateListInvoice {
	return s.Invoice
}

func (s *FlightOrderListQueryV2ResponseBodyModuleUserAffiliateList) GetProject() *FlightOrderListQueryV2ResponseBodyModuleUserAffiliateListProject {
	return s.Project
}

func (s *FlightOrderListQueryV2ResponseBodyModuleUserAffiliateList) GetUserId() *string {
	return s.UserId
}

func (s *FlightOrderListQueryV2ResponseBodyModuleUserAffiliateList) GetUserName() *string {
	return s.UserName
}

func (s *FlightOrderListQueryV2ResponseBodyModuleUserAffiliateList) SetCostCenter(v *FlightOrderListQueryV2ResponseBodyModuleUserAffiliateListCostCenter) *FlightOrderListQueryV2ResponseBodyModuleUserAffiliateList {
	s.CostCenter = v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleUserAffiliateList) SetDepartment(v *FlightOrderListQueryV2ResponseBodyModuleUserAffiliateListDepartment) *FlightOrderListQueryV2ResponseBodyModuleUserAffiliateList {
	s.Department = v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleUserAffiliateList) SetInvoice(v *FlightOrderListQueryV2ResponseBodyModuleUserAffiliateListInvoice) *FlightOrderListQueryV2ResponseBodyModuleUserAffiliateList {
	s.Invoice = v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleUserAffiliateList) SetProject(v *FlightOrderListQueryV2ResponseBodyModuleUserAffiliateListProject) *FlightOrderListQueryV2ResponseBodyModuleUserAffiliateList {
	s.Project = v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleUserAffiliateList) SetUserId(v string) *FlightOrderListQueryV2ResponseBodyModuleUserAffiliateList {
	s.UserId = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleUserAffiliateList) SetUserName(v string) *FlightOrderListQueryV2ResponseBodyModuleUserAffiliateList {
	s.UserName = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleUserAffiliateList) Validate() error {
	return dara.Validate(s)
}

type FlightOrderListQueryV2ResponseBodyModuleUserAffiliateListCostCenter struct {
	// example:
	//
	// alitrip
	CorpId *string `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
	// example:
	//
	// alitripTest
	Id   *int64  `json:"id,omitempty" xml:"id,omitempty"`
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// test_cost_center
	Number *string `json:"number,omitempty" xml:"number,omitempty"`
}

func (s FlightOrderListQueryV2ResponseBodyModuleUserAffiliateListCostCenter) String() string {
	return dara.Prettify(s)
}

func (s FlightOrderListQueryV2ResponseBodyModuleUserAffiliateListCostCenter) GoString() string {
	return s.String()
}

func (s *FlightOrderListQueryV2ResponseBodyModuleUserAffiliateListCostCenter) GetCorpId() *string {
	return s.CorpId
}

func (s *FlightOrderListQueryV2ResponseBodyModuleUserAffiliateListCostCenter) GetId() *int64 {
	return s.Id
}

func (s *FlightOrderListQueryV2ResponseBodyModuleUserAffiliateListCostCenter) GetName() *string {
	return s.Name
}

func (s *FlightOrderListQueryV2ResponseBodyModuleUserAffiliateListCostCenter) GetNumber() *string {
	return s.Number
}

func (s *FlightOrderListQueryV2ResponseBodyModuleUserAffiliateListCostCenter) SetCorpId(v string) *FlightOrderListQueryV2ResponseBodyModuleUserAffiliateListCostCenter {
	s.CorpId = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleUserAffiliateListCostCenter) SetId(v int64) *FlightOrderListQueryV2ResponseBodyModuleUserAffiliateListCostCenter {
	s.Id = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleUserAffiliateListCostCenter) SetName(v string) *FlightOrderListQueryV2ResponseBodyModuleUserAffiliateListCostCenter {
	s.Name = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleUserAffiliateListCostCenter) SetNumber(v string) *FlightOrderListQueryV2ResponseBodyModuleUserAffiliateListCostCenter {
	s.Number = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleUserAffiliateListCostCenter) Validate() error {
	return dara.Validate(s)
}

type FlightOrderListQueryV2ResponseBodyModuleUserAffiliateListDepartment struct {
	// example:
	//
	// alitrip
	DepartId   *string `json:"depart_id,omitempty" xml:"depart_id,omitempty"`
	DepartName *string `json:"depart_name,omitempty" xml:"depart_name,omitempty"`
}

func (s FlightOrderListQueryV2ResponseBodyModuleUserAffiliateListDepartment) String() string {
	return dara.Prettify(s)
}

func (s FlightOrderListQueryV2ResponseBodyModuleUserAffiliateListDepartment) GoString() string {
	return s.String()
}

func (s *FlightOrderListQueryV2ResponseBodyModuleUserAffiliateListDepartment) GetDepartId() *string {
	return s.DepartId
}

func (s *FlightOrderListQueryV2ResponseBodyModuleUserAffiliateListDepartment) GetDepartName() *string {
	return s.DepartName
}

func (s *FlightOrderListQueryV2ResponseBodyModuleUserAffiliateListDepartment) SetDepartId(v string) *FlightOrderListQueryV2ResponseBodyModuleUserAffiliateListDepartment {
	s.DepartId = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleUserAffiliateListDepartment) SetDepartName(v string) *FlightOrderListQueryV2ResponseBodyModuleUserAffiliateListDepartment {
	s.DepartName = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleUserAffiliateListDepartment) Validate() error {
	return dara.Validate(s)
}

type FlightOrderListQueryV2ResponseBodyModuleUserAffiliateListInvoice struct {
	// example:
	//
	// test1233
	Id    *int64  `json:"id,omitempty" xml:"id,omitempty"`
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s FlightOrderListQueryV2ResponseBodyModuleUserAffiliateListInvoice) String() string {
	return dara.Prettify(s)
}

func (s FlightOrderListQueryV2ResponseBodyModuleUserAffiliateListInvoice) GoString() string {
	return s.String()
}

func (s *FlightOrderListQueryV2ResponseBodyModuleUserAffiliateListInvoice) GetId() *int64 {
	return s.Id
}

func (s *FlightOrderListQueryV2ResponseBodyModuleUserAffiliateListInvoice) GetTitle() *string {
	return s.Title
}

func (s *FlightOrderListQueryV2ResponseBodyModuleUserAffiliateListInvoice) SetId(v int64) *FlightOrderListQueryV2ResponseBodyModuleUserAffiliateListInvoice {
	s.Id = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleUserAffiliateListInvoice) SetTitle(v string) *FlightOrderListQueryV2ResponseBodyModuleUserAffiliateListInvoice {
	s.Title = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleUserAffiliateListInvoice) Validate() error {
	return dara.Validate(s)
}

type FlightOrderListQueryV2ResponseBodyModuleUserAffiliateListProject struct {
	// example:
	//
	// test_project_id
	ProjectId    *string `json:"project_id,omitempty" xml:"project_id,omitempty"`
	ProjectTitle *string `json:"project_title,omitempty" xml:"project_title,omitempty"`
	// example:
	//
	// test_third_part_project_id
	ThirdpartProjectId *string `json:"thirdpart_project_id,omitempty" xml:"thirdpart_project_id,omitempty"`
}

func (s FlightOrderListQueryV2ResponseBodyModuleUserAffiliateListProject) String() string {
	return dara.Prettify(s)
}

func (s FlightOrderListQueryV2ResponseBodyModuleUserAffiliateListProject) GoString() string {
	return s.String()
}

func (s *FlightOrderListQueryV2ResponseBodyModuleUserAffiliateListProject) GetProjectId() *string {
	return s.ProjectId
}

func (s *FlightOrderListQueryV2ResponseBodyModuleUserAffiliateListProject) GetProjectTitle() *string {
	return s.ProjectTitle
}

func (s *FlightOrderListQueryV2ResponseBodyModuleUserAffiliateListProject) GetThirdpartProjectId() *string {
	return s.ThirdpartProjectId
}

func (s *FlightOrderListQueryV2ResponseBodyModuleUserAffiliateListProject) SetProjectId(v string) *FlightOrderListQueryV2ResponseBodyModuleUserAffiliateListProject {
	s.ProjectId = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleUserAffiliateListProject) SetProjectTitle(v string) *FlightOrderListQueryV2ResponseBodyModuleUserAffiliateListProject {
	s.ProjectTitle = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleUserAffiliateListProject) SetThirdpartProjectId(v string) *FlightOrderListQueryV2ResponseBodyModuleUserAffiliateListProject {
	s.ThirdpartProjectId = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyModuleUserAffiliateListProject) Validate() error {
	return dara.Validate(s)
}

type FlightOrderListQueryV2ResponseBodyPageInfo struct {
	Number *int32 `json:"number,omitempty" xml:"number,omitempty"`
	// example:
	//
	// CAESBgoEIgIIABgAIhkKFwMSAAAAMUw4ZGViODFlYmM3MYzM4
	ScrollId *string `json:"scroll_id,omitempty" xml:"scroll_id,omitempty"`
	// example:
	//
	// 100
	TotalNumber *int32 `json:"total_number,omitempty" xml:"total_number,omitempty"`
}

func (s FlightOrderListQueryV2ResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s FlightOrderListQueryV2ResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *FlightOrderListQueryV2ResponseBodyPageInfo) GetNumber() *int32 {
	return s.Number
}

func (s *FlightOrderListQueryV2ResponseBodyPageInfo) GetScrollId() *string {
	return s.ScrollId
}

func (s *FlightOrderListQueryV2ResponseBodyPageInfo) GetTotalNumber() *int32 {
	return s.TotalNumber
}

func (s *FlightOrderListQueryV2ResponseBodyPageInfo) SetNumber(v int32) *FlightOrderListQueryV2ResponseBodyPageInfo {
	s.Number = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyPageInfo) SetScrollId(v string) *FlightOrderListQueryV2ResponseBodyPageInfo {
	s.ScrollId = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyPageInfo) SetTotalNumber(v int32) *FlightOrderListQueryV2ResponseBodyPageInfo {
	s.TotalNumber = &v
	return s
}

func (s *FlightOrderListQueryV2ResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}
