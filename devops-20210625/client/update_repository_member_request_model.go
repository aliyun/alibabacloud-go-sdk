// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRepositoryMemberRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessToken(v string) *UpdateRepositoryMemberRequest
	GetAccessToken() *string
	SetAccessLevel(v int32) *UpdateRepositoryMemberRequest
	GetAccessLevel() *int32
	SetExpireAt(v string) *UpdateRepositoryMemberRequest
	GetExpireAt() *string
	SetMemberType(v string) *UpdateRepositoryMemberRequest
	GetMemberType() *string
	SetRelatedId(v string) *UpdateRepositoryMemberRequest
	GetRelatedId() *string
	SetRelatedInfos(v []*UpdateRepositoryMemberRequestRelatedInfos) *UpdateRepositoryMemberRequest
	GetRelatedInfos() []*UpdateRepositoryMemberRequestRelatedInfos
	SetOrganizationId(v string) *UpdateRepositoryMemberRequest
	GetOrganizationId() *string
}

type UpdateRepositoryMemberRequest struct {
	// example:
	//
	// f0b1e61db5961df5975a93f9129d2513
	AccessToken *string `json:"accessToken,omitempty" xml:"accessToken,omitempty"`
	// example:
	//
	// 30
	AccessLevel *int32 `json:"accessLevel,omitempty" xml:"accessLevel,omitempty"`
	// example:
	//
	// 2020-08-08 08:08:08
	ExpireAt *string `json:"expireAt,omitempty" xml:"expireAt,omitempty"`
	// example:
	//
	// USERS
	MemberType *string `json:"memberType,omitempty" xml:"memberType,omitempty"`
	// example:
	//
	// 10010
	RelatedId    *string                                      `json:"relatedId,omitempty" xml:"relatedId,omitempty"`
	RelatedInfos []*UpdateRepositoryMemberRequestRelatedInfos `json:"relatedInfos,omitempty" xml:"relatedInfos,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 60de7a6852743a5162b5f957
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
}

func (s UpdateRepositoryMemberRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateRepositoryMemberRequest) GoString() string {
	return s.String()
}

func (s *UpdateRepositoryMemberRequest) GetAccessToken() *string {
	return s.AccessToken
}

func (s *UpdateRepositoryMemberRequest) GetAccessLevel() *int32 {
	return s.AccessLevel
}

func (s *UpdateRepositoryMemberRequest) GetExpireAt() *string {
	return s.ExpireAt
}

func (s *UpdateRepositoryMemberRequest) GetMemberType() *string {
	return s.MemberType
}

func (s *UpdateRepositoryMemberRequest) GetRelatedId() *string {
	return s.RelatedId
}

func (s *UpdateRepositoryMemberRequest) GetRelatedInfos() []*UpdateRepositoryMemberRequestRelatedInfos {
	return s.RelatedInfos
}

func (s *UpdateRepositoryMemberRequest) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *UpdateRepositoryMemberRequest) SetAccessToken(v string) *UpdateRepositoryMemberRequest {
	s.AccessToken = &v
	return s
}

func (s *UpdateRepositoryMemberRequest) SetAccessLevel(v int32) *UpdateRepositoryMemberRequest {
	s.AccessLevel = &v
	return s
}

func (s *UpdateRepositoryMemberRequest) SetExpireAt(v string) *UpdateRepositoryMemberRequest {
	s.ExpireAt = &v
	return s
}

func (s *UpdateRepositoryMemberRequest) SetMemberType(v string) *UpdateRepositoryMemberRequest {
	s.MemberType = &v
	return s
}

func (s *UpdateRepositoryMemberRequest) SetRelatedId(v string) *UpdateRepositoryMemberRequest {
	s.RelatedId = &v
	return s
}

func (s *UpdateRepositoryMemberRequest) SetRelatedInfos(v []*UpdateRepositoryMemberRequestRelatedInfos) *UpdateRepositoryMemberRequest {
	s.RelatedInfos = v
	return s
}

func (s *UpdateRepositoryMemberRequest) SetOrganizationId(v string) *UpdateRepositoryMemberRequest {
	s.OrganizationId = &v
	return s
}

func (s *UpdateRepositoryMemberRequest) Validate() error {
	if s.RelatedInfos != nil {
		for _, item := range s.RelatedInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateRepositoryMemberRequestRelatedInfos struct {
	// example:
	//
	// 10011
	RelatedId *string `json:"relatedId,omitempty" xml:"relatedId,omitempty"`
	// example:
	//
	// 24790
	SourceId *int64 `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// example:
	//
	// Project
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
}

func (s UpdateRepositoryMemberRequestRelatedInfos) String() string {
	return dara.Prettify(s)
}

func (s UpdateRepositoryMemberRequestRelatedInfos) GoString() string {
	return s.String()
}

func (s *UpdateRepositoryMemberRequestRelatedInfos) GetRelatedId() *string {
	return s.RelatedId
}

func (s *UpdateRepositoryMemberRequestRelatedInfos) GetSourceId() *int64 {
	return s.SourceId
}

func (s *UpdateRepositoryMemberRequestRelatedInfos) GetSourceType() *string {
	return s.SourceType
}

func (s *UpdateRepositoryMemberRequestRelatedInfos) SetRelatedId(v string) *UpdateRepositoryMemberRequestRelatedInfos {
	s.RelatedId = &v
	return s
}

func (s *UpdateRepositoryMemberRequestRelatedInfos) SetSourceId(v int64) *UpdateRepositoryMemberRequestRelatedInfos {
	s.SourceId = &v
	return s
}

func (s *UpdateRepositoryMemberRequestRelatedInfos) SetSourceType(v string) *UpdateRepositoryMemberRequestRelatedInfos {
	s.SourceType = &v
	return s
}

func (s *UpdateRepositoryMemberRequestRelatedInfos) Validate() error {
	return dara.Validate(s)
}
