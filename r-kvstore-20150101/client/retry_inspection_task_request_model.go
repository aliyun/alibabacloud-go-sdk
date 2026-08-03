// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRetryInspectionTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *RetryInspectionTaskRequest
	GetInstanceId() *string
	SetSecurityToken(v string) *RetryInspectionTaskRequest
	GetSecurityToken() *string
	SetTaskId(v string) *RetryInspectionTaskRequest
	GetTaskId() *string
}

type RetryInspectionTaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// r-bp1zxszhcgatnx****
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	TaskId        *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s RetryInspectionTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s RetryInspectionTaskRequest) GoString() string {
	return s.String()
}

func (s *RetryInspectionTaskRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RetryInspectionTaskRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *RetryInspectionTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *RetryInspectionTaskRequest) SetInstanceId(v string) *RetryInspectionTaskRequest {
	s.InstanceId = &v
	return s
}

func (s *RetryInspectionTaskRequest) SetSecurityToken(v string) *RetryInspectionTaskRequest {
	s.SecurityToken = &v
	return s
}

func (s *RetryInspectionTaskRequest) SetTaskId(v string) *RetryInspectionTaskRequest {
	s.TaskId = &v
	return s
}

func (s *RetryInspectionTaskRequest) Validate() error {
	return dara.Validate(s)
}
