// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateMarkSuccessWorkflowExecutionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *OperateMarkSuccessWorkflowExecutionRequest
	GetAppName() *string
	SetClusterId(v string) *OperateMarkSuccessWorkflowExecutionRequest
	GetClusterId() *string
	SetWorkflowExecutionId(v int64) *OperateMarkSuccessWorkflowExecutionRequest
	GetWorkflowExecutionId() *int64
}

type OperateMarkSuccessWorkflowExecutionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// xxl-job-executor-sample
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxljob-d6a5243b6fa
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 100
	WorkflowExecutionId *int64 `json:"WorkflowExecutionId,omitempty" xml:"WorkflowExecutionId,omitempty"`
}

func (s OperateMarkSuccessWorkflowExecutionRequest) String() string {
	return dara.Prettify(s)
}

func (s OperateMarkSuccessWorkflowExecutionRequest) GoString() string {
	return s.String()
}

func (s *OperateMarkSuccessWorkflowExecutionRequest) GetAppName() *string {
	return s.AppName
}

func (s *OperateMarkSuccessWorkflowExecutionRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *OperateMarkSuccessWorkflowExecutionRequest) GetWorkflowExecutionId() *int64 {
	return s.WorkflowExecutionId
}

func (s *OperateMarkSuccessWorkflowExecutionRequest) SetAppName(v string) *OperateMarkSuccessWorkflowExecutionRequest {
	s.AppName = &v
	return s
}

func (s *OperateMarkSuccessWorkflowExecutionRequest) SetClusterId(v string) *OperateMarkSuccessWorkflowExecutionRequest {
	s.ClusterId = &v
	return s
}

func (s *OperateMarkSuccessWorkflowExecutionRequest) SetWorkflowExecutionId(v int64) *OperateMarkSuccessWorkflowExecutionRequest {
	s.WorkflowExecutionId = &v
	return s
}

func (s *OperateMarkSuccessWorkflowExecutionRequest) Validate() error {
	return dara.Validate(s)
}
