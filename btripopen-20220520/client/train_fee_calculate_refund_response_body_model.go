// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTrainFeeCalculateRefundResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *TrainFeeCalculateRefundResponseBody
	GetCode() *string
	SetMessage(v string) *TrainFeeCalculateRefundResponseBody
	GetMessage() *string
	SetModule(v *TrainFeeCalculateRefundResponseBodyModule) *TrainFeeCalculateRefundResponseBody
	GetModule() *TrainFeeCalculateRefundResponseBodyModule
	SetRequestId(v string) *TrainFeeCalculateRefundResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *TrainFeeCalculateRefundResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *TrainFeeCalculateRefundResponseBody
	GetTraceId() *string
}

type TrainFeeCalculateRefundResponseBody struct {
	// example:
	//
	// 200
	Code    *string `json:"code,omitempty" xml:"code,omitempty"`
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// module
	Module *TrainFeeCalculateRefundResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	// example:
	//
	// 210bc4c416859457407135025d509c
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// traceId
	//
	// example:
	//
	// 21041ce316577904808056433edbb2
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s TrainFeeCalculateRefundResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TrainFeeCalculateRefundResponseBody) GoString() string {
	return s.String()
}

func (s *TrainFeeCalculateRefundResponseBody) GetCode() *string {
	return s.Code
}

func (s *TrainFeeCalculateRefundResponseBody) GetMessage() *string {
	return s.Message
}

func (s *TrainFeeCalculateRefundResponseBody) GetModule() *TrainFeeCalculateRefundResponseBodyModule {
	return s.Module
}

func (s *TrainFeeCalculateRefundResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TrainFeeCalculateRefundResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *TrainFeeCalculateRefundResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *TrainFeeCalculateRefundResponseBody) SetCode(v string) *TrainFeeCalculateRefundResponseBody {
	s.Code = &v
	return s
}

func (s *TrainFeeCalculateRefundResponseBody) SetMessage(v string) *TrainFeeCalculateRefundResponseBody {
	s.Message = &v
	return s
}

func (s *TrainFeeCalculateRefundResponseBody) SetModule(v *TrainFeeCalculateRefundResponseBodyModule) *TrainFeeCalculateRefundResponseBody {
	s.Module = v
	return s
}

func (s *TrainFeeCalculateRefundResponseBody) SetRequestId(v string) *TrainFeeCalculateRefundResponseBody {
	s.RequestId = &v
	return s
}

func (s *TrainFeeCalculateRefundResponseBody) SetSuccess(v bool) *TrainFeeCalculateRefundResponseBody {
	s.Success = &v
	return s
}

func (s *TrainFeeCalculateRefundResponseBody) SetTraceId(v string) *TrainFeeCalculateRefundResponseBody {
	s.TraceId = &v
	return s
}

func (s *TrainFeeCalculateRefundResponseBody) Validate() error {
	return dara.Validate(s)
}

type TrainFeeCalculateRefundResponseBodyModule struct {
	// example:
	//
	// 123445
	DistributeOrderId *string `json:"distribute_order_id,omitempty" xml:"distribute_order_id,omitempty"`
	// example:
	//
	// 1017002195370467200
	OrderId            *string                                                        `json:"order_id,omitempty" xml:"order_id,omitempty"`
	RefundTrainDetails []*TrainFeeCalculateRefundResponseBodyModuleRefundTrainDetails `json:"refund_train_details,omitempty" xml:"refund_train_details,omitempty" type:"Repeated"`
}

func (s TrainFeeCalculateRefundResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s TrainFeeCalculateRefundResponseBodyModule) GoString() string {
	return s.String()
}

func (s *TrainFeeCalculateRefundResponseBodyModule) GetDistributeOrderId() *string {
	return s.DistributeOrderId
}

func (s *TrainFeeCalculateRefundResponseBodyModule) GetOrderId() *string {
	return s.OrderId
}

func (s *TrainFeeCalculateRefundResponseBodyModule) GetRefundTrainDetails() []*TrainFeeCalculateRefundResponseBodyModuleRefundTrainDetails {
	return s.RefundTrainDetails
}

