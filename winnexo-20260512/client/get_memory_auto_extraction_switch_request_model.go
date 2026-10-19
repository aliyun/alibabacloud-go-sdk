// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMemoryAutoExtractionSwitchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v string) *GetMemoryAutoExtractionSwitchRequest
	GetGroupId() *string
	SetOwnerType(v string) *GetMemoryAutoExtractionSwitchRequest
	GetOwnerType() *string
	SetTenantId(v string) *GetMemoryAutoExtractionSwitchRequest
	GetTenantId() *string
}

type GetMemoryAutoExtractionSwitchRequest struct {
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

func (s GetMemoryAutoExtractionSwitchRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMemoryAutoExtractionSwitchRequest) GoString() string {
	return s.String()
}

func (s *GetMemoryAutoExtractionSwitchRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *GetMemoryAutoExtractionSwitchRequest) GetOwnerType() *string {
	return s.OwnerType
}

func (s *GetMemoryAutoExtractionSwitchRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *GetMemoryAutoExtractionSwitchRequest) SetGroupId(v string) *GetMemoryAutoExtractionSwitchRequest {
	s.GroupId = &v
	return s
}

func (s *GetMemoryAutoExtractionSwitchRequest) SetOwnerType(v string) *GetMemoryAutoExtractionSwitchRequest {
	s.OwnerType = &v
	return s
}

func (s *GetMemoryAutoExtractionSwitchRequest) SetTenantId(v string) *GetMemoryAutoExtractionSwitchRequest {
	s.TenantId = &v
	return s
}

func (s *GetMemoryAutoExtractionSwitchRequest) Validate() error {
	return dara.Validate(s)
}
