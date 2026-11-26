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
	SetAppToken(v string) *CreateFreeExplorationWorkItemRequest
	GetAppToken() *string
	SetCollaborationGroupId(v string) *CreateFreeExplorationWorkItemRequest
	GetCollaborationGroupId() *string
	SetMemberUserIds(v []*int64) *CreateFreeExplorationWorkItemRequest
	GetMemberUserIds() []*int64
	SetOpsUserCode(v string) *CreateFreeExplorationWorkItemRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *CreateFreeExplorationWorkItemRequest
	GetOpsUserId() *string
	SetTenantId(v string) *CreateFreeExplorationWorkItemRequest
	GetTenantId() *string
	SetTitle(v string) *CreateFreeExplorationWorkItemRequest
	GetTitle() *string
	SetVisibility(v string) *CreateFreeExplorationWorkItemRequest
	GetVisibility() *string
}

type CreateFreeExplorationWorkItemRequest struct {
	// The technical names of the selected digital workers.
	//
	// This parameter is required.
	//
	// example:
	//
	// ["sales-agent"]
	AgentNames []*string `json:"agentNames,omitempty" xml:"agentNames,omitempty" type:"Repeated"`
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
	MemberUserIds []*int64 `json:"memberUserIds,omitempty" xml:"memberUserIds,omitempty" type:"Repeated"`
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

func (s CreateFreeExplorationWorkItemRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateFreeExplorationWorkItemRequest) GoString() string {
	return s.String()
}

func (s *CreateFreeExplorationWorkItemRequest) GetAgentNames() []*string {
	return s.AgentNames
}

func (s *CreateFreeExplorationWorkItemRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *CreateFreeExplorationWorkItemRequest) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *CreateFreeExplorationWorkItemRequest) GetMemberUserIds() []*int64 {
	return s.MemberUserIds
}

func (s *CreateFreeExplorationWorkItemRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *CreateFreeExplorationWorkItemRequest) GetOpsUserId() *string {
	return s.OpsUserId
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

func (s *CreateFreeExplorationWorkItemRequest) SetAppToken(v string) *CreateFreeExplorationWorkItemRequest {
	s.AppToken = &v
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

func (s *CreateFreeExplorationWorkItemRequest) SetOpsUserCode(v string) *CreateFreeExplorationWorkItemRequest {
	s.OpsUserCode = &v
	return s
}

func (s *CreateFreeExplorationWorkItemRequest) SetOpsUserId(v string) *CreateFreeExplorationWorkItemRequest {
	s.OpsUserId = &v
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
