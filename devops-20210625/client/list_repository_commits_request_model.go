// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRepositoryCommitsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessToken(v string) *ListRepositoryCommitsRequest
	GetAccessToken() *string
	SetEnd(v string) *ListRepositoryCommitsRequest
	GetEnd() *string
	SetOrganizationId(v string) *ListRepositoryCommitsRequest
	GetOrganizationId() *string
	SetPage(v int64) *ListRepositoryCommitsRequest
	GetPage() *int64
	SetPageSize(v int64) *ListRepositoryCommitsRequest
	GetPageSize() *int64
	SetPath(v string) *ListRepositoryCommitsRequest
	GetPath() *string
	SetRefName(v string) *ListRepositoryCommitsRequest
	GetRefName() *string
	SetSearch(v string) *ListRepositoryCommitsRequest
	GetSearch() *string
	SetShowCommentsCount(v bool) *ListRepositoryCommitsRequest
	GetShowCommentsCount() *bool
	SetShowSignature(v bool) *ListRepositoryCommitsRequest
	GetShowSignature() *bool
	SetStart(v string) *ListRepositoryCommitsRequest
	GetStart() *string
}

type ListRepositoryCommitsRequest struct {
	// example:
	//
	// f0b1e61db5961df5975a93f9129d2513
	AccessToken *string `json:"accessToken,omitempty" xml:"accessToken,omitempty"`
	// example:
	//
	// 2022-08-18 08:00:00
	End *string `json:"end,omitempty" xml:"end,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 60de7a6852743a5162b5f957
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
	// example:
	//
	// 1
	Page *int64 `json:"page,omitempty" xml:"page,omitempty"`
	// example:
	//
	// 20
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// src/cpp/main.cpp
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// This parameter is required.
	RefName *string `json:"refName,omitempty" xml:"refName,omitempty"`
	// example:
	//
	// search
	Search *string `json:"search,omitempty" xml:"search,omitempty"`
	// example:
	//
	// false
	ShowCommentsCount *bool `json:"showCommentsCount,omitempty" xml:"showCommentsCount,omitempty"`
	// example:
	//
	// false
	ShowSignature *bool `json:"showSignature,omitempty" xml:"showSignature,omitempty"`
	// example:
	//
	// 2022-03-18 08:00:00
	Start *string `json:"start,omitempty" xml:"start,omitempty"`
}

func (s ListRepositoryCommitsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRepositoryCommitsRequest) GoString() string {
	return s.String()
}

func (s *ListRepositoryCommitsRequest) GetAccessToken() *string {
	return s.AccessToken
}

func (s *ListRepositoryCommitsRequest) GetEnd() *string {
	return s.End
}

func (s *ListRepositoryCommitsRequest) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *ListRepositoryCommitsRequest) GetPage() *int64 {
	return s.Page
}

func (s *ListRepositoryCommitsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListRepositoryCommitsRequest) GetPath() *string {
	return s.Path
}

func (s *ListRepositoryCommitsRequest) GetRefName() *string {
	return s.RefName
}

func (s *ListRepositoryCommitsRequest) GetSearch() *string {
	return s.Search
}

func (s *ListRepositoryCommitsRequest) GetShowCommentsCount() *bool {
	return s.ShowCommentsCount
}

func (s *ListRepositoryCommitsRequest) GetShowSignature() *bool {
	return s.ShowSignature
}

func (s *ListRepositoryCommitsRequest) GetStart() *string {
	return s.Start
}

func (s *ListRepositoryCommitsRequest) SetAccessToken(v string) *ListRepositoryCommitsRequest {
	s.AccessToken = &v
	return s
}

func (s *ListRepositoryCommitsRequest) SetEnd(v string) *ListRepositoryCommitsRequest {
	s.End = &v
	return s
}

func (s *ListRepositoryCommitsRequest) SetOrganizationId(v string) *ListRepositoryCommitsRequest {
	s.OrganizationId = &v
	return s
}

func (s *ListRepositoryCommitsRequest) SetPage(v int64) *ListRepositoryCommitsRequest {
	s.Page = &v
	return s
}

func (s *ListRepositoryCommitsRequest) SetPageSize(v int64) *ListRepositoryCommitsRequest {
	s.PageSize = &v
	return s
}

func (s *ListRepositoryCommitsRequest) SetPath(v string) *ListRepositoryCommitsRequest {
	s.Path = &v
	return s
}

func (s *ListRepositoryCommitsRequest) SetRefName(v string) *ListRepositoryCommitsRequest {
	s.RefName = &v
	return s
}

func (s *ListRepositoryCommitsRequest) SetSearch(v string) *ListRepositoryCommitsRequest {
	s.Search = &v
	return s
}

func (s *ListRepositoryCommitsRequest) SetShowCommentsCount(v bool) *ListRepositoryCommitsRequest {
	s.ShowCommentsCount = &v
	return s
}

func (s *ListRepositoryCommitsRequest) SetShowSignature(v bool) *ListRepositoryCommitsRequest {
	s.ShowSignature = &v
	return s
}

func (s *ListRepositoryCommitsRequest) SetStart(v string) *ListRepositoryCommitsRequest {
	s.Start = &v
	return s
}

func (s *ListRepositoryCommitsRequest) Validate() error {
	return dara.Validate(s)
}
