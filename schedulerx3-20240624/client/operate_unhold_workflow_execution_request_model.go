// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateUnholdWorkflowExecutionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *OperateUnholdWorkflowExecutionRequest
	GetAppName() *string
	SetClusterId(v string) *OperateUnholdWorkflowExecutionRequest
	GetClusterId() *string
	SetWorkflowExecutionId(v int64) *OperateUnholdWorkflowExecutionRequest
	GetWorkflowExecutionId() *int64
}

type OperateUnholdWorkflowExecutionRequest struct {
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
	// 100
	WorkflowExecutionId *int64 `json:"WorkflowExecutionId,omitempty" xml:"WorkflowExecutionId,omitempty"`
}

func (s OperateUnholdWorkflowExecutionRequest) String() string {
	return dara.Prettify(s)
}

func (s OperateUnholdWorkflowExecutionRequest) GoString() string {
	return s.String()
}

func (s *OperateUnholdWorkflowExecutionRequest) GetAppName() *string {
	return s.AppName
}

func (s *OperateUnholdWorkflowExecutionRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *OperateUnholdWorkflowExecutionRequest) GetWorkflowExecutionId() *int64 {
	return s.WorkflowExecutionId
}

func (s *OperateUnholdWorkflowExecutionRequest) SetAppName(v string) *OperateUnholdWorkflowExecutionRequest {
	s.AppName = &v
	return s
}

func (s *OperateUnholdWorkflowExecutionRequest) SetClusterId(v string) *OperateUnholdWorkflowExecutionRequest {
	s.ClusterId = &v
	return s
}

func (s *OperateUnholdWorkflowExecutionRequest) SetWorkflowExecutionId(v int64) *OperateUnholdWorkflowExecutionRequest {
	s.WorkflowExecutionId = &v
	return s
}

func (s *OperateUnholdWorkflowExecutionRequest) Validate() error {
	return dara.Validate(s)
}
