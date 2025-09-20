// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVideoModerationResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProjectName(v string) *GetVideoModerationResultRequest
	GetProjectName() *string
	SetTaskId(v string) *GetVideoModerationResultRequest
	GetTaskId() *string
	SetTaskType(v string) *GetVideoModerationResultRequest
	GetTaskType() *string
}

type GetVideoModerationResultRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// test-project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// VideoModeration-d0f0df1d-531d-4ab4-b353-e7f475******
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// VideoModeration
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s GetVideoModerationResultRequest) String() string {
	return dara.Prettify(s)
}

func (s GetVideoModerationResultRequest) GoString() string {
	return s.String()
}

func (s *GetVideoModerationResultRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *GetVideoModerationResultRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetVideoModerationResultRequest) GetTaskType() *string {
	return s.TaskType
}

func (s *GetVideoModerationResultRequest) SetProjectName(v string) *GetVideoModerationResultRequest {
	s.ProjectName = &v
	return s
}

func (s *GetVideoModerationResultRequest) SetTaskId(v string) *GetVideoModerationResultRequest {
	s.TaskId = &v
	return s
}

func (s *GetVideoModerationResultRequest) SetTaskType(v string) *GetVideoModerationResultRequest {
	s.TaskType = &v
	return s
}

func (s *GetVideoModerationResultRequest) Validate() error {
	return dara.Validate(s)
}
