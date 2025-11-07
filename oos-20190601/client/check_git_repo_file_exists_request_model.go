// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckGitRepoFileExistsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBranch(v string) *CheckGitRepoFileExistsRequest
	GetBranch() *string
	SetClientToken(v string) *CheckGitRepoFileExistsRequest
	GetClientToken() *string
	SetFilePath(v string) *CheckGitRepoFileExistsRequest
	GetFilePath() *string
	SetOrgId(v string) *CheckGitRepoFileExistsRequest
	GetOrgId() *string
	SetOwner(v string) *CheckGitRepoFileExistsRequest
	GetOwner() *string
	SetPlatform(v string) *CheckGitRepoFileExistsRequest
	GetPlatform() *string
	SetRegionId(v string) *CheckGitRepoFileExistsRequest
	GetRegionId() *string
	SetRepoFullName(v string) *CheckGitRepoFileExistsRequest
	GetRepoFullName() *string
	SetRepoId(v int64) *CheckGitRepoFileExistsRequest
	GetRepoId() *int64
}

type CheckGitRepoFileExistsRequest struct {
	// example:
	//
	// test
	Branch *string `json:"Branch,omitempty" xml:"Branch,omitempty"`
	// example:
	//
	// -
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// src/main/resources/config/promql_capacity.yaml
	FilePath *string `json:"FilePath,omitempty" xml:"FilePath,omitempty"`
	// example:
	//
	// 5f9f9f6122a8c7ff3934f99a
	OrgId *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// hujue
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
	// 2835387
	RepoId *int64 `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
}

func (s CheckGitRepoFileExistsRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckGitRepoFileExistsRequest) GoString() string {
	return s.String()
}

func (s *CheckGitRepoFileExistsRequest) GetBranch() *string {
	return s.Branch
}

func (s *CheckGitRepoFileExistsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CheckGitRepoFileExistsRequest) GetFilePath() *string {
	return s.FilePath
}

func (s *CheckGitRepoFileExistsRequest) GetOrgId() *string {
	return s.OrgId
}

func (s *CheckGitRepoFileExistsRequest) GetOwner() *string {
	return s.Owner
}

func (s *CheckGitRepoFileExistsRequest) GetPlatform() *string {
	return s.Platform
}

func (s *CheckGitRepoFileExistsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CheckGitRepoFileExistsRequest) GetRepoFullName() *string {
	return s.RepoFullName
}

func (s *CheckGitRepoFileExistsRequest) GetRepoId() *int64 {
	return s.RepoId
}

func (s *CheckGitRepoFileExistsRequest) SetBranch(v string) *CheckGitRepoFileExistsRequest {
	s.Branch = &v
	return s
}

func (s *CheckGitRepoFileExistsRequest) SetClientToken(v string) *CheckGitRepoFileExistsRequest {
	s.ClientToken = &v
	return s
}

func (s *CheckGitRepoFileExistsRequest) SetFilePath(v string) *CheckGitRepoFileExistsRequest {
	s.FilePath = &v
	return s
}

func (s *CheckGitRepoFileExistsRequest) SetOrgId(v string) *CheckGitRepoFileExistsRequest {
	s.OrgId = &v
	return s
}

func (s *CheckGitRepoFileExistsRequest) SetOwner(v string) *CheckGitRepoFileExistsRequest {
	s.Owner = &v
	return s
}

func (s *CheckGitRepoFileExistsRequest) SetPlatform(v string) *CheckGitRepoFileExistsRequest {
	s.Platform = &v
	return s
}

func (s *CheckGitRepoFileExistsRequest) SetRegionId(v string) *CheckGitRepoFileExistsRequest {
	s.RegionId = &v
	return s
}

func (s *CheckGitRepoFileExistsRequest) SetRepoFullName(v string) *CheckGitRepoFileExistsRequest {
	s.RepoFullName = &v
	return s
}

func (s *CheckGitRepoFileExistsRequest) SetRepoId(v int64) *CheckGitRepoFileExistsRequest {
	s.RepoId = &v
	return s
}

func (s *CheckGitRepoFileExistsRequest) Validate() error {
	return dara.Validate(s)
}
