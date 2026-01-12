// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFindSvcMapBindRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrent(v int32) *FindSvcMapBindRequest
	GetCurrent() *int32
	SetJwtToken(v string) *FindSvcMapBindRequest
	GetJwtToken() *string
	SetSize(v int32) *FindSvcMapBindRequest
	GetSize() *int32
	SetSort(v string) *FindSvcMapBindRequest
	GetSort() *string
	SetSortField(v string) *FindSvcMapBindRequest
	GetSortField() *string
	SetSvcId(v int64) *FindSvcMapBindRequest
	GetSvcId() *int64
}

type FindSvcMapBindRequest struct {
	Current   *int32  `json:"Current,omitempty" xml:"Current,omitempty"`
	JwtToken  *string `json:"JwtToken,omitempty" xml:"JwtToken,omitempty"`
	Size      *int32  `json:"Size,omitempty" xml:"Size,omitempty"`
	Sort      *string `json:"Sort,omitempty" xml:"Sort,omitempty"`
	SortField *string `json:"SortField,omitempty" xml:"SortField,omitempty"`
	SvcId     *int64  `json:"SvcId,omitempty" xml:"SvcId,omitempty"`
}

func (s FindSvcMapBindRequest) String() string {
	return dara.Prettify(s)
}

func (s FindSvcMapBindRequest) GoString() string {
	return s.String()
}

func (s *FindSvcMapBindRequest) GetCurrent() *int32 {
	return s.Current
}

func (s *FindSvcMapBindRequest) GetJwtToken() *string {
	return s.JwtToken
}

func (s *FindSvcMapBindRequest) GetSize() *int32 {
	return s.Size
}

func (s *FindSvcMapBindRequest) GetSort() *string {
	return s.Sort
}

func (s *FindSvcMapBindRequest) GetSortField() *string {
	return s.SortField
}

func (s *FindSvcMapBindRequest) GetSvcId() *int64 {
	return s.SvcId
}

func (s *FindSvcMapBindRequest) SetCurrent(v int32) *FindSvcMapBindRequest {
	s.Current = &v
	return s
}

func (s *FindSvcMapBindRequest) SetJwtToken(v string) *FindSvcMapBindRequest {
	s.JwtToken = &v
	return s
}

func (s *FindSvcMapBindRequest) SetSize(v int32) *FindSvcMapBindRequest {
	s.Size = &v
	return s
}

func (s *FindSvcMapBindRequest) SetSort(v string) *FindSvcMapBindRequest {
	s.Sort = &v
	return s
}

func (s *FindSvcMapBindRequest) SetSortField(v string) *FindSvcMapBindRequest {
	s.SortField = &v
	return s
}

func (s *FindSvcMapBindRequest) SetSvcId(v int64) *FindSvcMapBindRequest {
	s.SvcId = &v
	return s
}

func (s *FindSvcMapBindRequest) Validate() error {
	return dara.Validate(s)
}
