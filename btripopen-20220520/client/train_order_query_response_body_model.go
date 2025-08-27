// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTrainOrderQueryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *TrainOrderQueryResponseBody
	GetCode() *string
	SetMessage(v string) *TrainOrderQueryResponseBody
	GetMessage() *string
	SetModule(v *TrainOrderQueryResponseBodyModule) *TrainOrderQueryResponseBody
	GetModule() *TrainOrderQueryResponseBodyModule
	SetRequestId(v string) *TrainOrderQueryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *TrainOrderQueryResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *TrainOrderQueryResponseBody
	GetTraceId() *string
}

type TrainOrderQueryResponseBody struct {
	// example:
	//
	// SUCCESS
	Code    *string                            `json:"code,omitempty" xml:"code,omitempty"`
	Message *string                            `json:"message,omitempty" xml:"message,omitempty"`
	Module  *TrainOrderQueryResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	// example:
	//
	// 407543AF-2BD9-5890-BD92-9D1AB7218B27
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

func (s TrainOrderQueryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TrainOrderQueryResponseBody) GoString() string {
	return s.String()
}

func (s *TrainOrderQueryResponseBody) GetCode() *string {
	return s.Code
}

func (s *TrainOrderQueryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *TrainOrderQueryResponseBody) GetModule() *TrainOrderQueryResponseBodyModule {
	return s.Module
}

func (s *TrainOrderQueryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TrainOrderQueryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *TrainOrderQueryResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *TrainOrderQueryResponseBody) SetCode(v string) *TrainOrderQueryResponseBody {
	s.Code = &v
	return s
}

func (s *TrainOrderQueryResponseBody) SetMessage(v string) *TrainOrderQueryResponseBody {
	s.Message = &v
	return s
}

func (s *TrainOrderQueryResponseBody) SetModule(v *TrainOrderQueryResponseBodyModule) *TrainOrderQueryResponseBody {
	s.Module = v
	return s
}

func (s *TrainOrderQueryResponseBody) SetRequestId(v string) *TrainOrderQueryResponseBody {
	s.RequestId = &v
	return s
}

func (s *TrainOrderQueryResponseBody) SetSuccess(v bool) *TrainOrderQueryResponseBody {
	s.Success = &v
	return s
}

func (s *TrainOrderQueryResponseBody) SetTraceId(v string) *TrainOrderQueryResponseBody {
	s.TraceId = &v
	return s
}

func (s *TrainOrderQueryResponseBody) Validate() error {
	return dara.Validate(s)
}

type TrainOrderQueryResponseBodyModule struct {
	ChangeTicketInfoList []*TrainOrderQueryResponseBodyModuleChangeTicketInfoList `json:"change_ticket_info_list,omitempty" xml:"change_ticket_info_list,omitempty" type:"Repeated"`
	InvoiceInfo          *TrainOrderQueryResponseBodyModuleInvoiceInfo            `json:"invoice_info,omitempty" xml:"invoice_info,omitempty" type:"Struct"`
	OrderBaseInfo        *TrainOrderQueryResponseBodyModuleOrderBaseInfo          `json:"order_base_info,omitempty" xml:"order_base_info,omitempty" type:"Struct"`
	PassengerInfoList    []*TrainOrderQueryResponseBodyModulePassengerInfoList    `json:"passenger_info_list,omitempty" xml:"passenger_info_list,omitempty" type:"Repeated"`
	PriceInfoList        []*TrainOrderQueryResponseBodyModulePriceInfoList        `json:"price_info_list,omitempty" xml:"price_info_list,omitempty" type:"Repeated"`
	RefundTicketInfoList []*TrainOrderQueryResponseBodyModuleRefundTicketInfoList `json:"refund_ticket_info_list,omitempty" xml:"refund_ticket_info_list,omitempty" type:"Repeated"`
	TicketInfoList       []*TrainOrderQueryResponseBodyModuleTicketInfoList       `json:"ticket_info_list,omitempty" xml:"ticket_info_list,omitempty" type:"Repeated"`
	TrainInfo            *TrainOrderQueryResponseBodyModuleTrainInfo              `json:"train_info,omitempty" xml:"train_info,omitempty" type:"Struct"`
}

func (s TrainOrderQueryResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s TrainOrderQueryResponseBodyModule) GoString() string {
	return s.String()
}

func (s *TrainOrderQueryResponseBodyModule) GetChangeTicketInfoList() []*TrainOrderQueryResponseBodyModuleChangeTicketInfoList {
	return s.ChangeTicketInfoList
}

func (s *TrainOrderQueryResponseBodyModule) GetInvoiceInfo() *TrainOrderQueryResponseBodyModuleInvoiceInfo {
	return s.InvoiceInfo
}

func (s *TrainOrderQueryResponseBodyModule) GetOrderBaseInfo() *TrainOrderQueryResponseBodyModuleOrderBaseInfo {
	return s.OrderBaseInfo
}

func (s *TrainOrderQueryResponseBodyModule) GetPassengerInfoList() []*TrainOrderQueryResponseBodyModulePassengerInfoList {
	return s.PassengerInfoList
}

func (s *TrainOrderQueryResponseBodyModule) GetPriceInfoList() []*TrainOrderQueryResponseBodyModulePriceInfoList {
	return s.PriceInfoList
}

func (s *TrainOrderQueryResponseBodyModule) GetRefundTicketInfoList() []*TrainOrderQueryResponseBodyModuleRefundTicketInfoList {
	return s.RefundTicketInfoList
}

func (s *TrainOrderQueryResponseBodyModule) GetTicketInfoList() []*TrainOrderQueryResponseBodyModuleTicketInfoList {
	return s.TicketInfoList
}

func (s *TrainOrderQueryResponseBodyModule) GetTrainInfo() *TrainOrderQueryResponseBodyModuleTrainInfo {
	return s.TrainInfo
}

func (s *TrainOrderQueryResponseBodyModule) SetChangeTicketInfoList(v []*TrainOrderQueryResponseBodyModuleChangeTicketInfoList) *TrainOrderQueryResponseBodyModule {
	s.ChangeTicketInfoList = v
	return s
}

func (s *TrainOrderQueryResponseBodyModule) SetInvoiceInfo(v *TrainOrderQueryResponseBodyModuleInvoiceInfo) *TrainOrderQueryResponseBodyModule {
	s.InvoiceInfo = v
	return s
}

func (s *TrainOrderQueryResponseBodyModule) SetOrderBaseInfo(v *TrainOrderQueryResponseBodyModuleOrderBaseInfo) *TrainOrderQueryResponseBodyModule {
	s.OrderBaseInfo = v
	return s
}

func (s *TrainOrderQueryResponseBodyModule) SetPassengerInfoList(v []*TrainOrderQueryResponseBodyModulePassengerInfoList) *TrainOrderQueryResponseBodyModule {
	s.PassengerInfoList = v
	return s
}

func (s *TrainOrderQueryResponseBodyModule) SetPriceInfoList(v []*TrainOrderQueryResponseBodyModulePriceInfoList) *TrainOrderQueryResponseBodyModule {
	s.PriceInfoList = v
	return s
}

func (s *TrainOrderQueryResponseBodyModule) SetRefundTicketInfoList(v []*TrainOrderQueryResponseBodyModuleRefundTicketInfoList) *TrainOrderQueryResponseBodyModule {
	s.RefundTicketInfoList = v
	return s
}

func (s *TrainOrderQueryResponseBodyModule) SetTicketInfoList(v []*TrainOrderQueryResponseBodyModuleTicketInfoList) *TrainOrderQueryResponseBodyModule {
	s.TicketInfoList = v
	return s
}

func (s *TrainOrderQueryResponseBodyModule) SetTrainInfo(v *TrainOrderQueryResponseBodyModuleTrainInfo) *TrainOrderQueryResponseBodyModule {
	s.TrainInfo = v
	return s
}

func (s *TrainOrderQueryResponseBodyModule) Validate() error {
	return dara.Validate(s)
}

type TrainOrderQueryResponseBodyModuleChangeTicketInfoList struct {
	// example:
	//
	// 03
	ChangeCoachNo *string `json:"change_coach_no,omitempty" xml:"change_coach_no,omitempty"`
	// example:
	//
	// 100
	ChangeGapFee *float64 `json:"change_gap_fee,omitempty" xml:"change_gap_fee,omitempty"`
	// example:
	//
	// 100
	ChangeHandlingFee *float64 `json:"change_handling_fee,omitempty" xml:"change_handling_fee,omitempty"`
	// example:
	//
	// 004C
	ChangeSeatNo       *string `json:"change_seat_no,omitempty" xml:"change_seat_no,omitempty"`
	ChangeSeatTypeName *string `json:"change_seat_type_name,omitempty" xml:"change_seat_type_name,omitempty"`
	// example:
	//
	// 100
	ChangeServiceFee *float64 `json:"change_service_fee,omitempty" xml:"change_service_fee,omitempty"`
	// example:
	//
	// D103
	ChangeTrainNo       *string `json:"change_train_no,omitempty" xml:"change_train_no,omitempty"`
	ChangeTrainTypeName *string `json:"change_train_type_name,omitempty" xml:"change_train_type_name,omitempty"`
	// example:
	//
	// 2022-05-15T22:27Z
	CheckInTime *string `json:"check_in_time,omitempty" xml:"check_in_time,omitempty"`
	// example:
	//
	// 2022-05-15T22:27Z
	CheckOutTime *string `json:"check_out_time,omitempty" xml:"check_out_time,omitempty"`
	// example:
	//
	// 2022-05-15T22:27Z
	EndTime         *string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	FromStationName *string `json:"from_station_name,omitempty" xml:"from_station_name,omitempty"`
	// example:
	//
	// 2022-05-15T22:27Z
	GmtCreate *string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// example:
	//
	// 2022-05-15T22:27Z
	GmtModify *string `json:"gmt_modify,omitempty" xml:"gmt_modify,omitempty"`
	// example:
	//
	// CS987JKDF
	OriginTicketNo *string `json:"origin_ticket_no,omitempty" xml:"origin_ticket_no,omitempty"`
	// example:
	//
	// m
	OutTicketStatus *string `json:"out_ticket_status,omitempty" xml:"out_ticket_status,omitempty"`
	// example:
	//
	// 2022-05-15T22:27Z
	StartTime *string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// example:
	//
	// CS987JKDF
	TicketNo      *string `json:"ticket_no,omitempty" xml:"ticket_no,omitempty"`
	ToStationName *string `json:"to_station_name,omitempty" xml:"to_station_name,omitempty"`
}

func (s TrainOrderQueryResponseBodyModuleChangeTicketInfoList) String() string {
	return dara.Prettify(s)
}

func (s TrainOrderQueryResponseBodyModuleChangeTicketInfoList) GoString() string {
	return s.String()
}

func (s *TrainOrderQueryResponseBodyModuleChangeTicketInfoList) GetChangeCoachNo() *string {
	return s.ChangeCoachNo
}

func (s *TrainOrderQueryResponseBodyModuleChangeTicketInfoList) GetChangeGapFee() *float64 {
	return s.ChangeGapFee
}

func (s *TrainOrderQueryResponseBodyModuleChangeTicketInfoList) GetChangeHandlingFee() *float64 {
	return s.ChangeHandlingFee
}

func (s *TrainOrderQueryResponseBodyModuleChangeTicketInfoList) GetChangeSeatNo() *string {
	return s.ChangeSeatNo
}

func (s *TrainOrderQueryResponseBodyModuleChangeTicketInfoList) GetChangeSeatTypeName() *string {
	return s.ChangeSeatTypeName
}

func (s *TrainOrderQueryResponseBodyModuleChangeTicketInfoList) GetChangeServiceFee() *float64 {
	return s.ChangeServiceFee
}

func (s *TrainOrderQueryResponseBodyModuleChangeTicketInfoList) GetChangeTrainNo() *string {
	return s.ChangeTrainNo
}

func (s *TrainOrderQueryResponseBodyModuleChangeTicketInfoList) GetChangeTrainTypeName() *string {
	return s.ChangeTrainTypeName
}

func (s *TrainOrderQueryResponseBodyModuleChangeTicketInfoList) GetCheckInTime() *string {
	return s.CheckInTime
}

func (s *TrainOrderQueryResponseBodyModuleChangeTicketInfoList) GetCheckOutTime() *string {
	return s.CheckOutTime
}

func (s *TrainOrderQueryResponseBodyModuleChangeTicketInfoList) GetEndTime() *string {
	return s.EndTime
}

func (s *TrainOrderQueryResponseBodyModuleChangeTicketInfoList) GetFromStationName() *string {
	return s.FromStationName
}

func (s *TrainOrderQueryResponseBodyModuleChangeTicketInfoList) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *TrainOrderQueryResponseBodyModuleChangeTicketInfoList) GetGmtModify() *string {
	return s.GmtModify
}

