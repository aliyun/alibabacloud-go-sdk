// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkflowExecutionDAGRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *GetWorkflowExecutionDAGRequest
	GetAppName() *string
	SetClusterId(v string) *GetWorkflowExecutionDAGRequest
	GetClusterId() *string
	SetWorkflowExecutionId(v int64) *GetWorkflowExecutionDAGRequest
	GetWorkflowExecutionId() *int64
}

type GetWorkflowExecutionDAGRequest struct {
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

func (s GetWorkflowExecutionDAGRequest) String() string {
	return dara.Prettify(s)
}

func (s GetWorkflowExecutionDAGRequest) GoString() string {
	return s.String()
}

func (s *GetWorkflowExecutionDAGRequest) GetAppName() *string {
	return s.AppName
}

func (s *GetWorkflowExecutionDAGRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetWorkflowExecutionDAGRequest) GetWorkflowExecutionId() *int64 {
	return s.WorkflowExecutionId
}

func (s *GetWorkflowExecutionDAGRequest) SetAppName(v string) *GetWorkflowExecutionDAGRequest {
	s.AppName = &v
	return s
}

func (s *GetWorkflowExecutionDAGRequest) SetClusterId(v string) *GetWorkflowExecutionDAGRequest {
	s.ClusterId = &v
	return s
}

func (s *GetWorkflowExecutionDAGRequest) SetWorkflowExecutionId(v int64) *GetWorkflowExecutionDAGRequest {
	s.WorkflowExecutionId = &v
	return s
}

func (s *GetWorkflowExecutionDAGRequest) Validate() error {
	return dara.Validate(s)
}
