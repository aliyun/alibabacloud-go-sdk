// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateProjectShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoBuild(v bool) *CreateProjectShrinkRequest
	GetAutoBuild() *bool
	SetDependencies(v string) *CreateProjectShrinkRequest
	GetDependencies() *string
	SetExtInfo(v string) *CreateProjectShrinkRequest
	GetExtInfo() *string
	SetGender(v string) *CreateProjectShrinkRequest
	GetGender() *string
	SetHeight(v float64) *CreateProjectShrinkRequest
	GetHeight() *float64
	SetIntro(v string) *CreateProjectShrinkRequest
	GetIntro() *string
	SetJwtToken(v string) *CreateProjectShrinkRequest
	GetJwtToken() *string
	SetMapUuid(v string) *CreateProjectShrinkRequest
	GetMapUuid() *string
	SetMethod(v string) *CreateProjectShrinkRequest
	GetMethod() *string
	SetMode(v string) *CreateProjectShrinkRequest
	GetMode() *string
	SetTitle(v string) *CreateProjectShrinkRequest
	GetTitle() *string
	SetTryOnParamsShrink(v string) *CreateProjectShrinkRequest
	GetTryOnParamsShrink() *string
	SetType(v string) *CreateProjectShrinkRequest
	GetType() *string
	SetWeight(v float64) *CreateProjectShrinkRequest
	GetWeight() *float64
}

type CreateProjectShrinkRequest struct {
	AutoBuild         *bool    `json:"AutoBuild,omitempty" xml:"AutoBuild,omitempty"`
	Dependencies      *string  `json:"Dependencies,omitempty" xml:"Dependencies,omitempty"`
	ExtInfo           *string  `json:"ExtInfo,omitempty" xml:"ExtInfo,omitempty"`
	Gender            *string  `json:"Gender,omitempty" xml:"Gender,omitempty"`
	Height            *float64 `json:"Height,omitempty" xml:"Height,omitempty"`
	Intro             *string  `json:"Intro,omitempty" xml:"Intro,omitempty"`
	JwtToken          *string  `json:"JwtToken,omitempty" xml:"JwtToken,omitempty"`
	MapUuid           *string  `json:"MapUuid,omitempty" xml:"MapUuid,omitempty"`
	Method            *string  `json:"Method,omitempty" xml:"Method,omitempty"`
	Mode              *string  `json:"Mode,omitempty" xml:"Mode,omitempty"`
	Title             *string  `json:"Title,omitempty" xml:"Title,omitempty"`
	TryOnParamsShrink *string  `json:"TryOnParams,omitempty" xml:"TryOnParams,omitempty"`
	Type              *string  `json:"Type,omitempty" xml:"Type,omitempty"`
	Weight            *float64 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s CreateProjectShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateProjectShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateProjectShrinkRequest) GetAutoBuild() *bool {
	return s.AutoBuild
}

func (s *CreateProjectShrinkRequest) GetDependencies() *string {
	return s.Dependencies
}

func (s *CreateProjectShrinkRequest) GetExtInfo() *string {
	return s.ExtInfo
}

func (s *CreateProjectShrinkRequest) GetGender() *string {
	return s.Gender
}

func (s *CreateProjectShrinkRequest) GetHeight() *float64 {
	return s.Height
}

func (s *CreateProjectShrinkRequest) GetIntro() *string {
	return s.Intro
}

func (s *CreateProjectShrinkRequest) GetJwtToken() *string {
	return s.JwtToken
}

func (s *CreateProjectShrinkRequest) GetMapUuid() *string {
	return s.MapUuid
}

func (s *CreateProjectShrinkRequest) GetMethod() *string {
	return s.Method
}

func (s *CreateProjectShrinkRequest) GetMode() *string {
	return s.Mode
}

func (s *CreateProjectShrinkRequest) GetTitle() *string {
	return s.Title
}

func (s *CreateProjectShrinkRequest) GetTryOnParamsShrink() *string {
	return s.TryOnParamsShrink
}

func (s *CreateProjectShrinkRequest) GetType() *string {
	return s.Type
}

func (s *CreateProjectShrinkRequest) GetWeight() *float64 {
	return s.Weight
}

func (s *CreateProjectShrinkRequest) SetAutoBuild(v bool) *CreateProjectShrinkRequest {
	s.AutoBuild = &v
	return s
}

func (s *CreateProjectShrinkRequest) SetDependencies(v string) *CreateProjectShrinkRequest {
	s.Dependencies = &v
	return s
}

func (s *CreateProjectShrinkRequest) SetExtInfo(v string) *CreateProjectShrinkRequest {
	s.ExtInfo = &v
	return s
}

func (s *CreateProjectShrinkRequest) SetGender(v string) *CreateProjectShrinkRequest {
	s.Gender = &v
	return s
}

func (s *CreateProjectShrinkRequest) SetHeight(v float64) *CreateProjectShrinkRequest {
	s.Height = &v
	return s
}

func (s *CreateProjectShrinkRequest) SetIntro(v string) *CreateProjectShrinkRequest {
	s.Intro = &v
	return s
}

func (s *CreateProjectShrinkRequest) SetJwtToken(v string) *CreateProjectShrinkRequest {
	s.JwtToken = &v
	return s
}

func (s *CreateProjectShrinkRequest) SetMapUuid(v string) *CreateProjectShrinkRequest {
	s.MapUuid = &v
	return s
}

func (s *CreateProjectShrinkRequest) SetMethod(v string) *CreateProjectShrinkRequest {
	s.Method = &v
	return s
}

func (s *CreateProjectShrinkRequest) SetMode(v string) *CreateProjectShrinkRequest {
	s.Mode = &v
	return s
}

func (s *CreateProjectShrinkRequest) SetTitle(v string) *CreateProjectShrinkRequest {
	s.Title = &v
	return s
}

func (s *CreateProjectShrinkRequest) SetTryOnParamsShrink(v string) *CreateProjectShrinkRequest {
	s.TryOnParamsShrink = &v
	return s
}

func (s *CreateProjectShrinkRequest) SetType(v string) *CreateProjectShrinkRequest {
	s.Type = &v
	return s
}

func (s *CreateProjectShrinkRequest) SetWeight(v float64) *CreateProjectShrinkRequest {
	s.Weight = &v
	return s
}

func (s *CreateProjectShrinkRequest) Validate() error {
	return dara.Validate(s)
}
