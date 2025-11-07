// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGitBranchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBranch(v string) *GetGitBranchRequest
	GetBranch() *string
	SetOrgId(v string) *GetGitBranchRequest
	GetOrgId() *string
	SetOwner(v string) *GetGitBranchRequest
	GetOwner() *string
	SetPlatform(v string) *GetGitBranchRequest
	GetPlatform() *string
	SetRegionId(v string) *GetGitBranchRequest
	GetRegionId() *string
	SetRepoFullName(v string) *GetGitBranchRequest
	GetRepoFullName() *string
	SetRepoId(v int64) *GetGitBranchRequest
	GetRepoId() *int64
}

type GetGitBranchRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// main
	Branch *string `json:"Branch,omitempty" xml:"Branch,omitempty"`
	// example:
	//
	// 5ffd468b1e45db3c1cc26ad6
	OrgId *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	// example:
	//
	// namexxx
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
	// 2823742
	RepoId *int64 `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
}

func (s GetGitBranchRequest) String() string {
	return dara.Prettify(s)
}

func (s GetGitBranchRequest) GoString() string {
	return s.String()
}

func (s *GetGitBranchRequest) GetBranch() *string {
	return s.Branch
}

func (s *GetGitBranchRequest) GetOrgId() *string {
	return s.OrgId
}

func (s *GetGitBranchRequest) GetOwner() *string {
	return s.Owner
}

func (s *GetGitBranchRequest) GetPlatform() *string {
	return s.Platform
}

func (s *GetGitBranchRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetGitBranchRequest) GetRepoFullName() *string {
	return s.RepoFullName
}

func (s *GetGitBranchRequest) GetRepoId() *int64 {
	return s.RepoId
}

func (s *GetGitBranchRequest) SetBranch(v string) *GetGitBranchRequest {
	s.Branch = &v
	return s
}

func (s *GetGitBranchRequest) SetOrgId(v string) *GetGitBranchRequest {
	s.OrgId = &v
	return s
}

func (s *GetGitBranchRequest) SetOwner(v string) *GetGitBranchRequest {
	s.Owner = &v
	return s
}

func (s *GetGitBranchRequest) SetPlatform(v string) *GetGitBranchRequest {
	s.Platform = &v
	return s
}

func (s *GetGitBranchRequest) SetRegionId(v string) *GetGitBranchRequest {
	s.RegionId = &v
	return s
}

func (s *GetGitBranchRequest) SetRepoFullName(v string) *GetGitBranchRequest {
	s.RepoFullName = &v
	return s
}

func (s *GetGitBranchRequest) SetRepoId(v int64) *GetGitBranchRequest {
	s.RepoId = &v
	return s
}

func (s *GetGitBranchRequest) Validate() error {
	return dara.Validate(s)
}
