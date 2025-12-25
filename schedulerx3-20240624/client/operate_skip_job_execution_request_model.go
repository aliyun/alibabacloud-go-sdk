// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateSkipJobExecutionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *OperateSkipJobExecutionRequest
	GetAppName() *string
	SetClusterId(v string) *OperateSkipJobExecutionRequest
	GetClusterId() *string
	SetJobExecutionId(v string) *OperateSkipJobExecutionRequest
	GetJobExecutionId() *string
}

type OperateSkipJobExecutionRequest struct {
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
	// xxljob-d6a5243b6fa
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// 1310630367761285120
	JobExecutionId *string `json:"JobExecutionId,omitempty" xml:"JobExecutionId,omitempty"`
}

func (s OperateSkipJobExecutionRequest) String() string {
	return dara.Prettify(s)
}

func (s OperateSkipJobExecutionRequest) GoString() string {
	return s.String()
}

func (s *OperateSkipJobExecutionRequest) GetAppName() *string {
	return s.AppName
}

func (s *OperateSkipJobExecutionRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *OperateSkipJobExecutionRequest) GetJobExecutionId() *string {
	return s.JobExecutionId
}

func (s *OperateSkipJobExecutionRequest) SetAppName(v string) *OperateSkipJobExecutionRequest {
	s.AppName = &v
	return s
}

func (s *OperateSkipJobExecutionRequest) SetClusterId(v string) *OperateSkipJobExecutionRequest {
	s.ClusterId = &v
	return s
}

func (s *OperateSkipJobExecutionRequest) SetJobExecutionId(v string) *OperateSkipJobExecutionRequest {
	s.JobExecutionId = &v
	return s
}

func (s *OperateSkipJobExecutionRequest) Validate() error {
	return dara.Validate(s)
}
