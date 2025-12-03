// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRepositoryMemberRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessToken(v string) *DeleteRepositoryMemberRequest
	GetAccessToken() *string
	SetMemberType(v string) *DeleteRepositoryMemberRequest
	GetMemberType() *string
	SetOrganizationId(v string) *DeleteRepositoryMemberRequest
	GetOrganizationId() *string
}

type DeleteRepositoryMemberRequest struct {
	// example:
	//
	// f0b1e61db5961df5975a93f9129d2513
	AccessToken *string `json:"accessToken,omitempty" xml:"accessToken,omitempty"`
	// example:
	//
	// USERS
	MemberType *string `json:"memberType,omitempty" xml:"memberType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 632bbfdf419338aaa2b1360a
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
}

func (s DeleteRepositoryMemberRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteRepositoryMemberRequest) GoString() string {
	return s.String()
}

func (s *DeleteRepositoryMemberRequest) GetAccessToken() *string {
	return s.AccessToken
}

func (s *DeleteRepositoryMemberRequest) GetMemberType() *string {
	return s.MemberType
}

func (s *DeleteRepositoryMemberRequest) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *DeleteRepositoryMemberRequest) SetAccessToken(v string) *DeleteRepositoryMemberRequest {
	s.AccessToken = &v
	return s
}

func (s *DeleteRepositoryMemberRequest) SetMemberType(v string) *DeleteRepositoryMemberRequest {
	s.MemberType = &v
	return s
}

func (s *DeleteRepositoryMemberRequest) SetOrganizationId(v string) *DeleteRepositoryMemberRequest {
	s.OrganizationId = &v
	return s
}

func (s *DeleteRepositoryMemberRequest) Validate() error {
	return dara.Validate(s)
}
