// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReplaceManualMemoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *ReplaceManualMemoryRequest
	GetContent() *string
	SetGroupId(v string) *ReplaceManualMemoryRequest
	GetGroupId() *string
	SetOwnerType(v string) *ReplaceManualMemoryRequest
	GetOwnerType() *string
	SetTenantId(v string) *ReplaceManualMemoryRequest
	GetTenantId() *string
}

type ReplaceManualMemoryRequest struct {
	// 手动记忆 profile.md 的完整新内容，将覆盖原内容；传空字符串表示清空，字段必须传入
	//
	// This parameter is required.
	//
	// example:
	//
	// 喜欢简洁的回答。
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
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

func (s ReplaceManualMemoryRequest) String() string {
	return dara.Prettify(s)
}

func (s ReplaceManualMemoryRequest) GoString() string {
	return s.String()
}

func (s *ReplaceManualMemoryRequest) GetContent() *string {
	return s.Content
}

func (s *ReplaceManualMemoryRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *ReplaceManualMemoryRequest) GetOwnerType() *string {
	return s.OwnerType
}

func (s *ReplaceManualMemoryRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ReplaceManualMemoryRequest) SetContent(v string) *ReplaceManualMemoryRequest {
	s.Content = &v
	return s
}

func (s *ReplaceManualMemoryRequest) SetGroupId(v string) *ReplaceManualMemoryRequest {
	s.GroupId = &v
	return s
}

func (s *ReplaceManualMemoryRequest) SetOwnerType(v string) *ReplaceManualMemoryRequest {
	s.OwnerType = &v
	return s
}

func (s *ReplaceManualMemoryRequest) SetTenantId(v string) *ReplaceManualMemoryRequest {
	s.TenantId = &v
	return s
}

func (s *ReplaceManualMemoryRequest) Validate() error {
	return dara.Validate(s)
}
