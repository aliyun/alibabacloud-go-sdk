// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateScheduledTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCollaborationGroupId(v string) *CreateScheduledTaskRequest
	GetCollaborationGroupId() *string
	SetDescription(v []*CreateScheduledTaskRequestDescription) *CreateScheduledTaskRequest
	GetDescription() []*CreateScheduledTaskRequestDescription
	SetDigitalEmployeeName(v []*string) *CreateScheduledTaskRequest
	GetDigitalEmployeeName() []*string
	SetIsOpen(v bool) *CreateScheduledTaskRequest
	GetIsOpen() *bool
	SetModel(v string) *CreateScheduledTaskRequest
	GetModel() *string
	SetName(v string) *CreateScheduledTaskRequest
	GetName() *string
	SetSegments(v []*CreateScheduledTaskRequestSegments) *CreateScheduledTaskRequest
	GetSegments() []*CreateScheduledTaskRequestSegments
	SetTaskDetail(v *CreateScheduledTaskRequestTaskDetail) *CreateScheduledTaskRequest
	GetTaskDetail() *CreateScheduledTaskRequestTaskDetail
	SetTenantId(v string) *CreateScheduledTaskRequest
	GetTenantId() *string
	SetTriggerConfig(v *CreateScheduledTaskRequestTriggerConfig) *CreateScheduledTaskRequest
	GetTriggerConfig() *CreateScheduledTaskRequestTriggerConfig
	SetVisibility(v string) *CreateScheduledTaskRequest
	GetVisibility() *string
	SetVisibleMemberUserIds(v []*string) *CreateScheduledTaskRequest
	GetVisibleMemberUserIds() []*string
}

type CreateScheduledTaskRequest struct {
	// The ID of the collaboration group (such as cg_101). If specified, a group space task is created (the caller must be a valid group member). If empty, a personal task is created.
	//
	// example:
	//
	// exampleCollaborationGroupId
	CollaborationGroupId *string `json:"collaborationGroupId,omitempty" xml:"collaborationGroupId,omitempty"`
	// The description of the to-do card type.
	Description []*CreateScheduledTaskRequestDescription `json:"description,omitempty" xml:"description,omitempty" type:"Repeated"`
	// The name of the current effective digital employee. This parameter is empty if not configured.
	//
	// example:
	//
	// string_value
	DigitalEmployeeName []*string `json:"digitalEmployeeName,omitempty" xml:"digitalEmployeeName,omitempty" type:"Repeated"`
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
	Segments []*CreateScheduledTaskRequestSegments `json:"segments,omitempty" xml:"segments,omitempty" type:"Repeated"`
	// The task details.
	TaskDetail *CreateScheduledTaskRequestTaskDetail `json:"taskDetail,omitempty" xml:"taskDetail,omitempty" type:"Struct"`
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
	TriggerConfig *CreateScheduledTaskRequestTriggerConfig `json:"triggerConfig,omitempty" xml:"triggerConfig,omitempty" type:"Struct"`
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
	VisibleMemberUserIds []*string `json:"visibleMemberUserIds,omitempty" xml:"visibleMemberUserIds,omitempty" type:"Repeated"`
}

func (s CreateScheduledTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateScheduledTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateScheduledTaskRequest) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *CreateScheduledTaskRequest) GetDescription() []*CreateScheduledTaskRequestDescription {
	return s.Description
}

func (s *CreateScheduledTaskRequest) GetDigitalEmployeeName() []*string {
	return s.DigitalEmployeeName
}

func (s *CreateScheduledTaskRequest) GetIsOpen() *bool {
	return s.IsOpen
}

func (s *CreateScheduledTaskRequest) GetModel() *string {
	return s.Model
}

func (s *CreateScheduledTaskRequest) GetName() *string {
	return s.Name
}

func (s *CreateScheduledTaskRequest) GetSegments() []*CreateScheduledTaskRequestSegments {
	return s.Segments
}

func (s *CreateScheduledTaskRequest) GetTaskDetail() *CreateScheduledTaskRequestTaskDetail {
	return s.TaskDetail
}

func (s *CreateScheduledTaskRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreateScheduledTaskRequest) GetTriggerConfig() *CreateScheduledTaskRequestTriggerConfig {
	return s.TriggerConfig
}

