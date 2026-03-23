// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTicketRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTicketId(v string) *GetTicketRequest
	GetTicketId() *string
}

type GetTicketRequest struct {
	// This parameter is required.
	TicketId *string `json:"TicketId,omitempty" xml:"TicketId,omitempty"`
}

func (s GetTicketRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTicketRequest) GoString() string {
	return s.String()
}

func (s *GetTicketRequest) GetTicketId() *string {
	return s.TicketId
}

func (s *GetTicketRequest) SetTicketId(v string) *GetTicketRequest {
	s.TicketId = &v
	return s
}

func (s *GetTicketRequest) Validate() error {
	return dara.Validate(s)
}
