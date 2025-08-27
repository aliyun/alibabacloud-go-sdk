// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHotelOrderChangeDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *HotelOrderChangeDetailResponseBody
	GetCode() *string
	SetMessage(v string) *HotelOrderChangeDetailResponseBody
	GetMessage() *string
	SetModule(v *HotelOrderChangeDetailResponseBodyModule) *HotelOrderChangeDetailResponseBody
	GetModule() *HotelOrderChangeDetailResponseBodyModule
	SetRequestId(v string) *HotelOrderChangeDetailResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *HotelOrderChangeDetailResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *HotelOrderChangeDetailResponseBody
	GetTraceId() *string
}

type HotelOrderChangeDetailResponseBody struct {
	// example:
	//
	// 0
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// operation success
	Message *string                                   `json:"message,omitempty" xml:"message,omitempty"`
	Module  *HotelOrderChangeDetailResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
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

func (s HotelOrderChangeDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s HotelOrderChangeDetailResponseBody) GoString() string {
	return s.String()
}

func (s *HotelOrderChangeDetailResponseBody) GetCode() *string {
	return s.Code
}

func (s *HotelOrderChangeDetailResponseBody) GetMessage() *string {
	return s.Message
}

func (s *HotelOrderChangeDetailResponseBody) GetModule() *HotelOrderChangeDetailResponseBodyModule {
	return s.Module
}

func (s *HotelOrderChangeDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *HotelOrderChangeDetailResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *HotelOrderChangeDetailResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *HotelOrderChangeDetailResponseBody) SetCode(v string) *HotelOrderChangeDetailResponseBody {
	s.Code = &v
	return s
}

func (s *HotelOrderChangeDetailResponseBody) SetMessage(v string) *HotelOrderChangeDetailResponseBody {
	s.Message = &v
	return s
}

func (s *HotelOrderChangeDetailResponseBody) SetModule(v *HotelOrderChangeDetailResponseBodyModule) *HotelOrderChangeDetailResponseBody {
	s.Module = v
	return s
}

func (s *HotelOrderChangeDetailResponseBody) SetRequestId(v string) *HotelOrderChangeDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *HotelOrderChangeDetailResponseBody) SetSuccess(v bool) *HotelOrderChangeDetailResponseBody {
	s.Success = &v
	return s
}

func (s *HotelOrderChangeDetailResponseBody) SetTraceId(v string) *HotelOrderChangeDetailResponseBody {
	s.TraceId = &v
	return s
}

func (s *HotelOrderChangeDetailResponseBody) Validate() error {
	return dara.Validate(s)
}

type HotelOrderChangeDetailResponseBodyModule struct {
	// example:
	//
	// 1234
	ChangeOrderId *string `json:"change_order_id,omitempty" xml:"change_order_id,omitempty"`
	// example:
	//
	// 0
	ChangeType *int32 `json:"change_type,omitempty" xml:"change_type,omitempty"`
	// example:
	//
	// open12ih3c8jb8o47v10B4r4josN00
	CorpId *string `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
	// example:
	//
	// dis1234
	DisOrderId *string `json:"dis_order_id,omitempty" xml:"dis_order_id,omitempty"`
	// example:
	//
	// 2024-07-07 13:42:49
	GmtCreate *string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// example:
	//
	// 2024-07-07 13:42:49
	GmtModified  *string                                                 `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	Reason       *string                                                 `json:"reason,omitempty" xml:"reason,omitempty"`
	Remarks      *string                                                 `json:"remarks,omitempty" xml:"remarks,omitempty"`
	RoomInfoList []*HotelOrderChangeDetailResponseBodyModuleRoomInfoList `json:"room_info_list,omitempty" xml:"room_info_list,omitempty" type:"Repeated"`
	// example:
	//
	// 1402002197440511306
	SaleOrderId *string `json:"sale_order_id,omitempty" xml:"sale_order_id,omitempty"`
	// example:
	//
	// 0
	Source *int32 `json:"source,omitempty" xml:"source,omitempty"`
	// example:
	//
	// 1
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 2505048378320666
	WorkOrderId *string `json:"work_order_id,omitempty" xml:"work_order_id,omitempty"`
}

func (s HotelOrderChangeDetailResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s HotelOrderChangeDetailResponseBodyModule) GoString() string {
	return s.String()
}

func (s *HotelOrderChangeDetailResponseBodyModule) GetChangeOrderId() *string {
	return s.ChangeOrderId
}

func (s *HotelOrderChangeDetailResponseBodyModule) GetChangeType() *int32 {
	return s.ChangeType
}

func (s *HotelOrderChangeDetailResponseBodyModule) GetCorpId() *string {
	return s.CorpId
}

