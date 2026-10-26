// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAgentConversationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConversationId(v string) *CreateAgentConversationResponseBody
	GetConversationId() *string
	SetCreatedAt(v string) *CreateAgentConversationResponseBody
	GetCreatedAt() *string
	SetCurrentRunStatus(v string) *CreateAgentConversationResponseBody
	GetCurrentRunStatus() *string
	SetId(v string) *CreateAgentConversationResponseBody
	GetId() *string
	SetRunId(v string) *CreateAgentConversationResponseBody
	GetRunId() *string
	SetScenarioId(v string) *CreateAgentConversationResponseBody
	GetScenarioId() *string
	SetTitle(v string) *CreateAgentConversationResponseBody
	GetTitle() *string
	SetUpdatedAt(v string) *CreateAgentConversationResponseBody
	GetUpdatedAt() *string
}

type CreateAgentConversationResponseBody struct {
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
	// running
	CurrentRunStatus *string `json:"currentRunStatus,omitempty" xml:"currentRunStatus,omitempty"`
	// example:
	//
	// 0198c8a91f4d7f01a3c21e97
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 0198c8a91f4d7f01a3c21e98
	RunId *string `json:"runId,omitempty" xml:"runId,omitempty"`
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
	// 2026-08-19T12:00:00.000Z
	UpdatedAt *string `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
}

func (s CreateAgentConversationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAgentConversationResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAgentConversationResponseBody) GetConversationId() *string {
	return s.ConversationId
}

func (s *CreateAgentConversationResponseBody) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *CreateAgentConversationResponseBody) GetCurrentRunStatus() *string {
	return s.CurrentRunStatus
}

func (s *CreateAgentConversationResponseBody) GetId() *string {
	return s.Id
}

func (s *CreateAgentConversationResponseBody) GetRunId() *string {
	return s.RunId
}

func (s *CreateAgentConversationResponseBody) GetScenarioId() *string {
	return s.ScenarioId
}

func (s *CreateAgentConversationResponseBody) GetTitle() *string {
	return s.Title
}

func (s *CreateAgentConversationResponseBody) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *CreateAgentConversationResponseBody) SetConversationId(v string) *CreateAgentConversationResponseBody {
	s.ConversationId = &v
	return s
}

func (s *CreateAgentConversationResponseBody) SetCreatedAt(v string) *CreateAgentConversationResponseBody {
	s.CreatedAt = &v
	return s
}

func (s *CreateAgentConversationResponseBody) SetCurrentRunStatus(v string) *CreateAgentConversationResponseBody {
	s.CurrentRunStatus = &v
	return s
}

func (s *CreateAgentConversationResponseBody) SetId(v string) *CreateAgentConversationResponseBody {
	s.Id = &v
	return s
}

func (s *CreateAgentConversationResponseBody) SetRunId(v string) *CreateAgentConversationResponseBody {
	s.RunId = &v
	return s
}

func (s *CreateAgentConversationResponseBody) SetScenarioId(v string) *CreateAgentConversationResponseBody {
	s.ScenarioId = &v
	return s
}

func (s *CreateAgentConversationResponseBody) SetTitle(v string) *CreateAgentConversationResponseBody {
	s.Title = &v
	return s
}

func (s *CreateAgentConversationResponseBody) SetUpdatedAt(v string) *CreateAgentConversationResponseBody {
	s.UpdatedAt = &v
	return s
}

func (s *CreateAgentConversationResponseBody) Validate() error {
	return dara.Validate(s)
}