func (s *CreateScheduledTaskRequest) GetVisibility() *string {
	return s.Visibility
}

func (s *CreateScheduledTaskRequest) GetVisibleMemberUserIds() []*string {
	return s.VisibleMemberUserIds
}

func (s *CreateScheduledTaskRequest) SetCollaborationGroupId(v string) *CreateScheduledTaskRequest {
	s.CollaborationGroupId = &v
	return s
}

func (s *CreateScheduledTaskRequest) SetDescription(v []*CreateScheduledTaskRequestDescription) *CreateScheduledTaskRequest {
	s.Description = v
	return s
}

func (s *CreateScheduledTaskRequest) SetDigitalEmployeeName(v []*string) *CreateScheduledTaskRequest {
	s.DigitalEmployeeName = v
	return s
}

func (s *CreateScheduledTaskRequest) SetIsOpen(v bool) *CreateScheduledTaskRequest {
	s.IsOpen = &v
	return s
}

func (s *CreateScheduledTaskRequest) SetModel(v string) *CreateScheduledTaskRequest {
	s.Model = &v
	return s
}

func (s *CreateScheduledTaskRequest) SetName(v string) *CreateScheduledTaskRequest {
	s.Name = &v
	return s
}

func (s *CreateScheduledTaskRequest) SetSegments(v []*CreateScheduledTaskRequestSegments) *CreateScheduledTaskRequest {
	s.Segments = v
	return s
}

func (s *CreateScheduledTaskRequest) SetTaskDetail(v *CreateScheduledTaskRequestTaskDetail) *CreateScheduledTaskRequest {
	s.TaskDetail = v
	return s
}

func (s *CreateScheduledTaskRequest) SetTenantId(v string) *CreateScheduledTaskRequest {
	s.TenantId = &v
	return s
}

func (s *CreateScheduledTaskRequest) SetTriggerConfig(v *CreateScheduledTaskRequestTriggerConfig) *CreateScheduledTaskRequest {
	s.TriggerConfig = v
	return s
}

func (s *CreateScheduledTaskRequest) SetVisibility(v string) *CreateScheduledTaskRequest {
	s.Visibility = &v
	return s
}

func (s *CreateScheduledTaskRequest) SetVisibleMemberUserIds(v []*string) *CreateScheduledTaskRequest {
	s.VisibleMemberUserIds = v
	return s
}

