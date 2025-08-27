// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHotelOrderDetailInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *HotelOrderDetailInfoResponseBody
	GetCode() *string
	SetMessage(v string) *HotelOrderDetailInfoResponseBody
	GetMessage() *string
	SetModule(v *HotelOrderDetailInfoResponseBodyModule) *HotelOrderDetailInfoResponseBody
	GetModule() *HotelOrderDetailInfoResponseBodyModule
	SetRequestId(v string) *HotelOrderDetailInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *HotelOrderDetailInfoResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *HotelOrderDetailInfoResponseBody
	GetTraceId() *string
}

type HotelOrderDetailInfoResponseBody struct {
	// example:
	//
	// 0
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// succeed in handling request
	Message *string                                 `json:"message,omitempty" xml:"message,omitempty"`
	Module  *HotelOrderDetailInfoResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
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

func (s HotelOrderDetailInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s HotelOrderDetailInfoResponseBody) GoString() string {
	return s.String()
}

func (s *HotelOrderDetailInfoResponseBody) GetCode() *string {
	return s.Code
}

func (s *HotelOrderDetailInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *HotelOrderDetailInfoResponseBody) GetModule() *HotelOrderDetailInfoResponseBodyModule {
	return s.Module
}

func (s *HotelOrderDetailInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *HotelOrderDetailInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *HotelOrderDetailInfoResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *HotelOrderDetailInfoResponseBody) SetCode(v string) *HotelOrderDetailInfoResponseBody {
	s.Code = &v
	return s
}

func (s *HotelOrderDetailInfoResponseBody) SetMessage(v string) *HotelOrderDetailInfoResponseBody {
	s.Message = &v
	return s
}

func (s *HotelOrderDetailInfoResponseBody) SetModule(v *HotelOrderDetailInfoResponseBodyModule) *HotelOrderDetailInfoResponseBody {
	s.Module = v
	return s
}

func (s *HotelOrderDetailInfoResponseBody) SetRequestId(v string) *HotelOrderDetailInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *HotelOrderDetailInfoResponseBody) SetSuccess(v bool) *HotelOrderDetailInfoResponseBody {
	s.Success = &v
	return s
}

func (s *HotelOrderDetailInfoResponseBody) SetTraceId(v string) *HotelOrderDetailInfoResponseBody {
	s.TraceId = &v
	return s
}

func (s *HotelOrderDetailInfoResponseBody) Validate() error {
	return dara.Validate(s)
}

