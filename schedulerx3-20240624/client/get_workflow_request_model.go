// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkflowRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *GetWorkflowRequest
	GetAppName() *string
	SetClusterId(v string) *GetWorkflowRequest
	GetClusterId() *string
	SetWorkflowId(v int64) *GetWorkflowRequest
	GetWorkflowId() *int64
}

type GetWorkflowRequest struct {
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
	// xxljob-b6ec1xxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20
	WorkflowId *int64 `json:"WorkflowId,omitempty" xml:"WorkflowId,omitempty"`
}

func (s GetWorkflowRequest) String() string {
	return dara.Prettify(s)
}

func (s GetWorkflowRequest) GoString() string {
	return s.String()
}

func (s *GetWorkflowRequest) GetAppName() *string {
	return s.AppName
}

func (s *GetWorkflowRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetWorkflowRequest) GetWorkflowId() *int64 {
	return s.WorkflowId
}

func (s *GetWorkflowRequest) SetAppName(v string) *GetWorkflowRequest {
	s.AppName = &v
	return s
}

func (s *GetWorkflowRequest) SetClusterId(v string) *GetWorkflowRequest {
	s.ClusterId = &v
	return s
}

func (s *GetWorkflowRequest) SetWorkflowId(v int64) *GetWorkflowRequest {
	s.WorkflowId = &v
	return s
}

func (s *GetWorkflowRequest) Validate() error {
	return dara.Validate(s)
}
