// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTicketChangingFlightListShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetArrCity(v string) *TicketChangingFlightListShrinkRequest
	GetArrCity() *string
	SetDepCity(v string) *TicketChangingFlightListShrinkRequest
	GetDepCity() *string
	SetDepDate(v string) *TicketChangingFlightListShrinkRequest
	GetDepDate() *string
	SetDisOrderId(v string) *TicketChangingFlightListShrinkRequest
	GetDisOrderId() *string
	SetIsVoluntary(v int32) *TicketChangingFlightListShrinkRequest
	GetIsVoluntary() *int32
	SetTravelerInfoListShrink(v string) *TicketChangingFlightListShrinkRequest
	GetTravelerInfoListShrink() *string
}

type TicketChangingFlightListShrinkRequest struct {
	// example:
	//
	// BJS
	ArrCity *string `json:"arr_city,omitempty" xml:"arr_city,omitempty"`
	// example:
	//
	// HGH
	DepCity *string `json:"dep_city,omitempty" xml:"dep_city,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2000-00-00 00:00:00
	DepDate *string `json:"dep_date,omitempty" xml:"dep_date,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// dis123
	DisOrderId             *string `json:"dis_order_id,omitempty" xml:"dis_order_id,omitempty"`
	IsVoluntary            *int32  `json:"is_voluntary,omitempty" xml:"is_voluntary,omitempty"`
	TravelerInfoListShrink *string `json:"traveler_info_list,omitempty" xml:"traveler_info_list,omitempty"`
}

func (s TicketChangingFlightListShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s TicketChangingFlightListShrinkRequest) GoString() string {
	return s.String()
}

func (s *TicketChangingFlightListShrinkRequest) GetArrCity() *string {
	return s.ArrCity
}

func (s *TicketChangingFlightListShrinkRequest) GetDepCity() *string {
	return s.DepCity
}

func (s *TicketChangingFlightListShrinkRequest) GetDepDate() *string {
	return s.DepDate
}

func (s *TicketChangingFlightListShrinkRequest) GetDisOrderId() *string {
	return s.DisOrderId
}

func (s *TicketChangingFlightListShrinkRequest) GetIsVoluntary() *int32 {
	return s.IsVoluntary
}

func (s *TicketChangingFlightListShrinkRequest) GetTravelerInfoListShrink() *string {
	return s.TravelerInfoListShrink
}

func (s *TicketChangingFlightListShrinkRequest) SetArrCity(v string) *TicketChangingFlightListShrinkRequest {
	s.ArrCity = &v
	return s
}

func (s *TicketChangingFlightListShrinkRequest) SetDepCity(v string) *TicketChangingFlightListShrinkRequest {
	s.DepCity = &v
	return s
}

func (s *TicketChangingFlightListShrinkRequest) SetDepDate(v string) *TicketChangingFlightListShrinkRequest {
	s.DepDate = &v
	return s
}

func (s *TicketChangingFlightListShrinkRequest) SetDisOrderId(v string) *TicketChangingFlightListShrinkRequest {
	s.DisOrderId = &v
	return s
}

func (s *TicketChangingFlightListShrinkRequest) SetIsVoluntary(v int32) *TicketChangingFlightListShrinkRequest {
	s.IsVoluntary = &v
	return s
}

func (s *TicketChangingFlightListShrinkRequest) SetTravelerInfoListShrink(v string) *TicketChangingFlightListShrinkRequest {
	s.TravelerInfoListShrink = &v
	return s
}

func (s *TicketChangingFlightListShrinkRequest) Validate() error {
	return dara.Validate(s)
}
