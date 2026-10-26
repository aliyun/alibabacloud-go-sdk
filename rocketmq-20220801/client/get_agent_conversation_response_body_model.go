// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentConversationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConversationId(v string) *GetAgentConversationResponseBody
	GetConversationId() *string
	SetCreatedAt(v string) *GetAgentConversationResponseBody
	GetCreatedAt() *string
	SetCurrentRunStatus(v string) *GetAgentConversationResponseBody
	GetCurrentRunStatus() *string
	SetId(v string) *GetAgentConversationResponseBody
	GetId() *string
	SetPendingTool(v interface{}) *GetAgentConversationResponseBody
	GetPendingTool() interface{}
	SetScenarioId(v string) *GetAgentConversationResponseBody
	GetScenarioId() *string
	SetTitle(v string) *GetAgentConversationResponseBody
	GetTitle() *string
	SetUpdatedAt(v string) *GetAgentConversationResponseBody
	GetUpdatedAt() *string
}

type GetAgentConversationResponseBody struct {
	// example:
	//
	// 0198c8a91f4d7f01a3c21e97
	ConversationId *string `json:"conversationId,omitempty" xml:"conversationId,omitempty"`
	// example:
	//
	// 2026-08-19T12:00:00.000Z
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// example:
	//
	// waiting_input
	CurrentRunStatus *string `json:"currentRunStatus,omitempty" xml:"currentRunStatus,omitempty"`
	// example:
	//
	// 0198c8a91f4d7f01a3c21e97
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// {"toolUseId":"tool_01","tool":"AskUserQuestion"}
	PendingTool interface{} `json:"pendingTool,omitempty" xml:"pendingTool,omitempty"`
	// example:
	//
	// diagnosis
	ScenarioId *string `json:"scenarioId,omitempty" xml:"scenarioId,omitempty"`
	// example:
	//
	// 消息堆积排查
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// example:
	//
	// 2026-08-19T12:05:00.000Z
	UpdatedAt *string `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
}

func (s GetAgentConversationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAgentConversationResponseBody) GoString() string {
	return s.String()
}

func (s *GetAgentConversationResponseBody) GetConversationId() *string {
	return s.ConversationId
}

func (s *GetAgentConversationResponseBody) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *GetAgentConversationResponseBody) GetCurrentRunStatus() *string {
	return s.CurrentRunStatus
}

func (s *GetAgentConversationResponseBody) GetId() *string {
	return s.Id
}

func (s *GetAgentConversationResponseBody) GetPendingTool() interface{} {
	return s.PendingTool
}

func (s *GetAgentConversationResponseBody) GetScenarioId() *string {
	return s.ScenarioId
}

func (s *GetAgentConversationResponseBody) GetTitle() *string {
	return s.Title
}

func (s *GetAgentConversationResponseBody) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *GetAgentConversationResponseBody) SetConversationId(v string) *GetAgentConversationResponseBody {
	s.ConversationId = &v
	return s
}

func (s *GetAgentConversationResponseBody) SetCreatedAt(v string) *GetAgentConversationResponseBody {
	s.CreatedAt = &v
	return s
}

func (s *GetAgentConversationResponseBody) SetCurrentRunStatus(v string) *GetAgentConversationResponseBody {
	s.CurrentRunStatus = &v
	return s
}

func (s *GetAgentConversationResponseBody) SetId(v string) *GetAgentConversationResponseBody {
	s.Id = &v
	return s
}

func (s *GetAgentConversationResponseBody) SetPendingTool(v interface{}) *GetAgentConversationResponseBody {
	s.PendingTool = v
	return s
}

func (s *GetAgentConversationResponseBody) SetScenarioId(v string) *GetAgentConversationResponseBody {
	s.ScenarioId = &v
	return s
}

func (s *GetAgentConversationResponseBody) SetTitle(v string) *GetAgentConversationResponseBody {
	s.Title = &v
	return s
}

func (s *GetAgentConversationResponseBody) SetUpdatedAt(v string) *GetAgentConversationResponseBody {
	s.UpdatedAt = &v
	return s
}

func (s *GetAgentConversationResponseBody) Validate() error {
	return dara.Validate(s)
}
