// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetJobExecutionThreadDumpRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *GetJobExecutionThreadDumpRequest
	GetAppName() *string
	SetClusterId(v string) *GetJobExecutionThreadDumpRequest
	GetClusterId() *string
	SetExecutorAddr(v string) *GetJobExecutionThreadDumpRequest
	GetExecutorAddr() *string
	SetJobExecutionId(v string) *GetJobExecutionThreadDumpRequest
	GetJobExecutionId() *string
}

type GetJobExecutionThreadDumpRequest struct {
	// example:
	//
	// xxl-job-executor-sample
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// xxljob-b6ec1xxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// http://192.168.0.215:9966/
	ExecutorAddr *string `json:"ExecutorAddr,omitempty" xml:"ExecutorAddr,omitempty"`
	// example:
	//
	// 1310630367761285120
	JobExecutionId *string `json:"JobExecutionId,omitempty" xml:"JobExecutionId,omitempty"`
}

func (s GetJobExecutionThreadDumpRequest) String() string {
	return dara.Prettify(s)
}

func (s GetJobExecutionThreadDumpRequest) GoString() string {
	return s.String()
}

func (s *GetJobExecutionThreadDumpRequest) GetAppName() *string {
	return s.AppName
}

func (s *GetJobExecutionThreadDumpRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetJobExecutionThreadDumpRequest) GetExecutorAddr() *string {
	return s.ExecutorAddr
}

func (s *GetJobExecutionThreadDumpRequest) GetJobExecutionId() *string {
	return s.JobExecutionId
}

func (s *GetJobExecutionThreadDumpRequest) SetAppName(v string) *GetJobExecutionThreadDumpRequest {
	s.AppName = &v
	return s
}

func (s *GetJobExecutionThreadDumpRequest) SetClusterId(v string) *GetJobExecutionThreadDumpRequest {
	s.ClusterId = &v
	return s
}

func (s *GetJobExecutionThreadDumpRequest) SetExecutorAddr(v string) *GetJobExecutionThreadDumpRequest {
	s.ExecutorAddr = &v
	return s
}

func (s *GetJobExecutionThreadDumpRequest) SetJobExecutionId(v string) *GetJobExecutionThreadDumpRequest {
	s.JobExecutionId = &v
	return s
}

func (s *GetJobExecutionThreadDumpRequest) Validate() error {
	return dara.Validate(s)
}