func (s *TrainFeeCalculateRefundResponseBodyModule) SetDistributeOrderId(v string) *TrainFeeCalculateRefundResponseBodyModule {
	s.DistributeOrderId = &v
	return s
}

func (s *TrainFeeCalculateRefundResponseBodyModule) SetOrderId(v string) *TrainFeeCalculateRefundResponseBodyModule {
	s.OrderId = &v
	return s
}

func (s *TrainFeeCalculateRefundResponseBodyModule) SetRefundTrainDetails(v []*TrainFeeCalculateRefundResponseBodyModuleRefundTrainDetails) *TrainFeeCalculateRefundResponseBodyModule {
	s.RefundTrainDetails = v
	return s
}

func (s *TrainFeeCalculateRefundResponseBodyModule) Validate() error {
	return dara.Validate(s)
}

type TrainFeeCalculateRefundResponseBodyModuleRefundTrainDetails struct {
	// example:
	//
	// BTC
	ArrStationCode *string `json:"arr_station_code,omitempty" xml:"arr_station_code,omitempty"`
	// example:
	//
	// BDC
	DepStationCode *string `json:"dep_station_code,omitempty" xml:"dep_station_code,omitempty"`
	// example:
	//
	// 2024-05-06 15:19:01
	DepTime             *string                                                                           `json:"dep_time,omitempty" xml:"dep_time,omitempty"`
	RefundTicketDetails []*TrainFeeCalculateRefundResponseBodyModuleRefundTrainDetailsRefundTicketDetails `json:"refund_ticket_details,omitempty" xml:"refund_ticket_details,omitempty" type:"Repeated"`
	// example:
	//
	// K1234
	TrainNo *string `json:"train_no,omitempty" xml:"train_no,omitempty"`
}

func (s TrainFeeCalculateRefundResponseBodyModuleRefundTrainDetails) String() string {
	return dara.Prettify(s)
}

func (s TrainFeeCalculateRefundResponseBodyModuleRefundTrainDetails) GoString() string {
	return s.String()
}

func (s *TrainFeeCalculateRefundResponseBodyModuleRefundTrainDetails) GetArrStationCode() *string {
	return s.ArrStationCode
}

func (s *TrainFeeCalculateRefundResponseBodyModuleRefundTrainDetails) GetDepStationCode() *string {
	return s.DepStationCode
}

func (s *TrainFeeCalculateRefundResponseBodyModuleRefundTrainDetails) GetDepTime() *string {
	return s.DepTime
}

func (s *TrainFeeCalculateRefundResponseBodyModuleRefundTrainDetails) GetRefundTicketDetails() []*TrainFeeCalculateRefundResponseBodyModuleRefundTrainDetailsRefundTicketDetails {
	return s.RefundTicketDetails
}

func (s *TrainFeeCalculateRefundResponseBodyModuleRefundTrainDetails) GetTrainNo() *string {
	return s.TrainNo
}

func (s *TrainFeeCalculateRefundResponseBodyModuleRefundTrainDetails) SetArrStationCode(v string) *TrainFeeCalculateRefundResponseBodyModuleRefundTrainDetails {
	s.ArrStationCode = &v
	return s
}

func (s *TrainFeeCalculateRefundResponseBodyModuleRefundTrainDetails) SetDepStationCode(v string) *TrainFeeCalculateRefundResponseBodyModuleRefundTrainDetails {
	s.DepStationCode = &v
	return s
}

func (s *TrainFeeCalculateRefundResponseBodyModuleRefundTrainDetails) SetDepTime(v string) *TrainFeeCalculateRefundResponseBodyModuleRefundTrainDetails {
	s.DepTime = &v
	return s
}

func (s *TrainFeeCalculateRefundResponseBodyModuleRefundTrainDetails) SetRefundTicketDetails(v []*TrainFeeCalculateRefundResponseBodyModuleRefundTrainDetailsRefundTicketDetails) *TrainFeeCalculateRefundResponseBodyModuleRefundTrainDetails {
	s.RefundTicketDetails = v
	return s
}

