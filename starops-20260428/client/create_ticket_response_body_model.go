// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTicketResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetTicket(v string) *CreateTicketResponseBody
	GetTicket() *string
}

type CreateTicketResponseBody struct {
	// example:
	//
	// eyJ***************.eyJ******************.KUT****************
	Ticket *string `json:"ticket,omitempty" xml:"ticket,omitempty"`
}

func (s CreateTicketResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateTicketResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTicketResponseBody) GetTicket() *string {
	return s.Ticket
}

func (s *CreateTicketResponseBody) SetTicket(v string) *CreateTicketResponseBody {
	s.Ticket = &v
	return s
}

func (s *CreateTicketResponseBody) Validate() error {
	return dara.Validate(s)
}
