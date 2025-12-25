// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateExecuteWorkflowRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *OperateExecuteWorkflowRequest
	GetAppName() *string
	SetClusterId(v string) *OperateExecuteWorkflowRequest
	GetClusterId() *string
	SetWorkflowId(v int64) *OperateExecuteWorkflowRequest
	GetWorkflowId() *int64
}

type OperateExecuteWorkflowRequest struct {
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
	// 20
	WorkflowId *int64 `json:"WorkflowId,omitempty" xml:"WorkflowId,omitempty"`
}

func (s OperateExecuteWorkflowRequest) String() string {
	return dara.Prettify(s)
}

func (s OperateExecuteWorkflowRequest) GoString() string {
	return s.String()
}

func (s *OperateExecuteWorkflowRequest) GetAppName() *string {
	return s.AppName
}

func (s *OperateExecuteWorkflowRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *OperateExecuteWorkflowRequest) GetWorkflowId() *int64 {
	return s.WorkflowId
}

func (s *OperateExecuteWorkflowRequest) SetAppName(v string) *OperateExecuteWorkflowRequest {
	s.AppName = &v
	return s
}

func (s *OperateExecuteWorkflowRequest) SetClusterId(v string) *OperateExecuteWorkflowRequest {
	s.ClusterId = &v
	return s
}

func (s *OperateExecuteWorkflowRequest) SetWorkflowId(v int64) *OperateExecuteWorkflowRequest {
	s.WorkflowId = &v
	return s
}

func (s *OperateExecuteWorkflowRequest) Validate() error {
	return dara.Validate(s)
}
