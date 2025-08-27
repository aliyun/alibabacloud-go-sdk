// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHotelOrderChangeApplyShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBtripUserId(v string) *HotelOrderChangeApplyShrinkRequest
	GetBtripUserId() *string
	SetDisOrderId(v string) *HotelOrderChangeApplyShrinkRequest
	GetDisOrderId() *string
	SetReason(v string) *HotelOrderChangeApplyShrinkRequest
	GetReason() *string
	SetRoomInfoListShrink(v string) *HotelOrderChangeApplyShrinkRequest
	GetRoomInfoListShrink() *string
	SetSaleOrderId(v string) *HotelOrderChangeApplyShrinkRequest
	GetSaleOrderId() *string
}

type HotelOrderChangeApplyShrinkRequest struct {
	// example:
	//
	// 123122
	BtripUserId *string `json:"btrip_user_id,omitempty" xml:"btrip_user_id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 3685792244476194816
	DisOrderId *string `json:"dis_order_id,omitempty" xml:"dis_order_id,omitempty"`
	// This parameter is required.
	Reason *string `json:"reason,omitempty" xml:"reason,omitempty"`
	// This parameter is required.
	RoomInfoListShrink *string `json:"room_info_list,omitempty" xml:"room_info_list,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1402002197440511306
	SaleOrderId *string `json:"sale_order_id,omitempty" xml:"sale_order_id,omitempty"`
}

func (s HotelOrderChangeApplyShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s HotelOrderChangeApplyShrinkRequest) GoString() string {
	return s.String()
}

func (s *HotelOrderChangeApplyShrinkRequest) GetBtripUserId() *string {
	return s.BtripUserId
}

func (s *HotelOrderChangeApplyShrinkRequest) GetDisOrderId() *string {
	return s.DisOrderId
}

func (s *HotelOrderChangeApplyShrinkRequest) GetReason() *string {
	return s.Reason
}

func (s *HotelOrderChangeApplyShrinkRequest) GetRoomInfoListShrink() *string {
	return s.RoomInfoListShrink
}

func (s *HotelOrderChangeApplyShrinkRequest) GetSaleOrderId() *string {
	return s.SaleOrderId
}

func (s *HotelOrderChangeApplyShrinkRequest) SetBtripUserId(v string) *HotelOrderChangeApplyShrinkRequest {
	s.BtripUserId = &v
	return s
}

func (s *HotelOrderChangeApplyShrinkRequest) SetDisOrderId(v string) *HotelOrderChangeApplyShrinkRequest {
	s.DisOrderId = &v
	return s
}

func (s *HotelOrderChangeApplyShrinkRequest) SetReason(v string) *HotelOrderChangeApplyShrinkRequest {
	s.Reason = &v
	return s
}

func (s *HotelOrderChangeApplyShrinkRequest) SetRoomInfoListShrink(v string) *HotelOrderChangeApplyShrinkRequest {
	s.RoomInfoListShrink = &v
	return s
}

func (s *HotelOrderChangeApplyShrinkRequest) SetSaleOrderId(v string) *HotelOrderChangeApplyShrinkRequest {
	s.SaleOrderId = &v
	return s
}

func (s *HotelOrderChangeApplyShrinkRequest) Validate() error {
	return dara.Validate(s)
}
