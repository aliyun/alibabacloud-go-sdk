// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProjectMemberRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessToken(v string) *GetProjectMemberRequest
	GetAccessToken() *string
	SetOrganizationId(v string) *GetProjectMemberRequest
	GetOrganizationId() *string
}

type GetProjectMemberRequest struct {
	// example:
	//
	// f0b1e61db5961df5975a93f9129d2513
	AccessToken *string `json:"accessToken,omitempty" xml:"accessToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 5eb53bb338076f00011bcfd5
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
}

func (s GetProjectMemberRequest) String() string {
	return dara.Prettify(s)
}

func (s GetProjectMemberRequest) GoString() string {
	return s.String()
}

func (s *GetProjectMemberRequest) GetAccessToken() *string {
	return s.AccessToken
}

func (s *GetProjectMemberRequest) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *GetProjectMemberRequest) SetAccessToken(v string) *GetProjectMemberRequest {
	s.AccessToken = &v
	return s
}

func (s *GetProjectMemberRequest) SetOrganizationId(v string) *GetProjectMemberRequest {
	s.OrganizationId = &v
	return s
}

func (s *GetProjectMemberRequest) Validate() error {
	return dara.Validate(s)
}
