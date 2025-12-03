// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGroupMemberRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessToken(v string) *DeleteGroupMemberRequest
	GetAccessToken() *string
	SetAliyunPk(v string) *DeleteGroupMemberRequest
	GetAliyunPk() *string
	SetMemberType(v string) *DeleteGroupMemberRequest
	GetMemberType() *string
	SetOrganizationId(v string) *DeleteGroupMemberRequest
	GetOrganizationId() *string
}

type DeleteGroupMemberRequest struct {
	// example:
	//
	// f0b1e61db5961df5975a93f9129d2513
	AccessToken *string `json:"accessToken,omitempty" xml:"accessToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1789095186553536
	AliyunPk *string `json:"aliyunPk,omitempty" xml:"aliyunPk,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// USERS
	MemberType *string `json:"memberType,omitempty" xml:"memberType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 6082a9b0c7972588ac363793
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
}

func (s DeleteGroupMemberRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteGroupMemberRequest) GoString() string {
	return s.String()
}

func (s *DeleteGroupMemberRequest) GetAccessToken() *string {
	return s.AccessToken
}

func (s *DeleteGroupMemberRequest) GetAliyunPk() *string {
	return s.AliyunPk
}

func (s *DeleteGroupMemberRequest) GetMemberType() *string {
	return s.MemberType
}

func (s *DeleteGroupMemberRequest) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *DeleteGroupMemberRequest) SetAccessToken(v string) *DeleteGroupMemberRequest {
	s.AccessToken = &v
	return s
}

func (s *DeleteGroupMemberRequest) SetAliyunPk(v string) *DeleteGroupMemberRequest {
	s.AliyunPk = &v
	return s
}

func (s *DeleteGroupMemberRequest) SetMemberType(v string) *DeleteGroupMemberRequest {
	s.MemberType = &v
	return s
}

func (s *DeleteGroupMemberRequest) SetOrganizationId(v string) *DeleteGroupMemberRequest {
	s.OrganizationId = &v
	return s
}

func (s *DeleteGroupMemberRequest) Validate() error {
	return dara.Validate(s)
}
