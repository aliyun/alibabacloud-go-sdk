// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRepositoryCommitRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessToken(v string) *GetRepositoryCommitRequest
	GetAccessToken() *string
	SetOrganizationId(v string) *GetRepositoryCommitRequest
	GetOrganizationId() *string
	SetShowSignature(v bool) *GetRepositoryCommitRequest
	GetShowSignature() *bool
}

type GetRepositoryCommitRequest struct {
	// example:
	//
	// f0b1e61db5961df5975a93f9129d2513
	AccessToken *string `json:"accessToken,omitempty" xml:"accessToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 60de7a6852743a5162b5f957
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
	// example:
	//
	// false
	ShowSignature *bool `json:"showSignature,omitempty" xml:"showSignature,omitempty"`
}

func (s GetRepositoryCommitRequest) String() string {
	return dara.Prettify(s)
}

func (s GetRepositoryCommitRequest) GoString() string {
	return s.String()
}

func (s *GetRepositoryCommitRequest) GetAccessToken() *string {
	return s.AccessToken
}

func (s *GetRepositoryCommitRequest) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *GetRepositoryCommitRequest) GetShowSignature() *bool {
	return s.ShowSignature
}

func (s *GetRepositoryCommitRequest) SetAccessToken(v string) *GetRepositoryCommitRequest {
	s.AccessToken = &v
	return s
}

func (s *GetRepositoryCommitRequest) SetOrganizationId(v string) *GetRepositoryCommitRequest {
	s.OrganizationId = &v
	return s
}

func (s *GetRepositoryCommitRequest) SetShowSignature(v bool) *GetRepositoryCommitRequest {
	s.ShowSignature = &v
	return s
}

func (s *GetRepositoryCommitRequest) Validate() error {
	return dara.Validate(s)
}
