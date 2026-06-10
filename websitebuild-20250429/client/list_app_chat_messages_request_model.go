// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAppChatMessagesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChatId(v string) *ListAppChatMessagesRequest
	GetChatId() *string
	SetConversationId(v string) *ListAppChatMessagesRequest
	GetConversationId() *string
	SetMaxResults(v int32) *ListAppChatMessagesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListAppChatMessagesRequest
	GetNextToken() *string
	SetSectionId(v string) *ListAppChatMessagesRequest
	GetSectionId() *string
}

type ListAppChatMessagesRequest struct {
	// Chat ID
	//
	// example:
	//
	// 20833ba4-d189-4c50-9a44-a6bcbda2c93b
	ChatId *string `json:"ChatId,omitempty" xml:"ChatId,omitempty"`
	// Session ID
	//
	// example:
	//
	// 593fe1a2-d0b4-4fde-a2b0-78ad6a438d41
	ConversationId *string `json:"ConversationId,omitempty" xml:"ConversationId,omitempty"`
	// Number of results per query.
	//
	// Valid range: 10 to 100. Default Value: 20.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// Token indicating the start of the next query. This value is empty if there is no subsequent query.
	//
	// example:
	//
	// FFh3Xqm+JgZ/U9Jyb7wdVr9LWk80Tghn5UZjbcWEVEderBcbVF+Y6PS0i8PpCL4PQZ3e0C9oEH0Asd4tJEuGtkl2WuKdiWZpEwadNydQdJPFM=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Segment ID
	//
	// example:
	//
	// 169
	SectionId *string `json:"SectionId,omitempty" xml:"SectionId,omitempty"`
}

func (s ListAppChatMessagesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAppChatMessagesRequest) GoString() string {
	return s.String()
}

func (s *ListAppChatMessagesRequest) GetChatId() *string {
	return s.ChatId
}

func (s *ListAppChatMessagesRequest) GetConversationId() *string {
	return s.ConversationId
}

func (s *ListAppChatMessagesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAppChatMessagesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAppChatMessagesRequest) GetSectionId() *string {
	return s.SectionId
}

func (s *ListAppChatMessagesRequest) SetChatId(v string) *ListAppChatMessagesRequest {
	s.ChatId = &v
	return s
}

func (s *ListAppChatMessagesRequest) SetConversationId(v string) *ListAppChatMessagesRequest {
	s.ConversationId = &v
	return s
}

func (s *ListAppChatMessagesRequest) SetMaxResults(v int32) *ListAppChatMessagesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListAppChatMessagesRequest) SetNextToken(v string) *ListAppChatMessagesRequest {
	s.NextToken = &v
	return s
}

func (s *ListAppChatMessagesRequest) SetSectionId(v string) *ListAppChatMessagesRequest {
	s.SectionId = &v
	return s
}

func (s *ListAppChatMessagesRequest) Validate() error {
	return dara.Validate(s)
}
