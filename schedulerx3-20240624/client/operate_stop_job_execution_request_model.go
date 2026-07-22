// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateStopJobExecutionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppGroupId(v int64) *OperateStopJobExecutionRequest
	GetAppGroupId() *int64
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
	// The subtask IDs.
	TaskList []*string `json:"TaskList,omitempty" xml:"TaskList,omitempty" type:"Repeated"`
}

func (s OperateStopJobExecutionRequest) String() string {
	return dara.Prettify(s)
}

func (s OperateStopJobExecutionRequest) GoString() string {
	return s.String()
}

func (s *OperateStopJobExecutionRequest) GetAppGroupId() *int64 {
	return s.AppGroupId
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

func (s *OperateStopJobExecutionRequest) SetAppGroupId(v int64) *OperateStopJobExecutionRequest {
	s.AppGroupId = &v
	return s
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
