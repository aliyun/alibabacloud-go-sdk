// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTicketChangingApplyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDisOrderId(v string) *TicketChangingApplyRequest
	GetDisOrderId() *string
	SetDisSubOrderId(v string) *TicketChangingApplyRequest
	GetDisSubOrderId() *string
	SetIsVoluntary(v int32) *TicketChangingApplyRequest
	GetIsVoluntary() *int32
	SetModifyFlightInfoList(v []*TicketChangingApplyRequestModifyFlightInfoList) *TicketChangingApplyRequest
	GetModifyFlightInfoList() []*TicketChangingApplyRequestModifyFlightInfoList
	SetOtaItemId(v string) *TicketChangingApplyRequest
	GetOtaItemId() *string
	SetReason(v string) *TicketChangingApplyRequest
	GetReason() *string
	SetSessionId(v string) *TicketChangingApplyRequest
	GetSessionId() *string
	SetWhetherRetry(v bool) *TicketChangingApplyRequest
	GetWhetherRetry() *bool
}

type TicketChangingApplyRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// dis1234
	DisOrderId *string `json:"dis_order_id,omitempty" xml:"dis_order_id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// mid1243
	DisSubOrderId *string `json:"dis_sub_order_id,omitempty" xml:"dis_sub_order_id,omitempty"`
	IsVoluntary   *int32  `json:"is_voluntary,omitempty" xml:"is_voluntary,omitempty"`
	// This parameter is required.
	ModifyFlightInfoList []*TicketChangingApplyRequestModifyFlightInfoList `json:"modify_flight_info_list,omitempty" xml:"modify_flight_info_list,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 1s8837sh991hsj92h
	OtaItemId *string `json:"ota_item_id,omitempty" xml:"ota_item_id,omitempty"`
	Reason    *string `json:"reason,omitempty" xml:"reason,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// us88s2bsbin22hjusd8i
	SessionId *string `json:"session_id,omitempty" xml:"session_id,omitempty"`
	// example:
	//
	// true
	WhetherRetry *bool `json:"whether_retry,omitempty" xml:"whether_retry,omitempty"`
}

func (s TicketChangingApplyRequest) String() string {
	return dara.Prettify(s)
}

func (s TicketChangingApplyRequest) GoString() string {
	return s.String()
}

func (s *TicketChangingApplyRequest) GetDisOrderId() *string {
	return s.DisOrderId
}

func (s *TicketChangingApplyRequest) GetDisSubOrderId() *string {
	return s.DisSubOrderId
}

func (s *TicketChangingApplyRequest) GetIsVoluntary() *int32 {
	return s.IsVoluntary
}

func (s *TicketChangingApplyRequest) GetModifyFlightInfoList() []*TicketChangingApplyRequestModifyFlightInfoList {
	return s.ModifyFlightInfoList
}

func (s *TicketChangingApplyRequest) GetOtaItemId() *string {
	return s.OtaItemId
}

func (s *TicketChangingApplyRequest) GetReason() *string {
	return s.Reason
}

func (s *TicketChangingApplyRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *TicketChangingApplyRequest) GetWhetherRetry() *bool {
	return s.WhetherRetry
}

func (s *TicketChangingApplyRequest) SetDisOrderId(v string) *TicketChangingApplyRequest {
	s.DisOrderId = &v
	return s
}

func (s *TicketChangingApplyRequest) SetDisSubOrderId(v string) *TicketChangingApplyRequest {
	s.DisSubOrderId = &v
	return s
}

func (s *TicketChangingApplyRequest) SetIsVoluntary(v int32) *TicketChangingApplyRequest {
	s.IsVoluntary = &v
	return s
}

func (s *TicketChangingApplyRequest) SetModifyFlightInfoList(v []*TicketChangingApplyRequestModifyFlightInfoList) *TicketChangingApplyRequest {
	s.ModifyFlightInfoList = v
	return s
}

func (s *TicketChangingApplyRequest) SetOtaItemId(v string) *TicketChangingApplyRequest {
	s.OtaItemId = &v
	return s
}

func (s *TicketChangingApplyRequest) SetReason(v string) *TicketChangingApplyRequest {
	s.Reason = &v
	return s
}

func (s *TicketChangingApplyRequest) SetSessionId(v string) *TicketChangingApplyRequest {
	s.SessionId = &v
	return s
}

func (s *TicketChangingApplyRequest) SetWhetherRetry(v bool) *TicketChangingApplyRequest {
	s.WhetherRetry = &v
	return s
}

func (s *TicketChangingApplyRequest) Validate() error {
	return dara.Validate(s)
}

type TicketChangingApplyRequestModifyFlightInfoList struct {
	// This parameter is required.
	//
	// example:
	//
	// BJS
	ArrCity *string `json:"arr_city,omitempty" xml:"arr_city,omitempty"`
	// example:
	//
	// Y
	Cabin *string `json:"cabin,omitempty" xml:"cabin,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// HGH
	DepCity *string `json:"dep_city,omitempty" xml:"dep_city,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0000-00-00 00:00:00
	DepDate *string `json:"dep_date,omitempty" xml:"dep_date,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// CA1704
	FlightNo *string `json:"flight_no,omitempty" xml:"flight_no,omitempty"`
	// This parameter is required.
	PassengerInfoList []*TicketChangingApplyRequestModifyFlightInfoListPassengerInfoList `json:"passenger_info_list,omitempty" xml:"passenger_info_list,omitempty" type:"Repeated"`
}

