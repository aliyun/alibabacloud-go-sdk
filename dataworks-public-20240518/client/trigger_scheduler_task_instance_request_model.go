// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTriggerSchedulerTaskInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnvType(v string) *TriggerSchedulerTaskInstanceRequest
	GetEnvType() *string
	SetTaskId(v int64) *TriggerSchedulerTaskInstanceRequest
	GetTaskId() *int64
	SetTriggerTime(v int64) *TriggerSchedulerTaskInstanceRequest
	GetTriggerTime() *int64
}

type TriggerSchedulerTaskInstanceRequest struct {
	// The environment of the workspace. Valid values:
	//
	// 	- Prod: production environment
	//
	// 	- Dev: development environment
	//
	// example:
	//
	// Prod
	EnvType *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	// The task ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The time defined by the HTTP Trigger node. This value is a UNIX timestamp.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1710239005403
	TriggerTime *int64 `json:"TriggerTime,omitempty" xml:"TriggerTime,omitempty"`
}

func (s TriggerSchedulerTaskInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s TriggerSchedulerTaskInstanceRequest) GoString() string {
	return s.String()
}

func (s *TriggerSchedulerTaskInstanceRequest) GetEnvType() *string {
	return s.EnvType
}

func (s *TriggerSchedulerTaskInstanceRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *TriggerSchedulerTaskInstanceRequest) GetTriggerTime() *int64 {
	return s.TriggerTime
}

func (s *TriggerSchedulerTaskInstanceRequest) SetEnvType(v string) *TriggerSchedulerTaskInstanceRequest {
	s.EnvType = &v
	return s
}

func (s *TriggerSchedulerTaskInstanceRequest) SetTaskId(v int64) *TriggerSchedulerTaskInstanceRequest {
	s.TaskId = &v
	return s
}

func (s *TriggerSchedulerTaskInstanceRequest) SetTriggerTime(v int64) *TriggerSchedulerTaskInstanceRequest {
	s.TriggerTime = &v
	return s
}

func (s *TriggerSchedulerTaskInstanceRequest) Validate() error {
	return dara.Validate(s)
}
