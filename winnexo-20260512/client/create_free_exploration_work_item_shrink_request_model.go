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
	SetAppToken(v string) *CreateFreeExplorationWorkItemShrinkRequest
	GetAppToken() *string
	SetCollaborationGroupId(v string) *CreateFreeExplorationWorkItemShrinkRequest
	GetCollaborationGroupId() *string
	SetMemberUserIdsShrink(v string) *CreateFreeExplorationWorkItemShrinkRequest
	GetMemberUserIdsShrink() *string
	SetOpsUserCode(v string) *CreateFreeExplorationWorkItemShrinkRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *CreateFreeExplorationWorkItemShrinkRequest
	GetOpsUserId() *string
	SetTenantId(v string) *CreateFreeExplorationWorkItemShrinkRequest
	GetTenantId() *string
	SetTitle(v string) *CreateFreeExplorationWorkItemShrinkRequest
	GetTitle() *string
	SetVisibility(v string) *CreateFreeExplorationWorkItemShrinkRequest
	GetVisibility() *string
}

type CreateFreeExplorationWorkItemShrinkRequest struct {
	// The technical names of the selected digital workers.
	//
	// This parameter is required.
	//
	// example:
	//
	// ["sales-agent"]
	AgentNamesShrink *string `json:"agentNames,omitempty" xml:"agentNames,omitempty"`
	// 平台应用 Token（控制台签发）。传入后将以 opsUserId 或 opsUserCode 指定的用户身份执行
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// The ID of the collaboration workspace to which the work item belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cg_123
	CollaborationGroupId *string `json:"collaborationGroupId,omitempty" xml:"collaborationGroupId,omitempty"`
	// The actual members of the topic, including the creator. For public topics, members of the entire workspace are not enumerated.
	//
	// example:
	//
	// [43]
	MemberUserIdsShrink *string `json:"memberUserIds,omitempty" xml:"memberUserIds,omitempty"`
	// 代理用户编码（租户内唯一），仅当应用凭证配置为 source_uid 时使用
	//
	// example:
	//
	// alice
	OpsUserCode *string `json:"opsUserCode,omitempty" xml:"opsUserCode,omitempty"`
	// 代理用户ID（WINNEXO 平台用户ID），仅当应用凭证配置为 winnexo_uid 时使用
	//
	// example:
	//
	// 10001
	OpsUserId *string `json:"opsUserId,omitempty" xml:"opsUserId,omitempty"`
	// The tenant ID. This is a common parameter. If this parameter is not specified, the default tenant of the caller is used.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The title of the work item.
	//
	// This parameter is required.
	//
	// example:
	//
	// Investigate sales opportunities
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// The actual visibility scope. Valid values: PRIVATE, COLLABORATIVE, and PUBLIC.
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

func (s *CreateFreeExplorationWorkItemShrinkRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *CreateFreeExplorationWorkItemShrinkRequest) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *CreateFreeExplorationWorkItemShrinkRequest) GetMemberUserIdsShrink() *string {
	return s.MemberUserIdsShrink
}

func (s *CreateFreeExplorationWorkItemShrinkRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *CreateFreeExplorationWorkItemShrinkRequest) GetOpsUserId() *string {
	return s.OpsUserId
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

func (s *CreateFreeExplorationWorkItemShrinkRequest) SetAppToken(v string) *CreateFreeExplorationWorkItemShrinkRequest {
	s.AppToken = &v
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

func (s *CreateFreeExplorationWorkItemShrinkRequest) SetOpsUserCode(v string) *CreateFreeExplorationWorkItemShrinkRequest {
	s.OpsUserCode = &v
	return s
}

func (s *CreateFreeExplorationWorkItemShrinkRequest) SetOpsUserId(v string) *CreateFreeExplorationWorkItemShrinkRequest {
	s.OpsUserId = &v
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
