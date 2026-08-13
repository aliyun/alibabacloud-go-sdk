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
}

type UpdateScheduledTaskRequest struct {
	Description []*UpdateScheduledTaskRequestDescription `json:"description,omitempty" xml:"description,omitempty" type:"Repeated"`
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
	// 执行模型档位；不传则不更新
	//
	// example:
	//
	// quick
	Model *string `json:"model,omitempty" xml:"model,omitempty"`
	// 文件名
	//
	// example:
	//
	// 示例名称.pdf
	Name       *string                               `json:"name,omitempty" xml:"name,omitempty"`
	Segments   []*UpdateScheduledTaskRequestSegments `json:"segments,omitempty" xml:"segments,omitempty" type:"Repeated"`
	TaskDetail *UpdateScheduledTaskRequestTaskDetail `json:"taskDetail,omitempty" xml:"taskDetail,omitempty" type:"Struct"`
	// 任务 ID
	//
	// This parameter is required.
	//
	// example:
	//
	// exampleTaskId
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// 租户ID，公共参数，缺省时使用调用方默认租户
	//
	// example:
	//
	// 10000
	TenantId      *string                                  `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	TriggerConfig *UpdateScheduledTaskRequestTriggerConfig `json:"triggerConfig,omitempty" xml:"triggerConfig,omitempty" type:"Struct"`
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
	RelatedObjects   []*UpdateScheduledTaskRequestTaskDetailRelatedObjects   `json:"relatedObjects,omitempty" xml:"relatedObjects,omitempty" type:"Repeated"`
	RelatedSemantics []*UpdateScheduledTaskRequestTaskDetailRelatedSemantics `json:"relatedSemantics,omitempty" xml:"relatedSemantics,omitempty" type:"Repeated"`
	RelatedSkills    []*UpdateScheduledTaskRequestTaskDetailRelatedSkills    `json:"relatedSkills,omitempty" xml:"relatedSkills,omitempty" type:"Repeated"`
	// LLM 润色后的任务理解描述
	//
	// This parameter is required.
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
	PushConfig []*UpdateScheduledTaskRequestTriggerConfigPushConfig `json:"pushConfig,omitempty" xml:"pushConfig,omitempty" type:"Repeated"`
	// 时区如 Asia/Shanghai，由服务端自动注入
	//
	// example:
	//
	// Asia/Shanghai
	Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
	// 触发模式：manual|scheduled
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