func (s *TrainOrderQueryResponseBodyModuleChangeTicketInfoList) GetOriginTicketNo() *string {
	return s.OriginTicketNo
}

func (s *TrainOrderQueryResponseBodyModuleChangeTicketInfoList) GetOutTicketStatus() *string {
	return s.OutTicketStatus
}

func (s *TrainOrderQueryResponseBodyModuleChangeTicketInfoList) GetStartTime() *string {
	return s.StartTime
}

func (s *TrainOrderQueryResponseBodyModuleChangeTicketInfoList) GetTicketNo() *string {
	return s.TicketNo
}

func (s *TrainOrderQueryResponseBodyModuleChangeTicketInfoList) GetToStationName() *string {
	return s.ToStationName
}

func (s *TrainOrderQueryResponseBodyModuleChangeTicketInfoList) SetChangeCoachNo(v string) *TrainOrderQueryResponseBodyModuleChangeTicketInfoList {
	s.ChangeCoachNo = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModuleChangeTicketInfoList) SetChangeGapFee(v float64) *TrainOrderQueryResponseBodyModuleChangeTicketInfoList {
	s.ChangeGapFee = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModuleChangeTicketInfoList) SetChangeHandlingFee(v float64) *TrainOrderQueryResponseBodyModuleChangeTicketInfoList {
	s.ChangeHandlingFee = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModuleChangeTicketInfoList) SetChangeSeatNo(v string) *TrainOrderQueryResponseBodyModuleChangeTicketInfoList {
	s.ChangeSeatNo = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModuleChangeTicketInfoList) SetChangeSeatTypeName(v string) *TrainOrderQueryResponseBodyModuleChangeTicketInfoList {
	s.ChangeSeatTypeName = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModuleChangeTicketInfoList) SetChangeServiceFee(v float64) *TrainOrderQueryResponseBodyModuleChangeTicketInfoList {
	s.ChangeServiceFee = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModuleChangeTicketInfoList) SetChangeTrainNo(v string) *TrainOrderQueryResponseBodyModuleChangeTicketInfoList {
	s.ChangeTrainNo = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModuleChangeTicketInfoList) SetChangeTrainTypeName(v string) *TrainOrderQueryResponseBodyModuleChangeTicketInfoList {
	s.ChangeTrainTypeName = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModuleChangeTicketInfoList) SetCheckInTime(v string) *TrainOrderQueryResponseBodyModuleChangeTicketInfoList {
	s.CheckInTime = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModuleChangeTicketInfoList) SetCheckOutTime(v string) *TrainOrderQueryResponseBodyModuleChangeTicketInfoList {
	s.CheckOutTime = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModuleChangeTicketInfoList) SetEndTime(v string) *TrainOrderQueryResponseBodyModuleChangeTicketInfoList {
	s.EndTime = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModuleChangeTicketInfoList) SetFromStationName(v string) *TrainOrderQueryResponseBodyModuleChangeTicketInfoList {
	s.FromStationName = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModuleChangeTicketInfoList) SetGmtCreate(v string) *TrainOrderQueryResponseBodyModuleChangeTicketInfoList {
	s.GmtCreate = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModuleChangeTicketInfoList) SetGmtModify(v string) *TrainOrderQueryResponseBodyModuleChangeTicketInfoList {
	s.GmtModify = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModuleChangeTicketInfoList) SetOriginTicketNo(v string) *TrainOrderQueryResponseBodyModuleChangeTicketInfoList {
	s.OriginTicketNo = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModuleChangeTicketInfoList) SetOutTicketStatus(v string) *TrainOrderQueryResponseBodyModuleChangeTicketInfoList {
	s.OutTicketStatus = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModuleChangeTicketInfoList) SetStartTime(v string) *TrainOrderQueryResponseBodyModuleChangeTicketInfoList {
	s.StartTime = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModuleChangeTicketInfoList) SetTicketNo(v string) *TrainOrderQueryResponseBodyModuleChangeTicketInfoList {
	s.TicketNo = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModuleChangeTicketInfoList) SetToStationName(v string) *TrainOrderQueryResponseBodyModuleChangeTicketInfoList {
	s.ToStationName = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModuleChangeTicketInfoList) Validate() error {
	return dara.Validate(s)
}

