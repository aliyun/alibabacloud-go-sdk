// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateHotelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppKey(v string) *UpdateHotelRequest
	GetAppKey() *string
	SetEstOpenTime(v string) *UpdateHotelRequest
	GetEstOpenTime() *string
	SetHotelAddress(v string) *UpdateHotelRequest
	GetHotelAddress() *string
	SetHotelEmail(v string) *UpdateHotelRequest
	GetHotelEmail() *string
	SetHotelId(v string) *UpdateHotelRequest
	GetHotelId() *string
	SetHotelName(v string) *UpdateHotelRequest
	GetHotelName() *string
	SetPhoneNumber(v string) *UpdateHotelRequest
	GetPhoneNumber() *string
	SetRelatedPks(v []*string) *UpdateHotelRequest
	GetRelatedPks() []*string
	SetRemark(v string) *UpdateHotelRequest
	GetRemark() *string
	SetRoomNum(v int32) *UpdateHotelRequest
	GetRoomNum() *int32
	SetTbOpenId(v string) *UpdateHotelRequest
	GetTbOpenId() *string
}

type UpdateHotelRequest struct {
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
	PhoneNumber *string   `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	RelatedPks  []*string `json:"RelatedPks,omitempty" xml:"RelatedPks,omitempty" type:"Repeated"`
	Remark      *string   `json:"Remark,omitempty" xml:"Remark,omitempty"`
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

func (s UpdateHotelRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateHotelRequest) GoString() string {
	return s.String()
}

func (s *UpdateHotelRequest) GetAppKey() *string {
	return s.AppKey
}

func (s *UpdateHotelRequest) GetEstOpenTime() *string {
	return s.EstOpenTime
}

func (s *UpdateHotelRequest) GetHotelAddress() *string {
	return s.HotelAddress
}

func (s *UpdateHotelRequest) GetHotelEmail() *string {
	return s.HotelEmail
}

func (s *UpdateHotelRequest) GetHotelId() *string {
	return s.HotelId
}

func (s *UpdateHotelRequest) GetHotelName() *string {
	return s.HotelName
}

func (s *UpdateHotelRequest) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *UpdateHotelRequest) GetRelatedPks() []*string {
	return s.RelatedPks
}

func (s *UpdateHotelRequest) GetRemark() *string {
	return s.Remark
}

func (s *UpdateHotelRequest) GetRoomNum() *int32 {
	return s.RoomNum
}

func (s *UpdateHotelRequest) GetTbOpenId() *string {
	return s.TbOpenId
}

func (s *UpdateHotelRequest) SetAppKey(v string) *UpdateHotelRequest {
	s.AppKey = &v
	return s
}

func (s *UpdateHotelRequest) SetEstOpenTime(v string) *UpdateHotelRequest {
	s.EstOpenTime = &v
	return s
}

func (s *UpdateHotelRequest) SetHotelAddress(v string) *UpdateHotelRequest {
	s.HotelAddress = &v
	return s
}

func (s *UpdateHotelRequest) SetHotelEmail(v string) *UpdateHotelRequest {
	s.HotelEmail = &v
	return s
}

func (s *UpdateHotelRequest) SetHotelId(v string) *UpdateHotelRequest {
	s.HotelId = &v
	return s
}

func (s *UpdateHotelRequest) SetHotelName(v string) *UpdateHotelRequest {
	s.HotelName = &v
	return s
}

func (s *UpdateHotelRequest) SetPhoneNumber(v string) *UpdateHotelRequest {
	s.PhoneNumber = &v
	return s
}

func (s *UpdateHotelRequest) SetRelatedPks(v []*string) *UpdateHotelRequest {
	s.RelatedPks = v
	return s
}

func (s *UpdateHotelRequest) SetRemark(v string) *UpdateHotelRequest {
	s.Remark = &v
	return s
}

func (s *UpdateHotelRequest) SetRoomNum(v int32) *UpdateHotelRequest {
	s.RoomNum = &v
	return s
}

func (s *UpdateHotelRequest) SetTbOpenId(v string) *UpdateHotelRequest {
	s.TbOpenId = &v
	return s
}

func (s *UpdateHotelRequest) Validate() error {
	return dara.Validate(s)
}
