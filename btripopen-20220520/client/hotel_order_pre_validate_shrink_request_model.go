// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHotelOrderPreValidateShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBtripUserId(v string) *HotelOrderPreValidateShrinkRequest
	GetBtripUserId() *string
	SetCheckIn(v string) *HotelOrderPreValidateShrinkRequest
	GetCheckIn() *string
	SetCheckOut(v string) *HotelOrderPreValidateShrinkRequest
	GetCheckOut() *string
	SetDailyListShrink(v string) *HotelOrderPreValidateShrinkRequest
	GetDailyListShrink() *string
	SetItemId(v int64) *HotelOrderPreValidateShrinkRequest
	GetItemId() *int64
	SetNumberOfAdultsPerRoom(v int32) *HotelOrderPreValidateShrinkRequest
	GetNumberOfAdultsPerRoom() *int32
	SetOccupantInfoListShrink(v string) *HotelOrderPreValidateShrinkRequest
	GetOccupantInfoListShrink() *string
	SetRatePlanId(v int64) *HotelOrderPreValidateShrinkRequest
	GetRatePlanId() *int64
	SetRoomId(v int64) *HotelOrderPreValidateShrinkRequest
	GetRoomId() *int64
	SetRoomNum(v int32) *HotelOrderPreValidateShrinkRequest
	GetRoomNum() *int32
	SetSearchRoomPrice(v int64) *HotelOrderPreValidateShrinkRequest
	GetSearchRoomPrice() *int64
	SetSellerId(v int64) *HotelOrderPreValidateShrinkRequest
	GetSellerId() *int64
	SetShid(v int64) *HotelOrderPreValidateShrinkRequest
	GetShid() *int64
}

type HotelOrderPreValidateShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 23141
	BtripUserId *string `json:"btrip_user_id,omitempty" xml:"btrip_user_id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2022-05-15
	CheckIn *string `json:"check_in,omitempty" xml:"check_in,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2022-05-15
	CheckOut *string `json:"check_out,omitempty" xml:"check_out,omitempty"`
	// This parameter is required.
	DailyListShrink *string `json:"daily_list,omitempty" xml:"daily_list,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 612673015638
	ItemId *int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// example:
	//
	// 1
	NumberOfAdultsPerRoom  *int32  `json:"number_of_adults_per_room,omitempty" xml:"number_of_adults_per_room,omitempty"`
	OccupantInfoListShrink *string `json:"occupant_info_list,omitempty" xml:"occupant_info_list,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 239872781
	RatePlanId *int64 `json:"rate_plan_id,omitempty" xml:"rate_plan_id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 71652158
	RoomId *int64 `json:"room_id,omitempty" xml:"room_id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	RoomNum *int32 `json:"room_num,omitempty" xml:"room_num,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1000
	SearchRoomPrice *int64 `json:"search_room_price,omitempty" xml:"search_room_price,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2088441675613762
	SellerId *int64 `json:"seller_id,omitempty" xml:"seller_id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 52302073
	Shid *int64 `json:"shid,omitempty" xml:"shid,omitempty"`
}

func (s HotelOrderPreValidateShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s HotelOrderPreValidateShrinkRequest) GoString() string {
	return s.String()
}

func (s *HotelOrderPreValidateShrinkRequest) GetBtripUserId() *string {
	return s.BtripUserId
}

func (s *HotelOrderPreValidateShrinkRequest) GetCheckIn() *string {
	return s.CheckIn
}

func (s *HotelOrderPreValidateShrinkRequest) GetCheckOut() *string {
	return s.CheckOut
}

func (s *HotelOrderPreValidateShrinkRequest) GetDailyListShrink() *string {
	return s.DailyListShrink
}

func (s *HotelOrderPreValidateShrinkRequest) GetItemId() *int64 {
	return s.ItemId
}

func (s *HotelOrderPreValidateShrinkRequest) GetNumberOfAdultsPerRoom() *int32 {
	return s.NumberOfAdultsPerRoom
}

func (s *HotelOrderPreValidateShrinkRequest) GetOccupantInfoListShrink() *string {
	return s.OccupantInfoListShrink
}

func (s *HotelOrderPreValidateShrinkRequest) GetRatePlanId() *int64 {
	return s.RatePlanId
}

func (s *HotelOrderPreValidateShrinkRequest) GetRoomId() *int64 {
	return s.RoomId
}

func (s *HotelOrderPreValidateShrinkRequest) GetRoomNum() *int32 {
	return s.RoomNum
}

func (s *HotelOrderPreValidateShrinkRequest) GetSearchRoomPrice() *int64 {
	return s.SearchRoomPrice
}

func (s *HotelOrderPreValidateShrinkRequest) GetSellerId() *int64 {
	return s.SellerId
}

func (s *HotelOrderPreValidateShrinkRequest) GetShid() *int64 {
	return s.Shid
}

func (s *HotelOrderPreValidateShrinkRequest) SetBtripUserId(v string) *HotelOrderPreValidateShrinkRequest {
	s.BtripUserId = &v
	return s
}

func (s *HotelOrderPreValidateShrinkRequest) SetCheckIn(v string) *HotelOrderPreValidateShrinkRequest {
	s.CheckIn = &v
	return s
}

func (s *HotelOrderPreValidateShrinkRequest) SetCheckOut(v string) *HotelOrderPreValidateShrinkRequest {
	s.CheckOut = &v
	return s
}

func (s *HotelOrderPreValidateShrinkRequest) SetDailyListShrink(v string) *HotelOrderPreValidateShrinkRequest {
	s.DailyListShrink = &v
	return s
}

func (s *HotelOrderPreValidateShrinkRequest) SetItemId(v int64) *HotelOrderPreValidateShrinkRequest {
	s.ItemId = &v
	return s
}

func (s *HotelOrderPreValidateShrinkRequest) SetNumberOfAdultsPerRoom(v int32) *HotelOrderPreValidateShrinkRequest {
	s.NumberOfAdultsPerRoom = &v
	return s
}

func (s *HotelOrderPreValidateShrinkRequest) SetOccupantInfoListShrink(v string) *HotelOrderPreValidateShrinkRequest {
	s.OccupantInfoListShrink = &v
	return s
}

func (s *HotelOrderPreValidateShrinkRequest) SetRatePlanId(v int64) *HotelOrderPreValidateShrinkRequest {
	s.RatePlanId = &v
	return s
}

func (s *HotelOrderPreValidateShrinkRequest) SetRoomId(v int64) *HotelOrderPreValidateShrinkRequest {
	s.RoomId = &v
	return s
}

func (s *HotelOrderPreValidateShrinkRequest) SetRoomNum(v int32) *HotelOrderPreValidateShrinkRequest {
	s.RoomNum = &v
	return s
}

func (s *HotelOrderPreValidateShrinkRequest) SetSearchRoomPrice(v int64) *HotelOrderPreValidateShrinkRequest {
	s.SearchRoomPrice = &v
	return s
}

func (s *HotelOrderPreValidateShrinkRequest) SetSellerId(v int64) *HotelOrderPreValidateShrinkRequest {
	s.SellerId = &v
	return s
}

func (s *HotelOrderPreValidateShrinkRequest) SetShid(v int64) *HotelOrderPreValidateShrinkRequest {
	s.Shid = &v
	return s
}

func (s *HotelOrderPreValidateShrinkRequest) Validate() error {
	return dara.Validate(s)
}
