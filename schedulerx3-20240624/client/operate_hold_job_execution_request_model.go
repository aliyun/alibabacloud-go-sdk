// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateHoldJobExecutionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *OperateHoldJobExecutionRequest
	GetAppName() *string
	SetClusterId(v string) *OperateHoldJobExecutionRequest
	GetClusterId() *string
	SetJobExecutionId(v string) *OperateHoldJobExecutionRequest
	GetJobExecutionId() *string
}

type OperateHoldJobExecutionRequest struct {
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

func (s OperateHoldJobExecutionRequest) String() string {
	return dara.Prettify(s)
}

func (s OperateHoldJobExecutionRequest) GoString() string {
	return s.String()
}

func (s *OperateHoldJobExecutionRequest) GetAppName() *string {
	return s.AppName
}

func (s *OperateHoldJobExecutionRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *OperateHoldJobExecutionRequest) GetJobExecutionId() *string {
	return s.JobExecutionId
}

func (s *OperateHoldJobExecutionRequest) SetAppName(v string) *OperateHoldJobExecutionRequest {
	s.AppName = &v
	return s
}

func (s *OperateHoldJobExecutionRequest) SetClusterId(v string) *OperateHoldJobExecutionRequest {
	s.ClusterId = &v
	return s
}

func (s *OperateHoldJobExecutionRequest) SetJobExecutionId(v string) *OperateHoldJobExecutionRequest {
	s.JobExecutionId = &v
	return s
}

func (s *OperateHoldJobExecutionRequest) Validate() error {
	return dara.Validate(s)
}
