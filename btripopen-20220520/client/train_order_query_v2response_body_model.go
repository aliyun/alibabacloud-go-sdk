// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTrainOrderQueryV2ResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *TrainOrderQueryV2ResponseBody
	GetCode() *string
	SetMessage(v string) *TrainOrderQueryV2ResponseBody
	GetMessage() *string
	SetModule(v *TrainOrderQueryV2ResponseBodyModule) *TrainOrderQueryV2ResponseBody
	GetModule() *TrainOrderQueryV2ResponseBodyModule
	SetRequestId(v string) *TrainOrderQueryV2ResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *TrainOrderQueryV2ResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *TrainOrderQueryV2ResponseBody
	GetTraceId() *string
}

type TrainOrderQueryV2ResponseBody struct {
	// example:
	//
	// 0
	Code    *string                              `json:"code,omitempty" xml:"code,omitempty"`
	Message *string                              `json:"message,omitempty" xml:"message,omitempty"`
	Module  *TrainOrderQueryV2ResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	// requestId
	//
	// example:
	//
	// 1213ds1d
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// traceId
	//
	// example:
	//
	// 123412dcdsac sd
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s TrainOrderQueryV2ResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TrainOrderQueryV2ResponseBody) GoString() string {
	return s.String()
}

func (s *TrainOrderQueryV2ResponseBody) GetCode() *string {
	return s.Code
}

func (s *TrainOrderQueryV2ResponseBody) GetMessage() *string {
	return s.Message
}

func (s *TrainOrderQueryV2ResponseBody) GetModule() *TrainOrderQueryV2ResponseBodyModule {
	return s.Module
}

func (s *TrainOrderQueryV2ResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TrainOrderQueryV2ResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *TrainOrderQueryV2ResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *TrainOrderQueryV2ResponseBody) SetCode(v string) *TrainOrderQueryV2ResponseBody {
	s.Code = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBody) SetMessage(v string) *TrainOrderQueryV2ResponseBody {
	s.Message = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBody) SetModule(v *TrainOrderQueryV2ResponseBodyModule) *TrainOrderQueryV2ResponseBody {
	s.Module = v
	return s
}

func (s *TrainOrderQueryV2ResponseBody) SetRequestId(v string) *TrainOrderQueryV2ResponseBody {
	s.RequestId = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBody) SetSuccess(v bool) *TrainOrderQueryV2ResponseBody {
	s.Success = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBody) SetTraceId(v string) *TrainOrderQueryV2ResponseBody {
	s.TraceId = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBody) Validate() error {
	return dara.Validate(s)
}

type TrainOrderQueryV2ResponseBodyModule struct {
	ChangeTicketInfoList []*TrainOrderQueryV2ResponseBodyModuleChangeTicketInfoList `json:"change_ticket_info_list,omitempty" xml:"change_ticket_info_list,omitempty" type:"Repeated"`
	InvoiceInfo          *TrainOrderQueryV2ResponseBodyModuleInvoiceInfo            `json:"invoice_info,omitempty" xml:"invoice_info,omitempty" type:"Struct"`
	OrderBaseInfo        *TrainOrderQueryV2ResponseBodyModuleOrderBaseInfo          `json:"order_base_info,omitempty" xml:"order_base_info,omitempty" type:"Struct"`
	PassengerInfoList    []*TrainOrderQueryV2ResponseBodyModulePassengerInfoList    `json:"passenger_info_list,omitempty" xml:"passenger_info_list,omitempty" type:"Repeated"`
	PriceInfoList        []*TrainOrderQueryV2ResponseBodyModulePriceInfoList        `json:"price_info_list,omitempty" xml:"price_info_list,omitempty" type:"Repeated"`
	RefundTicketInfoList []*TrainOrderQueryV2ResponseBodyModuleRefundTicketInfoList `json:"refund_ticket_info_list,omitempty" xml:"refund_ticket_info_list,omitempty" type:"Repeated"`
	TrainOrderInfo       *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfo         `json:"train_order_info,omitempty" xml:"train_order_info,omitempty" type:"Struct"`
}

func (s TrainOrderQueryV2ResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s TrainOrderQueryV2ResponseBodyModule) GoString() string {
	return s.String()
}

func (s *TrainOrderQueryV2ResponseBodyModule) GetChangeTicketInfoList() []*TrainOrderQueryV2ResponseBodyModuleChangeTicketInfoList {
	return s.ChangeTicketInfoList
}

func (s *TrainOrderQueryV2ResponseBodyModule) GetInvoiceInfo() *TrainOrderQueryV2ResponseBodyModuleInvoiceInfo {
	return s.InvoiceInfo
}

func (s *TrainOrderQueryV2ResponseBodyModule) GetOrderBaseInfo() *TrainOrderQueryV2ResponseBodyModuleOrderBaseInfo {
	return s.OrderBaseInfo
}

func (s *TrainOrderQueryV2ResponseBodyModule) GetPassengerInfoList() []*TrainOrderQueryV2ResponseBodyModulePassengerInfoList {
	return s.PassengerInfoList
}

func (s *TrainOrderQueryV2ResponseBodyModule) GetPriceInfoList() []*TrainOrderQueryV2ResponseBodyModulePriceInfoList {
	return s.PriceInfoList
}

func (s *TrainOrderQueryV2ResponseBodyModule) GetRefundTicketInfoList() []*TrainOrderQueryV2ResponseBodyModuleRefundTicketInfoList {
	return s.RefundTicketInfoList
}

func (s *TrainOrderQueryV2ResponseBodyModule) GetTrainOrderInfo() *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfo {
	return s.TrainOrderInfo
}

func (s *TrainOrderQueryV2ResponseBodyModule) SetChangeTicketInfoList(v []*TrainOrderQueryV2ResponseBodyModuleChangeTicketInfoList) *TrainOrderQueryV2ResponseBodyModule {
	s.ChangeTicketInfoList = v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModule) SetInvoiceInfo(v *TrainOrderQueryV2ResponseBodyModuleInvoiceInfo) *TrainOrderQueryV2ResponseBodyModule {
	s.InvoiceInfo = v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModule) SetOrderBaseInfo(v *TrainOrderQueryV2ResponseBodyModuleOrderBaseInfo) *TrainOrderQueryV2ResponseBodyModule {
	s.OrderBaseInfo = v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModule) SetPassengerInfoList(v []*TrainOrderQueryV2ResponseBodyModulePassengerInfoList) *TrainOrderQueryV2ResponseBodyModule {
	s.PassengerInfoList = v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModule) SetPriceInfoList(v []*TrainOrderQueryV2ResponseBodyModulePriceInfoList) *TrainOrderQueryV2ResponseBodyModule {
	s.PriceInfoList = v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModule) SetRefundTicketInfoList(v []*TrainOrderQueryV2ResponseBodyModuleRefundTicketInfoList) *TrainOrderQueryV2ResponseBodyModule {
	s.RefundTicketInfoList = v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModule) SetTrainOrderInfo(v *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfo) *TrainOrderQueryV2ResponseBodyModule {
	s.TrainOrderInfo = v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModule) Validate() error {
	return dara.Validate(s)
}

