// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWorkflowDAGVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *UpdateWorkflowDAGVersionRequest
	GetAppName() *string
	SetClusterId(v string) *UpdateWorkflowDAGVersionRequest
	GetClusterId() *string
	SetDagVersion(v string) *UpdateWorkflowDAGVersionRequest
	GetDagVersion() *string
	SetWorkflowId(v int64) *UpdateWorkflowDAGVersionRequest
	GetWorkflowId() *int64
}

type UpdateWorkflowDAGVersionRequest struct {
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
	// v1
	DagVersion *string `json:"DagVersion,omitempty" xml:"DagVersion,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	WorkflowId *int64 `json:"WorkflowId,omitempty" xml:"WorkflowId,omitempty"`
}

func (s UpdateWorkflowDAGVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkflowDAGVersionRequest) GoString() string {
	return s.String()
}

func (s *UpdateWorkflowDAGVersionRequest) GetAppName() *string {
	return s.AppName
}

func (s *UpdateWorkflowDAGVersionRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *UpdateWorkflowDAGVersionRequest) GetDagVersion() *string {
	return s.DagVersion
}

func (s *UpdateWorkflowDAGVersionRequest) GetWorkflowId() *int64 {
	return s.WorkflowId
}

func (s *UpdateWorkflowDAGVersionRequest) SetAppName(v string) *UpdateWorkflowDAGVersionRequest {
	s.AppName = &v
	return s
}

func (s *UpdateWorkflowDAGVersionRequest) SetClusterId(v string) *UpdateWorkflowDAGVersionRequest {
	s.ClusterId = &v
	return s
}

func (s *UpdateWorkflowDAGVersionRequest) SetDagVersion(v string) *UpdateWorkflowDAGVersionRequest {
	s.DagVersion = &v
	return s
}

func (s *UpdateWorkflowDAGVersionRequest) SetWorkflowId(v int64) *UpdateWorkflowDAGVersionRequest {
	s.WorkflowId = &v
	return s
}

func (s *UpdateWorkflowDAGVersionRequest) Validate() error {
	return dara.Validate(s)
}
