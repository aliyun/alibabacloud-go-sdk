// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateScheduledTaskShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescriptionShrink(v string) *UpdateScheduledTaskShrinkRequest
	GetDescriptionShrink() *string
	SetDigitalEmployeeNameShrink(v string) *UpdateScheduledTaskShrinkRequest
	GetDigitalEmployeeNameShrink() *string
	SetIsOpen(v bool) *UpdateScheduledTaskShrinkRequest
	GetIsOpen() *bool
	SetModel(v string) *UpdateScheduledTaskShrinkRequest
	GetModel() *string
	SetName(v string) *UpdateScheduledTaskShrinkRequest
	GetName() *string
	SetSegmentsShrink(v string) *UpdateScheduledTaskShrinkRequest
	GetSegmentsShrink() *string
	SetTaskDetailShrink(v string) *UpdateScheduledTaskShrinkRequest
	GetTaskDetailShrink() *string
	SetTaskId(v string) *UpdateScheduledTaskShrinkRequest
	GetTaskId() *string
	SetTenantId(v string) *UpdateScheduledTaskShrinkRequest
	GetTenantId() *string
	SetTriggerConfigShrink(v string) *UpdateScheduledTaskShrinkRequest
	GetTriggerConfigShrink() *string
	SetVisibility(v string) *UpdateScheduledTaskShrinkRequest
	GetVisibility() *string
	SetVisibleMemberUserIdsShrink(v string) *UpdateScheduledTaskShrinkRequest
	GetVisibleMemberUserIdsShrink() *string
}

type UpdateScheduledTaskShrinkRequest struct {
	// The description information.
	DescriptionShrink *string `json:"description,omitempty" xml:"description,omitempty"`
	// The list of digital human names.
	//
	// example:
	//
	// string_value
	DigitalEmployeeNameShrink *string `json:"digitalEmployeeName,omitempty" xml:"digitalEmployeeName,omitempty"`
	// Specifies whether the task is publicly accessible.
	//
	// example:
	//
	// true
	IsOpen *bool `json:"isOpen,omitempty" xml:"isOpen,omitempty"`
	// The execution model tier. If not specified, the model tier is not updated.
	//
	// example:
	//
	// quick
	Model *string `json:"model,omitempty" xml:"model,omitempty"`
	// The file name.
	//
	// example:
	//
	// SampleName.pdf
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The segments.
	SegmentsShrink *string `json:"segments,omitempty" xml:"segments,omitempty"`
	// The task details.
	TaskDetailShrink *string `json:"taskDetail,omitempty" xml:"taskDetail,omitempty"`
	// The task ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// exampleTaskId
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// The tenant ID. This is a common parameter. If not specified, the default tenant of the caller is used.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The trigger configuration. The configuration varies depending on the trigger type.
	TriggerConfigShrink *string `json:"triggerConfig,omitempty" xml:"triggerConfig,omitempty"`
	// The visibility scope for group tasks. Valid values: PRIVATE (visible only to the creator and group owner), COLLABORATIVE (visible to specified collaborators), and PUBLIC (visible to all group members). If not specified, the visibility is not updated. This parameter is ignored for personal tasks.
	//
	// example:
	//
	// COLLABORATIVE
	Visibility *string `json:"visibility,omitempty" xml:"visibility,omitempty"`
	// The full replacement list of collaborator member user IDs. This parameter takes effect only when visibility is set to COLLABORATIVE. The list is cleared when switching away from the COLLABORATIVE tier. A maximum of 1000 members are supported. If not specified, the member list is not updated. The task creator and group creator do not need to be included because they are covered by the authentication layer. This parameter is ignored for personal tasks.
	//
	// example:
	//
	// string_value
	VisibleMemberUserIdsShrink *string `json:"visibleMemberUserIds,omitempty" xml:"visibleMemberUserIds,omitempty"`
}

func (s UpdateScheduledTaskShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateScheduledTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateScheduledTaskShrinkRequest) GetDescriptionShrink() *string {
	return s.DescriptionShrink
}

func (s *UpdateScheduledTaskShrinkRequest) GetDigitalEmployeeNameShrink() *string {
	return s.DigitalEmployeeNameShrink
}

func (s *UpdateScheduledTaskShrinkRequest) GetIsOpen() *bool {
	return s.IsOpen
}

func (s *UpdateScheduledTaskShrinkRequest) GetModel() *string {
	return s.Model
}

func (s *UpdateScheduledTaskShrinkRequest) GetName() *string {
	return s.Name
}

func (s *UpdateScheduledTaskShrinkRequest) GetSegmentsShrink() *string {
	return s.SegmentsShrink
}

func (s *UpdateScheduledTaskShrinkRequest) GetTaskDetailShrink() *string {
	return s.TaskDetailShrink
}

func (s *UpdateScheduledTaskShrinkRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *UpdateScheduledTaskShrinkRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *UpdateScheduledTaskShrinkRequest) GetTriggerConfigShrink() *string {
	return s.TriggerConfigShrink
}

func (s *UpdateScheduledTaskShrinkRequest) GetVisibility() *string {
	return s.Visibility
}

func (s *UpdateScheduledTaskShrinkRequest) GetVisibleMemberUserIdsShrink() *string {
	return s.VisibleMemberUserIdsShrink
}

func (s *UpdateScheduledTaskShrinkRequest) SetDescriptionShrink(v string) *UpdateScheduledTaskShrinkRequest {
	s.DescriptionShrink = &v
	return s
}

func (s *UpdateScheduledTaskShrinkRequest) SetDigitalEmployeeNameShrink(v string) *UpdateScheduledTaskShrinkRequest {
	s.DigitalEmployeeNameShrink = &v
	return s
}

func (s *UpdateScheduledTaskShrinkRequest) SetIsOpen(v bool) *UpdateScheduledTaskShrinkRequest {
	s.IsOpen = &v
	return s
}

func (s *UpdateScheduledTaskShrinkRequest) SetModel(v string) *UpdateScheduledTaskShrinkRequest {
	s.Model = &v
	return s
}

func (s *UpdateScheduledTaskShrinkRequest) SetName(v string) *UpdateScheduledTaskShrinkRequest {
	s.Name = &v
	return s
}

func (s *UpdateScheduledTaskShrinkRequest) SetSegmentsShrink(v string) *UpdateScheduledTaskShrinkRequest {
	s.SegmentsShrink = &v
	return s
}

func (s *UpdateScheduledTaskShrinkRequest) SetTaskDetailShrink(v string) *UpdateScheduledTaskShrinkRequest {
	s.TaskDetailShrink = &v
	return s
}

func (s *UpdateScheduledTaskShrinkRequest) SetTaskId(v string) *UpdateScheduledTaskShrinkRequest {
	s.TaskId = &v
	return s
}

func (s *UpdateScheduledTaskShrinkRequest) SetTenantId(v string) *UpdateScheduledTaskShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *UpdateScheduledTaskShrinkRequest) SetTriggerConfigShrink(v string) *UpdateScheduledTaskShrinkRequest {
	s.TriggerConfigShrink = &v
	return s
}

func (s *UpdateScheduledTaskShrinkRequest) SetVisibility(v string) *UpdateScheduledTaskShrinkRequest {
	s.Visibility = &v
	return s
}

func (s *UpdateScheduledTaskShrinkRequest) SetVisibleMemberUserIdsShrink(v string) *UpdateScheduledTaskShrinkRequest {
	s.VisibleMemberUserIdsShrink = &v
	return s
}

func (s *UpdateScheduledTaskShrinkRequest) Validate() error {
	return dara.Validate(s)
}
