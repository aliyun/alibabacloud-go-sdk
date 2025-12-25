// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWorkflowsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *DeleteWorkflowsShrinkRequest
	GetAppName() *string
	SetClusterId(v string) *DeleteWorkflowsShrinkRequest
	GetClusterId() *string
	SetDeleteJobs(v bool) *DeleteWorkflowsShrinkRequest
	GetDeleteJobs() *bool
	SetWorkflowIdsShrink(v string) *DeleteWorkflowsShrinkRequest
	GetWorkflowIdsShrink() *string
}

type DeleteWorkflowsShrinkRequest struct {
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
	// false
	DeleteJobs *bool `json:"DeleteJobs,omitempty" xml:"DeleteJobs,omitempty"`
	// This parameter is required.
	WorkflowIdsShrink *string `json:"WorkflowIds,omitempty" xml:"WorkflowIds,omitempty"`
}

func (s DeleteWorkflowsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteWorkflowsShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteWorkflowsShrinkRequest) GetAppName() *string {
	return s.AppName
}

func (s *DeleteWorkflowsShrinkRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DeleteWorkflowsShrinkRequest) GetDeleteJobs() *bool {
	return s.DeleteJobs
}

func (s *DeleteWorkflowsShrinkRequest) GetWorkflowIdsShrink() *string {
	return s.WorkflowIdsShrink
}

func (s *DeleteWorkflowsShrinkRequest) SetAppName(v string) *DeleteWorkflowsShrinkRequest {
	s.AppName = &v
	return s
}

func (s *DeleteWorkflowsShrinkRequest) SetClusterId(v string) *DeleteWorkflowsShrinkRequest {
	s.ClusterId = &v
	return s
}

func (s *DeleteWorkflowsShrinkRequest) SetDeleteJobs(v bool) *DeleteWorkflowsShrinkRequest {
	s.DeleteJobs = &v
	return s
}

func (s *DeleteWorkflowsShrinkRequest) SetWorkflowIdsShrink(v string) *DeleteWorkflowsShrinkRequest {
	s.WorkflowIdsShrink = &v
	return s
}

func (s *DeleteWorkflowsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
