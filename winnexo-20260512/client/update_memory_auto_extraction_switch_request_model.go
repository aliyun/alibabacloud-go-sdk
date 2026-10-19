// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMemoryAutoExtractionSwitchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClearExisting(v bool) *UpdateMemoryAutoExtractionSwitchRequest
	GetClearExisting() *bool
	SetEnabled(v bool) *UpdateMemoryAutoExtractionSwitchRequest
	GetEnabled() *bool
	SetGroupId(v string) *UpdateMemoryAutoExtractionSwitchRequest
	GetGroupId() *string
	SetOwnerType(v string) *UpdateMemoryAutoExtractionSwitchRequest
	GetOwnerType() *string
	SetTenantId(v string) *UpdateMemoryAutoExtractionSwitchRequest
	GetTenantId() *string
}

type UpdateMemoryAutoExtractionSwitchRequest struct {
	// 是否同时清空当前记忆范围内的全部自动记忆；开启或关闭时均可使用，不影响手动记忆；不传时为 false
	//
	// example:
	//
	// true
	ClearExisting *bool `json:"clearExisting,omitempty" xml:"clearExisting,omitempty"`
	// 目标开关状态；true 为开启自动记忆拾取，false 为关闭
	//
	// This parameter is required.
	//
	// example:
	//
	// false
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// 写作协作空间 ID；ownerType 为 group 时必填，且当前鉴权用户须为该空间的活跃成员；ownerType 为 user 时不传
	//
	// example:
	//
	// cg_12345
	GroupId *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// 记忆归属范围；user 为当前鉴权用户的个人记忆，group 为指定写作协作空间的记忆；不传时默认为 user
	//
	// example:
	//
	// group
	OwnerType *string `json:"ownerType,omitempty" xml:"ownerType,omitempty"`
	// 租户 ID，查询参数；不传时使用调用方的默认租户
	//
	// example:
	//
	// 10001
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s UpdateMemoryAutoExtractionSwitchRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMemoryAutoExtractionSwitchRequest) GoString() string {
	return s.String()
}

func (s *UpdateMemoryAutoExtractionSwitchRequest) GetClearExisting() *bool {
	return s.ClearExisting
}

func (s *UpdateMemoryAutoExtractionSwitchRequest) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpdateMemoryAutoExtractionSwitchRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *UpdateMemoryAutoExtractionSwitchRequest) GetOwnerType() *string {
	return s.OwnerType
}

func (s *UpdateMemoryAutoExtractionSwitchRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *UpdateMemoryAutoExtractionSwitchRequest) SetClearExisting(v bool) *UpdateMemoryAutoExtractionSwitchRequest {
	s.ClearExisting = &v
	return s
}

func (s *UpdateMemoryAutoExtractionSwitchRequest) SetEnabled(v bool) *UpdateMemoryAutoExtractionSwitchRequest {
	s.Enabled = &v
	return s
}

func (s *UpdateMemoryAutoExtractionSwitchRequest) SetGroupId(v string) *UpdateMemoryAutoExtractionSwitchRequest {
	s.GroupId = &v
	return s
}

func (s *UpdateMemoryAutoExtractionSwitchRequest) SetOwnerType(v string) *UpdateMemoryAutoExtractionSwitchRequest {
	s.OwnerType = &v
	return s
}

func (s *UpdateMemoryAutoExtractionSwitchRequest) SetTenantId(v string) *UpdateMemoryAutoExtractionSwitchRequest {
	s.TenantId = &v
	return s
}

func (s *UpdateMemoryAutoExtractionSwitchRequest) Validate() error {
	return dara.Validate(s)
}
