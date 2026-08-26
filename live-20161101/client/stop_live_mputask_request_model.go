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
	// The application ID. Only a single ID can be specified. The ID can contain uppercase and lowercase letters, digits, underscores, and hyphens (-), with a maximum of 64 characters. You can view your application IDs by navigating to **ApsaraVideo Live > Live+ > ApsaraVideo Real-time Communication > Application Management**.
	//
	// This parameter is required.
	//
	// example:
	//
	// yourAppId
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The task ID. Only a single ID can be specified. The ID can contain uppercase and lowercase letters, digits, underscores, and hyphens (-), with a maximum of 55 characters. This ID serves as the identifier for the bypass forwarding task and must be unique.
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
