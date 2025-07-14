// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryTicketInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTicket(v string) *QueryTicketInfoRequest
	GetTicket() *string
}

type QueryTicketInfoRequest struct {
	// The value of the bill.
	//
	// This parameter is required.
	//
	// example:
	//
	// a27a9aec-****-****-bd40-1a21ea41d7c5
	Ticket *string `json:"Ticket,omitempty" xml:"Ticket,omitempty"`
}

func (s QueryTicketInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryTicketInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryTicketInfoRequest) GetTicket() *string {
	return s.Ticket
}

func (s *QueryTicketInfoRequest) SetTicket(v string) *QueryTicketInfoRequest {
	s.Ticket = &v
	return s
}

func (s *QueryTicketInfoRequest) Validate() error {
	return dara.Validate(s)
}
