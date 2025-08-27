// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTrainOrderDetailQueryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *TrainOrderDetailQueryResponseBody
	GetCode() *string
	SetMessage(v string) *TrainOrderDetailQueryResponseBody
	GetMessage() *string
	SetModule(v *TrainOrderDetailQueryResponseBodyModule) *TrainOrderDetailQueryResponseBody
	GetModule() *TrainOrderDetailQueryResponseBodyModule
	SetRequestId(v string) *TrainOrderDetailQueryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *TrainOrderDetailQueryResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *TrainOrderDetailQueryResponseBody
	GetTraceId() *string
}

type TrainOrderDetailQueryResponseBody struct {
	// example:
	//
	// 200
	Code    *string `json:"code,omitempty" xml:"code,omitempty"`
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// module
	Module *TrainOrderDetailQueryResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	// example:
	//
	// 210bc81a17090871660176894d008c
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// traceId
	//
	// example:
	//
	// 2103a05c16872420814992343d8a09
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s TrainOrderDetailQueryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TrainOrderDetailQueryResponseBody) GoString() string {
	return s.String()
}

func (s *TrainOrderDetailQueryResponseBody) GetCode() *string {
	return s.Code
}

func (s *TrainOrderDetailQueryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *TrainOrderDetailQueryResponseBody) GetModule() *TrainOrderDetailQueryResponseBodyModule {
	return s.Module
}

func (s *TrainOrderDetailQueryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TrainOrderDetailQueryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *TrainOrderDetailQueryResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *TrainOrderDetailQueryResponseBody) SetCode(v string) *TrainOrderDetailQueryResponseBody {
	s.Code = &v
	return s
}

func (s *TrainOrderDetailQueryResponseBody) SetMessage(v string) *TrainOrderDetailQueryResponseBody {
	s.Message = &v
	return s
}

func (s *TrainOrderDetailQueryResponseBody) SetModule(v *TrainOrderDetailQueryResponseBodyModule) *TrainOrderDetailQueryResponseBody {
	s.Module = v
	return s
}

func (s *TrainOrderDetailQueryResponseBody) SetRequestId(v string) *TrainOrderDetailQueryResponseBody {
	s.RequestId = &v
	return s
}

func (s *TrainOrderDetailQueryResponseBody) SetSuccess(v bool) *TrainOrderDetailQueryResponseBody {
	s.Success = &v
	return s
}

func (s *TrainOrderDetailQueryResponseBody) SetTraceId(v string) *TrainOrderDetailQueryResponseBody {
	s.TraceId = &v
	return s
}

func (s *TrainOrderDetailQueryResponseBody) Validate() error {
	return dara.Validate(s)
}

