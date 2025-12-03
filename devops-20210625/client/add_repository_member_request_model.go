// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddRepositoryMemberRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessToken(v string) *AddRepositoryMemberRequest
	GetAccessToken() *string
	SetAccessLevel(v int32) *AddRepositoryMemberRequest
	GetAccessLevel() *int32
	SetAliyunPks(v string) *AddRepositoryMemberRequest
	GetAliyunPks() *string
	SetOrganizationId(v string) *AddRepositoryMemberRequest
	GetOrganizationId() *string
}

type AddRepositoryMemberRequest struct {
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

func (s AddRepositoryMemberRequest) String() string {
	return dara.Prettify(s)
}

func (s AddRepositoryMemberRequest) GoString() string {
	return s.String()
}

func (s *AddRepositoryMemberRequest) GetAccessToken() *string {
	return s.AccessToken
}

func (s *AddRepositoryMemberRequest) GetAccessLevel() *int32 {
	return s.AccessLevel
}

func (s *AddRepositoryMemberRequest) GetAliyunPks() *string {
	return s.AliyunPks
}

func (s *AddRepositoryMemberRequest) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *AddRepositoryMemberRequest) SetAccessToken(v string) *AddRepositoryMemberRequest {
	s.AccessToken = &v
	return s
}

func (s *AddRepositoryMemberRequest) SetAccessLevel(v int32) *AddRepositoryMemberRequest {
	s.AccessLevel = &v
	return s
}

func (s *AddRepositoryMemberRequest) SetAliyunPks(v string) *AddRepositoryMemberRequest {
	s.AliyunPks = &v
	return s
}

func (s *AddRepositoryMemberRequest) SetOrganizationId(v string) *AddRepositoryMemberRequest {
	s.OrganizationId = &v
	return s
}

func (s *AddRepositoryMemberRequest) Validate() error {
	return dara.Validate(s)
}