type TrainOrderQueryV2ResponseBodyModuleChangeTicketInfoList struct {
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
	ChangeOrderId     *string  `json:"change_order_id,omitempty" xml:"change_order_id,omitempty"`
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
	CheckInTime         *string `json:"check_in_time,omitempty" xml:"check_in_time,omitempty"`
	CheckOutTime        *string `json:"check_out_time,omitempty" xml:"check_out_time,omitempty"`
	EndTime             *string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	FromCityName        *string `json:"from_city_name,omitempty" xml:"from_city_name,omitempty"`
	FromStationName     *string `json:"from_station_name,omitempty" xml:"from_station_name,omitempty"`
	GmtCreate           *string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	GmtModify           *string `json:"gmt_modify,omitempty" xml:"gmt_modify,omitempty"`
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
	// 0
	SegmentIndex  *int32  `json:"segment_index,omitempty" xml:"segment_index,omitempty"`
	ShortTicketNo *string `json:"short_ticket_no,omitempty" xml:"short_ticket_no,omitempty"`
	// example:
	//
	// 2022-05-15T22:27Z
	StartTime      *string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	TicketEntrance *string `json:"ticket_entrance,omitempty" xml:"ticket_entrance,omitempty"`
	// example:
	//
	// CS987JKDF
	TicketNo      *string `json:"ticket_no,omitempty" xml:"ticket_no,omitempty"`
	TicketStatus  *int32  `json:"ticket_status,omitempty" xml:"ticket_status,omitempty"`
	ToCityName    *string `json:"to_city_name,omitempty" xml:"to_city_name,omitempty"`
	ToStationName *string `json:"to_station_name,omitempty" xml:"to_station_name,omitempty"`
	// example:
	//
	// 0
	UseTicket *string `json:"use_ticket,omitempty" xml:"use_ticket,omitempty"`
	// example:
	//
	// 12312
	UserId *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

func (s TrainOrderQueryV2ResponseBodyModuleChangeTicketInfoList) String() string {
	return dara.Prettify(s)
}

func (s TrainOrderQueryV2ResponseBodyModuleChangeTicketInfoList) GoString() string {
	return s.String()
}

func (s *TrainOrderQueryV2ResponseBodyModuleChangeTicketInfoList) GetChangeCoachNo() *string {
	return s.ChangeCoachNo
}

func (s *TrainOrderQueryV2ResponseBodyModuleChangeTicketInfoList) GetChangeGapFee() *float64 {
	return s.ChangeGapFee
}

func (s *TrainOrderQueryV2ResponseBodyModuleChangeTicketInfoList) GetChangeHandlingFee() *float64 {
	return s.ChangeHandlingFee
}

func (s *TrainOrderQueryV2ResponseBodyModuleChangeTicketInfoList) GetChangeOrderId() *string {
	return s.ChangeOrderId
}

func (s *TrainOrderQueryV2ResponseBodyModuleChangeTicketInfoList) GetChangeSeatNo() *string {
	return s.ChangeSeatNo
}

func (s *TrainOrderQueryV2ResponseBodyModuleChangeTicketInfoList) GetChangeSeatTypeName() *string {
	return s.ChangeSeatTypeName
}

func (s *TrainOrderQueryV2ResponseBodyModuleChangeTicketInfoList) GetChangeServiceFee() *float64 {
	return s.ChangeServiceFee
}

func (s *TrainOrderQueryV2ResponseBodyModuleChangeTicketInfoList) GetChangeTrainNo() *string {
	return s.ChangeTrainNo
}

func (s *TrainOrderQueryV2ResponseBodyModuleChangeTicketInfoList) GetChangeTrainTypeName() *string {
	return s.ChangeTrainTypeName
}

func (s *TrainOrderQueryV2ResponseBodyModuleChangeTicketInfoList) GetCheckInTime() *string {
	return s.CheckInTime
}

func (s *TrainOrderQueryV2ResponseBodyModuleChangeTicketInfoList) GetCheckOutTime() *string {
	return s.CheckOutTime
}

func (s *TrainOrderQueryV2ResponseBodyModuleChangeTicketInfoList) GetEndTime() *string {
	return s.EndTime
}

func (s *TrainOrderQueryV2ResponseBodyModuleChangeTicketInfoList) GetFromCityName() *string {
	return s.FromCityName
}

func (s *TrainOrderQueryV2ResponseBodyModuleChangeTicketInfoList) GetFromStationName() *string {
	return s.FromStationName
}

func (s *TrainOrderQueryV2ResponseBodyModuleChangeTicketInfoList) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *TrainOrderQueryV2ResponseBodyModuleChangeTicketInfoList) GetGmtModify() *string {
	return s.GmtModify
}

func (s *TrainOrderQueryV2ResponseBodyModuleChangeTicketInfoList) GetOriginTicketNo() *string {
	return s.OriginTicketNo
}

func (s *TrainOrderQueryV2ResponseBodyModuleChangeTicketInfoList) GetOutTicketStatus() *string {
	return s.OutTicketStatus
}

func (s *TrainOrderQueryV2ResponseBodyModuleChangeTicketInfoList) GetSegmentIndex() *int32 {
	return s.SegmentIndex
}

func (s *TrainOrderQueryV2ResponseBodyModuleChangeTicketInfoList) GetShortTicketNo() *string {
	return s.ShortTicketNo
}

func (s *TrainOrderQueryV2ResponseBodyModuleChangeTicketInfoList) GetStartTime() *string {
	return s.StartTime
}

func (s *TrainOrderQueryV2ResponseBodyModuleChangeTicketInfoList) GetTicketEntrance() *string {
	return s.TicketEntrance
}

func (s *TrainOrderQueryV2ResponseBodyModuleChangeTicketInfoList) GetTicketNo() *string {
	return s.TicketNo
}

func (s *TrainOrderQueryV2ResponseBodyModuleChangeTicketInfoList) GetTicketStatus() *int32 {
	return s.TicketStatus
}

func (s *TrainOrderQueryV2ResponseBodyModuleChangeTicketInfoList) GetToCityName() *string {
	return s.ToCityName
}

func (s *TrainOrderQueryV2ResponseBodyModuleChangeTicketInfoList) GetToStationName() *string {
	return s.ToStationName
}

func (s *TrainOrderQueryV2ResponseBodyModuleChangeTicketInfoList) GetUseTicket() *string {
	return s.UseTicket
}

func (s *TrainOrderQueryV2ResponseBodyModuleChangeTicketInfoList) GetUserId() *string {
	return s.UserId
}

