// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteComfyWorkflowRequest interface {
	dara.Model
	String() string
	GoString() string
	SetWorkflowId(v string) *DeleteComfyWorkflowRequest
	GetWorkflowId() *string
}

type DeleteComfyWorkflowRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// wf_adb32aed-ccdc-42ae-b4d4-a21181ac8a5f
	WorkflowId *string `json:"WorkflowId,omitempty" xml:"WorkflowId,omitempty"`
}

func (s DeleteComfyWorkflowRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteComfyWorkflowRequest) GoString() string {
	return s.String()
}

func (s *DeleteComfyWorkflowRequest) GetWorkflowId() *string {
	return s.WorkflowId
}

func (s *DeleteComfyWorkflowRequest) SetWorkflowId(v string) *DeleteComfyWorkflowRequest {
	s.WorkflowId = &v
	return s
}

func (s *DeleteComfyWorkflowRequest) Validate() error {
	return dara.Validate(s)
}
