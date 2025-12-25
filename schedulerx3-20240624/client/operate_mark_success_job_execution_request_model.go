// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateMarkSuccessJobExecutionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *OperateMarkSuccessJobExecutionRequest
	GetAppName() *string
	SetClusterId(v string) *OperateMarkSuccessJobExecutionRequest
	GetClusterId() *string
	SetJobExecutionId(v string) *OperateMarkSuccessJobExecutionRequest
	GetJobExecutionId() *string
}

type OperateMarkSuccessJobExecutionRequest struct {
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
	// example:
	//
	// 1310630367761285120
	JobExecutionId *string `json:"JobExecutionId,omitempty" xml:"JobExecutionId,omitempty"`
}

func (s OperateMarkSuccessJobExecutionRequest) String() string {
	return dara.Prettify(s)
}

func (s OperateMarkSuccessJobExecutionRequest) GoString() string {
	return s.String()
}

func (s *OperateMarkSuccessJobExecutionRequest) GetAppName() *string {
	return s.AppName
}

func (s *OperateMarkSuccessJobExecutionRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *OperateMarkSuccessJobExecutionRequest) GetJobExecutionId() *string {
	return s.JobExecutionId
}

func (s *OperateMarkSuccessJobExecutionRequest) SetAppName(v string) *OperateMarkSuccessJobExecutionRequest {
	s.AppName = &v
	return s
}

func (s *OperateMarkSuccessJobExecutionRequest) SetClusterId(v string) *OperateMarkSuccessJobExecutionRequest {
	s.ClusterId = &v
	return s
}

func (s *OperateMarkSuccessJobExecutionRequest) SetJobExecutionId(v string) *OperateMarkSuccessJobExecutionRequest {
	s.JobExecutionId = &v
	return s
}

func (s *OperateMarkSuccessJobExecutionRequest) Validate() error {
	return dara.Validate(s)
}