type TrainOrderQueryResponseBodyModuleInvoiceInfo struct {
	// example:
	//
	// 11754
	Id    *int64  `json:"id,omitempty" xml:"id,omitempty"`
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s TrainOrderQueryResponseBodyModuleInvoiceInfo) String() string {
	return dara.Prettify(s)
}

func (s TrainOrderQueryResponseBodyModuleInvoiceInfo) GoString() string {
	return s.String()
}

func (s *TrainOrderQueryResponseBodyModuleInvoiceInfo) GetId() *int64 {
	return s.Id
}

func (s *TrainOrderQueryResponseBodyModuleInvoiceInfo) GetTitle() *string {
	return s.Title
}

func (s *TrainOrderQueryResponseBodyModuleInvoiceInfo) SetId(v int64) *TrainOrderQueryResponseBodyModuleInvoiceInfo {
	s.Id = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModuleInvoiceInfo) SetTitle(v string) *TrainOrderQueryResponseBodyModuleInvoiceInfo {
	s.Title = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModuleInvoiceInfo) Validate() error {
	return dara.Validate(s)
}

type TrainOrderQueryResponseBodyModuleOrderBaseInfo struct {
	// example:
	//
	// 11657
	ApplyId     *string `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
	BtripTitle  *string `json:"btrip_title,omitempty" xml:"btrip_title,omitempty"`
	ContactName *string `json:"contact_name,omitempty" xml:"contact_name,omitempty"`
	CorpId      *string `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
	CorpName    *string `json:"corp_name,omitempty" xml:"corp_name,omitempty"`
	DepartId    *string `json:"depart_id,omitempty" xml:"depart_id,omitempty"`
	DepartName  *string `json:"depart_name,omitempty" xml:"depart_name,omitempty"`
	// 火车票超标审批id
	//
	// example:
	//
	// 1422113021033961000
	ExceedApplyId *string `json:"exceed_apply_id,omitempty" xml:"exceed_apply_id,omitempty"`
	// 火车票超标审批三方id
	//
	// example:
	//
	// 2022113021030003600001715
	ExceedThirdPartApplyId *string `json:"exceed_third_part_apply_id,omitempty" xml:"exceed_third_part_apply_id,omitempty"`
	// example:
	//
	// 2022-05-15T22:27Z
	GmtCreate *string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// example:
	//
	// 2022-05-15T22:27Z
	GmtModify *string `json:"gmt_modify,omitempty" xml:"gmt_modify,omitempty"`
	// example:
	//
	// kaxasevesguikxn123kixnghid
	ItineraryId *string `json:"itinerary_id,omitempty" xml:"itinerary_id,omitempty"`
	// example:
	//
	// 2627694109810885616
	OrderId *int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// example:
	//
	// 4
	OrderStatus *int32 `json:"order_status,omitempty" xml:"order_status,omitempty"`
	// example:
	//
	// CS-EDES9898
	ThirdpartApplyId *string `json:"thirdpart_apply_id,omitempty" xml:"thirdpart_apply_id,omitempty"`
	ThirdpartCorpId  *string `json:"thirdpart_corp_id,omitempty" xml:"thirdpart_corp_id,omitempty"`
	// example:
	//
	// kaxasevesguikxn123kixnghid
	ThirdpartItineraryId *string `json:"thirdpart_itinerary_id,omitempty" xml:"thirdpart_itinerary_id,omitempty"`
	// example:
	//
	// 0
	TripType *int32  `json:"trip_type,omitempty" xml:"trip_type,omitempty"`
	UserId   *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

func (s TrainOrderQueryResponseBodyModuleOrderBaseInfo) String() string {
	return dara.Prettify(s)
}

func (s TrainOrderQueryResponseBodyModuleOrderBaseInfo) GoString() string {
	return s.String()
}

func (s *TrainOrderQueryResponseBodyModuleOrderBaseInfo) GetApplyId() *string {
	return s.ApplyId
}

func (s *TrainOrderQueryResponseBodyModuleOrderBaseInfo) GetBtripTitle() *string {
	return s.BtripTitle
}

func (s *TrainOrderQueryResponseBodyModuleOrderBaseInfo) GetContactName() *string {
	return s.ContactName
}

func (s *TrainOrderQueryResponseBodyModuleOrderBaseInfo) GetCorpId() *string {
	return s.CorpId
}

func (s *TrainOrderQueryResponseBodyModuleOrderBaseInfo) GetCorpName() *string {
	return s.CorpName
}

func (s *TrainOrderQueryResponseBodyModuleOrderBaseInfo) GetDepartId() *string {
	return s.DepartId
}

func (s *TrainOrderQueryResponseBodyModuleOrderBaseInfo) GetDepartName() *string {
	return s.DepartName
}

func (s *TrainOrderQueryResponseBodyModuleOrderBaseInfo) GetExceedApplyId() *string {
	return s.ExceedApplyId
}

func (s *TrainOrderQueryResponseBodyModuleOrderBaseInfo) GetExceedThirdPartApplyId() *string {
	return s.ExceedThirdPartApplyId
}

func (s *TrainOrderQueryResponseBodyModuleOrderBaseInfo) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *TrainOrderQueryResponseBodyModuleOrderBaseInfo) GetGmtModify() *string {
	return s.GmtModify
}

