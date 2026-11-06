// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFreeExplorationWorkItemShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentNamesShrink(v string) *CreateFreeExplorationWorkItemShrinkRequest
	GetAgentNamesShrink() *string
	SetCollaborationGroupId(v string) *CreateFreeExplorationWorkItemShrinkRequest
	GetCollaborationGroupId() *string
	SetMemberUserIdsShrink(v string) *CreateFreeExplorationWorkItemShrinkRequest
	GetMemberUserIdsShrink() *string
	SetTenantId(v string) *CreateFreeExplorationWorkItemShrinkRequest
	GetTenantId() *string
	SetTitle(v string) *CreateFreeExplorationWorkItemShrinkRequest
	GetTitle() *string
	SetVisibility(v string) *CreateFreeExplorationWorkItemShrinkRequest
	GetVisibility() *string
}

type CreateFreeExplorationWorkItemShrinkRequest struct {
	// 空间内数字员工技术名称，1–5 个，不可重复
	//
	// This parameter is required.
	//
	// example:
	//
	// ["sales-agent"]
	AgentNamesShrink *string `json:"agentNames,omitempty" xml:"agentNames,omitempty"`
	// 协作空间 ID；调用者须为有效成员
	//
	// This parameter is required.
	//
	// example:
	//
	// cg_123
	CollaborationGroupId *string `json:"collaborationGroupId,omitempty" xml:"collaborationGroupId,omitempty"`
	// 协作档至少一个创建人以外的有效空间成员；最多 1000 人；私密和公开档忽略；创建人始终加入
	//
	// example:
	//
	// [43]
	MemberUserIdsShrink *string `json:"memberUserIds,omitempty" xml:"memberUserIds,omitempty"`
	// 租户ID，公共参数；缺省时使用调用方默认租户
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// 自由探索工作项标题，去首尾空格后 1–200 字符
	//
	// This parameter is required.
	//
	// example:
	//
	// 调查销售机会
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// 可见范围：PRIVATE 私密；COLLABORATIVE 指定成员协作；PUBLIC 空间内公开
	//
	// example:
	//
	// PRIVATE
	Visibility *string `json:"visibility,omitempty" xml:"visibility,omitempty"`
}

func (s CreateFreeExplorationWorkItemShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateFreeExplorationWorkItemShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateFreeExplorationWorkItemShrinkRequest) GetAgentNamesShrink() *string {
	return s.AgentNamesShrink
}

func (s *CreateFreeExplorationWorkItemShrinkRequest) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *CreateFreeExplorationWorkItemShrinkRequest) GetMemberUserIdsShrink() *string {
	return s.MemberUserIdsShrink
}

func (s *CreateFreeExplorationWorkItemShrinkRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreateFreeExplorationWorkItemShrinkRequest) GetTitle() *string {
	return s.Title
}

func (s *CreateFreeExplorationWorkItemShrinkRequest) GetVisibility() *string {
	return s.Visibility
}

func (s *CreateFreeExplorationWorkItemShrinkRequest) SetAgentNamesShrink(v string) *CreateFreeExplorationWorkItemShrinkRequest {
	s.AgentNamesShrink = &v
	return s
}

func (s *CreateFreeExplorationWorkItemShrinkRequest) SetCollaborationGroupId(v string) *CreateFreeExplorationWorkItemShrinkRequest {
	s.CollaborationGroupId = &v
	return s
}

func (s *CreateFreeExplorationWorkItemShrinkRequest) SetMemberUserIdsShrink(v string) *CreateFreeExplorationWorkItemShrinkRequest {
	s.MemberUserIdsShrink = &v
	return s
}

func (s *CreateFreeExplorationWorkItemShrinkRequest) SetTenantId(v string) *CreateFreeExplorationWorkItemShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *CreateFreeExplorationWorkItemShrinkRequest) SetTitle(v string) *CreateFreeExplorationWorkItemShrinkRequest {
	s.Title = &v
	return s
}

func (s *CreateFreeExplorationWorkItemShrinkRequest) SetVisibility(v string) *CreateFreeExplorationWorkItemShrinkRequest {
	s.Visibility = &v
	return s
}

func (s *CreateFreeExplorationWorkItemShrinkRequest) Validate() error {
	return dara.Validate(s)
}
