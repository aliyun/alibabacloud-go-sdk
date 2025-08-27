// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTicketChangingFlightListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetArrCity(v string) *TicketChangingFlightListRequest
	GetArrCity() *string
	SetDepCity(v string) *TicketChangingFlightListRequest
	GetDepCity() *string
	SetDepDate(v string) *TicketChangingFlightListRequest
	GetDepDate() *string
	SetDisOrderId(v string) *TicketChangingFlightListRequest
	GetDisOrderId() *string
	SetIsVoluntary(v int32) *TicketChangingFlightListRequest
	GetIsVoluntary() *int32
	SetTravelerInfoList(v []*TicketChangingFlightListRequestTravelerInfoList) *TicketChangingFlightListRequest
	GetTravelerInfoList() []*TicketChangingFlightListRequestTravelerInfoList
}

type TicketChangingFlightListRequest struct {
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
	DisOrderId       *string                                            `json:"dis_order_id,omitempty" xml:"dis_order_id,omitempty"`
	IsVoluntary      *int32                                             `json:"is_voluntary,omitempty" xml:"is_voluntary,omitempty"`
	TravelerInfoList []*TicketChangingFlightListRequestTravelerInfoList `json:"traveler_info_list,omitempty" xml:"traveler_info_list,omitempty" type:"Repeated"`
}

func (s TicketChangingFlightListRequest) String() string {
	return dara.Prettify(s)
}

func (s TicketChangingFlightListRequest) GoString() string {
	return s.String()
}

func (s *TicketChangingFlightListRequest) GetArrCity() *string {
	return s.ArrCity
}

func (s *TicketChangingFlightListRequest) GetDepCity() *string {
	return s.DepCity
}

func (s *TicketChangingFlightListRequest) GetDepDate() *string {
	return s.DepDate
}

func (s *TicketChangingFlightListRequest) GetDisOrderId() *string {
	return s.DisOrderId
}

func (s *TicketChangingFlightListRequest) GetIsVoluntary() *int32 {
	return s.IsVoluntary
}

func (s *TicketChangingFlightListRequest) GetTravelerInfoList() []*TicketChangingFlightListRequestTravelerInfoList {
	return s.TravelerInfoList
}

func (s *TicketChangingFlightListRequest) SetArrCity(v string) *TicketChangingFlightListRequest {
	s.ArrCity = &v
	return s
}

func (s *TicketChangingFlightListRequest) SetDepCity(v string) *TicketChangingFlightListRequest {
	s.DepCity = &v
	return s
}

func (s *TicketChangingFlightListRequest) SetDepDate(v string) *TicketChangingFlightListRequest {
	s.DepDate = &v
	return s
}

func (s *TicketChangingFlightListRequest) SetDisOrderId(v string) *TicketChangingFlightListRequest {
	s.DisOrderId = &v
	return s
}

func (s *TicketChangingFlightListRequest) SetIsVoluntary(v int32) *TicketChangingFlightListRequest {
	s.IsVoluntary = &v
	return s
}

func (s *TicketChangingFlightListRequest) SetTravelerInfoList(v []*TicketChangingFlightListRequestTravelerInfoList) *TicketChangingFlightListRequest {
	s.TravelerInfoList = v
	return s
}

func (s *TicketChangingFlightListRequest) Validate() error {
	return dara.Validate(s)
}

type TicketChangingFlightListRequestTravelerInfoList struct {
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
	Name    *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// ADULT
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 23231
	UserId *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

func (s TicketChangingFlightListRequestTravelerInfoList) String() string {
	return dara.Prettify(s)
}

func (s TicketChangingFlightListRequestTravelerInfoList) GoString() string {
	return s.String()
}

func (s *TicketChangingFlightListRequestTravelerInfoList) GetArrCity() *string {
	return s.ArrCity
}

func (s *TicketChangingFlightListRequestTravelerInfoList) GetDepCity() *string {
	return s.DepCity
}

func (s *TicketChangingFlightListRequestTravelerInfoList) GetName() *string {
	return s.Name
}

func (s *TicketChangingFlightListRequestTravelerInfoList) GetType() *string {
	return s.Type
}

func (s *TicketChangingFlightListRequestTravelerInfoList) GetUserId() *string {
	return s.UserId
}

func (s *TicketChangingFlightListRequestTravelerInfoList) SetArrCity(v string) *TicketChangingFlightListRequestTravelerInfoList {
	s.ArrCity = &v
	return s
}

func (s *TicketChangingFlightListRequestTravelerInfoList) SetDepCity(v string) *TicketChangingFlightListRequestTravelerInfoList {
	s.DepCity = &v
	return s
}

func (s *TicketChangingFlightListRequestTravelerInfoList) SetName(v string) *TicketChangingFlightListRequestTravelerInfoList {
	s.Name = &v
	return s
}

func (s *TicketChangingFlightListRequestTravelerInfoList) SetType(v string) *TicketChangingFlightListRequestTravelerInfoList {
	s.Type = &v
	return s
}

func (s *TicketChangingFlightListRequestTravelerInfoList) SetUserId(v string) *TicketChangingFlightListRequestTravelerInfoList {
	s.UserId = &v
	return s
}

func (s *TicketChangingFlightListRequestTravelerInfoList) Validate() error {
	return dara.Validate(s)
}
