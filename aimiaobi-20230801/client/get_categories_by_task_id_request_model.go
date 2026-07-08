// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCategoriesByTaskIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v string) *GetCategoriesByTaskIdRequest
	GetTaskId() *string
	SetWorkspaceId(v string) *GetCategoriesByTaskIdRequest
	GetWorkspaceId() *string
}

type GetCategoriesByTaskIdRequest struct {
	// The unique ID of the task.
	//
	// > You do not need to specify this parameter. The system automatically generates a task ID. If you specify the same task ID for subsequent tasks, the tasks are considered part of the same conversation group.
	//
	// This parameter is required.
	//
	// example:
	//
	// xxx
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The unique ID of the Alibaba Cloud Model Studio workspace. For more information, see [Get a Workspace ID](https://help.aliyun.com/document_detail/2782167.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// xxxx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetCategoriesByTaskIdRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCategoriesByTaskIdRequest) GoString() string {
	return s.String()
}

func (s *GetCategoriesByTaskIdRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetCategoriesByTaskIdRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetCategoriesByTaskIdRequest) SetTaskId(v string) *GetCategoriesByTaskIdRequest {
	s.TaskId = &v
	return s
}

func (s *GetCategoriesByTaskIdRequest) SetWorkspaceId(v string) *GetCategoriesByTaskIdRequest {
	s.WorkspaceId = &v
	return s
}

func (s *GetCategoriesByTaskIdRequest) Validate() error {
	return dara.Validate(s)
}
