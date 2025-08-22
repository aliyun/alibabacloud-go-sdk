// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateRetryJobExecutionShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *OperateRetryJobExecutionShrinkRequest
	GetAppName() *string
	SetClusterId(v string) *OperateRetryJobExecutionShrinkRequest
	GetClusterId() *string
	SetJobExecutionId(v string) *OperateRetryJobExecutionShrinkRequest
	GetJobExecutionId() *string
	SetTaskListShrink(v string) *OperateRetryJobExecutionShrinkRequest
	GetTaskListShrink() *string
}

type OperateRetryJobExecutionShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// test-app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxljob-b6ec1xxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1310630367761285120
	JobExecutionId *string `json:"JobExecutionId,omitempty" xml:"JobExecutionId,omitempty"`
	TaskListShrink *string `json:"TaskList,omitempty" xml:"TaskList,omitempty"`
}

func (s OperateRetryJobExecutionShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s OperateRetryJobExecutionShrinkRequest) GoString() string {
	return s.String()
}

func (s *OperateRetryJobExecutionShrinkRequest) GetAppName() *string {
	return s.AppName
}

func (s *OperateRetryJobExecutionShrinkRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *OperateRetryJobExecutionShrinkRequest) GetJobExecutionId() *string {
	return s.JobExecutionId
}

func (s *OperateRetryJobExecutionShrinkRequest) GetTaskListShrink() *string {
	return s.TaskListShrink
}

func (s *OperateRetryJobExecutionShrinkRequest) SetAppName(v string) *OperateRetryJobExecutionShrinkRequest {
	s.AppName = &v
	return s
}

func (s *OperateRetryJobExecutionShrinkRequest) SetClusterId(v string) *OperateRetryJobExecutionShrinkRequest {
	s.ClusterId = &v
	return s
}

func (s *OperateRetryJobExecutionShrinkRequest) SetJobExecutionId(v string) *OperateRetryJobExecutionShrinkRequest {
	s.JobExecutionId = &v
	return s
}

func (s *OperateRetryJobExecutionShrinkRequest) SetTaskListShrink(v string) *OperateRetryJobExecutionShrinkRequest {
	s.TaskListShrink = &v
	return s
}

func (s *OperateRetryJobExecutionShrinkRequest) Validate() error {
	return dara.Validate(s)
}
