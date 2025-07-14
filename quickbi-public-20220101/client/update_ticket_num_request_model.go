// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTicketNumRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTicket(v string) *UpdateTicketNumRequest
	GetTicket() *string
	SetTicketNum(v int32) *UpdateTicketNumRequest
	GetTicketNum() *int32
}

type UpdateTicketNumRequest struct {
	// The value of the third-party embedded ticket, that is, the accessTicket value in the URL.
	//
	// This parameter is required.
	//
	// example:
	//
	// 040e6f79d33444838***83c7206c070
	Ticket *string `json:"Ticket,omitempty" xml:"Ticket,omitempty"`
	// The number of bills.
	//
	// 	- Valid values: 1 to 99998. Recommended value: 1.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	TicketNum *int32 `json:"TicketNum,omitempty" xml:"TicketNum,omitempty"`
}

func (s UpdateTicketNumRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateTicketNumRequest) GoString() string {
	return s.String()
}

func (s *UpdateTicketNumRequest) GetTicket() *string {
	return s.Ticket
}

func (s *UpdateTicketNumRequest) GetTicketNum() *int32 {
	return s.TicketNum
}

func (s *UpdateTicketNumRequest) SetTicket(v string) *UpdateTicketNumRequest {
	s.Ticket = &v
	return s
}

func (s *UpdateTicketNumRequest) SetTicketNum(v int32) *UpdateTicketNumRequest {
	s.TicketNum = &v
	return s
}

func (s *UpdateTicketNumRequest) Validate() error {
	return dara.Validate(s)
}
