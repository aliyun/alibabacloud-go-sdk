// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAnalyzeGitRepositoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBranch(v string) *AnalyzeGitRepositoryRequest
	GetBranch() *string
	SetClientToken(v string) *AnalyzeGitRepositoryRequest
	GetClientToken() *string
	SetOrgId(v string) *AnalyzeGitRepositoryRequest
	GetOrgId() *string
	SetOwner(v string) *AnalyzeGitRepositoryRequest
	GetOwner() *string
	SetPlatform(v string) *AnalyzeGitRepositoryRequest
	GetPlatform() *string
	SetRegionId(v string) *AnalyzeGitRepositoryRequest
	GetRegionId() *string
	SetRepoFullName(v string) *AnalyzeGitRepositoryRequest
	GetRepoFullName() *string
	SetRepoId(v string) *AnalyzeGitRepositoryRequest
	GetRepoId() *string
}

type AnalyzeGitRepositoryRequest struct {
	Branch      *string `json:"Branch,omitempty" xml:"Branch,omitempty"`
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OrgId       *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	// This parameter is required.
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// This parameter is required.
	Platform     *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RepoFullName *string `json:"RepoFullName,omitempty" xml:"RepoFullName,omitempty"`
	RepoId       *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
}

func (s AnalyzeGitRepositoryRequest) String() string {
	return dara.Prettify(s)
}

func (s AnalyzeGitRepositoryRequest) GoString() string {
	return s.String()
}

func (s *AnalyzeGitRepositoryRequest) GetBranch() *string {
	return s.Branch
}

func (s *AnalyzeGitRepositoryRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *AnalyzeGitRepositoryRequest) GetOrgId() *string {
	return s.OrgId
}

func (s *AnalyzeGitRepositoryRequest) GetOwner() *string {
	return s.Owner
}

func (s *AnalyzeGitRepositoryRequest) GetPlatform() *string {
	return s.Platform
}

func (s *AnalyzeGitRepositoryRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AnalyzeGitRepositoryRequest) GetRepoFullName() *string {
	return s.RepoFullName
}

func (s *AnalyzeGitRepositoryRequest) GetRepoId() *string {
	return s.RepoId
}

func (s *AnalyzeGitRepositoryRequest) SetBranch(v string) *AnalyzeGitRepositoryRequest {
	s.Branch = &v
	return s
}

func (s *AnalyzeGitRepositoryRequest) SetClientToken(v string) *AnalyzeGitRepositoryRequest {
	s.ClientToken = &v
	return s
}

func (s *AnalyzeGitRepositoryRequest) SetOrgId(v string) *AnalyzeGitRepositoryRequest {
	s.OrgId = &v
	return s
}

func (s *AnalyzeGitRepositoryRequest) SetOwner(v string) *AnalyzeGitRepositoryRequest {
	s.Owner = &v
	return s
}

func (s *AnalyzeGitRepositoryRequest) SetPlatform(v string) *AnalyzeGitRepositoryRequest {
	s.Platform = &v
	return s
}

func (s *AnalyzeGitRepositoryRequest) SetRegionId(v string) *AnalyzeGitRepositoryRequest {
	s.RegionId = &v
	return s
}

func (s *AnalyzeGitRepositoryRequest) SetRepoFullName(v string) *AnalyzeGitRepositoryRequest {
	s.RepoFullName = &v
	return s
}

func (s *AnalyzeGitRepositoryRequest) SetRepoId(v string) *AnalyzeGitRepositoryRequest {
	s.RepoId = &v
	return s
}

func (s *AnalyzeGitRepositoryRequest) Validate() error {
	return dara.Validate(s)
}
