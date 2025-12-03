// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessToken(v string) *UpdateGroupRequest
	GetAccessToken() *string
	SetAvatarUrl(v string) *UpdateGroupRequest
	GetAvatarUrl() *string
	SetDescription(v string) *UpdateGroupRequest
	GetDescription() *string
	SetName(v string) *UpdateGroupRequest
	GetName() *string
	SetPath(v string) *UpdateGroupRequest
	GetPath() *string
	SetPathWithNamespace(v string) *UpdateGroupRequest
	GetPathWithNamespace() *string
	SetVisibilityLevel(v int32) *UpdateGroupRequest
	GetVisibilityLevel() *int32
	SetOrganizationId(v string) *UpdateGroupRequest
	GetOrganizationId() *string
}

type UpdateGroupRequest struct {
	// example:
	//
	// f0b1e61dxxxxxxx975a93f9129d2513
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
	// codeup_group
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// codeup_group
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// org/group/subgroup/here
	PathWithNamespace *string `json:"pathWithNamespace,omitempty" xml:"pathWithNamespace,omitempty"`
	// example:
	//
	// 10
	VisibilityLevel *int32 `json:"visibilityLevel,omitempty" xml:"visibilityLevel,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1e9903d8b3f1xxxxxf9286ef5
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
}

func (s UpdateGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdateGroupRequest) GetAccessToken() *string {
	return s.AccessToken
}

func (s *UpdateGroupRequest) GetAvatarUrl() *string {
	return s.AvatarUrl
}

func (s *UpdateGroupRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateGroupRequest) GetName() *string {
	return s.Name
}

func (s *UpdateGroupRequest) GetPath() *string {
	return s.Path
}

func (s *UpdateGroupRequest) GetPathWithNamespace() *string {
	return s.PathWithNamespace
}

func (s *UpdateGroupRequest) GetVisibilityLevel() *int32 {
	return s.VisibilityLevel
}

func (s *UpdateGroupRequest) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *UpdateGroupRequest) SetAccessToken(v string) *UpdateGroupRequest {
	s.AccessToken = &v
	return s
}

func (s *UpdateGroupRequest) SetAvatarUrl(v string) *UpdateGroupRequest {
	s.AvatarUrl = &v
	return s
}

func (s *UpdateGroupRequest) SetDescription(v string) *UpdateGroupRequest {
	s.Description = &v
	return s
}

func (s *UpdateGroupRequest) SetName(v string) *UpdateGroupRequest {
	s.Name = &v
	return s
}

func (s *UpdateGroupRequest) SetPath(v string) *UpdateGroupRequest {
	s.Path = &v
	return s
}

func (s *UpdateGroupRequest) SetPathWithNamespace(v string) *UpdateGroupRequest {
	s.PathWithNamespace = &v
	return s
}

func (s *UpdateGroupRequest) SetVisibilityLevel(v int32) *UpdateGroupRequest {
	s.VisibilityLevel = &v
	return s
}

func (s *UpdateGroupRequest) SetOrganizationId(v string) *UpdateGroupRequest {
	s.OrganizationId = &v
	return s
}

func (s *UpdateGroupRequest) Validate() error {
	return dara.Validate(s)
}
