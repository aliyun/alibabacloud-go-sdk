// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateProjectLabelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessToken(v string) *UpdateProjectLabelRequest
	GetAccessToken() *string
	SetColor(v string) *UpdateProjectLabelRequest
	GetColor() *string
	SetDescription(v string) *UpdateProjectLabelRequest
	GetDescription() *string
	SetName(v string) *UpdateProjectLabelRequest
	GetName() *string
	SetOrganizationId(v string) *UpdateProjectLabelRequest
	GetOrganizationId() *string
	SetRepositoryIdentity(v string) *UpdateProjectLabelRequest
	GetRepositoryIdentity() *string
}

type UpdateProjectLabelRequest struct {
	// example:
	//
	// f0b1e61db5961df5975a93f9129d2513
	AccessToken *string `json:"accessToken,omitempty" xml:"accessToken,omitempty"`
	// example:
	//
	// #EF433B
	Color       *string `json:"color,omitempty" xml:"color,omitempty"`
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	Name        *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 5ebbc0228123212b59xxxxx
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
	// This parameter is required.
	RepositoryIdentity *string `json:"repositoryIdentity,omitempty" xml:"repositoryIdentity,omitempty"`
}

func (s UpdateProjectLabelRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateProjectLabelRequest) GoString() string {
	return s.String()
}

func (s *UpdateProjectLabelRequest) GetAccessToken() *string {
	return s.AccessToken
}

func (s *UpdateProjectLabelRequest) GetColor() *string {
	return s.Color
}

func (s *UpdateProjectLabelRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateProjectLabelRequest) GetName() *string {
	return s.Name
}

func (s *UpdateProjectLabelRequest) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *UpdateProjectLabelRequest) GetRepositoryIdentity() *string {
	return s.RepositoryIdentity
}

func (s *UpdateProjectLabelRequest) SetAccessToken(v string) *UpdateProjectLabelRequest {
	s.AccessToken = &v
	return s
}

func (s *UpdateProjectLabelRequest) SetColor(v string) *UpdateProjectLabelRequest {
	s.Color = &v
	return s
}

func (s *UpdateProjectLabelRequest) SetDescription(v string) *UpdateProjectLabelRequest {
	s.Description = &v
	return s
}

func (s *UpdateProjectLabelRequest) SetName(v string) *UpdateProjectLabelRequest {
	s.Name = &v
	return s
}

func (s *UpdateProjectLabelRequest) SetOrganizationId(v string) *UpdateProjectLabelRequest {
	s.OrganizationId = &v
	return s
}

func (s *UpdateProjectLabelRequest) SetRepositoryIdentity(v string) *UpdateProjectLabelRequest {
	s.RepositoryIdentity = &v
	return s
}

func (s *UpdateProjectLabelRequest) Validate() error {
	return dara.Validate(s)
}
