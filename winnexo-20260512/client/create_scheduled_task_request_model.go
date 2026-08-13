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
}

type CreateScheduledTaskRequest struct {
	// 所属协作群组 ID（如 cg_101）；传入时创建群空间任务（调用者需为有效群成员），为空创建个人任务
	//
	// example:
	//
	// exampleCollaborationGroupId
	CollaborationGroupId *string                                  `json:"collaborationGroupId,omitempty" xml:"collaborationGroupId,omitempty"`
	Description          []*CreateScheduledTaskRequestDescription `json:"description,omitempty" xml:"description,omitempty" type:"Repeated"`
	// 数字员工名称列表
	//
	// example:
	//
	// string_value
	DigitalEmployeeName []*string `json:"digitalEmployeeName,omitempty" xml:"digitalEmployeeName,omitempty" type:"Repeated"`
	// 是否公开访问
	//
	// example:
	//
	// true
	IsOpen *bool `json:"isOpen,omitempty" xml:"isOpen,omitempty"`
	// 执行模型档位，不传默认 standard
	//
	// example:
	//
	// quick
	Model *string `json:"model,omitempty" xml:"model,omitempty"`
	// 文件名
	//
	// This parameter is required.
	//
	// example:
	//
	// 示例名称.pdf
	Name       *string                               `json:"name,omitempty" xml:"name,omitempty"`
	Segments   []*CreateScheduledTaskRequestSegments `json:"segments,omitempty" xml:"segments,omitempty" type:"Repeated"`
	TaskDetail *CreateScheduledTaskRequestTaskDetail `json:"taskDetail,omitempty" xml:"taskDetail,omitempty" type:"Struct"`
	// 租户ID，公共参数，缺省时使用调用方默认租户
	//
	// example:
	//
	// 10000
	TenantId      *string                                  `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	TriggerConfig *CreateScheduledTaskRequestTriggerConfig `json:"triggerConfig,omitempty" xml:"triggerConfig,omitempty" type:"Struct"`
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
	// 文本内容，type=text 时必填
	//
	// example:
	//
	// 示例内容
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// 功能开关，type=web_search 时可选
	//
	// example:
	//
	// true
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// 文件名
	//
	// example:
	//
	// 示例名称.pdf
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 对象 ID，type=mention 时有值
	//
	// example:
	//
	// exampleObjectId
	ObjectId *string `json:"objectId,omitempty" xml:"objectId,omitempty"`
	// 对象类型如 customer，type=mention 时有值
	//
	// example:
	//
	// string_value
	ObjectType *string `json:"objectType,omitempty" xml:"objectType,omitempty"`
	// 技能编码，type=skill 时有值
	//
	// example:
	//
	// string_value
	SkillCode *string `json:"skillCode,omitempty" xml:"skillCode,omitempty"`
	// 元素类型：text|web_search|mention|skill
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
	// 文本内容，type=text 时必填
	//
	// example:
	//
	// 示例内容
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// 功能开关，type=web_search 时可选
	//
	// example:
	//
	// true
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// 文件名
	//
	// example:
	//
	// 示例名称.pdf
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 对象 ID，type=mention 时有值
	//
	// example:
	//
	// exampleObjectId
	ObjectId *string `json:"objectId,omitempty" xml:"objectId,omitempty"`
	// 对象类型如 customer，type=mention 时有值
	//
	// example:
	//
	// string_value
	ObjectType *string `json:"objectType,omitempty" xml:"objectType,omitempty"`
	// 技能编码，type=skill 时有值
	//
	// example:
	//
	// string_value
	SkillCode *string `json:"skillCode,omitempty" xml:"skillCode,omitempty"`
	// 元素类型：text|web_search|mention|skill
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
	RelatedObjects   []*CreateScheduledTaskRequestTaskDetailRelatedObjects   `json:"relatedObjects,omitempty" xml:"relatedObjects,omitempty" type:"Repeated"`
	RelatedSemantics []*CreateScheduledTaskRequestTaskDetailRelatedSemantics `json:"relatedSemantics,omitempty" xml:"relatedSemantics,omitempty" type:"Repeated"`
	RelatedSkills    []*CreateScheduledTaskRequestTaskDetailRelatedSkills    `json:"relatedSkills,omitempty" xml:"relatedSkills,omitempty" type:"Repeated"`
	// LLM 润色后的任务理解描述
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
	// 提及类型，如 objects
	//
	// example:
	//
	// string_value
	MentionType *string `json:"mentionType,omitempty" xml:"mentionType,omitempty"`
	// 文件名
	//
	// example:
	//
	// 示例名称.pdf
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 对象 ID（@指定时有值）
	//
	// example:
	//
	// exampleObjectId
	ObjectId *string `json:"objectId,omitempty" xml:"objectId,omitempty"`
	// 对象类型，如 customer、company
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
	// 语义属性（JSON 字符串），用于语义检索时过滤
	//
	// example:
	//
	// {"level": "VIP"}
	Attributes *string `json:"attributes,omitempty" xml:"attributes,omitempty"`
	// 语义实体名，如客户/机会
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
	// 技能展示名称
	//
	// example:
	//
	// string_value
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// 文件名
	//
	// example:
	//
	// 示例名称.pdf
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 技能代码
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
	// Cron 表达式，trigger_mode=scheduled 时必填，如 \"00 09 	- 	- *\"
	//
	// example:
	//
	// string_value
	Cron *string `json:"cron,omitempty" xml:"cron,omitempty"`
	// 语言如 zh-CN|en-US，由服务端自动注入
	//
	// example:
	//
	// zh-CN
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
	// 任务推送频道列表；为空或无启用频道时不推送
	PushConfig []*CreateScheduledTaskRequestTriggerConfigPushConfig `json:"pushConfig,omitempty" xml:"pushConfig,omitempty" type:"Repeated"`
	// 时区如 Asia/Shanghai，由服务端自动注入
	//
	// example:
	//
	// Asia/Shanghai
	Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
	// 触发模式：manual|scheduled
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
	// 推送渠道
	//
	// example:
	//
	// DINGTALK
	ChannelType *string `json:"channelType,omitempty" xml:"channelType,omitempty"`
	// 推送内容范围，默认 all_replies
	//
	// example:
	//
	// all_replies
	ContentScope *string `json:"contentScope,omitempty" xml:"contentScope,omitempty"`
	// 推送方式，默认 channel_bot
	//
	// example:
	//
	// channel_bot
	DeliveryMethod *string `json:"deliveryMethod,omitempty" xml:"deliveryMethod,omitempty"`
	// 是否推送该频道，默认关闭
	//
	// example:
	//
	// true
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// 产出文件推送格式，默认 file
	//
	// example:
	//
	// file
	FileFormat *string `json:"fileFormat,omitempty" xml:"fileFormat,omitempty"`
	// 发送机器人所属数字员工，必传且不可为空
	//
	// example:
	//
	// string_value
	OperatingObjectName *string `json:"operatingObjectName,omitempty" xml:"operatingObjectName,omitempty"`
	// 接收人，当前仅支持 self
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
