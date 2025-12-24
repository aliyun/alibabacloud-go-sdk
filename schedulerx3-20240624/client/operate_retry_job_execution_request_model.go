// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateRetryJobExecutionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *OperateRetryJobExecutionRequest
	GetAppName() *string
	SetClusterId(v string) *OperateRetryJobExecutionRequest
	GetClusterId() *string
	SetJobExecutionId(v string) *OperateRetryJobExecutionRequest
	GetJobExecutionId() *string
	SetTaskList(v []*string) *OperateRetryJobExecutionRequest
	GetTaskList() []*string
	SetTriggerChild(v bool) *OperateRetryJobExecutionRequest
	GetTriggerChild() *bool
}

type OperateRetryJobExecutionRequest struct {
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
	JobExecutionId *string   `json:"JobExecutionId,omitempty" xml:"JobExecutionId,omitempty"`
	TaskList       []*string `json:"TaskList,omitempty" xml:"TaskList,omitempty" type:"Repeated"`
	TriggerChild   *bool     `json:"TriggerChild,omitempty" xml:"TriggerChild,omitempty"`
}

func (s OperateRetryJobExecutionRequest) String() string {
	return dara.Prettify(s)
}

func (s OperateRetryJobExecutionRequest) GoString() string {
	return s.String()
}

func (s *OperateRetryJobExecutionRequest) GetAppName() *string {
	return s.AppName
}

func (s *OperateRetryJobExecutionRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *OperateRetryJobExecutionRequest) GetJobExecutionId() *string {
	return s.JobExecutionId
}

func (s *OperateRetryJobExecutionRequest) GetTaskList() []*string {
	return s.TaskList
}

func (s *OperateRetryJobExecutionRequest) GetTriggerChild() *bool {
	return s.TriggerChild
}

func (s *OperateRetryJobExecutionRequest) SetAppName(v string) *OperateRetryJobExecutionRequest {
	s.AppName = &v
	return s
}

func (s *OperateRetryJobExecutionRequest) SetClusterId(v string) *OperateRetryJobExecutionRequest {
	s.ClusterId = &v
	return s
}

func (s *OperateRetryJobExecutionRequest) SetJobExecutionId(v string) *OperateRetryJobExecutionRequest {
	s.JobExecutionId = &v
	return s
}

func (s *OperateRetryJobExecutionRequest) SetTaskList(v []*string) *OperateRetryJobExecutionRequest {
	s.TaskList = v
	return s
}

func (s *OperateRetryJobExecutionRequest) SetTriggerChild(v bool) *OperateRetryJobExecutionRequest {
	s.TriggerChild = &v
	return s
}

func (s *OperateRetryJobExecutionRequest) Validate() error {
	return dara.Validate(s)
}
