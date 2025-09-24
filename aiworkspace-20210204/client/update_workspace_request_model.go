// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWorkspaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateWorkspaceRequest
	GetDescription() *string
	SetDisplayName(v string) *UpdateWorkspaceRequest
	GetDisplayName() *string
}

type UpdateWorkspaceRequest struct {
	// The workspace description.
	//
	// example:
	//
	// Description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The display name of the workspace.
	//
	// 	- The name must be 3 to 23 characters in length, and can contain letters, underscores (_), and digits.
	//
	// 	- The name must start with a letter.
	//
	// 	- The name must be unique in the current region.
	//
	// example:
	//
	// workspace-example
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
}

func (s UpdateWorkspaceRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkspaceRequest) GoString() string {
	return s.String()
}

func (s *UpdateWorkspaceRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateWorkspaceRequest) GetDisplayName() *string {
	return s.DisplayName
}

func (s *UpdateWorkspaceRequest) SetDescription(v string) *UpdateWorkspaceRequest {
	s.Description = &v
	return s
}

func (s *UpdateWorkspaceRequest) SetDisplayName(v string) *UpdateWorkspaceRequest {
	s.DisplayName = &v
	return s
}

func (s *UpdateWorkspaceRequest) Validate() error {
	return dara.Validate(s)
}
