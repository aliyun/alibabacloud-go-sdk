// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGroupFeishuChatShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChatId(v string) *CreateGroupFeishuChatShrinkRequest
	GetChatId() *string
	SetDescription(v string) *CreateGroupFeishuChatShrinkRequest
	GetDescription() *string
	SetDirectoryId(v string) *CreateGroupFeishuChatShrinkRequest
	GetDirectoryId() *string
	SetGroupId(v string) *CreateGroupFeishuChatShrinkRequest
	GetGroupId() *string
	SetHistoryStartTime(v string) *CreateGroupFeishuChatShrinkRequest
	GetHistoryStartTime() *string
	SetNotes(v string) *CreateGroupFeishuChatShrinkRequest
	GetNotes() *string
	SetOperatingObjectName(v string) *CreateGroupFeishuChatShrinkRequest
	GetOperatingObjectName() *string
	SetSourceTags(v string) *CreateGroupFeishuChatShrinkRequest
	GetSourceTags() *string
	SetTenantId(v string) *CreateGroupFeishuChatShrinkRequest
	GetTenantId() *string
	SetUpdateFrequencyShrink(v string) *CreateGroupFeishuChatShrinkRequest
	GetUpdateFrequencyShrink() *string
}

type CreateGroupFeishuChatShrinkRequest struct {
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
	UpdateFrequencyShrink *string `json:"updateFrequency,omitempty" xml:"updateFrequency,omitempty"`
}

func (s CreateGroupFeishuChatShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateGroupFeishuChatShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateGroupFeishuChatShrinkRequest) GetChatId() *string {
	return s.ChatId
}

func (s *CreateGroupFeishuChatShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateGroupFeishuChatShrinkRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *CreateGroupFeishuChatShrinkRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *CreateGroupFeishuChatShrinkRequest) GetHistoryStartTime() *string {
	return s.HistoryStartTime
}

func (s *CreateGroupFeishuChatShrinkRequest) GetNotes() *string {
	return s.Notes
}

func (s *CreateGroupFeishuChatShrinkRequest) GetOperatingObjectName() *string {
	return s.OperatingObjectName
}

func (s *CreateGroupFeishuChatShrinkRequest) GetSourceTags() *string {
	return s.SourceTags
}

func (s *CreateGroupFeishuChatShrinkRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreateGroupFeishuChatShrinkRequest) GetUpdateFrequencyShrink() *string {
	return s.UpdateFrequencyShrink
}

func (s *CreateGroupFeishuChatShrinkRequest) SetChatId(v string) *CreateGroupFeishuChatShrinkRequest {
	s.ChatId = &v
	return s
}

func (s *CreateGroupFeishuChatShrinkRequest) SetDescription(v string) *CreateGroupFeishuChatShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateGroupFeishuChatShrinkRequest) SetDirectoryId(v string) *CreateGroupFeishuChatShrinkRequest {
	s.DirectoryId = &v
	return s
}

func (s *CreateGroupFeishuChatShrinkRequest) SetGroupId(v string) *CreateGroupFeishuChatShrinkRequest {
	s.GroupId = &v
	return s
}

func (s *CreateGroupFeishuChatShrinkRequest) SetHistoryStartTime(v string) *CreateGroupFeishuChatShrinkRequest {
	s.HistoryStartTime = &v
	return s
}

func (s *CreateGroupFeishuChatShrinkRequest) SetNotes(v string) *CreateGroupFeishuChatShrinkRequest {
	s.Notes = &v
	return s
}

func (s *CreateGroupFeishuChatShrinkRequest) SetOperatingObjectName(v string) *CreateGroupFeishuChatShrinkRequest {
	s.OperatingObjectName = &v
	return s
}

func (s *CreateGroupFeishuChatShrinkRequest) SetSourceTags(v string) *CreateGroupFeishuChatShrinkRequest {
	s.SourceTags = &v
	return s
}

func (s *CreateGroupFeishuChatShrinkRequest) SetTenantId(v string) *CreateGroupFeishuChatShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *CreateGroupFeishuChatShrinkRequest) SetUpdateFrequencyShrink(v string) *CreateGroupFeishuChatShrinkRequest {
	s.UpdateFrequencyShrink = &v
	return s
}

func (s *CreateGroupFeishuChatShrinkRequest) Validate() error {
	return dara.Validate(s)
}
