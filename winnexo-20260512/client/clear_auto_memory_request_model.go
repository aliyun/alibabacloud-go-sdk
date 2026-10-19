// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClearAutoMemoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v string) *ClearAutoMemoryRequest
	GetGroupId() *string
	SetOwnerType(v string) *ClearAutoMemoryRequest
	GetOwnerType() *string
	SetTenantId(v string) *ClearAutoMemoryRequest
	GetTenantId() *string
	SetType(v string) *ClearAutoMemoryRequest
	GetType() *string
}

type ClearAutoMemoryRequest struct {
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
	// 要清空的自动记忆类型：user、feedback、reference、project；all 表示清空全部自动记忆，但不清空手动记忆
	//
	// This parameter is required.
	//
	// example:
	//
	// project
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ClearAutoMemoryRequest) String() string {
	return dara.Prettify(s)
}

func (s ClearAutoMemoryRequest) GoString() string {
	return s.String()
}

func (s *ClearAutoMemoryRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *ClearAutoMemoryRequest) GetOwnerType() *string {
	return s.OwnerType
}

func (s *ClearAutoMemoryRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ClearAutoMemoryRequest) GetType() *string {
	return s.Type
}

func (s *ClearAutoMemoryRequest) SetGroupId(v string) *ClearAutoMemoryRequest {
	s.GroupId = &v
	return s
}

func (s *ClearAutoMemoryRequest) SetOwnerType(v string) *ClearAutoMemoryRequest {
	s.OwnerType = &v
	return s
}

func (s *ClearAutoMemoryRequest) SetTenantId(v string) *ClearAutoMemoryRequest {
	s.TenantId = &v
	return s
}

func (s *ClearAutoMemoryRequest) SetType(v string) *ClearAutoMemoryRequest {
	s.Type = &v
	return s
}

func (s *ClearAutoMemoryRequest) Validate() error {
	return dara.Validate(s)
}