func (s *HotelOrderChangeDetailResponseBodyModule) GetDisOrderId() *string {
	return s.DisOrderId
}

func (s *HotelOrderChangeDetailResponseBodyModule) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *HotelOrderChangeDetailResponseBodyModule) GetGmtModified() *string {
	return s.GmtModified
}

func (s *HotelOrderChangeDetailResponseBodyModule) GetReason() *string {
	return s.Reason
}

func (s *HotelOrderChangeDetailResponseBodyModule) GetRemarks() *string {
	return s.Remarks
}

func (s *HotelOrderChangeDetailResponseBodyModule) GetRoomInfoList() []*HotelOrderChangeDetailResponseBodyModuleRoomInfoList {
	return s.RoomInfoList
}

func (s *HotelOrderChangeDetailResponseBodyModule) GetSaleOrderId() *string {
	return s.SaleOrderId
}

func (s *HotelOrderChangeDetailResponseBodyModule) GetSource() *int32 {
	return s.Source
}

func (s *HotelOrderChangeDetailResponseBodyModule) GetStatus() *int32 {
	return s.Status
}

func (s *HotelOrderChangeDetailResponseBodyModule) GetWorkOrderId() *string {
	return s.WorkOrderId
}

func (s *HotelOrderChangeDetailResponseBodyModule) SetChangeOrderId(v string) *HotelOrderChangeDetailResponseBodyModule {
	s.ChangeOrderId = &v
	return s
}

func (s *HotelOrderChangeDetailResponseBodyModule) SetChangeType(v int32) *HotelOrderChangeDetailResponseBodyModule {
	s.ChangeType = &v
	return s
}

func (s *HotelOrderChangeDetailResponseBodyModule) SetCorpId(v string) *HotelOrderChangeDetailResponseBodyModule {
	s.CorpId = &v
	return s
}

func (s *HotelOrderChangeDetailResponseBodyModule) SetDisOrderId(v string) *HotelOrderChangeDetailResponseBodyModule {
	s.DisOrderId = &v
	return s
}

func (s *HotelOrderChangeDetailResponseBodyModule) SetGmtCreate(v string) *HotelOrderChangeDetailResponseBodyModule {
	s.GmtCreate = &v
	return s
}

func (s *HotelOrderChangeDetailResponseBodyModule) SetGmtModified(v string) *HotelOrderChangeDetailResponseBodyModule {
	s.GmtModified = &v
	return s
}

func (s *HotelOrderChangeDetailResponseBodyModule) SetReason(v string) *HotelOrderChangeDetailResponseBodyModule {
	s.Reason = &v
	return s
}

func (s *HotelOrderChangeDetailResponseBodyModule) SetRemarks(v string) *HotelOrderChangeDetailResponseBodyModule {
	s.Remarks = &v
	return s
}

func (s *HotelOrderChangeDetailResponseBodyModule) SetRoomInfoList(v []*HotelOrderChangeDetailResponseBodyModuleRoomInfoList) *HotelOrderChangeDetailResponseBodyModule {
	s.RoomInfoList = v
	return s
}

func (s *HotelOrderChangeDetailResponseBodyModule) SetSaleOrderId(v string) *HotelOrderChangeDetailResponseBodyModule {
	s.SaleOrderId = &v
	return s
}

func (s *HotelOrderChangeDetailResponseBodyModule) SetSource(v int32) *HotelOrderChangeDetailResponseBodyModule {
	s.Source = &v
	return s
}

func (s *HotelOrderChangeDetailResponseBodyModule) SetStatus(v int32) *HotelOrderChangeDetailResponseBodyModule {
	s.Status = &v
	return s
}

func (s *HotelOrderChangeDetailResponseBodyModule) SetWorkOrderId(v string) *HotelOrderChangeDetailResponseBodyModule {
	s.WorkOrderId = &v
	return s
}

func (s *HotelOrderChangeDetailResponseBodyModule) Validate() error {
	return dara.Validate(s)
}

