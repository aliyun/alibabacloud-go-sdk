// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightRefundPreCalRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDisOrderId(v string) *FlightRefundPreCalRequest
	GetDisOrderId() *string
	SetIsVoluntary(v string) *FlightRefundPreCalRequest
	GetIsVoluntary() *string
	SetPassengerSegmentInfoList(v []*FlightRefundPreCalRequestPassengerSegmentInfoList) *FlightRefundPreCalRequest
	GetPassengerSegmentInfoList() []*FlightRefundPreCalRequestPassengerSegmentInfoList
}

type FlightRefundPreCalRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// dis123
	DisOrderId *string `json:"dis_order_id,omitempty" xml:"dis_order_id,omitempty"`
	// example:
	//
	// 1
	IsVoluntary *string `json:"is_voluntary,omitempty" xml:"is_voluntary,omitempty"`
	// This parameter is required.
	PassengerSegmentInfoList []*FlightRefundPreCalRequestPassengerSegmentInfoList `json:"passenger_segment_info_list,omitempty" xml:"passenger_segment_info_list,omitempty" type:"Repeated"`
}

func (s FlightRefundPreCalRequest) String() string {
	return dara.Prettify(s)
}

func (s FlightRefundPreCalRequest) GoString() string {
	return s.String()
}

func (s *FlightRefundPreCalRequest) GetDisOrderId() *string {
	return s.DisOrderId
}

func (s *FlightRefundPreCalRequest) GetIsVoluntary() *string {
	return s.IsVoluntary
}

func (s *FlightRefundPreCalRequest) GetPassengerSegmentInfoList() []*FlightRefundPreCalRequestPassengerSegmentInfoList {
	return s.PassengerSegmentInfoList
}

func (s *FlightRefundPreCalRequest) SetDisOrderId(v string) *FlightRefundPreCalRequest {
	s.DisOrderId = &v
	return s
}

func (s *FlightRefundPreCalRequest) SetIsVoluntary(v string) *FlightRefundPreCalRequest {
	s.IsVoluntary = &v
	return s
}

func (s *FlightRefundPreCalRequest) SetPassengerSegmentInfoList(v []*FlightRefundPreCalRequestPassengerSegmentInfoList) *FlightRefundPreCalRequest {
	s.PassengerSegmentInfoList = v
	return s
}

func (s *FlightRefundPreCalRequest) Validate() error {
	return dara.Validate(s)
}

type FlightRefundPreCalRequestPassengerSegmentInfoList struct {
	// This parameter is required.
	//
	// example:
	//
	// CA1982
	FlightNo *string `json:"flight_no,omitempty" xml:"flight_no,omitempty"`
	// This parameter is required.
	PassengerName *string `json:"passenger_name,omitempty" xml:"passenger_name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 23112
	UserId *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

func (s FlightRefundPreCalRequestPassengerSegmentInfoList) String() string {
	return dara.Prettify(s)
}

func (s FlightRefundPreCalRequestPassengerSegmentInfoList) GoString() string {
	return s.String()
}

func (s *FlightRefundPreCalRequestPassengerSegmentInfoList) GetFlightNo() *string {
	return s.FlightNo
}

func (s *FlightRefundPreCalRequestPassengerSegmentInfoList) GetPassengerName() *string {
	return s.PassengerName
}

func (s *FlightRefundPreCalRequestPassengerSegmentInfoList) GetUserId() *string {
	return s.UserId
}

func (s *FlightRefundPreCalRequestPassengerSegmentInfoList) SetFlightNo(v string) *FlightRefundPreCalRequestPassengerSegmentInfoList {
	s.FlightNo = &v
	return s
}

func (s *FlightRefundPreCalRequestPassengerSegmentInfoList) SetPassengerName(v string) *FlightRefundPreCalRequestPassengerSegmentInfoList {
	s.PassengerName = &v
	return s
}

func (s *FlightRefundPreCalRequestPassengerSegmentInfoList) SetUserId(v string) *FlightRefundPreCalRequestPassengerSegmentInfoList {
	s.UserId = &v
	return s
}

func (s *FlightRefundPreCalRequestPassengerSegmentInfoList) Validate() error {
	return dara.Validate(s)
}
