// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateScheduledTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v []*UpdateScheduledTaskRequestDescription) *UpdateScheduledTaskRequest
	GetDescription() []*UpdateScheduledTaskRequestDescription
	SetDigitalEmployeeName(v []*string) *UpdateScheduledTaskRequest
	GetDigitalEmployeeName() []*string
	SetIsOpen(v bool) *UpdateScheduledTaskRequest
	GetIsOpen() *bool
	SetModel(v string) *UpdateScheduledTaskRequest
	GetModel() *string
	SetName(v string) *UpdateScheduledTaskRequest
	GetName() *string
	SetSegments(v []*UpdateScheduledTaskRequestSegments) *UpdateScheduledTaskRequest
	GetSegments() []*UpdateScheduledTaskRequestSegments
	SetTaskDetail(v *UpdateScheduledTaskRequestTaskDetail) *UpdateScheduledTaskRequest
	GetTaskDetail() *UpdateScheduledTaskRequestTaskDetail
	SetTaskId(v string) *UpdateScheduledTaskRequest
	GetTaskId() *string
	SetTenantId(v string) *UpdateScheduledTaskRequest
	GetTenantId() *string
	SetTriggerConfig(v *UpdateScheduledTaskRequestTriggerConfig) *UpdateScheduledTaskRequest
	GetTriggerConfig() *UpdateScheduledTaskRequestTriggerConfig
	SetVisibility(v string) *UpdateScheduledTaskRequest
	GetVisibility() *string
	SetVisibleMemberUserIds(v []*string) *UpdateScheduledTaskRequest
	GetVisibleMemberUserIds() []*string
}

type UpdateScheduledTaskRequest struct {
	// The description information.
	Description []*UpdateScheduledTaskRequestDescription `json:"description,omitempty" xml:"description,omitempty" type:"Repeated"`
	// The list of digital human names.
	//
	// example:
	//
	// string_value
	DigitalEmployeeName []*string `json:"digitalEmployeeName,omitempty" xml:"digitalEmployeeName,omitempty" type:"Repeated"`
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
	Segments []*UpdateScheduledTaskRequestSegments `json:"segments,omitempty" xml:"segments,omitempty" type:"Repeated"`
	// The task details.
	TaskDetail *UpdateScheduledTaskRequestTaskDetail `json:"taskDetail,omitempty" xml:"taskDetail,omitempty" type:"Struct"`
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
	TriggerConfig *UpdateScheduledTaskRequestTriggerConfig `json:"triggerConfig,omitempty" xml:"triggerConfig,omitempty" type:"Struct"`
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
	VisibleMemberUserIds []*string `json:"visibleMemberUserIds,omitempty" xml:"visibleMemberUserIds,omitempty" type:"Repeated"`
}

func (s UpdateScheduledTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateScheduledTaskRequest) GoString() string {
	return s.String()
}

func (s *UpdateScheduledTaskRequest) GetDescription() []*UpdateScheduledTaskRequestDescription {
	return s.Description
}

func (s *UpdateScheduledTaskRequest) GetDigitalEmployeeName() []*string {
	return s.DigitalEmployeeName
}

func (s *UpdateScheduledTaskRequest) GetIsOpen() *bool {
	return s.IsOpen
}

func (s *UpdateScheduledTaskRequest) GetModel() *string {
	return s.Model
}

func (s *UpdateScheduledTaskRequest) GetName() *string {
	return s.Name
}

func (s *UpdateScheduledTaskRequest) GetSegments() []*UpdateScheduledTaskRequestSegments {
	return s.Segments
}

func (s *UpdateScheduledTaskRequest) GetTaskDetail() *UpdateScheduledTaskRequestTaskDetail {
	return s.TaskDetail
}

func (s *UpdateScheduledTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *UpdateScheduledTaskRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *UpdateScheduledTaskRequest) GetTriggerConfig() *UpdateScheduledTaskRequestTriggerConfig {
	return s.TriggerConfig
}

func (s *UpdateScheduledTaskRequest) GetVisibility() *string {
	return s.Visibility
}

func (s *UpdateScheduledTaskRequest) GetVisibleMemberUserIds() []*string {
	return s.VisibleMemberUserIds
}

func (s *UpdateScheduledTaskRequest) SetDescription(v []*UpdateScheduledTaskRequestDescription) *UpdateScheduledTaskRequest {
	s.Description = v
	return s
}

