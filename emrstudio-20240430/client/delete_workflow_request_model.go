// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWorkflowRequest interface {
	dara.Model
	String() string
	GoString() string
	SetWorkspaceId(v string) *DeleteWorkflowRequest
	GetWorkspaceId() *string
}

type DeleteWorkflowRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 12345
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s DeleteWorkflowRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteWorkflowRequest) GoString() string {
	return s.String()
}

func (s *DeleteWorkflowRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *DeleteWorkflowRequest) SetWorkspaceId(v string) *DeleteWorkflowRequest {
	s.WorkspaceId = &v
	return s
}

func (s *DeleteWorkflowRequest) Validate() error {
	return dara.Validate(s)
}
