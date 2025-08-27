// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHotelOrderChangeDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBtripUserId(v string) *HotelOrderChangeDetailRequest
	GetBtripUserId() *string
	SetChangeOrderId(v string) *HotelOrderChangeDetailRequest
	GetChangeOrderId() *string
	SetDisOrderId(v string) *HotelOrderChangeDetailRequest
	GetDisOrderId() *string
	SetSaleOrderId(v string) *HotelOrderChangeDetailRequest
	GetSaleOrderId() *string
}

type HotelOrderChangeDetailRequest struct {
	// example:
	//
	// 123455
	BtripUserId *string `json:"btrip_user_id,omitempty" xml:"btrip_user_id,omitempty"`
	// example:
	//
	// 21351235134
	ChangeOrderId *string `json:"change_order_id,omitempty" xml:"change_order_id,omitempty"`
	// example:
	//
	// dis123
	DisOrderId *string `json:"dis_order_id,omitempty" xml:"dis_order_id,omitempty"`
	// example:
	//
	// 1402002197440511306
	SaleOrderId *string `json:"sale_order_id,omitempty" xml:"sale_order_id,omitempty"`
}

func (s HotelOrderChangeDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s HotelOrderChangeDetailRequest) GoString() string {
	return s.String()
}

func (s *HotelOrderChangeDetailRequest) GetBtripUserId() *string {
	return s.BtripUserId
}

func (s *HotelOrderChangeDetailRequest) GetChangeOrderId() *string {
	return s.ChangeOrderId
}

func (s *HotelOrderChangeDetailRequest) GetDisOrderId() *string {
	return s.DisOrderId
}

func (s *HotelOrderChangeDetailRequest) GetSaleOrderId() *string {
	return s.SaleOrderId
}

func (s *HotelOrderChangeDetailRequest) SetBtripUserId(v string) *HotelOrderChangeDetailRequest {
	s.BtripUserId = &v
	return s
}

func (s *HotelOrderChangeDetailRequest) SetChangeOrderId(v string) *HotelOrderChangeDetailRequest {
	s.ChangeOrderId = &v
	return s
}

func (s *HotelOrderChangeDetailRequest) SetDisOrderId(v string) *HotelOrderChangeDetailRequest {
	s.DisOrderId = &v
	return s
}

func (s *HotelOrderChangeDetailRequest) SetSaleOrderId(v string) *HotelOrderChangeDetailRequest {
	s.SaleOrderId = &v
	return s
}

func (s *HotelOrderChangeDetailRequest) Validate() error {
	return dara.Validate(s)
}
