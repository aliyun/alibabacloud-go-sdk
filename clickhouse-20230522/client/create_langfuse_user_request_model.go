// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLangfuseUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *CreateLangfuseUserRequest
	GetDBInstanceId() *string
	SetEmail(v string) *CreateLangfuseUserRequest
	GetEmail() *string
	SetName(v string) *CreateLangfuseUserRequest
	GetName() *string
	SetOrgRole(v string) *CreateLangfuseUserRequest
	GetOrgRole() *string
	SetOrganizationId(v string) *CreateLangfuseUserRequest
	GetOrganizationId() *string
	SetPassword(v string) *CreateLangfuseUserRequest
	GetPassword() *string
	SetRegionId(v string) *CreateLangfuseUserRequest
	GetRegionId() *string
}

type CreateLangfuseUserRequest struct {
	// The Langfuse instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// lfs-****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The email address of the user.
	//
	// This parameter is required.
	//
	// example:
	//
	// john@company.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The username.
	//
	// This parameter is required.
	//
	// example:
	//
	// john
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The role of the user in the organization.
	//
	// example:
	//
	// VIEWER
	OrgRole *string `json:"OrgRole,omitempty" xml:"OrgRole,omitempty"`
	// The Langfuse organization ID.
	//
	// example:
	//
	// cmrbhzx930005jw2q****
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	// The password of the user account. The password must meet the following requirements:
	//
	// - Contains at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters.
	//
	// - The following special characters are supported: !@#$%^&*()_+-=
	//
	// - The password must be 8 to 32 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2F9e9@******
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateLangfuseUserRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateLangfuseUserRequest) GoString() string {
	return s.String()
}

func (s *CreateLangfuseUserRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *CreateLangfuseUserRequest) GetEmail() *string {
	return s.Email
}

func (s *CreateLangfuseUserRequest) GetName() *string {
	return s.Name
}

func (s *CreateLangfuseUserRequest) GetOrgRole() *string {
	return s.OrgRole
}

func (s *CreateLangfuseUserRequest) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *CreateLangfuseUserRequest) GetPassword() *string {
	return s.Password
}

func (s *CreateLangfuseUserRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateLangfuseUserRequest) SetDBInstanceId(v string) *CreateLangfuseUserRequest {
	s.DBInstanceId = &v
	return s
}

func (s *CreateLangfuseUserRequest) SetEmail(v string) *CreateLangfuseUserRequest {
	s.Email = &v
	return s
}

func (s *CreateLangfuseUserRequest) SetName(v string) *CreateLangfuseUserRequest {
	s.Name = &v
	return s
}

func (s *CreateLangfuseUserRequest) SetOrgRole(v string) *CreateLangfuseUserRequest {
	s.OrgRole = &v
	return s
}

func (s *CreateLangfuseUserRequest) SetOrganizationId(v string) *CreateLangfuseUserRequest {
	s.OrganizationId = &v
	return s
}

func (s *CreateLangfuseUserRequest) SetPassword(v string) *CreateLangfuseUserRequest {
	s.Password = &v
	return s
}

func (s *CreateLangfuseUserRequest) SetRegionId(v string) *CreateLangfuseUserRequest {
	s.RegionId = &v
	return s
}

func (s *CreateLangfuseUserRequest) Validate() error {
	return dara.Validate(s)
}
