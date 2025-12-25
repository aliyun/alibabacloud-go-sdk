// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateRetryWorkflowExecutionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *OperateRetryWorkflowExecutionRequest
	GetAppName() *string
	SetClusterId(v string) *OperateRetryWorkflowExecutionRequest
	GetClusterId() *string
	SetOnlyFailed(v bool) *OperateRetryWorkflowExecutionRequest
	GetOnlyFailed() *bool
	SetWorkflowExecutionId(v int64) *OperateRetryWorkflowExecutionRequest
	GetWorkflowExecutionId() *int64
}

type OperateRetryWorkflowExecutionRequest struct {
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
	// true
	OnlyFailed *bool `json:"OnlyFailed,omitempty" xml:"OnlyFailed,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1450568762586578000
	WorkflowExecutionId *int64 `json:"WorkflowExecutionId,omitempty" xml:"WorkflowExecutionId,omitempty"`
}

func (s OperateRetryWorkflowExecutionRequest) String() string {
	return dara.Prettify(s)
}

func (s OperateRetryWorkflowExecutionRequest) GoString() string {
	return s.String()
}

func (s *OperateRetryWorkflowExecutionRequest) GetAppName() *string {
	return s.AppName
}

func (s *OperateRetryWorkflowExecutionRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *OperateRetryWorkflowExecutionRequest) GetOnlyFailed() *bool {
	return s.OnlyFailed
}

func (s *OperateRetryWorkflowExecutionRequest) GetWorkflowExecutionId() *int64 {
	return s.WorkflowExecutionId
}

func (s *OperateRetryWorkflowExecutionRequest) SetAppName(v string) *OperateRetryWorkflowExecutionRequest {
	s.AppName = &v
	return s
}

func (s *OperateRetryWorkflowExecutionRequest) SetClusterId(v string) *OperateRetryWorkflowExecutionRequest {
	s.ClusterId = &v
	return s
}

func (s *OperateRetryWorkflowExecutionRequest) SetOnlyFailed(v bool) *OperateRetryWorkflowExecutionRequest {
	s.OnlyFailed = &v
	return s
}

func (s *OperateRetryWorkflowExecutionRequest) SetWorkflowExecutionId(v int64) *OperateRetryWorkflowExecutionRequest {
	s.WorkflowExecutionId = &v
	return s
}

func (s *OperateRetryWorkflowExecutionRequest) Validate() error {
	return dara.Validate(s)
}
