// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddGroupMemberRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessToken(v string) *AddGroupMemberRequest
	GetAccessToken() *string
	SetAccessLevel(v int32) *AddGroupMemberRequest
	GetAccessLevel() *int32
	SetAliyunPks(v string) *AddGroupMemberRequest
	GetAliyunPks() *string
	SetOrganizationId(v string) *AddGroupMemberRequest
	GetOrganizationId() *string
}

type AddGroupMemberRequest struct {
	// example:
	//
	// f0b1e61db5961df5975a93f9129d2513
	AccessToken *string `json:"accessToken,omitempty" xml:"accessToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30
	AccessLevel *int32 `json:"accessLevel,omitempty" xml:"accessLevel,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 18745637472884
	AliyunPks *string `json:"aliyunPks,omitempty" xml:"aliyunPks,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 60de7a6852743a5162b5f957
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
}

func (s AddGroupMemberRequest) String() string {
	return dara.Prettify(s)
}

func (s AddGroupMemberRequest) GoString() string {
	return s.String()
}

func (s *AddGroupMemberRequest) GetAccessToken() *string {
	return s.AccessToken
}

func (s *AddGroupMemberRequest) GetAccessLevel() *int32 {
	return s.AccessLevel
}

func (s *AddGroupMemberRequest) GetAliyunPks() *string {
	return s.AliyunPks
}

func (s *AddGroupMemberRequest) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *AddGroupMemberRequest) SetAccessToken(v string) *AddGroupMemberRequest {
	s.AccessToken = &v
	return s
}

func (s *AddGroupMemberRequest) SetAccessLevel(v int32) *AddGroupMemberRequest {
	s.AccessLevel = &v
	return s
}

func (s *AddGroupMemberRequest) SetAliyunPks(v string) *AddGroupMemberRequest {
	s.AliyunPks = &v
	return s
}

func (s *AddGroupMemberRequest) SetOrganizationId(v string) *AddGroupMemberRequest {
	s.OrganizationId = &v
	return s
}

func (s *AddGroupMemberRequest) Validate() error {
	return dara.Validate(s)
}