func (s TicketChangingApplyRequestModifyFlightInfoList) String() string {
	return dara.Prettify(s)
}

func (s TicketChangingApplyRequestModifyFlightInfoList) GoString() string {
	return s.String()
}

func (s *TicketChangingApplyRequestModifyFlightInfoList) GetArrCity() *string {
	return s.ArrCity
}

func (s *TicketChangingApplyRequestModifyFlightInfoList) GetCabin() *string {
	return s.Cabin
}

func (s *TicketChangingApplyRequestModifyFlightInfoList) GetDepCity() *string {
	return s.DepCity
}

func (s *TicketChangingApplyRequestModifyFlightInfoList) GetDepDate() *string {
	return s.DepDate
}

func (s *TicketChangingApplyRequestModifyFlightInfoList) GetFlightNo() *string {
	return s.FlightNo
}

func (s *TicketChangingApplyRequestModifyFlightInfoList) GetPassengerInfoList() []*TicketChangingApplyRequestModifyFlightInfoListPassengerInfoList {
	return s.PassengerInfoList
}

func (s *TicketChangingApplyRequestModifyFlightInfoList) SetArrCity(v string) *TicketChangingApplyRequestModifyFlightInfoList {
	s.ArrCity = &v
	return s
}

func (s *TicketChangingApplyRequestModifyFlightInfoList) SetCabin(v string) *TicketChangingApplyRequestModifyFlightInfoList {
	s.Cabin = &v
	return s
}

func (s *TicketChangingApplyRequestModifyFlightInfoList) SetDepCity(v string) *TicketChangingApplyRequestModifyFlightInfoList {
	s.DepCity = &v
	return s
}

func (s *TicketChangingApplyRequestModifyFlightInfoList) SetDepDate(v string) *TicketChangingApplyRequestModifyFlightInfoList {
	s.DepDate = &v
	return s
}

func (s *TicketChangingApplyRequestModifyFlightInfoList) SetFlightNo(v string) *TicketChangingApplyRequestModifyFlightInfoList {
	s.FlightNo = &v
	return s
}

func (s *TicketChangingApplyRequestModifyFlightInfoList) SetPassengerInfoList(v []*TicketChangingApplyRequestModifyFlightInfoListPassengerInfoList) *TicketChangingApplyRequestModifyFlightInfoList {
	s.PassengerInfoList = v
	return s
}

func (s *TicketChangingApplyRequestModifyFlightInfoList) Validate() error {
	return dara.Validate(s)
}

type TicketChangingApplyRequestModifyFlightInfoListPassengerInfoList struct {
	// This parameter is required.
	//
	// example:
	//
	// CA1703
	OriginFlightNo *string `json:"origin_flight_no,omitempty" xml:"origin_flight_no,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1234
	OutUserId *string `json:"out_user_id,omitempty" xml:"out_user_id,omitempty"`
	// This parameter is required.
	PassengerName *string `json:"passenger_name,omitempty" xml:"passenger_name,omitempty"`
}

func (s TicketChangingApplyRequestModifyFlightInfoListPassengerInfoList) String() string {
	return dara.Prettify(s)
}

func (s TicketChangingApplyRequestModifyFlightInfoListPassengerInfoList) GoString() string {
	return s.String()
}

func (s *TicketChangingApplyRequestModifyFlightInfoListPassengerInfoList) GetOriginFlightNo() *string {
	return s.OriginFlightNo
}

func (s *TicketChangingApplyRequestModifyFlightInfoListPassengerInfoList) GetOutUserId() *string {
	return s.OutUserId
}

func (s *TicketChangingApplyRequestModifyFlightInfoListPassengerInfoList) GetPassengerName() *string {
	return s.PassengerName
}

func (s *TicketChangingApplyRequestModifyFlightInfoListPassengerInfoList) SetOriginFlightNo(v string) *TicketChangingApplyRequestModifyFlightInfoListPassengerInfoList {
	s.OriginFlightNo = &v
	return s
}

func (s *TicketChangingApplyRequestModifyFlightInfoListPassengerInfoList) SetOutUserId(v string) *TicketChangingApplyRequestModifyFlightInfoListPassengerInfoList {
	s.OutUserId = &v
	return s
}

func (s *TicketChangingApplyRequestModifyFlightInfoListPassengerInfoList) SetPassengerName(v string) *TicketChangingApplyRequestModifyFlightInfoListPassengerInfoList {
	s.PassengerName = &v
	return s
}

func (s *TicketChangingApplyRequestModifyFlightInfoListPassengerInfoList) Validate() error {
	return dara.Validate(s)
}
