// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHotelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppKey(v string) *CreateHotelRequest
	GetAppKey() *string
	SetEstOpenTime(v string) *CreateHotelRequest
	GetEstOpenTime() *string
	SetHotelAddress(v string) *CreateHotelRequest
	GetHotelAddress() *string
	SetHotelEmail(v string) *CreateHotelRequest
	GetHotelEmail() *string
	SetHotelName(v string) *CreateHotelRequest
	GetHotelName() *string
	SetPhoneNumber(v string) *CreateHotelRequest
	GetPhoneNumber() *string
	SetRelatedPk(v string) *CreateHotelRequest
	GetRelatedPk() *string
	SetRelatedPks(v []*string) *CreateHotelRequest
	GetRelatedPks() []*string
	SetRemark(v string) *CreateHotelRequest
	GetRemark() *string
	SetRoomNum(v int32) *CreateHotelRequest
	GetRoomNum() *int32
	SetTbOpenId(v string) *CreateHotelRequest
	GetTbOpenId() *string
}

type CreateHotelRequest struct {
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
	RelatedPks []*string `json:"RelatedPks,omitempty" xml:"RelatedPks,omitempty" type:"Repeated"`
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

func (s CreateHotelRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateHotelRequest) GoString() string {
	return s.String()
}

func (s *CreateHotelRequest) GetAppKey() *string {
	return s.AppKey
}

func (s *CreateHotelRequest) GetEstOpenTime() *string {
	return s.EstOpenTime
}

func (s *CreateHotelRequest) GetHotelAddress() *string {
	return s.HotelAddress
}

func (s *CreateHotelRequest) GetHotelEmail() *string {
	return s.HotelEmail
}

func (s *CreateHotelRequest) GetHotelName() *string {
	return s.HotelName
}

func (s *CreateHotelRequest) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *CreateHotelRequest) GetRelatedPk() *string {
	return s.RelatedPk
}

func (s *CreateHotelRequest) GetRelatedPks() []*string {
	return s.RelatedPks
}

func (s *CreateHotelRequest) GetRemark() *string {
	return s.Remark
}

func (s *CreateHotelRequest) GetRoomNum() *int32 {
	return s.RoomNum
}

func (s *CreateHotelRequest) GetTbOpenId() *string {
	return s.TbOpenId
}

func (s *CreateHotelRequest) SetAppKey(v string) *CreateHotelRequest {
	s.AppKey = &v
	return s
}

func (s *CreateHotelRequest) SetEstOpenTime(v string) *CreateHotelRequest {
	s.EstOpenTime = &v
	return s
}

func (s *CreateHotelRequest) SetHotelAddress(v string) *CreateHotelRequest {
	s.HotelAddress = &v
	return s
}

func (s *CreateHotelRequest) SetHotelEmail(v string) *CreateHotelRequest {
	s.HotelEmail = &v
	return s
}

func (s *CreateHotelRequest) SetHotelName(v string) *CreateHotelRequest {
	s.HotelName = &v
	return s
}

func (s *CreateHotelRequest) SetPhoneNumber(v string) *CreateHotelRequest {
	s.PhoneNumber = &v
	return s
}

func (s *CreateHotelRequest) SetRelatedPk(v string) *CreateHotelRequest {
	s.RelatedPk = &v
	return s
}

func (s *CreateHotelRequest) SetRelatedPks(v []*string) *CreateHotelRequest {
	s.RelatedPks = v
	return s
}

func (s *CreateHotelRequest) SetRemark(v string) *CreateHotelRequest {
	s.Remark = &v
	return s
}

func (s *CreateHotelRequest) SetRoomNum(v int32) *CreateHotelRequest {
	s.RoomNum = &v
	return s
}

func (s *CreateHotelRequest) SetTbOpenId(v string) *CreateHotelRequest {
	s.TbOpenId = &v
	return s
}

func (s *CreateHotelRequest) Validate() error {
	return dara.Validate(s)
}
