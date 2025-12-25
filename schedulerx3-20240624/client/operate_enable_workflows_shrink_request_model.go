// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateEnableWorkflowsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *OperateEnableWorkflowsShrinkRequest
	GetAppName() *string
	SetClusterId(v string) *OperateEnableWorkflowsShrinkRequest
	GetClusterId() *string
	SetWorkflowIdsShrink(v string) *OperateEnableWorkflowsShrinkRequest
	GetWorkflowIdsShrink() *string
}

type OperateEnableWorkflowsShrinkRequest struct {
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
	WorkflowIdsShrink *string `json:"WorkflowIds,omitempty" xml:"WorkflowIds,omitempty"`
}

func (s OperateEnableWorkflowsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s OperateEnableWorkflowsShrinkRequest) GoString() string {
	return s.String()
}

func (s *OperateEnableWorkflowsShrinkRequest) GetAppName() *string {
	return s.AppName
}

func (s *OperateEnableWorkflowsShrinkRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *OperateEnableWorkflowsShrinkRequest) GetWorkflowIdsShrink() *string {
	return s.WorkflowIdsShrink
}

func (s *OperateEnableWorkflowsShrinkRequest) SetAppName(v string) *OperateEnableWorkflowsShrinkRequest {
	s.AppName = &v
	return s
}

func (s *OperateEnableWorkflowsShrinkRequest) SetClusterId(v string) *OperateEnableWorkflowsShrinkRequest {
	s.ClusterId = &v
	return s
}

func (s *OperateEnableWorkflowsShrinkRequest) SetWorkflowIdsShrink(v string) *OperateEnableWorkflowsShrinkRequest {
	s.WorkflowIdsShrink = &v
	return s
}

func (s *OperateEnableWorkflowsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
