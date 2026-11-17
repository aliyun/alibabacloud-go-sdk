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
	// The technical names of the selected digital workers.
	//
	// This parameter is required.
	//
	// example:
	//
	// ["sales-agent"]
	AgentNames []*string `json:"agentNames,omitempty" xml:"agentNames,omitempty" type:"Repeated"`
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
