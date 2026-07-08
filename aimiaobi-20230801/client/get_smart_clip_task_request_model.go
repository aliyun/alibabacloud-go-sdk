// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSmartClipTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v string) *GetSmartClipTaskRequest
	GetTaskId() *string
	SetWorkspaceId(v string) *GetSmartClipTaskRequest
	GetWorkspaceId() *string
}

type GetSmartClipTaskRequest struct {
	// Unique identifier of the task.
	//
	// > You do not need to specify TaskId. The system generates it automatically. If you use the same TaskId for multiple tasks, those tasks belong to the same conversation group.
	//
	// This parameter is required.
	//
	// example:
	//
	// task-03d46184ee7d8749
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// Unique identifier of your Alibaba Cloud Model Studio workspace. To get the workspace ID, see [Get the workspace ID](https://help.aliyun.com/document_detail/2782167.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// xxxx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetSmartClipTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSmartClipTaskRequest) GoString() string {
	return s.String()
}

func (s *GetSmartClipTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetSmartClipTaskRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetSmartClipTaskRequest) SetTaskId(v string) *GetSmartClipTaskRequest {
	s.TaskId = &v
	return s
}

func (s *GetSmartClipTaskRequest) SetWorkspaceId(v string) *GetSmartClipTaskRequest {
	s.WorkspaceId = &v
	return s
}

func (s *GetSmartClipTaskRequest) Validate() error {
	return dara.Validate(s)
}
