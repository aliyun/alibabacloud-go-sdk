// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateJobExecutionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *UpdateJobExecutionRequest
	GetAppName() *string
	SetClusterId(v string) *UpdateJobExecutionRequest
	GetClusterId() *string
	SetJobExecutionId(v string) *UpdateJobExecutionRequest
	GetJobExecutionId() *string
	SetScheduleTime(v int64) *UpdateJobExecutionRequest
	GetScheduleTime() *int64
}

type UpdateJobExecutionRequest struct {
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
	// example:
	//
	// 2023-10-01 12:00:00
	ScheduleTime *int64 `json:"ScheduleTime,omitempty" xml:"ScheduleTime,omitempty"`
}

func (s UpdateJobExecutionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateJobExecutionRequest) GoString() string {
	return s.String()
}

func (s *UpdateJobExecutionRequest) GetAppName() *string {
	return s.AppName
}

func (s *UpdateJobExecutionRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *UpdateJobExecutionRequest) GetJobExecutionId() *string {
	return s.JobExecutionId
}

func (s *UpdateJobExecutionRequest) GetScheduleTime() *int64 {
	return s.ScheduleTime
}

func (s *UpdateJobExecutionRequest) SetAppName(v string) *UpdateJobExecutionRequest {
	s.AppName = &v
	return s
}

func (s *UpdateJobExecutionRequest) SetClusterId(v string) *UpdateJobExecutionRequest {
	s.ClusterId = &v
	return s
}

func (s *UpdateJobExecutionRequest) SetJobExecutionId(v string) *UpdateJobExecutionRequest {
	s.JobExecutionId = &v
	return s
}

func (s *UpdateJobExecutionRequest) SetScheduleTime(v int64) *UpdateJobExecutionRequest {
	s.ScheduleTime = &v
	return s
}

func (s *UpdateJobExecutionRequest) Validate() error {
	return dara.Validate(s)
}
