// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizUsage(v string) *ListProjectRequest
	GetBizUsage() *string
	SetCurrent(v int32) *ListProjectRequest
	GetCurrent() *int32
	SetExcludedBizUsage(v string) *ListProjectRequest
	GetExcludedBizUsage() *string
	SetJwtToken(v string) *ListProjectRequest
	GetJwtToken() *string
	SetSize(v int32) *ListProjectRequest
	GetSize() *int32
	SetSortField(v string) *ListProjectRequest
	GetSortField() *string
	SetStatus(v string) *ListProjectRequest
	GetStatus() *string
	SetTitle(v string) *ListProjectRequest
	GetTitle() *string
	SetType(v string) *ListProjectRequest
	GetType() *string
	SetWithSource(v bool) *ListProjectRequest
	GetWithSource() *bool
	SetWithUser(v bool) *ListProjectRequest
	GetWithUser() *bool
}

type ListProjectRequest struct {
	BizUsage         *string `json:"BizUsage,omitempty" xml:"BizUsage,omitempty"`
	Current          *int32  `json:"Current,omitempty" xml:"Current,omitempty"`
	ExcludedBizUsage *string `json:"ExcludedBizUsage,omitempty" xml:"ExcludedBizUsage,omitempty"`
	JwtToken         *string `json:"JwtToken,omitempty" xml:"JwtToken,omitempty"`
	Size             *int32  `json:"Size,omitempty" xml:"Size,omitempty"`
	SortField        *string `json:"SortField,omitempty" xml:"SortField,omitempty"`
	Status           *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Title            *string `json:"Title,omitempty" xml:"Title,omitempty"`
	Type             *string `json:"Type,omitempty" xml:"Type,omitempty"`
	WithSource       *bool   `json:"WithSource,omitempty" xml:"WithSource,omitempty"`
	WithUser         *bool   `json:"WithUser,omitempty" xml:"WithUser,omitempty"`
}

func (s ListProjectRequest) String() string {
	return dara.Prettify(s)
}

func (s ListProjectRequest) GoString() string {
	return s.String()
}

func (s *ListProjectRequest) GetBizUsage() *string {
	return s.BizUsage
}

func (s *ListProjectRequest) GetCurrent() *int32 {
	return s.Current
}

func (s *ListProjectRequest) GetExcludedBizUsage() *string {
	return s.ExcludedBizUsage
}

func (s *ListProjectRequest) GetJwtToken() *string {
	return s.JwtToken
}

func (s *ListProjectRequest) GetSize() *int32 {
	return s.Size
}

func (s *ListProjectRequest) GetSortField() *string {
	return s.SortField
}

func (s *ListProjectRequest) GetStatus() *string {
	return s.Status
}

func (s *ListProjectRequest) GetTitle() *string {
	return s.Title
}

func (s *ListProjectRequest) GetType() *string {
	return s.Type
}

func (s *ListProjectRequest) GetWithSource() *bool {
	return s.WithSource
}

func (s *ListProjectRequest) GetWithUser() *bool {
	return s.WithUser
}

func (s *ListProjectRequest) SetBizUsage(v string) *ListProjectRequest {
	s.BizUsage = &v
	return s
}

func (s *ListProjectRequest) SetCurrent(v int32) *ListProjectRequest {
	s.Current = &v
	return s
}

func (s *ListProjectRequest) SetExcludedBizUsage(v string) *ListProjectRequest {
	s.ExcludedBizUsage = &v
	return s
}

func (s *ListProjectRequest) SetJwtToken(v string) *ListProjectRequest {
	s.JwtToken = &v
	return s
}

func (s *ListProjectRequest) SetSize(v int32) *ListProjectRequest {
	s.Size = &v
	return s
}

func (s *ListProjectRequest) SetSortField(v string) *ListProjectRequest {
	s.SortField = &v
	return s
}

func (s *ListProjectRequest) SetStatus(v string) *ListProjectRequest {
	s.Status = &v
	return s
}

func (s *ListProjectRequest) SetTitle(v string) *ListProjectRequest {
	s.Title = &v
	return s
}

func (s *ListProjectRequest) SetType(v string) *ListProjectRequest {
	s.Type = &v
	return s
}

func (s *ListProjectRequest) SetWithSource(v bool) *ListProjectRequest {
	s.WithSource = &v
	return s
}

func (s *ListProjectRequest) SetWithUser(v bool) *ListProjectRequest {
	s.WithUser = &v
	return s
}

func (s *ListProjectRequest) Validate() error {
	return dara.Validate(s)
}
