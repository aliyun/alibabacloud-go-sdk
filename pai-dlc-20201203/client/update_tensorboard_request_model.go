// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTensorboardRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessibility(v string) *UpdateTensorboardRequest
	GetAccessibility() *string
	SetMaxRunningTimeMinutes(v int64) *UpdateTensorboardRequest
	GetMaxRunningTimeMinutes() *int64
	SetPriority(v string) *UpdateTensorboardRequest
	GetPriority() *string
	SetWorkspaceId(v string) *UpdateTensorboardRequest
	GetWorkspaceId() *string
}

type UpdateTensorboardRequest struct {
	// Visibility of the task. Valid values:
	//
	// - PUBLIC: Visible to all users in this workspace.
	//
	// - PRIVATE: Visible only to you and administrators in this workspace.
	//
	// example:
	//
	// PRIVATE
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	// Maximum runtime. Unit: minutes.
	//
	// example:
	//
	// 300
	MaxRunningTimeMinutes *int64  `json:"MaxRunningTimeMinutes,omitempty" xml:"MaxRunningTimeMinutes,omitempty"`
	Priority              *string `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// Workspace ID. For more information about how to get a workspace ID, see [ListWorkspaces]().
	//
	// example:
	//
	// 380
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s UpdateTensorboardRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateTensorboardRequest) GoString() string {
	return s.String()
}

func (s *UpdateTensorboardRequest) GetAccessibility() *string {
	return s.Accessibility
}

func (s *UpdateTensorboardRequest) GetMaxRunningTimeMinutes() *int64 {
	return s.MaxRunningTimeMinutes
}

func (s *UpdateTensorboardRequest) GetPriority() *string {
	return s.Priority
}

func (s *UpdateTensorboardRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *UpdateTensorboardRequest) SetAccessibility(v string) *UpdateTensorboardRequest {
	s.Accessibility = &v
	return s
}

func (s *UpdateTensorboardRequest) SetMaxRunningTimeMinutes(v int64) *UpdateTensorboardRequest {
	s.MaxRunningTimeMinutes = &v
	return s
}

func (s *UpdateTensorboardRequest) SetPriority(v string) *UpdateTensorboardRequest {
	s.Priority = &v
	return s
}

func (s *UpdateTensorboardRequest) SetWorkspaceId(v string) *UpdateTensorboardRequest {
	s.WorkspaceId = &v
	return s
}

func (s *UpdateTensorboardRequest) Validate() error {
	return dara.Validate(s)
}