type TrainOrderDetailQueryResponseBodyModule struct {
	BookInfos            *TrainOrderDetailQueryResponseBodyModuleBookInfos              `json:"book_infos,omitempty" xml:"book_infos,omitempty" type:"Struct"`
	ChangeInfos          []*TrainOrderDetailQueryResponseBodyModuleChangeInfos          `json:"change_infos,omitempty" xml:"change_infos,omitempty" type:"Repeated"`
	OfflineRefundDetails []*TrainOrderDetailQueryResponseBodyModuleOfflineRefundDetails `json:"offlineRefundDetails,omitempty" xml:"offlineRefundDetails,omitempty" type:"Repeated"`
	// example:
	//
	// 1017028198411054446
	OrderId *string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// example:
	//
	// 3703184209587306496
	OutOrderId     *string                                                  `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
	PassengerInfoS []*TrainOrderDetailQueryResponseBodyModulePassengerInfoS `json:"passenger_info_s,omitempty" xml:"passenger_info_s,omitempty" type:"Repeated"`
	RefundInfos    []*TrainOrderDetailQueryResponseBodyModuleRefundInfos    `json:"refund_infos,omitempty" xml:"refund_infos,omitempty" type:"Repeated"`
}

func (s TrainOrderDetailQueryResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s TrainOrderDetailQueryResponseBodyModule) GoString() string {
	return s.String()
}

func (s *TrainOrderDetailQueryResponseBodyModule) GetBookInfos() *TrainOrderDetailQueryResponseBodyModuleBookInfos {
	return s.BookInfos
}

func (s *TrainOrderDetailQueryResponseBodyModule) GetChangeInfos() []*TrainOrderDetailQueryResponseBodyModuleChangeInfos {
	return s.ChangeInfos
}

func (s *TrainOrderDetailQueryResponseBodyModule) GetOfflineRefundDetails() []*TrainOrderDetailQueryResponseBodyModuleOfflineRefundDetails {
	return s.OfflineRefundDetails
}

func (s *TrainOrderDetailQueryResponseBodyModule) GetOrderId() *string {
	return s.OrderId
}

func (s *TrainOrderDetailQueryResponseBodyModule) GetOutOrderId() *string {
	return s.OutOrderId
}

func (s *TrainOrderDetailQueryResponseBodyModule) GetPassengerInfoS() []*TrainOrderDetailQueryResponseBodyModulePassengerInfoS {
	return s.PassengerInfoS
}

func (s *TrainOrderDetailQueryResponseBodyModule) GetRefundInfos() []*TrainOrderDetailQueryResponseBodyModuleRefundInfos {
	return s.RefundInfos
}

func (s *TrainOrderDetailQueryResponseBodyModule) SetBookInfos(v *TrainOrderDetailQueryResponseBodyModuleBookInfos) *TrainOrderDetailQueryResponseBodyModule {
	s.BookInfos = v
	return s
}

func (s *TrainOrderDetailQueryResponseBodyModule) SetChangeInfos(v []*TrainOrderDetailQueryResponseBodyModuleChangeInfos) *TrainOrderDetailQueryResponseBodyModule {
	s.ChangeInfos = v
	return s
}

func (s *TrainOrderDetailQueryResponseBodyModule) SetOfflineRefundDetails(v []*TrainOrderDetailQueryResponseBodyModuleOfflineRefundDetails) *TrainOrderDetailQueryResponseBodyModule {
	s.OfflineRefundDetails = v
	return s
}

func (s *TrainOrderDetailQueryResponseBodyModule) SetOrderId(v string) *TrainOrderDetailQueryResponseBodyModule {
	s.OrderId = &v
	return s
}

func (s *TrainOrderDetailQueryResponseBodyModule) SetOutOrderId(v string) *TrainOrderDetailQueryResponseBodyModule {
	s.OutOrderId = &v
	return s
}

func (s *TrainOrderDetailQueryResponseBodyModule) SetPassengerInfoS(v []*TrainOrderDetailQueryResponseBodyModulePassengerInfoS) *TrainOrderDetailQueryResponseBodyModule {
	s.PassengerInfoS = v
	return s
}

func (s *TrainOrderDetailQueryResponseBodyModule) SetRefundInfos(v []*TrainOrderDetailQueryResponseBodyModuleRefundInfos) *TrainOrderDetailQueryResponseBodyModule {
	s.RefundInfos = v
	return s
}

func (s *TrainOrderDetailQueryResponseBodyModule) Validate() error {
	return dara.Validate(s)
}

type TrainOrderDetailQueryResponseBodyModuleBookInfos struct {
	BookTrainInfos []*TrainOrderDetailQueryResponseBodyModuleBookInfosBookTrainInfos `json:"book_train_infos,omitempty" xml:"book_train_infos,omitempty" type:"Repeated"`
	// example:
	//
	// 1111
	FailCode *string `json:"fail_code,omitempty" xml:"fail_code,omitempty"`
	FailMsg  *string `json:"fail_msg,omitempty" xml:"fail_msg,omitempty"`
	// example:
	//
	// 2024-05-06 15:19:01
	LastPayTime *string `json:"last_pay_time,omitempty" xml:"last_pay_time,omitempty"`
	// example:
	//
	// 1
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// GW123456
	TicketNo *string `json:"ticket_no,omitempty" xml:"ticket_no,omitempty"`
}

func (s TrainOrderDetailQueryResponseBodyModuleBookInfos) String() string {
	return dara.Prettify(s)
}

func (s TrainOrderDetailQueryResponseBodyModuleBookInfos) GoString() string {
	return s.String()
}

func (s *TrainOrderDetailQueryResponseBodyModuleBookInfos) GetBookTrainInfos() []*TrainOrderDetailQueryResponseBodyModuleBookInfosBookTrainInfos {
	return s.BookTrainInfos
}

func (s *TrainOrderDetailQueryResponseBodyModuleBookInfos) GetFailCode() *string {
	return s.FailCode
}

func (s *TrainOrderDetailQueryResponseBodyModuleBookInfos) GetFailMsg() *string {
	return s.FailMsg
}

func (s *TrainOrderDetailQueryResponseBodyModuleBookInfos) GetLastPayTime() *string {
	return s.LastPayTime
}

func (s *TrainOrderDetailQueryResponseBodyModuleBookInfos) GetStatus() *int32 {
	return s.Status
}

func (s *TrainOrderDetailQueryResponseBodyModuleBookInfos) GetTicketNo() *string {
	return s.TicketNo
}

func (s *TrainOrderDetailQueryResponseBodyModuleBookInfos) SetBookTrainInfos(v []*TrainOrderDetailQueryResponseBodyModuleBookInfosBookTrainInfos) *TrainOrderDetailQueryResponseBodyModuleBookInfos {
	s.BookTrainInfos = v
	return s
}

func (s *TrainOrderDetailQueryResponseBodyModuleBookInfos) SetFailCode(v string) *TrainOrderDetailQueryResponseBodyModuleBookInfos {
	s.FailCode = &v
	return s
}

func (s *TrainOrderDetailQueryResponseBodyModuleBookInfos) SetFailMsg(v string) *TrainOrderDetailQueryResponseBodyModuleBookInfos {
	s.FailMsg = &v
	return s
}

func (s *TrainOrderDetailQueryResponseBodyModuleBookInfos) SetLastPayTime(v string) *TrainOrderDetailQueryResponseBodyModuleBookInfos {
	s.LastPayTime = &v
	return s
}

func (s *TrainOrderDetailQueryResponseBodyModuleBookInfos) SetStatus(v int32) *TrainOrderDetailQueryResponseBodyModuleBookInfos {
	s.Status = &v
	return s
}

func (s *TrainOrderDetailQueryResponseBodyModuleBookInfos) SetTicketNo(v string) *TrainOrderDetailQueryResponseBodyModuleBookInfos {
	s.TicketNo = &v
	return s
}

func (s *TrainOrderDetailQueryResponseBodyModuleBookInfos) Validate() error {
	return dara.Validate(s)
}

type TrainOrderDetailQueryResponseBodyModuleBookInfosBookTrainInfos struct {
	// example:
	//
	// BTC
	ArrStationCode *string `json:"arr_station_code,omitempty" xml:"arr_station_code,omitempty"`
	ArrStationName *string `json:"arr_station_name,omitempty" xml:"arr_station_name,omitempty"`
	// example:
	//
	// 2024-05-07 15:19:01
	ArriveTime      *string                                                                          `json:"arrive_time,omitempty" xml:"arrive_time,omitempty"`
	BookTicketInfos []*TrainOrderDetailQueryResponseBodyModuleBookInfosBookTrainInfosBookTicketInfos `json:"book_ticket_infos,omitempty" xml:"book_ticket_infos,omitempty" type:"Repeated"`
	// example:
	//
	// BDC
	DepStationCode *string `json:"dep_station_code,omitempty" xml:"dep_station_code,omitempty"`
	DepStationName *string `json:"dep_station_name,omitempty" xml:"dep_station_name,omitempty"`
	// example:
	//
	// 2024-05-06 15:19:01
	DepTime *string `json:"dep_time,omitempty" xml:"dep_time,omitempty"`
	// example:
	//
	// D1234
	TrainNo *string `json:"train_no,omitempty" xml:"train_no,omitempty"`
}

func (s TrainOrderDetailQueryResponseBodyModuleBookInfosBookTrainInfos) String() string {
	return dara.Prettify(s)
}

func (s TrainOrderDetailQueryResponseBodyModuleBookInfosBookTrainInfos) GoString() string {
	return s.String()
}

func (s *TrainOrderDetailQueryResponseBodyModuleBookInfosBookTrainInfos) GetArrStationCode() *string {
	return s.ArrStationCode
}

func (s *TrainOrderDetailQueryResponseBodyModuleBookInfosBookTrainInfos) GetArrStationName() *string {
	return s.ArrStationName
}

func (s *TrainOrderDetailQueryResponseBodyModuleBookInfosBookTrainInfos) GetArriveTime() *string {
	return s.ArriveTime
}

func (s *TrainOrderDetailQueryResponseBodyModuleBookInfosBookTrainInfos) GetBookTicketInfos() []*TrainOrderDetailQueryResponseBodyModuleBookInfosBookTrainInfosBookTicketInfos {
	return s.BookTicketInfos
}

func (s *TrainOrderDetailQueryResponseBodyModuleBookInfosBookTrainInfos) GetDepStationCode() *string {
	return s.DepStationCode
}

func (s *TrainOrderDetailQueryResponseBodyModuleBookInfosBookTrainInfos) GetDepStationName() *string {
	return s.DepStationName
}

func (s *TrainOrderDetailQueryResponseBodyModuleBookInfosBookTrainInfos) GetDepTime() *string {
	return s.DepTime
}

func (s *TrainOrderDetailQueryResponseBodyModuleBookInfosBookTrainInfos) GetTrainNo() *string {
	return s.TrainNo
}

func (s *TrainOrderDetailQueryResponseBodyModuleBookInfosBookTrainInfos) SetArrStationCode(v string) *TrainOrderDetailQueryResponseBodyModuleBookInfosBookTrainInfos {
	s.ArrStationCode = &v
	return s
}

func (s *TrainOrderDetailQueryResponseBodyModuleBookInfosBookTrainInfos) SetArrStationName(v string) *TrainOrderDetailQueryResponseBodyModuleBookInfosBookTrainInfos {
	s.ArrStationName = &v
	return s
}

func (s *TrainOrderDetailQueryResponseBodyModuleBookInfosBookTrainInfos) SetArriveTime(v string) *TrainOrderDetailQueryResponseBodyModuleBookInfosBookTrainInfos {
	s.ArriveTime = &v
	return s
}

func (s *TrainOrderDetailQueryResponseBodyModuleBookInfosBookTrainInfos) SetBookTicketInfos(v []*TrainOrderDetailQueryResponseBodyModuleBookInfosBookTrainInfosBookTicketInfos) *TrainOrderDetailQueryResponseBodyModuleBookInfosBookTrainInfos {
	s.BookTicketInfos = v
	return s
}

func (s *TrainOrderDetailQueryResponseBodyModuleBookInfosBookTrainInfos) SetDepStationCode(v string) *TrainOrderDetailQueryResponseBodyModuleBookInfosBookTrainInfos {
	s.DepStationCode = &v
	return s
}

func (s *TrainOrderDetailQueryResponseBodyModuleBookInfosBookTrainInfos) SetDepStationName(v string) *TrainOrderDetailQueryResponseBodyModuleBookInfosBookTrainInfos {
	s.DepStationName = &v
	return s
}

func (s *TrainOrderDetailQueryResponseBodyModuleBookInfosBookTrainInfos) SetDepTime(v string) *TrainOrderDetailQueryResponseBodyModuleBookInfosBookTrainInfos {
	s.DepTime = &v
	return s
}

func (s *TrainOrderDetailQueryResponseBodyModuleBookInfosBookTrainInfos) SetTrainNo(v string) *TrainOrderDetailQueryResponseBodyModuleBookInfosBookTrainInfos {
	s.TrainNo = &v
	return s
}

func (s *TrainOrderDetailQueryResponseBodyModuleBookInfosBookTrainInfos) Validate() error {
	return dara.Validate(s)
}

type TrainOrderDetailQueryResponseBodyModuleBookInfosBookTrainInfosBookTicketInfos struct {
	// example:
	//
	// 04
	CoachNo *string `json:"coach_no,omitempty" xml:"coach_no,omitempty"`
	// example:
	//
	// 1111
	FailCode   *string `json:"fail_code,omitempty" xml:"fail_code,omitempty"`
	FailReason *string `json:"fail_reason,omitempty" xml:"fail_reason,omitempty"`
	// example:
	//
	// 123456
	PassengerId *string `json:"passenger_id,omitempty" xml:"passenger_id,omitempty"`
	// example:
	//
	// 10000
	RealTicketPrice *int64 `json:"real_ticket_price,omitempty" xml:"real_ticket_price,omitempty"`
	// example:
	//
	// 1A
	SeatNo *string `json:"seat_no,omitempty" xml:"seat_no,omitempty"`
	// example:
	//
	// 14
	SeatType *string `json:"seat_type,omitempty" xml:"seat_type,omitempty"`
	// example:
	//
	// null
	TicketEntrance *string `json:"ticket_entrance,omitempty" xml:"ticket_entrance,omitempty"`
	// example:
	//
	// 100
	TicketPrice *int64 `json:"ticket_price,omitempty" xml:"ticket_price,omitempty"`
	// example:
	//
	// 11
	TicketStatus *int32 `json:"ticket_status,omitempty" xml:"ticket_status,omitempty"`
	// example:
	//
	// 0
	TicketType *string `json:"ticket_type,omitempty" xml:"ticket_type,omitempty"`
}

func (s TrainOrderDetailQueryResponseBodyModuleBookInfosBookTrainInfosBookTicketInfos) String() string {
	return dara.Prettify(s)
}

func (s TrainOrderDetailQueryResponseBodyModuleBookInfosBookTrainInfosBookTicketInfos) GoString() string {
	return s.String()
}

func (s *TrainOrderDetailQueryResponseBodyModuleBookInfosBookTrainInfosBookTicketInfos) GetCoachNo() *string {
	return s.CoachNo
}

func (s *TrainOrderDetailQueryResponseBodyModuleBookInfosBookTrainInfosBookTicketInfos) GetFailCode() *string {
	return s.FailCode
}

func (s *TrainOrderDetailQueryResponseBodyModuleBookInfosBookTrainInfosBookTicketInfos) GetFailReason() *string {
	return s.FailReason
}

func (s *TrainOrderDetailQueryResponseBodyModuleBookInfosBookTrainInfosBookTicketInfos) GetPassengerId() *string {
	return s.PassengerId
}

func (s *TrainOrderDetailQueryResponseBodyModuleBookInfosBookTrainInfosBookTicketInfos) GetRealTicketPrice() *int64 {
	return s.RealTicketPrice
}

func (s *TrainOrderDetailQueryResponseBodyModuleBookInfosBookTrainInfosBookTicketInfos) GetSeatNo() *string {
	return s.SeatNo
}

func (s *TrainOrderDetailQueryResponseBodyModuleBookInfosBookTrainInfosBookTicketInfos) GetSeatType() *string {
	return s.SeatType
}

func (s *TrainOrderDetailQueryResponseBodyModuleBookInfosBookTrainInfosBookTicketInfos) GetTicketEntrance() *string {
	return s.TicketEntrance
}

func (s *TrainOrderDetailQueryResponseBodyModuleBookInfosBookTrainInfosBookTicketInfos) GetTicketPrice() *int64 {
	return s.TicketPrice
}

func (s *TrainOrderDetailQueryResponseBodyModuleBookInfosBookTrainInfosBookTicketInfos) GetTicketStatus() *int32 {
	return s.TicketStatus
}

func (s *TrainOrderDetailQueryResponseBodyModuleBookInfosBookTrainInfosBookTicketInfos) GetTicketType() *string {
	return s.TicketType
}

func (s *TrainOrderDetailQueryResponseBodyModuleBookInfosBookTrainInfosBookTicketInfos) SetCoachNo(v string) *TrainOrderDetailQueryResponseBodyModuleBookInfosBookTrainInfosBookTicketInfos {
	s.CoachNo = &v
	return s
}

func (s *TrainOrderDetailQueryResponseBodyModuleBookInfosBookTrainInfosBookTicketInfos) SetFailCode(v string) *TrainOrderDetailQueryResponseBodyModuleBookInfosBookTrainInfosBookTicketInfos {
	s.FailCode = &v
	return s
}

func (s *TrainOrderDetailQueryResponseBodyModuleBookInfosBookTrainInfosBookTicketInfos) SetFailReason(v string) *TrainOrderDetailQueryResponseBodyModuleBookInfosBookTrainInfosBookTicketInfos {
	s.FailReason = &v
	return s
}

func (s *TrainOrderDetailQueryResponseBodyModuleBookInfosBookTrainInfosBookTicketInfos) SetPassengerId(v string) *TrainOrderDetailQueryResponseBodyModuleBookInfosBookTrainInfosBookTicketInfos {
	s.PassengerId = &v
	return s
}

func (s *TrainOrderDetailQueryResponseBodyModuleBookInfosBookTrainInfosBookTicketInfos) SetRealTicketPrice(v int64) *TrainOrderDetailQueryResponseBodyModuleBookInfosBookTrainInfosBookTicketInfos {
	s.RealTicketPrice = &v
	return s
}

func (s *TrainOrderDetailQueryResponseBodyModuleBookInfosBookTrainInfosBookTicketInfos) SetSeatNo(v string) *TrainOrderDetailQueryResponseBodyModuleBookInfosBookTrainInfosBookTicketInfos {
	s.SeatNo = &v
	return s
}

func (s *TrainOrderDetailQueryResponseBodyModuleBookInfosBookTrainInfosBookTicketInfos) SetSeatType(v string) *TrainOrderDetailQueryResponseBodyModuleBookInfosBookTrainInfosBookTicketInfos {
	s.SeatType = &v
	return s
}

func (s *TrainOrderDetailQueryResponseBodyModuleBookInfosBookTrainInfosBookTicketInfos) SetTicketEntrance(v string) *TrainOrderDetailQueryResponseBodyModuleBookInfosBookTrainInfosBookTicketInfos {
	s.TicketEntrance = &v
	return s
}

func (s *TrainOrderDetailQueryResponseBodyModuleBookInfosBookTrainInfosBookTicketInfos) SetTicketPrice(v int64) *TrainOrderDetailQueryResponseBodyModuleBookInfosBookTrainInfosBookTicketInfos {
	s.TicketPrice = &v
	return s
}

func (s *TrainOrderDetailQueryResponseBodyModuleBookInfosBookTrainInfosBookTicketInfos) SetTicketStatus(v int32) *TrainOrderDetailQueryResponseBodyModuleBookInfosBookTrainInfosBookTicketInfos {
	s.TicketStatus = &v
	return s
}

func (s *TrainOrderDetailQueryResponseBodyModuleBookInfosBookTrainInfosBookTicketInfos) SetTicketType(v string) *TrainOrderDetailQueryResponseBodyModuleBookInfosBookTrainInfosBookTicketInfos {
	s.TicketType = &v
	return s
}

func (s *TrainOrderDetailQueryResponseBodyModuleBookInfosBookTrainInfosBookTicketInfos) Validate() error {
	return dara.Validate(s)
}

type TrainOrderDetailQueryResponseBodyModuleChangeInfos struct {
	// example:
	//
	// 1234567890
	ChangeApplyId    *string                                                               `json:"change_apply_id,omitempty" xml:"change_apply_id,omitempty"`
	ChangeTrainInfos []*TrainOrderDetailQueryResponseBodyModuleChangeInfosChangeTrainInfos `json:"change_train_infos,omitempty" xml:"change_train_infos,omitempty" type:"Repeated"`
	// example:
	//
	// 2024-05-06 15:19:01
	LimitPayTime *string `json:"limit_pay_time,omitempty" xml:"limit_pay_time,omitempty"`
	// example:
	//
	// 12345
	OutChangeApplyId *string `json:"out_change_apply_id,omitempty" xml:"out_change_apply_id,omitempty"`
	// example:
	//
	// 1
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s TrainOrderDetailQueryResponseBodyModuleChangeInfos) String() string {
	return dara.Prettify(s)
}

func (s TrainOrderDetailQueryResponseBodyModuleChangeInfos) GoString() string {
	return s.String()
}

func (s *TrainOrderDetailQueryResponseBodyModuleChangeInfos) GetChangeApplyId() *string {
	return s.ChangeApplyId
}

func (s *TrainOrderDetailQueryResponseBodyModuleChangeInfos) GetChangeTrainInfos() []*TrainOrderDetailQueryResponseBodyModuleChangeInfosChangeTrainInfos {
	return s.ChangeTrainInfos
}

func (s *TrainOrderDetailQueryResponseBodyModuleChangeInfos) GetLimitPayTime() *string {
	return s.LimitPayTime
}

func (s *TrainOrderDetailQueryResponseBodyModuleChangeInfos) GetOutChangeApplyId() *string {
	return s.OutChangeApplyId
}

func (s *TrainOrderDetailQueryResponseBodyModuleChangeInfos) GetStatus() *string {
	return s.Status
}

func (s *TrainOrderDetailQueryResponseBodyModuleChangeInfos) SetChangeApplyId(v string) *TrainOrderDetailQueryResponseBodyModuleChangeInfos {
	s.ChangeApplyId = &v
	return s
}

func (s *TrainOrderDetailQueryResponseBodyModuleChangeInfos) SetChangeTrainInfos(v []*TrainOrderDetailQueryResponseBodyModuleChangeInfosChangeTrainInfos) *TrainOrderDetailQueryResponseBodyModuleChangeInfos {
	s.ChangeTrainInfos = v
	return s
}

func (s *TrainOrderDetailQueryResponseBodyModuleChangeInfos) SetLimitPayTime(v string) *TrainOrderDetailQueryResponseBodyModuleChangeInfos {
	s.LimitPayTime = &v
	return s
}

func (s *TrainOrderDetailQueryResponseBodyModuleChangeInfos) SetOutChangeApplyId(v string) *TrainOrderDetailQueryResponseBodyModuleChangeInfos {
	s.OutChangeApplyId = &v
	return s
}

func (s *TrainOrderDetailQueryResponseBodyModuleChangeInfos) SetStatus(v string) *TrainOrderDetailQueryResponseBodyModuleChangeInfos {
	s.Status = &v
	return s
}

func (s *TrainOrderDetailQueryResponseBodyModuleChangeInfos) Validate() error {
	return dara.Validate(s)
}

type TrainOrderDetailQueryResponseBodyModuleChangeInfosChangeTrainInfos struct {
	// example:
	//
	// BDC
	ArrStationCode *string `json:"arr_station_code,omitempty" xml:"arr_station_code,omitempty"`
	ArrStationName *string `json:"arr_station_name,omitempty" xml:"arr_station_name,omitempty"`
	// example:
	//
	// 2024-05-06 15:19:01
	ArriveTime        *string                                                                                `json:"arrive_time,omitempty" xml:"arrive_time,omitempty"`
	ChangeTicketInfos []*TrainOrderDetailQueryResponseBodyModuleChangeInfosChangeTrainInfosChangeTicketInfos `json:"change_ticket_infos,omitempty" xml:"change_ticket_infos,omitempty" type:"Repeated"`
	// example:
	//
	// BTC
	DepStationCode *string `json:"dep_station_code,omitempty" xml:"dep_station_code,omitempty"`
	DepStationName *string `json:"dep_station_name,omitempty" xml:"dep_station_name,omitempty"`
	// example:
	//
	// 2024-05-06 15:19:01
	DepTime *string `json:"dep_time,omitempty" xml:"dep_time,omitempty"`
	// example:
	//
	// D1234
	TrainNo *string `json:"train_no,omitempty" xml:"train_no,omitempty"`
}

func (s TrainOrderDetailQueryResponseBodyModuleChangeInfosChangeTrainInfos) String() string {
	return dara.Prettify(s)
}

func (s TrainOrderDetailQueryResponseBodyModuleChangeInfosChangeTrainInfos) GoString() string {
	return s.String()
}

func (s *TrainOrderDetailQueryResponseBodyModuleChangeInfosChangeTrainInfos) GetArrStationCode() *string {
	return s.ArrStationCode
}

func (s *TrainOrderDetailQueryResponseBodyModuleChangeInfosChangeTrainInfos) GetArrStationName() *string {
	return s.ArrStationName
}

func (s *TrainOrderDetailQueryResponseBodyModuleChangeInfosChangeTrainInfos) GetArriveTime() *string {
	return s.ArriveTime
}

func (s *TrainOrderDetailQueryResponseBodyModuleChangeInfosChangeTrainInfos) GetChangeTicketInfos() []*TrainOrderDetailQueryResponseBodyModuleChangeInfosChangeTrainInfosChangeTicketInfos {
	return s.ChangeTicketInfos
}

func (s *TrainOrderDetailQueryResponseBodyModuleChangeInfosChangeTrainInfos) GetDepStationCode() *string {
	return s.DepStationCode
}

func (s *TrainOrderDetailQueryResponseBodyModuleChangeInfosChangeTrainInfos) GetDepStationName() *string {
	return s.DepStationName
}

func (s *TrainOrderDetailQueryResponseBodyModuleChangeInfosChangeTrainInfos) GetDepTime() *string {
	return s.DepTime
}

func (s *TrainOrderDetailQueryResponseBodyModuleChangeInfosChangeTrainInfos) GetTrainNo() *string {
	return s.TrainNo
}

func (s *TrainOrderDetailQueryResponseBodyModuleChangeInfosChangeTrainInfos) SetArrStationCode(v string) *TrainOrderDetailQueryResponseBodyModuleChangeInfosChangeTrainInfos {
	s.ArrStationCode = &v
	return s
}

func (s *TrainOrderDetailQueryResponseBodyModuleChangeInfosChangeTrainInfos) SetArrStationName(v string) *TrainOrderDetailQueryResponseBodyModuleChangeInfosChangeTrainInfos {
	s.ArrStationName = &v
	return s
}

func (s *TrainOrderDetailQueryResponseBodyModuleChangeInfosChangeTrainInfos) SetArriveTime(v string) *TrainOrderDetailQueryResponseBodyModuleChangeInfosChangeTrainInfos {
	s.ArriveTime = &v
	return s
}

func (s *TrainOrderDetailQueryResponseBodyModuleChangeInfosChangeTrainInfos) SetChangeTicketInfos(v []*TrainOrderDetailQueryResponseBodyModuleChangeInfosChangeTrainInfosChangeTicketInfos) *TrainOrderDetailQueryResponseBodyModuleChangeInfosChangeTrainInfos {
	s.ChangeTicketInfos = v
	return s
}

func (s *TrainOrderDetailQueryResponseBodyModuleChangeInfosChangeTrainInfos) SetDepStationCode(v string) *TrainOrderDetailQueryResponseBodyModuleChangeInfosChangeTrainInfos {
	s.DepStationCode = &v
	return s
}

func (s *TrainOrderDetailQueryResponseBodyModuleChangeInfosChangeTrainInfos) SetDepStationName(v string) *TrainOrderDetailQueryResponseBodyModuleChangeInfosChangeTrainInfos {
	s.DepStationName = &v
	return s
}

func (s *TrainOrderDetailQueryResponseBodyModuleChangeInfosChangeTrainInfos) SetDepTime(v string) *TrainOrderDetailQueryResponseBodyModuleChangeInfosChangeTrainInfos {
	s.DepTime = &v
	return s
}

func (s *TrainOrderDetailQueryResponseBodyModuleChangeInfosChangeTrainInfos) SetTrainNo(v string) *TrainOrderDetailQueryResponseBodyModuleChangeInfosChangeTrainInfos {
	s.TrainNo = &v
	return s
}

func (s *TrainOrderDetailQueryResponseBodyModuleChangeInfosChangeTrainInfos) Validate() error {
	return dara.Validate(s)
}

type TrainOrderDetailQueryResponseBodyModuleChangeInfosChangeTrainInfosChangeTicketInfos struct {
	// example:
	//
	// 0
	ChangeCost *int64 `json:"change_cost,omitempty" xml:"change_cost,omitempty"`
	// example:
	//
	// 0
	ChangeDiff *int64 `json:"change_diff,omitempty" xml:"change_diff,omitempty"`
	// example:
	//
	// 0
	ChangeGapHandingFee *int64 `json:"change_gap_handing_fee,omitempty" xml:"change_gap_handing_fee,omitempty"`
	// example:
	//
	// 0
	ChangeMinTicketAmountHandingFee *int64 `json:"change_min_ticket_amount_handing_fee,omitempty" xml:"change_min_ticket_amount_handing_fee,omitempty"`
	// example:
	//
	// 04
	CoachNo *string `json:"coach_no,omitempty" xml:"coach_no,omitempty"`
	// example:
	//
	// 4000
	FailCode   *string `json:"fail_code,omitempty" xml:"fail_code,omitempty"`
	FailReason *string `json:"fail_reason,omitempty" xml:"fail_reason,omitempty"`
	// example:
	//
	// 123456
	PassengerId *string `json:"passenger_id,omitempty" xml:"passenger_id,omitempty"`
	// example:
	//
	// 10000
	RealTicketPrice *int64 `json:"real_ticket_price,omitempty" xml:"real_ticket_price,omitempty"`
	// example:
	//
	// 1A
	SeatNo *string `json:"seat_no,omitempty" xml:"seat_no,omitempty"`
	// example:
	//
	// 14
	SeatType *string `json:"seat_type,omitempty" xml:"seat_type,omitempty"`
	// example:
	//
	// null
	TicketEntrance *string `json:"ticket_entrance,omitempty" xml:"ticket_entrance,omitempty"`
	// example:
	//
	// 1234
	TicketPrice *int64 `json:"ticket_price,omitempty" xml:"ticket_price,omitempty"`
	// example:
	//
	// 11
	TicketStatus *string `json:"ticket_status,omitempty" xml:"ticket_status,omitempty"`
}

func (s TrainOrderDetailQueryResponseBodyModuleChangeInfosChangeTrainInfosChangeTicketInfos) String() string {
	return dara.Prettify(s)
}

func (s TrainOrderDetailQueryResponseBodyModuleChangeInfosChangeTrainInfosChangeTicketInfos) GoString() string {
	return s.String()
}

func (s *TrainOrderDetailQueryResponseBodyModuleChangeInfosChangeTrainInfosChangeTicketInfos) GetChangeCost() *int64 {
	return s.ChangeCost
}

func (s *TrainOrderDetailQueryResponseBodyModuleChangeInfosChangeTrainInfosChangeTicketInfos) GetChangeDiff() *int64 {
	return s.ChangeDiff
}

func (s *TrainOrderDetailQueryResponseBodyModuleChangeInfosChangeTrainInfosChangeTicketInfos) GetChangeGapHandingFee() *int64 {
	return s.ChangeGapHandingFee
}

func (s *TrainOrderDetailQueryResponseBodyModuleChangeInfosChangeTrainInfosChangeTicketInfos) GetChangeMinTicketAmountHandingFee() *int64 {
	return s.ChangeMinTicketAmountHandingFee
}

func (s *TrainOrderDetailQueryResponseBodyModuleChangeInfosChangeTrainInfosChangeTicketInfos) GetCoachNo() *string {
	return s.CoachNo
}

func (s *TrainOrderDetailQueryResponseBodyModuleChangeInfosChangeTrainInfosChangeTicketInfos) GetFailCode() *string {
	return s.FailCode
}

func (s *TrainOrderDetailQueryResponseBodyModuleChangeInfosChangeTrainInfosChangeTicketInfos) GetFailReason() *string {
	return s.FailReason
}

func (s *TrainOrderDetailQueryResponseBodyModuleChangeInfosChangeTrainInfosChangeTicketInfos) GetPassengerId() *string {
	return s.PassengerId
}

func (s *TrainOrderDetailQueryResponseBodyModuleChangeInfosChangeTrainInfosChangeTicketInfos) GetRealTicketPrice() *int64 {
	return s.RealTicketPrice
}

func (s *TrainOrderDetailQueryResponseBodyModuleChangeInfosChangeTrainInfosChangeTicketInfos) GetSeatNo() *string {
	return s.SeatNo
}

func (s *TrainOrderDetailQueryResponseBodyModuleChangeInfosChangeTrainInfosChangeTicketInfos) GetSeatType() *string {
	return s.SeatType
}

func (s *TrainOrderDetailQueryResponseBodyModuleChangeInfosChangeTrainInfosChangeTicketInfos) GetTicketEntrance() *string {
	return s.TicketEntrance
}

func (s *TrainOrderDetailQueryResponseBodyModuleChangeInfosChangeTrainInfosChangeTicketInfos) GetTicketPrice() *int64 {
	return s.TicketPrice
}

func (s *TrainOrderDetailQueryResponseBodyModuleChangeInfosChangeTrainInfosChangeTicketInfos) GetTicketStatus() *string {
	return s.TicketStatus
}

func (s *TrainOrderDetailQueryResponseBodyModuleChangeInfosChangeTrainInfosChangeTicketInfos) SetChangeCost(v int64) *TrainOrderDetailQueryResponseBodyModuleChangeInfosChangeTrainInfosChangeTicketInfos {
	s.ChangeCost = &v
	return s
}

func (s *TrainOrderDetailQueryResponseBodyModuleChangeInfosChangeTrainInfosChangeTicketInfos) SetChangeDiff(v int64) *TrainOrderDetailQueryResponseBodyModuleChangeInfosChangeTrainInfosChangeTicketInfos {
	s.ChangeDiff = &v
	return s
}

func (s *TrainOrderDetailQueryResponseBodyModuleChangeInfosChangeTrainInfosChangeTicketInfos) SetChangeGapHandingFee(v int64) *TrainOrderDetailQueryResponseBodyModuleChangeInfosChangeTrainInfosChangeTicketInfos {
	s.ChangeGapHandingFee = &v
	return s
}

func (s *TrainOrderDetailQueryResponseBodyModuleChangeInfosChangeTrainInfosChangeTicketInfos) SetChangeMinTicketAmountHandingFee(v int64) *TrainOrderDetailQueryResponseBodyModuleChangeInfosChangeTrainInfosChangeTicketInfos {
	s.ChangeMinTicketAmountHandingFee = &v
	return s
}

func (s *TrainOrderDetailQueryResponseBodyModuleChangeInfosChangeTrainInfosChangeTicketInfos) SetCoachNo(v string) *TrainOrderDetailQueryResponseBodyModuleChangeInfosChangeTrainInfosChangeTicketInfos {
	s.CoachNo = &v
	return s
}

func (s *TrainOrderDetailQueryResponseBodyModuleChangeInfosChangeTrainInfosChangeTicketInfos) SetFailCode(v string) *TrainOrderDetailQueryResponseBodyModuleChangeInfosChangeTrainInfosChangeTicketInfos {
	s.FailCode = &v
	return s
}

func (s *TrainOrderDetailQueryResponseBodyModuleChangeInfosChangeTrainInfosChangeTicketInfos) SetFailReason(v string) *TrainOrderDetailQueryResponseBodyModuleChangeInfosChangeTrainInfosChangeTicketInfos {
	s.FailReason = &v
	return s
}

func (s *TrainOrderDetailQueryResponseBodyModuleChangeInfosChangeTrainInfosChangeTicketInfos) SetPassengerId(v string) *TrainOrderDetailQueryResponseBodyModuleChangeInfosChangeTrainInfosChangeTicketInfos {
	s.PassengerId = &v
	return s
}

func (s *TrainOrderDetailQueryResponseBodyModuleChangeInfosChangeTrainInfosChangeTicketInfos) SetRealTicketPrice(v int64) *TrainOrderDetailQueryResponseBodyModuleChangeInfosChangeTrainInfosChangeTicketInfos {
	s.RealTicketPrice = &v
	return s
}

func (s *TrainOrderDetailQueryResponseBodyModuleChangeInfosChangeTrainInfosChangeTicketInfos) SetSeatNo(v string) *TrainOrderDetailQueryResponseBodyModuleChangeInfosChangeTrainInfosChangeTicketInfos {
	s.SeatNo = &v
	return s
}

func (s *TrainOrderDetailQueryResponseBodyModuleChangeInfosChangeTrainInfosChangeTicketInfos) SetSeatType(v string) *TrainOrderDetailQueryResponseBodyModuleChangeInfosChangeTrainInfosChangeTicketInfos {
	s.SeatType = &v
	return s
}

func (s *TrainOrderDetailQueryResponseBodyModuleChangeInfosChangeTrainInfosChangeTicketInfos) SetTicketEntrance(v string) *TrainOrderDetailQueryResponseBodyModuleChangeInfosChangeTrainInfosChangeTicketInfos {
	s.TicketEntrance = &v
	return s
}

func (s *TrainOrderDetailQueryResponseBodyModuleChangeInfosChangeTrainInfosChangeTicketInfos) SetTicketPrice(v int64) *TrainOrderDetailQueryResponseBodyModuleChangeInfosChangeTrainInfosChangeTicketInfos {
	s.TicketPrice = &v
	return s
}

func (s *TrainOrderDetailQueryResponseBodyModuleChangeInfosChangeTrainInfosChangeTicketInfos) SetTicketStatus(v string) *TrainOrderDetailQueryResponseBodyModuleChangeInfosChangeTrainInfosChangeTicketInfos {
	s.TicketStatus = &v
	return s
}

func (s *TrainOrderDetailQueryResponseBodyModuleChangeInfosChangeTrainInfosChangeTicketInfos) Validate() error {
	return dara.Validate(s)
}

type TrainOrderDetailQueryResponseBodyModuleOfflineRefundDetails struct {
	// example:
	//
	// 123456
	OfflineRefundId    *string                                                                          `json:"offline_refund_id,omitempty" xml:"offline_refund_id,omitempty"`
	OfflineRefundInfos []*TrainOrderDetailQueryResponseBodyModuleOfflineRefundDetailsOfflineRefundInfos `json:"offline_refund_infos,omitempty" xml:"offline_refund_infos,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	OfflineRefundType *string `json:"offline_refund_type,omitempty" xml:"offline_refund_type,omitempty"`
	// example:
	//
	// 10000
	RefundTotalPrice *int64 `json:"refund_total_price,omitempty" xml:"refund_total_price,omitempty"`
}

