// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkflowDAGRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *GetWorkflowDAGRequest
	GetAppName() *string
	SetClusterId(v string) *GetWorkflowDAGRequest
	GetClusterId() *string
	SetWorkflowId(v int64) *GetWorkflowDAGRequest
	GetWorkflowId() *int64
}

type GetWorkflowDAGRequest struct {
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
	// example:
	//
	// 20
	WorkflowId *int64 `json:"WorkflowId,omitempty" xml:"WorkflowId,omitempty"`
}

func (s GetWorkflowDAGRequest) String() string {
	return dara.Prettify(s)
}

func (s GetWorkflowDAGRequest) GoString() string {
	return s.String()
}

func (s *GetWorkflowDAGRequest) GetAppName() *string {
	return s.AppName
}

func (s *GetWorkflowDAGRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetWorkflowDAGRequest) GetWorkflowId() *int64 {
	return s.WorkflowId
}

func (s *GetWorkflowDAGRequest) SetAppName(v string) *GetWorkflowDAGRequest {
	s.AppName = &v
	return s
}

func (s *GetWorkflowDAGRequest) SetClusterId(v string) *GetWorkflowDAGRequest {
	s.ClusterId = &v
	return s
}

func (s *GetWorkflowDAGRequest) SetWorkflowId(v int64) *GetWorkflowDAGRequest {
	s.WorkflowId = &v
	return s
}

func (s *GetWorkflowDAGRequest) Validate() error {
	return dara.Validate(s)
}
