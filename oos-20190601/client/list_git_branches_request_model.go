// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGitBranchesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ListGitBranchesRequest
	GetClientToken() *string
	SetOrgId(v string) *ListGitBranchesRequest
	GetOrgId() *string
	SetOwner(v string) *ListGitBranchesRequest
	GetOwner() *string
	SetPlatform(v string) *ListGitBranchesRequest
	GetPlatform() *string
	SetRegionId(v string) *ListGitBranchesRequest
	GetRegionId() *string
	SetRepoFullName(v string) *ListGitBranchesRequest
	GetRepoFullName() *string
	SetRepoId(v int64) *ListGitBranchesRequest
	GetRepoId() *int64
}

type ListGitBranchesRequest struct {
	// example:
	//
	// TF-CreateApplication-1647587475-84104b89-eba5-47a8-b2fd-807b8b7d
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// 6fekfmnfll135123kdf33
	OrgId *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// geegenw-j-imwinmtuej
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// alibaba/fastjson
	RepoFullName *string `json:"RepoFullName,omitempty" xml:"RepoFullName,omitempty"`
	// example:
	//
	// 2642213
	RepoId *int64 `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
}

func (s ListGitBranchesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListGitBranchesRequest) GoString() string {
	return s.String()
}

func (s *ListGitBranchesRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ListGitBranchesRequest) GetOrgId() *string {
	return s.OrgId
}

func (s *ListGitBranchesRequest) GetOwner() *string {
	return s.Owner
}

func (s *ListGitBranchesRequest) GetPlatform() *string {
	return s.Platform
}

func (s *ListGitBranchesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListGitBranchesRequest) GetRepoFullName() *string {
	return s.RepoFullName
}

func (s *ListGitBranchesRequest) GetRepoId() *int64 {
	return s.RepoId
}

func (s *ListGitBranchesRequest) SetClientToken(v string) *ListGitBranchesRequest {
	s.ClientToken = &v
	return s
}

func (s *ListGitBranchesRequest) SetOrgId(v string) *ListGitBranchesRequest {
	s.OrgId = &v
	return s
}

func (s *ListGitBranchesRequest) SetOwner(v string) *ListGitBranchesRequest {
	s.Owner = &v
	return s
}

func (s *ListGitBranchesRequest) SetPlatform(v string) *ListGitBranchesRequest {
	s.Platform = &v
	return s
}

func (s *ListGitBranchesRequest) SetRegionId(v string) *ListGitBranchesRequest {
	s.RegionId = &v
	return s
}

func (s *ListGitBranchesRequest) SetRepoFullName(v string) *ListGitBranchesRequest {
	s.RepoFullName = &v
	return s
}

func (s *ListGitBranchesRequest) SetRepoId(v int64) *ListGitBranchesRequest {
	s.RepoId = &v
	return s
}

func (s *ListGitBranchesRequest) Validate() error {
	return dara.Validate(s)
}