func (s *TrainOrderQueryResponseBodyModuleOrderBaseInfo) GetItineraryId() *string {
	return s.ItineraryId
}

func (s *TrainOrderQueryResponseBodyModuleOrderBaseInfo) GetOrderId() *int64 {
	return s.OrderId
}

func (s *TrainOrderQueryResponseBodyModuleOrderBaseInfo) GetOrderStatus() *int32 {
	return s.OrderStatus
}

func (s *TrainOrderQueryResponseBodyModuleOrderBaseInfo) GetThirdpartApplyId() *string {
	return s.ThirdpartApplyId
}

func (s *TrainOrderQueryResponseBodyModuleOrderBaseInfo) GetThirdpartCorpId() *string {
	return s.ThirdpartCorpId
}

func (s *TrainOrderQueryResponseBodyModuleOrderBaseInfo) GetThirdpartItineraryId() *string {
	return s.ThirdpartItineraryId
}

func (s *TrainOrderQueryResponseBodyModuleOrderBaseInfo) GetTripType() *int32 {
	return s.TripType
}

func (s *TrainOrderQueryResponseBodyModuleOrderBaseInfo) GetUserId() *string {
	return s.UserId
}

func (s *TrainOrderQueryResponseBodyModuleOrderBaseInfo) SetApplyId(v string) *TrainOrderQueryResponseBodyModuleOrderBaseInfo {
	s.ApplyId = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModuleOrderBaseInfo) SetBtripTitle(v string) *TrainOrderQueryResponseBodyModuleOrderBaseInfo {
	s.BtripTitle = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModuleOrderBaseInfo) SetContactName(v string) *TrainOrderQueryResponseBodyModuleOrderBaseInfo {
	s.ContactName = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModuleOrderBaseInfo) SetCorpId(v string) *TrainOrderQueryResponseBodyModuleOrderBaseInfo {
	s.CorpId = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModuleOrderBaseInfo) SetCorpName(v string) *TrainOrderQueryResponseBodyModuleOrderBaseInfo {
	s.CorpName = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModuleOrderBaseInfo) SetDepartId(v string) *TrainOrderQueryResponseBodyModuleOrderBaseInfo {
	s.DepartId = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModuleOrderBaseInfo) SetDepartName(v string) *TrainOrderQueryResponseBodyModuleOrderBaseInfo {
	s.DepartName = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModuleOrderBaseInfo) SetExceedApplyId(v string) *TrainOrderQueryResponseBodyModuleOrderBaseInfo {
	s.ExceedApplyId = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModuleOrderBaseInfo) SetExceedThirdPartApplyId(v string) *TrainOrderQueryResponseBodyModuleOrderBaseInfo {
	s.ExceedThirdPartApplyId = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModuleOrderBaseInfo) SetGmtCreate(v string) *TrainOrderQueryResponseBodyModuleOrderBaseInfo {
	s.GmtCreate = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModuleOrderBaseInfo) SetGmtModify(v string) *TrainOrderQueryResponseBodyModuleOrderBaseInfo {
	s.GmtModify = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModuleOrderBaseInfo) SetItineraryId(v string) *TrainOrderQueryResponseBodyModuleOrderBaseInfo {
	s.ItineraryId = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModuleOrderBaseInfo) SetOrderId(v int64) *TrainOrderQueryResponseBodyModuleOrderBaseInfo {
	s.OrderId = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModuleOrderBaseInfo) SetOrderStatus(v int32) *TrainOrderQueryResponseBodyModuleOrderBaseInfo {
	s.OrderStatus = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModuleOrderBaseInfo) SetThirdpartApplyId(v string) *TrainOrderQueryResponseBodyModuleOrderBaseInfo {
	s.ThirdpartApplyId = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModuleOrderBaseInfo) SetThirdpartCorpId(v string) *TrainOrderQueryResponseBodyModuleOrderBaseInfo {
	s.ThirdpartCorpId = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModuleOrderBaseInfo) SetThirdpartItineraryId(v string) *TrainOrderQueryResponseBodyModuleOrderBaseInfo {
	s.ThirdpartItineraryId = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModuleOrderBaseInfo) SetTripType(v int32) *TrainOrderQueryResponseBodyModuleOrderBaseInfo {
	s.TripType = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModuleOrderBaseInfo) SetUserId(v string) *TrainOrderQueryResponseBodyModuleOrderBaseInfo {
	s.UserId = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModuleOrderBaseInfo) Validate() error {
	return dara.Validate(s)
}

