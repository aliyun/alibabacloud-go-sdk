// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProjectRoleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetProjectRoleRequest
	GetCode() *string
	SetProjectId(v int64) *GetProjectRoleRequest
	GetProjectId() *int64
}

type GetProjectRoleRequest struct {
	// The code of the role in the DataWorks workspace. Valid values:
	//
	// 	- role_project_admin: workspace administrator
	//
	// 	- role_project_dev: developer
	//
	// 	- role_project_dg_admin: data governance administrator
	//
	// 	- role_project_guest: visitor
	//
	// 	- role_project_security: security administrator
	//
	// 	- role_project_deploy: deployer
	//
	// 	- role_project_owner: workspace owner
	//
	// 	- role_project_data_analyst: data analyst
	//
	// 	- role_project_pe: O\\&M engineer
	//
	// 	- role_project_erd: model designer
	//
	// This parameter is required.
	//
	// example:
	//
	// role_project_guest
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The DataWorks workspace ID. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace page to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10002
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s GetProjectRoleRequest) String() string {
	return dara.Prettify(s)
}

func (s GetProjectRoleRequest) GoString() string {
	return s.String()
}

func (s *GetProjectRoleRequest) GetCode() *string {
	return s.Code
}

func (s *GetProjectRoleRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetProjectRoleRequest) SetCode(v string) *GetProjectRoleRequest {
	s.Code = &v
	return s
}

func (s *GetProjectRoleRequest) SetProjectId(v int64) *GetProjectRoleRequest {
	s.ProjectId = &v
	return s
}

func (s *GetProjectRoleRequest) Validate() error {
	return dara.Validate(s)
}
