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
	// The ID of the DataWorks workspace. You can log on to the [DataWorks console](https://dataworks.console.aliyun.com/workspace/list) and go to the workspace management page to obtain the ID.
	//
	// This parameter specifies the DataWorks workspace used to invoke this API.
	//
	// This parameter is required.
	//
	// example:
	//
	// 88757
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The ID of the DataWorks account. You can log on to the [DataWorks console - Administration](https://dataworks.console.aliyun.com/product/ms_menu), select the workspace that you want to manage, go to the Administration page, and then go to the Tenant Members and Roles page to view the account ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123422****
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
