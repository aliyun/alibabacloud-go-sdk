// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateStopWorkflowExecutionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *OperateStopWorkflowExecutionRequest
	GetAppName() *string
	SetClusterId(v string) *OperateStopWorkflowExecutionRequest
	GetClusterId() *string
	SetWorkflowExecutionId(v int64) *OperateStopWorkflowExecutionRequest
	GetWorkflowExecutionId() *int64
}

type OperateStopWorkflowExecutionRequest struct {
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
	// xxljob-a1804a3226d
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 100
	WorkflowExecutionId *int64 `json:"WorkflowExecutionId,omitempty" xml:"WorkflowExecutionId,omitempty"`
}

func (s OperateStopWorkflowExecutionRequest) String() string {
	return dara.Prettify(s)
}

func (s OperateStopWorkflowExecutionRequest) GoString() string {
	return s.String()
}

func (s *OperateStopWorkflowExecutionRequest) GetAppName() *string {
	return s.AppName
}

func (s *OperateStopWorkflowExecutionRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *OperateStopWorkflowExecutionRequest) GetWorkflowExecutionId() *int64 {
	return s.WorkflowExecutionId
}

func (s *OperateStopWorkflowExecutionRequest) SetAppName(v string) *OperateStopWorkflowExecutionRequest {
	s.AppName = &v
	return s
}

func (s *OperateStopWorkflowExecutionRequest) SetClusterId(v string) *OperateStopWorkflowExecutionRequest {
	s.ClusterId = &v
	return s
}

func (s *OperateStopWorkflowExecutionRequest) SetWorkflowExecutionId(v int64) *OperateStopWorkflowExecutionRequest {
	s.WorkflowExecutionId = &v
	return s
}

func (s *OperateStopWorkflowExecutionRequest) Validate() error {
	return dara.Validate(s)
}