type HotelOrderChangeDetailResponseBodyModuleRoomInfoList struct {
	CancelDate           []*string                                                                   `json:"cancel_date,omitempty" xml:"cancel_date,omitempty" type:"Repeated"`
	RoomDailyRefundInfos []*HotelOrderChangeDetailResponseBodyModuleRoomInfoListRoomDailyRefundInfos `json:"room_daily_refund_infos,omitempty" xml:"room_daily_refund_infos,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	RoomNo *int32 `json:"room_no,omitempty" xml:"room_no,omitempty"`
}

func (s HotelOrderChangeDetailResponseBodyModuleRoomInfoList) String() string {
	return dara.Prettify(s)
}

func (s HotelOrderChangeDetailResponseBodyModuleRoomInfoList) GoString() string {
	return s.String()
}

func (s *HotelOrderChangeDetailResponseBodyModuleRoomInfoList) GetCancelDate() []*string {
	return s.CancelDate
}

func (s *HotelOrderChangeDetailResponseBodyModuleRoomInfoList) GetRoomDailyRefundInfos() []*HotelOrderChangeDetailResponseBodyModuleRoomInfoListRoomDailyRefundInfos {
	return s.RoomDailyRefundInfos
}

func (s *HotelOrderChangeDetailResponseBodyModuleRoomInfoList) GetRoomNo() *int32 {
	return s.RoomNo
}

func (s *HotelOrderChangeDetailResponseBodyModuleRoomInfoList) SetCancelDate(v []*string) *HotelOrderChangeDetailResponseBodyModuleRoomInfoList {
	s.CancelDate = v
	return s
}

func (s *HotelOrderChangeDetailResponseBodyModuleRoomInfoList) SetRoomDailyRefundInfos(v []*HotelOrderChangeDetailResponseBodyModuleRoomInfoListRoomDailyRefundInfos) *HotelOrderChangeDetailResponseBodyModuleRoomInfoList {
	s.RoomDailyRefundInfos = v
	return s
}

func (s *HotelOrderChangeDetailResponseBodyModuleRoomInfoList) SetRoomNo(v int32) *HotelOrderChangeDetailResponseBodyModuleRoomInfoList {
	s.RoomNo = &v
	return s
}

func (s *HotelOrderChangeDetailResponseBodyModuleRoomInfoList) Validate() error {
	return dara.Validate(s)
}

type HotelOrderChangeDetailResponseBodyModuleRoomInfoListRoomDailyRefundInfos struct {
	// example:
	//
	// 2024-02-10
	CheckInDate *string `json:"check_in_date,omitempty" xml:"check_in_date,omitempty"`
	// example:
	//
	// 71000
	PenaltyPrice *int64 `json:"penalty_price,omitempty" xml:"penalty_price,omitempty"`
	// example:
	//
	// 71000
	Price *int64 `json:"price,omitempty" xml:"price,omitempty"`
	// example:
	//
	// 71000
	RefundPrice *int64 `json:"refund_price,omitempty" xml:"refund_price,omitempty"`
}

func (s HotelOrderChangeDetailResponseBodyModuleRoomInfoListRoomDailyRefundInfos) String() string {
	return dara.Prettify(s)
}

func (s HotelOrderChangeDetailResponseBodyModuleRoomInfoListRoomDailyRefundInfos) GoString() string {
	return s.String()
}

func (s *HotelOrderChangeDetailResponseBodyModuleRoomInfoListRoomDailyRefundInfos) GetCheckInDate() *string {
	return s.CheckInDate
}

func (s *HotelOrderChangeDetailResponseBodyModuleRoomInfoListRoomDailyRefundInfos) GetPenaltyPrice() *int64 {
	return s.PenaltyPrice
}

func (s *HotelOrderChangeDetailResponseBodyModuleRoomInfoListRoomDailyRefundInfos) GetPrice() *int64 {
	return s.Price
}

func (s *HotelOrderChangeDetailResponseBodyModuleRoomInfoListRoomDailyRefundInfos) GetRefundPrice() *int64 {
	return s.RefundPrice
}

func (s *HotelOrderChangeDetailResponseBodyModuleRoomInfoListRoomDailyRefundInfos) SetCheckInDate(v string) *HotelOrderChangeDetailResponseBodyModuleRoomInfoListRoomDailyRefundInfos {
	s.CheckInDate = &v
	return s
}

func (s *HotelOrderChangeDetailResponseBodyModuleRoomInfoListRoomDailyRefundInfos) SetPenaltyPrice(v int64) *HotelOrderChangeDetailResponseBodyModuleRoomInfoListRoomDailyRefundInfos {
	s.PenaltyPrice = &v
	return s
}

func (s *HotelOrderChangeDetailResponseBodyModuleRoomInfoListRoomDailyRefundInfos) SetPrice(v int64) *HotelOrderChangeDetailResponseBodyModuleRoomInfoListRoomDailyRefundInfos {
	s.Price = &v
	return s
}

func (s *HotelOrderChangeDetailResponseBodyModuleRoomInfoListRoomDailyRefundInfos) SetRefundPrice(v int64) *HotelOrderChangeDetailResponseBodyModuleRoomInfoListRoomDailyRefundInfos {
	s.RefundPrice = &v
	return s
}

func (s *HotelOrderChangeDetailResponseBodyModuleRoomInfoListRoomDailyRefundInfos) Validate() error {
	return dara.Validate(s)
}