func (s *TrainOrderQueryV2ResponseBodyModuleChangeTicketInfoList) SetChangeCoachNo(v string) *TrainOrderQueryV2ResponseBodyModuleChangeTicketInfoList {
	s.ChangeCoachNo = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModuleChangeTicketInfoList) SetChangeGapFee(v float64) *TrainOrderQueryV2ResponseBodyModuleChangeTicketInfoList {
	s.ChangeGapFee = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModuleChangeTicketInfoList) SetChangeHandlingFee(v float64) *TrainOrderQueryV2ResponseBodyModuleChangeTicketInfoList {
	s.ChangeHandlingFee = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModuleChangeTicketInfoList) SetChangeOrderId(v string) *TrainOrderQueryV2ResponseBodyModuleChangeTicketInfoList {
	s.ChangeOrderId = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModuleChangeTicketInfoList) SetChangeSeatNo(v string) *TrainOrderQueryV2ResponseBodyModuleChangeTicketInfoList {
	s.ChangeSeatNo = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModuleChangeTicketInfoList) SetChangeSeatTypeName(v string) *TrainOrderQueryV2ResponseBodyModuleChangeTicketInfoList {
	s.ChangeSeatTypeName = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModuleChangeTicketInfoList) SetChangeServiceFee(v float64) *TrainOrderQueryV2ResponseBodyModuleChangeTicketInfoList {
	s.ChangeServiceFee = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModuleChangeTicketInfoList) SetChangeTrainNo(v string) *TrainOrderQueryV2ResponseBodyModuleChangeTicketInfoList {
	s.ChangeTrainNo = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModuleChangeTicketInfoList) SetChangeTrainTypeName(v string) *TrainOrderQueryV2ResponseBodyModuleChangeTicketInfoList {
	s.ChangeTrainTypeName = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModuleChangeTicketInfoList) SetCheckInTime(v string) *TrainOrderQueryV2ResponseBodyModuleChangeTicketInfoList {
	s.CheckInTime = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModuleChangeTicketInfoList) SetCheckOutTime(v string) *TrainOrderQueryV2ResponseBodyModuleChangeTicketInfoList {
	s.CheckOutTime = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModuleChangeTicketInfoList) SetEndTime(v string) *TrainOrderQueryV2ResponseBodyModuleChangeTicketInfoList {
	s.EndTime = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModuleChangeTicketInfoList) SetFromCityName(v string) *TrainOrderQueryV2ResponseBodyModuleChangeTicketInfoList {
	s.FromCityName = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModuleChangeTicketInfoList) SetFromStationName(v string) *TrainOrderQueryV2ResponseBodyModuleChangeTicketInfoList {
	s.FromStationName = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModuleChangeTicketInfoList) SetGmtCreate(v string) *TrainOrderQueryV2ResponseBodyModuleChangeTicketInfoList {
	s.GmtCreate = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModuleChangeTicketInfoList) SetGmtModify(v string) *TrainOrderQueryV2ResponseBodyModuleChangeTicketInfoList {
	s.GmtModify = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModuleChangeTicketInfoList) SetOriginTicketNo(v string) *TrainOrderQueryV2ResponseBodyModuleChangeTicketInfoList {
	s.OriginTicketNo = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModuleChangeTicketInfoList) SetOutTicketStatus(v string) *TrainOrderQueryV2ResponseBodyModuleChangeTicketInfoList {
	s.OutTicketStatus = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModuleChangeTicketInfoList) SetSegmentIndex(v int32) *TrainOrderQueryV2ResponseBodyModuleChangeTicketInfoList {
	s.SegmentIndex = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModuleChangeTicketInfoList) SetShortTicketNo(v string) *TrainOrderQueryV2ResponseBodyModuleChangeTicketInfoList {
	s.ShortTicketNo = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModuleChangeTicketInfoList) SetStartTime(v string) *TrainOrderQueryV2ResponseBodyModuleChangeTicketInfoList {
	s.StartTime = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModuleChangeTicketInfoList) SetTicketEntrance(v string) *TrainOrderQueryV2ResponseBodyModuleChangeTicketInfoList {
	s.TicketEntrance = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModuleChangeTicketInfoList) SetTicketNo(v string) *TrainOrderQueryV2ResponseBodyModuleChangeTicketInfoList {
	s.TicketNo = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModuleChangeTicketInfoList) SetTicketStatus(v int32) *TrainOrderQueryV2ResponseBodyModuleChangeTicketInfoList {
	s.TicketStatus = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModuleChangeTicketInfoList) SetToCityName(v string) *TrainOrderQueryV2ResponseBodyModuleChangeTicketInfoList {
	s.ToCityName = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModuleChangeTicketInfoList) SetToStationName(v string) *TrainOrderQueryV2ResponseBodyModuleChangeTicketInfoList {
	s.ToStationName = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModuleChangeTicketInfoList) SetUseTicket(v string) *TrainOrderQueryV2ResponseBodyModuleChangeTicketInfoList {
	s.UseTicket = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModuleChangeTicketInfoList) SetUserId(v string) *TrainOrderQueryV2ResponseBodyModuleChangeTicketInfoList {
	s.UserId = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModuleChangeTicketInfoList) Validate() error {
	return dara.Validate(s)
}

type TrainOrderQueryV2ResponseBodyModuleInvoiceInfo struct {
	// example:
	//
	// 3815504
	Id    *int64  `json:"id,omitempty" xml:"id,omitempty"`
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s TrainOrderQueryV2ResponseBodyModuleInvoiceInfo) String() string {
	return dara.Prettify(s)
}

func (s TrainOrderQueryV2ResponseBodyModuleInvoiceInfo) GoString() string {
	return s.String()
}

func (s *TrainOrderQueryV2ResponseBodyModuleInvoiceInfo) GetId() *int64 {
	return s.Id
}

func (s *TrainOrderQueryV2ResponseBodyModuleInvoiceInfo) GetTitle() *string {
	return s.Title
}

func (s *TrainOrderQueryV2ResponseBodyModuleInvoiceInfo) SetId(v int64) *TrainOrderQueryV2ResponseBodyModuleInvoiceInfo {
	s.Id = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModuleInvoiceInfo) SetTitle(v string) *TrainOrderQueryV2ResponseBodyModuleInvoiceInfo {
	s.Title = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModuleInvoiceInfo) Validate() error {
	return dara.Validate(s)
}

type TrainOrderQueryV2ResponseBodyModuleOrderBaseInfo struct {
	// example:
	//
	// 1003784135
	ApplyId     *string `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
	BtripTitle  *string `json:"btrip_title,omitempty" xml:"btrip_title,omitempty"`
	ContactName *string `json:"contact_name,omitempty" xml:"contact_name,omitempty"`
	// example:
	//
	// 12110002222
	ContactPhone *string `json:"contact_phone,omitempty" xml:"contact_phone,omitempty"`
	// example:
	//
	// btripsy4yd7v0gdpdntpp
	CorpId     *string `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
	CorpName   *string `json:"corp_name,omitempty" xml:"corp_name,omitempty"`
	DepartId   *string `json:"depart_id,omitempty" xml:"depart_id,omitempty"`
	DepartName *string `json:"depart_name,omitempty" xml:"depart_name,omitempty"`
	// example:
	//
	// 1422113021033961000
	ExceedApplyId *string `json:"exceed_apply_id,omitempty" xml:"exceed_apply_id,omitempty"`
	// example:
	//
	// 1422113021033961000
	ExceedThirdpartApplyId *string `json:"exceed_thirdpart_apply_id,omitempty" xml:"exceed_thirdpart_apply_id,omitempty"`
	GmtCreate              *string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	GmtModify              *string `json:"gmt_modify,omitempty" xml:"gmt_modify,omitempty"`
	// example:
	//
	// 42942924
	ItineraryId *string `json:"itinerary_id,omitempty" xml:"itinerary_id,omitempty"`
	// example:
	//
	// 2849819724653209258
	OrderId *int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// example:
	//
	// 4
	OrderStatus         *int32  `json:"order_status,omitempty" xml:"order_status,omitempty"`
	ThirdPartBusinessId *string `json:"thirdPart_business_id,omitempty" xml:"thirdPart_business_id,omitempty"`
	// example:
	//
	// 01-2023-01214
	ThirdpartApplyId *string `json:"thirdpart_apply_id,omitempty" xml:"thirdpart_apply_id,omitempty"`
	// example:
	//
	// DHDI2209141OEGHWRN
	ThirdpartDepartId *string `json:"thirdpart_depart_id,omitempty" xml:"thirdpart_depart_id,omitempty"`
	// example:
	//
	// CS-FLIGHT
	ThirdpartItineraryId *string `json:"thirdpart_itinerary_id,omitempty" xml:"thirdpart_itinerary_id,omitempty"`
	// example:
	//
	// 0
	TripType *int32 `json:"trip_type,omitempty" xml:"trip_type,omitempty"`
	// example:
	//
	// 1231212
	UserId   *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	UserNick *string `json:"user_nick,omitempty" xml:"user_nick,omitempty"`
}

func (s TrainOrderQueryV2ResponseBodyModuleOrderBaseInfo) String() string {
	return dara.Prettify(s)
}

func (s TrainOrderQueryV2ResponseBodyModuleOrderBaseInfo) GoString() string {
	return s.String()
}

func (s *TrainOrderQueryV2ResponseBodyModuleOrderBaseInfo) GetApplyId() *string {
	return s.ApplyId
}

func (s *TrainOrderQueryV2ResponseBodyModuleOrderBaseInfo) GetBtripTitle() *string {
	return s.BtripTitle
}

func (s *TrainOrderQueryV2ResponseBodyModuleOrderBaseInfo) GetContactName() *string {
	return s.ContactName
}

func (s *TrainOrderQueryV2ResponseBodyModuleOrderBaseInfo) GetContactPhone() *string {
	return s.ContactPhone
}

func (s *TrainOrderQueryV2ResponseBodyModuleOrderBaseInfo) GetCorpId() *string {
	return s.CorpId
}

func (s *TrainOrderQueryV2ResponseBodyModuleOrderBaseInfo) GetCorpName() *string {
	return s.CorpName
}

func (s *TrainOrderQueryV2ResponseBodyModuleOrderBaseInfo) GetDepartId() *string {
	return s.DepartId
}