type HotelOrderDetailInfoResponseBodyModule struct {
	// example:
	//
	// 2022-01-29
	ActualCheckInTime *string `json:"actual_check_in_time,omitempty" xml:"actual_check_in_time,omitempty"`
	// example:
	//
	// 2022-01-29
	ActualCheckOutTime        *string                                                          `json:"actual_check_out_time,omitempty" xml:"actual_check_out_time,omitempty"`
	BtripHotelCancelPolicyDTO *HotelOrderDetailInfoResponseBodyModuleBtripHotelCancelPolicyDTO `json:"btrip_hotel_cancel_policy_d_t_o,omitempty" xml:"btrip_hotel_cancel_policy_d_t_o,omitempty" type:"Struct"`
	// example:
	//
	// 123
	BtripOrderId *string `json:"btrip_order_id,omitempty" xml:"btrip_order_id,omitempty"`
	// example:
	//
	// 100
	CancelFine *int64                                            `json:"cancel_fine,omitempty" xml:"cancel_fine,omitempty"`
	CancelInfo *HotelOrderDetailInfoResponseBodyModuleCancelInfo `json:"cancel_info,omitempty" xml:"cancel_info,omitempty" type:"Struct"`
	// example:
	//
	// 2022-05-15T22:27Z
	CheckIn *string `json:"check_in,omitempty" xml:"check_in,omitempty"`
	// example:
	//
	// 2022-05-15T22:27Z
	CheckOut *string `json:"check_out,omitempty" xml:"check_out,omitempty"`
	// example:
	//
	// 2000-00-00 00:00:00
	ConfirmOrderTime *string `json:"confirm_order_time,omitempty" xml:"confirm_order_time,omitempty"`
	ContractName     *string `json:"contract_name,omitempty" xml:"contract_name,omitempty"`
	// example:
	//
	// 12316261873
	ContractTel *string `json:"contract_tel,omitempty" xml:"contract_tel,omitempty"`
	// example:
	//
	// 2020-01-21
	CreateOrderTime *string `json:"create_order_time,omitempty" xml:"create_order_time,omitempty"`
	// example:
	//
	// 2000-00-00 00:00:00
	EarlyArrivalTime *string `json:"early_arrival_time,omitempty" xml:"early_arrival_time,omitempty"`
	// example:
	//
	// true
	EarlyDeparture *bool `json:"early_departure,omitempty" xml:"early_departure,omitempty"`
	// example:
	//
	// 2
	GuestCount              *int32                                                           `json:"guest_count,omitempty" xml:"guest_count,omitempty"`
	HotelDetailInfo         *HotelOrderDetailInfoResponseBodyModuleHotelDetailInfo           `json:"hotel_detail_info,omitempty" xml:"hotel_detail_info,omitempty" type:"Struct"`
	HotelSaleOrderRoomInfos []*HotelOrderDetailInfoResponseBodyModuleHotelSaleOrderRoomInfos `json:"hotel_sale_order_room_infos,omitempty" xml:"hotel_sale_order_room_infos,omitempty" type:"Repeated"`
	InvoiceInfo             *HotelOrderDetailInfoResponseBodyModuleInvoiceInfo               `json:"invoice_info,omitempty" xml:"invoice_info,omitempty" type:"Struct"`
	// example:
	//
	// 1289918
	ItemId *string `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// example:
	//
	// 2000-00-00 00:00:00
	LastArrivalTime  *string                                                   `json:"last_arrival_time,omitempty" xml:"last_arrival_time,omitempty"`
	OccupantInfoList []*HotelOrderDetailInfoResponseBodyModuleOccupantInfoList `json:"occupant_info_list,omitempty" xml:"occupant_info_list,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	OrderStatus     *int32  `json:"order_status,omitempty" xml:"order_status,omitempty"`
	OrderStatusDesc *string `json:"order_status_desc,omitempty" xml:"order_status_desc,omitempty"`
	OutConfirmCode  *string `json:"out_confirm_code,omitempty" xml:"out_confirm_code,omitempty"`
	// example:
	//
	// 2000-00-00 00:00:00
	PayTime *string `json:"pay_time,omitempty" xml:"pay_time,omitempty"`
	// example:
	//
	// 1
	ProductType *int32 `json:"product_type,omitempty" xml:"product_type,omitempty"`
	// example:
	//
	// 100328718728171
	PurchaseOrderId *string `json:"purchase_order_id,omitempty" xml:"purchase_order_id,omitempty"`
	// example:
	//
	// 100
	RefundPrice *int64 `json:"refund_price,omitempty" xml:"refund_price,omitempty"`
	// example:
	//
	// demo
	RefundReason *string `json:"refund_reason,omitempty" xml:"refund_reason,omitempty"`
	// example:
	//
	// 100
	RefundServiceFee       *int64                                                          `json:"refund_service_fee,omitempty" xml:"refund_service_fee,omitempty"`
	RoomNightPriceInfoList []*HotelOrderDetailInfoResponseBodyModuleRoomNightPriceInfoList `json:"room_night_price_info_list,omitempty" xml:"room_night_price_info_list,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	RoomNumber   *int32  `json:"room_number,omitempty" xml:"room_number,omitempty"`
	RoomTypeName *string `json:"room_type_name,omitempty" xml:"room_type_name,omitempty"`
	// example:
	//
	// 2088441675613762
	SellerId *string `json:"seller_id,omitempty" xml:"seller_id,omitempty"`
	// example:
	//
	// taobao
	SellerName *string `json:"seller_name,omitempty" xml:"seller_name,omitempty"`
	// example:
	//
	// 1.02
	ServiceFee *int64 `json:"service_fee,omitempty" xml:"service_fee,omitempty"`
	// example:
	//
	// 0
	SettleType *string `json:"settle_type,omitempty" xml:"settle_type,omitempty"`
	// example:
	//
	// dis123
	SupplierOrderId *string `json:"supplier_order_id,omitempty" xml:"supplier_order_id,omitempty"`
	// example:
	//
	// 100
	TotalPrice *int64 `json:"total_price,omitempty" xml:"total_price,omitempty"`
}

func (s HotelOrderDetailInfoResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s HotelOrderDetailInfoResponseBodyModule) GoString() string {
	return s.String()
}

func (s *HotelOrderDetailInfoResponseBodyModule) GetActualCheckInTime() *string {
	return s.ActualCheckInTime
}

func (s *HotelOrderDetailInfoResponseBodyModule) GetActualCheckOutTime() *string {
	return s.ActualCheckOutTime
}

func (s *HotelOrderDetailInfoResponseBodyModule) GetBtripHotelCancelPolicyDTO() *HotelOrderDetailInfoResponseBodyModuleBtripHotelCancelPolicyDTO {
	return s.BtripHotelCancelPolicyDTO
}

func (s *HotelOrderDetailInfoResponseBodyModule) GetBtripOrderId() *string {
	return s.BtripOrderId
}

func (s *HotelOrderDetailInfoResponseBodyModule) GetCancelFine() *int64 {
	return s.CancelFine
}

func (s *HotelOrderDetailInfoResponseBodyModule) GetCancelInfo() *HotelOrderDetailInfoResponseBodyModuleCancelInfo {
	return s.CancelInfo
}

func (s *HotelOrderDetailInfoResponseBodyModule) GetCheckIn() *string {
	return s.CheckIn
}

func (s *HotelOrderDetailInfoResponseBodyModule) GetCheckOut() *string {
	return s.CheckOut
}

func (s *HotelOrderDetailInfoResponseBodyModule) GetConfirmOrderTime() *string {
	return s.ConfirmOrderTime
}

func (s *HotelOrderDetailInfoResponseBodyModule) GetContractName() *string {
	return s.ContractName
}

func (s *HotelOrderDetailInfoResponseBodyModule) GetContractTel() *string {
	return s.ContractTel
}

func (s *HotelOrderDetailInfoResponseBodyModule) GetCreateOrderTime() *string {
	return s.CreateOrderTime
}

func (s *HotelOrderDetailInfoResponseBodyModule) GetEarlyArrivalTime() *string {
	return s.EarlyArrivalTime
}

func (s *HotelOrderDetailInfoResponseBodyModule) GetEarlyDeparture() *bool {
	return s.EarlyDeparture
}

func (s *HotelOrderDetailInfoResponseBodyModule) GetGuestCount() *int32 {
	return s.GuestCount
}

func (s *HotelOrderDetailInfoResponseBodyModule) GetHotelDetailInfo() *HotelOrderDetailInfoResponseBodyModuleHotelDetailInfo {
	return s.HotelDetailInfo
}

func (s *HotelOrderDetailInfoResponseBodyModule) GetHotelSaleOrderRoomInfos() []*HotelOrderDetailInfoResponseBodyModuleHotelSaleOrderRoomInfos {
	return s.HotelSaleOrderRoomInfos
}

func (s *HotelOrderDetailInfoResponseBodyModule) GetInvoiceInfo() *HotelOrderDetailInfoResponseBodyModuleInvoiceInfo {
	return s.InvoiceInfo
}

func (s *HotelOrderDetailInfoResponseBodyModule) GetItemId() *string {
	return s.ItemId
}

func (s *HotelOrderDetailInfoResponseBodyModule) GetLastArrivalTime() *string {
	return s.LastArrivalTime
}

func (s *HotelOrderDetailInfoResponseBodyModule) GetOccupantInfoList() []*HotelOrderDetailInfoResponseBodyModuleOccupantInfoList {
	return s.OccupantInfoList
}

func (s *HotelOrderDetailInfoResponseBodyModule) GetOrderStatus() *int32 {
	return s.OrderStatus
}

func (s *HotelOrderDetailInfoResponseBodyModule) GetOrderStatusDesc() *string {
	return s.OrderStatusDesc
}

func (s *HotelOrderDetailInfoResponseBodyModule) GetOutConfirmCode() *string {
	return s.OutConfirmCode
}

func (s *HotelOrderDetailInfoResponseBodyModule) GetPayTime() *string {
	return s.PayTime
}

func (s *HotelOrderDetailInfoResponseBodyModule) GetProductType() *int32 {
	return s.ProductType
}

func (s *HotelOrderDetailInfoResponseBodyModule) GetPurchaseOrderId() *string {
	return s.PurchaseOrderId
}

func (s *HotelOrderDetailInfoResponseBodyModule) GetRefundPrice() *int64 {
	return s.RefundPrice
}

func (s *HotelOrderDetailInfoResponseBodyModule) GetRefundReason() *string {
	return s.RefundReason
}

func (s *HotelOrderDetailInfoResponseBodyModule) GetRefundServiceFee() *int64 {
	return s.RefundServiceFee
}

func (s *HotelOrderDetailInfoResponseBodyModule) GetRoomNightPriceInfoList() []*HotelOrderDetailInfoResponseBodyModuleRoomNightPriceInfoList {
	return s.RoomNightPriceInfoList
}

func (s *HotelOrderDetailInfoResponseBodyModule) GetRoomNumber() *int32 {
	return s.RoomNumber
}

func (s *HotelOrderDetailInfoResponseBodyModule) GetRoomTypeName() *string {
	return s.RoomTypeName
}

func (s *HotelOrderDetailInfoResponseBodyModule) GetSellerId() *string {
	return s.SellerId
}

func (s *HotelOrderDetailInfoResponseBodyModule) GetSellerName() *string {
	return s.SellerName
}

func (s *HotelOrderDetailInfoResponseBodyModule) GetServiceFee() *int64 {
	return s.ServiceFee
}

func (s *HotelOrderDetailInfoResponseBodyModule) GetSettleType() *string {
	return s.SettleType
}

func (s *HotelOrderDetailInfoResponseBodyModule) GetSupplierOrderId() *string {
	return s.SupplierOrderId
}

func (s *HotelOrderDetailInfoResponseBodyModule) GetTotalPrice() *int64 {
	return s.TotalPrice
}

func (s *HotelOrderDetailInfoResponseBodyModule) SetActualCheckInTime(v string) *HotelOrderDetailInfoResponseBodyModule {
	s.ActualCheckInTime = &v
	return s
}

func (s *HotelOrderDetailInfoResponseBodyModule) SetActualCheckOutTime(v string) *HotelOrderDetailInfoResponseBodyModule {
	s.ActualCheckOutTime = &v
	return s
}

func (s *HotelOrderDetailInfoResponseBodyModule) SetBtripHotelCancelPolicyDTO(v *HotelOrderDetailInfoResponseBodyModuleBtripHotelCancelPolicyDTO) *HotelOrderDetailInfoResponseBodyModule {
	s.BtripHotelCancelPolicyDTO = v
	return s
}

func (s *HotelOrderDetailInfoResponseBodyModule) SetBtripOrderId(v string) *HotelOrderDetailInfoResponseBodyModule {
	s.BtripOrderId = &v
	return s
}

func (s *HotelOrderDetailInfoResponseBodyModule) SetCancelFine(v int64) *HotelOrderDetailInfoResponseBodyModule {
	s.CancelFine = &v
	return s
}

func (s *HotelOrderDetailInfoResponseBodyModule) SetCancelInfo(v *HotelOrderDetailInfoResponseBodyModuleCancelInfo) *HotelOrderDetailInfoResponseBodyModule {
	s.CancelInfo = v
	return s
}

func (s *HotelOrderDetailInfoResponseBodyModule) SetCheckIn(v string) *HotelOrderDetailInfoResponseBodyModule {
	s.CheckIn = &v
	return s
}

func (s *HotelOrderDetailInfoResponseBodyModule) SetCheckOut(v string) *HotelOrderDetailInfoResponseBodyModule {
	s.CheckOut = &v
	return s
}

func (s *HotelOrderDetailInfoResponseBodyModule) SetConfirmOrderTime(v string) *HotelOrderDetailInfoResponseBodyModule {
	s.ConfirmOrderTime = &v
	return s
}

func (s *HotelOrderDetailInfoResponseBodyModule) SetContractName(v string) *HotelOrderDetailInfoResponseBodyModule {
	s.ContractName = &v
	return s
}

func (s *HotelOrderDetailInfoResponseBodyModule) SetContractTel(v string) *HotelOrderDetailInfoResponseBodyModule {
	s.ContractTel = &v
	return s
}

func (s *HotelOrderDetailInfoResponseBodyModule) SetCreateOrderTime(v string) *HotelOrderDetailInfoResponseBodyModule {
	s.CreateOrderTime = &v
	return s
}

func (s *HotelOrderDetailInfoResponseBodyModule) SetEarlyArrivalTime(v string) *HotelOrderDetailInfoResponseBodyModule {
	s.EarlyArrivalTime = &v
	return s
}

func (s *HotelOrderDetailInfoResponseBodyModule) SetEarlyDeparture(v bool) *HotelOrderDetailInfoResponseBodyModule {
	s.EarlyDeparture = &v
	return s
}

func (s *HotelOrderDetailInfoResponseBodyModule) SetGuestCount(v int32) *HotelOrderDetailInfoResponseBodyModule {
	s.GuestCount = &v
	return s
}

func (s *HotelOrderDetailInfoResponseBodyModule) SetHotelDetailInfo(v *HotelOrderDetailInfoResponseBodyModuleHotelDetailInfo) *HotelOrderDetailInfoResponseBodyModule {
	s.HotelDetailInfo = v
	return s
}

func (s *HotelOrderDetailInfoResponseBodyModule) SetHotelSaleOrderRoomInfos(v []*HotelOrderDetailInfoResponseBodyModuleHotelSaleOrderRoomInfos) *HotelOrderDetailInfoResponseBodyModule {
	s.HotelSaleOrderRoomInfos = v
	return s
}

func (s *HotelOrderDetailInfoResponseBodyModule) SetInvoiceInfo(v *HotelOrderDetailInfoResponseBodyModuleInvoiceInfo) *HotelOrderDetailInfoResponseBodyModule {
	s.InvoiceInfo = v
	return s
}

func (s *HotelOrderDetailInfoResponseBodyModule) SetItemId(v string) *HotelOrderDetailInfoResponseBodyModule {
	s.ItemId = &v
	return s
}

func (s *HotelOrderDetailInfoResponseBodyModule) SetLastArrivalTime(v string) *HotelOrderDetailInfoResponseBodyModule {
	s.LastArrivalTime = &v
	return s
}

func (s *HotelOrderDetailInfoResponseBodyModule) SetOccupantInfoList(v []*HotelOrderDetailInfoResponseBodyModuleOccupantInfoList) *HotelOrderDetailInfoResponseBodyModule {
	s.OccupantInfoList = v
	return s
}

func (s *HotelOrderDetailInfoResponseBodyModule) SetOrderStatus(v int32) *HotelOrderDetailInfoResponseBodyModule {
	s.OrderStatus = &v
	return s
}

func (s *HotelOrderDetailInfoResponseBodyModule) SetOrderStatusDesc(v string) *HotelOrderDetailInfoResponseBodyModule {
	s.OrderStatusDesc = &v
	return s
}

func (s *HotelOrderDetailInfoResponseBodyModule) SetOutConfirmCode(v string) *HotelOrderDetailInfoResponseBodyModule {
	s.OutConfirmCode = &v
	return s
}

func (s *HotelOrderDetailInfoResponseBodyModule) SetPayTime(v string) *HotelOrderDetailInfoResponseBodyModule {
	s.PayTime = &v
	return s
}

func (s *HotelOrderDetailInfoResponseBodyModule) SetProductType(v int32) *HotelOrderDetailInfoResponseBodyModule {
	s.ProductType = &v
	return s
}

func (s *HotelOrderDetailInfoResponseBodyModule) SetPurchaseOrderId(v string) *HotelOrderDetailInfoResponseBodyModule {
	s.PurchaseOrderId = &v
	return s
}

func (s *HotelOrderDetailInfoResponseBodyModule) SetRefundPrice(v int64) *HotelOrderDetailInfoResponseBodyModule {
	s.RefundPrice = &v
	return s
}

func (s *HotelOrderDetailInfoResponseBodyModule) SetRefundReason(v string) *HotelOrderDetailInfoResponseBodyModule {
	s.RefundReason = &v
	return s
}

func (s *HotelOrderDetailInfoResponseBodyModule) SetRefundServiceFee(v int64) *HotelOrderDetailInfoResponseBodyModule {
	s.RefundServiceFee = &v
	return s
}

func (s *HotelOrderDetailInfoResponseBodyModule) SetRoomNightPriceInfoList(v []*HotelOrderDetailInfoResponseBodyModuleRoomNightPriceInfoList) *HotelOrderDetailInfoResponseBodyModule {
	s.RoomNightPriceInfoList = v
	return s
}

func (s *HotelOrderDetailInfoResponseBodyModule) SetRoomNumber(v int32) *HotelOrderDetailInfoResponseBodyModule {
	s.RoomNumber = &v
	return s
}

func (s *HotelOrderDetailInfoResponseBodyModule) SetRoomTypeName(v string) *HotelOrderDetailInfoResponseBodyModule {
	s.RoomTypeName = &v
	return s
}

func (s *HotelOrderDetailInfoResponseBodyModule) SetSellerId(v string) *HotelOrderDetailInfoResponseBodyModule {
	s.SellerId = &v
	return s
}

func (s *HotelOrderDetailInfoResponseBodyModule) SetSellerName(v string) *HotelOrderDetailInfoResponseBodyModule {
	s.SellerName = &v
	return s
}

func (s *HotelOrderDetailInfoResponseBodyModule) SetServiceFee(v int64) *HotelOrderDetailInfoResponseBodyModule {
	s.ServiceFee = &v
	return s
}

func (s *HotelOrderDetailInfoResponseBodyModule) SetSettleType(v string) *HotelOrderDetailInfoResponseBodyModule {
	s.SettleType = &v
	return s
}

func (s *HotelOrderDetailInfoResponseBodyModule) SetSupplierOrderId(v string) *HotelOrderDetailInfoResponseBodyModule {
	s.SupplierOrderId = &v
	return s
}

func (s *HotelOrderDetailInfoResponseBodyModule) SetTotalPrice(v int64) *HotelOrderDetailInfoResponseBodyModule {
	s.TotalPrice = &v
	return s
}

func (s *HotelOrderDetailInfoResponseBodyModule) Validate() error {
	return dara.Validate(s)
}

type HotelOrderDetailInfoResponseBodyModuleBtripHotelCancelPolicyDTO struct {
	BtripHotelCancelPolicyInfoDTOList []*HotelOrderDetailInfoResponseBodyModuleBtripHotelCancelPolicyDTOBtripHotelCancelPolicyInfoDTOList `json:"btrip_hotel_cancel_policy_info_d_t_o_list,omitempty" xml:"btrip_hotel_cancel_policy_info_d_t_o_list,omitempty" type:"Repeated"`
	CancelPolicyType                  *int32                                                                                              `json:"cancel_policy_type,omitempty" xml:"cancel_policy_type,omitempty"`
}

func (s HotelOrderDetailInfoResponseBodyModuleBtripHotelCancelPolicyDTO) String() string {
	return dara.Prettify(s)
}

func (s HotelOrderDetailInfoResponseBodyModuleBtripHotelCancelPolicyDTO) GoString() string {
	return s.String()
}

func (s *HotelOrderDetailInfoResponseBodyModuleBtripHotelCancelPolicyDTO) GetBtripHotelCancelPolicyInfoDTOList() []*HotelOrderDetailInfoResponseBodyModuleBtripHotelCancelPolicyDTOBtripHotelCancelPolicyInfoDTOList {
	return s.BtripHotelCancelPolicyInfoDTOList
}

func (s *HotelOrderDetailInfoResponseBodyModuleBtripHotelCancelPolicyDTO) GetCancelPolicyType() *int32 {
	return s.CancelPolicyType
}

func (s *HotelOrderDetailInfoResponseBodyModuleBtripHotelCancelPolicyDTO) SetBtripHotelCancelPolicyInfoDTOList(v []*HotelOrderDetailInfoResponseBodyModuleBtripHotelCancelPolicyDTOBtripHotelCancelPolicyInfoDTOList) *HotelOrderDetailInfoResponseBodyModuleBtripHotelCancelPolicyDTO {
	s.BtripHotelCancelPolicyInfoDTOList = v
	return s
}

func (s *HotelOrderDetailInfoResponseBodyModuleBtripHotelCancelPolicyDTO) SetCancelPolicyType(v int32) *HotelOrderDetailInfoResponseBodyModuleBtripHotelCancelPolicyDTO {
	s.CancelPolicyType = &v
	return s
}

func (s *HotelOrderDetailInfoResponseBodyModuleBtripHotelCancelPolicyDTO) Validate() error {
	return dara.Validate(s)
}

type HotelOrderDetailInfoResponseBodyModuleBtripHotelCancelPolicyDTOBtripHotelCancelPolicyInfoDTOList struct {
	// example:
	//
	// 1
	Hour *int64 `json:"hour,omitempty" xml:"hour,omitempty"`
	// example:
	//
	// 1
	Value *int64 `json:"value,omitempty" xml:"value,omitempty"`
}

func (s HotelOrderDetailInfoResponseBodyModuleBtripHotelCancelPolicyDTOBtripHotelCancelPolicyInfoDTOList) String() string {
	return dara.Prettify(s)
}

func (s HotelOrderDetailInfoResponseBodyModuleBtripHotelCancelPolicyDTOBtripHotelCancelPolicyInfoDTOList) GoString() string {
	return s.String()
}

func (s *HotelOrderDetailInfoResponseBodyModuleBtripHotelCancelPolicyDTOBtripHotelCancelPolicyInfoDTOList) GetHour() *int64 {
	return s.Hour
}

func (s *HotelOrderDetailInfoResponseBodyModuleBtripHotelCancelPolicyDTOBtripHotelCancelPolicyInfoDTOList) GetValue() *int64 {
	return s.Value
}

func (s *HotelOrderDetailInfoResponseBodyModuleBtripHotelCancelPolicyDTOBtripHotelCancelPolicyInfoDTOList) SetHour(v int64) *HotelOrderDetailInfoResponseBodyModuleBtripHotelCancelPolicyDTOBtripHotelCancelPolicyInfoDTOList {
	s.Hour = &v
	return s
}

func (s *HotelOrderDetailInfoResponseBodyModuleBtripHotelCancelPolicyDTOBtripHotelCancelPolicyInfoDTOList) SetValue(v int64) *HotelOrderDetailInfoResponseBodyModuleBtripHotelCancelPolicyDTOBtripHotelCancelPolicyInfoDTOList {
	s.Value = &v
	return s
}

func (s *HotelOrderDetailInfoResponseBodyModuleBtripHotelCancelPolicyDTOBtripHotelCancelPolicyInfoDTOList) Validate() error {
	return dara.Validate(s)
}

type HotelOrderDetailInfoResponseBodyModuleCancelInfo struct {
	// example:
	//
	// 2022-01-29
	CancelEndTime *string `json:"cancel_end_time,omitempty" xml:"cancel_end_time,omitempty"`
	// example:
	//
	// 2022-01-29
	CancelStartTime *string `json:"cancel_start_time,omitempty" xml:"cancel_start_time,omitempty"`
}

func (s HotelOrderDetailInfoResponseBodyModuleCancelInfo) String() string {
	return dara.Prettify(s)
}

func (s HotelOrderDetailInfoResponseBodyModuleCancelInfo) GoString() string {
	return s.String()
}

func (s *HotelOrderDetailInfoResponseBodyModuleCancelInfo) GetCancelEndTime() *string {
	return s.CancelEndTime
}

func (s *HotelOrderDetailInfoResponseBodyModuleCancelInfo) GetCancelStartTime() *string {
	return s.CancelStartTime
}

func (s *HotelOrderDetailInfoResponseBodyModuleCancelInfo) SetCancelEndTime(v string) *HotelOrderDetailInfoResponseBodyModuleCancelInfo {
	s.CancelEndTime = &v
	return s
}

func (s *HotelOrderDetailInfoResponseBodyModuleCancelInfo) SetCancelStartTime(v string) *HotelOrderDetailInfoResponseBodyModuleCancelInfo {
	s.CancelStartTime = &v
	return s
}

func (s *HotelOrderDetailInfoResponseBodyModuleCancelInfo) Validate() error {
	return dara.Validate(s)
}

type HotelOrderDetailInfoResponseBodyModuleHotelDetailInfo struct {
	Address   *string `json:"address,omitempty" xml:"address,omitempty"`
	CityName  *string `json:"city_name,omitempty" xml:"city_name,omitempty"`
	HotelName *string `json:"hotel_name,omitempty" xml:"hotel_name,omitempty"`
	// example:
	//
	// 0571-87217
	HotelTel *string `json:"hotel_tel,omitempty" xml:"hotel_tel,omitempty"`
	// example:
	//
	// 2198781
	Shid *int64 `json:"shid,omitempty" xml:"shid,omitempty"`
}

func (s HotelOrderDetailInfoResponseBodyModuleHotelDetailInfo) String() string {
	return dara.Prettify(s)
}

func (s HotelOrderDetailInfoResponseBodyModuleHotelDetailInfo) GoString() string {
	return s.String()
}

func (s *HotelOrderDetailInfoResponseBodyModuleHotelDetailInfo) GetAddress() *string {
	return s.Address
}

func (s *HotelOrderDetailInfoResponseBodyModuleHotelDetailInfo) GetCityName() *string {
	return s.CityName
}

func (s *HotelOrderDetailInfoResponseBodyModuleHotelDetailInfo) GetHotelName() *string {
	return s.HotelName
}

func (s *HotelOrderDetailInfoResponseBodyModuleHotelDetailInfo) GetHotelTel() *string {
	return s.HotelTel
}

func (s *HotelOrderDetailInfoResponseBodyModuleHotelDetailInfo) GetShid() *int64 {
	return s.Shid
}

func (s *HotelOrderDetailInfoResponseBodyModuleHotelDetailInfo) SetAddress(v string) *HotelOrderDetailInfoResponseBodyModuleHotelDetailInfo {
	s.Address = &v
	return s
}

func (s *HotelOrderDetailInfoResponseBodyModuleHotelDetailInfo) SetCityName(v string) *HotelOrderDetailInfoResponseBodyModuleHotelDetailInfo {
	s.CityName = &v
	return s
}

func (s *HotelOrderDetailInfoResponseBodyModuleHotelDetailInfo) SetHotelName(v string) *HotelOrderDetailInfoResponseBodyModuleHotelDetailInfo {
	s.HotelName = &v
	return s
}

func (s *HotelOrderDetailInfoResponseBodyModuleHotelDetailInfo) SetHotelTel(v string) *HotelOrderDetailInfoResponseBodyModuleHotelDetailInfo {
	s.HotelTel = &v
	return s
}

func (s *HotelOrderDetailInfoResponseBodyModuleHotelDetailInfo) SetShid(v int64) *HotelOrderDetailInfoResponseBodyModuleHotelDetailInfo {
	s.Shid = &v
	return s
}

func (s *HotelOrderDetailInfoResponseBodyModuleHotelDetailInfo) Validate() error {
	return dara.Validate(s)
}

type HotelOrderDetailInfoResponseBodyModuleHotelSaleOrderRoomInfos struct {
	CheckinDate      *string `json:"checkin_date,omitempty" xml:"checkin_date,omitempty"`
	CheckoutDate     *string `json:"checkout_date,omitempty" xml:"checkout_date,omitempty"`
	PenalSum         *int64  `json:"penal_sum,omitempty" xml:"penal_sum,omitempty"`
	RealCheckoutDate *string `json:"real_checkout_date,omitempty" xml:"real_checkout_date,omitempty"`
	RefundStatus     *int32  `json:"refund_status,omitempty" xml:"refund_status,omitempty"`
	RoomNo           *int32  `json:"room_no,omitempty" xml:"room_no,omitempty"`
	RoomPrice        *int64  `json:"room_price,omitempty" xml:"room_price,omitempty"`
	RoomRefundPrice  *int64  `json:"room_refund_price,omitempty" xml:"room_refund_price,omitempty"`
	TravelerId       *string `json:"traveler_id,omitempty" xml:"traveler_id,omitempty"`
	TravelerName     *string `json:"traveler_name,omitempty" xml:"traveler_name,omitempty"`
}

func (s HotelOrderDetailInfoResponseBodyModuleHotelSaleOrderRoomInfos) String() string {
	return dara.Prettify(s)
}

func (s HotelOrderDetailInfoResponseBodyModuleHotelSaleOrderRoomInfos) GoString() string {
	return s.String()
}

func (s *HotelOrderDetailInfoResponseBodyModuleHotelSaleOrderRoomInfos) GetCheckinDate() *string {
	return s.CheckinDate
}

func (s *HotelOrderDetailInfoResponseBodyModuleHotelSaleOrderRoomInfos) GetCheckoutDate() *string {
	return s.CheckoutDate
}

func (s *HotelOrderDetailInfoResponseBodyModuleHotelSaleOrderRoomInfos) GetPenalSum() *int64 {
	return s.PenalSum
}

func (s *HotelOrderDetailInfoResponseBodyModuleHotelSaleOrderRoomInfos) GetRealCheckoutDate() *string {
	return s.RealCheckoutDate
}

func (s *HotelOrderDetailInfoResponseBodyModuleHotelSaleOrderRoomInfos) GetRefundStatus() *int32 {
	return s.RefundStatus
}

func (s *HotelOrderDetailInfoResponseBodyModuleHotelSaleOrderRoomInfos) GetRoomNo() *int32 {
	return s.RoomNo
}

func (s *HotelOrderDetailInfoResponseBodyModuleHotelSaleOrderRoomInfos) GetRoomPrice() *int64 {
	return s.RoomPrice
}

func (s *HotelOrderDetailInfoResponseBodyModuleHotelSaleOrderRoomInfos) GetRoomRefundPrice() *int64 {
	return s.RoomRefundPrice
}

func (s *HotelOrderDetailInfoResponseBodyModuleHotelSaleOrderRoomInfos) GetTravelerId() *string {
	return s.TravelerId
}

func (s *HotelOrderDetailInfoResponseBodyModuleHotelSaleOrderRoomInfos) GetTravelerName() *string {
	return s.TravelerName
}

func (s *HotelOrderDetailInfoResponseBodyModuleHotelSaleOrderRoomInfos) SetCheckinDate(v string) *HotelOrderDetailInfoResponseBodyModuleHotelSaleOrderRoomInfos {
	s.CheckinDate = &v
	return s
}

func (s *HotelOrderDetailInfoResponseBodyModuleHotelSaleOrderRoomInfos) SetCheckoutDate(v string) *HotelOrderDetailInfoResponseBodyModuleHotelSaleOrderRoomInfos {
	s.CheckoutDate = &v
	return s
}

func (s *HotelOrderDetailInfoResponseBodyModuleHotelSaleOrderRoomInfos) SetPenalSum(v int64) *HotelOrderDetailInfoResponseBodyModuleHotelSaleOrderRoomInfos {
	s.PenalSum = &v
	return s
}

func (s *HotelOrderDetailInfoResponseBodyModuleHotelSaleOrderRoomInfos) SetRealCheckoutDate(v string) *HotelOrderDetailInfoResponseBodyModuleHotelSaleOrderRoomInfos {
	s.RealCheckoutDate = &v
	return s
}

func (s *HotelOrderDetailInfoResponseBodyModuleHotelSaleOrderRoomInfos) SetRefundStatus(v int32) *HotelOrderDetailInfoResponseBodyModuleHotelSaleOrderRoomInfos {
	s.RefundStatus = &v
	return s
}

func (s *HotelOrderDetailInfoResponseBodyModuleHotelSaleOrderRoomInfos) SetRoomNo(v int32) *HotelOrderDetailInfoResponseBodyModuleHotelSaleOrderRoomInfos {
	s.RoomNo = &v
	return s
}

func (s *HotelOrderDetailInfoResponseBodyModuleHotelSaleOrderRoomInfos) SetRoomPrice(v int64) *HotelOrderDetailInfoResponseBodyModuleHotelSaleOrderRoomInfos {
	s.RoomPrice = &v
	return s
}

func (s *HotelOrderDetailInfoResponseBodyModuleHotelSaleOrderRoomInfos) SetRoomRefundPrice(v int64) *HotelOrderDetailInfoResponseBodyModuleHotelSaleOrderRoomInfos {
	s.RoomRefundPrice = &v
	return s
}

func (s *HotelOrderDetailInfoResponseBodyModuleHotelSaleOrderRoomInfos) SetTravelerId(v string) *HotelOrderDetailInfoResponseBodyModuleHotelSaleOrderRoomInfos {
	s.TravelerId = &v
	return s
}

func (s *HotelOrderDetailInfoResponseBodyModuleHotelSaleOrderRoomInfos) SetTravelerName(v string) *HotelOrderDetailInfoResponseBodyModuleHotelSaleOrderRoomInfos {
	s.TravelerName = &v
	return s
}

func (s *HotelOrderDetailInfoResponseBodyModuleHotelSaleOrderRoomInfos) Validate() error {
	return dara.Validate(s)
}

type HotelOrderDetailInfoResponseBodyModuleInvoiceInfo struct {
	// example:
	//
	// 100
	BillingMoney *int64 `json:"billing_money,omitempty" xml:"billing_money,omitempty"`
	// example:
	//
	// demo
	BuyerAdd *string `json:"buyer_add,omitempty" xml:"buyer_add,omitempty"`
	// example:
	//
	// demo
	BuyerBankAcc *string `json:"buyer_bank_acc,omitempty" xml:"buyer_bank_acc,omitempty"`
	// example:
	//
	// demo
	BuyerBankAdd *string `json:"buyer_bank_add,omitempty" xml:"buyer_bank_add,omitempty"`
	// example:
	//
	// 0571-82321777
	BuyerPhone *string `json:"buyer_phone,omitempty" xml:"buyer_phone,omitempty"`
	// example:
	//
	// 1
	BuyerTaxNum      *string `json:"buyer_tax_num,omitempty" xml:"buyer_tax_num,omitempty"`
	DeliveryAddress  *string `json:"delivery_address,omitempty" xml:"delivery_address,omitempty"`
	DeliveryArea     *string `json:"delivery_area,omitempty" xml:"delivery_area,omitempty"`
	DeliveryCity     *string `json:"delivery_city,omitempty" xml:"delivery_city,omitempty"`
	DeliveryProvince *string `json:"delivery_province,omitempty" xml:"delivery_province,omitempty"`
	DeliveryStreet   *string `json:"delivery_street,omitempty" xml:"delivery_street,omitempty"`
	// example:
	//
	// demo
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
	// example:
	//
	// 1
	InvoiceMaterial *int32 `json:"invoice_material,omitempty" xml:"invoice_material,omitempty"`
	// example:
	//
	// demo
	InvoiceTitle *string `json:"invoice_title,omitempty" xml:"invoice_title,omitempty"`
	// example:
	//
	// 1
	InvoiceType *int32 `json:"invoice_type,omitempty" xml:"invoice_type,omitempty"`
	// example:
	//
	// 100
	Postage      *int64  `json:"postage,omitempty" xml:"postage,omitempty"`
	ReceiverName *string `json:"receiver_name,omitempty" xml:"receiver_name,omitempty"`
	// example:
	//
	// 0571-82321777
	ReceiverPhone *string `json:"receiver_phone,omitempty" xml:"receiver_phone,omitempty"`
	// example:
	//
	// tmf closeCloneTask
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
}

func (s HotelOrderDetailInfoResponseBodyModuleInvoiceInfo) String() string {
	return dara.Prettify(s)
}

func (s HotelOrderDetailInfoResponseBodyModuleInvoiceInfo) GoString() string {
	return s.String()
}

func (s *HotelOrderDetailInfoResponseBodyModuleInvoiceInfo) GetBillingMoney() *int64 {
	return s.BillingMoney
}

func (s *HotelOrderDetailInfoResponseBodyModuleInvoiceInfo) GetBuyerAdd() *string {
	return s.BuyerAdd
}

func (s *HotelOrderDetailInfoResponseBodyModuleInvoiceInfo) GetBuyerBankAcc() *string {
	return s.BuyerBankAcc
}

func (s *HotelOrderDetailInfoResponseBodyModuleInvoiceInfo) GetBuyerBankAdd() *string {
	return s.BuyerBankAdd
}

func (s *HotelOrderDetailInfoResponseBodyModuleInvoiceInfo) GetBuyerPhone() *string {
	return s.BuyerPhone
}

func (s *HotelOrderDetailInfoResponseBodyModuleInvoiceInfo) GetBuyerTaxNum() *string {
	return s.BuyerTaxNum
}

func (s *HotelOrderDetailInfoResponseBodyModuleInvoiceInfo) GetDeliveryAddress() *string {
	return s.DeliveryAddress
}

func (s *HotelOrderDetailInfoResponseBodyModuleInvoiceInfo) GetDeliveryArea() *string {
	return s.DeliveryArea
}

func (s *HotelOrderDetailInfoResponseBodyModuleInvoiceInfo) GetDeliveryCity() *string {
	return s.DeliveryCity
}

func (s *HotelOrderDetailInfoResponseBodyModuleInvoiceInfo) GetDeliveryProvince() *string {
	return s.DeliveryProvince
}

func (s *HotelOrderDetailInfoResponseBodyModuleInvoiceInfo) GetDeliveryStreet() *string {
	return s.DeliveryStreet
}

func (s *HotelOrderDetailInfoResponseBodyModuleInvoiceInfo) GetEmail() *string {
	return s.Email
}

func (s *HotelOrderDetailInfoResponseBodyModuleInvoiceInfo) GetInvoiceMaterial() *int32 {
	return s.InvoiceMaterial
}

func (s *HotelOrderDetailInfoResponseBodyModuleInvoiceInfo) GetInvoiceTitle() *string {
	return s.InvoiceTitle
}

func (s *HotelOrderDetailInfoResponseBodyModuleInvoiceInfo) GetInvoiceType() *int32 {
	return s.InvoiceType
}

func (s *HotelOrderDetailInfoResponseBodyModuleInvoiceInfo) GetPostage() *int64 {
	return s.Postage
}

func (s *HotelOrderDetailInfoResponseBodyModuleInvoiceInfo) GetReceiverName() *string {
	return s.ReceiverName
}

func (s *HotelOrderDetailInfoResponseBodyModuleInvoiceInfo) GetReceiverPhone() *string {
	return s.ReceiverPhone
}

func (s *HotelOrderDetailInfoResponseBodyModuleInvoiceInfo) GetRemark() *string {
	return s.Remark
}

func (s *HotelOrderDetailInfoResponseBodyModuleInvoiceInfo) SetBillingMoney(v int64) *HotelOrderDetailInfoResponseBodyModuleInvoiceInfo {
	s.BillingMoney = &v
	return s
}

func (s *HotelOrderDetailInfoResponseBodyModuleInvoiceInfo) SetBuyerAdd(v string) *HotelOrderDetailInfoResponseBodyModuleInvoiceInfo {
	s.BuyerAdd = &v
	return s
}

func (s *HotelOrderDetailInfoResponseBodyModuleInvoiceInfo) SetBuyerBankAcc(v string) *HotelOrderDetailInfoResponseBodyModuleInvoiceInfo {
	s.BuyerBankAcc = &v
	return s
}

func (s *HotelOrderDetailInfoResponseBodyModuleInvoiceInfo) SetBuyerBankAdd(v string) *HotelOrderDetailInfoResponseBodyModuleInvoiceInfo {
	s.BuyerBankAdd = &v
	return s
}

func (s *HotelOrderDetailInfoResponseBodyModuleInvoiceInfo) SetBuyerPhone(v string) *HotelOrderDetailInfoResponseBodyModuleInvoiceInfo {
	s.BuyerPhone = &v
	return s
}

func (s *HotelOrderDetailInfoResponseBodyModuleInvoiceInfo) SetBuyerTaxNum(v string) *HotelOrderDetailInfoResponseBodyModuleInvoiceInfo {
	s.BuyerTaxNum = &v
	return s
}

func (s *HotelOrderDetailInfoResponseBodyModuleInvoiceInfo) SetDeliveryAddress(v string) *HotelOrderDetailInfoResponseBodyModuleInvoiceInfo {
	s.DeliveryAddress = &v
	return s
}

func (s *HotelOrderDetailInfoResponseBodyModuleInvoiceInfo) SetDeliveryArea(v string) *HotelOrderDetailInfoResponseBodyModuleInvoiceInfo {
	s.DeliveryArea = &v
	return s
}

func (s *HotelOrderDetailInfoResponseBodyModuleInvoiceInfo) SetDeliveryCity(v string) *HotelOrderDetailInfoResponseBodyModuleInvoiceInfo {
	s.DeliveryCity = &v
	return s
}

func (s *HotelOrderDetailInfoResponseBodyModuleInvoiceInfo) SetDeliveryProvince(v string) *HotelOrderDetailInfoResponseBodyModuleInvoiceInfo {
	s.DeliveryProvince = &v
	return s
}

func (s *HotelOrderDetailInfoResponseBodyModuleInvoiceInfo) SetDeliveryStreet(v string) *HotelOrderDetailInfoResponseBodyModuleInvoiceInfo {
	s.DeliveryStreet = &v
	return s
}

func (s *HotelOrderDetailInfoResponseBodyModuleInvoiceInfo) SetEmail(v string) *HotelOrderDetailInfoResponseBodyModuleInvoiceInfo {
	s.Email = &v
	return s
}

func (s *HotelOrderDetailInfoResponseBodyModuleInvoiceInfo) SetInvoiceMaterial(v int32) *HotelOrderDetailInfoResponseBodyModuleInvoiceInfo {
	s.InvoiceMaterial = &v
	return s
}

func (s *HotelOrderDetailInfoResponseBodyModuleInvoiceInfo) SetInvoiceTitle(v string) *HotelOrderDetailInfoResponseBodyModuleInvoiceInfo {
	s.InvoiceTitle = &v
	return s
}

func (s *HotelOrderDetailInfoResponseBodyModuleInvoiceInfo) SetInvoiceType(v int32) *HotelOrderDetailInfoResponseBodyModuleInvoiceInfo {
	s.InvoiceType = &v
	return s
}

func (s *HotelOrderDetailInfoResponseBodyModuleInvoiceInfo) SetPostage(v int64) *HotelOrderDetailInfoResponseBodyModuleInvoiceInfo {
	s.Postage = &v
	return s
}

func (s *HotelOrderDetailInfoResponseBodyModuleInvoiceInfo) SetReceiverName(v string) *HotelOrderDetailInfoResponseBodyModuleInvoiceInfo {
	s.ReceiverName = &v
	return s
}

func (s *HotelOrderDetailInfoResponseBodyModuleInvoiceInfo) SetReceiverPhone(v string) *HotelOrderDetailInfoResponseBodyModuleInvoiceInfo {
	s.ReceiverPhone = &v
	return s
}

func (s *HotelOrderDetailInfoResponseBodyModuleInvoiceInfo) SetRemark(v string) *HotelOrderDetailInfoResponseBodyModuleInvoiceInfo {
	s.Remark = &v
	return s
}

func (s *HotelOrderDetailInfoResponseBodyModuleInvoiceInfo) Validate() error {
	return dara.Validate(s)
}

type HotelOrderDetailInfoResponseBodyModuleOccupantInfoList struct {
	// example:
	//
	// 1235615612424222
	CardNo *string `json:"card_no,omitempty" xml:"card_no,omitempty"`
	// example:
	//
	// 1
	CardType           *int32                                                                      `json:"card_type,omitempty" xml:"card_type,omitempty"`
	CostCenterInfoList []*HotelOrderDetailInfoResponseBodyModuleOccupantInfoListCostCenterInfoList `json:"cost_center_info_list,omitempty" xml:"cost_center_info_list,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	CustomerType *int32 `json:"customer_type,omitempty" xml:"customer_type,omitempty"`
	// example:
	//
	// 123112
	DepartmentId *string `json:"department_id,omitempty" xml:"department_id,omitempty"`
	// example:
	//
	// demo
	DepartmentName *string `json:"department_name,omitempty" xml:"department_name,omitempty"`
	// example:
	//
	// demo
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
	// example:
	//
	// 1
	EmployeeType *int32 `json:"employee_type,omitempty" xml:"employee_type,omitempty"`
	// example:
	//
	// zhangsan
	FirstName *string `json:"first_name,omitempty" xml:"first_name,omitempty"`
	// example:
	//
	// true
	IsBooker *bool `json:"is_booker,omitempty" xml:"is_booker,omitempty"`
	// example:
	//
	// zhang
	LastName *string `json:"last_name,omitempty" xml:"last_name,omitempty"`
	Name     *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 0571-872237
	Phone  *string `json:"phone,omitempty" xml:"phone,omitempty"`
	RoomNo *int32  `json:"room_no,omitempty" xml:"room_no,omitempty"`
	// example:
	//
	// true
	Selected *bool `json:"selected,omitempty" xml:"selected,omitempty"`
	// example:
	//
	// 12817218
	StaffNo *string `json:"staff_no,omitempty" xml:"staff_no,omitempty"`
	// example:
	//
	// 1
	UserType *int32 `json:"user_type,omitempty" xml:"user_type,omitempty"`
}

