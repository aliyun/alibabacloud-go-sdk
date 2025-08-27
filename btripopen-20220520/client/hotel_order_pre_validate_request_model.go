// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHotelOrderPreValidateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBtripUserId(v string) *HotelOrderPreValidateRequest
	GetBtripUserId() *string
	SetCheckIn(v string) *HotelOrderPreValidateRequest
	GetCheckIn() *string
	SetCheckOut(v string) *HotelOrderPreValidateRequest
	GetCheckOut() *string
	SetDailyList(v []*HotelOrderPreValidateRequestDailyList) *HotelOrderPreValidateRequest
	GetDailyList() []*HotelOrderPreValidateRequestDailyList
	SetItemId(v int64) *HotelOrderPreValidateRequest
	GetItemId() *int64
	SetNumberOfAdultsPerRoom(v int32) *HotelOrderPreValidateRequest
	GetNumberOfAdultsPerRoom() *int32
	SetOccupantInfoList(v []*HotelOrderPreValidateRequestOccupantInfoList) *HotelOrderPreValidateRequest
	GetOccupantInfoList() []*HotelOrderPreValidateRequestOccupantInfoList
	SetRatePlanId(v int64) *HotelOrderPreValidateRequest
	GetRatePlanId() *int64
	SetRoomId(v int64) *HotelOrderPreValidateRequest
	GetRoomId() *int64
	SetRoomNum(v int32) *HotelOrderPreValidateRequest
	GetRoomNum() *int32
	SetSearchRoomPrice(v int64) *HotelOrderPreValidateRequest
	GetSearchRoomPrice() *int64
	SetSellerId(v int64) *HotelOrderPreValidateRequest
	GetSellerId() *int64
	SetShid(v int64) *HotelOrderPreValidateRequest
	GetShid() *int64
}

