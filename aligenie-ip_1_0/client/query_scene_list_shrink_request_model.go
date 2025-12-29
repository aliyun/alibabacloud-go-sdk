// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySceneListShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHotelId(v string) *QuerySceneListShrinkRequest
	GetHotelId() *string
	SetSceneStatesShrink(v string) *QuerySceneListShrinkRequest
	GetSceneStatesShrink() *string
	SetSceneTypesShrink(v string) *QuerySceneListShrinkRequest
	GetSceneTypesShrink() *string
	SetTemplateInfoIdsShrink(v string) *QuerySceneListShrinkRequest
	GetTemplateInfoIdsShrink() *string
}

type QuerySceneListShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// af7***536
	HotelId               *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	SceneStatesShrink     *string `json:"SceneStates,omitempty" xml:"SceneStates,omitempty"`
	SceneTypesShrink      *string `json:"SceneTypes,omitempty" xml:"SceneTypes,omitempty"`
	TemplateInfoIdsShrink *string `json:"TemplateInfoIds,omitempty" xml:"TemplateInfoIds,omitempty"`
}

func (s QuerySceneListShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s QuerySceneListShrinkRequest) GoString() string {
	return s.String()
}

func (s *QuerySceneListShrinkRequest) GetHotelId() *string {
	return s.HotelId
}

func (s *QuerySceneListShrinkRequest) GetSceneStatesShrink() *string {
	return s.SceneStatesShrink
}

func (s *QuerySceneListShrinkRequest) GetSceneTypesShrink() *string {
	return s.SceneTypesShrink
}

func (s *QuerySceneListShrinkRequest) GetTemplateInfoIdsShrink() *string {
	return s.TemplateInfoIdsShrink
}

func (s *QuerySceneListShrinkRequest) SetHotelId(v string) *QuerySceneListShrinkRequest {
	s.HotelId = &v
	return s
}

func (s *QuerySceneListShrinkRequest) SetSceneStatesShrink(v string) *QuerySceneListShrinkRequest {
	s.SceneStatesShrink = &v
	return s
}

func (s *QuerySceneListShrinkRequest) SetSceneTypesShrink(v string) *QuerySceneListShrinkRequest {
	s.SceneTypesShrink = &v
	return s
}

func (s *QuerySceneListShrinkRequest) SetTemplateInfoIdsShrink(v string) *QuerySceneListShrinkRequest {
	s.TemplateInfoIdsShrink = &v
	return s
}

func (s *QuerySceneListShrinkRequest) Validate() error {
	return dara.Validate(s)
}
