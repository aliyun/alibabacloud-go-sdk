// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateHotelSceneItemShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHotelId(v string) *UpdateHotelSceneItemShrinkRequest
	GetHotelId() *string
	SetUpdateHotelSceneOperateReqShrink(v string) *UpdateHotelSceneItemShrinkRequest
	GetUpdateHotelSceneOperateReqShrink() *string
	SetUpdateHotelSceneReqShrink(v string) *UpdateHotelSceneItemShrinkRequest
	GetUpdateHotelSceneReqShrink() *string
}

type UpdateHotelSceneItemShrinkRequest struct {
	// hotelID
	//
	// This parameter is required.
	//
	// example:
	//
	// 80d84ea8ed9e422fbad52715c8fc56f1
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	// updateHotelSceneReq
	//
	// This parameter is required.
	UpdateHotelSceneOperateReqShrink *string `json:"UpdateHotelSceneOperateReq,omitempty" xml:"UpdateHotelSceneOperateReq,omitempty"`
	// UpdateHotelSceneReq
	//
	// This parameter is required.
	UpdateHotelSceneReqShrink *string `json:"UpdateHotelSceneReq,omitempty" xml:"UpdateHotelSceneReq,omitempty"`
}

func (s UpdateHotelSceneItemShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateHotelSceneItemShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateHotelSceneItemShrinkRequest) GetHotelId() *string {
	return s.HotelId
}

func (s *UpdateHotelSceneItemShrinkRequest) GetUpdateHotelSceneOperateReqShrink() *string {
	return s.UpdateHotelSceneOperateReqShrink
}

func (s *UpdateHotelSceneItemShrinkRequest) GetUpdateHotelSceneReqShrink() *string {
	return s.UpdateHotelSceneReqShrink
}

func (s *UpdateHotelSceneItemShrinkRequest) SetHotelId(v string) *UpdateHotelSceneItemShrinkRequest {
	s.HotelId = &v
	return s
}

func (s *UpdateHotelSceneItemShrinkRequest) SetUpdateHotelSceneOperateReqShrink(v string) *UpdateHotelSceneItemShrinkRequest {
	s.UpdateHotelSceneOperateReqShrink = &v
	return s
}

func (s *UpdateHotelSceneItemShrinkRequest) SetUpdateHotelSceneReqShrink(v string) *UpdateHotelSceneItemShrinkRequest {
	s.UpdateHotelSceneReqShrink = &v
	return s
}

func (s *UpdateHotelSceneItemShrinkRequest) Validate() error {
	return dara.Validate(s)
}
