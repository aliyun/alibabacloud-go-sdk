// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryHotelRoomDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHotelId(v string) *QueryHotelRoomDetailRequest
	GetHotelId() *string
	SetMac(v string) *QueryHotelRoomDetailRequest
	GetMac() *string
	SetRoomNo(v string) *QueryHotelRoomDetailRequest
	GetRoomNo() *string
	SetSn(v string) *QueryHotelRoomDetailRequest
	GetSn() *string
	SetUuid(v string) *QueryHotelRoomDetailRequest
	GetUuid() *string
}

type QueryHotelRoomDetailRequest struct {
	// example:
	//
	// 520a0c0***5eb
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	// example:
	//
	// 38:c8:**:**:f5:22
	Mac *string `json:"Mac,omitempty" xml:"Mac,omitempty"`
	// example:
	//
	// 1211
	RoomNo *string `json:"RoomNo,omitempty" xml:"RoomNo,omitempty"`
	// 设备sn信息
	//
	// 注：若在mac uuid sn全都输入的情况下 按照输入正确的内容查询 若全输入都是正确的 则 按照 uuid > mac > sn 优先级查询
	//
	// 传入mac uuid sn其中一个 则酒店id和房间号可不传
	//
	// example:
	//
	// 280**28
	Sn *string `json:"Sn,omitempty" xml:"Sn,omitempty"`
	// example:
	//
	// 588***96j5WU
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s QueryHotelRoomDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryHotelRoomDetailRequest) GoString() string {
	return s.String()
}

func (s *QueryHotelRoomDetailRequest) GetHotelId() *string {
	return s.HotelId
}

func (s *QueryHotelRoomDetailRequest) GetMac() *string {
	return s.Mac
}

func (s *QueryHotelRoomDetailRequest) GetRoomNo() *string {
	return s.RoomNo
}

func (s *QueryHotelRoomDetailRequest) GetSn() *string {
	return s.Sn
}

func (s *QueryHotelRoomDetailRequest) GetUuid() *string {
	return s.Uuid
}

func (s *QueryHotelRoomDetailRequest) SetHotelId(v string) *QueryHotelRoomDetailRequest {
	s.HotelId = &v
	return s
}

func (s *QueryHotelRoomDetailRequest) SetMac(v string) *QueryHotelRoomDetailRequest {
	s.Mac = &v
	return s
}

func (s *QueryHotelRoomDetailRequest) SetRoomNo(v string) *QueryHotelRoomDetailRequest {
	s.RoomNo = &v
	return s
}

func (s *QueryHotelRoomDetailRequest) SetSn(v string) *QueryHotelRoomDetailRequest {
	s.Sn = &v
	return s
}

func (s *QueryHotelRoomDetailRequest) SetUuid(v string) *QueryHotelRoomDetailRequest {
	s.Uuid = &v
	return s
}

func (s *QueryHotelRoomDetailRequest) Validate() error {
	return dara.Validate(s)
}
