// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotelSceneItemDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHotelId(v string) *GetHotelSceneItemDetailRequest
	GetHotelId() *string
	SetItemId(v int64) *GetHotelSceneItemDetailRequest
	GetItemId() *int64
	SetName(v string) *GetHotelSceneItemDetailRequest
	GetName() *string
}

type GetHotelSceneItemDetailRequest struct {
	// hotelID
	//
	// This parameter is required.
	//
	// example:
	//
	// 80d84ea8ed9e422fbad52715c8fc56f1
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	// example:
	//
	// 10336
	ItemId *int64  `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	Name   *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetHotelSceneItemDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s GetHotelSceneItemDetailRequest) GoString() string {
	return s.String()
}

func (s *GetHotelSceneItemDetailRequest) GetHotelId() *string {
	return s.HotelId
}

func (s *GetHotelSceneItemDetailRequest) GetItemId() *int64 {
	return s.ItemId
}

func (s *GetHotelSceneItemDetailRequest) GetName() *string {
	return s.Name
}

func (s *GetHotelSceneItemDetailRequest) SetHotelId(v string) *GetHotelSceneItemDetailRequest {
	s.HotelId = &v
	return s
}

func (s *GetHotelSceneItemDetailRequest) SetItemId(v int64) *GetHotelSceneItemDetailRequest {
	s.ItemId = &v
	return s
}

func (s *GetHotelSceneItemDetailRequest) SetName(v string) *GetHotelSceneItemDetailRequest {
	s.Name = &v
	return s
}

func (s *GetHotelSceneItemDetailRequest) Validate() error {
	return dara.Validate(s)
}
