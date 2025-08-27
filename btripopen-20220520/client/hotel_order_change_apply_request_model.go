// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHotelOrderChangeApplyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBtripUserId(v string) *HotelOrderChangeApplyRequest
	GetBtripUserId() *string
	SetDisOrderId(v string) *HotelOrderChangeApplyRequest
	GetDisOrderId() *string
	SetReason(v string) *HotelOrderChangeApplyRequest
	GetReason() *string
	SetRoomInfoList(v []*HotelOrderChangeApplyRequestRoomInfoList) *HotelOrderChangeApplyRequest
	GetRoomInfoList() []*HotelOrderChangeApplyRequestRoomInfoList
	SetSaleOrderId(v string) *HotelOrderChangeApplyRequest
	GetSaleOrderId() *string
}

type HotelOrderChangeApplyRequest struct {
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
	RoomInfoList []*HotelOrderChangeApplyRequestRoomInfoList `json:"room_info_list,omitempty" xml:"room_info_list,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 1402002197440511306
	SaleOrderId *string `json:"sale_order_id,omitempty" xml:"sale_order_id,omitempty"`
}

func (s HotelOrderChangeApplyRequest) String() string {
	return dara.Prettify(s)
}

func (s HotelOrderChangeApplyRequest) GoString() string {
	return s.String()
}

func (s *HotelOrderChangeApplyRequest) GetBtripUserId() *string {
	return s.BtripUserId
}

func (s *HotelOrderChangeApplyRequest) GetDisOrderId() *string {
	return s.DisOrderId
}

func (s *HotelOrderChangeApplyRequest) GetReason() *string {
	return s.Reason
}

func (s *HotelOrderChangeApplyRequest) GetRoomInfoList() []*HotelOrderChangeApplyRequestRoomInfoList {
	return s.RoomInfoList
}

func (s *HotelOrderChangeApplyRequest) GetSaleOrderId() *string {
	return s.SaleOrderId
}

func (s *HotelOrderChangeApplyRequest) SetBtripUserId(v string) *HotelOrderChangeApplyRequest {
	s.BtripUserId = &v
	return s
}

func (s *HotelOrderChangeApplyRequest) SetDisOrderId(v string) *HotelOrderChangeApplyRequest {
	s.DisOrderId = &v
	return s
}

func (s *HotelOrderChangeApplyRequest) SetReason(v string) *HotelOrderChangeApplyRequest {
	s.Reason = &v
	return s
}

func (s *HotelOrderChangeApplyRequest) SetRoomInfoList(v []*HotelOrderChangeApplyRequestRoomInfoList) *HotelOrderChangeApplyRequest {
	s.RoomInfoList = v
	return s
}

func (s *HotelOrderChangeApplyRequest) SetSaleOrderId(v string) *HotelOrderChangeApplyRequest {
	s.SaleOrderId = &v
	return s
}

func (s *HotelOrderChangeApplyRequest) Validate() error {
	return dara.Validate(s)
}

type HotelOrderChangeApplyRequestRoomInfoList struct {
	// This parameter is required.
	CancelDate []*string `json:"cancel_date,omitempty" xml:"cancel_date,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 112
	RoomNo *int32 `json:"room_no,omitempty" xml:"room_no,omitempty"`
}

func (s HotelOrderChangeApplyRequestRoomInfoList) String() string {
	return dara.Prettify(s)
}

func (s HotelOrderChangeApplyRequestRoomInfoList) GoString() string {
	return s.String()
}

func (s *HotelOrderChangeApplyRequestRoomInfoList) GetCancelDate() []*string {
	return s.CancelDate
}

func (s *HotelOrderChangeApplyRequestRoomInfoList) GetRoomNo() *int32 {
	return s.RoomNo
}

func (s *HotelOrderChangeApplyRequestRoomInfoList) SetCancelDate(v []*string) *HotelOrderChangeApplyRequestRoomInfoList {
	s.CancelDate = v
	return s
}

func (s *HotelOrderChangeApplyRequestRoomInfoList) SetRoomNo(v int32) *HotelOrderChangeApplyRequestRoomInfoList {
	s.RoomNo = &v
	return s
}

func (s *HotelOrderChangeApplyRequestRoomInfoList) Validate() error {
	return dara.Validate(s)
}
