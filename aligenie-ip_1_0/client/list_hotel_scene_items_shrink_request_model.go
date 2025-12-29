// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHotelSceneItemsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHotelId(v string) *ListHotelSceneItemsShrinkRequest
	GetHotelId() *string
	SetListHotelSceneReqShrink(v string) *ListHotelSceneItemsShrinkRequest
	GetListHotelSceneReqShrink() *string
}

type ListHotelSceneItemsShrinkRequest struct {
	// hotelID
	//
	// This parameter is required.
	//
	// example:
	//
	// 80d84ea8ed9e422fbad52715c8fc56f1
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	// ListHotelSceneReq
	//
	// This parameter is required.
	ListHotelSceneReqShrink *string `json:"ListHotelSceneReq,omitempty" xml:"ListHotelSceneReq,omitempty"`
}

func (s ListHotelSceneItemsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListHotelSceneItemsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListHotelSceneItemsShrinkRequest) GetHotelId() *string {
	return s.HotelId
}

func (s *ListHotelSceneItemsShrinkRequest) GetListHotelSceneReqShrink() *string {
	return s.ListHotelSceneReqShrink
}

func (s *ListHotelSceneItemsShrinkRequest) SetHotelId(v string) *ListHotelSceneItemsShrinkRequest {
	s.HotelId = &v
	return s
}

func (s *ListHotelSceneItemsShrinkRequest) SetListHotelSceneReqShrink(v string) *ListHotelSceneItemsShrinkRequest {
	s.ListHotelSceneReqShrink = &v
	return s
}

func (s *ListHotelSceneItemsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