type TrainOrderQueryResponseBodyModulePassengerInfoList struct {
	// example:
	//
	// 11564
	CostCenterId   *int64  `json:"cost_center_id,omitempty" xml:"cost_center_id,omitempty"`
	CostCenterName *string `json:"cost_center_name,omitempty" xml:"cost_center_name,omitempty"`
	// example:
	//
	// 01
	CostCenterNumber *string `json:"cost_center_number,omitempty" xml:"cost_center_number,omitempty"`
	// example:
	//
	// CSP-01
	ProjectCode *string `json:"project_code,omitempty" xml:"project_code,omitempty"`
	// example:
	//
	// 22562
	ProjectId    *int64  `json:"project_id,omitempty" xml:"project_id,omitempty"`
	ProjectTitle *string `json:"project_title,omitempty" xml:"project_title,omitempty"`
	// example:
	//
	// CS-22562
	ThirdpartProjectId *string `json:"thirdpart_project_id,omitempty" xml:"thirdpart_project_id,omitempty"`
	UserId             *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	UserName           *string `json:"user_name,omitempty" xml:"user_name,omitempty"`
	// example:
	//
	// 0
	UserType *int32 `json:"user_type,omitempty" xml:"user_type,omitempty"`
}

func (s TrainOrderQueryResponseBodyModulePassengerInfoList) String() string {
	return dara.Prettify(s)
}

func (s TrainOrderQueryResponseBodyModulePassengerInfoList) GoString() string {
	return s.String()
}

func (s *TrainOrderQueryResponseBodyModulePassengerInfoList) GetCostCenterId() *int64 {
	return s.CostCenterId
}

func (s *TrainOrderQueryResponseBodyModulePassengerInfoList) GetCostCenterName() *string {
	return s.CostCenterName
}

func (s *TrainOrderQueryResponseBodyModulePassengerInfoList) GetCostCenterNumber() *string {
	return s.CostCenterNumber
}

func (s *TrainOrderQueryResponseBodyModulePassengerInfoList) GetProjectCode() *string {
	return s.ProjectCode
}

func (s *TrainOrderQueryResponseBodyModulePassengerInfoList) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *TrainOrderQueryResponseBodyModulePassengerInfoList) GetProjectTitle() *string {
	return s.ProjectTitle
}

func (s *TrainOrderQueryResponseBodyModulePassengerInfoList) GetThirdpartProjectId() *string {
	return s.ThirdpartProjectId
}

func (s *TrainOrderQueryResponseBodyModulePassengerInfoList) GetUserId() *string {
	return s.UserId
}

func (s *TrainOrderQueryResponseBodyModulePassengerInfoList) GetUserName() *string {
	return s.UserName
}

func (s *TrainOrderQueryResponseBodyModulePassengerInfoList) GetUserType() *int32 {
	return s.UserType
}