func (s TrainOrderDetailQueryResponseBodyModuleOfflineRefundDetails) String() string {
	return dara.Prettify(s)
}

func (s TrainOrderDetailQueryResponseBodyModuleOfflineRefundDetails) GoString() string {
	return s.String()
}

func (s *TrainOrderDetailQueryResponseBodyModuleOfflineRefundDetails) GetOfflineRefundId() *string {
	return s.OfflineRefundId
}

func (s *TrainOrderDetailQueryResponseBodyModuleOfflineRefundDetails) GetOfflineRefundInfos() []*TrainOrderDetailQueryResponseBodyModuleOfflineRefundDetailsOfflineRefundInfos {
	return s.OfflineRefundInfos
}

func (s *TrainOrderDetailQueryResponseBodyModuleOfflineRefundDetails) GetOfflineRefundType() *string {
	return s.OfflineRefundType
}

func (s *TrainOrderDetailQueryResponseBodyModuleOfflineRefundDetails) GetRefundTotalPrice() *int64 {
	return s.RefundTotalPrice
}

func (s *TrainOrderDetailQueryResponseBodyModuleOfflineRefundDetails) SetOfflineRefundId(v string) *TrainOrderDetailQueryResponseBodyModuleOfflineRefundDetails {
	s.OfflineRefundId = &v
	return s
}

