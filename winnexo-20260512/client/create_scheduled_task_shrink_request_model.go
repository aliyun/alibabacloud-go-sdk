// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateScheduledTaskShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCollaborationGroupId(v string) *CreateScheduledTaskShrinkRequest
	GetCollaborationGroupId() *string
	SetDescriptionShrink(v string) *CreateScheduledTaskShrinkRequest
	GetDescriptionShrink() *string
	SetDigitalEmployeeNameShrink(v string) *CreateScheduledTaskShrinkRequest
	GetDigitalEmployeeNameShrink() *string
	SetIsOpen(v bool) *CreateScheduledTaskShrinkRequest
	GetIsOpen() *bool
	SetModel(v string) *CreateScheduledTaskShrinkRequest
	GetModel() *string
	SetName(v string) *CreateScheduledTaskShrinkRequest
	GetName() *string
	SetOutputMode(v string) *CreateScheduledTaskShrinkRequest
	GetOutputMode() *string
	SetSegmentsShrink(v string) *CreateScheduledTaskShrinkRequest
	GetSegmentsShrink() *string
	SetTaskDetailShrink(v string) *CreateScheduledTaskShrinkRequest
	GetTaskDetailShrink() *string
	SetTenantId(v string) *CreateScheduledTaskShrinkRequest
	GetTenantId() *string
	SetTriggerConfigShrink(v string) *CreateScheduledTaskShrinkRequest
	GetTriggerConfigShrink() *string
	SetVisibility(v string) *CreateScheduledTaskShrinkRequest
	GetVisibility() *string
	SetVisibleMemberUserIdsShrink(v string) *CreateScheduledTaskShrinkRequest
	GetVisibleMemberUserIdsShrink() *string
}

type CreateScheduledTaskShrinkRequest struct {
	// The ID of the collaboration group to which the task belongs, such as cg_101. If this parameter is passed, a group space task is created, and the caller must be a valid group member. If this parameter is empty, a personal task is created.
	//
	// example:
	//
	// exampleCollaborationGroupId
	CollaborationGroupId *string `json:"collaborationGroupId,omitempty" xml:"collaborationGroupId,omitempty"`
	// The description of the to-do card type.
	DescriptionShrink *string `json:"description,omitempty" xml:"description,omitempty"`
	// The name of the current effective digital employee. This parameter is empty if it is not configured.
	//
	// example:
	//
	// ["pcitc-magent"]
	DigitalEmployeeNameShrink *string `json:"digitalEmployeeName,omitempty" xml:"digitalEmployeeName,omitempty"`
	// Specifies whether public access is allowed.
	//
	// example:
	//
	// true
	IsOpen *bool `json:"isOpen,omitempty" xml:"isOpen,omitempty"`
	// The large language model used by the assistant. The default value is empty, which indicates that DingTalk automatically selects the model.
	//
	// example:
	//
	// quick
	Model *string `json:"model,omitempty" xml:"model,omitempty"`
	// The name.
	//
	// This parameter is required.
	//
	// example:
	//
	// SampleName.pdf
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 输出会话模式，默认 NEW_CONVERSATION
	//
	// example:
	//
	// NEW_CONVERSATION
	OutputMode *string `json:"outputMode,omitempty" xml:"outputMode,omitempty"`
	// The site ID.
	SegmentsShrink *string `json:"segments,omitempty" xml:"segments,omitempty"`
	// The task details.
	TaskDetailShrink *string `json:"taskDetail,omitempty" xml:"taskDetail,omitempty"`
	// The effective tenant ID.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The trigger configuration of the scheduled task. The triggerMode parameter supports manual (manual trigger) and scheduled (scheduled trigger). If scheduled is used, the cron parameter must be specified. The timezone and language parameters are injected by the server, and the client does not need to pass them. The pushConfig parameter is the list of task push channels. If the list is empty or no channel is enabled, no push is performed.
	TriggerConfigShrink *string `json:"triggerConfig,omitempty" xml:"triggerConfig,omitempty"`
	// The visibility scope of the group task. Valid values: PRIVATE (visible only to the creator and group owner), COLLABORATIVE (visible to specified collaboration members), and PUBLIC (visible to all group members). If this parameter is not passed for a group task, the default value is PRIVATE. This field is ignored for personal tasks.
	//
	// example:
	//
	// PRIVATE
	Visibility *string `json:"visibility,omitempty" xml:"visibility,omitempty"`
	// The list of user IDs of collaboration members. This parameter takes effect only when visibility is set to COLLABORATIVE and is ignored for other values. A maximum of 1,000 user IDs can be specified. You do not need to pass the IDs of the task creator and group creator because they are covered by the authentication layer. This field is ignored for personal tasks.
	//
	// example:
	//
	// ["1002"]
	VisibleMemberUserIdsShrink *string `json:"visibleMemberUserIds,omitempty" xml:"visibleMemberUserIds,omitempty"`
}

