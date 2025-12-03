// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRepositoryGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessToken(v string) *CreateRepositoryGroupRequest
	GetAccessToken() *string
	SetAvatarUrl(v string) *CreateRepositoryGroupRequest
	GetAvatarUrl() *string
	SetDescription(v string) *CreateRepositoryGroupRequest
	GetDescription() *string
	SetName(v string) *CreateRepositoryGroupRequest
	GetName() *string
	SetParentId(v int64) *CreateRepositoryGroupRequest
	GetParentId() *int64
	SetPath(v string) *CreateRepositoryGroupRequest
	GetPath() *string
	SetVisibilityLevel(v int32) *CreateRepositoryGroupRequest
	GetVisibilityLevel() *int32
	SetOrganizationId(v string) *CreateRepositoryGroupRequest
	GetOrganizationId() *string
}

type CreateRepositoryGroupRequest struct {
	// example:
	//
	// f0b1e61db5961df5975a93f9129d2513
	AccessToken *string `json:"accessToken,omitempty" xml:"accessToken,omitempty"`
	// example:
	//
	// https://tcs-devops.aliyuncs.com/thumbnail/112afcb7a6a35c3f67f1bea827c4/w/100/h/100
	AvatarUrl   *string `json:"avatarUrl,omitempty" xml:"avatarUrl,omitempty"`
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test-create-group
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 26842
	ParentId *int64 `json:"parentId,omitempty" xml:"parentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test-create-group
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	VisibilityLevel *int32 `json:"visibilityLevel,omitempty" xml:"visibilityLevel,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 60de7a6852743a5162b5f957
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
}

func (s CreateRepositoryGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRepositoryGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateRepositoryGroupRequest) GetAccessToken() *string {
	return s.AccessToken
}

func (s *CreateRepositoryGroupRequest) GetAvatarUrl() *string {
	return s.AvatarUrl
}

func (s *CreateRepositoryGroupRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateRepositoryGroupRequest) GetName() *string {
	return s.Name
}

func (s *CreateRepositoryGroupRequest) GetParentId() *int64 {
	return s.ParentId
}

func (s *CreateRepositoryGroupRequest) GetPath() *string {
	return s.Path
}

func (s *CreateRepositoryGroupRequest) GetVisibilityLevel() *int32 {
	return s.VisibilityLevel
}

func (s *CreateRepositoryGroupRequest) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *CreateRepositoryGroupRequest) SetAccessToken(v string) *CreateRepositoryGroupRequest {
	s.AccessToken = &v
	return s
}

func (s *CreateRepositoryGroupRequest) SetAvatarUrl(v string) *CreateRepositoryGroupRequest {
	s.AvatarUrl = &v
	return s
}

func (s *CreateRepositoryGroupRequest) SetDescription(v string) *CreateRepositoryGroupRequest {
	s.Description = &v
	return s
}

func (s *CreateRepositoryGroupRequest) SetName(v string) *CreateRepositoryGroupRequest {
	s.Name = &v
	return s
}

func (s *CreateRepositoryGroupRequest) SetParentId(v int64) *CreateRepositoryGroupRequest {
	s.ParentId = &v
	return s
}

func (s *CreateRepositoryGroupRequest) SetPath(v string) *CreateRepositoryGroupRequest {
	s.Path = &v
	return s
}

func (s *CreateRepositoryGroupRequest) SetVisibilityLevel(v int32) *CreateRepositoryGroupRequest {
	s.VisibilityLevel = &v
	return s
}

func (s *CreateRepositoryGroupRequest) SetOrganizationId(v string) *CreateRepositoryGroupRequest {
	s.OrganizationId = &v
	return s
}

func (s *CreateRepositoryGroupRequest) Validate() error {
	return dara.Validate(s)
}