func (s *TrainOrderDetailQueryResponseBodyModuleOfflineRefundDetails) SetOfflineRefundInfos(v []*TrainOrderDetailQueryResponseBodyModuleOfflineRefundDetailsOfflineRefundInfos) *TrainOrderDetailQueryResponseBodyModuleOfflineRefundDetails {
	s.OfflineRefundInfos = v
	return s
}

func (s *TrainOrderDetailQueryResponseBodyModuleOfflineRefundDetails) SetOfflineRefundType(v string) *TrainOrderDetailQueryResponseBodyModuleOfflineRefundDetails {
	s.OfflineRefundType = &v
	return s
}

func (s *TrainOrderDetailQueryResponseBodyModuleOfflineRefundDetails) SetRefundTotalPrice(v int64) *TrainOrderDetailQueryResponseBodyModuleOfflineRefundDetails {
	s.RefundTotalPrice = &v
	return s
}

func (s *TrainOrderDetailQueryResponseBodyModuleOfflineRefundDetails) Validate() error {
	return dara.Validate(s)
}

type TrainOrderDetailQueryResponseBodyModuleOfflineRefundDetailsOfflineRefundInfos struct {
	// example:
	//
	// 123456
	PassengerId *string `json:"passenger_id,omitempty" xml:"passenger_id,omitempty"`
	// example:
	//
	// 111
	RefundPrice *int64 `json:"refund_price,omitempty" xml:"refund_price,omitempty"`
}

