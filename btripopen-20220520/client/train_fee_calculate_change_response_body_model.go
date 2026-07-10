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
	Code      *string                                    `json:"code,omitempty" xml:"code,omitempty"`
	Message   *string                                    `json:"message,omitempty" xml:"message,omitempty"`
	Module    *TrainFeeCalculateChangeResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	RequestId *string                                    `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool                                      `json:"success,omitempty" xml:"success,omitempty"`
	TraceId   *string                                    `json:"traceId,omitempty" xml:"traceId,omitempty"`
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
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type TrainFeeCalculateChangeResponseBodyModule struct {
	ChangeTrainDetails []*TrainFeeCalculateChangeResponseBodyModuleChangeTrainDetails `json:"change_train_details,omitempty" xml:"change_train_details,omitempty" type:"Repeated"`
	DistributeOrderId  *string                                                        `json:"distribute_order_id,omitempty" xml:"distribute_order_id,omitempty"`
	OrderId            *string                                                        `json:"order_id,omitempty" xml:"order_id,omitempty"`
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
	if s.ChangeTrainDetails != nil {
		for _, item := range s.ChangeTrainDetails {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type TrainFeeCalculateChangeResponseBodyModuleChangeTrainDetails struct {
	ArrStationCode      *string                                                                           `json:"arr_station_code,omitempty" xml:"arr_station_code,omitempty"`
	ChangeTicketDetails []*TrainFeeCalculateChangeResponseBodyModuleChangeTrainDetailsChangeTicketDetails `json:"change_ticket_details,omitempty" xml:"change_ticket_details,omitempty" type:"Repeated"`
	DepStationCode      *string                                                                           `json:"dep_station_code,omitempty" xml:"dep_station_code,omitempty"`
	DepTime             *string                                                                           `json:"dep_time,omitempty" xml:"dep_time,omitempty"`
	TrainNo             *string                                                                           `json:"train_no,omitempty" xml:"train_no,omitempty"`
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
	if s.ChangeTicketDetails != nil {
		for _, item := range s.ChangeTicketDetails {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type TrainFeeCalculateChangeResponseBodyModuleChangeTrainDetailsChangeTicketDetails struct {
	ChangeFee        *int64                                                                                       `json:"change_fee,omitempty" xml:"change_fee,omitempty"`
	ChangeRate       *int64                                                                                       `json:"change_rate,omitempty" xml:"change_rate,omitempty"`
	ChangeRefundFee  *int64                                                                                       `json:"change_refund_fee,omitempty" xml:"change_refund_fee,omitempty"`
	ChangeRefundRate *int64                                                                                       `json:"change_refund_rate,omitempty" xml:"change_refund_rate,omitempty"`
	PassengerInfo    *TrainFeeCalculateChangeResponseBodyModuleChangeTrainDetailsChangeTicketDetailsPassengerInfo `json:"passenger_info,omitempty" xml:"passenger_info,omitempty" type:"Struct"`
	SeatType         *string                                                                                      `json:"seat_type,omitempty" xml:"seat_type,omitempty"`
	TicketPrice      *int64                                                                                       `json:"ticket_price,omitempty" xml:"ticket_price,omitempty"`
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
	if s.PassengerInfo != nil {
		if err := s.PassengerInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type TrainFeeCalculateChangeResponseBodyModuleChangeTrainDetailsChangeTicketDetailsPassengerInfo struct {
	PassengerCertNo   *string `json:"passenger_cert_no,omitempty" xml:"passenger_cert_no,omitempty"`
	PassengerCertType *string `json:"passenger_cert_type,omitempty" xml:"passenger_cert_type,omitempty"`
	PassengerId       *string `json:"passenger_id,omitempty" xml:"passenger_id,omitempty"`
	PassengerName     *string `json:"passenger_name,omitempty" xml:"passenger_name,omitempty"`
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
