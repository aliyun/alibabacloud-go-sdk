// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitAsyncTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *SubmitAsyncTaskRequest
	GetAgentKey() *string
	SetTaskCode(v string) *SubmitAsyncTaskRequest
	GetTaskCode() *string
	SetTaskExecuteTime(v string) *SubmitAsyncTaskRequest
	GetTaskExecuteTime() *string
	SetTaskName(v string) *SubmitAsyncTaskRequest
	GetTaskName() *string
	SetTaskParam(v string) *SubmitAsyncTaskRequest
	GetTaskParam() *string
}

type SubmitAsyncTaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 2daaa2e0c209xb26acb97009ea77bd4b_p_efm
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// MaterialDocumentUpload
	TaskCode *string `json:"TaskCode,omitempty" xml:"TaskCode,omitempty"`
	// example:
	//
	// 2023-10-14 14:30:00
	TaskExecuteTime *string `json:"TaskExecuteTime,omitempty" xml:"TaskExecuteTime,omitempty"`
	// example:
	//
	// 任务名称
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// example:
	//
	// 任务提交参数
	TaskParam *string `json:"TaskParam,omitempty" xml:"TaskParam,omitempty"`
}

func (s SubmitAsyncTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitAsyncTaskRequest) GoString() string {
	return s.String()
}

func (s *SubmitAsyncTaskRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *SubmitAsyncTaskRequest) GetTaskCode() *string {
	return s.TaskCode
}

func (s *SubmitAsyncTaskRequest) GetTaskExecuteTime() *string {
	return s.TaskExecuteTime
}

func (s *SubmitAsyncTaskRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *SubmitAsyncTaskRequest) GetTaskParam() *string {
	return s.TaskParam
}

func (s *SubmitAsyncTaskRequest) SetAgentKey(v string) *SubmitAsyncTaskRequest {
	s.AgentKey = &v
	return s
}

func (s *SubmitAsyncTaskRequest) SetTaskCode(v string) *SubmitAsyncTaskRequest {
	s.TaskCode = &v
	return s
}

func (s *SubmitAsyncTaskRequest) SetTaskExecuteTime(v string) *SubmitAsyncTaskRequest {
	s.TaskExecuteTime = &v
	return s
}

func (s *SubmitAsyncTaskRequest) SetTaskName(v string) *SubmitAsyncTaskRequest {
	s.TaskName = &v
	return s
}

func (s *SubmitAsyncTaskRequest) SetTaskParam(v string) *SubmitAsyncTaskRequest {
	s.TaskParam = &v
	return s
}

func (s *SubmitAsyncTaskRequest) Validate() error {
	return dara.Validate(s)
}
