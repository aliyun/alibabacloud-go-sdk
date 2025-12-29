// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportRoomGenieScenesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHotelId(v string) *ImportRoomGenieScenesShrinkRequest
	GetHotelId() *string
	SetRoomNo(v string) *ImportRoomGenieScenesShrinkRequest
	GetRoomNo() *string
	SetSceneListShrink(v string) *ImportRoomGenieScenesShrinkRequest
	GetSceneListShrink() *string
}

type ImportRoomGenieScenesShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// a7a3***013
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1211
	RoomNo          *string `json:"RoomNo,omitempty" xml:"RoomNo,omitempty"`
	SceneListShrink *string `json:"SceneList,omitempty" xml:"SceneList,omitempty"`
}

func (s ImportRoomGenieScenesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ImportRoomGenieScenesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ImportRoomGenieScenesShrinkRequest) GetHotelId() *string {
	return s.HotelId
}

func (s *ImportRoomGenieScenesShrinkRequest) GetRoomNo() *string {
	return s.RoomNo
}

func (s *ImportRoomGenieScenesShrinkRequest) GetSceneListShrink() *string {
	return s.SceneListShrink
}

func (s *ImportRoomGenieScenesShrinkRequest) SetHotelId(v string) *ImportRoomGenieScenesShrinkRequest {
	s.HotelId = &v
	return s
}

func (s *ImportRoomGenieScenesShrinkRequest) SetRoomNo(v string) *ImportRoomGenieScenesShrinkRequest {
	s.RoomNo = &v
	return s
}

func (s *ImportRoomGenieScenesShrinkRequest) SetSceneListShrink(v string) *ImportRoomGenieScenesShrinkRequest {
	s.SceneListShrink = &v
	return s
}

func (s *ImportRoomGenieScenesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