func (s *TrainOrderQueryResponseBodyModulePassengerInfoList) SetCostCenterId(v int64) *TrainOrderQueryResponseBodyModulePassengerInfoList {
	s.CostCenterId = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModulePassengerInfoList) SetCostCenterName(v string) *TrainOrderQueryResponseBodyModulePassengerInfoList {
	s.CostCenterName = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModulePassengerInfoList) SetCostCenterNumber(v string) *TrainOrderQueryResponseBodyModulePassengerInfoList {
	s.CostCenterNumber = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModulePassengerInfoList) SetProjectCode(v string) *TrainOrderQueryResponseBodyModulePassengerInfoList {
	s.ProjectCode = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModulePassengerInfoList) SetProjectId(v int64) *TrainOrderQueryResponseBodyModulePassengerInfoList {
	s.ProjectId = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModulePassengerInfoList) SetProjectTitle(v string) *TrainOrderQueryResponseBodyModulePassengerInfoList {
	s.ProjectTitle = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModulePassengerInfoList) SetThirdpartProjectId(v string) *TrainOrderQueryResponseBodyModulePassengerInfoList {
	s.ThirdpartProjectId = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModulePassengerInfoList) SetUserId(v string) *TrainOrderQueryResponseBodyModulePassengerInfoList {
	s.UserId = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModulePassengerInfoList) SetUserName(v string) *TrainOrderQueryResponseBodyModulePassengerInfoList {
	s.UserName = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModulePassengerInfoList) SetUserType(v int32) *TrainOrderQueryResponseBodyModulePassengerInfoList {
	s.UserType = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModulePassengerInfoList) Validate() error {
	return dara.Validate(s)
}

type TrainOrderQueryResponseBodyModulePriceInfoList struct {
	// example:
	//
	// 1
	CategoryCode *int32 `json:"category_code,omitempty" xml:"category_code,omitempty"`
	// example:
	//
	// 2022-05-15T22:27Z
	GmtCreate     *string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	PassengerName *string `json:"passenger_name,omitempty" xml:"passenger_name,omitempty"`
	// example:
	//
	// 1
	PayType *int32 `json:"pay_type,omitempty" xml:"pay_type,omitempty"`
	// example:
	//
	// 0
	Price *float64 `json:"price,omitempty" xml:"price,omitempty"`
	// example:
	//
	// 12312312001
	TradeId *string `json:"trade_id,omitempty" xml:"trade_id,omitempty"`
	// example:
	//
	// 1
	Type *int32 `json:"type,omitempty" xml:"type,omitempty"`
}

func (s TrainOrderQueryResponseBodyModulePriceInfoList) String() string {
	return dara.Prettify(s)
}

func (s TrainOrderQueryResponseBodyModulePriceInfoList) GoString() string {
	return s.String()
}

func (s *TrainOrderQueryResponseBodyModulePriceInfoList) GetCategoryCode() *int32 {
	return s.CategoryCode
}

func (s *TrainOrderQueryResponseBodyModulePriceInfoList) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *TrainOrderQueryResponseBodyModulePriceInfoList) GetPassengerName() *string {
	return s.PassengerName
}

func (s *TrainOrderQueryResponseBodyModulePriceInfoList) GetPayType() *int32 {
	return s.PayType
}

func (s *TrainOrderQueryResponseBodyModulePriceInfoList) GetPrice() *float64 {
	return s.Price
}

func (s *TrainOrderQueryResponseBodyModulePriceInfoList) GetTradeId() *string {
	return s.TradeId
}

func (s *TrainOrderQueryResponseBodyModulePriceInfoList) GetType() *int32 {
	return s.Type
}

func (s *TrainOrderQueryResponseBodyModulePriceInfoList) SetCategoryCode(v int32) *TrainOrderQueryResponseBodyModulePriceInfoList {
	s.CategoryCode = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModulePriceInfoList) SetGmtCreate(v string) *TrainOrderQueryResponseBodyModulePriceInfoList {
	s.GmtCreate = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModulePriceInfoList) SetPassengerName(v string) *TrainOrderQueryResponseBodyModulePriceInfoList {
	s.PassengerName = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModulePriceInfoList) SetPayType(v int32) *TrainOrderQueryResponseBodyModulePriceInfoList {
	s.PayType = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModulePriceInfoList) SetPrice(v float64) *TrainOrderQueryResponseBodyModulePriceInfoList {
	s.Price = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModulePriceInfoList) SetTradeId(v string) *TrainOrderQueryResponseBodyModulePriceInfoList {
	s.TradeId = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModulePriceInfoList) SetType(v int32) *TrainOrderQueryResponseBodyModulePriceInfoList {
	s.Type = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModulePriceInfoList) Validate() error {
	return dara.Validate(s)
}

type TrainOrderQueryResponseBodyModuleRefundTicketInfoList struct {
	// example:
	//
	// 2022-05-15T22:27Z
	GmtCreate *string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// example:
	//
	// 2022-05-15T22:27Z
	GmtModify *string `json:"gmt_modify,omitempty" xml:"gmt_modify,omitempty"`
	// example:
	//
	// 10
	RefundFee *float64 `json:"refund_fee,omitempty" xml:"refund_fee,omitempty"`
	// example:
	//
	// 10
	RefundServiceFee *float64 `json:"refund_service_fee,omitempty" xml:"refund_service_fee,omitempty"`
	// example:
	//
	// CS987JKDF
	TicketNo *string `json:"ticket_no,omitempty" xml:"ticket_no,omitempty"`
}

func (s TrainOrderQueryResponseBodyModuleRefundTicketInfoList) String() string {
	return dara.Prettify(s)
}

func (s TrainOrderQueryResponseBodyModuleRefundTicketInfoList) GoString() string {
	return s.String()
}

func (s *TrainOrderQueryResponseBodyModuleRefundTicketInfoList) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *TrainOrderQueryResponseBodyModuleRefundTicketInfoList) GetGmtModify() *string {
	return s.GmtModify
}

func (s *TrainOrderQueryResponseBodyModuleRefundTicketInfoList) GetRefundFee() *float64 {
	return s.RefundFee
}

func (s *TrainOrderQueryResponseBodyModuleRefundTicketInfoList) GetRefundServiceFee() *float64 {
	return s.RefundServiceFee
}

func (s *TrainOrderQueryResponseBodyModuleRefundTicketInfoList) GetTicketNo() *string {
	return s.TicketNo
}

