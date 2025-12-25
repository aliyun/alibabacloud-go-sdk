// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateDisableWorkflowsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *OperateDisableWorkflowsShrinkRequest
	GetAppName() *string
	SetClusterId(v string) *OperateDisableWorkflowsShrinkRequest
	GetClusterId() *string
	SetWorkflowIdsShrink(v string) *OperateDisableWorkflowsShrinkRequest
	GetWorkflowIdsShrink() *string
}

type OperateDisableWorkflowsShrinkRequest struct {
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
	WorkflowIdsShrink *string `json:"WorkflowIds,omitempty" xml:"WorkflowIds,omitempty"`
}

func (s OperateDisableWorkflowsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s OperateDisableWorkflowsShrinkRequest) GoString() string {
	return s.String()
}

func (s *OperateDisableWorkflowsShrinkRequest) GetAppName() *string {
	return s.AppName
}

func (s *OperateDisableWorkflowsShrinkRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *OperateDisableWorkflowsShrinkRequest) GetWorkflowIdsShrink() *string {
	return s.WorkflowIdsShrink
}

func (s *OperateDisableWorkflowsShrinkRequest) SetAppName(v string) *OperateDisableWorkflowsShrinkRequest {
	s.AppName = &v
	return s
}

func (s *OperateDisableWorkflowsShrinkRequest) SetClusterId(v string) *OperateDisableWorkflowsShrinkRequest {
	s.ClusterId = &v
	return s
}

func (s *OperateDisableWorkflowsShrinkRequest) SetWorkflowIdsShrink(v string) *OperateDisableWorkflowsShrinkRequest {
	s.WorkflowIdsShrink = &v
	return s
}

func (s *OperateDisableWorkflowsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
