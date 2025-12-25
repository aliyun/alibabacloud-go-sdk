// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateUnholdJobExecutionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *OperateUnholdJobExecutionRequest
	GetAppName() *string
	SetClusterId(v string) *OperateUnholdJobExecutionRequest
	GetClusterId() *string
	SetJobExecutionId(v string) *OperateUnholdJobExecutionRequest
	GetJobExecutionId() *string
}

type OperateUnholdJobExecutionRequest struct {
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

func (s OperateUnholdJobExecutionRequest) String() string {
	return dara.Prettify(s)
}

func (s OperateUnholdJobExecutionRequest) GoString() string {
	return s.String()
}

func (s *OperateUnholdJobExecutionRequest) GetAppName() *string {
	return s.AppName
}

func (s *OperateUnholdJobExecutionRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *OperateUnholdJobExecutionRequest) GetJobExecutionId() *string {
	return s.JobExecutionId
}

func (s *OperateUnholdJobExecutionRequest) SetAppName(v string) *OperateUnholdJobExecutionRequest {
	s.AppName = &v
	return s
}

func (s *OperateUnholdJobExecutionRequest) SetClusterId(v string) *OperateUnholdJobExecutionRequest {
	s.ClusterId = &v
	return s
}

func (s *OperateUnholdJobExecutionRequest) SetJobExecutionId(v string) *OperateUnholdJobExecutionRequest {
	s.JobExecutionId = &v
	return s
}

func (s *OperateUnholdJobExecutionRequest) Validate() error {
	return dara.Validate(s)
}
