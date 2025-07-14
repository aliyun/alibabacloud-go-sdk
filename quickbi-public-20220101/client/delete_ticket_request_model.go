// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTicketRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTicket(v string) *DeleteTicketRequest
	GetTicket() *string
}

type DeleteTicketRequest struct {
	// The value of the third-party embedded ticket, which is the `accessTicket` in the URL.
	//
	// This parameter is required.
	//
	// example:
	//
	// 040e6f79d****7d283c7206c070
	Ticket *string `json:"Ticket,omitempty" xml:"Ticket,omitempty"`
}

func (s DeleteTicketRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteTicketRequest) GoString() string {
	return s.String()
}

func (s *DeleteTicketRequest) GetTicket() *string {
	return s.Ticket
}

func (s *DeleteTicketRequest) SetTicket(v string) *DeleteTicketRequest {
	s.Ticket = &v
	return s
}

func (s *DeleteTicketRequest) Validate() error {
	return dara.Validate(s)
}