func (s HotelOrderDetailInfoResponseBodyModuleOccupantInfoList) String() string {
	return dara.Prettify(s)
}

func (s HotelOrderDetailInfoResponseBodyModuleOccupantInfoList) GoString() string {
	return s.String()
}

func (s *HotelOrderDetailInfoResponseBodyModuleOccupantInfoList) GetCardNo() *string {
	return s.CardNo
}

func (s *HotelOrderDetailInfoResponseBodyModuleOccupantInfoList) GetCardType() *int32 {
	return s.CardType
}

func (s *HotelOrderDetailInfoResponseBodyModuleOccupantInfoList) GetCostCenterInfoList() []*HotelOrderDetailInfoResponseBodyModuleOccupantInfoListCostCenterInfoList {
	return s.CostCenterInfoList
}

func (s *HotelOrderDetailInfoResponseBodyModuleOccupantInfoList) GetCustomerType() *int32 {
	return s.CustomerType
}

func (s *HotelOrderDetailInfoResponseBodyModuleOccupantInfoList) GetDepartmentId() *string {
	return s.DepartmentId
}

func (s *HotelOrderDetailInfoResponseBodyModuleOccupantInfoList) GetDepartmentName() *string {
	return s.DepartmentName
}

func (s *HotelOrderDetailInfoResponseBodyModuleOccupantInfoList) GetEmail() *string {
	return s.Email
}

