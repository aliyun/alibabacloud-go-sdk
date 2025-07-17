// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProjectMemberRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProjectId(v int64) *GetProjectMemberRequest
	GetProjectId() *int64
	SetUserId(v string) *GetProjectMemberRequest
	GetUserId() *string
}

type GetProjectMemberRequest struct {
	// The DataWorks workspace ID. You can log on to the [DataWorks console](https://dataworks.console.aliyun.com/workspace/list) and go to the Workspace page to query the ID.
	//
	// You must configure this parameter to specify the DataWorks workspace to which the API operation is applied.
	//
	// This parameter is required.
	//
	// example:
	//
	// 88757
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The ID of the account used by the member in the workspace. You can log on to the [DataWorks console](https://dataworks.console.aliyun.com/product/ms_menu), choose More > Management Center in the left-side navigation pane, select the desired workspace on the Management Center page, and then click Go to Management Center. In the left-side navigation pane of the SettingCenter page, click Tenant Members and Roles. On the Tenant Members and Roles page, view the ID of the account used by the member in the workspace.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123422344899
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetProjectMemberRequest) String() string {
	return dara.Prettify(s)
}

func (s GetProjectMemberRequest) GoString() string {
	return s.String()
}

func (s *GetProjectMemberRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetProjectMemberRequest) GetUserId() *string {
	return s.UserId
}

func (s *GetProjectMemberRequest) SetProjectId(v int64) *GetProjectMemberRequest {
	s.ProjectId = &v
	return s
}

func (s *GetProjectMemberRequest) SetUserId(v string) *GetProjectMemberRequest {
	s.UserId = &v
	return s
}

func (s *GetProjectMemberRequest) Validate() error {
	return dara.Validate(s)
}
