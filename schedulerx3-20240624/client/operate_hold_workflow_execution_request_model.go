// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateHoldWorkflowExecutionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *OperateHoldWorkflowExecutionRequest
	GetAppName() *string
	SetClusterId(v string) *OperateHoldWorkflowExecutionRequest
	GetClusterId() *string
	SetWorkflowExecutionId(v int64) *OperateHoldWorkflowExecutionRequest
	GetWorkflowExecutionId() *int64
}

type OperateHoldWorkflowExecutionRequest struct {
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

func (s OperateHoldWorkflowExecutionRequest) String() string {
	return dara.Prettify(s)
}

func (s OperateHoldWorkflowExecutionRequest) GoString() string {
	return s.String()
}

func (s *OperateHoldWorkflowExecutionRequest) GetAppName() *string {
	return s.AppName
}

func (s *OperateHoldWorkflowExecutionRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *OperateHoldWorkflowExecutionRequest) GetWorkflowExecutionId() *int64 {
	return s.WorkflowExecutionId
}

func (s *OperateHoldWorkflowExecutionRequest) SetAppName(v string) *OperateHoldWorkflowExecutionRequest {
	s.AppName = &v
	return s
}

func (s *OperateHoldWorkflowExecutionRequest) SetClusterId(v string) *OperateHoldWorkflowExecutionRequest {
	s.ClusterId = &v
	return s
}

func (s *OperateHoldWorkflowExecutionRequest) SetWorkflowExecutionId(v int64) *OperateHoldWorkflowExecutionRequest {
	s.WorkflowExecutionId = &v
	return s
}

func (s *OperateHoldWorkflowExecutionRequest) Validate() error {
	return dara.Validate(s)
}
