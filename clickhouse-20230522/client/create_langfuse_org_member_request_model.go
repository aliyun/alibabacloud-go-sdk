// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLangfuseOrgMemberRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *CreateLangfuseOrgMemberRequest
	GetDBInstanceId() *string
	SetEmail(v string) *CreateLangfuseOrgMemberRequest
	GetEmail() *string
	SetOrganizationId(v string) *CreateLangfuseOrgMemberRequest
	GetOrganizationId() *string
	SetRegionId(v string) *CreateLangfuseOrgMemberRequest
	GetRegionId() *string
	SetRole(v string) *CreateLangfuseOrgMemberRequest
	GetRole() *string
}

type CreateLangfuseOrgMemberRequest struct {
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
	// The Langfuse organization ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cmrbhzx930005jw2q****
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The role of the user in the organization.
	//
	// This parameter is required.
	//
	// example:
	//
	// VIEWER
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
}

func (s CreateLangfuseOrgMemberRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateLangfuseOrgMemberRequest) GoString() string {
	return s.String()
}

func (s *CreateLangfuseOrgMemberRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *CreateLangfuseOrgMemberRequest) GetEmail() *string {
	return s.Email
}

func (s *CreateLangfuseOrgMemberRequest) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *CreateLangfuseOrgMemberRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateLangfuseOrgMemberRequest) GetRole() *string {
	return s.Role
}

func (s *CreateLangfuseOrgMemberRequest) SetDBInstanceId(v string) *CreateLangfuseOrgMemberRequest {
	s.DBInstanceId = &v
	return s
}

func (s *CreateLangfuseOrgMemberRequest) SetEmail(v string) *CreateLangfuseOrgMemberRequest {
	s.Email = &v
	return s
}

func (s *CreateLangfuseOrgMemberRequest) SetOrganizationId(v string) *CreateLangfuseOrgMemberRequest {
	s.OrganizationId = &v
	return s
}

func (s *CreateLangfuseOrgMemberRequest) SetRegionId(v string) *CreateLangfuseOrgMemberRequest {
	s.RegionId = &v
	return s
}

func (s *CreateLangfuseOrgMemberRequest) SetRole(v string) *CreateLangfuseOrgMemberRequest {
	s.Role = &v
	return s
}

func (s *CreateLangfuseOrgMemberRequest) Validate() error {
	return dara.Validate(s)
}
