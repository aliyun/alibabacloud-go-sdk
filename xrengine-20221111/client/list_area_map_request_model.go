// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAreaMapRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrent(v int32) *ListAreaMapRequest
	GetCurrent() *int32
	SetJwtToken(v string) *ListAreaMapRequest
	GetJwtToken() *string
	SetSize(v int32) *ListAreaMapRequest
	GetSize() *int32
	SetSort(v string) *ListAreaMapRequest
	GetSort() *string
	SetSortField(v string) *ListAreaMapRequest
	GetSortField() *string
}

type ListAreaMapRequest struct {
	Current   *int32  `json:"Current,omitempty" xml:"Current,omitempty"`
	JwtToken  *string `json:"JwtToken,omitempty" xml:"JwtToken,omitempty"`
	Size      *int32  `json:"Size,omitempty" xml:"Size,omitempty"`
	Sort      *string `json:"Sort,omitempty" xml:"Sort,omitempty"`
	SortField *string `json:"SortField,omitempty" xml:"SortField,omitempty"`
}

func (s ListAreaMapRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAreaMapRequest) GoString() string {
	return s.String()
}

func (s *ListAreaMapRequest) GetCurrent() *int32 {
	return s.Current
}

func (s *ListAreaMapRequest) GetJwtToken() *string {
	return s.JwtToken
}

func (s *ListAreaMapRequest) GetSize() *int32 {
	return s.Size
}

func (s *ListAreaMapRequest) GetSort() *string {
	return s.Sort
}

func (s *ListAreaMapRequest) GetSortField() *string {
	return s.SortField
}

func (s *ListAreaMapRequest) SetCurrent(v int32) *ListAreaMapRequest {
	s.Current = &v
	return s
}

func (s *ListAreaMapRequest) SetJwtToken(v string) *ListAreaMapRequest {
	s.JwtToken = &v
	return s
}

func (s *ListAreaMapRequest) SetSize(v int32) *ListAreaMapRequest {
	s.Size = &v
	return s
}

func (s *ListAreaMapRequest) SetSort(v string) *ListAreaMapRequest {
	s.Sort = &v
	return s
}

func (s *ListAreaMapRequest) SetSortField(v string) *ListAreaMapRequest {
	s.SortField = &v
	return s
}

func (s *ListAreaMapRequest) Validate() error {
	return dara.Validate(s)
}
