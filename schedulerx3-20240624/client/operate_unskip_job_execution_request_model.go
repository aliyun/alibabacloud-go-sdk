// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateUnskipJobExecutionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *OperateUnskipJobExecutionRequest
	GetAppName() *string
	SetClusterId(v string) *OperateUnskipJobExecutionRequest
	GetClusterId() *string
	SetJobExecutionId(v string) *OperateUnskipJobExecutionRequest
	GetJobExecutionId() *string
}

type OperateUnskipJobExecutionRequest struct {
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

func (s OperateUnskipJobExecutionRequest) String() string {
	return dara.Prettify(s)
}

func (s OperateUnskipJobExecutionRequest) GoString() string {
	return s.String()
}

func (s *OperateUnskipJobExecutionRequest) GetAppName() *string {
	return s.AppName
}

func (s *OperateUnskipJobExecutionRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *OperateUnskipJobExecutionRequest) GetJobExecutionId() *string {
	return s.JobExecutionId
}

func (s *OperateUnskipJobExecutionRequest) SetAppName(v string) *OperateUnskipJobExecutionRequest {
	s.AppName = &v
	return s
}

func (s *OperateUnskipJobExecutionRequest) SetClusterId(v string) *OperateUnskipJobExecutionRequest {
	s.ClusterId = &v
	return s
}

func (s *OperateUnskipJobExecutionRequest) SetJobExecutionId(v string) *OperateUnskipJobExecutionRequest {
	s.JobExecutionId = &v
	return s
}

func (s *OperateUnskipJobExecutionRequest) Validate() error {
	return dara.Validate(s)
}