func (s TrainOrderDetailQueryResponseBodyModuleOfflineRefundDetailsOfflineRefundInfos) String() string {
	return dara.Prettify(s)
}

func (s TrainOrderDetailQueryResponseBodyModuleOfflineRefundDetailsOfflineRefundInfos) GoString() string {
	return s.String()
}

func (s *TrainOrderDetailQueryResponseBodyModuleOfflineRefundDetailsOfflineRefundInfos) GetPassengerId() *string {
	return s.PassengerId
}

func (s *TrainOrderDetailQueryResponseBodyModuleOfflineRefundDetailsOfflineRefundInfos) GetRefundPrice() *int64 {
	return s.RefundPrice
}

func (s *TrainOrderDetailQueryResponseBodyModuleOfflineRefundDetailsOfflineRefundInfos) SetPassengerId(v string) *TrainOrderDetailQueryResponseBodyModuleOfflineRefundDetailsOfflineRefundInfos {
	s.PassengerId = &v
	return s
}

func (s *TrainOrderDetailQueryResponseBodyModuleOfflineRefundDetailsOfflineRefundInfos) SetRefundPrice(v int64) *TrainOrderDetailQueryResponseBodyModuleOfflineRefundDetailsOfflineRefundInfos {
	s.RefundPrice = &v
	return s
}

