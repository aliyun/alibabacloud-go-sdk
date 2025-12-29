// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRcuSceneRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHotelId(v string) *CreateRcuSceneRequest
	GetHotelId() *string
	SetSceneId(v string) *CreateRcuSceneRequest
	GetSceneId() *string
	SetSceneRelationExtDTO(v *CreateRcuSceneRequestSceneRelationExtDTO) *CreateRcuSceneRequest
	GetSceneRelationExtDTO() *CreateRcuSceneRequestSceneRelationExtDTO
}

type CreateRcuSceneRequest struct {
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
	SceneRelationExtDTO *CreateRcuSceneRequestSceneRelationExtDTO `json:"SceneRelationExtDTO,omitempty" xml:"SceneRelationExtDTO,omitempty" type:"Struct"`
}

func (s CreateRcuSceneRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRcuSceneRequest) GoString() string {
	return s.String()
}

func (s *CreateRcuSceneRequest) GetHotelId() *string {
	return s.HotelId
}

func (s *CreateRcuSceneRequest) GetSceneId() *string {
	return s.SceneId
}

func (s *CreateRcuSceneRequest) GetSceneRelationExtDTO() *CreateRcuSceneRequestSceneRelationExtDTO {
	return s.SceneRelationExtDTO
}

func (s *CreateRcuSceneRequest) SetHotelId(v string) *CreateRcuSceneRequest {
	s.HotelId = &v
	return s
}

func (s *CreateRcuSceneRequest) SetSceneId(v string) *CreateRcuSceneRequest {
	s.SceneId = &v
	return s
}

func (s *CreateRcuSceneRequest) SetSceneRelationExtDTO(v *CreateRcuSceneRequestSceneRelationExtDTO) *CreateRcuSceneRequest {
	s.SceneRelationExtDTO = v
	return s
}

func (s *CreateRcuSceneRequest) Validate() error {
	if s.SceneRelationExtDTO != nil {
		if err := s.SceneRelationExtDTO.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateRcuSceneRequestSceneRelationExtDTO struct {
	// This parameter is required.
	CorpusList []*string `json:"CorpusList,omitempty" xml:"CorpusList,omitempty" type:"Repeated"`
	// This parameter is required.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// https://ailabsaicloudservice.alicdn.com/hotel/icon/changjingmoshi/shuimian.png
	Icon *string `json:"Icon,omitempty" xml:"Icon,omitempty"`
	// This parameter is required.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateRcuSceneRequestSceneRelationExtDTO) String() string {
	return dara.Prettify(s)
}

func (s CreateRcuSceneRequestSceneRelationExtDTO) GoString() string {
	return s.String()
}

func (s *CreateRcuSceneRequestSceneRelationExtDTO) GetCorpusList() []*string {
	return s.CorpusList
}

func (s *CreateRcuSceneRequestSceneRelationExtDTO) GetDescription() *string {
	return s.Description
}

func (s *CreateRcuSceneRequestSceneRelationExtDTO) GetIcon() *string {
	return s.Icon
}

func (s *CreateRcuSceneRequestSceneRelationExtDTO) GetName() *string {
	return s.Name
}

func (s *CreateRcuSceneRequestSceneRelationExtDTO) SetCorpusList(v []*string) *CreateRcuSceneRequestSceneRelationExtDTO {
	s.CorpusList = v
	return s
}

func (s *CreateRcuSceneRequestSceneRelationExtDTO) SetDescription(v string) *CreateRcuSceneRequestSceneRelationExtDTO {
	s.Description = &v
	return s
}

func (s *CreateRcuSceneRequestSceneRelationExtDTO) SetIcon(v string) *CreateRcuSceneRequestSceneRelationExtDTO {
	s.Icon = &v
	return s
}

func (s *CreateRcuSceneRequestSceneRelationExtDTO) SetName(v string) *CreateRcuSceneRequestSceneRelationExtDTO {
	s.Name = &v
	return s
}

func (s *CreateRcuSceneRequestSceneRelationExtDTO) Validate() error {
	return dara.Validate(s)
}
