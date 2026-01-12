// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClothesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategoriesShrink(v string) *ListClothesShrinkRequest
	GetCategoriesShrink() *string
	SetClothSize(v string) *ListClothesShrinkRequest
	GetClothSize() *string
	SetCurrent(v int32) *ListClothesShrinkRequest
	GetCurrent() *int32
	SetJwtToken(v string) *ListClothesShrinkRequest
	GetJwtToken() *string
	SetName(v string) *ListClothesShrinkRequest
	GetName() *string
	SetSize(v int32) *ListClothesShrinkRequest
	GetSize() *int32
	SetType(v string) *ListClothesShrinkRequest
	GetType() *string
}

type ListClothesShrinkRequest struct {
	CategoriesShrink *string `json:"Categories,omitempty" xml:"Categories,omitempty"`
	ClothSize        *string `json:"ClothSize,omitempty" xml:"ClothSize,omitempty"`
	Current          *int32  `json:"Current,omitempty" xml:"Current,omitempty"`
	JwtToken         *string `json:"JwtToken,omitempty" xml:"JwtToken,omitempty"`
	Name             *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Size             *int32  `json:"Size,omitempty" xml:"Size,omitempty"`
	Type             *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListClothesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListClothesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListClothesShrinkRequest) GetCategoriesShrink() *string {
	return s.CategoriesShrink
}

func (s *ListClothesShrinkRequest) GetClothSize() *string {
	return s.ClothSize
}

func (s *ListClothesShrinkRequest) GetCurrent() *int32 {
	return s.Current
}

func (s *ListClothesShrinkRequest) GetJwtToken() *string {
	return s.JwtToken
}

func (s *ListClothesShrinkRequest) GetName() *string {
	return s.Name
}

func (s *ListClothesShrinkRequest) GetSize() *int32 {
	return s.Size
}

func (s *ListClothesShrinkRequest) GetType() *string {
	return s.Type
}

func (s *ListClothesShrinkRequest) SetCategoriesShrink(v string) *ListClothesShrinkRequest {
	s.CategoriesShrink = &v
	return s
}

func (s *ListClothesShrinkRequest) SetClothSize(v string) *ListClothesShrinkRequest {
	s.ClothSize = &v
	return s
}

func (s *ListClothesShrinkRequest) SetCurrent(v int32) *ListClothesShrinkRequest {
	s.Current = &v
	return s
}

func (s *ListClothesShrinkRequest) SetJwtToken(v string) *ListClothesShrinkRequest {
	s.JwtToken = &v
	return s
}

func (s *ListClothesShrinkRequest) SetName(v string) *ListClothesShrinkRequest {
	s.Name = &v
	return s
}

func (s *ListClothesShrinkRequest) SetSize(v int32) *ListClothesShrinkRequest {
	s.Size = &v
	return s
}

func (s *ListClothesShrinkRequest) SetType(v string) *ListClothesShrinkRequest {
	s.Type = &v
	return s
}

func (s *ListClothesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
