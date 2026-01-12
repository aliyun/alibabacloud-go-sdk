// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLocationPaiImageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrent(v int32) *ListLocationPaiImageRequest
	GetCurrent() *int32
	SetJwtToken(v string) *ListLocationPaiImageRequest
	GetJwtToken() *string
	SetSize(v int32) *ListLocationPaiImageRequest
	GetSize() *int32
	SetSort(v string) *ListLocationPaiImageRequest
	GetSort() *string
	SetSortField(v string) *ListLocationPaiImageRequest
	GetSortField() *string
}

type ListLocationPaiImageRequest struct {
	Current   *int32  `json:"Current,omitempty" xml:"Current,omitempty"`
	JwtToken  *string `json:"JwtToken,omitempty" xml:"JwtToken,omitempty"`
	Size      *int32  `json:"Size,omitempty" xml:"Size,omitempty"`
	Sort      *string `json:"Sort,omitempty" xml:"Sort,omitempty"`
	SortField *string `json:"SortField,omitempty" xml:"SortField,omitempty"`
}

func (s ListLocationPaiImageRequest) String() string {
	return dara.Prettify(s)
}

func (s ListLocationPaiImageRequest) GoString() string {
	return s.String()
}

func (s *ListLocationPaiImageRequest) GetCurrent() *int32 {
	return s.Current
}

func (s *ListLocationPaiImageRequest) GetJwtToken() *string {
	return s.JwtToken
}

func (s *ListLocationPaiImageRequest) GetSize() *int32 {
	return s.Size
}

func (s *ListLocationPaiImageRequest) GetSort() *string {
	return s.Sort
}

func (s *ListLocationPaiImageRequest) GetSortField() *string {
	return s.SortField
}

func (s *ListLocationPaiImageRequest) SetCurrent(v int32) *ListLocationPaiImageRequest {
	s.Current = &v
	return s
}

func (s *ListLocationPaiImageRequest) SetJwtToken(v string) *ListLocationPaiImageRequest {
	s.JwtToken = &v
	return s
}

func (s *ListLocationPaiImageRequest) SetSize(v int32) *ListLocationPaiImageRequest {
	s.Size = &v
	return s
}

func (s *ListLocationPaiImageRequest) SetSort(v string) *ListLocationPaiImageRequest {
	s.Sort = &v
	return s
}

func (s *ListLocationPaiImageRequest) SetSortField(v string) *ListLocationPaiImageRequest {
	s.SortField = &v
	return s
}

func (s *ListLocationPaiImageRequest) Validate() error {
	return dara.Validate(s)
}
