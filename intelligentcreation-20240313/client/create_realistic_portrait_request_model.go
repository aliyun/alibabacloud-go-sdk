// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRealisticPortraitRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAges(v []*int32) *CreateRealisticPortraitRequest
	GetAges() []*int32
	SetCloth(v int32) *CreateRealisticPortraitRequest
	GetCloth() *int32
	SetColor(v int32) *CreateRealisticPortraitRequest
	GetColor() *int32
	SetCustom(v string) *CreateRealisticPortraitRequest
	GetCustom() *string
	SetFace(v []*int32) *CreateRealisticPortraitRequest
	GetFace() []*int32
	SetFigure(v int32) *CreateRealisticPortraitRequest
	GetFigure() *int32
	SetGender(v int32) *CreateRealisticPortraitRequest
	GetGender() *int32
	SetHairColor(v int32) *CreateRealisticPortraitRequest
	GetHairColor() *int32
	SetHairstyle(v int32) *CreateRealisticPortraitRequest
	GetHairstyle() *int32
	SetHeight(v int32) *CreateRealisticPortraitRequest
	GetHeight() *int32
	SetImageUrl(v string) *CreateRealisticPortraitRequest
	GetImageUrl() *string
	SetNumbers(v int32) *CreateRealisticPortraitRequest
	GetNumbers() *int32
	SetRatio(v string) *CreateRealisticPortraitRequest
	GetRatio() *string
	SetWidth(v int32) *CreateRealisticPortraitRequest
	GetWidth() *int32
}

type CreateRealisticPortraitRequest struct {
	Ages []*int32 `json:"ages,omitempty" xml:"ages,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	Cloth *int32 `json:"cloth,omitempty" xml:"cloth,omitempty"`
	// example:
	//
	// 1
	Color *int32 `json:"color,omitempty" xml:"color,omitempty"`
	// example:
	//
	// 11
	Custom *string  `json:"custom,omitempty" xml:"custom,omitempty"`
	Face   []*int32 `json:"face,omitempty" xml:"face,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	Figure *int32 `json:"figure,omitempty" xml:"figure,omitempty"`
	// example:
	//
	// 1
	Gender *int32 `json:"gender,omitempty" xml:"gender,omitempty"`
	// example:
	//
	// 1
	HairColor *int32 `json:"hairColor,omitempty" xml:"hairColor,omitempty"`
	// example:
	//
	// 1
	Hairstyle *int32 `json:"hairstyle,omitempty" xml:"hairstyle,omitempty"`
	// example:
	//
	// 500
	Height   *int32  `json:"height,omitempty" xml:"height,omitempty"`
	ImageUrl *string `json:"imageUrl,omitempty" xml:"imageUrl,omitempty"`
	// example:
	//
	// 4
	Numbers *int32 `json:"numbers,omitempty" xml:"numbers,omitempty"`
	// example:
	//
	// 1:1
	Ratio *string `json:"ratio,omitempty" xml:"ratio,omitempty"`
	// example:
	//
	// 500
	Width *int32 `json:"width,omitempty" xml:"width,omitempty"`
}

func (s CreateRealisticPortraitRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRealisticPortraitRequest) GoString() string {
	return s.String()
}

func (s *CreateRealisticPortraitRequest) GetAges() []*int32 {
	return s.Ages
}

func (s *CreateRealisticPortraitRequest) GetCloth() *int32 {
	return s.Cloth
}

func (s *CreateRealisticPortraitRequest) GetColor() *int32 {
	return s.Color
}

func (s *CreateRealisticPortraitRequest) GetCustom() *string {
	return s.Custom
}

func (s *CreateRealisticPortraitRequest) GetFace() []*int32 {
	return s.Face
}

func (s *CreateRealisticPortraitRequest) GetFigure() *int32 {
	return s.Figure
}

func (s *CreateRealisticPortraitRequest) GetGender() *int32 {
	return s.Gender
}

func (s *CreateRealisticPortraitRequest) GetHairColor() *int32 {
	return s.HairColor
}

func (s *CreateRealisticPortraitRequest) GetHairstyle() *int32 {
	return s.Hairstyle
}

func (s *CreateRealisticPortraitRequest) GetHeight() *int32 {
	return s.Height
}

func (s *CreateRealisticPortraitRequest) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *CreateRealisticPortraitRequest) GetNumbers() *int32 {
	return s.Numbers
}

func (s *CreateRealisticPortraitRequest) GetRatio() *string {
	return s.Ratio
}

func (s *CreateRealisticPortraitRequest) GetWidth() *int32 {
	return s.Width
}

func (s *CreateRealisticPortraitRequest) SetAges(v []*int32) *CreateRealisticPortraitRequest {
	s.Ages = v
	return s
}

func (s *CreateRealisticPortraitRequest) SetCloth(v int32) *CreateRealisticPortraitRequest {
	s.Cloth = &v
	return s
}

func (s *CreateRealisticPortraitRequest) SetColor(v int32) *CreateRealisticPortraitRequest {
	s.Color = &v
	return s
}

func (s *CreateRealisticPortraitRequest) SetCustom(v string) *CreateRealisticPortraitRequest {
	s.Custom = &v
	return s
}

func (s *CreateRealisticPortraitRequest) SetFace(v []*int32) *CreateRealisticPortraitRequest {
	s.Face = v
	return s
}

func (s *CreateRealisticPortraitRequest) SetFigure(v int32) *CreateRealisticPortraitRequest {
	s.Figure = &v
	return s
}

func (s *CreateRealisticPortraitRequest) SetGender(v int32) *CreateRealisticPortraitRequest {
	s.Gender = &v
	return s
}

func (s *CreateRealisticPortraitRequest) SetHairColor(v int32) *CreateRealisticPortraitRequest {
	s.HairColor = &v
	return s
}

func (s *CreateRealisticPortraitRequest) SetHairstyle(v int32) *CreateRealisticPortraitRequest {
	s.Hairstyle = &v
	return s
}

func (s *CreateRealisticPortraitRequest) SetHeight(v int32) *CreateRealisticPortraitRequest {
	s.Height = &v
	return s
}

func (s *CreateRealisticPortraitRequest) SetImageUrl(v string) *CreateRealisticPortraitRequest {
	s.ImageUrl = &v
	return s
}

func (s *CreateRealisticPortraitRequest) SetNumbers(v int32) *CreateRealisticPortraitRequest {
	s.Numbers = &v
	return s
}

func (s *CreateRealisticPortraitRequest) SetRatio(v string) *CreateRealisticPortraitRequest {
	s.Ratio = &v
	return s
}

func (s *CreateRealisticPortraitRequest) SetWidth(v int32) *CreateRealisticPortraitRequest {
	s.Width = &v
	return s
}

func (s *CreateRealisticPortraitRequest) Validate() error {
	return dara.Validate(s)
}