func (s *TrainOrderDetailQueryResponseBodyModuleOfflineRefundDetailsOfflineRefundInfos) Validate() error {
	return dara.Validate(s)
}

type TrainOrderDetailQueryResponseBodyModulePassengerInfoS struct {
	CostCenterInfo *TrainOrderDetailQueryResponseBodyModulePassengerInfoSCostCenterInfo `json:"cost_center_info,omitempty" xml:"cost_center_info,omitempty" type:"Struct"`
	// example:
	//
	// 291487e553c5abde3b611aae283e2526f0d733ab55094aadc0b5ba587222a233c
	CountryCode *string `json:"country_code,omitempty" xml:"country_code,omitempty"`
	// example:
	//
	// 291487e553c5abde3b611aae283e2526f0d733ab55094aadc0b5ba587222a233c
	PassengerCertNo *string `json:"passenger_cert_no,omitempty" xml:"passenger_cert_no,omitempty"`
	// example:
	//
	// 170d9ac6f8807f9ec603c688f45f78a41
	PassengerCertType *string `json:"passenger_cert_type,omitempty" xml:"passenger_cert_type,omitempty"`
	// example:
	//
	// 123456
	PassengerId *string `json:"passenger_id,omitempty" xml:"passenger_id,omitempty"`
	// example:
	//
	// b6a6fc1bdf1ba60e25c2e132b612c8819
	PassengerMobile *string `json:"passenger_mobile,omitempty" xml:"passenger_mobile,omitempty"`
	// example:
	//
	// 949c9f34f677a0e5d249dfc94f5e62cc7
	PassengerName *string `json:"passenger_name,omitempty" xml:"passenger_name,omitempty"`
	// example:
	//
	// b6a6fc1bdf1ba60e25c2e132b612c8819
	ValidDateEnd *string `json:"valid_date_end,omitempty" xml:"valid_date_end,omitempty"`
}

func (s TrainOrderDetailQueryResponseBodyModulePassengerInfoS) String() string {
	return dara.Prettify(s)
}

func (s TrainOrderDetailQueryResponseBodyModulePassengerInfoS) GoString() string {
	return s.String()
}

func (s *TrainOrderDetailQueryResponseBodyModulePassengerInfoS) GetCostCenterInfo() *TrainOrderDetailQueryResponseBodyModulePassengerInfoSCostCenterInfo {
	return s.CostCenterInfo
}

func (s *TrainOrderDetailQueryResponseBodyModulePassengerInfoS) GetCountryCode() *string {
	return s.CountryCode
}

func (s *TrainOrderDetailQueryResponseBodyModulePassengerInfoS) GetPassengerCertNo() *string {
	return s.PassengerCertNo
}

func (s *TrainOrderDetailQueryResponseBodyModulePassengerInfoS) GetPassengerCertType() *string {
	return s.PassengerCertType
}

func (s *TrainOrderDetailQueryResponseBodyModulePassengerInfoS) GetPassengerId() *string {
	return s.PassengerId
}

func (s *TrainOrderDetailQueryResponseBodyModulePassengerInfoS) GetPassengerMobile() *string {
	return s.PassengerMobile
}

func (s *TrainOrderDetailQueryResponseBodyModulePassengerInfoS) GetPassengerName() *string {
	return s.PassengerName
}

func (s *TrainOrderDetailQueryResponseBodyModulePassengerInfoS) GetValidDateEnd() *string {
	return s.ValidDateEnd
}

func (s *TrainOrderDetailQueryResponseBodyModulePassengerInfoS) SetCostCenterInfo(v *TrainOrderDetailQueryResponseBodyModulePassengerInfoSCostCenterInfo) *TrainOrderDetailQueryResponseBodyModulePassengerInfoS {
	s.CostCenterInfo = v
	return s
}

func (s *TrainOrderDetailQueryResponseBodyModulePassengerInfoS) SetCountryCode(v string) *TrainOrderDetailQueryResponseBodyModulePassengerInfoS {
	s.CountryCode = &v
	return s
}

func (s *TrainOrderDetailQueryResponseBodyModulePassengerInfoS) SetPassengerCertNo(v string) *TrainOrderDetailQueryResponseBodyModulePassengerInfoS {
	s.PassengerCertNo = &v
	return s
}

func (s *TrainOrderDetailQueryResponseBodyModulePassengerInfoS) SetPassengerCertType(v string) *TrainOrderDetailQueryResponseBodyModulePassengerInfoS {
	s.PassengerCertType = &v
	return s
}

func (s *TrainOrderDetailQueryResponseBodyModulePassengerInfoS) SetPassengerId(v string) *TrainOrderDetailQueryResponseBodyModulePassengerInfoS {
	s.PassengerId = &v
	return s
}

func (s *TrainOrderDetailQueryResponseBodyModulePassengerInfoS) SetPassengerMobile(v string) *TrainOrderDetailQueryResponseBodyModulePassengerInfoS {
	s.PassengerMobile = &v
	return s
}

func (s *TrainOrderDetailQueryResponseBodyModulePassengerInfoS) SetPassengerName(v string) *TrainOrderDetailQueryResponseBodyModulePassengerInfoS {
	s.PassengerName = &v
	return s
}

func (s *TrainOrderDetailQueryResponseBodyModulePassengerInfoS) SetValidDateEnd(v string) *TrainOrderDetailQueryResponseBodyModulePassengerInfoS {
	s.ValidDateEnd = &v
	return s
}

func (s *TrainOrderDetailQueryResponseBodyModulePassengerInfoS) Validate() error {
	return dara.Validate(s)
}

type TrainOrderDetailQueryResponseBodyModulePassengerInfoSCostCenterInfo struct {
	CascadeDeptName *string `json:"cascade_dept_name,omitempty" xml:"cascade_dept_name,omitempty"`
	// example:
	//
	// 111111
	CostCenterId   *string `json:"cost_center_id,omitempty" xml:"cost_center_id,omitempty"`
	CostCenterName *string `json:"cost_center_name,omitempty" xml:"cost_center_name,omitempty"`
	// example:
	//
	// 123456
	CostCenterNo *string `json:"cost_center_no,omitempty" xml:"cost_center_no,omitempty"`
	// example:
	//
	// 582000002311
	DepartId   *string `json:"depart_id,omitempty" xml:"depart_id,omitempty"`
	DepartName *string `json:"depart_name,omitempty" xml:"depart_name,omitempty"`
	// example:
	//
	// 11111
	InvoiceId    *string `json:"invoice_id,omitempty" xml:"invoice_id,omitempty"`
	InvoiceTitle *string `json:"invoice_title,omitempty" xml:"invoice_title,omitempty"`
	// example:
	//
	// 123456
	PassengerId *string `json:"passenger_id,omitempty" xml:"passenger_id,omitempty"`
	// example:
	//
	// 1234
	ProjectCode  *string `json:"project_code,omitempty" xml:"project_code,omitempty"`
	ProjectTitle *string `json:"project_title,omitempty" xml:"project_title,omitempty"`
}

