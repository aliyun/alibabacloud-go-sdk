// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVideoAnalysisTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v string) *GetVideoAnalysisTaskRequest
	GetTaskId() *string
}

type GetVideoAnalysisTaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// a3d1c2ac-f086-4a21-9069-f5631542f5a2
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s GetVideoAnalysisTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s GetVideoAnalysisTaskRequest) GoString() string {
	return s.String()
}

func (s *GetVideoAnalysisTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetVideoAnalysisTaskRequest) SetTaskId(v string) *GetVideoAnalysisTaskRequest {
	s.TaskId = &v
	return s
}

func (s *GetVideoAnalysisTaskRequest) Validate() error {
	return dara.Validate(s)
}