func (s *TrainOrderQueryV2ResponseBodyModuleOrderBaseInfo) GetDepartName() *string {
	return s.DepartName
}

func (s *TrainOrderQueryV2ResponseBodyModuleOrderBaseInfo) GetExceedApplyId() *string {
	return s.ExceedApplyId
}

func (s *TrainOrderQueryV2ResponseBodyModuleOrderBaseInfo) GetExceedThirdpartApplyId() *string {
	return s.ExceedThirdpartApplyId
}

func (s *TrainOrderQueryV2ResponseBodyModuleOrderBaseInfo) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *TrainOrderQueryV2ResponseBodyModuleOrderBaseInfo) GetGmtModify() *string {
	return s.GmtModify
}

func (s *TrainOrderQueryV2ResponseBodyModuleOrderBaseInfo) GetItineraryId() *string {
	return s.ItineraryId
}

func (s *TrainOrderQueryV2ResponseBodyModuleOrderBaseInfo) GetOrderId() *int64 {
	return s.OrderId
}

func (s *TrainOrderQueryV2ResponseBodyModuleOrderBaseInfo) GetOrderStatus() *int32 {
	return s.OrderStatus
}

func (s *TrainOrderQueryV2ResponseBodyModuleOrderBaseInfo) GetThirdPartBusinessId() *string {
	return s.ThirdPartBusinessId
}

func (s *TrainOrderQueryV2ResponseBodyModuleOrderBaseInfo) GetThirdpartApplyId() *string {
	return s.ThirdpartApplyId
}

func (s *TrainOrderQueryV2ResponseBodyModuleOrderBaseInfo) GetThirdpartDepartId() *string {
	return s.ThirdpartDepartId
}

func (s *TrainOrderQueryV2ResponseBodyModuleOrderBaseInfo) GetThirdpartItineraryId() *string {
	return s.ThirdpartItineraryId
}

func (s *TrainOrderQueryV2ResponseBodyModuleOrderBaseInfo) GetTripType() *int32 {
	return s.TripType
}

func (s *TrainOrderQueryV2ResponseBodyModuleOrderBaseInfo) GetUserId() *string {
	return s.UserId
}

func (s *TrainOrderQueryV2ResponseBodyModuleOrderBaseInfo) GetUserNick() *string {
	return s.UserNick
}

func (s *TrainOrderQueryV2ResponseBodyModuleOrderBaseInfo) SetApplyId(v string) *TrainOrderQueryV2ResponseBodyModuleOrderBaseInfo {
	s.ApplyId = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModuleOrderBaseInfo) SetBtripTitle(v string) *TrainOrderQueryV2ResponseBodyModuleOrderBaseInfo {
	s.BtripTitle = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModuleOrderBaseInfo) SetContactName(v string) *TrainOrderQueryV2ResponseBodyModuleOrderBaseInfo {
	s.ContactName = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModuleOrderBaseInfo) SetContactPhone(v string) *TrainOrderQueryV2ResponseBodyModuleOrderBaseInfo {
	s.ContactPhone = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModuleOrderBaseInfo) SetCorpId(v string) *TrainOrderQueryV2ResponseBodyModuleOrderBaseInfo {
	s.CorpId = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModuleOrderBaseInfo) SetCorpName(v string) *TrainOrderQueryV2ResponseBodyModuleOrderBaseInfo {
	s.CorpName = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModuleOrderBaseInfo) SetDepartId(v string) *TrainOrderQueryV2ResponseBodyModuleOrderBaseInfo {
	s.DepartId = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModuleOrderBaseInfo) SetDepartName(v string) *TrainOrderQueryV2ResponseBodyModuleOrderBaseInfo {
	s.DepartName = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModuleOrderBaseInfo) SetExceedApplyId(v string) *TrainOrderQueryV2ResponseBodyModuleOrderBaseInfo {
	s.ExceedApplyId = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModuleOrderBaseInfo) SetExceedThirdpartApplyId(v string) *TrainOrderQueryV2ResponseBodyModuleOrderBaseInfo {
	s.ExceedThirdpartApplyId = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModuleOrderBaseInfo) SetGmtCreate(v string) *TrainOrderQueryV2ResponseBodyModuleOrderBaseInfo {
	s.GmtCreate = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModuleOrderBaseInfo) SetGmtModify(v string) *TrainOrderQueryV2ResponseBodyModuleOrderBaseInfo {
	s.GmtModify = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModuleOrderBaseInfo) SetItineraryId(v string) *TrainOrderQueryV2ResponseBodyModuleOrderBaseInfo {
	s.ItineraryId = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModuleOrderBaseInfo) SetOrderId(v int64) *TrainOrderQueryV2ResponseBodyModuleOrderBaseInfo {
	s.OrderId = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModuleOrderBaseInfo) SetOrderStatus(v int32) *TrainOrderQueryV2ResponseBodyModuleOrderBaseInfo {
	s.OrderStatus = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModuleOrderBaseInfo) SetThirdPartBusinessId(v string) *TrainOrderQueryV2ResponseBodyModuleOrderBaseInfo {
	s.ThirdPartBusinessId = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModuleOrderBaseInfo) SetThirdpartApplyId(v string) *TrainOrderQueryV2ResponseBodyModuleOrderBaseInfo {
	s.ThirdpartApplyId = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModuleOrderBaseInfo) SetThirdpartDepartId(v string) *TrainOrderQueryV2ResponseBodyModuleOrderBaseInfo {
	s.ThirdpartDepartId = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModuleOrderBaseInfo) SetThirdpartItineraryId(v string) *TrainOrderQueryV2ResponseBodyModuleOrderBaseInfo {
	s.ThirdpartItineraryId = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModuleOrderBaseInfo) SetTripType(v int32) *TrainOrderQueryV2ResponseBodyModuleOrderBaseInfo {
	s.TripType = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModuleOrderBaseInfo) SetUserId(v string) *TrainOrderQueryV2ResponseBodyModuleOrderBaseInfo {
	s.UserId = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModuleOrderBaseInfo) SetUserNick(v string) *TrainOrderQueryV2ResponseBodyModuleOrderBaseInfo {
	s.UserNick = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModuleOrderBaseInfo) Validate() error {
	return dara.Validate(s)
}

type TrainOrderQueryV2ResponseBodyModulePassengerInfoList struct {
	// example:
	//
	// 11564
	CostCenterId   *int64  `json:"cost_center_id,omitempty" xml:"cost_center_id,omitempty"`
	CostCenterName *string `json:"cost_center_name,omitempty" xml:"cost_center_name,omitempty"`
	// example:
	//
	// CS-3345
	CostCenterNumber *string `json:"cost_center_number,omitempty" xml:"cost_center_number,omitempty"`
	// example:
	//
	// CS-PROJECT
	ProjectCode *string `json:"project_code,omitempty" xml:"project_code,omitempty"`
	// example:
	//
	// 133576
	ProjectId    *int64  `json:"project_id,omitempty" xml:"project_id,omitempty"`
	ProjectTitle *string `json:"project_title,omitempty" xml:"project_title,omitempty"`
	// example:
	//
	// CS-KDISL
	ThirdpartCostCenterId *string `json:"thirdpart_cost_center_id,omitempty" xml:"thirdpart_cost_center_id,omitempty"`
	// example:
	//
	// CS-22562
	ThirdpartProjectId *string `json:"thirdpart_project_id,omitempty" xml:"thirdpart_project_id,omitempty"`
	// example:
	//
	// 1231231
	UserId   *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	UserName *string `json:"user_name,omitempty" xml:"user_name,omitempty"`
	// example:
	//
	// 0
	UserType *int32 `json:"user_type,omitempty" xml:"user_type,omitempty"`
}

func (s TrainOrderQueryV2ResponseBodyModulePassengerInfoList) String() string {
	return dara.Prettify(s)
}

func (s TrainOrderQueryV2ResponseBodyModulePassengerInfoList) GoString() string {
	return s.String()
}

