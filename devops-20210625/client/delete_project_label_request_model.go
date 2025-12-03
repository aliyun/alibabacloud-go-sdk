// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteProjectLabelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessToken(v string) *DeleteProjectLabelRequest
	GetAccessToken() *string
	SetOrganizationId(v string) *DeleteProjectLabelRequest
	GetOrganizationId() *string
	SetRepositoryIdentity(v string) *DeleteProjectLabelRequest
	GetRepositoryIdentity() *string
}

type DeleteProjectLabelRequest struct {
	// example:
	//
	// f0b1e61db5961df5975a93f9129d2513
	AccessToken *string `json:"accessToken,omitempty" xml:"accessToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 5ebbc0228123212b59xxxxx
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
	// This parameter is required.
	RepositoryIdentity *string `json:"repositoryIdentity,omitempty" xml:"repositoryIdentity,omitempty"`
}

func (s DeleteProjectLabelRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteProjectLabelRequest) GoString() string {
	return s.String()
}

func (s *DeleteProjectLabelRequest) GetAccessToken() *string {
	return s.AccessToken
}

func (s *DeleteProjectLabelRequest) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *DeleteProjectLabelRequest) GetRepositoryIdentity() *string {
	return s.RepositoryIdentity
}

func (s *DeleteProjectLabelRequest) SetAccessToken(v string) *DeleteProjectLabelRequest {
	s.AccessToken = &v
	return s
}

func (s *DeleteProjectLabelRequest) SetOrganizationId(v string) *DeleteProjectLabelRequest {
	s.OrganizationId = &v
	return s
}

func (s *DeleteProjectLabelRequest) SetRepositoryIdentity(v string) *DeleteProjectLabelRequest {
	s.RepositoryIdentity = &v
	return s
}

func (s *DeleteProjectLabelRequest) Validate() error {
	return dara.Validate(s)
}
