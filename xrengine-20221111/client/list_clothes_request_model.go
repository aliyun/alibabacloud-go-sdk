// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClothesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategories(v []*int64) *ListClothesRequest
	GetCategories() []*int64
	SetClothSize(v string) *ListClothesRequest
	GetClothSize() *string
	SetCurrent(v int32) *ListClothesRequest
	GetCurrent() *int32
	SetJwtToken(v string) *ListClothesRequest
	GetJwtToken() *string
	SetName(v string) *ListClothesRequest
	GetName() *string
	SetSize(v int32) *ListClothesRequest
	GetSize() *int32
	SetType(v string) *ListClothesRequest
	GetType() *string
}

type ListClothesRequest struct {
	Categories []*int64 `json:"Categories,omitempty" xml:"Categories,omitempty" type:"Repeated"`
	ClothSize  *string  `json:"ClothSize,omitempty" xml:"ClothSize,omitempty"`
	Current    *int32   `json:"Current,omitempty" xml:"Current,omitempty"`
	JwtToken   *string  `json:"JwtToken,omitempty" xml:"JwtToken,omitempty"`
	Name       *string  `json:"Name,omitempty" xml:"Name,omitempty"`
	Size       *int32   `json:"Size,omitempty" xml:"Size,omitempty"`
	Type       *string  `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListClothesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListClothesRequest) GoString() string {
	return s.String()
}

func (s *ListClothesRequest) GetCategories() []*int64 {
	return s.Categories
}

func (s *ListClothesRequest) GetClothSize() *string {
	return s.ClothSize
}

func (s *ListClothesRequest) GetCurrent() *int32 {
	return s.Current
}

func (s *ListClothesRequest) GetJwtToken() *string {
	return s.JwtToken
}

func (s *ListClothesRequest) GetName() *string {
	return s.Name
}

func (s *ListClothesRequest) GetSize() *int32 {
	return s.Size
}

func (s *ListClothesRequest) GetType() *string {
	return s.Type
}

func (s *ListClothesRequest) SetCategories(v []*int64) *ListClothesRequest {
	s.Categories = v
	return s
}

func (s *ListClothesRequest) SetClothSize(v string) *ListClothesRequest {
	s.ClothSize = &v
	return s
}

func (s *ListClothesRequest) SetCurrent(v int32) *ListClothesRequest {
	s.Current = &v
	return s
}

func (s *ListClothesRequest) SetJwtToken(v string) *ListClothesRequest {
	s.JwtToken = &v
	return s
}

func (s *ListClothesRequest) SetName(v string) *ListClothesRequest {
	s.Name = &v
	return s
}

func (s *ListClothesRequest) SetSize(v int32) *ListClothesRequest {
	s.Size = &v
	return s
}

func (s *ListClothesRequest) SetType(v string) *ListClothesRequest {
	s.Type = &v
	return s
}

func (s *ListClothesRequest) Validate() error {
	return dara.Validate(s)
}
