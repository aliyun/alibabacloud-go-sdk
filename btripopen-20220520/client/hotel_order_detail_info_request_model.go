// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHotelOrderDetailInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBtripOrderId(v string) *HotelOrderDetailInfoRequest
	GetBtripOrderId() *string
	SetDisOrderId(v string) *HotelOrderDetailInfoRequest
	GetDisOrderId() *string
}

type HotelOrderDetailInfoRequest struct {
	BtripOrderId *string `json:"btrip_order_id,omitempty" xml:"btrip_order_id,omitempty"`
	// example:
	//
	// dis123
	DisOrderId *string `json:"dis_order_id,omitempty" xml:"dis_order_id,omitempty"`
}

func (s HotelOrderDetailInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s HotelOrderDetailInfoRequest) GoString() string {
	return s.String()
}

func (s *HotelOrderDetailInfoRequest) GetBtripOrderId() *string {
	return s.BtripOrderId
}

func (s *HotelOrderDetailInfoRequest) GetDisOrderId() *string {
	return s.DisOrderId
}

func (s *HotelOrderDetailInfoRequest) SetBtripOrderId(v string) *HotelOrderDetailInfoRequest {
	s.BtripOrderId = &v
	return s
}

func (s *HotelOrderDetailInfoRequest) SetDisOrderId(v string) *HotelOrderDetailInfoRequest {
	s.DisOrderId = &v
	return s
}

func (s *HotelOrderDetailInfoRequest) Validate() error {
	return dara.Validate(s)
}
