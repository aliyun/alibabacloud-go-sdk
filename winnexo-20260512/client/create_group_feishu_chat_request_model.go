// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGroupFeishuChatRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChatId(v string) *CreateGroupFeishuChatRequest
	GetChatId() *string
	SetDescription(v string) *CreateGroupFeishuChatRequest
	GetDescription() *string
	SetDirectoryId(v string) *CreateGroupFeishuChatRequest
	GetDirectoryId() *string
	SetGroupId(v string) *CreateGroupFeishuChatRequest
	GetGroupId() *string
	SetHistoryStartTime(v string) *CreateGroupFeishuChatRequest
	GetHistoryStartTime() *string
	SetNotes(v string) *CreateGroupFeishuChatRequest
	GetNotes() *string
	SetOperatingObjectName(v string) *CreateGroupFeishuChatRequest
	GetOperatingObjectName() *string
	SetSourceTags(v string) *CreateGroupFeishuChatRequest
	GetSourceTags() *string
	SetTenantId(v string) *CreateGroupFeishuChatRequest
	GetTenantId() *string
	SetUpdateFrequency(v *CreateGroupFeishuChatRequestUpdateFrequency) *CreateGroupFeishuChatRequest
	GetUpdateFrequency() *CreateGroupFeishuChatRequestUpdateFrequency
}

type CreateGroupFeishuChatRequest struct {
	// 飞书群聊ID，以oc_开头，需当前用户有权读取
	//
	// This parameter is required.
	//
	// example:
	//
	// cidxxxxxxxx
	ChatId *string `json:"chatId,omitempty" xml:"chatId,omitempty"`
	// 资料描述
	//
	// example:
	//
	// string_value
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 空间物理目录ID；省略/root使用空间根，首次可能初始化根目录
	//
	// example:
	//
	// exampleDirectoryId
	DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
	// 协作空间 ID
	//
	// This parameter is required.
	//
	// example:
	//
	// exampleGroupId
	GroupId *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// 历史起始时间，YYYY-MM-DD或YYYY-MM-DD HH:MM:SS；省略读取全部可见历史
	//
	// example:
	//
	// 2026-08-01
	HistoryStartTime *string `json:"historyStartTime,omitempty" xml:"historyStartTime,omitempty"`
	// 分析指令
	//
	// example:
	//
	// 重点识别客户诉求与待办
	Notes *string `json:"notes,omitempty" xml:"notes,omitempty"`
	// 运营对象名称，用于来源追溯
	//
	// example:
	//
	// string_value
	OperatingObjectName *string `json:"operatingObjectName,omitempty" xml:"operatingObjectName,omitempty"`
	// 资料标签JSON字符串列表
	//
	// example:
	//
	// ["重点","文件"]
	SourceTags *string `json:"sourceTags,omitempty" xml:"sourceTags,omitempty"`
	// 租户ID，公共参数；缺省时使用调用方默认租户
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// Source级同步配置
	UpdateFrequency *CreateGroupFeishuChatRequestUpdateFrequency `json:"updateFrequency,omitempty" xml:"updateFrequency,omitempty" type:"Struct"`
}

func (s CreateGroupFeishuChatRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateGroupFeishuChatRequest) GoString() string {
	return s.String()
}

func (s *CreateGroupFeishuChatRequest) GetChatId() *string {
	return s.ChatId
}

func (s *CreateGroupFeishuChatRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateGroupFeishuChatRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *CreateGroupFeishuChatRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *CreateGroupFeishuChatRequest) GetHistoryStartTime() *string {
	return s.HistoryStartTime
}

func (s *CreateGroupFeishuChatRequest) GetNotes() *string {
	return s.Notes
}

func (s *CreateGroupFeishuChatRequest) GetOperatingObjectName() *string {
	return s.OperatingObjectName
}

func (s *CreateGroupFeishuChatRequest) GetSourceTags() *string {
	return s.SourceTags
}

func (s *CreateGroupFeishuChatRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreateGroupFeishuChatRequest) GetUpdateFrequency() *CreateGroupFeishuChatRequestUpdateFrequency {
	return s.UpdateFrequency
}

func (s *CreateGroupFeishuChatRequest) SetChatId(v string) *CreateGroupFeishuChatRequest {
	s.ChatId = &v
	return s
}

func (s *CreateGroupFeishuChatRequest) SetDescription(v string) *CreateGroupFeishuChatRequest {
	s.Description = &v
	return s
}

func (s *CreateGroupFeishuChatRequest) SetDirectoryId(v string) *CreateGroupFeishuChatRequest {
	s.DirectoryId = &v
	return s
}

func (s *CreateGroupFeishuChatRequest) SetGroupId(v string) *CreateGroupFeishuChatRequest {
	s.GroupId = &v
	return s
}

func (s *CreateGroupFeishuChatRequest) SetHistoryStartTime(v string) *CreateGroupFeishuChatRequest {
	s.HistoryStartTime = &v
	return s
}

func (s *CreateGroupFeishuChatRequest) SetNotes(v string) *CreateGroupFeishuChatRequest {
	s.Notes = &v
	return s
}

func (s *CreateGroupFeishuChatRequest) SetOperatingObjectName(v string) *CreateGroupFeishuChatRequest {
	s.OperatingObjectName = &v
	return s
}

func (s *CreateGroupFeishuChatRequest) SetSourceTags(v string) *CreateGroupFeishuChatRequest {
	s.SourceTags = &v
	return s
}

func (s *CreateGroupFeishuChatRequest) SetTenantId(v string) *CreateGroupFeishuChatRequest {
	s.TenantId = &v
	return s
}

func (s *CreateGroupFeishuChatRequest) SetUpdateFrequency(v *CreateGroupFeishuChatRequestUpdateFrequency) *CreateGroupFeishuChatRequest {
	s.UpdateFrequency = v
	return s
}

func (s *CreateGroupFeishuChatRequest) Validate() error {
	if s.UpdateFrequency != nil {
		if err := s.UpdateFrequency.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateGroupFeishuChatRequestUpdateFrequency struct {
	// 五段 cron，优先于 preset
	//
	// example:
	//
	// 0 2 	- 	- *
	Cron *string `json:"cron,omitempty" xml:"cron,omitempty"`
	// 是否启用同步，默认true
	//
	// example:
	//
	// true
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// 同步预设：hourly 或 daily_2am
	//
	// example:
	//
	// hourly
	Preset *string `json:"preset,omitempty" xml:"preset,omitempty"`
}

func (s CreateGroupFeishuChatRequestUpdateFrequency) String() string {
	return dara.Prettify(s)
}

func (s CreateGroupFeishuChatRequestUpdateFrequency) GoString() string {
	return s.String()
}

func (s *CreateGroupFeishuChatRequestUpdateFrequency) GetCron() *string {
	return s.Cron
}

func (s *CreateGroupFeishuChatRequestUpdateFrequency) GetEnabled() *bool {
	return s.Enabled
}

func (s *CreateGroupFeishuChatRequestUpdateFrequency) GetPreset() *string {
	return s.Preset
}

func (s *CreateGroupFeishuChatRequestUpdateFrequency) SetCron(v string) *CreateGroupFeishuChatRequestUpdateFrequency {
	s.Cron = &v
	return s
}

func (s *CreateGroupFeishuChatRequestUpdateFrequency) SetEnabled(v bool) *CreateGroupFeishuChatRequestUpdateFrequency {
	s.Enabled = &v
	return s
}

func (s *CreateGroupFeishuChatRequestUpdateFrequency) SetPreset(v string) *CreateGroupFeishuChatRequestUpdateFrequency {
	s.Preset = &v
	return s
}

func (s *CreateGroupFeishuChatRequestUpdateFrequency) Validate() error {
	return dara.Validate(s)
}
