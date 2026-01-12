// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLocationServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrent(v int32) *ListLocationServiceRequest
	GetCurrent() *int32
	SetJwtToken(v string) *ListLocationServiceRequest
	GetJwtToken() *string
	SetSize(v int32) *ListLocationServiceRequest
	GetSize() *int32
	SetSort(v string) *ListLocationServiceRequest
	GetSort() *string
	SetSortField(v string) *ListLocationServiceRequest
	GetSortField() *string
}

type ListLocationServiceRequest struct {
	Current   *int32  `json:"Current,omitempty" xml:"Current,omitempty"`
	JwtToken  *string `json:"JwtToken,omitempty" xml:"JwtToken,omitempty"`
	Size      *int32  `json:"Size,omitempty" xml:"Size,omitempty"`
	Sort      *string `json:"Sort,omitempty" xml:"Sort,omitempty"`
	SortField *string `json:"SortField,omitempty" xml:"SortField,omitempty"`
}

func (s ListLocationServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s ListLocationServiceRequest) GoString() string {
	return s.String()
}

func (s *ListLocationServiceRequest) GetCurrent() *int32 {
	return s.Current
}

func (s *ListLocationServiceRequest) GetJwtToken() *string {
	return s.JwtToken
}

func (s *ListLocationServiceRequest) GetSize() *int32 {
	return s.Size
}

func (s *ListLocationServiceRequest) GetSort() *string {
	return s.Sort
}

func (s *ListLocationServiceRequest) GetSortField() *string {
	return s.SortField
}

func (s *ListLocationServiceRequest) SetCurrent(v int32) *ListLocationServiceRequest {
	s.Current = &v
	return s
}

func (s *ListLocationServiceRequest) SetJwtToken(v string) *ListLocationServiceRequest {
	s.JwtToken = &v
	return s
}

func (s *ListLocationServiceRequest) SetSize(v int32) *ListLocationServiceRequest {
	s.Size = &v
	return s
}

func (s *ListLocationServiceRequest) SetSort(v string) *ListLocationServiceRequest {
	s.Sort = &v
	return s
}

func (s *ListLocationServiceRequest) SetSortField(v string) *ListLocationServiceRequest {
	s.SortField = &v
	return s
}

func (s *ListLocationServiceRequest) Validate() error {
	return dara.Validate(s)
}
