// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGitRepositoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateGitRepositoryRequest
	GetClientToken() *string
	SetIsPrivate(v bool) *CreateGitRepositoryRequest
	GetIsPrivate() *bool
	SetOrgId(v string) *CreateGitRepositoryRequest
	GetOrgId() *string
	SetOwner(v string) *CreateGitRepositoryRequest
	GetOwner() *string
	SetPlatform(v string) *CreateGitRepositoryRequest
	GetPlatform() *string
	SetRegionId(v string) *CreateGitRepositoryRequest
	GetRegionId() *string
	SetSourceRepoBranch(v string) *CreateGitRepositoryRequest
	GetSourceRepoBranch() *string
	SetSourceRepoName(v string) *CreateGitRepositoryRequest
	GetSourceRepoName() *string
	SetSourceRepoOwner(v string) *CreateGitRepositoryRequest
	GetSourceRepoOwner() *string
	SetTargetRepoName(v string) *CreateGitRepositoryRequest
	GetTargetRepoName() *string
	SetTargetRepoOwner(v string) *CreateGitRepositoryRequest
	GetTargetRepoOwner() *string
}

type CreateGitRepositoryRequest struct {
	// example:
	//
	// TF-CreateApplication-1647587475-84104b89-eba5-47a8-b2fd-807b8b7d
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// False
	IsPrivate *bool `json:"IsPrivate,omitempty" xml:"IsPrivate,omitempty"`
	// example:
	//
	// 5ffd468b1e45db3c1cc26ad6
	OrgId *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// JJGGu
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
	// example:
	//
	// main
	SourceRepoBranch *string `json:"SourceRepoBranch,omitempty" xml:"SourceRepoBranch,omitempty"`
	// example:
	//
	// test-repo
	SourceRepoName *string `json:"SourceRepoName,omitempty" xml:"SourceRepoName,omitempty"`
	// example:
	//
	// aliyun-computenest
	SourceRepoOwner *string `json:"SourceRepoOwner,omitempty" xml:"SourceRepoOwner,omitempty"`
	// example:
	//
	// test-repo
	TargetRepoName *string `json:"TargetRepoName,omitempty" xml:"TargetRepoName,omitempty"`
	// example:
	//
	// namexx
	TargetRepoOwner *string `json:"TargetRepoOwner,omitempty" xml:"TargetRepoOwner,omitempty"`
}

func (s CreateGitRepositoryRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateGitRepositoryRequest) GoString() string {
	return s.String()
}

func (s *CreateGitRepositoryRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateGitRepositoryRequest) GetIsPrivate() *bool {
	return s.IsPrivate
}

func (s *CreateGitRepositoryRequest) GetOrgId() *string {
	return s.OrgId
}

func (s *CreateGitRepositoryRequest) GetOwner() *string {
	return s.Owner
}

func (s *CreateGitRepositoryRequest) GetPlatform() *string {
	return s.Platform
}

func (s *CreateGitRepositoryRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateGitRepositoryRequest) GetSourceRepoBranch() *string {
	return s.SourceRepoBranch
}

func (s *CreateGitRepositoryRequest) GetSourceRepoName() *string {
	return s.SourceRepoName
}

func (s *CreateGitRepositoryRequest) GetSourceRepoOwner() *string {
	return s.SourceRepoOwner
}

func (s *CreateGitRepositoryRequest) GetTargetRepoName() *string {
	return s.TargetRepoName
}

func (s *CreateGitRepositoryRequest) GetTargetRepoOwner() *string {
	return s.TargetRepoOwner
}

func (s *CreateGitRepositoryRequest) SetClientToken(v string) *CreateGitRepositoryRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateGitRepositoryRequest) SetIsPrivate(v bool) *CreateGitRepositoryRequest {
	s.IsPrivate = &v
	return s
}

func (s *CreateGitRepositoryRequest) SetOrgId(v string) *CreateGitRepositoryRequest {
	s.OrgId = &v
	return s
}

func (s *CreateGitRepositoryRequest) SetOwner(v string) *CreateGitRepositoryRequest {
	s.Owner = &v
	return s
}

func (s *CreateGitRepositoryRequest) SetPlatform(v string) *CreateGitRepositoryRequest {
	s.Platform = &v
	return s
}

func (s *CreateGitRepositoryRequest) SetRegionId(v string) *CreateGitRepositoryRequest {
	s.RegionId = &v
	return s
}

func (s *CreateGitRepositoryRequest) SetSourceRepoBranch(v string) *CreateGitRepositoryRequest {
	s.SourceRepoBranch = &v
	return s
}

func (s *CreateGitRepositoryRequest) SetSourceRepoName(v string) *CreateGitRepositoryRequest {
	s.SourceRepoName = &v
	return s
}

func (s *CreateGitRepositoryRequest) SetSourceRepoOwner(v string) *CreateGitRepositoryRequest {
	s.SourceRepoOwner = &v
	return s
}

func (s *CreateGitRepositoryRequest) SetTargetRepoName(v string) *CreateGitRepositoryRequest {
	s.TargetRepoName = &v
	return s
}

func (s *CreateGitRepositoryRequest) SetTargetRepoOwner(v string) *CreateGitRepositoryRequest {
	s.TargetRepoOwner = &v
	return s
}

func (s *CreateGitRepositoryRequest) Validate() error {
	return dara.Validate(s)
}
