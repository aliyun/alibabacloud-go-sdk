// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRcuSceneShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHotelId(v string) *CreateRcuSceneShrinkRequest
	GetHotelId() *string
	SetSceneId(v string) *CreateRcuSceneShrinkRequest
	GetSceneId() *string
	SetSceneRelationExtDTOShrink(v string) *CreateRcuSceneShrinkRequest
	GetSceneRelationExtDTOShrink() *string
}

type CreateRcuSceneShrinkRequest struct {
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

func (s CreateRcuSceneShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRcuSceneShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateRcuSceneShrinkRequest) GetHotelId() *string {
	return s.HotelId
}

func (s *CreateRcuSceneShrinkRequest) GetSceneId() *string {
	return s.SceneId
}

func (s *CreateRcuSceneShrinkRequest) GetSceneRelationExtDTOShrink() *string {
	return s.SceneRelationExtDTOShrink
}

func (s *CreateRcuSceneShrinkRequest) SetHotelId(v string) *CreateRcuSceneShrinkRequest {
	s.HotelId = &v
	return s
}

func (s *CreateRcuSceneShrinkRequest) SetSceneId(v string) *CreateRcuSceneShrinkRequest {
	s.SceneId = &v
	return s
}

func (s *CreateRcuSceneShrinkRequest) SetSceneRelationExtDTOShrink(v string) *CreateRcuSceneShrinkRequest {
	s.SceneRelationExtDTOShrink = &v
	return s
}

func (s *CreateRcuSceneShrinkRequest) Validate() error {
	return dara.Validate(s)
}
