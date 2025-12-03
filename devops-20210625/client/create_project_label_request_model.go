// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateProjectLabelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessToken(v string) *CreateProjectLabelRequest
	GetAccessToken() *string
	SetColor(v string) *CreateProjectLabelRequest
	GetColor() *string
	SetDescription(v string) *CreateProjectLabelRequest
	GetDescription() *string
	SetName(v string) *CreateProjectLabelRequest
	GetName() *string
	SetOrganizationId(v string) *CreateProjectLabelRequest
	GetOrganizationId() *string
	SetRepositoryIdentity(v string) *CreateProjectLabelRequest
	GetRepositoryIdentity() *string
}

type CreateProjectLabelRequest struct {
	// example:
	//
	// f0b1e61db5961df5975a93f9129d2513
	AccessToken *string `json:"accessToken,omitempty" xml:"accessToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// #006AD4
	Color       *string `json:"color,omitempty" xml:"color,omitempty"`
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// This parameter is required.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 5ebbc0228123212b59xxxxx
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
	// This parameter is required.
	RepositoryIdentity *string `json:"repositoryIdentity,omitempty" xml:"repositoryIdentity,omitempty"`
}

func (s CreateProjectLabelRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateProjectLabelRequest) GoString() string {
	return s.String()
}

func (s *CreateProjectLabelRequest) GetAccessToken() *string {
	return s.AccessToken
}

func (s *CreateProjectLabelRequest) GetColor() *string {
	return s.Color
}

func (s *CreateProjectLabelRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateProjectLabelRequest) GetName() *string {
	return s.Name
}

func (s *CreateProjectLabelRequest) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *CreateProjectLabelRequest) GetRepositoryIdentity() *string {
	return s.RepositoryIdentity
}

func (s *CreateProjectLabelRequest) SetAccessToken(v string) *CreateProjectLabelRequest {
	s.AccessToken = &v
	return s
}

func (s *CreateProjectLabelRequest) SetColor(v string) *CreateProjectLabelRequest {
	s.Color = &v
	return s
}

func (s *CreateProjectLabelRequest) SetDescription(v string) *CreateProjectLabelRequest {
	s.Description = &v
	return s
}

func (s *CreateProjectLabelRequest) SetName(v string) *CreateProjectLabelRequest {
	s.Name = &v
	return s
}

func (s *CreateProjectLabelRequest) SetOrganizationId(v string) *CreateProjectLabelRequest {
	s.OrganizationId = &v
	return s
}

func (s *CreateProjectLabelRequest) SetRepositoryIdentity(v string) *CreateProjectLabelRequest {
	s.RepositoryIdentity = &v
	return s
}

func (s *CreateProjectLabelRequest) Validate() error {
	return dara.Validate(s)
}
