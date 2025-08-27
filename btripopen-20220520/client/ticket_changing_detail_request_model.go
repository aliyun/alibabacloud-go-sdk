// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTicketChangingDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDisOrderId(v string) *TicketChangingDetailRequest
	GetDisOrderId() *string
	SetDisSubOrderId(v string) *TicketChangingDetailRequest
	GetDisSubOrderId() *string
}

type TicketChangingDetailRequest struct {
	// example:
	//
	// dis123
	DisOrderId *string `json:"dis_order_id,omitempty" xml:"dis_order_id,omitempty"`
	// example:
	//
	// chang123
	DisSubOrderId *string `json:"dis_sub_order_id,omitempty" xml:"dis_sub_order_id,omitempty"`
}

func (s TicketChangingDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s TicketChangingDetailRequest) GoString() string {
	return s.String()
}

func (s *TicketChangingDetailRequest) GetDisOrderId() *string {
	return s.DisOrderId
}

func (s *TicketChangingDetailRequest) GetDisSubOrderId() *string {
	return s.DisSubOrderId
}

func (s *TicketChangingDetailRequest) SetDisOrderId(v string) *TicketChangingDetailRequest {
	s.DisOrderId = &v
	return s
}

func (s *TicketChangingDetailRequest) SetDisSubOrderId(v string) *TicketChangingDetailRequest {
	s.DisSubOrderId = &v
	return s
}

func (s *TicketChangingDetailRequest) Validate() error {
	return dara.Validate(s)
}