func (s *UpdateScheduledTaskRequest) SetDigitalEmployeeName(v []*string) *UpdateScheduledTaskRequest {
	s.DigitalEmployeeName = v
	return s
}

func (s *UpdateScheduledTaskRequest) SetIsOpen(v bool) *UpdateScheduledTaskRequest {
	s.IsOpen = &v
	return s
}

func (s *UpdateScheduledTaskRequest) SetModel(v string) *UpdateScheduledTaskRequest {
	s.Model = &v
	return s
}

func (s *UpdateScheduledTaskRequest) SetName(v string) *UpdateScheduledTaskRequest {
	s.Name = &v
	return s
}

func (s *UpdateScheduledTaskRequest) SetSegments(v []*UpdateScheduledTaskRequestSegments) *UpdateScheduledTaskRequest {
	s.Segments = v
	return s
}

func (s *UpdateScheduledTaskRequest) SetTaskDetail(v *UpdateScheduledTaskRequestTaskDetail) *UpdateScheduledTaskRequest {
	s.TaskDetail = v
	return s
}

func (s *UpdateScheduledTaskRequest) SetTaskId(v string) *UpdateScheduledTaskRequest {
	s.TaskId = &v
	return s
}

func (s *UpdateScheduledTaskRequest) SetTenantId(v string) *UpdateScheduledTaskRequest {
	s.TenantId = &v
	return s
}

func (s *UpdateScheduledTaskRequest) SetTriggerConfig(v *UpdateScheduledTaskRequestTriggerConfig) *UpdateScheduledTaskRequest {
	s.TriggerConfig = v
	return s
}

func (s *UpdateScheduledTaskRequest) SetVisibility(v string) *UpdateScheduledTaskRequest {
	s.Visibility = &v
	return s
}

func (s *UpdateScheduledTaskRequest) SetVisibleMemberUserIds(v []*string) *UpdateScheduledTaskRequest {
	s.VisibleMemberUserIds = v
	return s
}

