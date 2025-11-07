// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGitRepositoryContentsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBranch(v string) *ListGitRepositoryContentsRequest
	GetBranch() *string
	SetClientToken(v string) *ListGitRepositoryContentsRequest
	GetClientToken() *string
	SetContentType(v string) *ListGitRepositoryContentsRequest
	GetContentType() *string
	SetOrgId(v string) *ListGitRepositoryContentsRequest
	GetOrgId() *string
	SetOwner(v string) *ListGitRepositoryContentsRequest
	GetOwner() *string
	SetPath(v string) *ListGitRepositoryContentsRequest
	GetPath() *string
	SetPlatform(v string) *ListGitRepositoryContentsRequest
	GetPlatform() *string
	SetRegionId(v string) *ListGitRepositoryContentsRequest
	GetRegionId() *string
	SetRepoFullName(v string) *ListGitRepositoryContentsRequest
	GetRepoFullName() *string
	SetRepoId(v int64) *ListGitRepositoryContentsRequest
	GetRepoId() *int64
}

type ListGitRepositoryContentsRequest struct {
	// example:
	//
	// dev
	Branch *string `json:"Branch,omitempty" xml:"Branch,omitempty"`
	// example:
	//
	// -
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// dir
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	// example:
	//
	// 6fekfmnfll135123kdf33
	OrgId *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// dhuai
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// example:
	//
	// dir1
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// gitee
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// LYH-RAIN/vue-color-avatar
	RepoFullName *string `json:"RepoFullName,omitempty" xml:"RepoFullName,omitempty"`
	// example:
	//
	// 2642213
	RepoId *int64 `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
}

func (s ListGitRepositoryContentsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListGitRepositoryContentsRequest) GoString() string {
	return s.String()
}

func (s *ListGitRepositoryContentsRequest) GetBranch() *string {
	return s.Branch
}

func (s *ListGitRepositoryContentsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ListGitRepositoryContentsRequest) GetContentType() *string {
	return s.ContentType
}

func (s *ListGitRepositoryContentsRequest) GetOrgId() *string {
	return s.OrgId
}

func (s *ListGitRepositoryContentsRequest) GetOwner() *string {
	return s.Owner
}

func (s *ListGitRepositoryContentsRequest) GetPath() *string {
	return s.Path
}

func (s *ListGitRepositoryContentsRequest) GetPlatform() *string {
	return s.Platform
}

func (s *ListGitRepositoryContentsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListGitRepositoryContentsRequest) GetRepoFullName() *string {
	return s.RepoFullName
}

func (s *ListGitRepositoryContentsRequest) GetRepoId() *int64 {
	return s.RepoId
}

func (s *ListGitRepositoryContentsRequest) SetBranch(v string) *ListGitRepositoryContentsRequest {
	s.Branch = &v
	return s
}

func (s *ListGitRepositoryContentsRequest) SetClientToken(v string) *ListGitRepositoryContentsRequest {
	s.ClientToken = &v
	return s
}

func (s *ListGitRepositoryContentsRequest) SetContentType(v string) *ListGitRepositoryContentsRequest {
	s.ContentType = &v
	return s
}

func (s *ListGitRepositoryContentsRequest) SetOrgId(v string) *ListGitRepositoryContentsRequest {
	s.OrgId = &v
	return s
}

func (s *ListGitRepositoryContentsRequest) SetOwner(v string) *ListGitRepositoryContentsRequest {
	s.Owner = &v
	return s
}

func (s *ListGitRepositoryContentsRequest) SetPath(v string) *ListGitRepositoryContentsRequest {
	s.Path = &v
	return s
}

func (s *ListGitRepositoryContentsRequest) SetPlatform(v string) *ListGitRepositoryContentsRequest {
	s.Platform = &v
	return s
}

func (s *ListGitRepositoryContentsRequest) SetRegionId(v string) *ListGitRepositoryContentsRequest {
	s.RegionId = &v
	return s
}

func (s *ListGitRepositoryContentsRequest) SetRepoFullName(v string) *ListGitRepositoryContentsRequest {
	s.RepoFullName = &v
	return s
}

func (s *ListGitRepositoryContentsRequest) SetRepoId(v int64) *ListGitRepositoryContentsRequest {
	s.RepoId = &v
	return s
}

func (s *ListGitRepositoryContentsRequest) Validate() error {
	return dara.Validate(s)
}
