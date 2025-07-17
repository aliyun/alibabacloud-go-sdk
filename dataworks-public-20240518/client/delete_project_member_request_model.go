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
	// The DataWorks workspace ID. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace page to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 534752
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The ID of the account used by the member in the workspace. You can log on to the [DataWorks console](https://dataworks.console.aliyun.com/product/ms_menu), choose More > Management Center in the left-side navigation pane, select the desired workspace on the Management Center page, and then click Go to Management Center. In the left-side navigation pane of the SettingCenter page, click Tenant Members and Roles. On the **Tenant Members and Roles*	- page, view the ID of the account used by the member in the workspace.
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
