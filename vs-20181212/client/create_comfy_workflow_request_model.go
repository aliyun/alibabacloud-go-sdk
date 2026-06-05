// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateComfyWorkflowRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateComfyWorkflowRequest
	GetDescription() *string
	SetName(v string) *CreateComfyWorkflowRequest
	GetName() *string
	SetWorkflow(v string) *CreateComfyWorkflowRequest
	GetWorkflow() *string
}

type CreateComfyWorkflowRequest struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	Workflow *string `json:"Workflow,omitempty" xml:"Workflow,omitempty"`
}

func (s CreateComfyWorkflowRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateComfyWorkflowRequest) GoString() string {
	return s.String()
}

func (s *CreateComfyWorkflowRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateComfyWorkflowRequest) GetName() *string {
	return s.Name
}

func (s *CreateComfyWorkflowRequest) GetWorkflow() *string {
	return s.Workflow
}

func (s *CreateComfyWorkflowRequest) SetDescription(v string) *CreateComfyWorkflowRequest {
	s.Description = &v
	return s
}

func (s *CreateComfyWorkflowRequest) SetName(v string) *CreateComfyWorkflowRequest {
	s.Name = &v
	return s
}

func (s *CreateComfyWorkflowRequest) SetWorkflow(v string) *CreateComfyWorkflowRequest {
	s.Workflow = &v
	return s
}

func (s *CreateComfyWorkflowRequest) Validate() error {
	return dara.Validate(s)
}
