// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTrainFeeCalculateChangeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *TrainFeeCalculateChangeResponseBody
	GetCode() *string
	SetMessage(v string) *TrainFeeCalculateChangeResponseBody
	GetMessage() *string
	SetModule(v *TrainFeeCalculateChangeResponseBodyModule) *TrainFeeCalculateChangeResponseBody
	GetModule() *TrainFeeCalculateChangeResponseBodyModule
	SetRequestId(v string) *TrainFeeCalculateChangeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *TrainFeeCalculateChangeResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *TrainFeeCalculateChangeResponseBody
	GetTraceId() *string
}

type TrainFeeCalculateChangeResponseBody struct {
	// example:
	//
	// 200
	Code    *string `json:"code,omitempty" xml:"code,omitempty"`
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// module
	Module *TrainFeeCalculateChangeResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	// example:
	//
	// 407543AF-2BD9-5890-BD92-9D1AB7218B27
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

func (s TrainFeeCalculateChangeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TrainFeeCalculateChangeResponseBody) GoString() string {
	return s.String()
}

func (s *TrainFeeCalculateChangeResponseBody) GetCode() *string {
	return s.Code
}

func (s *TrainFeeCalculateChangeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *TrainFeeCalculateChangeResponseBody) GetModule() *TrainFeeCalculateChangeResponseBodyModule {
	return s.Module
}

func (s *TrainFeeCalculateChangeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TrainFeeCalculateChangeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *TrainFeeCalculateChangeResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *TrainFeeCalculateChangeResponseBody) SetCode(v string) *TrainFeeCalculateChangeResponseBody {
	s.Code = &v
	return s
}

func (s *TrainFeeCalculateChangeResponseBody) SetMessage(v string) *TrainFeeCalculateChangeResponseBody {
	s.Message = &v
	return s
}

func (s *TrainFeeCalculateChangeResponseBody) SetModule(v *TrainFeeCalculateChangeResponseBodyModule) *TrainFeeCalculateChangeResponseBody {
	s.Module = v
	return s
}

func (s *TrainFeeCalculateChangeResponseBody) SetRequestId(v string) *TrainFeeCalculateChangeResponseBody {
	s.RequestId = &v
	return s
}

func (s *TrainFeeCalculateChangeResponseBody) SetSuccess(v bool) *TrainFeeCalculateChangeResponseBody {
	s.Success = &v
	return s
}

func (s *TrainFeeCalculateChangeResponseBody) SetTraceId(v string) *TrainFeeCalculateChangeResponseBody {
	s.TraceId = &v
	return s
}

func (s *TrainFeeCalculateChangeResponseBody) Validate() error {
	return dara.Validate(s)
}

type TrainFeeCalculateChangeResponseBodyModule struct {
	ChangeTrainDetails []*TrainFeeCalculateChangeResponseBodyModuleChangeTrainDetails `json:"change_train_details,omitempty" xml:"change_train_details,omitempty" type:"Repeated"`
	// example:
	//
	// 123456
	DistributeOrderId *string `json:"distribute_order_id,omitempty" xml:"distribute_order_id,omitempty"`
	// example:
	//
	// 1683901850297448200
	OrderId *string `json:"order_id,omitempty" xml:"order_id,omitempty"`
}

func (s TrainFeeCalculateChangeResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s TrainFeeCalculateChangeResponseBodyModule) GoString() string {
	return s.String()
}

func (s *TrainFeeCalculateChangeResponseBodyModule) GetChangeTrainDetails() []*TrainFeeCalculateChangeResponseBodyModuleChangeTrainDetails {
	return s.ChangeTrainDetails
}

func (s *TrainFeeCalculateChangeResponseBodyModule) GetDistributeOrderId() *string {
	return s.DistributeOrderId
}

func (s *TrainFeeCalculateChangeResponseBodyModule) GetOrderId() *string {
	return s.OrderId
}

func (s *TrainFeeCalculateChangeResponseBodyModule) SetChangeTrainDetails(v []*TrainFeeCalculateChangeResponseBodyModuleChangeTrainDetails) *TrainFeeCalculateChangeResponseBodyModule {
	s.ChangeTrainDetails = v
	return s
}

func (s *TrainFeeCalculateChangeResponseBodyModule) SetDistributeOrderId(v string) *TrainFeeCalculateChangeResponseBodyModule {
	s.DistributeOrderId = &v
	return s
}

func (s *TrainFeeCalculateChangeResponseBodyModule) SetOrderId(v string) *TrainFeeCalculateChangeResponseBodyModule {
	s.OrderId = &v
	return s
}

func (s *TrainFeeCalculateChangeResponseBodyModule) Validate() error {
	return dara.Validate(s)
}

type TrainFeeCalculateChangeResponseBodyModuleChangeTrainDetails struct {
	// example:
	//
	// BTC
	ArrStationCode      *string                                                                           `json:"arr_station_code,omitempty" xml:"arr_station_code,omitempty"`
	ChangeTicketDetails []*TrainFeeCalculateChangeResponseBodyModuleChangeTrainDetailsChangeTicketDetails `json:"change_ticket_details,omitempty" xml:"change_ticket_details,omitempty" type:"Repeated"`
	// example:
	//
	// BDC
	DepStationCode *string `json:"dep_station_code,omitempty" xml:"dep_station_code,omitempty"`
	// example:
	//
	// 2024-05-06 15:19:01
	DepTime *string `json:"dep_time,omitempty" xml:"dep_time,omitempty"`
	// example:
	//
	// K2345
	TrainNo *string `json:"train_no,omitempty" xml:"train_no,omitempty"`
}

func (s TrainFeeCalculateChangeResponseBodyModuleChangeTrainDetails) String() string {
	return dara.Prettify(s)
}

func (s TrainFeeCalculateChangeResponseBodyModuleChangeTrainDetails) GoString() string {
	return s.String()
}

func (s *TrainFeeCalculateChangeResponseBodyModuleChangeTrainDetails) GetArrStationCode() *string {
	return s.ArrStationCode
}

func (s *TrainFeeCalculateChangeResponseBodyModuleChangeTrainDetails) GetChangeTicketDetails() []*TrainFeeCalculateChangeResponseBodyModuleChangeTrainDetailsChangeTicketDetails {
	return s.ChangeTicketDetails
}

func (s *TrainFeeCalculateChangeResponseBodyModuleChangeTrainDetails) GetDepStationCode() *string {
	return s.DepStationCode
}

func (s *TrainFeeCalculateChangeResponseBodyModuleChangeTrainDetails) GetDepTime() *string {
	return s.DepTime
}

func (s *TrainFeeCalculateChangeResponseBodyModuleChangeTrainDetails) GetTrainNo() *string {
	return s.TrainNo
}

func (s *TrainFeeCalculateChangeResponseBodyModuleChangeTrainDetails) SetArrStationCode(v string) *TrainFeeCalculateChangeResponseBodyModuleChangeTrainDetails {
	s.ArrStationCode = &v
	return s
}

func (s *TrainFeeCalculateChangeResponseBodyModuleChangeTrainDetails) SetChangeTicketDetails(v []*TrainFeeCalculateChangeResponseBodyModuleChangeTrainDetailsChangeTicketDetails) *TrainFeeCalculateChangeResponseBodyModuleChangeTrainDetails {
	s.ChangeTicketDetails = v
	return s
}

func (s *TrainFeeCalculateChangeResponseBodyModuleChangeTrainDetails) SetDepStationCode(v string) *TrainFeeCalculateChangeResponseBodyModuleChangeTrainDetails {
	s.DepStationCode = &v
	return s
}

func (s *TrainFeeCalculateChangeResponseBodyModuleChangeTrainDetails) SetDepTime(v string) *TrainFeeCalculateChangeResponseBodyModuleChangeTrainDetails {
	s.DepTime = &v
	return s
}

func (s *TrainFeeCalculateChangeResponseBodyModuleChangeTrainDetails) SetTrainNo(v string) *TrainFeeCalculateChangeResponseBodyModuleChangeTrainDetails {
	s.TrainNo = &v
	return s
}

func (s *TrainFeeCalculateChangeResponseBodyModuleChangeTrainDetails) Validate() error {
	return dara.Validate(s)
}

type TrainFeeCalculateChangeResponseBodyModuleChangeTrainDetailsChangeTicketDetails struct {
	// example:
	//
	// 0
	ChangeFee *int64 `json:"change_fee,omitempty" xml:"change_fee,omitempty"`
	// example:
	//
	// 5
	ChangeRate *int64 `json:"change_rate,omitempty" xml:"change_rate,omitempty"`
	// example:
	//
	// 10000
	ChangeRefundFee *int64 `json:"change_refund_fee,omitempty" xml:"change_refund_fee,omitempty"`
	// example:
	//
	// 10
	ChangeRefundRate *int64                                                                                       `json:"change_refund_rate,omitempty" xml:"change_refund_rate,omitempty"`
	PassengerInfo    *TrainFeeCalculateChangeResponseBodyModuleChangeTrainDetailsChangeTicketDetailsPassengerInfo `json:"passenger_info,omitempty" xml:"passenger_info,omitempty" type:"Struct"`
	// example:
	//
	// 14
	SeatType *string `json:"seat_type,omitempty" xml:"seat_type,omitempty"`
	// example:
	//
	// 100
	TicketPrice *int64 `json:"ticket_price,omitempty" xml:"ticket_price,omitempty"`
}

func (s TrainFeeCalculateChangeResponseBodyModuleChangeTrainDetailsChangeTicketDetails) String() string {
	return dara.Prettify(s)
}

func (s TrainFeeCalculateChangeResponseBodyModuleChangeTrainDetailsChangeTicketDetails) GoString() string {
	return s.String()
}

func (s *TrainFeeCalculateChangeResponseBodyModuleChangeTrainDetailsChangeTicketDetails) GetChangeFee() *int64 {
	return s.ChangeFee
}

func (s *TrainFeeCalculateChangeResponseBodyModuleChangeTrainDetailsChangeTicketDetails) GetChangeRate() *int64 {
	return s.ChangeRate
}

func (s *TrainFeeCalculateChangeResponseBodyModuleChangeTrainDetailsChangeTicketDetails) GetChangeRefundFee() *int64 {
	return s.ChangeRefundFee
}

func (s *TrainFeeCalculateChangeResponseBodyModuleChangeTrainDetailsChangeTicketDetails) GetChangeRefundRate() *int64 {
	return s.ChangeRefundRate
}

func (s *TrainFeeCalculateChangeResponseBodyModuleChangeTrainDetailsChangeTicketDetails) GetPassengerInfo() *TrainFeeCalculateChangeResponseBodyModuleChangeTrainDetailsChangeTicketDetailsPassengerInfo {
	return s.PassengerInfo
}

func (s *TrainFeeCalculateChangeResponseBodyModuleChangeTrainDetailsChangeTicketDetails) GetSeatType() *string {
	return s.SeatType
}

func (s *TrainFeeCalculateChangeResponseBodyModuleChangeTrainDetailsChangeTicketDetails) GetTicketPrice() *int64 {
	return s.TicketPrice
}

func (s *TrainFeeCalculateChangeResponseBodyModuleChangeTrainDetailsChangeTicketDetails) SetChangeFee(v int64) *TrainFeeCalculateChangeResponseBodyModuleChangeTrainDetailsChangeTicketDetails {
	s.ChangeFee = &v
	return s
}

func (s *TrainFeeCalculateChangeResponseBodyModuleChangeTrainDetailsChangeTicketDetails) SetChangeRate(v int64) *TrainFeeCalculateChangeResponseBodyModuleChangeTrainDetailsChangeTicketDetails {
	s.ChangeRate = &v
	return s
}

func (s *TrainFeeCalculateChangeResponseBodyModuleChangeTrainDetailsChangeTicketDetails) SetChangeRefundFee(v int64) *TrainFeeCalculateChangeResponseBodyModuleChangeTrainDetailsChangeTicketDetails {
	s.ChangeRefundFee = &v
	return s
}

func (s *TrainFeeCalculateChangeResponseBodyModuleChangeTrainDetailsChangeTicketDetails) SetChangeRefundRate(v int64) *TrainFeeCalculateChangeResponseBodyModuleChangeTrainDetailsChangeTicketDetails {
	s.ChangeRefundRate = &v
	return s
}

func (s *TrainFeeCalculateChangeResponseBodyModuleChangeTrainDetailsChangeTicketDetails) SetPassengerInfo(v *TrainFeeCalculateChangeResponseBodyModuleChangeTrainDetailsChangeTicketDetailsPassengerInfo) *TrainFeeCalculateChangeResponseBodyModuleChangeTrainDetailsChangeTicketDetails {
	s.PassengerInfo = v
	return s
}

func (s *TrainFeeCalculateChangeResponseBodyModuleChangeTrainDetailsChangeTicketDetails) SetSeatType(v string) *TrainFeeCalculateChangeResponseBodyModuleChangeTrainDetailsChangeTicketDetails {
	s.SeatType = &v
	return s
}

func (s *TrainFeeCalculateChangeResponseBodyModuleChangeTrainDetailsChangeTicketDetails) SetTicketPrice(v int64) *TrainFeeCalculateChangeResponseBodyModuleChangeTrainDetailsChangeTicketDetails {
	s.TicketPrice = &v
	return s
}

func (s *TrainFeeCalculateChangeResponseBodyModuleChangeTrainDetailsChangeTicketDetails) Validate() error {
	return dara.Validate(s)
}

type TrainFeeCalculateChangeResponseBodyModuleChangeTrainDetailsChangeTicketDetailsPassengerInfo struct {
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
	// 123
	PassengerId *string `json:"passenger_id,omitempty" xml:"passenger_id,omitempty"`
	// example:
	//
	// 949c9f34f677a0e5d249dfc94f5e62cc7
	PassengerName *string `json:"passenger_name,omitempty" xml:"passenger_name,omitempty"`
}

func (s TrainFeeCalculateChangeResponseBodyModuleChangeTrainDetailsChangeTicketDetailsPassengerInfo) String() string {
	return dara.Prettify(s)
}

func (s TrainFeeCalculateChangeResponseBodyModuleChangeTrainDetailsChangeTicketDetailsPassengerInfo) GoString() string {
	return s.String()
}

func (s *TrainFeeCalculateChangeResponseBodyModuleChangeTrainDetailsChangeTicketDetailsPassengerInfo) GetPassengerCertNo() *string {
	return s.PassengerCertNo
}

func (s *TrainFeeCalculateChangeResponseBodyModuleChangeTrainDetailsChangeTicketDetailsPassengerInfo) GetPassengerCertType() *string {
	return s.PassengerCertType
}

func (s *TrainFeeCalculateChangeResponseBodyModuleChangeTrainDetailsChangeTicketDetailsPassengerInfo) GetPassengerId() *string {
	return s.PassengerId
}

func (s *TrainFeeCalculateChangeResponseBodyModuleChangeTrainDetailsChangeTicketDetailsPassengerInfo) GetPassengerName() *string {
	return s.PassengerName
}

func (s *TrainFeeCalculateChangeResponseBodyModuleChangeTrainDetailsChangeTicketDetailsPassengerInfo) SetPassengerCertNo(v string) *TrainFeeCalculateChangeResponseBodyModuleChangeTrainDetailsChangeTicketDetailsPassengerInfo {
	s.PassengerCertNo = &v
	return s
}

func (s *TrainFeeCalculateChangeResponseBodyModuleChangeTrainDetailsChangeTicketDetailsPassengerInfo) SetPassengerCertType(v string) *TrainFeeCalculateChangeResponseBodyModuleChangeTrainDetailsChangeTicketDetailsPassengerInfo {
	s.PassengerCertType = &v
	return s
}

func (s *TrainFeeCalculateChangeResponseBodyModuleChangeTrainDetailsChangeTicketDetailsPassengerInfo) SetPassengerId(v string) *TrainFeeCalculateChangeResponseBodyModuleChangeTrainDetailsChangeTicketDetailsPassengerInfo {
	s.PassengerId = &v
	return s
}

func (s *TrainFeeCalculateChangeResponseBodyModuleChangeTrainDetailsChangeTicketDetailsPassengerInfo) SetPassengerName(v string) *TrainFeeCalculateChangeResponseBodyModuleChangeTrainDetailsChangeTicketDetailsPassengerInfo {
	s.PassengerName = &v
	return s
}

func (s *TrainFeeCalculateChangeResponseBodyModuleChangeTrainDetailsChangeTicketDetailsPassengerInfo) Validate() error {
	return dara.Validate(s)
}
