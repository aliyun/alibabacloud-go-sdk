// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateProjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoBuild(v bool) *CreateProjectRequest
	GetAutoBuild() *bool
	SetDependencies(v string) *CreateProjectRequest
	GetDependencies() *string
	SetExtInfo(v string) *CreateProjectRequest
	GetExtInfo() *string
	SetGender(v string) *CreateProjectRequest
	GetGender() *string
	SetHeight(v float64) *CreateProjectRequest
	GetHeight() *float64
	SetIntro(v string) *CreateProjectRequest
	GetIntro() *string
	SetJwtToken(v string) *CreateProjectRequest
	GetJwtToken() *string
	SetMapUuid(v string) *CreateProjectRequest
	GetMapUuid() *string
	SetMethod(v string) *CreateProjectRequest
	GetMethod() *string
	SetMode(v string) *CreateProjectRequest
	GetMode() *string
	SetTitle(v string) *CreateProjectRequest
	GetTitle() *string
	SetTryOnParams(v *CreateProjectRequestTryOnParams) *CreateProjectRequest
	GetTryOnParams() *CreateProjectRequestTryOnParams
	SetType(v string) *CreateProjectRequest
	GetType() *string
	SetWeight(v float64) *CreateProjectRequest
	GetWeight() *float64
}