func (s TrainOrderDetailQueryResponseBodyModulePassengerInfoSCostCenterInfo) String() string {
	return dara.Prettify(s)
}

func (s TrainOrderDetailQueryResponseBodyModulePassengerInfoSCostCenterInfo) GoString() string {
	return s.String()
}

func (s *TrainOrderDetailQueryResponseBodyModulePassengerInfoSCostCenterInfo) GetCascadeDeptName() *string {
	return s.CascadeDeptName
}

func (s *TrainOrderDetailQueryResponseBodyModulePassengerInfoSCostCenterInfo) GetCostCenterId() *string {
	return s.CostCenterId
}

func (s *TrainOrderDetailQueryResponseBodyModulePassengerInfoSCostCenterInfo) GetCostCenterName() *string {
	return s.CostCenterName
}

func (s *TrainOrderDetailQueryResponseBodyModulePassengerInfoSCostCenterInfo) GetCostCenterNo() *string {
	return s.CostCenterNo
}

func (s *TrainOrderDetailQueryResponseBodyModulePassengerInfoSCostCenterInfo) GetDepartId() *string {
	return s.DepartId
}

func (s *TrainOrderDetailQueryResponseBodyModulePassengerInfoSCostCenterInfo) GetDepartName() *string {
	return s.DepartName
}

func (s *TrainOrderDetailQueryResponseBodyModulePassengerInfoSCostCenterInfo) GetInvoiceId() *string {
	return s.InvoiceId
}

func (s *TrainOrderDetailQueryResponseBodyModulePassengerInfoSCostCenterInfo) GetInvoiceTitle() *string {
	return s.InvoiceTitle
}

func (s *TrainOrderDetailQueryResponseBodyModulePassengerInfoSCostCenterInfo) GetPassengerId() *string {
	return s.PassengerId
}

func (s *TrainOrderDetailQueryResponseBodyModulePassengerInfoSCostCenterInfo) GetProjectCode() *string {
	return s.ProjectCode
}

func (s *TrainOrderDetailQueryResponseBodyModulePassengerInfoSCostCenterInfo) GetProjectTitle() *string {
	return s.ProjectTitle
}

func (s *TrainOrderDetailQueryResponseBodyModulePassengerInfoSCostCenterInfo) SetCascadeDeptName(v string) *TrainOrderDetailQueryResponseBodyModulePassengerInfoSCostCenterInfo {
	s.CascadeDeptName = &v
	return s
}

func (s *TrainOrderDetailQueryResponseBodyModulePassengerInfoSCostCenterInfo) SetCostCenterId(v string) *TrainOrderDetailQueryResponseBodyModulePassengerInfoSCostCenterInfo {
	s.CostCenterId = &v
	return s
}

func (s *TrainOrderDetailQueryResponseBodyModulePassengerInfoSCostCenterInfo) SetCostCenterName(v string) *TrainOrderDetailQueryResponseBodyModulePassengerInfoSCostCenterInfo {
	s.CostCenterName = &v
	return s
}

func (s *TrainOrderDetailQueryResponseBodyModulePassengerInfoSCostCenterInfo) SetCostCenterNo(v string) *TrainOrderDetailQueryResponseBodyModulePassengerInfoSCostCenterInfo {
	s.CostCenterNo = &v
	return s
}

func (s *TrainOrderDetailQueryResponseBodyModulePassengerInfoSCostCenterInfo) SetDepartId(v string) *TrainOrderDetailQueryResponseBodyModulePassengerInfoSCostCenterInfo {
	s.DepartId = &v
	return s
}

func (s *TrainOrderDetailQueryResponseBodyModulePassengerInfoSCostCenterInfo) SetDepartName(v string) *TrainOrderDetailQueryResponseBodyModulePassengerInfoSCostCenterInfo {
	s.DepartName = &v
	return s
}

func (s *TrainOrderDetailQueryResponseBodyModulePassengerInfoSCostCenterInfo) SetInvoiceId(v string) *TrainOrderDetailQueryResponseBodyModulePassengerInfoSCostCenterInfo {
	s.InvoiceId = &v
	return s
}

func (s *TrainOrderDetailQueryResponseBodyModulePassengerInfoSCostCenterInfo) SetInvoiceTitle(v string) *TrainOrderDetailQueryResponseBodyModulePassengerInfoSCostCenterInfo {
	s.InvoiceTitle = &v
	return s
}

func (s *TrainOrderDetailQueryResponseBodyModulePassengerInfoSCostCenterInfo) SetPassengerId(v string) *TrainOrderDetailQueryResponseBodyModulePassengerInfoSCostCenterInfo {
	s.PassengerId = &v
	return s
}

func (s *TrainOrderDetailQueryResponseBodyModulePassengerInfoSCostCenterInfo) SetProjectCode(v string) *TrainOrderDetailQueryResponseBodyModulePassengerInfoSCostCenterInfo {
	s.ProjectCode = &v
	return s
}

func (s *TrainOrderDetailQueryResponseBodyModulePassengerInfoSCostCenterInfo) SetProjectTitle(v string) *TrainOrderDetailQueryResponseBodyModulePassengerInfoSCostCenterInfo {
	s.ProjectTitle = &v
	return s
}

func (s *TrainOrderDetailQueryResponseBodyModulePassengerInfoSCostCenterInfo) Validate() error {
	return dara.Validate(s)
}

