// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateVideoAnalysisTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v string) *UpdateVideoAnalysisTaskRequest
	GetTaskId() *string
	SetTaskStatus(v string) *UpdateVideoAnalysisTaskRequest
	GetTaskStatus() *string
}

type UpdateVideoAnalysisTaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// xxxx
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// CANCELED
	TaskStatus *string `json:"taskStatus,omitempty" xml:"taskStatus,omitempty"`
}

func (s UpdateVideoAnalysisTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateVideoAnalysisTaskRequest) GoString() string {
	return s.String()
}

func (s *UpdateVideoAnalysisTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *UpdateVideoAnalysisTaskRequest) GetTaskStatus() *string {
	return s.TaskStatus
}

func (s *UpdateVideoAnalysisTaskRequest) SetTaskId(v string) *UpdateVideoAnalysisTaskRequest {
	s.TaskId = &v
	return s
}

func (s *UpdateVideoAnalysisTaskRequest) SetTaskStatus(v string) *UpdateVideoAnalysisTaskRequest {
	s.TaskStatus = &v
	return s
}

func (s *UpdateVideoAnalysisTaskRequest) Validate() error {
	return dara.Validate(s)
}
