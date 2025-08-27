// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTicketChangingCancelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDisOrderId(v string) *TicketChangingCancelRequest
	GetDisOrderId() *string
	SetDisSubOrderId(v string) *TicketChangingCancelRequest
	GetDisSubOrderId() *string
}

type TicketChangingCancelRequest struct {
	// example:
	//
	// dis123
	DisOrderId *string `json:"dis_order_id,omitempty" xml:"dis_order_id,omitempty"`
	// example:
	//
	// refun123
	DisSubOrderId *string `json:"dis_sub_order_id,omitempty" xml:"dis_sub_order_id,omitempty"`
}

func (s TicketChangingCancelRequest) String() string {
	return dara.Prettify(s)
}

func (s TicketChangingCancelRequest) GoString() string {
	return s.String()
}

func (s *TicketChangingCancelRequest) GetDisOrderId() *string {
	return s.DisOrderId
}

func (s *TicketChangingCancelRequest) GetDisSubOrderId() *string {
	return s.DisSubOrderId
}

func (s *TicketChangingCancelRequest) SetDisOrderId(v string) *TicketChangingCancelRequest {
	s.DisOrderId = &v
	return s
}

func (s *TicketChangingCancelRequest) SetDisSubOrderId(v string) *TicketChangingCancelRequest {
	s.DisSubOrderId = &v
	return s
}

func (s *TicketChangingCancelRequest) Validate() error {
	return dara.Validate(s)
}