func (s CreateScheduledTaskShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateScheduledTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateScheduledTaskShrinkRequest) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *CreateScheduledTaskShrinkRequest) GetDescriptionShrink() *string {
	return s.DescriptionShrink
}

func (s *CreateScheduledTaskShrinkRequest) GetDigitalEmployeeNameShrink() *string {
	return s.DigitalEmployeeNameShrink
}

func (s *CreateScheduledTaskShrinkRequest) GetIsOpen() *bool {
	return s.IsOpen
}

func (s *CreateScheduledTaskShrinkRequest) GetModel() *string {
	return s.Model
}

func (s *CreateScheduledTaskShrinkRequest) GetName() *string {
	return s.Name
}

func (s *CreateScheduledTaskShrinkRequest) GetOutputMode() *string {
	return s.OutputMode
}

func (s *CreateScheduledTaskShrinkRequest) GetSegmentsShrink() *string {
	return s.SegmentsShrink
}

func (s *CreateScheduledTaskShrinkRequest) GetTaskDetailShrink() *string {
	return s.TaskDetailShrink
}

func (s *CreateScheduledTaskShrinkRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreateScheduledTaskShrinkRequest) GetTriggerConfigShrink() *string {
	return s.TriggerConfigShrink
}

func (s *CreateScheduledTaskShrinkRequest) GetVisibility() *string {
	return s.Visibility
}

func (s *CreateScheduledTaskShrinkRequest) GetVisibleMemberUserIdsShrink() *string {
	return s.VisibleMemberUserIdsShrink
}

func (s *CreateScheduledTaskShrinkRequest) SetCollaborationGroupId(v string) *CreateScheduledTaskShrinkRequest {
	s.CollaborationGroupId = &v
	return s
}

func (s *CreateScheduledTaskShrinkRequest) SetDescriptionShrink(v string) *CreateScheduledTaskShrinkRequest {
	s.DescriptionShrink = &v
	return s
}

func (s *CreateScheduledTaskShrinkRequest) SetDigitalEmployeeNameShrink(v string) *CreateScheduledTaskShrinkRequest {
	s.DigitalEmployeeNameShrink = &v
	return s
}

func (s *CreateScheduledTaskShrinkRequest) SetIsOpen(v bool) *CreateScheduledTaskShrinkRequest {
	s.IsOpen = &v
	return s
}

func (s *CreateScheduledTaskShrinkRequest) SetModel(v string) *CreateScheduledTaskShrinkRequest {
	s.Model = &v
	return s
}

func (s *CreateScheduledTaskShrinkRequest) SetName(v string) *CreateScheduledTaskShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateScheduledTaskShrinkRequest) SetOutputMode(v string) *CreateScheduledTaskShrinkRequest {
	s.OutputMode = &v
	return s
}

func (s *CreateScheduledTaskShrinkRequest) SetSegmentsShrink(v string) *CreateScheduledTaskShrinkRequest {
	s.SegmentsShrink = &v
	return s
}

func (s *CreateScheduledTaskShrinkRequest) SetTaskDetailShrink(v string) *CreateScheduledTaskShrinkRequest {
	s.TaskDetailShrink = &v
	return s
}

func (s *CreateScheduledTaskShrinkRequest) SetTenantId(v string) *CreateScheduledTaskShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *CreateScheduledTaskShrinkRequest) SetTriggerConfigShrink(v string) *CreateScheduledTaskShrinkRequest {
	s.TriggerConfigShrink = &v
	return s
}

func (s *CreateScheduledTaskShrinkRequest) SetVisibility(v string) *CreateScheduledTaskShrinkRequest {
	s.Visibility = &v
	return s
}

func (s *CreateScheduledTaskShrinkRequest) SetVisibleMemberUserIdsShrink(v string) *CreateScheduledTaskShrinkRequest {
	s.VisibleMemberUserIdsShrink = &v
	return s
}

func (s *CreateScheduledTaskShrinkRequest) Validate() error {
	return dara.Validate(s)
}
