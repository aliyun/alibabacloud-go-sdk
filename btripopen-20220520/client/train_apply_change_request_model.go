// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTrainApplyChangeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptNoSeat(v string) *TrainApplyChangeRequest
	GetAcceptNoSeat() *string
	SetChangeTrainInfoS(v []*TrainApplyChangeRequestChangeTrainInfoS) *TrainApplyChangeRequest
	GetChangeTrainInfoS() []*TrainApplyChangeRequestChangeTrainInfoS
	SetForceMatch(v string) *TrainApplyChangeRequest
	GetForceMatch() *string
	SetIsPayNow(v bool) *TrainApplyChangeRequest
	GetIsPayNow() *bool
	SetOrderId(v string) *TrainApplyChangeRequest
	GetOrderId() *string
	SetOutChangeApplyId(v string) *TrainApplyChangeRequest
	GetOutChangeApplyId() *string
	SetOutOrderId(v string) *TrainApplyChangeRequest
	GetOutOrderId() *string
}

type TrainApplyChangeRequest struct {
	// example:
	//
	// 0
	AcceptNoSeat *string `json:"accept_no_seat,omitempty" xml:"accept_no_seat,omitempty"`
	// This parameter is required.
	ChangeTrainInfoS []*TrainApplyChangeRequestChangeTrainInfoS `json:"change_train_info_s,omitempty" xml:"change_train_info_s,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	ForceMatch *string `json:"force_match,omitempty" xml:"force_match,omitempty"`
	// example:
	//
	// false
	IsPayNow *bool `json:"is_pay_now,omitempty" xml:"is_pay_now,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1017028198411054446
	OrderId *string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 12345
	OutChangeApplyId *string `json:"out_change_apply_id,omitempty" xml:"out_change_apply_id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123456
	OutOrderId *string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
}

func (s TrainApplyChangeRequest) String() string {
	return dara.Prettify(s)
}

func (s TrainApplyChangeRequest) GoString() string {
	return s.String()
}

func (s *TrainApplyChangeRequest) GetAcceptNoSeat() *string {
	return s.AcceptNoSeat
}

func (s *TrainApplyChangeRequest) GetChangeTrainInfoS() []*TrainApplyChangeRequestChangeTrainInfoS {
	return s.ChangeTrainInfoS
}

func (s *TrainApplyChangeRequest) GetForceMatch() *string {
	return s.ForceMatch
}

func (s *TrainApplyChangeRequest) GetIsPayNow() *bool {
	return s.IsPayNow
}

func (s *TrainApplyChangeRequest) GetOrderId() *string {
	return s.OrderId
}

func (s *TrainApplyChangeRequest) GetOutChangeApplyId() *string {
	return s.OutChangeApplyId
}

func (s *TrainApplyChangeRequest) GetOutOrderId() *string {
	return s.OutOrderId
}

func (s *TrainApplyChangeRequest) SetAcceptNoSeat(v string) *TrainApplyChangeRequest {
	s.AcceptNoSeat = &v
	return s
}

func (s *TrainApplyChangeRequest) SetChangeTrainInfoS(v []*TrainApplyChangeRequestChangeTrainInfoS) *TrainApplyChangeRequest {
	s.ChangeTrainInfoS = v
	return s
}

func (s *TrainApplyChangeRequest) SetForceMatch(v string) *TrainApplyChangeRequest {
	s.ForceMatch = &v
	return s
}

func (s *TrainApplyChangeRequest) SetIsPayNow(v bool) *TrainApplyChangeRequest {
	s.IsPayNow = &v
	return s
}

func (s *TrainApplyChangeRequest) SetOrderId(v string) *TrainApplyChangeRequest {
	s.OrderId = &v
	return s
}

func (s *TrainApplyChangeRequest) SetOutChangeApplyId(v string) *TrainApplyChangeRequest {
	s.OutChangeApplyId = &v
	return s
}

func (s *TrainApplyChangeRequest) SetOutOrderId(v string) *TrainApplyChangeRequest {
	s.OutOrderId = &v
	return s
}

func (s *TrainApplyChangeRequest) Validate() error {
	return dara.Validate(s)
}

type TrainApplyChangeRequestChangeTrainInfoS struct {
	// This parameter is required.
	//
	// example:
	//
	// BTC
	ArrStationCode *string `json:"arr_station_code,omitempty" xml:"arr_station_code,omitempty"`
	// This parameter is required.
	ChangeTicketInfoS []*TrainApplyChangeRequestChangeTrainInfoSChangeTicketInfoS `json:"change_ticket_info_s,omitempty" xml:"change_ticket_info_s,omitempty" type:"Repeated"`
	// example:
	//
	// null
	ChooseBedS *string `json:"choose_bed_s,omitempty" xml:"choose_bed_s,omitempty"`
	// example:
	//
	// 1T
	ChooseSeatS *string `json:"choose_seat_s,omitempty" xml:"choose_seat_s,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// BDC
	DepStationCode *string `json:"dep_station_code,omitempty" xml:"dep_station_code,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2024-05-06 15:19:01
	DepTime *string `json:"dep_time,omitempty" xml:"dep_time,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2024-05-06 15:19:01
	OriginalDepTime *string `json:"original_dep_time,omitempty" xml:"original_dep_time,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// K234
	OriginalTrainNo *string `json:"original_train_no,omitempty" xml:"original_train_no,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// K2345
	TrainNo *string `json:"train_no,omitempty" xml:"train_no,omitempty"`
}

