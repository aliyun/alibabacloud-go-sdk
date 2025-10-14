// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTicketingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrderNum(v int64) *TicketingRequest
	GetOrderNum() *int64
}

type TicketingRequest struct {
	// order number
	//
	// This parameter is required.
	//
	// example:
	//
	// 496***2617111
	OrderNum *int64 `json:"order_num,omitempty" xml:"order_num,omitempty"`
}

func (s TicketingRequest) String() string {
	return dara.Prettify(s)
}

func (s TicketingRequest) GoString() string {
	return s.String()
}

func (s *TicketingRequest) GetOrderNum() *int64 {
	return s.OrderNum
}

func (s *TicketingRequest) SetOrderNum(v int64) *TicketingRequest {
	s.OrderNum = &v
	return s
}

func (s *TicketingRequest) Validate() error {
	return dara.Validate(s)
}