func (s *TrainOrderQueryResponseBodyModuleRefundTicketInfoList) SetGmtCreate(v string) *TrainOrderQueryResponseBodyModuleRefundTicketInfoList {
	s.GmtCreate = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModuleRefundTicketInfoList) SetGmtModify(v string) *TrainOrderQueryResponseBodyModuleRefundTicketInfoList {
	s.GmtModify = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModuleRefundTicketInfoList) SetRefundFee(v float64) *TrainOrderQueryResponseBodyModuleRefundTicketInfoList {
	s.RefundFee = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModuleRefundTicketInfoList) SetRefundServiceFee(v float64) *TrainOrderQueryResponseBodyModuleRefundTicketInfoList {
	s.RefundServiceFee = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModuleRefundTicketInfoList) SetTicketNo(v string) *TrainOrderQueryResponseBodyModuleRefundTicketInfoList {
	s.TicketNo = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModuleRefundTicketInfoList) Validate() error {
	return dara.Validate(s)
}

type TrainOrderQueryResponseBodyModuleTicketInfoList struct {
	// example:
	//
	// false
	Changed *bool `json:"changed,omitempty" xml:"changed,omitempty"`
	// example:
	//
	// 2022-05-15T22:27Z
	CheckInTime *string `json:"check_in_time,omitempty" xml:"check_in_time,omitempty"`
	// example:
	//
	// 2022-05-15T22:27Z
	CheckOutTime *string `json:"check_out_time,omitempty" xml:"check_out_time,omitempty"`
	// example:
	//
	// 01
	CoachNo *string `json:"coach_no,omitempty" xml:"coach_no,omitempty"`
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
	// 2022-05-15T22:27Z
	GmtModify *string `json:"gmt_modify,omitempty" xml:"gmt_modify,omitempty"`
	// example:
	//
	// m
	OutTicketStatus *string `json:"out_ticket_status,omitempty" xml:"out_ticket_status,omitempty"`
	// example:
	//
	// 1
	PayType *int32 `json:"pay_type,omitempty" xml:"pay_type,omitempty"`
	// example:
	//
	// A001
	SeatNo       *string `json:"seat_no,omitempty" xml:"seat_no,omitempty"`
	SeatTypeName *string `json:"seat_type_name,omitempty" xml:"seat_type_name,omitempty"`
	// example:
	//
	// 10
	ServiceFee *float64 `json:"service_fee,omitempty" xml:"service_fee,omitempty"`
	// example:
	//
	// 2022-05-15T22:27Z
	StartTime *string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// example:
	//
	// CS987JKDF
	TicketNo *string `json:"ticket_no,omitempty" xml:"ticket_no,omitempty"`
	// example:
	//
	// 100
	TicketPrice *float64 `json:"ticket_price,omitempty" xml:"ticket_price,omitempty"`
	// example:
	//
	// 1
	TicketStatus  *int32  `json:"ticket_status,omitempty" xml:"ticket_status,omitempty"`
	TrainTypeName *string `json:"train_type_name,omitempty" xml:"train_type_name,omitempty"`
	UserId        *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

func (s TrainOrderQueryResponseBodyModuleTicketInfoList) String() string {
	return dara.Prettify(s)
}

func (s TrainOrderQueryResponseBodyModuleTicketInfoList) GoString() string {
	return s.String()
}

func (s *TrainOrderQueryResponseBodyModuleTicketInfoList) GetChanged() *bool {
	return s.Changed
}

func (s *TrainOrderQueryResponseBodyModuleTicketInfoList) GetCheckInTime() *string {
	return s.CheckInTime
}

func (s *TrainOrderQueryResponseBodyModuleTicketInfoList) GetCheckOutTime() *string {
	return s.CheckOutTime
}

func (s *TrainOrderQueryResponseBodyModuleTicketInfoList) GetCoachNo() *string {
	return s.CoachNo
}

func (s *TrainOrderQueryResponseBodyModuleTicketInfoList) GetEndTime() *string {
	return s.EndTime
}

func (s *TrainOrderQueryResponseBodyModuleTicketInfoList) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *TrainOrderQueryResponseBodyModuleTicketInfoList) GetGmtModify() *string {
	return s.GmtModify
}

func (s *TrainOrderQueryResponseBodyModuleTicketInfoList) GetOutTicketStatus() *string {
	return s.OutTicketStatus
}

func (s *TrainOrderQueryResponseBodyModuleTicketInfoList) GetPayType() *int32 {
	return s.PayType
}

func (s *TrainOrderQueryResponseBodyModuleTicketInfoList) GetSeatNo() *string {
	return s.SeatNo
}

func (s *TrainOrderQueryResponseBodyModuleTicketInfoList) GetSeatTypeName() *string {
	return s.SeatTypeName
}

func (s *TrainOrderQueryResponseBodyModuleTicketInfoList) GetServiceFee() *float64 {
	return s.ServiceFee
}

func (s *TrainOrderQueryResponseBodyModuleTicketInfoList) GetStartTime() *string {
	return s.StartTime
}

func (s *TrainOrderQueryResponseBodyModuleTicketInfoList) GetTicketNo() *string {
	return s.TicketNo
}

func (s *TrainOrderQueryResponseBodyModuleTicketInfoList) GetTicketPrice() *float64 {
	return s.TicketPrice
}

func (s *TrainOrderQueryResponseBodyModuleTicketInfoList) GetTicketStatus() *int32 {
	return s.TicketStatus
}

func (s *TrainOrderQueryResponseBodyModuleTicketInfoList) GetTrainTypeName() *string {
	return s.TrainTypeName
}

func (s *TrainOrderQueryResponseBodyModuleTicketInfoList) GetUserId() *string {
	return s.UserId
}

