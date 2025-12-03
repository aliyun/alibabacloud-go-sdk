// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateProjectMemberRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRoleIdentifier(v string) *UpdateProjectMemberRequest
	GetRoleIdentifier() *string
	SetTargetIdentifier(v string) *UpdateProjectMemberRequest
	GetTargetIdentifier() *string
	SetTargetType(v string) *UpdateProjectMemberRequest
	GetTargetType() *string
	SetUserIdentifier(v string) *UpdateProjectMemberRequest
	GetUserIdentifier() *string
	SetUserType(v string) *UpdateProjectMemberRequest
	GetUserType() *string
}

type UpdateProjectMemberRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// project.admin
	RoleIdentifier *string `json:"roleIdentifier,omitempty" xml:"roleIdentifier,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 5e70xxxxxxcd000xxxxe96
	TargetIdentifier *string `json:"targetIdentifier,omitempty" xml:"targetIdentifier,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Space
	TargetType *string `json:"targetType,omitempty" xml:"targetType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 19xx7043xxxxxxx914
	UserIdentifier *string `json:"userIdentifier,omitempty" xml:"userIdentifier,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// user
	UserType *string `json:"userType,omitempty" xml:"userType,omitempty"`
}

func (s UpdateProjectMemberRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateProjectMemberRequest) GoString() string {
	return s.String()
}

func (s *UpdateProjectMemberRequest) GetRoleIdentifier() *string {
	return s.RoleIdentifier
}

func (s *UpdateProjectMemberRequest) GetTargetIdentifier() *string {
	return s.TargetIdentifier
}

func (s *UpdateProjectMemberRequest) GetTargetType() *string {
	return s.TargetType
}

func (s *UpdateProjectMemberRequest) GetUserIdentifier() *string {
	return s.UserIdentifier
}

func (s *UpdateProjectMemberRequest) GetUserType() *string {
	return s.UserType
}

func (s *UpdateProjectMemberRequest) SetRoleIdentifier(v string) *UpdateProjectMemberRequest {
	s.RoleIdentifier = &v
	return s
}

func (s *UpdateProjectMemberRequest) SetTargetIdentifier(v string) *UpdateProjectMemberRequest {
	s.TargetIdentifier = &v
	return s
}

func (s *UpdateProjectMemberRequest) SetTargetType(v string) *UpdateProjectMemberRequest {
	s.TargetType = &v
	return s
}

func (s *UpdateProjectMemberRequest) SetUserIdentifier(v string) *UpdateProjectMemberRequest {
	s.UserIdentifier = &v
	return s
}

func (s *UpdateProjectMemberRequest) SetUserType(v string) *UpdateProjectMemberRequest {
	s.UserType = &v
	return s
}

func (s *UpdateProjectMemberRequest) Validate() error {
	return dara.Validate(s)
}
