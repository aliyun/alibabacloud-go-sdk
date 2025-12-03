// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGroupMemberRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessToken(v string) *UpdateGroupMemberRequest
	GetAccessToken() *string
	SetAliyunPk(v string) *UpdateGroupMemberRequest
	GetAliyunPk() *string
	SetAccessLevel(v int32) *UpdateGroupMemberRequest
	GetAccessLevel() *int32
	SetMemberType(v string) *UpdateGroupMemberRequest
	GetMemberType() *string
	SetOrganizationId(v string) *UpdateGroupMemberRequest
	GetOrganizationId() *string
}

type UpdateGroupMemberRequest struct {
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
	// 30
	AccessLevel *int32 `json:"accessLevel,omitempty" xml:"accessLevel,omitempty"`
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

func (s UpdateGroupMemberRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateGroupMemberRequest) GoString() string {
	return s.String()
}

func (s *UpdateGroupMemberRequest) GetAccessToken() *string {
	return s.AccessToken
}

func (s *UpdateGroupMemberRequest) GetAliyunPk() *string {
	return s.AliyunPk
}

func (s *UpdateGroupMemberRequest) GetAccessLevel() *int32 {
	return s.AccessLevel
}

func (s *UpdateGroupMemberRequest) GetMemberType() *string {
	return s.MemberType
}

func (s *UpdateGroupMemberRequest) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *UpdateGroupMemberRequest) SetAccessToken(v string) *UpdateGroupMemberRequest {
	s.AccessToken = &v
	return s
}

func (s *UpdateGroupMemberRequest) SetAliyunPk(v string) *UpdateGroupMemberRequest {
	s.AliyunPk = &v
	return s
}

func (s *UpdateGroupMemberRequest) SetAccessLevel(v int32) *UpdateGroupMemberRequest {
	s.AccessLevel = &v
	return s
}

func (s *UpdateGroupMemberRequest) SetMemberType(v string) *UpdateGroupMemberRequest {
	s.MemberType = &v
	return s
}

func (s *UpdateGroupMemberRequest) SetOrganizationId(v string) *UpdateGroupMemberRequest {
	s.OrganizationId = &v
	return s
}

func (s *UpdateGroupMemberRequest) Validate() error {
	return dara.Validate(s)
}
