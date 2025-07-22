// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopMPUTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *StopMPUTaskRequest
	GetAppId() *string
	SetOwnerId(v int64) *StopMPUTaskRequest
	GetOwnerId() *int64
	SetTaskId(v string) *StopMPUTaskRequest
	GetTaskId() *string
}

type StopMPUTaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// yourAppId
	AppId   *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// yourTaskId
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s StopMPUTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s StopMPUTaskRequest) GoString() string {
	return s.String()
}

func (s *StopMPUTaskRequest) GetAppId() *string {
	return s.AppId
}

func (s *StopMPUTaskRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *StopMPUTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *StopMPUTaskRequest) SetAppId(v string) *StopMPUTaskRequest {
	s.AppId = &v
	return s
}

func (s *StopMPUTaskRequest) SetOwnerId(v int64) *StopMPUTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *StopMPUTaskRequest) SetTaskId(v string) *StopMPUTaskRequest {
	s.TaskId = &v
	return s
}

func (s *StopMPUTaskRequest) Validate() error {
	return dara.Validate(s)
}
