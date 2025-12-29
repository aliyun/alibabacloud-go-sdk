// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInsertHotelSceneBookItemShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddHotelSceneItemReqShrink(v string) *InsertHotelSceneBookItemShrinkRequest
	GetAddHotelSceneItemReqShrink() *string
	SetHotelId(v string) *InsertHotelSceneBookItemShrinkRequest
	GetHotelId() *string
}

type InsertHotelSceneBookItemShrinkRequest struct {
	// addHotelSceneItemReq
	//
	// This parameter is required.
	AddHotelSceneItemReqShrink *string `json:"AddHotelSceneItemReq,omitempty" xml:"AddHotelSceneItemReq,omitempty"`
	// hotelID
	//
	// This parameter is required.
	//
	// example:
	//
	// 80d84ea8ed9e422fbad52715c8fc56f1
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
}

func (s InsertHotelSceneBookItemShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s InsertHotelSceneBookItemShrinkRequest) GoString() string {
	return s.String()
}

func (s *InsertHotelSceneBookItemShrinkRequest) GetAddHotelSceneItemReqShrink() *string {
	return s.AddHotelSceneItemReqShrink
}

func (s *InsertHotelSceneBookItemShrinkRequest) GetHotelId() *string {
	return s.HotelId
}

func (s *InsertHotelSceneBookItemShrinkRequest) SetAddHotelSceneItemReqShrink(v string) *InsertHotelSceneBookItemShrinkRequest {
	s.AddHotelSceneItemReqShrink = &v
	return s
}

func (s *InsertHotelSceneBookItemShrinkRequest) SetHotelId(v string) *InsertHotelSceneBookItemShrinkRequest {
	s.HotelId = &v
	return s
}

func (s *InsertHotelSceneBookItemShrinkRequest) Validate() error {
	return dara.Validate(s)
}
