// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHotelOrderCancelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBtripOrderId(v string) *HotelOrderCancelRequest
	GetBtripOrderId() *string
	SetDisOrderId(v string) *HotelOrderCancelRequest
	GetDisOrderId() *string
}

type HotelOrderCancelRequest struct {
	BtripOrderId *string `json:"btrip_order_id,omitempty" xml:"btrip_order_id,omitempty"`
	// example:
	//
	// dis123
	DisOrderId *string `json:"dis_order_id,omitempty" xml:"dis_order_id,omitempty"`
}

func (s HotelOrderCancelRequest) String() string {
	return dara.Prettify(s)
}

func (s HotelOrderCancelRequest) GoString() string {
	return s.String()
}

func (s *HotelOrderCancelRequest) GetBtripOrderId() *string {
	return s.BtripOrderId
}

func (s *HotelOrderCancelRequest) GetDisOrderId() *string {
	return s.DisOrderId
}

func (s *HotelOrderCancelRequest) SetBtripOrderId(v string) *HotelOrderCancelRequest {
	s.BtripOrderId = &v
	return s
}

func (s *HotelOrderCancelRequest) SetDisOrderId(v string) *HotelOrderCancelRequest {
	s.DisOrderId = &v
	return s
}

func (s *HotelOrderCancelRequest) Validate() error {
	return dara.Validate(s)
}
