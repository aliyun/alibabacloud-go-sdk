// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHotelShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppKey(v string) *CreateHotelShrinkRequest
	GetAppKey() *string
	SetEstOpenTime(v string) *CreateHotelShrinkRequest
	GetEstOpenTime() *string
	SetHotelAddress(v string) *CreateHotelShrinkRequest
	GetHotelAddress() *string
	SetHotelEmail(v string) *CreateHotelShrinkRequest
	GetHotelEmail() *string
	SetHotelName(v string) *CreateHotelShrinkRequest
	GetHotelName() *string
	SetPhoneNumber(v string) *CreateHotelShrinkRequest
	GetPhoneNumber() *string
	SetRelatedPk(v string) *CreateHotelShrinkRequest
	GetRelatedPk() *string
	SetRelatedPksShrink(v string) *CreateHotelShrinkRequest
	GetRelatedPksShrink() *string
	SetRemark(v string) *CreateHotelShrinkRequest
	GetRemark() *string
	SetRoomNum(v int32) *CreateHotelShrinkRequest
	GetRoomNum() *int32
	SetTbOpenId(v string) *CreateHotelShrinkRequest
	GetTbOpenId() *string
}

type CreateHotelShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 333566791
	AppKey *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2022-10-1 00:00:00
	EstOpenTime *string `json:"EstOpenTime,omitempty" xml:"EstOpenTime,omitempty"`
	// This parameter is required.
	HotelAddress *string `json:"HotelAddress,omitempty" xml:"HotelAddress,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test@hotel.com
	HotelEmail *string `json:"HotelEmail,omitempty" xml:"HotelEmail,omitempty"`
	// This parameter is required.
	HotelName *string `json:"HotelName,omitempty" xml:"HotelName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 13xxxxxxxx
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// jTO****Rw
	RelatedPk *string `json:"RelatedPk,omitempty" xml:"RelatedPk,omitempty"`
	// 酒店关联产品列表
	RelatedPksShrink *string `json:"RelatedPks,omitempty" xml:"RelatedPks,omitempty"`
	// example:
	//
	// test
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 100
	RoomNum *int32 `json:"RoomNum,omitempty" xml:"RoomNum,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// AAEV***E3d3Z2ETwh
	TbOpenId *string `json:"TbOpenId,omitempty" xml:"TbOpenId,omitempty"`
}

func (s CreateHotelShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateHotelShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateHotelShrinkRequest) GetAppKey() *string {
	return s.AppKey
}

func (s *CreateHotelShrinkRequest) GetEstOpenTime() *string {
	return s.EstOpenTime
}

func (s *CreateHotelShrinkRequest) GetHotelAddress() *string {
	return s.HotelAddress
}

func (s *CreateHotelShrinkRequest) GetHotelEmail() *string {
	return s.HotelEmail
}

func (s *CreateHotelShrinkRequest) GetHotelName() *string {
	return s.HotelName
}

func (s *CreateHotelShrinkRequest) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *CreateHotelShrinkRequest) GetRelatedPk() *string {
	return s.RelatedPk
}

func (s *CreateHotelShrinkRequest) GetRelatedPksShrink() *string {
	return s.RelatedPksShrink
}

func (s *CreateHotelShrinkRequest) GetRemark() *string {
	return s.Remark
}

func (s *CreateHotelShrinkRequest) GetRoomNum() *int32 {
	return s.RoomNum
}

func (s *CreateHotelShrinkRequest) GetTbOpenId() *string {
	return s.TbOpenId
}

func (s *CreateHotelShrinkRequest) SetAppKey(v string) *CreateHotelShrinkRequest {
	s.AppKey = &v
	return s
}

func (s *CreateHotelShrinkRequest) SetEstOpenTime(v string) *CreateHotelShrinkRequest {
	s.EstOpenTime = &v
	return s
}

func (s *CreateHotelShrinkRequest) SetHotelAddress(v string) *CreateHotelShrinkRequest {
	s.HotelAddress = &v
	return s
}

func (s *CreateHotelShrinkRequest) SetHotelEmail(v string) *CreateHotelShrinkRequest {
	s.HotelEmail = &v
	return s
}

func (s *CreateHotelShrinkRequest) SetHotelName(v string) *CreateHotelShrinkRequest {
	s.HotelName = &v
	return s
}

func (s *CreateHotelShrinkRequest) SetPhoneNumber(v string) *CreateHotelShrinkRequest {
	s.PhoneNumber = &v
	return s
}

func (s *CreateHotelShrinkRequest) SetRelatedPk(v string) *CreateHotelShrinkRequest {
	s.RelatedPk = &v
	return s
}

func (s *CreateHotelShrinkRequest) SetRelatedPksShrink(v string) *CreateHotelShrinkRequest {
	s.RelatedPksShrink = &v
	return s
}

func (s *CreateHotelShrinkRequest) SetRemark(v string) *CreateHotelShrinkRequest {
	s.Remark = &v
	return s
}

func (s *CreateHotelShrinkRequest) SetRoomNum(v int32) *CreateHotelShrinkRequest {
	s.RoomNum = &v
	return s
}

func (s *CreateHotelShrinkRequest) SetTbOpenId(v string) *CreateHotelShrinkRequest {
	s.TbOpenId = &v
	return s
}

func (s *CreateHotelShrinkRequest) Validate() error {
	return dara.Validate(s)
}