func (s *HotelOrderDetailInfoResponseBodyModuleOccupantInfoList) GetEmployeeType() *int32 {
	return s.EmployeeType
}

func (s *HotelOrderDetailInfoResponseBodyModuleOccupantInfoList) GetFirstName() *string {
	return s.FirstName
}

func (s *HotelOrderDetailInfoResponseBodyModuleOccupantInfoList) GetIsBooker() *bool {
	return s.IsBooker
}

func (s *HotelOrderDetailInfoResponseBodyModuleOccupantInfoList) GetLastName() *string {
	return s.LastName
}

func (s *HotelOrderDetailInfoResponseBodyModuleOccupantInfoList) GetName() *string {
	return s.Name
}

func (s *HotelOrderDetailInfoResponseBodyModuleOccupantInfoList) GetPhone() *string {
	return s.Phone
}

func (s *HotelOrderDetailInfoResponseBodyModuleOccupantInfoList) GetRoomNo() *int32 {
	return s.RoomNo
}

func (s *HotelOrderDetailInfoResponseBodyModuleOccupantInfoList) GetSelected() *bool {
	return s.Selected
}

func (s *HotelOrderDetailInfoResponseBodyModuleOccupantInfoList) GetStaffNo() *string {
	return s.StaffNo
}

func (s *HotelOrderDetailInfoResponseBodyModuleOccupantInfoList) GetUserType() *int32 {
	return s.UserType
}

