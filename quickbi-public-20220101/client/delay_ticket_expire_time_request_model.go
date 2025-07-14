// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDelayTicketExpireTimeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExpireTime(v int32) *DelayTicketExpireTimeRequest
	GetExpireTime() *int32
	SetTicket(v string) *DelayTicketExpireTimeRequest
	GetTicket() *string
}

type DelayTicketExpireTimeRequest struct {
	// The time to postpone.
	//
	// 	- Unit: minutes. Valid values: 0 to 240. Unit: minutes. Valid values: 4 hours.
	//
	// 	- Expired bills cannot be extended.
	//
	// This parameter is required.
	//
	// example:
	//
	// 200
	ExpireTime *int32 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The value of the third-party embedded ticket, that is, the accessTicket value in the URL.
	//
	// This parameter is required.
	//
	// example:
	//
	// 040e6f79d33444838e*****c7206c070
	Ticket *string `json:"Ticket,omitempty" xml:"Ticket,omitempty"`
}

func (s DelayTicketExpireTimeRequest) String() string {
	return dara.Prettify(s)
}

func (s DelayTicketExpireTimeRequest) GoString() string {
	return s.String()
}

func (s *DelayTicketExpireTimeRequest) GetExpireTime() *int32 {
	return s.ExpireTime
}

func (s *DelayTicketExpireTimeRequest) GetTicket() *string {
	return s.Ticket
}

func (s *DelayTicketExpireTimeRequest) SetExpireTime(v int32) *DelayTicketExpireTimeRequest {
	s.ExpireTime = &v
	return s
}

func (s *DelayTicketExpireTimeRequest) SetTicket(v string) *DelayTicketExpireTimeRequest {
	s.Ticket = &v
	return s
}

func (s *DelayTicketExpireTimeRequest) Validate() error {
	return dara.Validate(s)
}
