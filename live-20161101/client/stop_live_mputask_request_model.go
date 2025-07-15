// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopLiveMPUTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *StopLiveMPUTaskRequest
	GetAppId() *string
	SetTaskId(v string) *StopLiveMPUTaskRequest
	GetTaskId() *string
}

type StopLiveMPUTaskRequest struct {
	// The application ID. You can specify only one application ID. The ID can be up to 64 characters in length and can contain letters, digits, underscores (_), and hyphens (-).
	//
	// This parameter is required.
	//
	// example:
	//
	// yourAppId
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The task ID. You can specify only one task ID. The ID can be up to 55 characters in length and can contain letters, digits, underscores (_), and hyphens (-). The ID must be unique.
	//
	// This parameter is required.
	//
	// example:
	//
	// yourTaskId
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s StopLiveMPUTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s StopLiveMPUTaskRequest) GoString() string {
	return s.String()
}

func (s *StopLiveMPUTaskRequest) GetAppId() *string {
	return s.AppId
}

func (s *StopLiveMPUTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *StopLiveMPUTaskRequest) SetAppId(v string) *StopLiveMPUTaskRequest {
	s.AppId = &v
	return s
}

func (s *StopLiveMPUTaskRequest) SetTaskId(v string) *StopLiveMPUTaskRequest {
	s.TaskId = &v
	return s
}

func (s *StopLiveMPUTaskRequest) Validate() error {
	return dara.Validate(s)
}