func (s *TrainFeeCalculateRefundResponseBodyModuleRefundTrainDetails) SetTrainNo(v string) *TrainFeeCalculateRefundResponseBodyModuleRefundTrainDetails {
	s.TrainNo = &v
	return s
}

func (s *TrainFeeCalculateRefundResponseBodyModuleRefundTrainDetails) Validate() error {
	return dara.Validate(s)
}

type TrainFeeCalculateRefundResponseBodyModuleRefundTrainDetailsRefundTicketDetails struct {
	// example:
	//
	// true
	CanRefund     *bool                                                                                        `json:"can_refund,omitempty" xml:"can_refund,omitempty"`
	PassengerInfo *TrainFeeCalculateRefundResponseBodyModuleRefundTrainDetailsRefundTicketDetailsPassengerInfo `json:"passenger_info,omitempty" xml:"passenger_info,omitempty" type:"Struct"`
	// example:
	//
	// 90
	RefundCostFee *int64 `json:"refund_cost_fee,omitempty" xml:"refund_cost_fee,omitempty"`
	// example:
	//
	// 10
	RefundPrice *int64 `json:"refund_price,omitempty" xml:"refund_price,omitempty"`
	// example:
	//
	// 90
	RefundRate *int64 `json:"refund_rate,omitempty" xml:"refund_rate,omitempty"`
	// example:
	//
	// 100
	TicketPrice *int64 `json:"ticket_price,omitempty" xml:"ticket_price,omitempty"`
}

func (s TrainFeeCalculateRefundResponseBodyModuleRefundTrainDetailsRefundTicketDetails) String() string {
	return dara.Prettify(s)
}

func (s TrainFeeCalculateRefundResponseBodyModuleRefundTrainDetailsRefundTicketDetails) GoString() string {
	return s.String()
}

func (s *TrainFeeCalculateRefundResponseBodyModuleRefundTrainDetailsRefundTicketDetails) GetCanRefund() *bool {
	return s.CanRefund
}

func (s *TrainFeeCalculateRefundResponseBodyModuleRefundTrainDetailsRefundTicketDetails) GetPassengerInfo() *TrainFeeCalculateRefundResponseBodyModuleRefundTrainDetailsRefundTicketDetailsPassengerInfo {
	return s.PassengerInfo
}

func (s *TrainFeeCalculateRefundResponseBodyModuleRefundTrainDetailsRefundTicketDetails) GetRefundCostFee() *int64 {
	return s.RefundCostFee
}

func (s *TrainFeeCalculateRefundResponseBodyModuleRefundTrainDetailsRefundTicketDetails) GetRefundPrice() *int64 {
	return s.RefundPrice
}

func (s *TrainFeeCalculateRefundResponseBodyModuleRefundTrainDetailsRefundTicketDetails) GetRefundRate() *int64 {
	return s.RefundRate
}

func (s *TrainFeeCalculateRefundResponseBodyModuleRefundTrainDetailsRefundTicketDetails) GetTicketPrice() *int64 {
	return s.TicketPrice
}

func (s *TrainFeeCalculateRefundResponseBodyModuleRefundTrainDetailsRefundTicketDetails) SetCanRefund(v bool) *TrainFeeCalculateRefundResponseBodyModuleRefundTrainDetailsRefundTicketDetails {
	s.CanRefund = &v
	return s
}

func (s *TrainFeeCalculateRefundResponseBodyModuleRefundTrainDetailsRefundTicketDetails) SetPassengerInfo(v *TrainFeeCalculateRefundResponseBodyModuleRefundTrainDetailsRefundTicketDetailsPassengerInfo) *TrainFeeCalculateRefundResponseBodyModuleRefundTrainDetailsRefundTicketDetails {
	s.PassengerInfo = v
	return s
}

func (s *TrainFeeCalculateRefundResponseBodyModuleRefundTrainDetailsRefundTicketDetails) SetRefundCostFee(v int64) *TrainFeeCalculateRefundResponseBodyModuleRefundTrainDetailsRefundTicketDetails {
	s.RefundCostFee = &v
	return s
}

