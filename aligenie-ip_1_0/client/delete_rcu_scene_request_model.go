// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRcuSceneRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHotelId(v string) *DeleteRcuSceneRequest
	GetHotelId() *string
	SetSceneId(v string) *DeleteRcuSceneRequest
	GetSceneId() *string
}

type DeleteRcuSceneRequest struct {
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
	// yoga
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s DeleteRcuSceneRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteRcuSceneRequest) GoString() string {
	return s.String()
}

func (s *DeleteRcuSceneRequest) GetHotelId() *string {
	return s.HotelId
}

func (s *DeleteRcuSceneRequest) GetSceneId() *string {
	return s.SceneId
}

func (s *DeleteRcuSceneRequest) SetHotelId(v string) *DeleteRcuSceneRequest {
	s.HotelId = &v
	return s
}

func (s *DeleteRcuSceneRequest) SetSceneId(v string) *DeleteRcuSceneRequest {
	s.SceneId = &v
	return s
}

func (s *DeleteRcuSceneRequest) Validate() error {
	return dara.Validate(s)
}
