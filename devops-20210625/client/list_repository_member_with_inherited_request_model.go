// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRepositoryMemberWithInheritedRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessToken(v string) *ListRepositoryMemberWithInheritedRequest
	GetAccessToken() *string
	SetOrganizationId(v string) *ListRepositoryMemberWithInheritedRequest
	GetOrganizationId() *string
}

type ListRepositoryMemberWithInheritedRequest struct {
	// example:
	//
	// f0b1e61db5961df5975a93f9129d2513
	AccessToken *string `json:"accessToken,omitempty" xml:"accessToken,omitempty"`
	// This parameter is required.
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
}

func (s ListRepositoryMemberWithInheritedRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRepositoryMemberWithInheritedRequest) GoString() string {
	return s.String()
}

func (s *ListRepositoryMemberWithInheritedRequest) GetAccessToken() *string {
	return s.AccessToken
}

func (s *ListRepositoryMemberWithInheritedRequest) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *ListRepositoryMemberWithInheritedRequest) SetAccessToken(v string) *ListRepositoryMemberWithInheritedRequest {
	s.AccessToken = &v
	return s
}

func (s *ListRepositoryMemberWithInheritedRequest) SetOrganizationId(v string) *ListRepositoryMemberWithInheritedRequest {
	s.OrganizationId = &v
	return s
}

func (s *ListRepositoryMemberWithInheritedRequest) Validate() error {
	return dara.Validate(s)
}
