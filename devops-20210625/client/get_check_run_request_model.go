// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCheckRunRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessToken(v string) *GetCheckRunRequest
	GetAccessToken() *string
	SetCheckRunId(v int64) *GetCheckRunRequest
	GetCheckRunId() *int64
	SetOrganizationId(v string) *GetCheckRunRequest
	GetOrganizationId() *string
	SetRepositoryIdentity(v string) *GetCheckRunRequest
	GetRepositoryIdentity() *string
}

type GetCheckRunRequest struct {
	// example:
	//
	// f0b1e61db5961df5975a93f9129d2513
	AccessToken *string `json:"accessToken,omitempty" xml:"accessToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	CheckRunId *int64 `json:"checkRunId,omitempty" xml:"checkRunId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 5ebbc0228123212b59xxxxx
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
	// This parameter is required.
	RepositoryIdentity *string `json:"repositoryIdentity,omitempty" xml:"repositoryIdentity,omitempty"`
}

func (s GetCheckRunRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCheckRunRequest) GoString() string {
	return s.String()
}

func (s *GetCheckRunRequest) GetAccessToken() *string {
	return s.AccessToken
}

func (s *GetCheckRunRequest) GetCheckRunId() *int64 {
	return s.CheckRunId
}

func (s *GetCheckRunRequest) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *GetCheckRunRequest) GetRepositoryIdentity() *string {
	return s.RepositoryIdentity
}

func (s *GetCheckRunRequest) SetAccessToken(v string) *GetCheckRunRequest {
	s.AccessToken = &v
	return s
}

func (s *GetCheckRunRequest) SetCheckRunId(v int64) *GetCheckRunRequest {
	s.CheckRunId = &v
	return s
}

func (s *GetCheckRunRequest) SetOrganizationId(v string) *GetCheckRunRequest {
	s.OrganizationId = &v
	return s
}

func (s *GetCheckRunRequest) SetRepositoryIdentity(v string) *GetCheckRunRequest {
	s.RepositoryIdentity = &v
	return s
}

func (s *GetCheckRunRequest) Validate() error {
	return dara.Validate(s)
}
