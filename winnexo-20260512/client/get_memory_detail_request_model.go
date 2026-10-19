// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMemoryDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v string) *GetMemoryDetailRequest
	GetGroupId() *string
	SetOwnerType(v string) *GetMemoryDetailRequest
	GetOwnerType() *string
	SetTenantId(v string) *GetMemoryDetailRequest
	GetTenantId() *string
}

type GetMemoryDetailRequest struct {
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

func (s GetMemoryDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMemoryDetailRequest) GoString() string {
	return s.String()
}

func (s *GetMemoryDetailRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *GetMemoryDetailRequest) GetOwnerType() *string {
	return s.OwnerType
}

func (s *GetMemoryDetailRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *GetMemoryDetailRequest) SetGroupId(v string) *GetMemoryDetailRequest {
	s.GroupId = &v
	return s
}

func (s *GetMemoryDetailRequest) SetOwnerType(v string) *GetMemoryDetailRequest {
	s.OwnerType = &v
	return s
}

func (s *GetMemoryDetailRequest) SetTenantId(v string) *GetMemoryDetailRequest {
	s.TenantId = &v
	return s
}

func (s *GetMemoryDetailRequest) Validate() error {
	return dara.Validate(s)
}
