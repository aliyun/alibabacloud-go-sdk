// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTicketingCheckRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrderNum(v int64) *TicketingCheckRequest
	GetOrderNum() *int64
}

type TicketingCheckRequest struct {
	// order number
	//
	// This parameter is required.
	//
	// example:
	//
	// 496***2617111
	OrderNum *int64 `json:"order_num,omitempty" xml:"order_num,omitempty"`
}

func (s TicketingCheckRequest) String() string {
	return dara.Prettify(s)
}

func (s TicketingCheckRequest) GoString() string {
	return s.String()
}

func (s *TicketingCheckRequest) GetOrderNum() *int64 {
	return s.OrderNum
}

func (s *TicketingCheckRequest) SetOrderNum(v int64) *TicketingCheckRequest {
	s.OrderNum = &v
	return s
}

func (s *TicketingCheckRequest) Validate() error {
	return dara.Validate(s)
}