type TrainOrderDetailQueryResponseBodyModuleRefundInfos struct {
	// example:
	//
	// 1111
	FailCode *string `json:"fail_code,omitempty" xml:"fail_code,omitempty"`
	FailMsg  *string `json:"fail_msg,omitempty" xml:"fail_msg,omitempty"`
	// example:
	//
	// 123456778
	OutRefundId *string `json:"out_refund_id,omitempty" xml:"out_refund_id,omitempty"`
	// String
	//
	// example:
	//
	// 123456
	RefundId        *string                                                              `json:"refund_id,omitempty" xml:"refund_id,omitempty"`
	RefundTrainInfo []*TrainOrderDetailQueryResponseBodyModuleRefundInfosRefundTrainInfo `json:"refund_train_info,omitempty" xml:"refund_train_info,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s TrainOrderDetailQueryResponseBodyModuleRefundInfos) String() string {
	return dara.Prettify(s)
}

func (s TrainOrderDetailQueryResponseBodyModuleRefundInfos) GoString() string {
	return s.String()
}

func (s *TrainOrderDetailQueryResponseBodyModuleRefundInfos) GetFailCode() *string {
	return s.FailCode
}

func (s *TrainOrderDetailQueryResponseBodyModuleRefundInfos) GetFailMsg() *string {
	return s.FailMsg
}

func (s *TrainOrderDetailQueryResponseBodyModuleRefundInfos) GetOutRefundId() *string {
	return s.OutRefundId
}

func (s *TrainOrderDetailQueryResponseBodyModuleRefundInfos) GetRefundId() *string {
	return s.RefundId
}

func (s *TrainOrderDetailQueryResponseBodyModuleRefundInfos) GetRefundTrainInfo() []*TrainOrderDetailQueryResponseBodyModuleRefundInfosRefundTrainInfo {
	return s.RefundTrainInfo
}

func (s *TrainOrderDetailQueryResponseBodyModuleRefundInfos) GetStatus() *string {
	return s.Status
}

func (s *TrainOrderDetailQueryResponseBodyModuleRefundInfos) SetFailCode(v string) *TrainOrderDetailQueryResponseBodyModuleRefundInfos {
	s.FailCode = &v
	return s
}

func (s *TrainOrderDetailQueryResponseBodyModuleRefundInfos) SetFailMsg(v string) *TrainOrderDetailQueryResponseBodyModuleRefundInfos {
	s.FailMsg = &v
	return s
}

func (s *TrainOrderDetailQueryResponseBodyModuleRefundInfos) SetOutRefundId(v string) *TrainOrderDetailQueryResponseBodyModuleRefundInfos {
	s.OutRefundId = &v
	return s
}

func (s *TrainOrderDetailQueryResponseBodyModuleRefundInfos) SetRefundId(v string) *TrainOrderDetailQueryResponseBodyModuleRefundInfos {
	s.RefundId = &v
	return s
}

func (s *TrainOrderDetailQueryResponseBodyModuleRefundInfos) SetRefundTrainInfo(v []*TrainOrderDetailQueryResponseBodyModuleRefundInfosRefundTrainInfo) *TrainOrderDetailQueryResponseBodyModuleRefundInfos {
	s.RefundTrainInfo = v
	return s
}

func (s *TrainOrderDetailQueryResponseBodyModuleRefundInfos) SetStatus(v string) *TrainOrderDetailQueryResponseBodyModuleRefundInfos {
	s.Status = &v
	return s
}

func (s *TrainOrderDetailQueryResponseBodyModuleRefundInfos) Validate() error {
	return dara.Validate(s)
}

type TrainOrderDetailQueryResponseBodyModuleRefundInfosRefundTrainInfo struct {
	ArrStationName *string `json:"arr_station_name,omitempty" xml:"arr_station_name,omitempty"`
	// example:
	//
	// 2024-05-06 15:19:01
	ArrTime *string `json:"arr_time,omitempty" xml:"arr_time,omitempty"`
	// example:
	//
	// BTC
	DepStationCode *string `json:"dep_station_code,omitempty" xml:"dep_station_code,omitempty"`
	DepStationName *string `json:"dep_station_name,omitempty" xml:"dep_station_name,omitempty"`
	// example:
	//
	// 2024-05-06 15:19:01
	DepTime           *string                                                                               `json:"dep_time,omitempty" xml:"dep_time,omitempty"`
	RefundTicketInfos []*TrainOrderDetailQueryResponseBodyModuleRefundInfosRefundTrainInfoRefundTicketInfos `json:"refund_ticket_infos,omitempty" xml:"refund_ticket_infos,omitempty" type:"Repeated"`
	// example:
	//
	// K1234
	TrainNo *string `json:"train_no,omitempty" xml:"train_no,omitempty"`
}

func (s TrainOrderDetailQueryResponseBodyModuleRefundInfosRefundTrainInfo) String() string {
	return dara.Prettify(s)
}

func (s TrainOrderDetailQueryResponseBodyModuleRefundInfosRefundTrainInfo) GoString() string {
	return s.String()
}

func (s *TrainOrderDetailQueryResponseBodyModuleRefundInfosRefundTrainInfo) GetArrStationName() *string {
	return s.ArrStationName
}

func (s *TrainOrderDetailQueryResponseBodyModuleRefundInfosRefundTrainInfo) GetArrTime() *string {
	return s.ArrTime
}

func (s *TrainOrderDetailQueryResponseBodyModuleRefundInfosRefundTrainInfo) GetDepStationCode() *string {
	return s.DepStationCode
}

func (s *TrainOrderDetailQueryResponseBodyModuleRefundInfosRefundTrainInfo) GetDepStationName() *string {
	return s.DepStationName
}

func (s *TrainOrderDetailQueryResponseBodyModuleRefundInfosRefundTrainInfo) GetDepTime() *string {
	return s.DepTime
}

func (s *TrainOrderDetailQueryResponseBodyModuleRefundInfosRefundTrainInfo) GetRefundTicketInfos() []*TrainOrderDetailQueryResponseBodyModuleRefundInfosRefundTrainInfoRefundTicketInfos {
	return s.RefundTicketInfos
}

func (s *TrainOrderDetailQueryResponseBodyModuleRefundInfosRefundTrainInfo) GetTrainNo() *string {
	return s.TrainNo
}

func (s *TrainOrderDetailQueryResponseBodyModuleRefundInfosRefundTrainInfo) SetArrStationName(v string) *TrainOrderDetailQueryResponseBodyModuleRefundInfosRefundTrainInfo {
	s.ArrStationName = &v
	return s
}

func (s *TrainOrderDetailQueryResponseBodyModuleRefundInfosRefundTrainInfo) SetArrTime(v string) *TrainOrderDetailQueryResponseBodyModuleRefundInfosRefundTrainInfo {
	s.ArrTime = &v
	return s
}

func (s *TrainOrderDetailQueryResponseBodyModuleRefundInfosRefundTrainInfo) SetDepStationCode(v string) *TrainOrderDetailQueryResponseBodyModuleRefundInfosRefundTrainInfo {
	s.DepStationCode = &v
	return s
}

func (s *TrainOrderDetailQueryResponseBodyModuleRefundInfosRefundTrainInfo) SetDepStationName(v string) *TrainOrderDetailQueryResponseBodyModuleRefundInfosRefundTrainInfo {
	s.DepStationName = &v
	return s
}

func (s *TrainOrderDetailQueryResponseBodyModuleRefundInfosRefundTrainInfo) SetDepTime(v string) *TrainOrderDetailQueryResponseBodyModuleRefundInfosRefundTrainInfo {
	s.DepTime = &v
	return s
}

func (s *TrainOrderDetailQueryResponseBodyModuleRefundInfosRefundTrainInfo) SetRefundTicketInfos(v []*TrainOrderDetailQueryResponseBodyModuleRefundInfosRefundTrainInfoRefundTicketInfos) *TrainOrderDetailQueryResponseBodyModuleRefundInfosRefundTrainInfo {
	s.RefundTicketInfos = v
	return s
}

func (s *TrainOrderDetailQueryResponseBodyModuleRefundInfosRefundTrainInfo) SetTrainNo(v string) *TrainOrderDetailQueryResponseBodyModuleRefundInfosRefundTrainInfo {
	s.TrainNo = &v
	return s
}

func (s *TrainOrderDetailQueryResponseBodyModuleRefundInfosRefundTrainInfo) Validate() error {
	return dara.Validate(s)
}

type TrainOrderDetailQueryResponseBodyModuleRefundInfosRefundTrainInfoRefundTicketInfos struct {
	// example:
	//
	// 123456
	PassengerId *string `json:"passenger_id,omitempty" xml:"passenger_id,omitempty"`
	// example:
	//
	// 0
	RefundCost *int64 `json:"refund_cost,omitempty" xml:"refund_cost,omitempty"`
	// example:
	//
	// 10000
	RefundPrice *int64 `json:"refund_price,omitempty" xml:"refund_price,omitempty"`
	// example:
	//
	// 10000
	TicketPrice *int64 `json:"ticket_price,omitempty" xml:"ticket_price,omitempty"`
}

func (s TrainOrderDetailQueryResponseBodyModuleRefundInfosRefundTrainInfoRefundTicketInfos) String() string {
	return dara.Prettify(s)
}

func (s TrainOrderDetailQueryResponseBodyModuleRefundInfosRefundTrainInfoRefundTicketInfos) GoString() string {
	return s.String()
}

func (s *TrainOrderDetailQueryResponseBodyModuleRefundInfosRefundTrainInfoRefundTicketInfos) GetPassengerId() *string {
	return s.PassengerId
}

func (s *TrainOrderDetailQueryResponseBodyModuleRefundInfosRefundTrainInfoRefundTicketInfos) GetRefundCost() *int64 {
	return s.RefundCost
}

func (s *TrainOrderDetailQueryResponseBodyModuleRefundInfosRefundTrainInfoRefundTicketInfos) GetRefundPrice() *int64 {
	return s.RefundPrice
}

func (s *TrainOrderDetailQueryResponseBodyModuleRefundInfosRefundTrainInfoRefundTicketInfos) GetTicketPrice() *int64 {
	return s.TicketPrice
}

func (s *TrainOrderDetailQueryResponseBodyModuleRefundInfosRefundTrainInfoRefundTicketInfos) SetPassengerId(v string) *TrainOrderDetailQueryResponseBodyModuleRefundInfosRefundTrainInfoRefundTicketInfos {
	s.PassengerId = &v
	return s
}

func (s *TrainOrderDetailQueryResponseBodyModuleRefundInfosRefundTrainInfoRefundTicketInfos) SetRefundCost(v int64) *TrainOrderDetailQueryResponseBodyModuleRefundInfosRefundTrainInfoRefundTicketInfos {
	s.RefundCost = &v
	return s
}

func (s *TrainOrderDetailQueryResponseBodyModuleRefundInfosRefundTrainInfoRefundTicketInfos) SetRefundPrice(v int64) *TrainOrderDetailQueryResponseBodyModuleRefundInfosRefundTrainInfoRefundTicketInfos {
	s.RefundPrice = &v
	return s
}

func (s *TrainOrderDetailQueryResponseBodyModuleRefundInfosRefundTrainInfoRefundTicketInfos) SetTicketPrice(v int64) *TrainOrderDetailQueryResponseBodyModuleRefundInfosRefundTrainInfoRefundTicketInfos {
	s.TicketPrice = &v
	return s
}

func (s *TrainOrderDetailQueryResponseBodyModuleRefundInfosRefundTrainInfoRefundTicketInfos) Validate() error {
	return dara.Validate(s)
}
