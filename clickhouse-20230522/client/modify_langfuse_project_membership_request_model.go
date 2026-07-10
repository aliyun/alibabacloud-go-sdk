// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyLangfuseProjectMembershipRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *ModifyLangfuseProjectMembershipRequest
	GetDBInstanceId() *string
	SetEmail(v string) *ModifyLangfuseProjectMembershipRequest
	GetEmail() *string
	SetOrganizationId(v string) *ModifyLangfuseProjectMembershipRequest
	GetOrganizationId() *string
	SetProjectId(v string) *ModifyLangfuseProjectMembershipRequest
	GetProjectId() *string
	SetRegionId(v string) *ModifyLangfuseProjectMembershipRequest
	GetRegionId() *string
	SetRole(v string) *ModifyLangfuseProjectMembershipRequest
	GetRole() *string
}

type ModifyLangfuseProjectMembershipRequest struct {
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
	// The Langfuse project ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cmrbhzx930005jw2q****
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The role of the user in the project.
	//
	// This parameter is required.
	//
	// example:
	//
	// VIEWER
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
}

func (s ModifyLangfuseProjectMembershipRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyLangfuseProjectMembershipRequest) GoString() string {
	return s.String()
}

func (s *ModifyLangfuseProjectMembershipRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyLangfuseProjectMembershipRequest) GetEmail() *string {
	return s.Email
}

func (s *ModifyLangfuseProjectMembershipRequest) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *ModifyLangfuseProjectMembershipRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *ModifyLangfuseProjectMembershipRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyLangfuseProjectMembershipRequest) GetRole() *string {
	return s.Role
}

func (s *ModifyLangfuseProjectMembershipRequest) SetDBInstanceId(v string) *ModifyLangfuseProjectMembershipRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyLangfuseProjectMembershipRequest) SetEmail(v string) *ModifyLangfuseProjectMembershipRequest {
	s.Email = &v
	return s
}

func (s *ModifyLangfuseProjectMembershipRequest) SetOrganizationId(v string) *ModifyLangfuseProjectMembershipRequest {
	s.OrganizationId = &v
	return s
}

func (s *ModifyLangfuseProjectMembershipRequest) SetProjectId(v string) *ModifyLangfuseProjectMembershipRequest {
	s.ProjectId = &v
	return s
}

func (s *ModifyLangfuseProjectMembershipRequest) SetRegionId(v string) *ModifyLangfuseProjectMembershipRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyLangfuseProjectMembershipRequest) SetRole(v string) *ModifyLangfuseProjectMembershipRequest {
	s.Role = &v
	return s
}

func (s *ModifyLangfuseProjectMembershipRequest) Validate() error {
	return dara.Validate(s)
}
