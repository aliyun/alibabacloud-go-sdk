// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRepositoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessToken(v string) *GetRepositoryRequest
	GetAccessToken() *string
	SetIdentity(v string) *GetRepositoryRequest
	GetIdentity() *string
	SetOrganizationId(v string) *GetRepositoryRequest
	GetOrganizationId() *string
}

type GetRepositoryRequest struct {
	// example:
	//
	// f0b1e61db5961df5975a93f9129d2513
	AccessToken *string `json:"accessToken,omitempty" xml:"accessToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 100
	Identity *string `json:"identity,omitempty" xml:"identity,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 60de7a6852743a5162b5f957
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
}

func (s GetRepositoryRequest) String() string {
	return dara.Prettify(s)
}

func (s GetRepositoryRequest) GoString() string {
	return s.String()
}

func (s *GetRepositoryRequest) GetAccessToken() *string {
	return s.AccessToken
}

func (s *GetRepositoryRequest) GetIdentity() *string {
	return s.Identity
}

func (s *GetRepositoryRequest) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *GetRepositoryRequest) SetAccessToken(v string) *GetRepositoryRequest {
	s.AccessToken = &v
	return s
}

func (s *GetRepositoryRequest) SetIdentity(v string) *GetRepositoryRequest {
	s.Identity = &v
	return s
}

func (s *GetRepositoryRequest) SetOrganizationId(v string) *GetRepositoryRequest {
	s.OrganizationId = &v
	return s
}

func (s *GetRepositoryRequest) Validate() error {
	return dara.Validate(s)
}
