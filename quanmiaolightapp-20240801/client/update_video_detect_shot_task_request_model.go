// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateVideoDetectShotTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v string) *UpdateVideoDetectShotTaskRequest
	GetTaskId() *string
	SetTaskStatus(v string) *UpdateVideoDetectShotTaskRequest
	GetTaskStatus() *string
}

type UpdateVideoDetectShotTaskRequest struct {
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

func (s UpdateVideoDetectShotTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateVideoDetectShotTaskRequest) GoString() string {
	return s.String()
}

func (s *UpdateVideoDetectShotTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *UpdateVideoDetectShotTaskRequest) GetTaskStatus() *string {
	return s.TaskStatus
}

func (s *UpdateVideoDetectShotTaskRequest) SetTaskId(v string) *UpdateVideoDetectShotTaskRequest {
	s.TaskId = &v
	return s
}

func (s *UpdateVideoDetectShotTaskRequest) SetTaskStatus(v string) *UpdateVideoDetectShotTaskRequest {
	s.TaskStatus = &v
	return s
}

func (s *UpdateVideoDetectShotTaskRequest) Validate() error {
	return dara.Validate(s)
}