func (s *HotelOrderDetailInfoResponseBodyModuleOccupantInfoList) SetCardNo(v string) *HotelOrderDetailInfoResponseBodyModuleOccupantInfoList {
	s.CardNo = &v
	return s
}

func (s *HotelOrderDetailInfoResponseBodyModuleOccupantInfoList) SetCardType(v int32) *HotelOrderDetailInfoResponseBodyModuleOccupantInfoList {
	s.CardType = &v
	return s
}

func (s *HotelOrderDetailInfoResponseBodyModuleOccupantInfoList) SetCostCenterInfoList(v []*HotelOrderDetailInfoResponseBodyModuleOccupantInfoListCostCenterInfoList) *HotelOrderDetailInfoResponseBodyModuleOccupantInfoList {
	s.CostCenterInfoList = v
	return s
}

func (s *HotelOrderDetailInfoResponseBodyModuleOccupantInfoList) SetCustomerType(v int32) *HotelOrderDetailInfoResponseBodyModuleOccupantInfoList {
	s.CustomerType = &v
	return s
}

func (s *HotelOrderDetailInfoResponseBodyModuleOccupantInfoList) SetDepartmentId(v string) *HotelOrderDetailInfoResponseBodyModuleOccupantInfoList {
	s.DepartmentId = &v
	return s
}

