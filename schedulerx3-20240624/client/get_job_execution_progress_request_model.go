// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetJobExecutionProgressRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *GetJobExecutionProgressRequest
	GetAppName() *string
	SetClusterId(v string) *GetJobExecutionProgressRequest
	GetClusterId() *string
	SetJobExecutionId(v string) *GetJobExecutionProgressRequest
	GetJobExecutionId() *string
}

type GetJobExecutionProgressRequest struct {
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
}

func (s GetJobExecutionProgressRequest) String() string {
	return dara.Prettify(s)
}

func (s GetJobExecutionProgressRequest) GoString() string {
	return s.String()
}

func (s *GetJobExecutionProgressRequest) GetAppName() *string {
	return s.AppName
}

func (s *GetJobExecutionProgressRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetJobExecutionProgressRequest) GetJobExecutionId() *string {
	return s.JobExecutionId
}

func (s *GetJobExecutionProgressRequest) SetAppName(v string) *GetJobExecutionProgressRequest {
	s.AppName = &v
	return s
}

func (s *GetJobExecutionProgressRequest) SetClusterId(v string) *GetJobExecutionProgressRequest {
	s.ClusterId = &v
	return s
}

func (s *GetJobExecutionProgressRequest) SetJobExecutionId(v string) *GetJobExecutionProgressRequest {
	s.JobExecutionId = &v
	return s
}

func (s *GetJobExecutionProgressRequest) Validate() error {
	return dara.Validate(s)
}
