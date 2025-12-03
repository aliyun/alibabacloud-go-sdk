// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCheckRunsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessToken(v string) *ListCheckRunsRequest
	GetAccessToken() *string
	SetOrganizationId(v string) *ListCheckRunsRequest
	GetOrganizationId() *string
	SetPage(v int64) *ListCheckRunsRequest
	GetPage() *int64
	SetPageSize(v int64) *ListCheckRunsRequest
	GetPageSize() *int64
	SetRef(v string) *ListCheckRunsRequest
	GetRef() *string
	SetRepositoryIdentity(v string) *ListCheckRunsRequest
	GetRepositoryIdentity() *string
}

type ListCheckRunsRequest struct {
	// example:
	//
	// f0b1e61db5961df5975a93f9129d2513
	AccessToken *string `json:"accessToken,omitempty" xml:"accessToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 5ebbc0228123212b59xxxxx
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
	// example:
	//
	// 1
	Page *int64 `json:"page,omitempty" xml:"page,omitempty"`
	// example:
	//
	// 20
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 40f4ccfe019cdd4a62d4acb0c57130106fc7e1be
	Ref *string `json:"ref,omitempty" xml:"ref,omitempty"`
	// This parameter is required.
	RepositoryIdentity *string `json:"repositoryIdentity,omitempty" xml:"repositoryIdentity,omitempty"`
}

func (s ListCheckRunsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCheckRunsRequest) GoString() string {
	return s.String()
}

func (s *ListCheckRunsRequest) GetAccessToken() *string {
	return s.AccessToken
}

func (s *ListCheckRunsRequest) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *ListCheckRunsRequest) GetPage() *int64 {
	return s.Page
}

func (s *ListCheckRunsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListCheckRunsRequest) GetRef() *string {
	return s.Ref
}

func (s *ListCheckRunsRequest) GetRepositoryIdentity() *string {
	return s.RepositoryIdentity
}

func (s *ListCheckRunsRequest) SetAccessToken(v string) *ListCheckRunsRequest {
	s.AccessToken = &v
	return s
}

func (s *ListCheckRunsRequest) SetOrganizationId(v string) *ListCheckRunsRequest {
	s.OrganizationId = &v
	return s
}

func (s *ListCheckRunsRequest) SetPage(v int64) *ListCheckRunsRequest {
	s.Page = &v
	return s
}

func (s *ListCheckRunsRequest) SetPageSize(v int64) *ListCheckRunsRequest {
	s.PageSize = &v
	return s
}

func (s *ListCheckRunsRequest) SetRef(v string) *ListCheckRunsRequest {
	s.Ref = &v
	return s
}

func (s *ListCheckRunsRequest) SetRepositoryIdentity(v string) *ListCheckRunsRequest {
	s.RepositoryIdentity = &v
	return s
}

func (s *ListCheckRunsRequest) Validate() error {
	return dara.Validate(s)
}
