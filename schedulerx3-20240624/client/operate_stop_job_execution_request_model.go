// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateStopJobExecutionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *OperateStopJobExecutionRequest
	GetAppName() *string
	SetClusterId(v string) *OperateStopJobExecutionRequest
	GetClusterId() *string
	SetJobExecutionId(v string) *OperateStopJobExecutionRequest
	GetJobExecutionId() *string
	SetTaskList(v []*string) *OperateStopJobExecutionRequest
	GetTaskList() []*string
}

type OperateStopJobExecutionRequest struct {
	// The name of the Application.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The ID of the Cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// xxljob-b6ec1xxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The ID of the Job Execution.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1310630367761285120
	JobExecutionId *string `json:"JobExecutionId,omitempty" xml:"JobExecutionId,omitempty"`
	// A list of Task IDs.
	TaskList []*string `json:"TaskList,omitempty" xml:"TaskList,omitempty" type:"Repeated"`
}

func (s OperateStopJobExecutionRequest) String() string {
	return dara.Prettify(s)
}

func (s OperateStopJobExecutionRequest) GoString() string {
	return s.String()
}

func (s *OperateStopJobExecutionRequest) GetAppName() *string {
	return s.AppName
}

func (s *OperateStopJobExecutionRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *OperateStopJobExecutionRequest) GetJobExecutionId() *string {
	return s.JobExecutionId
}

func (s *OperateStopJobExecutionRequest) GetTaskList() []*string {
	return s.TaskList
}

func (s *OperateStopJobExecutionRequest) SetAppName(v string) *OperateStopJobExecutionRequest {
	s.AppName = &v
	return s
}

func (s *OperateStopJobExecutionRequest) SetClusterId(v string) *OperateStopJobExecutionRequest {
	s.ClusterId = &v
	return s
}

func (s *OperateStopJobExecutionRequest) SetJobExecutionId(v string) *OperateStopJobExecutionRequest {
	s.JobExecutionId = &v
	return s
}

func (s *OperateStopJobExecutionRequest) SetTaskList(v []*string) *OperateStopJobExecutionRequest {
	s.TaskList = v
	return s
}

func (s *OperateStopJobExecutionRequest) Validate() error {
	return dara.Validate(s)
}
