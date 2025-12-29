// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRcuSceneShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHotelId(v string) *UpdateRcuSceneShrinkRequest
	GetHotelId() *string
	SetSceneId(v string) *UpdateRcuSceneShrinkRequest
	GetSceneId() *string
	SetSceneRelationExtDTOShrink(v string) *UpdateRcuSceneShrinkRequest
	GetSceneRelationExtDTOShrink() *string
}

type UpdateRcuSceneShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 520a0c0***5eb
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// yoga
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// This parameter is required.
	SceneRelationExtDTOShrink *string `json:"SceneRelationExtDTO,omitempty" xml:"SceneRelationExtDTO,omitempty"`
}

func (s UpdateRcuSceneShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateRcuSceneShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateRcuSceneShrinkRequest) GetHotelId() *string {
	return s.HotelId
}

func (s *UpdateRcuSceneShrinkRequest) GetSceneId() *string {
	return s.SceneId
}

func (s *UpdateRcuSceneShrinkRequest) GetSceneRelationExtDTOShrink() *string {
	return s.SceneRelationExtDTOShrink
}

func (s *UpdateRcuSceneShrinkRequest) SetHotelId(v string) *UpdateRcuSceneShrinkRequest {
	s.HotelId = &v
	return s
}

func (s *UpdateRcuSceneShrinkRequest) SetSceneId(v string) *UpdateRcuSceneShrinkRequest {
	s.SceneId = &v
	return s
}

func (s *UpdateRcuSceneShrinkRequest) SetSceneRelationExtDTOShrink(v string) *UpdateRcuSceneShrinkRequest {
	s.SceneRelationExtDTOShrink = &v
	return s
}

func (s *UpdateRcuSceneShrinkRequest) Validate() error {
	return dara.Validate(s)
}
