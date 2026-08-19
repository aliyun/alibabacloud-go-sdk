// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateProjectMemberRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProjectId(v int64) *CreateProjectMemberRequest
	GetProjectId() *int64
	SetRoleCodes(v []*string) *CreateProjectMemberRequest
	GetRoleCodes() []*string
	SetUserId(v string) *CreateProjectMemberRequest
	GetUserId() *string
}

type CreateProjectMemberRequest struct {
	// The ID of the DataWorks workspace. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace Management page to obtain the ID.
	//
	// This parameter specifies the DataWorks workspace for this API call operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 24054
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The list of workspace role codes. You can call [ListProjectRoles](https://help.aliyun.com/document_detail/2853930.html) to obtain the role codes.
	//
	// This parameter is used to grant workspace roles to the member when adding the member to the workspace.
	//
	// This parameter is required.
	RoleCodes []*string `json:"RoleCodes,omitempty" xml:"RoleCodes,omitempty" type:"Repeated"`
	// The ID of the DataWorks account. You can log on to the [DataWorks console - Management Center](https://dataworks.console.aliyun.com/product/ms_menu), select the workspace to which you want to add a member, go to the Management Center page, and then navigate to the **Tenant Members and Roles*	- page to view the account ID of the user you want to add to the workspace.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123422344899
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s CreateProjectMemberRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateProjectMemberRequest) GoString() string {
	return s.String()
}

func (s *CreateProjectMemberRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *CreateProjectMemberRequest) GetRoleCodes() []*string {
	return s.RoleCodes
}

func (s *CreateProjectMemberRequest) GetUserId() *string {
	return s.UserId
}

func (s *CreateProjectMemberRequest) SetProjectId(v int64) *CreateProjectMemberRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateProjectMemberRequest) SetRoleCodes(v []*string) *CreateProjectMemberRequest {
	s.RoleCodes = v
	return s
}

func (s *CreateProjectMemberRequest) SetUserId(v string) *CreateProjectMemberRequest {
	s.UserId = &v
	return s
}

func (s *CreateProjectMemberRequest) Validate() error {
	return dara.Validate(s)
}