func (s *HotelOrderDetailInfoResponseBodyModuleOccupantInfoList) SetDepartmentName(v string) *HotelOrderDetailInfoResponseBodyModuleOccupantInfoList {
	s.DepartmentName = &v
	return s
}

func (s *HotelOrderDetailInfoResponseBodyModuleOccupantInfoList) SetEmail(v string) *HotelOrderDetailInfoResponseBodyModuleOccupantInfoList {
	s.Email = &v
	return s
}

func (s *HotelOrderDetailInfoResponseBodyModuleOccupantInfoList) SetEmployeeType(v int32) *HotelOrderDetailInfoResponseBodyModuleOccupantInfoList {
	s.EmployeeType = &v
	return s
}

func (s *HotelOrderDetailInfoResponseBodyModuleOccupantInfoList) SetFirstName(v string) *HotelOrderDetailInfoResponseBodyModuleOccupantInfoList {
	s.FirstName = &v
	return s
}

func (s *HotelOrderDetailInfoResponseBodyModuleOccupantInfoList) SetIsBooker(v bool) *HotelOrderDetailInfoResponseBodyModuleOccupantInfoList {
	s.IsBooker = &v
	return s
}

func (s *HotelOrderDetailInfoResponseBodyModuleOccupantInfoList) SetLastName(v string) *HotelOrderDetailInfoResponseBodyModuleOccupantInfoList {
	s.LastName = &v
	return s
}

