// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTicketChangingEnquiryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetArrCity(v string) *TicketChangingEnquiryRequest
	GetArrCity() *string
	SetDepCity(v string) *TicketChangingEnquiryRequest
	GetDepCity() *string
	SetDisOrderId(v string) *TicketChangingEnquiryRequest
	GetDisOrderId() *string
	SetIsVoluntary(v int32) *TicketChangingEnquiryRequest
	GetIsVoluntary() *int32
	SetModifyDepartDate(v string) *TicketChangingEnquiryRequest
	GetModifyDepartDate() *string
	SetModifyFlightNo(v string) *TicketChangingEnquiryRequest
	GetModifyFlightNo() *string
	SetSessionId(v string) *TicketChangingEnquiryRequest
	GetSessionId() *string
}

type TicketChangingEnquiryRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// BJS
	ArrCity *string `json:"arr_city,omitempty" xml:"arr_city,omitempty"`
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
	// dis123
	DisOrderId *string `json:"dis_order_id,omitempty" xml:"dis_order_id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	IsVoluntary *int32 `json:"is_voluntary,omitempty" xml:"is_voluntary,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2000-00-00 00:00:00
	ModifyDepartDate *string `json:"modify_depart_date,omitempty" xml:"modify_depart_date,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// CA1704
	ModifyFlightNo *string `json:"modify_flight_no,omitempty" xml:"modify_flight_no,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ece19e8b1047898a5a98b6487348c2
	SessionId *string `json:"session_id,omitempty" xml:"session_id,omitempty"`
}

func (s TicketChangingEnquiryRequest) String() string {
	return dara.Prettify(s)
}

func (s TicketChangingEnquiryRequest) GoString() string {
	return s.String()
}

func (s *TicketChangingEnquiryRequest) GetArrCity() *string {
	return s.ArrCity
}

func (s *TicketChangingEnquiryRequest) GetDepCity() *string {
	return s.DepCity
}

func (s *TicketChangingEnquiryRequest) GetDisOrderId() *string {
	return s.DisOrderId
}

func (s *TicketChangingEnquiryRequest) GetIsVoluntary() *int32 {
	return s.IsVoluntary
}

func (s *TicketChangingEnquiryRequest) GetModifyDepartDate() *string {
	return s.ModifyDepartDate
}

func (s *TicketChangingEnquiryRequest) GetModifyFlightNo() *string {
	return s.ModifyFlightNo
}

func (s *TicketChangingEnquiryRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *TicketChangingEnquiryRequest) SetArrCity(v string) *TicketChangingEnquiryRequest {
	s.ArrCity = &v
	return s
}

func (s *TicketChangingEnquiryRequest) SetDepCity(v string) *TicketChangingEnquiryRequest {
	s.DepCity = &v
	return s
}

func (s *TicketChangingEnquiryRequest) SetDisOrderId(v string) *TicketChangingEnquiryRequest {
	s.DisOrderId = &v
	return s
}

func (s *TicketChangingEnquiryRequest) SetIsVoluntary(v int32) *TicketChangingEnquiryRequest {
	s.IsVoluntary = &v
	return s
}

func (s *TicketChangingEnquiryRequest) SetModifyDepartDate(v string) *TicketChangingEnquiryRequest {
	s.ModifyDepartDate = &v
	return s
}

func (s *TicketChangingEnquiryRequest) SetModifyFlightNo(v string) *TicketChangingEnquiryRequest {
	s.ModifyFlightNo = &v
	return s
}

func (s *TicketChangingEnquiryRequest) SetSessionId(v string) *TicketChangingEnquiryRequest {
	s.SessionId = &v
	return s
}

func (s *TicketChangingEnquiryRequest) Validate() error {
	return dara.Validate(s)
}
