// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTicketRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupKey(v string) *UpdateTicketRequest
	GetGroupKey() *string
	SetHotelId(v string) *UpdateTicketRequest
	GetHotelId() *string
	SetStatus(v string) *UpdateTicketRequest
	GetStatus() *string
}

type UpdateTicketRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 2023***93975
	GroupKey *string `json:"GroupKey,omitempty" xml:"GroupKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 520a0c0***5eb
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// waiting
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UpdateTicketRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateTicketRequest) GoString() string {
	return s.String()
}

func (s *UpdateTicketRequest) GetGroupKey() *string {
	return s.GroupKey
}

func (s *UpdateTicketRequest) GetHotelId() *string {
	return s.HotelId
}

func (s *UpdateTicketRequest) GetStatus() *string {
	return s.Status
}

func (s *UpdateTicketRequest) SetGroupKey(v string) *UpdateTicketRequest {
	s.GroupKey = &v
	return s
}

func (s *UpdateTicketRequest) SetHotelId(v string) *UpdateTicketRequest {
	s.HotelId = &v
	return s
}

func (s *UpdateTicketRequest) SetStatus(v string) *UpdateTicketRequest {
	s.Status = &v
	return s
}

func (s *UpdateTicketRequest) Validate() error {
	return dara.Validate(s)
}
