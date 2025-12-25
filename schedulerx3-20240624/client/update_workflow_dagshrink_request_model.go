// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWorkflowDAGShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *UpdateWorkflowDAGShrinkRequest
	GetAppName() *string
	SetClusterId(v string) *UpdateWorkflowDAGShrinkRequest
	GetClusterId() *string
	SetDagShrink(v string) *UpdateWorkflowDAGShrinkRequest
	GetDagShrink() *string
	SetDagVersion(v string) *UpdateWorkflowDAGShrinkRequest
	GetDagVersion() *string
	SetWorkflowId(v int64) *UpdateWorkflowDAGShrinkRequest
	GetWorkflowId() *int64
}

type UpdateWorkflowDAGShrinkRequest struct {
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
	DagShrink *string `json:"Dag,omitempty" xml:"Dag,omitempty"`
	// example:
	//
	// 1137005
	DagVersion *string `json:"DagVersion,omitempty" xml:"DagVersion,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	WorkflowId *int64 `json:"WorkflowId,omitempty" xml:"WorkflowId,omitempty"`
}

func (s UpdateWorkflowDAGShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkflowDAGShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateWorkflowDAGShrinkRequest) GetAppName() *string {
	return s.AppName
}

func (s *UpdateWorkflowDAGShrinkRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *UpdateWorkflowDAGShrinkRequest) GetDagShrink() *string {
	return s.DagShrink
}

func (s *UpdateWorkflowDAGShrinkRequest) GetDagVersion() *string {
	return s.DagVersion
}

func (s *UpdateWorkflowDAGShrinkRequest) GetWorkflowId() *int64 {
	return s.WorkflowId
}

func (s *UpdateWorkflowDAGShrinkRequest) SetAppName(v string) *UpdateWorkflowDAGShrinkRequest {
	s.AppName = &v
	return s
}

func (s *UpdateWorkflowDAGShrinkRequest) SetClusterId(v string) *UpdateWorkflowDAGShrinkRequest {
	s.ClusterId = &v
	return s
}

func (s *UpdateWorkflowDAGShrinkRequest) SetDagShrink(v string) *UpdateWorkflowDAGShrinkRequest {
	s.DagShrink = &v
	return s
}

func (s *UpdateWorkflowDAGShrinkRequest) SetDagVersion(v string) *UpdateWorkflowDAGShrinkRequest {
	s.DagVersion = &v
	return s
}

func (s *UpdateWorkflowDAGShrinkRequest) SetWorkflowId(v int64) *UpdateWorkflowDAGShrinkRequest {
	s.WorkflowId = &v
	return s
}

func (s *UpdateWorkflowDAGShrinkRequest) Validate() error {
	return dara.Validate(s)
}