func (s *TrainFeeCalculateRefundResponseBodyModuleRefundTrainDetailsRefundTicketDetails) SetRefundPrice(v int64) *TrainFeeCalculateRefundResponseBodyModuleRefundTrainDetailsRefundTicketDetails {
	s.RefundPrice = &v
	return s
}

func (s *TrainFeeCalculateRefundResponseBodyModuleRefundTrainDetailsRefundTicketDetails) SetRefundRate(v int64) *TrainFeeCalculateRefundResponseBodyModuleRefundTrainDetailsRefundTicketDetails {
	s.RefundRate = &v
	return s
}

func (s *TrainFeeCalculateRefundResponseBodyModuleRefundTrainDetailsRefundTicketDetails) SetTicketPrice(v int64) *TrainFeeCalculateRefundResponseBodyModuleRefundTrainDetailsRefundTicketDetails {
	s.TicketPrice = &v
	return s
}

func (s *TrainFeeCalculateRefundResponseBodyModuleRefundTrainDetailsRefundTicketDetails) Validate() error {
	return dara.Validate(s)
}

type TrainFeeCalculateRefundResponseBodyModuleRefundTrainDetailsRefundTicketDetailsPassengerInfo struct {
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
	// 949c9f34f677a0e5d249dfc94f5e62cc7
	PassengerName *string `json:"passenger_name,omitempty" xml:"passenger_name,omitempty"`
}

func (s TrainFeeCalculateRefundResponseBodyModuleRefundTrainDetailsRefundTicketDetailsPassengerInfo) String() string {
	return dara.Prettify(s)
}

func (s TrainFeeCalculateRefundResponseBodyModuleRefundTrainDetailsRefundTicketDetailsPassengerInfo) GoString() string {
	return s.String()
}

func (s *TrainFeeCalculateRefundResponseBodyModuleRefundTrainDetailsRefundTicketDetailsPassengerInfo) GetPassengerCertNo() *string {
	return s.PassengerCertNo
}

func (s *TrainFeeCalculateRefundResponseBodyModuleRefundTrainDetailsRefundTicketDetailsPassengerInfo) GetPassengerCertType() *string {
	return s.PassengerCertType
}

func (s *TrainFeeCalculateRefundResponseBodyModuleRefundTrainDetailsRefundTicketDetailsPassengerInfo) GetPassengerId() *string {
	return s.PassengerId
}

func (s *TrainFeeCalculateRefundResponseBodyModuleRefundTrainDetailsRefundTicketDetailsPassengerInfo) GetPassengerName() *string {
	return s.PassengerName
}

func (s *TrainFeeCalculateRefundResponseBodyModuleRefundTrainDetailsRefundTicketDetailsPassengerInfo) SetPassengerCertNo(v string) *TrainFeeCalculateRefundResponseBodyModuleRefundTrainDetailsRefundTicketDetailsPassengerInfo {
	s.PassengerCertNo = &v
	return s
}

func (s *TrainFeeCalculateRefundResponseBodyModuleRefundTrainDetailsRefundTicketDetailsPassengerInfo) SetPassengerCertType(v string) *TrainFeeCalculateRefundResponseBodyModuleRefundTrainDetailsRefundTicketDetailsPassengerInfo {
	s.PassengerCertType = &v
	return s
}

func (s *TrainFeeCalculateRefundResponseBodyModuleRefundTrainDetailsRefundTicketDetailsPassengerInfo) SetPassengerId(v string) *TrainFeeCalculateRefundResponseBodyModuleRefundTrainDetailsRefundTicketDetailsPassengerInfo {
	s.PassengerId = &v
	return s
}

func (s *TrainFeeCalculateRefundResponseBodyModuleRefundTrainDetailsRefundTicketDetailsPassengerInfo) SetPassengerName(v string) *TrainFeeCalculateRefundResponseBodyModuleRefundTrainDetailsRefundTicketDetailsPassengerInfo {
	s.PassengerName = &v
	return s
}

func (s *TrainFeeCalculateRefundResponseBodyModuleRefundTrainDetailsRefundTicketDetailsPassengerInfo) Validate() error {
	return dara.Validate(s)
}