func (s *HotelOrderDetailInfoResponseBodyModuleOccupantInfoList) SetName(v string) *HotelOrderDetailInfoResponseBodyModuleOccupantInfoList {
	s.Name = &v
	return s
}

func (s *HotelOrderDetailInfoResponseBodyModuleOccupantInfoList) SetPhone(v string) *HotelOrderDetailInfoResponseBodyModuleOccupantInfoList {
	s.Phone = &v
	return s
}

func (s *HotelOrderDetailInfoResponseBodyModuleOccupantInfoList) SetRoomNo(v int32) *HotelOrderDetailInfoResponseBodyModuleOccupantInfoList {
	s.RoomNo = &v
	return s
}

func (s *HotelOrderDetailInfoResponseBodyModuleOccupantInfoList) SetSelected(v bool) *HotelOrderDetailInfoResponseBodyModuleOccupantInfoList {
	s.Selected = &v
	return s
}

func (s *HotelOrderDetailInfoResponseBodyModuleOccupantInfoList) SetStaffNo(v string) *HotelOrderDetailInfoResponseBodyModuleOccupantInfoList {
	s.StaffNo = &v
	return s
}

func (s *HotelOrderDetailInfoResponseBodyModuleOccupantInfoList) SetUserType(v int32) *HotelOrderDetailInfoResponseBodyModuleOccupantInfoList {
	s.UserType = &v
	return s
}

func (s *HotelOrderDetailInfoResponseBodyModuleOccupantInfoList) Validate() error {
	return dara.Validate(s)
}

type HotelOrderDetailInfoResponseBodyModuleOccupantInfoListCostCenterInfoList struct {
	// example:
	//
	// demo
	CostCenterId *string `json:"cost_center_id,omitempty" xml:"cost_center_id,omitempty"`
	// example:
	//
	// demo
	CostCenterName *string `json:"cost_center_name,omitempty" xml:"cost_center_name,omitempty"`
	// example:
	//
	// demo
	CostCenterNo *string `json:"cost_center_no,omitempty" xml:"cost_center_no,omitempty"`
	// example:
	//
	// demo
	CostCenterPrices *int64 `json:"cost_center_prices,omitempty" xml:"cost_center_prices,omitempty"`
	// example:
	//
	// demo
	CostCenterRatios *int64 `json:"cost_center_ratios,omitempty" xml:"cost_center_ratios,omitempty"`
	// example:
	//
	// demo
	CostCenterSubjectCode *string `json:"cost_center_subject_code,omitempty" xml:"cost_center_subject_code,omitempty"`
	// example:
	//
	// demo
	CostCenterSubjectName *string `json:"cost_center_subject_name,omitempty" xml:"cost_center_subject_name,omitempty"`
	// example:
	//
	// demo
	SettleSubjectId *string `json:"settle_subject_id,omitempty" xml:"settle_subject_id,omitempty"`
	// example:
	//
	// demo
	SettleSubjectName *string `json:"settle_subject_name,omitempty" xml:"settle_subject_name,omitempty"`
	// example:
	//
	// demo
	SettleSubjectNo *string `json:"settle_subject_no,omitempty" xml:"settle_subject_no,omitempty"`
}

func (s HotelOrderDetailInfoResponseBodyModuleOccupantInfoListCostCenterInfoList) String() string {
	return dara.Prettify(s)
}

func (s HotelOrderDetailInfoResponseBodyModuleOccupantInfoListCostCenterInfoList) GoString() string {
	return s.String()
}

func (s *HotelOrderDetailInfoResponseBodyModuleOccupantInfoListCostCenterInfoList) GetCostCenterId() *string {
	return s.CostCenterId
}

func (s *HotelOrderDetailInfoResponseBodyModuleOccupantInfoListCostCenterInfoList) GetCostCenterName() *string {
	return s.CostCenterName
}

func (s *HotelOrderDetailInfoResponseBodyModuleOccupantInfoListCostCenterInfoList) GetCostCenterNo() *string {
	return s.CostCenterNo
}

func (s *HotelOrderDetailInfoResponseBodyModuleOccupantInfoListCostCenterInfoList) GetCostCenterPrices() *int64 {
	return s.CostCenterPrices
}

func (s *HotelOrderDetailInfoResponseBodyModuleOccupantInfoListCostCenterInfoList) GetCostCenterRatios() *int64 {
	return s.CostCenterRatios
}

func (s *HotelOrderDetailInfoResponseBodyModuleOccupantInfoListCostCenterInfoList) GetCostCenterSubjectCode() *string {
	return s.CostCenterSubjectCode
}