func (s TrainApplyChangeRequestChangeTrainInfoS) String() string {
	return dara.Prettify(s)
}

func (s TrainApplyChangeRequestChangeTrainInfoS) GoString() string {
	return s.String()
}

func (s *TrainApplyChangeRequestChangeTrainInfoS) GetArrStationCode() *string {
	return s.ArrStationCode
}

func (s *TrainApplyChangeRequestChangeTrainInfoS) GetChangeTicketInfoS() []*TrainApplyChangeRequestChangeTrainInfoSChangeTicketInfoS {
	return s.ChangeTicketInfoS
}

func (s *TrainApplyChangeRequestChangeTrainInfoS) GetChooseBedS() *string {
	return s.ChooseBedS
}

func (s *TrainApplyChangeRequestChangeTrainInfoS) GetChooseSeatS() *string {
	return s.ChooseSeatS
}

func (s *TrainApplyChangeRequestChangeTrainInfoS) GetDepStationCode() *string {
	return s.DepStationCode
}

func (s *TrainApplyChangeRequestChangeTrainInfoS) GetDepTime() *string {
	return s.DepTime
}

func (s *TrainApplyChangeRequestChangeTrainInfoS) GetOriginalDepTime() *string {
	return s.OriginalDepTime
}

func (s *TrainApplyChangeRequestChangeTrainInfoS) GetOriginalTrainNo() *string {
	return s.OriginalTrainNo
}

func (s *TrainApplyChangeRequestChangeTrainInfoS) GetTrainNo() *string {
	return s.TrainNo
}

func (s *TrainApplyChangeRequestChangeTrainInfoS) SetArrStationCode(v string) *TrainApplyChangeRequestChangeTrainInfoS {
	s.ArrStationCode = &v
	return s
}

func (s *TrainApplyChangeRequestChangeTrainInfoS) SetChangeTicketInfoS(v []*TrainApplyChangeRequestChangeTrainInfoSChangeTicketInfoS) *TrainApplyChangeRequestChangeTrainInfoS {
	s.ChangeTicketInfoS = v
	return s
}

func (s *TrainApplyChangeRequestChangeTrainInfoS) SetChooseBedS(v string) *TrainApplyChangeRequestChangeTrainInfoS {
	s.ChooseBedS = &v
	return s
}

func (s *TrainApplyChangeRequestChangeTrainInfoS) SetChooseSeatS(v string) *TrainApplyChangeRequestChangeTrainInfoS {
	s.ChooseSeatS = &v
	return s
}

func (s *TrainApplyChangeRequestChangeTrainInfoS) SetDepStationCode(v string) *TrainApplyChangeRequestChangeTrainInfoS {
	s.DepStationCode = &v
	return s
}

func (s *TrainApplyChangeRequestChangeTrainInfoS) SetDepTime(v string) *TrainApplyChangeRequestChangeTrainInfoS {
	s.DepTime = &v
	return s
}

func (s *TrainApplyChangeRequestChangeTrainInfoS) SetOriginalDepTime(v string) *TrainApplyChangeRequestChangeTrainInfoS {
	s.OriginalDepTime = &v
	return s
}

func (s *TrainApplyChangeRequestChangeTrainInfoS) SetOriginalTrainNo(v string) *TrainApplyChangeRequestChangeTrainInfoS {
	s.OriginalTrainNo = &v
	return s
}

func (s *TrainApplyChangeRequestChangeTrainInfoS) SetTrainNo(v string) *TrainApplyChangeRequestChangeTrainInfoS {
	s.TrainNo = &v
	return s
}

func (s *TrainApplyChangeRequestChangeTrainInfoS) Validate() error {
	return dara.Validate(s)
}

