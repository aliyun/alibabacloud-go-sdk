// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckGitRepositoryExistsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CheckGitRepositoryExistsRequest
	GetClientToken() *string
	SetOrgId(v string) *CheckGitRepositoryExistsRequest
	GetOrgId() *string
	SetOwner(v string) *CheckGitRepositoryExistsRequest
	GetOwner() *string
	SetPlatform(v string) *CheckGitRepositoryExistsRequest
	GetPlatform() *string
	SetRegionId(v string) *CheckGitRepositoryExistsRequest
	GetRegionId() *string
	SetRepoFullName(v string) *CheckGitRepositoryExistsRequest
	GetRepoFullName() *string
	SetRepoId(v int64) *CheckGitRepositoryExistsRequest
	GetRepoId() *int64
}

type CheckGitRepositoryExistsRequest struct {
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// 5ffd468b1e45db3c1cc26ad6
	OrgId *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ir9SK9n1u
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
	// LYH-RAIN/vue-color-avatar
	RepoFullName *string `json:"RepoFullName,omitempty" xml:"RepoFullName,omitempty"`
	// example:
	//
	// 2859382
	RepoId *int64 `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
}

func (s CheckGitRepositoryExistsRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckGitRepositoryExistsRequest) GoString() string {
	return s.String()
}

func (s *CheckGitRepositoryExistsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CheckGitRepositoryExistsRequest) GetOrgId() *string {
	return s.OrgId
}

func (s *CheckGitRepositoryExistsRequest) GetOwner() *string {
	return s.Owner
}

func (s *CheckGitRepositoryExistsRequest) GetPlatform() *string {
	return s.Platform
}

func (s *CheckGitRepositoryExistsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CheckGitRepositoryExistsRequest) GetRepoFullName() *string {
	return s.RepoFullName
}

func (s *CheckGitRepositoryExistsRequest) GetRepoId() *int64 {
	return s.RepoId
}

func (s *CheckGitRepositoryExistsRequest) SetClientToken(v string) *CheckGitRepositoryExistsRequest {
	s.ClientToken = &v
	return s
}

func (s *CheckGitRepositoryExistsRequest) SetOrgId(v string) *CheckGitRepositoryExistsRequest {
	s.OrgId = &v
	return s
}

func (s *CheckGitRepositoryExistsRequest) SetOwner(v string) *CheckGitRepositoryExistsRequest {
	s.Owner = &v
	return s
}

func (s *CheckGitRepositoryExistsRequest) SetPlatform(v string) *CheckGitRepositoryExistsRequest {
	s.Platform = &v
	return s
}

func (s *CheckGitRepositoryExistsRequest) SetRegionId(v string) *CheckGitRepositoryExistsRequest {
	s.RegionId = &v
	return s
}

func (s *CheckGitRepositoryExistsRequest) SetRepoFullName(v string) *CheckGitRepositoryExistsRequest {
	s.RepoFullName = &v
	return s
}

func (s *CheckGitRepositoryExistsRequest) SetRepoId(v int64) *CheckGitRepositoryExistsRequest {
	s.RepoId = &v
	return s
}

func (s *CheckGitRepositoryExistsRequest) Validate() error {
	return dara.Validate(s)
}
