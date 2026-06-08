// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAskLumaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentName(v string) *AskLumaRequest
	GetAgentName() *string
	SetConversationId(v string) *AskLumaRequest
	GetConversationId() *string
	SetMaxRows(v int32) *AskLumaRequest
	GetMaxRows() *int32
	SetQuestion(v string) *AskLumaRequest
	GetQuestion() *string
}

type AskLumaRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// demo-luma-agent
	AgentName *string `json:"AgentName,omitempty" xml:"AgentName,omitempty"`
	// example:
	//
	// conv_1474xxx32_593b9d08-9
	ConversationId *string `json:"ConversationId,omitempty" xml:"ConversationId,omitempty"`
	// example:
	//
	// 100
	MaxRows *int32 `json:"MaxRows,omitempty" xml:"MaxRows,omitempty"`
	// This parameter is required.
	Question *string `json:"Question,omitempty" xml:"Question,omitempty"`
}

func (s AskLumaRequest) String() string {
	return dara.Prettify(s)
}

func (s AskLumaRequest) GoString() string {
	return s.String()
}

func (s *AskLumaRequest) GetAgentName() *string {
	return s.AgentName
}

func (s *AskLumaRequest) GetConversationId() *string {
	return s.ConversationId
}

func (s *AskLumaRequest) GetMaxRows() *int32 {
	return s.MaxRows
}

func (s *AskLumaRequest) GetQuestion() *string {
	return s.Question
}

func (s *AskLumaRequest) SetAgentName(v string) *AskLumaRequest {
	s.AgentName = &v
	return s
}

func (s *AskLumaRequest) SetConversationId(v string) *AskLumaRequest {
	s.ConversationId = &v
	return s
}

func (s *AskLumaRequest) SetMaxRows(v int32) *AskLumaRequest {
	s.MaxRows = &v
	return s
}

func (s *AskLumaRequest) SetQuestion(v string) *AskLumaRequest {
	s.Question = &v
	return s
}

func (s *AskLumaRequest) Validate() error {
	return dara.Validate(s)
}
