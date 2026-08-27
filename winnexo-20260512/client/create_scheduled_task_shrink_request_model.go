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
	// The ID of the collaboration group (such as cg_101). If specified, a group space task is created (the caller must be a valid group member). If empty, a personal task is created.
	//
	// example:
	//
	// exampleCollaborationGroupId
	CollaborationGroupId *string `json:"collaborationGroupId,omitempty" xml:"collaborationGroupId,omitempty"`
	// The description of the to-do card type.
	DescriptionShrink *string `json:"description,omitempty" xml:"description,omitempty"`
	// The name of the current effective digital employee. This parameter is empty if not configured.
	//
	// example:
	//
	// string_value
	DigitalEmployeeNameShrink *string `json:"digitalEmployeeName,omitempty" xml:"digitalEmployeeName,omitempty"`
	// Specifies whether public access is enabled.
	//
	// example:
	//
	// true
	IsOpen *bool `json:"isOpen,omitempty" xml:"isOpen,omitempty"`
	// The large model used by the assistant. An empty value indicates that DingTalk automatically selects the model.
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
	// The site ID.
	SegmentsShrink *string `json:"segments,omitempty" xml:"segments,omitempty"`
	// The task details.
	TaskDetailShrink *string `json:"taskDetail,omitempty" xml:"taskDetail,omitempty"`
	// The ID of the effective tenant.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The trigger configuration. The configuration varies depending on the trigger type. For the specific format, refer to the following data structures:
	//
	//   - OSS trigger: See [OSSTriggerConfig](https://help.aliyun.com/document_detail/415697.html).
	//
	//   - Simple Log Service trigger: See [LogTriggerConfig](https://help.aliyun.com/document_detail/415694.html).
	//
	//   - Time trigger: See [TimeTriggerConfig](https://help.aliyun.com/document_detail/415712.html).
	//
	//   - HTTP trigger: See [HTTPTriggerConfig](https://help.aliyun.com/document_detail/415685.html).
	//
	//   - Tablestore trigger: You only need to specify the complete **SourceArn*	- parameter. No additional configuration is required. Set the value to an empty object {}.
	//
	//   - CDN event trigger: See [CDNEventsTriggerConfig](https://help.aliyun.com/document_detail/415674.html).
	//
	//   - MNS topic trigger: See [MnsTopicTriggerConfig](https://help.aliyun.com/document_detail/415695.html).
	//
	//   - EventBridge trigger: See [EventBridgeTriggerConfig](https://help.aliyun.com/document_detail/2508622.html).
	TriggerConfigShrink *string `json:"triggerConfig,omitempty" xml:"triggerConfig,omitempty"`
	// The visibility scope of the group task. Valid values: PRIVATE (visible only to the creator and group owner), COLLABORATIVE (visible to specified collaborators), and PUBLIC (visible to all group members). Default value for group tasks: PRIVATE. This parameter is ignored for personal tasks.
	//
	// example:
	//
	// PRIVATE
	Visibility *string `json:"visibility,omitempty" xml:"visibility,omitempty"`
	// The list of collaborator user IDs. This parameter takes effect only when visibility is set to COLLABORATIVE. It is ignored for other visibility levels. A maximum of 1000 IDs are supported. The task creator and group creator do not need to be included (covered by the authentication layer). This parameter is ignored for personal tasks.
	//
	// example:
	//
	// string_value
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
