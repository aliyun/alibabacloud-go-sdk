// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCommitStatusesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessToken(v string) *ListCommitStatusesRequest
	GetAccessToken() *string
	SetOrganizationId(v string) *ListCommitStatusesRequest
	GetOrganizationId() *string
	SetPage(v int64) *ListCommitStatusesRequest
	GetPage() *int64
	SetPageSize(v int64) *ListCommitStatusesRequest
	GetPageSize() *int64
	SetRepositoryIdentity(v string) *ListCommitStatusesRequest
	GetRepositoryIdentity() *string
	SetSha(v string) *ListCommitStatusesRequest
	GetSha() *string
}

type ListCommitStatusesRequest struct {
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
	RepositoryIdentity *string `json:"repositoryIdentity,omitempty" xml:"repositoryIdentity,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 61cc69557962d29f737a91730b3e86f497f083a3
	Sha *string `json:"sha,omitempty" xml:"sha,omitempty"`
}

func (s ListCommitStatusesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCommitStatusesRequest) GoString() string {
	return s.String()
}

func (s *ListCommitStatusesRequest) GetAccessToken() *string {
	return s.AccessToken
}

func (s *ListCommitStatusesRequest) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *ListCommitStatusesRequest) GetPage() *int64 {
	return s.Page
}

func (s *ListCommitStatusesRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListCommitStatusesRequest) GetRepositoryIdentity() *string {
	return s.RepositoryIdentity
}

func (s *ListCommitStatusesRequest) GetSha() *string {
	return s.Sha
}

func (s *ListCommitStatusesRequest) SetAccessToken(v string) *ListCommitStatusesRequest {
	s.AccessToken = &v
	return s
}

func (s *ListCommitStatusesRequest) SetOrganizationId(v string) *ListCommitStatusesRequest {
	s.OrganizationId = &v
	return s
}

func (s *ListCommitStatusesRequest) SetPage(v int64) *ListCommitStatusesRequest {
	s.Page = &v
	return s
}

func (s *ListCommitStatusesRequest) SetPageSize(v int64) *ListCommitStatusesRequest {
	s.PageSize = &v
	return s
}

func (s *ListCommitStatusesRequest) SetRepositoryIdentity(v string) *ListCommitStatusesRequest {
	s.RepositoryIdentity = &v
	return s
}

func (s *ListCommitStatusesRequest) SetSha(v string) *ListCommitStatusesRequest {
	s.Sha = &v
	return s
}

func (s *ListCommitStatusesRequest) Validate() error {
	return dara.Validate(s)
}
