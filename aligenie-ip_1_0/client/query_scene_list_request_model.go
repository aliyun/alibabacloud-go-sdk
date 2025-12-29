// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySceneListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHotelId(v string) *QuerySceneListRequest
	GetHotelId() *string
	SetSceneStates(v []*int32) *QuerySceneListRequest
	GetSceneStates() []*int32
	SetSceneTypes(v []*string) *QuerySceneListRequest
	GetSceneTypes() []*string
	SetTemplateInfoIds(v []*string) *QuerySceneListRequest
	GetTemplateInfoIds() []*string
}

type QuerySceneListRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// af7***536
	HotelId         *string   `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	SceneStates     []*int32  `json:"SceneStates,omitempty" xml:"SceneStates,omitempty" type:"Repeated"`
	SceneTypes      []*string `json:"SceneTypes,omitempty" xml:"SceneTypes,omitempty" type:"Repeated"`
	TemplateInfoIds []*string `json:"TemplateInfoIds,omitempty" xml:"TemplateInfoIds,omitempty" type:"Repeated"`
}

func (s QuerySceneListRequest) String() string {
	return dara.Prettify(s)
}

func (s QuerySceneListRequest) GoString() string {
	return s.String()
}

func (s *QuerySceneListRequest) GetHotelId() *string {
	return s.HotelId
}

func (s *QuerySceneListRequest) GetSceneStates() []*int32 {
	return s.SceneStates
}

func (s *QuerySceneListRequest) GetSceneTypes() []*string {
	return s.SceneTypes
}

func (s *QuerySceneListRequest) GetTemplateInfoIds() []*string {
	return s.TemplateInfoIds
}

func (s *QuerySceneListRequest) SetHotelId(v string) *QuerySceneListRequest {
	s.HotelId = &v
	return s
}

func (s *QuerySceneListRequest) SetSceneStates(v []*int32) *QuerySceneListRequest {
	s.SceneStates = v
	return s
}

func (s *QuerySceneListRequest) SetSceneTypes(v []*string) *QuerySceneListRequest {
	s.SceneTypes = v
	return s
}

func (s *QuerySceneListRequest) SetTemplateInfoIds(v []*string) *QuerySceneListRequest {
	s.TemplateInfoIds = v
	return s
}

func (s *QuerySceneListRequest) Validate() error {
	return dara.Validate(s)
}
