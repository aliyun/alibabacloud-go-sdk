// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCommitStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessToken(v string) *CreateCommitStatusRequest
	GetAccessToken() *string
	SetContext(v string) *CreateCommitStatusRequest
	GetContext() *string
	SetDescription(v string) *CreateCommitStatusRequest
	GetDescription() *string
	SetState(v string) *CreateCommitStatusRequest
	GetState() *string
	SetTargetUrl(v string) *CreateCommitStatusRequest
	GetTargetUrl() *string
	SetOrganizationId(v string) *CreateCommitStatusRequest
	GetOrganizationId() *string
	SetRepositoryIdentity(v string) *CreateCommitStatusRequest
	GetRepositoryIdentity() *string
	SetSha(v string) *CreateCommitStatusRequest
	GetSha() *string
}

type CreateCommitStatusRequest struct {
	// example:
	//
	// f0b1e61db5961df5975a93f9129d2513
	AccessToken *string `json:"accessToken,omitempty" xml:"accessToken,omitempty"`
	// example:
	//
	// default
	Context     *string `json:"context,omitempty" xml:"context,omitempty"`
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// success
	State *string `json:"state,omitempty" xml:"state,omitempty"`
	// example:
	//
	// xxx
	TargetUrl *string `json:"targetUrl,omitempty" xml:"targetUrl,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 5ebbc0228123212b59xxxxx
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2080972
	RepositoryIdentity *string `json:"repositoryIdentity,omitempty" xml:"repositoryIdentity,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// e0c1a1299a2656bfc155650bbd2df5e628fa1f4c
	Sha *string `json:"sha,omitempty" xml:"sha,omitempty"`
}

func (s CreateCommitStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCommitStatusRequest) GoString() string {
	return s.String()
}

func (s *CreateCommitStatusRequest) GetAccessToken() *string {
	return s.AccessToken
}

func (s *CreateCommitStatusRequest) GetContext() *string {
	return s.Context
}

func (s *CreateCommitStatusRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateCommitStatusRequest) GetState() *string {
	return s.State
}

func (s *CreateCommitStatusRequest) GetTargetUrl() *string {
	return s.TargetUrl
}

func (s *CreateCommitStatusRequest) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *CreateCommitStatusRequest) GetRepositoryIdentity() *string {
	return s.RepositoryIdentity
}

func (s *CreateCommitStatusRequest) GetSha() *string {
	return s.Sha
}

func (s *CreateCommitStatusRequest) SetAccessToken(v string) *CreateCommitStatusRequest {
	s.AccessToken = &v
	return s
}

func (s *CreateCommitStatusRequest) SetContext(v string) *CreateCommitStatusRequest {
	s.Context = &v
	return s
}

func (s *CreateCommitStatusRequest) SetDescription(v string) *CreateCommitStatusRequest {
	s.Description = &v
	return s
}

func (s *CreateCommitStatusRequest) SetState(v string) *CreateCommitStatusRequest {
	s.State = &v
	return s
}

func (s *CreateCommitStatusRequest) SetTargetUrl(v string) *CreateCommitStatusRequest {
	s.TargetUrl = &v
	return s
}

func (s *CreateCommitStatusRequest) SetOrganizationId(v string) *CreateCommitStatusRequest {
	s.OrganizationId = &v
	return s
}

func (s *CreateCommitStatusRequest) SetRepositoryIdentity(v string) *CreateCommitStatusRequest {
	s.RepositoryIdentity = &v
	return s
}

func (s *CreateCommitStatusRequest) SetSha(v string) *CreateCommitStatusRequest {
	s.Sha = &v
	return s
}

func (s *CreateCommitStatusRequest) Validate() error {
	return dara.Validate(s)
}
