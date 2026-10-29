// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyComfyWorkflowRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *ModifyComfyWorkflowRequest
	GetDescription() *string
	SetWorkflowId(v string) *ModifyComfyWorkflowRequest
	GetWorkflowId() *string
	SetWorkflowName(v string) *ModifyComfyWorkflowRequest
	GetWorkflowName() *string
}

type ModifyComfyWorkflowRequest struct {
	// The description of the workflow template.
	//
	// example:
	//
	// This is an image-to-video workflow
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The workflow ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// wf_3de1eb6e-1dfe-45aa-8f88-2269d0a30f53
	WorkflowId *string `json:"WorkflowId,omitempty" xml:"WorkflowId,omitempty"`
	// The name of the workflow template.
	//
	// example:
	//
	// Image-to-video workflow example
	WorkflowName *string `json:"WorkflowName,omitempty" xml:"WorkflowName,omitempty"`
}

func (s ModifyComfyWorkflowRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyComfyWorkflowRequest) GoString() string {
	return s.String()
}

func (s *ModifyComfyWorkflowRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyComfyWorkflowRequest) GetWorkflowId() *string {
	return s.WorkflowId
}

func (s *ModifyComfyWorkflowRequest) GetWorkflowName() *string {
	return s.WorkflowName
}

func (s *ModifyComfyWorkflowRequest) SetDescription(v string) *ModifyComfyWorkflowRequest {
	s.Description = &v
	return s
}

func (s *ModifyComfyWorkflowRequest) SetWorkflowId(v string) *ModifyComfyWorkflowRequest {
	s.WorkflowId = &v
	return s
}

func (s *ModifyComfyWorkflowRequest) SetWorkflowName(v string) *ModifyComfyWorkflowRequest {
	s.WorkflowName = &v
	return s
}

func (s *ModifyComfyWorkflowRequest) Validate() error {
	return dara.Validate(s)
}