func (s *UpdateScheduledTaskRequest) Validate() error {
	if s.Description != nil {
		for _, item := range s.Description {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Segments != nil {
		for _, item := range s.Segments {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.TaskDetail != nil {
		if err := s.TaskDetail.Validate(); err != nil {
			return err
		}
	}
	if s.TriggerConfig != nil {
		if err := s.TriggerConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateScheduledTaskRequestDescription struct {
	// The text content. Required when type is set to text.
	//
	// example:
	//
	// Sample content
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// The feature switch. Optional when type is set to web_search.
	//
	// example:
	//
	// true
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The file name.
	//
	// example:
	//
	// SampleName.pdf
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The object ID. This parameter has a value when type is set to mention.
	//
	// example:
	//
	// exampleObjectId
	ObjectId *string `json:"objectId,omitempty" xml:"objectId,omitempty"`
	// The object type, such as customer. This parameter has a value when type is set to mention.
	//
	// example:
	//
	// string_value
	ObjectType *string `json:"objectType,omitempty" xml:"objectType,omitempty"`
	// The skill code. This parameter has a value when type is set to skill.
	//
	// example:
	//
	// string_value
	SkillCode *string `json:"skillCode,omitempty" xml:"skillCode,omitempty"`
	// The element type. Valid values: text, web_search, mention, and skill.
	//
	// This parameter is required.
	//
	// example:
	//
	// text
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s UpdateScheduledTaskRequestDescription) String() string {
	return dara.Prettify(s)
}

func (s UpdateScheduledTaskRequestDescription) GoString() string {
	return s.String()
}

func (s *UpdateScheduledTaskRequestDescription) GetContent() *string {
	return s.Content
}

func (s *UpdateScheduledTaskRequestDescription) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpdateScheduledTaskRequestDescription) GetName() *string {
	return s.Name
}

func (s *UpdateScheduledTaskRequestDescription) GetObjectId() *string {
	return s.ObjectId
}

func (s *UpdateScheduledTaskRequestDescription) GetObjectType() *string {
	return s.ObjectType
}

func (s *UpdateScheduledTaskRequestDescription) GetSkillCode() *string {
	return s.SkillCode
}

func (s *UpdateScheduledTaskRequestDescription) GetType() *string {
	return s.Type
}

func (s *UpdateScheduledTaskRequestDescription) SetContent(v string) *UpdateScheduledTaskRequestDescription {
	s.Content = &v
	return s
}

func (s *UpdateScheduledTaskRequestDescription) SetEnabled(v bool) *UpdateScheduledTaskRequestDescription {
	s.Enabled = &v
	return s
}

func (s *UpdateScheduledTaskRequestDescription) SetName(v string) *UpdateScheduledTaskRequestDescription {
	s.Name = &v
	return s
}

func (s *UpdateScheduledTaskRequestDescription) SetObjectId(v string) *UpdateScheduledTaskRequestDescription {
	s.ObjectId = &v
	return s
}

func (s *UpdateScheduledTaskRequestDescription) SetObjectType(v string) *UpdateScheduledTaskRequestDescription {
	s.ObjectType = &v
	return s
}

func (s *UpdateScheduledTaskRequestDescription) SetSkillCode(v string) *UpdateScheduledTaskRequestDescription {
	s.SkillCode = &v
	return s
}

func (s *UpdateScheduledTaskRequestDescription) SetType(v string) *UpdateScheduledTaskRequestDescription {
	s.Type = &v
	return s
}

func (s *UpdateScheduledTaskRequestDescription) Validate() error {
	return dara.Validate(s)
}

type UpdateScheduledTaskRequestSegments struct {
	// The text content. Required when type is set to text.
	//
	// example:
	//
	// Sample content
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// The feature switch. Optional when type is set to web_search.
	//
	// example:
	//
	// true
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The file name.
	//
	// example:
	//
	// SampleName.pdf
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The object ID. This parameter has a value when type is set to mention.
	//
	// example:
	//
	// exampleObjectId
	ObjectId *string `json:"objectId,omitempty" xml:"objectId,omitempty"`
	// The object type, such as customer. This parameter has a value when type is set to mention.
	//
	// example:
	//
	// string_value
	ObjectType *string `json:"objectType,omitempty" xml:"objectType,omitempty"`
	// The skill code. This parameter has a value when type is set to skill.
	//
	// example:
	//
	// string_value
	SkillCode *string `json:"skillCode,omitempty" xml:"skillCode,omitempty"`
	// The element type. Valid values: text, web_search, mention, and skill.
	//
	// This parameter is required.
	//
	// example:
	//
	// text
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s UpdateScheduledTaskRequestSegments) String() string {
	return dara.Prettify(s)
}

func (s UpdateScheduledTaskRequestSegments) GoString() string {
	return s.String()
}

func (s *UpdateScheduledTaskRequestSegments) GetContent() *string {
	return s.Content
}

func (s *UpdateScheduledTaskRequestSegments) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpdateScheduledTaskRequestSegments) GetName() *string {
	return s.Name
}

func (s *UpdateScheduledTaskRequestSegments) GetObjectId() *string {
	return s.ObjectId
}

func (s *UpdateScheduledTaskRequestSegments) GetObjectType() *string {
	return s.ObjectType
}

func (s *UpdateScheduledTaskRequestSegments) GetSkillCode() *string {
	return s.SkillCode
}

func (s *UpdateScheduledTaskRequestSegments) GetType() *string {
	return s.Type
}

func (s *UpdateScheduledTaskRequestSegments) SetContent(v string) *UpdateScheduledTaskRequestSegments {
	s.Content = &v
	return s
}

func (s *UpdateScheduledTaskRequestSegments) SetEnabled(v bool) *UpdateScheduledTaskRequestSegments {
	s.Enabled = &v
	return s
}

func (s *UpdateScheduledTaskRequestSegments) SetName(v string) *UpdateScheduledTaskRequestSegments {
	s.Name = &v
	return s
}

func (s *UpdateScheduledTaskRequestSegments) SetObjectId(v string) *UpdateScheduledTaskRequestSegments {
	s.ObjectId = &v
	return s
}

func (s *UpdateScheduledTaskRequestSegments) SetObjectType(v string) *UpdateScheduledTaskRequestSegments {
	s.ObjectType = &v
	return s
}

func (s *UpdateScheduledTaskRequestSegments) SetSkillCode(v string) *UpdateScheduledTaskRequestSegments {
	s.SkillCode = &v
	return s
}

func (s *UpdateScheduledTaskRequestSegments) SetType(v string) *UpdateScheduledTaskRequestSegments {
	s.Type = &v
	return s
}

func (s *UpdateScheduledTaskRequestSegments) Validate() error {
	return dara.Validate(s)
}

type UpdateScheduledTaskRequestTaskDetail struct {
	// The related objects.
	RelatedObjects []*UpdateScheduledTaskRequestTaskDetailRelatedObjects `json:"relatedObjects,omitempty" xml:"relatedObjects,omitempty" type:"Repeated"`
	// The related semantics.
	RelatedSemantics []*UpdateScheduledTaskRequestTaskDetailRelatedSemantics `json:"relatedSemantics,omitempty" xml:"relatedSemantics,omitempty" type:"Repeated"`
	// The related skills.
	RelatedSkills []*UpdateScheduledTaskRequestTaskDetailRelatedSkills `json:"relatedSkills,omitempty" xml:"relatedSkills,omitempty" type:"Repeated"`
	// The task understanding description polished by the LLM.
	//
	// example:
	//
	// string_value
	TaskUnderstand *string `json:"taskUnderstand,omitempty" xml:"taskUnderstand,omitempty"`
}

func (s UpdateScheduledTaskRequestTaskDetail) String() string {
	return dara.Prettify(s)
}

func (s UpdateScheduledTaskRequestTaskDetail) GoString() string {
	return s.String()
}

func (s *UpdateScheduledTaskRequestTaskDetail) GetRelatedObjects() []*UpdateScheduledTaskRequestTaskDetailRelatedObjects {
	return s.RelatedObjects
}

func (s *UpdateScheduledTaskRequestTaskDetail) GetRelatedSemantics() []*UpdateScheduledTaskRequestTaskDetailRelatedSemantics {
	return s.RelatedSemantics
}

func (s *UpdateScheduledTaskRequestTaskDetail) GetRelatedSkills() []*UpdateScheduledTaskRequestTaskDetailRelatedSkills {
	return s.RelatedSkills
}

func (s *UpdateScheduledTaskRequestTaskDetail) GetTaskUnderstand() *string {
	return s.TaskUnderstand
}

func (s *UpdateScheduledTaskRequestTaskDetail) SetRelatedObjects(v []*UpdateScheduledTaskRequestTaskDetailRelatedObjects) *UpdateScheduledTaskRequestTaskDetail {
	s.RelatedObjects = v
	return s
}

func (s *UpdateScheduledTaskRequestTaskDetail) SetRelatedSemantics(v []*UpdateScheduledTaskRequestTaskDetailRelatedSemantics) *UpdateScheduledTaskRequestTaskDetail {
	s.RelatedSemantics = v
	return s
}

func (s *UpdateScheduledTaskRequestTaskDetail) SetRelatedSkills(v []*UpdateScheduledTaskRequestTaskDetailRelatedSkills) *UpdateScheduledTaskRequestTaskDetail {
	s.RelatedSkills = v
	return s
}

func (s *UpdateScheduledTaskRequestTaskDetail) SetTaskUnderstand(v string) *UpdateScheduledTaskRequestTaskDetail {
	s.TaskUnderstand = &v
	return s
}

func (s *UpdateScheduledTaskRequestTaskDetail) Validate() error {
	if s.RelatedObjects != nil {
		for _, item := range s.RelatedObjects {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.RelatedSemantics != nil {
		for _, item := range s.RelatedSemantics {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.RelatedSkills != nil {
		for _, item := range s.RelatedSkills {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateScheduledTaskRequestTaskDetailRelatedObjects struct {
	// The mention type, such as objects.
	//
	// example:
	//
	// string_value
	MentionType *string `json:"mentionType,omitempty" xml:"mentionType,omitempty"`
	// The file name.
	//
	// example:
	//
	// SampleName.pdf
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The object ID. This parameter has a value when an object is mentioned using @.
	//
	// example:
	//
	// exampleObjectId
	ObjectId *string `json:"objectId,omitempty" xml:"objectId,omitempty"`
	// The object type, such as customer or company.
	//
	// example:
	//
	// string_value
	ObjectType *string `json:"objectType,omitempty" xml:"objectType,omitempty"`
}

func (s UpdateScheduledTaskRequestTaskDetailRelatedObjects) String() string {
	return dara.Prettify(s)
}

func (s UpdateScheduledTaskRequestTaskDetailRelatedObjects) GoString() string {
	return s.String()
}

func (s *UpdateScheduledTaskRequestTaskDetailRelatedObjects) GetMentionType() *string {
	return s.MentionType
}

func (s *UpdateScheduledTaskRequestTaskDetailRelatedObjects) GetName() *string {
	return s.Name
}

func (s *UpdateScheduledTaskRequestTaskDetailRelatedObjects) GetObjectId() *string {
	return s.ObjectId
}

func (s *UpdateScheduledTaskRequestTaskDetailRelatedObjects) GetObjectType() *string {
	return s.ObjectType
}

func (s *UpdateScheduledTaskRequestTaskDetailRelatedObjects) SetMentionType(v string) *UpdateScheduledTaskRequestTaskDetailRelatedObjects {
	s.MentionType = &v
	return s
}

func (s *UpdateScheduledTaskRequestTaskDetailRelatedObjects) SetName(v string) *UpdateScheduledTaskRequestTaskDetailRelatedObjects {
	s.Name = &v
	return s
}

func (s *UpdateScheduledTaskRequestTaskDetailRelatedObjects) SetObjectId(v string) *UpdateScheduledTaskRequestTaskDetailRelatedObjects {
	s.ObjectId = &v
	return s
}

func (s *UpdateScheduledTaskRequestTaskDetailRelatedObjects) SetObjectType(v string) *UpdateScheduledTaskRequestTaskDetailRelatedObjects {
	s.ObjectType = &v
	return s
}

func (s *UpdateScheduledTaskRequestTaskDetailRelatedObjects) Validate() error {
	return dara.Validate(s)
}

type UpdateScheduledTaskRequestTaskDetailRelatedSemantics struct {
	// The semantic attributes (JSON string) used for filtering during semantic retrieval.
	//
	// example:
	//
	// {"level": "VIP"}
	Attributes *string `json:"attributes,omitempty" xml:"attributes,omitempty"`
	// The semantic entity name, such as customer or opportunity.
	//
	// example:
	//
	// customer
	Entity *string `json:"entity,omitempty" xml:"entity,omitempty"`
}

func (s UpdateScheduledTaskRequestTaskDetailRelatedSemantics) String() string {
	return dara.Prettify(s)
}

func (s UpdateScheduledTaskRequestTaskDetailRelatedSemantics) GoString() string {
	return s.String()
}

func (s *UpdateScheduledTaskRequestTaskDetailRelatedSemantics) GetAttributes() *string {
	return s.Attributes
}

func (s *UpdateScheduledTaskRequestTaskDetailRelatedSemantics) GetEntity() *string {
	return s.Entity
}

func (s *UpdateScheduledTaskRequestTaskDetailRelatedSemantics) SetAttributes(v string) *UpdateScheduledTaskRequestTaskDetailRelatedSemantics {
	s.Attributes = &v
	return s
}

func (s *UpdateScheduledTaskRequestTaskDetailRelatedSemantics) SetEntity(v string) *UpdateScheduledTaskRequestTaskDetailRelatedSemantics {
	s.Entity = &v
	return s
}

func (s *UpdateScheduledTaskRequestTaskDetailRelatedSemantics) Validate() error {
	return dara.Validate(s)
}

type UpdateScheduledTaskRequestTaskDetailRelatedSkills struct {
	// The display name of the skill.
	//
	// example:
	//
	// string_value
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// The file name.
	//
	// example:
	//
	// SampleName.pdf
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The skill code.
	//
	// example:
	//
	// string_value
	SkillCode *string `json:"skillCode,omitempty" xml:"skillCode,omitempty"`
	// sourceIds
	//
	// example:
	//
	// string_value
	SourceIds []*string `json:"sourceIds,omitempty" xml:"sourceIds,omitempty" type:"Repeated"`
}

func (s UpdateScheduledTaskRequestTaskDetailRelatedSkills) String() string {
	return dara.Prettify(s)
}

func (s UpdateScheduledTaskRequestTaskDetailRelatedSkills) GoString() string {
	return s.String()
}

func (s *UpdateScheduledTaskRequestTaskDetailRelatedSkills) GetDisplayName() *string {
	return s.DisplayName
}

func (s *UpdateScheduledTaskRequestTaskDetailRelatedSkills) GetName() *string {
	return s.Name
}

func (s *UpdateScheduledTaskRequestTaskDetailRelatedSkills) GetSkillCode() *string {
	return s.SkillCode
}

func (s *UpdateScheduledTaskRequestTaskDetailRelatedSkills) GetSourceIds() []*string {
	return s.SourceIds
}

func (s *UpdateScheduledTaskRequestTaskDetailRelatedSkills) SetDisplayName(v string) *UpdateScheduledTaskRequestTaskDetailRelatedSkills {
	s.DisplayName = &v
	return s
}

func (s *UpdateScheduledTaskRequestTaskDetailRelatedSkills) SetName(v string) *UpdateScheduledTaskRequestTaskDetailRelatedSkills {
	s.Name = &v
	return s
}

func (s *UpdateScheduledTaskRequestTaskDetailRelatedSkills) SetSkillCode(v string) *UpdateScheduledTaskRequestTaskDetailRelatedSkills {
	s.SkillCode = &v
	return s
}

func (s *UpdateScheduledTaskRequestTaskDetailRelatedSkills) SetSourceIds(v []*string) *UpdateScheduledTaskRequestTaskDetailRelatedSkills {
	s.SourceIds = v
	return s
}

func (s *UpdateScheduledTaskRequestTaskDetailRelatedSkills) Validate() error {
	return dara.Validate(s)
}

type UpdateScheduledTaskRequestTriggerConfig struct {
	// The cron expression. Required when trigger_mode is set to scheduled. Example: \\"00 09 	- 	- *\\".
	//
	// example:
	//
	// string_value
	Cron *string `json:"cron,omitempty" xml:"cron,omitempty"`
	// The language, such as zh-CN or en-US. Automatically injected by the server.
	//
	// example:
	//
	// zh-CN
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
	// The list of push channels for the task. No push notifications are sent if the list is empty or no channel is enabled.
	PushConfig []*UpdateScheduledTaskRequestTriggerConfigPushConfig `json:"pushConfig,omitempty" xml:"pushConfig,omitempty" type:"Repeated"`
	// The time zone, such as Asia/Shanghai. Automatically injected by the server.
	//
	// example:
	//
	// Asia/Shanghai
	Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
	// The trigger mode. Valid values: manual and scheduled.
	//
	// This parameter is required.
	//
	// example:
	//
	// manual
	TriggerMode *string `json:"triggerMode,omitempty" xml:"triggerMode,omitempty"`
}

func (s UpdateScheduledTaskRequestTriggerConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateScheduledTaskRequestTriggerConfig) GoString() string {
	return s.String()
}

func (s *UpdateScheduledTaskRequestTriggerConfig) GetCron() *string {
	return s.Cron
}

func (s *UpdateScheduledTaskRequestTriggerConfig) GetLanguage() *string {
	return s.Language
}

func (s *UpdateScheduledTaskRequestTriggerConfig) GetPushConfig() []*UpdateScheduledTaskRequestTriggerConfigPushConfig {
	return s.PushConfig
}

func (s *UpdateScheduledTaskRequestTriggerConfig) GetTimezone() *string {
	return s.Timezone
}

func (s *UpdateScheduledTaskRequestTriggerConfig) GetTriggerMode() *string {
	return s.TriggerMode
}

func (s *UpdateScheduledTaskRequestTriggerConfig) SetCron(v string) *UpdateScheduledTaskRequestTriggerConfig {
	s.Cron = &v
	return s
}

func (s *UpdateScheduledTaskRequestTriggerConfig) SetLanguage(v string) *UpdateScheduledTaskRequestTriggerConfig {
	s.Language = &v
	return s
}

func (s *UpdateScheduledTaskRequestTriggerConfig) SetPushConfig(v []*UpdateScheduledTaskRequestTriggerConfigPushConfig) *UpdateScheduledTaskRequestTriggerConfig {
	s.PushConfig = v
	return s
}

func (s *UpdateScheduledTaskRequestTriggerConfig) SetTimezone(v string) *UpdateScheduledTaskRequestTriggerConfig {
	s.Timezone = &v
	return s
}

func (s *UpdateScheduledTaskRequestTriggerConfig) SetTriggerMode(v string) *UpdateScheduledTaskRequestTriggerConfig {
	s.TriggerMode = &v
	return s
}

func (s *UpdateScheduledTaskRequestTriggerConfig) Validate() error {
	if s.PushConfig != nil {
		for _, item := range s.PushConfig {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateScheduledTaskRequestTriggerConfigPushConfig struct {
	// The push channel type.
	//
	// example:
	//
	// DINGTALK
	ChannelType *string `json:"channelType,omitempty" xml:"channelType,omitempty"`
	// The scope of push content. Default value: all_replies.
	//
	// example:
	//
	// all_replies
	ContentScope *string `json:"contentScope,omitempty" xml:"contentScope,omitempty"`
	// The push method. Default value: channel_bot.
	//
	// example:
	//
	// channel_bot
	DeliveryMethod *string `json:"deliveryMethod,omitempty" xml:"deliveryMethod,omitempty"`
	// Specifies whether to push to this channel. Default value: false.
	//
	// example:
	//
	// true
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The format for pushing output files. Default value: file.
	//
	// example:
	//
	// file
	FileFormat *string `json:"fileFormat,omitempty" xml:"fileFormat,omitempty"`
	// The digital human to which the sending bot belongs. This parameter is required and cannot be empty.
	//
	// example:
	//
	// string_value
	OperatingObjectName *string `json:"operatingObjectName,omitempty" xml:"operatingObjectName,omitempty"`
	// The receiver type. Currently only self is supported.
	//
	// example:
	//
	// string_value
	ReceiverType *string `json:"receiverType,omitempty" xml:"receiverType,omitempty"`
}

func (s UpdateScheduledTaskRequestTriggerConfigPushConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateScheduledTaskRequestTriggerConfigPushConfig) GoString() string {
	return s.String()
}

func (s *UpdateScheduledTaskRequestTriggerConfigPushConfig) GetChannelType() *string {
	return s.ChannelType
}

func (s *UpdateScheduledTaskRequestTriggerConfigPushConfig) GetContentScope() *string {
	return s.ContentScope
}

func (s *UpdateScheduledTaskRequestTriggerConfigPushConfig) GetDeliveryMethod() *string {
	return s.DeliveryMethod
}

func (s *UpdateScheduledTaskRequestTriggerConfigPushConfig) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpdateScheduledTaskRequestTriggerConfigPushConfig) GetFileFormat() *string {
	return s.FileFormat
}

func (s *UpdateScheduledTaskRequestTriggerConfigPushConfig) GetOperatingObjectName() *string {
	return s.OperatingObjectName
}

func (s *UpdateScheduledTaskRequestTriggerConfigPushConfig) GetReceiverType() *string {
	return s.ReceiverType
}

func (s *UpdateScheduledTaskRequestTriggerConfigPushConfig) SetChannelType(v string) *UpdateScheduledTaskRequestTriggerConfigPushConfig {
	s.ChannelType = &v
	return s
}

func (s *UpdateScheduledTaskRequestTriggerConfigPushConfig) SetContentScope(v string) *UpdateScheduledTaskRequestTriggerConfigPushConfig {
	s.ContentScope = &v
	return s
}

func (s *UpdateScheduledTaskRequestTriggerConfigPushConfig) SetDeliveryMethod(v string) *UpdateScheduledTaskRequestTriggerConfigPushConfig {
	s.DeliveryMethod = &v
	return s
}

func (s *UpdateScheduledTaskRequestTriggerConfigPushConfig) SetEnabled(v bool) *UpdateScheduledTaskRequestTriggerConfigPushConfig {
	s.Enabled = &v
	return s
}

func (s *UpdateScheduledTaskRequestTriggerConfigPushConfig) SetFileFormat(v string) *UpdateScheduledTaskRequestTriggerConfigPushConfig {
	s.FileFormat = &v
	return s
}

func (s *UpdateScheduledTaskRequestTriggerConfigPushConfig) SetOperatingObjectName(v string) *UpdateScheduledTaskRequestTriggerConfigPushConfig {
	s.OperatingObjectName = &v
	return s
}

func (s *UpdateScheduledTaskRequestTriggerConfigPushConfig) SetReceiverType(v string) *UpdateScheduledTaskRequestTriggerConfigPushConfig {
	s.ReceiverType = &v
	return s
}

func (s *UpdateScheduledTaskRequestTriggerConfigPushConfig) Validate() error {
	return dara.Validate(s)
}
