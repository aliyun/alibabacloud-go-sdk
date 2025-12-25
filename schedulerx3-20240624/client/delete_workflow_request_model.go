// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWorkflowRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *DeleteWorkflowRequest
	GetAppName() *string
	SetClusterId(v string) *DeleteWorkflowRequest
	GetClusterId() *string
	SetDeleteJobs(v bool) *DeleteWorkflowRequest
	GetDeleteJobs() *bool
	SetWorkflowId(v int64) *DeleteWorkflowRequest
	GetWorkflowId() *int64
}

type DeleteWorkflowRequest struct {
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
	// false
	DeleteJobs *bool `json:"DeleteJobs,omitempty" xml:"DeleteJobs,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	WorkflowId *int64 `json:"WorkflowId,omitempty" xml:"WorkflowId,omitempty"`
}

func (s DeleteWorkflowRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteWorkflowRequest) GoString() string {
	return s.String()
}

func (s *DeleteWorkflowRequest) GetAppName() *string {
	return s.AppName
}

func (s *DeleteWorkflowRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DeleteWorkflowRequest) GetDeleteJobs() *bool {
	return s.DeleteJobs
}

func (s *DeleteWorkflowRequest) GetWorkflowId() *int64 {
	return s.WorkflowId
}

func (s *DeleteWorkflowRequest) SetAppName(v string) *DeleteWorkflowRequest {
	s.AppName = &v
	return s
}

func (s *DeleteWorkflowRequest) SetClusterId(v string) *DeleteWorkflowRequest {
	s.ClusterId = &v
	return s
}

func (s *DeleteWorkflowRequest) SetDeleteJobs(v bool) *DeleteWorkflowRequest {
	s.DeleteJobs = &v
	return s
}

func (s *DeleteWorkflowRequest) SetWorkflowId(v int64) *DeleteWorkflowRequest {
	s.WorkflowId = &v
	return s
}

func (s *DeleteWorkflowRequest) Validate() error {
	return dara.Validate(s)
}
