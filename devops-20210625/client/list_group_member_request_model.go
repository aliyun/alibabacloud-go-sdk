// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGroupMemberRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessToken(v string) *ListGroupMemberRequest
	GetAccessToken() *string
	SetOrganizationId(v string) *ListGroupMemberRequest
	GetOrganizationId() *string
}

type ListGroupMemberRequest struct {
	// example:
	//
	// f0b1e61db5961df5975a93f9129d2513
	AccessToken *string `json:"accessToken,omitempty" xml:"accessToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 632bbfdf419338aaa2b1360a
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
}

func (s ListGroupMemberRequest) String() string {
	return dara.Prettify(s)
}

func (s ListGroupMemberRequest) GoString() string {
	return s.String()
}

func (s *ListGroupMemberRequest) GetAccessToken() *string {
	return s.AccessToken
}

func (s *ListGroupMemberRequest) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *ListGroupMemberRequest) SetAccessToken(v string) *ListGroupMemberRequest {
	s.AccessToken = &v
	return s
}

func (s *ListGroupMemberRequest) SetOrganizationId(v string) *ListGroupMemberRequest {
	s.OrganizationId = &v
	return s
}

func (s *ListGroupMemberRequest) Validate() error {
	return dara.Validate(s)
}
