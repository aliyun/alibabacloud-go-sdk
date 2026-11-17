// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFreeExplorationWorkItemResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAgentNames(v []*string) *CreateFreeExplorationWorkItemResponseBody
	GetAgentNames() []*string
	SetCode(v string) *CreateFreeExplorationWorkItemResponseBody
	GetCode() *string
	SetCollaborationGroupId(v string) *CreateFreeExplorationWorkItemResponseBody
	GetCollaborationGroupId() *string
	SetConversationId(v string) *CreateFreeExplorationWorkItemResponseBody
	GetConversationId() *string
	SetCreatedByUserId(v int64) *CreateFreeExplorationWorkItemResponseBody
	GetCreatedByUserId() *int64
	SetItemType(v string) *CreateFreeExplorationWorkItemResponseBody
	GetItemType() *string
	SetMemberUserIds(v []*int64) *CreateFreeExplorationWorkItemResponseBody
	GetMemberUserIds() []*int64
	SetMessage(v string) *CreateFreeExplorationWorkItemResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateFreeExplorationWorkItemResponseBody
	GetRequestId() *string
	SetTitle(v string) *CreateFreeExplorationWorkItemResponseBody
	GetTitle() *string
	SetVisibility(v string) *CreateFreeExplorationWorkItemResponseBody
	GetVisibility() *string
}

type CreateFreeExplorationWorkItemResponseBody struct {
	// The technical names of the selected digital workers.
	//
	// example:
	//
	// ["sales-agent"]
	AgentNames []*string `json:"agentNames,omitempty" xml:"agentNames,omitempty" type:"Repeated"`
	// The business status code. A value of 200 indicates success.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The ID of the collaboration workspace to which the work item belongs.
	//
	// example:
	//
	// cg_9401
	CollaborationGroupId *string `json:"collaborationGroupId,omitempty" xml:"collaborationGroupId,omitempty"`
	// The actual group topic ID, which can be associated in queryWorkItemFeed. It cannot be used as a workItemId in Workflow.
	//
	// example:
	//
	// conv_20261105_9f2c
	ConversationId *string `json:"conversationId,omitempty" xml:"conversationId,omitempty"`
	// The user ID of the creator on the platform.
	//
	// example:
	//
	// 10001
	CreatedByUserId *int64 `json:"createdByUserId,omitempty" xml:"createdByUserId,omitempty"`
	// The work item mode. The value is fixed to FREE_EXPLORATION.
	//
	// example:
	//
	// FREE_EXPLORATION
	ItemType *string `json:"itemType,omitempty" xml:"itemType,omitempty"`
	// The actual members of the topic, including the creator. For public topics, members of the entire workspace are not enumerated.
	//
	// example:
	//
	// [43]
	MemberUserIds []*int64 `json:"memberUserIds,omitempty" xml:"memberUserIds,omitempty" type:"Repeated"`
	// The error description.
	//
	// example:
	//
	// The requested resource does not exist
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The title of the work item.
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

func (s CreateFreeExplorationWorkItemResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateFreeExplorationWorkItemResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFreeExplorationWorkItemResponseBody) GetAgentNames() []*string {
	return s.AgentNames
}

func (s *CreateFreeExplorationWorkItemResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateFreeExplorationWorkItemResponseBody) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *CreateFreeExplorationWorkItemResponseBody) GetConversationId() *string {
	return s.ConversationId
}

func (s *CreateFreeExplorationWorkItemResponseBody) GetCreatedByUserId() *int64 {
	return s.CreatedByUserId
}

func (s *CreateFreeExplorationWorkItemResponseBody) GetItemType() *string {
	return s.ItemType
}

func (s *CreateFreeExplorationWorkItemResponseBody) GetMemberUserIds() []*int64 {
	return s.MemberUserIds
}

func (s *CreateFreeExplorationWorkItemResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateFreeExplorationWorkItemResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateFreeExplorationWorkItemResponseBody) GetTitle() *string {
	return s.Title
}

func (s *CreateFreeExplorationWorkItemResponseBody) GetVisibility() *string {
	return s.Visibility
}

func (s *CreateFreeExplorationWorkItemResponseBody) SetAgentNames(v []*string) *CreateFreeExplorationWorkItemResponseBody {
	s.AgentNames = v
	return s
}

func (s *CreateFreeExplorationWorkItemResponseBody) SetCode(v string) *CreateFreeExplorationWorkItemResponseBody {
	s.Code = &v
	return s
}

func (s *CreateFreeExplorationWorkItemResponseBody) SetCollaborationGroupId(v string) *CreateFreeExplorationWorkItemResponseBody {
	s.CollaborationGroupId = &v
	return s
}

func (s *CreateFreeExplorationWorkItemResponseBody) SetConversationId(v string) *CreateFreeExplorationWorkItemResponseBody {
	s.ConversationId = &v
	return s
}

func (s *CreateFreeExplorationWorkItemResponseBody) SetCreatedByUserId(v int64) *CreateFreeExplorationWorkItemResponseBody {
	s.CreatedByUserId = &v
	return s
}

func (s *CreateFreeExplorationWorkItemResponseBody) SetItemType(v string) *CreateFreeExplorationWorkItemResponseBody {
	s.ItemType = &v
	return s
}

func (s *CreateFreeExplorationWorkItemResponseBody) SetMemberUserIds(v []*int64) *CreateFreeExplorationWorkItemResponseBody {
	s.MemberUserIds = v
	return s
}

func (s *CreateFreeExplorationWorkItemResponseBody) SetMessage(v string) *CreateFreeExplorationWorkItemResponseBody {
	s.Message = &v
	return s
}

func (s *CreateFreeExplorationWorkItemResponseBody) SetRequestId(v string) *CreateFreeExplorationWorkItemResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateFreeExplorationWorkItemResponseBody) SetTitle(v string) *CreateFreeExplorationWorkItemResponseBody {
	s.Title = &v
	return s
}

func (s *CreateFreeExplorationWorkItemResponseBody) SetVisibility(v string) *CreateFreeExplorationWorkItemResponseBody {
	s.Visibility = &v
	return s
}

func (s *CreateFreeExplorationWorkItemResponseBody) Validate() error {
	return dara.Validate(s)
}
