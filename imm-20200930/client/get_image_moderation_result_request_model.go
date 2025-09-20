// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetImageModerationResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProjectName(v string) *GetImageModerationResultRequest
	GetProjectName() *string
	SetTaskId(v string) *GetImageModerationResultRequest
	GetTaskId() *string
	SetTaskType(v string) *GetImageModerationResultRequest
	GetTaskType() *string
}

type GetImageModerationResultRequest struct {
	// The name of the project.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The task ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ImageModeration-ff207203-3f93-4645-a041-7b8f0f******
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The type of the task.
	//
	// This parameter is required.
	//
	// example:
	//
	// ImageModeration
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s GetImageModerationResultRequest) String() string {
	return dara.Prettify(s)
}

func (s GetImageModerationResultRequest) GoString() string {
	return s.String()
}

func (s *GetImageModerationResultRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *GetImageModerationResultRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetImageModerationResultRequest) GetTaskType() *string {
	return s.TaskType
}

func (s *GetImageModerationResultRequest) SetProjectName(v string) *GetImageModerationResultRequest {
	s.ProjectName = &v
	return s
}

func (s *GetImageModerationResultRequest) SetTaskId(v string) *GetImageModerationResultRequest {
	s.TaskId = &v
	return s
}

func (s *GetImageModerationResultRequest) SetTaskType(v string) *GetImageModerationResultRequest {
	s.TaskType = &v
	return s
}

func (s *GetImageModerationResultRequest) Validate() error {
	return dara.Validate(s)
}
