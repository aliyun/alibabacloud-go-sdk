// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeMemberProjectRolesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProjectId(v int64) *RevokeMemberProjectRolesRequest
	GetProjectId() *int64
	SetRoleCodes(v []*string) *RevokeMemberProjectRolesRequest
	GetRoleCodes() []*string
	SetUserId(v string) *RevokeMemberProjectRolesRequest
	GetUserId() *string
}

type RevokeMemberProjectRolesRequest struct {
	// The ID of the DataWorks workspace. You can log on to the [DataWorks console](https://dataworks.console.aliyun.com/workspace/list) and go to the workspace settings page to obtain the workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 105149
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The list of workspace role codes. You can call [ListProjectRoles](https://help.aliyun.com/document_detail/2853930.html) to obtain the role codes.
	//
	// This parameter specifies the workspace roles to be revoked by this API call.
	//
	// This parameter is required.
	RoleCodes []*string `json:"RoleCodes,omitempty" xml:"RoleCodes,omitempty" type:"Repeated"`
	// The ID of the DataWorks account. You can log on to the [DataWorks console - Management Center](https://dataworks.console.aliyun.com/product/ms_menu), select the workspace that you want to manage, go to the Tenant Members and Roles page, and view the account ID of the member whose roles you want to revoke.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123422344899
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s RevokeMemberProjectRolesRequest) String() string {
	return dara.Prettify(s)
}

func (s RevokeMemberProjectRolesRequest) GoString() string {
	return s.String()
}

func (s *RevokeMemberProjectRolesRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *RevokeMemberProjectRolesRequest) GetRoleCodes() []*string {
	return s.RoleCodes
}

func (s *RevokeMemberProjectRolesRequest) GetUserId() *string {
	return s.UserId
}

func (s *RevokeMemberProjectRolesRequest) SetProjectId(v int64) *RevokeMemberProjectRolesRequest {
	s.ProjectId = &v
	return s
}

func (s *RevokeMemberProjectRolesRequest) SetRoleCodes(v []*string) *RevokeMemberProjectRolesRequest {
	s.RoleCodes = v
	return s
}

func (s *RevokeMemberProjectRolesRequest) SetUserId(v string) *RevokeMemberProjectRolesRequest {
	s.UserId = &v
	return s
}

func (s *RevokeMemberProjectRolesRequest) Validate() error {
	return dara.Validate(s)
}
