// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetJobExecutionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *GetJobExecutionRequest
	GetAppName() *string
	SetClusterId(v string) *GetJobExecutionRequest
	GetClusterId() *string
	SetJobExecutionId(v string) *GetJobExecutionRequest
	GetJobExecutionId() *string
	SetMseSessionId(v string) *GetJobExecutionRequest
	GetMseSessionId() *string
}

type GetJobExecutionRequest struct {
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
	MseSessionId   *string `json:"MseSessionId,omitempty" xml:"MseSessionId,omitempty"`
}

func (s GetJobExecutionRequest) String() string {
	return dara.Prettify(s)
}

func (s GetJobExecutionRequest) GoString() string {
	return s.String()
}

func (s *GetJobExecutionRequest) GetAppName() *string {
	return s.AppName
}

func (s *GetJobExecutionRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetJobExecutionRequest) GetJobExecutionId() *string {
	return s.JobExecutionId
}

func (s *GetJobExecutionRequest) GetMseSessionId() *string {
	return s.MseSessionId
}

func (s *GetJobExecutionRequest) SetAppName(v string) *GetJobExecutionRequest {
	s.AppName = &v
	return s
}

func (s *GetJobExecutionRequest) SetClusterId(v string) *GetJobExecutionRequest {
	s.ClusterId = &v
	return s
}

func (s *GetJobExecutionRequest) SetJobExecutionId(v string) *GetJobExecutionRequest {
	s.JobExecutionId = &v
	return s
}

func (s *GetJobExecutionRequest) SetMseSessionId(v string) *GetJobExecutionRequest {
	s.MseSessionId = &v
	return s
}

func (s *GetJobExecutionRequest) Validate() error {
	return dara.Validate(s)
}