func (s *TrainOrderQueryV2ResponseBodyModulePassengerInfoList) GetCostCenterId() *int64 {
	return s.CostCenterId
}

func (s *TrainOrderQueryV2ResponseBodyModulePassengerInfoList) GetCostCenterName() *string {
	return s.CostCenterName
}

func (s *TrainOrderQueryV2ResponseBodyModulePassengerInfoList) GetCostCenterNumber() *string {
	return s.CostCenterNumber
}

func (s *TrainOrderQueryV2ResponseBodyModulePassengerInfoList) GetProjectCode() *string {
	return s.ProjectCode
}

func (s *TrainOrderQueryV2ResponseBodyModulePassengerInfoList) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *TrainOrderQueryV2ResponseBodyModulePassengerInfoList) GetProjectTitle() *string {
	return s.ProjectTitle
}

func (s *TrainOrderQueryV2ResponseBodyModulePassengerInfoList) GetThirdpartCostCenterId() *string {
	return s.ThirdpartCostCenterId
}

func (s *TrainOrderQueryV2ResponseBodyModulePassengerInfoList) GetThirdpartProjectId() *string {
	return s.ThirdpartProjectId
}

func (s *TrainOrderQueryV2ResponseBodyModulePassengerInfoList) GetUserId() *string {
	return s.UserId
}

func (s *TrainOrderQueryV2ResponseBodyModulePassengerInfoList) GetUserName() *string {
	return s.UserName
}

func (s *TrainOrderQueryV2ResponseBodyModulePassengerInfoList) GetUserType() *int32 {
	return s.UserType
}

func (s *TrainOrderQueryV2ResponseBodyModulePassengerInfoList) SetCostCenterId(v int64) *TrainOrderQueryV2ResponseBodyModulePassengerInfoList {
	s.CostCenterId = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModulePassengerInfoList) SetCostCenterName(v string) *TrainOrderQueryV2ResponseBodyModulePassengerInfoList {
	s.CostCenterName = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModulePassengerInfoList) SetCostCenterNumber(v string) *TrainOrderQueryV2ResponseBodyModulePassengerInfoList {
	s.CostCenterNumber = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModulePassengerInfoList) SetProjectCode(v string) *TrainOrderQueryV2ResponseBodyModulePassengerInfoList {
	s.ProjectCode = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModulePassengerInfoList) SetProjectId(v int64) *TrainOrderQueryV2ResponseBodyModulePassengerInfoList {
	s.ProjectId = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModulePassengerInfoList) SetProjectTitle(v string) *TrainOrderQueryV2ResponseBodyModulePassengerInfoList {
	s.ProjectTitle = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModulePassengerInfoList) SetThirdpartCostCenterId(v string) *TrainOrderQueryV2ResponseBodyModulePassengerInfoList {
	s.ThirdpartCostCenterId = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModulePassengerInfoList) SetThirdpartProjectId(v string) *TrainOrderQueryV2ResponseBodyModulePassengerInfoList {
	s.ThirdpartProjectId = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModulePassengerInfoList) SetUserId(v string) *TrainOrderQueryV2ResponseBodyModulePassengerInfoList {
	s.UserId = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModulePassengerInfoList) SetUserName(v string) *TrainOrderQueryV2ResponseBodyModulePassengerInfoList {
	s.UserName = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModulePassengerInfoList) SetUserType(v int32) *TrainOrderQueryV2ResponseBodyModulePassengerInfoList {
	s.UserType = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModulePassengerInfoList) Validate() error {
	return dara.Validate(s)
}

type TrainOrderQueryV2ResponseBodyModulePriceInfoList struct {
	CategoryCode *int32 `json:"category_code,omitempty" xml:"category_code,omitempty"`
	// example:
	//
	// 1669344020
	GmtCreate     *string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	PassengerName *string `json:"passenger_name,omitempty" xml:"passenger_name,omitempty"`
	// example:
	//
	// 4
	PayType *int32 `json:"pay_type,omitempty" xml:"pay_type,omitempty"`
	// example:
	//
	// 176000
	Price *float64 `json:"price,omitempty" xml:"price,omitempty"`
	// example:
	//
	// 12345678910987654321
	TradeId *string `json:"trade_id,omitempty" xml:"trade_id,omitempty"`
	Type    *int32  `json:"type,omitempty" xml:"type,omitempty"`
}

func (s TrainOrderQueryV2ResponseBodyModulePriceInfoList) String() string {
	return dara.Prettify(s)
}

func (s TrainOrderQueryV2ResponseBodyModulePriceInfoList) GoString() string {
	return s.String()
}

func (s *TrainOrderQueryV2ResponseBodyModulePriceInfoList) GetCategoryCode() *int32 {
	return s.CategoryCode
}

func (s *TrainOrderQueryV2ResponseBodyModulePriceInfoList) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *TrainOrderQueryV2ResponseBodyModulePriceInfoList) GetPassengerName() *string {
	return s.PassengerName
}

func (s *TrainOrderQueryV2ResponseBodyModulePriceInfoList) GetPayType() *int32 {
	return s.PayType
}

func (s *TrainOrderQueryV2ResponseBodyModulePriceInfoList) GetPrice() *float64 {
	return s.Price
}

func (s *TrainOrderQueryV2ResponseBodyModulePriceInfoList) GetTradeId() *string {
	return s.TradeId
}

func (s *TrainOrderQueryV2ResponseBodyModulePriceInfoList) GetType() *int32 {
	return s.Type
}

func (s *TrainOrderQueryV2ResponseBodyModulePriceInfoList) SetCategoryCode(v int32) *TrainOrderQueryV2ResponseBodyModulePriceInfoList {
	s.CategoryCode = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModulePriceInfoList) SetGmtCreate(v string) *TrainOrderQueryV2ResponseBodyModulePriceInfoList {
	s.GmtCreate = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModulePriceInfoList) SetPassengerName(v string) *TrainOrderQueryV2ResponseBodyModulePriceInfoList {
	s.PassengerName = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModulePriceInfoList) SetPayType(v int32) *TrainOrderQueryV2ResponseBodyModulePriceInfoList {
	s.PayType = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModulePriceInfoList) SetPrice(v float64) *TrainOrderQueryV2ResponseBodyModulePriceInfoList {
	s.Price = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModulePriceInfoList) SetTradeId(v string) *TrainOrderQueryV2ResponseBodyModulePriceInfoList {
	s.TradeId = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModulePriceInfoList) SetType(v int32) *TrainOrderQueryV2ResponseBodyModulePriceInfoList {
	s.Type = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModulePriceInfoList) Validate() error {
	return dara.Validate(s)
}

type TrainOrderQueryV2ResponseBodyModuleRefundTicketInfoList struct {
	GmtCreate     *string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	GmtModify     *string `json:"gmt_modify,omitempty" xml:"gmt_modify,omitempty"`
	RefundApplyId *string `json:"refund_apply_id,omitempty" xml:"refund_apply_id,omitempty"`
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
	// example:
	//
	// 1231231
	UserId *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

func (s TrainOrderQueryV2ResponseBodyModuleRefundTicketInfoList) String() string {
	return dara.Prettify(s)
}

func (s TrainOrderQueryV2ResponseBodyModuleRefundTicketInfoList) GoString() string {
	return s.String()
}

func (s *TrainOrderQueryV2ResponseBodyModuleRefundTicketInfoList) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *TrainOrderQueryV2ResponseBodyModuleRefundTicketInfoList) GetGmtModify() *string {
	return s.GmtModify
}

func (s *TrainOrderQueryV2ResponseBodyModuleRefundTicketInfoList) GetRefundApplyId() *string {
	return s.RefundApplyId
}

func (s *TrainOrderQueryV2ResponseBodyModuleRefundTicketInfoList) GetRefundFee() *float64 {
	return s.RefundFee
}

func (s *TrainOrderQueryV2ResponseBodyModuleRefundTicketInfoList) GetRefundServiceFee() *float64 {
	return s.RefundServiceFee
}