func (s *HotelOrderDetailInfoResponseBodyModuleOccupantInfoListCostCenterInfoList) GetCostCenterSubjectName() *string {
	return s.CostCenterSubjectName
}

func (s *HotelOrderDetailInfoResponseBodyModuleOccupantInfoListCostCenterInfoList) GetSettleSubjectId() *string {
	return s.SettleSubjectId
}

func (s *HotelOrderDetailInfoResponseBodyModuleOccupantInfoListCostCenterInfoList) GetSettleSubjectName() *string {
	return s.SettleSubjectName
}

func (s *HotelOrderDetailInfoResponseBodyModuleOccupantInfoListCostCenterInfoList) GetSettleSubjectNo() *string {
	return s.SettleSubjectNo
}

func (s *HotelOrderDetailInfoResponseBodyModuleOccupantInfoListCostCenterInfoList) SetCostCenterId(v string) *HotelOrderDetailInfoResponseBodyModuleOccupantInfoListCostCenterInfoList {
	s.CostCenterId = &v
	return s
}

func (s *HotelOrderDetailInfoResponseBodyModuleOccupantInfoListCostCenterInfoList) SetCostCenterName(v string) *HotelOrderDetailInfoResponseBodyModuleOccupantInfoListCostCenterInfoList {
	s.CostCenterName = &v
	return s
}

func (s *HotelOrderDetailInfoResponseBodyModuleOccupantInfoListCostCenterInfoList) SetCostCenterNo(v string) *HotelOrderDetailInfoResponseBodyModuleOccupantInfoListCostCenterInfoList {
	s.CostCenterNo = &v
	return s
}

func (s *HotelOrderDetailInfoResponseBodyModuleOccupantInfoListCostCenterInfoList) SetCostCenterPrices(v int64) *HotelOrderDetailInfoResponseBodyModuleOccupantInfoListCostCenterInfoList {
	s.CostCenterPrices = &v
	return s
}

func (s *HotelOrderDetailInfoResponseBodyModuleOccupantInfoListCostCenterInfoList) SetCostCenterRatios(v int64) *HotelOrderDetailInfoResponseBodyModuleOccupantInfoListCostCenterInfoList {
	s.CostCenterRatios = &v
	return s
}

func (s *HotelOrderDetailInfoResponseBodyModuleOccupantInfoListCostCenterInfoList) SetCostCenterSubjectCode(v string) *HotelOrderDetailInfoResponseBodyModuleOccupantInfoListCostCenterInfoList {
	s.CostCenterSubjectCode = &v
	return s
}

func (s *HotelOrderDetailInfoResponseBodyModuleOccupantInfoListCostCenterInfoList) SetCostCenterSubjectName(v string) *HotelOrderDetailInfoResponseBodyModuleOccupantInfoListCostCenterInfoList {
	s.CostCenterSubjectName = &v
	return s
}

func (s *HotelOrderDetailInfoResponseBodyModuleOccupantInfoListCostCenterInfoList) SetSettleSubjectId(v string) *HotelOrderDetailInfoResponseBodyModuleOccupantInfoListCostCenterInfoList {
	s.SettleSubjectId = &v
	return s
}

func (s *HotelOrderDetailInfoResponseBodyModuleOccupantInfoListCostCenterInfoList) SetSettleSubjectName(v string) *HotelOrderDetailInfoResponseBodyModuleOccupantInfoListCostCenterInfoList {
	s.SettleSubjectName = &v
	return s
}

func (s *HotelOrderDetailInfoResponseBodyModuleOccupantInfoListCostCenterInfoList) SetSettleSubjectNo(v string) *HotelOrderDetailInfoResponseBodyModuleOccupantInfoListCostCenterInfoList {
	s.SettleSubjectNo = &v
	return s
}

func (s *HotelOrderDetailInfoResponseBodyModuleOccupantInfoListCostCenterInfoList) Validate() error {
	return dara.Validate(s)
}

type HotelOrderDetailInfoResponseBodyModuleRoomNightPriceInfoList struct {
	Board *string `json:"board,omitempty" xml:"board,omitempty"`
	// example:
	//
	// 2
	BoardNum *int32 `json:"board_num,omitempty" xml:"board_num,omitempty"`
	// example:
	//
	// 1677600000000
	CheckIn *string `json:"check_in,omitempty" xml:"check_in,omitempty"`
	// example:
	//
	// 1399417428510
	RatePlanId   *string `json:"rate_plan_id,omitempty" xml:"rate_plan_id,omitempty"`
	RatePlanName *string `json:"rate_plan_name,omitempty" xml:"rate_plan_name,omitempty"`
	// example:
	//
	// 545993154510
	RoomId   *string `json:"room_id,omitempty" xml:"room_id,omitempty"`
	RoomName *string `json:"room_name,omitempty" xml:"room_name,omitempty"`
	// example:
	//
	// 50000
	RoomPrice *int64 `json:"room_price,omitempty" xml:"room_price,omitempty"`
}

func (s HotelOrderDetailInfoResponseBodyModuleRoomNightPriceInfoList) String() string {
	return dara.Prettify(s)
}

func (s HotelOrderDetailInfoResponseBodyModuleRoomNightPriceInfoList) GoString() string {
	return s.String()
}

func (s *HotelOrderDetailInfoResponseBodyModuleRoomNightPriceInfoList) GetBoard() *string {
	return s.Board
}

func (s *HotelOrderDetailInfoResponseBodyModuleRoomNightPriceInfoList) GetBoardNum() *int32 {
	return s.BoardNum
}

func (s *HotelOrderDetailInfoResponseBodyModuleRoomNightPriceInfoList) GetCheckIn() *string {
	return s.CheckIn
}

func (s *HotelOrderDetailInfoResponseBodyModuleRoomNightPriceInfoList) GetRatePlanId() *string {
	return s.RatePlanId
}

func (s *HotelOrderDetailInfoResponseBodyModuleRoomNightPriceInfoList) GetRatePlanName() *string {
	return s.RatePlanName
}

func (s *HotelOrderDetailInfoResponseBodyModuleRoomNightPriceInfoList) GetRoomId() *string {
	return s.RoomId
}

func (s *HotelOrderDetailInfoResponseBodyModuleRoomNightPriceInfoList) GetRoomName() *string {
	return s.RoomName
}

func (s *HotelOrderDetailInfoResponseBodyModuleRoomNightPriceInfoList) GetRoomPrice() *int64 {
	return s.RoomPrice
}

func (s *HotelOrderDetailInfoResponseBodyModuleRoomNightPriceInfoList) SetBoard(v string) *HotelOrderDetailInfoResponseBodyModuleRoomNightPriceInfoList {
	s.Board = &v
	return s
}

func (s *HotelOrderDetailInfoResponseBodyModuleRoomNightPriceInfoList) SetBoardNum(v int32) *HotelOrderDetailInfoResponseBodyModuleRoomNightPriceInfoList {
	s.BoardNum = &v
	return s
}

func (s *HotelOrderDetailInfoResponseBodyModuleRoomNightPriceInfoList) SetCheckIn(v string) *HotelOrderDetailInfoResponseBodyModuleRoomNightPriceInfoList {
	s.CheckIn = &v
	return s
}

func (s *HotelOrderDetailInfoResponseBodyModuleRoomNightPriceInfoList) SetRatePlanId(v string) *HotelOrderDetailInfoResponseBodyModuleRoomNightPriceInfoList {
	s.RatePlanId = &v
	return s
}

func (s *HotelOrderDetailInfoResponseBodyModuleRoomNightPriceInfoList) SetRatePlanName(v string) *HotelOrderDetailInfoResponseBodyModuleRoomNightPriceInfoList {
	s.RatePlanName = &v
	return s
}

func (s *HotelOrderDetailInfoResponseBodyModuleRoomNightPriceInfoList) SetRoomId(v string) *HotelOrderDetailInfoResponseBodyModuleRoomNightPriceInfoList {
	s.RoomId = &v
	return s
}

func (s *HotelOrderDetailInfoResponseBodyModuleRoomNightPriceInfoList) SetRoomName(v string) *HotelOrderDetailInfoResponseBodyModuleRoomNightPriceInfoList {
	s.RoomName = &v
	return s
}

func (s *HotelOrderDetailInfoResponseBodyModuleRoomNightPriceInfoList) SetRoomPrice(v int64) *HotelOrderDetailInfoResponseBodyModuleRoomNightPriceInfoList {
	s.RoomPrice = &v
	return s
}

func (s *HotelOrderDetailInfoResponseBodyModuleRoomNightPriceInfoList) Validate() error {
	return dara.Validate(s)
}