func (s *CreateScheduledTaskRequest) Validate() error {
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

type CreateScheduledTaskRequestDescription struct {
	// The streaming output message.
	//
	// example:
	//
	// Sample content
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// Specifies whether the throttling rule is enabled. A value of true indicates enabled, and a value of false indicates disabled.
	//
	// example:
	//
	// true
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The name.
	//
	// example:
	//
	// SampleName.pdf
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The object ID. Pass the project task ID.
	//
	// - For internal enterprise applications, use the taskId obtained by calling the [Create a project task](https://open.dingtalk.com/document/orgapp-server/create-a-project-task) operation.
	//
	// - For third-party enterprise applications, use the taskId obtained by calling the [Create a project task](https://open.dingtalk.com/document/isvapp-server/create-a-project-task) operation.
	//
	// example:
	//
	// exampleObjectId
	ObjectId *string `json:"objectId,omitempty" xml:"objectId,omitempty"`
	// The object type. Fixed value: task, indicating a project task.
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
	// The HTTP API type. Valid values: Http (standard HTTP API), Rest (RESTful API), WebSocket (WebSocket API), HttpIngress (HTTP API accessed through Ingress), LLM (large language model API), and Agent (Agent proxy API).
	//
	// example:
	//
	// text
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateScheduledTaskRequestDescription) String() string {
	return dara.Prettify(s)
}

func (s CreateScheduledTaskRequestDescription) GoString() string {
	return s.String()
}

func (s *CreateScheduledTaskRequestDescription) GetContent() *string {
	return s.Content
}

func (s *CreateScheduledTaskRequestDescription) GetEnabled() *bool {
	return s.Enabled
}

func (s *CreateScheduledTaskRequestDescription) GetName() *string {
	return s.Name
}

func (s *CreateScheduledTaskRequestDescription) GetObjectId() *string {
	return s.ObjectId
}

func (s *CreateScheduledTaskRequestDescription) GetObjectType() *string {
	return s.ObjectType
}

func (s *CreateScheduledTaskRequestDescription) GetSkillCode() *string {
	return s.SkillCode
}

func (s *CreateScheduledTaskRequestDescription) GetType() *string {
	return s.Type
}

func (s *CreateScheduledTaskRequestDescription) SetContent(v string) *CreateScheduledTaskRequestDescription {
	s.Content = &v
	return s
}

func (s *CreateScheduledTaskRequestDescription) SetEnabled(v bool) *CreateScheduledTaskRequestDescription {
	s.Enabled = &v
	return s
}

func (s *CreateScheduledTaskRequestDescription) SetName(v string) *CreateScheduledTaskRequestDescription {
	s.Name = &v
	return s
}

func (s *CreateScheduledTaskRequestDescription) SetObjectId(v string) *CreateScheduledTaskRequestDescription {
	s.ObjectId = &v
	return s
}

func (s *CreateScheduledTaskRequestDescription) SetObjectType(v string) *CreateScheduledTaskRequestDescription {
	s.ObjectType = &v
	return s
}

func (s *CreateScheduledTaskRequestDescription) SetSkillCode(v string) *CreateScheduledTaskRequestDescription {
	s.SkillCode = &v
	return s
}

func (s *CreateScheduledTaskRequestDescription) SetType(v string) *CreateScheduledTaskRequestDescription {
	s.Type = &v
	return s
}

func (s *CreateScheduledTaskRequestDescription) Validate() error {
	return dara.Validate(s)
}

type CreateScheduledTaskRequestSegments struct {
	// The card callback content.
	//
	// example:
	//
	// Sample content
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// Specifies whether to enable this feature.
	//
	// example:
	//
	// true
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The name.
	//
	// example:
	//
	// SampleName.pdf
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The ID of the recommended item, which can be a **feedId*	- or a micro-application ID.
	//
	// example:
	//
	// exampleObjectId
	ObjectId *string `json:"objectId,omitempty" xml:"objectId,omitempty"`
	// The customer type to save.
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
	// The billing type. Only fixed is supported.
	//
	// example:
	//
	// text
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateScheduledTaskRequestSegments) String() string {
	return dara.Prettify(s)
}

func (s CreateScheduledTaskRequestSegments) GoString() string {
	return s.String()
}

func (s *CreateScheduledTaskRequestSegments) GetContent() *string {
	return s.Content
}

func (s *CreateScheduledTaskRequestSegments) GetEnabled() *bool {
	return s.Enabled
}

func (s *CreateScheduledTaskRequestSegments) GetName() *string {
	return s.Name
}

func (s *CreateScheduledTaskRequestSegments) GetObjectId() *string {
	return s.ObjectId
}

func (s *CreateScheduledTaskRequestSegments) GetObjectType() *string {
	return s.ObjectType
}

func (s *CreateScheduledTaskRequestSegments) GetSkillCode() *string {
	return s.SkillCode
}

func (s *CreateScheduledTaskRequestSegments) GetType() *string {
	return s.Type
}

func (s *CreateScheduledTaskRequestSegments) SetContent(v string) *CreateScheduledTaskRequestSegments {
	s.Content = &v
	return s
}

func (s *CreateScheduledTaskRequestSegments) SetEnabled(v bool) *CreateScheduledTaskRequestSegments {
	s.Enabled = &v
	return s
}

func (s *CreateScheduledTaskRequestSegments) SetName(v string) *CreateScheduledTaskRequestSegments {
	s.Name = &v
	return s
}

func (s *CreateScheduledTaskRequestSegments) SetObjectId(v string) *CreateScheduledTaskRequestSegments {
	s.ObjectId = &v
	return s
}

func (s *CreateScheduledTaskRequestSegments) SetObjectType(v string) *CreateScheduledTaskRequestSegments {
	s.ObjectType = &v
	return s
}

func (s *CreateScheduledTaskRequestSegments) SetSkillCode(v string) *CreateScheduledTaskRequestSegments {
	s.SkillCode = &v
	return s
}

func (s *CreateScheduledTaskRequestSegments) SetType(v string) *CreateScheduledTaskRequestSegments {
	s.Type = &v
	return s
}

func (s *CreateScheduledTaskRequestSegments) Validate() error {
	return dara.Validate(s)
}

type CreateScheduledTaskRequestTaskDetail struct {
	// The related objects.
	RelatedObjects []*CreateScheduledTaskRequestTaskDetailRelatedObjects `json:"relatedObjects,omitempty" xml:"relatedObjects,omitempty" type:"Repeated"`
	// The related semantics.
	RelatedSemantics []*CreateScheduledTaskRequestTaskDetailRelatedSemantics `json:"relatedSemantics,omitempty" xml:"relatedSemantics,omitempty" type:"Repeated"`
	// The related skills.
	RelatedSkills []*CreateScheduledTaskRequestTaskDetailRelatedSkills `json:"relatedSkills,omitempty" xml:"relatedSkills,omitempty" type:"Repeated"`
	// The task understanding description polished by the LLM.
	//
	// example:
	//
	// string_value
	TaskUnderstand *string `json:"taskUnderstand,omitempty" xml:"taskUnderstand,omitempty"`
}

func (s CreateScheduledTaskRequestTaskDetail) String() string {
	return dara.Prettify(s)
}

func (s CreateScheduledTaskRequestTaskDetail) GoString() string {
	return s.String()
}

func (s *CreateScheduledTaskRequestTaskDetail) GetRelatedObjects() []*CreateScheduledTaskRequestTaskDetailRelatedObjects {
	return s.RelatedObjects
}

func (s *CreateScheduledTaskRequestTaskDetail) GetRelatedSemantics() []*CreateScheduledTaskRequestTaskDetailRelatedSemantics {
	return s.RelatedSemantics
}

func (s *CreateScheduledTaskRequestTaskDetail) GetRelatedSkills() []*CreateScheduledTaskRequestTaskDetailRelatedSkills {
	return s.RelatedSkills
}

func (s *CreateScheduledTaskRequestTaskDetail) GetTaskUnderstand() *string {
	return s.TaskUnderstand
}

func (s *CreateScheduledTaskRequestTaskDetail) SetRelatedObjects(v []*CreateScheduledTaskRequestTaskDetailRelatedObjects) *CreateScheduledTaskRequestTaskDetail {
	s.RelatedObjects = v
	return s
}

func (s *CreateScheduledTaskRequestTaskDetail) SetRelatedSemantics(v []*CreateScheduledTaskRequestTaskDetailRelatedSemantics) *CreateScheduledTaskRequestTaskDetail {
	s.RelatedSemantics = v
	return s
}

func (s *CreateScheduledTaskRequestTaskDetail) SetRelatedSkills(v []*CreateScheduledTaskRequestTaskDetailRelatedSkills) *CreateScheduledTaskRequestTaskDetail {
	s.RelatedSkills = v
	return s
}

func (s *CreateScheduledTaskRequestTaskDetail) SetTaskUnderstand(v string) *CreateScheduledTaskRequestTaskDetail {
	s.TaskUnderstand = &v
	return s
}

func (s *CreateScheduledTaskRequestTaskDetail) Validate() error {
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

type CreateScheduledTaskRequestTaskDetailRelatedObjects struct {
	// The mention type, such as objects.
	//
	// example:
	//
	// string_value
	MentionType *string `json:"mentionType,omitempty" xml:"mentionType,omitempty"`
	// The name.
	//
	// example:
	//
	// SampleName.pdf
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The object ID. Pass the project task ID.
	//
	// - For internal enterprise applications, use the taskId obtained by calling the [Create a project task](https://open.dingtalk.com/document/orgapp-server/create-a-project-task) operation.
	//
	// - For third-party enterprise applications, use the taskId obtained by calling the [Create a project task](https://open.dingtalk.com/document/isvapp-server/create-a-project-task) operation.
	//
	// example:
	//
	// exampleObjectId
	ObjectId *string `json:"objectId,omitempty" xml:"objectId,omitempty"`
	// The relationship type. Valid values:
	//
	// - crm_customer: enterprise customer.
	//
	// - crm_customer_personal: individual customer.
	//
	// example:
	//
	// string_value
	ObjectType *string `json:"objectType,omitempty" xml:"objectType,omitempty"`
}

func (s CreateScheduledTaskRequestTaskDetailRelatedObjects) String() string {
	return dara.Prettify(s)
}

func (s CreateScheduledTaskRequestTaskDetailRelatedObjects) GoString() string {
	return s.String()
}

func (s *CreateScheduledTaskRequestTaskDetailRelatedObjects) GetMentionType() *string {
	return s.MentionType
}

func (s *CreateScheduledTaskRequestTaskDetailRelatedObjects) GetName() *string {
	return s.Name
}

func (s *CreateScheduledTaskRequestTaskDetailRelatedObjects) GetObjectId() *string {
	return s.ObjectId
}

func (s *CreateScheduledTaskRequestTaskDetailRelatedObjects) GetObjectType() *string {
	return s.ObjectType
}

func (s *CreateScheduledTaskRequestTaskDetailRelatedObjects) SetMentionType(v string) *CreateScheduledTaskRequestTaskDetailRelatedObjects {
	s.MentionType = &v
	return s
}

func (s *CreateScheduledTaskRequestTaskDetailRelatedObjects) SetName(v string) *CreateScheduledTaskRequestTaskDetailRelatedObjects {
	s.Name = &v
	return s
}

func (s *CreateScheduledTaskRequestTaskDetailRelatedObjects) SetObjectId(v string) *CreateScheduledTaskRequestTaskDetailRelatedObjects {
	s.ObjectId = &v
	return s
}

func (s *CreateScheduledTaskRequestTaskDetailRelatedObjects) SetObjectType(v string) *CreateScheduledTaskRequestTaskDetailRelatedObjects {
	s.ObjectType = &v
	return s
}

func (s *CreateScheduledTaskRequestTaskDetailRelatedObjects) Validate() error {
	return dara.Validate(s)
}

type CreateScheduledTaskRequestTaskDetailRelatedSemantics struct {
	// The file extension information.
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

func (s CreateScheduledTaskRequestTaskDetailRelatedSemantics) String() string {
	return dara.Prettify(s)
}

func (s CreateScheduledTaskRequestTaskDetailRelatedSemantics) GoString() string {
	return s.String()
}

func (s *CreateScheduledTaskRequestTaskDetailRelatedSemantics) GetAttributes() *string {
	return s.Attributes
}

func (s *CreateScheduledTaskRequestTaskDetailRelatedSemantics) GetEntity() *string {
	return s.Entity
}

func (s *CreateScheduledTaskRequestTaskDetailRelatedSemantics) SetAttributes(v string) *CreateScheduledTaskRequestTaskDetailRelatedSemantics {
	s.Attributes = &v
	return s
}

func (s *CreateScheduledTaskRequestTaskDetailRelatedSemantics) SetEntity(v string) *CreateScheduledTaskRequestTaskDetailRelatedSemantics {
	s.Entity = &v
	return s
}

func (s *CreateScheduledTaskRequestTaskDetailRelatedSemantics) Validate() error {
	return dara.Validate(s)
}

type CreateScheduledTaskRequestTaskDetailRelatedSkills struct {
	// The display name.
	//
	// example:
	//
	// string_value
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// The name.
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

func (s CreateScheduledTaskRequestTaskDetailRelatedSkills) String() string {
	return dara.Prettify(s)
}

func (s CreateScheduledTaskRequestTaskDetailRelatedSkills) GoString() string {
	return s.String()
}

func (s *CreateScheduledTaskRequestTaskDetailRelatedSkills) GetDisplayName() *string {
	return s.DisplayName
}

func (s *CreateScheduledTaskRequestTaskDetailRelatedSkills) GetName() *string {
	return s.Name
}

func (s *CreateScheduledTaskRequestTaskDetailRelatedSkills) GetSkillCode() *string {
	return s.SkillCode
}

func (s *CreateScheduledTaskRequestTaskDetailRelatedSkills) GetSourceIds() []*string {
	return s.SourceIds
}

func (s *CreateScheduledTaskRequestTaskDetailRelatedSkills) SetDisplayName(v string) *CreateScheduledTaskRequestTaskDetailRelatedSkills {
	s.DisplayName = &v
	return s
}

func (s *CreateScheduledTaskRequestTaskDetailRelatedSkills) SetName(v string) *CreateScheduledTaskRequestTaskDetailRelatedSkills {
	s.Name = &v
	return s
}

func (s *CreateScheduledTaskRequestTaskDetailRelatedSkills) SetSkillCode(v string) *CreateScheduledTaskRequestTaskDetailRelatedSkills {
	s.SkillCode = &v
	return s
}

func (s *CreateScheduledTaskRequestTaskDetailRelatedSkills) SetSourceIds(v []*string) *CreateScheduledTaskRequestTaskDetailRelatedSkills {
	s.SourceIds = v
	return s
}

func (s *CreateScheduledTaskRequestTaskDetailRelatedSkills) Validate() error {
	return dara.Validate(s)
}

type CreateScheduledTaskRequestTriggerConfig struct {
	// The periodic training information in cron syntax (Minutes Hours DayofMonth Month DayofWeek). An empty value indicates that periodic training is not performed (default). In DayofWeek, 0 indicates Sunday.
	//
	// example:
	//
	// string_value
	Cron *string `json:"cron,omitempty" xml:"cron,omitempty"`
	// The language. Valid values:
	//
	// - zh_CN: Chinese (default)
	//
	// - en_US: English
	//
	// example:
	//
	// zh-CN
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
	// The list of task push channels. No push is performed if the list is empty or no channel is enabled.
	PushConfig []*CreateScheduledTaskRequestTriggerConfigPushConfig `json:"pushConfig,omitempty" xml:"pushConfig,omitempty" type:"Repeated"`
	// The time zone.
	//
	// example:
	//
	// Asia/Shanghai
	Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
	// The trigger mode.
	//
	//
	//
	//   1: Manual trigger
	//
	//
	//
	//   2: Scheduled trigger
	//
	//   3: Code commit trigger
	//
	//
	//
	//   5: Pipeline trigger
	//
	//   6: WEBHOOK trigger
	//
	// example:
	//
	// manual
	TriggerMode *string `json:"triggerMode,omitempty" xml:"triggerMode,omitempty"`
}

func (s CreateScheduledTaskRequestTriggerConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateScheduledTaskRequestTriggerConfig) GoString() string {
	return s.String()
}

func (s *CreateScheduledTaskRequestTriggerConfig) GetCron() *string {
	return s.Cron
}

func (s *CreateScheduledTaskRequestTriggerConfig) GetLanguage() *string {
	return s.Language
}

func (s *CreateScheduledTaskRequestTriggerConfig) GetPushConfig() []*CreateScheduledTaskRequestTriggerConfigPushConfig {
	return s.PushConfig
}

func (s *CreateScheduledTaskRequestTriggerConfig) GetTimezone() *string {
	return s.Timezone
}

func (s *CreateScheduledTaskRequestTriggerConfig) GetTriggerMode() *string {
	return s.TriggerMode
}

func (s *CreateScheduledTaskRequestTriggerConfig) SetCron(v string) *CreateScheduledTaskRequestTriggerConfig {
	s.Cron = &v
	return s
}

func (s *CreateScheduledTaskRequestTriggerConfig) SetLanguage(v string) *CreateScheduledTaskRequestTriggerConfig {
	s.Language = &v
	return s
}

func (s *CreateScheduledTaskRequestTriggerConfig) SetPushConfig(v []*CreateScheduledTaskRequestTriggerConfigPushConfig) *CreateScheduledTaskRequestTriggerConfig {
	s.PushConfig = v
	return s
}

func (s *CreateScheduledTaskRequestTriggerConfig) SetTimezone(v string) *CreateScheduledTaskRequestTriggerConfig {
	s.Timezone = &v
	return s
}

func (s *CreateScheduledTaskRequestTriggerConfig) SetTriggerMode(v string) *CreateScheduledTaskRequestTriggerConfig {
	s.TriggerMode = &v
	return s
}

func (s *CreateScheduledTaskRequestTriggerConfig) Validate() error {
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

type CreateScheduledTaskRequestTriggerConfigPushConfig struct {
	// The notification method. Valid values:
	//
	// - **hdm_alarm_sms**: SMS.
	//
	// - **dingtalk**: DingTalk chatbot.
	//
	// - **hdm_alarm_sms_and_email**: SMS and email.
	//
	// - **hdm_alarm_sms,dingtalk**: SMS and DingTalk chatbot.
	//
	// example:
	//
	// DINGTALK
	ChannelType *string `json:"channelType,omitempty" xml:"channelType,omitempty"`
	// The push content scope. Default value: all_replies.
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
	// Specifies whether the credential is enabled. Valid values:
	//
	// - true: Enabled.
	//
	// - false: Disabled.
	//
	// example:
	//
	// true
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The file format. Valid values: Excel and CSV.
	//
	// example:
	//
	// file
	FileFormat *string `json:"fileFormat,omitempty" xml:"fileFormat,omitempty"`
	// The digital employee name (operating object name, optional).
	//
	// example:
	//
	// string_value
	OperatingObjectName *string `json:"operatingObjectName,omitempty" xml:"operatingObjectName,omitempty"`
	// The file receiver type. Valid values:
	//
	// - 0: One-on-one chat.
	//
	// - 1: Group chat.
	//
	// - 2: DingTalk Drive.
	//
	// - 3: Document.
	//
	// example:
	//
	// string_value
	ReceiverType *string `json:"receiverType,omitempty" xml:"receiverType,omitempty"`
}

func (s CreateScheduledTaskRequestTriggerConfigPushConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateScheduledTaskRequestTriggerConfigPushConfig) GoString() string {
	return s.String()
}

func (s *CreateScheduledTaskRequestTriggerConfigPushConfig) GetChannelType() *string {
	return s.ChannelType
}

func (s *CreateScheduledTaskRequestTriggerConfigPushConfig) GetContentScope() *string {
	return s.ContentScope
}

func (s *CreateScheduledTaskRequestTriggerConfigPushConfig) GetDeliveryMethod() *string {
	return s.DeliveryMethod
}

func (s *CreateScheduledTaskRequestTriggerConfigPushConfig) GetEnabled() *bool {
	return s.Enabled
}

func (s *CreateScheduledTaskRequestTriggerConfigPushConfig) GetFileFormat() *string {
	return s.FileFormat
}

func (s *CreateScheduledTaskRequestTriggerConfigPushConfig) GetOperatingObjectName() *string {
	return s.OperatingObjectName
}

func (s *CreateScheduledTaskRequestTriggerConfigPushConfig) GetReceiverType() *string {
	return s.ReceiverType
}

func (s *CreateScheduledTaskRequestTriggerConfigPushConfig) SetChannelType(v string) *CreateScheduledTaskRequestTriggerConfigPushConfig {
	s.ChannelType = &v
	return s
}

func (s *CreateScheduledTaskRequestTriggerConfigPushConfig) SetContentScope(v string) *CreateScheduledTaskRequestTriggerConfigPushConfig {
	s.ContentScope = &v
	return s
}

func (s *CreateScheduledTaskRequestTriggerConfigPushConfig) SetDeliveryMethod(v string) *CreateScheduledTaskRequestTriggerConfigPushConfig {
	s.DeliveryMethod = &v
	return s
}

func (s *CreateScheduledTaskRequestTriggerConfigPushConfig) SetEnabled(v bool) *CreateScheduledTaskRequestTriggerConfigPushConfig {
	s.Enabled = &v
	return s
}

func (s *CreateScheduledTaskRequestTriggerConfigPushConfig) SetFileFormat(v string) *CreateScheduledTaskRequestTriggerConfigPushConfig {
	s.FileFormat = &v
	return s
}

func (s *CreateScheduledTaskRequestTriggerConfigPushConfig) SetOperatingObjectName(v string) *CreateScheduledTaskRequestTriggerConfigPushConfig {
	s.OperatingObjectName = &v
	return s
}

func (s *CreateScheduledTaskRequestTriggerConfigPushConfig) SetReceiverType(v string) *CreateScheduledTaskRequestTriggerConfigPushConfig {
	s.ReceiverType = &v
	return s
}

func (s *CreateScheduledTaskRequestTriggerConfigPushConfig) Validate() error {
	return dara.Validate(s)
}
