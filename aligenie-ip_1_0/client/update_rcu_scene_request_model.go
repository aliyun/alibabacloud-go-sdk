// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRcuSceneRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHotelId(v string) *UpdateRcuSceneRequest
	GetHotelId() *string
	SetSceneId(v string) *UpdateRcuSceneRequest
	GetSceneId() *string
	SetSceneRelationExtDTO(v *UpdateRcuSceneRequestSceneRelationExtDTO) *UpdateRcuSceneRequest
	GetSceneRelationExtDTO() *UpdateRcuSceneRequestSceneRelationExtDTO
}

type UpdateRcuSceneRequest struct {
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
	SceneRelationExtDTO *UpdateRcuSceneRequestSceneRelationExtDTO `json:"SceneRelationExtDTO,omitempty" xml:"SceneRelationExtDTO,omitempty" type:"Struct"`
}

func (s UpdateRcuSceneRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateRcuSceneRequest) GoString() string {
	return s.String()
}

func (s *UpdateRcuSceneRequest) GetHotelId() *string {
	return s.HotelId
}

func (s *UpdateRcuSceneRequest) GetSceneId() *string {
	return s.SceneId
}

func (s *UpdateRcuSceneRequest) GetSceneRelationExtDTO() *UpdateRcuSceneRequestSceneRelationExtDTO {
	return s.SceneRelationExtDTO
}

func (s *UpdateRcuSceneRequest) SetHotelId(v string) *UpdateRcuSceneRequest {
	s.HotelId = &v
	return s
}

func (s *UpdateRcuSceneRequest) SetSceneId(v string) *UpdateRcuSceneRequest {
	s.SceneId = &v
	return s
}

func (s *UpdateRcuSceneRequest) SetSceneRelationExtDTO(v *UpdateRcuSceneRequestSceneRelationExtDTO) *UpdateRcuSceneRequest {
	s.SceneRelationExtDTO = v
	return s
}

func (s *UpdateRcuSceneRequest) Validate() error {
	if s.SceneRelationExtDTO != nil {
		if err := s.SceneRelationExtDTO.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateRcuSceneRequestSceneRelationExtDTO struct {
	CorpusList  []*string `json:"CorpusList,omitempty" xml:"CorpusList,omitempty" type:"Repeated"`
	Description *string   `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// https://ailabsaicloudservice.alicdn.com/hotel/icon/changjingmoshi/shuimian.png
	Icon *string `json:"Icon,omitempty" xml:"Icon,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateRcuSceneRequestSceneRelationExtDTO) String() string {
	return dara.Prettify(s)
}

func (s UpdateRcuSceneRequestSceneRelationExtDTO) GoString() string {
	return s.String()
}

func (s *UpdateRcuSceneRequestSceneRelationExtDTO) GetCorpusList() []*string {
	return s.CorpusList
}

func (s *UpdateRcuSceneRequestSceneRelationExtDTO) GetDescription() *string {
	return s.Description
}

func (s *UpdateRcuSceneRequestSceneRelationExtDTO) GetIcon() *string {
	return s.Icon
}

func (s *UpdateRcuSceneRequestSceneRelationExtDTO) GetName() *string {
	return s.Name
}

func (s *UpdateRcuSceneRequestSceneRelationExtDTO) SetCorpusList(v []*string) *UpdateRcuSceneRequestSceneRelationExtDTO {
	s.CorpusList = v
	return s
}

func (s *UpdateRcuSceneRequestSceneRelationExtDTO) SetDescription(v string) *UpdateRcuSceneRequestSceneRelationExtDTO {
	s.Description = &v
	return s
}

func (s *UpdateRcuSceneRequestSceneRelationExtDTO) SetIcon(v string) *UpdateRcuSceneRequestSceneRelationExtDTO {
	s.Icon = &v
	return s
}

func (s *UpdateRcuSceneRequestSceneRelationExtDTO) SetName(v string) *UpdateRcuSceneRequestSceneRelationExtDTO {
	s.Name = &v
	return s
}

func (s *UpdateRcuSceneRequestSceneRelationExtDTO) Validate() error {
	return dara.Validate(s)
}