type HotelOrderPreValidateRequest struct {
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
	DailyList []*HotelOrderPreValidateRequestDailyList `json:"daily_list,omitempty" xml:"daily_list,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 612673015638
	ItemId *int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// example:
	//
	// 1
	NumberOfAdultsPerRoom *int32                                          `json:"number_of_adults_per_room,omitempty" xml:"number_of_adults_per_room,omitempty"`
	OccupantInfoList      []*HotelOrderPreValidateRequestOccupantInfoList `json:"occupant_info_list,omitempty" xml:"occupant_info_list,omitempty" type:"Repeated"`
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

func (s HotelOrderPreValidateRequest) String() string {
	return dara.Prettify(s)
}

func (s HotelOrderPreValidateRequest) GoString() string {
	return s.String()
}

func (s *HotelOrderPreValidateRequest) GetBtripUserId() *string {
	return s.BtripUserId
}

func (s *HotelOrderPreValidateRequest) GetCheckIn() *string {
	return s.CheckIn
}

func (s *HotelOrderPreValidateRequest) GetCheckOut() *string {
	return s.CheckOut
}

func (s *HotelOrderPreValidateRequest) GetDailyList() []*HotelOrderPreValidateRequestDailyList {
	return s.DailyList
}

func (s *HotelOrderPreValidateRequest) GetItemId() *int64 {
	return s.ItemId
}

func (s *HotelOrderPreValidateRequest) GetNumberOfAdultsPerRoom() *int32 {
	return s.NumberOfAdultsPerRoom
}

func (s *HotelOrderPreValidateRequest) GetOccupantInfoList() []*HotelOrderPreValidateRequestOccupantInfoList {
	return s.OccupantInfoList
}

func (s *HotelOrderPreValidateRequest) GetRatePlanId() *int64 {
	return s.RatePlanId
}

func (s *HotelOrderPreValidateRequest) GetRoomId() *int64 {
	return s.RoomId
}

func (s *HotelOrderPreValidateRequest) GetRoomNum() *int32 {
	return s.RoomNum
}

func (s *HotelOrderPreValidateRequest) GetSearchRoomPrice() *int64 {
	return s.SearchRoomPrice
}

func (s *HotelOrderPreValidateRequest) GetSellerId() *int64 {
	return s.SellerId
}

func (s *HotelOrderPreValidateRequest) GetShid() *int64 {
	return s.Shid
}

func (s *HotelOrderPreValidateRequest) SetBtripUserId(v string) *HotelOrderPreValidateRequest {
	s.BtripUserId = &v
	return s
}

func (s *HotelOrderPreValidateRequest) SetCheckIn(v string) *HotelOrderPreValidateRequest {
	s.CheckIn = &v
	return s
}

func (s *HotelOrderPreValidateRequest) SetCheckOut(v string) *HotelOrderPreValidateRequest {
	s.CheckOut = &v
	return s
}

func (s *HotelOrderPreValidateRequest) SetDailyList(v []*HotelOrderPreValidateRequestDailyList) *HotelOrderPreValidateRequest {
	s.DailyList = v
	return s
}

func (s *HotelOrderPreValidateRequest) SetItemId(v int64) *HotelOrderPreValidateRequest {
	s.ItemId = &v
	return s
}

func (s *HotelOrderPreValidateRequest) SetNumberOfAdultsPerRoom(v int32) *HotelOrderPreValidateRequest {
	s.NumberOfAdultsPerRoom = &v
	return s
}

func (s *HotelOrderPreValidateRequest) SetOccupantInfoList(v []*HotelOrderPreValidateRequestOccupantInfoList) *HotelOrderPreValidateRequest {
	s.OccupantInfoList = v
	return s
}

func (s *HotelOrderPreValidateRequest) SetRatePlanId(v int64) *HotelOrderPreValidateRequest {
	s.RatePlanId = &v
	return s
}

func (s *HotelOrderPreValidateRequest) SetRoomId(v int64) *HotelOrderPreValidateRequest {
	s.RoomId = &v
	return s
}

func (s *HotelOrderPreValidateRequest) SetRoomNum(v int32) *HotelOrderPreValidateRequest {
	s.RoomNum = &v
	return s
}

func (s *HotelOrderPreValidateRequest) SetSearchRoomPrice(v int64) *HotelOrderPreValidateRequest {
	s.SearchRoomPrice = &v
	return s
}

func (s *HotelOrderPreValidateRequest) SetSellerId(v int64) *HotelOrderPreValidateRequest {
	s.SellerId = &v
	return s
}

func (s *HotelOrderPreValidateRequest) SetShid(v int64) *HotelOrderPreValidateRequest {
	s.Shid = &v
	return s
}

func (s *HotelOrderPreValidateRequest) Validate() error {
	return dara.Validate(s)
}

type HotelOrderPreValidateRequestDailyList struct {
	// example:
	//
	// 1
	Board *string `json:"board,omitempty" xml:"board,omitempty"`
	// example:
	//
	// 100
	Price *int64 `json:"price,omitempty" xml:"price,omitempty"`
	// example:
	//
	// 2022-05-15
	RateStartTime *string `json:"rate_start_time,omitempty" xml:"rate_start_time,omitempty"`
	// example:
	//
	// 10
	RoomCount *int32 `json:"room_count,omitempty" xml:"room_count,omitempty"`
}

func (s HotelOrderPreValidateRequestDailyList) String() string {
	return dara.Prettify(s)
}

func (s HotelOrderPreValidateRequestDailyList) GoString() string {
	return s.String()
}

func (s *HotelOrderPreValidateRequestDailyList) GetBoard() *string {
	return s.Board
}

func (s *HotelOrderPreValidateRequestDailyList) GetPrice() *int64 {
	return s.Price
}

func (s *HotelOrderPreValidateRequestDailyList) GetRateStartTime() *string {
	return s.RateStartTime
}

func (s *HotelOrderPreValidateRequestDailyList) GetRoomCount() *int32 {
	return s.RoomCount
}

func (s *HotelOrderPreValidateRequestDailyList) SetBoard(v string) *HotelOrderPreValidateRequestDailyList {
	s.Board = &v
	return s
}

func (s *HotelOrderPreValidateRequestDailyList) SetPrice(v int64) *HotelOrderPreValidateRequestDailyList {
	s.Price = &v
	return s
}

func (s *HotelOrderPreValidateRequestDailyList) SetRateStartTime(v string) *HotelOrderPreValidateRequestDailyList {
	s.RateStartTime = &v
	return s
}

func (s *HotelOrderPreValidateRequestDailyList) SetRoomCount(v int32) *HotelOrderPreValidateRequestDailyList {
	s.RoomCount = &v
	return s
}

func (s *HotelOrderPreValidateRequestDailyList) Validate() error {
	return dara.Validate(s)
}

type HotelOrderPreValidateRequestOccupantInfoList struct {
	// example:
	//
	// 232871871822
	CardNo *string `json:"card_no,omitempty" xml:"card_no,omitempty"`
	// example:
	//
	// 1
	CardType *int32  `json:"card_type,omitempty" xml:"card_type,omitempty"`
	Name     *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 12392827121
	Phone *string `json:"phone,omitempty" xml:"phone,omitempty"`
	// example:
	//
	// 128918
	StaffNo *string `json:"staff_no,omitempty" xml:"staff_no,omitempty"`
	// example:
	//
	// 1
	UserType *int32 `json:"user_type,omitempty" xml:"user_type,omitempty"`
}

func (s HotelOrderPreValidateRequestOccupantInfoList) String() string {
	return dara.Prettify(s)
}

func (s HotelOrderPreValidateRequestOccupantInfoList) GoString() string {
	return s.String()
}

func (s *HotelOrderPreValidateRequestOccupantInfoList) GetCardNo() *string {
	return s.CardNo
}

func (s *HotelOrderPreValidateRequestOccupantInfoList) GetCardType() *int32 {
	return s.CardType
}

func (s *HotelOrderPreValidateRequestOccupantInfoList) GetName() *string {
	return s.Name
}

func (s *HotelOrderPreValidateRequestOccupantInfoList) GetPhone() *string {
	return s.Phone
}

func (s *HotelOrderPreValidateRequestOccupantInfoList) GetStaffNo() *string {
	return s.StaffNo
}

func (s *HotelOrderPreValidateRequestOccupantInfoList) GetUserType() *int32 {
	return s.UserType
}

func (s *HotelOrderPreValidateRequestOccupantInfoList) SetCardNo(v string) *HotelOrderPreValidateRequestOccupantInfoList {
	s.CardNo = &v
	return s
}

func (s *HotelOrderPreValidateRequestOccupantInfoList) SetCardType(v int32) *HotelOrderPreValidateRequestOccupantInfoList {
	s.CardType = &v
	return s
}

func (s *HotelOrderPreValidateRequestOccupantInfoList) SetName(v string) *HotelOrderPreValidateRequestOccupantInfoList {
	s.Name = &v
	return s
}

func (s *HotelOrderPreValidateRequestOccupantInfoList) SetPhone(v string) *HotelOrderPreValidateRequestOccupantInfoList {
	s.Phone = &v
	return s
}

func (s *HotelOrderPreValidateRequestOccupantInfoList) SetStaffNo(v string) *HotelOrderPreValidateRequestOccupantInfoList {
	s.StaffNo = &v
	return s
}

func (s *HotelOrderPreValidateRequestOccupantInfoList) SetUserType(v int32) *HotelOrderPreValidateRequestOccupantInfoList {
	s.UserType = &v
	return s
}

func (s *HotelOrderPreValidateRequestOccupantInfoList) Validate() error {
	return dara.Validate(s)
}
