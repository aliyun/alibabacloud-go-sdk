// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateHotelShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppKey(v string) *UpdateHotelShrinkRequest
	GetAppKey() *string
	SetEstOpenTime(v string) *UpdateHotelShrinkRequest
	GetEstOpenTime() *string
	SetHotelAddress(v string) *UpdateHotelShrinkRequest
	GetHotelAddress() *string
	SetHotelEmail(v string) *UpdateHotelShrinkRequest
	GetHotelEmail() *string
	SetHotelId(v string) *UpdateHotelShrinkRequest
	GetHotelId() *string
	SetHotelName(v string) *UpdateHotelShrinkRequest
	GetHotelName() *string
	SetPhoneNumber(v string) *UpdateHotelShrinkRequest
	GetPhoneNumber() *string
	SetRelatedPksShrink(v string) *UpdateHotelShrinkRequest
	GetRelatedPksShrink() *string
	SetRemark(v string) *UpdateHotelShrinkRequest
	GetRemark() *string
	SetRoomNum(v int32) *UpdateHotelShrinkRequest
	GetRoomNum() *int32
	SetTbOpenId(v string) *UpdateHotelShrinkRequest
	GetTbOpenId() *string
}

type UpdateHotelShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 31342884
	AppKey *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// example:
	//
	// 2022-02-22 00:00:00
	EstOpenTime  *string `json:"EstOpenTime,omitempty" xml:"EstOpenTime,omitempty"`
	HotelAddress *string `json:"HotelAddress,omitempty" xml:"HotelAddress,omitempty"`
	// example:
	//
	// a*****@hotel.com
	HotelEmail *string `json:"HotelEmail,omitempty" xml:"HotelEmail,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// e6dd44fd16084db8a60d69fd625d9f0f
	HotelId   *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	HotelName *string `json:"HotelName,omitempty" xml:"HotelName,omitempty"`
	// example:
	//
	// 130***
	PhoneNumber      *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	RelatedPksShrink *string `json:"RelatedPks,omitempty" xml:"RelatedPks,omitempty"`
	Remark           *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// example:
	//
	// 4
	RoomNum *int32 `json:"RoomNum,omitempty" xml:"RoomNum,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// AAEVK***UE3d3Z2ETwh
	TbOpenId *string `json:"TbOpenId,omitempty" xml:"TbOpenId,omitempty"`
}

func (s UpdateHotelShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateHotelShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateHotelShrinkRequest) GetAppKey() *string {
	return s.AppKey
}

func (s *UpdateHotelShrinkRequest) GetEstOpenTime() *string {
	return s.EstOpenTime
}

func (s *UpdateHotelShrinkRequest) GetHotelAddress() *string {
	return s.HotelAddress
}

func (s *UpdateHotelShrinkRequest) GetHotelEmail() *string {
	return s.HotelEmail
}

func (s *UpdateHotelShrinkRequest) GetHotelId() *string {
	return s.HotelId
}

func (s *UpdateHotelShrinkRequest) GetHotelName() *string {
	return s.HotelName
}

func (s *UpdateHotelShrinkRequest) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *UpdateHotelShrinkRequest) GetRelatedPksShrink() *string {
	return s.RelatedPksShrink
}

func (s *UpdateHotelShrinkRequest) GetRemark() *string {
	return s.Remark
}

func (s *UpdateHotelShrinkRequest) GetRoomNum() *int32 {
	return s.RoomNum
}

func (s *UpdateHotelShrinkRequest) GetTbOpenId() *string {
	return s.TbOpenId
}

func (s *UpdateHotelShrinkRequest) SetAppKey(v string) *UpdateHotelShrinkRequest {
	s.AppKey = &v
	return s
}

func (s *UpdateHotelShrinkRequest) SetEstOpenTime(v string) *UpdateHotelShrinkRequest {
	s.EstOpenTime = &v
	return s
}

func (s *UpdateHotelShrinkRequest) SetHotelAddress(v string) *UpdateHotelShrinkRequest {
	s.HotelAddress = &v
	return s
}

func (s *UpdateHotelShrinkRequest) SetHotelEmail(v string) *UpdateHotelShrinkRequest {
	s.HotelEmail = &v
	return s
}

func (s *UpdateHotelShrinkRequest) SetHotelId(v string) *UpdateHotelShrinkRequest {
	s.HotelId = &v
	return s
}

func (s *UpdateHotelShrinkRequest) SetHotelName(v string) *UpdateHotelShrinkRequest {
	s.HotelName = &v
	return s
}

func (s *UpdateHotelShrinkRequest) SetPhoneNumber(v string) *UpdateHotelShrinkRequest {
	s.PhoneNumber = &v
	return s
}

func (s *UpdateHotelShrinkRequest) SetRelatedPksShrink(v string) *UpdateHotelShrinkRequest {
	s.RelatedPksShrink = &v
	return s
}

func (s *UpdateHotelShrinkRequest) SetRemark(v string) *UpdateHotelShrinkRequest {
	s.Remark = &v
	return s
}

func (s *UpdateHotelShrinkRequest) SetRoomNum(v int32) *UpdateHotelShrinkRequest {
	s.RoomNum = &v
	return s
}

func (s *UpdateHotelShrinkRequest) SetTbOpenId(v string) *UpdateHotelShrinkRequest {
	s.TbOpenId = &v
	return s
}

func (s *UpdateHotelShrinkRequest) Validate() error {
	return dara.Validate(s)
}
