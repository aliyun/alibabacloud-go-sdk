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
	// The visibility of the task. Valid values:
	//
	// - PUBLIC: The task is visible to all users in the workspace.
	//
	// - PRIVATE: The task is visible only to you and administrators in the workspace.
	//
	// example:
	//
	// PRIVATE
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	// The maximum running duration in minutes.
	//
	// example:
	//
	// 300
	MaxRunningTimeMinutes *int64 `json:"MaxRunningTimeMinutes,omitempty" xml:"MaxRunningTimeMinutes,omitempty"`
	// The scheduling priority of the TensorBoard task. Valid values: 1 to 9. A value of 1 indicates the lowest priority, and a value of 9 indicates the highest priority.
	//
	// example:
	//
	// 1
	Priority *string `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The workspace ID. <props="china">For more information about how to obtain the workspace ID, see [ListWorkspaces](https://help.aliyun.com/document_detail/449124.html).
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