func (s *TrainOrderQueryResponseBodyModuleTicketInfoList) SetChanged(v bool) *TrainOrderQueryResponseBodyModuleTicketInfoList {
	s.Changed = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModuleTicketInfoList) SetCheckInTime(v string) *TrainOrderQueryResponseBodyModuleTicketInfoList {
	s.CheckInTime = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModuleTicketInfoList) SetCheckOutTime(v string) *TrainOrderQueryResponseBodyModuleTicketInfoList {
	s.CheckOutTime = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModuleTicketInfoList) SetCoachNo(v string) *TrainOrderQueryResponseBodyModuleTicketInfoList {
	s.CoachNo = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModuleTicketInfoList) SetEndTime(v string) *TrainOrderQueryResponseBodyModuleTicketInfoList {
	s.EndTime = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModuleTicketInfoList) SetGmtCreate(v string) *TrainOrderQueryResponseBodyModuleTicketInfoList {
	s.GmtCreate = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModuleTicketInfoList) SetGmtModify(v string) *TrainOrderQueryResponseBodyModuleTicketInfoList {
	s.GmtModify = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModuleTicketInfoList) SetOutTicketStatus(v string) *TrainOrderQueryResponseBodyModuleTicketInfoList {
	s.OutTicketStatus = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModuleTicketInfoList) SetPayType(v int32) *TrainOrderQueryResponseBodyModuleTicketInfoList {
	s.PayType = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModuleTicketInfoList) SetSeatNo(v string) *TrainOrderQueryResponseBodyModuleTicketInfoList {
	s.SeatNo = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModuleTicketInfoList) SetSeatTypeName(v string) *TrainOrderQueryResponseBodyModuleTicketInfoList {
	s.SeatTypeName = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModuleTicketInfoList) SetServiceFee(v float64) *TrainOrderQueryResponseBodyModuleTicketInfoList {
	s.ServiceFee = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModuleTicketInfoList) SetStartTime(v string) *TrainOrderQueryResponseBodyModuleTicketInfoList {
	s.StartTime = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModuleTicketInfoList) SetTicketNo(v string) *TrainOrderQueryResponseBodyModuleTicketInfoList {
	s.TicketNo = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModuleTicketInfoList) SetTicketPrice(v float64) *TrainOrderQueryResponseBodyModuleTicketInfoList {
	s.TicketPrice = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModuleTicketInfoList) SetTicketStatus(v int32) *TrainOrderQueryResponseBodyModuleTicketInfoList {
	s.TicketStatus = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModuleTicketInfoList) SetTrainTypeName(v string) *TrainOrderQueryResponseBodyModuleTicketInfoList {
	s.TrainTypeName = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModuleTicketInfoList) SetUserId(v string) *TrainOrderQueryResponseBodyModuleTicketInfoList {
	s.UserId = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModuleTicketInfoList) Validate() error {
	return dara.Validate(s)
}

type TrainOrderQueryResponseBodyModuleTrainInfo struct {
	// example:
	//
	// 2022-05-15T22:27Z
	ArrTime *string `json:"arr_time,omitempty" xml:"arr_time,omitempty"`
	// example:
	//
	// 2022-05-15T22:27Z
	DepTime         *string `json:"dep_time,omitempty" xml:"dep_time,omitempty"`
	FromCityAdCode  *string `json:"from_city_ad_code,omitempty" xml:"from_city_ad_code,omitempty"`
	FromStationName *string `json:"from_station_name,omitempty" xml:"from_station_name,omitempty"`
	// example:
	//
	// 100
	RunTime       *int64  `json:"run_time,omitempty" xml:"run_time,omitempty"`
	ToCityAdCode  *string `json:"to_city_ad_code,omitempty" xml:"to_city_ad_code,omitempty"`
	ToStationName *string `json:"to_station_name,omitempty" xml:"to_station_name,omitempty"`
	// example:
	//
	// CS-150
	TrainNo *string `json:"train_no,omitempty" xml:"train_no,omitempty"`
}

func (s TrainOrderQueryResponseBodyModuleTrainInfo) String() string {
	return dara.Prettify(s)
}

func (s TrainOrderQueryResponseBodyModuleTrainInfo) GoString() string {
	return s.String()
}

func (s *TrainOrderQueryResponseBodyModuleTrainInfo) GetArrTime() *string {
	return s.ArrTime
}

func (s *TrainOrderQueryResponseBodyModuleTrainInfo) GetDepTime() *string {
	return s.DepTime
}

func (s *TrainOrderQueryResponseBodyModuleTrainInfo) GetFromCityAdCode() *string {
	return s.FromCityAdCode
}

func (s *TrainOrderQueryResponseBodyModuleTrainInfo) GetFromStationName() *string {
	return s.FromStationName
}

func (s *TrainOrderQueryResponseBodyModuleTrainInfo) GetRunTime() *int64 {
	return s.RunTime
}

func (s *TrainOrderQueryResponseBodyModuleTrainInfo) GetToCityAdCode() *string {
	return s.ToCityAdCode
}

func (s *TrainOrderQueryResponseBodyModuleTrainInfo) GetToStationName() *string {
	return s.ToStationName
}

func (s *TrainOrderQueryResponseBodyModuleTrainInfo) GetTrainNo() *string {
	return s.TrainNo
}

func (s *TrainOrderQueryResponseBodyModuleTrainInfo) SetArrTime(v string) *TrainOrderQueryResponseBodyModuleTrainInfo {
	s.ArrTime = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModuleTrainInfo) SetDepTime(v string) *TrainOrderQueryResponseBodyModuleTrainInfo {
	s.DepTime = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModuleTrainInfo) SetFromCityAdCode(v string) *TrainOrderQueryResponseBodyModuleTrainInfo {
	s.FromCityAdCode = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModuleTrainInfo) SetFromStationName(v string) *TrainOrderQueryResponseBodyModuleTrainInfo {
	s.FromStationName = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModuleTrainInfo) SetRunTime(v int64) *TrainOrderQueryResponseBodyModuleTrainInfo {
	s.RunTime = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModuleTrainInfo) SetToCityAdCode(v string) *TrainOrderQueryResponseBodyModuleTrainInfo {
	s.ToCityAdCode = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModuleTrainInfo) SetToStationName(v string) *TrainOrderQueryResponseBodyModuleTrainInfo {
	s.ToStationName = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModuleTrainInfo) SetTrainNo(v string) *TrainOrderQueryResponseBodyModuleTrainInfo {
	s.TrainNo = &v
	return s
}

func (s *TrainOrderQueryResponseBodyModuleTrainInfo) Validate() error {
	return dara.Validate(s)
}
