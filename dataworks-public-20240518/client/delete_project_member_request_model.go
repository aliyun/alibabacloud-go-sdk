// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteProjectMemberRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProjectId(v int64) *DeleteProjectMemberRequest
	GetProjectId() *int64
	SetUserId(v string) *DeleteProjectMemberRequest
	GetUserId() *string
}

type DeleteProjectMemberRequest struct {
	// The ID of the DataWorks workspace. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the workspace configuration page to obtain the workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 534752
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The ID of the DataWorks account. You can log on to the [DataWorks console - Management Center](https://dataworks.console.aliyun.com/product/ms_menu), select the workspace from which you want to remove the member, go to the **Tenant Members and Roles*	- page, and view the account ID of the member to be removed from the workspace.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123422344899
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DeleteProjectMemberRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteProjectMemberRequest) GoString() string {
	return s.String()
}

func (s *DeleteProjectMemberRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *DeleteProjectMemberRequest) GetUserId() *string {
	return s.UserId
}

func (s *DeleteProjectMemberRequest) SetProjectId(v int64) *DeleteProjectMemberRequest {
	s.ProjectId = &v
	return s
}

func (s *DeleteProjectMemberRequest) SetUserId(v string) *DeleteProjectMemberRequest {
	s.UserId = &v
	return s
}

func (s *DeleteProjectMemberRequest) Validate() error {
	return dara.Validate(s)
}
