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
	// The technical names of the selected digital workers.
	//
	// This parameter is required.
	//
	// example:
	//
	// ["sales-agent"]
	AgentNamesShrink *string `json:"agentNames,omitempty" xml:"agentNames,omitempty"`
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