type TrainApplyChangeRequestChangeTrainInfoSChangeTicketInfoS struct {
	// This parameter is required.
	PassengerInfo *TrainApplyChangeRequestChangeTrainInfoSChangeTicketInfoSPassengerInfo `json:"passenger_info,omitempty" xml:"passenger_info,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 14
	SeatType *string `json:"seat_type,omitempty" xml:"seat_type,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 100
	TicketPrice *string `json:"ticket_price,omitempty" xml:"ticket_price,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0
	TicketType *string `json:"ticket_type,omitempty" xml:"ticket_type,omitempty"`
}

func (s TrainApplyChangeRequestChangeTrainInfoSChangeTicketInfoS) String() string {
	return dara.Prettify(s)
}

func (s TrainApplyChangeRequestChangeTrainInfoSChangeTicketInfoS) GoString() string {
	return s.String()
}

func (s *TrainApplyChangeRequestChangeTrainInfoSChangeTicketInfoS) GetPassengerInfo() *TrainApplyChangeRequestChangeTrainInfoSChangeTicketInfoSPassengerInfo {
	return s.PassengerInfo
}

func (s *TrainApplyChangeRequestChangeTrainInfoSChangeTicketInfoS) GetSeatType() *string {
	return s.SeatType
}

func (s *TrainApplyChangeRequestChangeTrainInfoSChangeTicketInfoS) GetTicketPrice() *string {
	return s.TicketPrice
}

func (s *TrainApplyChangeRequestChangeTrainInfoSChangeTicketInfoS) GetTicketType() *string {
	return s.TicketType
}

func (s *TrainApplyChangeRequestChangeTrainInfoSChangeTicketInfoS) SetPassengerInfo(v *TrainApplyChangeRequestChangeTrainInfoSChangeTicketInfoSPassengerInfo) *TrainApplyChangeRequestChangeTrainInfoSChangeTicketInfoS {
	s.PassengerInfo = v
	return s
}

func (s *TrainApplyChangeRequestChangeTrainInfoSChangeTicketInfoS) SetSeatType(v string) *TrainApplyChangeRequestChangeTrainInfoSChangeTicketInfoS {
	s.SeatType = &v
	return s
}

func (s *TrainApplyChangeRequestChangeTrainInfoSChangeTicketInfoS) SetTicketPrice(v string) *TrainApplyChangeRequestChangeTrainInfoSChangeTicketInfoS {
	s.TicketPrice = &v
	return s
}

func (s *TrainApplyChangeRequestChangeTrainInfoSChangeTicketInfoS) SetTicketType(v string) *TrainApplyChangeRequestChangeTrainInfoSChangeTicketInfoS {
	s.TicketType = &v
	return s
}

func (s *TrainApplyChangeRequestChangeTrainInfoSChangeTicketInfoS) Validate() error {
	return dara.Validate(s)
}

type TrainApplyChangeRequestChangeTrainInfoSChangeTicketInfoSPassengerInfo struct {
	// This parameter is required.
	//
	// example:
	//
	// 291487e553c5abde3b611aae283e2526f0d733ab55094aadc0b5ba587222a233c
	PassengerCertNo *string `json:"passenger_cert_no,omitempty" xml:"passenger_cert_no,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 170d9ac6f8807f9ec603c688f45f78a41
	PassengerCertType *string `json:"passenger_cert_type,omitempty" xml:"passenger_cert_type,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123456
	PassengerId *string `json:"passenger_id,omitempty" xml:"passenger_id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 949c9f34f677a0e5d249dfc94f5e62cc7
	PassengerName *string `json:"passenger_name,omitempty" xml:"passenger_name,omitempty"`
}

func (s TrainApplyChangeRequestChangeTrainInfoSChangeTicketInfoSPassengerInfo) String() string {
	return dara.Prettify(s)
}

func (s TrainApplyChangeRequestChangeTrainInfoSChangeTicketInfoSPassengerInfo) GoString() string {
	return s.String()
}

func (s *TrainApplyChangeRequestChangeTrainInfoSChangeTicketInfoSPassengerInfo) GetPassengerCertNo() *string {
	return s.PassengerCertNo
}

func (s *TrainApplyChangeRequestChangeTrainInfoSChangeTicketInfoSPassengerInfo) GetPassengerCertType() *string {
	return s.PassengerCertType
}

func (s *TrainApplyChangeRequestChangeTrainInfoSChangeTicketInfoSPassengerInfo) GetPassengerId() *string {
	return s.PassengerId
}

func (s *TrainApplyChangeRequestChangeTrainInfoSChangeTicketInfoSPassengerInfo) GetPassengerName() *string {
	return s.PassengerName
}

func (s *TrainApplyChangeRequestChangeTrainInfoSChangeTicketInfoSPassengerInfo) SetPassengerCertNo(v string) *TrainApplyChangeRequestChangeTrainInfoSChangeTicketInfoSPassengerInfo {
	s.PassengerCertNo = &v
	return s
}

func (s *TrainApplyChangeRequestChangeTrainInfoSChangeTicketInfoSPassengerInfo) SetPassengerCertType(v string) *TrainApplyChangeRequestChangeTrainInfoSChangeTicketInfoSPassengerInfo {
	s.PassengerCertType = &v
	return s
}

func (s *TrainApplyChangeRequestChangeTrainInfoSChangeTicketInfoSPassengerInfo) SetPassengerId(v string) *TrainApplyChangeRequestChangeTrainInfoSChangeTicketInfoSPassengerInfo {
	s.PassengerId = &v
	return s
}

func (s *TrainApplyChangeRequestChangeTrainInfoSChangeTicketInfoSPassengerInfo) SetPassengerName(v string) *TrainApplyChangeRequestChangeTrainInfoSChangeTicketInfoSPassengerInfo {
	s.PassengerName = &v
	return s
}

func (s *TrainApplyChangeRequestChangeTrainInfoSChangeTicketInfoSPassengerInfo) Validate() error {
	return dara.Validate(s)
}
