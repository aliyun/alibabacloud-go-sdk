// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHotelSceneBookItemsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHotelId(v string) *ListHotelSceneBookItemsShrinkRequest
	GetHotelId() *string
	SetPageShrink(v string) *ListHotelSceneBookItemsShrinkRequest
	GetPageShrink() *string
	SetType(v string) *ListHotelSceneBookItemsShrinkRequest
	GetType() *string
}

type ListHotelSceneBookItemsShrinkRequest struct {
	// hotelID
	//
	// This parameter is required.
	//
	// example:
	//
	// 80d84ea8ed9e422fbad52715c8fc56f1
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	// This parameter is required.
	PageShrink *string `json:"Page,omitempty" xml:"Page,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// FOOD
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListHotelSceneBookItemsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListHotelSceneBookItemsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListHotelSceneBookItemsShrinkRequest) GetHotelId() *string {
	return s.HotelId
}

func (s *ListHotelSceneBookItemsShrinkRequest) GetPageShrink() *string {
	return s.PageShrink
}

func (s *ListHotelSceneBookItemsShrinkRequest) GetType() *string {
	return s.Type
}

func (s *ListHotelSceneBookItemsShrinkRequest) SetHotelId(v string) *ListHotelSceneBookItemsShrinkRequest {
	s.HotelId = &v
	return s
}

func (s *ListHotelSceneBookItemsShrinkRequest) SetPageShrink(v string) *ListHotelSceneBookItemsShrinkRequest {
	s.PageShrink = &v
	return s
}

func (s *ListHotelSceneBookItemsShrinkRequest) SetType(v string) *ListHotelSceneBookItemsShrinkRequest {
	s.Type = &v
	return s
}

func (s *ListHotelSceneBookItemsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