func (s *TrainOrderQueryV2ResponseBodyModuleRefundTicketInfoList) GetTicketNo() *string {
	return s.TicketNo
}

func (s *TrainOrderQueryV2ResponseBodyModuleRefundTicketInfoList) GetUserId() *string {
	return s.UserId
}

func (s *TrainOrderQueryV2ResponseBodyModuleRefundTicketInfoList) SetGmtCreate(v string) *TrainOrderQueryV2ResponseBodyModuleRefundTicketInfoList {
	s.GmtCreate = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModuleRefundTicketInfoList) SetGmtModify(v string) *TrainOrderQueryV2ResponseBodyModuleRefundTicketInfoList {
	s.GmtModify = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModuleRefundTicketInfoList) SetRefundApplyId(v string) *TrainOrderQueryV2ResponseBodyModuleRefundTicketInfoList {
	s.RefundApplyId = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModuleRefundTicketInfoList) SetRefundFee(v float64) *TrainOrderQueryV2ResponseBodyModuleRefundTicketInfoList {
	s.RefundFee = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModuleRefundTicketInfoList) SetRefundServiceFee(v float64) *TrainOrderQueryV2ResponseBodyModuleRefundTicketInfoList {
	s.RefundServiceFee = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModuleRefundTicketInfoList) SetTicketNo(v string) *TrainOrderQueryV2ResponseBodyModuleRefundTicketInfoList {
	s.TicketNo = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModuleRefundTicketInfoList) SetUserId(v string) *TrainOrderQueryV2ResponseBodyModuleRefundTicketInfoList {
	s.UserId = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModuleRefundTicketInfoList) Validate() error {
	return dara.Validate(s)
}

type TrainOrderQueryV2ResponseBodyModuleTrainOrderInfo struct {
	TrainInfoList     []*TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainInfoList   `json:"train_info_list,omitempty" xml:"train_info_list,omitempty" type:"Repeated"`
	TrainTransferInfo *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainTransferInfo `json:"train_transfer_info,omitempty" xml:"train_transfer_info,omitempty" type:"Struct"`
}

func (s TrainOrderQueryV2ResponseBodyModuleTrainOrderInfo) String() string {
	return dara.Prettify(s)
}

func (s TrainOrderQueryV2ResponseBodyModuleTrainOrderInfo) GoString() string {
	return s.String()
}

func (s *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfo) GetTrainInfoList() []*TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainInfoList {
	return s.TrainInfoList
}

func (s *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfo) GetTrainTransferInfo() *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainTransferInfo {
	return s.TrainTransferInfo
}

func (s *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfo) SetTrainInfoList(v []*TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainInfoList) *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfo {
	s.TrainInfoList = v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfo) SetTrainTransferInfo(v *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainTransferInfo) *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfo {
	s.TrainTransferInfo = v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfo) Validate() error {
	return dara.Validate(s)
}

type TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainInfoList struct {
	ArrTime         *string `json:"arr_time,omitempty" xml:"arr_time,omitempty"`
	DepTime         *string `json:"dep_time,omitempty" xml:"dep_time,omitempty"`
	FromCityAdCode  *string `json:"from_city_ad_code,omitempty" xml:"from_city_ad_code,omitempty"`
	FromCityName    *string `json:"from_city_name,omitempty" xml:"from_city_name,omitempty"`
	FromStationName *string `json:"from_station_name,omitempty" xml:"from_station_name,omitempty"`
	// itemId
	//
	// example:
	//
	// 12312
	ItemId *int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// example:
	//
	// 120
	RunTime           *int64  `json:"run_time,omitempty" xml:"run_time,omitempty"`
	SubFromCityAdCode *string `json:"sub_from_city_ad_code,omitempty" xml:"sub_from_city_ad_code,omitempty"`
	SubFromCityAdName *string `json:"sub_from_city_ad_name,omitempty" xml:"sub_from_city_ad_name,omitempty"`
	SubToCityCode     *string `json:"sub_to_city_code,omitempty" xml:"sub_to_city_code,omitempty"`
	SubToCityName     *string `json:"sub_to_city_name,omitempty" xml:"sub_to_city_name,omitempty"`
	ToCityAdCode      *string `json:"to_city_ad_code,omitempty" xml:"to_city_ad_code,omitempty"`
	ToCityName        *string `json:"to_city_name,omitempty" xml:"to_city_name,omitempty"`
	ToStationName     *string `json:"to_station_name,omitempty" xml:"to_station_name,omitempty"`
	// example:
	//
	// D11
	TrainNo          *string                                                                           `json:"train_no,omitempty" xml:"train_no,omitempty"`
	TrainTicketInfos []*TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainInfoListTrainTicketInfos `json:"train_ticket_infos,omitempty" xml:"train_ticket_infos,omitempty" type:"Repeated"`
}

func (s TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainInfoList) String() string {
	return dara.Prettify(s)
}

func (s TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainInfoList) GoString() string {
	return s.String()
}

func (s *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainInfoList) GetArrTime() *string {
	return s.ArrTime
}

func (s *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainInfoList) GetDepTime() *string {
	return s.DepTime
}

func (s *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainInfoList) GetFromCityAdCode() *string {
	return s.FromCityAdCode
}

func (s *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainInfoList) GetFromCityName() *string {
	return s.FromCityName
}

func (s *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainInfoList) GetFromStationName() *string {
	return s.FromStationName
}

func (s *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainInfoList) GetItemId() *int64 {
	return s.ItemId
}

func (s *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainInfoList) GetRunTime() *int64 {
	return s.RunTime
}

func (s *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainInfoList) GetSubFromCityAdCode() *string {
	return s.SubFromCityAdCode
}

func (s *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainInfoList) GetSubFromCityAdName() *string {
	return s.SubFromCityAdName
}

func (s *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainInfoList) GetSubToCityCode() *string {
	return s.SubToCityCode
}

func (s *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainInfoList) GetSubToCityName() *string {
	return s.SubToCityName
}

func (s *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainInfoList) GetToCityAdCode() *string {
	return s.ToCityAdCode
}

func (s *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainInfoList) GetToCityName() *string {
	return s.ToCityName
}

func (s *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainInfoList) GetToStationName() *string {
	return s.ToStationName
}

func (s *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainInfoList) GetTrainNo() *string {
	return s.TrainNo
}

func (s *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainInfoList) GetTrainTicketInfos() []*TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainInfoListTrainTicketInfos {
	return s.TrainTicketInfos
}

func (s *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainInfoList) SetArrTime(v string) *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainInfoList {
	s.ArrTime = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainInfoList) SetDepTime(v string) *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainInfoList {
	s.DepTime = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainInfoList) SetFromCityAdCode(v string) *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainInfoList {
	s.FromCityAdCode = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainInfoList) SetFromCityName(v string) *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainInfoList {
	s.FromCityName = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainInfoList) SetFromStationName(v string) *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainInfoList {
	s.FromStationName = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainInfoList) SetItemId(v int64) *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainInfoList {
	s.ItemId = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainInfoList) SetRunTime(v int64) *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainInfoList {
	s.RunTime = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainInfoList) SetSubFromCityAdCode(v string) *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainInfoList {
	s.SubFromCityAdCode = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainInfoList) SetSubFromCityAdName(v string) *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainInfoList {
	s.SubFromCityAdName = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainInfoList) SetSubToCityCode(v string) *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainInfoList {
	s.SubToCityCode = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainInfoList) SetSubToCityName(v string) *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainInfoList {
	s.SubToCityName = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainInfoList) SetToCityAdCode(v string) *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainInfoList {
	s.ToCityAdCode = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainInfoList) SetToCityName(v string) *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainInfoList {
	s.ToCityName = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainInfoList) SetToStationName(v string) *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainInfoList {
	s.ToStationName = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainInfoList) SetTrainNo(v string) *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainInfoList {
	s.TrainNo = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainInfoList) SetTrainTicketInfos(v []*TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainInfoListTrainTicketInfos) *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainInfoList {
	s.TrainTicketInfos = v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainInfoList) Validate() error {
	return dara.Validate(s)
}

type TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainInfoListTrainTicketInfos struct {
	// example:
	//
	// false
	Changed      *bool   `json:"changed,omitempty" xml:"changed,omitempty"`
	CheckInTime  *string `json:"check_in_time,omitempty" xml:"check_in_time,omitempty"`
	CheckOutTime *string `json:"check_out_time,omitempty" xml:"check_out_time,omitempty"`
	// example:
	//
	// 01
	CoachNo   *string `json:"coach_no,omitempty" xml:"coach_no,omitempty"`
	EndTime   *string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	GmtCreate *string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	GmtModify *string `json:"gmt_modify,omitempty" xml:"gmt_modify,omitempty"`
	// example:
	//
	// 1
	OutTicketStatus *string `json:"out_ticket_status,omitempty" xml:"out_ticket_status,omitempty"`
	// example:
	//
	// 2
	PayType *int32 `json:"pay_type,omitempty" xml:"pay_type,omitempty"`
	// example:
	//
	// 02A
	SeatNo       *string `json:"seat_no,omitempty" xml:"seat_no,omitempty"`
	SeatTypeName *string `json:"seat_type_name,omitempty" xml:"seat_type_name,omitempty"`
	// example:
	//
	// 0
	SegmentIndex *int32 `json:"segment_index,omitempty" xml:"segment_index,omitempty"`
	// example:
	//
	// 100
	ServiceFee     *float64 `json:"service_fee,omitempty" xml:"service_fee,omitempty"`
	ShortTicketNo  *string  `json:"short_ticket_no,omitempty" xml:"short_ticket_no,omitempty"`
	StartTime      *string  `json:"start_time,omitempty" xml:"start_time,omitempty"`
	TicketEntrance *string  `json:"ticket_entrance,omitempty" xml:"ticket_entrance,omitempty"`
	// example:
	//
	// ew123121
	TicketNo *string `json:"ticket_no,omitempty" xml:"ticket_no,omitempty"`
	// example:
	//
	// 400
	TicketPrice *float64 `json:"ticket_price,omitempty" xml:"ticket_price,omitempty"`
	// example:
	//
	// 1
	TicketStatus  *int32  `json:"ticket_status,omitempty" xml:"ticket_status,omitempty"`
	TrainTypeName *string `json:"train_type_name,omitempty" xml:"train_type_name,omitempty"`
	// example:
	//
	// 0
	UseTicket *string `json:"use_ticket,omitempty" xml:"use_ticket,omitempty"`
	// example:
	//
	// 231212
	UserId *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

func (s TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainInfoListTrainTicketInfos) String() string {
	return dara.Prettify(s)
}

func (s TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainInfoListTrainTicketInfos) GoString() string {
	return s.String()
}

func (s *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainInfoListTrainTicketInfos) GetChanged() *bool {
	return s.Changed
}

func (s *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainInfoListTrainTicketInfos) GetCheckInTime() *string {
	return s.CheckInTime
}

func (s *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainInfoListTrainTicketInfos) GetCheckOutTime() *string {
	return s.CheckOutTime
}

func (s *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainInfoListTrainTicketInfos) GetCoachNo() *string {
	return s.CoachNo
}

func (s *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainInfoListTrainTicketInfos) GetEndTime() *string {
	return s.EndTime
}

func (s *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainInfoListTrainTicketInfos) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainInfoListTrainTicketInfos) GetGmtModify() *string {
	return s.GmtModify
}

func (s *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainInfoListTrainTicketInfos) GetOutTicketStatus() *string {
	return s.OutTicketStatus
}

func (s *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainInfoListTrainTicketInfos) GetPayType() *int32 {
	return s.PayType
}

func (s *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainInfoListTrainTicketInfos) GetSeatNo() *string {
	return s.SeatNo
}

func (s *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainInfoListTrainTicketInfos) GetSeatTypeName() *string {
	return s.SeatTypeName
}

func (s *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainInfoListTrainTicketInfos) GetSegmentIndex() *int32 {
	return s.SegmentIndex
}

func (s *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainInfoListTrainTicketInfos) GetServiceFee() *float64 {
	return s.ServiceFee
}

func (s *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainInfoListTrainTicketInfos) GetShortTicketNo() *string {
	return s.ShortTicketNo
}

func (s *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainInfoListTrainTicketInfos) GetStartTime() *string {
	return s.StartTime
}

func (s *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainInfoListTrainTicketInfos) GetTicketEntrance() *string {
	return s.TicketEntrance
}

func (s *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainInfoListTrainTicketInfos) GetTicketNo() *string {
	return s.TicketNo
}

func (s *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainInfoListTrainTicketInfos) GetTicketPrice() *float64 {
	return s.TicketPrice
}

func (s *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainInfoListTrainTicketInfos) GetTicketStatus() *int32 {
	return s.TicketStatus
}

func (s *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainInfoListTrainTicketInfos) GetTrainTypeName() *string {
	return s.TrainTypeName
}

func (s *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainInfoListTrainTicketInfos) GetUseTicket() *string {
	return s.UseTicket
}

func (s *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainInfoListTrainTicketInfos) GetUserId() *string {
	return s.UserId
}

func (s *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainInfoListTrainTicketInfos) SetChanged(v bool) *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainInfoListTrainTicketInfos {
	s.Changed = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainInfoListTrainTicketInfos) SetCheckInTime(v string) *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainInfoListTrainTicketInfos {
	s.CheckInTime = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainInfoListTrainTicketInfos) SetCheckOutTime(v string) *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainInfoListTrainTicketInfos {
	s.CheckOutTime = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainInfoListTrainTicketInfos) SetCoachNo(v string) *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainInfoListTrainTicketInfos {
	s.CoachNo = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainInfoListTrainTicketInfos) SetEndTime(v string) *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainInfoListTrainTicketInfos {
	s.EndTime = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainInfoListTrainTicketInfos) SetGmtCreate(v string) *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainInfoListTrainTicketInfos {
	s.GmtCreate = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainInfoListTrainTicketInfos) SetGmtModify(v string) *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainInfoListTrainTicketInfos {
	s.GmtModify = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainInfoListTrainTicketInfos) SetOutTicketStatus(v string) *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainInfoListTrainTicketInfos {
	s.OutTicketStatus = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainInfoListTrainTicketInfos) SetPayType(v int32) *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainInfoListTrainTicketInfos {
	s.PayType = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainInfoListTrainTicketInfos) SetSeatNo(v string) *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainInfoListTrainTicketInfos {
	s.SeatNo = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainInfoListTrainTicketInfos) SetSeatTypeName(v string) *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainInfoListTrainTicketInfos {
	s.SeatTypeName = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainInfoListTrainTicketInfos) SetSegmentIndex(v int32) *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainInfoListTrainTicketInfos {
	s.SegmentIndex = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainInfoListTrainTicketInfos) SetServiceFee(v float64) *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainInfoListTrainTicketInfos {
	s.ServiceFee = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainInfoListTrainTicketInfos) SetShortTicketNo(v string) *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainInfoListTrainTicketInfos {
	s.ShortTicketNo = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainInfoListTrainTicketInfos) SetStartTime(v string) *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainInfoListTrainTicketInfos {
	s.StartTime = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainInfoListTrainTicketInfos) SetTicketEntrance(v string) *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainInfoListTrainTicketInfos {
	s.TicketEntrance = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainInfoListTrainTicketInfos) SetTicketNo(v string) *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainInfoListTrainTicketInfos {
	s.TicketNo = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainInfoListTrainTicketInfos) SetTicketPrice(v float64) *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainInfoListTrainTicketInfos {
	s.TicketPrice = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainInfoListTrainTicketInfos) SetTicketStatus(v int32) *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainInfoListTrainTicketInfos {
	s.TicketStatus = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainInfoListTrainTicketInfos) SetTrainTypeName(v string) *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainInfoListTrainTicketInfos {
	s.TrainTypeName = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainInfoListTrainTicketInfos) SetUseTicket(v string) *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainInfoListTrainTicketInfos {
	s.UseTicket = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainInfoListTrainTicketInfos) SetUserId(v string) *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainInfoListTrainTicketInfos {
	s.UserId = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainInfoListTrainTicketInfos) Validate() error {
	return dara.Validate(s)
}

type TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainTransferInfo struct {
	// example:
	//
	// 200
	CostTime *string `json:"cost_time,omitempty" xml:"cost_time,omitempty"`
	// example:
	//
	// 2022-11-15 00:00:00
	EndTime          *string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	FromCityAdCode   *string `json:"from_city_ad_code,omitempty" xml:"from_city_ad_code,omitempty"`
	FromCityName     *string `json:"from_city_name,omitempty" xml:"from_city_name,omitempty"`
	FromStationName  *string `json:"from_station_name,omitempty" xml:"from_station_name,omitempty"`
	MiddleCity       *string `json:"middle_city,omitempty" xml:"middle_city,omitempty"`
	MiddleCityAdCode *string `json:"middle_city_ad_code,omitempty" xml:"middle_city_ad_code,omitempty"`
	// example:
	//
	// 2023-01-29 18:10:00
	MiddleDate    *string `json:"middle_date,omitempty" xml:"middle_date,omitempty"`
	MiddleStation *string `json:"middle_station,omitempty" xml:"middle_station,omitempty"`
	MiddleType    *string `json:"middle_type,omitempty" xml:"middle_type,omitempty"`
	// example:
	//
	// 2022-11-01 00:00:00
	StartTime         *string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	SubFromCityAdCode *string `json:"sub_from_city_ad_code,omitempty" xml:"sub_from_city_ad_code,omitempty"`
	SubFromCityAdName *string `json:"sub_from_city_ad_name,omitempty" xml:"sub_from_city_ad_name,omitempty"`
	SubMiddleCityCode *string `json:"sub_middle_city_code,omitempty" xml:"sub_middle_city_code,omitempty"`
	SubMiddleCityName *string `json:"sub_middle_city_name,omitempty" xml:"sub_middle_city_name,omitempty"`
	SubToCityCode     *string `json:"sub_to_city_code,omitempty" xml:"sub_to_city_code,omitempty"`
	SubToCityName     *string `json:"sub_to_city_name,omitempty" xml:"sub_to_city_name,omitempty"`
	ToCityAdCode      *string `json:"to_city_ad_code,omitempty" xml:"to_city_ad_code,omitempty"`
	ToCityName        *string `json:"to_city_name,omitempty" xml:"to_city_name,omitempty"`
	ToStationName     *string `json:"to_station_name,omitempty" xml:"to_station_name,omitempty"`
	// example:
	//
	// 60
	WaitTime *string `json:"wait_time,omitempty" xml:"wait_time,omitempty"`
}

func (s TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainTransferInfo) String() string {
	return dara.Prettify(s)
}

func (s TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainTransferInfo) GoString() string {
	return s.String()
}

func (s *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainTransferInfo) GetCostTime() *string {
	return s.CostTime
}

func (s *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainTransferInfo) GetEndTime() *string {
	return s.EndTime
}

func (s *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainTransferInfo) GetFromCityAdCode() *string {
	return s.FromCityAdCode
}

func (s *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainTransferInfo) GetFromCityName() *string {
	return s.FromCityName
}

func (s *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainTransferInfo) GetFromStationName() *string {
	return s.FromStationName
}

func (s *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainTransferInfo) GetMiddleCity() *string {
	return s.MiddleCity
}

func (s *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainTransferInfo) GetMiddleCityAdCode() *string {
	return s.MiddleCityAdCode
}

func (s *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainTransferInfo) GetMiddleDate() *string {
	return s.MiddleDate
}

func (s *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainTransferInfo) GetMiddleStation() *string {
	return s.MiddleStation
}

func (s *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainTransferInfo) GetMiddleType() *string {
	return s.MiddleType
}

func (s *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainTransferInfo) GetStartTime() *string {
	return s.StartTime
}

func (s *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainTransferInfo) GetSubFromCityAdCode() *string {
	return s.SubFromCityAdCode
}

func (s *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainTransferInfo) GetSubFromCityAdName() *string {
	return s.SubFromCityAdName
}

func (s *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainTransferInfo) GetSubMiddleCityCode() *string {
	return s.SubMiddleCityCode
}

func (s *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainTransferInfo) GetSubMiddleCityName() *string {
	return s.SubMiddleCityName
}

func (s *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainTransferInfo) GetSubToCityCode() *string {
	return s.SubToCityCode
}

func (s *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainTransferInfo) GetSubToCityName() *string {
	return s.SubToCityName
}

func (s *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainTransferInfo) GetToCityAdCode() *string {
	return s.ToCityAdCode
}

func (s *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainTransferInfo) GetToCityName() *string {
	return s.ToCityName
}

func (s *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainTransferInfo) GetToStationName() *string {
	return s.ToStationName
}

func (s *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainTransferInfo) GetWaitTime() *string {
	return s.WaitTime
}

func (s *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainTransferInfo) SetCostTime(v string) *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainTransferInfo {
	s.CostTime = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainTransferInfo) SetEndTime(v string) *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainTransferInfo {
	s.EndTime = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainTransferInfo) SetFromCityAdCode(v string) *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainTransferInfo {
	s.FromCityAdCode = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainTransferInfo) SetFromCityName(v string) *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainTransferInfo {
	s.FromCityName = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainTransferInfo) SetFromStationName(v string) *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainTransferInfo {
	s.FromStationName = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainTransferInfo) SetMiddleCity(v string) *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainTransferInfo {
	s.MiddleCity = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainTransferInfo) SetMiddleCityAdCode(v string) *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainTransferInfo {
	s.MiddleCityAdCode = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainTransferInfo) SetMiddleDate(v string) *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainTransferInfo {
	s.MiddleDate = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainTransferInfo) SetMiddleStation(v string) *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainTransferInfo {
	s.MiddleStation = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainTransferInfo) SetMiddleType(v string) *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainTransferInfo {
	s.MiddleType = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainTransferInfo) SetStartTime(v string) *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainTransferInfo {
	s.StartTime = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainTransferInfo) SetSubFromCityAdCode(v string) *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainTransferInfo {
	s.SubFromCityAdCode = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainTransferInfo) SetSubFromCityAdName(v string) *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainTransferInfo {
	s.SubFromCityAdName = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainTransferInfo) SetSubMiddleCityCode(v string) *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainTransferInfo {
	s.SubMiddleCityCode = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainTransferInfo) SetSubMiddleCityName(v string) *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainTransferInfo {
	s.SubMiddleCityName = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainTransferInfo) SetSubToCityCode(v string) *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainTransferInfo {
	s.SubToCityCode = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainTransferInfo) SetSubToCityName(v string) *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainTransferInfo {
	s.SubToCityName = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainTransferInfo) SetToCityAdCode(v string) *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainTransferInfo {
	s.ToCityAdCode = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainTransferInfo) SetToCityName(v string) *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainTransferInfo {
	s.ToCityName = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainTransferInfo) SetToStationName(v string) *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainTransferInfo {
	s.ToStationName = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainTransferInfo) SetWaitTime(v string) *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainTransferInfo {
	s.WaitTime = &v
	return s
}

func (s *TrainOrderQueryV2ResponseBodyModuleTrainOrderInfoTrainTransferInfo) Validate() error {
	return dara.Validate(s)
}
