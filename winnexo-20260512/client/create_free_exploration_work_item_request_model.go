// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFreeExplorationWorkItemRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentNames(v []*string) *CreateFreeExplorationWorkItemRequest
	GetAgentNames() []*string
	SetCollaborationGroupId(v string) *CreateFreeExplorationWorkItemRequest
	GetCollaborationGroupId() *string
	SetMemberUserIds(v []*int64) *CreateFreeExplorationWorkItemRequest
	GetMemberUserIds() []*int64
	SetTenantId(v string) *CreateFreeExplorationWorkItemRequest
	GetTenantId() *string
	SetTitle(v string) *CreateFreeExplorationWorkItemRequest
	GetTitle() *string
	SetVisibility(v string) *CreateFreeExplorationWorkItemRequest
	GetVisibility() *string
}

type CreateFreeExplorationWorkItemRequest struct {
	// 空间内数字员工技术名称，1–5 个，不可重复
	//
	// This parameter is required.
	//
	// example:
	//
	// ["sales-agent"]
	AgentNames []*string `json:"agentNames,omitempty" xml:"agentNames,omitempty" type:"Repeated"`
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
	MemberUserIds []*int64 `json:"memberUserIds,omitempty" xml:"memberUserIds,omitempty" type:"Repeated"`
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

func (s CreateFreeExplorationWorkItemRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateFreeExplorationWorkItemRequest) GoString() string {
	return s.String()
}

func (s *CreateFreeExplorationWorkItemRequest) GetAgentNames() []*string {
	return s.AgentNames
}

func (s *CreateFreeExplorationWorkItemRequest) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *CreateFreeExplorationWorkItemRequest) GetMemberUserIds() []*int64 {
	return s.MemberUserIds
}

func (s *CreateFreeExplorationWorkItemRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreateFreeExplorationWorkItemRequest) GetTitle() *string {
	return s.Title
}

func (s *CreateFreeExplorationWorkItemRequest) GetVisibility() *string {
	return s.Visibility
}

func (s *CreateFreeExplorationWorkItemRequest) SetAgentNames(v []*string) *CreateFreeExplorationWorkItemRequest {
	s.AgentNames = v
	return s
}

func (s *CreateFreeExplorationWorkItemRequest) SetCollaborationGroupId(v string) *CreateFreeExplorationWorkItemRequest {
	s.CollaborationGroupId = &v
	return s
}

func (s *CreateFreeExplorationWorkItemRequest) SetMemberUserIds(v []*int64) *CreateFreeExplorationWorkItemRequest {
	s.MemberUserIds = v
	return s
}

func (s *CreateFreeExplorationWorkItemRequest) SetTenantId(v string) *CreateFreeExplorationWorkItemRequest {
	s.TenantId = &v
	return s
}

func (s *CreateFreeExplorationWorkItemRequest) SetTitle(v string) *CreateFreeExplorationWorkItemRequest {
	s.Title = &v
	return s
}

func (s *CreateFreeExplorationWorkItemRequest) SetVisibility(v string) *CreateFreeExplorationWorkItemRequest {
	s.Visibility = &v
	return s
}

func (s *CreateFreeExplorationWorkItemRequest) Validate() error {
	return dara.Validate(s)
}