type CreateProjectRequest struct {
	AutoBuild    *bool                            `json:"AutoBuild,omitempty" xml:"AutoBuild,omitempty"`
	Dependencies *string                          `json:"Dependencies,omitempty" xml:"Dependencies,omitempty"`
	ExtInfo      *string                          `json:"ExtInfo,omitempty" xml:"ExtInfo,omitempty"`
	Gender       *string                          `json:"Gender,omitempty" xml:"Gender,omitempty"`
	Height       *float64                         `json:"Height,omitempty" xml:"Height,omitempty"`
	Intro        *string                          `json:"Intro,omitempty" xml:"Intro,omitempty"`
	JwtToken     *string                          `json:"JwtToken,omitempty" xml:"JwtToken,omitempty"`
	MapUuid      *string                          `json:"MapUuid,omitempty" xml:"MapUuid,omitempty"`
	Method       *string                          `json:"Method,omitempty" xml:"Method,omitempty"`
	Mode         *string                          `json:"Mode,omitempty" xml:"Mode,omitempty"`
	Title        *string                          `json:"Title,omitempty" xml:"Title,omitempty"`
	TryOnParams  *CreateProjectRequestTryOnParams `json:"TryOnParams,omitempty" xml:"TryOnParams,omitempty" type:"Struct"`
	Type         *string                          `json:"Type,omitempty" xml:"Type,omitempty"`
	Weight       *float64                         `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s CreateProjectRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateProjectRequest) GoString() string {
	return s.String()
}

func (s *CreateProjectRequest) GetAutoBuild() *bool {
	return s.AutoBuild
}

func (s *CreateProjectRequest) GetDependencies() *string {
	return s.Dependencies
}

func (s *CreateProjectRequest) GetExtInfo() *string {
	return s.ExtInfo
}

func (s *CreateProjectRequest) GetGender() *string {
	return s.Gender
}

func (s *CreateProjectRequest) GetHeight() *float64 {
	return s.Height
}

func (s *CreateProjectRequest) GetIntro() *string {
	return s.Intro
}

func (s *CreateProjectRequest) GetJwtToken() *string {
	return s.JwtToken
}

func (s *CreateProjectRequest) GetMapUuid() *string {
	return s.MapUuid
}

func (s *CreateProjectRequest) GetMethod() *string {
	return s.Method
}

func (s *CreateProjectRequest) GetMode() *string {
	return s.Mode
}

func (s *CreateProjectRequest) GetTitle() *string {
	return s.Title
}

func (s *CreateProjectRequest) GetTryOnParams() *CreateProjectRequestTryOnParams {
	return s.TryOnParams
}

func (s *CreateProjectRequest) GetType() *string {
	return s.Type
}

func (s *CreateProjectRequest) GetWeight() *float64 {
	return s.Weight
}

func (s *CreateProjectRequest) SetAutoBuild(v bool) *CreateProjectRequest {
	s.AutoBuild = &v
	return s
}

func (s *CreateProjectRequest) SetDependencies(v string) *CreateProjectRequest {
	s.Dependencies = &v
	return s
}

func (s *CreateProjectRequest) SetExtInfo(v string) *CreateProjectRequest {
	s.ExtInfo = &v
	return s
}

func (s *CreateProjectRequest) SetGender(v string) *CreateProjectRequest {
	s.Gender = &v
	return s
}

func (s *CreateProjectRequest) SetHeight(v float64) *CreateProjectRequest {
	s.Height = &v
	return s
}

func (s *CreateProjectRequest) SetIntro(v string) *CreateProjectRequest {
	s.Intro = &v
	return s
}

func (s *CreateProjectRequest) SetJwtToken(v string) *CreateProjectRequest {
	s.JwtToken = &v
	return s
}

func (s *CreateProjectRequest) SetMapUuid(v string) *CreateProjectRequest {
	s.MapUuid = &v
	return s
}

func (s *CreateProjectRequest) SetMethod(v string) *CreateProjectRequest {
	s.Method = &v
	return s
}

func (s *CreateProjectRequest) SetMode(v string) *CreateProjectRequest {
	s.Mode = &v
	return s
}

func (s *CreateProjectRequest) SetTitle(v string) *CreateProjectRequest {
	s.Title = &v
	return s
}

func (s *CreateProjectRequest) SetTryOnParams(v *CreateProjectRequestTryOnParams) *CreateProjectRequest {
	s.TryOnParams = v
	return s
}

func (s *CreateProjectRequest) SetType(v string) *CreateProjectRequest {
	s.Type = &v
	return s
}

func (s *CreateProjectRequest) SetWeight(v float64) *CreateProjectRequest {
	s.Weight = &v
	return s
}

func (s *CreateProjectRequest) Validate() error {
	if s.TryOnParams != nil {
		if err := s.TryOnParams.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateProjectRequestTryOnParams struct {
	ClothIds   []*int64                                     `json:"ClothIds,omitempty" xml:"ClothIds,omitempty" type:"Repeated"`
	ClothInfos []*CreateProjectRequestTryOnParamsClothInfos `json:"ClothInfos,omitempty" xml:"ClothInfos,omitempty" type:"Repeated"`
}

func (s CreateProjectRequestTryOnParams) String() string {
	return dara.Prettify(s)
}

func (s CreateProjectRequestTryOnParams) GoString() string {
	return s.String()
}

func (s *CreateProjectRequestTryOnParams) GetClothIds() []*int64 {
	return s.ClothIds
}

func (s *CreateProjectRequestTryOnParams) GetClothInfos() []*CreateProjectRequestTryOnParamsClothInfos {
	return s.ClothInfos
}

func (s *CreateProjectRequestTryOnParams) SetClothIds(v []*int64) *CreateProjectRequestTryOnParams {
	s.ClothIds = v
	return s
}

func (s *CreateProjectRequestTryOnParams) SetClothInfos(v []*CreateProjectRequestTryOnParamsClothInfos) *CreateProjectRequestTryOnParams {
	s.ClothInfos = v
	return s
}

func (s *CreateProjectRequestTryOnParams) Validate() error {
	if s.ClothInfos != nil {
		for _, item := range s.ClothInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateProjectRequestTryOnParamsClothInfos struct {
	// This parameter is required.
	Id     *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Size   *string `json:"Size,omitempty" xml:"Size,omitempty"`
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreateProjectRequestTryOnParamsClothInfos) String() string {
	return dara.Prettify(s)
}

func (s CreateProjectRequestTryOnParamsClothInfos) GoString() string {
	return s.String()
}

func (s *CreateProjectRequestTryOnParamsClothInfos) GetId() *int64 {
	return s.Id
}

func (s *CreateProjectRequestTryOnParamsClothInfos) GetSize() *string {
	return s.Size
}

func (s *CreateProjectRequestTryOnParamsClothInfos) GetStatus() *string {
	return s.Status
}

func (s *CreateProjectRequestTryOnParamsClothInfos) SetId(v int64) *CreateProjectRequestTryOnParamsClothInfos {
	s.Id = &v
	return s
}

func (s *CreateProjectRequestTryOnParamsClothInfos) SetSize(v string) *CreateProjectRequestTryOnParamsClothInfos {
	s.Size = &v
	return s
}

func (s *CreateProjectRequestTryOnParamsClothInfos) SetStatus(v string) *CreateProjectRequestTryOnParamsClothInfos {
	s.Status = &v
	return s
}

func (s *CreateProjectRequestTryOnParamsClothInfos) Validate() error {
	return dara.Validate(s)
}
