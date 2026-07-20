// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHotelOrderQueryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v int64) *HotelOrderQueryRequest
	GetOrderId() *int64
}

type HotelOrderQueryRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1002145190081005400
	OrderId *int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
}

func (s HotelOrderQueryRequest) String() string {
	return dara.Prettify(s)
}

func (s HotelOrderQueryRequest) GoString() string {
	return s.String()
}

func (s *HotelOrderQueryRequest) GetOrderId() *int64 {
	return s.OrderId
}

func (s *HotelOrderQueryRequest) SetOrderId(v int64) *HotelOrderQueryRequest {
	s.OrderId = &v
	return s
}

func (s *HotelOrderQueryRequest) Validate() error {
	return dara.Validate(s)
}
