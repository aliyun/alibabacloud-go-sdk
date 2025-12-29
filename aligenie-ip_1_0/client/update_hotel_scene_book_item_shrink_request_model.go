// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateHotelSceneBookItemShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHotelId(v string) *UpdateHotelSceneBookItemShrinkRequest
	GetHotelId() *string
	SetUpdateHotelSceneBookReqShrink(v string) *UpdateHotelSceneBookItemShrinkRequest
	GetUpdateHotelSceneBookReqShrink() *string
}

type UpdateHotelSceneBookItemShrinkRequest struct {
	// hotelID
	//
	// This parameter is required.
	//
	// example:
	//
	// 80d84ea8ed9e422fbad52715c8fc56f1
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	// updateHotelSceneBookReq
	//
	// This parameter is required.
	UpdateHotelSceneBookReqShrink *string `json:"UpdateHotelSceneBookReq,omitempty" xml:"UpdateHotelSceneBookReq,omitempty"`
}

func (s UpdateHotelSceneBookItemShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateHotelSceneBookItemShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateHotelSceneBookItemShrinkRequest) GetHotelId() *string {
	return s.HotelId
}

func (s *UpdateHotelSceneBookItemShrinkRequest) GetUpdateHotelSceneBookReqShrink() *string {
	return s.UpdateHotelSceneBookReqShrink
}

func (s *UpdateHotelSceneBookItemShrinkRequest) SetHotelId(v string) *UpdateHotelSceneBookItemShrinkRequest {
	s.HotelId = &v
	return s
}

func (s *UpdateHotelSceneBookItemShrinkRequest) SetUpdateHotelSceneBookReqShrink(v string) *UpdateHotelSceneBookItemShrinkRequest {
	s.UpdateHotelSceneBookReqShrink = &v
	return s
}

func (s *UpdateHotelSceneBookItemShrinkRequest) Validate() error {
	return dara.Validate(s)
}
