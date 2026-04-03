// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIMBotInfo interface {
	dara.Model
	String() string
	GoString() string
	SetAgentRuntimeId(v string) *IMBotInfo
	GetAgentRuntimeId() *string
	SetBotBizId(v string) *IMBotInfo
	GetBotBizId() *string
	SetBotBizType(v string) *IMBotInfo
	GetBotBizType() *string
	SetBotMode(v string) *IMBotInfo
	GetBotMode() *string
	SetBotName(v string) *IMBotInfo
	GetBotName() *string
	SetCreatedAt(v string) *IMBotInfo
	GetCreatedAt() *string
	SetCurrentInstances(v int64) *IMBotInfo
	GetCurrentInstances() *int64
	SetDescription(v string) *IMBotInfo
	GetDescription() *string
	SetId(v string) *IMBotInfo
	GetId() *string
	SetImChannelServerResourceName(v string) *IMBotInfo
	GetImChannelServerResourceName() *string
	SetLastMessageTime(v string) *IMBotInfo
	GetLastMessageTime() *string
	SetMetadata(v *IMBotMetadata) *IMBotInfo
	GetMetadata() *IMBotMetadata
	SetStatus(v string) *IMBotInfo
	GetStatus() *string
	SetTenantId(v string) *IMBotInfo
	GetTenantId() *string
	SetUpdatedAt(v string) *IMBotInfo
	GetUpdatedAt() *string
}

type IMBotInfo struct {
	// if can be null:
	// true
	AgentRuntimeId *string `json:"agentRuntimeId,omitempty" xml:"agentRuntimeId,omitempty"`
	BotBizId       *string `json:"botBizId,omitempty" xml:"botBizId,omitempty"`
	// dingtalk、wecom、feishu、custom
	BotBizType *string `json:"botBizType,omitempty" xml:"botBizType,omitempty"`
	// standard、custom
	BotMode   *string `json:"botMode,omitempty" xml:"botMode,omitempty"`
	BotName   *string `json:"botName,omitempty" xml:"botName,omitempty"`
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// 可选；来自账号元数据 scaling.currentInstances，同一租户下各 Bot 响应中值相同
	//
	// if can be null:
	// true
	CurrentInstances *int64 `json:"currentInstances,omitempty" xml:"currentInstances,omitempty"`
	// Bot 描述信息
	//
	// if can be null:
	// true
	Description                 *string `json:"description,omitempty" xml:"description,omitempty"`
	Id                          *string `json:"id,omitempty" xml:"id,omitempty"`
	ImChannelServerResourceName *string `json:"imChannelServerResourceName,omitempty" xml:"imChannelServerResourceName,omitempty"`
	// 可选
	//
	// if can be null:
	// true
	LastMessageTime *string `json:"lastMessageTime,omitempty" xml:"lastMessageTime,omitempty"`
	// if can be null:
	// true
	Metadata  *IMBotMetadata `json:"metadata,omitempty" xml:"metadata,omitempty"`
	Status    *string        `json:"status,omitempty" xml:"status,omitempty"`
	TenantId  *string        `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	UpdatedAt *string        `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
}

func (s IMBotInfo) String() string {
	return dara.Prettify(s)
}

func (s IMBotInfo) GoString() string {
	return s.String()
}

func (s *IMBotInfo) GetAgentRuntimeId() *string {
	return s.AgentRuntimeId
}

func (s *IMBotInfo) GetBotBizId() *string {
	return s.BotBizId
}

func (s *IMBotInfo) GetBotBizType() *string {
	return s.BotBizType
}

func (s *IMBotInfo) GetBotMode() *string {
	return s.BotMode
}

func (s *IMBotInfo) GetBotName() *string {
	return s.BotName
}

func (s *IMBotInfo) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *IMBotInfo) GetCurrentInstances() *int64 {
	return s.CurrentInstances
}

func (s *IMBotInfo) GetDescription() *string {
	return s.Description
}

func (s *IMBotInfo) GetId() *string {
	return s.Id
}

func (s *IMBotInfo) GetImChannelServerResourceName() *string {
	return s.ImChannelServerResourceName
}

func (s *IMBotInfo) GetLastMessageTime() *string {
	return s.LastMessageTime
}

func (s *IMBotInfo) GetMetadata() *IMBotMetadata {
	return s.Metadata
}

func (s *IMBotInfo) GetStatus() *string {
	return s.Status
}

func (s *IMBotInfo) GetTenantId() *string {
	return s.TenantId
}

func (s *IMBotInfo) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *IMBotInfo) SetAgentRuntimeId(v string) *IMBotInfo {
	s.AgentRuntimeId = &v
	return s
}

func (s *IMBotInfo) SetBotBizId(v string) *IMBotInfo {
	s.BotBizId = &v
	return s
}

func (s *IMBotInfo) SetBotBizType(v string) *IMBotInfo {
	s.BotBizType = &v
	return s
}

func (s *IMBotInfo) SetBotMode(v string) *IMBotInfo {
	s.BotMode = &v
	return s
}

func (s *IMBotInfo) SetBotName(v string) *IMBotInfo {
	s.BotName = &v
	return s
}

func (s *IMBotInfo) SetCreatedAt(v string) *IMBotInfo {
	s.CreatedAt = &v
	return s
}

func (s *IMBotInfo) SetCurrentInstances(v int64) *IMBotInfo {
	s.CurrentInstances = &v
	return s
}

func (s *IMBotInfo) SetDescription(v string) *IMBotInfo {
	s.Description = &v
	return s
}

func (s *IMBotInfo) SetId(v string) *IMBotInfo {
	s.Id = &v
	return s
}

func (s *IMBotInfo) SetImChannelServerResourceName(v string) *IMBotInfo {
	s.ImChannelServerResourceName = &v
	return s
}

func (s *IMBotInfo) SetLastMessageTime(v string) *IMBotInfo {
	s.LastMessageTime = &v
	return s
}

func (s *IMBotInfo) SetMetadata(v *IMBotMetadata) *IMBotInfo {
	s.Metadata = v
	return s
}

func (s *IMBotInfo) SetStatus(v string) *IMBotInfo {
	s.Status = &v
	return s
}

func (s *IMBotInfo) SetTenantId(v string) *IMBotInfo {
	s.TenantId = &v
	return s
}

func (s *IMBotInfo) SetUpdatedAt(v string) *IMBotInfo {
	s.UpdatedAt = &v
	return s
}

func (s *IMBotInfo) Validate() error {
	if s.Metadata != nil {
		if err := s.Metadata.Validate(); err != nil {
			return err
		}
	}
	return nil
}
