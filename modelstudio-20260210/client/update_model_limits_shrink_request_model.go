// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateModelLimitsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetWorkspaceId(v string) *UpdateModelLimitsShrinkRequest
	GetWorkspaceId() *string
	SetWorkspaceLimitsShrink(v string) *UpdateModelLimitsShrinkRequest
	GetWorkspaceLimitsShrink() *string
}

type UpdateModelLimitsShrinkRequest struct {
	// The workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ws-ac3ef438bec22dc5
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
	// The throttling values for the workspace.
	WorkspaceLimitsShrink *string `json:"workspaceLimits,omitempty" xml:"workspaceLimits,omitempty"`
}

func (s UpdateModelLimitsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateModelLimitsShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateModelLimitsShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *UpdateModelLimitsShrinkRequest) GetWorkspaceLimitsShrink() *string {
	return s.WorkspaceLimitsShrink
}

func (s *UpdateModelLimitsShrinkRequest) SetWorkspaceId(v string) *UpdateModelLimitsShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *UpdateModelLimitsShrinkRequest) SetWorkspaceLimitsShrink(v string) *UpdateModelLimitsShrinkRequest {
	s.WorkspaceLimitsShrink = &v
	return s
}

func (s *UpdateModelLimitsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
