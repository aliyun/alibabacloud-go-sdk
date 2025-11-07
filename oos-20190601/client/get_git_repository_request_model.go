// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGitRepositoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrgId(v string) *GetGitRepositoryRequest
	GetOrgId() *string
	SetOwner(v string) *GetGitRepositoryRequest
	GetOwner() *string
	SetPlatform(v string) *GetGitRepositoryRequest
	GetPlatform() *string
	SetRegionId(v string) *GetGitRepositoryRequest
	GetRegionId() *string
	SetRepoFullName(v string) *GetGitRepositoryRequest
	GetRepoFullName() *string
	SetRepoId(v int64) *GetGitRepositoryRequest
	GetRepoId() *int64
}

type GetGitRepositoryRequest struct {
	// example:
	//
	// 5e9ee15af89c9700014a558a
	OrgId *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	// example:
	//
	// testowner
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// github
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// alibaba/fastjson
	RepoFullName *string `json:"RepoFullName,omitempty" xml:"RepoFullName,omitempty"`
	// example:
	//
	// 286584
	RepoId *int64 `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
}

func (s GetGitRepositoryRequest) String() string {
	return dara.Prettify(s)
}

func (s GetGitRepositoryRequest) GoString() string {
	return s.String()
}

func (s *GetGitRepositoryRequest) GetOrgId() *string {
	return s.OrgId
}

func (s *GetGitRepositoryRequest) GetOwner() *string {
	return s.Owner
}

func (s *GetGitRepositoryRequest) GetPlatform() *string {
	return s.Platform
}

func (s *GetGitRepositoryRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetGitRepositoryRequest) GetRepoFullName() *string {
	return s.RepoFullName
}

func (s *GetGitRepositoryRequest) GetRepoId() *int64 {
	return s.RepoId
}

func (s *GetGitRepositoryRequest) SetOrgId(v string) *GetGitRepositoryRequest {
	s.OrgId = &v
	return s
}

func (s *GetGitRepositoryRequest) SetOwner(v string) *GetGitRepositoryRequest {
	s.Owner = &v
	return s
}

func (s *GetGitRepositoryRequest) SetPlatform(v string) *GetGitRepositoryRequest {
	s.Platform = &v
	return s
}

func (s *GetGitRepositoryRequest) SetRegionId(v string) *GetGitRepositoryRequest {
	s.RegionId = &v
	return s
}

func (s *GetGitRepositoryRequest) SetRepoFullName(v string) *GetGitRepositoryRequest {
	s.RepoFullName = &v
	return s
}

func (s *GetGitRepositoryRequest) SetRepoId(v int64) *GetGitRepositoryRequest {
	s.RepoId = &v
	return s
}

func (s *GetGitRepositoryRequest) Validate() error {
	return dara.Validate(s)
}
