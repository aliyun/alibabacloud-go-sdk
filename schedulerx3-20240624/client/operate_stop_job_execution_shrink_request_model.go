// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateStopJobExecutionShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *OperateStopJobExecutionShrinkRequest
	GetAppName() *string
	SetClusterId(v string) *OperateStopJobExecutionShrinkRequest
	GetClusterId() *string
	SetJobExecutionId(v string) *OperateStopJobExecutionShrinkRequest
	GetJobExecutionId() *string
	SetTaskListShrink(v string) *OperateStopJobExecutionShrinkRequest
	GetTaskListShrink() *string
}

type OperateStopJobExecutionShrinkRequest struct {
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

func (s OperateStopJobExecutionShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s OperateStopJobExecutionShrinkRequest) GoString() string {
	return s.String()
}

func (s *OperateStopJobExecutionShrinkRequest) GetAppName() *string {
	return s.AppName
}

func (s *OperateStopJobExecutionShrinkRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *OperateStopJobExecutionShrinkRequest) GetJobExecutionId() *string {
	return s.JobExecutionId
}

func (s *OperateStopJobExecutionShrinkRequest) GetTaskListShrink() *string {
	return s.TaskListShrink
}

func (s *OperateStopJobExecutionShrinkRequest) SetAppName(v string) *OperateStopJobExecutionShrinkRequest {
	s.AppName = &v
	return s
}

func (s *OperateStopJobExecutionShrinkRequest) SetClusterId(v string) *OperateStopJobExecutionShrinkRequest {
	s.ClusterId = &v
	return s
}

func (s *OperateStopJobExecutionShrinkRequest) SetJobExecutionId(v string) *OperateStopJobExecutionShrinkRequest {
	s.JobExecutionId = &v
	return s
}

func (s *OperateStopJobExecutionShrinkRequest) SetTaskListShrink(v string) *OperateStopJobExecutionShrinkRequest {
	s.TaskListShrink = &v
	return s
}

func (s *OperateStopJobExecutionShrinkRequest) Validate() error {
	return dara.Validate(s)
}
