// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateRetryJobExecutionShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppGroupId(v int64) *OperateRetryJobExecutionShrinkRequest
	GetAppGroupId() *int64
	SetAppName(v string) *OperateRetryJobExecutionShrinkRequest
	GetAppName() *string
	SetClusterId(v string) *OperateRetryJobExecutionShrinkRequest
	GetClusterId() *string
	SetJobExecutionId(v string) *OperateRetryJobExecutionShrinkRequest
	GetJobExecutionId() *string
	SetTaskListShrink(v string) *OperateRetryJobExecutionShrinkRequest
	GetTaskListShrink() *string
	SetTriggerChild(v bool) *OperateRetryJobExecutionShrinkRequest
	GetTriggerChild() *bool
}

type OperateRetryJobExecutionShrinkRequest struct {
	AppGroupId *int64 `json:"AppGroupId,omitempty" xml:"AppGroupId,omitempty"`
	// The application name.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// xxljob-b6ec1xxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The job execution ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1310630367761285120
	JobExecutionId *string `json:"JobExecutionId,omitempty" xml:"JobExecutionId,omitempty"`
	// The list of subtask execution IDs (for broadcast jobs).
	//
	// >To rerun a subtask of a broadcast job, set this field to the execution ID of the corresponding subtask.
	TaskListShrink *string `json:"TaskList,omitempty" xml:"TaskList,omitempty"`
	// Specifies whether to trigger downstream nodes.
	TriggerChild *bool `json:"TriggerChild,omitempty" xml:"TriggerChild,omitempty"`
}

func (s OperateRetryJobExecutionShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s OperateRetryJobExecutionShrinkRequest) GoString() string {
	return s.String()
}

func (s *OperateRetryJobExecutionShrinkRequest) GetAppGroupId() *int64 {
	return s.AppGroupId
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

func (s *OperateRetryJobExecutionShrinkRequest) GetTriggerChild() *bool {
	return s.TriggerChild
}

func (s *OperateRetryJobExecutionShrinkRequest) SetAppGroupId(v int64) *OperateRetryJobExecutionShrinkRequest {
	s.AppGroupId = &v
	return s
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

func (s *OperateRetryJobExecutionShrinkRequest) SetTriggerChild(v bool) *OperateRetryJobExecutionShrinkRequest {
	s.TriggerChild = &v
	return s
}

func (s *OperateRetryJobExecutionShrinkRequest) Validate() error {
	return dara.Validate(s)
}
